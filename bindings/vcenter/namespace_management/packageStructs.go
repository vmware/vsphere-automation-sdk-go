/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for package: com.vmware.vcenter.namespace_management.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package namespace_management

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
)


// The ``SizingHint`` enumeration class determines the configuration of Kubernetes API server and the worker nodes. It also determines the default values associated with the maximum number of pods and services. Use ClusterSizeInfo#get to get information associated with a ``SizingHint``. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.

type SizingHint string

const (
    // Cluster size of 'tiny'. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     SizingHint_TINY SizingHint = "TINY"
    // Cluster size of 'small'. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     SizingHint_SMALL SizingHint = "SMALL"
    // Cluster size of 'medium'. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     SizingHint_MEDIUM SizingHint = "MEDIUM"
    // Cluster size of 'large'. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     SizingHint_LARGE SizingHint = "LARGE"
)

func (s SizingHint) SizingHint() bool {
    switch s {
        case SizingHint_TINY:
            return true
        case SizingHint_SMALL:
            return true
        case SizingHint_MEDIUM:
            return true
        case SizingHint_LARGE:
            return true
        default:
            return false
    }
}





// An ``EndpointAddress`` contains the data that describes an Endpoint's address in Kubernetes. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type EndpointAddress struct {
    Service string
    Ip string
    Port int64
    Protocol string
}






// An ``IngressRuleValue`` contains the data that describes an Ingress rule in Kubernetes. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type IngressRuleValue struct {
    Host string
    Path string
    Name string
    Port int64
}






// The ``Ipv4Cidr`` class contains the specification for representing CIDR notation of IP range. For example, this can be used to represent 256 IP addresses using 10.10.10.0/24. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type Ipv4Cidr struct {
    Address string
    Prefix int64
}










func EndpointAddressBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["service"] = bindings.NewStringType()
    fieldNameMap["service"] = "Service"
    fields["ip"] = bindings.NewStringType()
    fieldNameMap["ip"] = "Ip"
    fields["port"] = bindings.NewIntegerType()
    fieldNameMap["port"] = "Port"
    fields["protocol"] = bindings.NewStringType()
    fieldNameMap["protocol"] = "Protocol"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.endpoint_address",fields, reflect.TypeOf(EndpointAddress{}), fieldNameMap, validators)
}

func IngressRuleValueBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["host"] = bindings.NewStringType()
    fieldNameMap["host"] = "Host"
    fields["path"] = bindings.NewStringType()
    fieldNameMap["path"] = "Path"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["port"] = bindings.NewIntegerType()
    fieldNameMap["port"] = "Port"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.ingress_rule_value",fields, reflect.TypeOf(IngressRuleValue{}), fieldNameMap, validators)
}

func Ipv4CidrBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["address"] = bindings.NewStringType()
    fieldNameMap["address"] = "Address"
    fields["prefix"] = bindings.NewIntegerType()
    fieldNameMap["prefix"] = "Prefix"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.ipv4_cidr",fields, reflect.TypeOf(Ipv4Cidr{}), fieldNameMap, validators)
}


