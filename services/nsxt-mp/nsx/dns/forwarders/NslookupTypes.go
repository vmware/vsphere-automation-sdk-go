// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Nslookup.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package forwarders

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

func nslookupGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["forwarder_id"] = vapiBindings_.NewStringType()
	fields["address"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["server_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["source_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["forwarder_id"] = "ForwarderId"
	fieldNameMap["address"] = "Address"
	fieldNameMap["server_ip"] = "ServerIp"
	fieldNameMap["source_ip"] = "SourceIp"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func NslookupGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.DnsAnswerBindingType)
}

func nslookupGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["forwarder_id"] = vapiBindings_.NewStringType()
	fields["address"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["server_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["source_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["forwarder_id"] = "ForwarderId"
	fieldNameMap["address"] = "Address"
	fieldNameMap["server_ip"] = "ServerIp"
	fieldNameMap["source_ip"] = "SourceIp"
	paramsTypeMap["address"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["server_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["forwarder_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["source_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["forwarderId"] = vapiBindings_.NewStringType()
	pathParams["forwarder_id"] = "forwarderId"
	queryParams["address"] = "address"
	queryParams["server_ip"] = "server_ip"
	queryParams["source_ip"] = "source_ip"
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
		"/api/v1/dns/forwarders/{forwarderId}/nslookup",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
