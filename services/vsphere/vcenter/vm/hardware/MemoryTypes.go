/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Memory.
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



// The ``Info`` class contains memory-related information about a virtual machine.
type MemoryInfo struct {
    // Memory size in mebibytes.
	SizeMiB int64
    // Flag indicating whether adding memory while the virtual machine is running is enabled. 
    //
    //  Some guest operating systems may consume more resources or perform less efficiently when they run on hardware that supports adding memory while the machine is running.
	HotAddEnabled bool
    // The granularity, in mebibytes, at which memory can be added to a running virtual machine. 
    //
    //  When adding memory to a running virtual machine, the amount of memory added must be at least MemoryInfo#hotAddIncrementSizeMiB and the total memory size of the virtual machine must be a multiple of {\\\\@link>hotAddIncrementSize}.
	HotAddIncrementSizeMiB *int64
    // The maximum amount of memory, in mebibytes, that can be added to a running virtual machine.
	HotAddLimitMiB *int64
}

// The ``UpdateSpec`` class describes the updates to be made to the memory-related settings of a virtual machine.
type MemoryUpdateSpec struct {
    // New memory size in mebibytes. 
    //
    //  The supported range of memory sizes is constrained by the configured guest operating system and virtual hardware version of the virtual machine. 
    //
    //  If the virtual machine is running, this value may only be changed if MemoryInfo#hotAddEnabled is true, and the new memory size must satisfy the constraints specified by MemoryInfo#hotAddIncrementSizeMiB and MemoryInfo#hotAddLimitMiB.
	SizeMiB *int64
    // Flag indicating whether adding memory while the virtual machine is running should be enabled. 
    //
    //  Some guest operating systems may consume more resources or perform less efficiently when they run on hardware that supports adding memory while the machine is running. 
    //
    //  This property may only be modified if the virtual machine is not powered on.
	HotAddEnabled *bool
}



func memoryGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fieldNameMap["vm"] = "Vm"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func memoryGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(MemoryInfoBindingType)
}

func memoryGetRestMetadata() protocol.OperationRestMetadata {
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

func memoryUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["spec"] = bindings.NewReferenceType(MemoryUpdateSpecBindingType)
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func memoryUpdateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func memoryUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["spec"] = bindings.NewReferenceType(MemoryUpdateSpecBindingType)
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
		map[string]int{"Error": 500,"NotFound": 404,"InvalidArgument": 400,"NotAllowedInCurrentState": 400,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func MemoryInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["size_MiB"] = bindings.NewIntegerType()
	fieldNameMap["size_MiB"] = "SizeMiB"
	fields["hot_add_enabled"] = bindings.NewBooleanType()
	fieldNameMap["hot_add_enabled"] = "HotAddEnabled"
	fields["hot_add_increment_size_MiB"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["hot_add_increment_size_MiB"] = "HotAddIncrementSizeMiB"
	fields["hot_add_limit_MiB"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["hot_add_limit_MiB"] = "HotAddLimitMiB"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.memory.info", fields, reflect.TypeOf(MemoryInfo{}), fieldNameMap, validators)
}

func MemoryUpdateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["size_MiB"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["size_MiB"] = "SizeMiB"
	fields["hot_add_enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["hot_add_enabled"] = "HotAddEnabled"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.memory.update_spec", fields, reflect.TypeOf(MemoryUpdateSpec{}), fieldNameMap, validators)
}

