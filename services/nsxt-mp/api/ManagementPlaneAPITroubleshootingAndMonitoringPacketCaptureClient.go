// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ManagementPlaneAPITroubleshootingAndMonitoringPacketCapture
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

type ManagementPlaneAPITroubleshootingAndMonitoringPacketCaptureClient interface {

	// Create an new packet capture session on given node with specified options
	//
	// @param packetCaptureRequestParam (required)
	// @return com.vmware.model.PacketCaptureSession
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	CreatePacketCaptureSession(packetCaptureRequestParam model.PacketCaptureRequest) (model.PacketCaptureSession, error)

	// Delete all the packet capture sessions.
	// @return com.vmware.model.PacketCaptureSessionList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteAllCaptureSessionsDelete() (model.PacketCaptureSessionList, error)

	// Before calling this method, terminate any running capture session.
	//
	// @param sessionIdParam Packet capture session id (required)
	// @return com.vmware.model.PacketCaptureSession
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeletePacketCaptureSessionDelete(sessionIdParam string) (model.PacketCaptureSession, error)

	// You must provide the request header \"Accept:application/octet-stream\" when calling this API. The capture file can only be found in MP which receives the capture request.
	//
	// @param sessionIdParam Packet capture session id (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetCaptureFile(sessionIdParam string) error

	// Get the information of all packet capture sessions.
	// @return com.vmware.model.PacketCaptureSessionList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListPacketCaptureSessions() (model.PacketCaptureSessionList, error)

	// Get the packet capture status information by session id.
	//
	// @param sessionIdParam Packet capture session id (required)
	// @return com.vmware.model.PacketCaptureSession
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ReadPacketCaptureSession(sessionIdParam string) (model.PacketCaptureSession, error)

	// Restart the packet capture session
	//
	// @param sessionIdParam Packet capture session id (required)
	// @return com.vmware.model.PacketCaptureSession
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	RestartPacketCaptureSessionRestart(sessionIdParam string) (model.PacketCaptureSession, error)

	// Terminate the packet capture session by session id.
	//
	// @param sessionIdParam Packet capture session id (required)
	// @return com.vmware.model.PacketCaptureSession
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	TerminatePacketCaptureSessionTerminate(sessionIdParam string) (model.PacketCaptureSession, error)
}

type managementPlaneAPITroubleshootingAndMonitoringPacketCaptureClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewManagementPlaneAPITroubleshootingAndMonitoringPacketCaptureClient(connector client.Connector) *managementPlaneAPITroubleshootingAndMonitoringPacketCaptureClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.management_plane_API_troubleshooting_and_monitoring_packet_capture")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create_packet_capture_session":              core.NewMethodIdentifier(interfaceIdentifier, "create_packet_capture_session"),
		"delete_all_capture_sessions_delete":         core.NewMethodIdentifier(interfaceIdentifier, "delete_all_capture_sessions_delete"),
		"delete_packet_capture_session_delete":       core.NewMethodIdentifier(interfaceIdentifier, "delete_packet_capture_session_delete"),
		"get_capture_file":                           core.NewMethodIdentifier(interfaceIdentifier, "get_capture_file"),
		"list_packet_capture_sessions":               core.NewMethodIdentifier(interfaceIdentifier, "list_packet_capture_sessions"),
		"read_packet_capture_session":                core.NewMethodIdentifier(interfaceIdentifier, "read_packet_capture_session"),
		"restart_packet_capture_session_restart":     core.NewMethodIdentifier(interfaceIdentifier, "restart_packet_capture_session_restart"),
		"terminate_packet_capture_session_terminate": core.NewMethodIdentifier(interfaceIdentifier, "terminate_packet_capture_session_terminate"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	mIface := managementPlaneAPITroubleshootingAndMonitoringPacketCaptureClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &mIface
}

func (mIface *managementPlaneAPITroubleshootingAndMonitoringPacketCaptureClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := mIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (mIface *managementPlaneAPITroubleshootingAndMonitoringPacketCaptureClient) CreatePacketCaptureSession(packetCaptureRequestParam model.PacketCaptureRequest) (model.PacketCaptureSession, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPITroubleshootingAndMonitoringPacketCaptureCreatePacketCaptureSessionInputType(), typeConverter)
	sv.AddStructField("PacketCaptureRequest", packetCaptureRequestParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.PacketCaptureSession
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPITroubleshootingAndMonitoringPacketCaptureCreatePacketCaptureSessionRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_troubleshooting_and_monitoring_packet_capture", "create_packet_capture_session", inputDataValue, executionContext)
	var emptyOutput model.PacketCaptureSession
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPITroubleshootingAndMonitoringPacketCaptureCreatePacketCaptureSessionOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.PacketCaptureSession), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPITroubleshootingAndMonitoringPacketCaptureClient) DeleteAllCaptureSessionsDelete() (model.PacketCaptureSessionList, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPITroubleshootingAndMonitoringPacketCaptureDeleteAllCaptureSessionsDeleteInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.PacketCaptureSessionList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPITroubleshootingAndMonitoringPacketCaptureDeleteAllCaptureSessionsDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_troubleshooting_and_monitoring_packet_capture", "delete_all_capture_sessions_delete", inputDataValue, executionContext)
	var emptyOutput model.PacketCaptureSessionList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPITroubleshootingAndMonitoringPacketCaptureDeleteAllCaptureSessionsDeleteOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.PacketCaptureSessionList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPITroubleshootingAndMonitoringPacketCaptureClient) DeletePacketCaptureSessionDelete(sessionIdParam string) (model.PacketCaptureSession, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPITroubleshootingAndMonitoringPacketCaptureDeletePacketCaptureSessionDeleteInputType(), typeConverter)
	sv.AddStructField("SessionId", sessionIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.PacketCaptureSession
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPITroubleshootingAndMonitoringPacketCaptureDeletePacketCaptureSessionDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_troubleshooting_and_monitoring_packet_capture", "delete_packet_capture_session_delete", inputDataValue, executionContext)
	var emptyOutput model.PacketCaptureSession
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPITroubleshootingAndMonitoringPacketCaptureDeletePacketCaptureSessionDeleteOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.PacketCaptureSession), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPITroubleshootingAndMonitoringPacketCaptureClient) GetCaptureFile(sessionIdParam string) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPITroubleshootingAndMonitoringPacketCaptureGetCaptureFileInputType(), typeConverter)
	sv.AddStructField("SessionId", sessionIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPITroubleshootingAndMonitoringPacketCaptureGetCaptureFileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_troubleshooting_and_monitoring_packet_capture", "get_capture_file", inputDataValue, executionContext)
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

func (mIface *managementPlaneAPITroubleshootingAndMonitoringPacketCaptureClient) ListPacketCaptureSessions() (model.PacketCaptureSessionList, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPITroubleshootingAndMonitoringPacketCaptureListPacketCaptureSessionsInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.PacketCaptureSessionList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPITroubleshootingAndMonitoringPacketCaptureListPacketCaptureSessionsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_troubleshooting_and_monitoring_packet_capture", "list_packet_capture_sessions", inputDataValue, executionContext)
	var emptyOutput model.PacketCaptureSessionList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPITroubleshootingAndMonitoringPacketCaptureListPacketCaptureSessionsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.PacketCaptureSessionList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPITroubleshootingAndMonitoringPacketCaptureClient) ReadPacketCaptureSession(sessionIdParam string) (model.PacketCaptureSession, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPITroubleshootingAndMonitoringPacketCaptureReadPacketCaptureSessionInputType(), typeConverter)
	sv.AddStructField("SessionId", sessionIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.PacketCaptureSession
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPITroubleshootingAndMonitoringPacketCaptureReadPacketCaptureSessionRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_troubleshooting_and_monitoring_packet_capture", "read_packet_capture_session", inputDataValue, executionContext)
	var emptyOutput model.PacketCaptureSession
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPITroubleshootingAndMonitoringPacketCaptureReadPacketCaptureSessionOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.PacketCaptureSession), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPITroubleshootingAndMonitoringPacketCaptureClient) RestartPacketCaptureSessionRestart(sessionIdParam string) (model.PacketCaptureSession, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPITroubleshootingAndMonitoringPacketCaptureRestartPacketCaptureSessionRestartInputType(), typeConverter)
	sv.AddStructField("SessionId", sessionIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.PacketCaptureSession
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPITroubleshootingAndMonitoringPacketCaptureRestartPacketCaptureSessionRestartRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_troubleshooting_and_monitoring_packet_capture", "restart_packet_capture_session_restart", inputDataValue, executionContext)
	var emptyOutput model.PacketCaptureSession
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPITroubleshootingAndMonitoringPacketCaptureRestartPacketCaptureSessionRestartOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.PacketCaptureSession), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPITroubleshootingAndMonitoringPacketCaptureClient) TerminatePacketCaptureSessionTerminate(sessionIdParam string) (model.PacketCaptureSession, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPITroubleshootingAndMonitoringPacketCaptureTerminatePacketCaptureSessionTerminateInputType(), typeConverter)
	sv.AddStructField("SessionId", sessionIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.PacketCaptureSession
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPITroubleshootingAndMonitoringPacketCaptureTerminatePacketCaptureSessionTerminateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_troubleshooting_and_monitoring_packet_capture", "terminate_packet_capture_session_terminate", inputDataValue, executionContext)
	var emptyOutput model.PacketCaptureSession
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPITroubleshootingAndMonitoringPacketCaptureTerminatePacketCaptureSessionTerminateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.PacketCaptureSession), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
