// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Aggregated
// Used by client-side stubs.

package drafts

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

const _ = core.SupportedByRuntimeVersion1

type AggregatedClient interface {

	// Get an aggregated configuration that will get applied onto current configuration during publish of this draft. The response is a hierarchical payload containing the aggregated configuration differences from the latest auto draft till the specified draft.
	//
	// @param draftIdParam (required)
	// @return com.vmware.nsx_global_policy.model.Infra
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(draftIdParam string) (model.Infra, error)
}

type aggregatedClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewAggregatedClient(connector client.Connector) *aggregatedClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_global_policy.global_infra.drafts.aggregated")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get": core.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	aIface := aggregatedClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &aIface
}

func (aIface *aggregatedClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := aIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (aIface *aggregatedClient) Get(draftIdParam string) (model.Infra, error) {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(aggregatedGetInputType(), typeConverter)
	sv.AddStructField("DraftId", draftIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.Infra
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := aggregatedGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	aIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.global_infra.drafts.aggregated", "get", inputDataValue, executionContext)
	var emptyOutput model.Infra
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), aggregatedGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.Infra), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
