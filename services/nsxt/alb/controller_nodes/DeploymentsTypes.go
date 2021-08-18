// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Deployments.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package controller_nodes

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

// Possible value for ``state`` of method Deployments#list.
const Deployments_LIST_STATE_DEPLOYED = "DEPLOYED"

// Possible value for ``state`` of method Deployments#list.
const Deployments_LIST_STATE_PENDING = "PENDING"

func deploymentsCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["add_ALB_controller_node_VM_info"] = bindings.NewReferenceType(model.AddALBControllerNodeVMInfoBindingType)
	fieldNameMap["add_ALB_controller_node_VM_info"] = "AddALBControllerNodeVMInfo"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func deploymentsCreateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ALBControllerNodeVMDeploymentRequestListBindingType)
}

func deploymentsCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["add_ALB_controller_node_VM_info"] = bindings.NewReferenceType(model.AddALBControllerNodeVMInfoBindingType)
	fieldNameMap["add_ALB_controller_node_VM_info"] = "AddALBControllerNodeVMInfo"
	paramsTypeMap["add_ALB_controller_node_VM_info"] = bindings.NewReferenceType(model.AddALBControllerNodeVMInfoBindingType)
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
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
		"add_ALB_controller_node_VM_info",
		"POST",
		"/policy/api/v1/alb/controller-nodes/deployments",
		"",
		resultHeaders,
		201,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func deploymentsDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["node_id"] = bindings.NewStringType()
	fields["force_delete"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["force_delete"] = "ForceDelete"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func deploymentsDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func deploymentsDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["node_id"] = bindings.NewStringType()
	fields["force_delete"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["force_delete"] = "ForceDelete"
	paramsTypeMap["force_delete"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["node_id"] = bindings.NewStringType()
	paramsTypeMap["nodeId"] = bindings.NewStringType()
	pathParams["node_id"] = "nodeId"
	queryParams["force_delete"] = "force_delete"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"action=delete",
		"",
		"POST",
		"/policy/api/v1/alb/controller-nodes/deployments/{nodeId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func deploymentsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["node_id"] = bindings.NewStringType()
	fieldNameMap["node_id"] = "NodeId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func deploymentsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ALBControllerNodeVMDeploymentRequestBindingType)
}

func deploymentsGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["node_id"] = bindings.NewStringType()
	fieldNameMap["node_id"] = "NodeId"
	paramsTypeMap["node_id"] = bindings.NewStringType()
	paramsTypeMap["nodeId"] = bindings.NewStringType()
	pathParams["node_id"] = "nodeId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
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
		"/policy/api/v1/alb/controller-nodes/deployments/{nodeId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func deploymentsListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["state"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["state"] = "State"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func deploymentsListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ALBControllerNodeVMDeploymentRequestListBindingType)
}

func deploymentsListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["state"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["state"] = "State"
	paramsTypeMap["state"] = bindings.NewOptionalType(bindings.NewStringType())
	queryParams["state"] = "state"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
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
		"/policy/api/v1/alb/controller-nodes/deployments",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
