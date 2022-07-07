// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: TransportNodes
// Used by client-side stubs.

package nsx

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type TransportNodesClient interface {

	// Edge transport node maintains its entry in many internal tables. In some cases a few of these entries might not get cleaned up during edge transport node deletion. This api cleans up any stale entries that may exist in the internal tables that store the Edge Transport Node data.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Cleanstaleentries() error

	// Transport nodes are hypervisor hosts and NSX Edges that will participate in an NSX-T overlay. For a hypervisor host, this means that it hosts VMs that will communicate over NSX-T logical switches. For NSX Edges, this means that it will have logical router uplinks and downlinks. This API creates transport node for a host node (hypervisor) or edge node (router) in the transport network. When you run this command for a host, NSX Manager attempts to install the NSX kernel modules, which are packaged as VIB, RPM, or DEB files. For the installation to succeed, you must provide the host login credentials and the host thumbprint. To get the ESXi host thumbprint, SSH to the host and run the **openssl x509 -in /etc/vmware/ssl/rui.crt -fingerprint -sha256 -noout** command. To generate host key thumbprint using SHA-256 algorithm please follow the steps below. Log into the host, making sure that the connection is not vulnerable to a man in the middle attack. Check whether a public key already exists. Host public key is generally located at '/etc/ssh/ssh_host_rsa_key.pub'. If the key is not present then generate a new key by running the following command and follow the instructions. **ssh-keygen -t rsa** Now generate a SHA256 hash of the key using the following command. Please make sure to pass the appropriate file name if the public key is stored with a different file name other than the default 'id_rsa.pub'. **awk '{print $2}' id_rsa.pub | base64 -d | sha256sum -b | sed 's/ .\*$//' | xxd -r -p | base64** This api is deprecated as part of FN+TN unification. Please use Transport Node API to install NSX components on a node. Additional documentation on creating a transport node can be found in the NSX-T Installation Guide. In order for the transport node to forward packets, the host_switch_spec property must be specified. Host switches (called bridges in OVS on KVM hypervisors) are the individual switches within the host virtual switch. Virtual machines are connected to the host switches. When creating a transport node, you need to specify if the host switches are already manually preconfigured on the node, or if NSX should create and manage the host switches. You specify this choice by the type of host switches you pass in the host_switch_spec property of the TransportNode request payload. For a KVM host, you can preconfigure the host switch, or you can have NSX Manager perform the configuration. For an ESXi host or NSX Edge node, NSX Manager always configures the host switch. To preconfigure the host switches on a KVM host, pass an array of PreconfiguredHostSwitchSpec objects that describes those host switches. In the current NSX-T release, only one prefonfigured host switch can be specified. See the PreconfiguredHostSwitchSpec schema definition for documentation on the properties that must be provided. Preconfigured host switches are only supported on KVM hosts, not on ESXi hosts or NSX Edge nodes. To allow NSX to manage the host switch configuration on KVM hosts, ESXi hosts, or NSX Edge nodes, pass an array of StandardHostSwitchSpec objects in the host_switch_spec property, and NSX will automatically create host switches with the properties you provide. In the current NSX-T release, up to 16 host switches can be automatically managed. See the StandardHostSwitchSpec schema definition for documentation on the properties that must be provided. Note: Previous versions of NSX-T also used a property named transport_zone_endpoints at TransportNode level. This property is deprecated which creates some combinations of new client along with old client payloads. Examples [1] & [2] show old/existing client request and response by populating transport_zone_endpoints property at TransportNode level. Example [3] shows TransportNode creation request/response by populating transport_zone_endpoints property at StandardHostSwitch level and other new properties. The request should either provide node_deployement_info or node_id. If the host node (hypervisor) or edge node (router) is already added in system then it can be converted to transport node by providing node_id in request. If host node (hypervisor) or edge node (router) is not already present in system then information should be provided under node_deployment_info. This api is now deprecated. Please use new api - /infra/sites/<site-id>/enforcement-points/<enforcementpoint-id>/host-transport-nodes/<host-transport-node-id>
	//
	// @param transportNodeParam (required)
	// @return com.vmware.nsx.model.TransportNode
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(transportNodeParam model.TransportNode) (model.TransportNode, error)

	// Deletes the specified transport node. Query param force can be used to force delete the host nodes. Force deletion of edge and public cloud gateway nodes is not supported. It also removes the specified node (host or edge) from system. If unprepare_host option is set to false, then host will be deleted without uninstalling the NSX components from the host. This api is now deprecated. Please use new api - /infra/sites/<site-id>/enforcement-points/<enforcementpoint-id>/host-transport-nodes/<host-transport-node-id>
	//
	// @param transportNodeIdParam (required)
	// @param forceParam Force delete the resource even if it is being used somewhere (optional, default to false)
	// @param unprepareHostParam Uninstall NSX components from host while deleting (optional, default to true)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(transportNodeIdParam string, forceParam *bool, unprepareHostParam *bool) error

	// Invoke DELETE request on target transport node
	//
	// @param targetNodeIdParam Target node UUID (required)
	// @param targetUriParam URI of API to invoke on target node (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Deleteontransportnode(targetNodeIdParam string, targetUriParam string) error

	// Disable flow cache for edge transport node. Caution: This involves restart of the edge dataplane and hence may lead to network disruption.
	//
	// @param transportNodeIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Disableflowcache(transportNodeIdParam string) error

	// Enable flow cache for edge transport node. Caution: This involves restart of the edge dataplane and hence may lead to network disruption.
	//
	// @param transportNodeIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Enableflowcache(transportNodeIdParam string) error

	// Returns information about a specified transport node. This api is now deprecated. Please use new api - /infra/sites/<site-id>/enforcement-points/<enforcementpoint-id>/host-transport-nodes/<host-transport-node-id>
	//
	// @param transportNodeIdParam (required)
	// @return com.vmware.nsx.model.TransportNode
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(transportNodeIdParam string) (model.TransportNode, error)

	// Invoke GET request on target transport node
	//
	// @param targetNodeIdParam Target node UUID (required)
	// @param targetUriParam URI of API to invoke on target node (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Getontransportnode(targetNodeIdParam string, targetUriParam string) error

	// Returns information about all transport nodes along with underlying host or edge details. A transport node is a host or edge that contains hostswitches. A hostswitch can have virtual machines connected to them. Because each transport node has hostswitches, transport nodes can also have virtual tunnel endpoints, which means that they can be part of the overlay. This api is now deprecated. Please use new api - /infra/sites/<site-id>/enforcement-points/<enforcementpoint-id>/host-transport-nodes
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param inMaintenanceModeParam maintenance mode flag (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param nodeIdParam node identifier (optional)
	// @param nodeIpParam Fabric node IP address (optional)
	// @param nodeTypesParam a list of fabric node types separated by comma or a single type (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param transportZoneIdParam Transport zone identifier (optional)
	// @return com.vmware.nsx.model.TransportNodeListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(cursorParam *string, inMaintenanceModeParam *bool, includedFieldsParam *string, nodeIdParam *string, nodeIpParam *string, nodeTypesParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, transportZoneIdParam *string) (model.TransportNodeListResult, error)

	// Migrates all NVDS to VDS on given TransportNode. Upgrade precheck apis should have been run prior to invoking this API on transport node and a migration topology should be created. Please refer to Migration guide for details about migration APIs. This api is now deprecated. Please use new api - /infra/sites/<site-id>/enforcement-points/<enforcementpoint-id>/host-transport-nodes/<host-transport-node-id>?action=migrate_to_vds
	//
	// @param transportNodeIdParam (required)
	// @param skipMaintmodeParam Skip Maintenance mode check (optional, default to false)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Migratetovds(transportNodeIdParam string, skipMaintmodeParam *bool) error

	// Invoke POST request on target transport node
	//
	// @param targetNodeIdParam Target node UUID (required)
	// @param targetUriParam URI of API to invoke on target node (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Postontransportnode(targetNodeIdParam string, targetUriParam string) error

	// Invoke PUT request on target transport node
	//
	// @param targetNodeIdParam Target node UUID (required)
	// @param targetUriParam URI of API to invoke on target node (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Putontransportnode(targetNodeIdParam string, targetUriParam string) error

	// Redeploys an edge node at NSX Manager that replaces the edge node with identifier <node-id>. If NSX Manager can access the specified edge node, then the node is put into maintenance mode and then the associated VM is deleted. This is a means to reset all configuration on the edge node. The communication channel between NSX Manager and edge is established after this operation.
	//
	// @param nodeIdParam (required)
	// @param transportNodeParam (required)
	// @return com.vmware.nsx.model.TransportNode
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Redeploy(nodeIdParam string, transportNodeParam model.TransportNode) (model.TransportNode, error)

	// The API is applicable for Edge transport nodes. If you update the edge configuration and find a discrepancy in Edge configuration at NSX Manager in compare with realized, then use this API to refresh configuration at NSX Manager. It refreshes the Edge configuration from sources external to NSX Manager like vSphere Server or the Edge node CLI. After this action, Edge configuration at NSX Manager is updated and the API GET api/v1/transport-nodes will show refreshed data. From 3.2 release onwards, refresh API updates the MP intent by default.
	//
	// @param transportNodeIdParam (required)
	// @param readOnlyParam Read-only flag for Refresh API (optional, default to false)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Refreshnodeconfiguration(transportNodeIdParam string, readOnlyParam *bool) error

	// Restart the inventory sync for the node if it is currently internally paused. After this action the next inventory sync coming from the node is processed.
	//
	// @param transportNodeIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Restartinventorysync(transportNodeIdParam string) error

	// A host can be overridden to have different configuration than Transport Node Profile(TNP) on cluster. This action will restore such overridden host back to cluster level TNP. This API can be used in other case. When TNP is applied to a cluster, if any validation fails (e.g. VMs running on host) then existing transport node (TN) is not updated. In that case after the issue is resolved manually (e.g. VMs powered off), you can call this API to update TN as per cluster level TNP. This api is now deprecated. Please use new api - /infra/sites/<site-id>/enforcement-points/<enforcementpoint-id>/host-transport-nodes/<host-transport-node-id>?action=restore_cluster_config
	//
	// @param transportNodeIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Restoreclusterconfig(transportNodeIdParam string) error

	// Resync the TransportNode configuration on a host. It is similar to updating the TransportNode with existing configuration, but force synce these configurations to the host (no backend optimizations).
	//
	// @param transportnodeIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Resynchostconfig(transportnodeIdParam string) error

	// Modifies the transport node information. The host_switch_name field must match the host_switch_name value specified in the transport zone (API: transport-zones). You must create the associated uplink profile (API: host-switch-profiles) before you can specify an uplink_name here. If the host is an ESX and has only one physical NIC being used by a vSphere standard switch, TransportNodeUpdateParameters should be used to migrate the management interface and the physical NIC into a logical switch that is in a transport zone this transport node will join or has already joined. If the migration is already done, TransportNodeUpdateParameters can also be used to migrate the management interface and the physical NIC back to a vSphere standard switch. In other cases, the TransportNodeUpdateParameters should NOT be used. When updating transport node you should follow pattern where you should fetch the existing transport node and then only modify the required properties keeping other properties as is. It also modifies attributes of node (host or edge). Note: Previous versions of NSX-T also used a property named transport_zone_endpoints at TransportNode level. This property is deprecated which creates some combinations of new client along with old client payloads. Examples [1] shows old/existing client request and response by populating transport_zone_endpoints property at TransportNode level. Example [2] shows TransportNode updating TransportNode from exmaple [1] request/response by adding a new StandardHostSwitch by populating transport_zone_endpoints at StandardHostSwitch level. TransportNode level transport_zone_endpoints will ONLY have TransportZoneEndpoints that were originally specified here during create/update operation and does not include TransportZoneEndpoints that were directly specified at StandardHostSwitch level. This api is now deprecated. Please use new api - /infra/sites/<site-id>/enforcement-points/<enforcementpoint-id>/host-transport-nodes/<host-transport-node-id>
	//
	// @param transportNodeIdParam (required)
	// @param transportNodeParam (required)
	// @param esxMgmtIfMigrationDestParam The network ids to which the ESX vmk interfaces will be migrated (optional)
	// @param ifIdParam The ESX vmk interfaces to migrate (optional)
	// @param pingIpParam IP Addresses to ping right after ESX vmk interfaces were migrated. (optional)
	// @param skipValidationParam Whether to skip front-end validation for vmk/vnic/pnic migration (optional, default to false)
	// @param vnicParam The ESX vmk interfaces and/or VM NIC to migrate (optional)
	// @param vnicMigrationDestParam The migration destinations of ESX vmk interfaces and/or VM NIC (optional)
	// @return com.vmware.nsx.model.TransportNode
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(transportNodeIdParam string, transportNodeParam model.TransportNode, esxMgmtIfMigrationDestParam *string, ifIdParam *string, pingIpParam *string, skipValidationParam *bool, vnicParam *string, vnicMigrationDestParam *string) (model.TransportNode, error)

	// Put transport node into maintenance mode or exit from maintenance mode.
	//
	// @param transportnodeIdParam (required)
	// @param actionParam (optional)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Updatemaintenancemode(transportnodeIdParam string, actionParam *string) error
}

type transportNodesClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewTransportNodesClient(connector client.Connector) *transportNodesClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.transport_nodes")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"cleanstaleentries":        core.NewMethodIdentifier(interfaceIdentifier, "cleanstaleentries"),
		"create":                   core.NewMethodIdentifier(interfaceIdentifier, "create"),
		"delete":                   core.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"deleteontransportnode":    core.NewMethodIdentifier(interfaceIdentifier, "deleteontransportnode"),
		"disableflowcache":         core.NewMethodIdentifier(interfaceIdentifier, "disableflowcache"),
		"enableflowcache":          core.NewMethodIdentifier(interfaceIdentifier, "enableflowcache"),
		"get":                      core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"getontransportnode":       core.NewMethodIdentifier(interfaceIdentifier, "getontransportnode"),
		"list":                     core.NewMethodIdentifier(interfaceIdentifier, "list"),
		"migratetovds":             core.NewMethodIdentifier(interfaceIdentifier, "migratetovds"),
		"postontransportnode":      core.NewMethodIdentifier(interfaceIdentifier, "postontransportnode"),
		"putontransportnode":       core.NewMethodIdentifier(interfaceIdentifier, "putontransportnode"),
		"redeploy":                 core.NewMethodIdentifier(interfaceIdentifier, "redeploy"),
		"refreshnodeconfiguration": core.NewMethodIdentifier(interfaceIdentifier, "refreshnodeconfiguration"),
		"restartinventorysync":     core.NewMethodIdentifier(interfaceIdentifier, "restartinventorysync"),
		"restoreclusterconfig":     core.NewMethodIdentifier(interfaceIdentifier, "restoreclusterconfig"),
		"resynchostconfig":         core.NewMethodIdentifier(interfaceIdentifier, "resynchostconfig"),
		"update":                   core.NewMethodIdentifier(interfaceIdentifier, "update"),
		"updatemaintenancemode":    core.NewMethodIdentifier(interfaceIdentifier, "updatemaintenancemode"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	tIface := transportNodesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &tIface
}

func (tIface *transportNodesClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := tIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (tIface *transportNodesClient) Cleanstaleentries() error {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(transportNodesCleanstaleentriesInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := transportNodesCleanstaleentriesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_nodes", "cleanstaleentries", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (tIface *transportNodesClient) Create(transportNodeParam model.TransportNode) (model.TransportNode, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(transportNodesCreateInputType(), typeConverter)
	sv.AddStructField("TransportNode", transportNodeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.TransportNode
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := transportNodesCreateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_nodes", "create", inputDataValue, executionContext)
	var emptyOutput model.TransportNode
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), transportNodesCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.TransportNode), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *transportNodesClient) Delete(transportNodeIdParam string, forceParam *bool, unprepareHostParam *bool) error {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(transportNodesDeleteInputType(), typeConverter)
	sv.AddStructField("TransportNodeId", transportNodeIdParam)
	sv.AddStructField("Force", forceParam)
	sv.AddStructField("UnprepareHost", unprepareHostParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := transportNodesDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_nodes", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (tIface *transportNodesClient) Deleteontransportnode(targetNodeIdParam string, targetUriParam string) error {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(transportNodesDeleteontransportnodeInputType(), typeConverter)
	sv.AddStructField("TargetNodeId", targetNodeIdParam)
	sv.AddStructField("TargetUri", targetUriParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := transportNodesDeleteontransportnodeRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_nodes", "deleteontransportnode", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (tIface *transportNodesClient) Disableflowcache(transportNodeIdParam string) error {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(transportNodesDisableflowcacheInputType(), typeConverter)
	sv.AddStructField("TransportNodeId", transportNodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := transportNodesDisableflowcacheRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_nodes", "disableflowcache", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (tIface *transportNodesClient) Enableflowcache(transportNodeIdParam string) error {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(transportNodesEnableflowcacheInputType(), typeConverter)
	sv.AddStructField("TransportNodeId", transportNodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := transportNodesEnableflowcacheRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_nodes", "enableflowcache", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (tIface *transportNodesClient) Get(transportNodeIdParam string) (model.TransportNode, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(transportNodesGetInputType(), typeConverter)
	sv.AddStructField("TransportNodeId", transportNodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.TransportNode
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := transportNodesGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_nodes", "get", inputDataValue, executionContext)
	var emptyOutput model.TransportNode
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), transportNodesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.TransportNode), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *transportNodesClient) Getontransportnode(targetNodeIdParam string, targetUriParam string) error {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(transportNodesGetontransportnodeInputType(), typeConverter)
	sv.AddStructField("TargetNodeId", targetNodeIdParam)
	sv.AddStructField("TargetUri", targetUriParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := transportNodesGetontransportnodeRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_nodes", "getontransportnode", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (tIface *transportNodesClient) List(cursorParam *string, inMaintenanceModeParam *bool, includedFieldsParam *string, nodeIdParam *string, nodeIpParam *string, nodeTypesParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, transportZoneIdParam *string) (model.TransportNodeListResult, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(transportNodesListInputType(), typeConverter)
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
		var emptyOutput model.TransportNodeListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := transportNodesListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_nodes", "list", inputDataValue, executionContext)
	var emptyOutput model.TransportNodeListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), transportNodesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.TransportNodeListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *transportNodesClient) Migratetovds(transportNodeIdParam string, skipMaintmodeParam *bool) error {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(transportNodesMigratetovdsInputType(), typeConverter)
	sv.AddStructField("TransportNodeId", transportNodeIdParam)
	sv.AddStructField("SkipMaintmode", skipMaintmodeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := transportNodesMigratetovdsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_nodes", "migratetovds", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (tIface *transportNodesClient) Postontransportnode(targetNodeIdParam string, targetUriParam string) error {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(transportNodesPostontransportnodeInputType(), typeConverter)
	sv.AddStructField("TargetNodeId", targetNodeIdParam)
	sv.AddStructField("TargetUri", targetUriParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := transportNodesPostontransportnodeRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_nodes", "postontransportnode", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (tIface *transportNodesClient) Putontransportnode(targetNodeIdParam string, targetUriParam string) error {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(transportNodesPutontransportnodeInputType(), typeConverter)
	sv.AddStructField("TargetNodeId", targetNodeIdParam)
	sv.AddStructField("TargetUri", targetUriParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := transportNodesPutontransportnodeRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_nodes", "putontransportnode", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (tIface *transportNodesClient) Redeploy(nodeIdParam string, transportNodeParam model.TransportNode) (model.TransportNode, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(transportNodesRedeployInputType(), typeConverter)
	sv.AddStructField("NodeId", nodeIdParam)
	sv.AddStructField("TransportNode", transportNodeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.TransportNode
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := transportNodesRedeployRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_nodes", "redeploy", inputDataValue, executionContext)
	var emptyOutput model.TransportNode
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), transportNodesRedeployOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.TransportNode), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *transportNodesClient) Refreshnodeconfiguration(transportNodeIdParam string, readOnlyParam *bool) error {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(transportNodesRefreshnodeconfigurationInputType(), typeConverter)
	sv.AddStructField("TransportNodeId", transportNodeIdParam)
	sv.AddStructField("ReadOnly", readOnlyParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := transportNodesRefreshnodeconfigurationRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_nodes", "refreshnodeconfiguration", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (tIface *transportNodesClient) Restartinventorysync(transportNodeIdParam string) error {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(transportNodesRestartinventorysyncInputType(), typeConverter)
	sv.AddStructField("TransportNodeId", transportNodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := transportNodesRestartinventorysyncRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_nodes", "restartinventorysync", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (tIface *transportNodesClient) Restoreclusterconfig(transportNodeIdParam string) error {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(transportNodesRestoreclusterconfigInputType(), typeConverter)
	sv.AddStructField("TransportNodeId", transportNodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := transportNodesRestoreclusterconfigRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_nodes", "restoreclusterconfig", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (tIface *transportNodesClient) Resynchostconfig(transportnodeIdParam string) error {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(transportNodesResynchostconfigInputType(), typeConverter)
	sv.AddStructField("TransportnodeId", transportnodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := transportNodesResynchostconfigRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_nodes", "resynchostconfig", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (tIface *transportNodesClient) Update(transportNodeIdParam string, transportNodeParam model.TransportNode, esxMgmtIfMigrationDestParam *string, ifIdParam *string, pingIpParam *string, skipValidationParam *bool, vnicParam *string, vnicMigrationDestParam *string) (model.TransportNode, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(transportNodesUpdateInputType(), typeConverter)
	sv.AddStructField("TransportNodeId", transportNodeIdParam)
	sv.AddStructField("TransportNode", transportNodeParam)
	sv.AddStructField("EsxMgmtIfMigrationDest", esxMgmtIfMigrationDestParam)
	sv.AddStructField("IfId", ifIdParam)
	sv.AddStructField("PingIp", pingIpParam)
	sv.AddStructField("SkipValidation", skipValidationParam)
	sv.AddStructField("Vnic", vnicParam)
	sv.AddStructField("VnicMigrationDest", vnicMigrationDestParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.TransportNode
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := transportNodesUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_nodes", "update", inputDataValue, executionContext)
	var emptyOutput model.TransportNode
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), transportNodesUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.TransportNode), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *transportNodesClient) Updatemaintenancemode(transportnodeIdParam string, actionParam *string) error {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(transportNodesUpdatemaintenancemodeInputType(), typeConverter)
	sv.AddStructField("TransportnodeId", transportnodeIdParam)
	sv.AddStructField("Action", actionParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := transportNodesUpdatemaintenancemodeRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_nodes", "updatemaintenancemode", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
