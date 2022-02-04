package data

import (
	"context"
	pb "gflow-kratos/api/workflow/v1"
	"gflow-kratos/app/workflow/internal/biz"
	"gflow-kratos/pkg/id"
	"gflow-kratos/pkg/mysql"

	"github.com/go-kratos/kratos/v2/log"

	"gorm.io/gorm"
)

type Process struct {
	gorm.Model
	WorkflowId           mysql.ObjectId
	Type                 pb.ProcessType
	Name                 string
	WorkflowName         string
	ApprovalType         pb.ProcessApprovalType
	ApproverIds          []string // TODO: use identity define
	ApproverText         []string
	AllowCountersigned   bool
	AllowGoBack          bool
	NextProcessCondition string
	XCoordinate          int32
	YCoordinate          int32
	NextProcessIds       []id.WorkflowProcessId
	Dateline             int32
	Mode                 pb.ProcessMode
}

type ProcessRepo struct {
	data *Data
	log  *log.Helper
}

func NewProcessRepo(data *Data, logger log.Logger) biz.ProcessRepo {
	return &ProcessRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (p *ProcessRepo) GetProcess(ctx context.Context, pId id.WorkflowProcessId) (*biz.ProcessEntity, error) {
	panic("implement me")
}

func (p *ProcessRepo) GetProcessesByWorkflowId(ctx context.Context, fId id.WorkflowId) ([]*biz.ProcessEntity, error) {
	panic("implement me")
}
