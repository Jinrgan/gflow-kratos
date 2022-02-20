package shift

import (
	pb "gflow-kratos/api/shift/v1"
	"gflow-kratos/pkg/id"
	"time"
)

type Log struct {
	DateSec           time.Time `gorm:"index"`
	HandoverId        id.ShiftId
	OrganizationId    id.OrganizationId
	DepartmentId      id.DepartmentId
	LocationId        string
	ShiftTypeId       string
	TeamId            string
	RoleId            string
	ReporterId        id.UserId
	SpeechRecognition string
	SuccessorId       id.UserId
	PicturePaths      string
	PlanType          pb.PlanType
	TaskId            string
}
