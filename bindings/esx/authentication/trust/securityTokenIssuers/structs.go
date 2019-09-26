/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: SecurityTokenIssuers.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package securityTokenIssuers

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// The resource type for the SecurityTokenIssuers instances.
const SecurityTokenIssuers_RESOURCE_TYPE = "com.vmware.esx.authentication.trust.security-token-issuer"


// The ``SummaryType`` enumeration class defines the types of Summary members to return from the SecurityTokenIssuers#list method.
//
//  The profile information could include the access grants or be a shorter summary.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type SummaryType string

const (
    // The full security token isuer information, including certificates.
     SummaryType_FULL SummaryType = "FULL"
    // A summary containing only the security token issuer alias and the issuer string.
     SummaryType_NORMAL SummaryType = "NORMAL"
    // A brief summary, containing only the security token issuer alias.
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





// The ``X509CertChain`` class contains information about a chain of X.509 certificates.
//
//  The structure includes a list of strings, each string is a PEM encoded certificate.
 type X509CertChain struct {
    CertChain []string
}






// The ``IssuerAlreadyExistsInfo`` class contains the information under which alias the issuer is.
 type IssuerAlreadyExistsInfo struct {
    Issuer string
    IssuerAlias string
}






// The ``Info`` class contains information about an existing security token issuer trust.
//
//  The structure includes an issuer and a list of token signing certificate chains.
 type Info struct {
    Issuer string
    SigningCertChains []X509CertChain
}






// The ``Summary`` class contains summary from the list of existing security token issuer trusts.
//
//  The structure includes the alias identifier and the issuer string.
 type Summary struct {
    SummaryType SummaryType
    IssuerAlias *string
    Issuer *string
    SigningCertChains []X509CertChain
}






// The ``CreateSpec`` class contains fields to be specified for creating a new security token issuer trust. The structure includes an alias identifier, an issuer and a list of certificate chains.
 type CreateSpec struct {
    IssuerAlias string
    Issuer string
    SigningCertChains []X509CertChain
}






// The ``UpdateSpec`` class contains the fields of the existing security token issuer trust which can be updated.
//
//  The structure includes an issuer and a list of token signing certificate chains.
 type UpdateSpec struct {
    Issuer *string
    SigningCertChains []X509CertChain
}









func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["projection"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.esx.authentication.trust.security_token_issuers.summary_type", reflect.TypeOf(SummaryType(SummaryType_FULL))))
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
    paramsTypeMap["projection"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.esx.authentication.trust.security_token_issuers.summary_type", reflect.TypeOf(SummaryType(SummaryType_FULL))))
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
      "/esx/authentication/trust/security-token-issuers",
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
    return bindings.NewVoidType()
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
      "/esx/authentication/trust/security-token-issuers",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"AlreadyExists": 400,"InvalidArgument": 400,"Error": 500,"Unauthenticated": 401})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["issuer_alias"] = bindings.NewIdType([]string {"com.vmware.esx.authentication.trust.security-token-issuer"}, "")
    fieldNameMap["issuer_alias"] = "IssuerAlias"
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
    paramsTypeMap["issuer_alias"] = bindings.NewIdType([]string {"com.vmware.esx.authentication.trust.security-token-issuer"}, "")
    paramsTypeMap["issuerAlias"] = bindings.NewIdType([]string {"com.vmware.esx.authentication.trust.security-token-issuer"}, "")
    pathParams["issuer_alias"] = "issuerAlias"
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
      "/esx/authentication/trust/security-token-issuers/{issuerAlias}",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"NotFound": 404,"Error": 500,"Unauthenticated": 401})
}


func updateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["issuer_alias"] = bindings.NewIdType([]string {"com.vmware.esx.authentication.trust.security-token-issuer"}, "")
    fields["spec"] = bindings.NewReferenceType(UpdateSpecBindingType)
    fieldNameMap["issuer_alias"] = "IssuerAlias"
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
    paramsTypeMap["issuer_alias"] = bindings.NewIdType([]string {"com.vmware.esx.authentication.trust.security-token-issuer"}, "")
    paramsTypeMap["issuerAlias"] = bindings.NewIdType([]string {"com.vmware.esx.authentication.trust.security-token-issuer"}, "")
    pathParams["issuer_alias"] = "issuerAlias"
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
      "/esx/authentication/trust/security-token-issuers/{issuerAlias}",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"NotFound": 404,"Error": 500,"InvalidArgument": 400,"Unauthenticated": 401})
}


func deleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["issuer_alias"] = bindings.NewIdType([]string {"com.vmware.esx.authentication.trust.security-token-issuer"}, "")
    fieldNameMap["issuer_alias"] = "IssuerAlias"
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
    paramsTypeMap["issuer_alias"] = bindings.NewIdType([]string {"com.vmware.esx.authentication.trust.security-token-issuer"}, "")
    paramsTypeMap["issuerAlias"] = bindings.NewIdType([]string {"com.vmware.esx.authentication.trust.security-token-issuer"}, "")
    pathParams["issuer_alias"] = "issuerAlias"
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
      "/esx/authentication/trust/security-token-issuers/{issuerAlias}",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"NotFound": 404,"Error": 500,"Unauthenticated": 401})
}



func X509CertChainBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cert_chain"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["cert_chain"] = "CertChain"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.authentication.trust.security_token_issuers.x509_cert_chain",fields, reflect.TypeOf(X509CertChain{}), fieldNameMap, validators)
}

func IssuerAlreadyExistsInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["issuer"] = bindings.NewStringType()
    fieldNameMap["issuer"] = "Issuer"
    fields["issuer_alias"] = bindings.NewStringType()
    fieldNameMap["issuer_alias"] = "IssuerAlias"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.authentication.trust.security_token_issuers.issuer_already_exists_info",fields, reflect.TypeOf(IssuerAlreadyExistsInfo{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["issuer"] = bindings.NewStringType()
    fieldNameMap["issuer"] = "Issuer"
    fields["signing_cert_chains"] = bindings.NewListType(bindings.NewReferenceType(X509CertChainBindingType), reflect.TypeOf([]X509CertChain{}))
    fieldNameMap["signing_cert_chains"] = "SigningCertChains"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.authentication.trust.security_token_issuers.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["summary_type"] = bindings.NewEnumType("com.vmware.esx.authentication.trust.security_token_issuers.summary_type", reflect.TypeOf(SummaryType(SummaryType_FULL)))
    fieldNameMap["summary_type"] = "SummaryType"
    fields["issuer_alias"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.esx.authentication.trust.security-token-issuer"}, ""))
    fieldNameMap["issuer_alias"] = "IssuerAlias"
    fields["issuer"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["issuer"] = "Issuer"
    fields["signing_cert_chains"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(X509CertChainBindingType), reflect.TypeOf([]X509CertChain{})))
    fieldNameMap["signing_cert_chains"] = "SigningCertChains"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("summary_type",
        map[string][]bindings.FieldData {
            "BRIEF": []bindings.FieldData {
                 bindings.NewFieldData("issuer_alias", true),
            },
            "NORMAL": []bindings.FieldData {
                 bindings.NewFieldData("issuer_alias", true),
                 bindings.NewFieldData("issuer", true),
            },
            "FULL": []bindings.FieldData {
                 bindings.NewFieldData("issuer_alias", true),
                 bindings.NewFieldData("issuer", true),
                 bindings.NewFieldData("signing_cert_chains", true),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.esx.authentication.trust.security_token_issuers.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}

func CreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["issuer_alias"] = bindings.NewIdType([]string {"com.vmware.esx.authentication.trust.security-token-issuer"}, "")
    fieldNameMap["issuer_alias"] = "IssuerAlias"
    fields["issuer"] = bindings.NewStringType()
    fieldNameMap["issuer"] = "Issuer"
    fields["signing_cert_chains"] = bindings.NewListType(bindings.NewReferenceType(X509CertChainBindingType), reflect.TypeOf([]X509CertChain{}))
    fieldNameMap["signing_cert_chains"] = "SigningCertChains"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.authentication.trust.security_token_issuers.create_spec",fields, reflect.TypeOf(CreateSpec{}), fieldNameMap, validators)
}

func UpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["issuer"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["issuer"] = "Issuer"
    fields["signing_cert_chains"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(X509CertChainBindingType), reflect.TypeOf([]X509CertChain{})))
    fieldNameMap["signing_cert_chains"] = "SigningCertChains"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.authentication.trust.security_token_issuers.update_spec",fields, reflect.TypeOf(UpdateSpec{}), fieldNameMap, validators)
}


