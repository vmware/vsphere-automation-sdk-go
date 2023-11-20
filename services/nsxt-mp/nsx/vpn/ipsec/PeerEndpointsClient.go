// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: PeerEndpoints
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

type PeerEndpointsClient interface {

	// Create custom IPSec peer endpoint.
	//
	//  Please use below Policy APIs.
	//  PATCH /policy/api/v1/infra/tier-0s/<tier-0-id>/ipsec-vpn-services/<service-id>/sessions/<session-id>
	//  PATCH /policy/api/v1/infra/tier-1s/<tier-1-id>/ipsec-vpn-services/<service-id>/sessions/<session-id>
	//
	// Deprecated: This API element is deprecated.
	//
	// @param ipSecVPNPeerEndpointParam (required)
	// @return com.vmware.nsx.model.IPSecVPNPeerEndpoint
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(ipSecVPNPeerEndpointParam nsxModel.IPSecVPNPeerEndpoint) (nsxModel.IPSecVPNPeerEndpoint, error)

	// Delete custom IPSec VPN peer endpoint. All references are strong references and dependent peer endpoints can not be deleted if being referenced.
	//
	//  Please use below Policy APIs.
	//  DELETE /policy/api/v1/infra/tier-0s/<tier-0-id>/ipsec-vpn-services/<service-id>/sessions/<session-id>
	//  DELETE /policy/api/v1/infra/tier-1s/<tier-1-id>/ipsec-vpn-services/<service-id>/sessions/<session-id>
	//
	// Deprecated: This API element is deprecated.
	//
	// @param ipsecVpnPeerEndpointIdParam (required)
	// @param forceParam Force delete the resource even if it is being used somewhere (optional, default to false)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(ipsecVpnPeerEndpointIdParam string, forceParam *bool) error

	// Get custom IPSec VPN peer endpoint.
	//
	//  Please use below Policy APIs.
	//  GET /policy/api/v1/infra/tier-0s/<tier-0-id>/ipsec-vpn-services/<service-id>/sessions/<session-id>
	//  GET /policy/api/v1/infra/tier-1s/<tier-1-id>/ipsec-vpn-services/<service-id>/sessions/<session-id>
	//
	// Deprecated: This API element is deprecated.
	//
	// @param ipsecVpnPeerEndpointIdParam (required)
	// @return com.vmware.nsx.model.IPSecVPNPeerEndpoint
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(ipsecVpnPeerEndpointIdParam string) (nsxModel.IPSecVPNPeerEndpoint, error)

	// Get paginated list of all peer endpoint.
	//
	//  Please use below Policy APIs.
	//  GET /policy/api/v1/infra/tier-0s/<tier-0-id>/ipsec-vpn-services/<service-id>/sessions
	//  GET /policy/api/v1/infra/tier-1s/<tier-1-id>/ipsec-vpn-services/<service-id>/sessions
	//
	// Deprecated: This API element is deprecated.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx.model.IPSecVPNPeerEndpointListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsxModel.IPSecVPNPeerEndpointListResult, error)

	// Get custom IPSec VPN peer endpoint with PSK.
	//
	//  Please use below Policy APIs.
	//  GET /policy/api/v1/infra/tier-0s/<tier-0-id>/ipsec-vpn-services/<service-id>/sessions/<session-id>?action=show-sensitive-data
	//  GET /policy/api/v1/infra/tier-1s/<tier-1-id>/ipsec-vpn-services/<service-id>/sessions/<session-id>?action=show-sensitive-data
	//
	// Deprecated: This API element is deprecated.
	//
	// @param ipsecVpnPeerEndpointIdParam (required)
	// @return com.vmware.nsx.model.IPSecVPNPeerEndpoint
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Showsensitivedata(ipsecVpnPeerEndpointIdParam string) (nsxModel.IPSecVPNPeerEndpoint, error)

	// Edit custom IPSec peer endpoint. System owned endpoints are non editable.
	//
	//  Please use below Policy APIs.
	//  PUT /policy/api/v1/infra/tier-0s/<tier-0-id>/ipsec-vpn-services/<service-id>/sessions/<session-id>
	//  PUT /policy/api/v1/infra/tier-1s/<tier-1-id>/ipsec-vpn-services/<service-id>/sessions/<session-id>
	//
	// Deprecated: This API element is deprecated.
	//
	// @param ipsecVpnPeerEndpointIdParam (required)
	// @param ipSecVPNPeerEndpointParam (required)
	// @return com.vmware.nsx.model.IPSecVPNPeerEndpoint
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(ipsecVpnPeerEndpointIdParam string, ipSecVPNPeerEndpointParam nsxModel.IPSecVPNPeerEndpoint) (nsxModel.IPSecVPNPeerEndpoint, error)
}

type peerEndpointsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewPeerEndpointsClient(connector vapiProtocolClient_.Connector) *peerEndpointsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.vpn.ipsec.peer_endpoints")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"create":            vapiCore_.NewMethodIdentifier(interfaceIdentifier, "create"),
		"delete":            vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":               vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":              vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"showsensitivedata": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "showsensitivedata"),
		"update":            vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	pIface := peerEndpointsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &pIface
}

func (pIface *peerEndpointsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := pIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (pIface *peerEndpointsClient) Create(ipSecVPNPeerEndpointParam nsxModel.IPSecVPNPeerEndpoint) (nsxModel.IPSecVPNPeerEndpoint, error) {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	operationRestMetaData := peerEndpointsCreateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(peerEndpointsCreateInputType(), typeConverter)
	sv.AddStructField("IpSecVPNPeerEndpoint", ipSecVPNPeerEndpointParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.IPSecVPNPeerEndpoint
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx.vpn.ipsec.peer_endpoints", "create", inputDataValue, executionContext)
	var emptyOutput nsxModel.IPSecVPNPeerEndpoint
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), PeerEndpointsCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.IPSecVPNPeerEndpoint), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (pIface *peerEndpointsClient) Delete(ipsecVpnPeerEndpointIdParam string, forceParam *bool) error {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	operationRestMetaData := peerEndpointsDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(peerEndpointsDeleteInputType(), typeConverter)
	sv.AddStructField("IpsecVpnPeerEndpointId", ipsecVpnPeerEndpointIdParam)
	sv.AddStructField("Force", forceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx.vpn.ipsec.peer_endpoints", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (pIface *peerEndpointsClient) Get(ipsecVpnPeerEndpointIdParam string) (nsxModel.IPSecVPNPeerEndpoint, error) {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	operationRestMetaData := peerEndpointsGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(peerEndpointsGetInputType(), typeConverter)
	sv.AddStructField("IpsecVpnPeerEndpointId", ipsecVpnPeerEndpointIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.IPSecVPNPeerEndpoint
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx.vpn.ipsec.peer_endpoints", "get", inputDataValue, executionContext)
	var emptyOutput nsxModel.IPSecVPNPeerEndpoint
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), PeerEndpointsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.IPSecVPNPeerEndpoint), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (pIface *peerEndpointsClient) List(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsxModel.IPSecVPNPeerEndpointListResult, error) {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	operationRestMetaData := peerEndpointsListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(peerEndpointsListInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.IPSecVPNPeerEndpointListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx.vpn.ipsec.peer_endpoints", "list", inputDataValue, executionContext)
	var emptyOutput nsxModel.IPSecVPNPeerEndpointListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), PeerEndpointsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.IPSecVPNPeerEndpointListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (pIface *peerEndpointsClient) Showsensitivedata(ipsecVpnPeerEndpointIdParam string) (nsxModel.IPSecVPNPeerEndpoint, error) {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	operationRestMetaData := peerEndpointsShowsensitivedataRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(peerEndpointsShowsensitivedataInputType(), typeConverter)
	sv.AddStructField("IpsecVpnPeerEndpointId", ipsecVpnPeerEndpointIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.IPSecVPNPeerEndpoint
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx.vpn.ipsec.peer_endpoints", "showsensitivedata", inputDataValue, executionContext)
	var emptyOutput nsxModel.IPSecVPNPeerEndpoint
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), PeerEndpointsShowsensitivedataOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.IPSecVPNPeerEndpoint), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (pIface *peerEndpointsClient) Update(ipsecVpnPeerEndpointIdParam string, ipSecVPNPeerEndpointParam nsxModel.IPSecVPNPeerEndpoint) (nsxModel.IPSecVPNPeerEndpoint, error) {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	operationRestMetaData := peerEndpointsUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(peerEndpointsUpdateInputType(), typeConverter)
	sv.AddStructField("IpsecVpnPeerEndpointId", ipsecVpnPeerEndpointIdParam)
	sv.AddStructField("IpSecVPNPeerEndpoint", ipSecVPNPeerEndpointParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.IPSecVPNPeerEndpoint
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx.vpn.ipsec.peer_endpoints", "update", inputDataValue, executionContext)
	var emptyOutput nsxModel.IPSecVPNPeerEndpoint
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), PeerEndpointsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.IPSecVPNPeerEndpoint), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
