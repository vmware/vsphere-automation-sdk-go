// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: HostConfig
// Used by client-side stubs.

package intelligence

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type HostConfigClient interface {

	// Deprecated - Please use the intelligence API /napp/api/v1/intelligence/data-collection/host-config instead, after installing NSX Intelligence. Get the current NSX-Intelligence host configuration. Recommend to keep the value same for flow_data_collection_interval and context_data_collection_interval.
	// @return com.vmware.nsx.model.IntelligenceHostConfigurationInfo
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get() (model.IntelligenceHostConfigurationInfo, error)

	// Deprecated - Please use the intelligence API /napp/api/v1/intelligence/data-collection/host-config instead, after installing NSX Intelligence. Patch the current NSX-Intelligence host configuration. Return error if NSX-Intelligence is not registered with NSX.
	//
	// @param intelligenceHostConfigurationInfoParam (required)
	// @return com.vmware.nsx.model.IntelligenceHostConfigurationInfo
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(intelligenceHostConfigurationInfoParam model.IntelligenceHostConfigurationInfo) (model.IntelligenceHostConfigurationInfo, error)

	// Deprecated - Please use the intelligence API /napp/api/v1/intelligence/data-collection/host-config?action=reset instead, after installing NSX Intelligence. Reset NSX-Intelligence host configuration to the default setting. Clear NSX-Intelligence host configuration if NSX-Intelligence is not registered with NSX. Return the NSX-Intelligence host configuration after reset operation.
	// @return com.vmware.nsx.model.IntelligenceHostConfigurationInfo
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Reset() (model.IntelligenceHostConfigurationInfo, error)
}

type hostConfigClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewHostConfigClient(connector client.Connector) *hostConfigClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.intelligence.host_config")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get":   core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"patch": core.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"reset": core.NewMethodIdentifier(interfaceIdentifier, "reset"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	hIface := hostConfigClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &hIface
}

func (hIface *hostConfigClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := hIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (hIface *hostConfigClient) Get() (model.IntelligenceHostConfigurationInfo, error) {
	typeConverter := hIface.connector.TypeConverter()
	executionContext := hIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(hostConfigGetInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IntelligenceHostConfigurationInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := hostConfigGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	hIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := hIface.connector.GetApiProvider().Invoke("com.vmware.nsx.intelligence.host_config", "get", inputDataValue, executionContext)
	var emptyOutput model.IntelligenceHostConfigurationInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), hostConfigGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IntelligenceHostConfigurationInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), hIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (hIface *hostConfigClient) Patch(intelligenceHostConfigurationInfoParam model.IntelligenceHostConfigurationInfo) (model.IntelligenceHostConfigurationInfo, error) {
	typeConverter := hIface.connector.TypeConverter()
	executionContext := hIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(hostConfigPatchInputType(), typeConverter)
	sv.AddStructField("IntelligenceHostConfigurationInfo", intelligenceHostConfigurationInfoParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IntelligenceHostConfigurationInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := hostConfigPatchRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	hIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := hIface.connector.GetApiProvider().Invoke("com.vmware.nsx.intelligence.host_config", "patch", inputDataValue, executionContext)
	var emptyOutput model.IntelligenceHostConfigurationInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), hostConfigPatchOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IntelligenceHostConfigurationInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), hIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (hIface *hostConfigClient) Reset() (model.IntelligenceHostConfigurationInfo, error) {
	typeConverter := hIface.connector.TypeConverter()
	executionContext := hIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(hostConfigResetInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IntelligenceHostConfigurationInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := hostConfigResetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	hIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := hIface.connector.GetApiProvider().Invoke("com.vmware.nsx.intelligence.host_config", "reset", inputDataValue, executionContext)
	var emptyOutput model.IntelligenceHostConfigurationInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), hostConfigResetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IntelligenceHostConfigurationInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), hIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
