package data

import (
	"context"
	pb "gflow-kratos/api/workflow/v1"
	"gflow-kratos/app/workflow/internal/biz"
	"gflow-kratos/pkg/id"
	"gflow-kratos/pkg/mysql"

	"gorm.io/gorm"

	"github.com/go-kratos/kratos/v2/log"
)

type Workflow struct {
	gorm.Model
	Name      string
	Describe  string
	Type      pb.WorkflowType
	Status    pb.Status
	SortOrder uint32
	CreatedBy mysql.ObjectId
}

type WorkflowRepo struct {
	data *Data
	log  *log.Helper
}

// NewWorkflowRepo .
func NewWorkflowRepo(data *Data, logger log.Logger) biz.WorkflowRepo {
	return &WorkflowRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *WorkflowRepo) CreateWorkflow(ctx context.Context, wf *biz.Workflow) (id.WorkflowId, error) {
	panic("implement me")
}

func (r *WorkflowRepo) GetWorkflow(ctx context.Context, fId id.WorkflowId) (*biz.WorkflowEntity, error) {
	panic("implement me")
}

func (r *WorkflowRepo) GetWorkflows(ctx context.Context) ([]*biz.WorkflowEntity, error) {
	panic("implement me")
}

func (r *WorkflowRepo) UpdateWorkflow(ctx context.Context, fId id.WorkflowId, wf *biz.Workflow) (*biz.WorkflowEntity, error) {
	panic("implement me")
}

func (r *WorkflowRepo) DeleteWorkflow(ctx context.Context, fId id.WorkflowId) error {
	panic("implement me")
}
