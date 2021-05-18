// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ConnectedAccounts
// Used by client-side stubs.

package account_link

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/vmc/model"
)

const _ = core.SupportedByRuntimeVersion1

type ConnectedAccountsClient interface {

	// Delete a particular connected (linked) account.
	//
	// @param orgParam Organization identifier (required)
	// @param linkedAccountPathIdParam The linked connected account identifier (required)
	// @param forceEvenWhenSddcPresentParam When true, forcibly removes a connected account even when SDDC's are still linked to it. (optional)
	// @return com.vmware.vmc.model.AwsCustomerConnectedAccount
	// @throws Unauthenticated  Unauthorized
	// @throws InvalidRequest  An invalid connected account ID was specified, or the connection still has SDDCs active on it.
	// @throws Unauthorized  Forbidden
	Delete(orgParam string, linkedAccountPathIdParam string, forceEvenWhenSddcPresentParam *bool) (model.AwsCustomerConnectedAccount, error)

	// Get a list of connected accounts
	//
	// @param orgParam Organization identifier (required)
	// @param providerParam The cloud provider of the SDDC (AWS or ZeroCloud). Default value is AWS. (optional)
	// @throws Unauthenticated  Unauthorized
	// @throws Unauthorized  Forbidden
	Get(orgParam string, providerParam *string) ([]model.AwsCustomerConnectedAccount, error)
}

type connectedAccountsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewConnectedAccountsClient(connector client.Connector) *connectedAccountsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.vmc.orgs.account_link.connected_accounts")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"delete": core.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    core.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	cIface := connectedAccountsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &cIface
}

func (cIface *connectedAccountsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := cIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (cIface *connectedAccountsClient) Delete(orgParam string, linkedAccountPathIdParam string, forceEvenWhenSddcPresentParam *bool) (model.AwsCustomerConnectedAccount, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(connectedAccountsDeleteInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("LinkedAccountPathId", linkedAccountPathIdParam)
	sv.AddStructField("ForceEvenWhenSddcPresent", forceEvenWhenSddcPresentParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.AwsCustomerConnectedAccount
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := connectedAccountsDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.vmc.orgs.account_link.connected_accounts", "delete", inputDataValue, executionContext)
	var emptyOutput model.AwsCustomerConnectedAccount
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), connectedAccountsDeleteOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.AwsCustomerConnectedAccount), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *connectedAccountsClient) Get(orgParam string, providerParam *string) ([]model.AwsCustomerConnectedAccount, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(connectedAccountsGetInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Provider", providerParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []model.AwsCustomerConnectedAccount
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := connectedAccountsGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.vmc.orgs.account_link.connected_accounts", "get", inputDataValue, executionContext)
	var emptyOutput []model.AwsCustomerConnectedAccount
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), connectedAccountsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.([]model.AwsCustomerConnectedAccount), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
