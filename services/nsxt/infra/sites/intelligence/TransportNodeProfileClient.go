// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: TransportNodeProfile
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

type TransportNodeProfileClient interface {

	// Get the current Intelligence transport node profile. Recommend to keep the value same for flow_data_collection_interval and context_data_collection_interval.
	//
	// @param siteIdParam (required)
	// @return com.vmware.nsx_policy.model.IntelligenceTransportNodeProfile
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(siteIdParam string) (model.IntelligenceTransportNodeProfile, error)

	// Patch the current Intelligence transport node profile. Return error if NAPP is not registered with NSX or Intelligence feature is disabled.
	//
	// @param siteIdParam (required)
	// @param intelligenceTransportNodeProfileParam (required)
	// @return com.vmware.nsx_policy.model.IntelligenceTransportNodeProfile
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(siteIdParam string, intelligenceTransportNodeProfileParam model.IntelligenceTransportNodeProfile) (model.IntelligenceTransportNodeProfile, error)

	// Reset Intelligence transport node profile to the default setting. Clear Intelligence transport node profile if NAPP is not registered with NSX or Intelligence feature is disabled. Return the Intelligence transport node profile after reset operation.
	//
	// @param siteIdParam (required)
	// @return com.vmware.nsx_policy.model.IntelligenceTransportNodeProfile
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Reset(siteIdParam string) (model.IntelligenceTransportNodeProfile, error)
}

type transportNodeProfileClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewTransportNodeProfileClient(connector client.Connector) *transportNodeProfileClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.sites.intelligence.transport_node_profile")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get":   core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"patch": core.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"reset": core.NewMethodIdentifier(interfaceIdentifier, "reset"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	tIface := transportNodeProfileClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &tIface
}

func (tIface *transportNodeProfileClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := tIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (tIface *transportNodeProfileClient) Get(siteIdParam string) (model.IntelligenceTransportNodeProfile, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(transportNodeProfileGetInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IntelligenceTransportNodeProfile
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := transportNodeProfileGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.intelligence.transport_node_profile", "get", inputDataValue, executionContext)
	var emptyOutput model.IntelligenceTransportNodeProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), transportNodeProfileGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IntelligenceTransportNodeProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *transportNodeProfileClient) Patch(siteIdParam string, intelligenceTransportNodeProfileParam model.IntelligenceTransportNodeProfile) (model.IntelligenceTransportNodeProfile, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(transportNodeProfilePatchInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("IntelligenceTransportNodeProfile", intelligenceTransportNodeProfileParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IntelligenceTransportNodeProfile
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := transportNodeProfilePatchRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.intelligence.transport_node_profile", "patch", inputDataValue, executionContext)
	var emptyOutput model.IntelligenceTransportNodeProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), transportNodeProfilePatchOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IntelligenceTransportNodeProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *transportNodeProfileClient) Reset(siteIdParam string) (model.IntelligenceTransportNodeProfile, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(transportNodeProfileResetInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IntelligenceTransportNodeProfile
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := transportNodeProfileResetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.intelligence.transport_node_profile", "reset", inputDataValue, executionContext)
	var emptyOutput model.IntelligenceTransportNodeProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), transportNodeProfileResetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IntelligenceTransportNodeProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
