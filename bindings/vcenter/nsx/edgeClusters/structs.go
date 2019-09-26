/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: EdgeClusters.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package edgeClusters

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/vcenter/nsx"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)




// The ``ManagementNetworkSpec`` contains information related to network configuration for the management network interface on the NSX Edge node. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ManagementNetworkSpec struct {
    Network string
    SubnetMask string
    Gateway string
}






// The ``OverlayNetworkSpec`` contains information related to network configuration for the overlay network interface on the NSX Edge node. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type OverlayNetworkSpec struct {
    Vlan int64
    CreateIpPool bool
    IpPoolSpec *nsx.IpPoolCreateSpec
    IpPool *string
}






// The ``UplinkNetworkSpec`` contains information related to configuration of the uplink network interface on the NSX Edge node(s). **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type UplinkNetworkSpec struct {
    Vlan int64
    SubnetMask string
    Gateway string
}






// The ``PlacementSpec`` contains the placement specification required to set up NSX Edge node(s). **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type PlacementSpec struct {
    CreateResourcePool bool
    ResourcePool *string
    ResourcePoolName *string
}






// The ``NodeSpec`` class contains the specification required to set up NSX Edge node. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type NodeSpec struct {
    Name string
    Password string
    Hostname string
    ManagementIpAddress string
    Datastore string
}






// The ``UplinkNodeSpec`` class contains the specification required to set up an NSX Edge node configured with the uplink network. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type UplinkNodeSpec struct {
    UplinkIpAddress string
    Name string
    Password string
    Hostname string
    ManagementIpAddress string
    Datastore string
}






// The ``EnableSpec`` class contains the specification required to set up NSX Edge node(s). **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type EnableSpec struct {
    EdgeNodes []UplinkNodeSpec
    Placement PlacementSpec
    OverlayNetwork OverlayNetworkSpec
    ManagementNetwork ManagementNetworkSpec
    UplinkNetwork UplinkNetworkSpec
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
      "/vcenter/nsx/edge-clusters/{cluster}",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"AlreadyExists": 400,"InvalidArgument": 400,"NotAllowedInCurrentState": 400,"NotFound": 404,"Error": 500})
}



func ManagementNetworkSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["network"] = bindings.NewIdType([]string {"Network"}, "")
    fieldNameMap["network"] = "Network"
    fields["subnet_mask"] = bindings.NewStringType()
    fieldNameMap["subnet_mask"] = "SubnetMask"
    fields["gateway"] = bindings.NewStringType()
    fieldNameMap["gateway"] = "Gateway"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.edge_clusters.management_network_spec",fields, reflect.TypeOf(ManagementNetworkSpec{}), fieldNameMap, validators)
}

func OverlayNetworkSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vlan"] = bindings.NewIntegerType()
    fieldNameMap["vlan"] = "Vlan"
    fields["create_ip_pool"] = bindings.NewBooleanType()
    fieldNameMap["create_ip_pool"] = "CreateIpPool"
    fields["ip_pool_spec"] = bindings.NewOptionalType(bindings.NewReferenceType(nsx.IpPoolCreateSpecBindingType))
    fieldNameMap["ip_pool_spec"] = "IpPoolSpec"
    fields["ip_pool"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.nsx.pools.ip_pool"}, ""))
    fieldNameMap["ip_pool"] = "IpPool"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.edge_clusters.overlay_network_spec",fields, reflect.TypeOf(OverlayNetworkSpec{}), fieldNameMap, validators)
}

func UplinkNetworkSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vlan"] = bindings.NewIntegerType()
    fieldNameMap["vlan"] = "Vlan"
    fields["subnet_mask"] = bindings.NewStringType()
    fieldNameMap["subnet_mask"] = "SubnetMask"
    fields["gateway"] = bindings.NewStringType()
    fieldNameMap["gateway"] = "Gateway"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.edge_clusters.uplink_network_spec",fields, reflect.TypeOf(UplinkNetworkSpec{}), fieldNameMap, validators)
}

func PlacementSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["create_resource_pool"] = bindings.NewBooleanType()
    fieldNameMap["create_resource_pool"] = "CreateResourcePool"
    fields["resource_pool"] = bindings.NewOptionalType(bindings.NewIdType([]string {"ResourcePool"}, ""))
    fieldNameMap["resource_pool"] = "ResourcePool"
    fields["resource_pool_name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["resource_pool_name"] = "ResourcePoolName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.edge_clusters.placement_spec",fields, reflect.TypeOf(PlacementSpec{}), fieldNameMap, validators)
}

func NodeSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["password"] = bindings.NewSecretType()
    fieldNameMap["password"] = "Password"
    fields["hostname"] = bindings.NewStringType()
    fieldNameMap["hostname"] = "Hostname"
    fields["management_ip_address"] = bindings.NewStringType()
    fieldNameMap["management_ip_address"] = "ManagementIpAddress"
    fields["datastore"] = bindings.NewIdType([]string {"Datastore"}, "")
    fieldNameMap["datastore"] = "Datastore"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.edge_clusters.node_spec",fields, reflect.TypeOf(NodeSpec{}), fieldNameMap, validators)
}

func UplinkNodeSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["uplink_ip_address"] = bindings.NewStringType()
    fieldNameMap["uplink_ip_address"] = "UplinkIpAddress"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["password"] = bindings.NewSecretType()
    fieldNameMap["password"] = "Password"
    fields["hostname"] = bindings.NewStringType()
    fieldNameMap["hostname"] = "Hostname"
    fields["management_ip_address"] = bindings.NewStringType()
    fieldNameMap["management_ip_address"] = "ManagementIpAddress"
    fields["datastore"] = bindings.NewIdType([]string {"Datastore"}, "")
    fieldNameMap["datastore"] = "Datastore"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.edge_clusters.uplink_node_spec",fields, reflect.TypeOf(UplinkNodeSpec{}), fieldNameMap, validators)
}

func EnableSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["edge_nodes"] = bindings.NewListType(bindings.NewReferenceType(UplinkNodeSpecBindingType), reflect.TypeOf([]UplinkNodeSpec{}))
    fieldNameMap["edge_nodes"] = "EdgeNodes"
    fields["placement"] = bindings.NewReferenceType(PlacementSpecBindingType)
    fieldNameMap["placement"] = "Placement"
    fields["overlay_network"] = bindings.NewReferenceType(OverlayNetworkSpecBindingType)
    fieldNameMap["overlay_network"] = "OverlayNetwork"
    fields["management_network"] = bindings.NewReferenceType(ManagementNetworkSpecBindingType)
    fieldNameMap["management_network"] = "ManagementNetwork"
    fields["uplink_network"] = bindings.NewReferenceType(UplinkNetworkSpecBindingType)
    fieldNameMap["uplink_network"] = "UplinkNetwork"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.edge_clusters.enable_spec",fields, reflect.TypeOf(EnableSpec{}), fieldNameMap, validators)
}


