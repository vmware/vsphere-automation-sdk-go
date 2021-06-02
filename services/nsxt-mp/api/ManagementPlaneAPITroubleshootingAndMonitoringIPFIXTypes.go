// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: ManagementPlaneAPITroubleshootingAndMonitoringIPFIX.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package api

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/model"
	"reflect"
)

// Possible value for ``ipfixConfigType`` of method ManagementPlaneAPITroubleshootingAndMonitoringIPFIX#listIpfixConfig.
const ManagementPlaneAPITroubleshootingAndMonitoringIPFIX_LIST_IPFIX_CONFIG_IPFIX_CONFIG_TYPE_IPFIXSWITCHCONFIG = "IpfixSwitchConfig"

// Possible value for ``ipfixConfigType`` of method ManagementPlaneAPITroubleshootingAndMonitoringIPFIX#listIpfixConfig.
const ManagementPlaneAPITroubleshootingAndMonitoringIPFIX_LIST_IPFIX_CONFIG_IPFIX_CONFIG_TYPE_IPFIXDFWCONFIG = "IpfixDfwConfig"

// Possible value for ``appliedToEntityType`` of method ManagementPlaneAPITroubleshootingAndMonitoringIPFIX#listIpfixUpmProfiles.
const ManagementPlaneAPITroubleshootingAndMonitoringIPFIX_LIST_IPFIX_UPM_PROFILES_APPLIED_TO_ENTITY_TYPE_LOGICALPORT = "LogicalPort"

// Possible value for ``appliedToEntityType`` of method ManagementPlaneAPITroubleshootingAndMonitoringIPFIX#listIpfixUpmProfiles.
const ManagementPlaneAPITroubleshootingAndMonitoringIPFIX_LIST_IPFIX_UPM_PROFILES_APPLIED_TO_ENTITY_TYPE_LOGICALSWITCH = "LogicalSwitch"

// Possible value for ``appliedToEntityType`` of method ManagementPlaneAPITroubleshootingAndMonitoringIPFIX#listIpfixUpmProfiles.
const ManagementPlaneAPITroubleshootingAndMonitoringIPFIX_LIST_IPFIX_UPM_PROFILES_APPLIED_TO_ENTITY_TYPE_NSGROUP = "NSGroup"

func managementPlaneAPITroubleshootingAndMonitoringIPFIXCreateIpfixCollectorConfigInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipfix_collector_config"] = bindings.NewReferenceType(model.IpfixCollectorConfigBindingType)
	fieldNameMap["ipfix_collector_config"] = "IpfixCollectorConfig"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXCreateIpfixCollectorConfigOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IpfixCollectorConfigBindingType)
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXCreateIpfixCollectorConfigRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ipfix_collector_config"] = bindings.NewReferenceType(model.IpfixCollectorConfigBindingType)
	fieldNameMap["ipfix_collector_config"] = "IpfixCollectorConfig"
	paramsTypeMap["ipfix_collector_config"] = bindings.NewReferenceType(model.IpfixCollectorConfigBindingType)
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
		"ipfix_collector_config",
		"POST",
		"/api/v1/ipfix/collectorconfigs",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXCreateIpfixCollectorUpmProfileInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipfix_collector_upm_profile"] = bindings.NewReferenceType(model.IpfixCollectorUpmProfileBindingType)
	fieldNameMap["ipfix_collector_upm_profile"] = "IpfixCollectorUpmProfile"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXCreateIpfixCollectorUpmProfileOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IpfixCollectorUpmProfileBindingType)
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXCreateIpfixCollectorUpmProfileRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ipfix_collector_upm_profile"] = bindings.NewReferenceType(model.IpfixCollectorUpmProfileBindingType)
	fieldNameMap["ipfix_collector_upm_profile"] = "IpfixCollectorUpmProfile"
	paramsTypeMap["ipfix_collector_upm_profile"] = bindings.NewReferenceType(model.IpfixCollectorUpmProfileBindingType)
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
		"ipfix_collector_upm_profile",
		"POST",
		"/api/v1/ipfix-collector-profiles",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXCreateIpfixConfigInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipfix_config"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.IpfixConfigBindingType)}, bindings.REST)
	fieldNameMap["ipfix_config"] = "IpfixConfig"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXCreateIpfixConfigOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.IpfixConfigBindingType)}, bindings.REST)
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXCreateIpfixConfigRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ipfix_config"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.IpfixConfigBindingType)}, bindings.REST)
	fieldNameMap["ipfix_config"] = "IpfixConfig"
	paramsTypeMap["ipfix_config"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.IpfixConfigBindingType)}, bindings.REST)
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
		"ipfix_config",
		"POST",
		"/api/v1/ipfix/configs",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXCreateIpfixUpmProfileInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipfix_upm_profile"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.IpfixUpmProfileBindingType)}, bindings.REST)
	fieldNameMap["ipfix_upm_profile"] = "IpfixUpmProfile"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXCreateIpfixUpmProfileOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.IpfixUpmProfileBindingType)}, bindings.REST)
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXCreateIpfixUpmProfileRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ipfix_upm_profile"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.IpfixUpmProfileBindingType)}, bindings.REST)
	fieldNameMap["ipfix_upm_profile"] = "IpfixUpmProfile"
	paramsTypeMap["ipfix_upm_profile"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.IpfixUpmProfileBindingType)}, bindings.REST)
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
		"ipfix_upm_profile",
		"POST",
		"/api/v1/ipfix-profiles",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXDeleteIpfixCollectorConfigInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["collector_config_id"] = bindings.NewStringType()
	fieldNameMap["collector_config_id"] = "CollectorConfigId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXDeleteIpfixCollectorConfigOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXDeleteIpfixCollectorConfigRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["collector_config_id"] = bindings.NewStringType()
	fieldNameMap["collector_config_id"] = "CollectorConfigId"
	paramsTypeMap["collector_config_id"] = bindings.NewStringType()
	paramsTypeMap["collectorConfigId"] = bindings.NewStringType()
	pathParams["collector_config_id"] = "collectorConfigId"
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
		"/api/v1/ipfix/collectorconfigs/{collectorConfigId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXDeleteIpfixCollectorUpmProfileInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipfix_collector_profile_id"] = bindings.NewStringType()
	fieldNameMap["ipfix_collector_profile_id"] = "IpfixCollectorProfileId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXDeleteIpfixCollectorUpmProfileOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXDeleteIpfixCollectorUpmProfileRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ipfix_collector_profile_id"] = bindings.NewStringType()
	fieldNameMap["ipfix_collector_profile_id"] = "IpfixCollectorProfileId"
	paramsTypeMap["ipfix_collector_profile_id"] = bindings.NewStringType()
	paramsTypeMap["ipfixCollectorProfileId"] = bindings.NewStringType()
	pathParams["ipfix_collector_profile_id"] = "ipfixCollectorProfileId"
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
		"/api/v1/ipfix-collector-profiles/{ipfixCollectorProfileId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXDeleteIpfixConfigInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["config_id"] = bindings.NewStringType()
	fieldNameMap["config_id"] = "ConfigId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXDeleteIpfixConfigOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXDeleteIpfixConfigRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["config_id"] = bindings.NewStringType()
	fieldNameMap["config_id"] = "ConfigId"
	paramsTypeMap["config_id"] = bindings.NewStringType()
	paramsTypeMap["configId"] = bindings.NewStringType()
	pathParams["config_id"] = "configId"
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
		"/api/v1/ipfix/configs/{configId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXDeleteIpfixUpmProfileInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipfix_profile_id"] = bindings.NewStringType()
	fieldNameMap["ipfix_profile_id"] = "IpfixProfileId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXDeleteIpfixUpmProfileOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXDeleteIpfixUpmProfileRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ipfix_profile_id"] = bindings.NewStringType()
	fieldNameMap["ipfix_profile_id"] = "IpfixProfileId"
	paramsTypeMap["ipfix_profile_id"] = bindings.NewStringType()
	paramsTypeMap["ipfixProfileId"] = bindings.NewStringType()
	pathParams["ipfix_profile_id"] = "ipfixProfileId"
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
		"/api/v1/ipfix-profiles/{ipfixProfileId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXGetIpfixCollectorConfigInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["collector_config_id"] = bindings.NewStringType()
	fieldNameMap["collector_config_id"] = "CollectorConfigId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXGetIpfixCollectorConfigOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IpfixCollectorConfigBindingType)
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXGetIpfixCollectorConfigRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["collector_config_id"] = bindings.NewStringType()
	fieldNameMap["collector_config_id"] = "CollectorConfigId"
	paramsTypeMap["collector_config_id"] = bindings.NewStringType()
	paramsTypeMap["collectorConfigId"] = bindings.NewStringType()
	pathParams["collector_config_id"] = "collectorConfigId"
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
		"/api/v1/ipfix/collectorconfigs/{collectorConfigId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXGetIpfixCollectorUpmProfileInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipfix_collector_profile_id"] = bindings.NewStringType()
	fieldNameMap["ipfix_collector_profile_id"] = "IpfixCollectorProfileId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXGetIpfixCollectorUpmProfileOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IpfixCollectorUpmProfileBindingType)
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXGetIpfixCollectorUpmProfileRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ipfix_collector_profile_id"] = bindings.NewStringType()
	fieldNameMap["ipfix_collector_profile_id"] = "IpfixCollectorProfileId"
	paramsTypeMap["ipfix_collector_profile_id"] = bindings.NewStringType()
	paramsTypeMap["ipfixCollectorProfileId"] = bindings.NewStringType()
	pathParams["ipfix_collector_profile_id"] = "ipfixCollectorProfileId"
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
		"/api/v1/ipfix-collector-profiles/{ipfixCollectorProfileId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXGetIpfixConfigInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["config_id"] = bindings.NewStringType()
	fieldNameMap["config_id"] = "ConfigId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXGetIpfixConfigOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.IpfixConfigBindingType)}, bindings.REST)
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXGetIpfixConfigRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["config_id"] = bindings.NewStringType()
	fieldNameMap["config_id"] = "ConfigId"
	paramsTypeMap["config_id"] = bindings.NewStringType()
	paramsTypeMap["configId"] = bindings.NewStringType()
	pathParams["config_id"] = "configId"
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
		"/api/v1/ipfix/configs/{configId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXGetIpfixObsPointsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXGetIpfixObsPointsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IpfixObsPointsListResultBindingType)
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXGetIpfixObsPointsRestMetadata() protocol.OperationRestMetadata {
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
		"/api/v1/ipfix-obs-points",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXGetIpfixUpmProfileInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipfix_profile_id"] = bindings.NewStringType()
	fieldNameMap["ipfix_profile_id"] = "IpfixProfileId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXGetIpfixUpmProfileOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.IpfixUpmProfileBindingType)}, bindings.REST)
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXGetIpfixUpmProfileRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ipfix_profile_id"] = bindings.NewStringType()
	fieldNameMap["ipfix_profile_id"] = "IpfixProfileId"
	paramsTypeMap["ipfix_profile_id"] = bindings.NewStringType()
	paramsTypeMap["ipfixProfileId"] = bindings.NewStringType()
	pathParams["ipfix_profile_id"] = "ipfixProfileId"
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
		"/api/v1/ipfix-profiles/{ipfixProfileId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXGetSwitchIpfixConfigInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXGetSwitchIpfixConfigOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IpfixObsPointConfigBindingType)
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXGetSwitchIpfixConfigRestMetadata() protocol.OperationRestMetadata {
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
		"/api/v1/ipfix-obs-points/switch-global",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXListIpfixCollectorConfigInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXListIpfixCollectorConfigOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IpfixCollectorConfigListResultBindingType)
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXListIpfixCollectorConfigRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sort_by"] = "sort_by"
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
		"/api/v1/ipfix/collectorconfigs",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXListIpfixCollectorUpmProfilesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["profile_types"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["profile_types"] = "ProfileTypes"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXListIpfixCollectorUpmProfilesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IpfixCollectorUpmProfileListResultBindingType)
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXListIpfixCollectorUpmProfilesRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["profile_types"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["profile_types"] = "ProfileTypes"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["profile_types"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["profile_types"] = "profile_types"
	queryParams["sort_by"] = "sort_by"
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
		"/api/v1/ipfix-collector-profiles",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXListIpfixConfigInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["applied_to"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["ipfix_config_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["applied_to"] = "AppliedTo"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["ipfix_config_type"] = "IpfixConfigType"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXListIpfixConfigOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IpfixConfigListResultBindingType)
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXListIpfixConfigRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["applied_to"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["ipfix_config_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["applied_to"] = "AppliedTo"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["ipfix_config_type"] = "IpfixConfigType"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["applied_to"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["ipfix_config_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["ipfix_config_type"] = "ipfix_config_type"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["applied_to"] = "applied_to"
	queryParams["sort_by"] = "sort_by"
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
		"/api/v1/ipfix/configs",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXListIpfixUpmProfilesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["applied_to_entity_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["applied_to_entity_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["profile_types"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["applied_to_entity_id"] = "AppliedToEntityId"
	fieldNameMap["applied_to_entity_type"] = "AppliedToEntityType"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["profile_types"] = "ProfileTypes"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXListIpfixUpmProfilesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IpfixUpmProfileListResultBindingType)
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXListIpfixUpmProfilesRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["applied_to_entity_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["applied_to_entity_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["profile_types"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["applied_to_entity_id"] = "AppliedToEntityId"
	fieldNameMap["applied_to_entity_type"] = "AppliedToEntityType"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["profile_types"] = "ProfileTypes"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["profile_types"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["applied_to_entity_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["applied_to_entity_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["applied_to_entity_id"] = "applied_to_entity_id"
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["profile_types"] = "profile_types"
	queryParams["sort_by"] = "sort_by"
	queryParams["applied_to_entity_type"] = "applied_to_entity_type"
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
		"/api/v1/ipfix-profiles",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXUpdateIpfixCollectorConfigInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["collector_config_id"] = bindings.NewStringType()
	fields["ipfix_collector_config"] = bindings.NewReferenceType(model.IpfixCollectorConfigBindingType)
	fieldNameMap["collector_config_id"] = "CollectorConfigId"
	fieldNameMap["ipfix_collector_config"] = "IpfixCollectorConfig"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXUpdateIpfixCollectorConfigOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IpfixCollectorConfigBindingType)
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXUpdateIpfixCollectorConfigRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["collector_config_id"] = bindings.NewStringType()
	fields["ipfix_collector_config"] = bindings.NewReferenceType(model.IpfixCollectorConfigBindingType)
	fieldNameMap["collector_config_id"] = "CollectorConfigId"
	fieldNameMap["ipfix_collector_config"] = "IpfixCollectorConfig"
	paramsTypeMap["ipfix_collector_config"] = bindings.NewReferenceType(model.IpfixCollectorConfigBindingType)
	paramsTypeMap["collector_config_id"] = bindings.NewStringType()
	paramsTypeMap["collectorConfigId"] = bindings.NewStringType()
	pathParams["collector_config_id"] = "collectorConfigId"
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
		"ipfix_collector_config",
		"PUT",
		"/api/v1/ipfix/collectorconfigs/{collectorConfigId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXUpdateIpfixCollectorUpmProfileInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipfix_collector_profile_id"] = bindings.NewStringType()
	fields["ipfix_collector_upm_profile"] = bindings.NewReferenceType(model.IpfixCollectorUpmProfileBindingType)
	fieldNameMap["ipfix_collector_profile_id"] = "IpfixCollectorProfileId"
	fieldNameMap["ipfix_collector_upm_profile"] = "IpfixCollectorUpmProfile"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXUpdateIpfixCollectorUpmProfileOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IpfixCollectorUpmProfileBindingType)
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXUpdateIpfixCollectorUpmProfileRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ipfix_collector_profile_id"] = bindings.NewStringType()
	fields["ipfix_collector_upm_profile"] = bindings.NewReferenceType(model.IpfixCollectorUpmProfileBindingType)
	fieldNameMap["ipfix_collector_profile_id"] = "IpfixCollectorProfileId"
	fieldNameMap["ipfix_collector_upm_profile"] = "IpfixCollectorUpmProfile"
	paramsTypeMap["ipfix_collector_profile_id"] = bindings.NewStringType()
	paramsTypeMap["ipfix_collector_upm_profile"] = bindings.NewReferenceType(model.IpfixCollectorUpmProfileBindingType)
	paramsTypeMap["ipfixCollectorProfileId"] = bindings.NewStringType()
	pathParams["ipfix_collector_profile_id"] = "ipfixCollectorProfileId"
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
		"ipfix_collector_upm_profile",
		"PUT",
		"/api/v1/ipfix-collector-profiles/{ipfixCollectorProfileId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXUpdateIpfixConfigInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["config_id"] = bindings.NewStringType()
	fields["ipfix_config"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.IpfixConfigBindingType)}, bindings.REST)
	fieldNameMap["config_id"] = "ConfigId"
	fieldNameMap["ipfix_config"] = "IpfixConfig"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXUpdateIpfixConfigOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.IpfixConfigBindingType)}, bindings.REST)
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXUpdateIpfixConfigRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["config_id"] = bindings.NewStringType()
	fields["ipfix_config"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.IpfixConfigBindingType)}, bindings.REST)
	fieldNameMap["config_id"] = "ConfigId"
	fieldNameMap["ipfix_config"] = "IpfixConfig"
	paramsTypeMap["ipfix_config"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.IpfixConfigBindingType)}, bindings.REST)
	paramsTypeMap["config_id"] = bindings.NewStringType()
	paramsTypeMap["configId"] = bindings.NewStringType()
	pathParams["config_id"] = "configId"
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
		"ipfix_config",
		"PUT",
		"/api/v1/ipfix/configs/{configId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXUpdateIpfixUpmProfileInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipfix_profile_id"] = bindings.NewStringType()
	fields["ipfix_upm_profile"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.IpfixUpmProfileBindingType)}, bindings.REST)
	fieldNameMap["ipfix_profile_id"] = "IpfixProfileId"
	fieldNameMap["ipfix_upm_profile"] = "IpfixUpmProfile"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXUpdateIpfixUpmProfileOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.IpfixUpmProfileBindingType)}, bindings.REST)
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXUpdateIpfixUpmProfileRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ipfix_profile_id"] = bindings.NewStringType()
	fields["ipfix_upm_profile"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.IpfixUpmProfileBindingType)}, bindings.REST)
	fieldNameMap["ipfix_profile_id"] = "IpfixProfileId"
	fieldNameMap["ipfix_upm_profile"] = "IpfixUpmProfile"
	paramsTypeMap["ipfix_upm_profile"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.IpfixUpmProfileBindingType)}, bindings.REST)
	paramsTypeMap["ipfix_profile_id"] = bindings.NewStringType()
	paramsTypeMap["ipfixProfileId"] = bindings.NewStringType()
	pathParams["ipfix_profile_id"] = "ipfixProfileId"
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
		"ipfix_upm_profile",
		"PUT",
		"/api/v1/ipfix-profiles/{ipfixProfileId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXUpdateSwitchIpfixConfigInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipfix_obs_point_config"] = bindings.NewReferenceType(model.IpfixObsPointConfigBindingType)
	fieldNameMap["ipfix_obs_point_config"] = "IpfixObsPointConfig"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXUpdateSwitchIpfixConfigOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IpfixObsPointConfigBindingType)
}

func managementPlaneAPITroubleshootingAndMonitoringIPFIXUpdateSwitchIpfixConfigRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ipfix_obs_point_config"] = bindings.NewReferenceType(model.IpfixObsPointConfigBindingType)
	fieldNameMap["ipfix_obs_point_config"] = "IpfixObsPointConfig"
	paramsTypeMap["ipfix_obs_point_config"] = bindings.NewReferenceType(model.IpfixObsPointConfigBindingType)
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
		"ipfix_obs_point_config",
		"PUT",
		"/api/v1/ipfix-obs-points/switch-global",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
