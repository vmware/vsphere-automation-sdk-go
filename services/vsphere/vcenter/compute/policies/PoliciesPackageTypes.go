/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for package: com.vmware.vcenter.compute.policies.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package policies

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
)


// The ``ObjectCompliance`` enumeration class defines the compliance states a policy can be in on a particular object in VMware Cloud on AWS. Usage beyond VMware Cloud on AWS is not supported. **Warning:** This enumeration is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type ObjectCompliance string

const (
    // The object is in a state for which the policy does not apply. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	ObjectCompliance_NOT_APPLICABLE ObjectCompliance = "NOT_APPLICABLE"
    // The policy is in compliance on the object. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	ObjectCompliance_COMPLIANT ObjectCompliance = "COMPLIANT"
    // The policy is not in compliance on the object. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	ObjectCompliance_NOT_COMPLIANT ObjectCompliance = "NOT_COMPLIANT"
)

func (o ObjectCompliance) ObjectCompliance() bool {
	switch o {
	case ObjectCompliance_NOT_APPLICABLE:
		return true
	case ObjectCompliance_COMPLIANT:
		return true
	case ObjectCompliance_NOT_COMPLIANT:
		return true
	default:
		return false
	}
}


// The ``CreateSpec`` class contains common information used to create a new policy in VMware Cloud on AWS. Usage beyond VMware Cloud on AWS is not supported. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type CreateSpec struct {
    // Identifier of the capability this policy is based on. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Capability string
    // Name of the policy. The name needs to be unique within this vCenter server. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Name string
    // Description of the policy. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Description string
}

// The ``Info`` class contains common information about a compute policy in VMware Cloud on AWS. Usage beyond VMware Cloud on AWS is not supported. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type Info struct {
    // Name of the policy. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Name string
    // Description of the policy. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Description string
    // Identifier of the capability this policy is based on. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Capability string
}

// The ``Status`` class describes the current status of a compute policy in VMware Cloud on AWS. Usage beyond VMware Cloud on AWS is not supported. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type Status struct {
    // The compliance status of the policy on a specified object. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Status ObjectCompliance
}




func CreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["capability"] = bindings.NewIdType([]string{"com.vmware.vcenter.compute.policies.Capability"}, "")
	fieldNameMap["capability"] = "Capability"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["description"] = bindings.NewStringType()
	fieldNameMap["description"] = "Description"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.compute.policies.create_spec", fields, reflect.TypeOf(CreateSpec{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["description"] = bindings.NewStringType()
	fieldNameMap["description"] = "Description"
	fields["capability"] = bindings.NewIdType([]string{"com.vmware.vcenter.compute.policies.Capability"}, "")
	fieldNameMap["capability"] = "Capability"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.compute.policies.info", fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func StatusBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["status"] = bindings.NewEnumType("com.vmware.vcenter.compute.policies.object_compliance", reflect.TypeOf(ObjectCompliance(ObjectCompliance_NOT_APPLICABLE)))
	fieldNameMap["status"] = "Status"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.compute.policies.status", fields, reflect.TypeOf(Status{}), fieldNameMap, validators)
}


