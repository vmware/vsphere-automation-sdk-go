/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Networking.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package networking

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/appliance/networking/dns/servers"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/appliance/networking/interfaces"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/appliance/networking/interfaces/ipv4"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/appliance/networking/interfaces/ipv6"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)




// The ``DNSInfo`` class contains information about the DNS configuration of a virtual appliance.
 type DNSInfo struct {
    Mode DNSInfo_DNSMode
    Hostname string
    Servers []string
}




    
    // The ``DNSMode`` enumeration class describes the source of DNS servers.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type DNSInfo_DNSMode string

    const (
        // The DNS servers addresses are obtained from a DHCP server.
         DNSInfo_DNSMode_DHCP DNSInfo_DNSMode = "DHCP"
        // The DNS servers addresses are specified explicitly.
         DNSInfo_DNSMode_STATIC DNSInfo_DNSMode = "STATIC"
    )

    func (d DNSInfo_DNSMode) DNSInfo_DNSMode() bool {
        switch d {
            case DNSInfo_DNSMode_DHCP:
                return true
            case DNSInfo_DNSMode_STATIC:
                return true
            default:
                return false
        }
    }



// The ``Info`` class contains information about the network configuration of a virtual appliance.
 type Info struct {
    Dns DNSInfo
    Interfaces map[string]interfaces.InterfaceInfo
}






// The ``UpdateSpec`` class describes whether to enable or disable ipv6 on interfaces
 type UpdateSpec struct {
    Ipv6Enabled *bool
}






 type ChangeSpec struct {
    Hostname string
    SSOUser string
    SSOPassword string
    Dns *servers.DNSServerConfig
    Ipv4 *ipv4.Config
    Ipv6 *ipv6.Config
}









func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewReferenceType(InfoBindingType)
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


func updateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(UpdateSpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func updateOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func updateRestMetadata() protocol.OperationRestMetadata {
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


func resetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func resetOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func resetRestMetadata() protocol.OperationRestMetadata {
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


func changeInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(ChangeSpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func changeOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func changeRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
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
       map[string]int{"Unsupported": 400,"InvalidArgument": 400,"Unauthenticated": 401,"NotAllowedInCurrentState": 400})
}



func DNSInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["mode"] = bindings.NewEnumType("com.vmware.appliance.networking.DNS_info.DNS_mode", reflect.TypeOf(DNSInfo_DNSMode(DNSInfo_DNSMode_DHCP)))
    fieldNameMap["mode"] = "Mode"
    fields["hostname"] = bindings.NewStringType()
    fieldNameMap["hostname"] = "Hostname"
    fields["servers"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["servers"] = "Servers"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.networking.DNS_info",fields, reflect.TypeOf(DNSInfo{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["dns"] = bindings.NewReferenceType(DNSInfoBindingType)
    fieldNameMap["dns"] = "Dns"
    fields["interfaces"] = bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.appliance.networking.interfaces"}, ""), bindings.NewReferenceType(interfaces.InterfaceInfoBindingType),reflect.TypeOf(map[string]interfaces.InterfaceInfo{}))
    fieldNameMap["interfaces"] = "Interfaces"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.networking.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func UpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["ipv6_enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["ipv6_enabled"] = "Ipv6Enabled"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.networking.update_spec",fields, reflect.TypeOf(UpdateSpec{}), fieldNameMap, validators)
}

func ChangeSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["hostname"] = bindings.NewStringType()
    fieldNameMap["hostname"] = "Hostname"
    fields["SSO_user"] = bindings.NewStringType()
    fieldNameMap["SSO_user"] = "SSOUser"
    fields["SSO_password"] = bindings.NewSecretType()
    fieldNameMap["SSO_password"] = "SSOPassword"
    fields["dns"] = bindings.NewOptionalType(bindings.NewReferenceType(servers.DNSServerConfigBindingType))
    fieldNameMap["dns"] = "Dns"
    fields["ipv4"] = bindings.NewOptionalType(bindings.NewReferenceType(ipv4.ConfigBindingType))
    fieldNameMap["ipv4"] = "Ipv4"
    fields["ipv6"] = bindings.NewOptionalType(bindings.NewReferenceType(ipv6.ConfigBindingType))
    fieldNameMap["ipv6"] = "Ipv6"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.networking.change_spec",fields, reflect.TypeOf(ChangeSpec{}), fieldNameMap, validators)
}


