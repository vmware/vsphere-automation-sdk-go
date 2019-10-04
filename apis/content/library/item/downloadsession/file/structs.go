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



// The ``PrepareStatus`` enumeration class defines the state of the file in preparation for download.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type PrepareStatus string

const (
    // The file hasn't been requested for preparation.
     PrepareStatus_UNPREPARED PrepareStatus = "UNPREPARED"
    // A prepare has been requested, however the server hasn't started the preparation yet.
     PrepareStatus_PREPARE_REQUESTED PrepareStatus = "PREPARE_REQUESTED"
    // A prepare has been requested and the file is in the process of being prepared.
     PrepareStatus_PREPARING PrepareStatus = "PREPARING"
    // Prepare succeeded. The file is ready for download.
     PrepareStatus_PREPARED PrepareStatus = "PREPARED"
    // Prepare failed.
     PrepareStatus_ERROR PrepareStatus = "ERROR"
)

func (p PrepareStatus) PrepareStatus() bool {
    switch p {
        case PrepareStatus_UNPREPARED:
            return true
        case PrepareStatus_PREPARE_REQUESTED:
            return true
        case PrepareStatus_PREPARING:
            return true
        case PrepareStatus_PREPARED:
            return true
        case PrepareStatus_ERROR:
            return true
        default:
            return false
    }
}




// The ``EndpointType`` enumeration class defines the types of endpoints used to download the file.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type EndpointType string

const (
    // An https download endpoint.
     EndpointType_HTTPS EndpointType = "HTTPS"
    // A direct download endpoint indicating the location of the file on storage. The caller is responsible for retrieving the file from the storage location directly.
     EndpointType_DIRECT EndpointType = "DIRECT"
)

func (e EndpointType) EndpointType() bool {
    switch e {
        case EndpointType_HTTPS:
            return true
        case EndpointType_DIRECT:
            return true
        default:
            return false
    }
}





// The ``Info`` class defines the downloaded file.
 type Info struct {
    Name string
    Size *int64
    BytesTransferred int64
    Status PrepareStatus
    DownloadEndpoint *item.TransferEndpoint
    ChecksumInfo *file.ChecksumInfo
    ErrorMessage *std.LocalizableMessage
}









func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["download_session_id"] = bindings.NewIdType([]string {"com.vmware.content.library.item.DownloadSession"}, "")
    fieldNameMap["download_session_id"] = "DownloadSessionId"
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


func prepareInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["download_session_id"] = bindings.NewIdType([]string {"com.vmware.content.library.item.DownloadSession"}, "")
    fields["file_name"] = bindings.NewStringType()
    fields["endpoint_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.content.library.item.downloadsession.file.endpoint_type", reflect.TypeOf(EndpointType(EndpointType_HTTPS))))
    fieldNameMap["download_session_id"] = "DownloadSessionId"
    fieldNameMap["file_name"] = "FileName"
    fieldNameMap["endpoint_type"] = "EndpointType"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func prepareOutputType() bindings.BindingType {
    return bindings.NewReferenceType(InfoBindingType)
}

func prepareRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404,"InvalidArgument": 400,"Unauthorized": 403})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["download_session_id"] = bindings.NewIdType([]string {"com.vmware.content.library.item.DownloadSession"}, "")
    fields["file_name"] = bindings.NewStringType()
    fieldNameMap["download_session_id"] = "DownloadSessionId"
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



func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["size"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["size"] = "Size"
    fields["bytes_transferred"] = bindings.NewIntegerType()
    fieldNameMap["bytes_transferred"] = "BytesTransferred"
    fields["status"] = bindings.NewEnumType("com.vmware.content.library.item.downloadsession.file.prepare_status", reflect.TypeOf(PrepareStatus(PrepareStatus_UNPREPARED)))
    fieldNameMap["status"] = "Status"
    fields["download_endpoint"] = bindings.NewOptionalType(bindings.NewReferenceType(item.TransferEndpointBindingType))
    fieldNameMap["download_endpoint"] = "DownloadEndpoint"
    fields["checksum_info"] = bindings.NewOptionalType(bindings.NewReferenceType(file.ChecksumInfoBindingType))
    fieldNameMap["checksum_info"] = "ChecksumInfo"
    fields["error_message"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
    fieldNameMap["error_message"] = "ErrorMessage"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.content.library.item.downloadsession.file.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}


