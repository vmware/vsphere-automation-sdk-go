// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: DnsResolverConfig.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package firewall

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func dnsResolverConfigDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier0_id"] = vapiBindings_.NewStringType()
	fieldNameMap["tier0_id"] = "Tier0Id"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func DnsResolverConfigDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func dnsResolverConfigDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["tier0_id"] = vapiBindings_.NewStringType()
	fieldNameMap["tier0_id"] = "Tier0Id"
	paramsTypeMap["tier0_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["tier0Id"] = vapiBindings_.NewStringType()
	pathParams["tier0_id"] = "tier0Id"
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
		"/policy/api/v1/infra/tier-0s/{tier0Id}/firewall/dns-resolver-config",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func dnsResolverConfigGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier0_id"] = vapiBindings_.NewStringType()
	fieldNameMap["tier0_id"] = "Tier0Id"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func DnsResolverConfigGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.GatewayDnsResolverConfigBindingType)
}

func dnsResolverConfigGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["tier0_id"] = vapiBindings_.NewStringType()
	fieldNameMap["tier0_id"] = "Tier0Id"
	paramsTypeMap["tier0_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["tier0Id"] = vapiBindings_.NewStringType()
	pathParams["tier0_id"] = "tier0Id"
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
		"/policy/api/v1/infra/tier-0s/{tier0Id}/firewall/dns-resolver-config",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func dnsResolverConfigPatchInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier0_id"] = vapiBindings_.NewStringType()
	fields["gateway_dns_resolver_config"] = vapiBindings_.NewReferenceType(nsx_policyModel.GatewayDnsResolverConfigBindingType)
	fieldNameMap["tier0_id"] = "Tier0Id"
	fieldNameMap["gateway_dns_resolver_config"] = "GatewayDnsResolverConfig"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func DnsResolverConfigPatchOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func dnsResolverConfigPatchRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["tier0_id"] = vapiBindings_.NewStringType()
	fields["gateway_dns_resolver_config"] = vapiBindings_.NewReferenceType(nsx_policyModel.GatewayDnsResolverConfigBindingType)
	fieldNameMap["tier0_id"] = "Tier0Id"
	fieldNameMap["gateway_dns_resolver_config"] = "GatewayDnsResolverConfig"
	paramsTypeMap["tier0_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["gateway_dns_resolver_config"] = vapiBindings_.NewReferenceType(nsx_policyModel.GatewayDnsResolverConfigBindingType)
	paramsTypeMap["tier0Id"] = vapiBindings_.NewStringType()
	pathParams["tier0_id"] = "tier0Id"
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
		"gateway_dns_resolver_config",
		"PATCH",
		"/policy/api/v1/infra/tier-0s/{tier0Id}/firewall/dns-resolver-config",
		"application/json",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func dnsResolverConfigUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier0_id"] = vapiBindings_.NewStringType()
	fields["gateway_dns_resolver_config"] = vapiBindings_.NewReferenceType(nsx_policyModel.GatewayDnsResolverConfigBindingType)
	fieldNameMap["tier0_id"] = "Tier0Id"
	fieldNameMap["gateway_dns_resolver_config"] = "GatewayDnsResolverConfig"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func DnsResolverConfigUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.GatewayDnsResolverConfigBindingType)
}

func dnsResolverConfigUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["tier0_id"] = vapiBindings_.NewStringType()
	fields["gateway_dns_resolver_config"] = vapiBindings_.NewReferenceType(nsx_policyModel.GatewayDnsResolverConfigBindingType)
	fieldNameMap["tier0_id"] = "Tier0Id"
	fieldNameMap["gateway_dns_resolver_config"] = "GatewayDnsResolverConfig"
	paramsTypeMap["tier0_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["gateway_dns_resolver_config"] = vapiBindings_.NewReferenceType(nsx_policyModel.GatewayDnsResolverConfigBindingType)
	paramsTypeMap["tier0Id"] = vapiBindings_.NewStringType()
	pathParams["tier0_id"] = "tier0Id"
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
		"gateway_dns_resolver_config",
		"PUT",
		"/policy/api/v1/infra/tier-0s/{tier0Id}/firewall/dns-resolver-config",
		"application/json",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
