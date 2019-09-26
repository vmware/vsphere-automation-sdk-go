/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Nodes.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package nodes

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``ApplianceType`` enumeration class defines values for valid appliance types for the vCenter and Platform Services Controller node. See Info.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type ApplianceType string

const (
    // vCenter Server Appliance with an embedded Platform Services Controller.
     ApplianceType_VCSA_EMBEDDED ApplianceType = "VCSA_EMBEDDED"
    // vCenter Server Appliance with an external Platform Services Controller.
     ApplianceType_VCSA_EXTERNAL ApplianceType = "VCSA_EXTERNAL"
    // An external Platform Services Controller.
     ApplianceType_PSC_EXTERNAL ApplianceType = "PSC_EXTERNAL"
)

func (a ApplianceType) ApplianceType() bool {
    switch a {
        case ApplianceType_VCSA_EMBEDDED:
            return true
        case ApplianceType_VCSA_EXTERNAL:
            return true
        case ApplianceType_PSC_EXTERNAL:
            return true
        default:
            return false
    }
}





// The ``Info`` class contains vCenter or Platform Services Controller node details.
 type Info struct {
    Domain string
    Type_ ApplianceType
    ReplicationPartners []string
    ClientAffinity *string
}






// The ``Summary`` class contains commonly used information of vCenter or Platform Services Controller node.
 type Summary struct {
    Node string
    Type_ ApplianceType
    ReplicationPartners []string
    ClientAffinity *string
}






// The ``FilterSpec`` class contains property used to filter the results when listing vCenter and Platform Services Controller nodes (see Nodes#list).
 type FilterSpec struct {
    Types map[ApplianceType]bool
}









func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(FilterSpecBindingType))
    fieldNameMap["filter"] = "Filter"
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
       map[string]int{"Unauthenticated": 401,"Unauthorized": 403,"InvalidArgument": 400})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["node"] = bindings.NewIdType([]string {"com.vmware.vcenter.VCenter.name"}, "")
    fieldNameMap["node"] = "Node"
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
       map[string]int{"Unauthenticated": 401,"Unauthorized": 403,"NotFound": 404})
}



func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["domain"] = bindings.NewStringType()
    fieldNameMap["domain"] = "Domain"
    fields["type"] = bindings.NewEnumType("com.vmware.vcenter.topology.nodes.appliance_type", reflect.TypeOf(ApplianceType(ApplianceType_VCSA_EMBEDDED)))
    fieldNameMap["type"] = "Type_"
    fields["replication_partners"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewIdType([]string {"com.vmware.vcenter.VCenter.name"}, ""), reflect.TypeOf([]string{})))
    fieldNameMap["replication_partners"] = "ReplicationPartners"
    fields["client_affinity"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.vcenter.VCenter.name"}, ""))
    fieldNameMap["client_affinity"] = "ClientAffinity"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("type",
        map[string][]bindings.FieldData {
            "VCSA_EMBEDDED": []bindings.FieldData {
                 bindings.NewFieldData("replication_partners", true),
            },
            "PSC_EXTERNAL": []bindings.FieldData {
                 bindings.NewFieldData("replication_partners", true),
            },
            "VCSA_EXTERNAL": []bindings.FieldData {
                 bindings.NewFieldData("client_affinity", true),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.topology.nodes.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["node"] = bindings.NewIdType([]string {"com.vmware.vcenter.VCenter.name"}, "")
    fieldNameMap["node"] = "Node"
    fields["type"] = bindings.NewEnumType("com.vmware.vcenter.topology.nodes.appliance_type", reflect.TypeOf(ApplianceType(ApplianceType_VCSA_EMBEDDED)))
    fieldNameMap["type"] = "Type_"
    fields["replication_partners"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewIdType([]string {"com.vmware.vcenter.VCenter.name"}, ""), reflect.TypeOf([]string{})))
    fieldNameMap["replication_partners"] = "ReplicationPartners"
    fields["client_affinity"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.vcenter.VCenter.name"}, ""))
    fieldNameMap["client_affinity"] = "ClientAffinity"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("type",
        map[string][]bindings.FieldData {
            "VCSA_EMBEDDED": []bindings.FieldData {
                 bindings.NewFieldData("replication_partners", true),
            },
            "PSC_EXTERNAL": []bindings.FieldData {
                 bindings.NewFieldData("replication_partners", true),
            },
            "VCSA_EXTERNAL": []bindings.FieldData {
                 bindings.NewFieldData("client_affinity", true),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.topology.nodes.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}

func FilterSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["types"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewEnumType("com.vmware.vcenter.topology.nodes.appliance_type", reflect.TypeOf(ApplianceType(ApplianceType_VCSA_EMBEDDED))), reflect.TypeOf(map[ApplianceType]bool{})))
    fieldNameMap["types"] = "Types"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.topology.nodes.filter_spec",fields, reflect.TypeOf(FilterSpec{}), fieldNameMap, validators)
}


