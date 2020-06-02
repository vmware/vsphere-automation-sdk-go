/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Logical.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package networks

import (
	"reflect"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/vmc/model"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
)





func logicalCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = bindings.NewStringType()
	fields["sddc"] = bindings.NewStringType()
	fields["sddc_network"] = bindings.NewReferenceType(model.SddcNetworkBindingType)
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc"] = "Sddc"
	fieldNameMap["sddc_network"] = "SddcNetwork"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func logicalCreateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func logicalCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org"] = bindings.NewStringType()
	fields["sddc"] = bindings.NewStringType()
	fields["sddc_network"] = bindings.NewReferenceType(model.SddcNetworkBindingType)
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc"] = "Sddc"
	fieldNameMap["sddc_network"] = "SddcNetwork"
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["sddc"] = bindings.NewStringType()
	paramsTypeMap["sddc_network"] = bindings.NewReferenceType(model.SddcNetworkBindingType)
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["sddc"] = bindings.NewStringType()
	pathParams["org"] = "org"
	pathParams["sddc"] = "sddc"
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
		"sddc_network",
		"POST",
		"/vmc/api/orgs/{org}/sddcs/{sddc}/networks/4.0/sddc/networks",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400,"com.vmware.vapi.std.errors.unauthorized": 403,"com.vmware.vapi.std.errors.not_found": 404})
}

func logicalDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = bindings.NewStringType()
	fields["sddc"] = bindings.NewStringType()
	fields["network_id"] = bindings.NewStringType()
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc"] = "Sddc"
	fieldNameMap["network_id"] = "NetworkId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func logicalDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func logicalDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org"] = bindings.NewStringType()
	fields["sddc"] = bindings.NewStringType()
	fields["network_id"] = bindings.NewStringType()
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc"] = "Sddc"
	fieldNameMap["network_id"] = "NetworkId"
	paramsTypeMap["network_id"] = bindings.NewStringType()
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["sddc"] = bindings.NewStringType()
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["sddc"] = bindings.NewStringType()
	paramsTypeMap["networkId"] = bindings.NewStringType()
	pathParams["org"] = "org"
	pathParams["sddc"] = "sddc"
	pathParams["network_id"] = "networkId"
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
		"/vmc/api/orgs/{org}/sddcs/{sddc}/networks/4.0/sddc/networks/{networkId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400,"com.vmware.vapi.std.errors.unauthorized": 403,"com.vmware.vapi.std.errors.not_found": 404})
}

func logicalGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = bindings.NewStringType()
	fields["sddc"] = bindings.NewStringType()
	fields["network_id"] = bindings.NewStringType()
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc"] = "Sddc"
	fieldNameMap["network_id"] = "NetworkId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func logicalGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.SddcNetworkBindingType)
}

func logicalGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org"] = bindings.NewStringType()
	fields["sddc"] = bindings.NewStringType()
	fields["network_id"] = bindings.NewStringType()
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc"] = "Sddc"
	fieldNameMap["network_id"] = "NetworkId"
	paramsTypeMap["network_id"] = bindings.NewStringType()
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["sddc"] = bindings.NewStringType()
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["sddc"] = bindings.NewStringType()
	paramsTypeMap["networkId"] = bindings.NewStringType()
	pathParams["org"] = "org"
	pathParams["sddc"] = "sddc"
	pathParams["network_id"] = "networkId"
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
		"/vmc/api/orgs/{org}/sddcs/{sddc}/networks/4.0/sddc/networks/{networkId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400,"com.vmware.vapi.std.errors.unauthorized": 403,"com.vmware.vapi.std.errors.not_found": 404})
}

func logicalGet0InputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = bindings.NewStringType()
	fields["sddc"] = bindings.NewStringType()
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["start_index"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["prev_sddc_network_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_order_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc"] = "Sddc"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["start_index"] = "StartIndex"
	fieldNameMap["prev_sddc_network_id"] = "PrevSddcNetworkId"
	fieldNameMap["sort_order_ascending"] = "SortOrderAscending"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func logicalGet0OutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.DataPageSddcNetworkBindingType)
}

func logicalGet0RestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org"] = bindings.NewStringType()
	fields["sddc"] = bindings.NewStringType()
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["start_index"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["prev_sddc_network_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_order_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc"] = "Sddc"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["start_index"] = "StartIndex"
	fieldNameMap["prev_sddc_network_id"] = "PrevSddcNetworkId"
	fieldNameMap["sort_order_ascending"] = "SortOrderAscending"
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["sddc"] = bindings.NewStringType()
	paramsTypeMap["sort_order_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["start_index"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["prev_sddc_network_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["sddc"] = bindings.NewStringType()
	pathParams["org"] = "org"
	pathParams["sddc"] = "sddc"
	queryParams["start_index"] = "startIndex"
	queryParams["prev_sddc_network_id"] = "prevSddcNetworkId"
	queryParams["page_size"] = "pageSize"
	queryParams["sort_order_ascending"] = "sortOrderAscending"
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
		"/vmc/api/orgs/{org}/sddcs/{sddc}/networks/4.0/sddc/networks",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400,"com.vmware.vapi.std.errors.unauthorized": 403,"com.vmware.vapi.std.errors.not_found": 404})
}

func logicalUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = bindings.NewStringType()
	fields["sddc"] = bindings.NewStringType()
	fields["network_id"] = bindings.NewStringType()
	fields["sddc_network"] = bindings.NewReferenceType(model.SddcNetworkBindingType)
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc"] = "Sddc"
	fieldNameMap["network_id"] = "NetworkId"
	fieldNameMap["sddc_network"] = "SddcNetwork"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func logicalUpdateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func logicalUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org"] = bindings.NewStringType()
	fields["sddc"] = bindings.NewStringType()
	fields["network_id"] = bindings.NewStringType()
	fields["sddc_network"] = bindings.NewReferenceType(model.SddcNetworkBindingType)
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc"] = "Sddc"
	fieldNameMap["network_id"] = "NetworkId"
	fieldNameMap["sddc_network"] = "SddcNetwork"
	paramsTypeMap["network_id"] = bindings.NewStringType()
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["sddc"] = bindings.NewStringType()
	paramsTypeMap["sddc_network"] = bindings.NewReferenceType(model.SddcNetworkBindingType)
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["sddc"] = bindings.NewStringType()
	paramsTypeMap["networkId"] = bindings.NewStringType()
	pathParams["org"] = "org"
	pathParams["sddc"] = "sddc"
	pathParams["network_id"] = "networkId"
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
		"sddc_network",
		"PUT",
		"/vmc/api/orgs/{org}/sddcs/{sddc}/networks/4.0/sddc/networks/{networkId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400,"com.vmware.vapi.std.errors.unauthorized": 403,"com.vmware.vapi.std.errors.not_found": 404})
}


