// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: weather.proto

package pb

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

// WeatherServiceClient is the client API for WeatherService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WeatherServiceClient interface {
	GetWeatherDataByLatLongStream(ctx context.Context, opts ...grpc.CallOption) (WeatherService_GetWeatherDataByLatLongStreamClient, error)
	GetWeatherDataStream(ctx context.Context, in *RequestWeatherData, opts ...grpc.CallOption) (WeatherService_GetWeatherDataStreamClient, error)
	GetWeatherData(ctx context.Context, in *RequestWeatherData, opts ...grpc.CallOption) (*ResponseWeatherData, error)
	GetUserCity(ctx context.Context, in *RequstUserCity, opts ...grpc.CallOption) (*ResponseUserCity, error)
}

type weatherServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWeatherServiceClient(cc grpc.ClientConnInterface) WeatherServiceClient {
	return &weatherServiceClient{cc}
}

func (c *weatherServiceClient) GetWeatherDataByLatLongStream(ctx context.Context, opts ...grpc.CallOption) (WeatherService_GetWeatherDataByLatLongStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &WeatherService_ServiceDesc.Streams[0], "/weather.WeatherService/GetWeatherDataByLatLongStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &weatherServiceGetWeatherDataByLatLongStreamClient{stream}
	return x, nil
}

type WeatherService_GetWeatherDataByLatLongStreamClient interface {
	Send(*RequestStreamUserByLatLong) error
	Recv() (*ResponseStreamUserLatLong, error)
	grpc.ClientStream
}

type weatherServiceGetWeatherDataByLatLongStreamClient struct {
	grpc.ClientStream
}

func (x *weatherServiceGetWeatherDataByLatLongStreamClient) Send(m *RequestStreamUserByLatLong) error {
	return x.ClientStream.SendMsg(m)
}

func (x *weatherServiceGetWeatherDataByLatLongStreamClient) Recv() (*ResponseStreamUserLatLong, error) {
	m := new(ResponseStreamUserLatLong)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *weatherServiceClient) GetWeatherDataStream(ctx context.Context, in *RequestWeatherData, opts ...grpc.CallOption) (WeatherService_GetWeatherDataStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &WeatherService_ServiceDesc.Streams[1], "/weather.WeatherService/GetWeatherDataStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &weatherServiceGetWeatherDataStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WeatherService_GetWeatherDataStreamClient interface {
	Recv() (*ResponseStreamWeatherData, error)
	grpc.ClientStream
}

type weatherServiceGetWeatherDataStreamClient struct {
	grpc.ClientStream
}

func (x *weatherServiceGetWeatherDataStreamClient) Recv() (*ResponseStreamWeatherData, error) {
	m := new(ResponseStreamWeatherData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *weatherServiceClient) GetWeatherData(ctx context.Context, in *RequestWeatherData, opts ...grpc.CallOption) (*ResponseWeatherData, error) {
	out := new(ResponseWeatherData)
	err := c.cc.Invoke(ctx, "/weather.WeatherService/GetWeatherData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *weatherServiceClient) GetUserCity(ctx context.Context, in *RequstUserCity, opts ...grpc.CallOption) (*ResponseUserCity, error) {
	out := new(ResponseUserCity)
	err := c.cc.Invoke(ctx, "/weather.WeatherService/GetUserCity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WeatherServiceServer is the server API for WeatherService service.
// All implementations must embed UnimplementedWeatherServiceServer
// for forward compatibility
type WeatherServiceServer interface {
	GetWeatherDataByLatLongStream(WeatherService_GetWeatherDataByLatLongStreamServer) error
	GetWeatherDataStream(*RequestWeatherData, WeatherService_GetWeatherDataStreamServer) error
	GetWeatherData(context.Context, *RequestWeatherData) (*ResponseWeatherData, error)
	GetUserCity(context.Context, *RequstUserCity) (*ResponseUserCity, error)
	mustEmbedUnimplementedWeatherServiceServer()
}

// UnimplementedWeatherServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWeatherServiceServer struct {
}

func (UnimplementedWeatherServiceServer) GetWeatherDataByLatLongStream(WeatherService_GetWeatherDataByLatLongStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetWeatherDataByLatLongStream not implemented")
}
func (UnimplementedWeatherServiceServer) GetWeatherDataStream(*RequestWeatherData, WeatherService_GetWeatherDataStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetWeatherDataStream not implemented")
}
func (UnimplementedWeatherServiceServer) GetWeatherData(context.Context, *RequestWeatherData) (*ResponseWeatherData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWeatherData not implemented")
}
func (UnimplementedWeatherServiceServer) GetUserCity(context.Context, *RequstUserCity) (*ResponseUserCity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserCity not implemented")
}
func (UnimplementedWeatherServiceServer) mustEmbedUnimplementedWeatherServiceServer() {}

// UnsafeWeatherServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WeatherServiceServer will
// result in compilation errors.
type UnsafeWeatherServiceServer interface {
	mustEmbedUnimplementedWeatherServiceServer()
}

func RegisterWeatherServiceServer(s grpc.ServiceRegistrar, srv WeatherServiceServer) {
	s.RegisterService(&WeatherService_ServiceDesc, srv)
}

func _WeatherService_GetWeatherDataByLatLongStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WeatherServiceServer).GetWeatherDataByLatLongStream(&weatherServiceGetWeatherDataByLatLongStreamServer{stream})
}

type WeatherService_GetWeatherDataByLatLongStreamServer interface {
	Send(*ResponseStreamUserLatLong) error
	Recv() (*RequestStreamUserByLatLong, error)
	grpc.ServerStream
}

type weatherServiceGetWeatherDataByLatLongStreamServer struct {
	grpc.ServerStream
}

func (x *weatherServiceGetWeatherDataByLatLongStreamServer) Send(m *ResponseStreamUserLatLong) error {
	return x.ServerStream.SendMsg(m)
}

func (x *weatherServiceGetWeatherDataByLatLongStreamServer) Recv() (*RequestStreamUserByLatLong, error) {
	m := new(RequestStreamUserByLatLong)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _WeatherService_GetWeatherDataStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestWeatherData)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WeatherServiceServer).GetWeatherDataStream(m, &weatherServiceGetWeatherDataStreamServer{stream})
}

type WeatherService_GetWeatherDataStreamServer interface {
	Send(*ResponseStreamWeatherData) error
	grpc.ServerStream
}

type weatherServiceGetWeatherDataStreamServer struct {
	grpc.ServerStream
}

func (x *weatherServiceGetWeatherDataStreamServer) Send(m *ResponseStreamWeatherData) error {
	return x.ServerStream.SendMsg(m)
}

func _WeatherService_GetWeatherData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestWeatherData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeatherServiceServer).GetWeatherData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/weather.WeatherService/GetWeatherData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeatherServiceServer).GetWeatherData(ctx, req.(*RequestWeatherData))
	}
	return interceptor(ctx, in, info, handler)
}

func _WeatherService_GetUserCity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequstUserCity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeatherServiceServer).GetUserCity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/weather.WeatherService/GetUserCity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeatherServiceServer).GetUserCity(ctx, req.(*RequstUserCity))
	}
	return interceptor(ctx, in, info, handler)
}

// WeatherService_ServiceDesc is the grpc.ServiceDesc for WeatherService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WeatherService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "weather.WeatherService",
	HandlerType: (*WeatherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWeatherData",
			Handler:    _WeatherService_GetWeatherData_Handler,
		},
		{
			MethodName: "GetUserCity",
			Handler:    _WeatherService_GetUserCity_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetWeatherDataByLatLongStream",
			Handler:       _WeatherService_GetWeatherDataByLatLongStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "GetWeatherDataStream",
			Handler:       _WeatherService_GetWeatherDataStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "weather.proto",
}
