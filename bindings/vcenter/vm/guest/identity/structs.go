/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Identity.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package identity

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/vcenter/vm"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)




// The ``Info`` class contains information describing the guest operating system identification.
 type Info struct {
    Name vm.GuestOS
    Family vm.GuestOSFamily
    FullName std.LocalizableMessage
    HostName string
    IpAddress *string
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
    return bindings.NewReferenceType(InfoBindingType)
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
      "/vcenter/vm/{vm}/guest/identity",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"NotFound": 404,"ServiceUnavailable": 503})
}



func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewEnumType("com.vmware.vcenter.vm.guest_OS", reflect.TypeOf(vm.GuestOS(vm.GuestOS_DOS)))
    fieldNameMap["name"] = "Name"
    fields["family"] = bindings.NewEnumType("com.vmware.vcenter.vm.guest_OS_family", reflect.TypeOf(vm.GuestOSFamily(vm.GuestOSFamily_WINDOWS)))
    fieldNameMap["family"] = "Family"
    fields["full_name"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
    fieldNameMap["full_name"] = "FullName"
    fields["host_name"] = bindings.NewStringType()
    fieldNameMap["host_name"] = "HostName"
    fields["ip_address"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["ip_address"] = "IpAddress"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.guest.identity.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}


