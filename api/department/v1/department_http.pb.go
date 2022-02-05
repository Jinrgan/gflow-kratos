// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.3

package pb

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

type DepartmentHTTPServer interface {
	GetDepartments(context.Context, *emptypb.Empty) (*GetDepartmentsResponse, error)
}

func RegisterDepartmentHTTPServer(s *http.Server, srv DepartmentHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/departments", _Department_GetDepartments0_HTTP_Handler(srv))
}

func _Department_GetDepartments0_HTTP_Handler(srv DepartmentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/greek.gflow.department.v1.Department/GetDepartments")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetDepartments(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetDepartmentsResponse)
		return ctx.Result(200, reply)
	}
}

type DepartmentHTTPClient interface {
	GetDepartments(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *GetDepartmentsResponse, err error)
}

type DepartmentHTTPClientImpl struct {
	cc *http.Client
}

func NewDepartmentHTTPClient(client *http.Client) DepartmentHTTPClient {
	return &DepartmentHTTPClientImpl{client}
}

func (c *DepartmentHTTPClientImpl) GetDepartments(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*GetDepartmentsResponse, error) {
	var out GetDepartmentsResponse
	pattern := "/v1/departments"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/greek.gflow.department.v1.Department/GetDepartments"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}