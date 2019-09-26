/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Storage.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package storage

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/content/library"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/content/library/item/file"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
    "net/url"
)




// The ``Info`` class is the expanded form of file.Info that includes details about the storage backing for a file in a library item.
 type Info struct {
    StorageBacking library.StorageBacking
    StoragePolicyId *string
    StorageUris []url.URL
    ChecksumInfo *file.ChecksumInfo
    Name string
    Size int64
    Cached bool
    Version string
}









func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["library_item_id"] = bindings.NewIdType([]string {"com.vmware.content.library.Item"}, "")
    fields["file_name"] = bindings.NewStringType()
    fieldNameMap["library_item_id"] = "LibraryItemId"
    fieldNameMap["file_name"] = "FileName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(InfoBindingType), reflect.TypeOf([]Info{}))
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



func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["storage_backing"] = bindings.NewReferenceType(library.StorageBackingBindingType)
    fieldNameMap["storage_backing"] = "StorageBacking"
    fields["storage_policy_id"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.spbm.StorageProfile"}, ""))
    fieldNameMap["storage_policy_id"] = "StoragePolicyId"
    fields["storage_uris"] = bindings.NewListType(bindings.NewUriType(), reflect.TypeOf([]url.URL{}))
    fieldNameMap["storage_uris"] = "StorageUris"
    fields["checksum_info"] = bindings.NewOptionalType(bindings.NewReferenceType(file.ChecksumInfoBindingType))
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
    return bindings.NewStructType("com.vmware.content.library.item.storage.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}


