// Copyright (c) 2019-2025 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: BlockedEvents.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package geo_ip

import (
	"reflect"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
)





func blockedEventsGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["destination_country_code"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["destination_ip_address"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["direction"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_all_projects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["rule_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["source_country_code"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["source_ip_address"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["destination_country_code"] = "DestinationCountryCode"
	fieldNameMap["destination_ip_address"] = "DestinationIpAddress"
	fieldNameMap["direction"] = "Direction"
	fieldNameMap["include_all_projects"] = "IncludeAllProjects"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["rule_id"] = "RuleId"
	fieldNameMap["source_country_code"] = "SourceCountryCode"
	fieldNameMap["source_ip_address"] = "SourceIpAddress"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func BlockedEventsGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.GeoIpBlockedEventsListBindingType)
}

func blockedEventsGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["destination_country_code"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["destination_ip_address"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["direction"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_all_projects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["rule_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["source_country_code"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["source_ip_address"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["destination_country_code"] = "DestinationCountryCode"
	fieldNameMap["destination_ip_address"] = "DestinationIpAddress"
	fieldNameMap["direction"] = "Direction"
	fieldNameMap["include_all_projects"] = "IncludeAllProjects"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["rule_id"] = "RuleId"
	fieldNameMap["source_country_code"] = "SourceCountryCode"
	fieldNameMap["source_ip_address"] = "SourceIpAddress"
	paramsTypeMap["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["include_all_projects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["rule_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["source_ip_address"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["project_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["destination_country_code"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["source_country_code"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["destination_ip_address"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["direction"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	paramsTypeMap["projectId"] = vapiBindings_.NewStringType()
	pathParams["project_id"] = "projectId"
	pathParams["org_id"] = "orgId"
	queryParams["cursor"] = "cursor"
	queryParams["include_all_projects"] = "include_all_projects"
	queryParams["rule_id"] = "rule_id"
	queryParams["source_ip_address"] = "source_ip_address"
	queryParams["destination_country_code"] = "destination_country_code"
	queryParams["source_country_code"] = "source_country_code"
	queryParams["destination_ip_address"] = "destination_ip_address"
	queryParams["direction"] = "direction"
	queryParams["page_size"] = "page_size"
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
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/infra/geo-ip/blocked-events",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400,"com.vmware.vapi.std.errors.unauthorized": 403,"com.vmware.vapi.std.errors.service_unavailable": 503,"com.vmware.vapi.std.errors.internal_server_error": 500,"com.vmware.vapi.std.errors.not_found": 404})
}


