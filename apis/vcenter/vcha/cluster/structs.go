/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Cluster.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package cluster

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vcenter/vcha"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``Type`` enumeration class defines the possible deployment types for a VCHA Cluster.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Type string

const (
    // Passive and witness nodes are cloned automatically.
     Type_AUTO Type = "AUTO"
    // Passive and witness nodes are not cloned automatically. After deployment, the customer should clone the passive and witness virtual machines.
     Type_MANUAL Type = "MANUAL"
)

func (t Type) Type() bool {
    switch t {
        case Type_AUTO:
            return true
        case Type_MANUAL:
            return true
        default:
            return false
    }
}




// The ``ClusterMode`` enumeration class defines the possible modes for a VCHA Cluster.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type ClusterMode string

const (
    // VCHA Cluster is enabled. State replication between the Active and Passive node is enabled and automatic failover is allowed.
     ClusterMode_ENABLED ClusterMode = "ENABLED"
    // VCHA Cluster is disabled. State replication between the Active and Passive node is disabled and automatic failover is not allowed.
     ClusterMode_DISABLED ClusterMode = "DISABLED"
    // VCHA Cluster is in maintenance mode. State replication between the Active and Passive node is enabled but automatic failover is not allowed.
     ClusterMode_MAINTENANCE ClusterMode = "MAINTENANCE"
)

func (c ClusterMode) ClusterMode() bool {
    switch c {
        case ClusterMode_ENABLED:
            return true
        case ClusterMode_DISABLED:
            return true
        case ClusterMode_MAINTENANCE:
            return true
        default:
            return false
    }
}




// The ``ClusterState`` enumeration class defines the possible for a VCHA Cluster.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type ClusterState string

const (
    // All three nodes in a VCHA Cluster are healthy and connected. State replication between Active and Passive node is working and both nodes are in sync.
     ClusterState_HEALTHY ClusterState = "HEALTHY"
    // A VCHA Cluster is said to be in a degraded state for either or all of the following reasons: 
    //
    // * There is a node loss.
    // * State replication between the Active and Passive node fails.
     ClusterState_DEGRADED ClusterState = "DEGRADED"
    // All three nodes are isolated from each other.
     ClusterState_ISOLATED ClusterState = "ISOLATED"
)

func (c ClusterState) ClusterState() bool {
    switch c {
        case ClusterState_HEALTHY:
            return true
        case ClusterState_DEGRADED:
            return true
        case ClusterState_ISOLATED:
            return true
        default:
            return false
    }
}




// The ``NodeState`` enumeration class defines possible state a node can be in a VCHA Cluster.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type NodeState string

const (
    // Node is up and has joined the VCHA Cluster.
     NodeState_UP NodeState = "UP"
    // Node is down and has left the VCHA Cluster.
     NodeState_DOWN NodeState = "DOWN"
)

func (n NodeState) NodeState() bool {
    switch n {
        case NodeState_UP:
            return true
        case NodeState_DOWN:
            return true
        default:
            return false
    }
}




// The ``NodeRole`` enumeration class defines the role node can be in a VCHA Cluster.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type NodeRole string

const (
    // Node is having a role of Active. In this role, node runs a vCenter Server that serves client requests.
     NodeRole_ACTIVE NodeRole = "ACTIVE"
    // Node is having a role of Passive. In this role node, runs as a standby for the Active vCenter Server and receives state updates. This node takes over the role of Active vCenter Server upon failover.
     NodeRole_PASSIVE NodeRole = "PASSIVE"
    // Node is having a role of Witness. In this role, node acts as a quorum node for avoiding the classic split-brain problem.
     NodeRole_WITNESS NodeRole = "WITNESS"
)

func (n NodeRole) NodeRole() bool {
    switch n {
        case NodeRole_ACTIVE:
            return true
        case NodeRole_PASSIVE:
            return true
        case NodeRole_WITNESS:
            return true
        default:
            return false
    }
}




// The ``ConfigState`` enumeration class defines the VCHA configuration state.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type ConfigState string

const (
    // VCHA cluster is configured.
     ConfigState_CONFIGURED ConfigState = "CONFIGURED"
    // VCHA cluster is not configured.
     ConfigState_NOTCONFIGURED ConfigState = "NOTCONFIGURED"
    // VCHA cluster is in an invalid/dirty state.
     ConfigState_INVALID ConfigState = "INVALID"
    // vCenter server appliance has been prepared for VCHA cluster configuration.
     ConfigState_PREPARED ConfigState = "PREPARED"
)

func (c ConfigState) ConfigState() bool {
    switch c {
        case ConfigState_CONFIGURED:
            return true
        case ConfigState_NOTCONFIGURED:
            return true
        case ConfigState_INVALID:
            return true
        case ConfigState_PREPARED:
            return true
        default:
            return false
    }
}




// The ``IpFamily`` enumeration class defines the IP address family.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type IpFamily string

const (
    // IPV4 address family.
     IpFamily_IPV4 IpFamily = "IPV4"
    // IPV6 address family.
     IpFamily_IPV6 IpFamily = "IPV6"
)

func (i IpFamily) IpFamily() bool {
    switch i {
        case IpFamily_IPV4:
            return true
        case IpFamily_IPV6:
            return true
        default:
            return false
    }
}





// The ``ActiveSpec`` class contains the deploy specification for the Active Node of the VCHA cluster.
 type ActiveSpec struct {
    HaNetworkType *vcha.NetworkType
    HaNetwork *string
    HaIp vcha.IpSpec
}






// The ``PassiveSpec`` class contains the deploy specification for the Passive Node of the VCHA cluster.
 type PassiveSpec struct {
    Placement *vcha.PlacementSpec
    HaIp vcha.IpSpec
    FailoverIp *vcha.IpSpec
}






// The ``WitnessSpec`` class contains the deploy specification for the Witness Node of the VCHA cluster.
 type WitnessSpec struct {
    Placement *vcha.PlacementSpec
    HaIp vcha.IpSpec
}






// The ``DeploySpec`` class contains the deploy specification for the three nodes of a VCHA cluster.
 type DeploySpec struct {
    VcSpec *vcha.CredentialsSpec
    Deployment Type
    Active ActiveSpec
    Passive PassiveSpec
    Witness WitnessSpec
}






// The ``NodeRuntimeInfo`` class describes a node's runtime information in a VCHA Cluster.
 type NodeRuntimeInfo struct {
    State *NodeState
    Role *NodeRole
    Placement *vcha.PlacementInfo
}






// The ``Ipv4Info`` class contains properties to describe IPV4 information of the configured network interface.
 type Ipv4Info struct {
    Address string
    SubnetMask string
    Prefix *int64
}






// The ``Ipv6Info`` class contains properties to describe IPV6 information of the configured network interface.
 type Ipv6Info struct {
    Address string
    Prefix int64
}






// The ``IpInfo`` class contains properties related to an ip.
 type IpInfo struct {
    IpFamily IpFamily
    Ipv4 *Ipv4Info
    Ipv6 *Ipv6Info
    GatewayIp *string
}






// The ``NodeInfo`` class defines the configuration information for the active and passive nodes in the cluster.
 type NodeInfo struct {
    FailoverIp *IpInfo
    HaIp IpInfo
    Runtime *NodeRuntimeInfo
}






// The ``WitnessInfo`` class defines the configuration and runtime information for the witness node in the cluster.
 type WitnessInfo struct {
    HaIp IpInfo
    Runtime *NodeRuntimeInfo
}






// The ``ErrorCondition`` class contains an error condition and a recommendation to handle the error condition.
 type ErrorCondition struct {
    Error std.LocalizableMessage
    Recommendation *std.LocalizableMessage
}






// The ``Info`` class contains the configuration and health information of the three nodes in a VCHA Cluster.
 type Info struct {
    ConfigState *ConfigState
    Node1 *NodeInfo
    Node2 *NodeInfo
    Witness *WitnessInfo
    Mode *ClusterMode
    HealthState *ClusterState
    HealthException []std.LocalizableMessage
    HealthWarnings []ErrorCondition
    ManualFailoverAllowed *bool
    AutoFailoverAllowed *bool
}






// The ``NodeVmInfo`` class contains information to describe the Virtual Machine of a node of a VCHA cluster.
 type NodeVmInfo struct {
    Vm string
    BiosUuid string
}






// The ``VmInfo`` class contains information to describe the Virtual Machines of passive and witness nodes of a VCHA cluster.
 type VmInfo struct {
    Passive NodeVmInfo
    Witness NodeVmInfo
}






// The ``UndeploySpec`` class contains the undeploy specification for a VCHA cluster.
 type UndeploySpec struct {
    VcSpec *vcha.CredentialsSpec
    ForceDelete *bool
    Vms *VmInfo
}









func deployInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(DeploySpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func deployOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func deployRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
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
       map[string]int{"InvalidArgument": 400,"Unauthorized": 403,"UnverifiedPeer": 400,"Error": 500})
}


func failoverInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["planned"] = bindings.NewBooleanType()
    fieldNameMap["planned"] = "Planned"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func failoverOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func failoverRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
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
       map[string]int{"Unauthorized": 403,"Error": 500})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vc_spec"] = bindings.NewOptionalType(bindings.NewReferenceType(vcha.CredentialsSpecBindingType))
    fields["partial"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["vc_spec"] = "VcSpec"
    fieldNameMap["partial"] = "Partial"
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
       map[string]int{"InvalidArgument": 400,"Unauthorized": 403,"UnverifiedPeer": 400,"Error": 500})
}


func undeployInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(UndeploySpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func undeployOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func undeployRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
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
       map[string]int{"InvalidArgument": 400,"NotFound": 404,"Unauthorized": 403,"UnverifiedPeer": 400,"Error": 500})
}



func ActiveSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["ha_network_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vcha.network_type", reflect.TypeOf(vcha.NetworkType(vcha.NetworkType_STANDARD_PORTGROUP))))
    fieldNameMap["ha_network_type"] = "HaNetworkType"
    fields["ha_network"] = bindings.NewOptionalType(bindings.NewIdType([]string {"Network:VCenter"}, ""))
    fieldNameMap["ha_network"] = "HaNetwork"
    fields["ha_ip"] = bindings.NewReferenceType(vcha.IpSpecBindingType)
    fieldNameMap["ha_ip"] = "HaIp"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vcha.cluster.active_spec",fields, reflect.TypeOf(ActiveSpec{}), fieldNameMap, validators)
}

func PassiveSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["placement"] = bindings.NewOptionalType(bindings.NewReferenceType(vcha.PlacementSpecBindingType))
    fieldNameMap["placement"] = "Placement"
    fields["ha_ip"] = bindings.NewReferenceType(vcha.IpSpecBindingType)
    fieldNameMap["ha_ip"] = "HaIp"
    fields["failover_ip"] = bindings.NewOptionalType(bindings.NewReferenceType(vcha.IpSpecBindingType))
    fieldNameMap["failover_ip"] = "FailoverIp"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vcha.cluster.passive_spec",fields, reflect.TypeOf(PassiveSpec{}), fieldNameMap, validators)
}

func WitnessSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["placement"] = bindings.NewOptionalType(bindings.NewReferenceType(vcha.PlacementSpecBindingType))
    fieldNameMap["placement"] = "Placement"
    fields["ha_ip"] = bindings.NewReferenceType(vcha.IpSpecBindingType)
    fieldNameMap["ha_ip"] = "HaIp"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vcha.cluster.witness_spec",fields, reflect.TypeOf(WitnessSpec{}), fieldNameMap, validators)
}

func DeploySpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vc_spec"] = bindings.NewOptionalType(bindings.NewReferenceType(vcha.CredentialsSpecBindingType))
    fieldNameMap["vc_spec"] = "VcSpec"
    fields["deployment"] = bindings.NewEnumType("com.vmware.vcenter.vcha.cluster.type", reflect.TypeOf(Type(Type_AUTO)))
    fieldNameMap["deployment"] = "Deployment"
    fields["active"] = bindings.NewReferenceType(ActiveSpecBindingType)
    fieldNameMap["active"] = "Active"
    fields["passive"] = bindings.NewReferenceType(PassiveSpecBindingType)
    fieldNameMap["passive"] = "Passive"
    fields["witness"] = bindings.NewReferenceType(WitnessSpecBindingType)
    fieldNameMap["witness"] = "Witness"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vcha.cluster.deploy_spec",fields, reflect.TypeOf(DeploySpec{}), fieldNameMap, validators)
}

func NodeRuntimeInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["state"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vcha.cluster.node_state", reflect.TypeOf(NodeState(NodeState_UP))))
    fieldNameMap["state"] = "State"
    fields["role"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vcha.cluster.node_role", reflect.TypeOf(NodeRole(NodeRole_ACTIVE))))
    fieldNameMap["role"] = "Role"
    fields["placement"] = bindings.NewOptionalType(bindings.NewReferenceType(vcha.PlacementInfoBindingType))
    fieldNameMap["placement"] = "Placement"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vcha.cluster.node_runtime_info",fields, reflect.TypeOf(NodeRuntimeInfo{}), fieldNameMap, validators)
}

func Ipv4InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["address"] = bindings.NewStringType()
    fieldNameMap["address"] = "Address"
    fields["subnet_mask"] = bindings.NewStringType()
    fieldNameMap["subnet_mask"] = "SubnetMask"
    fields["prefix"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["prefix"] = "Prefix"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vcha.cluster.ipv4_info",fields, reflect.TypeOf(Ipv4Info{}), fieldNameMap, validators)
}

func Ipv6InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["address"] = bindings.NewStringType()
    fieldNameMap["address"] = "Address"
    fields["prefix"] = bindings.NewIntegerType()
    fieldNameMap["prefix"] = "Prefix"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vcha.cluster.ipv6_info",fields, reflect.TypeOf(Ipv6Info{}), fieldNameMap, validators)
}

func IpInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["ip_family"] = bindings.NewEnumType("com.vmware.vcenter.vcha.cluster.ip_family", reflect.TypeOf(IpFamily(IpFamily_IPV4)))
    fieldNameMap["ip_family"] = "IpFamily"
    fields["ipv4"] = bindings.NewOptionalType(bindings.NewReferenceType(Ipv4InfoBindingType))
    fieldNameMap["ipv4"] = "Ipv4"
    fields["ipv6"] = bindings.NewOptionalType(bindings.NewReferenceType(Ipv6InfoBindingType))
    fieldNameMap["ipv6"] = "Ipv6"
    fields["gateway_ip"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["gateway_ip"] = "GatewayIp"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("ip_family",
        map[string][]bindings.FieldData {
            "IPV4": []bindings.FieldData {
                 bindings.NewFieldData("ipv4", true),
            },
            "IPV6": []bindings.FieldData {
                 bindings.NewFieldData("ipv6", true),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.vcha.cluster.ip_info",fields, reflect.TypeOf(IpInfo{}), fieldNameMap, validators)
}

func NodeInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["failover_ip"] = bindings.NewOptionalType(bindings.NewReferenceType(IpInfoBindingType))
    fieldNameMap["failover_ip"] = "FailoverIp"
    fields["ha_ip"] = bindings.NewReferenceType(IpInfoBindingType)
    fieldNameMap["ha_ip"] = "HaIp"
    fields["runtime"] = bindings.NewOptionalType(bindings.NewReferenceType(NodeRuntimeInfoBindingType))
    fieldNameMap["runtime"] = "Runtime"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vcha.cluster.node_info",fields, reflect.TypeOf(NodeInfo{}), fieldNameMap, validators)
}

func WitnessInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["ha_ip"] = bindings.NewReferenceType(IpInfoBindingType)
    fieldNameMap["ha_ip"] = "HaIp"
    fields["runtime"] = bindings.NewOptionalType(bindings.NewReferenceType(NodeRuntimeInfoBindingType))
    fieldNameMap["runtime"] = "Runtime"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vcha.cluster.witness_info",fields, reflect.TypeOf(WitnessInfo{}), fieldNameMap, validators)
}

func ErrorConditionBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["error"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
    fieldNameMap["error"] = "Error"
    fields["recommendation"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
    fieldNameMap["recommendation"] = "Recommendation"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vcha.cluster.error_condition",fields, reflect.TypeOf(ErrorCondition{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["config_state"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vcha.cluster.config_state", reflect.TypeOf(ConfigState(ConfigState_CONFIGURED))))
    fieldNameMap["config_state"] = "ConfigState"
    fields["node1"] = bindings.NewOptionalType(bindings.NewReferenceType(NodeInfoBindingType))
    fieldNameMap["node1"] = "Node1"
    fields["node2"] = bindings.NewOptionalType(bindings.NewReferenceType(NodeInfoBindingType))
    fieldNameMap["node2"] = "Node2"
    fields["witness"] = bindings.NewOptionalType(bindings.NewReferenceType(WitnessInfoBindingType))
    fieldNameMap["witness"] = "Witness"
    fields["mode"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vcha.cluster.cluster_mode", reflect.TypeOf(ClusterMode(ClusterMode_ENABLED))))
    fieldNameMap["mode"] = "Mode"
    fields["health_state"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vcha.cluster.cluster_state", reflect.TypeOf(ClusterState(ClusterState_HEALTHY))))
    fieldNameMap["health_state"] = "HealthState"
    fields["health_exception"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{})))
    fieldNameMap["health_exception"] = "HealthException"
    fields["health_warnings"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ErrorConditionBindingType), reflect.TypeOf([]ErrorCondition{})))
    fieldNameMap["health_warnings"] = "HealthWarnings"
    fields["manual_failover_allowed"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["manual_failover_allowed"] = "ManualFailoverAllowed"
    fields["auto_failover_allowed"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["auto_failover_allowed"] = "AutoFailoverAllowed"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vcha.cluster.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func NodeVmInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine:VCenter"}, "")
    fieldNameMap["vm"] = "Vm"
    fields["bios_uuid"] = bindings.NewStringType()
    fieldNameMap["bios_uuid"] = "BiosUuid"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vcha.cluster.node_vm_info",fields, reflect.TypeOf(NodeVmInfo{}), fieldNameMap, validators)
}

func VmInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["passive"] = bindings.NewReferenceType(NodeVmInfoBindingType)
    fieldNameMap["passive"] = "Passive"
    fields["witness"] = bindings.NewReferenceType(NodeVmInfoBindingType)
    fieldNameMap["witness"] = "Witness"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vcha.cluster.vm_info",fields, reflect.TypeOf(VmInfo{}), fieldNameMap, validators)
}

func UndeploySpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vc_spec"] = bindings.NewOptionalType(bindings.NewReferenceType(vcha.CredentialsSpecBindingType))
    fieldNameMap["vc_spec"] = "VcSpec"
    fields["force_delete"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["force_delete"] = "ForceDelete"
    fields["vms"] = bindings.NewOptionalType(bindings.NewReferenceType(VmInfoBindingType))
    fieldNameMap["vms"] = "Vms"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vcha.cluster.undeploy_spec",fields, reflect.TypeOf(UndeploySpec{}), fieldNameMap, validators)
}


