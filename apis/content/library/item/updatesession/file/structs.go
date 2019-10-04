/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: File.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package file

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/content/library/item"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/content/library/item/file"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``SourceType`` enumeration class defines how the file content is retrieved.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type SourceType string

const (
    // No source type has been requested.
     SourceType_NONE SourceType = "NONE"
    // The client is uploading content using HTTP(S) PUT requests.
     SourceType_PUSH SourceType = "PUSH"
    // The server is pulling content from a URL. The URL scheme can be ``http``, ``https``, ``file``, or ``ds``.
     SourceType_PULL SourceType = "PULL"
)

func (s SourceType) SourceType() bool {
    switch s {
        case SourceType_NONE:
            return true
        case SourceType_PUSH:
            return true
        case SourceType_PULL:
            return true
        default:
            return false
    }
}





// The ``AddSpec`` class describes the properties of the file to be uploaded.
 type AddSpec struct {
    Name string
    SourceType SourceType
    SourceEndpoint *item.TransferEndpoint
    Size *int64
    ChecksumInfo *file.ChecksumInfo
}






// The ``Info`` class defines the uploaded file.
 type Info struct {
    Name string
    SourceType SourceType
    Size *int64
    ChecksumInfo *file.ChecksumInfo
    SourceEndpoint *item.TransferEndpoint
    UploadEndpoint *item.TransferEndpoint
    BytesTransferred int64
    Status item.TransferStatus
    ErrorMessage *std.LocalizableMessage
    KeepInStorage *bool
}






// The ``ValidationError`` class defines the validation error of a file in the session.
 type ValidationError struct {
    Name string
    ErrorMessage std.LocalizableMessage
}






// The ``ValidationResult`` class defines the result of validating the files in the session.
 type ValidationResult struct {
    HasErrors bool
    MissingFiles map[string]bool
    InvalidFiles []ValidationError
}






// The ``ProbeResult`` class contains information about the accessibility of a source endpoint URI. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ProbeResult struct {
    Status ProbeResult_Status
    SslThumbprint *string
    ErrorMessages []std.LocalizableMessage
}




    
    // The ``Status`` enumeration class defines the possible status values from an attempt to access a source endpoint URI. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type ProbeResult_Status string

    const (
        // Indicates that the probe was successful. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         ProbeResult_Status_SUCCESS ProbeResult_Status = "SUCCESS"
        // Indicates that the supplied URL was not valid. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         ProbeResult_Status_INVALID_URL ProbeResult_Status = "INVALID_URL"
        // Indicates that the probe timed out while attempting to connect to the URL. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         ProbeResult_Status_TIMED_OUT ProbeResult_Status = "TIMED_OUT"
        // Indicates that the host in the URL could not be found. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         ProbeResult_Status_HOST_NOT_FOUND ProbeResult_Status = "HOST_NOT_FOUND"
        // Indicates that the provided server certificate thumbprint in item.TransferEndpoint#sslCertificateThumbprint is invalid. In this case, the returned null should be set in item.TransferEndpoint#sslCertificateThumbprint. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         ProbeResult_Status_CERTIFICATE_ERROR ProbeResult_Status = "CERTIFICATE_ERROR"
        // Indicates an unspecified error different from the other error cases defined in ProbeResult_Status. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
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
            case ProbeResult_Status_CERTIFICATE_ERROR:
                return true
            case ProbeResult_Status_UNKNOWN_ERROR:
                return true
            default:
                return false
        }
    }






func validateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["update_session_id"] = bindings.NewIdType([]string {"com.vmware.content.library.item.UpdateSession"}, "")
    fieldNameMap["update_session_id"] = "UpdateSessionId"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func validateOutputType() bindings.BindingType {
    return bindings.NewReferenceType(ValidationResultBindingType)
}

func validateRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404,"NotAllowedInCurrentState": 400})
}


func addInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["update_session_id"] = bindings.NewIdType([]string {"com.vmware.content.library.item.UpdateSession"}, "")
    fields["file_spec"] = bindings.NewReferenceType(AddSpecBindingType)
    fieldNameMap["update_session_id"] = "UpdateSessionId"
    fieldNameMap["file_spec"] = "FileSpec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func addOutputType() bindings.BindingType {
    return bindings.NewReferenceType(InfoBindingType)
}

func addRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"InvalidArgument": 400,"NotFound": 404,"Unauthorized": 403,"NotAllowedInCurrentState": 400})
}


func removeInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["update_session_id"] = bindings.NewIdType([]string {"com.vmware.content.library.item.UpdateSession"}, "")
    fields["file_name"] = bindings.NewStringType()
    fieldNameMap["update_session_id"] = "UpdateSessionId"
    fieldNameMap["file_name"] = "FileName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func removeOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func removeRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404,"InvalidArgument": 400})
}


func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["update_session_id"] = bindings.NewIdType([]string {"com.vmware.content.library.item.UpdateSession"}, "")
    fieldNameMap["update_session_id"] = "UpdateSessionId"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(InfoBindingType), reflect.TypeOf([]Info{}))
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
       map[string]int{"NotFound": 404})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["update_session_id"] = bindings.NewIdType([]string {"com.vmware.content.library.item.UpdateSession"}, "")
    fields["file_name"] = bindings.NewStringType()
    fieldNameMap["update_session_id"] = "UpdateSessionId"
    fieldNameMap["file_name"] = "FileName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewReferenceType(InfoBindingType)
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
       map[string]int{"NotFound": 404,"InvalidArgument": 400})
}


func probeInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["source_endpoint"] = bindings.NewReferenceType(item.TransferEndpointBindingType)
    fieldNameMap["source_endpoint"] = "SourceEndpoint"
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



func AddSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["source_type"] = bindings.NewEnumType("com.vmware.content.library.item.updatesession.file.source_type", reflect.TypeOf(SourceType(SourceType_NONE)))
    fieldNameMap["source_type"] = "SourceType"
    fields["source_endpoint"] = bindings.NewOptionalType(bindings.NewReferenceType(item.TransferEndpointBindingType))
    fieldNameMap["source_endpoint"] = "SourceEndpoint"
    fields["size"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["size"] = "Size"
    fields["checksum_info"] = bindings.NewOptionalType(bindings.NewReferenceType(file.ChecksumInfoBindingType))
    fieldNameMap["checksum_info"] = "ChecksumInfo"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("source_type",
        map[string][]bindings.FieldData {
            "PULL": []bindings.FieldData {
                 bindings.NewFieldData("source_endpoint", true),
            },
            "NONE": []bindings.FieldData {},
            "PUSH": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.content.library.item.updatesession.file.add_spec",fields, reflect.TypeOf(AddSpec{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["source_type"] = bindings.NewEnumType("com.vmware.content.library.item.updatesession.file.source_type", reflect.TypeOf(SourceType(SourceType_NONE)))
    fieldNameMap["source_type"] = "SourceType"
    fields["size"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["size"] = "Size"
    fields["checksum_info"] = bindings.NewOptionalType(bindings.NewReferenceType(file.ChecksumInfoBindingType))
    fieldNameMap["checksum_info"] = "ChecksumInfo"
    fields["source_endpoint"] = bindings.NewOptionalType(bindings.NewReferenceType(item.TransferEndpointBindingType))
    fieldNameMap["source_endpoint"] = "SourceEndpoint"
    fields["upload_endpoint"] = bindings.NewOptionalType(bindings.NewReferenceType(item.TransferEndpointBindingType))
    fieldNameMap["upload_endpoint"] = "UploadEndpoint"
    fields["bytes_transferred"] = bindings.NewIntegerType()
    fieldNameMap["bytes_transferred"] = "BytesTransferred"
    fields["status"] = bindings.NewEnumType("com.vmware.content.library.item.transfer_status", reflect.TypeOf(item.TransferStatus(item.TransferStatus_WAITING_FOR_TRANSFER)))
    fieldNameMap["status"] = "Status"
    fields["error_message"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
    fieldNameMap["error_message"] = "ErrorMessage"
    fields["keep_in_storage"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["keep_in_storage"] = "KeepInStorage"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("source_type",
        map[string][]bindings.FieldData {
            "PULL": []bindings.FieldData {
                 bindings.NewFieldData("source_endpoint", true),
            },
            "PUSH": []bindings.FieldData {
                 bindings.NewFieldData("upload_endpoint", true),
            },
            "NONE": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.content.library.item.updatesession.file.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func ValidationErrorBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["error_message"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
    fieldNameMap["error_message"] = "ErrorMessage"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.content.library.item.updatesession.file.validation_error",fields, reflect.TypeOf(ValidationError{}), fieldNameMap, validators)
}

func ValidationResultBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["has_errors"] = bindings.NewBooleanType()
    fieldNameMap["has_errors"] = "HasErrors"
    fields["missing_files"] = bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{}))
    fieldNameMap["missing_files"] = "MissingFiles"
    fields["invalid_files"] = bindings.NewListType(bindings.NewReferenceType(ValidationErrorBindingType), reflect.TypeOf([]ValidationError{}))
    fieldNameMap["invalid_files"] = "InvalidFiles"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.content.library.item.updatesession.file.validation_result",fields, reflect.TypeOf(ValidationResult{}), fieldNameMap, validators)
}

func ProbeResultBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["status"] = bindings.NewEnumType("com.vmware.content.library.item.updatesession.file.probe_result.status", reflect.TypeOf(ProbeResult_Status(ProbeResult_Status_SUCCESS)))
    fieldNameMap["status"] = "Status"
    fields["ssl_thumbprint"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["ssl_thumbprint"] = "SslThumbprint"
    fields["error_messages"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{})))
    fieldNameMap["error_messages"] = "ErrorMessages"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("status",
        map[string][]bindings.FieldData {
            "CERTIFICATE_ERROR": []bindings.FieldData {
                 bindings.NewFieldData("ssl_thumbprint", false),
                 bindings.NewFieldData("error_messages", true),
            },
            "SUCCESS": []bindings.FieldData {
                 bindings.NewFieldData("ssl_thumbprint", false),
            },
            "UNKNOWN_ERROR": []bindings.FieldData {
                 bindings.NewFieldData("ssl_thumbprint", false),
                 bindings.NewFieldData("error_messages", true),
            },
            "INVALID_URL": []bindings.FieldData {
                 bindings.NewFieldData("error_messages", true),
            },
            "TIMED_OUT": []bindings.FieldData {
                 bindings.NewFieldData("error_messages", true),
            },
            "HOST_NOT_FOUND": []bindings.FieldData {
                 bindings.NewFieldData("error_messages", true),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.content.library.item.updatesession.file.probe_result",fields, reflect.TypeOf(ProbeResult{}), fieldNameMap, validators)
}


