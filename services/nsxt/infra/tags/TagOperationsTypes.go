/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: TagOperations.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package tags

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/nsxt/model"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)





func tagOperationsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["operation_id"] = bindings.NewStringType()
	fieldNameMap["operation_id"] = "OperationId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tagOperationsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TagBulkOperationBindingType)
}

func tagOperationsGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["operation_id"] = bindings.NewStringType()
	fieldNameMap["operation_id"] = "OperationId"
	paramsTypeMap["operation_id"] = bindings.NewStringType()
	paramsTypeMap["operationId"] = bindings.NewStringType()
	pathParams["operation_id"] = "operationId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"GET",
		"/policy/api/v1/infra/tags/tag-operations/{operationId}",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func tagOperationsUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["operation_id"] = bindings.NewStringType()
	fields["tag_bulk_operation"] = bindings.NewReferenceType(model.TagBulkOperationBindingType)
	fieldNameMap["operation_id"] = "OperationId"
	fieldNameMap["tag_bulk_operation"] = "TagBulkOperation"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tagOperationsUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TagBulkOperationBindingType)
}

func tagOperationsUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["operation_id"] = bindings.NewStringType()
	fields["tag_bulk_operation"] = bindings.NewReferenceType(model.TagBulkOperationBindingType)
	fieldNameMap["operation_id"] = "OperationId"
	fieldNameMap["tag_bulk_operation"] = "TagBulkOperation"
	paramsTypeMap["tag_bulk_operation"] = bindings.NewReferenceType(model.TagBulkOperationBindingType)
	paramsTypeMap["operation_id"] = bindings.NewStringType()
	paramsTypeMap["operationId"] = bindings.NewStringType()
	pathParams["operation_id"] = "operationId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"tag_bulk_operation",
		"PUT",
		"/policy/api/v1/infra/tags/tag-operations/{operationId}",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}


