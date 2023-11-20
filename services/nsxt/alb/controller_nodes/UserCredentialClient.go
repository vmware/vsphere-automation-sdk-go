// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: UserCredential
// Used by client-side stubs.

package controller_nodes

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type UserCredentialClient interface {

	// Use this API to create service user object credential in ALB controller clutser
	//
	// @param aLBControllerUserCredentialParam (required)
	// @return com.vmware.nsx_policy.model.ALBControllerUserCredentialResponse
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(aLBControllerUserCredentialParam nsx_policyModel.ALBControllerUserCredential) (nsx_policyModel.ALBControllerUserCredentialResponse, error)

	// Use this API to update admin user password or service user object password in ALB controller.
	//
	// @param aLBControllerUserCredentialParam (required)
	// @param runningConfigParam Update Advanced Load Balancer Controller runtime config as well (optional)
	// @return com.vmware.nsx_policy.model.ALBControllerUserCredentialResponse
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(aLBControllerUserCredentialParam nsx_policyModel.ALBControllerUserCredential, runningConfigParam *bool) (nsx_policyModel.ALBControllerUserCredentialResponse, error)
}

type userCredentialClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewUserCredentialClient(connector vapiProtocolClient_.Connector) *userCredentialClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.alb.controller_nodes.user_credential")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"create": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "create"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	uIface := userCredentialClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &uIface
}

func (uIface *userCredentialClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := uIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (uIface *userCredentialClient) Create(aLBControllerUserCredentialParam nsx_policyModel.ALBControllerUserCredential) (nsx_policyModel.ALBControllerUserCredentialResponse, error) {
	typeConverter := uIface.connector.TypeConverter()
	executionContext := uIface.connector.NewExecutionContext()
	operationRestMetaData := userCredentialCreateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(userCredentialCreateInputType(), typeConverter)
	sv.AddStructField("ALBControllerUserCredential", aLBControllerUserCredentialParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.ALBControllerUserCredentialResponse
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := uIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.alb.controller_nodes.user_credential", "create", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.ALBControllerUserCredentialResponse
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), UserCredentialCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.ALBControllerUserCredentialResponse), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), uIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (uIface *userCredentialClient) Update(aLBControllerUserCredentialParam nsx_policyModel.ALBControllerUserCredential, runningConfigParam *bool) (nsx_policyModel.ALBControllerUserCredentialResponse, error) {
	typeConverter := uIface.connector.TypeConverter()
	executionContext := uIface.connector.NewExecutionContext()
	operationRestMetaData := userCredentialUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(userCredentialUpdateInputType(), typeConverter)
	sv.AddStructField("ALBControllerUserCredential", aLBControllerUserCredentialParam)
	sv.AddStructField("RunningConfig", runningConfigParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.ALBControllerUserCredentialResponse
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := uIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.alb.controller_nodes.user_credential", "update", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.ALBControllerUserCredentialResponse
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), UserCredentialUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.ALBControllerUserCredentialResponse), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), uIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
