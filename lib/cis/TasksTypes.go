/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Tasks.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package cis

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/cis/task"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
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
	Status map[task.Status]bool
    // Identifiers of the targets the operation for the associated task created or was performed on (see task.CommonInfo#target).
	Targets []std.DynamicID
    // Users who must have initiated the operation for the associated task to match the filter (see task.CommonInfo#user).
	Users map[string]bool
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
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
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
		"null",
		"",
		"",
		resultHeaders,
		0,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.error": 500,"com.vmware.vapi.std.errors.not_found": 404,"com.vmware.vapi.std.errors.resource_inaccessible": 500,"com.vmware.vapi.std.errors.service_unavailable": 503,"com.vmware.vapi.std.errors.unauthenticated": 401,"com.vmware.vapi.std.errors.unauthorized": 403})
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
	return bindings.NewMapType(bindings.NewIdType([]string{"com.vmware.cis.task"}, ""), bindings.NewReferenceType(task.InfoBindingType),reflect.TypeOf(map[string]task.Info{}))
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
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
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
		"null",
		"",
		"",
		resultHeaders,
		0,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_argument": 400,"com.vmware.vapi.std.errors.resource_inaccessible": 500,"com.vmware.vapi.std.errors.service_unavailable": 503,"com.vmware.vapi.std.errors.unauthenticated": 401,"com.vmware.vapi.std.errors.unauthorized": 403})
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
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
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
		"null",
		"",
		"",
		resultHeaders,
		0,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.error": 500,"com.vmware.vapi.std.errors.not_allowed_in_current_state": 400,"com.vmware.vapi.std.errors.not_found": 404,"com.vmware.vapi.std.errors.resource_inaccessible": 500,"com.vmware.vapi.std.errors.service_unavailable": 503,"com.vmware.vapi.std.errors.unauthenticated": 401,"com.vmware.vapi.std.errors.unauthorized": 403,"com.vmware.vapi.std.errors.unsupported": 400})
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
	fields["status"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewEnumType("com.vmware.cis.task.status", reflect.TypeOf(task.Status(task.Status_PENDING))), reflect.TypeOf(map[task.Status]bool{})))
	fieldNameMap["status"] = "Status"
	fields["targets"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(std.DynamicIDBindingType), reflect.TypeOf([]std.DynamicID{})))
	fieldNameMap["targets"] = "Targets"
	fields["users"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["users"] = "Users"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.cis.tasks.filter_spec", fields, reflect.TypeOf(TasksFilterSpec{}), fieldNameMap, validators)
}

