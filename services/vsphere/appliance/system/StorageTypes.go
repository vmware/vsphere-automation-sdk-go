/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Storage.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package system

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``StorageMapping`` class describes the mapping between VCSA partitions and the Hard disk numbers visible in the vSphere Web Client.
type StorageStorageMapping struct {
    // The disk number in the vSphere Web Client.
	Disk string
    // Storage partition name.
	Partition string
    // Description of partition. This property was added in vSphere API 6.7.
	Description *std.LocalizableMessage
}

// The ``StorageChange`` class describes the changes in capasity of a storage partition. This class was added in vSphere API 6.7.
type StorageStorageChange struct {
    // Original size of the partition in MB. This property was added in vSphere API 6.7.
	OldSize int64
    // Nedw size of the partition in MB. This property was added in vSphere API 6.7.
	NewSize int64
}



func storageListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func storageListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(StorageStorageMappingBindingType), reflect.TypeOf([]StorageStorageMapping{}))
}

func storageListRestMetadata() protocol.OperationRestMetadata {
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

func storageResizeInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func storageResizeOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func storageResizeRestMetadata() protocol.OperationRestMetadata {
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

func storageResizeExInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func storageResizeExOutputType() bindings.BindingType {
	return bindings.NewMapType(bindings.NewStringType(), bindings.NewReferenceType(StorageStorageChangeBindingType),reflect.TypeOf(map[string]StorageStorageChange{}))
}

func storageResizeExRestMetadata() protocol.OperationRestMetadata {
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


func StorageStorageMappingBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["disk"] = bindings.NewIdType([]string{"com.vmware.appliance.system.storage"}, "")
	fieldNameMap["disk"] = "Disk"
	fields["partition"] = bindings.NewStringType()
	fieldNameMap["partition"] = "Partition"
	fields["description"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
	fieldNameMap["description"] = "Description"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.system.storage.storage_mapping", fields, reflect.TypeOf(StorageStorageMapping{}), fieldNameMap, validators)
}

func StorageStorageChangeBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["old_size"] = bindings.NewIntegerType()
	fieldNameMap["old_size"] = "OldSize"
	fields["new_size"] = bindings.NewIntegerType()
	fieldNameMap["new_size"] = "NewSize"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.system.storage.storage_change", fields, reflect.TypeOf(StorageStorageChange{}), fieldNameMap, validators)
}

