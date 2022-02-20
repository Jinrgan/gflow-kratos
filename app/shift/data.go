package shift

import (
	"context"
	"fmt"
	pb "gflow-kratos/api/shift/v1"
	"gflow-kratos/pkg/mysql"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type dbError interface {
	IsErrNoRowExist() bool
}

func IsErrNoRowExist(err error) bool {
	ex, ok := err.(dbError)

	return ok && ex.IsErrNoRowExist()
}

type errRowExist string

func (e errRowExist) Error() string {
	return string(e)
}

func (e errRowExist) IsErrNoRowExist() bool {
	return true
}

//Affair MCI schedule & plan define an affair
type Affair struct {
	MciScheduleId mysql.ObjectId `gorm:"primaryKey;type:int unsigned;not null"`
	PlanType      pb.PlanType    `gorm:"primaryKey;type:tinyint unsigned;not null;default:0"`
	PlanId        mysql.ObjectId `gorm:"primaryKey;type:int unsigned;not null"`
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
	DeletedAt     gorm.DeletedAt
}

type Mysql struct {
	DB *gorm.DB
}

//CreateAffair returns wrapped pkgMysql.ErDupKey when affair already exist
func (m *Mysql) CreateAffair(ctx context.Context, aff *pb.Affair) error {
	schId, err := strconv.Atoi(aff.MciScheduleId)
	if err != nil {
		return fmt.Errorf("cannot convey MCI schedule id from string: %w", err)
	}

	pId, err := strconv.Atoi(aff.PlanId)
	if err != nil {
		return fmt.Errorf("cannot convey plan id from string: %w", err)
	}

	data := &Affair{
		MciScheduleId: mysql.ObjectId(schId),
		PlanType:      aff.PlanType,
		PlanId:        mysql.ObjectId(pId),
	}

	if err := m.DB.Create(&data).Error; err != nil {
		return fmt.Errorf("cannot create affair: %w", err)
	}

	return nil
}

func (m *Mysql) GetShiftBySchedule(ctx context.Context, schId int32) (*pb.Shift, error) {
	panic("implement me")
}

func (m *Mysql) GetShifts(ctx context.Context) ([]*pb.Shift, error) {
	var shifts []*pb.Shift
	res := m.DB.Find(&shifts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return shifts, nil
}
