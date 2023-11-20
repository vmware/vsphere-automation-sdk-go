// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ConnectedAccounts
// Used by client-side stubs.

package account_link

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	vmcModel "github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type ConnectedAccountsClient interface {

	// Delete a particular connected (linked) account.
	//
	// @param orgParam Organization identifier (required)
	// @param linkedAccountPathIdParam The linked connected account identifier (required)
	// @param forceEvenWhenSddcPresentParam When true, forcibly removes a connected account even when SDDC's are still linked to it. (optional)
	// @return com.vmware.vmc.model.AwsCustomerConnectedAccount
	//
	// @throws Unauthenticated  Unauthorized
	// @throws InvalidRequest  An invalid connected account ID was specified, or the connection still has SDDCs active on it.
	// @throws Unauthorized  Forbidden
	Delete(orgParam string, linkedAccountPathIdParam string, forceEvenWhenSddcPresentParam *bool) (vmcModel.AwsCustomerConnectedAccount, error)

	// Get a list of connected accounts
	//
	// @param orgParam Organization identifier (required)
	// @param providerParam The cloud provider of the SDDC (AWS or ZeroCloud). Default value is AWS. (optional)
	//
	// @throws Unauthenticated  Unauthorized
	// @throws Unauthorized  Forbidden
	Get(orgParam string, providerParam *string) ([]vmcModel.AwsCustomerConnectedAccount, error)
}

type connectedAccountsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewConnectedAccountsClient(connector vapiProtocolClient_.Connector) *connectedAccountsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vmc.orgs.account_link.connected_accounts")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"delete": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	cIface := connectedAccountsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &cIface
}

func (cIface *connectedAccountsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := cIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (cIface *connectedAccountsClient) Delete(orgParam string, linkedAccountPathIdParam string, forceEvenWhenSddcPresentParam *bool) (vmcModel.AwsCustomerConnectedAccount, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := connectedAccountsDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(connectedAccountsDeleteInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("LinkedAccountPathId", linkedAccountPathIdParam)
	sv.AddStructField("ForceEvenWhenSddcPresent", forceEvenWhenSddcPresentParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmcModel.AwsCustomerConnectedAccount
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.vmc.orgs.account_link.connected_accounts", "delete", inputDataValue, executionContext)
	var emptyOutput vmcModel.AwsCustomerConnectedAccount
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ConnectedAccountsDeleteOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmcModel.AwsCustomerConnectedAccount), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *connectedAccountsClient) Get(orgParam string, providerParam *string) ([]vmcModel.AwsCustomerConnectedAccount, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := connectedAccountsGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(connectedAccountsGetInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Provider", providerParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []vmcModel.AwsCustomerConnectedAccount
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.vmc.orgs.account_link.connected_accounts", "get", inputDataValue, executionContext)
	var emptyOutput []vmcModel.AwsCustomerConnectedAccount
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ConnectedAccountsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.([]vmcModel.AwsCustomerConnectedAccount), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
