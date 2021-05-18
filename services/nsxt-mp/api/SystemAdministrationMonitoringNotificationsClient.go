// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: SystemAdministrationMonitoringNotifications
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

type SystemAdministrationMonitoringNotificationsClient interface {

	// Add a new notification watcher.
	//
	// @param notificationWatcherParam (required)
	// @return com.vmware.model.NotificationWatcher
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	AddNotificationWatcher(notificationWatcherParam model.NotificationWatcher) (model.NotificationWatcher, error)

	// Add uri filters for the specified watcher ID.
	//
	// @param watcherIdParam (required)
	// @param notificationParam (required)
	// @return com.vmware.model.NotificationsList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	AddUriFiltersAddUriFilters(watcherIdParam string, notificationParam model.Notification) (model.NotificationsList, error)

	// Delete notification watcher.
	//
	// @param watcherIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteNotificationWatcher(watcherIdParam string) error

	// Delete uri filters for the specified watcher ID.
	//
	// @param watcherIdParam (required)
	// @param notificationParam (required)
	// @return com.vmware.model.NotificationsList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteUriFiltersDeleteUriFilters(watcherIdParam string, notificationParam model.Notification) (model.NotificationsList, error)

	// Returns notification watcher by watcher id.
	//
	// @param watcherIdParam (required)
	// @return com.vmware.model.NotificationWatcher
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetNotificationWatcher(watcherIdParam string) (model.NotificationWatcher, error)

	// Get notifications for the specified watcher ID.
	//
	// @param watcherIdParam (required)
	// @return com.vmware.model.NotificationsList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetNotifications(watcherIdParam string) (model.NotificationsList, error)

	// Returns a list of registered notification watchers.
	// @return com.vmware.model.NotificationWatcherListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListNotificationWatchers() (model.NotificationWatcherListResult, error)

	// Update notification watcher.
	//
	// @param watcherIdParam (required)
	// @param notificationWatcherParam (required)
	// @return com.vmware.model.NotificationWatcher
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateNotificationWatcher(watcherIdParam string, notificationWatcherParam model.NotificationWatcher) (model.NotificationWatcher, error)

	// Update notifications for the specified watcher ID.
	//
	// @param watcherIdParam (required)
	// @param notificationsListParam (required)
	// @return com.vmware.model.NotificationsList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateNotifications(watcherIdParam string, notificationsListParam model.NotificationsList) (model.NotificationsList, error)
}

type systemAdministrationMonitoringNotificationsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSystemAdministrationMonitoringNotificationsClient(connector client.Connector) *systemAdministrationMonitoringNotificationsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.system_administration_monitoring_notifications")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"add_notification_watcher":              core.NewMethodIdentifier(interfaceIdentifier, "add_notification_watcher"),
		"add_uri_filters_add_uri_filters":       core.NewMethodIdentifier(interfaceIdentifier, "add_uri_filters_add_uri_filters"),
		"delete_notification_watcher":           core.NewMethodIdentifier(interfaceIdentifier, "delete_notification_watcher"),
		"delete_uri_filters_delete_uri_filters": core.NewMethodIdentifier(interfaceIdentifier, "delete_uri_filters_delete_uri_filters"),
		"get_notification_watcher":              core.NewMethodIdentifier(interfaceIdentifier, "get_notification_watcher"),
		"get_notifications":                     core.NewMethodIdentifier(interfaceIdentifier, "get_notifications"),
		"list_notification_watchers":            core.NewMethodIdentifier(interfaceIdentifier, "list_notification_watchers"),
		"update_notification_watcher":           core.NewMethodIdentifier(interfaceIdentifier, "update_notification_watcher"),
		"update_notifications":                  core.NewMethodIdentifier(interfaceIdentifier, "update_notifications"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := systemAdministrationMonitoringNotificationsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *systemAdministrationMonitoringNotificationsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *systemAdministrationMonitoringNotificationsClient) AddNotificationWatcher(notificationWatcherParam model.NotificationWatcher) (model.NotificationWatcher, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationMonitoringNotificationsAddNotificationWatcherInputType(), typeConverter)
	sv.AddStructField("NotificationWatcher", notificationWatcherParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NotificationWatcher
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationMonitoringNotificationsAddNotificationWatcherRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_monitoring_notifications", "add_notification_watcher", inputDataValue, executionContext)
	var emptyOutput model.NotificationWatcher
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationMonitoringNotificationsAddNotificationWatcherOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NotificationWatcher), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationMonitoringNotificationsClient) AddUriFiltersAddUriFilters(watcherIdParam string, notificationParam model.Notification) (model.NotificationsList, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationMonitoringNotificationsAddUriFiltersAddUriFiltersInputType(), typeConverter)
	sv.AddStructField("WatcherId", watcherIdParam)
	sv.AddStructField("Notification", notificationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NotificationsList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationMonitoringNotificationsAddUriFiltersAddUriFiltersRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_monitoring_notifications", "add_uri_filters_add_uri_filters", inputDataValue, executionContext)
	var emptyOutput model.NotificationsList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationMonitoringNotificationsAddUriFiltersAddUriFiltersOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NotificationsList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationMonitoringNotificationsClient) DeleteNotificationWatcher(watcherIdParam string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationMonitoringNotificationsDeleteNotificationWatcherInputType(), typeConverter)
	sv.AddStructField("WatcherId", watcherIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationMonitoringNotificationsDeleteNotificationWatcherRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_monitoring_notifications", "delete_notification_watcher", inputDataValue, executionContext)
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

func (sIface *systemAdministrationMonitoringNotificationsClient) DeleteUriFiltersDeleteUriFilters(watcherIdParam string, notificationParam model.Notification) (model.NotificationsList, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationMonitoringNotificationsDeleteUriFiltersDeleteUriFiltersInputType(), typeConverter)
	sv.AddStructField("WatcherId", watcherIdParam)
	sv.AddStructField("Notification", notificationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NotificationsList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationMonitoringNotificationsDeleteUriFiltersDeleteUriFiltersRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_monitoring_notifications", "delete_uri_filters_delete_uri_filters", inputDataValue, executionContext)
	var emptyOutput model.NotificationsList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationMonitoringNotificationsDeleteUriFiltersDeleteUriFiltersOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NotificationsList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationMonitoringNotificationsClient) GetNotificationWatcher(watcherIdParam string) (model.NotificationWatcher, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationMonitoringNotificationsGetNotificationWatcherInputType(), typeConverter)
	sv.AddStructField("WatcherId", watcherIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NotificationWatcher
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationMonitoringNotificationsGetNotificationWatcherRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_monitoring_notifications", "get_notification_watcher", inputDataValue, executionContext)
	var emptyOutput model.NotificationWatcher
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationMonitoringNotificationsGetNotificationWatcherOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NotificationWatcher), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationMonitoringNotificationsClient) GetNotifications(watcherIdParam string) (model.NotificationsList, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationMonitoringNotificationsGetNotificationsInputType(), typeConverter)
	sv.AddStructField("WatcherId", watcherIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NotificationsList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationMonitoringNotificationsGetNotificationsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_monitoring_notifications", "get_notifications", inputDataValue, executionContext)
	var emptyOutput model.NotificationsList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationMonitoringNotificationsGetNotificationsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NotificationsList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationMonitoringNotificationsClient) ListNotificationWatchers() (model.NotificationWatcherListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationMonitoringNotificationsListNotificationWatchersInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NotificationWatcherListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationMonitoringNotificationsListNotificationWatchersRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_monitoring_notifications", "list_notification_watchers", inputDataValue, executionContext)
	var emptyOutput model.NotificationWatcherListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationMonitoringNotificationsListNotificationWatchersOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NotificationWatcherListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationMonitoringNotificationsClient) UpdateNotificationWatcher(watcherIdParam string, notificationWatcherParam model.NotificationWatcher) (model.NotificationWatcher, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationMonitoringNotificationsUpdateNotificationWatcherInputType(), typeConverter)
	sv.AddStructField("WatcherId", watcherIdParam)
	sv.AddStructField("NotificationWatcher", notificationWatcherParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NotificationWatcher
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationMonitoringNotificationsUpdateNotificationWatcherRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_monitoring_notifications", "update_notification_watcher", inputDataValue, executionContext)
	var emptyOutput model.NotificationWatcher
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationMonitoringNotificationsUpdateNotificationWatcherOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NotificationWatcher), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationMonitoringNotificationsClient) UpdateNotifications(watcherIdParam string, notificationsListParam model.NotificationsList) (model.NotificationsList, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationMonitoringNotificationsUpdateNotificationsInputType(), typeConverter)
	sv.AddStructField("WatcherId", watcherIdParam)
	sv.AddStructField("NotificationsList", notificationsListParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NotificationsList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationMonitoringNotificationsUpdateNotificationsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_monitoring_notifications", "update_notifications", inputDataValue, executionContext)
	var emptyOutput model.NotificationsList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationMonitoringNotificationsUpdateNotificationsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NotificationsList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
