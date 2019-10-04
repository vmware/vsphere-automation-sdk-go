/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for package: com.vmware.vcenter.compute.policies.capabilities.cluster_scale_in_ignore_vm_capabilities.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package cluster_scale_in_ignore_vm_capabilities

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
)



// The ``CreateSpec`` class contains information used to create a new policy to ignore virtual machine capabilities when scaling-in a cluster, see Policies#create. When considering scaling-in a cluster, policies that have been created with one of the listed CreateSpec#vmCapabilities are ignored for virtual machines that have the tag indicated by CreateSpec#vmTag. **Warning:** This class is available as technical preview. It may be changed in a future release.
type CreateSpec struct {
    VmTag string
    VmCapabilities map[string]bool
    Capability string
    Name string
    Description string
}






// The ``Info`` class contains information about a policy to ignore virtual machine capabilities when scaling-in a cluster, see Policies#get. When considering scaling-in a cluster, policies that have been created with one of the listed Info#vmCapabilities are ignored for virtual machines that have the tag indicated by Info#vmTag. **Warning:** This class is available as technical preview. It may be changed in a future release.
type Info struct {
    VmTag string
    VmCapabilities map[string]bool
    Name string
    Description string
    Capability string
}










func CreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm_tag"] = bindings.NewIdType([]string {"com.vmware.cis.tagging.Tag:VirtualMachine"}, "")
    fieldNameMap["vm_tag"] = "VmTag"
    fields["vm_capabilities"] = bindings.NewSetType(bindings.NewIdType([]string {"com.vmware.vcenter.compute.policies.Capability:VirtualMachine"}, ""), reflect.TypeOf(map[string]bool{}))
    fieldNameMap["vm_capabilities"] = "VmCapabilities"
    fields["capability"] = bindings.NewIdType([]string {"com.vmware.vcenter.compute.policies.Capability"}, "")
    fieldNameMap["capability"] = "Capability"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["description"] = bindings.NewStringType()
    fieldNameMap["description"] = "Description"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.compute.policies.capabilities.cluster_scale_in_ignore_vm_capabilities.create_spec",fields, reflect.TypeOf(CreateSpec{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm_tag"] = bindings.NewIdType([]string {"com.vmware.cis.tagging.Tag:VirtualMachine"}, "")
    fieldNameMap["vm_tag"] = "VmTag"
    fields["vm_capabilities"] = bindings.NewSetType(bindings.NewIdType([]string {"com.vmware.vcenter.compute.policies.Capability:VirtualMachine"}, ""), reflect.TypeOf(map[string]bool{}))
    fieldNameMap["vm_capabilities"] = "VmCapabilities"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["description"] = bindings.NewStringType()
    fieldNameMap["description"] = "Description"
    fields["capability"] = bindings.NewIdType([]string {"com.vmware.vcenter.compute.policies.Capability"}, "")
    fieldNameMap["capability"] = "Capability"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.compute.policies.capabilities.cluster_scale_in_ignore_vm_capabilities.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}


