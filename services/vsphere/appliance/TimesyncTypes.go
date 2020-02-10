/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Timesync.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package appliance

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``TimeSyncMode`` enumeration class defines time synchronization modes. This enumeration was added in vSphere API 6.7.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type TimesyncTimeSyncMode string

const (
    // Time synchronization is disabled. This constant field was added in vSphere API 6.7.
	TimesyncTimeSyncMode_DISABLED TimesyncTimeSyncMode = "DISABLED"
    // NTP-based time synchronization. This constant field was added in vSphere API 6.7.
	TimesyncTimeSyncMode_NTP TimesyncTimeSyncMode = "NTP"
    // VMware Tool-based time synchronization. This constant field was added in vSphere API 6.7.
	TimesyncTimeSyncMode_HOST TimesyncTimeSyncMode = "HOST"
)

func (t TimesyncTimeSyncMode) TimesyncTimeSyncMode() bool {
	switch t {
	case TimesyncTimeSyncMode_DISABLED:
		return true
	case TimesyncTimeSyncMode_NTP:
		return true
	case TimesyncTimeSyncMode_HOST:
		return true
	default:
		return false
	}
}




func timesyncSetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["mode"] = bindings.NewEnumType("com.vmware.appliance.timesync.time_sync_mode", reflect.TypeOf(TimesyncTimeSyncMode(TimesyncTimeSyncMode_DISABLED)))
	fieldNameMap["mode"] = "Mode"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func timesyncSetOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func timesyncSetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["mode"] = bindings.NewEnumType("com.vmware.appliance.timesync.time_sync_mode", reflect.TypeOf(TimesyncTimeSyncMode(TimesyncTimeSyncMode_DISABLED)))
	fieldNameMap["mode"] = "Mode"
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

func timesyncGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func timesyncGetOutputType() bindings.BindingType {
	return bindings.NewEnumType("com.vmware.appliance.timesync.time_sync_mode", reflect.TypeOf(TimesyncTimeSyncMode(TimesyncTimeSyncMode_DISABLED)))
}

func timesyncGetRestMetadata() protocol.OperationRestMetadata {
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


