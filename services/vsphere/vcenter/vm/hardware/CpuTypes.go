/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Cpu.
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



// The ``Info`` class contains CPU-related information about a virtual machine.
type CpuInfo struct {
    // Number of CPU cores.
	Count int64
    // Number of CPU cores per socket.
	CoresPerSocket int64
    // Flag indicating whether adding CPUs while the virtual machine is running is enabled.
	HotAddEnabled bool
    // Flag indicating whether removing CPUs while the virtual machine is running is enabled.
	HotRemoveEnabled bool
}

// The ``UpdateSpec`` class describes the updates to be made to the CPU-related settings of a virtual machine.
type CpuUpdateSpec struct {
    // New number of CPU cores. The number of CPU cores in the virtual machine must be a multiple of the number of cores per socket. 
    //
    //  The supported range of CPU counts is constrained by the configured guest operating system and virtual hardware version of the virtual machine. 
    //
    //  If the virtual machine is running, the number of CPU cores may only be increased if CpuInfo#hotAddEnabled is true, and may only be decreased if CpuInfo#hotRemoveEnabled is true.
	Count *int64
    // New number of CPU cores per socket. The number of CPU cores in the virtual machine must be a multiple of the number of cores per socket.
	CoresPerSocket *int64
    // Flag indicating whether adding CPUs while the virtual machine is running is enabled. 
    //
    //  This property may only be modified if the virtual machine is powered off.
	HotAddEnabled *bool
    // Flag indicating whether removing CPUs while the virtual machine is running is enabled. 
    //
    //  This property may only be modified if the virtual machine is powered off.
	HotRemoveEnabled *bool
}



func cpuGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fieldNameMap["vm"] = "Vm"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cpuGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(CpuInfoBindingType)
}

func cpuGetRestMetadata() protocol.OperationRestMetadata {
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

func cpuUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["spec"] = bindings.NewReferenceType(CpuUpdateSpecBindingType)
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cpuUpdateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func cpuUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["spec"] = bindings.NewReferenceType(CpuUpdateSpecBindingType)
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


func CpuInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["count"] = bindings.NewIntegerType()
	fieldNameMap["count"] = "Count"
	fields["cores_per_socket"] = bindings.NewIntegerType()
	fieldNameMap["cores_per_socket"] = "CoresPerSocket"
	fields["hot_add_enabled"] = bindings.NewBooleanType()
	fieldNameMap["hot_add_enabled"] = "HotAddEnabled"
	fields["hot_remove_enabled"] = bindings.NewBooleanType()
	fieldNameMap["hot_remove_enabled"] = "HotRemoveEnabled"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.cpu.info", fields, reflect.TypeOf(CpuInfo{}), fieldNameMap, validators)
}

func CpuUpdateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["count"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["count"] = "Count"
	fields["cores_per_socket"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["cores_per_socket"] = "CoresPerSocket"
	fields["hot_add_enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["hot_add_enabled"] = "HotAddEnabled"
	fields["hot_remove_enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["hot_remove_enabled"] = "HotRemoveEnabled"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.cpu.update_spec", fields, reflect.TypeOf(CpuUpdateSpec{}), fieldNameMap, validators)
}

