// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ManagementPlaneAPISecurityIdentityFirewallMonitoring
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

type ManagementPlaneAPISecurityIdentityFirewallMonitoringClient interface {

	// Retrieve the compute collection status by ID.
	//
	// @param computeCollectionExtIdParam (required)
	// @return com.vmware.model.IdfwComputeCollectionStatus
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetComputeCollectionStatusById(computeCollectionExtIdParam string) (model.IdfwComputeCollectionStatus, error)

	// Retrieve all the Compute collection status.
	// @return com.vmware.model.IdfwComputeCollectionListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListComputeCollectionStatuses() (model.IdfwComputeCollectionListResult, error)

	// Retrieve all the transport node and status by idfw enabled ComputeCollection ID in the request.
	//
	// @param ccExtIdParam (required)
	// @return com.vmware.model.IdfwTransportNodeStatusListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListTransportNodeStatusesByComputeCollectionId(ccExtIdParam string) (model.IdfwTransportNodeStatusListResult, error)

	// Retrieve all the VM and status by transport node ID of idfw enabled compute collection in the request.
	//
	// @param transportNodeIdParam (required)
	// @return com.vmware.model.IdfwVirtualMachineStatusListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListVirtualMachineStatusesByTransportNodeId(transportNodeIdParam string) (model.IdfwVirtualMachineStatusListResult, error)
}

type managementPlaneAPISecurityIdentityFirewallMonitoringClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewManagementPlaneAPISecurityIdentityFirewallMonitoringClient(connector client.Connector) *managementPlaneAPISecurityIdentityFirewallMonitoringClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.management_plane_API_security_identity_firewall_monitoring")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get_compute_collection_status_by_id":                   core.NewMethodIdentifier(interfaceIdentifier, "get_compute_collection_status_by_id"),
		"list_compute_collection_statuses":                      core.NewMethodIdentifier(interfaceIdentifier, "list_compute_collection_statuses"),
		"list_transport_node_statuses_by_compute_collection_id": core.NewMethodIdentifier(interfaceIdentifier, "list_transport_node_statuses_by_compute_collection_id"),
		"list_virtual_machine_statuses_by_transport_node_id":    core.NewMethodIdentifier(interfaceIdentifier, "list_virtual_machine_statuses_by_transport_node_id"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	mIface := managementPlaneAPISecurityIdentityFirewallMonitoringClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &mIface
}

func (mIface *managementPlaneAPISecurityIdentityFirewallMonitoringClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := mIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (mIface *managementPlaneAPISecurityIdentityFirewallMonitoringClient) GetComputeCollectionStatusById(computeCollectionExtIdParam string) (model.IdfwComputeCollectionStatus, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityIdentityFirewallMonitoringGetComputeCollectionStatusByIdInputType(), typeConverter)
	sv.AddStructField("ComputeCollectionExtId", computeCollectionExtIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IdfwComputeCollectionStatus
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityIdentityFirewallMonitoringGetComputeCollectionStatusByIdRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_identity_firewall_monitoring", "get_compute_collection_status_by_id", inputDataValue, executionContext)
	var emptyOutput model.IdfwComputeCollectionStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityIdentityFirewallMonitoringGetComputeCollectionStatusByIdOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IdfwComputeCollectionStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityIdentityFirewallMonitoringClient) ListComputeCollectionStatuses() (model.IdfwComputeCollectionListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityIdentityFirewallMonitoringListComputeCollectionStatusesInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IdfwComputeCollectionListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityIdentityFirewallMonitoringListComputeCollectionStatusesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_identity_firewall_monitoring", "list_compute_collection_statuses", inputDataValue, executionContext)
	var emptyOutput model.IdfwComputeCollectionListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityIdentityFirewallMonitoringListComputeCollectionStatusesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IdfwComputeCollectionListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityIdentityFirewallMonitoringClient) ListTransportNodeStatusesByComputeCollectionId(ccExtIdParam string) (model.IdfwTransportNodeStatusListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityIdentityFirewallMonitoringListTransportNodeStatusesByComputeCollectionIdInputType(), typeConverter)
	sv.AddStructField("CcExtId", ccExtIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IdfwTransportNodeStatusListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityIdentityFirewallMonitoringListTransportNodeStatusesByComputeCollectionIdRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_identity_firewall_monitoring", "list_transport_node_statuses_by_compute_collection_id", inputDataValue, executionContext)
	var emptyOutput model.IdfwTransportNodeStatusListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityIdentityFirewallMonitoringListTransportNodeStatusesByComputeCollectionIdOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IdfwTransportNodeStatusListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityIdentityFirewallMonitoringClient) ListVirtualMachineStatusesByTransportNodeId(transportNodeIdParam string) (model.IdfwVirtualMachineStatusListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityIdentityFirewallMonitoringListVirtualMachineStatusesByTransportNodeIdInputType(), typeConverter)
	sv.AddStructField("TransportNodeId", transportNodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IdfwVirtualMachineStatusListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityIdentityFirewallMonitoringListVirtualMachineStatusesByTransportNodeIdRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_identity_firewall_monitoring", "list_virtual_machine_statuses_by_transport_node_id", inputDataValue, executionContext)
	var emptyOutput model.IdfwVirtualMachineStatusListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityIdentityFirewallMonitoringListVirtualMachineStatusesByTransportNodeIdOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IdfwVirtualMachineStatusListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
