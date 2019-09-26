/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for package: com.vmware.vcenter.compute.policies.capabilities.vm_host_affinity.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package vm_host_affinity

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
)



// The ``CreateSpec`` class contains information used to create a new virtual machine to host affinity policy, see Policies#create. Virtual machines that have the tag indicated by CreateSpec#vmTag will be affine to hosts that have the tag indicated by CreateSpec#hostTag. **Warning:** This class is available as technical preview. It may be changed in a future release.
type CreateSpec struct {
    VmTag string
    HostTag string
    Capability string
    Name string
    Description string
}






// The ``Info`` class contains information about a virtual machine to host affinity policy, see Policies#get. Virtual machines that have the tag indicated by Info#vmTag will be affine to hosts that have the tag indicated by Info#hostTag. **Warning:** This class is available as technical preview. It may be changed in a future release.
type Info struct {
    VmTag string
    HostTag string
    Name string
    Description string
    Capability string
}










func CreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm_tag"] = bindings.NewIdType([]string {"com.vmware.cis.tagging.Tag:VirtualMachine"}, "")
    fieldNameMap["vm_tag"] = "VmTag"
    fields["host_tag"] = bindings.NewIdType([]string {"com.vmware.cis.tagging.Tag:HostSystem"}, "")
    fieldNameMap["host_tag"] = "HostTag"
    fields["capability"] = bindings.NewIdType([]string {"com.vmware.vcenter.compute.policies.Capability"}, "")
    fieldNameMap["capability"] = "Capability"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["description"] = bindings.NewStringType()
    fieldNameMap["description"] = "Description"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.compute.policies.capabilities.vm_host_affinity.create_spec",fields, reflect.TypeOf(CreateSpec{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm_tag"] = bindings.NewIdType([]string {"com.vmware.cis.tagging.Tag:VirtualMachine"}, "")
    fieldNameMap["vm_tag"] = "VmTag"
    fields["host_tag"] = bindings.NewIdType([]string {"com.vmware.cis.tagging.Tag:HostSystem"}, "")
    fieldNameMap["host_tag"] = "HostTag"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["description"] = bindings.NewStringType()
    fieldNameMap["description"] = "Description"
    fields["capability"] = bindings.NewIdType([]string {"com.vmware.vcenter.compute.policies.Capability"}, "")
    fieldNameMap["capability"] = "Capability"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.compute.policies.capabilities.vm_host_affinity.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}


