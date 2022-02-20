package mysql

import (
	"fmt"
	"gflow-kratos/pkg/id"
	"strconv"
)

type ObjectId uint

func (id ObjectId) String() string {
	return strconv.FormatUint(uint64(id), 10)
}

// ObjectIDFromId converts an id to objected id.
func ObjectIDFromId(id fmt.Stringer) (ObjectId, error) {
	oid, err := strconv.Atoi(id.String())
	if err != nil {
		return 0, err
	}

	return ObjectId(oid), err
}

// ObjectIdMustFromId converts an id to objected id, panics on error.
func ObjectIdMustFromId(id fmt.Stringer) ObjectId {
	oid, err := ObjectIDFromId(id)
	if err != nil {
		panic(err)
	}

	return oid
}

// ObjectIdToUserId converts object id to user id.
func ObjectIdToUserId(oid ObjectId) id.UserId {
	return id.UserId(strconv.FormatUint(uint64(oid), 10))
}

func ObjectIdToOrganizationId(oid ObjectId) id.OrganizationId {
	return id.OrganizationId(strconv.FormatUint(uint64(oid), 10))
}

func ObjectIdToDepartmentId(oid ObjectId) id.DepartmentId {
	return id.DepartmentId(strconv.FormatUint(uint64(oid), 10))
}

// ObjectIdToRouterId converts object id to router id.
func ObjectIdToRouterId(oid ObjectId) id.RouterId {
	return id.RouterId(strconv.FormatUint(uint64(oid), 10))
}

// ObjectIdToWorkflowId converts object id to workflow id.
func ObjectIdToWorkflowId(oid ObjectId) id.WorkflowId {
	return id.WorkflowId(strconv.FormatUint(uint64(oid), 10))
}

func ObjectIdToWorkflowProcessId(oid ObjectId) id.WorkflowProcessId {
	return id.WorkflowProcessId(strconv.FormatUint(uint64(oid), 10))
}

func ObjectIdToWorkflowNewId(oid ObjectId) id.WorkflowNewId {
	return id.WorkflowNewId(strconv.FormatUint(uint64(oid), 10))
}

func ObjectIdToWorkflowInstanceId(oid ObjectId) id.WorkflowInstanceId {
	return id.WorkflowInstanceId(strconv.FormatUint(uint64(oid), 10))
}

func ObjectIdToWorkflowProcessInstanceId(oid ObjectId) id.WorkflowProcessInstanceId {
	return id.WorkflowProcessInstanceId(strconv.FormatUint(uint64(oid), 10))
}

func ObjectIdToWorkflowLogId(oid ObjectId) id.WorkflowLogId {
	return id.WorkflowLogId(strconv.FormatUint(uint64(oid), 10))
}

func ObjectIDToWorkflowApproverId(oid ObjectId) id.WorkflowApproverId {
	return id.WorkflowApproverId(strconv.FormatUint(uint64(oid), 10))
}

func ObjectIdToHandoverId(oid ObjectId) id.ShiftId {
	return id.ShiftId(strconv.FormatUint(uint64(oid), 10))
}
