/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for package: com.vmware.vcenter.trusted_infrastructure.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package trusted_infrastructure

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
)


// The ``StsPrincipalType`` enum can be either users or groups.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.

type StsPrincipalType string

const (
    // The principal is a user.
     StsPrincipalType_STS_USER StsPrincipalType = "STS_USER"
    // The principal is a group.
     StsPrincipalType_STS_GROUP StsPrincipalType = "STS_GROUP"
)

func (s StsPrincipalType) StsPrincipalType() bool {
    switch s {
        case StsPrincipalType_STS_USER:
            return true
        case StsPrincipalType_STS_GROUP:
            return true
        default:
            return false
    }
}





// The ``NetworkAddress`` class contains an IP address or DNS resolvable name and a port on which a connection can be established.
type NetworkAddress struct {
    Hostname string
    Port *int64
}






// The ``StsPrincipalId`` class contains an IDM principal ID.
type StsPrincipalId struct {
    Name string
    Domain string
}






// The ``StsPrincipal`` class contains a IDM principal.
type StsPrincipal struct {
    Id StsPrincipalId
    Type_ StsPrincipalType
}






// The ``X509CertChain`` class contains x509 certificate chain.
type X509CertChain struct {
    CertChain []string
}










func NetworkAddressBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["hostname"] = bindings.NewStringType()
    fieldNameMap["hostname"] = "Hostname"
    fields["port"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["port"] = "Port"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.network_address",fields, reflect.TypeOf(NetworkAddress{}), fieldNameMap, validators)
}

func StsPrincipalIdBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["domain"] = bindings.NewStringType()
    fieldNameMap["domain"] = "Domain"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.sts_principal_id",fields, reflect.TypeOf(StsPrincipalId{}), fieldNameMap, validators)
}

func StsPrincipalBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewReferenceType(StsPrincipalIdBindingType)
    fieldNameMap["id"] = "Id"
    fields["type"] = bindings.NewEnumType("com.vmware.vcenter.trusted_infrastructure.sts_principal_type", reflect.TypeOf(StsPrincipalType(StsPrincipalType_STS_USER)))
    fieldNameMap["type"] = "Type_"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.sts_principal",fields, reflect.TypeOf(StsPrincipal{}), fieldNameMap, validators)
}

func X509CertChainBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cert_chain"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["cert_chain"] = "CertChain"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.x509_cert_chain",fields, reflect.TypeOf(X509CertChain{}), fieldNameMap, validators)
}


