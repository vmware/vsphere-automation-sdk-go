// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ManagementPlaneAPINetworkingVPNStatisticsIPSECSessionsStatistics
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

type ManagementPlaneAPINetworkingVPNStatisticsIPSECSessionsStatisticsClient interface {

	// Get statistics of a vpn session across all tunnels and IKE session. Query parameter \"source=realtime\" is the only supported source.
	//
	// @param sessionIdParam (required)
	// @param sourceParam Data source type. (optional)
	// @return com.vmware.model.IPSecVPNSessionStatistics
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetIPSecVPNSessionStatistics(sessionIdParam string, sourceParam *string) (model.IPSecVPNSessionStatistics, error)
}

type managementPlaneAPINetworkingVPNStatisticsIPSECSessionsStatisticsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewManagementPlaneAPINetworkingVPNStatisticsIPSECSessionsStatisticsClient(connector client.Connector) *managementPlaneAPINetworkingVPNStatisticsIPSECSessionsStatisticsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.management_plane_API_networking_VPN_statistics_IPSEC_sessions_statistics")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get_IP_sec_VPN_session_statistics": core.NewMethodIdentifier(interfaceIdentifier, "get_IP_sec_VPN_session_statistics"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	mIface := managementPlaneAPINetworkingVPNStatisticsIPSECSessionsStatisticsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &mIface
}

func (mIface *managementPlaneAPINetworkingVPNStatisticsIPSECSessionsStatisticsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := mIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (mIface *managementPlaneAPINetworkingVPNStatisticsIPSECSessionsStatisticsClient) GetIPSecVPNSessionStatistics(sessionIdParam string, sourceParam *string) (model.IPSecVPNSessionStatistics, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingVPNStatisticsIPSECSessionsStatisticsGetIPSecVPNSessionStatisticsInputType(), typeConverter)
	sv.AddStructField("SessionId", sessionIdParam)
	sv.AddStructField("Source", sourceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IPSecVPNSessionStatistics
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingVPNStatisticsIPSECSessionsStatisticsGetIPSecVPNSessionStatisticsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_VPN_statistics_IPSEC_sessions_statistics", "get_IP_sec_VPN_session_statistics", inputDataValue, executionContext)
	var emptyOutput model.IPSecVPNSessionStatistics
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingVPNStatisticsIPSECSessionsStatisticsGetIPSecVPNSessionStatisticsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IPSecVPNSessionStatistics), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
