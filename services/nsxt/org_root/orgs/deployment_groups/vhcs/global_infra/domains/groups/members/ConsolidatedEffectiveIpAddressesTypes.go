// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: ConsolidatedEffectiveIpAddresses.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package members

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func consolidatedEffectiveIpAddressesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = bindings.NewStringType()
	fields["deployment_group_id"] = bindings.NewStringType()
	fields["vhc_id"] = bindings.NewStringType()
	fields["domain_id"] = bindings.NewStringType()
	fields["group_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["ip_filter"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["site_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["deployment_group_id"] = "DeploymentGroupId"
	fieldNameMap["vhc_id"] = "VhcId"
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["ip_filter"] = "IpFilter"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func consolidatedEffectiveIpAddressesListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ConsolidatedEffectiveIPAddressMemberListResultBindingType)
}

func consolidatedEffectiveIpAddressesListRestMetadata() protocol.OperationRestMetadata {
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
	fields["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["ip_filter"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["site_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["deployment_group_id"] = "DeploymentGroupId"
	fieldNameMap["vhc_id"] = "VhcId"
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["ip_filter"] = "IpFilter"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["vhc_id"] = bindings.NewStringType()
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["domain_id"] = bindings.NewStringType()
	paramsTypeMap["site_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["group_id"] = bindings.NewStringType()
	paramsTypeMap["org_id"] = bindings.NewStringType()
	paramsTypeMap["ip_filter"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["deployment_group_id"] = bindings.NewStringType()
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
	queryParams["site_id"] = "site_id"
	queryParams["enforcement_point_path"] = "enforcement_point_path"
	queryParams["sort_by"] = "sort_by"
	queryParams["ip_filter"] = "ip_filter"
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
		"/policy/api/v1/org-root/orgs/{orgId}/deployment-groups/{deploymentGroupId}/vhcs/{vhcId}/global-infra/domains/{domainId}/groups/{groupId}/members/consolidated-effective-ip-addresses",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
