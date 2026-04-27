// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Routing
// Used by client-side stubs.

package transit_gateways

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type RoutingClient interface {

	// Delete the routing configuration for the specified transit gateway.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param transitGatewayIdParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(orgIdParam string, projectIdParam string, transitGatewayIdParam string) error

	// Read the routing configuration for the specified transit gateway.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param transitGatewayIdParam (required)
	// @return com.vmware.nsx_policy.model.TransitGatewayRoutingConfig
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(orgIdParam string, projectIdParam string, transitGatewayIdParam string) (nsx_policyModel.TransitGatewayRoutingConfig, error)

	// Partially update the routing configuration for the specified transit gateway. Only the provided fields are updated.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param transitGatewayIdParam (required)
	// @param transitGatewayRoutingConfigParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(orgIdParam string, projectIdParam string, transitGatewayIdParam string, transitGatewayRoutingConfigParam nsx_policyModel.TransitGatewayRoutingConfig) error

	// Create or replace the routing configuration for the specified transit gateway. Controls forwarding timers and route redistribution rules for BGP neighbors on centralized network attachment connections.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param transitGatewayIdParam (required)
	// @param transitGatewayRoutingConfigParam (required)
	// @return com.vmware.nsx_policy.model.TransitGatewayRoutingConfig
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(orgIdParam string, projectIdParam string, transitGatewayIdParam string, transitGatewayRoutingConfigParam nsx_policyModel.TransitGatewayRoutingConfig) (nsx_policyModel.TransitGatewayRoutingConfig, error)
}

type routingClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewRoutingClient(connector vapiProtocolClient_.Connector) *routingClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.orgs.projects.transit_gateways.routing")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"delete": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"patch":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	rIface := routingClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &rIface
}

func (rIface *routingClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := rIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (rIface *routingClient) Delete(orgIdParam string, projectIdParam string, transitGatewayIdParam string) error {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	operationRestMetaData := routingDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(routingDeleteInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("TransitGatewayId", transitGatewayIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.transit_gateways.routing", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (rIface *routingClient) Get(orgIdParam string, projectIdParam string, transitGatewayIdParam string) (nsx_policyModel.TransitGatewayRoutingConfig, error) {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	operationRestMetaData := routingGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(routingGetInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("TransitGatewayId", transitGatewayIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.TransitGatewayRoutingConfig
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.transit_gateways.routing", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.TransitGatewayRoutingConfig
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), RoutingGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.TransitGatewayRoutingConfig), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (rIface *routingClient) Patch(orgIdParam string, projectIdParam string, transitGatewayIdParam string, transitGatewayRoutingConfigParam nsx_policyModel.TransitGatewayRoutingConfig) error {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	operationRestMetaData := routingPatchRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(routingPatchInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("TransitGatewayId", transitGatewayIdParam)
	sv.AddStructField("TransitGatewayRoutingConfig", transitGatewayRoutingConfigParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.transit_gateways.routing", "patch", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (rIface *routingClient) Update(orgIdParam string, projectIdParam string, transitGatewayIdParam string, transitGatewayRoutingConfigParam nsx_policyModel.TransitGatewayRoutingConfig) (nsx_policyModel.TransitGatewayRoutingConfig, error) {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	operationRestMetaData := routingUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(routingUpdateInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("TransitGatewayId", transitGatewayIdParam)
	sv.AddStructField("TransitGatewayRoutingConfig", transitGatewayRoutingConfigParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.TransitGatewayRoutingConfig
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.transit_gateways.routing", "update", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.TransitGatewayRoutingConfig
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), RoutingUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.TransitGatewayRoutingConfig), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
