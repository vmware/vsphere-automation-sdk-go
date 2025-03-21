// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: PrecheckByClusters.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package nvds_urt

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

func precheckByClustersCreateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["precheck_parameters"] = vapiBindings_.NewReferenceType(nsxModel.PrecheckParametersBindingType)
	fields["tolerate_different_configurations"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["precheck_parameters"] = "PrecheckParameters"
	fieldNameMap["tolerate_different_configurations"] = "TolerateDifferentConfigurations"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func PrecheckByClustersCreateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.NvdsUpgradePrecheckIdBindingType)
}

func precheckByClustersCreateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["precheck_parameters"] = vapiBindings_.NewReferenceType(nsxModel.PrecheckParametersBindingType)
	fields["tolerate_different_configurations"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["precheck_parameters"] = "PrecheckParameters"
	fieldNameMap["tolerate_different_configurations"] = "TolerateDifferentConfigurations"
	paramsTypeMap["precheck_parameters"] = vapiBindings_.NewReferenceType(nsxModel.PrecheckParametersBindingType)
	paramsTypeMap["tolerate_different_configurations"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	queryParams["tolerate_different_configurations"] = "tolerate_different_configurations"
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
		"precheck_parameters",
		"POST",
		"/api/v1/nvds-urt/precheck-by-clusters",
		"",
		resultHeaders,
		201,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
