/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Mem.
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
type MemHealthLevel string

const (
    // The service health is degraded. The service might have serious problems
	MemHealthLevel_orange MemHealthLevel = "orange"
    // No health data is available for this service.
	MemHealthLevel_gray MemHealthLevel = "gray"
    // Service is healthy.
	MemHealthLevel_green MemHealthLevel = "green"
    // The service is unavaiable, not functioning properly, or will stop functioning soon.
	MemHealthLevel_red MemHealthLevel = "red"
    // The service is healthy state, but experiencing some levels of problems.
	MemHealthLevel_yellow MemHealthLevel = "yellow"
)

func (h MemHealthLevel) MemHealthLevel() bool {
	switch h {
	case MemHealthLevel_orange:
		return true
	case MemHealthLevel_gray:
		return true
	case MemHealthLevel_green:
		return true
	case MemHealthLevel_red:
		return true
	case MemHealthLevel_yellow:
		return true
	default:
		return false
	}
}




func memGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func memGetOutputType() bindings.BindingType {
	return bindings.NewEnumType("com.vmware.appliance.health.mem.health_level", reflect.TypeOf(MemHealthLevel(MemHealthLevel_orange)))
}

func memGetRestMetadata() protocol.OperationRestMetadata {
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


