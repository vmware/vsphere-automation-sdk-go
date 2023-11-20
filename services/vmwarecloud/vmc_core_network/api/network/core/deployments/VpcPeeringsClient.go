// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: VpcPeerings
// Used by client-side stubs.

package deployments

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	vmwarecloudVmc_core_networkModel "github.com/vmware/vsphere-automation-sdk-go/services/vmwarecloud/vmc_core_network/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type VpcPeeringsClient interface {

	// Get All VPC peering of SDDC with current status
	//
	// @param orgIdParam organization identifier
	// @param deploymentIdParam deployment identifier
	// @return OK
	//
	// @throws InvalidRequest Bad Request
	// @throws Unauthenticated Unauthorized
	// @throws Unauthorized Forbidden
	// @throws NotFound Cannot find the vpc peering with given deployment Id
	GetAllVpcPeering(orgIdParam string, deploymentIdParam string) ([]vmwarecloudVmc_core_networkModel.VpcPeeringInfo, error)

	// Create VPC Peering DB object
	//
	// @param orgIdParam organization identifier
	// @param deploymentIdParam deployment identifier
	// @param vpcPeeringParam Vpc Peering DB Object
	// @return OK
	//
	// @throws Unauthenticated Unauthorized
	// @throws Unauthorized Forbidden
	CreateVpcPeeringEntity(orgIdParam string, deploymentIdParam string, vpcPeeringParam vmwarecloudVmc_core_networkModel.VpcPeering) (vmwarecloudVmc_core_networkModel.VpcPeering, error)

	// Get VPC peering of SDDC with given identifier
	//
	// @param orgIdParam organization identifier
	// @param deploymentIdParam deployment identifier
	// @param idParam resource identifier
	// @return OK
	//
	// @throws InvalidRequest Bad Request
	// @throws Unauthenticated Unauthorized
	// @throws Unauthorized Forbidden
	// @throws NotFound Cannot find the deployment with given identifier
	GetVpcPeeringById(orgIdParam string, deploymentIdParam string, idParam string) (vmwarecloudVmc_core_networkModel.VpcPeeringInfo, error)

	// Delete VPC Peering DB Object of given id.
	//
	// @param orgIdParam organization identifier
	// @param deploymentIdParam deployment identifier
	// @param idParam resource identifier
	// @return OK
	//
	// @throws Unauthenticated Unauthorized
	// @throws Unauthorized Forbidden
	DeleteVpcPeeringEntity(orgIdParam string, deploymentIdParam string, idParam string) (bool, error)

	// Patch VPC Peering DB object with given id.
	//
	// @param orgIdParam organization identifier
	// @param deploymentIdParam deployment identifier
	// @param idParam resource identifier
	// @param vpcPeeringParam VPC Peering DB object with updated properties.
	// @return OK
	//
	// @throws Unauthenticated Unauthorized
	// @throws Unauthorized Forbidden
	// @throws NotFound Not Found
	UpdateVpcPeeringEntity(orgIdParam string, deploymentIdParam string, idParam string, vpcPeeringParam vmwarecloudVmc_core_networkModel.VpcPeering) (vmwarecloudVmc_core_networkModel.VpcPeering, error)
}

type vpcPeeringsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewVpcPeeringsClient(connector vapiProtocolClient_.Connector) *vpcPeeringsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vmwarecloud.vmc_core_network.api.network.core.deployments.vpc_peerings")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get_all_vpc_peering":       vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get_all_vpc_peering"),
		"create_vpc_peering_entity": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "create_vpc_peering_entity"),
		"get_vpc_peering_by_id":     vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get_vpc_peering_by_id"),
		"delete_vpc_peering_entity": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete_vpc_peering_entity"),
		"update_vpc_peering_entity": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update_vpc_peering_entity"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	vIface := vpcPeeringsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &vIface
}

func (vIface *vpcPeeringsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := vIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (vIface *vpcPeeringsClient) GetAllVpcPeering(orgIdParam string, deploymentIdParam string) ([]vmwarecloudVmc_core_networkModel.VpcPeeringInfo, error) {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	operationRestMetaData := vpcPeeringsGetAllVpcPeeringRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(vpcPeeringsGetAllVpcPeeringInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("DeploymentId", deploymentIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []vmwarecloudVmc_core_networkModel.VpcPeeringInfo
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.vmwarecloud.vmc_core_network.api.network.core.deployments.vpc_peerings", "get_all_vpc_peering", inputDataValue, executionContext)
	var emptyOutput []vmwarecloudVmc_core_networkModel.VpcPeeringInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), VpcPeeringsGetAllVpcPeeringOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.([]vmwarecloudVmc_core_networkModel.VpcPeeringInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (vIface *vpcPeeringsClient) CreateVpcPeeringEntity(orgIdParam string, deploymentIdParam string, vpcPeeringParam vmwarecloudVmc_core_networkModel.VpcPeering) (vmwarecloudVmc_core_networkModel.VpcPeering, error) {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	operationRestMetaData := vpcPeeringsCreateVpcPeeringEntityRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(vpcPeeringsCreateVpcPeeringEntityInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("DeploymentId", deploymentIdParam)
	sv.AddStructField("VpcPeering", vpcPeeringParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmwarecloudVmc_core_networkModel.VpcPeering
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.vmwarecloud.vmc_core_network.api.network.core.deployments.vpc_peerings", "create_vpc_peering_entity", inputDataValue, executionContext)
	var emptyOutput vmwarecloudVmc_core_networkModel.VpcPeering
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), VpcPeeringsCreateVpcPeeringEntityOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmwarecloudVmc_core_networkModel.VpcPeering), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (vIface *vpcPeeringsClient) GetVpcPeeringById(orgIdParam string, deploymentIdParam string, idParam string) (vmwarecloudVmc_core_networkModel.VpcPeeringInfo, error) {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	operationRestMetaData := vpcPeeringsGetVpcPeeringByIdRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(vpcPeeringsGetVpcPeeringByIdInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("DeploymentId", deploymentIdParam)
	sv.AddStructField("Id", idParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmwarecloudVmc_core_networkModel.VpcPeeringInfo
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.vmwarecloud.vmc_core_network.api.network.core.deployments.vpc_peerings", "get_vpc_peering_by_id", inputDataValue, executionContext)
	var emptyOutput vmwarecloudVmc_core_networkModel.VpcPeeringInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), VpcPeeringsGetVpcPeeringByIdOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmwarecloudVmc_core_networkModel.VpcPeeringInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (vIface *vpcPeeringsClient) DeleteVpcPeeringEntity(orgIdParam string, deploymentIdParam string, idParam string) (bool, error) {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	operationRestMetaData := vpcPeeringsDeleteVpcPeeringEntityRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(vpcPeeringsDeleteVpcPeeringEntityInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("DeploymentId", deploymentIdParam)
	sv.AddStructField("Id", idParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput bool
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.vmwarecloud.vmc_core_network.api.network.core.deployments.vpc_peerings", "delete_vpc_peering_entity", inputDataValue, executionContext)
	var emptyOutput bool
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), VpcPeeringsDeleteVpcPeeringEntityOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(bool), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (vIface *vpcPeeringsClient) UpdateVpcPeeringEntity(orgIdParam string, deploymentIdParam string, idParam string, vpcPeeringParam vmwarecloudVmc_core_networkModel.VpcPeering) (vmwarecloudVmc_core_networkModel.VpcPeering, error) {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	operationRestMetaData := vpcPeeringsUpdateVpcPeeringEntityRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(vpcPeeringsUpdateVpcPeeringEntityInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("DeploymentId", deploymentIdParam)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("VpcPeering", vpcPeeringParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmwarecloudVmc_core_networkModel.VpcPeering
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.vmwarecloud.vmc_core_network.api.network.core.deployments.vpc_peerings", "update_vpc_peering_entity", inputDataValue, executionContext)
	var emptyOutput vmwarecloudVmc_core_networkModel.VpcPeering
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), VpcPeeringsUpdateVpcPeeringEntityOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmwarecloudVmc_core_networkModel.VpcPeering), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
