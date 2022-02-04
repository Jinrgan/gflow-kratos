package biz

import (
	"context"
	pb "gflow-kratos/api/workflow/v1"
	"gflow-kratos/pkg/id"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type Workflow struct {
	Name      string
	Describe  string
	Type      pb.WorkflowType
	Status    pb.Status
	SortOrder uint32
	CreatedBy id.UserId
	CreatedAt time.Time
}

type WorkflowEntity struct {
	Id       id.WorkflowId
	Workflow *Workflow
}

type WorkflowRepo interface {
	CreateWorkflow(ctx context.Context, workflow *Workflow) (id.WorkflowId, error)
	GetWorkflow(ctx context.Context, fId id.WorkflowId) (*WorkflowEntity, error)
	GetWorkflows(ctx context.Context) ([]*WorkflowEntity, error)
	UpdateWorkflow(ctx context.Context, fId id.WorkflowId, workflow *Workflow) (*WorkflowEntity, error)
	DeleteWorkflow(ctx context.Context, fId id.WorkflowId) error
}

type WorkflowUseCase struct {
	repo WorkflowRepo
	log  *log.Helper
}

func NewWorkflowUseCase(repo WorkflowRepo, logger log.Logger) *WorkflowUseCase {
	return &WorkflowUseCase{repo: repo, log: log.NewHelper(logger)}
}
