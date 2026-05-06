// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: VcWorkflowTimestampList.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package workflows

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

// Possible value for ``workflowType`` of method VcWorkflowTimestampList#get.
const VcWorkflowTimestampList_GET_WORKFLOW_TYPE_COLLECT_CONFIG = "COLLECT_CONFIG"

// Possible value for ``workflowType`` of method VcWorkflowTimestampList#get.
const VcWorkflowTimestampList_GET_WORKFLOW_TYPE_ASSESSMENT = "ASSESSMENT"

// Possible value for ``workflowType`` of method VcWorkflowTimestampList#get.
const VcWorkflowTimestampList_GET_WORKFLOW_TYPE_EXTEND_DVPG = "EXTEND_DVPG"

// Possible value for ``workflowType`` of method VcWorkflowTimestampList#get.
const VcWorkflowTimestampList_GET_WORKFLOW_TYPE_MIGRATE_VMS_TO_OVERLAY = "MIGRATE_VMS_TO_OVERLAY"

// Possible value for ``workflowType`` of method VcWorkflowTimestampList#get.
const VcWorkflowTimestampList_GET_WORKFLOW_TYPE_MIGRATE_VMS_TO_DVPG = "MIGRATE_VMS_TO_DVPG"

// Possible value for ``workflowType`` of method VcWorkflowTimestampList#get.
const VcWorkflowTimestampList_GET_WORKFLOW_TYPE_MIGRATE_GATEWAY = "MIGRATE_GATEWAY"

// Possible value for ``workflowType`` of method VcWorkflowTimestampList#get.
const VcWorkflowTimestampList_GET_WORKFLOW_TYPE_DISCONNECT_BRIDGE = "DISCONNECT_BRIDGE"

func vcWorkflowTimestampListGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["workflow_type"] = vapiBindings_.NewStringType()
	fields["vc_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["workflow_type"] = "WorkflowType"
	fieldNameMap["vc_id"] = "VcId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func VcWorkflowTimestampListGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.VcWorkflowTimeStampListBindingType)
}

func vcWorkflowTimestampListGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["workflow_type"] = vapiBindings_.NewStringType()
	fields["vc_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["workflow_type"] = "WorkflowType"
	fieldNameMap["vc_id"] = "VcId"
	paramsTypeMap["vc_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["workflow_type"] = vapiBindings_.NewStringType()
	queryParams["vc_id"] = "vc_id"
	queryParams["workflow_type"] = "workflow_type"
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
		"/api/v1/migration/overlay-adoption/workflows/vc-workflow-timestamp-list",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
