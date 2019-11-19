/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Interface.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package dashboard

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/vmc/model"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)





func interfaceGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = bindings.NewStringType()
	fields["sddc"] = bindings.NewStringType()
	fields["edge_id"] = bindings.NewStringType()
	fields["interval"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc"] = "Sddc"
	fieldNameMap["edge_id"] = "EdgeId"
	fieldNameMap["interval"] = "Interval"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func interfaceGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.DashboardStatisticsBindingType)
}

func interfaceGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["org"] = bindings.NewStringType()
	fields["sddc"] = bindings.NewStringType()
	fields["edge_id"] = bindings.NewStringType()
	fields["interval"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc"] = "Sddc"
	fieldNameMap["edge_id"] = "EdgeId"
	fieldNameMap["interval"] = "Interval"
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["sddc"] = bindings.NewStringType()
	paramsTypeMap["interval"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["edge_id"] = bindings.NewStringType()
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["sddc"] = bindings.NewStringType()
	paramsTypeMap["edgeId"] = bindings.NewStringType()
	pathParams["edge_id"] = "edgeId"
	pathParams["org"] = "org"
	pathParams["sddc"] = "sddc"
	queryParams["interval"] = "interval"
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
		"/vmc/api/orgs/{org}/sddcs/{sddc}/networks/4.0/edges/{edgeId}/statistics/dashboard/interface",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"NotFound": 404})
}


