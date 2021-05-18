// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: SystemAdministrationConfigurationFabricNodesFabricNodes.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package api

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-mp/model"
	"reflect"
)

// Possible value for ``hypervisorOsType`` of method SystemAdministrationConfigurationFabricNodesFabricNodes#listNodes.
const SystemAdministrationConfigurationFabricNodesFabricNodes_LIST_NODES_HYPERVISOR_OS_TYPE_ESXI = "ESXI"

// Possible value for ``hypervisorOsType`` of method SystemAdministrationConfigurationFabricNodesFabricNodes#listNodes.
const SystemAdministrationConfigurationFabricNodesFabricNodes_LIST_NODES_HYPERVISOR_OS_TYPE_RHELKVM = "RHELKVM"

// Possible value for ``hypervisorOsType`` of method SystemAdministrationConfigurationFabricNodesFabricNodes#listNodes.
const SystemAdministrationConfigurationFabricNodesFabricNodes_LIST_NODES_HYPERVISOR_OS_TYPE_UBUNTUKVM = "UBUNTUKVM"

// Possible value for ``hypervisorOsType`` of method SystemAdministrationConfigurationFabricNodesFabricNodes#listNodes.
const SystemAdministrationConfigurationFabricNodesFabricNodes_LIST_NODES_HYPERVISOR_OS_TYPE_HYPERV = "HYPERV"

// Possible value for ``hypervisorOsType`` of method SystemAdministrationConfigurationFabricNodesFabricNodes#listNodes.
const SystemAdministrationConfigurationFabricNodesFabricNodes_LIST_NODES_HYPERVISOR_OS_TYPE_RHELCONTAINER = "RHELCONTAINER"

// Possible value for ``hypervisorOsType`` of method SystemAdministrationConfigurationFabricNodesFabricNodes#listNodes.
const SystemAdministrationConfigurationFabricNodesFabricNodes_LIST_NODES_HYPERVISOR_OS_TYPE_CENTOSCONTAINER = "CENTOSCONTAINER"

// Possible value for ``hypervisorOsType`` of method SystemAdministrationConfigurationFabricNodesFabricNodes#listNodes.
const SystemAdministrationConfigurationFabricNodesFabricNodes_LIST_NODES_HYPERVISOR_OS_TYPE_RHELSERVER = "RHELSERVER"

// Possible value for ``hypervisorOsType`` of method SystemAdministrationConfigurationFabricNodesFabricNodes#listNodes.
const SystemAdministrationConfigurationFabricNodesFabricNodes_LIST_NODES_HYPERVISOR_OS_TYPE_WINDOWSSERVER = "WINDOWSSERVER"

// Possible value for ``hypervisorOsType`` of method SystemAdministrationConfigurationFabricNodesFabricNodes#listNodes.
const SystemAdministrationConfigurationFabricNodesFabricNodes_LIST_NODES_HYPERVISOR_OS_TYPE_UBUNTUSERVER = "UBUNTUSERVER"

// Possible value for ``hypervisorOsType`` of method SystemAdministrationConfigurationFabricNodesFabricNodes#listNodes.
const SystemAdministrationConfigurationFabricNodesFabricNodes_LIST_NODES_HYPERVISOR_OS_TYPE_CENTOSSERVER = "CENTOSSERVER"

// Possible value for ``hypervisorOsType`` of method SystemAdministrationConfigurationFabricNodesFabricNodes#listNodes.
const SystemAdministrationConfigurationFabricNodesFabricNodes_LIST_NODES_HYPERVISOR_OS_TYPE_CENTOSKVM = "CENTOSKVM"

// Possible value for ``hypervisorOsType`` of method SystemAdministrationConfigurationFabricNodesFabricNodes#listNodes.
const SystemAdministrationConfigurationFabricNodesFabricNodes_LIST_NODES_HYPERVISOR_OS_TYPE_SLESKVM = "SLESKVM"

// Possible value for ``hypervisorOsType`` of method SystemAdministrationConfigurationFabricNodesFabricNodes#listNodes.
const SystemAdministrationConfigurationFabricNodesFabricNodes_LIST_NODES_HYPERVISOR_OS_TYPE_SLESSERVER = "SLESSERVER"

// Possible value for ``hypervisorOsType`` of method SystemAdministrationConfigurationFabricNodesFabricNodes#listNodes.
const SystemAdministrationConfigurationFabricNodesFabricNodes_LIST_NODES_HYPERVISOR_OS_TYPE_OELSERVER = "OELSERVER"

// Possible value for ``hypervisorOsType`` of method SystemAdministrationConfigurationFabricNodesFabricNodes#listNodes.
const SystemAdministrationConfigurationFabricNodesFabricNodes_LIST_NODES_HYPERVISOR_OS_TYPE_RHELSMARTNIC = "RHELSMARTNIC"

// Possible value for ``hypervisorOsType`` of method SystemAdministrationConfigurationFabricNodesFabricNodes#listNodes.
const SystemAdministrationConfigurationFabricNodesFabricNodes_LIST_NODES_HYPERVISOR_OS_TYPE_UBUNTUSMARTNIC = "UBUNTUSMARTNIC"

// Possible value for ``resourceType`` of method SystemAdministrationConfigurationFabricNodesFabricNodes#listNodes.
const SystemAdministrationConfigurationFabricNodesFabricNodes_LIST_NODES_RESOURCE_TYPE_HOSTNODE = "HostNode"

// Possible value for ``resourceType`` of method SystemAdministrationConfigurationFabricNodesFabricNodes#listNodes.
const SystemAdministrationConfigurationFabricNodesFabricNodes_LIST_NODES_RESOURCE_TYPE_EDGENODE = "EdgeNode"

// Possible value for ``resourceType`` of method SystemAdministrationConfigurationFabricNodesFabricNodes#listNodes.
const SystemAdministrationConfigurationFabricNodesFabricNodes_LIST_NODES_RESOURCE_TYPE_PUBLICCLOUDGATEWAYNODE = "PublicCloudGatewayNode"

// Possible value for ``action`` of method SystemAdministrationConfigurationFabricNodesFabricNodes#performNodeAction.
const SystemAdministrationConfigurationFabricNodesFabricNodes_PERFORM_NODE_ACTION_ACTION_ENTER_MAINTENANCE_MODE = "enter_maintenance_mode"

// Possible value for ``action`` of method SystemAdministrationConfigurationFabricNodesFabricNodes#performNodeAction.
const SystemAdministrationConfigurationFabricNodesFabricNodes_PERFORM_NODE_ACTION_ACTION_EXIT_MAINTENANCE_MODE = "exit_maintenance_mode"

// Possible value for ``action`` of method SystemAdministrationConfigurationFabricNodesFabricNodes#performNodeAction.
const SystemAdministrationConfigurationFabricNodesFabricNodes_PERFORM_NODE_ACTION_ACTION_GET_MAINTENANCE_MODE_STATE = "get_maintenance_mode_state"

// Possible value for ``vsanMode`` of method SystemAdministrationConfigurationFabricNodesFabricNodes#performNodeAction.
const SystemAdministrationConfigurationFabricNodesFabricNodes_PERFORM_NODE_ACTION_VSAN_MODE_EVACUATE_ALL_DATA = "evacuate_all_data"

// Possible value for ``vsanMode`` of method SystemAdministrationConfigurationFabricNodesFabricNodes#performNodeAction.
const SystemAdministrationConfigurationFabricNodesFabricNodes_PERFORM_NODE_ACTION_VSAN_MODE_ENSURE_OBJECT_ACCESSIBILITY = "ensure_object_accessibility"

// Possible value for ``vsanMode`` of method SystemAdministrationConfigurationFabricNodesFabricNodes#performNodeAction.
const SystemAdministrationConfigurationFabricNodesFabricNodes_PERFORM_NODE_ACTION_VSAN_MODE_NO_ACTION = "no_action"

// Possible value for ``source`` of method SystemAdministrationConfigurationFabricNodesFabricNodes#readNodeStatus.
const SystemAdministrationConfigurationFabricNodesFabricNodes_READ_NODE_STATUS_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method SystemAdministrationConfigurationFabricNodesFabricNodes#readNodeStatus.
const SystemAdministrationConfigurationFabricNodesFabricNodes_READ_NODE_STATUS_SOURCE_CACHED = "cached"

func systemAdministrationConfigurationFabricNodesFabricNodesAddNodeInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["node"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.NodeBindingType)}, bindings.REST)
	fieldNameMap["node"] = "Node"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricNodesFabricNodesAddNodeOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.NodeBindingType)}, bindings.REST)
}

func systemAdministrationConfigurationFabricNodesFabricNodesAddNodeRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["node"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.NodeBindingType)}, bindings.REST)
	fieldNameMap["node"] = "Node"
	paramsTypeMap["node"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.NodeBindingType)}, bindings.REST)
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
		"node",
		"POST",
		"/api/v1/fabric/nodes",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricNodesFabricNodesDeleteNodeInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["node_id"] = bindings.NewStringType()
	fields["unprepare_host"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["unprepare_host"] = "UnprepareHost"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricNodesFabricNodesDeleteNodeOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func systemAdministrationConfigurationFabricNodesFabricNodesDeleteNodeRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["node_id"] = bindings.NewStringType()
	fields["unprepare_host"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["unprepare_host"] = "UnprepareHost"
	paramsTypeMap["node_id"] = bindings.NewStringType()
	paramsTypeMap["unprepare_host"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["nodeId"] = bindings.NewStringType()
	pathParams["node_id"] = "nodeId"
	queryParams["unprepare_host"] = "unprepare_host"
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
		"DELETE",
		"/api/v1/fabric/nodes/{nodeId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricNodesFabricNodesGetFabricNodeModulesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["node_id"] = bindings.NewStringType()
	fieldNameMap["node_id"] = "NodeId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricNodesFabricNodesGetFabricNodeModulesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.SoftwareModuleResultBindingType)
}

func systemAdministrationConfigurationFabricNodesFabricNodesGetFabricNodeModulesRestMetadata() protocol.OperationRestMetadata {
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
		"/api/v1/fabric/nodes/{nodeId}/modules",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricNodesFabricNodesGetFabricNodeStateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["node_id"] = bindings.NewStringType()
	fieldNameMap["node_id"] = "NodeId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricNodesFabricNodesGetFabricNodeStateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ConfigurationStateBindingType)
}

func systemAdministrationConfigurationFabricNodesFabricNodesGetFabricNodeStateRestMetadata() protocol.OperationRestMetadata {
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
		"/api/v1/fabric/nodes/{nodeId}/state",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricNodesFabricNodesGetSupportedHostOSTypesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricNodesFabricNodesGetSupportedHostOSTypesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.SupportedHostOSListResultBindingType)
}

func systemAdministrationConfigurationFabricNodesFabricNodesGetSupportedHostOSTypesRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
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
		"/api/v1/fabric/ostypes",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricNodesFabricNodesInvokeDeleteFabricCentralAPIInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["target_node_id"] = bindings.NewStringType()
	fields["target_uri"] = bindings.NewStringType()
	fieldNameMap["target_node_id"] = "TargetNodeId"
	fieldNameMap["target_uri"] = "TargetUri"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricNodesFabricNodesInvokeDeleteFabricCentralAPIOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func systemAdministrationConfigurationFabricNodesFabricNodesInvokeDeleteFabricCentralAPIRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["target_node_id"] = bindings.NewStringType()
	fields["target_uri"] = bindings.NewStringType()
	fieldNameMap["target_node_id"] = "TargetNodeId"
	fieldNameMap["target_uri"] = "TargetUri"
	paramsTypeMap["target_uri"] = bindings.NewStringType()
	paramsTypeMap["target_node_id"] = bindings.NewStringType()
	paramsTypeMap["targetNodeId"] = bindings.NewStringType()
	paramsTypeMap["targetUri"] = bindings.NewStringType()
	pathParams["target_uri"] = "targetUri"
	pathParams["target_node_id"] = "targetNodeId"
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
		"DELETE",
		"/api/v1/fabric/nodes/{targetNodeId}/{targetUri}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricNodesFabricNodesInvokeGetFabricCentralAPIInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["target_node_id"] = bindings.NewStringType()
	fields["target_uri"] = bindings.NewStringType()
	fieldNameMap["target_node_id"] = "TargetNodeId"
	fieldNameMap["target_uri"] = "TargetUri"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricNodesFabricNodesInvokeGetFabricCentralAPIOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func systemAdministrationConfigurationFabricNodesFabricNodesInvokeGetFabricCentralAPIRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["target_node_id"] = bindings.NewStringType()
	fields["target_uri"] = bindings.NewStringType()
	fieldNameMap["target_node_id"] = "TargetNodeId"
	fieldNameMap["target_uri"] = "TargetUri"
	paramsTypeMap["target_uri"] = bindings.NewStringType()
	paramsTypeMap["target_node_id"] = bindings.NewStringType()
	paramsTypeMap["targetNodeId"] = bindings.NewStringType()
	paramsTypeMap["targetUri"] = bindings.NewStringType()
	pathParams["target_uri"] = "targetUri"
	pathParams["target_node_id"] = "targetNodeId"
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
		"/api/v1/fabric/nodes/{targetNodeId}/{targetUri}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricNodesFabricNodesInvokePostFabricCentralAPIInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["target_node_id"] = bindings.NewStringType()
	fields["target_uri"] = bindings.NewStringType()
	fieldNameMap["target_node_id"] = "TargetNodeId"
	fieldNameMap["target_uri"] = "TargetUri"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricNodesFabricNodesInvokePostFabricCentralAPIOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func systemAdministrationConfigurationFabricNodesFabricNodesInvokePostFabricCentralAPIRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["target_node_id"] = bindings.NewStringType()
	fields["target_uri"] = bindings.NewStringType()
	fieldNameMap["target_node_id"] = "TargetNodeId"
	fieldNameMap["target_uri"] = "TargetUri"
	paramsTypeMap["target_uri"] = bindings.NewStringType()
	paramsTypeMap["target_node_id"] = bindings.NewStringType()
	paramsTypeMap["targetNodeId"] = bindings.NewStringType()
	paramsTypeMap["targetUri"] = bindings.NewStringType()
	pathParams["target_uri"] = "targetUri"
	pathParams["target_node_id"] = "targetNodeId"
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
		"POST",
		"/api/v1/fabric/nodes/{targetNodeId}/{targetUri}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricNodesFabricNodesInvokePutFabricCentralAPIInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["target_node_id"] = bindings.NewStringType()
	fields["target_uri"] = bindings.NewStringType()
	fieldNameMap["target_node_id"] = "TargetNodeId"
	fieldNameMap["target_uri"] = "TargetUri"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricNodesFabricNodesInvokePutFabricCentralAPIOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func systemAdministrationConfigurationFabricNodesFabricNodesInvokePutFabricCentralAPIRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["target_node_id"] = bindings.NewStringType()
	fields["target_uri"] = bindings.NewStringType()
	fieldNameMap["target_node_id"] = "TargetNodeId"
	fieldNameMap["target_uri"] = "TargetUri"
	paramsTypeMap["target_uri"] = bindings.NewStringType()
	paramsTypeMap["target_node_id"] = bindings.NewStringType()
	paramsTypeMap["targetNodeId"] = bindings.NewStringType()
	paramsTypeMap["targetUri"] = bindings.NewStringType()
	pathParams["target_uri"] = "targetUri"
	pathParams["target_node_id"] = "targetNodeId"
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
		"PUT",
		"/api/v1/fabric/nodes/{targetNodeId}/{targetUri}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricNodesFabricNodesListNodeCapabilitiesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["node_id"] = bindings.NewStringType()
	fieldNameMap["node_id"] = "NodeId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricNodesFabricNodesListNodeCapabilitiesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NodeCapabilitiesResultBindingType)
}

func systemAdministrationConfigurationFabricNodesFabricNodesListNodeCapabilitiesRestMetadata() protocol.OperationRestMetadata {
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
		"/api/v1/fabric/nodes/{nodeId}/capabilities",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricNodesFabricNodesListNodesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["discovered_node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["external_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["hardware_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["hypervisor_os_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["ip_address"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["discovered_node_id"] = "DiscoveredNodeId"
	fieldNameMap["display_name"] = "DisplayName"
	fieldNameMap["external_id"] = "ExternalId"
	fieldNameMap["hardware_id"] = "HardwareId"
	fieldNameMap["hypervisor_os_type"] = "HypervisorOsType"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["ip_address"] = "IpAddress"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["resource_type"] = "ResourceType"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricNodesFabricNodesListNodesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NodeListResultBindingType)
}

func systemAdministrationConfigurationFabricNodesFabricNodesListNodesRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["discovered_node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["external_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["hardware_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["hypervisor_os_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["ip_address"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["discovered_node_id"] = "DiscoveredNodeId"
	fieldNameMap["display_name"] = "DisplayName"
	fieldNameMap["external_id"] = "ExternalId"
	fieldNameMap["hardware_id"] = "HardwareId"
	fieldNameMap["hypervisor_os_type"] = "HypervisorOsType"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["ip_address"] = "IpAddress"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["resource_type"] = "ResourceType"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["external_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["discovered_node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["ip_address"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["hypervisor_os_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["hardware_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["discovered_node_id"] = "discovered_node_id"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["hypervisor_os_type"] = "hypervisor_os_type"
	queryParams["resource_type"] = "resource_type"
	queryParams["external_id"] = "external_id"
	queryParams["ip_address"] = "ip_address"
	queryParams["sort_by"] = "sort_by"
	queryParams["display_name"] = "display_name"
	queryParams["hardware_id"] = "hardware_id"
	queryParams["page_size"] = "page_size"
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
		"/api/v1/fabric/nodes",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricNodesFabricNodesPerformHostNodeUpgradeActionUpgradeInfraInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["node_id"] = bindings.NewStringType()
	fields["disable_vm_migration"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["disable_vm_migration"] = "DisableVmMigration"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricNodesFabricNodesPerformHostNodeUpgradeActionUpgradeInfraOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.NodeBindingType)}, bindings.REST)
}

func systemAdministrationConfigurationFabricNodesFabricNodesPerformHostNodeUpgradeActionUpgradeInfraRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["node_id"] = bindings.NewStringType()
	fields["disable_vm_migration"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["disable_vm_migration"] = "DisableVmMigration"
	paramsTypeMap["disable_vm_migration"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["node_id"] = bindings.NewStringType()
	paramsTypeMap["nodeId"] = bindings.NewStringType()
	pathParams["node_id"] = "nodeId"
	queryParams["disable_vm_migration"] = "disable_vm_migration"
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
		"action=upgrade_infra",
		"",
		"POST",
		"/api/v1/fabric/nodes/{nodeId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricNodesFabricNodesPerformNodeActionInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["node_id"] = bindings.NewStringType()
	fields["action"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["evacuate_powered_off_vms"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["vsan_mode"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["action"] = "Action"
	fieldNameMap["evacuate_powered_off_vms"] = "EvacuatePoweredOffVms"
	fieldNameMap["vsan_mode"] = "VsanMode"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricNodesFabricNodesPerformNodeActionOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.NodeBindingType)}, bindings.REST)
}

func systemAdministrationConfigurationFabricNodesFabricNodesPerformNodeActionRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["node_id"] = bindings.NewStringType()
	fields["action"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["evacuate_powered_off_vms"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["vsan_mode"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["action"] = "Action"
	fieldNameMap["evacuate_powered_off_vms"] = "EvacuatePoweredOffVms"
	fieldNameMap["vsan_mode"] = "VsanMode"
	paramsTypeMap["action"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["evacuate_powered_off_vms"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["vsan_mode"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["node_id"] = bindings.NewStringType()
	paramsTypeMap["nodeId"] = bindings.NewStringType()
	pathParams["node_id"] = "nodeId"
	queryParams["evacuate_powered_off_vms"] = "evacuate_powered_off_vms"
	queryParams["action"] = "action"
	queryParams["vsan_mode"] = "vsan_mode"
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
		"POST",
		"/api/v1/fabric/nodes/{nodeId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricNodesFabricNodesReadNodeInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["node_id"] = bindings.NewStringType()
	fieldNameMap["node_id"] = "NodeId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricNodesFabricNodesReadNodeOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.NodeBindingType)}, bindings.REST)
}

func systemAdministrationConfigurationFabricNodesFabricNodesReadNodeRestMetadata() protocol.OperationRestMetadata {
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
		"/api/v1/fabric/nodes/{nodeId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricNodesFabricNodesReadNodeStatusInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["node_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["source"] = "Source"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricNodesFabricNodesReadNodeStatusOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NodeStatusBindingType)
}

func systemAdministrationConfigurationFabricNodesFabricNodesReadNodeStatusRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["node_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["source"] = "Source"
	paramsTypeMap["source"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["node_id"] = bindings.NewStringType()
	paramsTypeMap["nodeId"] = bindings.NewStringType()
	pathParams["node_id"] = "nodeId"
	queryParams["source"] = "source"
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
		"/api/v1/fabric/nodes/{nodeId}/status",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricNodesFabricNodesReadNodesStatusInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["node_ids"] = bindings.NewStringType()
	fieldNameMap["node_ids"] = "NodeIds"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricNodesFabricNodesReadNodesStatusOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NodeStatusListResultBindingType)
}

func systemAdministrationConfigurationFabricNodesFabricNodesReadNodesStatusRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["node_ids"] = bindings.NewStringType()
	fieldNameMap["node_ids"] = "NodeIds"
	paramsTypeMap["node_ids"] = bindings.NewStringType()
	queryParams["node_ids"] = "node_ids"
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
		"/api/v1/fabric/nodes/status",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricNodesFabricNodesRestartInventorySyncRestartInventorySyncInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["node_id"] = bindings.NewStringType()
	fieldNameMap["node_id"] = "NodeId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricNodesFabricNodesRestartInventorySyncRestartInventorySyncOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func systemAdministrationConfigurationFabricNodesFabricNodesRestartInventorySyncRestartInventorySyncRestMetadata() protocol.OperationRestMetadata {
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
		"action=restart_inventory_sync",
		"",
		"POST",
		"/api/v1/fabric/nodes/{nodeId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricNodesFabricNodesUpdateNodeInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["node_id"] = bindings.NewStringType()
	fields["node"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.NodeBindingType)}, bindings.REST)
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["node"] = "Node"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricNodesFabricNodesUpdateNodeOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.NodeBindingType)}, bindings.REST)
}

func systemAdministrationConfigurationFabricNodesFabricNodesUpdateNodeRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["node_id"] = bindings.NewStringType()
	fields["node"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.NodeBindingType)}, bindings.REST)
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["node"] = "Node"
	paramsTypeMap["node"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.NodeBindingType)}, bindings.REST)
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
		"node",
		"PUT",
		"/api/v1/fabric/nodes/{nodeId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
