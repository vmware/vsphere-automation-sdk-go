// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: IdsEventsByCategory.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package intrusion_services

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

// Possible value for ``timeperiod`` of method IdsEventsByCategory#list.
const IdsEventsByCategory_LIST_TIMEPERIOD_24HOURS = "24HOURS"

// Possible value for ``timeperiod`` of method IdsEventsByCategory#list.
const IdsEventsByCategory_LIST_TIMEPERIOD_48HOURS = "48HOURS"

// Possible value for ``timeperiod`` of method IdsEventsByCategory#list.
const IdsEventsByCategory_LIST_TIMEPERIOD_7DAYS = "7DAYS"

func idsEventsByCategoryListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ids_category_name"] = bindings.NewStringType()
	fields["timeperiod"] = bindings.NewStringType()
	fields["detected_on_gateway"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["gateway"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["ids_category_name"] = "IdsCategoryName"
	fieldNameMap["timeperiod"] = "Timeperiod"
	fieldNameMap["detected_on_gateway"] = "DetectedOnGateway"
	fieldNameMap["gateway"] = "Gateway"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func idsEventsByCategoryListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PolicyIdsEventsByCategoryResultBindingType)
}

func idsEventsByCategoryListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ids_category_name"] = bindings.NewStringType()
	fields["timeperiod"] = bindings.NewStringType()
	fields["detected_on_gateway"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["gateway"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["ids_category_name"] = "IdsCategoryName"
	fieldNameMap["timeperiod"] = "Timeperiod"
	fieldNameMap["detected_on_gateway"] = "DetectedOnGateway"
	fieldNameMap["gateway"] = "Gateway"
	paramsTypeMap["ids_category_name"] = bindings.NewStringType()
	paramsTypeMap["timeperiod"] = bindings.NewStringType()
	paramsTypeMap["gateway"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["detected_on_gateway"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["idsCategoryName"] = bindings.NewStringType()
	pathParams["ids_category_name"] = "idsCategoryName"
	queryParams["detected_on_gateway"] = "detected_on_gateway"
	queryParams["gateway"] = "gateway"
	queryParams["timeperiod"] = "timeperiod"
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
		"/policy/api/v1/global-infra/settings/firewall/security/intrusion-services/ids-events-by-category/{idsCategoryName}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
