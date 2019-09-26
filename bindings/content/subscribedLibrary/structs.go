/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: SubscribedLibrary.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package subscribedLibrary

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/content"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/content/library"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)




// The ``ProbeResult`` class defines the subscription information probe result. This describes whether using a given subscription URL is successful or if there are access problems, such as SSL errors.
 type ProbeResult struct {
    Status ProbeResult_Status
    SslThumbprint *string
    ErrorMessages []std.LocalizableMessage
}




    
    // The ``Status`` enumeration class defines the error status constants for the probe result.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type ProbeResult_Status string

    const (
        // Indicates that the probe was successful.
         ProbeResult_Status_SUCCESS ProbeResult_Status = "SUCCESS"
        // Indicates that the supplied URL was not valid.
         ProbeResult_Status_INVALID_URL ProbeResult_Status = "INVALID_URL"
        // Indicates that the probe timed out while attempting to connect to the URL.
         ProbeResult_Status_TIMED_OUT ProbeResult_Status = "TIMED_OUT"
        // Indicates that the host in the URL could not be found.
         ProbeResult_Status_HOST_NOT_FOUND ProbeResult_Status = "HOST_NOT_FOUND"
        // Indicates that the given resource at the URL was not found.
         ProbeResult_Status_RESOURCE_NOT_FOUND ProbeResult_Status = "RESOURCE_NOT_FOUND"
        // Indicates that the connection was rejected due to invalid credentials.
         ProbeResult_Status_INVALID_CREDENTIALS ProbeResult_Status = "INVALID_CREDENTIALS"
        // Indicates that the provided server certificate thumbprint in library.SubscriptionInfo#sslThumbprint is invalid. In this case, the returned null should be set in library.SubscriptionInfo#sslThumbprint.
         ProbeResult_Status_CERTIFICATE_ERROR ProbeResult_Status = "CERTIFICATE_ERROR"
        // Indicates an unspecified error different from the other error cases defined in ProbeResult_Status.
         ProbeResult_Status_UNKNOWN_ERROR ProbeResult_Status = "UNKNOWN_ERROR"
    )

    func (s ProbeResult_Status) ProbeResult_Status() bool {
        switch s {
            case ProbeResult_Status_SUCCESS:
                return true
            case ProbeResult_Status_INVALID_URL:
                return true
            case ProbeResult_Status_TIMED_OUT:
                return true
            case ProbeResult_Status_HOST_NOT_FOUND:
                return true
            case ProbeResult_Status_RESOURCE_NOT_FOUND:
                return true
            case ProbeResult_Status_INVALID_CREDENTIALS:
                return true
            case ProbeResult_Status_CERTIFICATE_ERROR:
                return true
            case ProbeResult_Status_UNKNOWN_ERROR:
                return true
            default:
                return false
        }
    }






func createInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["client_token"] = bindings.NewOptionalType(bindings.NewStringType())
    fields["create_spec"] = bindings.NewReferenceType(content.LibraryModelBindingType)
    fieldNameMap["client_token"] = "ClientToken"
    fieldNameMap["create_spec"] = "CreateSpec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func createOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"com.vmware.content.Library"}, "")
}

func createRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"InvalidArgument": 400,"Unsupported": 400,"ResourceInaccessible": 500})
}


func deleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["library_id"] = bindings.NewIdType([]string {"com.vmware.content.Library"}, "")
    fieldNameMap["library_id"] = "LibraryId"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func deleteOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func deleteRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"InvalidElementType": 400,"NotFound": 404})
}


func evictInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["library_id"] = bindings.NewIdType([]string {"com.vmware.content.Library"}, "")
    fieldNameMap["library_id"] = "LibraryId"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func evictOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func evictRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"NotFound": 404,"InvalidElementType": 400,"NotAllowedInCurrentState": 400})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["library_id"] = bindings.NewIdType([]string {"com.vmware.content.Library"}, "")
    fieldNameMap["library_id"] = "LibraryId"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewReferenceType(content.LibraryModelBindingType)
}

func getRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"NotFound": 404,"InvalidElementType": 400})
}


func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewIdType([]string {"com.vmware.content.Library"}, ""), reflect.TypeOf([]string{}))
}

func listRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{})
}


func syncInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["library_id"] = bindings.NewIdType([]string {"com.vmware.content.Library"}, "")
    fieldNameMap["library_id"] = "LibraryId"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func syncOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func syncRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"NotFound": 404,"InvalidElementType": 400,"NotAllowedInCurrentState": 400,"InvalidArgument": 400,"ResourceInaccessible": 500})
}


func updateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["library_id"] = bindings.NewIdType([]string {"com.vmware.content.Library"}, "")
    fields["update_spec"] = bindings.NewReferenceType(content.LibraryModelBindingType)
    fieldNameMap["library_id"] = "LibraryId"
    fieldNameMap["update_spec"] = "UpdateSpec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func updateOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func updateRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"NotFound": 404,"NotAllowedInCurrentState": 400,"InvalidElementType": 400,"InvalidArgument": 400,"ResourceInaccessible": 500,"ResourceBusy": 500,"ConcurrentChange": 409})
}


func probeInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["subscription_info"] = bindings.NewReferenceType(library.SubscriptionInfoBindingType)
    fieldNameMap["subscription_info"] = "SubscriptionInfo"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func probeOutputType() bindings.BindingType {
    return bindings.NewReferenceType(ProbeResultBindingType)
}

func probeRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{})
}



func ProbeResultBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["status"] = bindings.NewEnumType("com.vmware.content.subscribed_library.probe_result.status", reflect.TypeOf(ProbeResult_Status(ProbeResult_Status_SUCCESS)))
    fieldNameMap["status"] = "Status"
    fields["ssl_thumbprint"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["ssl_thumbprint"] = "SslThumbprint"
    fields["error_messages"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["error_messages"] = "ErrorMessages"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.content.subscribed_library.probe_result",fields, reflect.TypeOf(ProbeResult{}), fieldNameMap, validators)
}


