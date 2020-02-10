/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Databasestorage.
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



// ``HealthLevel`` enumeration class Defines service health levels.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type DatabasestorageHealthLevel string

const (
    // The service health is degraded. The service might have serious problems.
	DatabasestorageHealthLevel_orange DatabasestorageHealthLevel = "orange"
    // No health data is available for this service.
	DatabasestorageHealthLevel_gray DatabasestorageHealthLevel = "gray"
    // The service is healthy.
	DatabasestorageHealthLevel_green DatabasestorageHealthLevel = "green"
    // The service is unavaiable, not functioning properly, or will stop functioning soon.
	DatabasestorageHealthLevel_red DatabasestorageHealthLevel = "red"
    // The service is healthy but experiencing some problems.
	DatabasestorageHealthLevel_yellow DatabasestorageHealthLevel = "yellow"
)

func (h DatabasestorageHealthLevel) DatabasestorageHealthLevel() bool {
	switch h {
	case DatabasestorageHealthLevel_orange:
		return true
	case DatabasestorageHealthLevel_gray:
		return true
	case DatabasestorageHealthLevel_green:
		return true
	case DatabasestorageHealthLevel_red:
		return true
	case DatabasestorageHealthLevel_yellow:
		return true
	default:
		return false
	}
}




func databasestorageGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func databasestorageGetOutputType() bindings.BindingType {
	return bindings.NewEnumType("com.vmware.appliance.health.databasestorage.health_level", reflect.TypeOf(DatabasestorageHealthLevel(DatabasestorageHealthLevel_orange)))
}

func databasestorageGetRestMetadata() protocol.OperationRestMetadata {
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


