// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: EventLogServers.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package identity_firewall_stores

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func eventLogServersDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["identity_firewall_store_id"] = vapiBindings_.NewStringType()
	fields["event_log_server_id"] = vapiBindings_.NewStringType()
	fields["enforcement_point_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["identity_firewall_store_id"] = "IdentityFirewallStoreId"
	fieldNameMap["event_log_server_id"] = "EventLogServerId"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func EventLogServersDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func eventLogServersDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["identity_firewall_store_id"] = vapiBindings_.NewStringType()
	fields["event_log_server_id"] = vapiBindings_.NewStringType()
	fields["enforcement_point_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["identity_firewall_store_id"] = "IdentityFirewallStoreId"
	fieldNameMap["event_log_server_id"] = "EventLogServerId"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	paramsTypeMap["event_log_server_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcement_point_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["identity_firewall_store_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["identityFirewallStoreId"] = vapiBindings_.NewStringType()
	paramsTypeMap["eventLogServerId"] = vapiBindings_.NewStringType()
	pathParams["identity_firewall_store_id"] = "identityFirewallStoreId"
	pathParams["event_log_server_id"] = "eventLogServerId"
	queryParams["enforcement_point_path"] = "enforcement_point_path"
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
		"DELETE",
		"/policy/api/v1/infra/identity-firewall-stores/{identityFirewallStoreId}/event-log-servers/{eventLogServerId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func eventLogServersGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["identity_firewall_store_id"] = vapiBindings_.NewStringType()
	fields["event_log_server_id"] = vapiBindings_.NewStringType()
	fields["enforcement_point_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["identity_firewall_store_id"] = "IdentityFirewallStoreId"
	fieldNameMap["event_log_server_id"] = "EventLogServerId"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func EventLogServersGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.IdentityFirewallStoreEventLogServerBindingType)
}

func eventLogServersGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["identity_firewall_store_id"] = vapiBindings_.NewStringType()
	fields["event_log_server_id"] = vapiBindings_.NewStringType()
	fields["enforcement_point_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["identity_firewall_store_id"] = "IdentityFirewallStoreId"
	fieldNameMap["event_log_server_id"] = "EventLogServerId"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	paramsTypeMap["event_log_server_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcement_point_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["identity_firewall_store_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["identityFirewallStoreId"] = vapiBindings_.NewStringType()
	paramsTypeMap["eventLogServerId"] = vapiBindings_.NewStringType()
	pathParams["identity_firewall_store_id"] = "identityFirewallStoreId"
	pathParams["event_log_server_id"] = "eventLogServerId"
	queryParams["enforcement_point_path"] = "enforcement_point_path"
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
		"/policy/api/v1/infra/identity-firewall-stores/{identityFirewallStoreId}/event-log-servers/{eventLogServerId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func eventLogServersPatchInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["identity_firewall_store_id"] = vapiBindings_.NewStringType()
	fields["event_log_server_id"] = vapiBindings_.NewStringType()
	fields["identity_firewall_store_event_log_server"] = vapiBindings_.NewReferenceType(nsx_policyModel.IdentityFirewallStoreEventLogServerBindingType)
	fields["enforcement_point_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["identity_firewall_store_id"] = "IdentityFirewallStoreId"
	fieldNameMap["event_log_server_id"] = "EventLogServerId"
	fieldNameMap["identity_firewall_store_event_log_server"] = "IdentityFirewallStoreEventLogServer"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func EventLogServersPatchOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func eventLogServersPatchRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["identity_firewall_store_id"] = vapiBindings_.NewStringType()
	fields["event_log_server_id"] = vapiBindings_.NewStringType()
	fields["identity_firewall_store_event_log_server"] = vapiBindings_.NewReferenceType(nsx_policyModel.IdentityFirewallStoreEventLogServerBindingType)
	fields["enforcement_point_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["identity_firewall_store_id"] = "IdentityFirewallStoreId"
	fieldNameMap["event_log_server_id"] = "EventLogServerId"
	fieldNameMap["identity_firewall_store_event_log_server"] = "IdentityFirewallStoreEventLogServer"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	paramsTypeMap["event_log_server_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcement_point_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["identity_firewall_store_event_log_server"] = vapiBindings_.NewReferenceType(nsx_policyModel.IdentityFirewallStoreEventLogServerBindingType)
	paramsTypeMap["identity_firewall_store_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["identityFirewallStoreId"] = vapiBindings_.NewStringType()
	paramsTypeMap["eventLogServerId"] = vapiBindings_.NewStringType()
	pathParams["identity_firewall_store_id"] = "identityFirewallStoreId"
	pathParams["event_log_server_id"] = "eventLogServerId"
	queryParams["enforcement_point_path"] = "enforcement_point_path"
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
		"identity_firewall_store_event_log_server",
		"PATCH",
		"/policy/api/v1/infra/identity-firewall-stores/{identityFirewallStoreId}/event-log-servers/{eventLogServerId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func eventLogServersUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["identity_firewall_store_id"] = vapiBindings_.NewStringType()
	fields["event_log_server_id"] = vapiBindings_.NewStringType()
	fields["identity_firewall_store_event_log_server"] = vapiBindings_.NewReferenceType(nsx_policyModel.IdentityFirewallStoreEventLogServerBindingType)
	fields["enforcement_point_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["identity_firewall_store_id"] = "IdentityFirewallStoreId"
	fieldNameMap["event_log_server_id"] = "EventLogServerId"
	fieldNameMap["identity_firewall_store_event_log_server"] = "IdentityFirewallStoreEventLogServer"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func EventLogServersUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.IdentityFirewallStoreEventLogServerBindingType)
}

func eventLogServersUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["identity_firewall_store_id"] = vapiBindings_.NewStringType()
	fields["event_log_server_id"] = vapiBindings_.NewStringType()
	fields["identity_firewall_store_event_log_server"] = vapiBindings_.NewReferenceType(nsx_policyModel.IdentityFirewallStoreEventLogServerBindingType)
	fields["enforcement_point_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["identity_firewall_store_id"] = "IdentityFirewallStoreId"
	fieldNameMap["event_log_server_id"] = "EventLogServerId"
	fieldNameMap["identity_firewall_store_event_log_server"] = "IdentityFirewallStoreEventLogServer"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	paramsTypeMap["event_log_server_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcement_point_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["identity_firewall_store_event_log_server"] = vapiBindings_.NewReferenceType(nsx_policyModel.IdentityFirewallStoreEventLogServerBindingType)
	paramsTypeMap["identity_firewall_store_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["identityFirewallStoreId"] = vapiBindings_.NewStringType()
	paramsTypeMap["eventLogServerId"] = vapiBindings_.NewStringType()
	pathParams["identity_firewall_store_id"] = "identityFirewallStoreId"
	pathParams["event_log_server_id"] = "eventLogServerId"
	queryParams["enforcement_point_path"] = "enforcement_point_path"
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
		"identity_firewall_store_event_log_server",
		"PUT",
		"/policy/api/v1/infra/identity-firewall-stores/{identityFirewallStoreId}/event-log-servers/{eventLogServerId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
