// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: VendorTemplates.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package services

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

func vendorTemplatesCreateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = vapiBindings_.NewStringType()
	fields["vendor_template"] = vapiBindings_.NewReferenceType(nsxModel.VendorTemplateBindingType)
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["vendor_template"] = "VendorTemplate"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func VendorTemplatesCreateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.VendorTemplateBindingType)
}

func vendorTemplatesCreateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = vapiBindings_.NewStringType()
	fields["vendor_template"] = vapiBindings_.NewReferenceType(nsxModel.VendorTemplateBindingType)
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["vendor_template"] = "VendorTemplate"
	paramsTypeMap["vendor_template"] = vapiBindings_.NewReferenceType(nsxModel.VendorTemplateBindingType)
	paramsTypeMap["service_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["serviceId"] = vapiBindings_.NewStringType()
	pathParams["service_id"] = "serviceId"
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
		"vendor_template",
		"POST",
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/vendor-templates",
		"",
		resultHeaders,
		201,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func vendorTemplatesDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = vapiBindings_.NewStringType()
	fields["vendor_template_id"] = vapiBindings_.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["vendor_template_id"] = "VendorTemplateId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func VendorTemplatesDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func vendorTemplatesDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = vapiBindings_.NewStringType()
	fields["vendor_template_id"] = vapiBindings_.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["vendor_template_id"] = "VendorTemplateId"
	paramsTypeMap["vendor_template_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["service_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["serviceId"] = vapiBindings_.NewStringType()
	paramsTypeMap["vendorTemplateId"] = vapiBindings_.NewStringType()
	pathParams["service_id"] = "serviceId"
	pathParams["vendor_template_id"] = "vendorTemplateId"
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
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/vendor-templates/{vendorTemplateId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func vendorTemplatesGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = vapiBindings_.NewStringType()
	fields["vendor_template_id"] = vapiBindings_.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["vendor_template_id"] = "VendorTemplateId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func VendorTemplatesGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.VendorTemplateBindingType)
}

func vendorTemplatesGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = vapiBindings_.NewStringType()
	fields["vendor_template_id"] = vapiBindings_.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["vendor_template_id"] = "VendorTemplateId"
	paramsTypeMap["vendor_template_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["service_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["serviceId"] = vapiBindings_.NewStringType()
	paramsTypeMap["vendorTemplateId"] = vapiBindings_.NewStringType()
	pathParams["service_id"] = "serviceId"
	pathParams["vendor_template_id"] = "vendorTemplateId"
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
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/vendor-templates/{vendorTemplateId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func vendorTemplatesListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = vapiBindings_.NewStringType()
	fields["vendor_template_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["vendor_template_name"] = "VendorTemplateName"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func VendorTemplatesListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.VendorTemplateListResultBindingType)
}

func vendorTemplatesListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = vapiBindings_.NewStringType()
	fields["vendor_template_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["vendor_template_name"] = "VendorTemplateName"
	paramsTypeMap["vendor_template_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["service_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["serviceId"] = vapiBindings_.NewStringType()
	pathParams["service_id"] = "serviceId"
	queryParams["vendor_template_name"] = "vendor_template_name"
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
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/vendor-templates",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
