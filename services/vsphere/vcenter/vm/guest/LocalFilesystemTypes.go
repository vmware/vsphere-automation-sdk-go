/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: LocalFilesystem.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package guest

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// Describes the virtual disk backing a local guest disk. This class was added in vSphere API 7.0.0.
type LocalFilesystemVirtualDiskMapping struct {
    // The virtual disk. This property was added in vSphere API 7.0.0.
	Disk string
}

// The ``Info`` class contains information about a local file system configured in the guest operating system. This class was added in vSphere API 6.7.
type LocalFilesystemInfo struct {
    // Total capacity of the file system, in bytes. This property was added in vSphere API 6.7.
	Capacity int64
    // Free space on the file system, in bytes. This property was added in vSphere API 6.7.
	FreeSpace int64
    // Filesystem type, if known. For example, ext3 or NTFS. This property was added in vSphere API 7.0.0.
	Filesystem *string
    // VirtualDisks backing the guest partition, if known. This property was added in vSphere API 7.0.0.
	Mappings []LocalFilesystemVirtualDiskMapping
}



func localFilesystemGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fieldNameMap["vm"] = "Vm"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func localFilesystemGetOutputType() bindings.BindingType {
	return bindings.NewMapType(bindings.NewStringType(), bindings.NewReferenceType(LocalFilesystemInfoBindingType),reflect.TypeOf(map[string]LocalFilesystemInfo{}))
}

func localFilesystemGetRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"Error": 500,"NotFound": 404,"ServiceUnavailable": 503})
}


func LocalFilesystemVirtualDiskMappingBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["disk"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Disk"}, "")
	fieldNameMap["disk"] = "Disk"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm.guest.local_filesystem.virtual_disk_mapping", fields, reflect.TypeOf(LocalFilesystemVirtualDiskMapping{}), fieldNameMap, validators)
}

func LocalFilesystemInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["capacity"] = bindings.NewIntegerType()
	fieldNameMap["capacity"] = "Capacity"
	fields["free_space"] = bindings.NewIntegerType()
	fieldNameMap["free_space"] = "FreeSpace"
	fields["filesystem"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["filesystem"] = "Filesystem"
	fields["mappings"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(LocalFilesystemVirtualDiskMappingBindingType), reflect.TypeOf([]LocalFilesystemVirtualDiskMapping{})))
	fieldNameMap["mappings"] = "Mappings"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm.guest.local_filesystem.info", fields, reflect.TypeOf(LocalFilesystemInfo{}), fieldNameMap, validators)
}

