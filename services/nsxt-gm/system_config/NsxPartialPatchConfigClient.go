// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: NsxPartialPatchConfig
// Used by client-side stubs.

package system_config

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_global_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type NsxPartialPatchConfigClient interface {

	// Get Configuration values for nsx-partial-patch. By default partial patch is deactivated (i.e false).
	// @return com.vmware.nsx_global_policy.model.PartialPatchConfig
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get() (nsx_global_policyModel.PartialPatchConfig, error)

	// Update partial patch configuration values. Only boolean value is allowed for enable_partial_patch
	//
	// @param partialPatchConfigParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(partialPatchConfigParam nsx_global_policyModel.PartialPatchConfig) error
}

type nsxPartialPatchConfigClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewNsxPartialPatchConfigClient(connector vapiProtocolClient_.Connector) *nsxPartialPatchConfigClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_global_policy.system_config.nsx_partial_patch_config")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"patch": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "patch"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	nIface := nsxPartialPatchConfigClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &nIface
}

func (nIface *nsxPartialPatchConfigClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := nIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (nIface *nsxPartialPatchConfigClient) Get() (nsx_global_policyModel.PartialPatchConfig, error) {
	typeConverter := nIface.connector.TypeConverter()
	executionContext := nIface.connector.NewExecutionContext()
	operationRestMetaData := nsxPartialPatchConfigGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(nsxPartialPatchConfigGetInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_global_policyModel.PartialPatchConfig
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := nIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.system_config.nsx_partial_patch_config", "get", inputDataValue, executionContext)
	var emptyOutput nsx_global_policyModel.PartialPatchConfig
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), NsxPartialPatchConfigGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_global_policyModel.PartialPatchConfig), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (nIface *nsxPartialPatchConfigClient) Patch(partialPatchConfigParam nsx_global_policyModel.PartialPatchConfig) error {
	typeConverter := nIface.connector.TypeConverter()
	executionContext := nIface.connector.NewExecutionContext()
	operationRestMetaData := nsxPartialPatchConfigPatchRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(nsxPartialPatchConfigPatchInputType(), typeConverter)
	sv.AddStructField("PartialPatchConfig", partialPatchConfigParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := nIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.system_config.nsx_partial_patch_config", "patch", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
