// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package log_v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// LogClient is the client API for Log service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogClient interface {
	Produce(ctx context.Context, in *ProduceRequest, opts ...grpc.CallOption) (*ProduceResponse, error)
	Consume(ctx context.Context, in *ConsumeRequest, opts ...grpc.CallOption) (*ConsumeResponse, error)
	ConsumeStream(ctx context.Context, in *ConsumeRequest, opts ...grpc.CallOption) (Log_ConsumeStreamClient, error)
	ProdcueStream(ctx context.Context, opts ...grpc.CallOption) (Log_ProdcueStreamClient, error)
}

type logClient struct {
	cc grpc.ClientConnInterface
}

func NewLogClient(cc grpc.ClientConnInterface) LogClient {
	return &logClient{cc}
}

func (c *logClient) Produce(ctx context.Context, in *ProduceRequest, opts ...grpc.CallOption) (*ProduceResponse, error) {
	out := new(ProduceResponse)
	err := c.cc.Invoke(ctx, "/log.v1.Log/Produce", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logClient) Consume(ctx context.Context, in *ConsumeRequest, opts ...grpc.CallOption) (*ConsumeResponse, error) {
	out := new(ConsumeResponse)
	err := c.cc.Invoke(ctx, "/log.v1.Log/Consume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logClient) ConsumeStream(ctx context.Context, in *ConsumeRequest, opts ...grpc.CallOption) (Log_ConsumeStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Log_serviceDesc.Streams[0], "/log.v1.Log/ConsumeStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &logConsumeStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Log_ConsumeStreamClient interface {
	Recv() (*ConsumeResponse, error)
	grpc.ClientStream
}

type logConsumeStreamClient struct {
	grpc.ClientStream
}

func (x *logConsumeStreamClient) Recv() (*ConsumeResponse, error) {
	m := new(ConsumeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *logClient) ProdcueStream(ctx context.Context, opts ...grpc.CallOption) (Log_ProdcueStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Log_serviceDesc.Streams[1], "/log.v1.Log/ProdcueStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &logProdcueStreamClient{stream}
	return x, nil
}

type Log_ProdcueStreamClient interface {
	Send(*ProduceRequest) error
	Recv() (*ProduceResponse, error)
	grpc.ClientStream
}

type logProdcueStreamClient struct {
	grpc.ClientStream
}

func (x *logProdcueStreamClient) Send(m *ProduceRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *logProdcueStreamClient) Recv() (*ProduceResponse, error) {
	m := new(ProduceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LogServer is the server API for Log service.
// All implementations must embed UnimplementedLogServer
// for forward compatibility
type LogServer interface {
	Produce(context.Context, *ProduceRequest) (*ProduceResponse, error)
	Consume(context.Context, *ConsumeRequest) (*ConsumeResponse, error)
	ConsumeStream(*ConsumeRequest, Log_ConsumeStreamServer) error
	ProdcueStream(Log_ProdcueStreamServer) error
	mustEmbedUnimplementedLogServer()
}

// UnimplementedLogServer must be embedded to have forward compatible implementations.
type UnimplementedLogServer struct {
}

func (UnimplementedLogServer) Produce(context.Context, *ProduceRequest) (*ProduceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Produce not implemented")
}
func (UnimplementedLogServer) Consume(context.Context, *ConsumeRequest) (*ConsumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Consume not implemented")
}
func (UnimplementedLogServer) ConsumeStream(*ConsumeRequest, Log_ConsumeStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ConsumeStream not implemented")
}
func (UnimplementedLogServer) ProdcueStream(Log_ProdcueStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ProdcueStream not implemented")
}
func (UnimplementedLogServer) mustEmbedUnimplementedLogServer() {}

// UnsafeLogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogServer will
// result in compilation errors.
type UnsafeLogServer interface {
	mustEmbedUnimplementedLogServer()
}

func RegisterLogServer(s grpc.ServiceRegistrar, srv LogServer) {
	s.RegisterService(&_Log_serviceDesc, srv)
}

func _Log_Produce_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProduceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServer).Produce(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/log.v1.Log/Produce",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServer).Produce(ctx, req.(*ProduceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Log_Consume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConsumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServer).Consume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/log.v1.Log/Consume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServer).Consume(ctx, req.(*ConsumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Log_ConsumeStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConsumeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LogServer).ConsumeStream(m, &logConsumeStreamServer{stream})
}

type Log_ConsumeStreamServer interface {
	Send(*ConsumeResponse) error
	grpc.ServerStream
}

type logConsumeStreamServer struct {
	grpc.ServerStream
}

func (x *logConsumeStreamServer) Send(m *ConsumeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Log_ProdcueStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LogServer).ProdcueStream(&logProdcueStreamServer{stream})
}

type Log_ProdcueStreamServer interface {
	Send(*ProduceResponse) error
	Recv() (*ProduceRequest, error)
	grpc.ServerStream
}

type logProdcueStreamServer struct {
	grpc.ServerStream
}

func (x *logProdcueStreamServer) Send(m *ProduceResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *logProdcueStreamServer) Recv() (*ProduceRequest, error) {
	m := new(ProduceRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Log_serviceDesc = grpc.ServiceDesc{
	ServiceName: "log.v1.Log",
	HandlerType: (*LogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Produce",
			Handler:    _Log_Produce_Handler,
		},
		{
			MethodName: "Consume",
			Handler:    _Log_Consume_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ConsumeStream",
			Handler:       _Log_ConsumeStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ProdcueStream",
			Handler:       _Log_ProdcueStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/v1/log.proto",
}
