// Copyright (c) 2019-2025 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Info
// Used by client-side stubs.

package controller_info

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type InfoClient interface {

	// Delete information of Avi controller in NSX
	//
	// @param controllerIdParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(controllerIdParam string) error

	// Read a controller information registered to NSX.
	//
	// @param controllerIdParam (required)
	// @return com.vmware.nsx_policy.model.AlbControllerInfo
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(controllerIdParam string) (nsx_policyModel.AlbControllerInfo, error)

	// List controllers registered to NSX for licensing.
	// @return com.vmware.nsx_policy.model.AlbControllerInfoList
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List() (nsx_policyModel.AlbControllerInfoList, error)

	// If the passed controller info does not already exist, create a new controller info. If it already exists, replace it.
	//
	// @param controllerIdParam (required)
	// @param albControllerInfoParam (required)
	// @return com.vmware.nsx_policy.model.AlbControllerInfo
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(controllerIdParam string, albControllerInfoParam nsx_policyModel.AlbControllerInfo) (nsx_policyModel.AlbControllerInfo, error)
}

type infoClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewInfoClient(connector vapiProtocolClient_.Connector) *infoClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.alb.controller_info.info")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"delete": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	iIface := infoClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &iIface
}

func (iIface *infoClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := iIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (iIface *infoClient) Delete(controllerIdParam string) error {
	typeConverter := iIface.connector.TypeConverter()
	executionContext := iIface.connector.NewExecutionContext()
	operationRestMetaData := infoDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(infoDeleteInputType(), typeConverter)
	sv.AddStructField("ControllerId", controllerIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := iIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.alb.controller_info.info", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (iIface *infoClient) Get(controllerIdParam string) (nsx_policyModel.AlbControllerInfo, error) {
	typeConverter := iIface.connector.TypeConverter()
	executionContext := iIface.connector.NewExecutionContext()
	operationRestMetaData := infoGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(infoGetInputType(), typeConverter)
	sv.AddStructField("ControllerId", controllerIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.AlbControllerInfo
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := iIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.alb.controller_info.info", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.AlbControllerInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), InfoGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.AlbControllerInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (iIface *infoClient) List() (nsx_policyModel.AlbControllerInfoList, error) {
	typeConverter := iIface.connector.TypeConverter()
	executionContext := iIface.connector.NewExecutionContext()
	operationRestMetaData := infoListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(infoListInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.AlbControllerInfoList
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := iIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.alb.controller_info.info", "list", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.AlbControllerInfoList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), InfoListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.AlbControllerInfoList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (iIface *infoClient) Update(controllerIdParam string, albControllerInfoParam nsx_policyModel.AlbControllerInfo) (nsx_policyModel.AlbControllerInfo, error) {
	typeConverter := iIface.connector.TypeConverter()
	executionContext := iIface.connector.NewExecutionContext()
	operationRestMetaData := infoUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(infoUpdateInputType(), typeConverter)
	sv.AddStructField("ControllerId", controllerIdParam)
	sv.AddStructField("AlbControllerInfo", albControllerInfoParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.AlbControllerInfo
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := iIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.alb.controller_info.info", "update", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.AlbControllerInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), InfoUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.AlbControllerInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
