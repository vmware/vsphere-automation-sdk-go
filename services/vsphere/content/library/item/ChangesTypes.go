/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Changes.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package item

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
	"time"
)

// Resource type for library item versions. This constant field was added in vSphere API 6.9.1.
const Changes_RESOURCE_TYPE = "com.vmware.content.library.item.Version"


// The ``Summary`` class contains commonly used information about a library item change. This class was added in vSphere API 6.9.1.
type ChangesSummary struct {
    // The version of the library item. This property was added in vSphere API 6.9.1.
	Version string
    // The date and time when the item content was changed. This property was added in vSphere API 6.9.1.
	Time time.Time
    // The user who made the content change. This property was added in vSphere API 6.9.1.
	User *string
    // The short message describing the content change. The message is truncated to the first 80 characters or first non-leading newline character, whichever length is shorter. This property was added in vSphere API 6.9.1.
	ShortMessage *string
}

// The ``Info`` class contains information about a library item change. This class was added in vSphere API 6.9.1.
type ChangesInfo struct {
    // The date and time when the item content was changed. This property was added in vSphere API 6.9.1.
	Time time.Time
    // The user who made the content change. This property was added in vSphere API 6.9.1.
	User *string
    // The full message describing the content change. This property was added in vSphere API 6.9.1.
	Message *string
}



func changesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["library_item"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fieldNameMap["library_item"] = "LibraryItem"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func changesListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(ChangesSummaryBindingType), reflect.TypeOf([]ChangesSummary{}))
}

func changesListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["library_item"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fieldNameMap["library_item"] = "LibraryItem"
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
		map[string]int{"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}

func changesGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["library_item"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fields["version"] = bindings.NewIdType([]string{"com.vmware.content.library.item.Version"}, "")
	fieldNameMap["library_item"] = "LibraryItem"
	fieldNameMap["version"] = "Version"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func changesGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(ChangesInfoBindingType)
}

func changesGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["library_item"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fields["version"] = bindings.NewIdType([]string{"com.vmware.content.library.item.Version"}, "")
	fieldNameMap["library_item"] = "LibraryItem"
	fieldNameMap["version"] = "Version"
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
		map[string]int{"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}


func ChangesSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["version"] = bindings.NewIdType([]string{"com.vmware.content.library.item.Version"}, "")
	fieldNameMap["version"] = "Version"
	fields["time"] = bindings.NewDateTimeType()
	fieldNameMap["time"] = "Time"
	fields["user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["user"] = "User"
	fields["short_message"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["short_message"] = "ShortMessage"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.content.library.item.changes.summary", fields, reflect.TypeOf(ChangesSummary{}), fieldNameMap, validators)
}

func ChangesInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["time"] = bindings.NewDateTimeType()
	fieldNameMap["time"] = "Time"
	fields["user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["user"] = "User"
	fields["message"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["message"] = "Message"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.content.library.item.changes.info", fields, reflect.TypeOf(ChangesInfo{}), fieldNameMap, validators)
}

