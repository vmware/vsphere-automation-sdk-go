/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: DetailedStatus.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package lb_virtual_servers

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/nsxt/model"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// Possible value for ``source`` of method DetailedStatus#get.
const DetailedStatus_GET_SOURCE_REALTIME = "realtime"
// Possible value for ``source`` of method DetailedStatus#get.
const DetailedStatus_GET_SOURCE_CACHED = "cached"




func detailedStatusGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["lb_service_id"] = bindings.NewStringType()
	fields["lb_virtual_server_id"] = bindings.NewStringType()
	fields["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["lb_service_id"] = "LbServiceId"
	fieldNameMap["lb_virtual_server_id"] = "LbVirtualServerId"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	fieldNameMap["source"] = "Source"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func detailedStatusGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.AggregateLBVirtualServerStatusBindingType)
}

func detailedStatusGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["lb_service_id"] = bindings.NewStringType()
	fields["lb_virtual_server_id"] = bindings.NewStringType()
	fields["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["lb_service_id"] = "LbServiceId"
	fieldNameMap["lb_virtual_server_id"] = "LbVirtualServerId"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	fieldNameMap["source"] = "Source"
	paramsTypeMap["source"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["lb_service_id"] = bindings.NewStringType()
	paramsTypeMap["lb_virtual_server_id"] = bindings.NewStringType()
	paramsTypeMap["lbServiceId"] = bindings.NewStringType()
	paramsTypeMap["lbVirtualServerId"] = bindings.NewStringType()
	pathParams["lb_service_id"] = "lbServiceId"
	pathParams["lb_virtual_server_id"] = "lbVirtualServerId"
	queryParams["enforcement_point_path"] = "enforcement_point_path"
	queryParams["source"] = "source"
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
		"/policy/api/v1/infra/lb-services/{lbServiceId}/lb-virtual-servers/{lbVirtualServerId}/detailed-status",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}


