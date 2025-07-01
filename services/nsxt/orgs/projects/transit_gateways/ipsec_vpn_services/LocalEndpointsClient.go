// Copyright (c) 2019-2025 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: LocalEndpoints
// Used by client-side stubs.

package ipsec_vpn_services

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type LocalEndpointsClient interface {

	// Delete IPSec VPN local endpoint for a given ipsec vpn service under Transit Gateway.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param transitGatewayIdParam (required)
	// @param ipsecVpnServiceIdParam (required)
	// @param localEndpointIdParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(orgIdParam string, projectIdParam string, transitGatewayIdParam string, ipsecVpnServiceIdParam string, localEndpointIdParam string) error

	// Get IPSec VPN local endpoint for a given ipsec vpn service under Transit Gateway.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param transitGatewayIdParam (required)
	// @param ipsecVpnServiceIdParam (required)
	// @param localEndpointIdParam (required)
	// @return com.vmware.nsx_policy.model.IPSecVpnLocalEndpoint
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(orgIdParam string, projectIdParam string, transitGatewayIdParam string, ipsecVpnServiceIdParam string, localEndpointIdParam string) (nsx_policyModel.IPSecVpnLocalEndpoint, error)

	// Get paginated list of all IPSec VPN local endpoints for a given ipsec vpn service under Transit Gateway.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param transitGatewayIdParam (required)
	// @param ipsecVpnServiceIdParam (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includeMarkForDeleteObjectsParam If true, resources that are marked for deletion will be included in the results. By default, these resources are not included. (optional, default to false)
	// @param includedFieldsParam Note - this parameter currently only works when used with the search APIs /policy/api/v1/search/query and /policy/api/v1/search/dsl. It is ignored for other list APIs. (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam If true, results are sorted in ascending order (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx_policy.model.IPSecVpnLocalEndpointListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(orgIdParam string, projectIdParam string, transitGatewayIdParam string, ipsecVpnServiceIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_policyModel.IPSecVpnLocalEndpointListResult, error)

	// Create or patch a custom IPSec VPN local endpoint under Transit Gateway.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param transitGatewayIdParam (required)
	// @param ipsecVpnServiceIdParam (required)
	// @param localEndpointIdParam (required)
	// @param ipSecVpnLocalEndpointParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(orgIdParam string, projectIdParam string, transitGatewayIdParam string, ipsecVpnServiceIdParam string, localEndpointIdParam string, ipSecVpnLocalEndpointParam nsx_policyModel.IPSecVpnLocalEndpoint) error

	// Create or fully replace IPSec VPN local endpoint for a given ipsec vpn service under Transit Gateway. Revision is optional for creation and required for update.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param transitGatewayIdParam (required)
	// @param ipsecVpnServiceIdParam (required)
	// @param localEndpointIdParam (required)
	// @param ipSecVpnLocalEndpointParam (required)
	// @return com.vmware.nsx_policy.model.IPSecVpnLocalEndpoint
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(orgIdParam string, projectIdParam string, transitGatewayIdParam string, ipsecVpnServiceIdParam string, localEndpointIdParam string, ipSecVpnLocalEndpointParam nsx_policyModel.IPSecVpnLocalEndpoint) (nsx_policyModel.IPSecVpnLocalEndpoint, error)
}

type localEndpointsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewLocalEndpointsClient(connector vapiProtocolClient_.Connector) *localEndpointsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.orgs.projects.transit_gateways.ipsec_vpn_services.local_endpoints")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"delete": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"patch":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	lIface := localEndpointsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &lIface
}

func (lIface *localEndpointsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := lIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (lIface *localEndpointsClient) Delete(orgIdParam string, projectIdParam string, transitGatewayIdParam string, ipsecVpnServiceIdParam string, localEndpointIdParam string) error {
	typeConverter := lIface.connector.TypeConverter()
	executionContext := lIface.connector.NewExecutionContext()
	operationRestMetaData := localEndpointsDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(localEndpointsDeleteInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("TransitGatewayId", transitGatewayIdParam)
	sv.AddStructField("IpsecVpnServiceId", ipsecVpnServiceIdParam)
	sv.AddStructField("LocalEndpointId", localEndpointIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := lIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.transit_gateways.ipsec_vpn_services.local_endpoints", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (lIface *localEndpointsClient) Get(orgIdParam string, projectIdParam string, transitGatewayIdParam string, ipsecVpnServiceIdParam string, localEndpointIdParam string) (nsx_policyModel.IPSecVpnLocalEndpoint, error) {
	typeConverter := lIface.connector.TypeConverter()
	executionContext := lIface.connector.NewExecutionContext()
	operationRestMetaData := localEndpointsGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(localEndpointsGetInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("TransitGatewayId", transitGatewayIdParam)
	sv.AddStructField("IpsecVpnServiceId", ipsecVpnServiceIdParam)
	sv.AddStructField("LocalEndpointId", localEndpointIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.IPSecVpnLocalEndpoint
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := lIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.transit_gateways.ipsec_vpn_services.local_endpoints", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.IPSecVpnLocalEndpoint
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), LocalEndpointsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.IPSecVpnLocalEndpoint), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (lIface *localEndpointsClient) List(orgIdParam string, projectIdParam string, transitGatewayIdParam string, ipsecVpnServiceIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_policyModel.IPSecVpnLocalEndpointListResult, error) {
	typeConverter := lIface.connector.TypeConverter()
	executionContext := lIface.connector.NewExecutionContext()
	operationRestMetaData := localEndpointsListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(localEndpointsListInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("TransitGatewayId", transitGatewayIdParam)
	sv.AddStructField("IpsecVpnServiceId", ipsecVpnServiceIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludeMarkForDeleteObjects", includeMarkForDeleteObjectsParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.IPSecVpnLocalEndpointListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := lIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.transit_gateways.ipsec_vpn_services.local_endpoints", "list", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.IPSecVpnLocalEndpointListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), LocalEndpointsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.IPSecVpnLocalEndpointListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (lIface *localEndpointsClient) Patch(orgIdParam string, projectIdParam string, transitGatewayIdParam string, ipsecVpnServiceIdParam string, localEndpointIdParam string, ipSecVpnLocalEndpointParam nsx_policyModel.IPSecVpnLocalEndpoint) error {
	typeConverter := lIface.connector.TypeConverter()
	executionContext := lIface.connector.NewExecutionContext()
	operationRestMetaData := localEndpointsPatchRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(localEndpointsPatchInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("TransitGatewayId", transitGatewayIdParam)
	sv.AddStructField("IpsecVpnServiceId", ipsecVpnServiceIdParam)
	sv.AddStructField("LocalEndpointId", localEndpointIdParam)
	sv.AddStructField("IpSecVpnLocalEndpoint", ipSecVpnLocalEndpointParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := lIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.transit_gateways.ipsec_vpn_services.local_endpoints", "patch", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (lIface *localEndpointsClient) Update(orgIdParam string, projectIdParam string, transitGatewayIdParam string, ipsecVpnServiceIdParam string, localEndpointIdParam string, ipSecVpnLocalEndpointParam nsx_policyModel.IPSecVpnLocalEndpoint) (nsx_policyModel.IPSecVpnLocalEndpoint, error) {
	typeConverter := lIface.connector.TypeConverter()
	executionContext := lIface.connector.NewExecutionContext()
	operationRestMetaData := localEndpointsUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(localEndpointsUpdateInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("TransitGatewayId", transitGatewayIdParam)
	sv.AddStructField("IpsecVpnServiceId", ipsecVpnServiceIdParam)
	sv.AddStructField("LocalEndpointId", localEndpointIdParam)
	sv.AddStructField("IpSecVpnLocalEndpoint", ipSecVpnLocalEndpointParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.IPSecVpnLocalEndpoint
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := lIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.transit_gateways.ipsec_vpn_services.local_endpoints", "update", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.IPSecVpnLocalEndpoint
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), LocalEndpointsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.IPSecVpnLocalEndpoint), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
