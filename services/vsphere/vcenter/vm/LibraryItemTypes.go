/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: LibraryItem.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package vm

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``Info`` class contains information about the library item associated with a virtual machine. This class was added in vSphere API 6.9.1.
type LibraryItemInfo struct {
    // Information about the checked out virtual machine. This property was added in vSphere API 6.9.1.
	CheckOut *LibraryItemCheckOutInfo
}

// The ``CheckOutInfo`` class contains information about a virtual machine checked out of a content library item. This class was added in vSphere API 6.9.1.
type LibraryItemCheckOutInfo struct {
    // Identifier of the library item that the virtual machine is checked out from. This property was added in vSphere API 6.9.1.
	LibraryItem string
}



func libraryItemGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fieldNameMap["vm"] = "Vm"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func libraryItemGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(LibraryItemInfoBindingType)
}

func libraryItemGetRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}


func LibraryItemInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["check_out"] = bindings.NewOptionalType(bindings.NewReferenceType(LibraryItemCheckOutInfoBindingType))
	fieldNameMap["check_out"] = "CheckOut"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm.library_item.info", fields, reflect.TypeOf(LibraryItemInfo{}), fieldNameMap, validators)
}

func LibraryItemCheckOutInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["library_item"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fieldNameMap["library_item"] = "LibraryItem"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm.library_item.check_out_info", fields, reflect.TypeOf(LibraryItemCheckOutInfo{}), fieldNameMap, validators)
}

