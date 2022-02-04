// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.3

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type WorkflowHTTPServer interface {
	CreateProcess(context.Context, *CreateProcessRequest) (*CreateProcessResponse, error)
	CreateWorkflow(context.Context, *CreateWorkflowRequest) (*CreateWorkflowResponse, error)
	DeleteProcess(context.Context, *DeleteProcessRequest) (*emptypb.Empty, error)
	DeleteProcesses(context.Context, *DeleteProcessesRequest) (*emptypb.Empty, error)
	DeleteWorkflow(context.Context, *DeleteWorkflowRequest) (*emptypb.Empty, error)
	DrawProcesses(context.Context, *DrawProcessesRequest) (*emptypb.Empty, error)
	GetWorkflowDetail(context.Context, *GetWorkflowDetailRequest) (*GetWorkflowDetailResponse, error)
	GetWorkflows(context.Context, *GetWorkflowsRequest) (*GetWorkflowsResponse, error)
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	UpdateProcess(context.Context, *UpdateProcessRequest) (*emptypb.Empty, error)
	UpdateWorkflow(context.Context, *UpdateWorkflowRequest) (*emptypb.Empty, error)
}

func RegisterWorkflowHTTPServer(s *http.Server, srv WorkflowHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/workflow", _Workflow_CreateWorkflow0_HTTP_Handler(srv))
	r.GET("/v1/workflow/detail/{id}", _Workflow_GetWorkflowDetail0_HTTP_Handler(srv))
	r.GET("/v1/workflows", _Workflow_GetWorkflows0_HTTP_Handler(srv))
	r.PUT("/v1/workflow/{id}", _Workflow_UpdateWorkflow0_HTTP_Handler(srv))
	r.DELETE("/v1/workflow/{id}", _Workflow_DeleteWorkflow0_HTTP_Handler(srv))
	r.POST("/v1/workflow/process", _Workflow_CreateProcess0_HTTP_Handler(srv))
	r.PUT("/v1/workflow/process/{id}", _Workflow_UpdateProcess0_HTTP_Handler(srv))
	r.PUT("/v1/workflow/processes/diagram", _Workflow_DrawProcesses0_HTTP_Handler(srv))
	r.DELETE("/v1/workflow/process/{id}", _Workflow_DeleteProcess0_HTTP_Handler(srv))
	r.DELETE("/v1/workflow/{workflow_id}/processes", _Workflow_DeleteProcesses0_HTTP_Handler(srv))
	r.GET("/helloworld/{name}", _Workflow_SayHello0_HTTP_Handler(srv))
}

func _Workflow_CreateWorkflow0_HTTP_Handler(srv WorkflowHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateWorkflowRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/greek.gflow.workflow.v1.Workflow/CreateWorkflow")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateWorkflow(ctx, req.(*CreateWorkflowRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateWorkflowResponse)
		return ctx.Result(200, reply)
	}
}

func _Workflow_GetWorkflowDetail0_HTTP_Handler(srv WorkflowHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetWorkflowDetailRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/greek.gflow.workflow.v1.Workflow/GetWorkflowDetail")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetWorkflowDetail(ctx, req.(*GetWorkflowDetailRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetWorkflowDetailResponse)
		return ctx.Result(200, reply)
	}
}

func _Workflow_GetWorkflows0_HTTP_Handler(srv WorkflowHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetWorkflowsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/greek.gflow.workflow.v1.Workflow/GetWorkflows")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetWorkflows(ctx, req.(*GetWorkflowsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetWorkflowsResponse)
		return ctx.Result(200, reply)
	}
}

func _Workflow_UpdateWorkflow0_HTTP_Handler(srv WorkflowHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateWorkflowRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/greek.gflow.workflow.v1.Workflow/UpdateWorkflow")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateWorkflow(ctx, req.(*UpdateWorkflowRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Workflow_DeleteWorkflow0_HTTP_Handler(srv WorkflowHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteWorkflowRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/greek.gflow.workflow.v1.Workflow/DeleteWorkflow")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteWorkflow(ctx, req.(*DeleteWorkflowRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Workflow_CreateProcess0_HTTP_Handler(srv WorkflowHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateProcessRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/greek.gflow.workflow.v1.Workflow/CreateProcess")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateProcess(ctx, req.(*CreateProcessRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateProcessResponse)
		return ctx.Result(200, reply)
	}
}

func _Workflow_UpdateProcess0_HTTP_Handler(srv WorkflowHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateProcessRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/greek.gflow.workflow.v1.Workflow/UpdateProcess")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateProcess(ctx, req.(*UpdateProcessRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Workflow_DrawProcesses0_HTTP_Handler(srv WorkflowHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DrawProcessesRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/greek.gflow.workflow.v1.Workflow/DrawProcesses")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DrawProcesses(ctx, req.(*DrawProcessesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Workflow_DeleteProcess0_HTTP_Handler(srv WorkflowHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteProcessRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/greek.gflow.workflow.v1.Workflow/DeleteProcess")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteProcess(ctx, req.(*DeleteProcessRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Workflow_DeleteProcesses0_HTTP_Handler(srv WorkflowHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteProcessesRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/greek.gflow.workflow.v1.Workflow/DeleteProcesses")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteProcesses(ctx, req.(*DeleteProcessesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Workflow_SayHello0_HTTP_Handler(srv WorkflowHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HelloRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/greek.gflow.workflow.v1.Workflow/SayHello")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SayHello(ctx, req.(*HelloRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HelloReply)
		return ctx.Result(200, reply)
	}
}

type WorkflowHTTPClient interface {
	CreateProcess(ctx context.Context, req *CreateProcessRequest, opts ...http.CallOption) (rsp *CreateProcessResponse, err error)
	CreateWorkflow(ctx context.Context, req *CreateWorkflowRequest, opts ...http.CallOption) (rsp *CreateWorkflowResponse, err error)
	DeleteProcess(ctx context.Context, req *DeleteProcessRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	DeleteProcesses(ctx context.Context, req *DeleteProcessesRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	DeleteWorkflow(ctx context.Context, req *DeleteWorkflowRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	DrawProcesses(ctx context.Context, req *DrawProcessesRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	GetWorkflowDetail(ctx context.Context, req *GetWorkflowDetailRequest, opts ...http.CallOption) (rsp *GetWorkflowDetailResponse, err error)
	GetWorkflows(ctx context.Context, req *GetWorkflowsRequest, opts ...http.CallOption) (rsp *GetWorkflowsResponse, err error)
	SayHello(ctx context.Context, req *HelloRequest, opts ...http.CallOption) (rsp *HelloReply, err error)
	UpdateProcess(ctx context.Context, req *UpdateProcessRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	UpdateWorkflow(ctx context.Context, req *UpdateWorkflowRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
}

type WorkflowHTTPClientImpl struct {
	cc *http.Client
}

func NewWorkflowHTTPClient(client *http.Client) WorkflowHTTPClient {
	return &WorkflowHTTPClientImpl{client}
}

func (c *WorkflowHTTPClientImpl) CreateProcess(ctx context.Context, in *CreateProcessRequest, opts ...http.CallOption) (*CreateProcessResponse, error) {
	var out CreateProcessResponse
	pattern := "/v1/workflow/process"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/greek.gflow.workflow.v1.Workflow/CreateProcess"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *WorkflowHTTPClientImpl) CreateWorkflow(ctx context.Context, in *CreateWorkflowRequest, opts ...http.CallOption) (*CreateWorkflowResponse, error) {
	var out CreateWorkflowResponse
	pattern := "/v1/workflow"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/greek.gflow.workflow.v1.Workflow/CreateWorkflow"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *WorkflowHTTPClientImpl) DeleteProcess(ctx context.Context, in *DeleteProcessRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/v1/workflow/process/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/greek.gflow.workflow.v1.Workflow/DeleteProcess"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *WorkflowHTTPClientImpl) DeleteProcesses(ctx context.Context, in *DeleteProcessesRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/v1/workflow/{workflow_id}/processes"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/greek.gflow.workflow.v1.Workflow/DeleteProcesses"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *WorkflowHTTPClientImpl) DeleteWorkflow(ctx context.Context, in *DeleteWorkflowRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/v1/workflow/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/greek.gflow.workflow.v1.Workflow/DeleteWorkflow"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *WorkflowHTTPClientImpl) DrawProcesses(ctx context.Context, in *DrawProcessesRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/v1/workflow/processes/diagram"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/greek.gflow.workflow.v1.Workflow/DrawProcesses"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *WorkflowHTTPClientImpl) GetWorkflowDetail(ctx context.Context, in *GetWorkflowDetailRequest, opts ...http.CallOption) (*GetWorkflowDetailResponse, error) {
	var out GetWorkflowDetailResponse
	pattern := "/v1/workflow/detail/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/greek.gflow.workflow.v1.Workflow/GetWorkflowDetail"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *WorkflowHTTPClientImpl) GetWorkflows(ctx context.Context, in *GetWorkflowsRequest, opts ...http.CallOption) (*GetWorkflowsResponse, error) {
	var out GetWorkflowsResponse
	pattern := "/v1/workflows"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/greek.gflow.workflow.v1.Workflow/GetWorkflows"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *WorkflowHTTPClientImpl) SayHello(ctx context.Context, in *HelloRequest, opts ...http.CallOption) (*HelloReply, error) {
	var out HelloReply
	pattern := "/helloworld/{name}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/greek.gflow.workflow.v1.Workflow/SayHello"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *WorkflowHTTPClientImpl) UpdateProcess(ctx context.Context, in *UpdateProcessRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/v1/workflow/process/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/greek.gflow.workflow.v1.Workflow/UpdateProcess"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *WorkflowHTTPClientImpl) UpdateWorkflow(ctx context.Context, in *UpdateWorkflowRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/v1/workflow/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/greek.gflow.workflow.v1.Workflow/UpdateWorkflow"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}