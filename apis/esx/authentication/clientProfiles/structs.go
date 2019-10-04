/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: ClientProfiles.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package clientProfiles

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// The resource type for the ClientProfiles instances.
const ClientProfiles_RESOURCE_TYPE = "com.vmware.esx.authentication.clientprofile"


// The ``ResourceType`` enumeration class defines the types of AccessGrant elements in a client profile. These are permission resource types. There is support for entitlements, but not for groups.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type ResourceType string

const (
    // Permission entitlements. 
    //
    //  These are coarse-grained permissions that are not associated with an object, i.e. they are system-wide.
     ResourceType_ENTITLEMENT ResourceType = "ENTITLEMENT"
)

func (r ResourceType) ResourceType() bool {
    switch r {
        case ResourceType_ENTITLEMENT:
            return true
        default:
            return false
    }
}




// The ``Entitlement`` enumeration class defines all permission entitlements supported on the ESX.
//
//  These are coarse-grained permissions that are not associated with an object, i.e. they are system-wide.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Entitlement string

const (
    // Allows modifying the identity configuration. 
    //
    //  For example: ClientProfiles, SecurityTokenIssuers.
     Entitlement_IDENTITY_MGMT Entitlement = "IDENTITY_MGMT"
    // Allows modifying security configuration. 
    //
    //  For example: KMS, Attestation.
     Entitlement_SECURITY_MGMT Entitlement = "SECURITY_MGMT"
    // Allows access to some read-only methods. Not all read-only methods are accessible with this entitlement. Check the specific method documentation for the required authorization.
     Entitlement_READ_ONLY Entitlement = "READ_ONLY"
)

func (e Entitlement) Entitlement() bool {
    switch e {
        case Entitlement_IDENTITY_MGMT:
            return true
        case Entitlement_SECURITY_MGMT:
            return true
        case Entitlement_READ_ONLY:
            return true
        default:
            return false
    }
}




// The ``SubjectType`` enumeration class defines the types of subject matching that a client profile is associated with.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type SubjectType string

const (
    // local user
     SubjectType_LOCAL_USER SubjectType = "LOCAL_USER"
    // A group from external source.
     SubjectType_EXTERNAL_GROUP SubjectType = "EXTERNAL_GROUP"
    // A user from external source.
     SubjectType_EXTERNAL_USER SubjectType = "EXTERNAL_USER"
)

func (s SubjectType) SubjectType() bool {
    switch s {
        case SubjectType_LOCAL_USER:
            return true
        case SubjectType_EXTERNAL_GROUP:
            return true
        case SubjectType_EXTERNAL_USER:
            return true
        default:
            return false
    }
}




// The ``SummaryType`` enumeration class defines the types of Summary members to return from the ClientProfiles#list method.
//
//  The profile information could include the access grants or be a shorter summary.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type SummaryType string

const (
    // The full profile information, including access grants.
     SummaryType_FULL SummaryType = "FULL"
    // A summary containing only the profile identifier and the subject information.
     SummaryType_NORMAL SummaryType = "NORMAL"
    // A brief summary, containing only the profile identifier.
     SummaryType_BRIEF SummaryType = "BRIEF"
)

func (s SummaryType) SummaryType() bool {
    switch s {
        case SummaryType_FULL:
            return true
        case SummaryType_NORMAL:
            return true
        case SummaryType_BRIEF:
            return true
        default:
            return false
    }
}





// The ``FilterSpec`` class contains information about filtering the list of existing ClientProfiles.
//
//  The structure includes local user, external group or external user specification.
//
//  Only one of the subject types can be specified. The filter is appied using exact match of all fields.
 type FilterSpec struct {
    LocalUserName *string
    ExternalGroupName *string
    ExternalUserName *string
    IssuerAlias *string
    Domain *string
}






// The ``AccessGrant`` class contains information about access permissions.
//
//  The structure includes resource type and the specific resource information - entitlement.
 type AccessGrant struct {
    ResourceType ResourceType
    Entitlement *Entitlement
}






// The ``Subject`` class contains information about the subject that a client profile is associated with.
//
//  The structure includes a name, issuer alias, and domain.
 type Subject struct {
    Type_ SubjectType
    Name *string
    IssuerAlias *string
    Domain *string
}






// The ``Summary`` class contains summary from the list of existing ClientProfiles.
//
//  The structure includes a client profile identifier, subject type, the details of the subject - local user, external user or external group.
 type Summary struct {
    SummaryType SummaryType
    ClientProfile *string
    Subject *Subject
    Grants []AccessGrant
}






// The ``Info`` class contains information about an existing client profile.
//
//  The structure includes a subject type, the details of the subject - local user, external user or external group, and a list of access grants. 
//
// * A local user is a user account configured on the ESX system.
// * An external user is a user account configured in an external for the ESX identity provider.
// * An external group is a group account configured in an external for the ESX identity provider.
 type Info struct {
    Subject Subject
    Grants []AccessGrant
}






// The ``CreateSpec`` class contains fields to be specified for creating a new client profile.
//
//  The structure includes a subject - local user, external user or external group, and a list of access grants.
 type CreateSpec struct {
    LocalUserName *string
    ExternalGroupName *string
    ExternalUserName *string
    IssuerAlias *string
    Domain *string
    Grants []AccessGrant
}






// The ``UpdateSpec`` class contains the fields of the existing client profile which can be updated.
//
//  The structure includes a list of access grants.
 type UpdateSpec struct {
    Grants []AccessGrant
}









func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(FilterSpecBindingType))
    fields["projection"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.esx.authentication.client_profiles.summary_type", reflect.TypeOf(SummaryType(SummaryType_FULL))))
    fieldNameMap["filter"] = "Filter"
    fieldNameMap["projection"] = "Projection"
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
    paramsTypeMap["projection"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.esx.authentication.client_profiles.summary_type", reflect.TypeOf(SummaryType(SummaryType_FULL))))
    paramsTypeMap["filter.domain"] = bindings.NewOptionalType(bindings.NewStringType())
    paramsTypeMap["filter.external_user_name"] = bindings.NewOptionalType(bindings.NewStringType())
    paramsTypeMap["filter.external_group_name"] = bindings.NewOptionalType(bindings.NewStringType())
    paramsTypeMap["filter.issuer_alias"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.esx.authentication.trust.security-token-issuer"}, ""))
    paramsTypeMap["filter.local_user_name"] = bindings.NewOptionalType(bindings.NewStringType())
    queryParams["filter.external_user_name"] = "external_user_name"
    queryParams["filter.local_user_name"] = "local_user_name"
    queryParams["filter.domain"] = "domain"
    queryParams["filter.external_group_name"] = "external_group_name"
    queryParams["filter.issuer_alias"] = "issuer_alias"
    queryParams["projection"] = "projection"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "GET",
      "/esx/authentication/client-profiles",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"Unauthenticated": 401})
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
    return bindings.NewIdType([]string {"com.vmware.esx.authentication.clientprofile"}, "")
}

func createRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["spec"] = bindings.NewReferenceType(CreateSpecBindingType)
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "spec",
      "POST",
      "/esx/authentication/client-profiles",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"AlreadyExists": 400,"InvalidArgument": 400,"Error": 500,"Unauthenticated": 401})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["profile"] = bindings.NewIdType([]string {"com.vmware.esx.authentication.clientprofile"}, "")
    fieldNameMap["profile"] = "Profile"
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
    paramsTypeMap["profile"] = bindings.NewIdType([]string {"com.vmware.esx.authentication.clientprofile"}, "")
    paramsTypeMap["profile"] = bindings.NewIdType([]string {"com.vmware.esx.authentication.clientprofile"}, "")
    pathParams["profile"] = "profile"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "GET",
      "/esx/authentication/client-profiles/{profile}",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"NotFound": 404,"Error": 500,"Unauthenticated": 401})
}


func updateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["profile"] = bindings.NewIdType([]string {"com.vmware.esx.authentication.clientprofile"}, "")
    fields["spec"] = bindings.NewReferenceType(UpdateSpecBindingType)
    fieldNameMap["profile"] = "Profile"
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
    paramsTypeMap["spec"] = bindings.NewReferenceType(UpdateSpecBindingType)
    paramsTypeMap["profile"] = bindings.NewIdType([]string {"com.vmware.esx.authentication.clientprofile"}, "")
    paramsTypeMap["profile"] = bindings.NewIdType([]string {"com.vmware.esx.authentication.clientprofile"}, "")
    pathParams["profile"] = "profile"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "spec",
      "PATCH",
      "/esx/authentication/client-profiles/{profile}",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"NotFound": 404,"InvalidArgument": 400,"Error": 500,"Unauthenticated": 401})
}


func deleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["profile"] = bindings.NewIdType([]string {"com.vmware.esx.authentication.clientprofile"}, "")
    fieldNameMap["profile"] = "Profile"
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
    paramsTypeMap["profile"] = bindings.NewIdType([]string {"com.vmware.esx.authentication.clientprofile"}, "")
    paramsTypeMap["profile"] = bindings.NewIdType([]string {"com.vmware.esx.authentication.clientprofile"}, "")
    pathParams["profile"] = "profile"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "DELETE",
      "/esx/authentication/client-profiles/{profile}",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"NotFound": 404,"Error": 500,"Unauthenticated": 401})
}



func FilterSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["local_user_name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["local_user_name"] = "LocalUserName"
    fields["external_group_name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["external_group_name"] = "ExternalGroupName"
    fields["external_user_name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["external_user_name"] = "ExternalUserName"
    fields["issuer_alias"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.esx.authentication.trust.security-token-issuer"}, ""))
    fieldNameMap["issuer_alias"] = "IssuerAlias"
    fields["domain"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["domain"] = "Domain"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.authentication.client_profiles.filter_spec",fields, reflect.TypeOf(FilterSpec{}), fieldNameMap, validators)
}

func AccessGrantBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["resource_type"] = bindings.NewEnumType("com.vmware.esx.authentication.client_profiles.resource_type", reflect.TypeOf(ResourceType(ResourceType_ENTITLEMENT)))
    fieldNameMap["resource_type"] = "ResourceType"
    fields["entitlement"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.esx.authentication.client_profiles.entitlement", reflect.TypeOf(Entitlement(Entitlement_IDENTITY_MGMT))))
    fieldNameMap["entitlement"] = "Entitlement"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("resource_type",
        map[string][]bindings.FieldData {
            "ENTITLEMENT": []bindings.FieldData {
                 bindings.NewFieldData("entitlement", true),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.esx.authentication.client_profiles.access_grant",fields, reflect.TypeOf(AccessGrant{}), fieldNameMap, validators)
}

func SubjectBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewEnumType("com.vmware.esx.authentication.client_profiles.subject_type", reflect.TypeOf(SubjectType(SubjectType_LOCAL_USER)))
    fieldNameMap["type"] = "Type_"
    fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["name"] = "Name"
    fields["issuer_alias"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.esx.authentication.trust.security-token-issuer"}, ""))
    fieldNameMap["issuer_alias"] = "IssuerAlias"
    fields["domain"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["domain"] = "Domain"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("type",
        map[string][]bindings.FieldData {
            "LOCAL_USER": []bindings.FieldData {
                 bindings.NewFieldData("name", true),
            },
            "EXTERNAL_GROUP": []bindings.FieldData {
                 bindings.NewFieldData("name", true),
                 bindings.NewFieldData("issuer_alias", true),
                 bindings.NewFieldData("domain", true),
            },
            "EXTERNAL_USER": []bindings.FieldData {
                 bindings.NewFieldData("name", true),
                 bindings.NewFieldData("issuer_alias", true),
                 bindings.NewFieldData("domain", true),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.esx.authentication.client_profiles.subject",fields, reflect.TypeOf(Subject{}), fieldNameMap, validators)
}

func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["summary_type"] = bindings.NewEnumType("com.vmware.esx.authentication.client_profiles.summary_type", reflect.TypeOf(SummaryType(SummaryType_FULL)))
    fieldNameMap["summary_type"] = "SummaryType"
    fields["client_profile"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.esx.authentication.clientprofile"}, ""))
    fieldNameMap["client_profile"] = "ClientProfile"
    fields["subject"] = bindings.NewOptionalType(bindings.NewReferenceType(SubjectBindingType))
    fieldNameMap["subject"] = "Subject"
    fields["grants"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(AccessGrantBindingType), reflect.TypeOf([]AccessGrant{})))
    fieldNameMap["grants"] = "Grants"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("summary_type",
        map[string][]bindings.FieldData {
            "BRIEF": []bindings.FieldData {
                 bindings.NewFieldData("client_profile", true),
            },
            "NORMAL": []bindings.FieldData {
                 bindings.NewFieldData("client_profile", true),
                 bindings.NewFieldData("subject", true),
            },
            "FULL": []bindings.FieldData {
                 bindings.NewFieldData("client_profile", true),
                 bindings.NewFieldData("subject", true),
                 bindings.NewFieldData("grants", true),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.esx.authentication.client_profiles.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["subject"] = bindings.NewReferenceType(SubjectBindingType)
    fieldNameMap["subject"] = "Subject"
    fields["grants"] = bindings.NewListType(bindings.NewReferenceType(AccessGrantBindingType), reflect.TypeOf([]AccessGrant{}))
    fieldNameMap["grants"] = "Grants"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.authentication.client_profiles.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func CreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["local_user_name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["local_user_name"] = "LocalUserName"
    fields["external_group_name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["external_group_name"] = "ExternalGroupName"
    fields["external_user_name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["external_user_name"] = "ExternalUserName"
    fields["issuer_alias"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.esx.authentication.trust.security-token-issuer"}, ""))
    fieldNameMap["issuer_alias"] = "IssuerAlias"
    fields["domain"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["domain"] = "Domain"
    fields["grants"] = bindings.NewListType(bindings.NewReferenceType(AccessGrantBindingType), reflect.TypeOf([]AccessGrant{}))
    fieldNameMap["grants"] = "Grants"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.authentication.client_profiles.create_spec",fields, reflect.TypeOf(CreateSpec{}), fieldNameMap, validators)
}

func UpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["grants"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(AccessGrantBindingType), reflect.TypeOf([]AccessGrant{})))
    fieldNameMap["grants"] = "Grants"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.authentication.client_profiles.update_spec",fields, reflect.TypeOf(UpdateSpec{}), fieldNameMap, validators)
}


