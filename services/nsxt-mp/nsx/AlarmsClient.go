// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Alarms
// Used by client-side stubs.

package nsx

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type AlarmsClient interface {

	// Returns alarm associated with alarm-id. If HTTP status 404 is returned, this means the specified alarm-id is invalid or the alarm with alarm-id has been deleted. An alarm is deleted by the system if it is RESOLVED and older than eight days. The system can also delete the remaining RESOLVED alarms sooner to free system resources when too many alarms are being generated. When this happens the oldest day's RESOLVED alarms are deleted first.
	//
	// @param alarmIdParam (required)
	// @return com.vmware.nsx.model.Alarm
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(alarmIdParam string) (model.Alarm, error)

	// Returns a list of all Alarms currently known to the system.
	//
	// @param afterParam Timestamp in milliseconds since epoch (optional)
	// @param beforeParam Timestamp in milliseconds since epoch (optional)
	// @param cursorParam Cursor for pagination (optional)
	// @param eventTypeParam Event Type Filter (optional)
	// @param featureNameParam Feature Name (optional)
	// @param idParam Alarm ID (optional)
	// @param intentPathParam Intent Path for entity ID (optional)
	// @param nodeIdParam Node ID (optional)
	// @param nodeResourceTypeParam Node Resource Type (optional)
	// @param pageSizeParam Page Size for pagination (optional)
	// @param severityParam Severity (optional)
	// @param sortAscendingParam Represents order of sorting the values (optional, default to true)
	// @param sortByParam Key for sorting on this column (optional)
	// @param statusParam Status (optional)
	// @return com.vmware.nsx.model.AlarmsListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(afterParam *int64, beforeParam *int64, cursorParam *string, eventTypeParam *string, featureNameParam *string, idParam *string, intentPathParam *string, nodeIdParam *string, nodeResourceTypeParam *string, pageSizeParam *int64, severityParam *string, sortAscendingParam *bool, sortByParam *string, statusParam *string) (model.AlarmsListResult, error)

	// Update status of an Alarm. The new_status value can be OPEN, ACKNOWLEDGED, SUPPRESSED, or RESOLVED. If new_status is SUPPRESSED, the suppress_duration query parameter must also be specified.
	//
	// @param alarmIdParam (required)
	// @param newStatusParam Status (required)
	// @param suppressDurationParam Duration in hours for which Alarm should be suppressed (optional)
	// @return com.vmware.nsx.model.Alarm
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Setstatus(alarmIdParam string, newStatusParam string, suppressDurationParam *int64) (model.Alarm, error)

	// Bulk update the status of zero or more Alarms that match the specified filters. The new_status value can be OPEN, ACKNOWLEDGED, SUPPRESSED, or RESOLVED. If new_status is SUPPRESSED, the suppress_duration query parameter must also be specified.
	//
	// @param newStatusParam Status (required)
	// @param afterParam Timestamp in milliseconds since epoch (optional)
	// @param beforeParam Timestamp in milliseconds since epoch (optional)
	// @param cursorParam Cursor for pagination (optional)
	// @param eventTypeParam Event Type Filter (optional)
	// @param featureNameParam Feature Name (optional)
	// @param idParam Alarm ID (optional)
	// @param intentPathParam Intent Path for entity ID (optional)
	// @param nodeIdParam Node ID (optional)
	// @param nodeResourceTypeParam Node Resource Type (optional)
	// @param pageSizeParam Page Size for pagination (optional)
	// @param severityParam Severity (optional)
	// @param sortAscendingParam Represents order of sorting the values (optional, default to true)
	// @param sortByParam Key for sorting on this column (optional)
	// @param statusParam Status (optional)
	// @param suppressDurationParam Duration in hours for which Alarm should be suppressed (optional)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Setstatus0(newStatusParam string, afterParam *int64, beforeParam *int64, cursorParam *string, eventTypeParam *string, featureNameParam *string, idParam *string, intentPathParam *string, nodeIdParam *string, nodeResourceTypeParam *string, pageSizeParam *int64, severityParam *string, sortAscendingParam *bool, sortByParam *string, statusParam *string, suppressDurationParam *int64) error
}

type alarmsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewAlarmsClient(connector client.Connector) *alarmsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.alarms")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get":         core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":        core.NewMethodIdentifier(interfaceIdentifier, "list"),
		"setstatus":   core.NewMethodIdentifier(interfaceIdentifier, "setstatus"),
		"setstatus_0": core.NewMethodIdentifier(interfaceIdentifier, "setstatus_0"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	aIface := alarmsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &aIface
}

func (aIface *alarmsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := aIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (aIface *alarmsClient) Get(alarmIdParam string) (model.Alarm, error) {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(alarmsGetInputType(), typeConverter)
	sv.AddStructField("AlarmId", alarmIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.Alarm
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := alarmsGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	aIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.nsx.alarms", "get", inputDataValue, executionContext)
	var emptyOutput model.Alarm
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), alarmsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.Alarm), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (aIface *alarmsClient) List(afterParam *int64, beforeParam *int64, cursorParam *string, eventTypeParam *string, featureNameParam *string, idParam *string, intentPathParam *string, nodeIdParam *string, nodeResourceTypeParam *string, pageSizeParam *int64, severityParam *string, sortAscendingParam *bool, sortByParam *string, statusParam *string) (model.AlarmsListResult, error) {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(alarmsListInputType(), typeConverter)
	sv.AddStructField("After", afterParam)
	sv.AddStructField("Before", beforeParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("EventType", eventTypeParam)
	sv.AddStructField("FeatureName", featureNameParam)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("IntentPath", intentPathParam)
	sv.AddStructField("NodeId", nodeIdParam)
	sv.AddStructField("NodeResourceType", nodeResourceTypeParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("Severity", severityParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("Status", statusParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.AlarmsListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := alarmsListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	aIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.nsx.alarms", "list", inputDataValue, executionContext)
	var emptyOutput model.AlarmsListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), alarmsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.AlarmsListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (aIface *alarmsClient) Setstatus(alarmIdParam string, newStatusParam string, suppressDurationParam *int64) (model.Alarm, error) {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(alarmsSetstatusInputType(), typeConverter)
	sv.AddStructField("AlarmId", alarmIdParam)
	sv.AddStructField("NewStatus", newStatusParam)
	sv.AddStructField("SuppressDuration", suppressDurationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.Alarm
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := alarmsSetstatusRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	aIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.nsx.alarms", "setstatus", inputDataValue, executionContext)
	var emptyOutput model.Alarm
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), alarmsSetstatusOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.Alarm), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (aIface *alarmsClient) Setstatus0(newStatusParam string, afterParam *int64, beforeParam *int64, cursorParam *string, eventTypeParam *string, featureNameParam *string, idParam *string, intentPathParam *string, nodeIdParam *string, nodeResourceTypeParam *string, pageSizeParam *int64, severityParam *string, sortAscendingParam *bool, sortByParam *string, statusParam *string, suppressDurationParam *int64) error {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(alarmsSetstatus0InputType(), typeConverter)
	sv.AddStructField("NewStatus", newStatusParam)
	sv.AddStructField("After", afterParam)
	sv.AddStructField("Before", beforeParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("EventType", eventTypeParam)
	sv.AddStructField("FeatureName", featureNameParam)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("IntentPath", intentPathParam)
	sv.AddStructField("NodeId", nodeIdParam)
	sv.AddStructField("NodeResourceType", nodeResourceTypeParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("Severity", severityParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("Status", statusParam)
	sv.AddStructField("SuppressDuration", suppressDurationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := alarmsSetstatus0RestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	aIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.nsx.alarms", "setstatus_0", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
