// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: GroupMonitoringProfileBindingMaps.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package groups

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func groupMonitoringProfileBindingMapsDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = bindings.NewStringType()
	fields["deployment_group_id"] = bindings.NewStringType()
	fields["vhc_id"] = bindings.NewStringType()
	fields["domain_id"] = bindings.NewStringType()
	fields["group_id"] = bindings.NewStringType()
	fields["group_monitoring_profile_binding_map_id"] = bindings.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["deployment_group_id"] = "DeploymentGroupId"
	fieldNameMap["vhc_id"] = "VhcId"
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["group_monitoring_profile_binding_map_id"] = "GroupMonitoringProfileBindingMapId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func groupMonitoringProfileBindingMapsDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func groupMonitoringProfileBindingMapsDeleteRestMetadata() protocol.OperationRestMetadata {
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
	fields["domain_id"] = bindings.NewStringType()
	fields["group_id"] = bindings.NewStringType()
	fields["group_monitoring_profile_binding_map_id"] = bindings.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["deployment_group_id"] = "DeploymentGroupId"
	fieldNameMap["vhc_id"] = "VhcId"
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["group_monitoring_profile_binding_map_id"] = "GroupMonitoringProfileBindingMapId"
	paramsTypeMap["domain_id"] = bindings.NewStringType()
	paramsTypeMap["group_id"] = bindings.NewStringType()
	paramsTypeMap["group_monitoring_profile_binding_map_id"] = bindings.NewStringType()
	paramsTypeMap["org_id"] = bindings.NewStringType()
	paramsTypeMap["vhc_id"] = bindings.NewStringType()
	paramsTypeMap["deployment_group_id"] = bindings.NewStringType()
	paramsTypeMap["orgId"] = bindings.NewStringType()
	paramsTypeMap["deploymentGroupId"] = bindings.NewStringType()
	paramsTypeMap["vhcId"] = bindings.NewStringType()
	paramsTypeMap["domainId"] = bindings.NewStringType()
	paramsTypeMap["groupId"] = bindings.NewStringType()
	paramsTypeMap["groupMonitoringProfileBindingMapId"] = bindings.NewStringType()
	pathParams["vhc_id"] = "vhcId"
	pathParams["group_id"] = "groupId"
	pathParams["group_monitoring_profile_binding_map_id"] = "groupMonitoringProfileBindingMapId"
	pathParams["deployment_group_id"] = "deploymentGroupId"
	pathParams["org_id"] = "orgId"
	pathParams["domain_id"] = "domainId"
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
		"/policy/api/v1/org-root/orgs/{orgId}/deployment-groups/{deploymentGroupId}/vhcs/{vhcId}/infra/domains/{domainId}/groups/{groupId}/group-monitoring-profile-binding-maps/{groupMonitoringProfileBindingMapId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func groupMonitoringProfileBindingMapsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = bindings.NewStringType()
	fields["deployment_group_id"] = bindings.NewStringType()
	fields["vhc_id"] = bindings.NewStringType()
	fields["domain_id"] = bindings.NewStringType()
	fields["group_id"] = bindings.NewStringType()
	fields["group_monitoring_profile_binding_map_id"] = bindings.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["deployment_group_id"] = "DeploymentGroupId"
	fieldNameMap["vhc_id"] = "VhcId"
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["group_monitoring_profile_binding_map_id"] = "GroupMonitoringProfileBindingMapId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func groupMonitoringProfileBindingMapsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.GroupMonitoringProfileBindingMapBindingType)
}

func groupMonitoringProfileBindingMapsGetRestMetadata() protocol.OperationRestMetadata {
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
	fields["domain_id"] = bindings.NewStringType()
	fields["group_id"] = bindings.NewStringType()
	fields["group_monitoring_profile_binding_map_id"] = bindings.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["deployment_group_id"] = "DeploymentGroupId"
	fieldNameMap["vhc_id"] = "VhcId"
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["group_monitoring_profile_binding_map_id"] = "GroupMonitoringProfileBindingMapId"
	paramsTypeMap["domain_id"] = bindings.NewStringType()
	paramsTypeMap["group_id"] = bindings.NewStringType()
	paramsTypeMap["group_monitoring_profile_binding_map_id"] = bindings.NewStringType()
	paramsTypeMap["org_id"] = bindings.NewStringType()
	paramsTypeMap["vhc_id"] = bindings.NewStringType()
	paramsTypeMap["deployment_group_id"] = bindings.NewStringType()
	paramsTypeMap["orgId"] = bindings.NewStringType()
	paramsTypeMap["deploymentGroupId"] = bindings.NewStringType()
	paramsTypeMap["vhcId"] = bindings.NewStringType()
	paramsTypeMap["domainId"] = bindings.NewStringType()
	paramsTypeMap["groupId"] = bindings.NewStringType()
	paramsTypeMap["groupMonitoringProfileBindingMapId"] = bindings.NewStringType()
	pathParams["vhc_id"] = "vhcId"
	pathParams["group_id"] = "groupId"
	pathParams["group_monitoring_profile_binding_map_id"] = "groupMonitoringProfileBindingMapId"
	pathParams["deployment_group_id"] = "deploymentGroupId"
	pathParams["org_id"] = "orgId"
	pathParams["domain_id"] = "domainId"
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
		"/policy/api/v1/org-root/orgs/{orgId}/deployment-groups/{deploymentGroupId}/vhcs/{vhcId}/infra/domains/{domainId}/groups/{groupId}/group-monitoring-profile-binding-maps/{groupMonitoringProfileBindingMapId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func groupMonitoringProfileBindingMapsListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = bindings.NewStringType()
	fields["deployment_group_id"] = bindings.NewStringType()
	fields["vhc_id"] = bindings.NewStringType()
	fields["domain_id"] = bindings.NewStringType()
	fields["group_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["deployment_group_id"] = "DeploymentGroupId"
	fieldNameMap["vhc_id"] = "VhcId"
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func groupMonitoringProfileBindingMapsListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.GroupMonitoringProfileBindingMapListResultBindingType)
}

func groupMonitoringProfileBindingMapsListRestMetadata() protocol.OperationRestMetadata {
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
	fields["domain_id"] = bindings.NewStringType()
	fields["group_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["deployment_group_id"] = "DeploymentGroupId"
	fieldNameMap["vhc_id"] = "VhcId"
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["domain_id"] = bindings.NewStringType()
	paramsTypeMap["group_id"] = bindings.NewStringType()
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
	paramsTypeMap["domainId"] = bindings.NewStringType()
	paramsTypeMap["groupId"] = bindings.NewStringType()
	pathParams["vhc_id"] = "vhcId"
	pathParams["group_id"] = "groupId"
	pathParams["deployment_group_id"] = "deploymentGroupId"
	pathParams["org_id"] = "orgId"
	pathParams["domain_id"] = "domainId"
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
		"/policy/api/v1/org-root/orgs/{orgId}/deployment-groups/{deploymentGroupId}/vhcs/{vhcId}/infra/domains/{domainId}/groups/{groupId}/group-monitoring-profile-binding-maps",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func groupMonitoringProfileBindingMapsPatchInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = bindings.NewStringType()
	fields["deployment_group_id"] = bindings.NewStringType()
	fields["vhc_id"] = bindings.NewStringType()
	fields["domain_id"] = bindings.NewStringType()
	fields["group_id"] = bindings.NewStringType()
	fields["group_monitoring_profile_binding_map_id"] = bindings.NewStringType()
	fields["group_monitoring_profile_binding_map"] = bindings.NewReferenceType(model.GroupMonitoringProfileBindingMapBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["deployment_group_id"] = "DeploymentGroupId"
	fieldNameMap["vhc_id"] = "VhcId"
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["group_monitoring_profile_binding_map_id"] = "GroupMonitoringProfileBindingMapId"
	fieldNameMap["group_monitoring_profile_binding_map"] = "GroupMonitoringProfileBindingMap"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func groupMonitoringProfileBindingMapsPatchOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func groupMonitoringProfileBindingMapsPatchRestMetadata() protocol.OperationRestMetadata {
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
	fields["domain_id"] = bindings.NewStringType()
	fields["group_id"] = bindings.NewStringType()
	fields["group_monitoring_profile_binding_map_id"] = bindings.NewStringType()
	fields["group_monitoring_profile_binding_map"] = bindings.NewReferenceType(model.GroupMonitoringProfileBindingMapBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["deployment_group_id"] = "DeploymentGroupId"
	fieldNameMap["vhc_id"] = "VhcId"
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["group_monitoring_profile_binding_map_id"] = "GroupMonitoringProfileBindingMapId"
	fieldNameMap["group_monitoring_profile_binding_map"] = "GroupMonitoringProfileBindingMap"
	paramsTypeMap["domain_id"] = bindings.NewStringType()
	paramsTypeMap["group_monitoring_profile_binding_map"] = bindings.NewReferenceType(model.GroupMonitoringProfileBindingMapBindingType)
	paramsTypeMap["group_id"] = bindings.NewStringType()
	paramsTypeMap["group_monitoring_profile_binding_map_id"] = bindings.NewStringType()
	paramsTypeMap["org_id"] = bindings.NewStringType()
	paramsTypeMap["vhc_id"] = bindings.NewStringType()
	paramsTypeMap["deployment_group_id"] = bindings.NewStringType()
	paramsTypeMap["orgId"] = bindings.NewStringType()
	paramsTypeMap["deploymentGroupId"] = bindings.NewStringType()
	paramsTypeMap["vhcId"] = bindings.NewStringType()
	paramsTypeMap["domainId"] = bindings.NewStringType()
	paramsTypeMap["groupId"] = bindings.NewStringType()
	paramsTypeMap["groupMonitoringProfileBindingMapId"] = bindings.NewStringType()
	pathParams["vhc_id"] = "vhcId"
	pathParams["group_id"] = "groupId"
	pathParams["group_monitoring_profile_binding_map_id"] = "groupMonitoringProfileBindingMapId"
	pathParams["deployment_group_id"] = "deploymentGroupId"
	pathParams["org_id"] = "orgId"
	pathParams["domain_id"] = "domainId"
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
		"group_monitoring_profile_binding_map",
		"PATCH",
		"/policy/api/v1/org-root/orgs/{orgId}/deployment-groups/{deploymentGroupId}/vhcs/{vhcId}/infra/domains/{domainId}/groups/{groupId}/group-monitoring-profile-binding-maps/{groupMonitoringProfileBindingMapId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func groupMonitoringProfileBindingMapsUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = bindings.NewStringType()
	fields["deployment_group_id"] = bindings.NewStringType()
	fields["vhc_id"] = bindings.NewStringType()
	fields["domain_id"] = bindings.NewStringType()
	fields["group_id"] = bindings.NewStringType()
	fields["group_monitoring_profile_binding_map_id"] = bindings.NewStringType()
	fields["group_monitoring_profile_binding_map"] = bindings.NewReferenceType(model.GroupMonitoringProfileBindingMapBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["deployment_group_id"] = "DeploymentGroupId"
	fieldNameMap["vhc_id"] = "VhcId"
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["group_monitoring_profile_binding_map_id"] = "GroupMonitoringProfileBindingMapId"
	fieldNameMap["group_monitoring_profile_binding_map"] = "GroupMonitoringProfileBindingMap"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func groupMonitoringProfileBindingMapsUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.GroupMonitoringProfileBindingMapBindingType)
}

func groupMonitoringProfileBindingMapsUpdateRestMetadata() protocol.OperationRestMetadata {
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
	fields["domain_id"] = bindings.NewStringType()
	fields["group_id"] = bindings.NewStringType()
	fields["group_monitoring_profile_binding_map_id"] = bindings.NewStringType()
	fields["group_monitoring_profile_binding_map"] = bindings.NewReferenceType(model.GroupMonitoringProfileBindingMapBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["deployment_group_id"] = "DeploymentGroupId"
	fieldNameMap["vhc_id"] = "VhcId"
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["group_monitoring_profile_binding_map_id"] = "GroupMonitoringProfileBindingMapId"
	fieldNameMap["group_monitoring_profile_binding_map"] = "GroupMonitoringProfileBindingMap"
	paramsTypeMap["domain_id"] = bindings.NewStringType()
	paramsTypeMap["group_monitoring_profile_binding_map"] = bindings.NewReferenceType(model.GroupMonitoringProfileBindingMapBindingType)
	paramsTypeMap["group_id"] = bindings.NewStringType()
	paramsTypeMap["group_monitoring_profile_binding_map_id"] = bindings.NewStringType()
	paramsTypeMap["org_id"] = bindings.NewStringType()
	paramsTypeMap["vhc_id"] = bindings.NewStringType()
	paramsTypeMap["deployment_group_id"] = bindings.NewStringType()
	paramsTypeMap["orgId"] = bindings.NewStringType()
	paramsTypeMap["deploymentGroupId"] = bindings.NewStringType()
	paramsTypeMap["vhcId"] = bindings.NewStringType()
	paramsTypeMap["domainId"] = bindings.NewStringType()
	paramsTypeMap["groupId"] = bindings.NewStringType()
	paramsTypeMap["groupMonitoringProfileBindingMapId"] = bindings.NewStringType()
	pathParams["vhc_id"] = "vhcId"
	pathParams["group_id"] = "groupId"
	pathParams["group_monitoring_profile_binding_map_id"] = "groupMonitoringProfileBindingMapId"
	pathParams["deployment_group_id"] = "deploymentGroupId"
	pathParams["org_id"] = "orgId"
	pathParams["domain_id"] = "domainId"
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
		"group_monitoring_profile_binding_map",
		"PUT",
		"/policy/api/v1/org-root/orgs/{orgId}/deployment-groups/{deploymentGroupId}/vhcs/{vhcId}/infra/domains/{domainId}/groups/{groupId}/group-monitoring-profile-binding-maps/{groupMonitoringProfileBindingMapId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
