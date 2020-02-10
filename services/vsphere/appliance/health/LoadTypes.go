/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Load.
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
type LoadHealthLevel string

const (
    // The service health is degraded. The service might have serious problems.
	LoadHealthLevel_orange LoadHealthLevel = "orange"
    // No health data is available for this service.
	LoadHealthLevel_gray LoadHealthLevel = "gray"
    // Service is healthy.
	LoadHealthLevel_green LoadHealthLevel = "green"
    // The service is unavaiable, not functioning properly, or will stop functioning soon.
	LoadHealthLevel_red LoadHealthLevel = "red"
    // The service is healthy state, but experiencing some levels of problems.
	LoadHealthLevel_yellow LoadHealthLevel = "yellow"
)

func (h LoadHealthLevel) LoadHealthLevel() bool {
	switch h {
	case LoadHealthLevel_orange:
		return true
	case LoadHealthLevel_gray:
		return true
	case LoadHealthLevel_green:
		return true
	case LoadHealthLevel_red:
		return true
	case LoadHealthLevel_yellow:
		return true
	default:
		return false
	}
}




func loadGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func loadGetOutputType() bindings.BindingType {
	return bindings.NewEnumType("com.vmware.appliance.health.load.health_level", reflect.TypeOf(LoadHealthLevel(LoadHealthLevel_orange)))
}

func loadGetRestMetadata() protocol.OperationRestMetadata {
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


