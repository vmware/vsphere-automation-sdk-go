// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: OfferInstances.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package subscriptions

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	vmcModel "github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
	"reflect"
)

func offerInstancesListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = vapiBindings_.NewStringType()
	fields["region"] = vapiBindings_.NewStringType()
	fields["product_type"] = vapiBindings_.NewStringType()
	fields["product"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org"] = "Org"
	fieldNameMap["region"] = "Region"
	fieldNameMap["product_type"] = "ProductType"
	fieldNameMap["product"] = "Product"
	fieldNameMap["type"] = "Type_"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func OfferInstancesListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(vmcModel.OfferInstancesHolderBindingType)
}

func offerInstancesListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org"] = vapiBindings_.NewStringType()
	fields["region"] = vapiBindings_.NewStringType()
	fields["product_type"] = vapiBindings_.NewStringType()
	fields["product"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org"] = "Org"
	fieldNameMap["region"] = "Region"
	fieldNameMap["product_type"] = "ProductType"
	fieldNameMap["product"] = "Product"
	fieldNameMap["type"] = "Type_"
	paramsTypeMap["product"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["product_type"] = vapiBindings_.NewStringType()
	paramsTypeMap["org"] = vapiBindings_.NewStringType()
	paramsTypeMap["region"] = vapiBindings_.NewStringType()
	paramsTypeMap["type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["org"] = vapiBindings_.NewStringType()
	pathParams["org"] = "org"
	queryParams["product"] = "product"
	queryParams["product_type"] = "product_type"
	queryParams["region"] = "region"
	queryParams["type"] = "type"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"] = make(map[string]string)
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"]["challenge"] = "WWW-Authenticate"
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
		"/vmc/api/orgs/{org}/subscriptions/offer-instances",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthenticated": 401, "com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403})
}
