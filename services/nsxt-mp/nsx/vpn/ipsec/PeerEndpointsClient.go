// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: PeerEndpoints
// Used by client-side stubs.

package ipsec

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type PeerEndpointsClient interface {

	// Create custom IPSec peer endpoint.
	//
	// @param ipSecVPNPeerEndpointParam (required)
	// @return com.vmware.nsx.model.IPSecVPNPeerEndpoint
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(ipSecVPNPeerEndpointParam model.IPSecVPNPeerEndpoint) (model.IPSecVPNPeerEndpoint, error)

	// Delete custom IPSec VPN peer endpoint. All references are strong references and dependent peer endpoints can not be deleted if being referenced.
	//
	// @param ipsecVpnPeerEndpointIdParam (required)
	// @param forceParam Force delete the resource even if it is being used somewhere (optional, default to false)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(ipsecVpnPeerEndpointIdParam string, forceParam *bool) error

	// Get custom IPSec VPN peer endpoint.
	//
	// @param ipsecVpnPeerEndpointIdParam (required)
	// @return com.vmware.nsx.model.IPSecVPNPeerEndpoint
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(ipsecVpnPeerEndpointIdParam string) (model.IPSecVPNPeerEndpoint, error)

	// Get paginated list of all peer endpoint.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx.model.IPSecVPNPeerEndpointListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.IPSecVPNPeerEndpointListResult, error)

	// Get custom IPSec VPN peer endpoint with PSK.
	//
	// @param ipsecVpnPeerEndpointIdParam (required)
	// @return com.vmware.nsx.model.IPSecVPNPeerEndpoint
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Showsensitivedata(ipsecVpnPeerEndpointIdParam string) (model.IPSecVPNPeerEndpoint, error)

	// Edit custom IPSec peer endpoint. System owned endpoints are non editable.
	//
	// @param ipsecVpnPeerEndpointIdParam (required)
	// @param ipSecVPNPeerEndpointParam (required)
	// @return com.vmware.nsx.model.IPSecVPNPeerEndpoint
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(ipsecVpnPeerEndpointIdParam string, ipSecVPNPeerEndpointParam model.IPSecVPNPeerEndpoint) (model.IPSecVPNPeerEndpoint, error)
}

type peerEndpointsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewPeerEndpointsClient(connector client.Connector) *peerEndpointsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.vpn.ipsec.peer_endpoints")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create":            core.NewMethodIdentifier(interfaceIdentifier, "create"),
		"delete":            core.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":               core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":              core.NewMethodIdentifier(interfaceIdentifier, "list"),
		"showsensitivedata": core.NewMethodIdentifier(interfaceIdentifier, "showsensitivedata"),
		"update":            core.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	pIface := peerEndpointsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &pIface
}

func (pIface *peerEndpointsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := pIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (pIface *peerEndpointsClient) Create(ipSecVPNPeerEndpointParam model.IPSecVPNPeerEndpoint) (model.IPSecVPNPeerEndpoint, error) {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(peerEndpointsCreateInputType(), typeConverter)
	sv.AddStructField("IpSecVPNPeerEndpoint", ipSecVPNPeerEndpointParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IPSecVPNPeerEndpoint
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := peerEndpointsCreateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx.vpn.ipsec.peer_endpoints", "create", inputDataValue, executionContext)
	var emptyOutput model.IPSecVPNPeerEndpoint
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), peerEndpointsCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IPSecVPNPeerEndpoint), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (pIface *peerEndpointsClient) Delete(ipsecVpnPeerEndpointIdParam string, forceParam *bool) error {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(peerEndpointsDeleteInputType(), typeConverter)
	sv.AddStructField("IpsecVpnPeerEndpointId", ipsecVpnPeerEndpointIdParam)
	sv.AddStructField("Force", forceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := peerEndpointsDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx.vpn.ipsec.peer_endpoints", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (pIface *peerEndpointsClient) Get(ipsecVpnPeerEndpointIdParam string) (model.IPSecVPNPeerEndpoint, error) {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(peerEndpointsGetInputType(), typeConverter)
	sv.AddStructField("IpsecVpnPeerEndpointId", ipsecVpnPeerEndpointIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IPSecVPNPeerEndpoint
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := peerEndpointsGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx.vpn.ipsec.peer_endpoints", "get", inputDataValue, executionContext)
	var emptyOutput model.IPSecVPNPeerEndpoint
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), peerEndpointsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IPSecVPNPeerEndpoint), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (pIface *peerEndpointsClient) List(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.IPSecVPNPeerEndpointListResult, error) {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(peerEndpointsListInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IPSecVPNPeerEndpointListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := peerEndpointsListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx.vpn.ipsec.peer_endpoints", "list", inputDataValue, executionContext)
	var emptyOutput model.IPSecVPNPeerEndpointListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), peerEndpointsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IPSecVPNPeerEndpointListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (pIface *peerEndpointsClient) Showsensitivedata(ipsecVpnPeerEndpointIdParam string) (model.IPSecVPNPeerEndpoint, error) {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(peerEndpointsShowsensitivedataInputType(), typeConverter)
	sv.AddStructField("IpsecVpnPeerEndpointId", ipsecVpnPeerEndpointIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IPSecVPNPeerEndpoint
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := peerEndpointsShowsensitivedataRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx.vpn.ipsec.peer_endpoints", "showsensitivedata", inputDataValue, executionContext)
	var emptyOutput model.IPSecVPNPeerEndpoint
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), peerEndpointsShowsensitivedataOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IPSecVPNPeerEndpoint), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (pIface *peerEndpointsClient) Update(ipsecVpnPeerEndpointIdParam string, ipSecVPNPeerEndpointParam model.IPSecVPNPeerEndpoint) (model.IPSecVPNPeerEndpoint, error) {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(peerEndpointsUpdateInputType(), typeConverter)
	sv.AddStructField("IpsecVpnPeerEndpointId", ipsecVpnPeerEndpointIdParam)
	sv.AddStructField("IpSecVPNPeerEndpoint", ipSecVPNPeerEndpointParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IPSecVPNPeerEndpoint
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := peerEndpointsUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx.vpn.ipsec.peer_endpoints", "update", inputDataValue, executionContext)
	var emptyOutput model.IPSecVPNPeerEndpoint
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), peerEndpointsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IPSecVPNPeerEndpoint), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
