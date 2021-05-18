// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ManagementPlaneAPITroubleshootingAndMonitoringPortConnection
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

type ManagementPlaneAPITroubleshootingAndMonitoringPortConnectionClient interface {

	// Get networking entities between two logical ports with VIF attachment
	//
	// @param lportIdParam ID of source port (required)
	// @param peerPortIdParam ID of peer port (required)
	// @return com.vmware.model.PortConnectionEntities
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetForwardingPath(lportIdParam string, peerPortIdParam string) (model.PortConnectionEntities, error)
}

type managementPlaneAPITroubleshootingAndMonitoringPortConnectionClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewManagementPlaneAPITroubleshootingAndMonitoringPortConnectionClient(connector client.Connector) *managementPlaneAPITroubleshootingAndMonitoringPortConnectionClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.management_plane_API_troubleshooting_and_monitoring_port_connection")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get_forwarding_path": core.NewMethodIdentifier(interfaceIdentifier, "get_forwarding_path"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	mIface := managementPlaneAPITroubleshootingAndMonitoringPortConnectionClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &mIface
}

func (mIface *managementPlaneAPITroubleshootingAndMonitoringPortConnectionClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := mIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (mIface *managementPlaneAPITroubleshootingAndMonitoringPortConnectionClient) GetForwardingPath(lportIdParam string, peerPortIdParam string) (model.PortConnectionEntities, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPITroubleshootingAndMonitoringPortConnectionGetForwardingPathInputType(), typeConverter)
	sv.AddStructField("LportId", lportIdParam)
	sv.AddStructField("PeerPortId", peerPortIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.PortConnectionEntities
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPITroubleshootingAndMonitoringPortConnectionGetForwardingPathRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_troubleshooting_and_monitoring_port_connection", "get_forwarding_path", inputDataValue, executionContext)
	var emptyOutput model.PortConnectionEntities
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPITroubleshootingAndMonitoringPortConnectionGetForwardingPathOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.PortConnectionEntities), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
