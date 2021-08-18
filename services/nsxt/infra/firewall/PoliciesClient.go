// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Policies
// Used by client-side stubs.

package firewall

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = core.SupportedByRuntimeVersion1

type PoliciesClient interface {

	// Get the list of policies filtered based on the given criteria.
	//
	// @param scopeParam Scope filter criteria (required)
	// @param enforcementPointPathParam Path of the enforcement point (optional)
	// @param parentPathParam Path of the parent object of the entities (optional)
	// @return com.vmware.nsx_policy.model.PolicyResourceReferenceForEPListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(scopeParam string, enforcementPointPathParam *string, parentPathParam *string) (model.PolicyResourceReferenceForEPListResult, error)
}

type policiesClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewPoliciesClient(connector client.Connector) *policiesClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.firewall.policies")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"list": core.NewMethodIdentifier(interfaceIdentifier, "list"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	pIface := policiesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &pIface
}

func (pIface *policiesClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := pIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (pIface *policiesClient) List(scopeParam string, enforcementPointPathParam *string, parentPathParam *string) (model.PolicyResourceReferenceForEPListResult, error) {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(policiesListInputType(), typeConverter)
	sv.AddStructField("Scope", scopeParam)
	sv.AddStructField("EnforcementPointPath", enforcementPointPathParam)
	sv.AddStructField("ParentPath", parentPathParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.PolicyResourceReferenceForEPListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := policiesListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.firewall.policies", "list", inputDataValue, executionContext)
	var emptyOutput model.PolicyResourceReferenceForEPListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), policiesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.PolicyResourceReferenceForEPListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
