/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Swap.
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
type SwapHealthLevel string

const (
    // The service health is degraded. The service might have serious problems.
	SwapHealthLevel_orange SwapHealthLevel = "orange"
    // No health data is available for this service.
	SwapHealthLevel_gray SwapHealthLevel = "gray"
    // Service is healthy.
	SwapHealthLevel_green SwapHealthLevel = "green"
    // The service is unavaiable, not functioning properly, or will stop functioning soon.
	SwapHealthLevel_red SwapHealthLevel = "red"
    // The service is healthy state, but experiencing some levels of problems.
	SwapHealthLevel_yellow SwapHealthLevel = "yellow"
)

func (h SwapHealthLevel) SwapHealthLevel() bool {
	switch h {
	case SwapHealthLevel_orange:
		return true
	case SwapHealthLevel_gray:
		return true
	case SwapHealthLevel_green:
		return true
	case SwapHealthLevel_red:
		return true
	case SwapHealthLevel_yellow:
		return true
	default:
		return false
	}
}




func swapGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func swapGetOutputType() bindings.BindingType {
	return bindings.NewEnumType("com.vmware.appliance.health.swap.health_level", reflect.TypeOf(SwapHealthLevel(SwapHealthLevel_orange)))
}

func swapGetRestMetadata() protocol.OperationRestMetadata {
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


