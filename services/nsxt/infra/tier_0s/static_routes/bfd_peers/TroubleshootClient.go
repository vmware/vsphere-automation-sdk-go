// Copyright (c) 2019-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Troubleshoot
// Used by client-side stubs.

package bfd_peers

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type TroubleshootClient interface {

	// Read a StaticRouteBfdTroubleshootConfig
	//
	// @param tier0IdParam Tier-0 ID (required)
	// @return com.vmware.nsx_policy.model.StaticRouteBfdTroubleshootConfig
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(tier0IdParam string) (nsx_policyModel.StaticRouteBfdTroubleshootConfig, error)

	// If a StaticRouteBfdTroubleshootConfig is not already present, create a new StaticRouteBfdTroubleshootConfig. If it already exists, update the StaticRouteBfdTroubleshootConfig. This is a full replace.
	//
	// @param tier0IdParam Tier-0 ID (required)
	// @param staticRouteBfdTroubleshootConfigParam (required)
	// @param overrideParam Locally override the global object (optional, default to false)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(tier0IdParam string, staticRouteBfdTroubleshootConfigParam nsx_policyModel.StaticRouteBfdTroubleshootConfig, overrideParam *bool) error

	// If a StaticRouteBfdTroubleshootConfig is not already present, create a new StaticRouteBfdTroubleshootConfig. If it already exists, update the StaticRouteBfdTroubleshootConfig. This operation will fully replace the object.
	//
	// @param tier0IdParam Tier-0 ID (required)
	// @param staticRouteBfdTroubleshootConfigParam (required)
	// @param overrideParam Locally override the global object (optional, default to false)
	// @return com.vmware.nsx_policy.model.StaticRouteBfdTroubleshootConfig
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(tier0IdParam string, staticRouteBfdTroubleshootConfigParam nsx_policyModel.StaticRouteBfdTroubleshootConfig, overrideParam *bool) (nsx_policyModel.StaticRouteBfdTroubleshootConfig, error)
}

type troubleshootClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewTroubleshootClient(connector vapiProtocolClient_.Connector) *troubleshootClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.tier_0s.static_routes.bfd_peers.troubleshoot")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"patch":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	tIface := troubleshootClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &tIface
}

func (tIface *troubleshootClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := tIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (tIface *troubleshootClient) Get(tier0IdParam string) (nsx_policyModel.StaticRouteBfdTroubleshootConfig, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := troubleshootGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(troubleshootGetInputType(), typeConverter)
	sv.AddStructField("Tier0Id", tier0IdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.StaticRouteBfdTroubleshootConfig
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.tier_0s.static_routes.bfd_peers.troubleshoot", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.StaticRouteBfdTroubleshootConfig
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), TroubleshootGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.StaticRouteBfdTroubleshootConfig), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *troubleshootClient) Patch(tier0IdParam string, staticRouteBfdTroubleshootConfigParam nsx_policyModel.StaticRouteBfdTroubleshootConfig, overrideParam *bool) error {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := troubleshootPatchRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(troubleshootPatchInputType(), typeConverter)
	sv.AddStructField("Tier0Id", tier0IdParam)
	sv.AddStructField("StaticRouteBfdTroubleshootConfig", staticRouteBfdTroubleshootConfigParam)
	sv.AddStructField("Override", overrideParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.tier_0s.static_routes.bfd_peers.troubleshoot", "patch", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (tIface *troubleshootClient) Update(tier0IdParam string, staticRouteBfdTroubleshootConfigParam nsx_policyModel.StaticRouteBfdTroubleshootConfig, overrideParam *bool) (nsx_policyModel.StaticRouteBfdTroubleshootConfig, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := troubleshootUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(troubleshootUpdateInputType(), typeConverter)
	sv.AddStructField("Tier0Id", tier0IdParam)
	sv.AddStructField("StaticRouteBfdTroubleshootConfig", staticRouteBfdTroubleshootConfigParam)
	sv.AddStructField("Override", overrideParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.StaticRouteBfdTroubleshootConfig
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.tier_0s.static_routes.bfd_peers.troubleshoot", "update", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.StaticRouteBfdTroubleshootConfig
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), TroubleshootUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.StaticRouteBfdTroubleshootConfig), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
