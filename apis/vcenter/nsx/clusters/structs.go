/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Clusters.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package clusters

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vcenter/nsx"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)




// The ``NetworkSpec`` contains information related to network configuration for the Tunnel Endpoints (TEPs) on the hosts in the VC cluster. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type NetworkSpec struct {
    Vlan int64
    Mode nsx.IpAllocationMode
    CreateIpPool *bool
    IpPoolSpec *nsx.IpPoolCreateSpec
    IpPool *string
}






// The ``EnableSpec`` class contains the specification to configure NSX networking on a vSphere cluster. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type EnableSpec struct {
    DvSwitch string
    OverlayNetwork NetworkSpec
}









func enableInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fields["spec"] = bindings.NewReferenceType(EnableSpecBindingType)
    fieldNameMap["cluster"] = "Cluster"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func enableOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func enableRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    paramsTypeMap["spec"] = bindings.NewReferenceType(EnableSpecBindingType)
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    pathParams["cluster"] = "cluster"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "spec",
      "POST",
      "/vcenter/nsx/clusters/{cluster}",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"NotFound": 404,"AlreadyExists": 400,"InvalidArgument": 400,"Error": 500})
}



func NetworkSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vlan"] = bindings.NewIntegerType()
    fieldNameMap["vlan"] = "Vlan"
    fields["mode"] = bindings.NewEnumType("com.vmware.vcenter.nsx.ip_allocation_mode", reflect.TypeOf(nsx.IpAllocationMode(nsx.IpAllocationMode_DHCP)))
    fieldNameMap["mode"] = "Mode"
    fields["create_ip_pool"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["create_ip_pool"] = "CreateIpPool"
    fields["ip_pool_spec"] = bindings.NewOptionalType(bindings.NewReferenceType(nsx.IpPoolCreateSpecBindingType))
    fieldNameMap["ip_pool_spec"] = "IpPoolSpec"
    fields["ip_pool"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.nsx.pools.ip_pool"}, ""))
    fieldNameMap["ip_pool"] = "IpPool"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("mode",
        map[string][]bindings.FieldData {
            "IP_POOL": []bindings.FieldData {
                 bindings.NewFieldData("create_ip_pool", true),
                 bindings.NewFieldData("ip_pool_spec", false),
                 bindings.NewFieldData("ip_pool", false),
            },
            "DHCP": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.nsx.clusters.network_spec",fields, reflect.TypeOf(NetworkSpec{}), fieldNameMap, validators)
}

func EnableSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["dv_switch"] = bindings.NewIdType([]string {"DistributedVirtualSwitch"}, "")
    fieldNameMap["dv_switch"] = "DvSwitch"
    fields["overlay_network"] = bindings.NewReferenceType(NetworkSpecBindingType)
    fieldNameMap["overlay_network"] = "OverlayNetwork"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.clusters.enable_spec",fields, reflect.TypeOf(EnableSpec{}), fieldNameMap, validators)
}


