/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for package: com.vmware.vcenter.compute.policies.capabilities.vm.evacuation.vmotion.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package vmotion

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
)



// The ``CreateSpec`` class contains information used to create a new vMotion policy that applies when a host is evacuated, see Policies#create. All virtual machines that share the tag indicated by CreateSpec#vmTag will be vMotioned from a host whenever the host is evacuated by vCenter. If vCenter cannot migrate a virtual machine, then it will remain running on its current host. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type CreateSpec struct {
    VmTag string
    Capability string
    Name string
    Description string
}






// The ``Info`` class contains information about a vMotion policy that applies when a host is evacuated, see Policies#get. All virtual machines that share the tag indicated by Info#vmTag will be vMotioned from a host whenever the host is evacuated by vCenter. If vCenter cannot migrate a virtual machine, then it will remain running on its current host. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type Info struct {
    VmTag string
    Name string
    Description string
    Capability string
}










func CreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm_tag"] = bindings.NewIdType([]string {"com.vmware.cis.tagging.Tag:VirtualMachine"}, "")
    fieldNameMap["vm_tag"] = "VmTag"
    fields["capability"] = bindings.NewIdType([]string {"com.vmware.vcenter.compute.policies.Capability"}, "")
    fieldNameMap["capability"] = "Capability"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["description"] = bindings.NewStringType()
    fieldNameMap["description"] = "Description"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.compute.policies.capabilities.vm.evacuation.vmotion.create_spec",fields, reflect.TypeOf(CreateSpec{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm_tag"] = bindings.NewIdType([]string {"com.vmware.cis.tagging.Tag:VirtualMachine"}, "")
    fieldNameMap["vm_tag"] = "VmTag"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["description"] = bindings.NewStringType()
    fieldNameMap["description"] = "Description"
    fields["capability"] = bindings.NewIdType([]string {"com.vmware.vcenter.compute.policies.Capability"}, "")
    fieldNameMap["capability"] = "Capability"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.compute.policies.capabilities.vm.evacuation.vmotion.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}


