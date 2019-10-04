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
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``ChecksumAlgorithm`` enumeration class defines the valid checksum algorithms.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type ChecksumAlgorithm string

const (
    // Checksum algorithm: SHA-1
     ChecksumAlgorithm_SHA1 ChecksumAlgorithm = "SHA1"
    // Checksum algorithm: MD5
     ChecksumAlgorithm_MD5 ChecksumAlgorithm = "MD5"
    // Checksum algorithm: SHA-256
     ChecksumAlgorithm_SHA256 ChecksumAlgorithm = "SHA256"
    // Checksum algorithm: SHA-512
     ChecksumAlgorithm_SHA512 ChecksumAlgorithm = "SHA512"
)

func (c ChecksumAlgorithm) ChecksumAlgorithm() bool {
    switch c {
        case ChecksumAlgorithm_SHA1:
            return true
        case ChecksumAlgorithm_MD5:
            return true
        case ChecksumAlgorithm_SHA256:
            return true
        case ChecksumAlgorithm_SHA512:
            return true
        default:
            return false
    }
}





// Provides checksums for a Info object.
 type ChecksumInfo struct {
    Algorithm *ChecksumAlgorithm
    Checksum string
}






// The ``Info`` class provides information about a file in Content Library Service storage. 
//
//  A file is an actual stored object for a library item. An item will have zero files initially, but one or more can be uploaded to the item.
 type Info struct {
    ChecksumInfo *ChecksumInfo
    Name string
    Size int64
    Cached bool
    Version string
}









func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["library_item_id"] = bindings.NewIdType([]string {"com.vmware.content.library.Item"}, "")
    fields["name"] = bindings.NewStringType()
    fieldNameMap["library_item_id"] = "LibraryItemId"
    fieldNameMap["name"] = "Name"
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
       map[string]int{"NotFound": 404})
}


func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["library_item_id"] = bindings.NewIdType([]string {"com.vmware.content.library.Item"}, "")
    fieldNameMap["library_item_id"] = "LibraryItemId"
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



func ChecksumInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["algorithm"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.content.library.item.file.checksum_algorithm", reflect.TypeOf(ChecksumAlgorithm(ChecksumAlgorithm_SHA1))))
    fieldNameMap["algorithm"] = "Algorithm"
    fields["checksum"] = bindings.NewStringType()
    fieldNameMap["checksum"] = "Checksum"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.content.library.item.file.checksum_info",fields, reflect.TypeOf(ChecksumInfo{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["checksum_info"] = bindings.NewOptionalType(bindings.NewReferenceType(ChecksumInfoBindingType))
    fieldNameMap["checksum_info"] = "ChecksumInfo"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["size"] = bindings.NewIntegerType()
    fieldNameMap["size"] = "Size"
    fields["cached"] = bindings.NewBooleanType()
    fieldNameMap["cached"] = "Cached"
    fields["version"] = bindings.NewStringType()
    fieldNameMap["version"] = "Version"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.content.library.item.file.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}


