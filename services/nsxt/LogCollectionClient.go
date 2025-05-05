// Copyright (c) 2019-2025 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: LogCollection
// Used by client-side stubs.

package nsx_policy

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type LogCollectionClient interface {

	// Returns list of VCF logging agent with status information.
	// @return com.vmware.nsx_policy.model.VcfLoggingAgentConfigurationStatus
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get() (nsx_policyModel.VcfLoggingAgentConfigurationStatus, error)

	// Create or update VCF logging agent's configuration and returns list of VCF logging agent's configuration with status information.
	//
	// @param vcfLoggingAgentConfigurationParam (required)
	// @return com.vmware.nsx_policy.model.VcfLoggingAgentConfigurationStatus
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(vcfLoggingAgentConfigurationParam nsx_policyModel.VcfLoggingAgentConfiguration) (nsx_policyModel.VcfLoggingAgentConfigurationStatus, error)
}

type logCollectionClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewLogCollectionClient(connector vapiProtocolClient_.Connector) *logCollectionClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.log_collection")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	lIface := logCollectionClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &lIface
}

func (lIface *logCollectionClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := lIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (lIface *logCollectionClient) Get() (nsx_policyModel.VcfLoggingAgentConfigurationStatus, error) {
	typeConverter := lIface.connector.TypeConverter()
	executionContext := lIface.connector.NewExecutionContext()
	operationRestMetaData := logCollectionGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(logCollectionGetInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.VcfLoggingAgentConfigurationStatus
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := lIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.log_collection", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.VcfLoggingAgentConfigurationStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), LogCollectionGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.VcfLoggingAgentConfigurationStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (lIface *logCollectionClient) Update(vcfLoggingAgentConfigurationParam nsx_policyModel.VcfLoggingAgentConfiguration) (nsx_policyModel.VcfLoggingAgentConfigurationStatus, error) {
	typeConverter := lIface.connector.TypeConverter()
	executionContext := lIface.connector.NewExecutionContext()
	operationRestMetaData := logCollectionUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(logCollectionUpdateInputType(), typeConverter)
	sv.AddStructField("VcfLoggingAgentConfiguration", vcfLoggingAgentConfigurationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.VcfLoggingAgentConfigurationStatus
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := lIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.log_collection", "update", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.VcfLoggingAgentConfigurationStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), LogCollectionUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.VcfLoggingAgentConfigurationStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
