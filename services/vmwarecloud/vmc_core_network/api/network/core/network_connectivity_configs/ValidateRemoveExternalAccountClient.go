// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ValidateRemoveExternalAccount
// Used by client-side stubs.

package network_connectivity_configs

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type ValidateRemoveExternalAccountClient interface {

	// Validates removal of external account by checking any existing attachments
	//
	// @param orgIdParam organization identifier
	// @param networkConnectivityConfigIdParam NetworkConnectivityConfig Id
	// @param accountNumberParam AWS account number
	// @return OK
	//
	// @throws InvalidRequest Bad Request
	// @throws Unauthenticated Unauthorized
	// @throws Unauthorized Forbidden
	// @throws NotFound Not Found
	// @throws ConcurrentChange Conflict
	ValidateNetworkConnectivityConfigsRemoveExternalAccount(orgIdParam string, networkConnectivityConfigIdParam string, accountNumberParam string) error
}

type validateRemoveExternalAccountClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewValidateRemoveExternalAccountClient(connector vapiProtocolClient_.Connector) *validateRemoveExternalAccountClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vmwarecloud.vmc_core_network.api.network.core.network_connectivity_configs.validate_remove_external_account")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"validate_network_connectivity_configs_remove_external_account": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "validate_network_connectivity_configs_remove_external_account"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	vIface := validateRemoveExternalAccountClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &vIface
}

func (vIface *validateRemoveExternalAccountClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := vIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (vIface *validateRemoveExternalAccountClient) ValidateNetworkConnectivityConfigsRemoveExternalAccount(orgIdParam string, networkConnectivityConfigIdParam string, accountNumberParam string) error {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	operationRestMetaData := validateRemoveExternalAccountValidateNetworkConnectivityConfigsRemoveExternalAccountRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(validateRemoveExternalAccountValidateNetworkConnectivityConfigsRemoveExternalAccountInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("NetworkConnectivityConfigId", networkConnectivityConfigIdParam)
	sv.AddStructField("AccountNumber", accountNumberParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.vmwarecloud.vmc_core_network.api.network.core.network_connectivity_configs.validate_remove_external_account", "validate_network_connectivity_configs_remove_external_account", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
