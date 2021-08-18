// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: DynamicPlugins
// Used by client-side stubs.

package sha

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = core.SupportedByRuntimeVersion1

type DynamicPluginsClient interface {

	// Read Sha dynamic plugin.
	//
	// @param pluginIdParam Plugin filename (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(pluginIdParam string) error

	// Read Sha dynamic plugin.
	//
	// @param pluginIdParam Plugin filename (required)
	// @return com.vmware.nsx_policy.model.ShaDynamicPlugin
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(pluginIdParam string) (model.ShaDynamicPlugin, error)

	// API will provide list of Sha dynamic plugins.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx_policy.model.ShaDynamicPluginListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.ShaDynamicPluginListResult, error)

	// Upload Sha dynamic plugin content.
	//
	// @param pluginIdParam Sha pre-defined plugin (required)
	// @param shaDynamicPluginParam (required)
	// @return com.vmware.nsx_policy.model.ShaDynamicPlugin
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(pluginIdParam string, shaDynamicPluginParam model.ShaDynamicPlugin) (model.ShaDynamicPlugin, error)

	// Create Sha dynamic plugin.
	//
	// @param pluginIdParam Sha plugin id (required)
	// @param shaDynamicPluginParam (required)
	// @return com.vmware.nsx_policy.model.ShaDynamicPlugin
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(pluginIdParam string, shaDynamicPluginParam model.ShaDynamicPlugin) (model.ShaDynamicPlugin, error)
}

type dynamicPluginsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewDynamicPluginsClient(connector client.Connector) *dynamicPluginsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.sha.dynamic_plugins")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"delete": core.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   core.NewMethodIdentifier(interfaceIdentifier, "list"),
		"patch":  core.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": core.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	dIface := dynamicPluginsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &dIface
}

func (dIface *dynamicPluginsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := dIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (dIface *dynamicPluginsClient) Delete(pluginIdParam string) error {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(dynamicPluginsDeleteInputType(), typeConverter)
	sv.AddStructField("PluginId", pluginIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := dynamicPluginsDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	dIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sha.dynamic_plugins", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (dIface *dynamicPluginsClient) Get(pluginIdParam string) (model.ShaDynamicPlugin, error) {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(dynamicPluginsGetInputType(), typeConverter)
	sv.AddStructField("PluginId", pluginIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ShaDynamicPlugin
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := dynamicPluginsGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	dIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sha.dynamic_plugins", "get", inputDataValue, executionContext)
	var emptyOutput model.ShaDynamicPlugin
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), dynamicPluginsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ShaDynamicPlugin), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (dIface *dynamicPluginsClient) List(cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.ShaDynamicPluginListResult, error) {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(dynamicPluginsListInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludeMarkForDeleteObjects", includeMarkForDeleteObjectsParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ShaDynamicPluginListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := dynamicPluginsListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	dIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sha.dynamic_plugins", "list", inputDataValue, executionContext)
	var emptyOutput model.ShaDynamicPluginListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), dynamicPluginsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ShaDynamicPluginListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (dIface *dynamicPluginsClient) Patch(pluginIdParam string, shaDynamicPluginParam model.ShaDynamicPlugin) (model.ShaDynamicPlugin, error) {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(dynamicPluginsPatchInputType(), typeConverter)
	sv.AddStructField("PluginId", pluginIdParam)
	sv.AddStructField("ShaDynamicPlugin", shaDynamicPluginParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ShaDynamicPlugin
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := dynamicPluginsPatchRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	dIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sha.dynamic_plugins", "patch", inputDataValue, executionContext)
	var emptyOutput model.ShaDynamicPlugin
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), dynamicPluginsPatchOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ShaDynamicPlugin), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (dIface *dynamicPluginsClient) Update(pluginIdParam string, shaDynamicPluginParam model.ShaDynamicPlugin) (model.ShaDynamicPlugin, error) {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(dynamicPluginsUpdateInputType(), typeConverter)
	sv.AddStructField("PluginId", pluginIdParam)
	sv.AddStructField("ShaDynamicPlugin", shaDynamicPluginParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ShaDynamicPlugin
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := dynamicPluginsUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	dIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sha.dynamic_plugins", "update", inputDataValue, executionContext)
	var emptyOutput model.ShaDynamicPlugin
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), dynamicPluginsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ShaDynamicPlugin), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
