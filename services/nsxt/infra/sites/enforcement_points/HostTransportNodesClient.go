// Copyright (c) 2019-2025 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: HostTransportNodes
// Used by client-side stubs.

package enforcement_points

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type HostTransportNodesClient interface {

	// Deletes the specified transport node. Query param force can be used to force delete the host nodes. Force delete is not supported if transport node is part of a cluster on which Transport node profile is applied. It also removes the specified host node from system. If unprepare_host option is set to false, then host will be deleted without uninstalling the NSX components from the host. If transport node delete is called with query param force not being set or set to false and uninstall of NSX components in the host fails, TransportNodeState object will be retained. If transport node delete is called with query param force set to true and uninstall of NSX components in the host fails, TransportNodeState object will be deleted.
	//
	// @param siteIdParam (required)
	// @param enforcementpointIdParam (required)
	// @param hostTransportNodeIdParam (required)
	// @param forceParam If true, deleting the resource succeeds even if it is being referred as a resource reference. (optional, default to false)
	// @param unprepareHostParam Uninstall NSX components from host while deleting (optional, default to true)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(siteIdParam string, enforcementpointIdParam string, hostTransportNodeIdParam string, forceParam *bool, unprepareHostParam *bool) error

	// Returns information about a specified transport node.
	//
	// @param siteIdParam (required)
	// @param enforcementpointIdParam (required)
	// @param hostTransportNodeIdParam (required)
	// @return com.vmware.nsx_policy.model.HostTransportNode
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(siteIdParam string, enforcementpointIdParam string, hostTransportNodeIdParam string) (nsx_policyModel.HostTransportNode, error)

	// Returns information about all host transport nodes along with underlying host details. A transport node is a host that contains hostswitches. A hostswitch can have virtual machines connected to them. Because each transport node has hostswitches, transport nodes can also have virtual tunnel endpoints, which means that they can be part of the overlay.
	//
	// @param siteIdParam (required)
	// @param enforcementpointIdParam (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param discoveredNodeIdParam This property can be used by itself or along with 'transport_zone_path'. This applies only to vCenter Managed hosts. For Unmanaged hosts use node_ip instead. These hosts are discovered by Nsx manager after adding a vCenter as Compute Manager. Refer to fabric discovered-nodes api to get discoverednode_id. eg. 6ab2278f-951d-471b-8d0f-510c825945f1:host-14 (optional)
	// @param inMaintenanceModeParam If the flag is true, transport node with 'ENABLED' or 'FORCE_ENABLED' desired state will be returned, otherwise transport nodes in 'DISABLED' will be returned. (optional)
	// @param includedFieldsParam Note - this parameter currently only works when used with the search APIs /policy/api/v1/search/query and /policy/api/v1/search/dsl. It is ignored for other list APIs. (optional)
	// @param nodeIpParam This property can only be used alone. It can not be combined with other filtering properties. If the ESX host has both IPv4 and IPv6 addresses, and the NSX Manager has both IPv4 and IPv6 addresses, then this filter will work only on the IPv6 address of the ESX host. In all other cases, this filter will work only on the IPv4 address of ESX host. (optional)
	// @param nodeTypesParam Specify types from [HostNode, EdgeNode, PublicCloudGatewayNode]. If a list of node types is given, all transport nodes of given types will be returned. (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam If true, results are sorted in ascending order (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param transportZonePathParam This property can be used along with 'node_id'. Valid Policy Tz path should be given, eg. /infra/sites/default/enforcement-points/default/transport-zones/web-tz1 (optional)
	// @return com.vmware.nsx_policy.model.HostTransportNodeListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(siteIdParam string, enforcementpointIdParam string, cursorParam *string, discoveredNodeIdParam *string, inMaintenanceModeParam *bool, includedFieldsParam *string, nodeIpParam *string, nodeTypesParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, transportZonePathParam *string) (nsx_policyModel.HostTransportNodeListResult, error)

	// Migrates all NVDS to VDS on given TransportNode. Upgrade precheck apis should have been run prior to invoking this API on transport node and a migration topology should be created. Please refer to Migration guide for details about migration APIs.
	//
	// Deprecated: This API element is deprecated.
	//
	// @param siteIdParam (required)
	// @param enforcementpointIdParam (required)
	// @param hostTransportNodeIdParam (required)
	// @param skipMaintmodeParam Skipping maintenance mode check before starting migration. This parameter is only used by SDDC environment. (optional, default to false)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Migratetovds(siteIdParam string, enforcementpointIdParam string, hostTransportNodeIdParam string, skipMaintmodeParam *bool) error

	// Transport nodes are hypervisor hosts that will participate in an NSX-T overlay. For a hypervisor host, this means that it hosts VMs that will communicate over NSX-T logical switches. This API creates transport node for a host node (hypervisor) in the transport network. When you run this command for a host, NSX Manager attempts to install the NSX kernel modules, which are packaged as VIB, RPM, or DEB files. For the installation to succeed, you must provide the host login credentials and the host thumbprint. To get the ESXi host thumbprint, SSH to the host and run the **openssl x509 -in /etc/vmware/ssl/rui.crt -fingerprint -sha256 -noout** command. To generate host key thumbprint using SHA-256 algorithm please follow the steps below. Log into the host, making sure that the connection is not vulnerable to a man in the middle attack. Check whether a public key already exists. Host public key is generally located at '/etc/ssh/ssh_host_rsa_key.pub'. If the key is not present then generate a new key by running the following command and follow the instructions. **ssh-keygen -t rsa** Now generate a SHA256 hash of the key using the following command. Please make sure to pass the appropriate file name if the public key is stored with a different file name other than the default 'id_rsa.pub'. **awk '{print $2}' id_rsa.pub | base64 -d | sha256sum -b | sed 's/ .\*$//' | xxd -r -p | base64** Additional documentation on creating a transport node can be found in the NSX-T Installation Guide. In order for the transport node to forward packets, the host_switch_spec property must be specified. Host switches (called bridges in OVS on KVM hypervisors) are the individual switches within the host virtual switch. Virtual machines are connected to the host switches. When creating a transport node, you need to specify if the host switches are already manually preconfigured on the node, or if NSX should create and manage the host switches. You specify this choice by the type of host switches you pass in the host_switch_spec property of the TransportNode request payload. For a KVM host, you can preconfigure the host switch, or you can have NSX Manager perform the configuration. For an ESXi host NSX Manager always configures the host switch. To preconfigure the host switches on a KVM host, pass an array of PreconfiguredHostSwitchSpec objects that describes those host switches. In the current NSX-T release, only one prefonfigured host switch can be specified. See the PreconfiguredHostSwitchSpec schema definition for documentation on the properties that must be provided. Preconfigured host switches are only supported on KVM hosts, not on ESXi hosts. To allow NSX to manage the host switch configuration on KVM hosts, ESXi hosts, pass an array of StandardHostSwitchSpec objects in the host_switch_spec property, and NSX will automatically create host switches with the properties you provide. In the current NSX-T release, up to 16 host switches can be automatically managed. See the StandardHostSwitchSpec schema definition for documentation on the properties that must be provided. The request should provide node_deployement_info.
	//
	// @param siteIdParam (required)
	// @param enforcementpointIdParam (required)
	// @param hostTransportNodeIdParam (required)
	// @param hostTransportNodeParam (required)
	// @param esxMgmtIfMigrationDestParam A comma separated list of network ids. When migrating vmks into logical switches, the ids are the logical switches's ids. When migrating out of logical switches, the ids are vSphere Standard Switch portgroup names in a single vSphere Standard Switch, or distributed virtual portgroup names in a single distributed virtual switch (DVS). This property can only used together with 'if_id'. (optional)
	// @param ifIdParam A comma separated list of vmk interfaces (for example, vmk0,vmk1). This property can only used along with 'esx_mgmt_if_migration_dest'. If all vmk interfaces will be migrated into the same logical switch or DV portgroup, the 'esx_mgmt_if_migration_dest' can be just one logical switch id or DV portgroup name. Otherwise the number of vmks in this list must equal the number of ids in 'esx_mgmt_if_migration_dest' list, and the orders of the two lists are important because the vmks match the network ids one by one in the same order. (optional)
	// @param overrideNsxOwnershipParam Flag indicating whether the NSX ownership constraints (on Managed Objects like Host/Cluster/DVS) should be overridden/bypassed. Note: Overriding/bypassing NSX ownership constraints is not recommended at all. This indicates, you want to use/configure/own certain Managed Objects (like Cluster, Host or DVS) which seem to be already in use/configured/owned by some other NSX instance. This option should be used with caution. It should only be used to come out of situations where: a. The other NSX instance no longer intends to use the Managed Objects (and has already unconfigured NSX configurations) but the ownership still lies with it (incorrectly) and you want those Managed Objects to be used/configured/owned by this NSX instance. b. The other NSX instance has crashed or decommisioned but the ownership still lies with it and you want those Managed Objects to be used/configured/owned by this NSX instance. Enabling this option, while the Managed Objects affected by this operation are actively used by other NSX, can lead to problematic states on both the NSX instances. For example, if a TN is forcefully reconfigured by this NSX instance (using override_nsx_ownership=true), while it was already configured and in use by the other NSX instance, it could corrupt the HostSwitch configurations pushed down by the other NSX instance. (optional, default to false)
	// @param pingIpParam A comma separated list of IP addresses that match the vmk interfaces given in property 'if_id\" or 'vnic' one-by-one in the same order. '0.0.0.0' is a special IP that indicates the pre-migration gateway of the vmk will be pinged post-migration. If a VMK does not need the ping ip or a VM NIC is given inside 'vnic', the ping ip must be skipped but the comma has to stay. For example, '0.0.0.0,,10.1.1.1' indicates the vmk or VM NIC at the 2nd position does not need ping post-migration. Right after all ESX vmk interfaces are migrated, ping packets will be sent through each vmk to its given ping_ip to check if the migraton will break the network connectivity or not. If any vmk_ping fails, the whole migration of all vmks will be rolled back and transport-node will be in failed state. (optional)
	// @param skipValidationParam If this property is set true, all front-end validation for vmk, vnic, and/or pnic migration will be skipped. This is useful when the remote host becomes unreachable as a result of a migration; in which case the front-end validation will always fail because data from the remote host is no longer available. Skipping the validation will allow user to undo the migration by updating the transport node first and then restoring the host network connectivity. (optional, default to false)
	// @param vnicParam A comma separated list of vmk interfaces and/or one VM NIC. Only one VM NIC is allowed in the list; the format must be vmInstanceUuid:DeviceId like '50ca5f2d-1fa2-432d-991e-f01e0e16d182:4000'. An example list is 'vmk0,vmk1,50ca5f2d-1fa2-432d-991e-f01e0e16d182:4000'. The property can only be used along with 'vnic_migration_dest'. (optional)
	// @param vnicMigrationDestParam A comma separated list of vif ids, or port group names. When migrating into logical switches, the ids are vif ids in the logical ports created in the logical switches. When migrating out of logical switches, the ids are vSphere Standard Switch portgroup names in a single vSphere Standard Switch, or distributed virtual portgroup names in a single distributed virtual switch (DVS). The property can only be used in combination with property 'vnic'. The number of vnic interfaces in 'vnic' must equal the number of vif ids or port-group names in this list. The items in the two lists match by the the order. (optional)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(siteIdParam string, enforcementpointIdParam string, hostTransportNodeIdParam string, hostTransportNodeParam nsx_policyModel.HostTransportNode, esxMgmtIfMigrationDestParam *string, ifIdParam *string, overrideNsxOwnershipParam *bool, pingIpParam *string, skipValidationParam *bool, vnicParam *string, vnicMigrationDestParam *string) error

	// A host can be overridden to have different configuration than Transport Node Profile(TNP) on cluster. This action will restore such overridden host back to cluster level TNP. This API can be used in other case. When TNP is applied to a cluster, if any validation fails (e.g. VMs running on host) then existing transport node (TN) is not updated. In that case after the issue is resolved manually (e.g. VMs powered off), you can call this API to update TN as per cluster level TNP.
	//
	// @param siteIdParam (required)
	// @param enforcementpointIdParam (required)
	// @param hostTransportNodeIdParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Restoreclusterconfig(siteIdParam string, enforcementpointIdParam string, hostTransportNodeIdParam string) error

	// Resync the TransportNode configuration on a host. It is similar to updating the TransportNode with existing configuration, but force synce these configurations to the host (no backend optimizations).
	//
	// @param siteIdParam (required)
	// @param enforcementpointIdParam (required)
	// @param hostTransportNodeIdParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Resynchostconfig(siteIdParam string, enforcementpointIdParam string, hostTransportNodeIdParam string) error

	// Transport nodes are hypervisor hosts that will participate in an NSX-T overlay. For a hypervisor host, this means that it hosts VMs that will communicate over NSX-T logical switches. This API creates transport node for a host node (hypervisor) in the transport network. When you run this command for a host, NSX Manager attempts to install the NSX kernel modules, which are packaged as VIB, RPM, or DEB files. For the installation to succeed, you must provide the host login credentials and the host thumbprint. To get the ESXi host thumbprint, SSH to the host and run the **openssl x509 -in /etc/vmware/ssl/rui.crt -fingerprint -sha256 -noout** command. To generate host key thumbprint using SHA-256 algorithm please follow the steps below. Log into the host, making sure that the connection is not vulnerable to a man in the middle attack. Check whether a public key already exists. Host public key is generally located at '/etc/ssh/ssh_host_rsa_key.pub'. If the key is not present then generate a new key by running the following command and follow the instructions. **ssh-keygen -t rsa** Now generate a SHA256 hash of the key using the following command. Please make sure to pass the appropriate file name if the public key is stored with a different file name other than the default 'id_rsa.pub'. **awk '{print $2}' id_rsa.pub | base64 -d | sha256sum -b | sed 's/ .\*$//' | xxd -r -p | base64** Additional documentation on creating a transport node can be found in the NSX-T Installation Guide. In order for the transport node to forward packets, the host_switch_spec property must be specified. Host switches (called bridges in OVS on KVM hypervisors) are the individual switches within the host virtual switch. Virtual machines are connected to the host switches. When creating a transport node, you need to specify if the host switches are already manually preconfigured on the node, or if NSX should create and manage the host switches. You specify this choice by the type of host switches you pass in the host_switch_spec property of the TransportNode request payload. For a KVM host, you can preconfigure the host switch, or you can have NSX Manager perform the configuration. For an ESXi host NSX Manager always configures the host switch. To preconfigure the host switches on a KVM host, pass an array of PreconfiguredHostSwitchSpec objects that describes those host switches. In the current NSX-T release, only one prefonfigured host switch can be specified. See the PreconfiguredHostSwitchSpec schema definition for documentation on the properties that must be provided. Preconfigured host switches are only supported on KVM hosts, not on ESXi hosts. To allow NSX to manage the host switch configuration on KVM hosts, ESXi hosts, pass an array of StandardHostSwitchSpec objects in the host_switch_spec property, and NSX will automatically create host switches with the properties you provide. In the current NSX-T release, up to 16 host switches can be automatically managed. See the StandardHostSwitchSpec schema definition for documentation on the properties that must be provided. The request should provide node_deployement_info.
	//
	// @param siteIdParam (required)
	// @param enforcementpointIdParam (required)
	// @param hostTransportNodeIdParam (required)
	// @param hostTransportNodeParam (required)
	// @param esxMgmtIfMigrationDestParam A comma separated list of network ids. When migrating vmks into logical switches, the ids are the logical switches's ids. When migrating out of logical switches, the ids are vSphere Standard Switch portgroup names in a single vSphere Standard Switch, or distributed virtual portgroup names in a single distributed virtual switch (DVS). This property can only used together with 'if_id'. (optional)
	// @param ifIdParam A comma separated list of vmk interfaces (for example, vmk0,vmk1). This property can only used along with 'esx_mgmt_if_migration_dest'. If all vmk interfaces will be migrated into the same logical switch or DV portgroup, the 'esx_mgmt_if_migration_dest' can be just one logical switch id or DV portgroup name. Otherwise the number of vmks in this list must equal the number of ids in 'esx_mgmt_if_migration_dest' list, and the orders of the two lists are important because the vmks match the network ids one by one in the same order. (optional)
	// @param overrideNsxOwnershipParam Flag indicating whether the NSX ownership constraints (on Managed Objects like Host/Cluster/DVS) should be overridden/bypassed. Note: Overriding/bypassing NSX ownership constraints is not recommended at all. This indicates, you want to use/configure/own certain Managed Objects (like Cluster, Host or DVS) which seem to be already in use/configured/owned by some other NSX instance. This option should be used with caution. It should only be used to come out of situations where: a. The other NSX instance no longer intends to use the Managed Objects (and has already unconfigured NSX configurations) but the ownership still lies with it (incorrectly) and you want those Managed Objects to be used/configured/owned by this NSX instance. b. The other NSX instance has crashed or decommisioned but the ownership still lies with it and you want those Managed Objects to be used/configured/owned by this NSX instance. Enabling this option, while the Managed Objects affected by this operation are actively used by other NSX, can lead to problematic states on both the NSX instances. For example, if a TN is forcefully reconfigured by this NSX instance (using override_nsx_ownership=true), while it was already configured and in use by the other NSX instance, it could corrupt the HostSwitch configurations pushed down by the other NSX instance. (optional, default to false)
	// @param pingIpParam A comma separated list of IP addresses that match the vmk interfaces given in property 'if_id\" or 'vnic' one-by-one in the same order. '0.0.0.0' is a special IP that indicates the pre-migration gateway of the vmk will be pinged post-migration. If a VMK does not need the ping ip or a VM NIC is given inside 'vnic', the ping ip must be skipped but the comma has to stay. For example, '0.0.0.0,,10.1.1.1' indicates the vmk or VM NIC at the 2nd position does not need ping post-migration. Right after all ESX vmk interfaces are migrated, ping packets will be sent through each vmk to its given ping_ip to check if the migraton will break the network connectivity or not. If any vmk_ping fails, the whole migration of all vmks will be rolled back and transport-node will be in failed state. (optional)
	// @param skipValidationParam If this property is set true, all front-end validation for vmk, vnic, and/or pnic migration will be skipped. This is useful when the remote host becomes unreachable as a result of a migration; in which case the front-end validation will always fail because data from the remote host is no longer available. Skipping the validation will allow user to undo the migration by updating the transport node first and then restoring the host network connectivity. (optional, default to false)
	// @param vnicParam A comma separated list of vmk interfaces and/or one VM NIC. Only one VM NIC is allowed in the list; the format must be vmInstanceUuid:DeviceId like '50ca5f2d-1fa2-432d-991e-f01e0e16d182:4000'. An example list is 'vmk0,vmk1,50ca5f2d-1fa2-432d-991e-f01e0e16d182:4000'. The property can only be used along with 'vnic_migration_dest'. (optional)
	// @param vnicMigrationDestParam A comma separated list of vif ids, or port group names. When migrating into logical switches, the ids are vif ids in the logical ports created in the logical switches. When migrating out of logical switches, the ids are vSphere Standard Switch portgroup names in a single vSphere Standard Switch, or distributed virtual portgroup names in a single distributed virtual switch (DVS). The property can only be used in combination with property 'vnic'. The number of vnic interfaces in 'vnic' must equal the number of vif ids or port-group names in this list. The items in the two lists match by the the order. (optional)
	// @return com.vmware.nsx_policy.model.HostTransportNode
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(siteIdParam string, enforcementpointIdParam string, hostTransportNodeIdParam string, hostTransportNodeParam nsx_policyModel.HostTransportNode, esxMgmtIfMigrationDestParam *string, ifIdParam *string, overrideNsxOwnershipParam *bool, pingIpParam *string, skipValidationParam *bool, vnicParam *string, vnicMigrationDestParam *string) (nsx_policyModel.HostTransportNode, error)

	// Put transport node into maintenance mode or exit from maintenance mode. When HostTransportNode is in maintenance mode no configuration changes are allowed
	//
	// @param siteIdParam (required)
	// @param enforcementpointIdParam (required)
	// @param hostTransportNodeIdParam (required)
	// @param actionParam User could use this parameter to put transport node into maintenance mode or exit from maintenance mode. 'enter_maintenance_mode' will put Transport Node into maintenance mode if there is no VIFs attached. 'forced_enter_maintenance_mode' will put transport node into maintenance mode forcibly regardless of whether or not VIF attached. 'exit_maintenance_mode' will exit from maintenance mode. (optional)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Updatemaintenancemode(siteIdParam string, enforcementpointIdParam string, hostTransportNodeIdParam string, actionParam *string) error
}

type hostTransportNodesClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewHostTransportNodesClient(connector vapiProtocolClient_.Connector) *hostTransportNodesClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.sites.enforcement_points.host_transport_nodes")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"delete":                vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":                   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":                  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"migratetovds":          vapiCore_.NewMethodIdentifier(interfaceIdentifier, "migratetovds"),
		"patch":                 vapiCore_.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"restoreclusterconfig":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "restoreclusterconfig"),
		"resynchostconfig":      vapiCore_.NewMethodIdentifier(interfaceIdentifier, "resynchostconfig"),
		"update":                vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
		"updatemaintenancemode": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "updatemaintenancemode"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	hIface := hostTransportNodesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &hIface
}

func (hIface *hostTransportNodesClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := hIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (hIface *hostTransportNodesClient) Delete(siteIdParam string, enforcementpointIdParam string, hostTransportNodeIdParam string, forceParam *bool, unprepareHostParam *bool) error {
	typeConverter := hIface.connector.TypeConverter()
	executionContext := hIface.connector.NewExecutionContext()
	operationRestMetaData := hostTransportNodesDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(hostTransportNodesDeleteInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("HostTransportNodeId", hostTransportNodeIdParam)
	sv.AddStructField("Force", forceParam)
	sv.AddStructField("UnprepareHost", unprepareHostParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := hIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.host_transport_nodes", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), hIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (hIface *hostTransportNodesClient) Get(siteIdParam string, enforcementpointIdParam string, hostTransportNodeIdParam string) (nsx_policyModel.HostTransportNode, error) {
	typeConverter := hIface.connector.TypeConverter()
	executionContext := hIface.connector.NewExecutionContext()
	operationRestMetaData := hostTransportNodesGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(hostTransportNodesGetInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("HostTransportNodeId", hostTransportNodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.HostTransportNode
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := hIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.host_transport_nodes", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.HostTransportNode
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), HostTransportNodesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.HostTransportNode), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), hIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (hIface *hostTransportNodesClient) List(siteIdParam string, enforcementpointIdParam string, cursorParam *string, discoveredNodeIdParam *string, inMaintenanceModeParam *bool, includedFieldsParam *string, nodeIpParam *string, nodeTypesParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, transportZonePathParam *string) (nsx_policyModel.HostTransportNodeListResult, error) {
	typeConverter := hIface.connector.TypeConverter()
	executionContext := hIface.connector.NewExecutionContext()
	operationRestMetaData := hostTransportNodesListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(hostTransportNodesListInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("DiscoveredNodeId", discoveredNodeIdParam)
	sv.AddStructField("InMaintenanceMode", inMaintenanceModeParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("NodeIp", nodeIpParam)
	sv.AddStructField("NodeTypes", nodeTypesParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("TransportZonePath", transportZonePathParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.HostTransportNodeListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := hIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.host_transport_nodes", "list", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.HostTransportNodeListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), HostTransportNodesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.HostTransportNodeListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), hIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (hIface *hostTransportNodesClient) Migratetovds(siteIdParam string, enforcementpointIdParam string, hostTransportNodeIdParam string, skipMaintmodeParam *bool) error {
	typeConverter := hIface.connector.TypeConverter()
	executionContext := hIface.connector.NewExecutionContext()
	operationRestMetaData := hostTransportNodesMigratetovdsRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(hostTransportNodesMigratetovdsInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("HostTransportNodeId", hostTransportNodeIdParam)
	sv.AddStructField("SkipMaintmode", skipMaintmodeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := hIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.host_transport_nodes", "migratetovds", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), hIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (hIface *hostTransportNodesClient) Patch(siteIdParam string, enforcementpointIdParam string, hostTransportNodeIdParam string, hostTransportNodeParam nsx_policyModel.HostTransportNode, esxMgmtIfMigrationDestParam *string, ifIdParam *string, overrideNsxOwnershipParam *bool, pingIpParam *string, skipValidationParam *bool, vnicParam *string, vnicMigrationDestParam *string) error {
	typeConverter := hIface.connector.TypeConverter()
	executionContext := hIface.connector.NewExecutionContext()
	operationRestMetaData := hostTransportNodesPatchRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(hostTransportNodesPatchInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("HostTransportNodeId", hostTransportNodeIdParam)
	sv.AddStructField("HostTransportNode", hostTransportNodeParam)
	sv.AddStructField("EsxMgmtIfMigrationDest", esxMgmtIfMigrationDestParam)
	sv.AddStructField("IfId", ifIdParam)
	sv.AddStructField("OverrideNsxOwnership", overrideNsxOwnershipParam)
	sv.AddStructField("PingIp", pingIpParam)
	sv.AddStructField("SkipValidation", skipValidationParam)
	sv.AddStructField("Vnic", vnicParam)
	sv.AddStructField("VnicMigrationDest", vnicMigrationDestParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := hIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.host_transport_nodes", "patch", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), hIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (hIface *hostTransportNodesClient) Restoreclusterconfig(siteIdParam string, enforcementpointIdParam string, hostTransportNodeIdParam string) error {
	typeConverter := hIface.connector.TypeConverter()
	executionContext := hIface.connector.NewExecutionContext()
	operationRestMetaData := hostTransportNodesRestoreclusterconfigRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(hostTransportNodesRestoreclusterconfigInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("HostTransportNodeId", hostTransportNodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := hIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.host_transport_nodes", "restoreclusterconfig", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), hIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (hIface *hostTransportNodesClient) Resynchostconfig(siteIdParam string, enforcementpointIdParam string, hostTransportNodeIdParam string) error {
	typeConverter := hIface.connector.TypeConverter()
	executionContext := hIface.connector.NewExecutionContext()
	operationRestMetaData := hostTransportNodesResynchostconfigRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(hostTransportNodesResynchostconfigInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("HostTransportNodeId", hostTransportNodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := hIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.host_transport_nodes", "resynchostconfig", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), hIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (hIface *hostTransportNodesClient) Update(siteIdParam string, enforcementpointIdParam string, hostTransportNodeIdParam string, hostTransportNodeParam nsx_policyModel.HostTransportNode, esxMgmtIfMigrationDestParam *string, ifIdParam *string, overrideNsxOwnershipParam *bool, pingIpParam *string, skipValidationParam *bool, vnicParam *string, vnicMigrationDestParam *string) (nsx_policyModel.HostTransportNode, error) {
	typeConverter := hIface.connector.TypeConverter()
	executionContext := hIface.connector.NewExecutionContext()
	operationRestMetaData := hostTransportNodesUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(hostTransportNodesUpdateInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("HostTransportNodeId", hostTransportNodeIdParam)
	sv.AddStructField("HostTransportNode", hostTransportNodeParam)
	sv.AddStructField("EsxMgmtIfMigrationDest", esxMgmtIfMigrationDestParam)
	sv.AddStructField("IfId", ifIdParam)
	sv.AddStructField("OverrideNsxOwnership", overrideNsxOwnershipParam)
	sv.AddStructField("PingIp", pingIpParam)
	sv.AddStructField("SkipValidation", skipValidationParam)
	sv.AddStructField("Vnic", vnicParam)
	sv.AddStructField("VnicMigrationDest", vnicMigrationDestParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.HostTransportNode
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := hIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.host_transport_nodes", "update", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.HostTransportNode
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), HostTransportNodesUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.HostTransportNode), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), hIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (hIface *hostTransportNodesClient) Updatemaintenancemode(siteIdParam string, enforcementpointIdParam string, hostTransportNodeIdParam string, actionParam *string) error {
	typeConverter := hIface.connector.TypeConverter()
	executionContext := hIface.connector.NewExecutionContext()
	operationRestMetaData := hostTransportNodesUpdatemaintenancemodeRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(hostTransportNodesUpdatemaintenancemodeInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("HostTransportNodeId", hostTransportNodeIdParam)
	sv.AddStructField("Action", actionParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := hIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.host_transport_nodes", "updatemaintenancemode", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), hIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
