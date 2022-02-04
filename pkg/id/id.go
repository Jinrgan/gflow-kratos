package id

//UserId defines user id object.
type UserId string

func (i UserId) String() string {
	return string(i)
}

type DepartmentId string

func (i DepartmentId) String() string {
	return string(i)
}

type RouterId string

func (i RouterId) String() string {
	return string(i)
}

//WorkflowId defines workflow id object.
type WorkflowId string

func (i WorkflowId) String() string {
	return string(i)
}

type WorkflowNewId string

func (i WorkflowNewId) String() string {
	return string(i)
}

type WorkflowProcessId string

func (i WorkflowProcessId) String() string {
	return string(i)
}

type WorkflowProcessStepId string

func (i WorkflowProcessStepId) String() string {
	return string(i)
}

type WorkflowInstanceId string

func (i WorkflowInstanceId) String() string {
	return string(i)
}

type WorkflowProcessInstanceId string

func (i WorkflowProcessInstanceId) String() string {
	return string(i)
}

type WorkflowLogId string

func (i WorkflowLogId) String() string {
	return string(i)
}

type WorkflowApproverId string

func (i WorkflowApproverId) String() string {
	return string(i)
}
