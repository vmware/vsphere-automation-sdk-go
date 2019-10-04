/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: ManagementNetworkCompatibility.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package managementNetworkCompatibility

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)




// The ``HostVmknicInfo`` class contains information about the network configuration of a VMKernel adapter of a host. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type HostVmknicInfo struct {
    Host string
    Vmknic string
    SubnetMask string
    Gateway string
    IncompatibleReasons []std.LocalizableMessage
    BestPracticeViolations []std.LocalizableMessage
}






// The ``Summary`` class contains information about the compatibility of a Distributed Virtual Portgroup (DVPG) and the network configuration of associated vmknics for setting up management networks. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Summary struct {
    DistributedPortgroup string
    Name string
    HostVmknics []HostVmknicInfo
}









func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fields["distributed_switch"] = bindings.NewIdType([]string {"DistributedVirtualSwitch"}, "")
    fieldNameMap["cluster"] = "Cluster"
    fieldNameMap["distributed_switch"] = "DistributedSwitch"
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
    paramsTypeMap["distributed_switch"] = bindings.NewIdType([]string {"DistributedVirtualSwitch"}, "")
    queryParams["cluster"] = "cluster"
    queryParams["distributed_switch"] = "distributed_switch"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "GET",
      "/vcenter/nsx/management-network-compatibility",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"NotFound": 404,"Error": 500})
}



func HostVmknicInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["host"] = bindings.NewIdType([]string {"HostSystem"}, "")
    fieldNameMap["host"] = "Host"
    fields["vmknic"] = bindings.NewStringType()
    fieldNameMap["vmknic"] = "Vmknic"
    fields["subnet_mask"] = bindings.NewStringType()
    fieldNameMap["subnet_mask"] = "SubnetMask"
    fields["gateway"] = bindings.NewStringType()
    fieldNameMap["gateway"] = "Gateway"
    fields["incompatible_reasons"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["incompatible_reasons"] = "IncompatibleReasons"
    fields["best_practice_violations"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["best_practice_violations"] = "BestPracticeViolations"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.management_network_compatibility.host_vmknic_info",fields, reflect.TypeOf(HostVmknicInfo{}), fieldNameMap, validators)
}

func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["distributed_portgroup"] = bindings.NewIdType([]string {"DistributedVirtualPortgroup"}, "")
    fieldNameMap["distributed_portgroup"] = "DistributedPortgroup"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["host_vmknics"] = bindings.NewListType(bindings.NewReferenceType(HostVmknicInfoBindingType), reflect.TypeOf([]HostVmknicInfo{}))
    fieldNameMap["host_vmknics"] = "HostVmknics"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.management_network_compatibility.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}


