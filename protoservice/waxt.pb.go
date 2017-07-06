package protoservice

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The request message containing the user's name.
type Customer struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Email string `protobuf:"bytes,2,opt,email=email" json:"email,omitempty"`
	Phone string `protobuf:"bytes,3,opt,phone=phone" json:"phone,omitempty"`
}

func (c *Customer) Reset()                    { *c = Customer{} }
func (c *Customer) String() string            { return proto.CompactTextString(c) }
func (*Customer) ProtoMessage()               {}
func (*Customer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*Customer)(nil), "protoservice.Customer")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Greeter service

type WaxtClient interface {
	Save(ctx context.Context, in *Customer, opts ...grpc.CallOption) (*Customer, error)
}

type waxtClient struct {
	cc *grpc.ClientConn
}

func NewWaxtClient(cc *grpc.ClientConn) WaxtClient {
	return &waxtClient{cc}
}

func (c *waxtClient) Save(ctx context.Context, in *Customer, opts ...grpc.CallOption) (*Customer, error) {
	out := new(Customer)
	err := grpc.Invoke(ctx, "/protoservice.Waxt/Save", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type WaxtServer interface {
	Save(context.Context, *Customer) (*Customer, error)
}

func RegisterWaxtServer(s *grpc.Server, srv WaxtServer) {
	s.RegisterService(&_waxt_serviceDesc, srv)
}

var _waxt_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protoservice.Waxt",
	HandlerType: (*WaxtServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Save",
			Handler:    _Waxt_Save_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protoservice.proto",
}

func _Waxt_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Customer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WaxtServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protoservice.Waxt/Save",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WaxtServer).Save(ctx, req.(*Customer))
	}
	return interceptor(ctx, in, info, handler)
}

func init() { proto.RegisterFile("protoservice.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0x48, 0xcd, 0xc9,
	0xc9, 0x2f, 0xcf, 0x2f, 0xca, 0x49, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x42, 0x88,
	0x28, 0x29, 0x71, 0xf1, 0x78, 0x80, 0x78, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x42,
	0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x92,
	0x1a, 0x17, 0x17, 0x54, 0x4d, 0x41, 0x4e, 0xa5, 0x90, 0x04, 0x17, 0x7b, 0x6e, 0x6a, 0x71, 0x71,
	0x62, 0x3a, 0x4c, 0x11, 0x8c, 0x6b, 0xe4, 0xc9, 0xc5, 0xee, 0x5e, 0x94, 0x9a, 0x5a, 0x92, 0x5a,
	0x24, 0x64, 0xc7, 0xc5, 0x11, 0x9c, 0x58, 0x09, 0xd6, 0x25, 0x24, 0xa1, 0x87, 0xe4, 0x02, 0x64,
	0xcb, 0xa4, 0xc4, 0xb0, 0xc8, 0x00, 0xad, 0x50, 0x62, 0x70, 0x32, 0xe0, 0x92, 0xce, 0xcc, 0xd7,
	0x4b, 0x2f, 0x2a, 0x48, 0xd6, 0x4b, 0xad, 0x48, 0xcc, 0x2d, 0xc8, 0x49, 0x2d, 0x46, 0x52, 0xeb,
	0xc4, 0x0f, 0x56, 0x1c, 0x0e, 0x62, 0x07, 0x80, 0xbc, 0x14, 0xc0, 0x98, 0xc4, 0x06, 0xf6, 0x9b,
	0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x0f, 0xb7, 0xcd, 0xf2, 0xef, 0x00, 0x00, 0x00,
}
