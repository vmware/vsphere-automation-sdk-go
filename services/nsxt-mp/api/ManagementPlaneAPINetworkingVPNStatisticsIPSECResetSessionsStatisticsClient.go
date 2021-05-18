// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ManagementPlaneAPINetworkingVPNStatisticsIPSECResetSessionsStatistics
// Used by client-side stubs.

package api

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
)

const _ = core.SupportedByRuntimeVersion1

type ManagementPlaneAPINetworkingVPNStatisticsIPSECResetSessionsStatisticsClient interface {

	// Reset the statistics of the given VPN session.
	//
	// @param sessionIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ResetIPSecVPNSessionStatisticsReset(sessionIdParam string) error
}

type managementPlaneAPINetworkingVPNStatisticsIPSECResetSessionsStatisticsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewManagementPlaneAPINetworkingVPNStatisticsIPSECResetSessionsStatisticsClient(connector client.Connector) *managementPlaneAPINetworkingVPNStatisticsIPSECResetSessionsStatisticsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.management_plane_API_networking_VPN_statistics_IPSEC_reset_sessions_statistics")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"reset_IP_sec_VPN_session_statistics_reset": core.NewMethodIdentifier(interfaceIdentifier, "reset_IP_sec_VPN_session_statistics_reset"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	mIface := managementPlaneAPINetworkingVPNStatisticsIPSECResetSessionsStatisticsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &mIface
}

func (mIface *managementPlaneAPINetworkingVPNStatisticsIPSECResetSessionsStatisticsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := mIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (mIface *managementPlaneAPINetworkingVPNStatisticsIPSECResetSessionsStatisticsClient) ResetIPSecVPNSessionStatisticsReset(sessionIdParam string) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingVPNStatisticsIPSECResetSessionsStatisticsResetIPSecVPNSessionStatisticsResetInputType(), typeConverter)
	sv.AddStructField("SessionId", sessionIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingVPNStatisticsIPSECResetSessionsStatisticsResetIPSecVPNSessionStatisticsResetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_VPN_statistics_IPSEC_reset_sessions_statistics", "reset_IP_sec_VPN_session_statistics_reset", inputDataValue, executionContext)
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
