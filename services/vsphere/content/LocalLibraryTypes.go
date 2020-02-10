/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: LocalLibrary.
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



// The ``DestinationSpec`` class contains information required to publish the library to a specific subscription. This class was added in vSphere API 6.7.2.
type LocalLibraryDestinationSpec struct {
    // Identifier of the subscription associated with the subscribed library. This property was added in vSphere API 6.7.2.
	Subscription string
}



func localLibraryCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["client_token"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["create_spec"] = bindings.NewReferenceType(LibraryModelBindingType)
	fieldNameMap["client_token"] = "ClientToken"
	fieldNameMap["create_spec"] = "CreateSpec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func localLibraryCreateOutputType() bindings.BindingType {
	return bindings.NewIdType([]string{"com.vmware.content.Library"}, "")
}

func localLibraryCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["client_token"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["create_spec"] = bindings.NewReferenceType(LibraryModelBindingType)
	fieldNameMap["client_token"] = "ClientToken"
	fieldNameMap["create_spec"] = "CreateSpec"
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
		map[string]int{"InvalidArgument": 400,"Unsupported": 400})
}

func localLibraryDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["library_id"] = bindings.NewIdType([]string{"com.vmware.content.Library"}, "")
	fieldNameMap["library_id"] = "LibraryId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func localLibraryDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func localLibraryDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["library_id"] = bindings.NewIdType([]string{"com.vmware.content.Library"}, "")
	fieldNameMap["library_id"] = "LibraryId"
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
		map[string]int{"InvalidElementType": 400,"NotFound": 404,"NotAllowedInCurrentState": 400})
}

func localLibraryGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["library_id"] = bindings.NewIdType([]string{"com.vmware.content.Library"}, "")
	fieldNameMap["library_id"] = "LibraryId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func localLibraryGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(LibraryModelBindingType)
}

func localLibraryGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["library_id"] = bindings.NewIdType([]string{"com.vmware.content.Library"}, "")
	fieldNameMap["library_id"] = "LibraryId"
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
		map[string]int{"NotFound": 404,"InvalidElementType": 400})
}

func localLibraryListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func localLibraryListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewIdType([]string{"com.vmware.content.Library"}, ""), reflect.TypeOf([]string{}))
}

func localLibraryListRestMetadata() protocol.OperationRestMetadata {
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

func localLibraryUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["library_id"] = bindings.NewIdType([]string{"com.vmware.content.Library"}, "")
	fields["update_spec"] = bindings.NewReferenceType(LibraryModelBindingType)
	fieldNameMap["library_id"] = "LibraryId"
	fieldNameMap["update_spec"] = "UpdateSpec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func localLibraryUpdateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func localLibraryUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["library_id"] = bindings.NewIdType([]string{"com.vmware.content.Library"}, "")
	fields["update_spec"] = bindings.NewReferenceType(LibraryModelBindingType)
	fieldNameMap["library_id"] = "LibraryId"
	fieldNameMap["update_spec"] = "UpdateSpec"
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
		map[string]int{"NotFound": 404,"NotAllowedInCurrentState": 400,"InvalidElementType": 400,"InvalidArgument": 400,"ResourceBusy": 500,"ConcurrentChange": 409})
}

func localLibraryPublishInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["library_id"] = bindings.NewIdType([]string{"com.vmware.content.Library"}, "")
	fields["subscriptions"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(LocalLibraryDestinationSpecBindingType), reflect.TypeOf([]LocalLibraryDestinationSpec{})))
	fieldNameMap["library_id"] = "LibraryId"
	fieldNameMap["subscriptions"] = "Subscriptions"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func localLibraryPublishOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func localLibraryPublishRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["library_id"] = bindings.NewIdType([]string{"com.vmware.content.Library"}, "")
	fields["subscriptions"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(LocalLibraryDestinationSpecBindingType), reflect.TypeOf([]LocalLibraryDestinationSpec{})))
	fieldNameMap["library_id"] = "LibraryId"
	fieldNameMap["subscriptions"] = "Subscriptions"
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
		map[string]int{"Error": 500,"NotFound": 404,"InvalidArgument": 400,"InvalidElementType": 400,"NotAllowedInCurrentState": 400,"Unauthenticated": 401,"Unauthorized": 403})
}


func LocalLibraryDestinationSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["subscription"] = bindings.NewIdType([]string{"com.vmware.content.library.Subscriptions"}, "")
	fieldNameMap["subscription"] = "Subscription"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.content.local_library.destination_spec", fields, reflect.TypeOf(LocalLibraryDestinationSpec{}), fieldNameMap, validators)
}

