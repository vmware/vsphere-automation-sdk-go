/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
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


// The ``CreateSpec`` class contains information used to create a new virtual machine to host affinity policy, see Policies#create. Virtual machines that have the tag indicated by CreateSpec#vmTag will be affine to hosts that have the tag indicated by CreateSpec#hostTag in VMware Cloud on AWS. Usage beyond VMware Cloud on AWS is not supported. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type CreateSpec struct {
    // Identifier of a tag that can be associated with a virtual machine. Virtual machines with this tag will be affine to the hosts indicated by CreateSpec#hostTag. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	VmTag string
    // Identifier of a tag that can be associated with a host. Virtual machines indicated by CreateSpec#vmTag will be affine to hosts with this tag. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	HostTag string
    // Identifier of the capability this policy is based on. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Capability string
    // Name of the policy. The name needs to be unique within this vCenter server. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Name string
    // Description of the policy. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Description string
}

// The ``Info`` class contains information about a virtual machine to host affinity policy, see Policies#get. Virtual machines that have the tag indicated by Info#vmTag will be affine to hosts that have the tag indicated by Info#hostTag in VMware Cloud on AWS. Usage beyond VMware Cloud on AWS is not supported. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type Info struct {
    // Identifier of a tag that can be associated with a virtual machine. Virtual machines with this tag will be affine to the hosts indicated by Info#hostTag. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	VmTag string
    // Identifier of a tag that can be associated with a host. Virtual machines indicated by Info#vmTag will be affine to hosts with this tag. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	HostTag string
    // Name of the policy. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Name string
    // Description of the policy. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Description string
    // Identifier of the capability this policy is based on. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Capability string
}




func CreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm_tag"] = bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag:VirtualMachine"}, "")
	fieldNameMap["vm_tag"] = "VmTag"
	fields["host_tag"] = bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag:HostSystem"}, "")
	fieldNameMap["host_tag"] = "HostTag"
	fields["capability"] = bindings.NewIdType([]string{"com.vmware.vcenter.compute.policies.Capability"}, "")
	fieldNameMap["capability"] = "Capability"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["description"] = bindings.NewStringType()
	fieldNameMap["description"] = "Description"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.compute.policies.capabilities.vm_host_affinity.create_spec", fields, reflect.TypeOf(CreateSpec{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm_tag"] = bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag:VirtualMachine"}, "")
	fieldNameMap["vm_tag"] = "VmTag"
	fields["host_tag"] = bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag:HostSystem"}, "")
	fieldNameMap["host_tag"] = "HostTag"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["description"] = bindings.NewStringType()
	fieldNameMap["description"] = "Description"
	fields["capability"] = bindings.NewIdType([]string{"com.vmware.vcenter.compute.policies.Capability"}, "")
	fieldNameMap["capability"] = "Capability"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.compute.policies.capabilities.vm_host_affinity.info", fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}


