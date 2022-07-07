// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: FeedbackResponse.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package migration

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
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

func feedbackResponseAcceptrecommendedInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["network_layer"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["network_layer"] = "NetworkLayer"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func feedbackResponseAcceptrecommendedOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func feedbackResponseAcceptrecommendedRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["network_layer"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["network_layer"] = "NetworkLayer"
	paramsTypeMap["network_layer"] = bindings.NewOptionalType(bindings.NewStringType())
	queryParams["network_layer"] = "network_layer"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"] = make(map[string]string)
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"]["challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
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

func feedbackResponseUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["migration_feedback_response_list"] = bindings.NewReferenceType(model.MigrationFeedbackResponseListBindingType)
	fields["network_layer"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["migration_feedback_response_list"] = "MigrationFeedbackResponseList"
	fieldNameMap["network_layer"] = "NetworkLayer"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func feedbackResponseUpdateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func feedbackResponseUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["migration_feedback_response_list"] = bindings.NewReferenceType(model.MigrationFeedbackResponseListBindingType)
	fields["network_layer"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["migration_feedback_response_list"] = "MigrationFeedbackResponseList"
	fieldNameMap["network_layer"] = "NetworkLayer"
	paramsTypeMap["network_layer"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["migration_feedback_response_list"] = bindings.NewReferenceType(model.MigrationFeedbackResponseListBindingType)
	queryParams["network_layer"] = "network_layer"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"] = make(map[string]string)
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"]["challenge"] = "WWW-Authenticate"
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
