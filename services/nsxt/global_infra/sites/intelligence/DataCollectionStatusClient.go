// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: DataCollectionStatus
// Used by client-side stubs.

package intelligence

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = core.SupportedByRuntimeVersion1

type DataCollectionStatusClient interface {

	// Get NSX Intelligence data collection status, including number of host nodes and number of host clusters that have data collection enabled. The number of host nodes includes both standalone hosts and hosts in clusters.
	//
	// @param siteIdParam (required)
	// @return com.vmware.nsx_policy.model.IntelligenceDataCollectionStatus
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(siteIdParam string) (model.IntelligenceDataCollectionStatus, error)
}

type dataCollectionStatusClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewDataCollectionStatusClient(connector client.Connector) *dataCollectionStatusClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.global_infra.sites.intelligence.data_collection_status")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get": core.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	dIface := dataCollectionStatusClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &dIface
}

func (dIface *dataCollectionStatusClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := dIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (dIface *dataCollectionStatusClient) Get(siteIdParam string) (model.IntelligenceDataCollectionStatus, error) {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(dataCollectionStatusGetInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IntelligenceDataCollectionStatus
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := dataCollectionStatusGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	dIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.global_infra.sites.intelligence.data_collection_status", "get", inputDataValue, executionContext)
	var emptyOutput model.IntelligenceDataCollectionStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), dataCollectionStatusGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IntelligenceDataCollectionStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
