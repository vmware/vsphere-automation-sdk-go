// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Topology
// Used by client-side stubs.

package nvds_urt

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type TopologyClient interface {

	// Upon successful preheck status goes to PENDING_TOPOLOGY and global recommended topology is generated which can be retrieved via GetRecommendedVdsTopology API. User can apply the entire recommeneded topology all together or can apply partial depending on which TrasportNodes user wants to be upgraded from NVDS to CVDS. User can change system generated vds_name field, all other fields cannot be changed when applying topology.
	//
	// Deprecated: This API element is deprecated.
	//
	// @param upgradeTopologyParam (required)
	// @param clusterIdParam cluster identifier (optional)
	// @param useRecommendedTopologyConfigParam Flag to indicate if use recommended topology got from the latest precheck (optional)
	// @return com.vmware.nsx.model.UpgradeTopology
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Apply(upgradeTopologyParam nsxModel.UpgradeTopology, clusterIdParam *string, useRecommendedTopologyConfigParam *bool) (nsxModel.UpgradeTopology, error)

	// This returns global recommended topology generated when precheck is successful.
	//
	// Deprecated: This API element is deprecated.
	//
	// @param precheckIdParam (required)
	// @param clusterIdParam cluster identifier (optional)
	// @param computeManagerIdParam vCenter identifier (optional)
	// @param showVdsConfigParam Flag to indicate if VdsTopology should contain VDS configuration (optional)
	// @return com.vmware.nsx.model.UpgradeTopology
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(precheckIdParam string, clusterIdParam *string, computeManagerIdParam *string, showVdsConfigParam *bool) (nsxModel.UpgradeTopology, error)

	// This will revert corresponding VDS to PENDING_TOPOLOGY state. User can revert the entire topology all together or can revert partially depending on which TrasportNodes user does not want to upgrade to VDS.
	//
	// Deprecated: This API element is deprecated.
	//
	// @param upgradeTopologyParam (required)
	// @return com.vmware.nsx.model.UpgradeTopology
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Revert(upgradeTopologyParam nsxModel.UpgradeTopology) (nsxModel.UpgradeTopology, error)
}

type topologyClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewTopologyClient(connector vapiProtocolClient_.Connector) *topologyClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.nvds_urt.topology")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"apply":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "apply"),
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"revert": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "revert"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	tIface := topologyClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &tIface
}

func (tIface *topologyClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := tIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (tIface *topologyClient) Apply(upgradeTopologyParam nsxModel.UpgradeTopology, clusterIdParam *string, useRecommendedTopologyConfigParam *bool) (nsxModel.UpgradeTopology, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := topologyApplyRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(topologyApplyInputType(), typeConverter)
	sv.AddStructField("UpgradeTopology", upgradeTopologyParam)
	sv.AddStructField("ClusterId", clusterIdParam)
	sv.AddStructField("UseRecommendedTopologyConfig", useRecommendedTopologyConfigParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.UpgradeTopology
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.nvds_urt.topology", "apply", inputDataValue, executionContext)
	var emptyOutput nsxModel.UpgradeTopology
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), TopologyApplyOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.UpgradeTopology), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *topologyClient) Get(precheckIdParam string, clusterIdParam *string, computeManagerIdParam *string, showVdsConfigParam *bool) (nsxModel.UpgradeTopology, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := topologyGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(topologyGetInputType(), typeConverter)
	sv.AddStructField("PrecheckId", precheckIdParam)
	sv.AddStructField("ClusterId", clusterIdParam)
	sv.AddStructField("ComputeManagerId", computeManagerIdParam)
	sv.AddStructField("ShowVdsConfig", showVdsConfigParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.UpgradeTopology
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.nvds_urt.topology", "get", inputDataValue, executionContext)
	var emptyOutput nsxModel.UpgradeTopology
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), TopologyGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.UpgradeTopology), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *topologyClient) Revert(upgradeTopologyParam nsxModel.UpgradeTopology) (nsxModel.UpgradeTopology, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := topologyRevertRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(topologyRevertInputType(), typeConverter)
	sv.AddStructField("UpgradeTopology", upgradeTopologyParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.UpgradeTopology
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.nvds_urt.topology", "revert", inputDataValue, executionContext)
	var emptyOutput nsxModel.UpgradeTopology
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), TopologyRevertOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.UpgradeTopology), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
