/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Storage.
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
type StorageHealthLevel string

const (
    // The service health is degraded. The service might have serious problems.
	StorageHealthLevel_orange StorageHealthLevel = "orange"
    // No health data is available for this service.
	StorageHealthLevel_gray StorageHealthLevel = "gray"
    // Service is healthy.
	StorageHealthLevel_green StorageHealthLevel = "green"
    // The service is unavaiable, not functioning properly, or will stop functioning soon.
	StorageHealthLevel_red StorageHealthLevel = "red"
    // The service is healthy state, but experiencing some levels of problems.
	StorageHealthLevel_yellow StorageHealthLevel = "yellow"
)

func (h StorageHealthLevel) StorageHealthLevel() bool {
	switch h {
	case StorageHealthLevel_orange:
		return true
	case StorageHealthLevel_gray:
		return true
	case StorageHealthLevel_green:
		return true
	case StorageHealthLevel_red:
		return true
	case StorageHealthLevel_yellow:
		return true
	default:
		return false
	}
}




func storageGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func storageGetOutputType() bindings.BindingType {
	return bindings.NewEnumType("com.vmware.appliance.health.storage.health_level", reflect.TypeOf(StorageHealthLevel(StorageHealthLevel_orange)))
}

func storageGetRestMetadata() protocol.OperationRestMetadata {
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


