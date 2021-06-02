// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ManagementPlaneAPINetworkingVPNStatisticsL2VPNRemoteMACS
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

type ManagementPlaneAPINetworkingVPNStatisticsL2VPNRemoteMACSClient interface {

	// Get L2VPN session remote mac for logical switch.
	//
	// @param sessionIdParam (required)
	// @param logicalSwitchIdParam logical switch identifier (optional)
	// @return com.vmware.model.L2VPNSessionRemoteMacs
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetL2VPNSessionRemoteMacsForLS(sessionIdParam string, logicalSwitchIdParam *string) (model.L2VPNSessionRemoteMacs, error)
}

type managementPlaneAPINetworkingVPNStatisticsL2VPNRemoteMACSClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewManagementPlaneAPINetworkingVPNStatisticsL2VPNRemoteMACSClient(connector client.Connector) *managementPlaneAPINetworkingVPNStatisticsL2VPNRemoteMACSClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.management_plane_API_networking_VPN_statistics_l2_VPN_remote_MACS")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get_l2_VPN_session_remote_macs_for_LS": core.NewMethodIdentifier(interfaceIdentifier, "get_l2_VPN_session_remote_macs_for_LS"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	mIface := managementPlaneAPINetworkingVPNStatisticsL2VPNRemoteMACSClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &mIface
}

func (mIface *managementPlaneAPINetworkingVPNStatisticsL2VPNRemoteMACSClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := mIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (mIface *managementPlaneAPINetworkingVPNStatisticsL2VPNRemoteMACSClient) GetL2VPNSessionRemoteMacsForLS(sessionIdParam string, logicalSwitchIdParam *string) (model.L2VPNSessionRemoteMacs, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingVPNStatisticsL2VPNRemoteMACSGetL2VPNSessionRemoteMacsForLSInputType(), typeConverter)
	sv.AddStructField("SessionId", sessionIdParam)
	sv.AddStructField("LogicalSwitchId", logicalSwitchIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.L2VPNSessionRemoteMacs
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingVPNStatisticsL2VPNRemoteMACSGetL2VPNSessionRemoteMacsForLSRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_VPN_statistics_l2_VPN_remote_MACS", "get_l2_VPN_session_remote_macs_for_LS", inputDataValue, executionContext)
	var emptyOutput model.L2VPNSessionRemoteMacs
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingVPNStatisticsL2VPNRemoteMACSGetL2VPNSessionRemoteMacsForLSOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.L2VPNSessionRemoteMacs), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
