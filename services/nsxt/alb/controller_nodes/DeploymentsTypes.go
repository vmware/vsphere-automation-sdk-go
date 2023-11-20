// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Deployments.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package controller_nodes

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

// Possible value for ``state`` of method Deployments#list.
const Deployments_LIST_STATE_DEPLOYED = "DEPLOYED"

// Possible value for ``state`` of method Deployments#list.
const Deployments_LIST_STATE_PENDING = "PENDING"

func deploymentsCreateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["add_ALB_controller_node_VM_info"] = vapiBindings_.NewReferenceType(nsx_policyModel.AddALBControllerNodeVMInfoBindingType)
	fieldNameMap["add_ALB_controller_node_VM_info"] = "AddALBControllerNodeVMInfo"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func DeploymentsCreateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.ALBControllerNodeVMDeploymentRequestListBindingType)
}

func deploymentsCreateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["add_ALB_controller_node_VM_info"] = vapiBindings_.NewReferenceType(nsx_policyModel.AddALBControllerNodeVMInfoBindingType)
	fieldNameMap["add_ALB_controller_node_VM_info"] = "AddALBControllerNodeVMInfo"
	paramsTypeMap["add_ALB_controller_node_VM_info"] = vapiBindings_.NewReferenceType(nsx_policyModel.AddALBControllerNodeVMInfoBindingType)
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
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

func deploymentsDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["node_id"] = vapiBindings_.NewStringType()
	fields["force_delete"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["inaccessible"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["force_delete"] = "ForceDelete"
	fieldNameMap["inaccessible"] = "Inaccessible"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func DeploymentsDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func deploymentsDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["node_id"] = vapiBindings_.NewStringType()
	fields["force_delete"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["inaccessible"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["force_delete"] = "ForceDelete"
	fieldNameMap["inaccessible"] = "Inaccessible"
	paramsTypeMap["inaccessible"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["force_delete"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["node_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["nodeId"] = vapiBindings_.NewStringType()
	pathParams["node_id"] = "nodeId"
	queryParams["inaccessible"] = "inaccessible"
	queryParams["force_delete"] = "force_delete"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return vapiProtocol_.NewOperationRestMetadata(
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

func deploymentsGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["node_id"] = vapiBindings_.NewStringType()
	fieldNameMap["node_id"] = "NodeId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func DeploymentsGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.ALBControllerNodeVMDeploymentRequestBindingType)
}

func deploymentsGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["node_id"] = vapiBindings_.NewStringType()
	fieldNameMap["node_id"] = "NodeId"
	paramsTypeMap["node_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["nodeId"] = vapiBindings_.NewStringType()
	pathParams["node_id"] = "nodeId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
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
		"/policy/api/v1/alb/controller-nodes/deployments/{nodeId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func deploymentsListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["clustering_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["clustering_id"] = "ClusteringId"
	fieldNameMap["state"] = "State"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func DeploymentsListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.ALBControllerNodeVMDeploymentRequestListBindingType)
}

func deploymentsListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["clustering_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["clustering_id"] = "ClusteringId"
	fieldNameMap["state"] = "State"
	paramsTypeMap["clustering_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	queryParams["clustering_id"] = "clustering_id"
	queryParams["state"] = "state"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
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
		"/policy/api/v1/alb/controller-nodes/deployments",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func deploymentsUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["node_id"] = vapiBindings_.NewStringType()
	fields["a_LB_controller_node_VM_deployment_request"] = vapiBindings_.NewReferenceType(nsx_policyModel.ALBControllerNodeVMDeploymentRequestBindingType)
	fields["running_config"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["a_LB_controller_node_VM_deployment_request"] = "ALBControllerNodeVMDeploymentRequest"
	fieldNameMap["running_config"] = "RunningConfig"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func DeploymentsUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.ALBControllerNodeVMDeploymentRequestBindingType)
}

func deploymentsUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["node_id"] = vapiBindings_.NewStringType()
	fields["a_LB_controller_node_VM_deployment_request"] = vapiBindings_.NewReferenceType(nsx_policyModel.ALBControllerNodeVMDeploymentRequestBindingType)
	fields["running_config"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["a_LB_controller_node_VM_deployment_request"] = "ALBControllerNodeVMDeploymentRequest"
	fieldNameMap["running_config"] = "RunningConfig"
	paramsTypeMap["a_LB_controller_node_VM_deployment_request"] = vapiBindings_.NewReferenceType(nsx_policyModel.ALBControllerNodeVMDeploymentRequestBindingType)
	paramsTypeMap["running_config"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["node_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["nodeId"] = vapiBindings_.NewStringType()
	pathParams["node_id"] = "nodeId"
	queryParams["running_config"] = "running_config"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
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
		"a_LB_controller_node_VM_deployment_request",
		"PUT",
		"/policy/api/v1/alb/controller-nodes/deployments/{nodeId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
