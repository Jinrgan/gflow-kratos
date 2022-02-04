package biz

import (
	"context"
	pb "gflow-kratos/api/workflow/v1"
	"gflow-kratos/pkg/id"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type ProcessDiagram struct {
	XCoordinate    int32
	YCoordinate    int32
	NextProcessIds []id.WorkflowProcessId
}

type Process struct {
	WorkflowId           id.WorkflowId
	Type                 pb.ProcessType
	Name                 string
	WorkflowName         string
	ApprovalType         pb.ProcessApprovalType
	ApproverIds          []string // TODO: use identity define
	ApproverText         []string
	AllowCountersigned   bool
	AllowGoBack          bool
	NextProcessCondition string
	Diagram              *ProcessDiagram
	UpdatedAt            *time.Time
	Dateline             int32
	Mode                 pb.ProcessMode
}

type ProcessEntity struct {
	Id      id.WorkflowProcessId
	Process *Process
}

type ProcessRepo interface {
	GetProcess(ctx context.Context, pId id.WorkflowProcessId) (*ProcessEntity, error)
	GetProcessesByWorkflowId(ctx context.Context, fId id.WorkflowId) ([]*ProcessEntity, error)
}

type ProcessUseCase struct {
	repo ProcessRepo
	log  *log.Helper
}

func NewProcessUseCase(repo ProcessRepo, logger log.Logger) *ProcessUseCase {
	return &ProcessUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "use-case/process")),
	}
}
