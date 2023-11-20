// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Subscriptions.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package orgs

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	vmcModel "github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
	"reflect"
)

// Possible value for ``offerType`` of method Subscriptions#get0.
const Subscriptions_GET_0_OFFER_TYPE_TERM = "TERM"

// Possible value for ``offerType`` of method Subscriptions#get0.
const Subscriptions_GET_0_OFFER_TYPE_ON_DEMAND = "ON_DEMAND"

func subscriptionsCreateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = vapiBindings_.NewStringType()
	fields["subscription_request"] = vapiBindings_.NewReferenceType(vmcModel.SubscriptionRequestBindingType)
	fieldNameMap["org"] = "Org"
	fieldNameMap["subscription_request"] = "SubscriptionRequest"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SubscriptionsCreateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(vmcModel.TaskBindingType)
}

func subscriptionsCreateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org"] = vapiBindings_.NewStringType()
	fields["subscription_request"] = vapiBindings_.NewReferenceType(vmcModel.SubscriptionRequestBindingType)
	fieldNameMap["org"] = "Org"
	fieldNameMap["subscription_request"] = "SubscriptionRequest"
	paramsTypeMap["subscription_request"] = vapiBindings_.NewReferenceType(vmcModel.SubscriptionRequestBindingType)
	paramsTypeMap["org"] = vapiBindings_.NewStringType()
	paramsTypeMap["org"] = vapiBindings_.NewStringType()
	pathParams["org"] = "org"
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
		"subscription_request",
		"POST",
		"/vmc/api/orgs/{org}/subscriptions",
		"",
		resultHeaders,
		201,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthenticated": 401, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.internal_server_error": 500})
}

func subscriptionsGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = vapiBindings_.NewStringType()
	fields["subscription"] = vapiBindings_.NewStringType()
	fieldNameMap["org"] = "Org"
	fieldNameMap["subscription"] = "Subscription"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SubscriptionsGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(vmcModel.SubscriptionDetailsBindingType)
}

func subscriptionsGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org"] = vapiBindings_.NewStringType()
	fields["subscription"] = vapiBindings_.NewStringType()
	fieldNameMap["org"] = "Org"
	fieldNameMap["subscription"] = "Subscription"
	paramsTypeMap["org"] = vapiBindings_.NewStringType()
	paramsTypeMap["subscription"] = vapiBindings_.NewStringType()
	paramsTypeMap["org"] = vapiBindings_.NewStringType()
	paramsTypeMap["subscription"] = vapiBindings_.NewStringType()
	pathParams["org"] = "org"
	pathParams["subscription"] = "subscription"
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
		"/vmc/api/orgs/{org}/subscriptions/{subscription}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func subscriptionsGet0InputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = vapiBindings_.NewStringType()
	fields["offer_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org"] = "Org"
	fieldNameMap["offer_type"] = "OfferType"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SubscriptionsGet0OutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewListType(vapiBindings_.NewReferenceType(vmcModel.SubscriptionDetailsBindingType), reflect.TypeOf([]vmcModel.SubscriptionDetails{}))
}

func subscriptionsGet0RestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org"] = vapiBindings_.NewStringType()
	fields["offer_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org"] = "Org"
	fieldNameMap["offer_type"] = "OfferType"
	paramsTypeMap["org"] = vapiBindings_.NewStringType()
	paramsTypeMap["offer_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["org"] = vapiBindings_.NewStringType()
	pathParams["org"] = "org"
	queryParams["offer_type"] = "offer_type"
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
		"/vmc/api/orgs/{org}/subscriptions",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthenticated": 401, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
