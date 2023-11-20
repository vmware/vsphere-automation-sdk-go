// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: AuthPolicy
// Used by client-side stubs.

package aaa

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type AuthPolicyClient interface {

	// Returns information about the currently configured authentication policies and password complexity on the node.
	// @return com.vmware.nsx.model.AuthenticationPolicyProperties
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get() (nsxModel.AuthenticationPolicyProperties, error)

	// Resets to default, currently configured authentication policy and password complexity on the node. Administrators need to enforce password change for existing user accounts in order to match newly configured complexity requirements in system.
	//
	// **reset-all**: resets configured Authentication policy and Password complexity
	//
	// **reset-auth-policies**: resets only configured Authentication policy
	//  *includes - {api_failed_auth_lockout_period, api_failed_auth_reset_period, api_max_auth_failures, cli_failed_auth_lockout_period, cli_max_auth_failures}*
	//
	// **reset-pwd-complexity**: resets only configured Password complexity
	// @return com.vmware.nsx.model.AuthenticationPolicyProperties
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Resetall() (nsxModel.AuthenticationPolicyProperties, error)

	// Resets to default, currently configured authentication policy and password complexity on the node. Administrators need to enforce password change for existing user accounts in order to match newly configured complexity requirements in system.
	//
	// **reset-all**: resets configured Authentication policy and Password complexity
	//
	// **reset-auth-policies**: resets only configured Authentication policy
	//  *includes - {api_failed_auth_lockout_period, api_failed_auth_reset_period, api_max_auth_failures, cli_failed_auth_lockout_period, cli_max_auth_failures}*
	//
	// **reset-pwd-complexity**: resets only configured Password complexity
	// @return com.vmware.nsx.model.AuthenticationPolicyProperties
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Resetauthpolicies() (nsxModel.AuthenticationPolicyProperties, error)

	// Resets to default, currently configured authentication policy and password complexity on the node. Administrators need to enforce password change for existing user accounts in order to match newly configured complexity requirements in system.
	//
	// **reset-all**: resets configured Authentication policy and Password complexity
	//
	// **reset-auth-policies**: resets only configured Authentication policy
	//  *includes - {api_failed_auth_lockout_period, api_failed_auth_reset_period, api_max_auth_failures, cli_failed_auth_lockout_period, cli_max_auth_failures}*
	//
	// **reset-pwd-complexity**: resets only configured Password complexity
	// @return com.vmware.nsx.model.AuthenticationPolicyProperties
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Resetpwdcomplexity() (nsxModel.AuthenticationPolicyProperties, error)

	// Update the currently configured authentication policy and password complexity on the node. If any of api_max_auth_failures, api_failed_auth_reset_period, or api_failed_auth_lockout_period are modified, the http service is automatically restarted. Whereas change in any password complexity will not be applicable on already configured user passwords. Administrators need to enforce password change for existing user accounts in order to match newly configured complexity requirements enforced in system. All values from AuthenticationPolicyProperties are in sync among the management cluster nodes.
	//
	// @param authenticationPolicyPropertiesParam (required)
	// @return com.vmware.nsx.model.AuthenticationPolicyProperties
	//
	// @throws ConcurrentChange  Conflict
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(authenticationPolicyPropertiesParam nsxModel.AuthenticationPolicyProperties) (nsxModel.AuthenticationPolicyProperties, error)
}

type authPolicyClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewAuthPolicyClient(connector vapiProtocolClient_.Connector) *authPolicyClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.node.aaa.auth_policy")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get":                vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"resetall":           vapiCore_.NewMethodIdentifier(interfaceIdentifier, "resetall"),
		"resetauthpolicies":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "resetauthpolicies"),
		"resetpwdcomplexity": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "resetpwdcomplexity"),
		"update":             vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	aIface := authPolicyClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &aIface
}

func (aIface *authPolicyClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := aIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (aIface *authPolicyClient) Get() (nsxModel.AuthenticationPolicyProperties, error) {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	operationRestMetaData := authPolicyGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(authPolicyGetInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.AuthenticationPolicyProperties
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.nsx.node.aaa.auth_policy", "get", inputDataValue, executionContext)
	var emptyOutput nsxModel.AuthenticationPolicyProperties
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), AuthPolicyGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.AuthenticationPolicyProperties), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (aIface *authPolicyClient) Resetall() (nsxModel.AuthenticationPolicyProperties, error) {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	operationRestMetaData := authPolicyResetallRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(authPolicyResetallInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.AuthenticationPolicyProperties
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.nsx.node.aaa.auth_policy", "resetall", inputDataValue, executionContext)
	var emptyOutput nsxModel.AuthenticationPolicyProperties
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), AuthPolicyResetallOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.AuthenticationPolicyProperties), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (aIface *authPolicyClient) Resetauthpolicies() (nsxModel.AuthenticationPolicyProperties, error) {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	operationRestMetaData := authPolicyResetauthpoliciesRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(authPolicyResetauthpoliciesInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.AuthenticationPolicyProperties
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.nsx.node.aaa.auth_policy", "resetauthpolicies", inputDataValue, executionContext)
	var emptyOutput nsxModel.AuthenticationPolicyProperties
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), AuthPolicyResetauthpoliciesOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.AuthenticationPolicyProperties), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (aIface *authPolicyClient) Resetpwdcomplexity() (nsxModel.AuthenticationPolicyProperties, error) {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	operationRestMetaData := authPolicyResetpwdcomplexityRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(authPolicyResetpwdcomplexityInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.AuthenticationPolicyProperties
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.nsx.node.aaa.auth_policy", "resetpwdcomplexity", inputDataValue, executionContext)
	var emptyOutput nsxModel.AuthenticationPolicyProperties
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), AuthPolicyResetpwdcomplexityOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.AuthenticationPolicyProperties), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (aIface *authPolicyClient) Update(authenticationPolicyPropertiesParam nsxModel.AuthenticationPolicyProperties) (nsxModel.AuthenticationPolicyProperties, error) {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	operationRestMetaData := authPolicyUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(authPolicyUpdateInputType(), typeConverter)
	sv.AddStructField("AuthenticationPolicyProperties", authenticationPolicyPropertiesParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.AuthenticationPolicyProperties
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.nsx.node.aaa.auth_policy", "update", inputDataValue, executionContext)
	var emptyOutput nsxModel.AuthenticationPolicyProperties
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), AuthPolicyUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.AuthenticationPolicyProperties), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
