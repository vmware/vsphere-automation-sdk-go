// Copyright (c) 2021-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

package tasks

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/vmware/vsphere-automation-sdk-go/lib/cis"
	"github.com/vmware/vsphere-automation-sdk-go/lib/cis/task"
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std"
	vapiErrors "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/common"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/log"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/server/tracing"
)

// CreateSpec provides information on tasks specifics when it is initialized by CreateTask method
//
//	BaseId - dictates how a task id would look like. If not set a newly generated UUID is used;
//		next to the BaseId service id is concatenated, separated by managerOptions.TaskIdDelimiter
//	Description - simple task description string
//	Cancelable - boolean specifying whether a task can be canceled or not
type CreateSpec struct {
	BaseId      string
	Description std.LocalizableMessage
	Cancelable  bool
}

// Manager provides functionality to create, update, and delete tasks.
// Interface is to be used by provider implementations to also start a task and get its info.
type Manager interface {
	core.TaskResultProvider

	// CreateTask creates new task object and returns its id.
	// To start already created task use StartTaskWithContext method.
	//
	// The given context must contain the values of ExecutionContext.Context(),
	// i.e. the context of the API operation which triggered the task.
	CreateTask(ctx context.Context, createSpec CreateSpec) (string, error)

	// StartTask by given task id marks task as started
	// Deprecated: Use StartTaskWithContext
	StartTask(taskId string, progress *task.Progress) error

	// StartTaskWithContext marks the given task as started.
	//
	// Normally a task is executed on a separate goroutine, asynchronously from
	// the API operation which triggered the task. StartTaskWithContext should
	// be invoked at the beginning of the task goroutine, so that time is
	// tracked correctly.
	//
	// The given context must contain the values of ExecutionContext.Context(),
	// i.e. the context of the API operation which triggered the task, so that
	// cross-cutting concerns work properly (e.g. distributed tracing).
	//
	// The returned context or its descendant must be passed to outgoing API
	// calls within the business logic of the task, so that cross-cutting
	// concerns work properly (e.g. distributed tracing).
	// The returned context contains the values of the input context without
	// cancellation/deadline/timeout.
	StartTaskWithContext(ctx context.Context, taskId string, progress *task.Progress) (context.Context, error)

	GetTaskInfo(taskId string) (*task.Info, error)

	GetTaskInfos(filterSpecParam *cis.TasksFilterSpec, resultSpecParam *cis.TasksGetSpec) (map[string]task.Info, error)

	// SetTaskInfo updates the status of a given task.
	// Do not use this method to start or end a task, instead use StartTaskWithContext, SetResult, and SetError
	SetTaskInfo(taskId string, taskInfo *task.Info) error

	// MarkForCancellation marks given task for cancellation. Task must be created as cancellable and have not finished yet.
	// Method is used in accordance with IsMarkedForCancellation. The latter is used to check if a task is marked for
	// cancellation in certain part of the code. If that is the case, implementor should do the necessary work to actually
	// cancel the task.
	MarkForCancellation(taskId string) error

	// IsMarkedForCancellation verifies whether a task has already been marked for cancellation by MarkForCancellation method.
	IsMarkedForCancellation(taskId string) bool

	// SetError ends a task in error.
	//
	// The given context must be the context returned by StartTaskWithContext
	// or a descendant of that context.
	SetError(_ context.Context, taskId string, errorParam error) error

	// SetResult ends a task and sets its result to provided parameter.
	//
	// The given context must be the context returned by StartTaskWithContext
	// or a descendant of that context.
	SetResult(ctx context.Context, taskId string, result interface{}) error
}

type serverTaskWrapper struct {
	info    *task.Info
	done    chan bool
	expired bool
	span    tracing.ServerSpan
	lock    sync.RWMutex
}

// InMemoryManager is a default implementation of core.TaskManager
// It provides functionality to:
//   - create task - by using CreateTask method and provided necessary CreateSpec information
//   - start task - by using StartTaskWithContext method
//   - edit task - done by SetError method
//   - end task - by using either SetError or SetResult methods
//   - retrieve task information - through GetTaskInfo or GetTasksInfo methods
//   - mark task for cancellation - by using MarkForCancellation, at which point a task provider can listen for this
//     event and do the actual task cancel actions
//   - retrieve task result information - GetTaskResult is used by bindings to block on waiting for a task end
type InMemoryManager struct {
	options       *managerOptions
	lock          sync.RWMutex
	tasks         map[string]*serverTaskWrapper
	tasksToCancel map[string]*serverTaskWrapper
}

var _ Manager = &InMemoryManager{}

// managerOptions defines basic task manager options
type managerOptions struct {
	TaskIdDelimiter      string
	TaskExpirationTime   time.Duration
	TaskPreservationTime time.Duration
	TypeConverter        *bindings.TypeConverter
	tracer               tracing.StartTaskSpan
}

// managerOption defines function which can modify default managerOptions
type managerOption func(options *managerOptions)

// WithTaskExpirationTime InMemoryManager helper method allowing to override default task expiration time
func WithTaskExpirationTime(taskExpirationTime time.Duration) managerOption {
	return func(options *managerOptions) {
		options.TaskExpirationTime = taskExpirationTime
	}
}

// WithTaskPreservationTime InMemoryManager helper method allowing to override default task preservation time
func WithTaskPreservationTime(taskPreservationTime time.Duration) managerOption {
	return func(options *managerOptions) {
		options.TaskPreservationTime = taskPreservationTime
	}
}

// WithTaskIdDelimiter InMemoryManager helper method allowing to override default task id delimiter
func WithTaskIdDelimiter(delimiter string) managerOption {
	return func(options *managerOptions) {
		options.TaskIdDelimiter = delimiter
	}
}

// WithTracer configures distributed tracing for long-running operations
func WithTracer(tracer tracing.StartTaskSpan) managerOption {
	return func(options *managerOptions) {
		options.tracer = tracer
	}
}

// NewInMemoryManager initializes new InMemoryManager with following default options
// - task expiration time - 24 hours
// - task preservation time or time for which task is preserved in memory after it has finished - 7 minutes
// - task id delimiter - ':'
func NewInMemoryManager(opts ...managerOption) *InMemoryManager {
	m := &InMemoryManager{
		// default options
		options: &managerOptions{
			TaskIdDelimiter:      ":",
			TaskExpirationTime:   24 * time.Hour,
			TaskPreservationTime: 7 * time.Minute,
			TypeConverter:        bindings.NewTypeConverter(),
			tracer:               tracing.NoopTaskTracer,
		},
		tasks:         make(map[string]*serverTaskWrapper),
		tasksToCancel: make(map[string]*serverTaskWrapper),
	}

	for _, opt := range opts {
		opt(m.options)
	}

	return m
}

func (m *InMemoryManager) GetTaskResult(taskId string) (core.MonoResult, error) {
	taskObj, err := m.getTask(taskId)
	if err != nil {
		return nil, err
	}

	<-taskObj.done

	return core.NewMonoResult(taskObj.info.Result, taskObj.info.Error_), nil
}

// CreateTask creates new task object and returns its id.
// To start already created task use StartTaskWithContext method.
func (m *InMemoryManager) CreateTask(ctx context.Context, createSpec CreateSpec) (string, error) {
	taskCtx, err := core.GetTaskInvocationContext(ctx)
	if err != nil {
		return "", err
	}

	taskInfo := &task.Info{
		Description: createSpec.Description,
		Service:     taskCtx.ServiceId,
		Operation:   taskCtx.OperationId,
		Status:      task.Status_PENDING,
		Cancelable:  createSpec.Cancelable,
	}

	baseId := createSpec.BaseId
	if baseId == "" {
		baseId = common.NewUUID()
	}
	taskId := baseId + m.options.TaskIdDelimiter + taskInfo.Service
	m.tasks[taskId] = &serverTaskWrapper{
		info: taskInfo,
		done: make(chan bool),
		// make sure the span is never nil
		span: tracing.NoopServerSpan,
	}

	return taskId, nil
}

func (m *InMemoryManager) getTask(taskId string) (*serverTaskWrapper, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	if val, ok := m.tasks[taskId]; ok {
		return val, nil
	}

	errMsg := fmt.Sprintf("Task with id %s not found", taskId)
	return nil, getLocalizableError(errMsg, "notFound", "vapi.task.not_found", []string{taskId})
}

// StartTask by given task id marks task as started
// Deprecated: Use StartTaskWithContext
func (m *InMemoryManager) StartTask(taskId string, progress *task.Progress) error {
	_, err := m.StartTaskWithContext(context.Background(), taskId, progress)
	return err
}

func (m *InMemoryManager) StartTaskWithContext(ctx context.Context, taskId string, progress *task.Progress) (context.Context, error) {
	taskObj, err := m.getTask(taskId)
	if err != nil {
		return nil, err
	}

	taskObj.lock.Lock()
	defer taskObj.lock.Unlock()

	traceCtx, span := m.options.tracer(ctx, taskObj.info.Service, taskObj.info.Operation, taskId)
	taskObj.span = span

	taskCtx := withoutCancellation(traceCtx)

	taskObj.info.Status = task.Status_RUNNING
	timeNow := time.Now().UTC()
	taskObj.info.StartTime = &timeNow
	taskObj.info.Progress = progress

	// start goroutine to track expiration of task
	go func(taskId string, taskObj *serverTaskWrapper) {
		log.Info("Starting go routine to track task expiration")
		select {
		case <-time.After(m.options.TaskExpirationTime):
			errMsg := fmt.Sprintf("Task %s expired", taskId)
			vapiErr := getLocalizableError(errMsg, "internalServer", "vapi.task.expired_task", []string{taskId})
			errVal, _ := vapiErr.(*vapiErrors.InternalServerError).GetDataValue__()
			_ = m.endTask(taskId, nil, errVal.(*data.ErrorValue), task.Status_FAILED)
			return
		case <-taskObj.done:
			// task finished, end routine
			return
		}
	}(taskId, taskObj)

	return taskCtx, nil
}

func (m *InMemoryManager) GetTaskInfo(taskId string) (*task.Info, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	if val, ok := m.tasks[taskId]; ok {
		taskInfoCopy, err := m.getTaskInfoCopy(val.info)
		if err != nil {
			return nil, err
		}
		return taskInfoCopy, nil
	}

	errMsg := fmt.Sprintf("Task with id %s not found", taskId)
	return nil, getLocalizableError(errMsg, "notFound", "vapi.task.not_found", []string{taskId})
}

func (m *InMemoryManager) GetTaskInfos(filterSpecParam *cis.TasksFilterSpec, resultSpecParam *cis.TasksGetSpec) (map[string]task.Info, error) {
	if filterSpecParam != nil {
		errMsg := fmt.Sprintf("filterSpecParam logic is not implemented yet")
		return nil, getLocalizableError(errMsg, "internalServer", "vapi.task.filterSpecParam_not_implemented_yet", nil)
	}

	result := make(map[string]task.Info, len(m.tasks))

	m.lock.RLock()
	defer m.lock.RUnlock()

	for taskId, taskItem := range m.tasks {
		taskCopy, err := m.getTaskInfoCopy(taskItem.info)
		if err != nil {
			return nil, err
		}
		if resultSpecParam != nil && resultSpecParam.ExcludeResult != nil && *resultSpecParam.ExcludeResult {
			taskCopy = nil
		}
		result[taskId] = *taskCopy
	}

	return result, nil
}

// getTaskInfoCopy uses type converter to produce copy instance by converting
// to data value and back to native type
func (m *InMemoryManager) getTaskInfoCopy(info *task.Info) (*task.Info, error) {
	dataVal, errs := info.GetDataValue__()
	if errs != nil {
		errsStr := common.GetErrorsString(errs)
		errMsg := fmt.Sprintf("Error occurred while getting task info data value: %s", errsStr)
		return nil, getLocalizableError(errMsg, "internalServer", "vapi.task.convert.data_value", []string{errsStr})
	}

	infoCopy, errs := m.options.TypeConverter.ConvertToGolang(dataVal, task.InfoBindingType())
	if errs != nil {
		errMsg := fmt.Sprintf("error occurred while convertin task info to bindings type")
		return nil, getLocalizableError(errMsg, "internalServer", "vapi.task.convert.binding_type", nil)
	}
	infoCopyVal := infoCopy.(task.Info)

	return &infoCopyVal, nil
}

// SetTaskInfo updates the status of a given task.
// Do not use this method to start or end a task, instead use StartTaskWithContext, SetResult, and SetError
func (m *InMemoryManager) SetTaskInfo(taskId string, taskInfo *task.Info) error {
	prevInfo, err := m.GetTaskInfo(taskId)
	if err != nil {
		return err
	}

	if prevInfo.Status != taskInfo.Status {
		switch taskInfo.Status {
		case task.Status_RUNNING:
			errMsg := "use manager's StartTaskWithContext to set running status"
			return getLocalizableError(errMsg, "internalServer", "vapi.task.invalid_status_running", nil)
		case task.Status_FAILED:
			errMsg := "you cannot set failed status directly; use manager's SetError instead"
			return getLocalizableError(errMsg, "internalServer", "vapi.task.invalid_status_failed", nil)
		case task.Status_SUCCEEDED:
			errMsg := "you cannot set succeeded status directly; use manager's SetResult instead"
			return getLocalizableError(errMsg, "internalServer", "vapi.task.invalid_status_succeeded", nil)
		}
	}

	taskObj, err := m.getTask(taskId)
	if err != nil {
		return err
	}

	taskObj.lock.Lock()
	defer taskObj.lock.Unlock()
	taskObj.info = taskInfo

	return nil
}

// MarkForCancellation marks given task for cancellation. Task must be created as cancellable and have not finished yet.
// Method is used in accordance with IsMarkedForCancellation. The latter is used to check if a task is marked for
// cancellation in certain part of the code. If that is the case, implementor should do the necessary work to actually
// cancel the task.
func (m *InMemoryManager) MarkForCancellation(taskId string) error {
	taskObj, err := m.getTask(taskId)
	if err != nil {
		log.Error("Error getting task: %v, error: %v", taskId, err)
		return err
	}

	m.lock.Lock()
	defer m.lock.Unlock()

	if !taskObj.info.Cancelable {
		errMsg := fmt.Sprintf("task %s is not cancelable", taskId)
		return getLocalizableError(errMsg, "invalidArgument", "vapi.task.task_non_cancelable", []string{taskId})
	}

	if taskObj.info.Status == task.Status_FAILED || taskObj.info.Status == task.Status_SUCCEEDED {
		errMsg := fmt.Sprintf("task %s is in invalid status for cancelation", taskId)
		return getLocalizableError(errMsg, "invalidArgument", "vapi.task.task_non_cancelable", []string{taskId})
	}

	m.tasksToCancel[taskId] = taskObj
	return nil
}

func (m *InMemoryManager) unmarkForCancellation(taskId string) {
	if m.IsMarkedForCancellation(taskId) {
		m.lock.Lock()
		delete(m.tasksToCancel, taskId)
		m.lock.Unlock()
	}
}

// IsMarkedForCancellation verifies whether a task has already be marked for cancellation by MarkForCancellation method.
func (m *InMemoryManager) IsMarkedForCancellation(taskId string) bool {
	m.lock.RLock()
	defer m.lock.RUnlock()

	if _, ok := m.tasksToCancel[taskId]; ok {
		return true
	}
	return false
}

// SetError ends a task in error.
func (m *InMemoryManager) SetError(_ context.Context, taskId string, errorParam error) error {
	err := m.endTask(taskId, nil, bindings.GetVAPIError(errorParam), task.Status_FAILED)
	if err != nil {
		return err
	}

	return nil
}

// SetResult ends a task and sets its result to provided parameter.
func (m *InMemoryManager) SetResult(ctx context.Context, taskId string, result interface{}) error {
	taskCtx, err := core.GetTaskInvocationContext(ctx)
	if err != nil {
		return err
	}

	dataValue, errs := m.options.TypeConverter.ConvertToVapi(result, taskCtx.OutputBindingType)
	if errs != nil {
		vapiErrs := common.GetErrorsString(errs)
		errMsg := fmt.Sprintf("error occurred while converting task info to data value: %s", vapiErrs)
		return getLocalizableError(errMsg, "internalServer", "vapi.task.convert.data_value_2", []string{vapiErrs})
	}

	err = m.endTask(taskId, dataValue, nil, task.Status_SUCCEEDED)
	if err != nil {
		return err
	}

	return nil
}

func (m *InMemoryManager) endTask(taskId string,
	output data.DataValue,
	errorValue *data.ErrorValue,
	status task.StatusEnum) error {

	taskObj, err := m.getTask(taskId)
	if err != nil {
		return err
	}

	taskObj.lock.Lock()
	defer taskObj.lock.Unlock()

	if taskObj.info.Status == task.Status_FAILED || taskObj.info.Status == task.Status_SUCCEEDED {
		errMsg := fmt.Sprintf("task %s has already finished with status %s.", taskId, taskObj.info.Status)
		return getLocalizableError(errMsg, "internalServer", "vapi.task.task_already_finished", []string{taskId})
	}

	timeNow := time.Now().UTC()
	if errorValue != nil {
		taskObj.info.Error_ = errorValue
		taskObj.span.LogVapiError(errorValue)
	}
	if output != nil {
		taskObj.info.Result = output
	}
	taskObj.info.Status = status
	taskObj.info.EndTime = &timeNow

	taskObj.done <- true
	close(taskObj.done)

	taskObj.span.Finish()

	m.unmarkForCancellation(taskId)
	m.igniteTaskRemoval(taskId)

	return nil
}

func (m *InMemoryManager) igniteTaskRemoval(taskId string) {
	log.Info("Starting go routine for task removal")
	time.AfterFunc(m.options.TaskPreservationTime, func() {
		log.Info(fmt.Sprintf("Removing task %s", taskId))

		m.lock.Lock()
		defer m.lock.Unlock()

		if _, ok := m.tasks[taskId]; ok {
			delete(m.tasks, taskId)
		}
	})
}

func getLocalizableError(defaultMessage, errorType, messageId string, arguments []string) error {

	log.Error(defaultMessage)

	switch errorType {
	case "notFound":
		err := vapiErrors.NewNotFound()
		err.Messages = []std.LocalizableMessage{
			{
				Id:             messageId,
				DefaultMessage: defaultMessage,
				Args:           arguments,
			},
		}
		return err
	case "internalServer":
		fallthrough
	default:
		err := vapiErrors.NewInternalServerError()
		err.Messages = []std.LocalizableMessage{
			{
				Id:             messageId,
				DefaultMessage: defaultMessage,
				Args:           arguments,
			},
		}
		return err
	}
}

var _ cis.TasksProvider = &InMemoryManager{}

// Get gets in-memory task by given id.
func (m *InMemoryManager) Get(taskId string, getSpec *cis.TasksGetSpec, _ *core.ExecutionContext) (task.Info, error) {
	log.Debugf("Getting info for task: %v", taskId)
	info, err := m.GetTaskInfo(taskId)
	if err != nil {
		log.Error("Error getting info for task: %v", taskId)
		return task.Info{}, err
	}
	if getSpec != nil && getSpec.ExcludeResult != nil && *getSpec.ExcludeResult {
		info.Result = nil
	}
	return *info, nil
}

// List lists all in-memory tasks
func (m *InMemoryManager) List(filterSpecParam *cis.TasksFilterSpec, resultSpecParam *cis.TasksGetSpec, _ *core.ExecutionContext) (map[string]task.Info, error) {
	log.Info("Listing tasks")
	return m.GetTaskInfos(filterSpecParam, resultSpecParam)
}

// Cancel attempts to cancel a task by given id
func (m *InMemoryManager) Cancel(taskId string, _ *core.ExecutionContext) error {
	log.Info("Canceling task: %v", taskId)
	return m.MarkForCancellation(taskId)
}

// withoutCancellation returns a context with the same values as the given
// context but without its cancellation/deadline/timeout.
//
// TODO: Remove this code if the standard library introduces an alternative:
// https://github.com/golang/go/issues/40221
func withoutCancellation(parent context.Context) context.Context {
	return &detachedCtx{parent}
}

type detachedCtx struct {
	context.Context
}

func (c *detachedCtx) Deadline() (deadline time.Time, ok bool) {
	return
}
func (c *detachedCtx) Done() <-chan struct{} {
	return nil
}
func (c *detachedCtx) Err() error {
	return nil
}
func (c *detachedCtx) Value(key interface{}) interface{} {
	return c.Context.Value(key)
}
