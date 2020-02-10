/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Parallel.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package hardware

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// Resource type for the virtual parallel port.
const Parallel_RESOURCE_TYPE = "com.vmware.vcenter.vm.hardware.ParallelPort"


// The ``BackingType`` enumeration class defines the valid backing types for a virtual parallel port.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type ParallelBackingType string

const (
    // Virtual parallel port is backed by a file.
	ParallelBackingType_FILE ParallelBackingType = "FILE"
    // Virtual parallel port is backed by a device on the host where the virtual machine is running.
	ParallelBackingType_HOST_DEVICE ParallelBackingType = "HOST_DEVICE"
)

func (b ParallelBackingType) ParallelBackingType() bool {
	switch b {
	case ParallelBackingType_FILE:
		return true
	case ParallelBackingType_HOST_DEVICE:
		return true
	default:
		return false
	}
}


// The ``BackingInfo`` class contains information about the physical resource backing a virtual parallel port.
type ParallelBackingInfo struct {
    // Backing type for the virtual parallel port.
	Type_ ParallelBackingType
    // Path of the file backing the virtual parallel port.
	File *string
    // Name of the device backing the virtual parallel port. 
	HostDevice *string
    // Flag indicating whether the virtual parallel port is configured to automatically detect a suitable host device.
	AutoDetect *bool
}

// The ``BackingSpec`` class provides a specification of the physical resource backing a virtual parallel port.
type ParallelBackingSpec struct {
    // Backing type for the virtual parallel port.
	Type_ ParallelBackingType
    // Path of the file that should be used as the virtual parallel port backing.
	File *string
    // Name of the device that should be used as the virtual parallel port backing.
	HostDevice *string
}

// The ``Info`` class contains information about a virtual parallel port.
type ParallelInfo struct {
    // Device label.
	Label string
    // Physical resource backing for the virtual parallel port.
	Backing ParallelBackingInfo
    // Connection status of the virtual device.
	State ConnectionState
    // Flag indicating whether the virtual device should be connected whenever the virtual machine is powered on.
	StartConnected bool
    // Flag indicating whether the guest can connect and disconnect the device.
	AllowGuestControl bool
}

// The ``CreateSpec`` class provides a specification for the configuration of a newly-created virtual parallel port.
type ParallelCreateSpec struct {
    // Physical resource backing for the virtual parallel port.
	Backing *ParallelBackingSpec
    // Flag indicating whether the virtual device should be connected whenever the virtual machine is powered on.
	StartConnected *bool
    // Flag indicating whether the guest can connect and disconnect the device.
	AllowGuestControl *bool
}

// The ``UpdateSpec`` class describes the updates to be made to the configuration of a virtual parallel port.
type ParallelUpdateSpec struct {
    // Physical resource backing for the virtual parallel port. 
    //
    //  This property may only be modified if the virtual machine is not powered on or the virtual parallel port is not connected.
	Backing *ParallelBackingSpec
    // Flag indicating whether the virtual device should be connected whenever the virtual machine is powered on.
	StartConnected *bool
    // Flag indicating whether the guest can connect and disconnect the device.
	AllowGuestControl *bool
}

// The ``Summary`` class contains commonly used information about a virtual parallel port.
type ParallelSummary struct {
    // Identifier of the virtual parallel port.
	Port string
}



func parallelListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fieldNameMap["vm"] = "Vm"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func parallelListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(ParallelSummaryBindingType), reflect.TypeOf([]ParallelSummary{}))
}

func parallelListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fieldNameMap["vm"] = "Vm"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
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
		map[string]int{"Error": 500,"NotFound": 404,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func parallelGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["port"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.ParallelPort"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["port"] = "Port"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func parallelGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(ParallelInfoBindingType)
}

func parallelGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["port"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.ParallelPort"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["port"] = "Port"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
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
		map[string]int{"Error": 500,"NotFound": 404,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func parallelCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["spec"] = bindings.NewReferenceType(ParallelCreateSpecBindingType)
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func parallelCreateOutputType() bindings.BindingType {
	return bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.ParallelPort"}, "")
}

func parallelCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["spec"] = bindings.NewReferenceType(ParallelCreateSpecBindingType)
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["spec"] = "Spec"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
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
		map[string]int{"Error": 500,"NotFound": 404,"NotAllowedInCurrentState": 400,"UnableToAllocateResource": 500,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func parallelUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["port"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.ParallelPort"}, "")
	fields["spec"] = bindings.NewReferenceType(ParallelUpdateSpecBindingType)
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["port"] = "Port"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func parallelUpdateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func parallelUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["port"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.ParallelPort"}, "")
	fields["spec"] = bindings.NewReferenceType(ParallelUpdateSpecBindingType)
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["port"] = "Port"
	fieldNameMap["spec"] = "Spec"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
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
		map[string]int{"Error": 500,"NotFound": 404,"NotAllowedInCurrentState": 400,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func parallelDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["port"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.ParallelPort"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["port"] = "Port"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func parallelDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func parallelDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["port"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.ParallelPort"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["port"] = "Port"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
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
		map[string]int{"Error": 500,"NotFound": 404,"NotAllowedInCurrentState": 400,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func parallelConnectInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["port"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.ParallelPort"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["port"] = "Port"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func parallelConnectOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func parallelConnectRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["port"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.ParallelPort"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["port"] = "Port"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
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
		map[string]int{"Error": 500,"NotFound": 404,"AlreadyInDesiredState": 400,"NotAllowedInCurrentState": 400,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func parallelDisconnectInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["port"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.ParallelPort"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["port"] = "Port"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func parallelDisconnectOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func parallelDisconnectRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["port"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.ParallelPort"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["port"] = "Port"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
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
		map[string]int{"Error": 500,"NotFound": 404,"AlreadyInDesiredState": 400,"NotAllowedInCurrentState": 400,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func ParallelBackingInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.parallel.backing_type", reflect.TypeOf(ParallelBackingType(ParallelBackingType_FILE)))
	fieldNameMap["type"] = "Type_"
	fields["file"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["file"] = "File"
	fields["host_device"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["host_device"] = "HostDevice"
	fields["auto_detect"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["auto_detect"] = "AutoDetect"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
		map[string][]bindings.FieldData{
			"FILE": []bindings.FieldData{
				bindings.NewFieldData("file", true),
			},
			"HOST_DEVICE": []bindings.FieldData{
				bindings.NewFieldData("host_device", false),
				bindings.NewFieldData("auto_detect", true),
			},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.parallel.backing_info", fields, reflect.TypeOf(ParallelBackingInfo{}), fieldNameMap, validators)
}

func ParallelBackingSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.parallel.backing_type", reflect.TypeOf(ParallelBackingType(ParallelBackingType_FILE)))
	fieldNameMap["type"] = "Type_"
	fields["file"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["file"] = "File"
	fields["host_device"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["host_device"] = "HostDevice"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
		map[string][]bindings.FieldData{
			"FILE": []bindings.FieldData{
				bindings.NewFieldData("file", true),
			},
			"HOST_DEVICE": []bindings.FieldData{
				bindings.NewFieldData("host_device", false),
			},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.parallel.backing_spec", fields, reflect.TypeOf(ParallelBackingSpec{}), fieldNameMap, validators)
}

func ParallelInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["label"] = bindings.NewStringType()
	fieldNameMap["label"] = "Label"
	fields["backing"] = bindings.NewReferenceType(ParallelBackingInfoBindingType)
	fieldNameMap["backing"] = "Backing"
	fields["state"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.connection_state", reflect.TypeOf(ConnectionState(ConnectionState_CONNECTED)))
	fieldNameMap["state"] = "State"
	fields["start_connected"] = bindings.NewBooleanType()
	fieldNameMap["start_connected"] = "StartConnected"
	fields["allow_guest_control"] = bindings.NewBooleanType()
	fieldNameMap["allow_guest_control"] = "AllowGuestControl"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.parallel.info", fields, reflect.TypeOf(ParallelInfo{}), fieldNameMap, validators)
}

func ParallelCreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["backing"] = bindings.NewOptionalType(bindings.NewReferenceType(ParallelBackingSpecBindingType))
	fieldNameMap["backing"] = "Backing"
	fields["start_connected"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["start_connected"] = "StartConnected"
	fields["allow_guest_control"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["allow_guest_control"] = "AllowGuestControl"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.parallel.create_spec", fields, reflect.TypeOf(ParallelCreateSpec{}), fieldNameMap, validators)
}

func ParallelUpdateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["backing"] = bindings.NewOptionalType(bindings.NewReferenceType(ParallelBackingSpecBindingType))
	fieldNameMap["backing"] = "Backing"
	fields["start_connected"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["start_connected"] = "StartConnected"
	fields["allow_guest_control"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["allow_guest_control"] = "AllowGuestControl"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.parallel.update_spec", fields, reflect.TypeOf(ParallelUpdateSpec{}), fieldNameMap, validators)
}

func ParallelSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["port"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.ParallelPort"}, "")
	fieldNameMap["port"] = "Port"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.parallel.summary", fields, reflect.TypeOf(ParallelSummary{}), fieldNameMap, validators)
}

