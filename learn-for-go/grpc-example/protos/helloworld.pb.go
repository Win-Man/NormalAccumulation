// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helloworld.proto

package protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// The request message containing the user's name.
type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{1}
}

func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (m *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(m, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Something struct {
	LineCode             int64    `protobuf:"varint,1,opt,name=lineCode,proto3" json:"lineCode,omitempty"`
	Line                 string   `protobuf:"bytes,2,opt,name=line,proto3" json:"line,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Something) Reset()         { *m = Something{} }
func (m *Something) String() string { return proto.CompactTextString(m) }
func (*Something) ProtoMessage()    {}
func (*Something) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{2}
}

func (m *Something) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Something.Unmarshal(m, b)
}
func (m *Something) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Something.Marshal(b, m, deterministic)
}
func (m *Something) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Something.Merge(m, src)
}
func (m *Something) XXX_Size() int {
	return xxx_messageInfo_Something.Size(m)
}
func (m *Something) XXX_DiscardUnknown() {
	xxx_messageInfo_Something.DiscardUnknown(m)
}

var xxx_messageInfo_Something proto.InternalMessageInfo

func (m *Something) GetLineCode() int64 {
	if m != nil {
		return m.LineCode
	}
	return 0
}

func (m *Something) GetLine() string {
	if m != nil {
		return m.Line
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "protos.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "protos.HelloReply")
	proto.RegisterType((*Something)(nil), "protos.Something")
}

func init() {
	proto.RegisterFile("helloworld.proto", fileDescriptor_17b8c58d586b62f2)
}

var fileDescriptor_17b8c58d586b62f2 = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0x17, 0x15, 0xb7, 0x3d, 0x04, 0xe7, 0xc3, 0x43, 0xd9, 0x49, 0x72, 0x90, 0x9e, 0xc6,
	0x50, 0x10, 0x41, 0xf4, 0xe2, 0x41, 0x2f, 0xbb, 0x74, 0x03, 0xf1, 0x18, 0xd9, 0x63, 0x2d, 0xbe,
	0x36, 0x35, 0xc9, 0x90, 0xfe, 0xeb, 0x9e, 0x24, 0x8d, 0x6d, 0x45, 0x73, 0xca, 0xf7, 0x5e, 0xbe,
	0x2f, 0xbf, 0x0f, 0x02, 0xb3, 0x9c, 0x98, 0xf5, 0xa7, 0x36, 0xbc, 0x5d, 0xd4, 0x46, 0x3b, 0x8d,
	0xc7, 0xed, 0x61, 0xa5, 0x84, 0x93, 0x67, 0x7f, 0x97, 0xd1, 0xc7, 0x9e, 0xac, 0x43, 0x84, 0xa3,
	0x4a, 0x95, 0x94, 0x88, 0x0b, 0x91, 0x4e, 0xb3, 0x56, 0xcb, 0x4b, 0x80, 0x1f, 0x4f, 0xcd, 0x0d,
	0x26, 0x30, 0x2e, 0xc9, 0x5a, 0xb5, 0xeb, 0x4c, 0xdd, 0x28, 0xef, 0x60, 0xba, 0xd6, 0x25, 0xb9,
	0xbc, 0xa8, 0x76, 0x38, 0x87, 0x09, 0x17, 0x15, 0x3d, 0xea, 0x6d, 0xf0, 0x1d, 0x66, 0xfd, 0xec,
	0x21, 0x5e, 0x27, 0x07, 0x01, 0xe2, 0xf5, 0xd5, 0x97, 0x80, 0xf1, 0x93, 0x21, 0x72, 0x64, 0xf0,
	0x06, 0x26, 0x6b, 0xd5, 0xb4, 0x4c, 0x3c, 0x0f, 0x85, 0xed, 0xe2, 0x77, 0xcd, 0x39, 0xfe, 0xd9,
	0xd6, 0xdc, 0xc8, 0x11, 0x3e, 0xc0, 0xe9, 0x86, 0x98, 0x57, 0x34, 0xd4, 0x88, 0xc7, 0xcf, 0xba,
	0x6d, 0x6f, 0x94, 0xa3, 0xa5, 0xc0, 0x7b, 0x98, 0xf9, 0xfc, 0xab, 0xde, 0x0f, 0x0f, 0xfc, 0xb7,
	0xc6, 0xe1, 0xa9, 0xc0, 0x5b, 0x80, 0x8d, 0xe2, 0xf7, 0x97, 0xc2, 0xe5, 0x2b, 0x8a, 0x05, 0x63,
	0xd8, 0x54, 0x2c, 0xc5, 0x5b, 0xf8, 0x8d, 0xeb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xef, 0xbe,
	0x26, 0xec, 0xa8, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	TellMeSomething(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (Greeter_TellMeSomethingClient, error)
	TellYouSomething(ctx context.Context, opts ...grpc.CallOption) (Greeter_TellYouSomethingClient, error)
	TalkWithMe(ctx context.Context, opts ...grpc.CallOption) (Greeter_TalkWithMeClient, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/protos.Greeter/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) TellMeSomething(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (Greeter_TellMeSomethingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Greeter_serviceDesc.Streams[0], "/protos.Greeter/TellMeSomething", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterTellMeSomethingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Greeter_TellMeSomethingClient interface {
	Recv() (*Something, error)
	grpc.ClientStream
}

type greeterTellMeSomethingClient struct {
	grpc.ClientStream
}

func (x *greeterTellMeSomethingClient) Recv() (*Something, error) {
	m := new(Something)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterClient) TellYouSomething(ctx context.Context, opts ...grpc.CallOption) (Greeter_TellYouSomethingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Greeter_serviceDesc.Streams[1], "/protos.Greeter/TellYouSomething", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterTellYouSomethingClient{stream}
	return x, nil
}

type Greeter_TellYouSomethingClient interface {
	Send(*Something) error
	CloseAndRecv() (*HelloReply, error)
	grpc.ClientStream
}

type greeterTellYouSomethingClient struct {
	grpc.ClientStream
}

func (x *greeterTellYouSomethingClient) Send(m *Something) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterTellYouSomethingClient) CloseAndRecv() (*HelloReply, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(HelloReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterClient) TalkWithMe(ctx context.Context, opts ...grpc.CallOption) (Greeter_TalkWithMeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Greeter_serviceDesc.Streams[2], "/protos.Greeter/TalkWithMe", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterTalkWithMeClient{stream}
	return x, nil
}

type Greeter_TalkWithMeClient interface {
	Send(*Something) error
	Recv() (*Something, error)
	grpc.ClientStream
}

type greeterTalkWithMeClient struct {
	grpc.ClientStream
}

func (x *greeterTalkWithMeClient) Send(m *Something) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterTalkWithMeClient) Recv() (*Something, error) {
	m := new(Something)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	TellMeSomething(*HelloRequest, Greeter_TellMeSomethingServer) error
	TellYouSomething(Greeter_TellYouSomethingServer) error
	TalkWithMe(Greeter_TalkWithMeServer) error
}

// UnimplementedGreeterServer can be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (*UnimplementedGreeterServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (*UnimplementedGreeterServer) TellMeSomething(req *HelloRequest, srv Greeter_TellMeSomethingServer) error {
	return status.Errorf(codes.Unimplemented, "method TellMeSomething not implemented")
}
func (*UnimplementedGreeterServer) TellYouSomething(srv Greeter_TellYouSomethingServer) error {
	return status.Errorf(codes.Unimplemented, "method TellYouSomething not implemented")
}
func (*UnimplementedGreeterServer) TalkWithMe(srv Greeter_TalkWithMeServer) error {
	return status.Errorf(codes.Unimplemented, "method TalkWithMe not implemented")
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_TellMeSomething_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreeterServer).TellMeSomething(m, &greeterTellMeSomethingServer{stream})
}

type Greeter_TellMeSomethingServer interface {
	Send(*Something) error
	grpc.ServerStream
}

type greeterTellMeSomethingServer struct {
	grpc.ServerStream
}

func (x *greeterTellMeSomethingServer) Send(m *Something) error {
	return x.ServerStream.SendMsg(m)
}

func _Greeter_TellYouSomething_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServer).TellYouSomething(&greeterTellYouSomethingServer{stream})
}

type Greeter_TellYouSomethingServer interface {
	SendAndClose(*HelloReply) error
	Recv() (*Something, error)
	grpc.ServerStream
}

type greeterTellYouSomethingServer struct {
	grpc.ServerStream
}

func (x *greeterTellYouSomethingServer) SendAndClose(m *HelloReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterTellYouSomethingServer) Recv() (*Something, error) {
	m := new(Something)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Greeter_TalkWithMe_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServer).TalkWithMe(&greeterTalkWithMeServer{stream})
}

type Greeter_TalkWithMeServer interface {
	Send(*Something) error
	Recv() (*Something, error)
	grpc.ServerStream
}

type greeterTalkWithMeServer struct {
	grpc.ServerStream
}

func (x *greeterTalkWithMeServer) Send(m *Something) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterTalkWithMeServer) Recv() (*Something, error) {
	m := new(Something)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TellMeSomething",
			Handler:       _Greeter_TellMeSomething_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "TellYouSomething",
			Handler:       _Greeter_TellYouSomething_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "TalkWithMe",
			Handler:       _Greeter_TalkWithMe_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "helloworld.proto",
}