// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: FirewallIdentityStores
// Used by client-side stubs.

package global_infra

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type FirewallIdentityStoresClient interface {

	// Invoke full sync, delta sync, or stop sync for a directory domain
	//
	//  Use the following Policy API -
	//  POST /infra/identity-firewall-stores/action/delta-sync; POST /infra/identity-firewall-stores/action/full-sync; POST /infra/identity-firewall-stores/action/stop-sync;
	//
	// Deprecated: This API element is deprecated.
	//
	// @param firewallIdentityStoreIdParam Firewall identity store identifier (required)
	// @param actionParam Sync type could be either FULL sync or DELTA sync. The full sync fetches all the objects under the configured sync nodes while delta sync will get the changed objects from previous sync time. FULL_SYNC - Perform a full synchronization, where the local state of all AD objects is updated. DELTA_SYNC - Perform a delta synchronization, where local AD objects that have changed since the last synchronization are updated. STOP_SYNC - Stop the synchronization process. (required)
	// @param delayParam The delay can be added to execute the sync action in the future. (optional, default to 0)
	// @param enforcementPointPathParam enforcement point path, forward slashes must be escaped using %2F. (optional)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(firewallIdentityStoreIdParam string, actionParam string, delayParam *int64, enforcementPointPathParam *string) error
}

type firewallIdentityStoresClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
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
