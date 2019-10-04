/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for package: com.vmware.vcenter.vm.guest.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package guest

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
)



// The ``Credentials`` class defines the guest credentials used for guest operation authentication. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type Credentials struct {
    InteractiveSession bool
    Type_ Credentials_Type
    UserName *string
    Password *string
    SamlToken *string
}




    
    // Types of guest credentials. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type Credentials_Type string

    const (
        // USERNAME_PASSWORD credentials contains the information necessary to authenticate within a guest using a username and password. This method of authentication is stateless. 
        //
        //  To use USERNAME_PASSWORD, populate userName and password with the appropriate login information. 
        //
        //  Once populated, you can use USERNAME_PASSWORD in any guest operations function call.. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         Credentials_Type_USERNAME_PASSWORD Credentials_Type = "USERNAME_PASSWORD"
        // SAML_BEARER_TOKEN contains the information necessary to authenticate within a guest using a SAML token. SAML Bearer token credentials relies on a guest alias that associates a guest account with the subject and certificate encoded in a SAML Bearer token obtained from the VMware SSO Server. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         Credentials_Type_SAML_BEARER_TOKEN Credentials_Type = "SAML_BEARER_TOKEN"
    )

    func (t Credentials_Type) Credentials_Type() bool {
        switch t {
            case Credentials_Type_USERNAME_PASSWORD:
                return true
            case Credentials_Type_SAML_BEARER_TOKEN:
                return true
            default:
                return false
        }
    }



// The {\\\\@name DnsAssignedValues) class describes values assigned by a Domain Name Server (DNS).
type DnsAssignedValues struct {
    HostName string
    DomainName string
}






// The ``DnsConfigInfo`` class describes the configuration of RFC 1034 DNS settings.
type DnsConfigInfo struct {
    IpAddresses []string
    SearchDomains []string
}






// The ``DhcpConfigInfo`` class specifies when Dynamic Host Configuration Protocol is enabled.
type DhcpConfigInfo struct {
    Ipv4Enabled bool
    Ipv6Enabled bool
}










func CredentialsBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["interactive_session"] = bindings.NewBooleanType()
    fieldNameMap["interactive_session"] = "InteractiveSession"
    fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm.guest.credentials.type", reflect.TypeOf(Credentials_Type(Credentials_Type_USERNAME_PASSWORD)))
    fieldNameMap["type"] = "Type_"
    fields["user_name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["user_name"] = "UserName"
    fields["password"] = bindings.NewOptionalType(bindings.NewSecretType())
    fieldNameMap["password"] = "Password"
    fields["saml_token"] = bindings.NewOptionalType(bindings.NewSecretType())
    fieldNameMap["saml_token"] = "SamlToken"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("type",
        map[string][]bindings.FieldData {
            "USERNAME_PASSWORD": []bindings.FieldData {
                 bindings.NewFieldData("user_name", false),
                 bindings.NewFieldData("password", true),
            },
            "SAML_BEARER_TOKEN": []bindings.FieldData {
                 bindings.NewFieldData("user_name", false),
                 bindings.NewFieldData("saml_token", true),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.vm.guest.credentials",fields, reflect.TypeOf(Credentials{}), fieldNameMap, validators)
}

func DnsAssignedValuesBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["host_name"] = bindings.NewStringType()
    fieldNameMap["host_name"] = "HostName"
    fields["domain_name"] = bindings.NewStringType()
    fieldNameMap["domain_name"] = "DomainName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.guest.dns_assigned_values",fields, reflect.TypeOf(DnsAssignedValues{}), fieldNameMap, validators)
}

func DnsConfigInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["ip_addresses"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["ip_addresses"] = "IpAddresses"
    fields["search_domains"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["search_domains"] = "SearchDomains"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.guest.dns_config_info",fields, reflect.TypeOf(DnsConfigInfo{}), fieldNameMap, validators)
}

func DhcpConfigInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["ipv4_enabled"] = bindings.NewBooleanType()
    fieldNameMap["ipv4_enabled"] = "Ipv4Enabled"
    fields["ipv6_enabled"] = bindings.NewBooleanType()
    fieldNameMap["ipv6_enabled"] = "Ipv6Enabled"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.guest.dhcp_config_info",fields, reflect.TypeOf(DhcpConfigInfo{}), fieldNameMap, validators)
}


