// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: SystemAdministrationMonitoringDashboardsWidgetConfiguration
// Used by client-side stubs.

package api

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/model"
)

const _ = core.SupportedByRuntimeVersion1

type SystemAdministrationMonitoringDashboardsWidgetConfigurationClient interface {

	// Creates a new Widget Configuration and adds it to the specified view. Supported resource_types are LabelValueConfiguration, DonutConfiguration, GridConfiguration, StatsConfiguration, MultiWidgetConfiguration, GraphConfiguration and ContainerConfiguration. Note: Expressions should be given in a single line. If an expression spans multiple lines, then form the expression in a single line. For label-value pairs, expressions are evaluated as follows: a. First, render configurations are evaluated in their order of appearance in the widget config. The 'field' is evaluated at the end. b. Second, when render configuration is provided then the order of evaluation is 1. If expressions provided in 'condition' and 'display value' are well-formed and free of runtime-errors such as 'null pointers' and evaluates to 'true'; Then remaining render configurations are not evaluated, and the current render configuration's 'display value' is taken as the final value. 2. If expression provided in 'condition' of render configuration is false, then next render configuration is evaluated. 3. Finally, 'field' is evaluated only when every render configuration evaluates to false and no error occurs during steps 1 and 2 above. If an error occurs during evaluation of render configuration, then an error message is shown. The display value corresponding to that label is not shown and evaluation of the remaining render configurations continues to collect and show all the error messages (marked with the 'Label' for identification) as 'Error_Messages: {}'. If during evaluation of expressions for any label-value pair an error occurs, then it is marked with error. The errors are shown in the report, along with the label value pairs that are error-free. Important: For elements that take expressions, strings should be provided by escaping them with a back-slash. These elements are - condition, field, tooltip text and render_configuration's display_value.
	//
	// @param viewIdParam (required)
	// @param widgetConfigurationParam (required)
	// The parameter must contain all the properties defined in model.WidgetConfiguration.
	// @return com.vmware.model.WidgetConfiguration
	// The return value will contain all the properties defined in model.WidgetConfiguration.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	CreateWidgetConfiguration(viewIdParam string, widgetConfigurationParam *data.StructValue) (*data.StructValue, error)

	// Detaches widget from a given view. If the widget is no longer part of any view, then it will be purged.
	//
	// @param viewIdParam (required)
	// @param widgetconfigurationIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteWidgetConfiguration(viewIdParam string, widgetconfigurationIdParam string) error

	// Returns Information about a specific Widget Configuration.
	//
	// @param viewIdParam (required)
	// @param widgetconfigurationIdParam (required)
	// @return com.vmware.model.WidgetConfiguration
	// The return value will contain all the properties defined in model.WidgetConfiguration.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetWidgetConfiguration(viewIdParam string, widgetconfigurationIdParam string) (*data.StructValue, error)

	// If no query params are specified then all the Widget Configurations of the specified view are returned.
	//
	// @param viewIdParam (required)
	// @param containerParam Id of the container (optional)
	// @param widgetIdsParam Ids of the WidgetConfigurations (optional)
	// @return com.vmware.model.WidgetConfigurationList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListWidgetConfigurations(viewIdParam string, containerParam *string, widgetIdsParam *string) (model.WidgetConfigurationList, error)

	// Updates the widget at the given view. If the widget is referenced by other views, then the widget will be updated in all the views that it is part of.
	//
	// @param viewIdParam (required)
	// @param widgetconfigurationIdParam (required)
	// @param widgetConfigurationParam (required)
	// The parameter must contain all the properties defined in model.WidgetConfiguration.
	// @return com.vmware.model.WidgetConfiguration
	// The return value will contain all the properties defined in model.WidgetConfiguration.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateWidgetConfiguration(viewIdParam string, widgetconfigurationIdParam string, widgetConfigurationParam *data.StructValue) (*data.StructValue, error)
}

type systemAdministrationMonitoringDashboardsWidgetConfigurationClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSystemAdministrationMonitoringDashboardsWidgetConfigurationClient(connector client.Connector) *systemAdministrationMonitoringDashboardsWidgetConfigurationClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.system_administration_monitoring_dashboards_widget_configuration")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create_widget_configuration": core.NewMethodIdentifier(interfaceIdentifier, "create_widget_configuration"),
		"delete_widget_configuration": core.NewMethodIdentifier(interfaceIdentifier, "delete_widget_configuration"),
		"get_widget_configuration":    core.NewMethodIdentifier(interfaceIdentifier, "get_widget_configuration"),
		"list_widget_configurations":  core.NewMethodIdentifier(interfaceIdentifier, "list_widget_configurations"),
		"update_widget_configuration": core.NewMethodIdentifier(interfaceIdentifier, "update_widget_configuration"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := systemAdministrationMonitoringDashboardsWidgetConfigurationClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *systemAdministrationMonitoringDashboardsWidgetConfigurationClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *systemAdministrationMonitoringDashboardsWidgetConfigurationClient) CreateWidgetConfiguration(viewIdParam string, widgetConfigurationParam *data.StructValue) (*data.StructValue, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationMonitoringDashboardsWidgetConfigurationCreateWidgetConfigurationInputType(), typeConverter)
	sv.AddStructField("ViewId", viewIdParam)
	sv.AddStructField("WidgetConfiguration", widgetConfigurationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationMonitoringDashboardsWidgetConfigurationCreateWidgetConfigurationRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_monitoring_dashboards_widget_configuration", "create_widget_configuration", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationMonitoringDashboardsWidgetConfigurationCreateWidgetConfigurationOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(*data.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationMonitoringDashboardsWidgetConfigurationClient) DeleteWidgetConfiguration(viewIdParam string, widgetconfigurationIdParam string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationMonitoringDashboardsWidgetConfigurationDeleteWidgetConfigurationInputType(), typeConverter)
	sv.AddStructField("ViewId", viewIdParam)
	sv.AddStructField("WidgetconfigurationId", widgetconfigurationIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationMonitoringDashboardsWidgetConfigurationDeleteWidgetConfigurationRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_monitoring_dashboards_widget_configuration", "delete_widget_configuration", inputDataValue, executionContext)
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

func (sIface *systemAdministrationMonitoringDashboardsWidgetConfigurationClient) GetWidgetConfiguration(viewIdParam string, widgetconfigurationIdParam string) (*data.StructValue, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationMonitoringDashboardsWidgetConfigurationGetWidgetConfigurationInputType(), typeConverter)
	sv.AddStructField("ViewId", viewIdParam)
	sv.AddStructField("WidgetconfigurationId", widgetconfigurationIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationMonitoringDashboardsWidgetConfigurationGetWidgetConfigurationRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_monitoring_dashboards_widget_configuration", "get_widget_configuration", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationMonitoringDashboardsWidgetConfigurationGetWidgetConfigurationOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(*data.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationMonitoringDashboardsWidgetConfigurationClient) ListWidgetConfigurations(viewIdParam string, containerParam *string, widgetIdsParam *string) (model.WidgetConfigurationList, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationMonitoringDashboardsWidgetConfigurationListWidgetConfigurationsInputType(), typeConverter)
	sv.AddStructField("ViewId", viewIdParam)
	sv.AddStructField("Container", containerParam)
	sv.AddStructField("WidgetIds", widgetIdsParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.WidgetConfigurationList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationMonitoringDashboardsWidgetConfigurationListWidgetConfigurationsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_monitoring_dashboards_widget_configuration", "list_widget_configurations", inputDataValue, executionContext)
	var emptyOutput model.WidgetConfigurationList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationMonitoringDashboardsWidgetConfigurationListWidgetConfigurationsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.WidgetConfigurationList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationMonitoringDashboardsWidgetConfigurationClient) UpdateWidgetConfiguration(viewIdParam string, widgetconfigurationIdParam string, widgetConfigurationParam *data.StructValue) (*data.StructValue, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationMonitoringDashboardsWidgetConfigurationUpdateWidgetConfigurationInputType(), typeConverter)
	sv.AddStructField("ViewId", viewIdParam)
	sv.AddStructField("WidgetconfigurationId", widgetconfigurationIdParam)
	sv.AddStructField("WidgetConfiguration", widgetConfigurationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationMonitoringDashboardsWidgetConfigurationUpdateWidgetConfigurationRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_monitoring_dashboards_widget_configuration", "update_widget_configuration", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationMonitoringDashboardsWidgetConfigurationUpdateWidgetConfigurationOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(*data.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
