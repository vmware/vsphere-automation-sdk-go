// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Tasks.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package cis

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/cis/task"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/log"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
	"reflect"
)

// Resource type for task.
const Tasks_RESOURCE_TYPE = "com.vmware.cis.task"

// The ``GetSpec`` class describes what data should be included when retrieving information about a task.
type TasksGetSpec struct {
	// If true, all data, including operation-specific data, will be returned, otherwise only the data described in task.Info will be returned.
	ReturnAll *bool
	// If true, the result will not be included in the task information, otherwise it will be included.
	ExcludeResult *bool
}

func (s *TasksGetSpec) GetType__() bindings.BindingType {
	return TasksGetSpecBindingType()
}

func (s *TasksGetSpec) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for TasksGetSpec._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// The ``FilterSpec`` class contains properties used to filter the results when listing tasks (see Tasks#list). If multiple properties are specified, only tasks matching all of the properties match the filter.
//
//  Currently at least one of TasksFilterSpec#tasks or TasksFilterSpec#services must be specified and not empty.
type TasksFilterSpec struct {
	// Identifiers of tasks that can match the filter.
	Tasks map[string]bool
	// Identifiers of services. Tasks created by operations in these services match the filter (see task.CommonInfo#service).
	Services map[string]bool
	// Identifiers of operations. Tasks created by these operations match the filter (see task.CommonInfo#operation).
	//
	//  Note that an operation identifier by itself is not globally unique. To filter on an operation, the identifier of the service interface containing the operation should also be specified in ``services``.
	Operations map[string]bool
	// Status that a task must have to match the filter (see task.CommonInfo#status).
	Status map[task.StatusEnum]bool
	// Identifiers of the targets the operation for the associated task created or was performed on (see task.CommonInfo#target).
	Targets []std.DynamicID
	// Users who must have initiated the operation for the associated task to match the filter (see task.CommonInfo#user).
	Users map[string]bool
}

func (s *TasksFilterSpec) GetType__() bindings.BindingType {
	return TasksFilterSpecBindingType()
}

func (s *TasksFilterSpec) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for TasksFilterSpec._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

func tasksGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["task"] = bindings.NewIdType([]string{"com.vmware.cis.task"}, "")
	fields["spec"] = bindings.NewOptionalType(bindings.NewReferenceType(TasksGetSpecBindingType))
	fieldNameMap["task"] = "Task"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tasksGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(task.InfoBindingType)
}

func tasksGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["task"] = bindings.NewIdType([]string{"com.vmware.cis.task"}, "")
	fields["spec"] = bindings.NewOptionalType(bindings.NewReferenceType(TasksGetSpecBindingType))
	fieldNameMap["task"] = "Task"
	fieldNameMap["spec"] = "Spec"
	paramsTypeMap["task"] = bindings.NewIdType([]string{"com.vmware.cis.task"}, "")
	paramsTypeMap["spec.return_all"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["spec.exclude_result"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["task"] = bindings.NewIdType([]string{"com.vmware.cis.task"}, "")
	pathParams["task"] = "task"
	queryParams["spec.exclude_result"] = "exclude_result"
	queryParams["spec.return_all"] = "return_all"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"] = make(map[string]string)
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"]["challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/cis/tasks/{task}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.error": 500, "com.vmware.vapi.std.errors.not_found": 404, "com.vmware.vapi.std.errors.resource_inaccessible": 500, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.unauthenticated": 401, "com.vmware.vapi.std.errors.unauthorized": 403})
}

func tasksListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["filter_spec"] = bindings.NewOptionalType(bindings.NewReferenceType(TasksFilterSpecBindingType))
	fields["result_spec"] = bindings.NewOptionalType(bindings.NewReferenceType(TasksGetSpecBindingType))
	fieldNameMap["filter_spec"] = "FilterSpec"
	fieldNameMap["result_spec"] = "ResultSpec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tasksListOutputType() bindings.BindingType {
	return bindings.NewMapType(bindings.NewIdType([]string{"com.vmware.cis.task"}, ""), bindings.NewReferenceType(task.InfoBindingType), reflect.TypeOf(map[string]task.Info{}))
}

func tasksListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["filter_spec"] = bindings.NewOptionalType(bindings.NewReferenceType(TasksFilterSpecBindingType))
	fields["result_spec"] = bindings.NewOptionalType(bindings.NewReferenceType(TasksGetSpecBindingType))
	fieldNameMap["filter_spec"] = "FilterSpec"
	fieldNameMap["result_spec"] = "ResultSpec"
	bodyFieldsMap["result_spec"] = "result_spec"
	bodyFieldsMap["filter_spec"] = "filter_spec"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"] = make(map[string]string)
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"]["challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"action=list",
		"",
		"POST",
		"/cis/tasks",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_argument": 400, "com.vmware.vapi.std.errors.resource_inaccessible": 500, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.unauthenticated": 401, "com.vmware.vapi.std.errors.unauthorized": 403})
}

func tasksCancelInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["task"] = bindings.NewIdType([]string{"com.vmware.cis.task"}, "")
	fieldNameMap["task"] = "Task"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tasksCancelOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func tasksCancelRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["task"] = bindings.NewIdType([]string{"com.vmware.cis.task"}, "")
	fieldNameMap["task"] = "Task"
	paramsTypeMap["task"] = bindings.NewIdType([]string{"com.vmware.cis.task"}, "")
	paramsTypeMap["task"] = bindings.NewIdType([]string{"com.vmware.cis.task"}, "")
	pathParams["task"] = "task"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"] = make(map[string]string)
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"]["challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"action=cancel",
		"",
		"POST",
		"/cis/tasks/{task}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.error": 500, "com.vmware.vapi.std.errors.not_allowed_in_current_state": 400, "com.vmware.vapi.std.errors.not_found": 404, "com.vmware.vapi.std.errors.resource_inaccessible": 500, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.unauthenticated": 401, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.unsupported": 400})
}

func TasksGetSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["return_all"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["return_all"] = "ReturnAll"
	fields["exclude_result"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["exclude_result"] = "ExcludeResult"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.cis.tasks.get_spec", fields, reflect.TypeOf(TasksGetSpec{}), fieldNameMap, validators)
}

func TasksFilterSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tasks"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string{"com.vmware.cis.task"}, ""), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["tasks"] = "Tasks"
	fields["services"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string{"com.vmware.vapi.service"}, ""), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["services"] = "Services"
	fields["operations"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string{"com.vmware.vapi.operation"}, ""), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["operations"] = "Operations"
	fields["status"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewEnumType("com.vmware.cis.task.status", reflect.TypeOf(task.StatusEnum(task.Status_PENDING))), reflect.TypeOf(map[task.StatusEnum]bool{})))
	fieldNameMap["status"] = "Status"
	fields["targets"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(std.DynamicIDBindingType), reflect.TypeOf([]std.DynamicID{})))
	fieldNameMap["targets"] = "Targets"
	fields["users"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["users"] = "Users"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.cis.tasks.filter_spec", fields, reflect.TypeOf(TasksFilterSpec{}), fieldNameMap, validators)
}
