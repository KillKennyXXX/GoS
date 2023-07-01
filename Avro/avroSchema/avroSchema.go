// Code generated by gen-grpc-protocol. DO NOT EDIT.

package avroSchema

import (
	context "context"
	goavro "github.com/linkedin/goavro/v2"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// TestClient is the client API for the test service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestClient interface {
	// Send a one way message
	Ack(ctx context.Context, in map[string]interface{}, opts ...grpc.CallOption) (interface{}, error)
	// Client and server stream counts bi-directionally
	BidirCounter(ctx context.Context, opts ...grpc.CallOption) (Test_BidirCounterClient, error)
	// Client streams counts to the server
	ClientCounter(ctx context.Context, opts ...grpc.CallOption) (Test_ClientCounterClient, error)
	// Pretend you're in a cave!
	Echo(ctx context.Context, in map[string]interface{}, opts ...grpc.CallOption) (interface{}, error)
	// Send a friendly greeting
	Hello(ctx context.Context, in map[string]interface{}, opts ...grpc.CallOption) (interface{}, error)
	// Server streams counts to the client
	ServerCounter(ctx context.Context, in map[string]interface{}, opts ...grpc.CallOption) (Test_ServerCounterClient, error)
	// Intentionally not implemented
	Unimplemented(ctx context.Context, in map[string]interface{}, opts ...grpc.CallOption) (interface{}, error)
}

type testClient struct {
	cc grpc.ClientConnInterface
}

func NewTestClient(cc grpc.ClientConnInterface) TestClient {
	return &testClient{cc}
}

func (c *testClient) Ack(ctx context.Context, in map[string]interface{}, opts ...grpc.CallOption) (interface{}, error) {
	buf, err := _Test_Ack_requestCodec.BinaryFromNative(nil, in)
	if err != nil {
		return nil, err
	}
	if err = c.cc.Invoke(ctx, "/test/ack", buf, &buf, opts...); err != nil {
		return nil, err
	}
	out, _, err := _Test_Ack_responseCodec.NativeFromBinary(buf)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testClient) BidirCounter(ctx context.Context, opts ...grpc.CallOption) (Test_BidirCounterClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Test_serviceDesc.Streams[0], "/test/bidirCounter", opts...)
	if err != nil {
		return nil, err
	}
	x := &testBidirCounterClient{stream}
	return x, nil
}

type Test_BidirCounterClient interface {
	Send(interface{}) error
	Recv() (interface{}, error)
	grpc.ClientStream
}

type testBidirCounterClient struct {
	grpc.ClientStream
}

func (x *testBidirCounterClient) Send(m interface{}) error {
	buf, err := _Test_BidirCounter_requestCodec.BinaryFromNative(nil, m)
	if err != nil {
		return err
	}
	return x.ClientStream.SendMsg(buf)
}

func (x *testBidirCounterClient) Recv() (interface{}, error) {
	var buf []byte
	if err := x.ClientStream.RecvMsg(&buf); err != nil {
		return nil, err
	}
	m, _, err := _Test_BidirCounter_responseCodec.NativeFromBinary(buf)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testClient) ClientCounter(ctx context.Context, opts ...grpc.CallOption) (Test_ClientCounterClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Test_serviceDesc.Streams[0], "/test/clientCounter", opts...)
	if err != nil {
		return nil, err
	}
	x := &testClientCounterClient{stream}
	return x, nil
}

type Test_ClientCounterClient interface {
	Send(interface{}) error
	CloseAndRecv() (interface{}, error)
	grpc.ClientStream
}

type testClientCounterClient struct {
	grpc.ClientStream
}

func (x *testClientCounterClient) Send(m interface{}) error {
	buf, err := _Test_ClientCounter_requestCodec.BinaryFromNative(nil, m)
	if err != nil {
		return err
	}
	return x.ClientStream.SendMsg(buf)
}

func (x *testClientCounterClient) CloseAndRecv() (interface{}, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	var buf []byte
	if err := x.ClientStream.RecvMsg(&buf); err != nil {
		return nil, err
	}
	m, _, err := _Test_ClientCounter_responseCodec.NativeFromBinary(buf)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testClient) Echo(ctx context.Context, in map[string]interface{}, opts ...grpc.CallOption) (interface{}, error) {
	buf, err := _Test_Echo_requestCodec.BinaryFromNative(nil, in)
	if err != nil {
		return nil, err
	}
	if err = c.cc.Invoke(ctx, "/test/echo", buf, &buf, opts...); err != nil {
		return nil, err
	}
	out, _, err := _Test_Echo_responseCodec.NativeFromBinary(buf)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testClient) Hello(ctx context.Context, in map[string]interface{}, opts ...grpc.CallOption) (interface{}, error) {
	buf, err := _Test_Hello_requestCodec.BinaryFromNative(nil, in)
	if err != nil {
		return nil, err
	}
	if err = c.cc.Invoke(ctx, "/test/hello", buf, &buf, opts...); err != nil {
		return nil, err
	}
	out, _, err := _Test_Hello_responseCodec.NativeFromBinary(buf)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testClient) ServerCounter(ctx context.Context, in map[string]interface{}, opts ...grpc.CallOption) (Test_ServerCounterClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Test_serviceDesc.Streams[0], "/test/serverCounter", opts...)
	if err != nil {
		return nil, err
	}
	buf, err := _Test_ServerCounter_requestCodec.BinaryFromNative(nil, in)
	if err != nil {
		return nil, err
	}
	x := &testServerCounterClient{stream}
	if err := x.ClientStream.SendMsg(buf); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Test_ServerCounterClient interface {
	Recv() (interface{}, error)
	grpc.ClientStream
}

type testServerCounterClient struct {
	grpc.ClientStream
}

func (x *testServerCounterClient) Recv() (interface{}, error) {
	var buf []byte
	if err := x.ClientStream.RecvMsg(&buf); err != nil {
		return nil, err
	}
	m, _, err := _Test_ServerCounter_responseCodec.NativeFromBinary(buf)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testClient) Unimplemented(ctx context.Context, in map[string]interface{}, opts ...grpc.CallOption) (interface{}, error) {
	buf, err := _Test_Unimplemented_requestCodec.BinaryFromNative(nil, in)
	if err != nil {
		return nil, err
	}
	if err = c.cc.Invoke(ctx, "/test/unimplemented", buf, &buf, opts...); err != nil {
		return nil, err
	}
	out, _, err := _Test_Unimplemented_responseCodec.NativeFromBinary(buf)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestServer is the server API for the test service.
// All implementations must embed UnimplementedTestServer
// for forward compatibility
type TestServer interface {
	// Send a one way message
	Ack(ctx context.Context, in map[string]interface{}) (interface{}, error)
	// Client and server stream counts bi-directionally
	BidirCounter(srv Test_BidirCounterServer) error
	// Client streams counts to the server
	ClientCounter(srv Test_ClientCounterServer) error
	// Pretend you're in a cave!
	Echo(ctx context.Context, in map[string]interface{}) (interface{}, error)
	// Send a friendly greeting
	Hello(ctx context.Context, in map[string]interface{}) (interface{}, error)
	// Server streams counts to the client
	ServerCounter(in map[string]interface{}, srv Test_ServerCounterServer) error
	// Intentionally not implemented
	Unimplemented(ctx context.Context, in map[string]interface{}) (interface{}, error)
	mustEmbedUnimplementedTestServer()
}

// UnimplementedTestServer must be embedded to have forward compatible implementations.
type UnimplementedTestServer struct {
}

func (UnimplementedTestServer) Ack(context.Context, map[string]interface{}) (interface{}, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ack not implemented")
}

func (UnimplementedTestServer) BidirCounter(Test_BidirCounterServer) error {
	return status.Errorf(codes.Unimplemented, "method BidirCounter not implemented")
}

func (UnimplementedTestServer) ClientCounter(Test_ClientCounterServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientCounter not implemented")
}

func (UnimplementedTestServer) Echo(context.Context, map[string]interface{}) (interface{}, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}

func (UnimplementedTestServer) Hello(context.Context, map[string]interface{}) (interface{}, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}

func (UnimplementedTestServer) ServerCounter(map[string]interface{}, Test_ServerCounterServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerCounter not implemented")
}

func (UnimplementedTestServer) Unimplemented(context.Context, map[string]interface{}) (interface{}, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unimplemented not implemented")
}

func (UnimplementedTestServer) mustEmbedUnimplementedTestServer() {}

// UnsafeTestServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestServer will
// result in compilation errors.
type UnsafeTestServer interface {
	mustEmbedUnimplementedTestServer()
}

func RegisterTestServer(s grpc.ServiceRegistrar, srv TestServer) {
	s.RegisterService(&_Test_serviceDesc, srv)
}

func _Test_Ack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	var buf []byte
	if err := dec(&buf); err != nil {
		return nil, err
	}
	blob, _, err := _Test_Ack_requestCodec.NativeFromBinary(buf)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error unmarshaling Ack request: "+err.Error())
	}
	in := blob.(map[string]interface{})
	if interceptor == nil {
		out, err := srv.(TestServer).Ack(ctx, in)
		if err != nil {
			return nil, err
		}
		if out, err = _Test_Ack_responseCodec.BinaryFromNative(nil, out); err != nil {
			return nil, status.Errorf(codes.Internal, "error marshaling Ack response: "+err.Error())
		}
		return out, nil
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test/ack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		out, err := srv.(TestServer).Ack(ctx, req.(map[string]interface{}))
		if err != nil {
			return nil, err
		}
		if out, err = _Test_Ack_responseCodec.BinaryFromNative(nil, out); err != nil {
			return nil, status.Errorf(codes.Internal, "error marshaling Ack response: "+err.Error())
		}
		return out, nil
	}
	return interceptor(ctx, in, info, handler)
}

func _Test_BidirCounter_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServer).BidirCounter(&testBidirCounterServer{stream})
}

type Test_BidirCounterServer interface {
	Send(interface{}) error
	Recv() (interface{}, error)
	grpc.ServerStream
}

type testBidirCounterServer struct {
	grpc.ServerStream
}

func (x *testBidirCounterServer) Send(m interface{}) error {
	buf, err := _Test_BidirCounter_responseCodec.BinaryFromNative(nil, m)
	if err != nil {
		return err
	}
	return x.ServerStream.SendMsg(buf)
}

func (x *testBidirCounterServer) Recv() (interface{}, error) {
	var buf []byte
	if err := x.ServerStream.RecvMsg(&buf); err != nil {
		return nil, err
	}
	m, _, err := _Test_BidirCounter_requestCodec.NativeFromBinary(buf)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func _Test_ClientCounter_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServer).ClientCounter(&testClientCounterServer{stream})
}

type Test_ClientCounterServer interface {
	SendAndClose(interface{}) error
	Recv() (interface{}, error)
	grpc.ServerStream
}

type testClientCounterServer struct {
	grpc.ServerStream
}

func (x *testClientCounterServer) SendAndClose(m interface{}) error {
	buf, err := _Test_ClientCounter_responseCodec.BinaryFromNative(nil, m)
	if err != nil {
		return err
	}
	return x.ServerStream.SendMsg(buf)
}

func (x *testClientCounterServer) Recv() (interface{}, error) {
	var buf []byte
	if err := x.ServerStream.RecvMsg(&buf); err != nil {
		return nil, err
	}
	m, _, err := _Test_ClientCounter_requestCodec.NativeFromBinary(buf)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func _Test_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	var buf []byte
	if err := dec(&buf); err != nil {
		return nil, err
	}
	blob, _, err := _Test_Echo_requestCodec.NativeFromBinary(buf)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error unmarshaling Echo request: "+err.Error())
	}
	in := blob.(map[string]interface{})
	if interceptor == nil {
		out, err := srv.(TestServer).Echo(ctx, in)
		if err != nil {
			return nil, err
		}
		if out, err = _Test_Echo_responseCodec.BinaryFromNative(nil, out); err != nil {
			return nil, status.Errorf(codes.Internal, "error marshaling Echo response: "+err.Error())
		}
		return out, nil
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test/echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		out, err := srv.(TestServer).Echo(ctx, req.(map[string]interface{}))
		if err != nil {
			return nil, err
		}
		if out, err = _Test_Echo_responseCodec.BinaryFromNative(nil, out); err != nil {
			return nil, status.Errorf(codes.Internal, "error marshaling Echo response: "+err.Error())
		}
		return out, nil
	}
	return interceptor(ctx, in, info, handler)
}

func _Test_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	var buf []byte
	if err := dec(&buf); err != nil {
		return nil, err
	}
	blob, _, err := _Test_Hello_requestCodec.NativeFromBinary(buf)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error unmarshaling Hello request: "+err.Error())
	}
	in := blob.(map[string]interface{})
	if interceptor == nil {
		out, err := srv.(TestServer).Hello(ctx, in)
		if err != nil {
			return nil, err
		}
		if out, err = _Test_Hello_responseCodec.BinaryFromNative(nil, out); err != nil {
			return nil, status.Errorf(codes.Internal, "error marshaling Hello response: "+err.Error())
		}
		return out, nil
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test/hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		out, err := srv.(TestServer).Hello(ctx, req.(map[string]interface{}))
		if err != nil {
			return nil, err
		}
		if out, err = _Test_Hello_responseCodec.BinaryFromNative(nil, out); err != nil {
			return nil, status.Errorf(codes.Internal, "error marshaling Hello response: "+err.Error())
		}
		return out, nil
	}
	return interceptor(ctx, in, info, handler)
}

func _Test_ServerCounter_Handler(srv interface{}, stream grpc.ServerStream) error {
	var buf []byte
	if err := stream.RecvMsg(&buf); err != nil {
		return err
	}
	blob, _, err := _Test_ServerCounter_requestCodec.NativeFromBinary(buf)
	if err != nil {
		return err
	}
	in := blob.(map[string]interface{})
	return srv.(TestServer).ServerCounter(in, &testServerCounterServer{stream})
}

type Test_ServerCounterServer interface {
	Send(interface{}) error
	grpc.ServerStream
}

type testServerCounterServer struct {
	grpc.ServerStream
}

func (x *testServerCounterServer) Send(m interface{}) error {
	buf, err := _Test_ServerCounter_responseCodec.BinaryFromNative(nil, m)
	if err != nil {
		return err
	}
	return x.ServerStream.SendMsg(buf)
}

func _Test_Unimplemented_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	var buf []byte
	if err := dec(&buf); err != nil {
		return nil, err
	}
	blob, _, err := _Test_Unimplemented_requestCodec.NativeFromBinary(buf)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error unmarshaling Unimplemented request: "+err.Error())
	}
	in := blob.(map[string]interface{})
	if interceptor == nil {
		out, err := srv.(TestServer).Unimplemented(ctx, in)
		if err != nil {
			return nil, err
		}
		if out, err = _Test_Unimplemented_responseCodec.BinaryFromNative(nil, out); err != nil {
			return nil, status.Errorf(codes.Internal, "error marshaling Unimplemented response: "+err.Error())
		}
		return out, nil
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test/unimplemented",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		out, err := srv.(TestServer).Unimplemented(ctx, req.(map[string]interface{}))
		if err != nil {
			return nil, err
		}
		if out, err = _Test_Unimplemented_responseCodec.BinaryFromNative(nil, out); err != nil {
			return nil, status.Errorf(codes.Internal, "error marshaling Unimplemented response: "+err.Error())
		}
		return out, nil
	}
	return interceptor(ctx, in, info, handler)
}

var (
	_Test_serviceDesc = grpc.ServiceDesc{
		ServiceName: "test",
		HandlerType: (*TestServer)(nil),
		Methods: []grpc.MethodDesc{
			{
				MethodName: "ack",
				Handler:    _Test_Ack_Handler,
			},
			{
				MethodName: "echo",
				Handler:    _Test_Echo_Handler,
			},
			{
				MethodName: "hello",
				Handler:    _Test_Hello_Handler,
			},
			{
				MethodName: "unimplemented",
				Handler:    _Test_Unimplemented_Handler,
			},
		},
		Streams: []grpc.StreamDesc{
			{
				StreamName:    "bidirCounter",
				Handler:       _Test_BidirCounter_Handler,
				ServerStreams: true,
				ClientStreams: true,
			},
			{
				StreamName:    "clientCounter",
				Handler:       _Test_ClientCounter_Handler,
				ClientStreams: true,
			},
			{
				StreamName:    "serverCounter",
				Handler:       _Test_ServerCounter_Handler,
				ServerStreams: true,
			},
		},
		Metadata: "test.avpr",
	}

	_Test_Ack_requestCodec            *goavro.Codec
	_Test_Ack_responseCodec           *goavro.Codec
	_Test_BidirCounter_requestCodec   *goavro.Codec
	_Test_BidirCounter_responseCodec  *goavro.Codec
	_Test_ClientCounter_requestCodec  *goavro.Codec
	_Test_ClientCounter_responseCodec *goavro.Codec
	_Test_Echo_requestCodec           *goavro.Codec
	_Test_Echo_responseCodec          *goavro.Codec
	_Test_Hello_requestCodec          *goavro.Codec
	_Test_Hello_responseCodec         *goavro.Codec
	_Test_ServerCounter_requestCodec  *goavro.Codec
	_Test_ServerCounter_responseCodec *goavro.Codec
	_Test_Unimplemented_requestCodec  *goavro.Codec
	_Test_Unimplemented_responseCodec *goavro.Codec
)

func init() {
	var err error
	if _Test_Ack_requestCodec, err = goavro.NewCodec(`{"doc":"Send a one way message","fields":[],"name":"ack","type":"record"}`); err != nil {
		panic(err)
	}
	if _Test_Ack_responseCodec, err = goavro.NewCodec(`"null"`); err != nil {
		panic(err)
	}
	if _Test_BidirCounter_requestCodec, err = goavro.NewCodec(`{"doc":"Client and server stream counts bi-directionally","fields":[{"name":"counter","type":"int"}],"name":"bidirCounter","type":"record"}`); err != nil {
		panic(err)
	}
	if _Test_BidirCounter_responseCodec, err = goavro.NewCodec(`"int"`); err != nil {
		panic(err)
	}
	if _Test_ClientCounter_requestCodec, err = goavro.NewCodec(`{"doc":"Client streams counts to the server","fields":[{"name":"counter","type":"int"}],"name":"clientCounter","type":"record"}`); err != nil {
		panic(err)
	}
	if _Test_ClientCounter_responseCodec, err = goavro.NewCodec(`{"items":"int","type":"array"}`); err != nil {
		panic(err)
	}
	if _Test_Echo_requestCodec, err = goavro.NewCodec(`{"doc":"Pretend you're in a cave!","fields":[{"fields":[{"name":"name","order":"ignore","type":"string"},{"name":"kind","order":"descending","symbols":["FOO","BAR","BAZ"],"type":"enum"},{"name":"hash","size":16,"type":"fixed"}],"name":"record","type":"record"}],"name":"echo","type":"record"}`); err != nil {
		panic(err)
	}
	if _Test_Echo_responseCodec, err = goavro.NewCodec(`{"fields":[{"name":"name","order":"ignore","type":"string"},{"name":"kind","order":"descending","symbols":["FOO","BAR","BAZ"],"type":"enum"},{"name":"hash","size":16,"type":"fixed"}],"name":"TestRecord","type":"record"}`); err != nil {
		panic(err)
	}
	if _Test_Hello_requestCodec, err = goavro.NewCodec(`{"doc":"Send a friendly greeting","fields":[{"default":"","name":"name","type":"string"}],"name":"hello","type":"record"}`); err != nil {
		panic(err)
	}
	if _Test_Hello_responseCodec, err = goavro.NewCodec(`"string"`); err != nil {
		panic(err)
	}
	if _Test_ServerCounter_requestCodec, err = goavro.NewCodec(`{"doc":"Server streams counts to the client","fields":[],"name":"serverCounter","type":"record"}`); err != nil {
		panic(err)
	}
	if _Test_ServerCounter_responseCodec, err = goavro.NewCodec(`"int"`); err != nil {
		panic(err)
	}
	if _Test_Unimplemented_requestCodec, err = goavro.NewCodec(`{"doc":"Intentionally not implemented","fields":[],"name":"unimplemented","type":"record"}`); err != nil {
		panic(err)
	}
	if _Test_Unimplemented_responseCodec, err = goavro.NewCodec(`"null"`); err != nil {
		panic(err)
	}
}
