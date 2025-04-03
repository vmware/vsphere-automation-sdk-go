// Copyright (c) 2019-2025 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: FirewallIdentityStores
// Used by client-side stubs.

package global_infra

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type FirewallIdentityStoresClient interface {

    // Invoke full sync or delta sync for a specific domain, with additional delay in seconds if needed. Stop sync will try to stop any pending sync if any to return to idle state.
    //
    // @param firewallIdentityStoreIdParam Firewall identity store identifier (required)
    // @param actionParam Sync type requested (required)
    // @param delayParam Request to execute the sync with some delay in seconds (optional, default to 0)
    // @param enforcementPointPathParam String Path of the enforcement point (optional)
    //
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Create(firewallIdentityStoreIdParam string, actionParam string, delayParam *int64, enforcementPointPathParam *string) error
}


type firewallIdentityStoresClient struct {
	connector           	   vapiProtocolClient_.Connector
	interfaceDefinition 	   vapiCore_.InterfaceDefinition
	errorsBindingMap           map[string]vapiBindings_.BindingType
}

func NewFirewallIdentityStoresClient(connector vapiProtocolClient_.Connector) *firewallIdentityStoresClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.global_infra.firewall_identity_stores")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"create": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "create"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	fIface := firewallIdentityStoresClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &fIface
}

func (fIface *firewallIdentityStoresClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := fIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (fIface *firewallIdentityStoresClient) Create(firewallIdentityStoreIdParam string, actionParam string, delayParam *int64, enforcementPointPathParam *string) error {
	typeConverter := fIface.connector.TypeConverter()
	executionContext := fIface.connector.NewExecutionContext()
	operationRestMetaData := firewallIdentityStoresCreateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(firewallIdentityStoresCreateInputType(), typeConverter)
	sv.AddStructField("FirewallIdentityStoreId", firewallIdentityStoreIdParam)
	sv.AddStructField("Action", actionParam)
	sv.AddStructField("Delay", delayParam)
	sv.AddStructField("EnforcementPointPath", enforcementPointPathParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := fIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.global_infra.firewall_identity_stores", "create", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), fIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

