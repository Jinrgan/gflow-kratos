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

type ShiftServiceHTTPServer interface {
	CreateAffair(context.Context, *Affair) (*emptypb.Empty, error)
	CreateShift(context.Context, *CreateShiftRequest) (*emptypb.Empty, error)
	GetAffairByCurrentUser(context.Context, *emptypb.Empty) (*Affair, error)
	GetAffairs(context.Context, *emptypb.Empty) (*GetAffairsResponse, error)
	GetShifts(context.Context, *emptypb.Empty) (*GetShiftsResponse, error)
}

func RegisterShiftServiceHTTPServer(s *http.Server, srv ShiftServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/affair", _ShiftService_CreateAffair0_HTTP_Handler(srv))
	r.GET("/v1/affair/current-user", _ShiftService_GetAffairByCurrentUser0_HTTP_Handler(srv))
	r.GET("/v1/affairs", _ShiftService_GetAffairs0_HTTP_Handler(srv))
	r.POST("/v1/shift", _ShiftService_CreateShift0_HTTP_Handler(srv))
	r.GET("/v1/shifts", _ShiftService_GetShifts0_HTTP_Handler(srv))
}

func _ShiftService_CreateAffair0_HTTP_Handler(srv ShiftServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in Affair
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/bbm.mci.shift.v1.ShiftService/CreateAffair")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateAffair(ctx, req.(*Affair))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _ShiftService_GetAffairByCurrentUser0_HTTP_Handler(srv ShiftServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/bbm.mci.shift.v1.ShiftService/GetAffairByCurrentUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetAffairByCurrentUser(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Affair)
		return ctx.Result(200, reply)
	}
}

func _ShiftService_GetAffairs0_HTTP_Handler(srv ShiftServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/bbm.mci.shift.v1.ShiftService/GetAffairs")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetAffairs(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetAffairsResponse)
		return ctx.Result(200, reply)
	}
}

func _ShiftService_CreateShift0_HTTP_Handler(srv ShiftServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateShiftRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/bbm.mci.shift.v1.ShiftService/CreateShift")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateShift(ctx, req.(*CreateShiftRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _ShiftService_GetShifts0_HTTP_Handler(srv ShiftServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/bbm.mci.shift.v1.ShiftService/GetShifts")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetShifts(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetShiftsResponse)
		return ctx.Result(200, reply)
	}
}

type ShiftServiceHTTPClient interface {
	CreateAffair(ctx context.Context, req *Affair, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	CreateShift(ctx context.Context, req *CreateShiftRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	GetAffairByCurrentUser(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *Affair, err error)
	GetAffairs(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *GetAffairsResponse, err error)
	GetShifts(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *GetShiftsResponse, err error)
}

type ShiftServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewShiftServiceHTTPClient(client *http.Client) ShiftServiceHTTPClient {
	return &ShiftServiceHTTPClientImpl{client}
}

func (c *ShiftServiceHTTPClientImpl) CreateAffair(ctx context.Context, in *Affair, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/v1/affair"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/bbm.mci.shift.v1.ShiftService/CreateAffair"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ShiftServiceHTTPClientImpl) CreateShift(ctx context.Context, in *CreateShiftRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/v1/shift"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/bbm.mci.shift.v1.ShiftService/CreateShift"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ShiftServiceHTTPClientImpl) GetAffairByCurrentUser(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*Affair, error) {
	var out Affair
	pattern := "/v1/affair/current-user"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/bbm.mci.shift.v1.ShiftService/GetAffairByCurrentUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ShiftServiceHTTPClientImpl) GetAffairs(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*GetAffairsResponse, error) {
	var out GetAffairsResponse
	pattern := "/v1/affairs"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/bbm.mci.shift.v1.ShiftService/GetAffairs"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ShiftServiceHTTPClientImpl) GetShifts(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*GetShiftsResponse, error) {
	var out GetShiftsResponse
	pattern := "/v1/shifts"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/bbm.mci.shift.v1.ShiftService/GetShifts"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}