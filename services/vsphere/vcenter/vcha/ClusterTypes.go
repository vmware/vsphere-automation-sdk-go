/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Cluster.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package vcha

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``Type`` enumeration class defines the possible deployment types for a VCHA Cluster. This enumeration was added in vSphere API 6.7.1.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type ClusterType string

const (
    // Passive and witness nodes are cloned automatically. This constant field was added in vSphere API 6.7.1.
	ClusterType_AUTO ClusterType = "AUTO"
    // Passive and witness nodes are not cloned automatically. After deployment, the customer should clone the passive and witness virtual machines. This constant field was added in vSphere API 6.7.1.
	ClusterType_MANUAL ClusterType = "MANUAL"
)

func (t ClusterType) ClusterType() bool {
	switch t {
	case ClusterType_AUTO:
		return true
	case ClusterType_MANUAL:
		return true
	default:
		return false
	}
}


// The ``ClusterMode`` enumeration class defines the possible modes for a VCHA Cluster. This enumeration was added in vSphere API 6.7.1.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type ClusterClusterMode string

const (
    // VCHA Cluster is enabled. State replication between the Active and Passive node is enabled and automatic failover is allowed. This constant field was added in vSphere API 6.7.1.
	ClusterClusterMode_ENABLED ClusterClusterMode = "ENABLED"
    // VCHA Cluster is disabled. State replication between the Active and Passive node is disabled and automatic failover is not allowed. This constant field was added in vSphere API 6.7.1.
	ClusterClusterMode_DISABLED ClusterClusterMode = "DISABLED"
    // VCHA Cluster is in maintenance mode. State replication between the Active and Passive node is enabled but automatic failover is not allowed. This constant field was added in vSphere API 6.7.1.
	ClusterClusterMode_MAINTENANCE ClusterClusterMode = "MAINTENANCE"
)

func (c ClusterClusterMode) ClusterClusterMode() bool {
	switch c {
	case ClusterClusterMode_ENABLED:
		return true
	case ClusterClusterMode_DISABLED:
		return true
	case ClusterClusterMode_MAINTENANCE:
		return true
	default:
		return false
	}
}


// The ``ClusterState`` enumeration class defines the possible for a VCHA Cluster. This enumeration was added in vSphere API 6.7.1.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type ClusterClusterState string

const (
    // All three nodes in a VCHA Cluster are healthy and connected. State replication between Active and Passive node is working and both nodes are in sync. This constant field was added in vSphere API 6.7.1.
	ClusterClusterState_HEALTHY ClusterClusterState = "HEALTHY"
    // A VCHA Cluster is said to be in a degraded state for either or all of the following reasons: 
    //
    // * There is a node loss.
    // * State replication between the Active and Passive node fails.
    //
    // . This constant field was added in vSphere API 6.7.1.
	ClusterClusterState_DEGRADED ClusterClusterState = "DEGRADED"
    // All three nodes are isolated from each other. This constant field was added in vSphere API 6.7.1.
	ClusterClusterState_ISOLATED ClusterClusterState = "ISOLATED"
)

func (c ClusterClusterState) ClusterClusterState() bool {
	switch c {
	case ClusterClusterState_HEALTHY:
		return true
	case ClusterClusterState_DEGRADED:
		return true
	case ClusterClusterState_ISOLATED:
		return true
	default:
		return false
	}
}


// The ``NodeState`` enumeration class defines possible state a node can be in a VCHA Cluster. This enumeration was added in vSphere API 6.7.1.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type ClusterNodeState string

const (
    // Node is up and has joined the VCHA Cluster. This constant field was added in vSphere API 6.7.1.
	ClusterNodeState_UP ClusterNodeState = "UP"
    // Node is down and has left the VCHA Cluster. This constant field was added in vSphere API 6.7.1.
	ClusterNodeState_DOWN ClusterNodeState = "DOWN"
)

func (n ClusterNodeState) ClusterNodeState() bool {
	switch n {
	case ClusterNodeState_UP:
		return true
	case ClusterNodeState_DOWN:
		return true
	default:
		return false
	}
}


// The ``NodeRole`` enumeration class defines the role node can be in a VCHA Cluster. This enumeration was added in vSphere API 6.7.1.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type ClusterNodeRole string

const (
    // Node is having a role of Active. In this role, node runs a vCenter Server that serves client requests. This constant field was added in vSphere API 6.7.1.
	ClusterNodeRole_ACTIVE ClusterNodeRole = "ACTIVE"
    // Node is having a role of Passive. In this role node, runs as a standby for the Active vCenter Server and receives state updates. This node takes over the role of Active vCenter Server upon failover. This constant field was added in vSphere API 6.7.1.
	ClusterNodeRole_PASSIVE ClusterNodeRole = "PASSIVE"
    // Node is having a role of Witness. In this role, node acts as a quorum node for avoiding the classic split-brain problem. This constant field was added in vSphere API 6.7.1.
	ClusterNodeRole_WITNESS ClusterNodeRole = "WITNESS"
)

func (n ClusterNodeRole) ClusterNodeRole() bool {
	switch n {
	case ClusterNodeRole_ACTIVE:
		return true
	case ClusterNodeRole_PASSIVE:
		return true
	case ClusterNodeRole_WITNESS:
		return true
	default:
		return false
	}
}


// The ``ConfigState`` enumeration class defines the VCHA configuration state. This enumeration was added in vSphere API 6.7.1.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type ClusterConfigState string

const (
    // VCHA cluster is configured. This constant field was added in vSphere API 6.7.1.
	ClusterConfigState_CONFIGURED ClusterConfigState = "CONFIGURED"
    // VCHA cluster is not configured. This constant field was added in vSphere API 6.7.1.
	ClusterConfigState_NOTCONFIGURED ClusterConfigState = "NOTCONFIGURED"
    // VCHA cluster is in an invalid/dirty state. This constant field was added in vSphere API 6.7.1.
	ClusterConfigState_INVALID ClusterConfigState = "INVALID"
    // vCenter server appliance has been prepared for VCHA cluster configuration. This constant field was added in vSphere API 6.7.1.
	ClusterConfigState_PREPARED ClusterConfigState = "PREPARED"
)

func (c ClusterConfigState) ClusterConfigState() bool {
	switch c {
	case ClusterConfigState_CONFIGURED:
		return true
	case ClusterConfigState_NOTCONFIGURED:
		return true
	case ClusterConfigState_INVALID:
		return true
	case ClusterConfigState_PREPARED:
		return true
	default:
		return false
	}
}


// The ``IpFamily`` enumeration class defines the IP address family. This enumeration was added in vSphere API 6.7.1.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type ClusterIpFamily string

const (
    // IPV4 address family. This constant field was added in vSphere API 6.7.1.
	ClusterIpFamily_IPV4 ClusterIpFamily = "IPV4"
    // IPV6 address family. This constant field was added in vSphere API 6.7.1.
	ClusterIpFamily_IPV6 ClusterIpFamily = "IPV6"
)

func (i ClusterIpFamily) ClusterIpFamily() bool {
	switch i {
	case ClusterIpFamily_IPV4:
		return true
	case ClusterIpFamily_IPV6:
		return true
	default:
		return false
	}
}


// The ``ActiveSpec`` class contains the deploy specification for the Active Node of the VCHA cluster. This class was added in vSphere API 6.7.1.
type ClusterActiveSpec struct {
    // The type of the Network object used by the HA network.
    //  If the ClusterActiveSpec#haNetwork property is set, then the ClusterActiveSpec#haNetworkType field must be set.
    //  If the ClusterActiveSpec#haNetwork property is null, then the ClusterActiveSpec#haNetworkType property is ignored. This property was added in vSphere API 6.7.1.
	HaNetworkType *NetworkType
    // The identifier of the Network object used for the HA network.
    //  If the ClusterActiveSpec#haNetwork property is set, then the ClusterActiveSpec#haNetworkType property must be set.
    //  If the ClusterActiveSpec#haNetwork property is null, then the ClusterActiveSpec#haNetworkType property is ignored. This property was added in vSphere API 6.7.1.
	HaNetwork *string
    // IP specification for the HA network. This property was added in vSphere API 6.7.1.
	HaIp IpSpec
}

// The ``PassiveSpec`` class contains the deploy specification for the Passive Node of the VCHA cluster. This class was added in vSphere API 6.7.1.
type ClusterPassiveSpec struct {
    // Contains the placement configuration of the node. This property was added in vSphere API 6.7.1.
	Placement *PlacementSpec
    // IP specification for the HA network. This property was added in vSphere API 6.7.1.
	HaIp IpSpec
    // IP specification for the management network. This property was added in vSphere API 6.7.1.
	FailoverIp *IpSpec
}

// The ``WitnessSpec`` class contains the deploy specification for the Witness Node of the VCHA cluster. This class was added in vSphere API 6.7.1.
type ClusterWitnessSpec struct {
    // Contains the placement configuration of the node. This property was added in vSphere API 6.7.1.
	Placement *PlacementSpec
    // IP specification for the HA network. This property was added in vSphere API 6.7.1.
	HaIp IpSpec
}

// The ``DeploySpec`` class contains the deploy specification for the three nodes of a VCHA cluster. This class was added in vSphere API 6.7.1.
type ClusterDeploySpec struct {
    // Contains the active node's management vCenter server credentials. This property was added in vSphere API 6.7.1.
	VcSpec *CredentialsSpec
    // Contains the deployment type. This property was added in vSphere API 6.7.1.
	Deployment ClusterType
    // Contains the active node's network configuration. This property was added in vSphere API 6.7.1.
	Active ClusterActiveSpec
    // Contains the passive node's placement configuration. This property was added in vSphere API 6.7.1.
	Passive ClusterPassiveSpec
    // Contains the witness node's placement configuration. This property was added in vSphere API 6.7.1.
	Witness ClusterWitnessSpec
}

// The ``NodeRuntimeInfo`` class describes a node's runtime information in a VCHA Cluster. This class was added in vSphere API 6.7.1.
type ClusterNodeRuntimeInfo struct {
    // Last known state of the node.
    //  The active node's management vCenter server credentials are not required to populate ClusterNodeRuntimeInfo#state. This property was added in vSphere API 6.7.1.
	State *ClusterNodeState
    // Last known role of the node.
    //  The active node's management vCenter server credentials are not required to populate ClusterNodeRuntimeInfo#role. This property was added in vSphere API 6.7.1.
	Role *ClusterNodeRole
    // Placement information of the node.
    //  The active node's management vCenter server credentials are required to populate most properties of ClusterNodeRuntimeInfo#placement. This property was added in vSphere API 6.7.1.
	Placement *PlacementInfo
}

// The ``Ipv4Info`` class contains properties to describe IPV4 information of the configured network interface. This class was added in vSphere API 6.7.1.
type ClusterIpv4Info struct {
    // IP address of the configured network interface. This property was added in vSphere API 6.7.1.
	Address string
    // The subnet mask of the interface. This property was added in vSphere API 6.7.1.
	SubnetMask string
    // The CIDR prefix of the interface. This property was added in vSphere API 6.7.1.
	Prefix *int64
}

// The ``Ipv6Info`` class contains properties to describe IPV6 information of the configured network interface. This class was added in vSphere API 6.7.1.
type ClusterIpv6Info struct {
    // IP address of the configured network interface. This property was added in vSphere API 6.7.1.
	Address string
    // The CIDR prefix of the interface. This property was added in vSphere API 6.7.1.
	Prefix int64
}

// The ``IpInfo`` class contains properties related to an ip. This class was added in vSphere API 6.7.1.
type ClusterIpInfo struct {
    // Family of the ip. This property was added in vSphere API 6.7.1.
	IpFamily ClusterIpFamily
    // If the type of the ip family is IPV4, then this will point to IPv4 address specification. This property was added in vSphere API 6.7.1.
	Ipv4 *ClusterIpv4Info
    // If the type of the ip family is IPV6, then this will point to IPv6 address specification. This property was added in vSphere API 6.7.1.
	Ipv6 *ClusterIpv6Info
    // Gateway IP address. This property was added in vSphere API 6.7.1.
	GatewayIp *string
}

// The ``NodeInfo`` class defines the configuration information for the active and passive nodes in the cluster. This class was added in vSphere API 6.7.1.
type ClusterNodeInfo struct {
    // Failover IP address that this node will assume after the failover to serve client requests. Each failover node can have a different failover IP address.
    //  The active node's management vCenter server credentials are not required to populate ClusterNodeInfo#failoverIp. This property was added in vSphere API 6.7.1.
	FailoverIp *ClusterIpInfo
    // VCHA Cluster network configuration of the node. All cluster communication (state replication, heartbeat, cluster messages) happens over this network.
    //  The active node's management vCenter server credentials are not required to populate this ClusterNodeInfo#haIp. This property was added in vSphere API 6.7.1.
	HaIp ClusterIpInfo
    // Runtime information for the node in the VCHA Cluster.
    //  The active node's management vCenter server credentials are required to populate some properties of ClusterNodeInfo#runtime. This property was added in vSphere API 6.7.1.
	Runtime *ClusterNodeRuntimeInfo
}

// The ``WitnessInfo`` class defines the configuration and runtime information for the witness node in the cluster. This class was added in vSphere API 6.7.1.
type ClusterWitnessInfo struct {
    // VCHA Cluster network configuration of the node. All cluster communication (state replication, heartbeat, cluster messages) happens over this network.
    //  The active node's management vCenter server credentials are not required to populate ClusterWitnessInfo#haIp. This property was added in vSphere API 6.7.1.
	HaIp ClusterIpInfo
    // Runtime information for the node in the VCHA Cluster.
    //  The active node's management vCenter server credentials are required to populate some properties of ClusterWitnessInfo#runtime. This property was added in vSphere API 6.7.1.
	Runtime *ClusterNodeRuntimeInfo
}

// The ``ErrorCondition`` class contains an error condition and a recommendation to handle the error condition. This class was added in vSphere API 6.7.1.
type ClusterErrorCondition struct {
    // Contains an error condition. This property was added in vSphere API 6.7.1.
	Error_ std.LocalizableMessage
    // Contains a recommendation on handling the error condition. This property was added in vSphere API 6.7.1.
	Recommendation *std.LocalizableMessage
}

// The ``Info`` class contains the configuration and health information of the three nodes in a VCHA Cluster. This class was added in vSphere API 6.7.1.
type ClusterInfo struct {
    // Configuration state of the VCHA cluster.
    //  The active node's management vCenter server credentials are not required to populate this property. This property was added in vSphere API 6.7.1.
	ConfigState *ClusterConfigState
    // Node configuration information for the VCHA cluster. This property was added in vSphere API 6.7.1.
	Node1 *ClusterNodeInfo
    // Node configuration information for the VCHA cluster. This property was added in vSphere API 6.7.1.
	Node2 *ClusterNodeInfo
    // Node configuration information for the VCHA cluster. This property was added in vSphere API 6.7.1.
	Witness *ClusterWitnessInfo
    // Operational mode of the VCHA Cluster. This property was added in vSphere API 6.7.1.
	Mode *ClusterClusterMode
    // Last known state of the VCHA Cluster. This property was added in vSphere API 6.7.1.
	HealthState *ClusterClusterState
    // Health warning messages if the health information is unavailable. This property was added in vSphere API 6.7.1.
	HealthException []std.LocalizableMessage
    // A collection of messages describing the reason for a non-healthy Cluster. This property was added in vSphere API 6.7.1.
	HealthWarnings []ClusterErrorCondition
    // Specifies if manual failover is allowed. This property was added in vSphere API 6.7.1.
	ManualFailoverAllowed *bool
    // Specifies if automatic failover is allowed. This property was added in vSphere API 6.7.1.
	AutoFailoverAllowed *bool
}

// The ``NodeVmInfo`` class contains information to describe the Virtual Machine of a node of a VCHA cluster. This class was added in vSphere API 6.7.1.
type ClusterNodeVmInfo struct {
    // The identifier of the virtual machine of the VCHA node. This property was added in vSphere API 6.7.1.
	Vm string
    // BIOS UUID for the node. This property was added in vSphere API 6.7.1.
	BiosUuid string
}

// The ``VmInfo`` class contains information to describe the Virtual Machines of passive and witness nodes of a VCHA cluster. This class was added in vSphere API 6.7.1.
type ClusterVmInfo struct {
    // The virtual machine information of the passive node. This property was added in vSphere API 6.7.1.
	Passive ClusterNodeVmInfo
    // The virtual machine information of the witness node. This property was added in vSphere API 6.7.1.
	Witness ClusterNodeVmInfo
}

// The ``UndeploySpec`` class contains the undeploy specification for a VCHA cluster. This class was added in vSphere API 6.7.1.
type ClusterUndeploySpec struct {
    // Contains the active node's management vCenter server credentials. This property was added in vSphere API 6.7.1.
	VcSpec *CredentialsSpec
    // Flag controlling in what circumstances the virtual machines will be deleted. For this flag to take effect, the VCHA cluster should have been successfully configured using automatic deployment. 
    //
    // * If true, the ClusterUndeploySpec#vms property will be ignored, the VCHA cluster specific information is removed, and the passive and witness virtual machines will be deleted.
    // * If false, the ClusterUndeploySpec#vms property contains the information identifying the passive and witness virtual machines.
    //
    //
    //     * If the ClusterUndeploySpec#vms property is set, then it will be validated prior to deleting the passive and witness virtual machines and VCHA cluster specific information is removed.
    //     * If the ClusterUndeploySpec#vms property is null, then the passive and witness virtual machines will not be deleted. The customer should delete them in order to cleanup completely. VCHA cluster specific information is removed.
    //
    //  
    // . This property was added in vSphere API 6.7.1.
	ForceDelete *bool
    // Contains virtual machine information for the passive and witness virtual machines. For this flag to take effect, the VCHA cluster should have been successfully configured using automatic deployment. 
    //
    //  If set, the ClusterUndeploySpec#forceDelete property controls whether this information is validated. 
    //
    // * If the ClusterUndeploySpec#forceDelete property is true, then this information is ignored, VCHA cluster specific information is removed and the passive and witness virtual machines will be deleted.
    // * If the ClusterUndeploySpec#forceDelete property is null or false, then this information is validated prior to deleting the passive and witness virtual machines. VCHA cluster specific information is removed.
    //
    // . This property was added in vSphere API 6.7.1.
	Vms *ClusterVmInfo
}



func clusterDeployInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewReferenceType(ClusterDeploySpecBindingType)
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clusterDeployOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func clusterDeployRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["spec"] = bindings.NewReferenceType(ClusterDeploySpecBindingType)
	fieldNameMap["spec"] = "Spec"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"InvalidArgument": 400,"Unauthorized": 403,"UnverifiedPeer": 400,"Error": 500})
}

func clusterFailoverInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["planned"] = bindings.NewBooleanType()
	fieldNameMap["planned"] = "Planned"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clusterFailoverOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func clusterFailoverRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["planned"] = bindings.NewBooleanType()
	fieldNameMap["planned"] = "Planned"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"Unauthorized": 403,"Error": 500})
}

func clusterGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vc_spec"] = bindings.NewOptionalType(bindings.NewReferenceType(CredentialsSpecBindingType))
	fields["partial"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["vc_spec"] = "VcSpec"
	fieldNameMap["partial"] = "Partial"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clusterGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(ClusterInfoBindingType)
}

func clusterGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["vc_spec"] = bindings.NewOptionalType(bindings.NewReferenceType(CredentialsSpecBindingType))
	fields["partial"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["vc_spec"] = "VcSpec"
	fieldNameMap["partial"] = "Partial"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"InvalidArgument": 400,"Unauthorized": 403,"UnverifiedPeer": 400,"Error": 500})
}

func clusterUndeployInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewReferenceType(ClusterUndeploySpecBindingType)
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clusterUndeployOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func clusterUndeployRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["spec"] = bindings.NewReferenceType(ClusterUndeploySpecBindingType)
	fieldNameMap["spec"] = "Spec"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"InvalidArgument": 400,"NotFound": 404,"Unauthorized": 403,"UnverifiedPeer": 400,"Error": 500})
}


func ClusterActiveSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ha_network_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vcha.network_type", reflect.TypeOf(NetworkType(NetworkType_STANDARD_PORTGROUP))))
	fieldNameMap["ha_network_type"] = "HaNetworkType"
	fields["ha_network"] = bindings.NewOptionalType(bindings.NewIdType([]string{"Network:VCenter"}, ""))
	fieldNameMap["ha_network"] = "HaNetwork"
	fields["ha_ip"] = bindings.NewReferenceType(IpSpecBindingType)
	fieldNameMap["ha_ip"] = "HaIp"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vcha.cluster.active_spec", fields, reflect.TypeOf(ClusterActiveSpec{}), fieldNameMap, validators)
}

func ClusterPassiveSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["placement"] = bindings.NewOptionalType(bindings.NewReferenceType(PlacementSpecBindingType))
	fieldNameMap["placement"] = "Placement"
	fields["ha_ip"] = bindings.NewReferenceType(IpSpecBindingType)
	fieldNameMap["ha_ip"] = "HaIp"
	fields["failover_ip"] = bindings.NewOptionalType(bindings.NewReferenceType(IpSpecBindingType))
	fieldNameMap["failover_ip"] = "FailoverIp"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vcha.cluster.passive_spec", fields, reflect.TypeOf(ClusterPassiveSpec{}), fieldNameMap, validators)
}

func ClusterWitnessSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["placement"] = bindings.NewOptionalType(bindings.NewReferenceType(PlacementSpecBindingType))
	fieldNameMap["placement"] = "Placement"
	fields["ha_ip"] = bindings.NewReferenceType(IpSpecBindingType)
	fieldNameMap["ha_ip"] = "HaIp"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vcha.cluster.witness_spec", fields, reflect.TypeOf(ClusterWitnessSpec{}), fieldNameMap, validators)
}

func ClusterDeploySpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vc_spec"] = bindings.NewOptionalType(bindings.NewReferenceType(CredentialsSpecBindingType))
	fieldNameMap["vc_spec"] = "VcSpec"
	fields["deployment"] = bindings.NewEnumType("com.vmware.vcenter.vcha.cluster.type", reflect.TypeOf(ClusterType(ClusterType_AUTO)))
	fieldNameMap["deployment"] = "Deployment"
	fields["active"] = bindings.NewReferenceType(ClusterActiveSpecBindingType)
	fieldNameMap["active"] = "Active"
	fields["passive"] = bindings.NewReferenceType(ClusterPassiveSpecBindingType)
	fieldNameMap["passive"] = "Passive"
	fields["witness"] = bindings.NewReferenceType(ClusterWitnessSpecBindingType)
	fieldNameMap["witness"] = "Witness"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vcha.cluster.deploy_spec", fields, reflect.TypeOf(ClusterDeploySpec{}), fieldNameMap, validators)
}

func ClusterNodeRuntimeInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["state"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vcha.cluster.node_state", reflect.TypeOf(ClusterNodeState(ClusterNodeState_UP))))
	fieldNameMap["state"] = "State"
	fields["role"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vcha.cluster.node_role", reflect.TypeOf(ClusterNodeRole(ClusterNodeRole_ACTIVE))))
	fieldNameMap["role"] = "Role"
	fields["placement"] = bindings.NewOptionalType(bindings.NewReferenceType(PlacementInfoBindingType))
	fieldNameMap["placement"] = "Placement"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vcha.cluster.node_runtime_info", fields, reflect.TypeOf(ClusterNodeRuntimeInfo{}), fieldNameMap, validators)
}

func ClusterIpv4InfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["address"] = bindings.NewStringType()
	fieldNameMap["address"] = "Address"
	fields["subnet_mask"] = bindings.NewStringType()
	fieldNameMap["subnet_mask"] = "SubnetMask"
	fields["prefix"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["prefix"] = "Prefix"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vcha.cluster.ipv4_info", fields, reflect.TypeOf(ClusterIpv4Info{}), fieldNameMap, validators)
}

func ClusterIpv6InfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["address"] = bindings.NewStringType()
	fieldNameMap["address"] = "Address"
	fields["prefix"] = bindings.NewIntegerType()
	fieldNameMap["prefix"] = "Prefix"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vcha.cluster.ipv6_info", fields, reflect.TypeOf(ClusterIpv6Info{}), fieldNameMap, validators)
}

func ClusterIpInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ip_family"] = bindings.NewEnumType("com.vmware.vcenter.vcha.cluster.ip_family", reflect.TypeOf(ClusterIpFamily(ClusterIpFamily_IPV4)))
	fieldNameMap["ip_family"] = "IpFamily"
	fields["ipv4"] = bindings.NewOptionalType(bindings.NewReferenceType(ClusterIpv4InfoBindingType))
	fieldNameMap["ipv4"] = "Ipv4"
	fields["ipv6"] = bindings.NewOptionalType(bindings.NewReferenceType(ClusterIpv6InfoBindingType))
	fieldNameMap["ipv6"] = "Ipv6"
	fields["gateway_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["gateway_ip"] = "GatewayIp"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("ip_family",
		map[string][]bindings.FieldData{
			"IPV4": []bindings.FieldData{
				bindings.NewFieldData("ipv4", true),
			},
			"IPV6": []bindings.FieldData{
				bindings.NewFieldData("ipv6", true),
			},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.vcha.cluster.ip_info", fields, reflect.TypeOf(ClusterIpInfo{}), fieldNameMap, validators)
}

func ClusterNodeInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["failover_ip"] = bindings.NewOptionalType(bindings.NewReferenceType(ClusterIpInfoBindingType))
	fieldNameMap["failover_ip"] = "FailoverIp"
	fields["ha_ip"] = bindings.NewReferenceType(ClusterIpInfoBindingType)
	fieldNameMap["ha_ip"] = "HaIp"
	fields["runtime"] = bindings.NewOptionalType(bindings.NewReferenceType(ClusterNodeRuntimeInfoBindingType))
	fieldNameMap["runtime"] = "Runtime"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vcha.cluster.node_info", fields, reflect.TypeOf(ClusterNodeInfo{}), fieldNameMap, validators)
}

func ClusterWitnessInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ha_ip"] = bindings.NewReferenceType(ClusterIpInfoBindingType)
	fieldNameMap["ha_ip"] = "HaIp"
	fields["runtime"] = bindings.NewOptionalType(bindings.NewReferenceType(ClusterNodeRuntimeInfoBindingType))
	fieldNameMap["runtime"] = "Runtime"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vcha.cluster.witness_info", fields, reflect.TypeOf(ClusterWitnessInfo{}), fieldNameMap, validators)
}

func ClusterErrorConditionBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["error"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
	fieldNameMap["error"] = "Error_"
	fields["recommendation"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
	fieldNameMap["recommendation"] = "Recommendation"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vcha.cluster.error_condition", fields, reflect.TypeOf(ClusterErrorCondition{}), fieldNameMap, validators)
}

func ClusterInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["config_state"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vcha.cluster.config_state", reflect.TypeOf(ClusterConfigState(ClusterConfigState_CONFIGURED))))
	fieldNameMap["config_state"] = "ConfigState"
	fields["node1"] = bindings.NewOptionalType(bindings.NewReferenceType(ClusterNodeInfoBindingType))
	fieldNameMap["node1"] = "Node1"
	fields["node2"] = bindings.NewOptionalType(bindings.NewReferenceType(ClusterNodeInfoBindingType))
	fieldNameMap["node2"] = "Node2"
	fields["witness"] = bindings.NewOptionalType(bindings.NewReferenceType(ClusterWitnessInfoBindingType))
	fieldNameMap["witness"] = "Witness"
	fields["mode"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vcha.cluster.cluster_mode", reflect.TypeOf(ClusterClusterMode(ClusterClusterMode_ENABLED))))
	fieldNameMap["mode"] = "Mode"
	fields["health_state"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vcha.cluster.cluster_state", reflect.TypeOf(ClusterClusterState(ClusterClusterState_HEALTHY))))
	fieldNameMap["health_state"] = "HealthState"
	fields["health_exception"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{})))
	fieldNameMap["health_exception"] = "HealthException"
	fields["health_warnings"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ClusterErrorConditionBindingType), reflect.TypeOf([]ClusterErrorCondition{})))
	fieldNameMap["health_warnings"] = "HealthWarnings"
	fields["manual_failover_allowed"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["manual_failover_allowed"] = "ManualFailoverAllowed"
	fields["auto_failover_allowed"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["auto_failover_allowed"] = "AutoFailoverAllowed"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vcha.cluster.info", fields, reflect.TypeOf(ClusterInfo{}), fieldNameMap, validators)
}

func ClusterNodeVmInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine:VCenter"}, "")
	fieldNameMap["vm"] = "Vm"
	fields["bios_uuid"] = bindings.NewStringType()
	fieldNameMap["bios_uuid"] = "BiosUuid"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vcha.cluster.node_vm_info", fields, reflect.TypeOf(ClusterNodeVmInfo{}), fieldNameMap, validators)
}

func ClusterVmInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["passive"] = bindings.NewReferenceType(ClusterNodeVmInfoBindingType)
	fieldNameMap["passive"] = "Passive"
	fields["witness"] = bindings.NewReferenceType(ClusterNodeVmInfoBindingType)
	fieldNameMap["witness"] = "Witness"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vcha.cluster.vm_info", fields, reflect.TypeOf(ClusterVmInfo{}), fieldNameMap, validators)
}

func ClusterUndeploySpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vc_spec"] = bindings.NewOptionalType(bindings.NewReferenceType(CredentialsSpecBindingType))
	fieldNameMap["vc_spec"] = "VcSpec"
	fields["force_delete"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["force_delete"] = "ForceDelete"
	fields["vms"] = bindings.NewOptionalType(bindings.NewReferenceType(ClusterVmInfoBindingType))
	fieldNameMap["vms"] = "Vms"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vcha.cluster.undeploy_spec", fields, reflect.TypeOf(ClusterUndeploySpec{}), fieldNameMap, validators)
}

