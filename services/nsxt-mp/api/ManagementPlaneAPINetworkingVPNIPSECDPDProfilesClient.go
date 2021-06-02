// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ManagementPlaneAPINetworkingVPNIPSECDPDProfiles
// Used by client-side stubs.

package api

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/model"
)

const _ = core.SupportedByRuntimeVersion1

type ManagementPlaneAPINetworkingVPNIPSECDPDProfilesClient interface {

	// Create dead peer detection (DPD) profile. Any change in profile affects all sessions consuming this profile.
	//
	// @param ipSecVPNDPDProfileParam (required)
	// @return com.vmware.model.IPSecVPNDPDProfile
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	CreateIPSecVPNDPDProfile(ipSecVPNDPDProfileParam model.IPSecVPNDPDProfile) (model.IPSecVPNDPDProfile, error)

	// Delete dead peer detection (DPD) profile.
	//
	// @param ipsecVpnDpdProfileIdParam (required)
	// @param forceParam Force delete the resource even if it is being used somewhere (optional, default to false)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteIPSecVPNDPDProfile(ipsecVpnDpdProfileIdParam string, forceParam *bool) error

	// Get IPSec dead peer detection (DPD) profile.
	//
	// @param ipsecVpnDpdProfileIdParam (required)
	// @return com.vmware.model.IPSecVPNDPDProfile
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetIPSecVPNDPDProfile(ipsecVpnDpdProfileIdParam string) (model.IPSecVPNDPDProfile, error)

	// Get paginated list of all dead peer detection (DPD) profiles.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.model.IPSecVPNDPDProfileListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListIPSecVPNDPDProfiles(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.IPSecVPNDPDProfileListResult, error)

	// Edit IPSec dead peer detection (DPD) profile.
	//
	// @param ipsecVpnDpdProfileIdParam (required)
	// @param ipSecVPNDPDProfileParam (required)
	// @return com.vmware.model.IPSecVPNDPDProfile
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateIPSecVPNDPDProfile(ipsecVpnDpdProfileIdParam string, ipSecVPNDPDProfileParam model.IPSecVPNDPDProfile) (model.IPSecVPNDPDProfile, error)
}

type managementPlaneAPINetworkingVPNIPSECDPDProfilesClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewManagementPlaneAPINetworkingVPNIPSECDPDProfilesClient(connector client.Connector) *managementPlaneAPINetworkingVPNIPSECDPDProfilesClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.management_plane_API_networking_VPNIPSECDPD_profiles")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create_IP_sec_VPNDPD_profile": core.NewMethodIdentifier(interfaceIdentifier, "create_IP_sec_VPNDPD_profile"),
		"delete_IP_sec_VPNDPD_profile": core.NewMethodIdentifier(interfaceIdentifier, "delete_IP_sec_VPNDPD_profile"),
		"get_IP_sec_VPNDPD_profile":    core.NewMethodIdentifier(interfaceIdentifier, "get_IP_sec_VPNDPD_profile"),
		"list_IP_sec_VPNDPD_profiles":  core.NewMethodIdentifier(interfaceIdentifier, "list_IP_sec_VPNDPD_profiles"),
		"update_IP_sec_VPNDPD_profile": core.NewMethodIdentifier(interfaceIdentifier, "update_IP_sec_VPNDPD_profile"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	mIface := managementPlaneAPINetworkingVPNIPSECDPDProfilesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &mIface
}

func (mIface *managementPlaneAPINetworkingVPNIPSECDPDProfilesClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := mIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (mIface *managementPlaneAPINetworkingVPNIPSECDPDProfilesClient) CreateIPSecVPNDPDProfile(ipSecVPNDPDProfileParam model.IPSecVPNDPDProfile) (model.IPSecVPNDPDProfile, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingVPNIPSECDPDProfilesCreateIPSecVPNDPDProfileInputType(), typeConverter)
	sv.AddStructField("IpSecVPNDPDProfile", ipSecVPNDPDProfileParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IPSecVPNDPDProfile
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingVPNIPSECDPDProfilesCreateIPSecVPNDPDProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_VPNIPSECDPD_profiles", "create_IP_sec_VPNDPD_profile", inputDataValue, executionContext)
	var emptyOutput model.IPSecVPNDPDProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingVPNIPSECDPDProfilesCreateIPSecVPNDPDProfileOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IPSecVPNDPDProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingVPNIPSECDPDProfilesClient) DeleteIPSecVPNDPDProfile(ipsecVpnDpdProfileIdParam string, forceParam *bool) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingVPNIPSECDPDProfilesDeleteIPSecVPNDPDProfileInputType(), typeConverter)
	sv.AddStructField("IpsecVpnDpdProfileId", ipsecVpnDpdProfileIdParam)
	sv.AddStructField("Force", forceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingVPNIPSECDPDProfilesDeleteIPSecVPNDPDProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_VPNIPSECDPD_profiles", "delete_IP_sec_VPNDPD_profile", inputDataValue, executionContext)
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

func (mIface *managementPlaneAPINetworkingVPNIPSECDPDProfilesClient) GetIPSecVPNDPDProfile(ipsecVpnDpdProfileIdParam string) (model.IPSecVPNDPDProfile, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingVPNIPSECDPDProfilesGetIPSecVPNDPDProfileInputType(), typeConverter)
	sv.AddStructField("IpsecVpnDpdProfileId", ipsecVpnDpdProfileIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IPSecVPNDPDProfile
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingVPNIPSECDPDProfilesGetIPSecVPNDPDProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_VPNIPSECDPD_profiles", "get_IP_sec_VPNDPD_profile", inputDataValue, executionContext)
	var emptyOutput model.IPSecVPNDPDProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingVPNIPSECDPDProfilesGetIPSecVPNDPDProfileOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IPSecVPNDPDProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingVPNIPSECDPDProfilesClient) ListIPSecVPNDPDProfiles(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.IPSecVPNDPDProfileListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingVPNIPSECDPDProfilesListIPSecVPNDPDProfilesInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IPSecVPNDPDProfileListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingVPNIPSECDPDProfilesListIPSecVPNDPDProfilesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_VPNIPSECDPD_profiles", "list_IP_sec_VPNDPD_profiles", inputDataValue, executionContext)
	var emptyOutput model.IPSecVPNDPDProfileListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingVPNIPSECDPDProfilesListIPSecVPNDPDProfilesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IPSecVPNDPDProfileListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingVPNIPSECDPDProfilesClient) UpdateIPSecVPNDPDProfile(ipsecVpnDpdProfileIdParam string, ipSecVPNDPDProfileParam model.IPSecVPNDPDProfile) (model.IPSecVPNDPDProfile, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingVPNIPSECDPDProfilesUpdateIPSecVPNDPDProfileInputType(), typeConverter)
	sv.AddStructField("IpsecVpnDpdProfileId", ipsecVpnDpdProfileIdParam)
	sv.AddStructField("IpSecVPNDPDProfile", ipSecVPNDPDProfileParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IPSecVPNDPDProfile
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingVPNIPSECDPDProfilesUpdateIPSecVPNDPDProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_VPNIPSECDPD_profiles", "update_IP_sec_VPNDPD_profile", inputDataValue, executionContext)
	var emptyOutput model.IPSecVPNDPDProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingVPNIPSECDPDProfilesUpdateIPSecVPNDPDProfileOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IPSecVPNDPDProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
