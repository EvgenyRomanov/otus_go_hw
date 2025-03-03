// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: EventService.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CalendarService_CreateEvent_FullMethodName       = "/event.CalendarService/CreateEvent"
	CalendarService_UpdateEvent_FullMethodName       = "/event.CalendarService/UpdateEvent"
	CalendarService_DeleteEvent_FullMethodName       = "/event.CalendarService/DeleteEvent"
	CalendarService_GetEvents_FullMethodName         = "/event.CalendarService/GetEvents"
	CalendarService_GetEvent_FullMethodName          = "/event.CalendarService/GetEvent"
	CalendarService_GetEventsForDay_FullMethodName   = "/event.CalendarService/GetEventsForDay"
	CalendarService_GetEventsForWeek_FullMethodName  = "/event.CalendarService/GetEventsForWeek"
	CalendarService_GetEventsForMonth_FullMethodName = "/event.CalendarService/GetEventsForMonth"
)

// CalendarServiceClient is the client API for CalendarService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalendarServiceClient interface {
	CreateEvent(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateEvent(ctx context.Context, in *EventUpdateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteEvent(ctx context.Context, in *EventIdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetEvents(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*EventsResponse, error)
	GetEvent(ctx context.Context, in *EventIdRequest, opts ...grpc.CallOption) (*EventResponse, error)
	GetEventsForDay(ctx context.Context, in *RangeRequest, opts ...grpc.CallOption) (*EventsResponse, error)
	GetEventsForWeek(ctx context.Context, in *RangeRequest, opts ...grpc.CallOption) (*EventsResponse, error)
	GetEventsForMonth(ctx context.Context, in *RangeRequest, opts ...grpc.CallOption) (*EventsResponse, error)
}

type calendarServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalendarServiceClient(cc grpc.ClientConnInterface) CalendarServiceClient {
	return &calendarServiceClient{cc}
}

func (c *calendarServiceClient) CreateEvent(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CalendarService_CreateEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarServiceClient) UpdateEvent(ctx context.Context, in *EventUpdateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CalendarService_UpdateEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarServiceClient) DeleteEvent(ctx context.Context, in *EventIdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CalendarService_DeleteEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarServiceClient) GetEvents(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*EventsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EventsResponse)
	err := c.cc.Invoke(ctx, CalendarService_GetEvents_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarServiceClient) GetEvent(ctx context.Context, in *EventIdRequest, opts ...grpc.CallOption) (*EventResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EventResponse)
	err := c.cc.Invoke(ctx, CalendarService_GetEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarServiceClient) GetEventsForDay(ctx context.Context, in *RangeRequest, opts ...grpc.CallOption) (*EventsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EventsResponse)
	err := c.cc.Invoke(ctx, CalendarService_GetEventsForDay_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarServiceClient) GetEventsForWeek(ctx context.Context, in *RangeRequest, opts ...grpc.CallOption) (*EventsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EventsResponse)
	err := c.cc.Invoke(ctx, CalendarService_GetEventsForWeek_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarServiceClient) GetEventsForMonth(ctx context.Context, in *RangeRequest, opts ...grpc.CallOption) (*EventsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EventsResponse)
	err := c.cc.Invoke(ctx, CalendarService_GetEventsForMonth_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalendarServiceServer is the server API for CalendarService service.
// All implementations must embed UnimplementedCalendarServiceServer
// for forward compatibility.
type CalendarServiceServer interface {
	CreateEvent(context.Context, *EventRequest) (*emptypb.Empty, error)
	UpdateEvent(context.Context, *EventUpdateRequest) (*emptypb.Empty, error)
	DeleteEvent(context.Context, *EventIdRequest) (*emptypb.Empty, error)
	GetEvents(context.Context, *emptypb.Empty) (*EventsResponse, error)
	GetEvent(context.Context, *EventIdRequest) (*EventResponse, error)
	GetEventsForDay(context.Context, *RangeRequest) (*EventsResponse, error)
	GetEventsForWeek(context.Context, *RangeRequest) (*EventsResponse, error)
	GetEventsForMonth(context.Context, *RangeRequest) (*EventsResponse, error)
	mustEmbedUnimplementedCalendarServiceServer()
}

// UnimplementedCalendarServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCalendarServiceServer struct{}

func (UnimplementedCalendarServiceServer) CreateEvent(context.Context, *EventRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEvent not implemented")
}
func (UnimplementedCalendarServiceServer) UpdateEvent(context.Context, *EventUpdateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEvent not implemented")
}
func (UnimplementedCalendarServiceServer) DeleteEvent(context.Context, *EventIdRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEvent not implemented")
}
func (UnimplementedCalendarServiceServer) GetEvents(context.Context, *emptypb.Empty) (*EventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvents not implemented")
}
func (UnimplementedCalendarServiceServer) GetEvent(context.Context, *EventIdRequest) (*EventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvent not implemented")
}
func (UnimplementedCalendarServiceServer) GetEventsForDay(context.Context, *RangeRequest) (*EventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventsForDay not implemented")
}
func (UnimplementedCalendarServiceServer) GetEventsForWeek(context.Context, *RangeRequest) (*EventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventsForWeek not implemented")
}
func (UnimplementedCalendarServiceServer) GetEventsForMonth(context.Context, *RangeRequest) (*EventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventsForMonth not implemented")
}
func (UnimplementedCalendarServiceServer) mustEmbedUnimplementedCalendarServiceServer() {}
func (UnimplementedCalendarServiceServer) testEmbeddedByValue()                         {}

// UnsafeCalendarServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalendarServiceServer will
// result in compilation errors.
type UnsafeCalendarServiceServer interface {
	mustEmbedUnimplementedCalendarServiceServer()
}

func RegisterCalendarServiceServer(s grpc.ServiceRegistrar, srv CalendarServiceServer) {
	// If the following call pancis, it indicates UnimplementedCalendarServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CalendarService_ServiceDesc, srv)
}

func _CalendarService_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServiceServer).CreateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CalendarService_CreateEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServiceServer).CreateEvent(ctx, req.(*EventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarService_UpdateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServiceServer).UpdateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CalendarService_UpdateEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServiceServer).UpdateEvent(ctx, req.(*EventUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarService_DeleteEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServiceServer).DeleteEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CalendarService_DeleteEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServiceServer).DeleteEvent(ctx, req.(*EventIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarService_GetEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServiceServer).GetEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CalendarService_GetEvents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServiceServer).GetEvents(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarService_GetEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServiceServer).GetEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CalendarService_GetEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServiceServer).GetEvent(ctx, req.(*EventIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarService_GetEventsForDay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServiceServer).GetEventsForDay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CalendarService_GetEventsForDay_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServiceServer).GetEventsForDay(ctx, req.(*RangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarService_GetEventsForWeek_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServiceServer).GetEventsForWeek(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CalendarService_GetEventsForWeek_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServiceServer).GetEventsForWeek(ctx, req.(*RangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarService_GetEventsForMonth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServiceServer).GetEventsForMonth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CalendarService_GetEventsForMonth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServiceServer).GetEventsForMonth(ctx, req.(*RangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CalendarService_ServiceDesc is the grpc.ServiceDesc for CalendarService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalendarService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "event.CalendarService",
	HandlerType: (*CalendarServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEvent",
			Handler:    _CalendarService_CreateEvent_Handler,
		},
		{
			MethodName: "UpdateEvent",
			Handler:    _CalendarService_UpdateEvent_Handler,
		},
		{
			MethodName: "DeleteEvent",
			Handler:    _CalendarService_DeleteEvent_Handler,
		},
		{
			MethodName: "GetEvents",
			Handler:    _CalendarService_GetEvents_Handler,
		},
		{
			MethodName: "GetEvent",
			Handler:    _CalendarService_GetEvent_Handler,
		},
		{
			MethodName: "GetEventsForDay",
			Handler:    _CalendarService_GetEventsForDay_Handler,
		},
		{
			MethodName: "GetEventsForWeek",
			Handler:    _CalendarService_GetEventsForWeek_Handler,
		},
		{
			MethodName: "GetEventsForMonth",
			Handler:    _CalendarService_GetEventsForMonth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "EventService.proto",
}
