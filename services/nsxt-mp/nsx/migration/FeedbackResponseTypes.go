// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: FeedbackResponse.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package migration

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

// Possible value for ``networkLayer`` of method FeedbackResponse#acceptrecommended.
const FeedbackResponse_ACCEPTRECOMMENDED_NETWORK_LAYER_L2 = "L2"

// Possible value for ``networkLayer`` of method FeedbackResponse#acceptrecommended.
const FeedbackResponse_ACCEPTRECOMMENDED_NETWORK_LAYER_L3_L7 = "L3_L7"

// Possible value for ``networkLayer`` of method FeedbackResponse#update.
const FeedbackResponse_UPDATE_NETWORK_LAYER_L2 = "L2"

// Possible value for ``networkLayer`` of method FeedbackResponse#update.
const FeedbackResponse_UPDATE_NETWORK_LAYER_L3_L7 = "L3_L7"

func feedbackResponseAcceptrecommendedInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["network_layer"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["network_layer"] = "NetworkLayer"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func FeedbackResponseAcceptrecommendedOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func feedbackResponseAcceptrecommendedRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["network_layer"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["network_layer"] = "NetworkLayer"
	paramsTypeMap["network_layer"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	queryParams["network_layer"] = "network_layer"
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
		"action=accept-recommended",
		"",
		"POST",
		"/api/v1/migration/feedback-response",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthenticated": 401, "com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func feedbackResponseUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["migration_feedback_response_list"] = vapiBindings_.NewReferenceType(nsxModel.MigrationFeedbackResponseListBindingType)
	fields["network_layer"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["migration_feedback_response_list"] = "MigrationFeedbackResponseList"
	fieldNameMap["network_layer"] = "NetworkLayer"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func FeedbackResponseUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func feedbackResponseUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["migration_feedback_response_list"] = vapiBindings_.NewReferenceType(nsxModel.MigrationFeedbackResponseListBindingType)
	fields["network_layer"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["migration_feedback_response_list"] = "MigrationFeedbackResponseList"
	fieldNameMap["network_layer"] = "NetworkLayer"
	paramsTypeMap["network_layer"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["migration_feedback_response_list"] = vapiBindings_.NewReferenceType(nsxModel.MigrationFeedbackResponseListBindingType)
	queryParams["network_layer"] = "network_layer"
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
		"migration_feedback_response_list",
		"PUT",
		"/api/v1/migration/feedback-response",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthenticated": 401, "com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
