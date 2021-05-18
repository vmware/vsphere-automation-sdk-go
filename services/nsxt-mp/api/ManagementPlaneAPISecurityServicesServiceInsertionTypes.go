// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: ManagementPlaneAPISecurityServicesServiceInsertion.
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

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesServiceInsertion#addServiceInsertionRuleInSection.
const ManagementPlaneAPISecurityServicesServiceInsertion_ADD_SERVICE_INSERTION_RULE_IN_SECTION_OPERATION_TOP = "insert_top"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesServiceInsertion#addServiceInsertionRuleInSection.
const ManagementPlaneAPISecurityServicesServiceInsertion_ADD_SERVICE_INSERTION_RULE_IN_SECTION_OPERATION_BOTTOM = "insert_bottom"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesServiceInsertion#addServiceInsertionRuleInSection.
const ManagementPlaneAPISecurityServicesServiceInsertion_ADD_SERVICE_INSERTION_RULE_IN_SECTION_OPERATION_AFTER = "insert_after"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesServiceInsertion#addServiceInsertionRuleInSection.
const ManagementPlaneAPISecurityServicesServiceInsertion_ADD_SERVICE_INSERTION_RULE_IN_SECTION_OPERATION_BEFORE = "insert_before"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesServiceInsertion#addServiceInsertionRulesInSectionCreateMultiple.
const ManagementPlaneAPISecurityServicesServiceInsertion_ADD_SERVICE_INSERTION_RULES_IN_SECTION_CREATE_MULTIPLE_OPERATION_TOP = "insert_top"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesServiceInsertion#addServiceInsertionRulesInSectionCreateMultiple.
const ManagementPlaneAPISecurityServicesServiceInsertion_ADD_SERVICE_INSERTION_RULES_IN_SECTION_CREATE_MULTIPLE_OPERATION_BOTTOM = "insert_bottom"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesServiceInsertion#addServiceInsertionRulesInSectionCreateMultiple.
const ManagementPlaneAPISecurityServicesServiceInsertion_ADD_SERVICE_INSERTION_RULES_IN_SECTION_CREATE_MULTIPLE_OPERATION_AFTER = "insert_after"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesServiceInsertion#addServiceInsertionRulesInSectionCreateMultiple.
const ManagementPlaneAPISecurityServicesServiceInsertion_ADD_SERVICE_INSERTION_RULES_IN_SECTION_CREATE_MULTIPLE_OPERATION_BEFORE = "insert_before"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesServiceInsertion#addServiceInsertionSection.
const ManagementPlaneAPISecurityServicesServiceInsertion_ADD_SERVICE_INSERTION_SECTION_OPERATION_TOP = "insert_top"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesServiceInsertion#addServiceInsertionSection.
const ManagementPlaneAPISecurityServicesServiceInsertion_ADD_SERVICE_INSERTION_SECTION_OPERATION_BOTTOM = "insert_bottom"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesServiceInsertion#addServiceInsertionSection.
const ManagementPlaneAPISecurityServicesServiceInsertion_ADD_SERVICE_INSERTION_SECTION_OPERATION_AFTER = "insert_after"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesServiceInsertion#addServiceInsertionSection.
const ManagementPlaneAPISecurityServicesServiceInsertion_ADD_SERVICE_INSERTION_SECTION_OPERATION_BEFORE = "insert_before"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesServiceInsertion#addServiceInsertionSectionWithRulesCreateWithRules.
const ManagementPlaneAPISecurityServicesServiceInsertion_ADD_SERVICE_INSERTION_SECTION_WITH_RULES_CREATE_WITH_RULES_OPERATION_TOP = "insert_top"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesServiceInsertion#addServiceInsertionSectionWithRulesCreateWithRules.
const ManagementPlaneAPISecurityServicesServiceInsertion_ADD_SERVICE_INSERTION_SECTION_WITH_RULES_CREATE_WITH_RULES_OPERATION_BOTTOM = "insert_bottom"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesServiceInsertion#addServiceInsertionSectionWithRulesCreateWithRules.
const ManagementPlaneAPISecurityServicesServiceInsertion_ADD_SERVICE_INSERTION_SECTION_WITH_RULES_CREATE_WITH_RULES_OPERATION_AFTER = "insert_after"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesServiceInsertion#addServiceInsertionSectionWithRulesCreateWithRules.
const ManagementPlaneAPISecurityServicesServiceInsertion_ADD_SERVICE_INSERTION_SECTION_WITH_RULES_CREATE_WITH_RULES_OPERATION_BEFORE = "insert_before"

// Possible value for ``source`` of method ManagementPlaneAPISecurityServicesServiceInsertion#getRuntimeInterfaceOperationalStatus.
const ManagementPlaneAPISecurityServicesServiceInsertion_GET_RUNTIME_INTERFACE_OPERATIONAL_STATUS_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method ManagementPlaneAPISecurityServicesServiceInsertion#getRuntimeInterfaceOperationalStatus.
const ManagementPlaneAPISecurityServicesServiceInsertion_GET_RUNTIME_INTERFACE_OPERATIONAL_STATUS_SOURCE_CACHED = "cached"

// Possible value for ``source`` of method ManagementPlaneAPISecurityServicesServiceInsertion#getRuntimeInterfaceStatistics.
const ManagementPlaneAPISecurityServicesServiceInsertion_GET_RUNTIME_INTERFACE_STATISTICS_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method ManagementPlaneAPISecurityServicesServiceInsertion#getRuntimeInterfaceStatistics.
const ManagementPlaneAPISecurityServicesServiceInsertion_GET_RUNTIME_INTERFACE_STATISTICS_SOURCE_CACHED = "cached"

// Possible value for ``source`` of method ManagementPlaneAPISecurityServicesServiceInsertion#getServiceDeploymentStatus.
const ManagementPlaneAPISecurityServicesServiceInsertion_GET_SERVICE_DEPLOYMENT_STATUS_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method ManagementPlaneAPISecurityServicesServiceInsertion#getServiceDeploymentStatus.
const ManagementPlaneAPISecurityServicesServiceInsertion_GET_SERVICE_DEPLOYMENT_STATUS_SOURCE_CACHED = "cached"

// Possible value for ``filterType`` of method ManagementPlaneAPISecurityServicesServiceInsertion#getServiceInsertionRules.
const ManagementPlaneAPISecurityServicesServiceInsertion_GET_SERVICE_INSERTION_RULES_FILTER_TYPE_FILTER = "FILTER"

// Possible value for ``filterType`` of method ManagementPlaneAPISecurityServicesServiceInsertion#getServiceInsertionRules.
const ManagementPlaneAPISecurityServicesServiceInsertion_GET_SERVICE_INSERTION_RULES_FILTER_TYPE_SEARCH = "SEARCH"

// Possible value for ``source`` of method ManagementPlaneAPISecurityServicesServiceInsertion#getServiceInstanceStatus.
const ManagementPlaneAPISecurityServicesServiceInsertion_GET_SERVICE_INSTANCE_STATUS_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method ManagementPlaneAPISecurityServicesServiceInsertion#getServiceInstanceStatus.
const ManagementPlaneAPISecurityServicesServiceInsertion_GET_SERVICE_INSTANCE_STATUS_SOURCE_CACHED = "cached"

// Possible value for ``excludeAppliedToType`` of method ManagementPlaneAPISecurityServicesServiceInsertion#listServiceInsertionSections.
const ManagementPlaneAPISecurityServicesServiceInsertion_LIST_SERVICE_INSERTION_SECTIONS_EXCLUDE_APPLIED_TO_TYPE_NSGROUP = "NSGroup"

// Possible value for ``excludeAppliedToType`` of method ManagementPlaneAPISecurityServicesServiceInsertion#listServiceInsertionSections.
const ManagementPlaneAPISecurityServicesServiceInsertion_LIST_SERVICE_INSERTION_SECTIONS_EXCLUDE_APPLIED_TO_TYPE_LOGICALSWITCH = "LogicalSwitch"

// Possible value for ``excludeAppliedToType`` of method ManagementPlaneAPISecurityServicesServiceInsertion#listServiceInsertionSections.
const ManagementPlaneAPISecurityServicesServiceInsertion_LIST_SERVICE_INSERTION_SECTIONS_EXCLUDE_APPLIED_TO_TYPE_LOGICALROUTER = "LogicalRouter"

// Possible value for ``excludeAppliedToType`` of method ManagementPlaneAPISecurityServicesServiceInsertion#listServiceInsertionSections.
const ManagementPlaneAPISecurityServicesServiceInsertion_LIST_SERVICE_INSERTION_SECTIONS_EXCLUDE_APPLIED_TO_TYPE_LOGICALPORT = "LogicalPort"

// Possible value for ``filterType`` of method ManagementPlaneAPISecurityServicesServiceInsertion#listServiceInsertionSections.
const ManagementPlaneAPISecurityServicesServiceInsertion_LIST_SERVICE_INSERTION_SECTIONS_FILTER_TYPE_FILTER = "FILTER"

// Possible value for ``filterType`` of method ManagementPlaneAPISecurityServicesServiceInsertion#listServiceInsertionSections.
const ManagementPlaneAPISecurityServicesServiceInsertion_LIST_SERVICE_INSERTION_SECTIONS_FILTER_TYPE_SEARCH = "SEARCH"

// Possible value for ``includeAppliedToType`` of method ManagementPlaneAPISecurityServicesServiceInsertion#listServiceInsertionSections.
const ManagementPlaneAPISecurityServicesServiceInsertion_LIST_SERVICE_INSERTION_SECTIONS_INCLUDE_APPLIED_TO_TYPE_NSGROUP = "NSGroup"

// Possible value for ``includeAppliedToType`` of method ManagementPlaneAPISecurityServicesServiceInsertion#listServiceInsertionSections.
const ManagementPlaneAPISecurityServicesServiceInsertion_LIST_SERVICE_INSERTION_SECTIONS_INCLUDE_APPLIED_TO_TYPE_LOGICALSWITCH = "LogicalSwitch"

// Possible value for ``includeAppliedToType`` of method ManagementPlaneAPISecurityServicesServiceInsertion#listServiceInsertionSections.
const ManagementPlaneAPISecurityServicesServiceInsertion_LIST_SERVICE_INSERTION_SECTIONS_INCLUDE_APPLIED_TO_TYPE_LOGICALROUTER = "LogicalRouter"

// Possible value for ``includeAppliedToType`` of method ManagementPlaneAPISecurityServicesServiceInsertion#listServiceInsertionSections.
const ManagementPlaneAPISecurityServicesServiceInsertion_LIST_SERVICE_INSERTION_SECTIONS_INCLUDE_APPLIED_TO_TYPE_LOGICALPORT = "LogicalPort"

// Possible value for ``type`` of method ManagementPlaneAPISecurityServicesServiceInsertion#listServiceInsertionSections.
const ManagementPlaneAPISecurityServicesServiceInsertion_LIST_SERVICE_INSERTION_SECTIONS_TYPE_L3REDIRECT = "L3REDIRECT"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesServiceInsertion#reviseServiceInsertionRuleRevise.
const ManagementPlaneAPISecurityServicesServiceInsertion_REVISE_SERVICE_INSERTION_RULE_REVISE_OPERATION_TOP = "insert_top"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesServiceInsertion#reviseServiceInsertionRuleRevise.
const ManagementPlaneAPISecurityServicesServiceInsertion_REVISE_SERVICE_INSERTION_RULE_REVISE_OPERATION_BOTTOM = "insert_bottom"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesServiceInsertion#reviseServiceInsertionRuleRevise.
const ManagementPlaneAPISecurityServicesServiceInsertion_REVISE_SERVICE_INSERTION_RULE_REVISE_OPERATION_AFTER = "insert_after"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesServiceInsertion#reviseServiceInsertionRuleRevise.
const ManagementPlaneAPISecurityServicesServiceInsertion_REVISE_SERVICE_INSERTION_RULE_REVISE_OPERATION_BEFORE = "insert_before"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesServiceInsertion#reviseServiceInsertionSectionRevise.
const ManagementPlaneAPISecurityServicesServiceInsertion_REVISE_SERVICE_INSERTION_SECTION_REVISE_OPERATION_TOP = "insert_top"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesServiceInsertion#reviseServiceInsertionSectionRevise.
const ManagementPlaneAPISecurityServicesServiceInsertion_REVISE_SERVICE_INSERTION_SECTION_REVISE_OPERATION_BOTTOM = "insert_bottom"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesServiceInsertion#reviseServiceInsertionSectionRevise.
const ManagementPlaneAPISecurityServicesServiceInsertion_REVISE_SERVICE_INSERTION_SECTION_REVISE_OPERATION_AFTER = "insert_after"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesServiceInsertion#reviseServiceInsertionSectionRevise.
const ManagementPlaneAPISecurityServicesServiceInsertion_REVISE_SERVICE_INSERTION_SECTION_REVISE_OPERATION_BEFORE = "insert_before"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesServiceInsertion#reviseServiceInsertionSectionWithRulesReviseWithRules.
const ManagementPlaneAPISecurityServicesServiceInsertion_REVISE_SERVICE_INSERTION_SECTION_WITH_RULES_REVISE_WITH_RULES_OPERATION_TOP = "insert_top"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesServiceInsertion#reviseServiceInsertionSectionWithRulesReviseWithRules.
const ManagementPlaneAPISecurityServicesServiceInsertion_REVISE_SERVICE_INSERTION_SECTION_WITH_RULES_REVISE_WITH_RULES_OPERATION_BOTTOM = "insert_bottom"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesServiceInsertion#reviseServiceInsertionSectionWithRulesReviseWithRules.
const ManagementPlaneAPISecurityServicesServiceInsertion_REVISE_SERVICE_INSERTION_SECTION_WITH_RULES_REVISE_WITH_RULES_OPERATION_AFTER = "insert_after"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesServiceInsertion#reviseServiceInsertionSectionWithRulesReviseWithRules.
const ManagementPlaneAPISecurityServicesServiceInsertion_REVISE_SERVICE_INSERTION_SECTION_WITH_RULES_REVISE_WITH_RULES_OPERATION_BEFORE = "insert_before"

// Possible value for ``action`` of method ManagementPlaneAPISecurityServicesServiceInsertion#updateServiceVmState.
const ManagementPlaneAPISecurityServicesServiceInsertion_UPDATE_SERVICE_VM_STATE_ACTION_ENABLE_MAINTENANCE_MODE = "enable_maintenance_mode"

// Possible value for ``action`` of method ManagementPlaneAPISecurityServicesServiceInsertion#updateServiceVmState.
const ManagementPlaneAPISecurityServicesServiceInsertion_UPDATE_SERVICE_VM_STATE_ACTION_DISABLE_MAINTENANCE_MODE = "disable_maintenance_mode"

// Possible value for ``action`` of method ManagementPlaneAPISecurityServicesServiceInsertion#updateServiceVmState.
const ManagementPlaneAPISecurityServicesServiceInsertion_UPDATE_SERVICE_VM_STATE_ACTION_IS_HEALTHY = "is_healthy"

// Possible value for ``action`` of method ManagementPlaneAPISecurityServicesServiceInsertion#updateServiceVmState.
const ManagementPlaneAPISecurityServicesServiceInsertion_UPDATE_SERVICE_VM_STATE_ACTION_IS_STOPPED = "is_stopped"

// Possible value for ``action`` of method ManagementPlaneAPISecurityServicesServiceInsertion#updateServiceVmState.
const ManagementPlaneAPISecurityServicesServiceInsertion_UPDATE_SERVICE_VM_STATE_ACTION_IS_NOT_RESPONDING = "is_not_responding"

func managementPlaneAPISecurityServicesServiceInsertionAddInstanceEndpointInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fields["instance_endpoint"] = bindings.NewReferenceType(model.InstanceEndpointBindingType)
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	fieldNameMap["instance_endpoint"] = "InstanceEndpoint"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionAddInstanceEndpointOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.InstanceEndpointBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionAddInstanceEndpointRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fields["instance_endpoint"] = bindings.NewReferenceType(model.InstanceEndpointBindingType)
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	fieldNameMap["instance_endpoint"] = "InstanceEndpoint"
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["service_instance_id"] = bindings.NewStringType()
	paramsTypeMap["instance_endpoint"] = bindings.NewReferenceType(model.InstanceEndpointBindingType)
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	paramsTypeMap["serviceInstanceId"] = bindings.NewStringType()
	pathParams["service_instance_id"] = "serviceInstanceId"
	pathParams["service_id"] = "serviceId"
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
		"instance_endpoint",
		"POST",
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/service-instances/{serviceInstanceId}/instance-endpoints",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionAddSIServiceProfileInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["base_service_profile"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.BaseServiceProfileBindingType)}, bindings.REST)
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["base_service_profile"] = "BaseServiceProfile"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionAddSIServiceProfileOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.BaseServiceProfileBindingType)}, bindings.REST)
}

func managementPlaneAPISecurityServicesServiceInsertionAddSIServiceProfileRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["base_service_profile"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.BaseServiceProfileBindingType)}, bindings.REST)
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["base_service_profile"] = "BaseServiceProfile"
	paramsTypeMap["base_service_profile"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.BaseServiceProfileBindingType)}, bindings.REST)
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	pathParams["service_id"] = "serviceId"
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
		"base_service_profile",
		"POST",
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/service-profiles",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionAddServiceAttachmentInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_attachment"] = bindings.NewReferenceType(model.ServiceAttachmentBindingType)
	fieldNameMap["service_attachment"] = "ServiceAttachment"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionAddServiceAttachmentOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceAttachmentBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionAddServiceAttachmentRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_attachment"] = bindings.NewReferenceType(model.ServiceAttachmentBindingType)
	fieldNameMap["service_attachment"] = "ServiceAttachment"
	paramsTypeMap["service_attachment"] = bindings.NewReferenceType(model.ServiceAttachmentBindingType)
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
		"service_attachment",
		"POST",
		"/api/v1/serviceinsertion/service-attachments",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionAddServiceChainInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_chain"] = bindings.NewReferenceType(model.ServiceChainBindingType)
	fieldNameMap["service_chain"] = "ServiceChain"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionAddServiceChainOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceChainBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionAddServiceChainRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_chain"] = bindings.NewReferenceType(model.ServiceChainBindingType)
	fieldNameMap["service_chain"] = "ServiceChain"
	paramsTypeMap["service_chain"] = bindings.NewReferenceType(model.ServiceChainBindingType)
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
		"service_chain",
		"POST",
		"/api/v1/serviceinsertion/service-chains",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionAddServiceInsertionExcludeListMemberAddMemberInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["resource_reference"] = bindings.NewReferenceType(model.ResourceReferenceBindingType)
	fieldNameMap["resource_reference"] = "ResourceReference"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionAddServiceInsertionExcludeListMemberAddMemberOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ResourceReferenceBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionAddServiceInsertionExcludeListMemberAddMemberRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["resource_reference"] = bindings.NewReferenceType(model.ResourceReferenceBindingType)
	fieldNameMap["resource_reference"] = "ResourceReference"
	paramsTypeMap["resource_reference"] = bindings.NewReferenceType(model.ResourceReferenceBindingType)
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
		"action=add_member",
		"resource_reference",
		"POST",
		"/api/v1/serviceinsertion/excludelist",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionAddServiceInsertionRuleInSectionInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["service_insertion_rule"] = bindings.NewReferenceType(model.ServiceInsertionRuleBindingType)
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["service_insertion_rule"] = "ServiceInsertionRule"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionAddServiceInsertionRuleInSectionOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceInsertionRuleBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionAddServiceInsertionRuleInSectionRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = bindings.NewStringType()
	fields["service_insertion_rule"] = bindings.NewReferenceType(model.ServiceInsertionRuleBindingType)
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["service_insertion_rule"] = "ServiceInsertionRule"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	paramsTypeMap["id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["section_id"] = bindings.NewStringType()
	paramsTypeMap["service_insertion_rule"] = bindings.NewReferenceType(model.ServiceInsertionRuleBindingType)
	paramsTypeMap["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sectionId"] = bindings.NewStringType()
	pathParams["section_id"] = "sectionId"
	queryParams["id"] = "id"
	queryParams["operation"] = "operation"
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
		"service_insertion_rule",
		"POST",
		"/api/v1/serviceinsertion/sections/{sectionId}/rules",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionAddServiceInsertionRulesInSectionCreateMultipleInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["service_insertion_rule_list"] = bindings.NewReferenceType(model.ServiceInsertionRuleListBindingType)
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["service_insertion_rule_list"] = "ServiceInsertionRuleList"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionAddServiceInsertionRulesInSectionCreateMultipleOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceInsertionRuleListBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionAddServiceInsertionRulesInSectionCreateMultipleRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = bindings.NewStringType()
	fields["service_insertion_rule_list"] = bindings.NewReferenceType(model.ServiceInsertionRuleListBindingType)
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["service_insertion_rule_list"] = "ServiceInsertionRuleList"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	paramsTypeMap["id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["service_insertion_rule_list"] = bindings.NewReferenceType(model.ServiceInsertionRuleListBindingType)
	paramsTypeMap["section_id"] = bindings.NewStringType()
	paramsTypeMap["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sectionId"] = bindings.NewStringType()
	pathParams["section_id"] = "sectionId"
	queryParams["id"] = "id"
	queryParams["operation"] = "operation"
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
		"action=create_multiple",
		"service_insertion_rule_list",
		"POST",
		"/api/v1/serviceinsertion/sections/{sectionId}/rules",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionAddServiceInsertionSectionInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_insertion_section"] = bindings.NewReferenceType(model.ServiceInsertionSectionBindingType)
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["service_insertion_section"] = "ServiceInsertionSection"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionAddServiceInsertionSectionOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceInsertionSectionBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionAddServiceInsertionSectionRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_insertion_section"] = bindings.NewReferenceType(model.ServiceInsertionSectionBindingType)
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["service_insertion_section"] = "ServiceInsertionSection"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	paramsTypeMap["id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["service_insertion_section"] = bindings.NewReferenceType(model.ServiceInsertionSectionBindingType)
	queryParams["id"] = "id"
	queryParams["operation"] = "operation"
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
		"service_insertion_section",
		"POST",
		"/api/v1/serviceinsertion/sections",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionAddServiceInsertionSectionWithRulesCreateWithRulesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_insertion_section_rule_list"] = bindings.NewReferenceType(model.ServiceInsertionSectionRuleListBindingType)
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["service_insertion_section_rule_list"] = "ServiceInsertionSectionRuleList"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionAddServiceInsertionSectionWithRulesCreateWithRulesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceInsertionSectionRuleListBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionAddServiceInsertionSectionWithRulesCreateWithRulesRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_insertion_section_rule_list"] = bindings.NewReferenceType(model.ServiceInsertionSectionRuleListBindingType)
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["service_insertion_section_rule_list"] = "ServiceInsertionSectionRuleList"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	paramsTypeMap["id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["service_insertion_section_rule_list"] = bindings.NewReferenceType(model.ServiceInsertionSectionRuleListBindingType)
	queryParams["id"] = "id"
	queryParams["operation"] = "operation"
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
		"action=create_with_rules",
		"service_insertion_section_rule_list",
		"POST",
		"/api/v1/serviceinsertion/sections",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionAddServiceInsertionServiceInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_definition"] = bindings.NewReferenceType(model.ServiceDefinitionBindingType)
	fieldNameMap["service_definition"] = "ServiceDefinition"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionAddServiceInsertionServiceOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceDefinitionBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionAddServiceInsertionServiceRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_definition"] = bindings.NewReferenceType(model.ServiceDefinitionBindingType)
	fieldNameMap["service_definition"] = "ServiceDefinition"
	paramsTypeMap["service_definition"] = bindings.NewReferenceType(model.ServiceDefinitionBindingType)
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
		"service_definition",
		"POST",
		"/api/v1/serviceinsertion/services/nsxt-mp",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionAddServiceInstanceInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["base_service_instance"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.BaseServiceInstanceBindingType)}, bindings.REST)
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["base_service_instance"] = "BaseServiceInstance"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionAddServiceInstanceOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.BaseServiceInstanceBindingType)}, bindings.REST)
}

func managementPlaneAPISecurityServicesServiceInsertionAddServiceInstanceRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["base_service_instance"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.BaseServiceInstanceBindingType)}, bindings.REST)
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["base_service_instance"] = "BaseServiceInstance"
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["base_service_instance"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.BaseServiceInstanceBindingType)}, bindings.REST)
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	pathParams["service_id"] = "serviceId"
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
		"base_service_instance",
		"POST",
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/service-instances",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionAddVendorTemplateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["vendor_template"] = bindings.NewReferenceType(model.VendorTemplateBindingType)
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["vendor_template"] = "VendorTemplate"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionAddVendorTemplateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.VendorTemplateBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionAddVendorTemplateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["vendor_template"] = bindings.NewReferenceType(model.VendorTemplateBindingType)
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["vendor_template"] = "VendorTemplate"
	paramsTypeMap["vendor_template"] = bindings.NewReferenceType(model.VendorTemplateBindingType)
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	pathParams["service_id"] = "serviceId"
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
		"vendor_template",
		"POST",
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/vendor-templates",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionCreateSolutionConfigInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["solution_config"] = bindings.NewReferenceType(model.SolutionConfigBindingType)
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["solution_config"] = "SolutionConfig"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionCreateSolutionConfigOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.SolutionConfigBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionCreateSolutionConfigRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["solution_config"] = bindings.NewReferenceType(model.SolutionConfigBindingType)
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["solution_config"] = "SolutionConfig"
	paramsTypeMap["solution_config"] = bindings.NewReferenceType(model.SolutionConfigBindingType)
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	pathParams["service_id"] = "serviceId"
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
		"solution_config",
		"POST",
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/solution-configs",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionDeleteInstanceEndpointInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fields["instance_endpoint_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	fieldNameMap["instance_endpoint_id"] = "InstanceEndpointId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionDeleteInstanceEndpointOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPISecurityServicesServiceInsertionDeleteInstanceEndpointRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fields["instance_endpoint_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	fieldNameMap["instance_endpoint_id"] = "InstanceEndpointId"
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["service_instance_id"] = bindings.NewStringType()
	paramsTypeMap["instance_endpoint_id"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	paramsTypeMap["serviceInstanceId"] = bindings.NewStringType()
	paramsTypeMap["instanceEndpointId"] = bindings.NewStringType()
	pathParams["instance_endpoint_id"] = "instanceEndpointId"
	pathParams["service_instance_id"] = "serviceInstanceId"
	pathParams["service_id"] = "serviceId"
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
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/service-instances/{serviceInstanceId}/instance-endpoints/{instanceEndpointId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionDeleteSIServiceProfileInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["service_profile_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_profile_id"] = "ServiceProfileId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionDeleteSIServiceProfileOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPISecurityServicesServiceInsertionDeleteSIServiceProfileRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["service_profile_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_profile_id"] = "ServiceProfileId"
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["service_profile_id"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	paramsTypeMap["serviceProfileId"] = bindings.NewStringType()
	pathParams["service_profile_id"] = "serviceProfileId"
	pathParams["service_id"] = "serviceId"
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
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/service-profiles/{serviceProfileId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionDeleteServiceAttachmentInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_attachment_id"] = bindings.NewStringType()
	fieldNameMap["service_attachment_id"] = "ServiceAttachmentId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionDeleteServiceAttachmentOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPISecurityServicesServiceInsertionDeleteServiceAttachmentRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_attachment_id"] = bindings.NewStringType()
	fieldNameMap["service_attachment_id"] = "ServiceAttachmentId"
	paramsTypeMap["service_attachment_id"] = bindings.NewStringType()
	paramsTypeMap["serviceAttachmentId"] = bindings.NewStringType()
	pathParams["service_attachment_id"] = "serviceAttachmentId"
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
		"/api/v1/serviceinsertion/service-attachments/{serviceAttachmentId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionDeleteServiceChainInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_chain_id"] = bindings.NewStringType()
	fieldNameMap["service_chain_id"] = "ServiceChainId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionDeleteServiceChainOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPISecurityServicesServiceInsertionDeleteServiceChainRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_chain_id"] = bindings.NewStringType()
	fieldNameMap["service_chain_id"] = "ServiceChainId"
	paramsTypeMap["service_chain_id"] = bindings.NewStringType()
	paramsTypeMap["serviceChainId"] = bindings.NewStringType()
	pathParams["service_chain_id"] = "serviceChainId"
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
		"/api/v1/serviceinsertion/service-chains/{serviceChainId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionDeleteServiceDeploymentInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["service_deployment_id"] = bindings.NewStringType()
	fields["force"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_deployment_id"] = "ServiceDeploymentId"
	fieldNameMap["force"] = "Force"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionDeleteServiceDeploymentOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPISecurityServicesServiceInsertionDeleteServiceDeploymentRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["service_deployment_id"] = bindings.NewStringType()
	fields["force"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_deployment_id"] = "ServiceDeploymentId"
	fieldNameMap["force"] = "Force"
	paramsTypeMap["force"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["service_deployment_id"] = bindings.NewStringType()
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	paramsTypeMap["serviceDeploymentId"] = bindings.NewStringType()
	pathParams["service_deployment_id"] = "serviceDeploymentId"
	pathParams["service_id"] = "serviceId"
	queryParams["force"] = "force"
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
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/service-deployments/{serviceDeploymentId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionDeleteServiceInsertionRuleInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["rule_id"] = bindings.NewStringType()
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["rule_id"] = "RuleId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionDeleteServiceInsertionRuleOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPISecurityServicesServiceInsertionDeleteServiceInsertionRuleRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = bindings.NewStringType()
	fields["rule_id"] = bindings.NewStringType()
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["rule_id"] = "RuleId"
	paramsTypeMap["rule_id"] = bindings.NewStringType()
	paramsTypeMap["section_id"] = bindings.NewStringType()
	paramsTypeMap["sectionId"] = bindings.NewStringType()
	paramsTypeMap["ruleId"] = bindings.NewStringType()
	pathParams["section_id"] = "sectionId"
	pathParams["rule_id"] = "ruleId"
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
		"/api/v1/serviceinsertion/sections/{sectionId}/rules/{ruleId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionDeleteServiceInsertionSectionInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["cascade"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["cascade"] = "Cascade"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionDeleteServiceInsertionSectionOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPISecurityServicesServiceInsertionDeleteServiceInsertionSectionRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = bindings.NewStringType()
	fields["cascade"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["cascade"] = "Cascade"
	paramsTypeMap["section_id"] = bindings.NewStringType()
	paramsTypeMap["cascade"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["sectionId"] = bindings.NewStringType()
	pathParams["section_id"] = "sectionId"
	queryParams["cascade"] = "cascade"
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
		"/api/v1/serviceinsertion/sections/{sectionId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionDeleteServiceInsertionServiceInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["cascade"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["cascade"] = "Cascade"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionDeleteServiceInsertionServiceOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPISecurityServicesServiceInsertionDeleteServiceInsertionServiceRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["cascade"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["cascade"] = "Cascade"
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["cascade"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	pathParams["service_id"] = "serviceId"
	queryParams["cascade"] = "cascade"
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
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionDeleteServiceInstanceInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionDeleteServiceInstanceOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPISecurityServicesServiceInsertionDeleteServiceInstanceRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["service_instance_id"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	paramsTypeMap["serviceInstanceId"] = bindings.NewStringType()
	pathParams["service_instance_id"] = "serviceInstanceId"
	pathParams["service_id"] = "serviceId"
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
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/service-instances/{serviceInstanceId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionDeleteServiceManagerInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_manager_id"] = bindings.NewStringType()
	fieldNameMap["service_manager_id"] = "ServiceManagerId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionDeleteServiceManagerOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPISecurityServicesServiceInsertionDeleteServiceManagerRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_manager_id"] = bindings.NewStringType()
	fieldNameMap["service_manager_id"] = "ServiceManagerId"
	paramsTypeMap["service_manager_id"] = bindings.NewStringType()
	paramsTypeMap["serviceManagerId"] = bindings.NewStringType()
	pathParams["service_manager_id"] = "serviceManagerId"
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
		"/api/v1/serviceinsertion/service-managers/{serviceManagerId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionDeleteServicevMsDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionDeleteServicevMsDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPISecurityServicesServiceInsertionDeleteServicevMsDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["service_instance_id"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	paramsTypeMap["serviceInstanceId"] = bindings.NewStringType()
	pathParams["service_instance_id"] = "serviceInstanceId"
	pathParams["service_id"] = "serviceId"
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
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/service-instances/{serviceInstanceId}/instance-runtimes",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionDeleteSolutionConfigInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["solution_config_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["solution_config_id"] = "SolutionConfigId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionDeleteSolutionConfigOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPISecurityServicesServiceInsertionDeleteSolutionConfigRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["solution_config_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["solution_config_id"] = "SolutionConfigId"
	paramsTypeMap["solution_config_id"] = bindings.NewStringType()
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	paramsTypeMap["solutionConfigId"] = bindings.NewStringType()
	pathParams["solution_config_id"] = "solutionConfigId"
	pathParams["service_id"] = "serviceId"
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
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/solution-configs/{solutionConfigId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionDeleteVendorTemplateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["vendor_template_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["vendor_template_id"] = "VendorTemplateId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionDeleteVendorTemplateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPISecurityServicesServiceInsertionDeleteVendorTemplateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["vendor_template_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["vendor_template_id"] = "VendorTemplateId"
	paramsTypeMap["vendor_template_id"] = bindings.NewStringType()
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	paramsTypeMap["vendorTemplateId"] = bindings.NewStringType()
	pathParams["service_id"] = "serviceId"
	pathParams["vendor_template_id"] = "vendorTemplateId"
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
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/vendor-templates/{vendorTemplateId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionDeployServiceInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["service_deployment"] = bindings.NewReferenceType(model.ServiceDeploymentBindingType)
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_deployment"] = "ServiceDeployment"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionDeployServiceOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceDeploymentBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionDeployServiceRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["service_deployment"] = bindings.NewReferenceType(model.ServiceDeploymentBindingType)
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_deployment"] = "ServiceDeployment"
	paramsTypeMap["service_deployment"] = bindings.NewReferenceType(model.ServiceDeploymentBindingType)
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	pathParams["service_id"] = "serviceId"
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
		"service_deployment",
		"POST",
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/service-deployments",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionDeployServicevMsDeployInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionDeployServicevMsDeployOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPISecurityServicesServiceInsertionDeployServicevMsDeployRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["service_instance_id"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	paramsTypeMap["serviceInstanceId"] = bindings.NewStringType()
	pathParams["service_instance_id"] = "serviceInstanceId"
	pathParams["service_id"] = "serviceId"
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
		"action=deploy",
		"",
		"POST",
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/service-instances/{serviceInstanceId}/instance-runtimes",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionGetInstanceEndpointInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fields["instance_endpoint_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	fieldNameMap["instance_endpoint_id"] = "InstanceEndpointId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionGetInstanceEndpointOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.InstanceEndpointBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionGetInstanceEndpointRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fields["instance_endpoint_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	fieldNameMap["instance_endpoint_id"] = "InstanceEndpointId"
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["service_instance_id"] = bindings.NewStringType()
	paramsTypeMap["instance_endpoint_id"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	paramsTypeMap["serviceInstanceId"] = bindings.NewStringType()
	paramsTypeMap["instanceEndpointId"] = bindings.NewStringType()
	pathParams["instance_endpoint_id"] = "instanceEndpointId"
	pathParams["service_instance_id"] = "serviceInstanceId"
	pathParams["service_id"] = "serviceId"
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
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/service-instances/{serviceInstanceId}/instance-endpoints/{instanceEndpointId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionGetRuntimeInterfaceOperationalStatusInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fields["instance_runtime_id"] = bindings.NewStringType()
	fields["interface_index"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	fieldNameMap["instance_runtime_id"] = "InstanceRuntimeId"
	fieldNameMap["interface_index"] = "InterfaceIndex"
	fieldNameMap["source"] = "Source"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionGetRuntimeInterfaceOperationalStatusOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.RuntimeInterfaceOperationalStatusBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionGetRuntimeInterfaceOperationalStatusRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fields["instance_runtime_id"] = bindings.NewStringType()
	fields["interface_index"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	fieldNameMap["instance_runtime_id"] = "InstanceRuntimeId"
	fieldNameMap["interface_index"] = "InterfaceIndex"
	fieldNameMap["source"] = "Source"
	paramsTypeMap["source"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["interface_index"] = bindings.NewStringType()
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["service_instance_id"] = bindings.NewStringType()
	paramsTypeMap["instance_runtime_id"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	paramsTypeMap["serviceInstanceId"] = bindings.NewStringType()
	paramsTypeMap["instanceRuntimeId"] = bindings.NewStringType()
	paramsTypeMap["interfaceIndex"] = bindings.NewStringType()
	pathParams["instance_runtime_id"] = "instanceRuntimeId"
	pathParams["service_instance_id"] = "serviceInstanceId"
	pathParams["interface_index"] = "interfaceIndex"
	pathParams["service_id"] = "serviceId"
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
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/service-instances/{serviceInstanceId}/instance-runtimes/{instanceRuntimeId}/interfaces/{interfaceIndex}/status",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionGetRuntimeInterfaceStatisticsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fields["instance_runtime_id"] = bindings.NewStringType()
	fields["interface_index"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	fieldNameMap["instance_runtime_id"] = "InstanceRuntimeId"
	fieldNameMap["interface_index"] = "InterfaceIndex"
	fieldNameMap["source"] = "Source"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionGetRuntimeInterfaceStatisticsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.RuntimeInterfaceStatisticsBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionGetRuntimeInterfaceStatisticsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fields["instance_runtime_id"] = bindings.NewStringType()
	fields["interface_index"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	fieldNameMap["instance_runtime_id"] = "InstanceRuntimeId"
	fieldNameMap["interface_index"] = "InterfaceIndex"
	fieldNameMap["source"] = "Source"
	paramsTypeMap["source"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["interface_index"] = bindings.NewStringType()
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["service_instance_id"] = bindings.NewStringType()
	paramsTypeMap["instance_runtime_id"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	paramsTypeMap["serviceInstanceId"] = bindings.NewStringType()
	paramsTypeMap["instanceRuntimeId"] = bindings.NewStringType()
	paramsTypeMap["interfaceIndex"] = bindings.NewStringType()
	pathParams["instance_runtime_id"] = "instanceRuntimeId"
	pathParams["service_instance_id"] = "serviceInstanceId"
	pathParams["interface_index"] = "interfaceIndex"
	pathParams["service_id"] = "serviceId"
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
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/service-instances/{serviceInstanceId}/instance-runtimes/{instanceRuntimeId}/interfaces/{interfaceIndex}/statistics",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionGetSIServiceProfileInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["service_profile_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_profile_id"] = "ServiceProfileId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionGetSIServiceProfileOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.BaseServiceProfileBindingType)}, bindings.REST)
}

func managementPlaneAPISecurityServicesServiceInsertionGetSIServiceProfileRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["service_profile_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_profile_id"] = "ServiceProfileId"
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["service_profile_id"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	paramsTypeMap["serviceProfileId"] = bindings.NewStringType()
	pathParams["service_profile_id"] = "serviceProfileId"
	pathParams["service_id"] = "serviceId"
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
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/service-profiles/{serviceProfileId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceAttachmentInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_attachment_id"] = bindings.NewStringType()
	fieldNameMap["service_attachment_id"] = "ServiceAttachmentId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceAttachmentOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceAttachmentBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceAttachmentRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_attachment_id"] = bindings.NewStringType()
	fieldNameMap["service_attachment_id"] = "ServiceAttachmentId"
	paramsTypeMap["service_attachment_id"] = bindings.NewStringType()
	paramsTypeMap["serviceAttachmentId"] = bindings.NewStringType()
	pathParams["service_attachment_id"] = "serviceAttachmentId"
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
		"/api/v1/serviceinsertion/service-attachments/{serviceAttachmentId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceChainInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_chain_id"] = bindings.NewStringType()
	fieldNameMap["service_chain_id"] = "ServiceChainId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceChainOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceChainBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceChainRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_chain_id"] = bindings.NewStringType()
	fieldNameMap["service_chain_id"] = "ServiceChainId"
	paramsTypeMap["service_chain_id"] = bindings.NewStringType()
	paramsTypeMap["serviceChainId"] = bindings.NewStringType()
	pathParams["service_chain_id"] = "serviceChainId"
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
		"/api/v1/serviceinsertion/service-chains/{serviceChainId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceDeploymentInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["service_deployment_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_deployment_id"] = "ServiceDeploymentId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceDeploymentOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceDeploymentBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceDeploymentRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["service_deployment_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_deployment_id"] = "ServiceDeploymentId"
	paramsTypeMap["service_deployment_id"] = bindings.NewStringType()
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	paramsTypeMap["serviceDeploymentId"] = bindings.NewStringType()
	pathParams["service_deployment_id"] = "serviceDeploymentId"
	pathParams["service_id"] = "serviceId"
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
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/service-deployments/{serviceDeploymentId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceDeploymentStateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["service_deployment_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_deployment_id"] = "ServiceDeploymentId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceDeploymentStateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ConfigurationStateBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceDeploymentStateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["service_deployment_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_deployment_id"] = "ServiceDeploymentId"
	paramsTypeMap["service_deployment_id"] = bindings.NewStringType()
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	paramsTypeMap["serviceDeploymentId"] = bindings.NewStringType()
	pathParams["service_deployment_id"] = "serviceDeploymentId"
	pathParams["service_id"] = "serviceId"
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
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/service-deployments/{serviceDeploymentId}/state",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceDeploymentStatusInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["service_deployment_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_deployment_id"] = "ServiceDeploymentId"
	fieldNameMap["source"] = "Source"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceDeploymentStatusOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceDeploymentStatusBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceDeploymentStatusRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["service_deployment_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_deployment_id"] = "ServiceDeploymentId"
	fieldNameMap["source"] = "Source"
	paramsTypeMap["source"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["service_deployment_id"] = bindings.NewStringType()
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	paramsTypeMap["serviceDeploymentId"] = bindings.NewStringType()
	pathParams["service_deployment_id"] = "serviceDeploymentId"
	pathParams["service_id"] = "serviceId"
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
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/service-deployments/{serviceDeploymentId}/status",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceDeploymentsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceDeploymentsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceDeploymentListResultBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceDeploymentsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	pathParams["service_id"] = "serviceId"
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
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/service-deployments",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceInsertionExcludeListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceInsertionExcludeListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.SIExcludeListBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceInsertionExcludeListRestMetadata() protocol.OperationRestMetadata {
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
		"/api/v1/serviceinsertion/excludelist",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceInsertionRuleInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["rule_id"] = bindings.NewStringType()
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["rule_id"] = "RuleId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceInsertionRuleOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceInsertionRuleBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceInsertionRuleRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = bindings.NewStringType()
	fields["rule_id"] = bindings.NewStringType()
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["rule_id"] = "RuleId"
	paramsTypeMap["rule_id"] = bindings.NewStringType()
	paramsTypeMap["section_id"] = bindings.NewStringType()
	paramsTypeMap["sectionId"] = bindings.NewStringType()
	paramsTypeMap["ruleId"] = bindings.NewStringType()
	pathParams["section_id"] = "sectionId"
	pathParams["rule_id"] = "ruleId"
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
		"/api/v1/serviceinsertion/sections/{sectionId}/rules/{ruleId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceInsertionRulesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["applied_tos"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["destinations"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["filter_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["services"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sources"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["applied_tos"] = "AppliedTos"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["destinations"] = "Destinations"
	fieldNameMap["filter_type"] = "FilterType"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["services"] = "Services"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["sources"] = "Sources"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceInsertionRulesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceInsertionRuleListResultBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceInsertionRulesRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = bindings.NewStringType()
	fields["applied_tos"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["destinations"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["filter_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["services"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sources"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["applied_tos"] = "AppliedTos"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["destinations"] = "Destinations"
	fieldNameMap["filter_type"] = "FilterType"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["services"] = "Services"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["sources"] = "Sources"
	paramsTypeMap["sources"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["filter_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["section_id"] = bindings.NewStringType()
	paramsTypeMap["applied_tos"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["services"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["destinations"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["sectionId"] = bindings.NewStringType()
	pathParams["section_id"] = "sectionId"
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["sources"] = "sources"
	queryParams["filter_type"] = "filter_type"
	queryParams["included_fields"] = "included_fields"
	queryParams["destinations"] = "destinations"
	queryParams["applied_tos"] = "applied_tos"
	queryParams["services"] = "services"
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
		"/api/v1/serviceinsertion/sections/{sectionId}/rules",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceInsertionSectionInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fieldNameMap["section_id"] = "SectionId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceInsertionSectionOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceInsertionSectionBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceInsertionSectionRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = bindings.NewStringType()
	fieldNameMap["section_id"] = "SectionId"
	paramsTypeMap["section_id"] = bindings.NewStringType()
	paramsTypeMap["sectionId"] = bindings.NewStringType()
	pathParams["section_id"] = "sectionId"
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
		"/api/v1/serviceinsertion/sections/{sectionId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceInsertionSectionWithRulesListWithRulesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fieldNameMap["section_id"] = "SectionId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceInsertionSectionWithRulesListWithRulesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceInsertionSectionRuleListBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceInsertionSectionWithRulesListWithRulesRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = bindings.NewStringType()
	fieldNameMap["section_id"] = "SectionId"
	paramsTypeMap["section_id"] = bindings.NewStringType()
	paramsTypeMap["sectionId"] = bindings.NewStringType()
	pathParams["section_id"] = "sectionId"
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
		"action=list_with_rules",
		"",
		"POST",
		"/api/v1/serviceinsertion/sections/{sectionId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceInsertionServiceInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceInsertionServiceOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceDefinitionBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceInsertionServiceRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	pathParams["service_id"] = "serviceId"
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
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceInsertionStatusInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["context_type"] = bindings.NewStringType()
	fieldNameMap["context_type"] = "ContextType"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceInsertionStatusOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceInsertionStatusBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceInsertionStatusRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["context_type"] = bindings.NewStringType()
	fieldNameMap["context_type"] = "ContextType"
	paramsTypeMap["context_type"] = bindings.NewStringType()
	paramsTypeMap["contextType"] = bindings.NewStringType()
	pathParams["context_type"] = "contextType"
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
		"/api/v1/serviceinsertion/status/{contextType}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceInstanceInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceInstanceOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.BaseServiceInstanceBindingType)}, bindings.REST)
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceInstanceRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["service_instance_id"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	paramsTypeMap["serviceInstanceId"] = bindings.NewStringType()
	pathParams["service_instance_id"] = "serviceInstanceId"
	pathParams["service_id"] = "serviceId"
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
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/service-instances/{serviceInstanceId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceInstanceNSGroupsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceInstanceNSGroupsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceInstanceNSGroupsBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceInstanceNSGroupsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["service_instance_id"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	paramsTypeMap["serviceInstanceId"] = bindings.NewStringType()
	pathParams["service_instance_id"] = "serviceInstanceId"
	pathParams["service_id"] = "serviceId"
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
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/service-instances/{serviceInstanceId}/group-associations",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceInstanceStateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceInstanceStateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ConfigurationStateBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceInstanceStateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["service_instance_id"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	paramsTypeMap["serviceInstanceId"] = bindings.NewStringType()
	pathParams["service_instance_id"] = "serviceInstanceId"
	pathParams["service_id"] = "serviceId"
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
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/service-instances/{serviceInstanceId}/state",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceInstanceStatusInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	fieldNameMap["source"] = "Source"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceInstanceStatusOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceInstanceStatusBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceInstanceStatusRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	fieldNameMap["source"] = "Source"
	paramsTypeMap["source"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["service_instance_id"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	paramsTypeMap["serviceInstanceId"] = bindings.NewStringType()
	pathParams["service_instance_id"] = "serviceInstanceId"
	pathParams["service_id"] = "serviceId"
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
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/service-instances/{serviceInstanceId}/status",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceManagerInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_manager_id"] = bindings.NewStringType()
	fieldNameMap["service_manager_id"] = "ServiceManagerId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceManagerOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceManagerBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceManagerRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_manager_id"] = bindings.NewStringType()
	fieldNameMap["service_manager_id"] = "ServiceManagerId"
	paramsTypeMap["service_manager_id"] = bindings.NewStringType()
	paramsTypeMap["serviceManagerId"] = bindings.NewStringType()
	pathParams["service_manager_id"] = "serviceManagerId"
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
		"/api/v1/serviceinsertion/service-managers/{serviceManagerId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceProfileNSGroupsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["service_profile_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_profile_id"] = "ServiceProfileId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceProfileNSGroupsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceProfileNSGroupsBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionGetServiceProfileNSGroupsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["service_profile_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_profile_id"] = "ServiceProfileId"
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["service_profile_id"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	paramsTypeMap["serviceProfileId"] = bindings.NewStringType()
	pathParams["service_profile_id"] = "serviceProfileId"
	pathParams["service_id"] = "serviceId"
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
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/service-profiles/{serviceProfileId}/nsgroups",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionGetSolutionConfigInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["solution_config_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["solution_config_id"] = "SolutionConfigId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionGetSolutionConfigOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.SolutionConfigBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionGetSolutionConfigRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["solution_config_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["solution_config_id"] = "SolutionConfigId"
	paramsTypeMap["solution_config_id"] = bindings.NewStringType()
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	paramsTypeMap["solutionConfigId"] = bindings.NewStringType()
	pathParams["solution_config_id"] = "solutionConfigId"
	pathParams["service_id"] = "serviceId"
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
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/solution-configs/{solutionConfigId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionGetVendorTemplateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["vendor_template_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["vendor_template_id"] = "VendorTemplateId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionGetVendorTemplateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.VendorTemplateBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionGetVendorTemplateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["vendor_template_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["vendor_template_id"] = "VendorTemplateId"
	paramsTypeMap["vendor_template_id"] = bindings.NewStringType()
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	paramsTypeMap["vendorTemplateId"] = bindings.NewStringType()
	pathParams["service_id"] = "serviceId"
	pathParams["vendor_template_id"] = "vendorTemplateId"
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
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/vendor-templates/{vendorTemplateId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionListInstanceEndpointsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionListInstanceEndpointsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.InstanceEndpointListResultBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionListInstanceEndpointsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["service_instance_id"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	paramsTypeMap["serviceInstanceId"] = bindings.NewStringType()
	pathParams["service_instance_id"] = "serviceInstanceId"
	pathParams["service_id"] = "serviceId"
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
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/service-instances/{serviceInstanceId}/instance-endpoints",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionListInstanceRuntimesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionListInstanceRuntimesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.InstanceRuntimeListResultBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionListInstanceRuntimesRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["service_instance_id"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	paramsTypeMap["serviceInstanceId"] = bindings.NewStringType()
	pathParams["service_instance_id"] = "serviceInstanceId"
	pathParams["service_id"] = "serviceId"
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
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/service-instances/{serviceInstanceId}/instance-runtimes",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionListSIServiceProfilesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionListSIServiceProfilesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.SIServiceProfileListResultBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionListSIServiceProfilesRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	pathParams["service_id"] = "serviceId"
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
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/service-profiles",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionListServiceAttachmentsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionListServiceAttachmentsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceAttachmentListResultBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionListServiceAttachmentsRestMetadata() protocol.OperationRestMetadata {
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
		"/api/v1/serviceinsertion/service-attachments",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionListServiceChainMappingsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["service_profile_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_profile_id"] = "ServiceProfileId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionListServiceChainMappingsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceChainMappingListResultBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionListServiceChainMappingsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["service_profile_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_profile_id"] = "ServiceProfileId"
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["service_profile_id"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	paramsTypeMap["serviceProfileId"] = bindings.NewStringType()
	pathParams["service_profile_id"] = "serviceProfileId"
	pathParams["service_id"] = "serviceId"
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
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/service-profiles/{serviceProfileId}/service-chain-mappings",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionListServiceChainsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionListServiceChainsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceChainListResultBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionListServiceChainsRestMetadata() protocol.OperationRestMetadata {
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
		"/api/v1/serviceinsertion/service-chains",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionListServiceInsertionSectionsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["applied_tos"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["destinations"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["exclude_applied_to_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["filter_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_applied_to_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["services"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sources"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["applied_tos"] = "AppliedTos"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["destinations"] = "Destinations"
	fieldNameMap["exclude_applied_to_type"] = "ExcludeAppliedToType"
	fieldNameMap["filter_type"] = "FilterType"
	fieldNameMap["include_applied_to_type"] = "IncludeAppliedToType"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["services"] = "Services"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["sources"] = "Sources"
	fieldNameMap["type"] = "Type_"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionListServiceInsertionSectionsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceInsertionSectionListResultBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionListServiceInsertionSectionsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["applied_tos"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["destinations"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["exclude_applied_to_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["filter_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_applied_to_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["services"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sources"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["applied_tos"] = "AppliedTos"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["destinations"] = "Destinations"
	fieldNameMap["exclude_applied_to_type"] = "ExcludeAppliedToType"
	fieldNameMap["filter_type"] = "FilterType"
	fieldNameMap["include_applied_to_type"] = "IncludeAppliedToType"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["services"] = "Services"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["sources"] = "Sources"
	fieldNameMap["type"] = "Type_"
	paramsTypeMap["sources"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["exclude_applied_to_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["include_applied_to_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["destinations"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["filter_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["applied_tos"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["services"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["sources"] = "sources"
	queryParams["destinations"] = "destinations"
	queryParams["exclude_applied_to_type"] = "exclude_applied_to_type"
	queryParams["include_applied_to_type"] = "include_applied_to_type"
	queryParams["services"] = "services"
	queryParams["sort_by"] = "sort_by"
	queryParams["type"] = "type"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["filter_type"] = "filter_type"
	queryParams["included_fields"] = "included_fields"
	queryParams["applied_tos"] = "applied_tos"
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
		"/api/v1/serviceinsertion/sections",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionListServiceInsertionServicesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionListServiceInsertionServicesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceInsertionServiceListResultBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionListServiceInsertionServicesRestMetadata() protocol.OperationRestMetadata {
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
		"/api/v1/serviceinsertion/services/nsxt-mp",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionListServiceInsertionStatusInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionListServiceInsertionStatusOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceInsertionStatusListResultBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionListServiceInsertionStatusRestMetadata() protocol.OperationRestMetadata {
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
		"/api/v1/serviceinsertion/status",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionListServiceInstancesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["deployed_to"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["service_deployment_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["deployed_to"] = "DeployedTo"
	fieldNameMap["service_deployment_id"] = "ServiceDeploymentId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionListServiceInstancesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceInstanceListResultBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionListServiceInstancesRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["deployed_to"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["service_deployment_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["deployed_to"] = "DeployedTo"
	fieldNameMap["service_deployment_id"] = "ServiceDeploymentId"
	paramsTypeMap["deployed_to"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["service_deployment_id"] = bindings.NewOptionalType(bindings.NewStringType())
	queryParams["service_deployment_id"] = "service_deployment_id"
	queryParams["deployed_to"] = "deployed_to"
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
		"/api/v1/serviceinsertion/service-instances",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionListServiceInstancesForServiceInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionListServiceInstancesForServiceOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceInstanceListResultBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionListServiceInstancesForServiceRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	pathParams["service_id"] = "serviceId"
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
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/service-instances",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionListServiceManagersInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionListServiceManagersOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceManagerListResultBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionListServiceManagersRestMetadata() protocol.OperationRestMetadata {
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
		"/api/v1/serviceinsertion/service-managers",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionListServicePathsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_chain_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["service_chain_id"] = "ServiceChainId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionListServicePathsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServicePathListResultBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionListServicePathsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_chain_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["service_chain_id"] = "ServiceChainId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["service_chain_id"] = bindings.NewStringType()
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["serviceChainId"] = bindings.NewStringType()
	pathParams["service_chain_id"] = "serviceChainId"
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
		"/api/v1/serviceinsertion/service-chains/{serviceChainId}/service-paths",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionListSolutionConfigsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionListSolutionConfigsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.SolutionConfigListResultBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionListSolutionConfigsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	pathParams["service_id"] = "serviceId"
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
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/solution-configs",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionListVendorTemplatesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["vendor_template_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["vendor_template_name"] = "VendorTemplateName"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionListVendorTemplatesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.VendorTemplateListResultBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionListVendorTemplatesRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["vendor_template_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["vendor_template_name"] = "VendorTemplateName"
	paramsTypeMap["vendor_template_name"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	pathParams["service_id"] = "serviceId"
	queryParams["vendor_template_name"] = "vendor_template_name"
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
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/vendor-templates",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionRegisterServiceManagerInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_manager"] = bindings.NewReferenceType(model.ServiceManagerBindingType)
	fieldNameMap["service_manager"] = "ServiceManager"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionRegisterServiceManagerOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceManagerBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionRegisterServiceManagerRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_manager"] = bindings.NewReferenceType(model.ServiceManagerBindingType)
	fieldNameMap["service_manager"] = "ServiceManager"
	paramsTypeMap["service_manager"] = bindings.NewReferenceType(model.ServiceManagerBindingType)
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
		"service_manager",
		"POST",
		"/api/v1/serviceinsertion/service-managers",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionRemoveServiceInsertionExcludeListMemberRemoveMemberInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["object_id"] = bindings.NewStringType()
	fieldNameMap["object_id"] = "ObjectId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionRemoveServiceInsertionExcludeListMemberRemoveMemberOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ResourceReferenceBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionRemoveServiceInsertionExcludeListMemberRemoveMemberRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["object_id"] = bindings.NewStringType()
	fieldNameMap["object_id"] = "ObjectId"
	paramsTypeMap["object_id"] = bindings.NewStringType()
	queryParams["object_id"] = "object_id"
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
		"action=remove_member",
		"",
		"POST",
		"/api/v1/serviceinsertion/excludelist",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionResolveSourceEntitiesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["source_node_value"] = bindings.NewStringType()
	fieldNameMap["source_node_value"] = "SourceNodeValue"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionResolveSourceEntitiesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.SourceEntityResultBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionResolveSourceEntitiesRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["source_node_value"] = bindings.NewStringType()
	fieldNameMap["source_node_value"] = "SourceNodeValue"
	paramsTypeMap["source_node_value"] = bindings.NewStringType()
	queryParams["source_node_value"] = "source_node_value"
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
		"/api/v1/serviceinsertion/source-entities",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionReviseServiceInsertionRuleReviseInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["rule_id"] = bindings.NewStringType()
	fields["service_insertion_rule"] = bindings.NewReferenceType(model.ServiceInsertionRuleBindingType)
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["rule_id"] = "RuleId"
	fieldNameMap["service_insertion_rule"] = "ServiceInsertionRule"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionReviseServiceInsertionRuleReviseOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceInsertionRuleBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionReviseServiceInsertionRuleReviseRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = bindings.NewStringType()
	fields["rule_id"] = bindings.NewStringType()
	fields["service_insertion_rule"] = bindings.NewReferenceType(model.ServiceInsertionRuleBindingType)
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["rule_id"] = "RuleId"
	fieldNameMap["service_insertion_rule"] = "ServiceInsertionRule"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	paramsTypeMap["rule_id"] = bindings.NewStringType()
	paramsTypeMap["id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["section_id"] = bindings.NewStringType()
	paramsTypeMap["service_insertion_rule"] = bindings.NewReferenceType(model.ServiceInsertionRuleBindingType)
	paramsTypeMap["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sectionId"] = bindings.NewStringType()
	paramsTypeMap["ruleId"] = bindings.NewStringType()
	pathParams["section_id"] = "sectionId"
	pathParams["rule_id"] = "ruleId"
	queryParams["id"] = "id"
	queryParams["operation"] = "operation"
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
		"action=revise",
		"service_insertion_rule",
		"POST",
		"/api/v1/serviceinsertion/sections/{sectionId}/rules/{ruleId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionReviseServiceInsertionSectionReviseInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["service_insertion_section"] = bindings.NewReferenceType(model.ServiceInsertionSectionBindingType)
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["service_insertion_section"] = "ServiceInsertionSection"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionReviseServiceInsertionSectionReviseOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceInsertionSectionBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionReviseServiceInsertionSectionReviseRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = bindings.NewStringType()
	fields["service_insertion_section"] = bindings.NewReferenceType(model.ServiceInsertionSectionBindingType)
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["service_insertion_section"] = "ServiceInsertionSection"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	paramsTypeMap["id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["section_id"] = bindings.NewStringType()
	paramsTypeMap["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["service_insertion_section"] = bindings.NewReferenceType(model.ServiceInsertionSectionBindingType)
	paramsTypeMap["sectionId"] = bindings.NewStringType()
	pathParams["section_id"] = "sectionId"
	queryParams["id"] = "id"
	queryParams["operation"] = "operation"
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
		"action=revise",
		"service_insertion_section",
		"POST",
		"/api/v1/serviceinsertion/sections/{sectionId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionReviseServiceInsertionSectionWithRulesReviseWithRulesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["service_insertion_section_rule_list"] = bindings.NewReferenceType(model.ServiceInsertionSectionRuleListBindingType)
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["service_insertion_section_rule_list"] = "ServiceInsertionSectionRuleList"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionReviseServiceInsertionSectionWithRulesReviseWithRulesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceInsertionSectionRuleListBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionReviseServiceInsertionSectionWithRulesReviseWithRulesRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = bindings.NewStringType()
	fields["service_insertion_section_rule_list"] = bindings.NewReferenceType(model.ServiceInsertionSectionRuleListBindingType)
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["service_insertion_section_rule_list"] = "ServiceInsertionSectionRuleList"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	paramsTypeMap["id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["section_id"] = bindings.NewStringType()
	paramsTypeMap["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["service_insertion_section_rule_list"] = bindings.NewReferenceType(model.ServiceInsertionSectionRuleListBindingType)
	paramsTypeMap["sectionId"] = bindings.NewStringType()
	pathParams["section_id"] = "sectionId"
	queryParams["id"] = "id"
	queryParams["operation"] = "operation"
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
		"action=revise_with_rules",
		"service_insertion_section_rule_list",
		"POST",
		"/api/v1/serviceinsertion/sections/{sectionId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionUpdateServiceAttachmentInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_attachment_id"] = bindings.NewStringType()
	fields["service_attachment"] = bindings.NewReferenceType(model.ServiceAttachmentBindingType)
	fieldNameMap["service_attachment_id"] = "ServiceAttachmentId"
	fieldNameMap["service_attachment"] = "ServiceAttachment"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionUpdateServiceAttachmentOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceAttachmentBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionUpdateServiceAttachmentRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_attachment_id"] = bindings.NewStringType()
	fields["service_attachment"] = bindings.NewReferenceType(model.ServiceAttachmentBindingType)
	fieldNameMap["service_attachment_id"] = "ServiceAttachmentId"
	fieldNameMap["service_attachment"] = "ServiceAttachment"
	paramsTypeMap["service_attachment"] = bindings.NewReferenceType(model.ServiceAttachmentBindingType)
	paramsTypeMap["service_attachment_id"] = bindings.NewStringType()
	paramsTypeMap["serviceAttachmentId"] = bindings.NewStringType()
	pathParams["service_attachment_id"] = "serviceAttachmentId"
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
		"service_attachment",
		"PUT",
		"/api/v1/serviceinsertion/service-attachments/{serviceAttachmentId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionUpdateServiceDeploymentInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["service_deployment_id"] = bindings.NewStringType()
	fields["service_deployment"] = bindings.NewReferenceType(model.ServiceDeploymentBindingType)
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_deployment_id"] = "ServiceDeploymentId"
	fieldNameMap["service_deployment"] = "ServiceDeployment"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionUpdateServiceDeploymentOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceDeploymentBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionUpdateServiceDeploymentRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["service_deployment_id"] = bindings.NewStringType()
	fields["service_deployment"] = bindings.NewReferenceType(model.ServiceDeploymentBindingType)
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_deployment_id"] = "ServiceDeploymentId"
	fieldNameMap["service_deployment"] = "ServiceDeployment"
	paramsTypeMap["service_deployment"] = bindings.NewReferenceType(model.ServiceDeploymentBindingType)
	paramsTypeMap["service_deployment_id"] = bindings.NewStringType()
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	paramsTypeMap["serviceDeploymentId"] = bindings.NewStringType()
	pathParams["service_deployment_id"] = "serviceDeploymentId"
	pathParams["service_id"] = "serviceId"
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
		"service_deployment",
		"PUT",
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/service-deployments/{serviceDeploymentId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionUpdateServiceInsertionExcludeListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["si_exclude_list"] = bindings.NewReferenceType(model.SIExcludeListBindingType)
	fieldNameMap["si_exclude_list"] = "SiExcludeList"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionUpdateServiceInsertionExcludeListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.SIExcludeListBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionUpdateServiceInsertionExcludeListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["si_exclude_list"] = bindings.NewReferenceType(model.SIExcludeListBindingType)
	fieldNameMap["si_exclude_list"] = "SiExcludeList"
	paramsTypeMap["si_exclude_list"] = bindings.NewReferenceType(model.SIExcludeListBindingType)
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
		"si_exclude_list",
		"PUT",
		"/api/v1/serviceinsertion/excludelist",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionUpdateServiceInsertionRuleInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["rule_id"] = bindings.NewStringType()
	fields["service_insertion_rule"] = bindings.NewReferenceType(model.ServiceInsertionRuleBindingType)
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["rule_id"] = "RuleId"
	fieldNameMap["service_insertion_rule"] = "ServiceInsertionRule"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionUpdateServiceInsertionRuleOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceInsertionRuleBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionUpdateServiceInsertionRuleRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = bindings.NewStringType()
	fields["rule_id"] = bindings.NewStringType()
	fields["service_insertion_rule"] = bindings.NewReferenceType(model.ServiceInsertionRuleBindingType)
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["rule_id"] = "RuleId"
	fieldNameMap["service_insertion_rule"] = "ServiceInsertionRule"
	paramsTypeMap["rule_id"] = bindings.NewStringType()
	paramsTypeMap["section_id"] = bindings.NewStringType()
	paramsTypeMap["service_insertion_rule"] = bindings.NewReferenceType(model.ServiceInsertionRuleBindingType)
	paramsTypeMap["sectionId"] = bindings.NewStringType()
	paramsTypeMap["ruleId"] = bindings.NewStringType()
	pathParams["section_id"] = "sectionId"
	pathParams["rule_id"] = "ruleId"
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
		"service_insertion_rule",
		"PUT",
		"/api/v1/serviceinsertion/sections/{sectionId}/rules/{ruleId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionUpdateServiceInsertionSectionInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["service_insertion_section"] = bindings.NewReferenceType(model.ServiceInsertionSectionBindingType)
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["service_insertion_section"] = "ServiceInsertionSection"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionUpdateServiceInsertionSectionOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceInsertionSectionBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionUpdateServiceInsertionSectionRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = bindings.NewStringType()
	fields["service_insertion_section"] = bindings.NewReferenceType(model.ServiceInsertionSectionBindingType)
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["service_insertion_section"] = "ServiceInsertionSection"
	paramsTypeMap["section_id"] = bindings.NewStringType()
	paramsTypeMap["service_insertion_section"] = bindings.NewReferenceType(model.ServiceInsertionSectionBindingType)
	paramsTypeMap["sectionId"] = bindings.NewStringType()
	pathParams["section_id"] = "sectionId"
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
		"service_insertion_section",
		"PUT",
		"/api/v1/serviceinsertion/sections/{sectionId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionUpdateServiceInsertionSectionWithRulesUpdateWithRulesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["service_insertion_section_rule_list"] = bindings.NewReferenceType(model.ServiceInsertionSectionRuleListBindingType)
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["service_insertion_section_rule_list"] = "ServiceInsertionSectionRuleList"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionUpdateServiceInsertionSectionWithRulesUpdateWithRulesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceInsertionSectionRuleListBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionUpdateServiceInsertionSectionWithRulesUpdateWithRulesRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = bindings.NewStringType()
	fields["service_insertion_section_rule_list"] = bindings.NewReferenceType(model.ServiceInsertionSectionRuleListBindingType)
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["service_insertion_section_rule_list"] = "ServiceInsertionSectionRuleList"
	paramsTypeMap["section_id"] = bindings.NewStringType()
	paramsTypeMap["service_insertion_section_rule_list"] = bindings.NewReferenceType(model.ServiceInsertionSectionRuleListBindingType)
	paramsTypeMap["sectionId"] = bindings.NewStringType()
	pathParams["section_id"] = "sectionId"
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
		"action=update_with_rules",
		"service_insertion_section_rule_list",
		"POST",
		"/api/v1/serviceinsertion/sections/{sectionId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionUpdateServiceInsertionServiceInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["service_definition"] = bindings.NewReferenceType(model.ServiceDefinitionBindingType)
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_definition"] = "ServiceDefinition"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionUpdateServiceInsertionServiceOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceDefinitionBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionUpdateServiceInsertionServiceRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["service_definition"] = bindings.NewReferenceType(model.ServiceDefinitionBindingType)
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_definition"] = "ServiceDefinition"
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["service_definition"] = bindings.NewReferenceType(model.ServiceDefinitionBindingType)
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	pathParams["service_id"] = "serviceId"
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
		"service_definition",
		"PUT",
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionUpdateServiceInsertionStatusInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["context_type"] = bindings.NewStringType()
	fields["service_insertion_status"] = bindings.NewReferenceType(model.ServiceInsertionStatusBindingType)
	fieldNameMap["context_type"] = "ContextType"
	fieldNameMap["service_insertion_status"] = "ServiceInsertionStatus"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionUpdateServiceInsertionStatusOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceInsertionStatusBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionUpdateServiceInsertionStatusRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["context_type"] = bindings.NewStringType()
	fields["service_insertion_status"] = bindings.NewReferenceType(model.ServiceInsertionStatusBindingType)
	fieldNameMap["context_type"] = "ContextType"
	fieldNameMap["service_insertion_status"] = "ServiceInsertionStatus"
	paramsTypeMap["service_insertion_status"] = bindings.NewReferenceType(model.ServiceInsertionStatusBindingType)
	paramsTypeMap["context_type"] = bindings.NewStringType()
	paramsTypeMap["contextType"] = bindings.NewStringType()
	pathParams["context_type"] = "contextType"
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
		"service_insertion_status",
		"PUT",
		"/api/v1/serviceinsertion/status/{contextType}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionUpdateServiceInstanceInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fields["base_service_instance"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.BaseServiceInstanceBindingType)}, bindings.REST)
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	fieldNameMap["base_service_instance"] = "BaseServiceInstance"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionUpdateServiceInstanceOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.BaseServiceInstanceBindingType)}, bindings.REST)
}

func managementPlaneAPISecurityServicesServiceInsertionUpdateServiceInstanceRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fields["base_service_instance"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.BaseServiceInstanceBindingType)}, bindings.REST)
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	fieldNameMap["base_service_instance"] = "BaseServiceInstance"
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["base_service_instance"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.BaseServiceInstanceBindingType)}, bindings.REST)
	paramsTypeMap["service_instance_id"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	paramsTypeMap["serviceInstanceId"] = bindings.NewStringType()
	pathParams["service_instance_id"] = "serviceInstanceId"
	pathParams["service_id"] = "serviceId"
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
		"base_service_instance",
		"PUT",
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/service-instances/{serviceInstanceId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionUpdateServiceManagerInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_manager_id"] = bindings.NewStringType()
	fields["service_manager"] = bindings.NewReferenceType(model.ServiceManagerBindingType)
	fieldNameMap["service_manager_id"] = "ServiceManagerId"
	fieldNameMap["service_manager"] = "ServiceManager"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionUpdateServiceManagerOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceManagerBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionUpdateServiceManagerRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_manager_id"] = bindings.NewStringType()
	fields["service_manager"] = bindings.NewReferenceType(model.ServiceManagerBindingType)
	fieldNameMap["service_manager_id"] = "ServiceManagerId"
	fieldNameMap["service_manager"] = "ServiceManager"
	paramsTypeMap["service_manager"] = bindings.NewReferenceType(model.ServiceManagerBindingType)
	paramsTypeMap["service_manager_id"] = bindings.NewStringType()
	paramsTypeMap["serviceManagerId"] = bindings.NewStringType()
	pathParams["service_manager_id"] = "serviceManagerId"
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
		"service_manager",
		"PUT",
		"/api/v1/serviceinsertion/service-managers/{serviceManagerId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionUpdateServiceVmStateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fields["instance_runtime_id"] = bindings.NewStringType()
	fields["action"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["unhealthy_reason"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	fieldNameMap["instance_runtime_id"] = "InstanceRuntimeId"
	fieldNameMap["action"] = "Action"
	fieldNameMap["unhealthy_reason"] = "UnhealthyReason"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionUpdateServiceVmStateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPISecurityServicesServiceInsertionUpdateServiceVmStateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fields["instance_runtime_id"] = bindings.NewStringType()
	fields["action"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["unhealthy_reason"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	fieldNameMap["instance_runtime_id"] = "InstanceRuntimeId"
	fieldNameMap["action"] = "Action"
	fieldNameMap["unhealthy_reason"] = "UnhealthyReason"
	paramsTypeMap["action"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["unhealthy_reason"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["service_instance_id"] = bindings.NewStringType()
	paramsTypeMap["instance_runtime_id"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	paramsTypeMap["serviceInstanceId"] = bindings.NewStringType()
	paramsTypeMap["instanceRuntimeId"] = bindings.NewStringType()
	pathParams["instance_runtime_id"] = "instanceRuntimeId"
	pathParams["service_instance_id"] = "serviceInstanceId"
	pathParams["service_id"] = "serviceId"
	queryParams["action"] = "action"
	queryParams["unhealthy_reason"] = "unhealthy_reason"
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
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/service-instances/{serviceInstanceId}/instance-runtimes/{instanceRuntimeId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionUpdateSolutionConfigInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["solution_config_id"] = bindings.NewStringType()
	fields["solution_config"] = bindings.NewReferenceType(model.SolutionConfigBindingType)
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["solution_config_id"] = "SolutionConfigId"
	fieldNameMap["solution_config"] = "SolutionConfig"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionUpdateSolutionConfigOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.SolutionConfigBindingType)
}

func managementPlaneAPISecurityServicesServiceInsertionUpdateSolutionConfigRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["solution_config_id"] = bindings.NewStringType()
	fields["solution_config"] = bindings.NewReferenceType(model.SolutionConfigBindingType)
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["solution_config_id"] = "SolutionConfigId"
	fieldNameMap["solution_config"] = "SolutionConfig"
	paramsTypeMap["solution_config"] = bindings.NewReferenceType(model.SolutionConfigBindingType)
	paramsTypeMap["solution_config_id"] = bindings.NewStringType()
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	paramsTypeMap["solutionConfigId"] = bindings.NewStringType()
	pathParams["solution_config_id"] = "solutionConfigId"
	pathParams["service_id"] = "serviceId"
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
		"solution_config",
		"PUT",
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/solution-configs/{solutionConfigId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionUpgradeServiceDeploymentUpgradeInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["service_deployment_id"] = bindings.NewStringType()
	fields["deployment_spec_name"] = bindings.NewReferenceType(model.DeploymentSpecNameBindingType)
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_deployment_id"] = "ServiceDeploymentId"
	fieldNameMap["deployment_spec_name"] = "DeploymentSpecName"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionUpgradeServiceDeploymentUpgradeOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPISecurityServicesServiceInsertionUpgradeServiceDeploymentUpgradeRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["service_deployment_id"] = bindings.NewStringType()
	fields["deployment_spec_name"] = bindings.NewReferenceType(model.DeploymentSpecNameBindingType)
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_deployment_id"] = "ServiceDeploymentId"
	fieldNameMap["deployment_spec_name"] = "DeploymentSpecName"
	paramsTypeMap["service_deployment_id"] = bindings.NewStringType()
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["deployment_spec_name"] = bindings.NewReferenceType(model.DeploymentSpecNameBindingType)
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	paramsTypeMap["serviceDeploymentId"] = bindings.NewStringType()
	pathParams["service_deployment_id"] = "serviceDeploymentId"
	pathParams["service_id"] = "serviceId"
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
		"action=upgrade",
		"deployment_spec_name",
		"POST",
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/service-deployments/{serviceDeploymentId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesServiceInsertionUpgradeServicevMsUpgradeInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesServiceInsertionUpgradeServicevMsUpgradeOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPISecurityServicesServiceInsertionUpgradeServicevMsUpgradeRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["service_instance_id"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	paramsTypeMap["serviceInstanceId"] = bindings.NewStringType()
	pathParams["service_instance_id"] = "serviceInstanceId"
	pathParams["service_id"] = "serviceId"
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
		"action=upgrade",
		"",
		"POST",
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/service-instances/{serviceInstanceId}/instance-runtimes",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
