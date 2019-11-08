// Code generated by protoc-gen-go. DO NOT EDIT.
// source: micro/micro/internal/token/proto/token.proto

package go_micro_token

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

// Token is for auth
type Token struct {
	// unique id of token
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// unix timestamp
	Expires uint64 `protobuf:"varint,2,opt,name=expires,proto3" json:"expires,omitempty"`
	// various values
	Claims map[string]string `protobuf:"bytes,3,rep,name=claims,proto3" json:"claims,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// encrypted key
	Key                  string   `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_785746f161be1dd6, []int{0}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Token) GetExpires() uint64 {
	if m != nil {
		return m.Expires
	}
	return 0
}

func (m *Token) GetClaims() map[string]string {
	if m != nil {
		return m.Claims
	}
	return nil
}

func (m *Token) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func init() {
	proto.RegisterType((*Token)(nil), "go.micro.token.Token")
	proto.RegisterMapType((map[string]string)(nil), "go.micro.token.Token.ClaimsEntry")
}

func init() {
	proto.RegisterFile("micro/micro/internal/token/proto/token.proto", fileDescriptor_785746f161be1dd6)
}

var fileDescriptor_785746f161be1dd6 = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xc9, 0xcd, 0x4c, 0x2e,
	0xca, 0xd7, 0x87, 0x90, 0x99, 0x79, 0x25, 0xa9, 0x45, 0x79, 0x89, 0x39, 0xfa, 0x25, 0xf9, 0xd9,
	0xa9, 0x79, 0xfa, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x10, 0xb6, 0x1e, 0x98, 0x2d, 0xc4, 0x97, 0x9e,
	0xaf, 0x07, 0x56, 0xaa, 0x07, 0x16, 0x55, 0xda, 0xc9, 0xc8, 0xc5, 0x1a, 0x02, 0x62, 0x09, 0xf1,
	0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x31, 0x65, 0xa6, 0x08, 0x49,
	0x70, 0xb1, 0xa7, 0x56, 0x14, 0x64, 0x16, 0xa5, 0x16, 0x4b, 0x30, 0x29, 0x30, 0x6a, 0xb0, 0x04,
	0xc1, 0xb8, 0x42, 0x96, 0x5c, 0x6c, 0xc9, 0x39, 0x89, 0x99, 0xb9, 0xc5, 0x12, 0xcc, 0x0a, 0xcc,
	0x1a, 0xdc, 0x46, 0x8a, 0x7a, 0xa8, 0x86, 0xea, 0x81, 0x0d, 0xd4, 0x73, 0x06, 0xab, 0x71, 0xcd,
	0x2b, 0x29, 0xaa, 0x0c, 0x82, 0x6a, 0x10, 0x12, 0xe0, 0x62, 0xce, 0x4e, 0xad, 0x94, 0x60, 0x01,
	0xdb, 0x02, 0x62, 0x4a, 0x59, 0x72, 0x71, 0x23, 0x29, 0x84, 0x29, 0x60, 0x84, 0x2b, 0x10, 0x12,
	0xe1, 0x62, 0x2d, 0x4b, 0xcc, 0x29, 0x4d, 0x05, 0xbb, 0x82, 0x33, 0x08, 0xc2, 0xb1, 0x62, 0xb2,
	0x60, 0x4c, 0x62, 0x03, 0x7b, 0xc9, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x90, 0x13, 0xd5, 0xa0,
	0x02, 0x01, 0x00, 0x00,
}