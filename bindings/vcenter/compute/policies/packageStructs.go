/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
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


// The ``ObjectCompliance`` enumeration class defines the compliance states a policy can be in on a particular object. **Warning:** This enumeration is available as technical preview. It may be changed in a future release.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.

type ObjectCompliance string

const (
    // The object has an unknown compliance state. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ObjectCompliance_UNKNOWN ObjectCompliance = "UNKNOWN"
    // The object is in a state for which the policy does not apply. **Warning:** This constant field is available as technical preview. It may be changed in a future release.
     ObjectCompliance_NOT_APPLICABLE ObjectCompliance = "NOT_APPLICABLE"
    // The policy is in compliance on the object. **Warning:** This constant field is available as technical preview. It may be changed in a future release.
     ObjectCompliance_COMPLIANT ObjectCompliance = "COMPLIANT"
    // The policy is not in compliance on the object. **Warning:** This constant field is available as technical preview. It may be changed in a future release.
     ObjectCompliance_NOT_COMPLIANT ObjectCompliance = "NOT_COMPLIANT"
)

func (o ObjectCompliance) ObjectCompliance() bool {
    switch o {
        case ObjectCompliance_UNKNOWN:
            return true
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




// The ``PolicyCompliance`` enumeration class defines the compliance states for a policy. This policy compliance status is calculated based on compliance status of all the objects associated with the policy and visible to the user. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.

type PolicyCompliance string

const (
    // The policy has unknown compliance state. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     PolicyCompliance_UNKNOWN PolicyCompliance = "UNKNOWN"
    // The policy is in not applicable state. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     PolicyCompliance_NOT_APPLICABLE PolicyCompliance = "NOT_APPLICABLE"
    // The policy is in compliance state. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     PolicyCompliance_COMPLIANT PolicyCompliance = "COMPLIANT"
    // The policy is not in compliance state. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     PolicyCompliance_NOT_COMPLIANT PolicyCompliance = "NOT_COMPLIANT"
)

func (p PolicyCompliance) PolicyCompliance() bool {
    switch p {
        case PolicyCompliance_UNKNOWN:
            return true
        case PolicyCompliance_NOT_APPLICABLE:
            return true
        case PolicyCompliance_COMPLIANT:
            return true
        case PolicyCompliance_NOT_COMPLIANT:
            return true
        default:
            return false
    }
}





// The ``CreateSpec`` class contains common information used to create a new policy. **Warning:** This class is available as technical preview. It may be changed in a future release.
type CreateSpec struct {
    Capability string
    Name string
    Description string
}






// The ``Info`` class contains common information about a compute policy. **Warning:** This class is available as technical preview. It may be changed in a future release.
type Info struct {
    Name string
    Description string
    Capability string
}






// The ``Status`` class describes the current status of a compute policy. **Warning:** This class is available as technical preview. It may be changed in a future release.
type Status struct {
    Status ObjectCompliance
}










func CreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["capability"] = bindings.NewIdType([]string {"com.vmware.vcenter.compute.policies.Capability"}, "")
    fieldNameMap["capability"] = "Capability"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["description"] = bindings.NewStringType()
    fieldNameMap["description"] = "Description"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.compute.policies.create_spec",fields, reflect.TypeOf(CreateSpec{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["description"] = bindings.NewStringType()
    fieldNameMap["description"] = "Description"
    fields["capability"] = bindings.NewIdType([]string {"com.vmware.vcenter.compute.policies.Capability"}, "")
    fieldNameMap["capability"] = "Capability"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.compute.policies.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func StatusBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["status"] = bindings.NewEnumType("com.vmware.vcenter.compute.policies.object_compliance", reflect.TypeOf(ObjectCompliance(ObjectCompliance_UNKNOWN)))
    fieldNameMap["status"] = "Status"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.compute.policies.status",fields, reflect.TypeOf(Status{}), fieldNameMap, validators)
}


