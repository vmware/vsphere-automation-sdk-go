/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Providers.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package providers

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vcenter/certificate_management"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
    "net/url"
)



// The ``ConfigType`` class contains the possible types of vCenter Server identity providers. **Warning:** This enumeration is available as technical preview. It may be changed in a future release.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type ConfigType string

const (
    // Config for OAuth2. **Warning:** This constant field is available as technical preview. It may be changed in a future release.
     ConfigType_Oauth2 ConfigType = "Oauth2"
    // Config for OIDC. **Warning:** This constant field is available as technical preview. It may be changed in a future release.
     ConfigType_Oidc ConfigType = "Oidc"
)

func (c ConfigType) ConfigType() bool {
    switch c {
        case ConfigType_Oauth2:
            return true
        case ConfigType_Oidc:
            return true
        default:
            return false
    }
}




// The ``IdmProtocol`` class contains the possible types of communication protocols to the identity management endpoints. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type IdmProtocol string

const (
    // REST protocol based identity management endpoints. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     IdmProtocol_REST IdmProtocol = "REST"
    // SCIM protocol based identity management endpoints. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     IdmProtocol_SCIM IdmProtocol = "SCIM"
    // LDAP protocol based identity management endpoints. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     IdmProtocol_LDAP IdmProtocol = "LDAP"
)

func (i IdmProtocol) IdmProtocol() bool {
    switch i {
        case IdmProtocol_REST:
            return true
        case IdmProtocol_SCIM:
            return true
        case IdmProtocol_LDAP:
            return true
        default:
            return false
    }
}




// The ``Oauth2AuthenticationMethod`` class contains the possible types of OAuth2 authentication methods. **Warning:** This enumeration is available as technical preview. It may be changed in a future release.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Oauth2AuthenticationMethod string

const (
    // Clients that have received a client_secret value from the Authorization Server, authenticate with the Authorization Server in accordance with Section 3.2.1 of OAuth 2.0 [RFC6749] using the HTTP Basic authentication scheme. **Warning:** This constant field is available as technical preview. It may be changed in a future release.
     Oauth2AuthenticationMethod_CLIENT_SECRET_BASIC Oauth2AuthenticationMethod = "CLIENT_SECRET_BASIC"
    // Clients that have received a client_secret value from the Authorization Server, authenticate with the Authorization Server in accordance with Section 3.2.1 of OAuth 2.0 [RFC6749] by including the Client Credentials in the request body. **Warning:** This constant field is available as technical preview. It may be changed in a future release.
     Oauth2AuthenticationMethod_CLIENT_SECRET_POST Oauth2AuthenticationMethod = "CLIENT_SECRET_POST"
    // Clients that have received a client_secret value from the Authorization Server, create a JWT using an HMAC SHA algorithm, such as HMAC SHA-256. The HMAC (Hash-based Message Authentication Code) is calculated using the octets of the UTF-8 representation of the client_secret as the shared key. **Warning:** This constant field is available as technical preview. It may be changed in a future release.
     Oauth2AuthenticationMethod_CLIENT_SECRET_JWT Oauth2AuthenticationMethod = "CLIENT_SECRET_JWT"
    // Clients that have registered a public key sign a JWT using that key. The client authenticates in accordance with JSON Web Token (JWT) Profile for OAuth 2.0 Client Authentication and Authorization Grants [OAuth.JWT] and Assertion Framework for OAuth 2.0 Client Authentication and Authorization Grants [OAuth.Assertions]. **Warning:** This constant field is available as technical preview. It may be changed in a future release.
     Oauth2AuthenticationMethod_PRIVATE_KEY_JWT Oauth2AuthenticationMethod = "PRIVATE_KEY_JWT"
)

func (o Oauth2AuthenticationMethod) Oauth2AuthenticationMethod() bool {
    switch o {
        case Oauth2AuthenticationMethod_CLIENT_SECRET_BASIC:
            return true
        case Oauth2AuthenticationMethod_CLIENT_SECRET_POST:
            return true
        case Oauth2AuthenticationMethod_CLIENT_SECRET_JWT:
            return true
        case Oauth2AuthenticationMethod_PRIVATE_KEY_JWT:
            return true
        default:
            return false
    }
}





// The ``Summary`` class contains commonly used information about an identity provider. **Warning:** This class is available as technical preview. It may be changed in a future release.
 type Summary struct {
    Provider string
    Name *string
    ConfigTag ConfigType
    Oauth2 *Oauth2Summary
    Oidc *OidcSummary
    IsDefault bool
    DomainNames map[string]bool
    AuthQueryParams map[string][]string
}






// The ``Info`` class contains the information about an identity provider. **Warning:** This class is available as technical preview. It may be changed in a future release.
 type Info struct {
    Name *string
    OrgIds map[string]bool
    ConfigTag ConfigType
    Oauth2 *Oauth2Info
    Oidc *OidcInfo
    IsDefault bool
    DomainNames map[string]bool
    AuthQueryParams map[string][]string
    IdmProtocol *IdmProtocol
    IdmEndpoints []url.URL
    ActiveDirectoryOverLdap *ActiveDirectoryOverLdap
    UpnClaim *string
    GroupsClaim *string
}






// The ``CreateSpec`` class contains the information used to create an identity provider. **Warning:** This class is available as technical preview. It may be changed in a future release.
 type CreateSpec struct {
    ConfigTag ConfigType
    Oauth2 *Oauth2CreateSpec
    Oidc *OidcCreateSpec
    OrgIds map[string]bool
    IsDefault *bool
    Name *string
    DomainNames map[string]bool
    AuthQueryParams map[string][]string
    IdmProtocol *IdmProtocol
    IdmEndpoints []url.URL
    ActiveDirectoryOverLdap *ActiveDirectoryOverLdap
    UpnClaim *string
    GroupsClaim *string
}






// The ``UpdateSpec`` class contains the information used to update the identity provider. **Warning:** This class is available as technical preview. It may be changed in a future release.
 type UpdateSpec struct {
    ConfigTag ConfigType
    Oauth2 *Oauth2UpdateSpec
    Oidc *OidcUpdateSpec
    OrgIds map[string]bool
    MakeDefault *bool
    Name *string
    DomainNames map[string]bool
    AuthQueryParams map[string][]string
    IdmProtocol *IdmProtocol
    IdmEndpoints []url.URL
    ActiveDirectoryOverLdap *ActiveDirectoryOverLdap
    UpnClaim *string
    ResetUpnClaim *bool
    GroupsClaim *string
    ResetGroupsClaim *bool
}






// The ``Oauth2Summary`` class contains commonly used information about an OAuth2 identity provider. **Warning:** This class is available as technical preview. It may be changed in a future release.
 type Oauth2Summary struct {
    AuthEndpoint url.URL
    TokenEndpoint url.URL
    ClientId string
    AuthenticationHeader string
    AuthQueryParams map[string][]string
}






// The ``Oauth2Info`` class contains the information about an OAuth2 identity provider. **Warning:** This class is available as technical preview. It may be changed in a future release.
 type Oauth2Info struct {
    AuthEndpoint url.URL
    TokenEndpoint url.URL
    PublicKeyUri url.URL
    ClientId string
    ClientSecret string
    ClaimMap map[string]map[string][]string
    Issuer string
    AuthenticationMethod Oauth2AuthenticationMethod
    AuthQueryParams map[string][]string
}






// The ``Oauth2CreateSpec`` class contains the information used to create an OAuth2 identity provider. **Warning:** This class is available as technical preview. It may be changed in a future release.
 type Oauth2CreateSpec struct {
    AuthEndpoint url.URL
    TokenEndpoint url.URL
    PublicKeyUri url.URL
    ClientId string
    ClientSecret string
    ClaimMap map[string]map[string][]string
    Issuer string
    AuthenticationMethod Oauth2AuthenticationMethod
    AuthQueryParams map[string][]string
}






// The ``Oauth2UpdateSpec`` class contains the information used to update the OAuth2 identity provider. **Warning:** This class is available as technical preview. It may be changed in a future release.
 type Oauth2UpdateSpec struct {
    AuthEndpoint *url.URL
    TokenEndpoint *url.URL
    PublicKeyUri *url.URL
    ClientId *string
    ClientSecret *string
    ClaimMap map[string]map[string][]string
    Issuer *string
    AuthenticationMethod *Oauth2AuthenticationMethod
    AuthQueryParams map[string][]string
}






// The ``OidcSummary`` class contains commonly used information about an OIDC identity provider. OIDC is a discovery protocol for OAuth2 configuration metadata, so ``OidcSummary`` contains discovered OAuth2 metadata. **Warning:** This class is available as technical preview. It may be changed in a future release.
 type OidcSummary struct {
    DiscoveryEndpoint *url.URL
    LogoutEndpoint *url.URL
    AuthEndpoint url.URL
    TokenEndpoint url.URL
    ClientId string
    AuthenticationHeader string
    AuthQueryParams map[string][]string
}






// The ``OidcInfo`` class contains information about an OIDC identity provider. OIDC is a discovery protocol for OAuth2 configuration metadata, so ``OidcInfo`` contains additional discovered OAuth2 metadata. **Warning:** This class is available as technical preview. It may be changed in a future release.
 type OidcInfo struct {
    DiscoveryEndpoint url.URL
    LogoutEndpoint *url.URL
    AuthEndpoint url.URL
    TokenEndpoint url.URL
    PublicKeyUri url.URL
    ClientId string
    ClientSecret string
    ClaimMap map[string]map[string][]string
    Issuer string
    AuthenticationMethod Oauth2AuthenticationMethod
    AuthQueryParams map[string][]string
}






// The ``OidcCreateSpec`` class contains the information used to create an OIDC identity provider. **Warning:** This class is available as technical preview. It may be changed in a future release.
 type OidcCreateSpec struct {
    DiscoveryEndpoint url.URL
    ClientId string
    ClientSecret string
    ClaimMap map[string]map[string][]string
}






// The ``OidcUpdateSpec`` class contains the information used to update the OIDC identity provider. **Warning:** This class is available as technical preview. It may be changed in a future release.
 type OidcUpdateSpec struct {
    DiscoveryEndpoint *url.URL
    ClientId *string
    ClientSecret *string
    ClaimMap map[string]map[string][]string
}






// The ``ActiveDirectoryOverLdap`` class contains the information about to how to use an Active Directory over LDAP connection to allow searching for users and groups if the identity provider is an On-Prem service. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ActiveDirectoryOverLdap struct {
    UserName string
    Password string
    UsersBaseDn string
    GroupsBaseDn string
    ServerEndpoints []url.URL
    CertChain *certificate_management.X509CertChain
}









func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(SummaryBindingType), reflect.TypeOf([]Summary{}))
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
       map[string]int{"Unauthorized": 403})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["provider"] = bindings.NewIdType([]string {"com.vmware.vcenter.identity.Providers"}, "")
    fieldNameMap["provider"] = "Provider"
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
       map[string]int{"Unauthorized": 403,"NotFound": 404})
}


func createInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(CreateSpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func createOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"com.vmware.vcenter.identity.Providers"}, "")
}

func createRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Unauthorized": 403,"InvalidArgument": 400})
}


func updateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["provider"] = bindings.NewIdType([]string {"com.vmware.vcenter.identity.Providers"}, "")
    fields["spec"] = bindings.NewReferenceType(UpdateSpecBindingType)
    fieldNameMap["provider"] = "Provider"
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
       map[string]int{"Unauthorized": 403,"InvalidArgument": 400,"NotFound": 404})
}


func deleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["provider"] = bindings.NewIdType([]string {"com.vmware.vcenter.identity.Providers"}, "")
    fieldNameMap["provider"] = "Provider"
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
       map[string]int{"Unauthorized": 403,"NotFound": 404})
}



func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["provider"] = bindings.NewIdType([]string {"com.vmware.vcenter.identity.Providers"}, "")
    fieldNameMap["provider"] = "Provider"
    fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["name"] = "Name"
    fields["config_tag"] = bindings.NewEnumType("com.vmware.vcenter.identity.providers.config_type", reflect.TypeOf(ConfigType(ConfigType_Oauth2)))
    fieldNameMap["config_tag"] = "ConfigTag"
    fields["oauth2"] = bindings.NewOptionalType(bindings.NewReferenceType(Oauth2SummaryBindingType))
    fieldNameMap["oauth2"] = "Oauth2"
    fields["oidc"] = bindings.NewOptionalType(bindings.NewReferenceType(OidcSummaryBindingType))
    fieldNameMap["oidc"] = "Oidc"
    fields["is_default"] = bindings.NewBooleanType()
    fieldNameMap["is_default"] = "IsDefault"
    fields["domain_names"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["domain_names"] = "DomainNames"
    fields["auth_query_params"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})),reflect.TypeOf(map[string][]string{})))
    fieldNameMap["auth_query_params"] = "AuthQueryParams"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("config_tag",
        map[string][]bindings.FieldData {
            "Oauth2": []bindings.FieldData {
                 bindings.NewFieldData("oauth2", true),
            },
            "Oidc": []bindings.FieldData {
                 bindings.NewFieldData("oidc", true),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.identity.providers.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["name"] = "Name"
    fields["org_ids"] = bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{}))
    fieldNameMap["org_ids"] = "OrgIds"
    fields["config_tag"] = bindings.NewEnumType("com.vmware.vcenter.identity.providers.config_type", reflect.TypeOf(ConfigType(ConfigType_Oauth2)))
    fieldNameMap["config_tag"] = "ConfigTag"
    fields["oauth2"] = bindings.NewOptionalType(bindings.NewReferenceType(Oauth2InfoBindingType))
    fieldNameMap["oauth2"] = "Oauth2"
    fields["oidc"] = bindings.NewOptionalType(bindings.NewReferenceType(OidcInfoBindingType))
    fieldNameMap["oidc"] = "Oidc"
    fields["is_default"] = bindings.NewBooleanType()
    fieldNameMap["is_default"] = "IsDefault"
    fields["domain_names"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["domain_names"] = "DomainNames"
    fields["auth_query_params"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})),reflect.TypeOf(map[string][]string{})))
    fieldNameMap["auth_query_params"] = "AuthQueryParams"
    fields["idm_protocol"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.identity.providers.idm_protocol", reflect.TypeOf(IdmProtocol(IdmProtocol_REST))))
    fieldNameMap["idm_protocol"] = "IdmProtocol"
    fields["idm_endpoints"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewUriType(), reflect.TypeOf([]url.URL{})))
    fieldNameMap["idm_endpoints"] = "IdmEndpoints"
    fields["active_directory_over_ldap"] = bindings.NewOptionalType(bindings.NewReferenceType(ActiveDirectoryOverLdapBindingType))
    fieldNameMap["active_directory_over_ldap"] = "ActiveDirectoryOverLdap"
    fields["upn_claim"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["upn_claim"] = "UpnClaim"
    fields["groups_claim"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["groups_claim"] = "GroupsClaim"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("config_tag",
        map[string][]bindings.FieldData {
            "Oauth2": []bindings.FieldData {
                 bindings.NewFieldData("oauth2", true),
            },
            "Oidc": []bindings.FieldData {
                 bindings.NewFieldData("oidc", true),
            },
        },
    )
    validators = append(validators, uv1)
    uv2 := bindings.NewUnionValidator("idm_protocol",
        map[string][]bindings.FieldData {
            "REST": []bindings.FieldData {
                 bindings.NewFieldData("idm_endpoints", true),
            },
            "SCIM": []bindings.FieldData {
                 bindings.NewFieldData("idm_endpoints", true),
            },
            "LDAP": []bindings.FieldData {
                 bindings.NewFieldData("active_directory_over_ldap", true),
            },
        },
    )
    validators = append(validators, uv2)
    return bindings.NewStructType("com.vmware.vcenter.identity.providers.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func CreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["config_tag"] = bindings.NewEnumType("com.vmware.vcenter.identity.providers.config_type", reflect.TypeOf(ConfigType(ConfigType_Oauth2)))
    fieldNameMap["config_tag"] = "ConfigTag"
    fields["oauth2"] = bindings.NewOptionalType(bindings.NewReferenceType(Oauth2CreateSpecBindingType))
    fieldNameMap["oauth2"] = "Oauth2"
    fields["oidc"] = bindings.NewOptionalType(bindings.NewReferenceType(OidcCreateSpecBindingType))
    fieldNameMap["oidc"] = "Oidc"
    fields["org_ids"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["org_ids"] = "OrgIds"
    fields["is_default"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["is_default"] = "IsDefault"
    fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["name"] = "Name"
    fields["domain_names"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["domain_names"] = "DomainNames"
    fields["auth_query_params"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})),reflect.TypeOf(map[string][]string{})))
    fieldNameMap["auth_query_params"] = "AuthQueryParams"
    fields["idm_protocol"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.identity.providers.idm_protocol", reflect.TypeOf(IdmProtocol(IdmProtocol_REST))))
    fieldNameMap["idm_protocol"] = "IdmProtocol"
    fields["idm_endpoints"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewUriType(), reflect.TypeOf([]url.URL{})))
    fieldNameMap["idm_endpoints"] = "IdmEndpoints"
    fields["active_directory_over_ldap"] = bindings.NewOptionalType(bindings.NewReferenceType(ActiveDirectoryOverLdapBindingType))
    fieldNameMap["active_directory_over_ldap"] = "ActiveDirectoryOverLdap"
    fields["upn_claim"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["upn_claim"] = "UpnClaim"
    fields["groups_claim"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["groups_claim"] = "GroupsClaim"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("config_tag",
        map[string][]bindings.FieldData {
            "Oauth2": []bindings.FieldData {
                 bindings.NewFieldData("oauth2", true),
            },
            "Oidc": []bindings.FieldData {
                 bindings.NewFieldData("oidc", true),
            },
        },
    )
    validators = append(validators, uv1)
    uv2 := bindings.NewUnionValidator("idm_protocol",
        map[string][]bindings.FieldData {
            "REST": []bindings.FieldData {
                 bindings.NewFieldData("idm_endpoints", true),
            },
            "SCIM": []bindings.FieldData {
                 bindings.NewFieldData("idm_endpoints", true),
            },
            "LDAP": []bindings.FieldData {
                 bindings.NewFieldData("active_directory_over_ldap", true),
            },
        },
    )
    validators = append(validators, uv2)
    return bindings.NewStructType("com.vmware.vcenter.identity.providers.create_spec",fields, reflect.TypeOf(CreateSpec{}), fieldNameMap, validators)
}

func UpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["config_tag"] = bindings.NewEnumType("com.vmware.vcenter.identity.providers.config_type", reflect.TypeOf(ConfigType(ConfigType_Oauth2)))
    fieldNameMap["config_tag"] = "ConfigTag"
    fields["oauth2"] = bindings.NewOptionalType(bindings.NewReferenceType(Oauth2UpdateSpecBindingType))
    fieldNameMap["oauth2"] = "Oauth2"
    fields["oidc"] = bindings.NewOptionalType(bindings.NewReferenceType(OidcUpdateSpecBindingType))
    fieldNameMap["oidc"] = "Oidc"
    fields["org_ids"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["org_ids"] = "OrgIds"
    fields["make_default"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["make_default"] = "MakeDefault"
    fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["name"] = "Name"
    fields["domain_names"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["domain_names"] = "DomainNames"
    fields["auth_query_params"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})),reflect.TypeOf(map[string][]string{})))
    fieldNameMap["auth_query_params"] = "AuthQueryParams"
    fields["idm_protocol"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.identity.providers.idm_protocol", reflect.TypeOf(IdmProtocol(IdmProtocol_REST))))
    fieldNameMap["idm_protocol"] = "IdmProtocol"
    fields["idm_endpoints"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewUriType(), reflect.TypeOf([]url.URL{})))
    fieldNameMap["idm_endpoints"] = "IdmEndpoints"
    fields["active_directory_over_ldap"] = bindings.NewOptionalType(bindings.NewReferenceType(ActiveDirectoryOverLdapBindingType))
    fieldNameMap["active_directory_over_ldap"] = "ActiveDirectoryOverLdap"
    fields["upn_claim"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["upn_claim"] = "UpnClaim"
    fields["reset_upn_claim"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["reset_upn_claim"] = "ResetUpnClaim"
    fields["groups_claim"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["groups_claim"] = "GroupsClaim"
    fields["reset_groups_claim"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["reset_groups_claim"] = "ResetGroupsClaim"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("config_tag",
        map[string][]bindings.FieldData {
            "Oauth2": []bindings.FieldData {
                 bindings.NewFieldData("oauth2", true),
            },
            "Oidc": []bindings.FieldData {
                 bindings.NewFieldData("oidc", true),
            },
        },
    )
    validators = append(validators, uv1)
    uv2 := bindings.NewUnionValidator("idm_protocol",
        map[string][]bindings.FieldData {
            "REST": []bindings.FieldData {
                 bindings.NewFieldData("idm_endpoints", true),
            },
            "SCIM": []bindings.FieldData {
                 bindings.NewFieldData("idm_endpoints", true),
            },
            "LDAP": []bindings.FieldData {
                 bindings.NewFieldData("active_directory_over_ldap", true),
            },
        },
    )
    validators = append(validators, uv2)
    return bindings.NewStructType("com.vmware.vcenter.identity.providers.update_spec",fields, reflect.TypeOf(UpdateSpec{}), fieldNameMap, validators)
}

func Oauth2SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["auth_endpoint"] = bindings.NewUriType()
    fieldNameMap["auth_endpoint"] = "AuthEndpoint"
    fields["token_endpoint"] = bindings.NewUriType()
    fieldNameMap["token_endpoint"] = "TokenEndpoint"
    fields["client_id"] = bindings.NewStringType()
    fieldNameMap["client_id"] = "ClientId"
    fields["authentication_header"] = bindings.NewStringType()
    fieldNameMap["authentication_header"] = "AuthenticationHeader"
    fields["auth_query_params"] = bindings.NewMapType(bindings.NewStringType(), bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})),reflect.TypeOf(map[string][]string{}))
    fieldNameMap["auth_query_params"] = "AuthQueryParams"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.identity.providers.oauth2_summary",fields, reflect.TypeOf(Oauth2Summary{}), fieldNameMap, validators)
}

func Oauth2InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["auth_endpoint"] = bindings.NewUriType()
    fieldNameMap["auth_endpoint"] = "AuthEndpoint"
    fields["token_endpoint"] = bindings.NewUriType()
    fieldNameMap["token_endpoint"] = "TokenEndpoint"
    fields["public_key_uri"] = bindings.NewUriType()
    fieldNameMap["public_key_uri"] = "PublicKeyUri"
    fields["client_id"] = bindings.NewStringType()
    fieldNameMap["client_id"] = "ClientId"
    fields["client_secret"] = bindings.NewStringType()
    fieldNameMap["client_secret"] = "ClientSecret"
    fields["claim_map"] = bindings.NewMapType(bindings.NewStringType(), bindings.NewMapType(bindings.NewStringType(), bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})),reflect.TypeOf(map[string][]string{})),reflect.TypeOf(map[string]map[string][]string{}))
    fieldNameMap["claim_map"] = "ClaimMap"
    fields["issuer"] = bindings.NewStringType()
    fieldNameMap["issuer"] = "Issuer"
    fields["authentication_method"] = bindings.NewEnumType("com.vmware.vcenter.identity.providers.oauth2_authentication_method", reflect.TypeOf(Oauth2AuthenticationMethod(Oauth2AuthenticationMethod_CLIENT_SECRET_BASIC)))
    fieldNameMap["authentication_method"] = "AuthenticationMethod"
    fields["auth_query_params"] = bindings.NewMapType(bindings.NewStringType(), bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})),reflect.TypeOf(map[string][]string{}))
    fieldNameMap["auth_query_params"] = "AuthQueryParams"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.identity.providers.oauth2_info",fields, reflect.TypeOf(Oauth2Info{}), fieldNameMap, validators)
}

func Oauth2CreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["auth_endpoint"] = bindings.NewUriType()
    fieldNameMap["auth_endpoint"] = "AuthEndpoint"
    fields["token_endpoint"] = bindings.NewUriType()
    fieldNameMap["token_endpoint"] = "TokenEndpoint"
    fields["public_key_uri"] = bindings.NewUriType()
    fieldNameMap["public_key_uri"] = "PublicKeyUri"
    fields["client_id"] = bindings.NewStringType()
    fieldNameMap["client_id"] = "ClientId"
    fields["client_secret"] = bindings.NewStringType()
    fieldNameMap["client_secret"] = "ClientSecret"
    fields["claim_map"] = bindings.NewMapType(bindings.NewStringType(), bindings.NewMapType(bindings.NewStringType(), bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})),reflect.TypeOf(map[string][]string{})),reflect.TypeOf(map[string]map[string][]string{}))
    fieldNameMap["claim_map"] = "ClaimMap"
    fields["issuer"] = bindings.NewStringType()
    fieldNameMap["issuer"] = "Issuer"
    fields["authentication_method"] = bindings.NewEnumType("com.vmware.vcenter.identity.providers.oauth2_authentication_method", reflect.TypeOf(Oauth2AuthenticationMethod(Oauth2AuthenticationMethod_CLIENT_SECRET_BASIC)))
    fieldNameMap["authentication_method"] = "AuthenticationMethod"
    fields["auth_query_params"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})),reflect.TypeOf(map[string][]string{})))
    fieldNameMap["auth_query_params"] = "AuthQueryParams"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.identity.providers.oauth2_create_spec",fields, reflect.TypeOf(Oauth2CreateSpec{}), fieldNameMap, validators)
}

func Oauth2UpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["auth_endpoint"] = bindings.NewOptionalType(bindings.NewUriType())
    fieldNameMap["auth_endpoint"] = "AuthEndpoint"
    fields["token_endpoint"] = bindings.NewOptionalType(bindings.NewUriType())
    fieldNameMap["token_endpoint"] = "TokenEndpoint"
    fields["public_key_uri"] = bindings.NewOptionalType(bindings.NewUriType())
    fieldNameMap["public_key_uri"] = "PublicKeyUri"
    fields["client_id"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["client_id"] = "ClientId"
    fields["client_secret"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["client_secret"] = "ClientSecret"
    fields["claim_map"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewMapType(bindings.NewStringType(), bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})),reflect.TypeOf(map[string][]string{})),reflect.TypeOf(map[string]map[string][]string{})))
    fieldNameMap["claim_map"] = "ClaimMap"
    fields["issuer"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["issuer"] = "Issuer"
    fields["authentication_method"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.identity.providers.oauth2_authentication_method", reflect.TypeOf(Oauth2AuthenticationMethod(Oauth2AuthenticationMethod_CLIENT_SECRET_BASIC))))
    fieldNameMap["authentication_method"] = "AuthenticationMethod"
    fields["auth_query_params"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})),reflect.TypeOf(map[string][]string{})))
    fieldNameMap["auth_query_params"] = "AuthQueryParams"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.identity.providers.oauth2_update_spec",fields, reflect.TypeOf(Oauth2UpdateSpec{}), fieldNameMap, validators)
}

func OidcSummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["discovery_endpoint"] = bindings.NewOptionalType(bindings.NewUriType())
    fieldNameMap["discovery_endpoint"] = "DiscoveryEndpoint"
    fields["logout_endpoint"] = bindings.NewOptionalType(bindings.NewUriType())
    fieldNameMap["logout_endpoint"] = "LogoutEndpoint"
    fields["auth_endpoint"] = bindings.NewUriType()
    fieldNameMap["auth_endpoint"] = "AuthEndpoint"
    fields["token_endpoint"] = bindings.NewUriType()
    fieldNameMap["token_endpoint"] = "TokenEndpoint"
    fields["client_id"] = bindings.NewStringType()
    fieldNameMap["client_id"] = "ClientId"
    fields["authentication_header"] = bindings.NewStringType()
    fieldNameMap["authentication_header"] = "AuthenticationHeader"
    fields["auth_query_params"] = bindings.NewMapType(bindings.NewStringType(), bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})),reflect.TypeOf(map[string][]string{}))
    fieldNameMap["auth_query_params"] = "AuthQueryParams"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.identity.providers.oidc_summary",fields, reflect.TypeOf(OidcSummary{}), fieldNameMap, validators)
}

func OidcInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["discovery_endpoint"] = bindings.NewUriType()
    fieldNameMap["discovery_endpoint"] = "DiscoveryEndpoint"
    fields["logout_endpoint"] = bindings.NewOptionalType(bindings.NewUriType())
    fieldNameMap["logout_endpoint"] = "LogoutEndpoint"
    fields["auth_endpoint"] = bindings.NewUriType()
    fieldNameMap["auth_endpoint"] = "AuthEndpoint"
    fields["token_endpoint"] = bindings.NewUriType()
    fieldNameMap["token_endpoint"] = "TokenEndpoint"
    fields["public_key_uri"] = bindings.NewUriType()
    fieldNameMap["public_key_uri"] = "PublicKeyUri"
    fields["client_id"] = bindings.NewStringType()
    fieldNameMap["client_id"] = "ClientId"
    fields["client_secret"] = bindings.NewStringType()
    fieldNameMap["client_secret"] = "ClientSecret"
    fields["claim_map"] = bindings.NewMapType(bindings.NewStringType(), bindings.NewMapType(bindings.NewStringType(), bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})),reflect.TypeOf(map[string][]string{})),reflect.TypeOf(map[string]map[string][]string{}))
    fieldNameMap["claim_map"] = "ClaimMap"
    fields["issuer"] = bindings.NewStringType()
    fieldNameMap["issuer"] = "Issuer"
    fields["authentication_method"] = bindings.NewEnumType("com.vmware.vcenter.identity.providers.oauth2_authentication_method", reflect.TypeOf(Oauth2AuthenticationMethod(Oauth2AuthenticationMethod_CLIENT_SECRET_BASIC)))
    fieldNameMap["authentication_method"] = "AuthenticationMethod"
    fields["auth_query_params"] = bindings.NewMapType(bindings.NewStringType(), bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})),reflect.TypeOf(map[string][]string{}))
    fieldNameMap["auth_query_params"] = "AuthQueryParams"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.identity.providers.oidc_info",fields, reflect.TypeOf(OidcInfo{}), fieldNameMap, validators)
}

func OidcCreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["discovery_endpoint"] = bindings.NewUriType()
    fieldNameMap["discovery_endpoint"] = "DiscoveryEndpoint"
    fields["client_id"] = bindings.NewStringType()
    fieldNameMap["client_id"] = "ClientId"
    fields["client_secret"] = bindings.NewStringType()
    fieldNameMap["client_secret"] = "ClientSecret"
    fields["claim_map"] = bindings.NewMapType(bindings.NewStringType(), bindings.NewMapType(bindings.NewStringType(), bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})),reflect.TypeOf(map[string][]string{})),reflect.TypeOf(map[string]map[string][]string{}))
    fieldNameMap["claim_map"] = "ClaimMap"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.identity.providers.oidc_create_spec",fields, reflect.TypeOf(OidcCreateSpec{}), fieldNameMap, validators)
}

func OidcUpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["discovery_endpoint"] = bindings.NewOptionalType(bindings.NewUriType())
    fieldNameMap["discovery_endpoint"] = "DiscoveryEndpoint"
    fields["client_id"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["client_id"] = "ClientId"
    fields["client_secret"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["client_secret"] = "ClientSecret"
    fields["claim_map"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewMapType(bindings.NewStringType(), bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})),reflect.TypeOf(map[string][]string{})),reflect.TypeOf(map[string]map[string][]string{})))
    fieldNameMap["claim_map"] = "ClaimMap"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.identity.providers.oidc_update_spec",fields, reflect.TypeOf(OidcUpdateSpec{}), fieldNameMap, validators)
}

func ActiveDirectoryOverLdapBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["user_name"] = bindings.NewStringType()
    fieldNameMap["user_name"] = "UserName"
    fields["password"] = bindings.NewSecretType()
    fieldNameMap["password"] = "Password"
    fields["users_base_dn"] = bindings.NewStringType()
    fieldNameMap["users_base_dn"] = "UsersBaseDn"
    fields["groups_base_dn"] = bindings.NewStringType()
    fieldNameMap["groups_base_dn"] = "GroupsBaseDn"
    fields["server_endpoints"] = bindings.NewListType(bindings.NewUriType(), reflect.TypeOf([]url.URL{}))
    fieldNameMap["server_endpoints"] = "ServerEndpoints"
    fields["cert_chain"] = bindings.NewOptionalType(bindings.NewReferenceType(certificate_management.X509CertChainBindingType))
    fieldNameMap["cert_chain"] = "CertChain"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.identity.providers.active_directory_over_ldap",fields, reflect.TypeOf(ActiveDirectoryOverLdap{}), fieldNameMap, validators)
}


