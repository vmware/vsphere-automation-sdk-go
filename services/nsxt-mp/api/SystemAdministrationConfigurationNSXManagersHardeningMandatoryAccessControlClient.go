// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: SystemAdministrationConfigurationNSXManagersHardeningMandatoryAccessControl
// Used by client-side stubs.

package api

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-mp/model"
)

const _ = core.SupportedByRuntimeVersion1

type SystemAdministrationConfigurationNSXManagersHardeningMandatoryAccessControlClient interface {

	// Gets the enable status for Mandatory Access Control
	// @return com.vmware.model.MandatoryAccessControlProperties
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetNodeMandatoryAccessControl() (model.MandatoryAccessControlProperties, error)

	// Get the report for Mandatory Access Control
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetNodeMandatoryAccessControlReport() error

	// Enable or disable Mandatory Access Control
	//
	// @param mandatoryAccessControlPropertiesParam (required)
	// @return com.vmware.model.MandatoryAccessControlProperties
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	SetNodeMandatoryAccessControl(mandatoryAccessControlPropertiesParam model.MandatoryAccessControlProperties) (model.MandatoryAccessControlProperties, error)
}

type systemAdministrationConfigurationNSXManagersHardeningMandatoryAccessControlClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSystemAdministrationConfigurationNSXManagersHardeningMandatoryAccessControlClient(connector client.Connector) *systemAdministrationConfigurationNSXManagersHardeningMandatoryAccessControlClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.system_administration_configuration_NSX_managers_hardening_mandatory_access_control")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get_node_mandatory_access_control":        core.NewMethodIdentifier(interfaceIdentifier, "get_node_mandatory_access_control"),
		"get_node_mandatory_access_control_report": core.NewMethodIdentifier(interfaceIdentifier, "get_node_mandatory_access_control_report"),
		"set_node_mandatory_access_control":        core.NewMethodIdentifier(interfaceIdentifier, "set_node_mandatory_access_control"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := systemAdministrationConfigurationNSXManagersHardeningMandatoryAccessControlClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *systemAdministrationConfigurationNSXManagersHardeningMandatoryAccessControlClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *systemAdministrationConfigurationNSXManagersHardeningMandatoryAccessControlClient) GetNodeMandatoryAccessControl() (model.MandatoryAccessControlProperties, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationNSXManagersHardeningMandatoryAccessControlGetNodeMandatoryAccessControlInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.MandatoryAccessControlProperties
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationNSXManagersHardeningMandatoryAccessControlGetNodeMandatoryAccessControlRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_NSX_managers_hardening_mandatory_access_control", "get_node_mandatory_access_control", inputDataValue, executionContext)
	var emptyOutput model.MandatoryAccessControlProperties
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationNSXManagersHardeningMandatoryAccessControlGetNodeMandatoryAccessControlOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.MandatoryAccessControlProperties), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationNSXManagersHardeningMandatoryAccessControlClient) GetNodeMandatoryAccessControlReport() error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationNSXManagersHardeningMandatoryAccessControlGetNodeMandatoryAccessControlReportInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationNSXManagersHardeningMandatoryAccessControlGetNodeMandatoryAccessControlReportRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_NSX_managers_hardening_mandatory_access_control", "get_node_mandatory_access_control_report", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationNSXManagersHardeningMandatoryAccessControlClient) SetNodeMandatoryAccessControl(mandatoryAccessControlPropertiesParam model.MandatoryAccessControlProperties) (model.MandatoryAccessControlProperties, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationNSXManagersHardeningMandatoryAccessControlSetNodeMandatoryAccessControlInputType(), typeConverter)
	sv.AddStructField("MandatoryAccessControlProperties", mandatoryAccessControlPropertiesParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.MandatoryAccessControlProperties
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationNSXManagersHardeningMandatoryAccessControlSetNodeMandatoryAccessControlRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_NSX_managers_hardening_mandatory_access_control", "set_node_mandatory_access_control", inputDataValue, executionContext)
	var emptyOutput model.MandatoryAccessControlProperties
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationNSXManagersHardeningMandatoryAccessControlSetNodeMandatoryAccessControlOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.MandatoryAccessControlProperties), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
