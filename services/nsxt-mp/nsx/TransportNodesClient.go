// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: TransportNodes
// Used by client-side stubs.

package nsx

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type TransportNodesClient interface {

	//
	//
	// @param nodeIdParam (required)
	// @param transportNodePlacementConfigParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Addorupdateplacementreferences(nodeIdParam string, transportNodePlacementConfigParam nsxModel.TransportNodePlacementConfig) error

	// Edge transport node maintains its entry in many internal tables. In some cases a few of these entries might not get cleaned up during edge transport node deletion. This api cleans up any stale entries that may exist in the internal tables that store the Edge Transport Node data.
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Cleanstaleentries() error

	//
	//
	// Deprecated: This API element is deprecated.
	//
	// @param transportNodeParam (required)
	// @param coreAllocationProfileParam This request parameter allows to create an edge transport node using a core profile. It enables users to configure the core allocation profile at the node level, with the understanding that the edge is initially not a part of the cluster and must be added to a cluster later. (optional)
	// @return com.vmware.nsx.model.TransportNode
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(transportNodeParam nsxModel.TransportNode, coreAllocationProfileParam *string) (nsxModel.TransportNode, error)

	//
	//
	// Deprecated: This API element is deprecated.
	//
	// @param transportNodeIdParam (required)
	// @param forceParam If true, deleting the resource succeeds even if it is being referred as a resource reference. (optional, default to false)
	// @param unprepareHostParam Uninstall NSX components from host while deleting (optional, default to true)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(transportNodeIdParam string, forceParam *bool, unprepareHostParam *bool) error

	// Invoke DELETE request on target transport node
	//
	// @param targetNodeIdParam Target node UUID (required)
	// @param targetUriParam URI of API to invoke on target node (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws UnableToAllocateResource  Insufficient Storage
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Deleteontransportnode(targetNodeIdParam string, targetUriParam string) error

	// Disable flow cache for edge transport node. Caution: This involves restart of the edge dataplane and hence may lead to network disruption.
	//
	// @param transportNodeIdParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Disableflowcache(transportNodeIdParam string) error

	// Enable flow cache for edge transport node. Caution: This involves restart of the edge dataplane and hence may lead to network disruption.
	//
	// @param transportNodeIdParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Enableflowcache(transportNodeIdParam string) error

	//
	//
	// Deprecated: This API element is deprecated.
	//
	// @param transportNodeIdParam (required)
	// @return com.vmware.nsx.model.TransportNode
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(transportNodeIdParam string) (nsxModel.TransportNode, error)

	// Invoke GET request on target transport node
	//
	// @param targetNodeIdParam Target node UUID (required)
	// @param targetUriParam URI of API to invoke on target node (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws UnableToAllocateResource  Insufficient Storage
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Getontransportnode(targetNodeIdParam string, targetUriParam string) error

	//
	//
	// Deprecated: This API element is deprecated.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param inMaintenanceModeParam If the flag is true, transport node with 'ENABLED' or 'FORCE_ENABLED' desired state will be returned, otherwise transport nodes in 'DISABLED' will be returned. (optional)
	// @param includedFieldsParam Note - this parameter currently only works when used with the search APIs /policy/api/v1/search/query and /policy/api/v1/search/dsl. It is ignored for other list APIs. (optional)
	// @param nodeIdParam This property can be used by itself or along with 'transport_zone_id'. (optional)
	// @param nodeIpParam This property can only be used alone. It can not be combined with other filtering properties. If the ESX host has both IPv4 and IPv6 addresses, and the NSX Manager has both IPv4 and IPv6 addresses, then this filter will work only on the IPv6 address of the ESX host. In all other cases, this filter will work only on the IPv4 address of ESX host. (optional)
	// @param nodeTypesParam The fabric node type is the resource_type of the Node such as HostNode, EdgeNode and PublicCloudGatewayNode. If a list of fabric node types are given, all transport nodes of all given types will be returned. Sorting the TransportNode list is supported only if this attribute is provided. (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam If true, results are sorted in ascending order (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param transportZoneIdParam This propery can be used along with 'node_id'. (optional)
	// @return com.vmware.nsx.model.TransportNodeListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(cursorParam *string, inMaintenanceModeParam *bool, includedFieldsParam *string, nodeIdParam *string, nodeIpParam *string, nodeTypesParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, transportZoneIdParam *string) (nsxModel.TransportNodeListResult, error)

	// Migrates all NVDS to VDS on given TransportNode. Upgrade precheck apis should have been run prior to invoking this API on transport node and a migration topology should be created. Please refer to Migration guide for details about migration APIs.
	//
	// Deprecated: This API element is deprecated.
	//
	// @param transportNodeIdParam (required)
	// @param skipMaintmodeParam Skipping maintenance mode check before starting migration. This parameter is only used by SDDC environment. (optional, default to false)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Migratetovds(transportNodeIdParam string, skipMaintmodeParam *bool) error

	// Invoke POST request on target transport node
	//
	// @param targetNodeIdParam Target node UUID (required)
	// @param targetUriParam URI of API to invoke on target node (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws UnableToAllocateResource  Insufficient Storage
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Postontransportnode(targetNodeIdParam string, targetUriParam string) error

	// Invoke PUT request on target transport node
	//
	// @param targetNodeIdParam Target node UUID (required)
	// @param targetUriParam URI of API to invoke on target node (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws UnableToAllocateResource  Insufficient Storage
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Putontransportnode(targetNodeIdParam string, targetUriParam string) error

	//
	//
	// Deprecated: This API element is deprecated.
	//
	// @param nodeIdParam (required)
	// @param transportNodeParam (required)
	// @return com.vmware.nsx.model.TransportNode
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Redeploy(nodeIdParam string, transportNodeParam nsxModel.TransportNode) (nsxModel.TransportNode, error)

	//
	//
	// Deprecated: This API element is deprecated.
	//
	// @param transportNodeIdParam (required)
	// @param readOnlyParam If readonly is true, then internal objects like resource reservation, allocation etc are updated. Note that Edge transport node intent is not updated. If it is false or if not set, then internal objects and Edge transport node intent both are updated (optional, default to false)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Refreshnodeconfiguration(transportNodeIdParam string, readOnlyParam *bool) error

	// Restart the inventory sync for the node if it is currently internally paused. After this action the next inventory sync coming from the node is processed.
	//
	// @param transportNodeIdParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Restartinventorysync(transportNodeIdParam string) error

	// A host can be overridden to have different configuration than Transport Node Profile(TNP) on cluster. This action will restore such overridden host back to cluster level TNP. This API can be used in other case. When TNP is applied to a cluster, if any validation fails (e.g. VMs running on host) then existing transport node (TN) is not updated. In that case after the issue is resolved manually (e.g. VMs powered off), you can call this API to update TN as per cluster level TNP.
	//  This api is now deprecated. Please use new api - /infra/sites/<site-id>/enforcement-points/<enforcementpoint-id>/host-transport-nodes/<host-transport-node-id>?action=restore_cluster_config
	//
	// Deprecated: This API element is deprecated.
	//
	// @param transportNodeIdParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Restoreclusterconfig(transportNodeIdParam string) error

	// Resync the TransportNode configuration on a host. It is similar to updating the TransportNode with existing configuration, but force synce these configurations to the host (no backend optimizations).
	//
	// @param transportnodeIdParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Resynchostconfig(transportnodeIdParam string) error

	//
	//
	// Deprecated: This API element is deprecated.
	//
	// @param transportNodeIdParam (required)
	// @param transportNodeParam (required)
	// @param esxMgmtIfMigrationDestParam A comma separated list of network ids. When migrating vmks into logical switches, the ids are the logical switches's ids. When migrating out of logical switches, the ids are vSphere Standard Switch portgroup names in a single vSphere Standard Switch, or distributed virtual portgroup names in a single distributed virtual switch (DVS). This property can only used together with 'if_id'. (optional)
	// @param ifIdParam A comma separated list of vmk interfaces (for example, vmk0,vmk1). This property can only used along with 'esx_mgmt_if_migration_dest'. If all vmk interfaces will be migrated into the same logical switch or DV portgroup, the 'esx_mgmt_if_migration_dest' can be just one logical switch id or DV portgroup name. Otherwise the number of vmks in this list must equal the number of ids in 'esx_mgmt_if_migration_dest' list, and the orders of the two lists are important because the vmks match the network ids one by one in the same order. (optional)
	// @param overrideNsxOwnershipParam Flag indicating whether the NSX ownership constraints (on Managed Objects like Host/Cluster/DVS) should be overridden/bypassed. Note: Overriding/bypassing NSX ownership constraints is not recommended at all. This indicates, you want to use/configure/own certain Managed Objects (like Cluster, Host or DVS) which seem to be already in use/configured/owned by some other NSX instance. This option should be used with caution. It should only be used to come out of situations where: a. The other NSX instance no longer intends to use the Managed Objects (and has already unconfigured NSX configurations) but the ownership still lies with it (incorrectly) and you want those Managed Objects to be used/configured/owned by this NSX instance. b. The other NSX instance has crashed or decommisioned but the ownership still lies with it and you want those Managed Objects to be used/configured/owned by this NSX instance. Enabling this option, while the Managed Objects affected by this operation are actively used by other NSX, can lead to problematic states on both the NSX instances. For example, if a TN is forcefully reconfigured by this NSX instance (using override_nsx_ownership=true), while it was already configured and in use by the other NSX instance, it could corrupt the HostSwitch configurations pushed down by the other NSX instance. (optional, default to false)
	// @param pingIpParam A comma separated list of IP addresses that match the vmk interfaces given in property 'if_id\" or 'vnic' one-by-one in the same order. '0.0.0.0' is a special IP that indicates the pre-migration gateway of the vmk will be pinged post-migration. If a VMK does not need the ping ip or a VM NIC is given inside 'vnic', the ping ip must be skipped but the comma has to stay. For example, '0.0.0.0,,10.1.1.1' indicates the vmk or VM NIC at the 2nd position does not need ping post-migration. Right after all ESX vmk interfaces are migrated, ping packets will be sent through each vmk to its given ping_ip to check if the migraton will break the network connectivity or not. If any vmk_ping fails, the whole migration of all vmks will be rolled back and transport-node will be in failed state. (optional)
	// @param skipValidationParam If this property is set true, all front-end validation for vmk, vnic, and/or pnic migration will be skipped. This is useful when the remote host becomes unreachable as a result of a migration; in which case the front-end validation will always fail because data from the remote host is no longer available. Skipping the validation will allow user to undo the migration by updating the transport node first and then restoring the host network connectivity. (optional, default to false)
	// @param vnicParam A comma separated list of vmk interfaces and/or one VM NIC. Only one VM NIC is allowed in the list; the format must be vmInstanceUuid:DeviceId like '50ca5f2d-1fa2-432d-991e-f01e0e16d182:4000'. An example list is 'vmk0,vmk1,50ca5f2d-1fa2-432d-991e-f01e0e16d182:4000'. The property can only be used along with 'vnic_migration_dest'. (optional)
	// @param vnicMigrationDestParam A comma separated list of vif ids, or port group names. When migrating into logical switches, the ids are vif ids in the logical ports created in the logical switches. When migrating out of logical switches, the ids are vSphere Standard Switch portgroup names in a single vSphere Standard Switch, or distributed virtual portgroup names in a single distributed virtual switch (DVS). The property can only be used in combination with property 'vnic'. The number of vnic interfaces in 'vnic' must equal the number of vif ids or port-group names in this list. The items in the two lists match by the the order. (optional)
	// @return com.vmware.nsx.model.TransportNode
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(transportNodeIdParam string, transportNodeParam nsxModel.TransportNode, esxMgmtIfMigrationDestParam *string, ifIdParam *string, overrideNsxOwnershipParam *bool, pingIpParam *string, skipValidationParam *bool, vnicParam *string, vnicMigrationDestParam *string) (nsxModel.TransportNode, error)

	//
	//
	// Deprecated: This API element is deprecated.
	//
	// @param transportnodeIdParam (required)
	// @param actionParam User could use this parameter to put transport node into maintenance mode or exit from maintenance mode. 'enter_maintenance_mode' will put Transport Node into maintenance mode if there is no VIFs attached. 'forced_enter_maintenance_mode' will put transport node into maintenance mode forcibly regardless of whether or not VIF attached. 'exit_maintenance_mode' will exit from maintenance mode. (optional)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Updatemaintenancemode(transportnodeIdParam string, actionParam *string) error
}

type transportNodesClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewTransportNodesClient(connector vapiProtocolClient_.Connector) *transportNodesClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.transport_nodes")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"addorupdateplacementreferences": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "addorupdateplacementreferences"),
		"cleanstaleentries":              vapiCore_.NewMethodIdentifier(interfaceIdentifier, "cleanstaleentries"),
		"create":                         vapiCore_.NewMethodIdentifier(interfaceIdentifier, "create"),
		"delete":                         vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"deleteontransportnode":          vapiCore_.NewMethodIdentifier(interfaceIdentifier, "deleteontransportnode"),
		"disableflowcache":               vapiCore_.NewMethodIdentifier(interfaceIdentifier, "disableflowcache"),
		"enableflowcache":                vapiCore_.NewMethodIdentifier(interfaceIdentifier, "enableflowcache"),
		"get":                            vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"getontransportnode":             vapiCore_.NewMethodIdentifier(interfaceIdentifier, "getontransportnode"),
		"list":                           vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"migratetovds":                   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "migratetovds"),
		"postontransportnode":            vapiCore_.NewMethodIdentifier(interfaceIdentifier, "postontransportnode"),
		"putontransportnode":             vapiCore_.NewMethodIdentifier(interfaceIdentifier, "putontransportnode"),
		"redeploy":                       vapiCore_.NewMethodIdentifier(interfaceIdentifier, "redeploy"),
		"refreshnodeconfiguration":       vapiCore_.NewMethodIdentifier(interfaceIdentifier, "refreshnodeconfiguration"),
		"restartinventorysync":           vapiCore_.NewMethodIdentifier(interfaceIdentifier, "restartinventorysync"),
		"restoreclusterconfig":           vapiCore_.NewMethodIdentifier(interfaceIdentifier, "restoreclusterconfig"),
		"resynchostconfig":               vapiCore_.NewMethodIdentifier(interfaceIdentifier, "resynchostconfig"),
		"update":                         vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
		"updatemaintenancemode":          vapiCore_.NewMethodIdentifier(interfaceIdentifier, "updatemaintenancemode"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	tIface := transportNodesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &tIface
}

func (tIface *transportNodesClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := tIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (tIface *transportNodesClient) Addorupdateplacementreferences(nodeIdParam string, transportNodePlacementConfigParam nsxModel.TransportNodePlacementConfig) error {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := transportNodesAddorupdateplacementreferencesRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(transportNodesAddorupdateplacementreferencesInputType(), typeConverter)
	sv.AddStructField("NodeId", nodeIdParam)
	sv.AddStructField("TransportNodePlacementConfig", transportNodePlacementConfigParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_nodes", "addorupdateplacementreferences", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (tIface *transportNodesClient) Cleanstaleentries() error {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := transportNodesCleanstaleentriesRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(transportNodesCleanstaleentriesInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_nodes", "cleanstaleentries", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (tIface *transportNodesClient) Create(transportNodeParam nsxModel.TransportNode, coreAllocationProfileParam *string) (nsxModel.TransportNode, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := transportNodesCreateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(transportNodesCreateInputType(), typeConverter)
	sv.AddStructField("TransportNode", transportNodeParam)
	sv.AddStructField("CoreAllocationProfile", coreAllocationProfileParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.TransportNode
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_nodes", "create", inputDataValue, executionContext)
	var emptyOutput nsxModel.TransportNode
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), TransportNodesCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.TransportNode), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *transportNodesClient) Delete(transportNodeIdParam string, forceParam *bool, unprepareHostParam *bool) error {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := transportNodesDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(transportNodesDeleteInputType(), typeConverter)
	sv.AddStructField("TransportNodeId", transportNodeIdParam)
	sv.AddStructField("Force", forceParam)
	sv.AddStructField("UnprepareHost", unprepareHostParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_nodes", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (tIface *transportNodesClient) Deleteontransportnode(targetNodeIdParam string, targetUriParam string) error {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := transportNodesDeleteontransportnodeRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(transportNodesDeleteontransportnodeInputType(), typeConverter)
	sv.AddStructField("TargetNodeId", targetNodeIdParam)
	sv.AddStructField("TargetUri", targetUriParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_nodes", "deleteontransportnode", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (tIface *transportNodesClient) Disableflowcache(transportNodeIdParam string) error {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := transportNodesDisableflowcacheRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(transportNodesDisableflowcacheInputType(), typeConverter)
	sv.AddStructField("TransportNodeId", transportNodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_nodes", "disableflowcache", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (tIface *transportNodesClient) Enableflowcache(transportNodeIdParam string) error {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := transportNodesEnableflowcacheRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(transportNodesEnableflowcacheInputType(), typeConverter)
	sv.AddStructField("TransportNodeId", transportNodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_nodes", "enableflowcache", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (tIface *transportNodesClient) Get(transportNodeIdParam string) (nsxModel.TransportNode, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := transportNodesGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(transportNodesGetInputType(), typeConverter)
	sv.AddStructField("TransportNodeId", transportNodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.TransportNode
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_nodes", "get", inputDataValue, executionContext)
	var emptyOutput nsxModel.TransportNode
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), TransportNodesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.TransportNode), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *transportNodesClient) Getontransportnode(targetNodeIdParam string, targetUriParam string) error {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := transportNodesGetontransportnodeRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(transportNodesGetontransportnodeInputType(), typeConverter)
	sv.AddStructField("TargetNodeId", targetNodeIdParam)
	sv.AddStructField("TargetUri", targetUriParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_nodes", "getontransportnode", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (tIface *transportNodesClient) List(cursorParam *string, inMaintenanceModeParam *bool, includedFieldsParam *string, nodeIdParam *string, nodeIpParam *string, nodeTypesParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, transportZoneIdParam *string) (nsxModel.TransportNodeListResult, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := transportNodesListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(transportNodesListInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("InMaintenanceMode", inMaintenanceModeParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("NodeId", nodeIdParam)
	sv.AddStructField("NodeIp", nodeIpParam)
	sv.AddStructField("NodeTypes", nodeTypesParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("TransportZoneId", transportZoneIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.TransportNodeListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_nodes", "list", inputDataValue, executionContext)
	var emptyOutput nsxModel.TransportNodeListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), TransportNodesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.TransportNodeListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *transportNodesClient) Migratetovds(transportNodeIdParam string, skipMaintmodeParam *bool) error {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := transportNodesMigratetovdsRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(transportNodesMigratetovdsInputType(), typeConverter)
	sv.AddStructField("TransportNodeId", transportNodeIdParam)
	sv.AddStructField("SkipMaintmode", skipMaintmodeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_nodes", "migratetovds", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (tIface *transportNodesClient) Postontransportnode(targetNodeIdParam string, targetUriParam string) error {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := transportNodesPostontransportnodeRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(transportNodesPostontransportnodeInputType(), typeConverter)
	sv.AddStructField("TargetNodeId", targetNodeIdParam)
	sv.AddStructField("TargetUri", targetUriParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_nodes", "postontransportnode", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (tIface *transportNodesClient) Putontransportnode(targetNodeIdParam string, targetUriParam string) error {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := transportNodesPutontransportnodeRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(transportNodesPutontransportnodeInputType(), typeConverter)
	sv.AddStructField("TargetNodeId", targetNodeIdParam)
	sv.AddStructField("TargetUri", targetUriParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_nodes", "putontransportnode", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (tIface *transportNodesClient) Redeploy(nodeIdParam string, transportNodeParam nsxModel.TransportNode) (nsxModel.TransportNode, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := transportNodesRedeployRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(transportNodesRedeployInputType(), typeConverter)
	sv.AddStructField("NodeId", nodeIdParam)
	sv.AddStructField("TransportNode", transportNodeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.TransportNode
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_nodes", "redeploy", inputDataValue, executionContext)
	var emptyOutput nsxModel.TransportNode
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), TransportNodesRedeployOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.TransportNode), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *transportNodesClient) Refreshnodeconfiguration(transportNodeIdParam string, readOnlyParam *bool) error {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := transportNodesRefreshnodeconfigurationRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(transportNodesRefreshnodeconfigurationInputType(), typeConverter)
	sv.AddStructField("TransportNodeId", transportNodeIdParam)
	sv.AddStructField("ReadOnly", readOnlyParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_nodes", "refreshnodeconfiguration", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (tIface *transportNodesClient) Restartinventorysync(transportNodeIdParam string) error {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := transportNodesRestartinventorysyncRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(transportNodesRestartinventorysyncInputType(), typeConverter)
	sv.AddStructField("TransportNodeId", transportNodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_nodes", "restartinventorysync", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (tIface *transportNodesClient) Restoreclusterconfig(transportNodeIdParam string) error {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := transportNodesRestoreclusterconfigRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(transportNodesRestoreclusterconfigInputType(), typeConverter)
	sv.AddStructField("TransportNodeId", transportNodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_nodes", "restoreclusterconfig", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (tIface *transportNodesClient) Resynchostconfig(transportnodeIdParam string) error {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := transportNodesResynchostconfigRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(transportNodesResynchostconfigInputType(), typeConverter)
	sv.AddStructField("TransportnodeId", transportnodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_nodes", "resynchostconfig", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (tIface *transportNodesClient) Update(transportNodeIdParam string, transportNodeParam nsxModel.TransportNode, esxMgmtIfMigrationDestParam *string, ifIdParam *string, overrideNsxOwnershipParam *bool, pingIpParam *string, skipValidationParam *bool, vnicParam *string, vnicMigrationDestParam *string) (nsxModel.TransportNode, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := transportNodesUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(transportNodesUpdateInputType(), typeConverter)
	sv.AddStructField("TransportNodeId", transportNodeIdParam)
	sv.AddStructField("TransportNode", transportNodeParam)
	sv.AddStructField("EsxMgmtIfMigrationDest", esxMgmtIfMigrationDestParam)
	sv.AddStructField("IfId", ifIdParam)
	sv.AddStructField("OverrideNsxOwnership", overrideNsxOwnershipParam)
	sv.AddStructField("PingIp", pingIpParam)
	sv.AddStructField("SkipValidation", skipValidationParam)
	sv.AddStructField("Vnic", vnicParam)
	sv.AddStructField("VnicMigrationDest", vnicMigrationDestParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.TransportNode
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_nodes", "update", inputDataValue, executionContext)
	var emptyOutput nsxModel.TransportNode
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), TransportNodesUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.TransportNode), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *transportNodesClient) Updatemaintenancemode(transportnodeIdParam string, actionParam *string) error {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := transportNodesUpdatemaintenancemodeRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(transportNodesUpdatemaintenancemodeInputType(), typeConverter)
	sv.AddStructField("TransportnodeId", transportnodeIdParam)
	sv.AddStructField("Action", actionParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_nodes", "updatemaintenancemode", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
