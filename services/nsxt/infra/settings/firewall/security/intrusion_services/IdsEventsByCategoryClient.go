// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: IdsEventsByCategory
// Used by client-side stubs.

package intrusion_services

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = core.SupportedByRuntimeVersion1

type IdsEventsByCategoryClient interface {

	// Get the precentage of the IDS events events detected grouped by the category provided. Supported categories could be Attack Type, Attack Target, and Severity.
	//
	// @param idsCategoryNameParam IDS category name (required)
	// @param timeperiodParam Time period over which data needs to be collected (required)
	// @param detectedOnGatewayParam Show events detected on gateways (optional)
	// @param gatewayParam Name of Gateway (optional)
	// @return com.vmware.nsx_policy.model.PolicyIdsEventsByCategoryResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(idsCategoryNameParam string, timeperiodParam string, detectedOnGatewayParam *bool, gatewayParam *string) (model.PolicyIdsEventsByCategoryResult, error)
}

type idsEventsByCategoryClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewIdsEventsByCategoryClient(connector client.Connector) *idsEventsByCategoryClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.settings.firewall.security.intrusion_services.ids_events_by_category")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"list": core.NewMethodIdentifier(interfaceIdentifier, "list"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	iIface := idsEventsByCategoryClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &iIface
}

func (iIface *idsEventsByCategoryClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := iIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (iIface *idsEventsByCategoryClient) List(idsCategoryNameParam string, timeperiodParam string, detectedOnGatewayParam *bool, gatewayParam *string) (model.PolicyIdsEventsByCategoryResult, error) {
	typeConverter := iIface.connector.TypeConverter()
	executionContext := iIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(idsEventsByCategoryListInputType(), typeConverter)
	sv.AddStructField("IdsCategoryName", idsCategoryNameParam)
	sv.AddStructField("Timeperiod", timeperiodParam)
	sv.AddStructField("DetectedOnGateway", detectedOnGatewayParam)
	sv.AddStructField("Gateway", gatewayParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.PolicyIdsEventsByCategoryResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := idsEventsByCategoryListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	iIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := iIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.settings.firewall.security.intrusion_services.ids_events_by_category", "list", inputDataValue, executionContext)
	var emptyOutput model.PolicyIdsEventsByCategoryResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), idsEventsByCategoryListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.PolicyIdsEventsByCategoryResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
