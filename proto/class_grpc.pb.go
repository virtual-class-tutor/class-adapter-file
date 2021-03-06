// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package class

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// AdapterClient is the client API for Adapter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdapterClient interface {
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*Classes, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Class, error)
	Create(ctx context.Context, in *Class, opts ...grpc.CallOption) (*Class, error)
	Update(ctx context.Context, in *Class, opts ...grpc.CallOption) (*Class, error)
	Delete(ctx context.Context, in *Class, opts ...grpc.CallOption) (*Empty, error)
}

type adapterClient struct {
	cc grpc.ClientConnInterface
}

func NewAdapterClient(cc grpc.ClientConnInterface) AdapterClient {
	return &adapterClient{cc}
}

func (c *adapterClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*Classes, error) {
	out := new(Classes)
	err := c.cc.Invoke(ctx, "/class.Adapter/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adapterClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Class, error) {
	out := new(Class)
	err := c.cc.Invoke(ctx, "/class.Adapter/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adapterClient) Create(ctx context.Context, in *Class, opts ...grpc.CallOption) (*Class, error) {
	out := new(Class)
	err := c.cc.Invoke(ctx, "/class.Adapter/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adapterClient) Update(ctx context.Context, in *Class, opts ...grpc.CallOption) (*Class, error) {
	out := new(Class)
	err := c.cc.Invoke(ctx, "/class.Adapter/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adapterClient) Delete(ctx context.Context, in *Class, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/class.Adapter/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdapterServer is the server API for Adapter service.
// All implementations must embed UnimplementedAdapterServer
// for forward compatibility
type AdapterServer interface {
	List(context.Context, *ListRequest) (*Classes, error)
	Get(context.Context, *GetRequest) (*Class, error)
	Create(context.Context, *Class) (*Class, error)
	Update(context.Context, *Class) (*Class, error)
	Delete(context.Context, *Class) (*Empty, error)
	mustEmbedUnimplementedAdapterServer()
}

// UnimplementedAdapterServer must be embedded to have forward compatible implementations.
type UnimplementedAdapterServer struct {
}

func (UnimplementedAdapterServer) List(context.Context, *ListRequest) (*Classes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedAdapterServer) Get(context.Context, *GetRequest) (*Class, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedAdapterServer) Create(context.Context, *Class) (*Class, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAdapterServer) Update(context.Context, *Class) (*Class, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAdapterServer) Delete(context.Context, *Class) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAdapterServer) mustEmbedUnimplementedAdapterServer() {}

// UnsafeAdapterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdapterServer will
// result in compilation errors.
type UnsafeAdapterServer interface {
	mustEmbedUnimplementedAdapterServer()
}

func RegisterAdapterServer(s *grpc.Server, srv AdapterServer) {
	s.RegisterService(&_Adapter_serviceDesc, srv)
}

func _Adapter_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdapterServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/class.Adapter/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdapterServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Adapter_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdapterServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/class.Adapter/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdapterServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Adapter_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Class)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdapterServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/class.Adapter/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdapterServer).Create(ctx, req.(*Class))
	}
	return interceptor(ctx, in, info, handler)
}

func _Adapter_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Class)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdapterServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/class.Adapter/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdapterServer).Update(ctx, req.(*Class))
	}
	return interceptor(ctx, in, info, handler)
}

func _Adapter_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Class)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdapterServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/class.Adapter/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdapterServer).Delete(ctx, req.(*Class))
	}
	return interceptor(ctx, in, info, handler)
}

var _Adapter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "class.Adapter",
	HandlerType: (*AdapterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Adapter_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Adapter_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Adapter_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Adapter_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Adapter_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/class.proto",
}
