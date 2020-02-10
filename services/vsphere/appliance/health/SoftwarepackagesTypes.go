/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Softwarepackages.
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
type SoftwarepackagesHealthLevel string

const (
    // The service health is degraded. The service might have serious problems.
	SoftwarepackagesHealthLevel_orange SoftwarepackagesHealthLevel = "orange"
    // No health data is available for this service.
	SoftwarepackagesHealthLevel_gray SoftwarepackagesHealthLevel = "gray"
    // Service is healthy.
	SoftwarepackagesHealthLevel_green SoftwarepackagesHealthLevel = "green"
    // The service is unavaiable, not functioning properly, or will stop functioning soon.
	SoftwarepackagesHealthLevel_red SoftwarepackagesHealthLevel = "red"
    // The service is healthy state, but experiencing some levels of problems.
	SoftwarepackagesHealthLevel_yellow SoftwarepackagesHealthLevel = "yellow"
)

func (h SoftwarepackagesHealthLevel) SoftwarepackagesHealthLevel() bool {
	switch h {
	case SoftwarepackagesHealthLevel_orange:
		return true
	case SoftwarepackagesHealthLevel_gray:
		return true
	case SoftwarepackagesHealthLevel_green:
		return true
	case SoftwarepackagesHealthLevel_red:
		return true
	case SoftwarepackagesHealthLevel_yellow:
		return true
	default:
		return false
	}
}




func softwarepackagesGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func softwarepackagesGetOutputType() bindings.BindingType {
	return bindings.NewEnumType("com.vmware.appliance.health.softwarepackages.health_level", reflect.TypeOf(SoftwarepackagesHealthLevel(SoftwarepackagesHealthLevel_orange)))
}

func softwarepackagesGetRestMetadata() protocol.OperationRestMetadata {
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


