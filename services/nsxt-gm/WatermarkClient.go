// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Watermark
// Used by client-side stubs.

package nsx_global_policy

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_global_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type WatermarkClient interface {

	// This is an API called by VCF to store or update deployment information.
	//
	// @param nSXManagerDeploymentInfoParam (required)
	// @return com.vmware.nsx_global_policy.model.NSXManagerDeploymentInfo
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(nSXManagerDeploymentInfoParam nsx_global_policyModel.NSXManagerDeploymentInfo) (nsx_global_policyModel.NSXManagerDeploymentInfo, error)

	// This is an API called by VCF to delete deployment information.
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete() error

	// This is an API called by VCF to get deployment information.
	// @return com.vmware.nsx_global_policy.model.NSXManagerDeploymentInfo
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get() (nsx_global_policyModel.NSXManagerDeploymentInfo, error)

	// This is an API called by VCF to update stored deployment information.
	//
	// @param nSXManagerDeploymentInfoParam (required)
	// @return com.vmware.nsx_global_policy.model.NSXManagerDeploymentInfo
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(nSXManagerDeploymentInfoParam nsx_global_policyModel.NSXManagerDeploymentInfo) (nsx_global_policyModel.NSXManagerDeploymentInfo, error)
}

type watermarkClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewWatermarkClient(connector vapiProtocolClient_.Connector) *watermarkClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_global_policy.watermark")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"create": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "create"),
		"delete": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	wIface := watermarkClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &wIface
}

func (wIface *watermarkClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := wIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (wIface *watermarkClient) Create(nSXManagerDeploymentInfoParam nsx_global_policyModel.NSXManagerDeploymentInfo) (nsx_global_policyModel.NSXManagerDeploymentInfo, error) {
	typeConverter := wIface.connector.TypeConverter()
	executionContext := wIface.connector.NewExecutionContext()
	operationRestMetaData := watermarkCreateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(watermarkCreateInputType(), typeConverter)
	sv.AddStructField("NSXManagerDeploymentInfo", nSXManagerDeploymentInfoParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_global_policyModel.NSXManagerDeploymentInfo
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := wIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.watermark", "create", inputDataValue, executionContext)
	var emptyOutput nsx_global_policyModel.NSXManagerDeploymentInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), WatermarkCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_global_policyModel.NSXManagerDeploymentInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), wIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (wIface *watermarkClient) Delete() error {
	typeConverter := wIface.connector.TypeConverter()
	executionContext := wIface.connector.NewExecutionContext()
	operationRestMetaData := watermarkDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(watermarkDeleteInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := wIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.watermark", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), wIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (wIface *watermarkClient) Get() (nsx_global_policyModel.NSXManagerDeploymentInfo, error) {
	typeConverter := wIface.connector.TypeConverter()
	executionContext := wIface.connector.NewExecutionContext()
	operationRestMetaData := watermarkGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(watermarkGetInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_global_policyModel.NSXManagerDeploymentInfo
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := wIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.watermark", "get", inputDataValue, executionContext)
	var emptyOutput nsx_global_policyModel.NSXManagerDeploymentInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), WatermarkGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_global_policyModel.NSXManagerDeploymentInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), wIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (wIface *watermarkClient) Update(nSXManagerDeploymentInfoParam nsx_global_policyModel.NSXManagerDeploymentInfo) (nsx_global_policyModel.NSXManagerDeploymentInfo, error) {
	typeConverter := wIface.connector.TypeConverter()
	executionContext := wIface.connector.NewExecutionContext()
	operationRestMetaData := watermarkUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(watermarkUpdateInputType(), typeConverter)
	sv.AddStructField("NSXManagerDeploymentInfo", nSXManagerDeploymentInfoParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_global_policyModel.NSXManagerDeploymentInfo
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := wIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.watermark", "update", inputDataValue, executionContext)
	var emptyOutput nsx_global_policyModel.NSXManagerDeploymentInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), WatermarkUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_global_policyModel.NSXManagerDeploymentInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), wIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
