// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: LocalEndpoints
// Used by client-side stubs.

package ipsec

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type LocalEndpointsClient interface {

	// Create custom IPSec local endpoint.
	//
	//  Please use below Policy APIs.
	//  PATCH /policy/api/v1/infra/tier-0s/<tier-0-id>/ipsec-vpn-services/<service-id>/local-endpoints/<local-endpoint-id>
	//  PATCH /policy/api/v1/infra/tier-1s/<tier-1-id>/ipsec-vpn-services/<service-id>/local-endpoints/<local-endpoint-id>
	//
	// Deprecated: This API element is deprecated.
	//
	// @param ipSecVPNLocalEndpointParam (required)
	// @return com.vmware.nsx.model.IPSecVPNLocalEndpoint
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(ipSecVPNLocalEndpointParam nsxModel.IPSecVPNLocalEndpoint) (nsxModel.IPSecVPNLocalEndpoint, error)

	// Delete custom IPSec local endpoint.
	//
	//  Please use below Policy APIs.
	//  DELETE /policy/api/v1/infra/tier-0s/<tier-0-id>/ipsec-vpn-services/<service-id>/local-endpoints/<local-endpoint-id>
	//  DELETE /policy/api/v1/infra/tier-1s/<tier-1-id>/ipsec-vpn-services/<service-id>/local-endpoints/<local-endpoint-id>
	//
	// Deprecated: This API element is deprecated.
	//
	// @param ipsecVpnLocalEndpointIdParam (required)
	// @param forceParam Force delete the resource even if it is being used somewhere (optional, default to false)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(ipsecVpnLocalEndpointIdParam string, forceParam *bool) error

	// Get custom IPSec local endpoint.
	//
	//  Please use below Policy APIs.
	//  GET /policy/api/v1/infra/tier-0s/<tier-0-id>/ipsec-vpn-services/<service-id>/local-endpoints/<local-endpoint-id>
	//  GET /policy/api/v1/infra/tier-1s/<tier-1-id>/ipsec-vpn-services/<service-id>/local-endpoints/<local-endpoint-id>
	//
	// Deprecated: This API element is deprecated.
	//
	// @param ipsecVpnLocalEndpointIdParam (required)
	// @return com.vmware.nsx.model.IPSecVPNLocalEndpoint
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(ipsecVpnLocalEndpointIdParam string) (nsxModel.IPSecVPNLocalEndpoint, error)

	// Get paginated list of all local endpoints.
	//
	//  Please use below Policy APIs.
	//  GET /policy/api/v1/infra/tier-0s/<tier-0-id>/ipsec-vpn-services/<service-id>/local-endpoints
	//  GET /policy/api/v1/infra/tier-1s/<tier-1-id>/ipsec-vpn-services/<service-id>/local-endpoints
	//
	// Deprecated: This API element is deprecated.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param ipsecVpnServiceIdParam Id of the IPSec VPN service (optional)
	// @param logicalRouterIdParam Id of logical router (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx.model.IPSecVPNLocalEndpointListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(cursorParam *string, includedFieldsParam *string, ipsecVpnServiceIdParam *string, logicalRouterIdParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsxModel.IPSecVPNLocalEndpointListResult, error)

	// Edit custom IPSec local endpoint.
	//
	//  Please use below Policy APIs.
	//  PUT /policy/api/v1/infra/tier-0s/<tier-0-id>/ipsec-vpn-services/<service-id>/local-endpoints/<local-endpoint-id>
	//  PUT /policy/api/v1/infra/tier-1s/<tier-1-id>/ipsec-vpn-services/<service-id>/local-endpoints/<local-endpoint-id>
	//
	// Deprecated: This API element is deprecated.
	//
	// @param ipsecVpnLocalEndpointIdParam (required)
	// @param ipSecVPNLocalEndpointParam (required)
	// @return com.vmware.nsx.model.IPSecVPNLocalEndpoint
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(ipsecVpnLocalEndpointIdParam string, ipSecVPNLocalEndpointParam nsxModel.IPSecVPNLocalEndpoint) (nsxModel.IPSecVPNLocalEndpoint, error)
}

type localEndpointsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewLocalEndpointsClient(connector vapiProtocolClient_.Connector) *localEndpointsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.vpn.ipsec.local_endpoints")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"create": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "create"),
		"delete": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
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

func (lIface *localEndpointsClient) Create(ipSecVPNLocalEndpointParam nsxModel.IPSecVPNLocalEndpoint) (nsxModel.IPSecVPNLocalEndpoint, error) {
	typeConverter := lIface.connector.TypeConverter()
	executionContext := lIface.connector.NewExecutionContext()
	operationRestMetaData := localEndpointsCreateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(localEndpointsCreateInputType(), typeConverter)
	sv.AddStructField("IpSecVPNLocalEndpoint", ipSecVPNLocalEndpointParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.IPSecVPNLocalEndpoint
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := lIface.connector.GetApiProvider().Invoke("com.vmware.nsx.vpn.ipsec.local_endpoints", "create", inputDataValue, executionContext)
	var emptyOutput nsxModel.IPSecVPNLocalEndpoint
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), LocalEndpointsCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.IPSecVPNLocalEndpoint), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (lIface *localEndpointsClient) Delete(ipsecVpnLocalEndpointIdParam string, forceParam *bool) error {
	typeConverter := lIface.connector.TypeConverter()
	executionContext := lIface.connector.NewExecutionContext()
	operationRestMetaData := localEndpointsDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(localEndpointsDeleteInputType(), typeConverter)
	sv.AddStructField("IpsecVpnLocalEndpointId", ipsecVpnLocalEndpointIdParam)
	sv.AddStructField("Force", forceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := lIface.connector.GetApiProvider().Invoke("com.vmware.nsx.vpn.ipsec.local_endpoints", "delete", inputDataValue, executionContext)
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

func (lIface *localEndpointsClient) Get(ipsecVpnLocalEndpointIdParam string) (nsxModel.IPSecVPNLocalEndpoint, error) {
	typeConverter := lIface.connector.TypeConverter()
	executionContext := lIface.connector.NewExecutionContext()
	operationRestMetaData := localEndpointsGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(localEndpointsGetInputType(), typeConverter)
	sv.AddStructField("IpsecVpnLocalEndpointId", ipsecVpnLocalEndpointIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.IPSecVPNLocalEndpoint
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := lIface.connector.GetApiProvider().Invoke("com.vmware.nsx.vpn.ipsec.local_endpoints", "get", inputDataValue, executionContext)
	var emptyOutput nsxModel.IPSecVPNLocalEndpoint
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), LocalEndpointsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.IPSecVPNLocalEndpoint), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (lIface *localEndpointsClient) List(cursorParam *string, includedFieldsParam *string, ipsecVpnServiceIdParam *string, logicalRouterIdParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsxModel.IPSecVPNLocalEndpointListResult, error) {
	typeConverter := lIface.connector.TypeConverter()
	executionContext := lIface.connector.NewExecutionContext()
	operationRestMetaData := localEndpointsListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(localEndpointsListInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("IpsecVpnServiceId", ipsecVpnServiceIdParam)
	sv.AddStructField("LogicalRouterId", logicalRouterIdParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.IPSecVPNLocalEndpointListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := lIface.connector.GetApiProvider().Invoke("com.vmware.nsx.vpn.ipsec.local_endpoints", "list", inputDataValue, executionContext)
	var emptyOutput nsxModel.IPSecVPNLocalEndpointListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), LocalEndpointsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.IPSecVPNLocalEndpointListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (lIface *localEndpointsClient) Update(ipsecVpnLocalEndpointIdParam string, ipSecVPNLocalEndpointParam nsxModel.IPSecVPNLocalEndpoint) (nsxModel.IPSecVPNLocalEndpoint, error) {
	typeConverter := lIface.connector.TypeConverter()
	executionContext := lIface.connector.NewExecutionContext()
	operationRestMetaData := localEndpointsUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(localEndpointsUpdateInputType(), typeConverter)
	sv.AddStructField("IpsecVpnLocalEndpointId", ipsecVpnLocalEndpointIdParam)
	sv.AddStructField("IpSecVPNLocalEndpoint", ipSecVPNLocalEndpointParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.IPSecVPNLocalEndpoint
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := lIface.connector.GetApiProvider().Invoke("com.vmware.nsx.vpn.ipsec.local_endpoints", "update", inputDataValue, executionContext)
	var emptyOutput nsxModel.IPSecVPNLocalEndpoint
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), LocalEndpointsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.IPSecVPNLocalEndpoint), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
