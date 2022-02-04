package service

import (
	"context"
	pb "gflow-kratos/api/workflow/v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *WorkflowService) CreateProcess(ctx context.Context, req *pb.CreateProcessRequest) (*pb.CreateProcessResponse, error) {
	panic("implement me")
}

func (s *WorkflowService) UpdateProcess(ctx context.Context, req *pb.UpdateProcessRequest) (*emptypb.Empty, error) {
	panic("implement me")
}

func (s *WorkflowService) DrawProcesses(ctx context.Context, req *pb.DrawProcessesRequest) (*emptypb.Empty, error) {
	panic("implement me")
}

func (s *WorkflowService) DeleteProcess(ctx context.Context, req *pb.DeleteProcessRequest) (*emptypb.Empty, error) {
	panic("implement me")
}

func (s *WorkflowService) DeleteProcesses(ctx context.Context, req *pb.DeleteProcessesRequest) (*emptypb.Empty, error) {
	panic("implement me")
}
