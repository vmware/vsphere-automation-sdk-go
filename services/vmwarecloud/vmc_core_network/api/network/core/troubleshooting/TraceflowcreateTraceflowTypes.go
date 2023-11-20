// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: TraceflowcreateTraceflow.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package troubleshooting

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	vmwarecloudVmc_core_networkModel "github.com/vmware/vsphere-automation-sdk-go/services/vmwarecloud/vmc_core_network/model"
	"reflect"
)

func traceflowcreateTraceflowCreateSddcTraceflowActivityInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["create_nsx_traceflow_request"] = vapiBindings_.NewReferenceType(vmwarecloudVmc_core_networkModel.CreateNsxTraceflowRequestBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["create_nsx_traceflow_request"] = "CreateNsxTraceflowRequest"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func TraceflowcreateTraceflowCreateSddcTraceflowActivityOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(vmwarecloudVmc_core_networkModel.ActivityBindingType)
}

func traceflowcreateTraceflowCreateSddcTraceflowActivityRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["create_nsx_traceflow_request"] = vapiBindings_.NewReferenceType(vmwarecloudVmc_core_networkModel.CreateNsxTraceflowRequestBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["create_nsx_traceflow_request"] = "CreateNsxTraceflowRequest"
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["create_nsx_traceflow_request"] = vapiBindings_.NewReferenceType(vmwarecloudVmc_core_networkModel.CreateNsxTraceflowRequestBindingType)
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	pathParams["org_id"] = "orgId"
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
		"create_nsx_traceflow_request",
		"POST",
		"/api/network/{orgId}/core/troubleshooting/traceflow:create-traceflow",
		"application/json",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthenticated": 401, "com.vmware.vapi.std.errors.unauthorized": 403})
}
