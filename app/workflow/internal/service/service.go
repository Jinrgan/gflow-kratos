package service

import (
	pb "gflow-kratos/api/workflow/v1"
	"gflow-kratos/app/workflow/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewWorkflowService)

// WorkflowService is a workflow service.
type WorkflowService struct {
	pb.UnimplementedWorkflowServer

	wc  *biz.WorkflowUseCase
	pc  *biz.ProcessUseCase
	log *log.Helper
}

// NewWorkflowService new a workflow service.
func NewWorkflowService(wc *biz.WorkflowUseCase, pc *biz.ProcessUseCase, logger log.Logger) *WorkflowService {
	return &WorkflowService{wc: wc, pc: pc, log: log.NewHelper(logger)}
}
