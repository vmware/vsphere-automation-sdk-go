// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: ServiceInterfaces.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package locale_services

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func serviceInterfacesDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = bindings.NewStringType()
	fields["deployment_group_id"] = bindings.NewStringType()
	fields["vhc_id"] = bindings.NewStringType()
	fields["tier1_id"] = bindings.NewStringType()
	fields["locale_service_id"] = bindings.NewStringType()
	fields["interface_id"] = bindings.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["deployment_group_id"] = "DeploymentGroupId"
	fieldNameMap["vhc_id"] = "VhcId"
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["locale_service_id"] = "LocaleServiceId"
	fieldNameMap["interface_id"] = "InterfaceId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serviceInterfacesDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func serviceInterfacesDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org_id"] = bindings.NewStringType()
	fields["deployment_group_id"] = bindings.NewStringType()
	fields["vhc_id"] = bindings.NewStringType()
	fields["tier1_id"] = bindings.NewStringType()
	fields["locale_service_id"] = bindings.NewStringType()
	fields["interface_id"] = bindings.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["deployment_group_id"] = "DeploymentGroupId"
	fieldNameMap["vhc_id"] = "VhcId"
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["locale_service_id"] = "LocaleServiceId"
	fieldNameMap["interface_id"] = "InterfaceId"
	paramsTypeMap["locale_service_id"] = bindings.NewStringType()
	paramsTypeMap["interface_id"] = bindings.NewStringType()
	paramsTypeMap["tier1_id"] = bindings.NewStringType()
	paramsTypeMap["org_id"] = bindings.NewStringType()
	paramsTypeMap["vhc_id"] = bindings.NewStringType()
	paramsTypeMap["deployment_group_id"] = bindings.NewStringType()
	paramsTypeMap["orgId"] = bindings.NewStringType()
	paramsTypeMap["deploymentGroupId"] = bindings.NewStringType()
	paramsTypeMap["vhcId"] = bindings.NewStringType()
	paramsTypeMap["tier1Id"] = bindings.NewStringType()
	paramsTypeMap["localeServiceId"] = bindings.NewStringType()
	paramsTypeMap["interfaceId"] = bindings.NewStringType()
	pathParams["vhc_id"] = "vhcId"
	pathParams["deployment_group_id"] = "deploymentGroupId"
	pathParams["tier1_id"] = "tier1Id"
	pathParams["interface_id"] = "interfaceId"
	pathParams["locale_service_id"] = "localeServiceId"
	pathParams["org_id"] = "orgId"
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
		"/policy/api/v1/org-root/orgs/{orgId}/deployment-groups/{deploymentGroupId}/vhcs/{vhcId}/infra/tier-1s/{tier1Id}/locale-services/{localeServiceId}/service-interfaces/{interfaceId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func serviceInterfacesGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = bindings.NewStringType()
	fields["deployment_group_id"] = bindings.NewStringType()
	fields["vhc_id"] = bindings.NewStringType()
	fields["tier1_id"] = bindings.NewStringType()
	fields["locale_service_id"] = bindings.NewStringType()
	fields["interface_id"] = bindings.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["deployment_group_id"] = "DeploymentGroupId"
	fieldNameMap["vhc_id"] = "VhcId"
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["locale_service_id"] = "LocaleServiceId"
	fieldNameMap["interface_id"] = "InterfaceId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serviceInterfacesGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceInterfaceBindingType)
}

func serviceInterfacesGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org_id"] = bindings.NewStringType()
	fields["deployment_group_id"] = bindings.NewStringType()
	fields["vhc_id"] = bindings.NewStringType()
	fields["tier1_id"] = bindings.NewStringType()
	fields["locale_service_id"] = bindings.NewStringType()
	fields["interface_id"] = bindings.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["deployment_group_id"] = "DeploymentGroupId"
	fieldNameMap["vhc_id"] = "VhcId"
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["locale_service_id"] = "LocaleServiceId"
	fieldNameMap["interface_id"] = "InterfaceId"
	paramsTypeMap["locale_service_id"] = bindings.NewStringType()
	paramsTypeMap["interface_id"] = bindings.NewStringType()
	paramsTypeMap["tier1_id"] = bindings.NewStringType()
	paramsTypeMap["org_id"] = bindings.NewStringType()
	paramsTypeMap["vhc_id"] = bindings.NewStringType()
	paramsTypeMap["deployment_group_id"] = bindings.NewStringType()
	paramsTypeMap["orgId"] = bindings.NewStringType()
	paramsTypeMap["deploymentGroupId"] = bindings.NewStringType()
	paramsTypeMap["vhcId"] = bindings.NewStringType()
	paramsTypeMap["tier1Id"] = bindings.NewStringType()
	paramsTypeMap["localeServiceId"] = bindings.NewStringType()
	paramsTypeMap["interfaceId"] = bindings.NewStringType()
	pathParams["vhc_id"] = "vhcId"
	pathParams["deployment_group_id"] = "deploymentGroupId"
	pathParams["tier1_id"] = "tier1Id"
	pathParams["interface_id"] = "interfaceId"
	pathParams["locale_service_id"] = "localeServiceId"
	pathParams["org_id"] = "orgId"
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
		"/policy/api/v1/org-root/orgs/{orgId}/deployment-groups/{deploymentGroupId}/vhcs/{vhcId}/infra/tier-1s/{tier1Id}/locale-services/{localeServiceId}/service-interfaces/{interfaceId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func serviceInterfacesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = bindings.NewStringType()
	fields["deployment_group_id"] = bindings.NewStringType()
	fields["vhc_id"] = bindings.NewStringType()
	fields["tier1_id"] = bindings.NewStringType()
	fields["locale_service_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["deployment_group_id"] = "DeploymentGroupId"
	fieldNameMap["vhc_id"] = "VhcId"
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["locale_service_id"] = "LocaleServiceId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serviceInterfacesListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceInterfaceListResultBindingType)
}

func serviceInterfacesListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org_id"] = bindings.NewStringType()
	fields["deployment_group_id"] = bindings.NewStringType()
	fields["vhc_id"] = bindings.NewStringType()
	fields["tier1_id"] = bindings.NewStringType()
	fields["locale_service_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["deployment_group_id"] = "DeploymentGroupId"
	fieldNameMap["vhc_id"] = "VhcId"
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["locale_service_id"] = "LocaleServiceId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["locale_service_id"] = bindings.NewStringType()
	paramsTypeMap["tier1_id"] = bindings.NewStringType()
	paramsTypeMap["org_id"] = bindings.NewStringType()
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["vhc_id"] = bindings.NewStringType()
	paramsTypeMap["deployment_group_id"] = bindings.NewStringType()
	paramsTypeMap["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["orgId"] = bindings.NewStringType()
	paramsTypeMap["deploymentGroupId"] = bindings.NewStringType()
	paramsTypeMap["vhcId"] = bindings.NewStringType()
	paramsTypeMap["tier1Id"] = bindings.NewStringType()
	paramsTypeMap["localeServiceId"] = bindings.NewStringType()
	pathParams["vhc_id"] = "vhcId"
	pathParams["deployment_group_id"] = "deploymentGroupId"
	pathParams["tier1_id"] = "tier1Id"
	pathParams["locale_service_id"] = "localeServiceId"
	pathParams["org_id"] = "orgId"
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sort_by"] = "sort_by"
	queryParams["include_mark_for_delete_objects"] = "include_mark_for_delete_objects"
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
		"/policy/api/v1/org-root/orgs/{orgId}/deployment-groups/{deploymentGroupId}/vhcs/{vhcId}/infra/tier-1s/{tier1Id}/locale-services/{localeServiceId}/service-interfaces",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func serviceInterfacesPatchInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = bindings.NewStringType()
	fields["deployment_group_id"] = bindings.NewStringType()
	fields["vhc_id"] = bindings.NewStringType()
	fields["tier1_id"] = bindings.NewStringType()
	fields["locale_service_id"] = bindings.NewStringType()
	fields["interface_id"] = bindings.NewStringType()
	fields["service_interface"] = bindings.NewReferenceType(model.ServiceInterfaceBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["deployment_group_id"] = "DeploymentGroupId"
	fieldNameMap["vhc_id"] = "VhcId"
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["locale_service_id"] = "LocaleServiceId"
	fieldNameMap["interface_id"] = "InterfaceId"
	fieldNameMap["service_interface"] = "ServiceInterface"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serviceInterfacesPatchOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func serviceInterfacesPatchRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org_id"] = bindings.NewStringType()
	fields["deployment_group_id"] = bindings.NewStringType()
	fields["vhc_id"] = bindings.NewStringType()
	fields["tier1_id"] = bindings.NewStringType()
	fields["locale_service_id"] = bindings.NewStringType()
	fields["interface_id"] = bindings.NewStringType()
	fields["service_interface"] = bindings.NewReferenceType(model.ServiceInterfaceBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["deployment_group_id"] = "DeploymentGroupId"
	fieldNameMap["vhc_id"] = "VhcId"
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["locale_service_id"] = "LocaleServiceId"
	fieldNameMap["interface_id"] = "InterfaceId"
	fieldNameMap["service_interface"] = "ServiceInterface"
	paramsTypeMap["locale_service_id"] = bindings.NewStringType()
	paramsTypeMap["interface_id"] = bindings.NewStringType()
	paramsTypeMap["tier1_id"] = bindings.NewStringType()
	paramsTypeMap["org_id"] = bindings.NewStringType()
	paramsTypeMap["service_interface"] = bindings.NewReferenceType(model.ServiceInterfaceBindingType)
	paramsTypeMap["vhc_id"] = bindings.NewStringType()
	paramsTypeMap["deployment_group_id"] = bindings.NewStringType()
	paramsTypeMap["orgId"] = bindings.NewStringType()
	paramsTypeMap["deploymentGroupId"] = bindings.NewStringType()
	paramsTypeMap["vhcId"] = bindings.NewStringType()
	paramsTypeMap["tier1Id"] = bindings.NewStringType()
	paramsTypeMap["localeServiceId"] = bindings.NewStringType()
	paramsTypeMap["interfaceId"] = bindings.NewStringType()
	pathParams["vhc_id"] = "vhcId"
	pathParams["deployment_group_id"] = "deploymentGroupId"
	pathParams["tier1_id"] = "tier1Id"
	pathParams["interface_id"] = "interfaceId"
	pathParams["locale_service_id"] = "localeServiceId"
	pathParams["org_id"] = "orgId"
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
		"service_interface",
		"PATCH",
		"/policy/api/v1/org-root/orgs/{orgId}/deployment-groups/{deploymentGroupId}/vhcs/{vhcId}/infra/tier-1s/{tier1Id}/locale-services/{localeServiceId}/service-interfaces/{interfaceId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func serviceInterfacesUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = bindings.NewStringType()
	fields["deployment_group_id"] = bindings.NewStringType()
	fields["vhc_id"] = bindings.NewStringType()
	fields["tier1_id"] = bindings.NewStringType()
	fields["locale_service_id"] = bindings.NewStringType()
	fields["interface_id"] = bindings.NewStringType()
	fields["service_interface"] = bindings.NewReferenceType(model.ServiceInterfaceBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["deployment_group_id"] = "DeploymentGroupId"
	fieldNameMap["vhc_id"] = "VhcId"
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["locale_service_id"] = "LocaleServiceId"
	fieldNameMap["interface_id"] = "InterfaceId"
	fieldNameMap["service_interface"] = "ServiceInterface"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serviceInterfacesUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceInterfaceBindingType)
}

func serviceInterfacesUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org_id"] = bindings.NewStringType()
	fields["deployment_group_id"] = bindings.NewStringType()
	fields["vhc_id"] = bindings.NewStringType()
	fields["tier1_id"] = bindings.NewStringType()
	fields["locale_service_id"] = bindings.NewStringType()
	fields["interface_id"] = bindings.NewStringType()
	fields["service_interface"] = bindings.NewReferenceType(model.ServiceInterfaceBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["deployment_group_id"] = "DeploymentGroupId"
	fieldNameMap["vhc_id"] = "VhcId"
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["locale_service_id"] = "LocaleServiceId"
	fieldNameMap["interface_id"] = "InterfaceId"
	fieldNameMap["service_interface"] = "ServiceInterface"
	paramsTypeMap["locale_service_id"] = bindings.NewStringType()
	paramsTypeMap["interface_id"] = bindings.NewStringType()
	paramsTypeMap["tier1_id"] = bindings.NewStringType()
	paramsTypeMap["org_id"] = bindings.NewStringType()
	paramsTypeMap["service_interface"] = bindings.NewReferenceType(model.ServiceInterfaceBindingType)
	paramsTypeMap["vhc_id"] = bindings.NewStringType()
	paramsTypeMap["deployment_group_id"] = bindings.NewStringType()
	paramsTypeMap["orgId"] = bindings.NewStringType()
	paramsTypeMap["deploymentGroupId"] = bindings.NewStringType()
	paramsTypeMap["vhcId"] = bindings.NewStringType()
	paramsTypeMap["tier1Id"] = bindings.NewStringType()
	paramsTypeMap["localeServiceId"] = bindings.NewStringType()
	paramsTypeMap["interfaceId"] = bindings.NewStringType()
	pathParams["vhc_id"] = "vhcId"
	pathParams["deployment_group_id"] = "deploymentGroupId"
	pathParams["tier1_id"] = "tier1Id"
	pathParams["interface_id"] = "interfaceId"
	pathParams["locale_service_id"] = "localeServiceId"
	pathParams["org_id"] = "orgId"
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
		"service_interface",
		"PUT",
		"/policy/api/v1/org-root/orgs/{orgId}/deployment-groups/{deploymentGroupId}/vhcs/{vhcId}/infra/tier-1s/{tier1Id}/locale-services/{localeServiceId}/service-interfaces/{interfaceId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
