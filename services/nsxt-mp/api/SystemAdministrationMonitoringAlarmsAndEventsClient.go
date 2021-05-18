// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: SystemAdministrationMonitoringAlarmsAndEvents
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

type SystemAdministrationMonitoringAlarmsAndEventsClient interface {

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
	BulkUpdateAlarmsSetStatus(newStatusParam string, afterParam *int64, beforeParam *int64, cursorParam *string, eventTypeParam *string, featureNameParam *string, idParam *string, intentPathParam *string, nodeIdParam *string, nodeResourceTypeParam *string, pageSizeParam *int64, severityParam *string, sortAscendingParam *bool, sortByParam *string, statusParam *string, suppressDurationParam *int64) error

	// Returns alarm associated with alarm-id. If HTTP status 404 is returned, this means the specified alarm-id is invalid or the alarm with alarm-id has been deleted. An alarm is deleted by the system if it is RESOLVED and older than eight days. The system can also delete the remaining RESOLVED alarms sooner to free system resources when too many alarms are being generated. When this happens the oldest day's RESOLVED alarms are deleted first.
	//
	// @param alarmIdParam (required)
	// @return com.vmware.model.Alarm
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetAlarm(alarmIdParam string) (model.Alarm, error)

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
	// @return com.vmware.model.AlarmsListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetAlarms(afterParam *int64, beforeParam *int64, cursorParam *string, eventTypeParam *string, featureNameParam *string, idParam *string, intentPathParam *string, nodeIdParam *string, nodeResourceTypeParam *string, pageSizeParam *int64, severityParam *string, sortAscendingParam *bool, sortByParam *string, statusParam *string) (model.AlarmsListResult, error)

	// Returns event associated with event-id.
	//
	// @param eventIdParam (required)
	// @return com.vmware.model.MonitoringEvent
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetEvent(eventIdParam string) (model.MonitoringEvent, error)

	// Returns a list of all Events defined in NSX.
	// @return com.vmware.model.EventListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListEvents() (model.EventListResult, error)

	// Reset all user configurable values for event identified by event-id to factory defaults.
	//
	// @param eventIdParam (required)
	// @return com.vmware.model.MonitoringEvent
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ResetEventValuesSetDefault(eventIdParam string) (model.MonitoringEvent, error)

	// Update status of an Alarm. The new_status value can be OPEN, ACKNOWLEDGED, SUPPRESSED, or RESOLVED. If new_status is SUPPRESSED, the suppress_duration query parameter must also be specified.
	//
	// @param alarmIdParam (required)
	// @param newStatusParam Status (required)
	// @param suppressDurationParam Duration in hours for which Alarm should be suppressed (optional)
	// @return com.vmware.model.Alarm
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateAlarmStatusSetStatus(alarmIdParam string, newStatusParam string, suppressDurationParam *int64) (model.Alarm, error)

	// Update event identified by event-id.
	//
	// @param eventIdParam (required)
	// @param monitoringEventParam (required)
	// @return com.vmware.model.MonitoringEvent
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateEvent(eventIdParam string, monitoringEventParam model.MonitoringEvent) (model.MonitoringEvent, error)
}

type systemAdministrationMonitoringAlarmsAndEventsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSystemAdministrationMonitoringAlarmsAndEventsClient(connector client.Connector) *systemAdministrationMonitoringAlarmsAndEventsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.system_administration_monitoring_alarms_and_events")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"bulk_update_alarms_set_status":  core.NewMethodIdentifier(interfaceIdentifier, "bulk_update_alarms_set_status"),
		"get_alarm":                      core.NewMethodIdentifier(interfaceIdentifier, "get_alarm"),
		"get_alarms":                     core.NewMethodIdentifier(interfaceIdentifier, "get_alarms"),
		"get_event":                      core.NewMethodIdentifier(interfaceIdentifier, "get_event"),
		"list_events":                    core.NewMethodIdentifier(interfaceIdentifier, "list_events"),
		"reset_event_values_set_default": core.NewMethodIdentifier(interfaceIdentifier, "reset_event_values_set_default"),
		"update_alarm_status_set_status": core.NewMethodIdentifier(interfaceIdentifier, "update_alarm_status_set_status"),
		"update_event":                   core.NewMethodIdentifier(interfaceIdentifier, "update_event"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := systemAdministrationMonitoringAlarmsAndEventsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *systemAdministrationMonitoringAlarmsAndEventsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *systemAdministrationMonitoringAlarmsAndEventsClient) BulkUpdateAlarmsSetStatus(newStatusParam string, afterParam *int64, beforeParam *int64, cursorParam *string, eventTypeParam *string, featureNameParam *string, idParam *string, intentPathParam *string, nodeIdParam *string, nodeResourceTypeParam *string, pageSizeParam *int64, severityParam *string, sortAscendingParam *bool, sortByParam *string, statusParam *string, suppressDurationParam *int64) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationMonitoringAlarmsAndEventsBulkUpdateAlarmsSetStatusInputType(), typeConverter)
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
	operationRestMetaData := systemAdministrationMonitoringAlarmsAndEventsBulkUpdateAlarmsSetStatusRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_monitoring_alarms_and_events", "bulk_update_alarms_set_status", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (sIface *systemAdministrationMonitoringAlarmsAndEventsClient) GetAlarm(alarmIdParam string) (model.Alarm, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationMonitoringAlarmsAndEventsGetAlarmInputType(), typeConverter)
	sv.AddStructField("AlarmId", alarmIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.Alarm
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationMonitoringAlarmsAndEventsGetAlarmRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_monitoring_alarms_and_events", "get_alarm", inputDataValue, executionContext)
	var emptyOutput model.Alarm
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationMonitoringAlarmsAndEventsGetAlarmOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.Alarm), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationMonitoringAlarmsAndEventsClient) GetAlarms(afterParam *int64, beforeParam *int64, cursorParam *string, eventTypeParam *string, featureNameParam *string, idParam *string, intentPathParam *string, nodeIdParam *string, nodeResourceTypeParam *string, pageSizeParam *int64, severityParam *string, sortAscendingParam *bool, sortByParam *string, statusParam *string) (model.AlarmsListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationMonitoringAlarmsAndEventsGetAlarmsInputType(), typeConverter)
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
	operationRestMetaData := systemAdministrationMonitoringAlarmsAndEventsGetAlarmsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_monitoring_alarms_and_events", "get_alarms", inputDataValue, executionContext)
	var emptyOutput model.AlarmsListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationMonitoringAlarmsAndEventsGetAlarmsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.AlarmsListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationMonitoringAlarmsAndEventsClient) GetEvent(eventIdParam string) (model.MonitoringEvent, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationMonitoringAlarmsAndEventsGetEventInputType(), typeConverter)
	sv.AddStructField("EventId", eventIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.MonitoringEvent
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationMonitoringAlarmsAndEventsGetEventRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_monitoring_alarms_and_events", "get_event", inputDataValue, executionContext)
	var emptyOutput model.MonitoringEvent
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationMonitoringAlarmsAndEventsGetEventOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.MonitoringEvent), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationMonitoringAlarmsAndEventsClient) ListEvents() (model.EventListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationMonitoringAlarmsAndEventsListEventsInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.EventListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationMonitoringAlarmsAndEventsListEventsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_monitoring_alarms_and_events", "list_events", inputDataValue, executionContext)
	var emptyOutput model.EventListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationMonitoringAlarmsAndEventsListEventsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.EventListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationMonitoringAlarmsAndEventsClient) ResetEventValuesSetDefault(eventIdParam string) (model.MonitoringEvent, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationMonitoringAlarmsAndEventsResetEventValuesSetDefaultInputType(), typeConverter)
	sv.AddStructField("EventId", eventIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.MonitoringEvent
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationMonitoringAlarmsAndEventsResetEventValuesSetDefaultRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_monitoring_alarms_and_events", "reset_event_values_set_default", inputDataValue, executionContext)
	var emptyOutput model.MonitoringEvent
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationMonitoringAlarmsAndEventsResetEventValuesSetDefaultOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.MonitoringEvent), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationMonitoringAlarmsAndEventsClient) UpdateAlarmStatusSetStatus(alarmIdParam string, newStatusParam string, suppressDurationParam *int64) (model.Alarm, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationMonitoringAlarmsAndEventsUpdateAlarmStatusSetStatusInputType(), typeConverter)
	sv.AddStructField("AlarmId", alarmIdParam)
	sv.AddStructField("NewStatus", newStatusParam)
	sv.AddStructField("SuppressDuration", suppressDurationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.Alarm
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationMonitoringAlarmsAndEventsUpdateAlarmStatusSetStatusRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_monitoring_alarms_and_events", "update_alarm_status_set_status", inputDataValue, executionContext)
	var emptyOutput model.Alarm
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationMonitoringAlarmsAndEventsUpdateAlarmStatusSetStatusOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.Alarm), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationMonitoringAlarmsAndEventsClient) UpdateEvent(eventIdParam string, monitoringEventParam model.MonitoringEvent) (model.MonitoringEvent, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationMonitoringAlarmsAndEventsUpdateEventInputType(), typeConverter)
	sv.AddStructField("EventId", eventIdParam)
	sv.AddStructField("MonitoringEvent", monitoringEventParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.MonitoringEvent
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationMonitoringAlarmsAndEventsUpdateEventRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_monitoring_alarms_and_events", "update_event", inputDataValue, executionContext)
	var emptyOutput model.MonitoringEvent
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationMonitoringAlarmsAndEventsUpdateEventOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.MonitoringEvent), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
