package service

import (
	"context"
	pb "gflow-kratos/api/workflow/v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *WorkflowService) CreateWorkflow(ctx context.Context, req *pb.CreateWorkflowRequest) (*pb.CreateWorkflowResponse, error) {
	panic("implement me")
}

func (s *WorkflowService) GetWorkflowDetail(ctx context.Context, req *pb.GetWorkflowDetailRequest) (*pb.GetWorkflowDetailResponse, error) {
	panic("implement me")
}

func (s *WorkflowService) GetWorkflows(ctx context.Context, req *pb.GetWorkflowsRequest) (*pb.GetWorkflowsResponse, error) {
	panic("implement me")
}

func (s *WorkflowService) UpdateWorkflow(ctx context.Context, req *pb.UpdateWorkflowRequest) (*emptypb.Empty, error) {
	panic("implement me")
}

func (s *WorkflowService) DeleteWorkflow(ctx context.Context, req *pb.DeleteWorkflowRequest) (*emptypb.Empty, error) {
	panic("implement me")
}
