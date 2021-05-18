// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ManagementPlaneAPIServicesDHCP
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

type ManagementPlaneAPIServicesDHCPClient interface {

	// Get specific leases of a given dhcp server. As a dhcp server could manage millions of leases, the API has to limit the number of the returned leases via two mutually-excluded request parameters, i.e. \"pool_id\" and \"address\". Either a \"pool_id\" or an \"address\" can be provided, but not both in a same call. If a \"pool_id\" is specified, the leases of the specific pool are returned. If an \"address\" is specified, only the lease(s) represented y this address is(are) returned. The \"address\" can be a single IP, an ip-range, or a mac address.
	//
	// @param serverIdParam (required)
	// @param addressParam can be an ip address, or an ip range, or a mac address (optional)
	// @param poolIdParam The uuid of dhcp ip pool (optional)
	// @param sourceParam Data source type. (optional)
	// @return com.vmware.model.DhcpLeases
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetDhcpLeaseInfo(serverIdParam string, addressParam *string, poolIdParam *string, sourceParam *string) (model.DhcpLeases, error)

	// Returns the statistics of the given dhcp server.
	//
	// @param serverIdParam (required)
	// @return com.vmware.model.DhcpStatistics
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetDhcpStatistics(serverIdParam string) (model.DhcpStatistics, error)
}

type managementPlaneAPIServicesDHCPClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewManagementPlaneAPIServicesDHCPClient(connector client.Connector) *managementPlaneAPIServicesDHCPClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.management_plane_API_services_DHCP")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get_dhcp_lease_info": core.NewMethodIdentifier(interfaceIdentifier, "get_dhcp_lease_info"),
		"get_dhcp_statistics": core.NewMethodIdentifier(interfaceIdentifier, "get_dhcp_statistics"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	mIface := managementPlaneAPIServicesDHCPClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &mIface
}

func (mIface *managementPlaneAPIServicesDHCPClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := mIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (mIface *managementPlaneAPIServicesDHCPClient) GetDhcpLeaseInfo(serverIdParam string, addressParam *string, poolIdParam *string, sourceParam *string) (model.DhcpLeases, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPIServicesDHCPGetDhcpLeaseInfoInputType(), typeConverter)
	sv.AddStructField("ServerId", serverIdParam)
	sv.AddStructField("Address", addressParam)
	sv.AddStructField("PoolId", poolIdParam)
	sv.AddStructField("Source", sourceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.DhcpLeases
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPIServicesDHCPGetDhcpLeaseInfoRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_services_DHCP", "get_dhcp_lease_info", inputDataValue, executionContext)
	var emptyOutput model.DhcpLeases
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPIServicesDHCPGetDhcpLeaseInfoOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.DhcpLeases), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPIServicesDHCPClient) GetDhcpStatistics(serverIdParam string) (model.DhcpStatistics, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPIServicesDHCPGetDhcpStatisticsInputType(), typeConverter)
	sv.AddStructField("ServerId", serverIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.DhcpStatistics
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPIServicesDHCPGetDhcpStatisticsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_services_DHCP", "get_dhcp_statistics", inputDataValue, executionContext)
	var emptyOutput model.DhcpStatistics
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPIServicesDHCPGetDhcpStatisticsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.DhcpStatistics), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
