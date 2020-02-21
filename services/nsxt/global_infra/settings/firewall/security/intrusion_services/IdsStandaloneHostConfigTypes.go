/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: IdsStandaloneHostConfig.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package intrusion_services

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/nsxt/model"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)





func idsStandaloneHostConfigGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func idsStandaloneHostConfigGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IdsStandaloneHostConfigBindingType)
}

func idsStandaloneHostConfigGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
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
		"/policy/api/v1/global-infra/settings/firewall/security/intrusion-services/ids-standalone-host-config",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func idsStandaloneHostConfigPatchInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ids_standalone_host_config"] = bindings.NewReferenceType(model.IdsStandaloneHostConfigBindingType)
	fieldNameMap["ids_standalone_host_config"] = "IdsStandaloneHostConfig"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func idsStandaloneHostConfigPatchOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func idsStandaloneHostConfigPatchRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["ids_standalone_host_config"] = bindings.NewReferenceType(model.IdsStandaloneHostConfigBindingType)
	fieldNameMap["ids_standalone_host_config"] = "IdsStandaloneHostConfig"
	paramsTypeMap["ids_standalone_host_config"] = bindings.NewReferenceType(model.IdsStandaloneHostConfigBindingType)
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
		"ids_standalone_host_config",
		"PATCH",
		"/policy/api/v1/global-infra/settings/firewall/security/intrusion-services/ids-standalone-host-config",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func idsStandaloneHostConfigUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ids_standalone_host_config"] = bindings.NewReferenceType(model.IdsStandaloneHostConfigBindingType)
	fieldNameMap["ids_standalone_host_config"] = "IdsStandaloneHostConfig"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func idsStandaloneHostConfigUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IdsStandaloneHostConfigBindingType)
}

func idsStandaloneHostConfigUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["ids_standalone_host_config"] = bindings.NewReferenceType(model.IdsStandaloneHostConfigBindingType)
	fieldNameMap["ids_standalone_host_config"] = "IdsStandaloneHostConfig"
	paramsTypeMap["ids_standalone_host_config"] = bindings.NewReferenceType(model.IdsStandaloneHostConfigBindingType)
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
		"ids_standalone_host_config",
		"PUT",
		"/policy/api/v1/global-infra/settings/firewall/security/intrusion-services/ids-standalone-host-config",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}


