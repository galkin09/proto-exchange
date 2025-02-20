// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.6
// source: exchange.proto

package proto_exchange

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ExchangeService_GetExchangeRates_FullMethodName           = "/exchange.ExchangeService/GetExchangeRates"
	ExchangeService_GetExchangeRateForCurrency_FullMethodName = "/exchange.ExchangeService/GetExchangeRateForCurrency"
)

// ExchangeServiceClient is the client API for ExchangeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Определение сервиса
type ExchangeServiceClient interface {
	// Получение курсов обмена всех валют
	GetExchangeRates(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ExchangeRatesResponse, error)
	// Получение курса обмена для конкретной валюты
	GetExchangeRateForCurrency(ctx context.Context, in *CurrencyRequest, opts ...grpc.CallOption) (*ExchangeRateResponse, error)
}

type exchangeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExchangeServiceClient(cc grpc.ClientConnInterface) ExchangeServiceClient {
	return &exchangeServiceClient{cc}
}

func (c *exchangeServiceClient) GetExchangeRates(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ExchangeRatesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExchangeRatesResponse)
	err := c.cc.Invoke(ctx, ExchangeService_GetExchangeRates_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeServiceClient) GetExchangeRateForCurrency(ctx context.Context, in *CurrencyRequest, opts ...grpc.CallOption) (*ExchangeRateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExchangeRateResponse)
	err := c.cc.Invoke(ctx, ExchangeService_GetExchangeRateForCurrency_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExchangeServiceServer is the server API for ExchangeService service.
// All implementations must embed UnimplementedExchangeServiceServer
// for forward compatibility.
//
// Определение сервиса
type ExchangeServiceServer interface {
	// Получение курсов обмена всех валют
	GetExchangeRates(context.Context, *Empty) (*ExchangeRatesResponse, error)
	// Получение курса обмена для конкретной валюты
	GetExchangeRateForCurrency(context.Context, *CurrencyRequest) (*ExchangeRateResponse, error)
	mustEmbedUnimplementedExchangeServiceServer()
}

// UnimplementedExchangeServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedExchangeServiceServer struct{}

func (UnimplementedExchangeServiceServer) GetExchangeRates(context.Context, *Empty) (*ExchangeRatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExchangeRates not implemented")
}
func (UnimplementedExchangeServiceServer) GetExchangeRateForCurrency(context.Context, *CurrencyRequest) (*ExchangeRateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExchangeRateForCurrency not implemented")
}
func (UnimplementedExchangeServiceServer) mustEmbedUnimplementedExchangeServiceServer() {}
func (UnimplementedExchangeServiceServer) testEmbeddedByValue()                         {}

// UnsafeExchangeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExchangeServiceServer will
// result in compilation errors.
type UnsafeExchangeServiceServer interface {
	mustEmbedUnimplementedExchangeServiceServer()
}

func RegisterExchangeServiceServer(s grpc.ServiceRegistrar, srv ExchangeServiceServer) {
	// If the following call pancis, it indicates UnimplementedExchangeServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ExchangeService_ServiceDesc, srv)
}

func _ExchangeService_GetExchangeRates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServiceServer).GetExchangeRates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExchangeService_GetExchangeRates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServiceServer).GetExchangeRates(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExchangeService_GetExchangeRateForCurrency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CurrencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServiceServer).GetExchangeRateForCurrency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExchangeService_GetExchangeRateForCurrency_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServiceServer).GetExchangeRateForCurrency(ctx, req.(*CurrencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExchangeService_ServiceDesc is the grpc.ServiceDesc for ExchangeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExchangeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "exchange.ExchangeService",
	HandlerType: (*ExchangeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetExchangeRates",
			Handler:    _ExchangeService_GetExchangeRates_Handler,
		},
		{
			MethodName: "GetExchangeRateForCurrency",
			Handler:    _ExchangeService_GetExchangeRateForCurrency_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "exchange.proto",
}
