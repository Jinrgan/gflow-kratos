// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: workflow.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// WorkflowClient is the client API for Workflow service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkflowClient interface {
	CreateWorkflow(ctx context.Context, in *CreateWorkflowRequest, opts ...grpc.CallOption) (*CreateWorkflowResponse, error)
	GetWorkflowDetail(ctx context.Context, in *GetWorkflowDetailRequest, opts ...grpc.CallOption) (*GetWorkflowDetailResponse, error)
	GetWorkflows(ctx context.Context, in *GetWorkflowsRequest, opts ...grpc.CallOption) (*GetWorkflowsResponse, error)
	UpdateWorkflow(ctx context.Context, in *UpdateWorkflowRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteWorkflow(ctx context.Context, in *DeleteWorkflowRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreateProcess(ctx context.Context, in *CreateProcessRequest, opts ...grpc.CallOption) (*CreateProcessResponse, error)
	UpdateProcess(ctx context.Context, in *UpdateProcessRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DrawProcesses(ctx context.Context, in *DrawProcessesRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteProcess(ctx context.Context, in *DeleteProcessRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteProcesses(ctx context.Context, in *DeleteProcessesRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type workflowClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkflowClient(cc grpc.ClientConnInterface) WorkflowClient {
	return &workflowClient{cc}
}

func (c *workflowClient) CreateWorkflow(ctx context.Context, in *CreateWorkflowRequest, opts ...grpc.CallOption) (*CreateWorkflowResponse, error) {
	out := new(CreateWorkflowResponse)
	err := c.cc.Invoke(ctx, "/greek.gflow.workflow.v1.Workflow/CreateWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowClient) GetWorkflowDetail(ctx context.Context, in *GetWorkflowDetailRequest, opts ...grpc.CallOption) (*GetWorkflowDetailResponse, error) {
	out := new(GetWorkflowDetailResponse)
	err := c.cc.Invoke(ctx, "/greek.gflow.workflow.v1.Workflow/GetWorkflowDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowClient) GetWorkflows(ctx context.Context, in *GetWorkflowsRequest, opts ...grpc.CallOption) (*GetWorkflowsResponse, error) {
	out := new(GetWorkflowsResponse)
	err := c.cc.Invoke(ctx, "/greek.gflow.workflow.v1.Workflow/GetWorkflows", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowClient) UpdateWorkflow(ctx context.Context, in *UpdateWorkflowRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/greek.gflow.workflow.v1.Workflow/UpdateWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowClient) DeleteWorkflow(ctx context.Context, in *DeleteWorkflowRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/greek.gflow.workflow.v1.Workflow/DeleteWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowClient) CreateProcess(ctx context.Context, in *CreateProcessRequest, opts ...grpc.CallOption) (*CreateProcessResponse, error) {
	out := new(CreateProcessResponse)
	err := c.cc.Invoke(ctx, "/greek.gflow.workflow.v1.Workflow/CreateProcess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowClient) UpdateProcess(ctx context.Context, in *UpdateProcessRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/greek.gflow.workflow.v1.Workflow/UpdateProcess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowClient) DrawProcesses(ctx context.Context, in *DrawProcessesRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/greek.gflow.workflow.v1.Workflow/DrawProcesses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowClient) DeleteProcess(ctx context.Context, in *DeleteProcessRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/greek.gflow.workflow.v1.Workflow/DeleteProcess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowClient) DeleteProcesses(ctx context.Context, in *DeleteProcessesRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/greek.gflow.workflow.v1.Workflow/DeleteProcesses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/greek.gflow.workflow.v1.Workflow/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkflowServer is the server API for Workflow service.
// All implementations must embed UnimplementedWorkflowServer
// for forward compatibility
type WorkflowServer interface {
	CreateWorkflow(context.Context, *CreateWorkflowRequest) (*CreateWorkflowResponse, error)
	GetWorkflowDetail(context.Context, *GetWorkflowDetailRequest) (*GetWorkflowDetailResponse, error)
	GetWorkflows(context.Context, *GetWorkflowsRequest) (*GetWorkflowsResponse, error)
	UpdateWorkflow(context.Context, *UpdateWorkflowRequest) (*emptypb.Empty, error)
	DeleteWorkflow(context.Context, *DeleteWorkflowRequest) (*emptypb.Empty, error)
	CreateProcess(context.Context, *CreateProcessRequest) (*CreateProcessResponse, error)
	UpdateProcess(context.Context, *UpdateProcessRequest) (*emptypb.Empty, error)
	DrawProcesses(context.Context, *DrawProcessesRequest) (*emptypb.Empty, error)
	DeleteProcess(context.Context, *DeleteProcessRequest) (*emptypb.Empty, error)
	DeleteProcesses(context.Context, *DeleteProcessesRequest) (*emptypb.Empty, error)
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedWorkflowServer()
}

// UnimplementedWorkflowServer must be embedded to have forward compatible implementations.
type UnimplementedWorkflowServer struct {
}

func (UnimplementedWorkflowServer) CreateWorkflow(context.Context, *CreateWorkflowRequest) (*CreateWorkflowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWorkflow not implemented")
}
func (UnimplementedWorkflowServer) GetWorkflowDetail(context.Context, *GetWorkflowDetailRequest) (*GetWorkflowDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkflowDetail not implemented")
}
func (UnimplementedWorkflowServer) GetWorkflows(context.Context, *GetWorkflowsRequest) (*GetWorkflowsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkflows not implemented")
}
func (UnimplementedWorkflowServer) UpdateWorkflow(context.Context, *UpdateWorkflowRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWorkflow not implemented")
}
func (UnimplementedWorkflowServer) DeleteWorkflow(context.Context, *DeleteWorkflowRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWorkflow not implemented")
}
func (UnimplementedWorkflowServer) CreateProcess(context.Context, *CreateProcessRequest) (*CreateProcessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProcess not implemented")
}
func (UnimplementedWorkflowServer) UpdateProcess(context.Context, *UpdateProcessRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProcess not implemented")
}
func (UnimplementedWorkflowServer) DrawProcesses(context.Context, *DrawProcessesRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DrawProcesses not implemented")
}
func (UnimplementedWorkflowServer) DeleteProcess(context.Context, *DeleteProcessRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProcess not implemented")
}
func (UnimplementedWorkflowServer) DeleteProcesses(context.Context, *DeleteProcessesRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProcesses not implemented")
}
func (UnimplementedWorkflowServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedWorkflowServer) mustEmbedUnimplementedWorkflowServer() {}

// UnsafeWorkflowServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkflowServer will
// result in compilation errors.
type UnsafeWorkflowServer interface {
	mustEmbedUnimplementedWorkflowServer()
}

func RegisterWorkflowServer(s grpc.ServiceRegistrar, srv WorkflowServer) {
	s.RegisterService(&Workflow_ServiceDesc, srv)
}

func _Workflow_CreateWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).CreateWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greek.gflow.workflow.v1.Workflow/CreateWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).CreateWorkflow(ctx, req.(*CreateWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflow_GetWorkflowDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkflowDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).GetWorkflowDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greek.gflow.workflow.v1.Workflow/GetWorkflowDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).GetWorkflowDetail(ctx, req.(*GetWorkflowDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflow_GetWorkflows_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkflowsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).GetWorkflows(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greek.gflow.workflow.v1.Workflow/GetWorkflows",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).GetWorkflows(ctx, req.(*GetWorkflowsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflow_UpdateWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).UpdateWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greek.gflow.workflow.v1.Workflow/UpdateWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).UpdateWorkflow(ctx, req.(*UpdateWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflow_DeleteWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).DeleteWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greek.gflow.workflow.v1.Workflow/DeleteWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).DeleteWorkflow(ctx, req.(*DeleteWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflow_CreateProcess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProcessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).CreateProcess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greek.gflow.workflow.v1.Workflow/CreateProcess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).CreateProcess(ctx, req.(*CreateProcessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflow_UpdateProcess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProcessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).UpdateProcess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greek.gflow.workflow.v1.Workflow/UpdateProcess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).UpdateProcess(ctx, req.(*UpdateProcessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflow_DrawProcesses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DrawProcessesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).DrawProcesses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greek.gflow.workflow.v1.Workflow/DrawProcesses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).DrawProcesses(ctx, req.(*DrawProcessesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflow_DeleteProcess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProcessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).DeleteProcess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greek.gflow.workflow.v1.Workflow/DeleteProcess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).DeleteProcess(ctx, req.(*DeleteProcessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflow_DeleteProcesses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProcessesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).DeleteProcesses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greek.gflow.workflow.v1.Workflow/DeleteProcesses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).DeleteProcesses(ctx, req.(*DeleteProcessesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflow_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greek.gflow.workflow.v1.Workflow/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Workflow_ServiceDesc is the grpc.ServiceDesc for Workflow service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Workflow_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "greek.gflow.workflow.v1.Workflow",
	HandlerType: (*WorkflowServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWorkflow",
			Handler:    _Workflow_CreateWorkflow_Handler,
		},
		{
			MethodName: "GetWorkflowDetail",
			Handler:    _Workflow_GetWorkflowDetail_Handler,
		},
		{
			MethodName: "GetWorkflows",
			Handler:    _Workflow_GetWorkflows_Handler,
		},
		{
			MethodName: "UpdateWorkflow",
			Handler:    _Workflow_UpdateWorkflow_Handler,
		},
		{
			MethodName: "DeleteWorkflow",
			Handler:    _Workflow_DeleteWorkflow_Handler,
		},
		{
			MethodName: "CreateProcess",
			Handler:    _Workflow_CreateProcess_Handler,
		},
		{
			MethodName: "UpdateProcess",
			Handler:    _Workflow_UpdateProcess_Handler,
		},
		{
			MethodName: "DrawProcesses",
			Handler:    _Workflow_DrawProcesses_Handler,
		},
		{
			MethodName: "DeleteProcess",
			Handler:    _Workflow_DeleteProcess_Handler,
		},
		{
			MethodName: "DeleteProcesses",
			Handler:    _Workflow_DeleteProcesses_Handler,
		},
		{
			MethodName: "SayHello",
			Handler:    _Workflow_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "workflow.proto",
}
