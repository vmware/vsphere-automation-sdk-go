// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: NodeConfigProfiles
// Used by client-side stubs.

package central_config

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type NodeConfigProfilesClient interface {

	// Returns properties in specified Central Node Config profile. Sensitive data (like SNMP v2c community strings) are included only if query parameter \"show_sensitive_data\" is true.
	//
	// @param profileIdParam Central Node Config profile id (required)
	// @param showSensitiveDataParam Show sensitive data in Central Node Config profile (optional, default to false)
	// @return com.vmware.nsx.model.CentralNodeConfigProfile
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(profileIdParam string, showSensitiveDataParam *bool) (model.CentralNodeConfigProfile, error)

	// Returns list of all Central Node Config profiles.
	// @return com.vmware.nsx.model.CentralNodeConfigProfileListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List() (model.CentralNodeConfigProfileListResult, error)

	// Updates properties in the specified Central Node Config profile.
	//
	// @param nodeConfigProfileIdParam (required)
	// @param centralNodeConfigProfileParam (required)
	// @return com.vmware.nsx.model.CentralNodeConfigProfile
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(nodeConfigProfileIdParam string, centralNodeConfigProfileParam model.CentralNodeConfigProfile) (model.CentralNodeConfigProfile, error)
}

type nodeConfigProfilesClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewNodeConfigProfilesClient(connector client.Connector) *nodeConfigProfilesClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.configs.central_config.node_config_profiles")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get":    core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   core.NewMethodIdentifier(interfaceIdentifier, "list"),
		"update": core.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	nIface := nodeConfigProfilesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &nIface
}

func (nIface *nodeConfigProfilesClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := nIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (nIface *nodeConfigProfilesClient) Get(profileIdParam string, showSensitiveDataParam *bool) (model.CentralNodeConfigProfile, error) {
	typeConverter := nIface.connector.TypeConverter()
	executionContext := nIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(nodeConfigProfilesGetInputType(), typeConverter)
	sv.AddStructField("ProfileId", profileIdParam)
	sv.AddStructField("ShowSensitiveData", showSensitiveDataParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.CentralNodeConfigProfile
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := nodeConfigProfilesGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	nIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := nIface.connector.GetApiProvider().Invoke("com.vmware.nsx.configs.central_config.node_config_profiles", "get", inputDataValue, executionContext)
	var emptyOutput model.CentralNodeConfigProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), nodeConfigProfilesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.CentralNodeConfigProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (nIface *nodeConfigProfilesClient) List() (model.CentralNodeConfigProfileListResult, error) {
	typeConverter := nIface.connector.TypeConverter()
	executionContext := nIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(nodeConfigProfilesListInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.CentralNodeConfigProfileListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := nodeConfigProfilesListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	nIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := nIface.connector.GetApiProvider().Invoke("com.vmware.nsx.configs.central_config.node_config_profiles", "list", inputDataValue, executionContext)
	var emptyOutput model.CentralNodeConfigProfileListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), nodeConfigProfilesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.CentralNodeConfigProfileListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (nIface *nodeConfigProfilesClient) Update(nodeConfigProfileIdParam string, centralNodeConfigProfileParam model.CentralNodeConfigProfile) (model.CentralNodeConfigProfile, error) {
	typeConverter := nIface.connector.TypeConverter()
	executionContext := nIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(nodeConfigProfilesUpdateInputType(), typeConverter)
	sv.AddStructField("NodeConfigProfileId", nodeConfigProfileIdParam)
	sv.AddStructField("CentralNodeConfigProfile", centralNodeConfigProfileParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.CentralNodeConfigProfile
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := nodeConfigProfilesUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	nIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := nIface.connector.GetApiProvider().Invoke("com.vmware.nsx.configs.central_config.node_config_profiles", "update", inputDataValue, executionContext)
	var emptyOutput model.CentralNodeConfigProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), nodeConfigProfilesUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.CentralNodeConfigProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
