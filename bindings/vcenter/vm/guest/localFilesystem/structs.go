/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: LocalFilesystem.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package localFilesystem

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)




// Describes the virtual disk backing a local guest disk.
 type VirtualDiskMapping struct {
    Disk string
}






// The ``Info`` class contains information about a local file system configured in the guest operating system.
 type Info struct {
    Capacity int64
    FreeSpace int64
    Filesystem *string
    Mappings []VirtualDiskMapping
}









func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fieldNameMap["vm"] = "Vm"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewMapType(bindings.NewStringType(), bindings.NewReferenceType(InfoBindingType),reflect.TypeOf(map[string]Info{}))
}

func getRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    paramsTypeMap["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    pathParams["vm"] = "vm"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "GET",
      "/vcenter/vm/{vm}/guest/local-filesystem",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"NotFound": 404,"ServiceUnavailable": 503})
}



func VirtualDiskMappingBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["disk"] = bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Disk"}, "")
    fieldNameMap["disk"] = "Disk"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.guest.local_filesystem.virtual_disk_mapping",fields, reflect.TypeOf(VirtualDiskMapping{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["capacity"] = bindings.NewIntegerType()
    fieldNameMap["capacity"] = "Capacity"
    fields["free_space"] = bindings.NewIntegerType()
    fieldNameMap["free_space"] = "FreeSpace"
    fields["filesystem"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["filesystem"] = "Filesystem"
    fields["mappings"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(VirtualDiskMappingBindingType), reflect.TypeOf([]VirtualDiskMapping{})))
    fieldNameMap["mappings"] = "Mappings"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.guest.local_filesystem.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}


