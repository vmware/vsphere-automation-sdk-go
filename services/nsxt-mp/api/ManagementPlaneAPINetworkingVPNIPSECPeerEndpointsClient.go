// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ManagementPlaneAPINetworkingVPNIPSECPeerEndpoints
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

type ManagementPlaneAPINetworkingVPNIPSECPeerEndpointsClient interface {

	// Create custom IPSec local endpoint.
	//
	// @param ipSecVPNLocalEndpointParam (required)
	// @return com.vmware.model.IPSecVPNLocalEndpoint
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	CreateIPSecVPNLocalEndpoint(ipSecVPNLocalEndpointParam model.IPSecVPNLocalEndpoint) (model.IPSecVPNLocalEndpoint, error)

	// Create custom IPSec peer endpoint.
	//
	// @param ipSecVPNPeerEndpointParam (required)
	// @return com.vmware.model.IPSecVPNPeerEndpoint
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	CreateIPSecVPNPeerEndPoint(ipSecVPNPeerEndpointParam model.IPSecVPNPeerEndpoint) (model.IPSecVPNPeerEndpoint, error)

	// Delete custom IPSec local endpoint.
	//
	// @param ipsecVpnLocalEndpointIdParam (required)
	// @param forceParam Force delete the resource even if it is being used somewhere (optional, default to false)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteIPSecVPNLocalEndpoint(ipsecVpnLocalEndpointIdParam string, forceParam *bool) error

	// Delete custom IPSec VPN peer endpoint. All references are strong references and dependent peer endpoints can not be deleted if being referenced.
	//
	// @param ipsecVpnPeerEndpointIdParam (required)
	// @param forceParam Force delete the resource even if it is being used somewhere (optional, default to false)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteIPSecVPNPeerEndpoint(ipsecVpnPeerEndpointIdParam string, forceParam *bool) error

	// Get custom IPSec local endpoint.
	//
	// @param ipsecVpnLocalEndpointIdParam (required)
	// @return com.vmware.model.IPSecVPNLocalEndpoint
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetIPSecVPNLocalEndpoint(ipsecVpnLocalEndpointIdParam string) (model.IPSecVPNLocalEndpoint, error)

	// Get custom IPSec VPN peer endpoint.
	//
	// @param ipsecVpnPeerEndpointIdParam (required)
	// @return com.vmware.model.IPSecVPNPeerEndpoint
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetIPSecVPNPeerEndpoint(ipsecVpnPeerEndpointIdParam string) (model.IPSecVPNPeerEndpoint, error)

	// Get custom IPSec VPN peer endpoint with PSK.
	//
	// @param ipsecVpnPeerEndpointIdParam (required)
	// @return com.vmware.model.IPSecVPNPeerEndpoint
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetIPSecVPNPeerEndpointWithPSKShowSensitiveData(ipsecVpnPeerEndpointIdParam string) (model.IPSecVPNPeerEndpoint, error)

	// Get paginated list of all local endpoints.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param ipsecVpnServiceIdParam Id of the IPSec VPN service (optional)
	// @param logicalRouterIdParam Id of logical router (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.model.IPSecVPNLocalEndpointListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListIPSecVPNLocalEndpoints(cursorParam *string, includedFieldsParam *string, ipsecVpnServiceIdParam *string, logicalRouterIdParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.IPSecVPNLocalEndpointListResult, error)

	// Get paginated list of all peer endpoint.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.model.IPSecVPNPeerEndpointListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListIPSecVPNPeerEndpoints(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.IPSecVPNPeerEndpointListResult, error)

	// Edit custom IPSec local endpoint.
	//
	// @param ipsecVpnLocalEndpointIdParam (required)
	// @param ipSecVPNLocalEndpointParam (required)
	// @return com.vmware.model.IPSecVPNLocalEndpoint
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateIPSecVPNLocalEndpoint(ipsecVpnLocalEndpointIdParam string, ipSecVPNLocalEndpointParam model.IPSecVPNLocalEndpoint) (model.IPSecVPNLocalEndpoint, error)

	// Edit custom IPSec peer endpoint. System owned endpoints are non editable.
	//
	// @param ipsecVpnPeerEndpointIdParam (required)
	// @param ipSecVPNPeerEndpointParam (required)
	// @return com.vmware.model.IPSecVPNPeerEndpoint
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateIPSecVPNPeerEndpoint(ipsecVpnPeerEndpointIdParam string, ipSecVPNPeerEndpointParam model.IPSecVPNPeerEndpoint) (model.IPSecVPNPeerEndpoint, error)
}

type managementPlaneAPINetworkingVPNIPSECPeerEndpointsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewManagementPlaneAPINetworkingVPNIPSECPeerEndpointsClient(connector client.Connector) *managementPlaneAPINetworkingVPNIPSECPeerEndpointsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.management_plane_API_networking_VPNIPSEC_peer_endpoints")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create_IP_sec_VPN_local_endpoint":                          core.NewMethodIdentifier(interfaceIdentifier, "create_IP_sec_VPN_local_endpoint"),
		"create_IP_sec_VPN_peer_end_point":                          core.NewMethodIdentifier(interfaceIdentifier, "create_IP_sec_VPN_peer_end_point"),
		"delete_IP_sec_VPN_local_endpoint":                          core.NewMethodIdentifier(interfaceIdentifier, "delete_IP_sec_VPN_local_endpoint"),
		"delete_IP_sec_VPN_peer_endpoint":                           core.NewMethodIdentifier(interfaceIdentifier, "delete_IP_sec_VPN_peer_endpoint"),
		"get_IP_sec_VPN_local_endpoint":                             core.NewMethodIdentifier(interfaceIdentifier, "get_IP_sec_VPN_local_endpoint"),
		"get_IP_sec_VPN_peer_endpoint":                              core.NewMethodIdentifier(interfaceIdentifier, "get_IP_sec_VPN_peer_endpoint"),
		"get_IP_sec_VPN_peer_endpoint_with_PSK_show_sensitive_data": core.NewMethodIdentifier(interfaceIdentifier, "get_IP_sec_VPN_peer_endpoint_with_PSK_show_sensitive_data"),
		"list_IP_sec_VPN_local_endpoints":                           core.NewMethodIdentifier(interfaceIdentifier, "list_IP_sec_VPN_local_endpoints"),
		"list_IP_sec_VPN_peer_endpoints":                            core.NewMethodIdentifier(interfaceIdentifier, "list_IP_sec_VPN_peer_endpoints"),
		"update_IP_sec_VPN_local_endpoint":                          core.NewMethodIdentifier(interfaceIdentifier, "update_IP_sec_VPN_local_endpoint"),
		"update_IP_sec_VPN_peer_endpoint":                           core.NewMethodIdentifier(interfaceIdentifier, "update_IP_sec_VPN_peer_endpoint"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	mIface := managementPlaneAPINetworkingVPNIPSECPeerEndpointsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &mIface
}

func (mIface *managementPlaneAPINetworkingVPNIPSECPeerEndpointsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := mIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (mIface *managementPlaneAPINetworkingVPNIPSECPeerEndpointsClient) CreateIPSecVPNLocalEndpoint(ipSecVPNLocalEndpointParam model.IPSecVPNLocalEndpoint) (model.IPSecVPNLocalEndpoint, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingVPNIPSECPeerEndpointsCreateIPSecVPNLocalEndpointInputType(), typeConverter)
	sv.AddStructField("IpSecVPNLocalEndpoint", ipSecVPNLocalEndpointParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IPSecVPNLocalEndpoint
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingVPNIPSECPeerEndpointsCreateIPSecVPNLocalEndpointRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_VPNIPSEC_peer_endpoints", "create_IP_sec_VPN_local_endpoint", inputDataValue, executionContext)
	var emptyOutput model.IPSecVPNLocalEndpoint
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingVPNIPSECPeerEndpointsCreateIPSecVPNLocalEndpointOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IPSecVPNLocalEndpoint), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingVPNIPSECPeerEndpointsClient) CreateIPSecVPNPeerEndPoint(ipSecVPNPeerEndpointParam model.IPSecVPNPeerEndpoint) (model.IPSecVPNPeerEndpoint, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingVPNIPSECPeerEndpointsCreateIPSecVPNPeerEndPointInputType(), typeConverter)
	sv.AddStructField("IpSecVPNPeerEndpoint", ipSecVPNPeerEndpointParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IPSecVPNPeerEndpoint
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingVPNIPSECPeerEndpointsCreateIPSecVPNPeerEndPointRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_VPNIPSEC_peer_endpoints", "create_IP_sec_VPN_peer_end_point", inputDataValue, executionContext)
	var emptyOutput model.IPSecVPNPeerEndpoint
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingVPNIPSECPeerEndpointsCreateIPSecVPNPeerEndPointOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IPSecVPNPeerEndpoint), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingVPNIPSECPeerEndpointsClient) DeleteIPSecVPNLocalEndpoint(ipsecVpnLocalEndpointIdParam string, forceParam *bool) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingVPNIPSECPeerEndpointsDeleteIPSecVPNLocalEndpointInputType(), typeConverter)
	sv.AddStructField("IpsecVpnLocalEndpointId", ipsecVpnLocalEndpointIdParam)
	sv.AddStructField("Force", forceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingVPNIPSECPeerEndpointsDeleteIPSecVPNLocalEndpointRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_VPNIPSEC_peer_endpoints", "delete_IP_sec_VPN_local_endpoint", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingVPNIPSECPeerEndpointsClient) DeleteIPSecVPNPeerEndpoint(ipsecVpnPeerEndpointIdParam string, forceParam *bool) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingVPNIPSECPeerEndpointsDeleteIPSecVPNPeerEndpointInputType(), typeConverter)
	sv.AddStructField("IpsecVpnPeerEndpointId", ipsecVpnPeerEndpointIdParam)
	sv.AddStructField("Force", forceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingVPNIPSECPeerEndpointsDeleteIPSecVPNPeerEndpointRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_VPNIPSEC_peer_endpoints", "delete_IP_sec_VPN_peer_endpoint", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingVPNIPSECPeerEndpointsClient) GetIPSecVPNLocalEndpoint(ipsecVpnLocalEndpointIdParam string) (model.IPSecVPNLocalEndpoint, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingVPNIPSECPeerEndpointsGetIPSecVPNLocalEndpointInputType(), typeConverter)
	sv.AddStructField("IpsecVpnLocalEndpointId", ipsecVpnLocalEndpointIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IPSecVPNLocalEndpoint
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingVPNIPSECPeerEndpointsGetIPSecVPNLocalEndpointRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_VPNIPSEC_peer_endpoints", "get_IP_sec_VPN_local_endpoint", inputDataValue, executionContext)
	var emptyOutput model.IPSecVPNLocalEndpoint
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingVPNIPSECPeerEndpointsGetIPSecVPNLocalEndpointOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IPSecVPNLocalEndpoint), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingVPNIPSECPeerEndpointsClient) GetIPSecVPNPeerEndpoint(ipsecVpnPeerEndpointIdParam string) (model.IPSecVPNPeerEndpoint, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingVPNIPSECPeerEndpointsGetIPSecVPNPeerEndpointInputType(), typeConverter)
	sv.AddStructField("IpsecVpnPeerEndpointId", ipsecVpnPeerEndpointIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IPSecVPNPeerEndpoint
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingVPNIPSECPeerEndpointsGetIPSecVPNPeerEndpointRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_VPNIPSEC_peer_endpoints", "get_IP_sec_VPN_peer_endpoint", inputDataValue, executionContext)
	var emptyOutput model.IPSecVPNPeerEndpoint
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingVPNIPSECPeerEndpointsGetIPSecVPNPeerEndpointOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IPSecVPNPeerEndpoint), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingVPNIPSECPeerEndpointsClient) GetIPSecVPNPeerEndpointWithPSKShowSensitiveData(ipsecVpnPeerEndpointIdParam string) (model.IPSecVPNPeerEndpoint, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingVPNIPSECPeerEndpointsGetIPSecVPNPeerEndpointWithPSKShowSensitiveDataInputType(), typeConverter)
	sv.AddStructField("IpsecVpnPeerEndpointId", ipsecVpnPeerEndpointIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IPSecVPNPeerEndpoint
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingVPNIPSECPeerEndpointsGetIPSecVPNPeerEndpointWithPSKShowSensitiveDataRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_VPNIPSEC_peer_endpoints", "get_IP_sec_VPN_peer_endpoint_with_PSK_show_sensitive_data", inputDataValue, executionContext)
	var emptyOutput model.IPSecVPNPeerEndpoint
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingVPNIPSECPeerEndpointsGetIPSecVPNPeerEndpointWithPSKShowSensitiveDataOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IPSecVPNPeerEndpoint), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingVPNIPSECPeerEndpointsClient) ListIPSecVPNLocalEndpoints(cursorParam *string, includedFieldsParam *string, ipsecVpnServiceIdParam *string, logicalRouterIdParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.IPSecVPNLocalEndpointListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingVPNIPSECPeerEndpointsListIPSecVPNLocalEndpointsInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("IpsecVpnServiceId", ipsecVpnServiceIdParam)
	sv.AddStructField("LogicalRouterId", logicalRouterIdParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IPSecVPNLocalEndpointListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingVPNIPSECPeerEndpointsListIPSecVPNLocalEndpointsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_VPNIPSEC_peer_endpoints", "list_IP_sec_VPN_local_endpoints", inputDataValue, executionContext)
	var emptyOutput model.IPSecVPNLocalEndpointListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingVPNIPSECPeerEndpointsListIPSecVPNLocalEndpointsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IPSecVPNLocalEndpointListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingVPNIPSECPeerEndpointsClient) ListIPSecVPNPeerEndpoints(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.IPSecVPNPeerEndpointListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingVPNIPSECPeerEndpointsListIPSecVPNPeerEndpointsInputType(), typeConverter)
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
	operationRestMetaData := managementPlaneAPINetworkingVPNIPSECPeerEndpointsListIPSecVPNPeerEndpointsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_VPNIPSEC_peer_endpoints", "list_IP_sec_VPN_peer_endpoints", inputDataValue, executionContext)
	var emptyOutput model.IPSecVPNPeerEndpointListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingVPNIPSECPeerEndpointsListIPSecVPNPeerEndpointsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IPSecVPNPeerEndpointListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingVPNIPSECPeerEndpointsClient) UpdateIPSecVPNLocalEndpoint(ipsecVpnLocalEndpointIdParam string, ipSecVPNLocalEndpointParam model.IPSecVPNLocalEndpoint) (model.IPSecVPNLocalEndpoint, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingVPNIPSECPeerEndpointsUpdateIPSecVPNLocalEndpointInputType(), typeConverter)
	sv.AddStructField("IpsecVpnLocalEndpointId", ipsecVpnLocalEndpointIdParam)
	sv.AddStructField("IpSecVPNLocalEndpoint", ipSecVPNLocalEndpointParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IPSecVPNLocalEndpoint
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingVPNIPSECPeerEndpointsUpdateIPSecVPNLocalEndpointRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_VPNIPSEC_peer_endpoints", "update_IP_sec_VPN_local_endpoint", inputDataValue, executionContext)
	var emptyOutput model.IPSecVPNLocalEndpoint
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingVPNIPSECPeerEndpointsUpdateIPSecVPNLocalEndpointOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IPSecVPNLocalEndpoint), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingVPNIPSECPeerEndpointsClient) UpdateIPSecVPNPeerEndpoint(ipsecVpnPeerEndpointIdParam string, ipSecVPNPeerEndpointParam model.IPSecVPNPeerEndpoint) (model.IPSecVPNPeerEndpoint, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingVPNIPSECPeerEndpointsUpdateIPSecVPNPeerEndpointInputType(), typeConverter)
	sv.AddStructField("IpsecVpnPeerEndpointId", ipsecVpnPeerEndpointIdParam)
	sv.AddStructField("IpSecVPNPeerEndpoint", ipSecVPNPeerEndpointParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IPSecVPNPeerEndpoint
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingVPNIPSECPeerEndpointsUpdateIPSecVPNPeerEndpointRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_VPNIPSEC_peer_endpoints", "update_IP_sec_VPN_peer_endpoint", inputDataValue, executionContext)
	var emptyOutput model.IPSecVPNPeerEndpoint
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingVPNIPSECPeerEndpointsUpdateIPSecVPNPeerEndpointOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IPSecVPNPeerEndpoint), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
