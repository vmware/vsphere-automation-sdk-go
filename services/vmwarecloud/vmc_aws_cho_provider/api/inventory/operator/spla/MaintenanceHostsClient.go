// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: MaintenanceHosts
// Used by client-side stubs.

package spla

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	vmwarecloudVmc_aws_cho_providerModel "github.com/vmware/vsphere-automation-sdk-go/services/vmwarecloud/vmc_aws_cho_provider/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type MaintenanceHostsClient interface {

	// Get SPLA maintenance hosts info for given region. Shall be used by operator.
	//
	// @param awsRegionParam AWS Region
	// @return OK
	//
	// @throws Unauthenticated Unauthorized
	// @throws Unauthorized Forbidden
	GetSplaMaintenanceHostsInfo(awsRegionParam string) (vmwarecloudVmc_aws_cho_providerModel.SplaMaintenanceHostsInfo, error)

	// Add maintenance hosts for SPLA info with given region. Shall be used by operator.
	//
	// @param addSplaMaintenanceHostsInfoRequestBodyParam Payload for updating SplaMaintenanceHostsInfo
	// @param awsRegionParam AWS Region
	// @return OK
	//
	// @throws Unauthenticated Unauthorized
	// @throws Unauthorized Forbidden
	AddSplaMaintenanceHostsInfo(addSplaMaintenanceHostsInfoRequestBodyParam []vmwarecloudVmc_aws_cho_providerModel.MaintenanceHost, awsRegionParam string) (vmwarecloudVmc_aws_cho_providerModel.SplaMaintenanceHostsInfo, error)

	// Delete maintenance hosts for SPLA info with given region. Shall be used by operator.
	//
	// @param removeSplaMaintenanceHostsInfoRequestBodyParam Payload for updating SplaMaintenanceHostsInfo
	// @param awsRegionParam AWS Region
	// @return OK
	//
	// @throws Unauthenticated Unauthorized
	// @throws Unauthorized Forbidden
	RemoveSplaMaintenanceHostsInfo(removeSplaMaintenanceHostsInfoRequestBodyParam []vmwarecloudVmc_aws_cho_providerModel.MaintenanceHost, awsRegionParam string) error
}

type maintenanceHostsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewMaintenanceHostsClient(connector vapiProtocolClient_.Connector) *maintenanceHostsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vmwarecloud.vmc_aws_cho_provider.api.inventory.operator.spla.maintenance_hosts")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get_spla_maintenance_hosts_info":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get_spla_maintenance_hosts_info"),
		"add_spla_maintenance_hosts_info":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "add_spla_maintenance_hosts_info"),
		"remove_spla_maintenance_hosts_info": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "remove_spla_maintenance_hosts_info"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	mIface := maintenanceHostsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &mIface
}

func (mIface *maintenanceHostsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := mIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (mIface *maintenanceHostsClient) GetSplaMaintenanceHostsInfo(awsRegionParam string) (vmwarecloudVmc_aws_cho_providerModel.SplaMaintenanceHostsInfo, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	operationRestMetaData := maintenanceHostsGetSplaMaintenanceHostsInfoRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(maintenanceHostsGetSplaMaintenanceHostsInfoInputType(), typeConverter)
	sv.AddStructField("AwsRegion", awsRegionParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmwarecloudVmc_aws_cho_providerModel.SplaMaintenanceHostsInfo
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.vmwarecloud.vmc_aws_cho_provider.api.inventory.operator.spla.maintenance_hosts", "get_spla_maintenance_hosts_info", inputDataValue, executionContext)
	var emptyOutput vmwarecloudVmc_aws_cho_providerModel.SplaMaintenanceHostsInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), MaintenanceHostsGetSplaMaintenanceHostsInfoOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmwarecloudVmc_aws_cho_providerModel.SplaMaintenanceHostsInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *maintenanceHostsClient) AddSplaMaintenanceHostsInfo(addSplaMaintenanceHostsInfoRequestBodyParam []vmwarecloudVmc_aws_cho_providerModel.MaintenanceHost, awsRegionParam string) (vmwarecloudVmc_aws_cho_providerModel.SplaMaintenanceHostsInfo, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	operationRestMetaData := maintenanceHostsAddSplaMaintenanceHostsInfoRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(maintenanceHostsAddSplaMaintenanceHostsInfoInputType(), typeConverter)
	sv.AddStructField("AddSplaMaintenanceHostsInfoRequestBody", addSplaMaintenanceHostsInfoRequestBodyParam)
	sv.AddStructField("AwsRegion", awsRegionParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmwarecloudVmc_aws_cho_providerModel.SplaMaintenanceHostsInfo
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.vmwarecloud.vmc_aws_cho_provider.api.inventory.operator.spla.maintenance_hosts", "add_spla_maintenance_hosts_info", inputDataValue, executionContext)
	var emptyOutput vmwarecloudVmc_aws_cho_providerModel.SplaMaintenanceHostsInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), MaintenanceHostsAddSplaMaintenanceHostsInfoOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmwarecloudVmc_aws_cho_providerModel.SplaMaintenanceHostsInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *maintenanceHostsClient) RemoveSplaMaintenanceHostsInfo(removeSplaMaintenanceHostsInfoRequestBodyParam []vmwarecloudVmc_aws_cho_providerModel.MaintenanceHost, awsRegionParam string) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	operationRestMetaData := maintenanceHostsRemoveSplaMaintenanceHostsInfoRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(maintenanceHostsRemoveSplaMaintenanceHostsInfoInputType(), typeConverter)
	sv.AddStructField("RemoveSplaMaintenanceHostsInfoRequestBody", removeSplaMaintenanceHostsInfoRequestBodyParam)
	sv.AddStructField("AwsRegion", awsRegionParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.vmwarecloud.vmc_aws_cho_provider.api.inventory.operator.spla.maintenance_hosts", "remove_spla_maintenance_hosts_info", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
