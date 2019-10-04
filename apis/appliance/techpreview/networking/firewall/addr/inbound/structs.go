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



// ``FirewallRulePolicy`` enumeration class Defines firewall rule policies
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type FirewallRulePolicy string

const (
    // Deny packet with correpsonding address.
     FirewallRulePolicy_deny FirewallRulePolicy = "deny"
    // Allow packet with corresponding address.
     FirewallRulePolicy_allow FirewallRulePolicy = "allow"
)

func (f FirewallRulePolicy) FirewallRulePolicy() bool {
    switch f {
        case FirewallRulePolicy_deny:
            return true
        case FirewallRulePolicy_allow:
            return true
        default:
            return false
    }
}





// ``FirewallAddressRule`` class Structure that defines a single address-based firewall rule.
 type FirewallAddressRule struct {
    Address string
    Prefix int64
    Policy FirewallRulePolicy
    InterfaceName string
}






// ``DeleteFirewallRule`` class Structure that defines networking.firewall.addr.inbound.delete api input argument
 type DeleteFirewallRule struct {
    Position int64
    All bool
}









func addInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["pos"] = bindings.NewIntegerType()
    fields["rule"] = bindings.NewReferenceType(FirewallAddressRuleBindingType)
    fieldNameMap["pos"] = "Pos"
    fieldNameMap["rule"] = "Rule"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func addOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func addRestMetadata() protocol.OperationRestMetadata {
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


func setInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["rules"] = bindings.NewListType(bindings.NewReferenceType(FirewallAddressRuleBindingType), reflect.TypeOf([]FirewallAddressRule{}))
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


func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(FirewallAddressRuleBindingType), reflect.TypeOf([]FirewallAddressRule{}))
}

func listRestMetadata() protocol.OperationRestMetadata {
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


func deleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["config"] = bindings.NewReferenceType(DeleteFirewallRuleBindingType)
    fieldNameMap["config"] = "Config"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func deleteOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func deleteRestMetadata() protocol.OperationRestMetadata {
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



func FirewallAddressRuleBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["address"] = bindings.NewStringType()
    fieldNameMap["address"] = "Address"
    fields["prefix"] = bindings.NewIntegerType()
    fieldNameMap["prefix"] = "Prefix"
    fields["policy"] = bindings.NewEnumType("com.vmware.appliance.techpreview.networking.firewall.addr.inbound.firewall_rule_policy", reflect.TypeOf(FirewallRulePolicy(FirewallRulePolicy_deny)))
    fieldNameMap["policy"] = "Policy"
    fields["interface_name"] = bindings.NewStringType()
    fieldNameMap["interface_name"] = "InterfaceName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.networking.firewall.addr.inbound.firewall_address_rule",fields, reflect.TypeOf(FirewallAddressRule{}), fieldNameMap, validators)
}

func DeleteFirewallRuleBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["position"] = bindings.NewIntegerType()
    fieldNameMap["position"] = "Position"
    fields["all"] = bindings.NewBooleanType()
    fieldNameMap["all"] = "All"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.networking.firewall.addr.inbound.delete_firewall_rule",fields, reflect.TypeOf(DeleteFirewallRule{}), fieldNameMap, validators)
}


