// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ManagementplaneapinetworkinglogicalswitchinginterSiteForwarder
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

type ManagementplaneapinetworkinglogicalswitchinginterSiteForwarderClient interface {

	// Returns remote mac addresses of the l2 forwarder on logical switch. It always returns realtime response.
	//
	// @param logicalSwitchIdParam (required)
	// @return com.vmware.model.L2ForwarderRemoteMacs
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetL2ForwarderRemoteMacs(logicalSwitchIdParam string) (model.L2ForwarderRemoteMacs, error)

	// Returns statistics of the l2 forwarder on logical switch. It always returns realtime response.
	//
	// @param logicalSwitchIdParam (required)
	// @return com.vmware.model.L2ForwarderStatistics
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetL2ForwarderStatistics(logicalSwitchIdParam string) (model.L2ForwarderStatistics, error)

	// Returns status per transport node of the l2 forwarder on logical switch.
	//
	// @param logicalSwitchIdParam (required)
	// @param sourceParam Data source type. (optional)
	// @return com.vmware.model.L2ForwarderStatus
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetL2ForwarderStatus(logicalSwitchIdParam string, sourceParam *string) (model.L2ForwarderStatus, error)
}

type managementplaneapinetworkinglogicalswitchinginterSiteForwarderClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewManagementplaneapinetworkinglogicalswitchinginterSiteForwarderClient(connector client.Connector) *managementplaneapinetworkinglogicalswitchinginterSiteForwarderClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.managementplaneapinetworkinglogicalswitchinginter_site_forwarder")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get_l2_forwarder_remote_macs": core.NewMethodIdentifier(interfaceIdentifier, "get_l2_forwarder_remote_macs"),
		"get_l2_forwarder_statistics":  core.NewMethodIdentifier(interfaceIdentifier, "get_l2_forwarder_statistics"),
		"get_l2_forwarder_status":      core.NewMethodIdentifier(interfaceIdentifier, "get_l2_forwarder_status"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	mIface := managementplaneapinetworkinglogicalswitchinginterSiteForwarderClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &mIface
}

func (mIface *managementplaneapinetworkinglogicalswitchinginterSiteForwarderClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := mIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (mIface *managementplaneapinetworkinglogicalswitchinginterSiteForwarderClient) GetL2ForwarderRemoteMacs(logicalSwitchIdParam string) (model.L2ForwarderRemoteMacs, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementplaneapinetworkinglogicalswitchinginterSiteForwarderGetL2ForwarderRemoteMacsInputType(), typeConverter)
	sv.AddStructField("LogicalSwitchId", logicalSwitchIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.L2ForwarderRemoteMacs
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementplaneapinetworkinglogicalswitchinginterSiteForwarderGetL2ForwarderRemoteMacsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.managementplaneapinetworkinglogicalswitchinginter_site_forwarder", "get_l2_forwarder_remote_macs", inputDataValue, executionContext)
	var emptyOutput model.L2ForwarderRemoteMacs
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementplaneapinetworkinglogicalswitchinginterSiteForwarderGetL2ForwarderRemoteMacsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.L2ForwarderRemoteMacs), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementplaneapinetworkinglogicalswitchinginterSiteForwarderClient) GetL2ForwarderStatistics(logicalSwitchIdParam string) (model.L2ForwarderStatistics, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementplaneapinetworkinglogicalswitchinginterSiteForwarderGetL2ForwarderStatisticsInputType(), typeConverter)
	sv.AddStructField("LogicalSwitchId", logicalSwitchIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.L2ForwarderStatistics
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementplaneapinetworkinglogicalswitchinginterSiteForwarderGetL2ForwarderStatisticsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.managementplaneapinetworkinglogicalswitchinginter_site_forwarder", "get_l2_forwarder_statistics", inputDataValue, executionContext)
	var emptyOutput model.L2ForwarderStatistics
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementplaneapinetworkinglogicalswitchinginterSiteForwarderGetL2ForwarderStatisticsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.L2ForwarderStatistics), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementplaneapinetworkinglogicalswitchinginterSiteForwarderClient) GetL2ForwarderStatus(logicalSwitchIdParam string, sourceParam *string) (model.L2ForwarderStatus, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementplaneapinetworkinglogicalswitchinginterSiteForwarderGetL2ForwarderStatusInputType(), typeConverter)
	sv.AddStructField("LogicalSwitchId", logicalSwitchIdParam)
	sv.AddStructField("Source", sourceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.L2ForwarderStatus
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementplaneapinetworkinglogicalswitchinginterSiteForwarderGetL2ForwarderStatusRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.managementplaneapinetworkinglogicalswitchinginter_site_forwarder", "get_l2_forwarder_status", inputDataValue, executionContext)
	var emptyOutput model.L2ForwarderStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementplaneapinetworkinglogicalswitchinginterSiteForwarderGetL2ForwarderStatusOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.L2ForwarderStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
