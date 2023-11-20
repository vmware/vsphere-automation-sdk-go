// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: TopologyByCluster
// Used by client-side stubs.

package nvds_urt

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type TopologyByClusterClient interface {

	// Set VDS configuration and create it in vCenter
	//  NVDS-to-CVDS Upgrade Readiness Tool feature won't be supported
	//
	// Deprecated: This API element is deprecated.
	//
	// @param precheckIdParam (required)
	// @param upgradeTopologyParam (required)
	// @param clusterIdParam cluster identifier (optional)
	// @param useRecommendedTopologyConfigParam Flag to indicate if use recommended topology got from the latest precheck (optional)
	// @return com.vmware.nsx_policy.model.UpgradeTopology
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Apply(precheckIdParam string, upgradeTopologyParam nsx_policyModel.UpgradeTopology, clusterIdParam *string, useRecommendedTopologyConfigParam *bool) (nsx_policyModel.UpgradeTopology, error)

	// Recommmended topology
	//  NVDS-to-CVDS Upgrade Readiness Tool feature won't be supported
	//
	// Deprecated: This API element is deprecated.
	//
	// @param precheckIdParam (required)
	// @param clusterIdParam cluster identifier (optional)
	// @param computeManagerIdParam vCenter identifier (optional)
	// @param showVdsConfigParam Flag to indicate if VdsTopology should contain VDS configuration (optional)
	// @return com.vmware.nsx_policy.model.UpgradeTopology
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(precheckIdParam string, clusterIdParam *string, computeManagerIdParam *string, showVdsConfigParam *bool) (nsx_policyModel.UpgradeTopology, error)
}

type topologyByClusterClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewTopologyByClusterClient(connector vapiProtocolClient_.Connector) *topologyByClusterClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.nvds_urt.topology_by_cluster")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"apply": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "apply"),
		"get":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	tIface := topologyByClusterClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &tIface
}

func (tIface *topologyByClusterClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := tIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (tIface *topologyByClusterClient) Apply(precheckIdParam string, upgradeTopologyParam nsx_policyModel.UpgradeTopology, clusterIdParam *string, useRecommendedTopologyConfigParam *bool) (nsx_policyModel.UpgradeTopology, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := topologyByClusterApplyRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(topologyByClusterApplyInputType(), typeConverter)
	sv.AddStructField("PrecheckId", precheckIdParam)
	sv.AddStructField("UpgradeTopology", upgradeTopologyParam)
	sv.AddStructField("ClusterId", clusterIdParam)
	sv.AddStructField("UseRecommendedTopologyConfig", useRecommendedTopologyConfigParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.UpgradeTopology
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.nvds_urt.topology_by_cluster", "apply", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.UpgradeTopology
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), TopologyByClusterApplyOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.UpgradeTopology), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *topologyByClusterClient) Get(precheckIdParam string, clusterIdParam *string, computeManagerIdParam *string, showVdsConfigParam *bool) (nsx_policyModel.UpgradeTopology, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := topologyByClusterGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(topologyByClusterGetInputType(), typeConverter)
	sv.AddStructField("PrecheckId", precheckIdParam)
	sv.AddStructField("ClusterId", clusterIdParam)
	sv.AddStructField("ComputeManagerId", computeManagerIdParam)
	sv.AddStructField("ShowVdsConfig", showVdsConfigParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.UpgradeTopology
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.nvds_urt.topology_by_cluster", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.UpgradeTopology
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), TopologyByClusterGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.UpgradeTopology), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
