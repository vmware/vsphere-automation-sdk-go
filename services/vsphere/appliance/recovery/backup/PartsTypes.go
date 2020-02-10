/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Parts.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package backup

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// ``LocalizableMessage`` class Structure representing message
type PartsLocalizableMessage struct {
    // id in message bundle
	Id string
    // text in english
	DefaultMessage string
    // nested data
	Args []string
}

// ``Part`` class Structure representing backup restore part
type PartsPart struct {
    // part ID
	Id string
    // part name id in message bundle
	Name PartsLocalizableMessage
    // part description id in message bundle
	Description PartsLocalizableMessage
    // Is this part selected by default in the user interface.
	SelectedByDefault bool
    // Is this part optional.
	Optional bool
}



func partsListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func partsListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(PartsPartBindingType), reflect.TypeOf([]PartsPart{}))
}

func partsListRestMetadata() protocol.OperationRestMetadata {
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

func partsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = bindings.NewIdType([]string{"com.vmware.appliance.recovery.backup.parts"}, "")
	fieldNameMap["id"] = "Id"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func partsGetOutputType() bindings.BindingType {
	return bindings.NewIntegerType()
}

func partsGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["id"] = bindings.NewIdType([]string{"com.vmware.appliance.recovery.backup.parts"}, "")
	fieldNameMap["id"] = "Id"
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


func PartsLocalizableMessageBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = bindings.NewStringType()
	fieldNameMap["id"] = "Id"
	fields["default_message"] = bindings.NewStringType()
	fieldNameMap["default_message"] = "DefaultMessage"
	fields["args"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["args"] = "Args"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.recovery.backup.parts.localizable_message", fields, reflect.TypeOf(PartsLocalizableMessage{}), fieldNameMap, validators)
}

func PartsPartBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = bindings.NewIdType([]string{"com.vmware.appliance.recovery.backup.parts"}, "")
	fieldNameMap["id"] = "Id"
	fields["name"] = bindings.NewReferenceType(PartsLocalizableMessageBindingType)
	fieldNameMap["name"] = "Name"
	fields["description"] = bindings.NewReferenceType(PartsLocalizableMessageBindingType)
	fieldNameMap["description"] = "Description"
	fields["selected_by_default"] = bindings.NewBooleanType()
	fieldNameMap["selected_by_default"] = "SelectedByDefault"
	fields["optional"] = bindings.NewBooleanType()
	fieldNameMap["optional"] = "Optional"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.recovery.backup.parts.part", fields, reflect.TypeOf(PartsPart{}), fieldNameMap, validators)
}

