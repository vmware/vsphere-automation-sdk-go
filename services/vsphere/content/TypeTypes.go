/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Type.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package content

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``Info`` class describes support for a specific type of data in an library.ItemModel. The ``Info`` can be queried through the Type interface. Type support describes plugins in the Content Library which can provide metadata on library items and help manage the transfer process by adding dependent files when a current file is added.
type TypeInfo struct {
    // A description of the type support offered by the plugin.
	Description string
    // The name of the plugin which provides the type support.
	Name string
    // The type which the plugin supports. 
    //
    //  To upload a library item of the type supported by the plugin, the library.ItemModel#type property of the item should be set to this value.
	Type_ string
    // The name of the vendor who created the type support plugin.
	Vendor string
    // The version number of the type support plugin.
	Version string
}



func typeListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func typeListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(TypeInfoBindingType), reflect.TypeOf([]TypeInfo{}))
}

func typeListRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{})
}


func TypeInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["description"] = bindings.NewStringType()
	fieldNameMap["description"] = "Description"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["type"] = bindings.NewStringType()
	fieldNameMap["type"] = "Type_"
	fields["vendor"] = bindings.NewStringType()
	fieldNameMap["vendor"] = "Vendor"
	fields["version"] = bindings.NewStringType()
	fieldNameMap["version"] = "Version"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.content.type.info", fields, reflect.TypeOf(TypeInfo{}), fieldNameMap, validators)
}

