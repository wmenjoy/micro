// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: client/bot/proto/bot.proto

package go_micro_bot

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v3/api"
	client "github.com/micro/go-micro/v3/client"
	server "github.com/micro/go-micro/v3/server"
	microService "github.com/micro/micro/v3/service"
	microClient "github.com/micro/micro/v3/service/client"
	microServer "github.com/micro/micro/v3/service/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option
var _ = microServer.Handle
var _ = microClient.Call

// Api Endpoints for Command service

func NewCommandEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Command service

type CommandService interface {
	Help(ctx context.Context, in *HelpRequest, opts ...client.CallOption) (*HelpResponse, error)
	Exec(ctx context.Context, in *ExecRequest, opts ...client.CallOption) (*ExecResponse, error)
}

type commandService struct {
	c    client.Client
	name string
}

func NewCommandService(name string, c client.Client) CommandService {
	return &commandService{
		c:    c,
		name: name,
	}
}

func CommandServiceClient() CommandService {
	return NewCommandService("command", microClient.DefaultClient)
}

func RunCommandService() {
	microService.Init(microService.Name("command"))
	microService.Run()
}

func (c *commandService) Help(ctx context.Context, in *HelpRequest, opts ...client.CallOption) (*HelpResponse, error) {
	req := microClient.NewRequest(c.name, "Command.Help", in)
	out := new(HelpResponse)
	err := microClient.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandService) Exec(ctx context.Context, in *ExecRequest, opts ...client.CallOption) (*ExecResponse, error) {
	req := microClient.NewRequest(c.name, "Command.Exec", in)
	out := new(ExecResponse)
	err := microClient.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Command service

type CommandHandler interface {
	Help(context.Context, *HelpRequest, *HelpResponse) error
	Exec(context.Context, *ExecRequest, *ExecResponse) error
}

func RegisterCommandService(hdlr CommandHandler, opts ...server.HandlerOption) error {
	return RegisterCommandHandler(microServer.DefaultServer, hdlr, opts...)
}

func RegisterCommandHandler(s server.Server, hdlr CommandHandler, opts ...server.HandlerOption) error {
	type command interface {
		Help(ctx context.Context, in *HelpRequest, out *HelpResponse) error
		Exec(ctx context.Context, in *ExecRequest, out *ExecResponse) error
	}
	type Command struct {
		command
	}
	h := &commandHandler{hdlr}
	return s.Handle(s.NewHandler(&Command{h}, opts...))
}

type commandHandler struct {
	CommandHandler
}

func (h *commandHandler) Help(ctx context.Context, in *HelpRequest, out *HelpResponse) error {
	return h.CommandHandler.Help(ctx, in, out)
}

func (h *commandHandler) Exec(ctx context.Context, in *ExecRequest, out *ExecResponse) error {
	return h.CommandHandler.Exec(ctx, in, out)
}