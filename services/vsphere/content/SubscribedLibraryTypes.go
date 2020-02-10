/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: SubscribedLibrary.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package content

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/vsphere/content/library"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``ProbeResult`` class defines the subscription information probe result. This describes whether using a given subscription URL is successful or if there are access problems, such as SSL errors.
type SubscribedLibraryProbeResult struct {
    // The status of probe result. This will be one of SUCCESS, INVALID_URL, TIMED_OUT, HOST_NOT_FOUND, RESOURCE_NOT_FOUND, INVALID_CREDENTIALS, CERTIFICATE_ERROR, UNKNOWN_ERROR.
	Status SubscribedLibraryProbeResultStatus
    // The SSL thumbprint for the remote endpoint.
	SslThumbprint *string
    // If the probe result is in an error status, this property will contain the detailed error messages.
	ErrorMessages []std.LocalizableMessage
}

// The ``Status`` enumeration class defines the error status constants for the probe result.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type SubscribedLibraryProbeResultStatus string

const (
    // Indicates that the probe was successful.
	SubscribedLibraryProbeResultStatus_SUCCESS SubscribedLibraryProbeResultStatus = "SUCCESS"
    // Indicates that the supplied URL was not valid.
	SubscribedLibraryProbeResultStatus_INVALID_URL SubscribedLibraryProbeResultStatus = "INVALID_URL"
    // Indicates that the probe timed out while attempting to connect to the URL.
	SubscribedLibraryProbeResultStatus_TIMED_OUT SubscribedLibraryProbeResultStatus = "TIMED_OUT"
    // Indicates that the host in the URL could not be found.
	SubscribedLibraryProbeResultStatus_HOST_NOT_FOUND SubscribedLibraryProbeResultStatus = "HOST_NOT_FOUND"
    // Indicates that the given resource at the URL was not found.
	SubscribedLibraryProbeResultStatus_RESOURCE_NOT_FOUND SubscribedLibraryProbeResultStatus = "RESOURCE_NOT_FOUND"
    // Indicates that the connection was rejected due to invalid credentials.
	SubscribedLibraryProbeResultStatus_INVALID_CREDENTIALS SubscribedLibraryProbeResultStatus = "INVALID_CREDENTIALS"
    // Indicates that the provided server certificate thumbprint in library.SubscriptionInfo#sslThumbprint is invalid. In this case, the returned null should be set in library.SubscriptionInfo#sslThumbprint.
	SubscribedLibraryProbeResultStatus_CERTIFICATE_ERROR SubscribedLibraryProbeResultStatus = "CERTIFICATE_ERROR"
    // Indicates an unspecified error different from the other error cases defined in SubscribedLibraryProbeResultStatus.
	SubscribedLibraryProbeResultStatus_UNKNOWN_ERROR SubscribedLibraryProbeResultStatus = "UNKNOWN_ERROR"
)

func (s SubscribedLibraryProbeResultStatus) SubscribedLibraryProbeResultStatus() bool {
	switch s {
	case SubscribedLibraryProbeResultStatus_SUCCESS:
		return true
	case SubscribedLibraryProbeResultStatus_INVALID_URL:
		return true
	case SubscribedLibraryProbeResultStatus_TIMED_OUT:
		return true
	case SubscribedLibraryProbeResultStatus_HOST_NOT_FOUND:
		return true
	case SubscribedLibraryProbeResultStatus_RESOURCE_NOT_FOUND:
		return true
	case SubscribedLibraryProbeResultStatus_INVALID_CREDENTIALS:
		return true
	case SubscribedLibraryProbeResultStatus_CERTIFICATE_ERROR:
		return true
	case SubscribedLibraryProbeResultStatus_UNKNOWN_ERROR:
		return true
	default:
		return false
	}
}




func subscribedLibraryCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["client_token"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["create_spec"] = bindings.NewReferenceType(LibraryModelBindingType)
	fieldNameMap["client_token"] = "ClientToken"
	fieldNameMap["create_spec"] = "CreateSpec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func subscribedLibraryCreateOutputType() bindings.BindingType {
	return bindings.NewIdType([]string{"com.vmware.content.Library"}, "")
}

func subscribedLibraryCreateRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"InvalidArgument": 400,"Unsupported": 400,"ResourceInaccessible": 500})
}

func subscribedLibraryDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["library_id"] = bindings.NewIdType([]string{"com.vmware.content.Library"}, "")
	fieldNameMap["library_id"] = "LibraryId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func subscribedLibraryDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func subscribedLibraryDeleteRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"InvalidElementType": 400,"NotFound": 404})
}

func subscribedLibraryEvictInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["library_id"] = bindings.NewIdType([]string{"com.vmware.content.Library"}, "")
	fieldNameMap["library_id"] = "LibraryId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func subscribedLibraryEvictOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func subscribedLibraryEvictRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"NotFound": 404,"InvalidElementType": 400,"NotAllowedInCurrentState": 400})
}

func subscribedLibraryGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["library_id"] = bindings.NewIdType([]string{"com.vmware.content.Library"}, "")
	fieldNameMap["library_id"] = "LibraryId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func subscribedLibraryGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(LibraryModelBindingType)
}

func subscribedLibraryGetRestMetadata() protocol.OperationRestMetadata {
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

func subscribedLibraryListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func subscribedLibraryListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewIdType([]string{"com.vmware.content.Library"}, ""), reflect.TypeOf([]string{}))
}

func subscribedLibraryListRestMetadata() protocol.OperationRestMetadata {
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

func subscribedLibrarySyncInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["library_id"] = bindings.NewIdType([]string{"com.vmware.content.Library"}, "")
	fieldNameMap["library_id"] = "LibraryId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func subscribedLibrarySyncOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func subscribedLibrarySyncRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"NotFound": 404,"InvalidElementType": 400,"NotAllowedInCurrentState": 400,"InvalidArgument": 400,"ResourceInaccessible": 500})
}

func subscribedLibraryUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["library_id"] = bindings.NewIdType([]string{"com.vmware.content.Library"}, "")
	fields["update_spec"] = bindings.NewReferenceType(LibraryModelBindingType)
	fieldNameMap["library_id"] = "LibraryId"
	fieldNameMap["update_spec"] = "UpdateSpec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func subscribedLibraryUpdateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func subscribedLibraryUpdateRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"NotFound": 404,"NotAllowedInCurrentState": 400,"InvalidElementType": 400,"InvalidArgument": 400,"ResourceInaccessible": 500,"ResourceBusy": 500,"ConcurrentChange": 409})
}

func subscribedLibraryProbeInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["subscription_info"] = bindings.NewReferenceType(library.SubscriptionInfoBindingType)
	fieldNameMap["subscription_info"] = "SubscriptionInfo"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func subscribedLibraryProbeOutputType() bindings.BindingType {
	return bindings.NewReferenceType(SubscribedLibraryProbeResultBindingType)
}

func subscribedLibraryProbeRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["subscription_info"] = bindings.NewReferenceType(library.SubscriptionInfoBindingType)
	fieldNameMap["subscription_info"] = "SubscriptionInfo"
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


func SubscribedLibraryProbeResultBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["status"] = bindings.NewEnumType("com.vmware.content.subscribed_library.probe_result.status", reflect.TypeOf(SubscribedLibraryProbeResultStatus(SubscribedLibraryProbeResultStatus_SUCCESS)))
	fieldNameMap["status"] = "Status"
	fields["ssl_thumbprint"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["ssl_thumbprint"] = "SslThumbprint"
	fields["error_messages"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
	fieldNameMap["error_messages"] = "ErrorMessages"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.content.subscribed_library.probe_result", fields, reflect.TypeOf(SubscribedLibraryProbeResult{}), fieldNameMap, validators)
}

