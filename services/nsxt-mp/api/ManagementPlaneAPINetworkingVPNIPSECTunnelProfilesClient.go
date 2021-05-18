// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ManagementPlaneAPINetworkingVPNIPSECTunnelProfiles
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

type ManagementPlaneAPINetworkingVPNIPSECTunnelProfilesClient interface {

	// Create custom IPSec tunnel profile. IPSec tunnel profile is a reusable profile that captures phase two negotiation parameters and tunnel properties. System will be provisioned with system owned non editable default IPSec tunnel profile. Any change in profile affects all sessions consuming this profile.
	//
	// @param ipSecVPNTunnelProfileParam (required)
	// @return com.vmware.model.IPSecVPNTunnelProfile
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	CreateIPSecVPNTunnelProfile(ipSecVPNTunnelProfileParam model.IPSecVPNTunnelProfile) (model.IPSecVPNTunnelProfile, error)

	// Delete custom IPSec Tunnel Profile.
	//
	// @param ipsecVpnTunnelProfileIdParam (required)
	// @param forceParam Force delete the resource even if it is being used somewhere (optional, default to false)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteIPSecVPNTunnelProfile(ipsecVpnTunnelProfileIdParam string, forceParam *bool) error

	// Get custom IPSec Tunnel Profile.
	//
	// @param ipsecVpnTunnelProfileIdParam (required)
	// @return com.vmware.model.IPSecVPNTunnelProfile
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetIPSecVPNTunnelProfile(ipsecVpnTunnelProfileIdParam string) (model.IPSecVPNTunnelProfile, error)

	// Get paginated list of all IPSecTunnelProfiles.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.model.IPSecVPNTunnelProfileListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListIPSecVPNTunnelProfiles(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.IPSecVPNTunnelProfileListResult, error)

	// Edit custom IPSec Tunnel Profile. System owned profiles are non editable.
	//
	// @param ipsecVpnTunnelProfileIdParam (required)
	// @param ipSecVPNTunnelProfileParam (required)
	// @return com.vmware.model.IPSecVPNTunnelProfile
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateIPSecVPNTunnelProfile(ipsecVpnTunnelProfileIdParam string, ipSecVPNTunnelProfileParam model.IPSecVPNTunnelProfile) (model.IPSecVPNTunnelProfile, error)
}

type managementPlaneAPINetworkingVPNIPSECTunnelProfilesClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewManagementPlaneAPINetworkingVPNIPSECTunnelProfilesClient(connector client.Connector) *managementPlaneAPINetworkingVPNIPSECTunnelProfilesClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.management_plane_API_networking_VPNIPSEC_tunnel_profiles")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create_IP_sec_VPN_tunnel_profile": core.NewMethodIdentifier(interfaceIdentifier, "create_IP_sec_VPN_tunnel_profile"),
		"delete_IP_sec_VPN_tunnel_profile": core.NewMethodIdentifier(interfaceIdentifier, "delete_IP_sec_VPN_tunnel_profile"),
		"get_IP_sec_VPN_tunnel_profile":    core.NewMethodIdentifier(interfaceIdentifier, "get_IP_sec_VPN_tunnel_profile"),
		"list_IP_sec_VPN_tunnel_profiles":  core.NewMethodIdentifier(interfaceIdentifier, "list_IP_sec_VPN_tunnel_profiles"),
		"update_IP_sec_VPN_tunnel_profile": core.NewMethodIdentifier(interfaceIdentifier, "update_IP_sec_VPN_tunnel_profile"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	mIface := managementPlaneAPINetworkingVPNIPSECTunnelProfilesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &mIface
}

func (mIface *managementPlaneAPINetworkingVPNIPSECTunnelProfilesClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := mIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (mIface *managementPlaneAPINetworkingVPNIPSECTunnelProfilesClient) CreateIPSecVPNTunnelProfile(ipSecVPNTunnelProfileParam model.IPSecVPNTunnelProfile) (model.IPSecVPNTunnelProfile, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingVPNIPSECTunnelProfilesCreateIPSecVPNTunnelProfileInputType(), typeConverter)
	sv.AddStructField("IpSecVPNTunnelProfile", ipSecVPNTunnelProfileParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IPSecVPNTunnelProfile
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingVPNIPSECTunnelProfilesCreateIPSecVPNTunnelProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_VPNIPSEC_tunnel_profiles", "create_IP_sec_VPN_tunnel_profile", inputDataValue, executionContext)
	var emptyOutput model.IPSecVPNTunnelProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingVPNIPSECTunnelProfilesCreateIPSecVPNTunnelProfileOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IPSecVPNTunnelProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingVPNIPSECTunnelProfilesClient) DeleteIPSecVPNTunnelProfile(ipsecVpnTunnelProfileIdParam string, forceParam *bool) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingVPNIPSECTunnelProfilesDeleteIPSecVPNTunnelProfileInputType(), typeConverter)
	sv.AddStructField("IpsecVpnTunnelProfileId", ipsecVpnTunnelProfileIdParam)
	sv.AddStructField("Force", forceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingVPNIPSECTunnelProfilesDeleteIPSecVPNTunnelProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_VPNIPSEC_tunnel_profiles", "delete_IP_sec_VPN_tunnel_profile", inputDataValue, executionContext)
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

func (mIface *managementPlaneAPINetworkingVPNIPSECTunnelProfilesClient) GetIPSecVPNTunnelProfile(ipsecVpnTunnelProfileIdParam string) (model.IPSecVPNTunnelProfile, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingVPNIPSECTunnelProfilesGetIPSecVPNTunnelProfileInputType(), typeConverter)
	sv.AddStructField("IpsecVpnTunnelProfileId", ipsecVpnTunnelProfileIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IPSecVPNTunnelProfile
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingVPNIPSECTunnelProfilesGetIPSecVPNTunnelProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_VPNIPSEC_tunnel_profiles", "get_IP_sec_VPN_tunnel_profile", inputDataValue, executionContext)
	var emptyOutput model.IPSecVPNTunnelProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingVPNIPSECTunnelProfilesGetIPSecVPNTunnelProfileOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IPSecVPNTunnelProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingVPNIPSECTunnelProfilesClient) ListIPSecVPNTunnelProfiles(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.IPSecVPNTunnelProfileListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingVPNIPSECTunnelProfilesListIPSecVPNTunnelProfilesInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IPSecVPNTunnelProfileListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingVPNIPSECTunnelProfilesListIPSecVPNTunnelProfilesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_VPNIPSEC_tunnel_profiles", "list_IP_sec_VPN_tunnel_profiles", inputDataValue, executionContext)
	var emptyOutput model.IPSecVPNTunnelProfileListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingVPNIPSECTunnelProfilesListIPSecVPNTunnelProfilesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IPSecVPNTunnelProfileListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingVPNIPSECTunnelProfilesClient) UpdateIPSecVPNTunnelProfile(ipsecVpnTunnelProfileIdParam string, ipSecVPNTunnelProfileParam model.IPSecVPNTunnelProfile) (model.IPSecVPNTunnelProfile, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingVPNIPSECTunnelProfilesUpdateIPSecVPNTunnelProfileInputType(), typeConverter)
	sv.AddStructField("IpsecVpnTunnelProfileId", ipsecVpnTunnelProfileIdParam)
	sv.AddStructField("IpSecVPNTunnelProfile", ipSecVPNTunnelProfileParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IPSecVPNTunnelProfile
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingVPNIPSECTunnelProfilesUpdateIPSecVPNTunnelProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_VPNIPSEC_tunnel_profiles", "update_IP_sec_VPN_tunnel_profile", inputDataValue, executionContext)
	var emptyOutput model.IPSecVPNTunnelProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingVPNIPSECTunnelProfilesUpdateIPSecVPNTunnelProfileOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IPSecVPNTunnelProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
