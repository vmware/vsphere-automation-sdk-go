/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Status.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package edges

import (
	"reflect"
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
)





func statusGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = bindings.NewStringType()
	fields["sddc"] = bindings.NewStringType()
	fields["edge_id"] = bindings.NewStringType()
	fields["getlatest"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["detailed"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc"] = "Sddc"
	fieldNameMap["edge_id"] = "EdgeId"
	fieldNameMap["getlatest"] = "Getlatest"
	fieldNameMap["detailed"] = "Detailed"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func statusGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.EdgeStatusBindingType)
}

func statusGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["org"] = bindings.NewStringType()
	fields["sddc"] = bindings.NewStringType()
	fields["edge_id"] = bindings.NewStringType()
	fields["getlatest"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["detailed"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc"] = "Sddc"
	fieldNameMap["edge_id"] = "EdgeId"
	fieldNameMap["getlatest"] = "Getlatest"
	fieldNameMap["detailed"] = "Detailed"
	paramsTypeMap["detailed"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["sddc"] = bindings.NewStringType()
	paramsTypeMap["getlatest"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["edge_id"] = bindings.NewStringType()
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["sddc"] = bindings.NewStringType()
	paramsTypeMap["edgeId"] = bindings.NewStringType()
	pathParams["edge_id"] = "edgeId"
	pathParams["org"] = "org"
	pathParams["sddc"] = "sddc"
	queryParams["detailed"] = "detailed"
	queryParams["getlatest"] = "getlatest"
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
		"/vmc/api/orgs/{org}/sddcs/{sddc}/networks/4.0/edges/{edgeId}/status",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"NotFound": 404})
}


