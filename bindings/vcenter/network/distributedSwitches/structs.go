/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: DistributedSwitches.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package distributedSwitches

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// The resource type for the vCenter Distributed Switch. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
const DistributedSwitches_RESOURCE_TYPE = "DistributedVirtualSwitch"



// The ``FilterSpec`` class contains properties used to filter the results when listing distributed switches (see DistributedSwitches#list). A flag is also provided to specify that only distributed switches connected to ALL specified hosts in a cluster should satisfy the filter. If multiple properties are specified, only distributed switches matching all the properties match the filter. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type FilterSpec struct {
    MinVersion *string
    MinMtu *int64
    Clusters map[string]bool
    ConnectedToAllHosts *bool
}






// The ``Summary`` class contains commonly used information about a VDS. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Summary struct {
    DistributedSwitch string
    Name string
    Uuid string
    Version string
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
       map[string]int{"InvalidArgument": 400,"UnableToAllocateResource": 500,"Unauthenticated": 401,"Unauthorized": 403})
}



func FilterSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["min_version"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["min_version"] = "MinVersion"
    fields["min_mtu"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["min_mtu"] = "MinMtu"
    fields["clusters"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"ClusterComputeResource"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["clusters"] = "Clusters"
    fields["connected_to_all_hosts"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["connected_to_all_hosts"] = "ConnectedToAllHosts"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.network.distributed_switches.filter_spec",fields, reflect.TypeOf(FilterSpec{}), fieldNameMap, validators)
}

func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["distributed_switch"] = bindings.NewIdType([]string {"DistributedVirtualSwitch"}, "")
    fieldNameMap["distributed_switch"] = "DistributedSwitch"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["uuid"] = bindings.NewStringType()
    fieldNameMap["uuid"] = "Uuid"
    fields["version"] = bindings.NewStringType()
    fieldNameMap["version"] = "Version"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.network.distributed_switches.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}


