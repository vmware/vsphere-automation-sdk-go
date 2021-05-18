// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ManagementPlaneAPINetworkingLogicalRoutingAndServicesRoutingBFDConfiguration
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

type ManagementPlaneAPINetworkingLogicalRoutingAndServicesRoutingBFDConfigurationClient interface {

	// Returns the BFD configuration for all routing BFD peers. This will be inherited | by all BFD peers for LogicalRouter unless overriden while configuring the Peer.
	//
	// @param logicalRouterIdParam (required)
	// @return com.vmware.model.BfdConfig
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ReadRoutingBfdConfig(logicalRouterIdParam string) (model.BfdConfig, error)

	// Modifies the BFD configuration for routing BFD peers. Note - the configuration | changes apply only to those routing BFD peers for which the BFD configuration has | not been overridden at Peer level.
	//
	// @param logicalRouterIdParam (required)
	// @param bfdConfigParam (required)
	// @return com.vmware.model.BfdConfig
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateRoutingBfdConfig(logicalRouterIdParam string, bfdConfigParam model.BfdConfig) (model.BfdConfig, error)
}

type managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingBFDConfigurationClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewManagementPlaneAPINetworkingLogicalRoutingAndServicesRoutingBFDConfigurationClient(connector client.Connector) *managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingBFDConfigurationClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.management_plane_API_networking_logical_routing_and_services_routing_BFD_configuration")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"read_routing_bfd_config":   core.NewMethodIdentifier(interfaceIdentifier, "read_routing_bfd_config"),
		"update_routing_bfd_config": core.NewMethodIdentifier(interfaceIdentifier, "update_routing_bfd_config"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	mIface := managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingBFDConfigurationClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &mIface
}

func (mIface *managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingBFDConfigurationClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := mIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (mIface *managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingBFDConfigurationClient) ReadRoutingBfdConfig(logicalRouterIdParam string) (model.BfdConfig, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingBFDConfigurationReadRoutingBfdConfigInputType(), typeConverter)
	sv.AddStructField("LogicalRouterId", logicalRouterIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.BfdConfig
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingBFDConfigurationReadRoutingBfdConfigRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_routing_and_services_routing_BFD_configuration", "read_routing_bfd_config", inputDataValue, executionContext)
	var emptyOutput model.BfdConfig
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingBFDConfigurationReadRoutingBfdConfigOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BfdConfig), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingBFDConfigurationClient) UpdateRoutingBfdConfig(logicalRouterIdParam string, bfdConfigParam model.BfdConfig) (model.BfdConfig, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingBFDConfigurationUpdateRoutingBfdConfigInputType(), typeConverter)
	sv.AddStructField("LogicalRouterId", logicalRouterIdParam)
	sv.AddStructField("BfdConfig", bfdConfigParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.BfdConfig
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingBFDConfigurationUpdateRoutingBfdConfigRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_routing_and_services_routing_BFD_configuration", "update_routing_bfd_config", inputDataValue, executionContext)
	var emptyOutput model.BfdConfig
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingBFDConfigurationUpdateRoutingBfdConfigOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BfdConfig), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
