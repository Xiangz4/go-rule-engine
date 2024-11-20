package entity

import (
	"fmt"
	"go-rule-engine/flowengine/context"
)

type Task struct {
	ID string `yaml:"id,omitempty" json:"id,omitempty"`
	// 任务名称
	Name        string    `yaml:"name,omitempty" json:"name,omitempty"`
	DependOn    []string  `yaml:"dependOn,omitempty" json:"dependOn,omitempty"`
	ActionName  string    `yaml:"actionName,omitempty" json:"actionName,omitempty"`
	TimeoutSecs int       `yaml:"timeoutSecs,omitempty" json:"timeoutSecs,omitempty"`
	PreChecks   PreChecks `yaml:"preChecks,omitempty" json:"preChecks,omitempty"`
}

// AllGetter
func (t *Task) GetID() string {
	return t.ID
}
func (t *Task) GetName() string {
	return t.Name
}
func (t *Task) GetDependOn() []string {
	return t.DependOn
}
func (t *Task) GetActionName() string {
	return t.ActionName
}
func (t *Task) GetPreChecks() PreChecks {
	return t.PreChecks
}

type PreChecks map[string]*Check

type Check struct {
	Conditions []TaskCondition `yaml:"conditions,omitempty" json:"conditions,omitempty"`
	Act        ActiveAction    `yaml:"active,omitempty" json:"active,omitempty"`
}

func (c *Check) IsMeet(dagIns *DagInstance) bool {
	for _, cd := range c.Conditions {
		if !cd.IsMeet(dagIns) {
			return false
		}
	}
	return true
}

type TaskInstance struct {
	BaseInfo    `bson:"inline"`
	TaskID      string             `json:"taskId,omitempty" bson:"taskID,omitempty"`
	DagInsID    string             `json:"dagInsID,omitempty" bson:"dagInsID,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	DependOn    []string           `json:"dependOn,omitempty" bson:"dependOn,omitempty"`
	ActionName  string             `json:"actionName,omitempty" bson:"actionName,omitempty"`
	TimeoutSecs int                `json:"timeoutSecs,omitempty" bson:"timeout_secs,omitempty"`
	Traces      []TraceInfo        `json:"traces,omitempty" bson:"traces,omitempty"`
	Status      TaskInstanceStatus `json:"status,omitempty" bson:"status,omitempty"`

	//used to save changes
	Patch              func(*TaskInstance) error `json:"-" bson:"-"`
	Context            context.ExecuteContext    `json:"-" bson:"-"`
	RelatedDagInstance *DagInstance              `json:"-" bson:"-"`

	bufTraceInfo []TraceInfo `json:"bufTraceInfo,omitempty" bson:"bufTraceInfo,omitempty"`
}

type TraceInfo struct {
	Time    int64  `json:"time,omitempty" bson:"time,omitempty"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}

// NewTaskInstance
func NewTaskInstance(dagInsId string, t Task) *TaskInstance {
	return &TaskInstance{
		TaskID:      t.ID,
		DagInsID:    dagInsId,
		Name:        t.Name,
		DependOn:    t.DependOn,
		ActionName:  t.ActionName,
		TimeoutSecs: t.TimeoutSecs,
		Status:      TaskInstanceStatusInit,
	}
}

// GetGraphID
func (t *TaskInstance) GetGraphID() string {
	return t.TaskID
}

// GetID
func (t *TaskInstance) GetID() string {
	return t.ID
}

// GetDepend
func (t *TaskInstance) GetDepend() []string {
	return t.DependOn
}

// GetStatus
func (t *TaskInstance) GetStatus() TaskInstanceStatus {
	return t.Status
}

func (t *TaskInstance) InitialDep(ctx context.ExecuteContext, patch func(*TaskInstance) error, dagIns *DagInstance) {
	t.Context = ctx
	t.Patch = patch
	t.RelatedDagInstance = dagIns
}

type TaskCondition struct {
	Key    string   `yaml:"key,omitempty" json:"key,omitempty"`
	Values []string `yaml:"values,omitempty" json:"values,omitempty"`
}

func (t *TaskInstance) SetStatus(s TaskInstanceStatus) error {
	t.Status = s
	patch := &TaskInstance{
		BaseInfo: BaseInfo{ID: t.ID},
		Status:   t.Status,
	}
	if len(t.bufTraceInfo) != 0 {
		patch.Traces = append(t.Traces, t.bufTraceInfo...)
	}
	return t.Patch(patch)
}

// Trace info
func (t *TaskInstance) Trace(msg string) {
	fmt.Println("traceInfo:", msg)
}

func (c *TaskCondition) IsMeet(dagIns *DagInstance) bool {
	return false
}

type ActiveAction string

type TaskInstanceStatus string

const (
	TaskInstanceStatusInit     TaskInstanceStatus = "init"
	TaskInstanceStatusCanceled TaskInstanceStatus = "canceled"
	TaskInstanceStatusRunning  TaskInstanceStatus = "running"
	TaskInstanceStatusEnding   TaskInstanceStatus = "ending"
	TaskInstanceStatusFailed   TaskInstanceStatus = "failed"
	TaskInstanceStatusRetrying TaskInstanceStatus = "retrying"
	TaskInstanceStatusSuccess  TaskInstanceStatus = "success"
	TaskInstanceStatusBlocked  TaskInstanceStatus = "blocked"
	TaskInstanceStatusContinue TaskInstanceStatus = "continue"
	TaskInstanceStatusSkipped  TaskInstanceStatus = "skipped"
)
