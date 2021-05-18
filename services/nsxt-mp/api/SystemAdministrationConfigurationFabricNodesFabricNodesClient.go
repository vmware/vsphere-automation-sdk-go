// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: SystemAdministrationConfigurationFabricNodesFabricNodes
// Used by client-side stubs.

package api

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-mp/model"
)

const _ = core.SupportedByRuntimeVersion1

type SystemAdministrationConfigurationFabricNodesFabricNodesClient interface {

	// Creates a host node (hypervisor) or edge node (router) in the transport network. When you run this command for a host, NSX Manager attempts to install the NSX kernel modules, which are packaged as VIB, RPM, or DEB files. For the installation to succeed, you must provide the host login credentials and the host thumbprint. To get the ESXi host thumbprint, SSH to the host and run the **openssl x509 -in /etc/vmware/ssl/rui.crt -fingerprint -sha256 -noout** command. To generate host key thumbprint using SHA-256 algorithm please follow the steps below. Log into the host, making sure that the connection is not vulnerable to a man in the middle attack. Check whether a public key already exists. Host public key is generally located at '/etc/ssh/ssh_host_rsa_key.pub'. If the key is not present then generate a new key by running the following command and follow the instructions. **ssh-keygen -t rsa** Now generate a SHA256 hash of the key using the following command. Please make sure to pass the appropriate file name if the public key is stored with a different file name other than the default 'id_rsa.pub'. **awk '{print $2}' id_rsa.pub | base64 -d | sha256sum -b | sed 's/ .\*$//' | xxd -r -p | base64** This api is deprecated as part of FN+TN unification. Please use Transport Node API POST /transport-nodes to install NSX components on a node.
	//
	// @param nodeParam (required)
	// The parameter must contain all the properties defined in model.Node.
	// @return com.vmware.model.Node
	// The return value will contain all the properties defined in model.Node.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	AddNode(nodeParam *data.StructValue) (*data.StructValue, error)

	//
	//
	// @param nodeIdParam (required)
	// @param unprepareHostParam Delete a host and uninstall NSX components (optional, default to true)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteNode(nodeIdParam string, unprepareHostParam *bool) error

	// Get the module details of a Fabric Node This api is deprecated, use Transport Node API GET /transport-nodes/<transportnode-id>/modules to get fabric node modules.
	//
	// @param nodeIdParam (required)
	// @return com.vmware.model.SoftwareModuleResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetFabricNodeModules(nodeIdParam string) (model.SoftwareModuleResult, error)

	// For edge nodes, returns the current install state when deployment is in progress, NODE_READY when deployment is complete and the failure state when deployment has failed. This api is deprecated. Please use /transport-nodes/<transportnode-id>/state to get realized state of a Fabric Node.
	//
	// @param nodeIdParam (required)
	// @return com.vmware.model.ConfigurationState
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetFabricNodeState(nodeIdParam string) (model.ConfigurationState, error)

	// Returns names of all supported host OS.
	// @return com.vmware.model.SupportedHostOSListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetSupportedHostOSTypes() (model.SupportedHostOSListResult, error)

	// Invoke DELETE request on target fabric node. This api is deprecated as part of FN+TN unification. Please use Transport Node API DELETE /transport-nodes/<transport-node-id>/<target-node-id>/<target-uri>
	//
	// @param targetNodeIdParam Target node UUID (required)
	// @param targetUriParam URI of API to invoke on target node (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	InvokeDeleteFabricCentralAPI(targetNodeIdParam string, targetUriParam string) error

	// Invoke GET request on target fabric node. This api is deprecated as part of FN+TN unification. Please use Transport Node API GET /transport-nodes/<transport-node-id>/<target-node-id>/<target-uri>
	//
	// @param targetNodeIdParam Target node UUID (required)
	// @param targetUriParam URI of API to invoke on target node (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	InvokeGetFabricCentralAPI(targetNodeIdParam string, targetUriParam string) error

	// Invoke POST request on target fabric node. This api is deprecated as part of FN+TN unification. Please use Transport Node API POST /transport-nodes/<transport-node-id>/<target-node-id>/<target-uri>
	//
	// @param targetNodeIdParam Target node UUID (required)
	// @param targetUriParam URI of API to invoke on target node (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	InvokePostFabricCentralAPI(targetNodeIdParam string, targetUriParam string) error

	// Invoke PUT request on target fabric node. This api is deprecated as part of FN+TN unification. Please use Transport Node API PUT /transport-nodes/<transport-node-id>/<target-node-id>/<target-uri>
	//
	// @param targetNodeIdParam Target node UUID (required)
	// @param targetUriParam URI of API to invoke on target node (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	InvokePutFabricCentralAPI(targetNodeIdParam string, targetUriParam string) error

	// Returns information about capabilities of a single fabric host node. Edge nodes do not have capabilities. This api is deprecated, use GET /transport-nodes/<transportnode-id>/capabilities if FN is converted to TN.
	//
	// @param nodeIdParam (required)
	// @return com.vmware.model.NodeCapabilitiesResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListNodeCapabilities(nodeIdParam string) (model.NodeCapabilitiesResult, error)

	// Returns information about all fabric nodes (hosts and edges). This api is deprecated as part of FN+TN unification. Please use Transport Node API GET /transport-nodes to list all fabric nodes.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param discoveredNodeIdParam Id of the discovered node which was converted to create this node (optional)
	// @param displayNameParam HostNode display name (optional)
	// @param externalIdParam HostNode external id (optional)
	// @param hardwareIdParam Hardware Id of the host (optional)
	// @param hypervisorOsTypeParam HostNode's Hypervisor type, for example ESXi, RHEL KVM or UBUNTU KVM. (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param ipAddressParam Management IP address of the node (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param resourceTypeParam Node type from 'HostNode', 'EdgeNode', 'PublicCloudGatewayNode' (optional)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.model.NodeListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListNodes(cursorParam *string, discoveredNodeIdParam *string, displayNameParam *string, externalIdParam *string, hardwareIdParam *string, hypervisorOsTypeParam *string, includedFieldsParam *string, ipAddressParam *string, pageSizeParam *int64, resourceTypeParam *string, sortAscendingParam *bool, sortByParam *string) (model.NodeListResult, error)

	// Perform a service deployment upgrade on a host node
	//
	// @param nodeIdParam (required)
	// @param disableVmMigrationParam Should VM migration be disabled during upgrade (optional, default to false)
	// @return com.vmware.model.Node
	// The return value will contain all the properties defined in model.Node.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	PerformHostNodeUpgradeActionUpgradeInfra(nodeIdParam string, disableVmMigrationParam *bool) (*data.StructValue, error)

	// The supported fabric node actions are enter_maintenance_mode, exit_maintenance_mode for EdgeNode. This API is deprecated, please call TransportNode maintenance mode API to update maintenance mode, refer to \"Update transport node maintenance mode\".
	//
	// @param nodeIdParam (required)
	// @param actionParam Supported fabric node actions (optional)
	// @param evacuatePoweredOffVmsParam Evacuate powered-off vms (optional, default to false)
	// @param vsanModeParam Vsan decommission mode (optional, default to ensure_object_accessibility)
	// @return com.vmware.model.Node
	// The return value will contain all the properties defined in model.Node.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	PerformNodeAction(nodeIdParam string, actionParam *string, evacuatePoweredOffVmsParam *bool, vsanModeParam *string) (*data.StructValue, error)

	// Returns information about a specific fabric node (host or edge). This api is deprecated, use Transport Node API GET /transport-nodes/<transport-node-id> to get fabric node information.
	//
	// @param nodeIdParam (required)
	// @return com.vmware.model.Node
	// The return value will contain all the properties defined in model.Node.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ReadNode(nodeIdParam string) (*data.StructValue, error)

	// Returns connectivity, heartbeat, and version information about a fabric node (host or edge). Note that the LCP connectivity status remains down until after the fabric node has been added as a transpot node and the NSX host switch has been successfully installed. See POST /api/v1/transport-nodes. This api is deprecated, use GET /api/v1/transport-nodes/<node-id>/status to get status information of a node with constraint FN is converted to TN.
	//
	// @param nodeIdParam (required)
	// @param sourceParam Data source type. (optional)
	// @return com.vmware.model.NodeStatus
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ReadNodeStatus(nodeIdParam string, sourceParam *string) (model.NodeStatus, error)

	// Returns connectivity, heartbeat, and version information about all fabric nodes (host or edge). This api is deprecated as part of FN+TN unification. Please use Transport Node Status API /transport-nodes/<node-id>/status to get status information of a node and to get all transport nodes ids use GET /transport-nodes.
	//
	// @param nodeIdsParam List of requested Nodes. (required)
	// @return com.vmware.model.NodeStatusListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ReadNodesStatus(nodeIdsParam string) (model.NodeStatusListResult, error)

	// Restart the inventory sync for the node if it is currently internally paused. After this action the next inventory sync coming from the node is processed. This api is deprecated as part of FN+TN unification. Please use Transport Node API POST /transport-nodes/<transport-node-id>?action=restart_inventory_sync to restart inventory sync of node.
	//
	// @param nodeIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	RestartInventorySyncRestartInventorySync(nodeIdParam string) error

	//
	//
	// @param nodeIdParam (required)
	// @param nodeParam (required)
	// The parameter must contain all the properties defined in model.Node.
	// @return com.vmware.model.Node
	// The return value will contain all the properties defined in model.Node.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateNode(nodeIdParam string, nodeParam *data.StructValue) (*data.StructValue, error)
}

type systemAdministrationConfigurationFabricNodesFabricNodesClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSystemAdministrationConfigurationFabricNodesFabricNodesClient(connector client.Connector) *systemAdministrationConfigurationFabricNodesFabricNodesClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.system_administration_configuration_fabric_nodes_fabric_nodes")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"add_node":                                       core.NewMethodIdentifier(interfaceIdentifier, "add_node"),
		"delete_node":                                    core.NewMethodIdentifier(interfaceIdentifier, "delete_node"),
		"get_fabric_node_modules":                        core.NewMethodIdentifier(interfaceIdentifier, "get_fabric_node_modules"),
		"get_fabric_node_state":                          core.NewMethodIdentifier(interfaceIdentifier, "get_fabric_node_state"),
		"get_supported_host_OS_types":                    core.NewMethodIdentifier(interfaceIdentifier, "get_supported_host_OS_types"),
		"invoke_delete_fabric_central_API":               core.NewMethodIdentifier(interfaceIdentifier, "invoke_delete_fabric_central_API"),
		"invoke_get_fabric_central_API":                  core.NewMethodIdentifier(interfaceIdentifier, "invoke_get_fabric_central_API"),
		"invoke_post_fabric_central_API":                 core.NewMethodIdentifier(interfaceIdentifier, "invoke_post_fabric_central_API"),
		"invoke_put_fabric_central_API":                  core.NewMethodIdentifier(interfaceIdentifier, "invoke_put_fabric_central_API"),
		"list_node_capabilities":                         core.NewMethodIdentifier(interfaceIdentifier, "list_node_capabilities"),
		"list_nodes":                                     core.NewMethodIdentifier(interfaceIdentifier, "list_nodes"),
		"perform_host_node_upgrade_action_upgrade_infra": core.NewMethodIdentifier(interfaceIdentifier, "perform_host_node_upgrade_action_upgrade_infra"),
		"perform_node_action":                            core.NewMethodIdentifier(interfaceIdentifier, "perform_node_action"),
		"read_node":                                      core.NewMethodIdentifier(interfaceIdentifier, "read_node"),
		"read_node_status":                               core.NewMethodIdentifier(interfaceIdentifier, "read_node_status"),
		"read_nodes_status":                              core.NewMethodIdentifier(interfaceIdentifier, "read_nodes_status"),
		"restart_inventory_sync_restart_inventory_sync":  core.NewMethodIdentifier(interfaceIdentifier, "restart_inventory_sync_restart_inventory_sync"),
		"update_node":                                    core.NewMethodIdentifier(interfaceIdentifier, "update_node"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := systemAdministrationConfigurationFabricNodesFabricNodesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *systemAdministrationConfigurationFabricNodesFabricNodesClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *systemAdministrationConfigurationFabricNodesFabricNodesClient) AddNode(nodeParam *data.StructValue) (*data.StructValue, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesFabricNodesAddNodeInputType(), typeConverter)
	sv.AddStructField("Node", nodeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesFabricNodesAddNodeRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_fabric_nodes", "add_node", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricNodesFabricNodesAddNodeOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(*data.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricNodesFabricNodesClient) DeleteNode(nodeIdParam string, unprepareHostParam *bool) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesFabricNodesDeleteNodeInputType(), typeConverter)
	sv.AddStructField("NodeId", nodeIdParam)
	sv.AddStructField("UnprepareHost", unprepareHostParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesFabricNodesDeleteNodeRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_fabric_nodes", "delete_node", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricNodesFabricNodesClient) GetFabricNodeModules(nodeIdParam string) (model.SoftwareModuleResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesFabricNodesGetFabricNodeModulesInputType(), typeConverter)
	sv.AddStructField("NodeId", nodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.SoftwareModuleResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesFabricNodesGetFabricNodeModulesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_fabric_nodes", "get_fabric_node_modules", inputDataValue, executionContext)
	var emptyOutput model.SoftwareModuleResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricNodesFabricNodesGetFabricNodeModulesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.SoftwareModuleResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricNodesFabricNodesClient) GetFabricNodeState(nodeIdParam string) (model.ConfigurationState, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesFabricNodesGetFabricNodeStateInputType(), typeConverter)
	sv.AddStructField("NodeId", nodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ConfigurationState
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesFabricNodesGetFabricNodeStateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_fabric_nodes", "get_fabric_node_state", inputDataValue, executionContext)
	var emptyOutput model.ConfigurationState
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricNodesFabricNodesGetFabricNodeStateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ConfigurationState), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricNodesFabricNodesClient) GetSupportedHostOSTypes() (model.SupportedHostOSListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesFabricNodesGetSupportedHostOSTypesInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.SupportedHostOSListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesFabricNodesGetSupportedHostOSTypesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_fabric_nodes", "get_supported_host_OS_types", inputDataValue, executionContext)
	var emptyOutput model.SupportedHostOSListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricNodesFabricNodesGetSupportedHostOSTypesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.SupportedHostOSListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricNodesFabricNodesClient) InvokeDeleteFabricCentralAPI(targetNodeIdParam string, targetUriParam string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesFabricNodesInvokeDeleteFabricCentralAPIInputType(), typeConverter)
	sv.AddStructField("TargetNodeId", targetNodeIdParam)
	sv.AddStructField("TargetUri", targetUriParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesFabricNodesInvokeDeleteFabricCentralAPIRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_fabric_nodes", "invoke_delete_fabric_central_API", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricNodesFabricNodesClient) InvokeGetFabricCentralAPI(targetNodeIdParam string, targetUriParam string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesFabricNodesInvokeGetFabricCentralAPIInputType(), typeConverter)
	sv.AddStructField("TargetNodeId", targetNodeIdParam)
	sv.AddStructField("TargetUri", targetUriParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesFabricNodesInvokeGetFabricCentralAPIRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_fabric_nodes", "invoke_get_fabric_central_API", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricNodesFabricNodesClient) InvokePostFabricCentralAPI(targetNodeIdParam string, targetUriParam string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesFabricNodesInvokePostFabricCentralAPIInputType(), typeConverter)
	sv.AddStructField("TargetNodeId", targetNodeIdParam)
	sv.AddStructField("TargetUri", targetUriParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesFabricNodesInvokePostFabricCentralAPIRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_fabric_nodes", "invoke_post_fabric_central_API", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricNodesFabricNodesClient) InvokePutFabricCentralAPI(targetNodeIdParam string, targetUriParam string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesFabricNodesInvokePutFabricCentralAPIInputType(), typeConverter)
	sv.AddStructField("TargetNodeId", targetNodeIdParam)
	sv.AddStructField("TargetUri", targetUriParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesFabricNodesInvokePutFabricCentralAPIRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_fabric_nodes", "invoke_put_fabric_central_API", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricNodesFabricNodesClient) ListNodeCapabilities(nodeIdParam string) (model.NodeCapabilitiesResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesFabricNodesListNodeCapabilitiesInputType(), typeConverter)
	sv.AddStructField("NodeId", nodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NodeCapabilitiesResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesFabricNodesListNodeCapabilitiesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_fabric_nodes", "list_node_capabilities", inputDataValue, executionContext)
	var emptyOutput model.NodeCapabilitiesResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricNodesFabricNodesListNodeCapabilitiesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NodeCapabilitiesResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricNodesFabricNodesClient) ListNodes(cursorParam *string, discoveredNodeIdParam *string, displayNameParam *string, externalIdParam *string, hardwareIdParam *string, hypervisorOsTypeParam *string, includedFieldsParam *string, ipAddressParam *string, pageSizeParam *int64, resourceTypeParam *string, sortAscendingParam *bool, sortByParam *string) (model.NodeListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesFabricNodesListNodesInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("DiscoveredNodeId", discoveredNodeIdParam)
	sv.AddStructField("DisplayName", displayNameParam)
	sv.AddStructField("ExternalId", externalIdParam)
	sv.AddStructField("HardwareId", hardwareIdParam)
	sv.AddStructField("HypervisorOsType", hypervisorOsTypeParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("IpAddress", ipAddressParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("ResourceType", resourceTypeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NodeListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesFabricNodesListNodesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_fabric_nodes", "list_nodes", inputDataValue, executionContext)
	var emptyOutput model.NodeListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricNodesFabricNodesListNodesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NodeListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricNodesFabricNodesClient) PerformHostNodeUpgradeActionUpgradeInfra(nodeIdParam string, disableVmMigrationParam *bool) (*data.StructValue, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesFabricNodesPerformHostNodeUpgradeActionUpgradeInfraInputType(), typeConverter)
	sv.AddStructField("NodeId", nodeIdParam)
	sv.AddStructField("DisableVmMigration", disableVmMigrationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesFabricNodesPerformHostNodeUpgradeActionUpgradeInfraRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_fabric_nodes", "perform_host_node_upgrade_action_upgrade_infra", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricNodesFabricNodesPerformHostNodeUpgradeActionUpgradeInfraOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(*data.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricNodesFabricNodesClient) PerformNodeAction(nodeIdParam string, actionParam *string, evacuatePoweredOffVmsParam *bool, vsanModeParam *string) (*data.StructValue, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesFabricNodesPerformNodeActionInputType(), typeConverter)
	sv.AddStructField("NodeId", nodeIdParam)
	sv.AddStructField("Action", actionParam)
	sv.AddStructField("EvacuatePoweredOffVms", evacuatePoweredOffVmsParam)
	sv.AddStructField("VsanMode", vsanModeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesFabricNodesPerformNodeActionRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_fabric_nodes", "perform_node_action", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricNodesFabricNodesPerformNodeActionOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(*data.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricNodesFabricNodesClient) ReadNode(nodeIdParam string) (*data.StructValue, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesFabricNodesReadNodeInputType(), typeConverter)
	sv.AddStructField("NodeId", nodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesFabricNodesReadNodeRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_fabric_nodes", "read_node", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricNodesFabricNodesReadNodeOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(*data.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricNodesFabricNodesClient) ReadNodeStatus(nodeIdParam string, sourceParam *string) (model.NodeStatus, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesFabricNodesReadNodeStatusInputType(), typeConverter)
	sv.AddStructField("NodeId", nodeIdParam)
	sv.AddStructField("Source", sourceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NodeStatus
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesFabricNodesReadNodeStatusRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_fabric_nodes", "read_node_status", inputDataValue, executionContext)
	var emptyOutput model.NodeStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricNodesFabricNodesReadNodeStatusOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NodeStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricNodesFabricNodesClient) ReadNodesStatus(nodeIdsParam string) (model.NodeStatusListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesFabricNodesReadNodesStatusInputType(), typeConverter)
	sv.AddStructField("NodeIds", nodeIdsParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NodeStatusListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesFabricNodesReadNodesStatusRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_fabric_nodes", "read_nodes_status", inputDataValue, executionContext)
	var emptyOutput model.NodeStatusListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricNodesFabricNodesReadNodesStatusOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NodeStatusListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricNodesFabricNodesClient) RestartInventorySyncRestartInventorySync(nodeIdParam string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesFabricNodesRestartInventorySyncRestartInventorySyncInputType(), typeConverter)
	sv.AddStructField("NodeId", nodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesFabricNodesRestartInventorySyncRestartInventorySyncRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_fabric_nodes", "restart_inventory_sync_restart_inventory_sync", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricNodesFabricNodesClient) UpdateNode(nodeIdParam string, nodeParam *data.StructValue) (*data.StructValue, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesFabricNodesUpdateNodeInputType(), typeConverter)
	sv.AddStructField("NodeId", nodeIdParam)
	sv.AddStructField("Node", nodeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesFabricNodesUpdateNodeRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_fabric_nodes", "update_node", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricNodesFabricNodesUpdateNodeOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(*data.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
