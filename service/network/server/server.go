package server

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/micro/cli/v2"
	net "github.com/micro/go-micro/v3/network"
	"github.com/micro/go-micro/v3/network/mucp"
	res "github.com/micro/go-micro/v3/network/resolver"
	"github.com/micro/go-micro/v3/network/resolver/dns"
	"github.com/micro/go-micro/v3/network/resolver/http"
	"github.com/micro/go-micro/v3/network/resolver/noop"
	"github.com/micro/go-micro/v3/network/resolver/registry"
	"github.com/micro/go-micro/v3/proxy"
	mucpProxy "github.com/micro/go-micro/v3/proxy/mucp"
	"github.com/micro/go-micro/v3/router"
	"github.com/micro/go-micro/v3/server"
	mucpServer "github.com/micro/go-micro/v3/server/mucp"
	"github.com/micro/go-micro/v3/transport"
	"github.com/micro/go-micro/v3/transport/quic"
	"github.com/micro/go-micro/v3/tunnel"
	tmucp "github.com/micro/go-micro/v3/tunnel/mucp"
	"github.com/micro/micro/v3/internal/helper"
	"github.com/micro/micro/v3/internal/muxer"
	"github.com/micro/micro/v3/service"
	log "github.com/micro/micro/v3/service/logger"
	muregistry "github.com/micro/micro/v3/service/registry"
	murouter "github.com/micro/micro/v3/service/router"
)

var (
	// name of the network service
	name = "network"
	// name of the micro network
	networkName = "micro"
	// address is the network address
	address = ":8443"
	// peerAddress is the address the network peers on
	peerAddress = ":8085"
	// set the advertise address
	advertise = ""
	// resolver is the network resolver
	resolver = "noop"
	// the tunnel token
	token = "micro"

	// Flags specific to the network
	Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "address",
			Usage:   "Set the address of the network service",
			EnvVars: []string{"MICRO_NETWORK_ADDRESS"},
		},
		&cli.StringFlag{
			Name:    "advertise",
			Usage:   "Set the micro network address to advertise",
			EnvVars: []string{"MICRO_NETWORK_ADVERTISE"},
		},
		&cli.StringFlag{
			Name:    "gateway",
			Usage:   "Set the default gateway",
			EnvVars: []string{"MICRO_NETWORK_GATEWAY"},
		},
		&cli.StringFlag{
			Name:    "network",
			Usage:   "Set the micro network name: micro",
			EnvVars: []string{"MICRO_NETWORK"},
		},
		&cli.StringFlag{
			Name:    "nodes",
			Usage:   "Set the micro network nodes to connect to. This can be a comma separated list.",
			EnvVars: []string{"MICRO_NETWORK_NODES"},
		},
		&cli.StringFlag{
			Name:    "token",
			Usage:   "Set the micro network token for authentication",
			EnvVars: []string{"MICRO_NETWORK_TOKEN"},
		},
		&cli.StringFlag{
			Name:    "resolver",
			Usage:   "Set the micro network resolver. This can be a comma separated list.",
			EnvVars: []string{"MICRO_NETWORK_RESOLVER"},
		},
	}
)

// Run runs the micro server
func Run(ctx *cli.Context) error {
	if len(ctx.String("server_name")) > 0 {
		name = ctx.String("server_name")
	}
	if len(ctx.String("address")) > 0 {
		address = ctx.String("address")
	}
	if len(ctx.String("peer_address")) > 0 {
		peerAddress = ctx.String("peer_address")
	}
	if len(ctx.String("advertise")) > 0 {
		advertise = ctx.String("advertise")
	}
	if len(ctx.String("network")) > 0 {
		networkName = ctx.String("network")
	}
	if len(ctx.String("token")) > 0 {
		token = ctx.String("token")
	}

	var nodes []string
	if len(ctx.String("nodes")) > 0 {
		nodes = strings.Split(ctx.String("nodes"), ",")
	}
	if len(ctx.String("resolver")) > 0 {
		resolver = ctx.String("resolver")
	}

	// setup the resolver used for internode networking
	var r res.Resolver

	switch resolver {
	case "dns":
		r = new(dns.Resolver)
	case "http":
		r = new(http.Resolver)
	case "registry":
		r = new(registry.Resolver)
	default:
		r = new(noop.Resolver)
	}

	// Initialise the local service
	service := service.New(
		service.Name(name),
		service.Address(address),
	)

	// create a tunnel
	tunOpts := []tunnel.Option{
		tunnel.Address(peerAddress),
		tunnel.Token(token),
	}

	if ctx.Bool("enable_tls") {
		config, err := helper.TLSConfig(ctx)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		config.InsecureSkipVerify = true

		tunOpts = append(tunOpts, tunnel.Transport(
			quic.NewTransport(transport.TLSConfig(config)),
		))
	}

	gateway := ctx.String("gateway")
	tun := tmucp.NewTunnel(tunOpts...)
	id := service.Server().Options().Id

	// local tunnel router
	rtr := murouter.DefaultRouter

	rtr.Init(
		router.Network(networkName),
		router.Id(id),
		router.Registry(muregistry.DefaultRegistry),
		router.Gateway(gateway),
		router.Precache(),
	)

	// create new network
	netService := mucp.NewNetwork(
		net.Id(id),
		net.Name(networkName),
		net.Address(peerAddress),
		net.Advertise(advertise),
		net.Nodes(nodes...),
		net.Tunnel(tun),
		net.Router(rtr),
		net.Resolver(r),
	)

	// network proxy
	// used by the network nodes to cluster
	// and share routes or route through
	// each other
	networkProxy := mucpProxy.NewProxy(
		proxy.WithRouter(rtr),
		proxy.WithClient(service.Client()),
		proxy.WithLink("network", netService.Client()),
	)

	// create a handler
	h := mucpServer.DefaultRouter.NewHandler(
		&Network{Network: netService},
	)

	// register the handler
	mucpServer.DefaultRouter.Handle(h)

	// network mux
	networkMux := muxer.New(name, networkProxy)

	// init the local server
	service.Server().Init(
		server.WithRouter(networkMux),
	)

	// set network server to proxy
	netService.Server().Init(
		server.WithRouter(networkMux),
	)

	// connect network
	if err := netService.Connect(); err != nil {
		log.Fatalf("Network failed to connect: %v", err)
	}

	// netClose hard exits if we have problems
	netClose := func(net net.Network) error {
		errChan := make(chan error, 1)

		go func() {
			errChan <- net.Close()
		}()

		select {
		case err := <-errChan:
			return err
		case <-time.After(time.Second):
			return errors.New("Network timeout closing")
		}
	}

	log.Infof("Network [%s] listening on %s", networkName, peerAddress)

	if err := service.Run(); err != nil {
		log.Errorf("Network %s failed: %v", networkName, err)
		netClose(netService)
		os.Exit(1)
	}

	// close the network
	netClose(netService)

	return nil
}
