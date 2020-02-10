/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Shutdown.
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
	"time"
)



// ``ShutdownConfig`` class Structure that defines shutdown configuration returned by the Shutdown.get operation. This class was added in vSphere API 6.7.
type ShutdownShutdownConfig struct {
    // Shutdown time. This property was added in vSphere API 6.7.
	ShutdownTime *time.Time
    // The pending shutdown operation. The string values for pending operations can be 'poweroff', 'reboot' or ''. This property was added in vSphere API 6.7.
	Action string
    // The reason behind the shutdown action. This property was added in vSphere API 6.7.
	Reason string
}



func shutdownCancelInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func shutdownCancelOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func shutdownCancelRestMetadata() protocol.OperationRestMetadata {
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

func shutdownPoweroffInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["delay"] = bindings.NewIntegerType()
	fields["reason"] = bindings.NewStringType()
	fieldNameMap["delay"] = "Delay"
	fieldNameMap["reason"] = "Reason"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func shutdownPoweroffOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func shutdownPoweroffRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["delay"] = bindings.NewIntegerType()
	fields["reason"] = bindings.NewStringType()
	fieldNameMap["delay"] = "Delay"
	fieldNameMap["reason"] = "Reason"
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

func shutdownRebootInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["delay"] = bindings.NewIntegerType()
	fields["reason"] = bindings.NewStringType()
	fieldNameMap["delay"] = "Delay"
	fieldNameMap["reason"] = "Reason"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func shutdownRebootOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func shutdownRebootRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["delay"] = bindings.NewIntegerType()
	fields["reason"] = bindings.NewStringType()
	fieldNameMap["delay"] = "Delay"
	fieldNameMap["reason"] = "Reason"
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

func shutdownGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func shutdownGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(ShutdownShutdownConfigBindingType)
}

func shutdownGetRestMetadata() protocol.OperationRestMetadata {
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


func ShutdownShutdownConfigBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["shutdown_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
	fieldNameMap["shutdown_time"] = "ShutdownTime"
	fields["action"] = bindings.NewStringType()
	fieldNameMap["action"] = "Action"
	fields["reason"] = bindings.NewStringType()
	fieldNameMap["reason"] = "Reason"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.shutdown.shutdown_config", fields, reflect.TypeOf(ShutdownShutdownConfig{}), fieldNameMap, validators)
}

