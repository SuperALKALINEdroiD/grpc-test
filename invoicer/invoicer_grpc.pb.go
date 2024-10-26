// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: invoicer.proto

package invoicer

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// InvoicerClient is the client API for Invoicer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InvoicerClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	ClientStreamTest(ctx context.Context, opts ...grpc.CallOption) (Invoicer_ClientStreamTestClient, error)
	ServerStreamTest(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (Invoicer_ServerStreamTestClient, error)
	BiDirectionalTest(ctx context.Context, opts ...grpc.CallOption) (Invoicer_BiDirectionalTestClient, error)
}

type invoicerClient struct {
	cc grpc.ClientConnInterface
}

func NewInvoicerClient(cc grpc.ClientConnInterface) InvoicerClient {
	return &invoicerClient{cc}
}

func (c *invoicerClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/Invoicer/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invoicerClient) ClientStreamTest(ctx context.Context, opts ...grpc.CallOption) (Invoicer_ClientStreamTestClient, error) {
	stream, err := c.cc.NewStream(ctx, &Invoicer_ServiceDesc.Streams[0], "/Invoicer/ClientStreamTest", opts...)
	if err != nil {
		return nil, err
	}
	x := &invoicerClientStreamTestClient{stream}
	return x, nil
}

type Invoicer_ClientStreamTestClient interface {
	Send(*CreateRequest) error
	CloseAndRecv() (*CreateResponse, error)
	grpc.ClientStream
}

type invoicerClientStreamTestClient struct {
	grpc.ClientStream
}

func (x *invoicerClientStreamTestClient) Send(m *CreateRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *invoicerClientStreamTestClient) CloseAndRecv() (*CreateResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CreateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *invoicerClient) ServerStreamTest(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (Invoicer_ServerStreamTestClient, error) {
	stream, err := c.cc.NewStream(ctx, &Invoicer_ServiceDesc.Streams[1], "/Invoicer/ServerStreamTest", opts...)
	if err != nil {
		return nil, err
	}
	x := &invoicerServerStreamTestClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Invoicer_ServerStreamTestClient interface {
	Recv() (*CreateResponse, error)
	grpc.ClientStream
}

type invoicerServerStreamTestClient struct {
	grpc.ClientStream
}

func (x *invoicerServerStreamTestClient) Recv() (*CreateResponse, error) {
	m := new(CreateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *invoicerClient) BiDirectionalTest(ctx context.Context, opts ...grpc.CallOption) (Invoicer_BiDirectionalTestClient, error) {
	stream, err := c.cc.NewStream(ctx, &Invoicer_ServiceDesc.Streams[2], "/Invoicer/BiDirectionalTest", opts...)
	if err != nil {
		return nil, err
	}
	x := &invoicerBiDirectionalTestClient{stream}
	return x, nil
}

type Invoicer_BiDirectionalTestClient interface {
	Send(*CreateRequest) error
	Recv() (*CreateResponse, error)
	grpc.ClientStream
}

type invoicerBiDirectionalTestClient struct {
	grpc.ClientStream
}

func (x *invoicerBiDirectionalTestClient) Send(m *CreateRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *invoicerBiDirectionalTestClient) Recv() (*CreateResponse, error) {
	m := new(CreateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// InvoicerServer is the server API for Invoicer service.
// All implementations must embed UnimplementedInvoicerServer
// for forward compatibility
type InvoicerServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	ClientStreamTest(Invoicer_ClientStreamTestServer) error
	ServerStreamTest(*CreateRequest, Invoicer_ServerStreamTestServer) error
	BiDirectionalTest(Invoicer_BiDirectionalTestServer) error
	mustEmbedUnimplementedInvoicerServer()
}

// UnimplementedInvoicerServer must be embedded to have forward compatible implementations.
type UnimplementedInvoicerServer struct {
}

func (UnimplementedInvoicerServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedInvoicerServer) ClientStreamTest(Invoicer_ClientStreamTestServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientStreamTest not implemented")
}
func (UnimplementedInvoicerServer) ServerStreamTest(*CreateRequest, Invoicer_ServerStreamTestServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerStreamTest not implemented")
}
func (UnimplementedInvoicerServer) BiDirectionalTest(Invoicer_BiDirectionalTestServer) error {
	return status.Errorf(codes.Unimplemented, "method BiDirectionalTest not implemented")
}
func (UnimplementedInvoicerServer) mustEmbedUnimplementedInvoicerServer() {}

// UnsafeInvoicerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InvoicerServer will
// result in compilation errors.
type UnsafeInvoicerServer interface {
	mustEmbedUnimplementedInvoicerServer()
}

func RegisterInvoicerServer(s grpc.ServiceRegistrar, srv InvoicerServer) {
	s.RegisterService(&Invoicer_ServiceDesc, srv)
}

func _Invoicer_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoicerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Invoicer/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoicerServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Invoicer_ClientStreamTest_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(InvoicerServer).ClientStreamTest(&invoicerClientStreamTestServer{stream})
}

type Invoicer_ClientStreamTestServer interface {
	SendAndClose(*CreateResponse) error
	Recv() (*CreateRequest, error)
	grpc.ServerStream
}

type invoicerClientStreamTestServer struct {
	grpc.ServerStream
}

func (x *invoicerClientStreamTestServer) SendAndClose(m *CreateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *invoicerClientStreamTestServer) Recv() (*CreateRequest, error) {
	m := new(CreateRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Invoicer_ServerStreamTest_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CreateRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(InvoicerServer).ServerStreamTest(m, &invoicerServerStreamTestServer{stream})
}

type Invoicer_ServerStreamTestServer interface {
	Send(*CreateResponse) error
	grpc.ServerStream
}

type invoicerServerStreamTestServer struct {
	grpc.ServerStream
}

func (x *invoicerServerStreamTestServer) Send(m *CreateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Invoicer_BiDirectionalTest_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(InvoicerServer).BiDirectionalTest(&invoicerBiDirectionalTestServer{stream})
}

type Invoicer_BiDirectionalTestServer interface {
	Send(*CreateResponse) error
	Recv() (*CreateRequest, error)
	grpc.ServerStream
}

type invoicerBiDirectionalTestServer struct {
	grpc.ServerStream
}

func (x *invoicerBiDirectionalTestServer) Send(m *CreateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *invoicerBiDirectionalTestServer) Recv() (*CreateRequest, error) {
	m := new(CreateRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Invoicer_ServiceDesc is the grpc.ServiceDesc for Invoicer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Invoicer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Invoicer",
	HandlerType: (*InvoicerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Invoicer_Create_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ClientStreamTest",
			Handler:       _Invoicer_ClientStreamTest_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ServerStreamTest",
			Handler:       _Invoicer_ServerStreamTest_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "BiDirectionalTest",
			Handler:       _Invoicer_BiDirectionalTest_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "invoicer.proto",
}
