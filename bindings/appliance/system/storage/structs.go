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
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)




// The ``StorageMapping`` class describes the mapping between VCSA partitions and the Hard disk numbers visible in the vSphere Web Client.
 type StorageMapping struct {
    Disk string
    Partition string
    Description *std.LocalizableMessage
}






// The ``StorageChange`` class describes the changes in capasity of a storage partition.
 type StorageChange struct {
    OldSize int64
    NewSize int64
}









func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(StorageMappingBindingType), reflect.TypeOf([]StorageMapping{}))
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
       map[string]int{"Error": 500})
}


func resizeInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func resizeOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func resizeRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500})
}


func resizeExInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func resizeExOutputType() bindings.BindingType {
    return bindings.NewMapType(bindings.NewStringType(), bindings.NewReferenceType(StorageChangeBindingType),reflect.TypeOf(map[string]StorageChange{}))
}

func resizeExRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500})
}



func StorageMappingBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["disk"] = bindings.NewIdType([]string {"com.vmware.appliance.system.storage"}, "")
    fieldNameMap["disk"] = "Disk"
    fields["partition"] = bindings.NewStringType()
    fieldNameMap["partition"] = "Partition"
    fields["description"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
    fieldNameMap["description"] = "Description"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.system.storage.storage_mapping",fields, reflect.TypeOf(StorageMapping{}), fieldNameMap, validators)
}

func StorageChangeBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["old_size"] = bindings.NewIntegerType()
    fieldNameMap["old_size"] = "OldSize"
    fields["new_size"] = bindings.NewIntegerType()
    fieldNameMap["new_size"] = "NewSize"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.system.storage.storage_change",fields, reflect.TypeOf(StorageChange{}), fieldNameMap, validators)
}


