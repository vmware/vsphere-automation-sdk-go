// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: SystemAdministrationConfigurationFabricEdgeClustersFailureDomains.
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

func systemAdministrationConfigurationFabricEdgeClustersFailureDomainsCreateFailureDomainInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["failure_domain"] = bindings.NewReferenceType(model.FailureDomainBindingType)
	fieldNameMap["failure_domain"] = "FailureDomain"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricEdgeClustersFailureDomainsCreateFailureDomainOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FailureDomainBindingType)
}

func systemAdministrationConfigurationFabricEdgeClustersFailureDomainsCreateFailureDomainRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["failure_domain"] = bindings.NewReferenceType(model.FailureDomainBindingType)
	fieldNameMap["failure_domain"] = "FailureDomain"
	paramsTypeMap["failure_domain"] = bindings.NewReferenceType(model.FailureDomainBindingType)
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
		"failure_domain",
		"POST",
		"/api/v1/failure-domains",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricEdgeClustersFailureDomainsDeleteFailureDomainInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["failure_domain_id"] = bindings.NewStringType()
	fieldNameMap["failure_domain_id"] = "FailureDomainId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricEdgeClustersFailureDomainsDeleteFailureDomainOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func systemAdministrationConfigurationFabricEdgeClustersFailureDomainsDeleteFailureDomainRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["failure_domain_id"] = bindings.NewStringType()
	fieldNameMap["failure_domain_id"] = "FailureDomainId"
	paramsTypeMap["failure_domain_id"] = bindings.NewStringType()
	paramsTypeMap["failureDomainId"] = bindings.NewStringType()
	pathParams["failure_domain_id"] = "failureDomainId"
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
		"/api/v1/failure-domains/{failureDomainId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricEdgeClustersFailureDomainsGetFailureDomainInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["failure_domain_id"] = bindings.NewStringType()
	fieldNameMap["failure_domain_id"] = "FailureDomainId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricEdgeClustersFailureDomainsGetFailureDomainOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FailureDomainBindingType)
}

func systemAdministrationConfigurationFabricEdgeClustersFailureDomainsGetFailureDomainRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["failure_domain_id"] = bindings.NewStringType()
	fieldNameMap["failure_domain_id"] = "FailureDomainId"
	paramsTypeMap["failure_domain_id"] = bindings.NewStringType()
	paramsTypeMap["failureDomainId"] = bindings.NewStringType()
	pathParams["failure_domain_id"] = "failureDomainId"
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
		"/api/v1/failure-domains/{failureDomainId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricEdgeClustersFailureDomainsListFailureDomainsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricEdgeClustersFailureDomainsListFailureDomainsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FailureDomainListResultBindingType)
}

func systemAdministrationConfigurationFabricEdgeClustersFailureDomainsListFailureDomainsRestMetadata() protocol.OperationRestMetadata {
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
		"/api/v1/failure-domains",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricEdgeClustersFailureDomainsUpdateFailureDomainInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["failure_domain_id"] = bindings.NewStringType()
	fields["failure_domain"] = bindings.NewReferenceType(model.FailureDomainBindingType)
	fieldNameMap["failure_domain_id"] = "FailureDomainId"
	fieldNameMap["failure_domain"] = "FailureDomain"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricEdgeClustersFailureDomainsUpdateFailureDomainOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FailureDomainBindingType)
}

func systemAdministrationConfigurationFabricEdgeClustersFailureDomainsUpdateFailureDomainRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["failure_domain_id"] = bindings.NewStringType()
	fields["failure_domain"] = bindings.NewReferenceType(model.FailureDomainBindingType)
	fieldNameMap["failure_domain_id"] = "FailureDomainId"
	fieldNameMap["failure_domain"] = "FailureDomain"
	paramsTypeMap["failure_domain_id"] = bindings.NewStringType()
	paramsTypeMap["failure_domain"] = bindings.NewReferenceType(model.FailureDomainBindingType)
	paramsTypeMap["failureDomainId"] = bindings.NewStringType()
	pathParams["failure_domain_id"] = "failureDomainId"
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
		"failure_domain",
		"PUT",
		"/api/v1/failure-domains/{failureDomainId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
