/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Inbound.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package inbound

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// ``Policy`` enumeration class Defines firewall rule policies.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Policy string

const (
    // Drop packet with correpsonding address.
     Policy_IGNORE Policy = "IGNORE"
    // Allow packet with corresponding address.
     Policy_ACCEPT Policy = "ACCEPT"
    // Drop packet with corresponding address sending destination is not reachable.
     Policy_REJECT Policy = "REJECT"
    // Apply default or port-specific rules to packet with corresponding address.
     Policy_RETURN Policy = "RETURN"
)

func (p Policy) Policy() bool {
    switch p {
        case Policy_IGNORE:
            return true
        case Policy_ACCEPT:
            return true
        case Policy_REJECT:
            return true
        case Policy_RETURN:
            return true
        default:
            return false
    }
}





// ``Rule`` class Structure that defines a single address-based firewall rule.
 type Rule struct {
    Address string
    Prefix int64
    Policy Policy
    InterfaceName *string
}









func setInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["rules"] = bindings.NewListType(bindings.NewReferenceType(RuleBindingType), reflect.TypeOf([]Rule{}))
    fieldNameMap["rules"] = "Rules"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func setOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func setRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(RuleBindingType), reflect.TypeOf([]Rule{}))
}

func getRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500})
}



func RuleBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["address"] = bindings.NewStringType()
    fieldNameMap["address"] = "Address"
    fields["prefix"] = bindings.NewIntegerType()
    fieldNameMap["prefix"] = "Prefix"
    fields["policy"] = bindings.NewEnumType("com.vmware.appliance.networking.firewall.inbound.policy", reflect.TypeOf(Policy(Policy_IGNORE)))
    fieldNameMap["policy"] = "Policy"
    fields["interface_name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["interface_name"] = "InterfaceName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.networking.firewall.inbound.rule",fields, reflect.TypeOf(Rule{}), fieldNameMap, validators)
}


