/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Attestation.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package attestation

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vcenter/trusted_infrastructure"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The connection information could include the certificates or be a shorter summary.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type SummaryType string

const (
    // The full connection information, including certificates.
     SummaryType_FULL SummaryType = "FULL"
    // A summary containing only the hostname, port, and the group ID which determines the Attestation Services this Attestation Service can communicate with.
     SummaryType_NORMAL SummaryType = "NORMAL"
    // A brief summary, containing only the hostname for the Attestation Service.
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





// The ``Summary`` class contains all the stored information about a Attestation Service.
 type Summary struct {
    SummaryType SummaryType
    Host *string
    Address *trusted_infrastructure.NetworkAddress
    Group *string
    Cluster *string
    TrustedCA *trusted_infrastructure.X509CertChain
}






// The ``Info`` class contains all the stored information about a Attestation Service.
 type Info struct {
    Host string
    Address trusted_infrastructure.NetworkAddress
    Group string
    Cluster string
    TrustedCA trusted_infrastructure.X509CertChain
}






// The ``FilterSpec`` class contains the data necessary for identifying a Attestation Service
 type FilterSpec struct {
    Hosts map[string]bool
    Clusters map[string]bool
    Address []trusted_infrastructure.NetworkAddress
    Groups map[string]bool
}









func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["host"] = bindings.NewIdType([]string {"HostSystem"}, "")
    fieldNameMap["host"] = "Host"
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
    paramsTypeMap["host"] = bindings.NewIdType([]string {"HostSystem"}, "")
    paramsTypeMap["host"] = bindings.NewIdType([]string {"HostSystem"}, "")
    pathParams["host"] = "host"
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
      "/vcenter/trusted-infrastructure/trust-authority-hosts/{host}/attestation/",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"NotFound": 404,"ResourceInaccessible": 500,"Unauthenticated": 401})
}


func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewOptionalType(bindings.NewReferenceType(FilterSpecBindingType))
    fields["projection"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.trusted_infrastructure.trust_authority_hosts.attestation.summary_type", reflect.TypeOf(SummaryType(SummaryType_FULL))))
    fieldNameMap["spec"] = "Spec"
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
    paramsTypeMap["projection"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.trusted_infrastructure.trust_authority_hosts.attestation.summary_type", reflect.TypeOf(SummaryType(SummaryType_FULL))))
    paramsTypeMap["spec"] = bindings.NewOptionalType(bindings.NewReferenceType(FilterSpecBindingType))
    queryParams["projection"] = "projection"
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
      "/vcenter/trusted-infrastructure/trust-authority-hosts/attestation",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"Unauthenticated": 401})
}



func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["summary_type"] = bindings.NewEnumType("com.vmware.vcenter.trusted_infrastructure.trust_authority_hosts.attestation.summary_type", reflect.TypeOf(SummaryType(SummaryType_FULL)))
    fieldNameMap["summary_type"] = "SummaryType"
    fields["host"] = bindings.NewOptionalType(bindings.NewIdType([]string {"HostSystem"}, ""))
    fieldNameMap["host"] = "Host"
    fields["address"] = bindings.NewOptionalType(bindings.NewReferenceType(trusted_infrastructure.NetworkAddressBindingType))
    fieldNameMap["address"] = "Address"
    fields["group"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["group"] = "Group"
    fields["cluster"] = bindings.NewOptionalType(bindings.NewIdType([]string {"HostSystem"}, ""))
    fieldNameMap["cluster"] = "Cluster"
    fields["trusted_CA"] = bindings.NewOptionalType(bindings.NewReferenceType(trusted_infrastructure.X509CertChainBindingType))
    fieldNameMap["trusted_CA"] = "TrustedCA"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("summary_type",
        map[string][]bindings.FieldData {
            "BRIEF": []bindings.FieldData {
                 bindings.NewFieldData("host", true),
                 bindings.NewFieldData("address", true),
            },
            "NORMAL": []bindings.FieldData {
                 bindings.NewFieldData("host", true),
                 bindings.NewFieldData("address", true),
                 bindings.NewFieldData("group", true),
                 bindings.NewFieldData("cluster", true),
            },
            "FULL": []bindings.FieldData {
                 bindings.NewFieldData("host", true),
                 bindings.NewFieldData("address", true),
                 bindings.NewFieldData("group", true),
                 bindings.NewFieldData("cluster", true),
                 bindings.NewFieldData("trusted_CA", true),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_hosts.attestation.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["host"] = bindings.NewIdType([]string {"HostSystem"}, "")
    fieldNameMap["host"] = "Host"
    fields["address"] = bindings.NewReferenceType(trusted_infrastructure.NetworkAddressBindingType)
    fieldNameMap["address"] = "Address"
    fields["group"] = bindings.NewStringType()
    fieldNameMap["group"] = "Group"
    fields["cluster"] = bindings.NewIdType([]string {"HostSystem"}, "")
    fieldNameMap["cluster"] = "Cluster"
    fields["trusted_CA"] = bindings.NewReferenceType(trusted_infrastructure.X509CertChainBindingType)
    fieldNameMap["trusted_CA"] = "TrustedCA"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_hosts.attestation.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func FilterSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["hosts"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"HostSystem"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["hosts"] = "Hosts"
    fields["clusters"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"ClusterComputeResource"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["clusters"] = "Clusters"
    fields["address"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(trusted_infrastructure.NetworkAddressBindingType), reflect.TypeOf([]trusted_infrastructure.NetworkAddress{})))
    fieldNameMap["address"] = "Address"
    fields["groups"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["groups"] = "Groups"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_hosts.attestation.filter_spec",fields, reflect.TypeOf(FilterSpec{}), fieldNameMap, validators)
}


