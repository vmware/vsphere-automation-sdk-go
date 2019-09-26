/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Policies.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package policies

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// Resource type for vAPI metadata policy
const Policies_RESOURCE_TYPE = "com.vmware.vcenter.StoragePolicy"



// The ``FilterSpec`` class contains properties used to filter the results when listing the storage policies (see Policies#list)
 type FilterSpec struct {
    Policies map[string]bool
}






// The ``Summary`` class contains commonly used information about a storage policy.
 type Summary struct {
    Policy string
    Name string
    Description string
}






// The ``CompatibleDatastoreInfo`` class contains compatible datastore's information.
 type CompatibleDatastoreInfo struct {
    Datastore string
}






// The ``CompatibilityInfo`` class contains info about a list of datastores compatible with a specific storage policy.
 type CompatibilityInfo struct {
    CompatibleDatastores []CompatibleDatastoreInfo
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
       map[string]int{"InvalidArgument": 400,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403,"UnableToAllocateResource": 500})
}


func checkCompatibilityInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["policy"] = bindings.NewIdType([]string {"com.vmware.vcenter.StoragePolicy"}, "")
    fields["datastores"] = bindings.NewSetType(bindings.NewIdType([]string {"Datastore"}, ""), reflect.TypeOf(map[string]bool{}))
    fieldNameMap["policy"] = "Policy"
    fieldNameMap["datastores"] = "Datastores"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func checkCompatibilityOutputType() bindings.BindingType {
    return bindings.NewReferenceType(CompatibilityInfoBindingType)
}

func checkCompatibilityRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"NotFound": 404,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403,"UnableToAllocateResource": 500})
}



func FilterSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["policies"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"com.vmware.vcenter.StoragePolicy"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["policies"] = "Policies"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.storage.policies.filter_spec",fields, reflect.TypeOf(FilterSpec{}), fieldNameMap, validators)
}

func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["policy"] = bindings.NewIdType([]string {"com.vmware.vcenter.StoragePolicy"}, "")
    fieldNameMap["policy"] = "Policy"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["description"] = bindings.NewStringType()
    fieldNameMap["description"] = "Description"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.storage.policies.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}

func CompatibleDatastoreInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["datastore"] = bindings.NewIdType([]string {"Datastore"}, "")
    fieldNameMap["datastore"] = "Datastore"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.storage.policies.compatible_datastore_info",fields, reflect.TypeOf(CompatibleDatastoreInfo{}), fieldNameMap, validators)
}

func CompatibilityInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["compatible_datastores"] = bindings.NewListType(bindings.NewReferenceType(CompatibleDatastoreInfoBindingType), reflect.TypeOf([]CompatibleDatastoreInfo{}))
    fieldNameMap["compatible_datastores"] = "CompatibleDatastores"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.storage.policies.compatibility_info",fields, reflect.TypeOf(CompatibilityInfo{}), fieldNameMap, validators)
}


