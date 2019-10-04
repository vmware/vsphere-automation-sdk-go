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
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
)




// The ``Info`` class contains information about the identity of a virtual machine.
 type Info struct {
    Name string
    BiosUuid string
    InstanceUuid string
}










func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["bios_uuid"] = bindings.NewStringType()
    fieldNameMap["bios_uuid"] = "BiosUuid"
    fields["instance_uuid"] = bindings.NewStringType()
    fieldNameMap["instance_uuid"] = "InstanceUuid"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.identity.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}


