// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: OrgUnits
// Used by client-side stubs.

package firewall_identity_stores

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = core.SupportedByRuntimeVersion1

type OrgUnitsClient interface {

	// Fetch all organization units for a Firewall Identity Store.
	//
	// @param firewallIdentityStoreIdParam Firewall Identity Store identifier (required)
	// @param enforcementPointPathParam String Path of the enforcement point (optional)
	// @return com.vmware.nsx_policy.model.DirectoryOrgUnitListResults
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(firewallIdentityStoreIdParam string, enforcementPointPathParam *string) (model.DirectoryOrgUnitListResults, error)
}

type orgUnitsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewOrgUnitsClient(connector client.Connector) *orgUnitsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.firewall_identity_stores.org_units")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"list": core.NewMethodIdentifier(interfaceIdentifier, "list"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	oIface := orgUnitsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &oIface
}

func (oIface *orgUnitsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := oIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (oIface *orgUnitsClient) List(firewallIdentityStoreIdParam string, enforcementPointPathParam *string) (model.DirectoryOrgUnitListResults, error) {
	typeConverter := oIface.connector.TypeConverter()
	executionContext := oIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(orgUnitsListInputType(), typeConverter)
	sv.AddStructField("FirewallIdentityStoreId", firewallIdentityStoreIdParam)
	sv.AddStructField("EnforcementPointPath", enforcementPointPathParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.DirectoryOrgUnitListResults
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := orgUnitsListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	oIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := oIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.firewall_identity_stores.org_units", "list", inputDataValue, executionContext)
	var emptyOutput model.DirectoryOrgUnitListResults
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), orgUnitsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.DirectoryOrgUnitListResults), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), oIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
