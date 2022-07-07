// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Widgetconfigurations
// Used by client-side stubs.

package ui_views

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type WidgetconfigurationsClient interface {

	// Creates a new Widget Configuration and adds it to the specified view. Supported resource_types are LabelValueConfiguration, DonutConfiguration, GridConfiguration, StatsConfiguration, MultiWidgetConfiguration, GraphConfiguration and ContainerConfiguration. Note: Expressions should be given in a single line. If an expression spans multiple lines, then form the expression in a single line. For label-value pairs, expressions are evaluated as follows: a. First, render configurations are evaluated in their order of appearance in the widget config. The 'field' is evaluated at the end. b. Second, when render configuration is provided then the order of evaluation is 1. If expressions provided in 'condition' and 'display value' are well-formed and free of runtime-errors such as 'null pointers' and evaluates to 'true'; Then remaining render configurations are not evaluated, and the current render configuration's 'display value' is taken as the final value. 2. If expression provided in 'condition' of render configuration is false, then next render configuration is evaluated. 3. Finally, 'field' is evaluated only when every render configuration evaluates to false and no error occurs during steps 1 and 2 above. If an error occurs during evaluation of render configuration, then an error message is shown. The display value corresponding to that label is not shown and evaluation of the remaining render configurations continues to collect and show all the error messages (marked with the 'Label' for identification) as 'Error_Messages: {}'. If during evaluation of expressions for any label-value pair an error occurs, then it is marked with error. The errors are shown in the report, along with the label value pairs that are error-free. Important: For elements that take expressions, strings should be provided by escaping them with a back-slash. These elements are - condition, field, tooltip text and render_configuration's display_value.
	//
	// @param viewIdParam (required)
	// @param widgetConfigurationParam (required)
	// The parameter must contain all the properties defined in model.WidgetConfiguration.
	// @return com.vmware.nsx.model.WidgetConfiguration
	// The return value will contain all the properties defined in model.WidgetConfiguration.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(viewIdParam string, widgetConfigurationParam *data.StructValue) (*data.StructValue, error)

	// Detaches widget from a given view. If the widget is no longer part of any view, then it will be purged.
	//
	// @param viewIdParam (required)
	// @param widgetconfigurationIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(viewIdParam string, widgetconfigurationIdParam string) error

	// If no query params are specified then all the Widget Configurations of the specified view are returned.
	//
	// @param viewIdParam (required)
	// @param containerParam Id of the container (optional)
	// @param widgetIdsParam Ids of the WidgetConfigurations (optional)
	// @return com.vmware.nsx.model.WidgetConfigurationList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(viewIdParam string, containerParam *string, widgetIdsParam *string) (model.WidgetConfigurationList, error)

	// Returns Information about a specific Widget Configuration.
	//
	// @param viewIdParam (required)
	// @param widgetconfigurationIdParam (required)
	// @return com.vmware.nsx.model.WidgetConfiguration
	// The return value will contain all the properties defined in model.WidgetConfiguration.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get0(viewIdParam string, widgetconfigurationIdParam string) (*data.StructValue, error)

	// Updates the widget at the given view. If the widget is referenced by other views, then the widget will be updated in all the views that it is part of.
	//
	// @param viewIdParam (required)
	// @param widgetconfigurationIdParam (required)
	// @param widgetConfigurationParam (required)
	// The parameter must contain all the properties defined in model.WidgetConfiguration.
	// @return com.vmware.nsx.model.WidgetConfiguration
	// The return value will contain all the properties defined in model.WidgetConfiguration.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(viewIdParam string, widgetconfigurationIdParam string, widgetConfigurationParam *data.StructValue) (*data.StructValue, error)
}

type widgetconfigurationsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewWidgetconfigurationsClient(connector client.Connector) *widgetconfigurationsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.ui_views.widgetconfigurations")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create": core.NewMethodIdentifier(interfaceIdentifier, "create"),
		"delete": core.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"get_0":  core.NewMethodIdentifier(interfaceIdentifier, "get_0"),
		"update": core.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	wIface := widgetconfigurationsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &wIface
}

func (wIface *widgetconfigurationsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := wIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (wIface *widgetconfigurationsClient) Create(viewIdParam string, widgetConfigurationParam *data.StructValue) (*data.StructValue, error) {
	typeConverter := wIface.connector.TypeConverter()
	executionContext := wIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(widgetconfigurationsCreateInputType(), typeConverter)
	sv.AddStructField("ViewId", viewIdParam)
	sv.AddStructField("WidgetConfiguration", widgetConfigurationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := widgetconfigurationsCreateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	wIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := wIface.connector.GetApiProvider().Invoke("com.vmware.nsx.ui_views.widgetconfigurations", "create", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), widgetconfigurationsCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(*data.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), wIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (wIface *widgetconfigurationsClient) Delete(viewIdParam string, widgetconfigurationIdParam string) error {
	typeConverter := wIface.connector.TypeConverter()
	executionContext := wIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(widgetconfigurationsDeleteInputType(), typeConverter)
	sv.AddStructField("ViewId", viewIdParam)
	sv.AddStructField("WidgetconfigurationId", widgetconfigurationIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := widgetconfigurationsDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	wIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := wIface.connector.GetApiProvider().Invoke("com.vmware.nsx.ui_views.widgetconfigurations", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), wIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (wIface *widgetconfigurationsClient) Get(viewIdParam string, containerParam *string, widgetIdsParam *string) (model.WidgetConfigurationList, error) {
	typeConverter := wIface.connector.TypeConverter()
	executionContext := wIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(widgetconfigurationsGetInputType(), typeConverter)
	sv.AddStructField("ViewId", viewIdParam)
	sv.AddStructField("Container", containerParam)
	sv.AddStructField("WidgetIds", widgetIdsParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.WidgetConfigurationList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := widgetconfigurationsGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	wIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := wIface.connector.GetApiProvider().Invoke("com.vmware.nsx.ui_views.widgetconfigurations", "get", inputDataValue, executionContext)
	var emptyOutput model.WidgetConfigurationList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), widgetconfigurationsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.WidgetConfigurationList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), wIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (wIface *widgetconfigurationsClient) Get0(viewIdParam string, widgetconfigurationIdParam string) (*data.StructValue, error) {
	typeConverter := wIface.connector.TypeConverter()
	executionContext := wIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(widgetconfigurationsGet0InputType(), typeConverter)
	sv.AddStructField("ViewId", viewIdParam)
	sv.AddStructField("WidgetconfigurationId", widgetconfigurationIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := widgetconfigurationsGet0RestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	wIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := wIface.connector.GetApiProvider().Invoke("com.vmware.nsx.ui_views.widgetconfigurations", "get_0", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), widgetconfigurationsGet0OutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(*data.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), wIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (wIface *widgetconfigurationsClient) Update(viewIdParam string, widgetconfigurationIdParam string, widgetConfigurationParam *data.StructValue) (*data.StructValue, error) {
	typeConverter := wIface.connector.TypeConverter()
	executionContext := wIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(widgetconfigurationsUpdateInputType(), typeConverter)
	sv.AddStructField("ViewId", viewIdParam)
	sv.AddStructField("WidgetconfigurationId", widgetconfigurationIdParam)
	sv.AddStructField("WidgetConfiguration", widgetConfigurationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := widgetconfigurationsUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	wIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := wIface.connector.GetApiProvider().Invoke("com.vmware.nsx.ui_views.widgetconfigurations", "update", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), widgetconfigurationsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(*data.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), wIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
