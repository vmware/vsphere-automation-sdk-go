// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Bgp
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

type BgpClient interface {

	// Delete the BGP routing configuration for the specified transit gateway.
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

	// Read the BGP routing configuration for the specified transit gateway.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param transitGatewayIdParam (required)
	// @return com.vmware.nsx_policy.model.BgpRoutingConfig
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(orgIdParam string, projectIdParam string, transitGatewayIdParam string) (nsx_policyModel.BgpRoutingConfig, error)

	// Partially update the BGP routing configuration for the specified transit gateway. Only the provided fields are updated.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param transitGatewayIdParam (required)
	// @param bgpRoutingConfigParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(orgIdParam string, projectIdParam string, transitGatewayIdParam string, bgpRoutingConfigParam nsx_policyModel.BgpRoutingConfig) error

	// Create or replace the BGP routing configuration for the specified transit gateway. Uses the same BGP routing config schema as tier-0 gateways.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param transitGatewayIdParam (required)
	// @param bgpRoutingConfigParam (required)
	// @return com.vmware.nsx_policy.model.BgpRoutingConfig
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(orgIdParam string, projectIdParam string, transitGatewayIdParam string, bgpRoutingConfigParam nsx_policyModel.BgpRoutingConfig) (nsx_policyModel.BgpRoutingConfig, error)
}

type bgpClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewBgpClient(connector vapiProtocolClient_.Connector) *bgpClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.orgs.projects.transit_gateways.bgp")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"delete": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"patch":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	bIface := bgpClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &bIface
}

func (bIface *bgpClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := bIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (bIface *bgpClient) Delete(orgIdParam string, projectIdParam string, transitGatewayIdParam string) error {
	typeConverter := bIface.connector.TypeConverter()
	executionContext := bIface.connector.NewExecutionContext()
	operationRestMetaData := bgpDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(bgpDeleteInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("TransitGatewayId", transitGatewayIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := bIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.transit_gateways.bgp", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), bIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (bIface *bgpClient) Get(orgIdParam string, projectIdParam string, transitGatewayIdParam string) (nsx_policyModel.BgpRoutingConfig, error) {
	typeConverter := bIface.connector.TypeConverter()
	executionContext := bIface.connector.NewExecutionContext()
	operationRestMetaData := bgpGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(bgpGetInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("TransitGatewayId", transitGatewayIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.BgpRoutingConfig
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := bIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.transit_gateways.bgp", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.BgpRoutingConfig
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), BgpGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.BgpRoutingConfig), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), bIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (bIface *bgpClient) Patch(orgIdParam string, projectIdParam string, transitGatewayIdParam string, bgpRoutingConfigParam nsx_policyModel.BgpRoutingConfig) error {
	typeConverter := bIface.connector.TypeConverter()
	executionContext := bIface.connector.NewExecutionContext()
	operationRestMetaData := bgpPatchRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(bgpPatchInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("TransitGatewayId", transitGatewayIdParam)
	sv.AddStructField("BgpRoutingConfig", bgpRoutingConfigParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := bIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.transit_gateways.bgp", "patch", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), bIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (bIface *bgpClient) Update(orgIdParam string, projectIdParam string, transitGatewayIdParam string, bgpRoutingConfigParam nsx_policyModel.BgpRoutingConfig) (nsx_policyModel.BgpRoutingConfig, error) {
	typeConverter := bIface.connector.TypeConverter()
	executionContext := bIface.connector.NewExecutionContext()
	operationRestMetaData := bgpUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(bgpUpdateInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("TransitGatewayId", transitGatewayIdParam)
	sv.AddStructField("BgpRoutingConfig", bgpRoutingConfigParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.BgpRoutingConfig
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := bIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.transit_gateways.bgp", "update", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.BgpRoutingConfig
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), BgpUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.BgpRoutingConfig), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), bIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
