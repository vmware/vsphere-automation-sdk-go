// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: DiscoveredNodes
// Used by client-side stubs.

package fabric

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = core.SupportedByRuntimeVersion1

type DiscoveredNodesClient interface {

	// NSX components are installaed on host and transport node is created with given configurations.
	//
	// @param siteIdParam (required)
	// @param enforcementpointIdParam (required)
	// @param nodeExtIdParam (required)
	// @param hostTransportNodeParam (required)
	// @return com.vmware.nsx_policy.model.HostTransportNode
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Createtransportnode(siteIdParam string, enforcementpointIdParam string, nodeExtIdParam string, hostTransportNodeParam model.HostTransportNode) (model.HostTransportNode, error)

	// When transport node profile (TNP) is applied to a cluster, if any validation fails (e.g. VMs running on host) then transport node (TN) is not created. In that case after the required action is taken (e.g. VMs powered off), you can call this API to try to create TN for that discovered node. Do not call this API if Transport Node already exists for the discovered node. In that case use API on transport node. /infra/sites/<site-id>/enforcement-points/<enforcementpoint-id>/host-transport-nodes/ <host-transport-node-id/<transport-node-id>?action=restore_cluster_config
	//
	// @param siteIdParam (required)
	// @param enforcementpointIdParam (required)
	// @param nodeExtIdParam (required)
	// @return com.vmware.nsx_policy.model.HostTransportNode
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Reapplyclusterconfig(siteIdParam string, enforcementpointIdParam string, nodeExtIdParam string) (model.HostTransportNode, error)
}

type discoveredNodesClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewDiscoveredNodesClient(connector client.Connector) *discoveredNodesClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.sites.enforcement_points.fabric.discovered_nodes")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"createtransportnode":  core.NewMethodIdentifier(interfaceIdentifier, "createtransportnode"),
		"reapplyclusterconfig": core.NewMethodIdentifier(interfaceIdentifier, "reapplyclusterconfig"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	dIface := discoveredNodesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &dIface
}

func (dIface *discoveredNodesClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := dIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (dIface *discoveredNodesClient) Createtransportnode(siteIdParam string, enforcementpointIdParam string, nodeExtIdParam string, hostTransportNodeParam model.HostTransportNode) (model.HostTransportNode, error) {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(discoveredNodesCreatetransportnodeInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("NodeExtId", nodeExtIdParam)
	sv.AddStructField("HostTransportNode", hostTransportNodeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.HostTransportNode
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := discoveredNodesCreatetransportnodeRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	dIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.fabric.discovered_nodes", "createtransportnode", inputDataValue, executionContext)
	var emptyOutput model.HostTransportNode
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), discoveredNodesCreatetransportnodeOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.HostTransportNode), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (dIface *discoveredNodesClient) Reapplyclusterconfig(siteIdParam string, enforcementpointIdParam string, nodeExtIdParam string) (model.HostTransportNode, error) {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(discoveredNodesReapplyclusterconfigInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("NodeExtId", nodeExtIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.HostTransportNode
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := discoveredNodesReapplyclusterconfigRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	dIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.fabric.discovered_nodes", "reapplyclusterconfig", inputDataValue, executionContext)
	var emptyOutput model.HostTransportNode
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), discoveredNodesReapplyclusterconfigOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.HostTransportNode), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
