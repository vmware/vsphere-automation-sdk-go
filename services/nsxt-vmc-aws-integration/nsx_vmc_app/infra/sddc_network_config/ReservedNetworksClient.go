// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ReservedNetworks
// Used by client-side stubs.

package sddc_network_config

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_vmc_appModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-vmc-aws-integration/nsx_vmc_app/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type ReservedNetworksClient interface {

	// Delete a reserved CIDR block, also remove this CIDR block from CGW segments creation constraint.
	//
	// @param reservedNetworkIdParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(reservedNetworkIdParam string) error

	// Get a reserved CIDR block by ID. Once a CIDR block is reserved, CGW segments and mgmt CIDRs cannot overlap with it.
	//
	// @param reservedNetworkIdParam (required)
	// @return com.vmware.nsx_vmc_app.model.ReservedCIDRBlock
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(reservedNetworkIdParam string) (nsx_vmc_appModel.ReservedCIDRBlock, error)

	// List all reserved CIDR blocks. Once a CIDR block is reserved, CGW segments and mgmt CIDRs cannot overlap with it.
	// @return com.vmware.nsx_vmc_app.model.ReservedCIDRBlocksListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List() (nsx_vmc_appModel.ReservedCIDRBlocksListResult, error)

	// Add a new reserved CIDR block and add this CIDR block to CGW segments creation constraint. The request will be rejected if the CIDR block's address range overlaps with any existing reserved network's CIDR range.
	//
	// @param reservedNetworkIdParam (required)
	// @param reservedCIDRBlockParam (required)
	// @return com.vmware.nsx_vmc_app.model.ReservedCIDRBlock
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(reservedNetworkIdParam string, reservedCIDRBlockParam nsx_vmc_appModel.ReservedCIDRBlock) (nsx_vmc_appModel.ReservedCIDRBlock, error)

	// This method is used only to check overlap between user inputs and existing reserved CIDR blocks, segments and mgmt CIDRs.
	//
	// @param reservedCIDRBlockParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Validateonly(reservedCIDRBlockParam nsx_vmc_appModel.ReservedCIDRBlock) error
}

type reservedNetworksClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewReservedNetworksClient(connector vapiProtocolClient_.Connector) *reservedNetworksClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_vmc_app.infra.sddc_network_config.reserved_networks")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"delete":       vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":          vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":         vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"update":       vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
		"validateonly": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "validateonly"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	rIface := reservedNetworksClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &rIface
}

func (rIface *reservedNetworksClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := rIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (rIface *reservedNetworksClient) Delete(reservedNetworkIdParam string) error {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	operationRestMetaData := reservedNetworksDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(reservedNetworksDeleteInputType(), typeConverter)
	sv.AddStructField("ReservedNetworkId", reservedNetworkIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx_vmc_app.infra.sddc_network_config.reserved_networks", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (rIface *reservedNetworksClient) Get(reservedNetworkIdParam string) (nsx_vmc_appModel.ReservedCIDRBlock, error) {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	operationRestMetaData := reservedNetworksGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(reservedNetworksGetInputType(), typeConverter)
	sv.AddStructField("ReservedNetworkId", reservedNetworkIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_vmc_appModel.ReservedCIDRBlock
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx_vmc_app.infra.sddc_network_config.reserved_networks", "get", inputDataValue, executionContext)
	var emptyOutput nsx_vmc_appModel.ReservedCIDRBlock
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ReservedNetworksGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_vmc_appModel.ReservedCIDRBlock), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (rIface *reservedNetworksClient) List() (nsx_vmc_appModel.ReservedCIDRBlocksListResult, error) {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	operationRestMetaData := reservedNetworksListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(reservedNetworksListInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_vmc_appModel.ReservedCIDRBlocksListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx_vmc_app.infra.sddc_network_config.reserved_networks", "list", inputDataValue, executionContext)
	var emptyOutput nsx_vmc_appModel.ReservedCIDRBlocksListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ReservedNetworksListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_vmc_appModel.ReservedCIDRBlocksListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (rIface *reservedNetworksClient) Update(reservedNetworkIdParam string, reservedCIDRBlockParam nsx_vmc_appModel.ReservedCIDRBlock) (nsx_vmc_appModel.ReservedCIDRBlock, error) {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	operationRestMetaData := reservedNetworksUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(reservedNetworksUpdateInputType(), typeConverter)
	sv.AddStructField("ReservedNetworkId", reservedNetworkIdParam)
	sv.AddStructField("ReservedCIDRBlock", reservedCIDRBlockParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_vmc_appModel.ReservedCIDRBlock
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx_vmc_app.infra.sddc_network_config.reserved_networks", "update", inputDataValue, executionContext)
	var emptyOutput nsx_vmc_appModel.ReservedCIDRBlock
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ReservedNetworksUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_vmc_appModel.ReservedCIDRBlock), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (rIface *reservedNetworksClient) Validateonly(reservedCIDRBlockParam nsx_vmc_appModel.ReservedCIDRBlock) error {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	operationRestMetaData := reservedNetworksValidateonlyRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(reservedNetworksValidateonlyInputType(), typeConverter)
	sv.AddStructField("ReservedCIDRBlock", reservedCIDRBlockParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx_vmc_app.infra.sddc_network_config.reserved_networks", "validateonly", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
