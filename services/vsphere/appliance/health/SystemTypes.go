/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: System.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package health

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// ``HealthLevel`` enumeration class Defines health levels.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type SystemHealthLevel string

const (
    // The service health is degraded. The service might have serious problems.
	SystemHealthLevel_orange SystemHealthLevel = "orange"
    // No health data is available for this service.
	SystemHealthLevel_gray SystemHealthLevel = "gray"
    // Service is healthy.
	SystemHealthLevel_green SystemHealthLevel = "green"
    // The service is unavaiable, not functioning properly, or will stop functioning soon.
	SystemHealthLevel_red SystemHealthLevel = "red"
    // The service is healthy state, but experiencing some levels of problems.
	SystemHealthLevel_yellow SystemHealthLevel = "yellow"
)

func (h SystemHealthLevel) SystemHealthLevel() bool {
	switch h {
	case SystemHealthLevel_orange:
		return true
	case SystemHealthLevel_gray:
		return true
	case SystemHealthLevel_green:
		return true
	case SystemHealthLevel_red:
		return true
	case SystemHealthLevel_yellow:
		return true
	default:
		return false
	}
}




func systemLastcheckInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemLastcheckOutputType() bindings.BindingType {
	return bindings.NewDateTimeType()
}

func systemLastcheckRestMetadata() protocol.OperationRestMetadata {
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
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"Error": 500})
}

func systemGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemGetOutputType() bindings.BindingType {
	return bindings.NewEnumType("com.vmware.appliance.health.system.health_level", reflect.TypeOf(SystemHealthLevel(SystemHealthLevel_orange)))
}

func systemGetRestMetadata() protocol.OperationRestMetadata {
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
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"Error": 500})
}


