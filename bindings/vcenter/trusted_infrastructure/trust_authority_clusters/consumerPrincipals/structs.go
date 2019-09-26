/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: ConsumerPrincipals.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package consumerPrincipals

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/vcenter/trusted_infrastructure"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``Health`` enumeration class defines the possible health states.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Health string

const (
    // None. No status available.
     Health_NONE Health = "NONE"
    // OK. Health is normal.
     Health_OK Health = "OK"
    // Warning. Health is normal, however there is an issue that requires attention.
     Health_WARNING Health = "WARNING"
    // Error. Not healthy.
     Health_ERROR Health = "ERROR"
)

func (h Health) Health() bool {
    switch h {
        case Health_NONE:
            return true
        case Health_OK:
            return true
        case Health_WARNING:
            return true
        case Health_ERROR:
            return true
        default:
            return false
    }
}





// The ``CreateSpec`` class contains the information necessary to establish trust between a workload vCenter and a Trust Authority Host.
 type CreateSpec struct {
    Certificates []trusted_infrastructure.X509CertChain
    IssuerAlias string
    Issuer string
    Principal trusted_infrastructure.StsPrincipal
}






// The ``FilterSpec`` class contains data which identifies a connection profile on the trusted vCenter.
 type FilterSpec struct {
    Id map[string]bool
    Principals []trusted_infrastructure.StsPrincipal
    Issuer map[string]bool
}






// The ``Info`` class contains the information necessary to establish trust between a workload vCenter and a Trust Authority Host.
 type Info struct {
    Id string
    Principal trusted_infrastructure.StsPrincipal
    IssuerAlias string
    Issuer string
    Certificates []trusted_infrastructure.X509CertChain
    Health Health
    Message *std.LocalizableMessage
}






// The ``Summary`` class contains a summary of the information necessary to establish trust between a workload vCenter and a Trust Authority Host.
 type Summary struct {
    Id string
    Principal trusted_infrastructure.StsPrincipal
    IssuerAlias string
    Issuer string
}









func createInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fields["spec"] = bindings.NewReferenceType(CreateSpecBindingType)
    fieldNameMap["cluster"] = "Cluster"
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
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    paramsTypeMap["spec"] = bindings.NewReferenceType(CreateSpecBindingType)
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    pathParams["cluster"] = "cluster"
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
      "/vcenter/trusted-infrastructure/trust-authority-clusters/{cluster}/consumer-principals",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"AlreadyExists": 400,"Error": 500,"NotFound": 404,"Unauthenticated": 401})
}


func deleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fields["profile"] = bindings.NewIdType([]string {"com.vmware.esx.authentication.clientprofile"}, "")
    fieldNameMap["cluster"] = "Cluster"
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
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    paramsTypeMap["profile"] = bindings.NewIdType([]string {"com.vmware.esx.authentication.clientprofile"}, "")
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    paramsTypeMap["profile"] = bindings.NewIdType([]string {"com.vmware.esx.authentication.clientprofile"}, "")
    pathParams["cluster"] = "cluster"
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
      "/vcenter/trusted-infrastructure/trust-authority-clusters/{cluster}/consumer-principals/{profile}",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"NotFound": 404,"Unauthenticated": 401})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fields["profile"] = bindings.NewIdType([]string {"com.vmware.esx.authentication.clientprofile"}, "")
    fieldNameMap["cluster"] = "Cluster"
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
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    paramsTypeMap["profile"] = bindings.NewIdType([]string {"com.vmware.esx.authentication.clientprofile"}, "")
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    paramsTypeMap["profile"] = bindings.NewIdType([]string {"com.vmware.esx.authentication.clientprofile"}, "")
    pathParams["cluster"] = "cluster"
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
      "/vcenter/trusted-infrastructure/trust-authority-clusters/{cluster}/consumer-principals/{profile}",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"NotFound": 404,"Unauthenticated": 401})
}


func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fields["spec"] = bindings.NewOptionalType(bindings.NewReferenceType(FilterSpecBindingType))
    fieldNameMap["cluster"] = "Cluster"
    fieldNameMap["spec"] = "Spec"
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
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    paramsTypeMap["spec"] = bindings.NewOptionalType(bindings.NewReferenceType(FilterSpecBindingType))
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    pathParams["cluster"] = "cluster"
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
      "/vcenter/trusted-infrastructure/trust-authority-clusters/{cluster}/consumer-principals",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"NotFound": 404,"Unauthenticated": 401})
}



func CreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["certificates"] = bindings.NewListType(bindings.NewReferenceType(trusted_infrastructure.X509CertChainBindingType), reflect.TypeOf([]trusted_infrastructure.X509CertChain{}))
    fieldNameMap["certificates"] = "Certificates"
    fields["issuer_alias"] = bindings.NewIdType([]string {"com.vmware.esx.authentication.trust.security-token-issuer"}, "")
    fieldNameMap["issuer_alias"] = "IssuerAlias"
    fields["issuer"] = bindings.NewStringType()
    fieldNameMap["issuer"] = "Issuer"
    fields["principal"] = bindings.NewReferenceType(trusted_infrastructure.StsPrincipalBindingType)
    fieldNameMap["principal"] = "Principal"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.consumer_principals.create_spec",fields, reflect.TypeOf(CreateSpec{}), fieldNameMap, validators)
}

func FilterSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"com.vmware.esx.authentication.clientprofile"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["id"] = "Id"
    fields["principals"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(trusted_infrastructure.StsPrincipalBindingType), reflect.TypeOf([]trusted_infrastructure.StsPrincipal{})))
    fieldNameMap["principals"] = "Principals"
    fields["issuer"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"com.vmware.esx.authentication.trust.security-token-issuer"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["issuer"] = "Issuer"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.consumer_principals.filter_spec",fields, reflect.TypeOf(FilterSpec{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.esx.authentication.clientprofile"}, "")
    fieldNameMap["id"] = "Id"
    fields["principal"] = bindings.NewReferenceType(trusted_infrastructure.StsPrincipalBindingType)
    fieldNameMap["principal"] = "Principal"
    fields["issuer_alias"] = bindings.NewIdType([]string {"com.vmware.esx.authentication.trust.security-token-issuer"}, "")
    fieldNameMap["issuer_alias"] = "IssuerAlias"
    fields["issuer"] = bindings.NewStringType()
    fieldNameMap["issuer"] = "Issuer"
    fields["certificates"] = bindings.NewListType(bindings.NewReferenceType(trusted_infrastructure.X509CertChainBindingType), reflect.TypeOf([]trusted_infrastructure.X509CertChain{}))
    fieldNameMap["certificates"] = "Certificates"
    fields["health"] = bindings.NewEnumType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.consumer_principals.health", reflect.TypeOf(Health(Health_NONE)))
    fieldNameMap["health"] = "Health"
    fields["message"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
    fieldNameMap["message"] = "Message"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.consumer_principals.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.esx.authentication.clientprofile"}, "")
    fieldNameMap["id"] = "Id"
    fields["principal"] = bindings.NewReferenceType(trusted_infrastructure.StsPrincipalBindingType)
    fieldNameMap["principal"] = "Principal"
    fields["issuer_alias"] = bindings.NewIdType([]string {"com.vmware.esx.authentication.trust.security-token-issuer"}, "")
    fieldNameMap["issuer_alias"] = "IssuerAlias"
    fields["issuer"] = bindings.NewStringType()
    fieldNameMap["issuer"] = "Issuer"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.consumer_principals.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}


