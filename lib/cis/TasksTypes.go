// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Tasks.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package cis

import (
	cisTask_ "github.com/vmware/vsphere-automation-sdk-go/lib/cis/task"
	vapiStd_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiLog_ "github.com/vmware/vsphere-automation-sdk-go/runtime/log"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"reflect"
)

// Resource type for task.
const Tasks_RESOURCE_TYPE = "com.vmware.cis.task"

// The ``GetSpec`` class describes what data should be included when retrieving information about a task.
type TasksGetSpec struct {
	// If true, all data, including operation-specific data, will be returned, otherwise only the data described in cisTask_.Info will be returned.
	ReturnAll *bool
	// If true, the result will not be included in the task information, otherwise it will be included.
	ExcludeResult *bool
}

func (s *TasksGetSpec) GetType__() vapiBindings_.BindingType {
	return TasksGetSpecBindingType()
}

func (s *TasksGetSpec) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for TasksGetSpec._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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
	// Identifiers of services. Tasks created by operations in these services match the filter (see cisTask_.CommonInfo#service).
	Services map[string]bool
	// Identifiers of operations. Tasks created by these operations match the filter (see cisTask_.CommonInfo#operation).
	//
	//  Note that an operation identifier by itself is not globally unique. To filter on an operation, the identifier of the service interface containing the operation should also be specified in ``services``.
	Operations map[string]bool
	// Status that a task must have to match the filter (see cisTask_.CommonInfo#status).
	Status map[cisTask_.StatusEnum]bool
	// Identifiers of the targets the operation for the associated task created or was performed on (see cisTask_.CommonInfo#target).
	Targets []vapiStd_.DynamicID
	// Users who must have initiated the operation for the associated task to match the filter (see cisTask_.CommonInfo#user).
	Users map[string]bool
}

func (s *TasksFilterSpec) GetType__() vapiBindings_.BindingType {
	return TasksFilterSpecBindingType()
}

func (s *TasksFilterSpec) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for TasksFilterSpec._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

func tasksGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["task"] = vapiBindings_.NewIdType([]string{"com.vmware.cis.task"}, "")
	fields["spec"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(TasksGetSpecBindingType))
	fieldNameMap["task"] = "Task"
	fieldNameMap["spec"] = "Spec"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func TasksGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(cisTask_.InfoBindingType)
}

func tasksGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["task"] = vapiBindings_.NewIdType([]string{"com.vmware.cis.task"}, "")
	fields["spec"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(TasksGetSpecBindingType))
	fieldNameMap["task"] = "Task"
	fieldNameMap["spec"] = "Spec"
	paramsTypeMap["task"] = vapiBindings_.NewIdType([]string{"com.vmware.cis.task"}, "")
	paramsTypeMap["spec.exclude_result"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["spec.return_all"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["task"] = vapiBindings_.NewIdType([]string{"com.vmware.cis.task"}, "")
	pathParams["task"] = "task"
	queryParams["spec.exclude_result"] = "exclude_result"
	queryParams["spec.return_all"] = "return_all"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"] = make(map[string]string)
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"]["challenge"] = "WWW-Authenticate"
	return vapiProtocol_.NewOperationRestMetadata(
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

func tasksListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["filter_spec"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(TasksFilterSpecBindingType))
	fields["result_spec"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(TasksGetSpecBindingType))
	fieldNameMap["filter_spec"] = "FilterSpec"
	fieldNameMap["result_spec"] = "ResultSpec"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func TasksListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewMapType(vapiBindings_.NewIdType([]string{"com.vmware.cis.task"}, ""), vapiBindings_.NewReferenceType(cisTask_.InfoBindingType), reflect.TypeOf(map[string]cisTask_.Info{}))
}

func tasksListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["filter_spec"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(TasksFilterSpecBindingType))
	fields["result_spec"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(TasksGetSpecBindingType))
	fieldNameMap["filter_spec"] = "FilterSpec"
	fieldNameMap["result_spec"] = "ResultSpec"
	bodyFieldsMap["result_spec"] = "result_spec"
	bodyFieldsMap["filter_spec"] = "filter_spec"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"] = make(map[string]string)
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"]["challenge"] = "WWW-Authenticate"
	return vapiProtocol_.NewOperationRestMetadata(
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

func tasksCancelInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["task"] = vapiBindings_.NewIdType([]string{"com.vmware.cis.task"}, "")
	fieldNameMap["task"] = "Task"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func TasksCancelOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func tasksCancelRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["task"] = vapiBindings_.NewIdType([]string{"com.vmware.cis.task"}, "")
	fieldNameMap["task"] = "Task"
	paramsTypeMap["task"] = vapiBindings_.NewIdType([]string{"com.vmware.cis.task"}, "")
	paramsTypeMap["task"] = vapiBindings_.NewIdType([]string{"com.vmware.cis.task"}, "")
	pathParams["task"] = "task"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"] = make(map[string]string)
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"]["challenge"] = "WWW-Authenticate"
	return vapiProtocol_.NewOperationRestMetadata(
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

func TasksGetSpecBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["return_all"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["return_all"] = "ReturnAll"
	fields["exclude_result"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["exclude_result"] = "ExcludeResult"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.cis.tasks.get_spec", fields, reflect.TypeOf(TasksGetSpec{}), fieldNameMap, validators)
}

func TasksFilterSpecBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tasks"] = vapiBindings_.NewOptionalType(vapiBindings_.NewSetType(vapiBindings_.NewIdType([]string{"com.vmware.cis.task"}, ""), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["tasks"] = "Tasks"
	fields["services"] = vapiBindings_.NewOptionalType(vapiBindings_.NewSetType(vapiBindings_.NewIdType([]string{"com.vmware.vapi.service"}, ""), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["services"] = "Services"
	fields["operations"] = vapiBindings_.NewOptionalType(vapiBindings_.NewSetType(vapiBindings_.NewIdType([]string{"com.vmware.vapi.operation"}, ""), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["operations"] = "Operations"
	fields["status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewSetType(vapiBindings_.NewEnumType("com.vmware.cis.task.status", reflect.TypeOf(cisTask_.StatusEnum(cisTask_.Status_PENDING))), reflect.TypeOf(map[cisTask_.StatusEnum]bool{})))
	fieldNameMap["status"] = "Status"
	fields["targets"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(vapiStd_.DynamicIDBindingType), reflect.TypeOf([]vapiStd_.DynamicID{})))
	fieldNameMap["targets"] = "Targets"
	fields["users"] = vapiBindings_.NewOptionalType(vapiBindings_.NewSetType(vapiBindings_.NewStringType(), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["users"] = "Users"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.cis.tasks.filter_spec", fields, reflect.TypeOf(TasksFilterSpec{}), fieldNameMap, validators)
}
