// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: PreDefinedPlugins
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

type PreDefinedPluginsClient interface {

	// Read SHA dynamic plugin.
	//
	// @param pluginIdParam Sha pre-defined plugin (required)
	// @return com.vmware.nsx_policy.model.ShaPredefinedPlugin
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(pluginIdParam string) (model.ShaPredefinedPlugin, error)

	// API will provide list of Sha dynamic plugins.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx_policy.model.ShaPreDefinedPluginListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.ShaPreDefinedPluginListResult, error)
}

type preDefinedPluginsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewPreDefinedPluginsClient(connector client.Connector) *preDefinedPluginsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.sha.pre_defined_plugins")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get":  core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list": core.NewMethodIdentifier(interfaceIdentifier, "list"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	pIface := preDefinedPluginsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &pIface
}

func (pIface *preDefinedPluginsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := pIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (pIface *preDefinedPluginsClient) Get(pluginIdParam string) (model.ShaPredefinedPlugin, error) {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(preDefinedPluginsGetInputType(), typeConverter)
	sv.AddStructField("PluginId", pluginIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ShaPredefinedPlugin
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := preDefinedPluginsGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sha.pre_defined_plugins", "get", inputDataValue, executionContext)
	var emptyOutput model.ShaPredefinedPlugin
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), preDefinedPluginsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ShaPredefinedPlugin), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (pIface *preDefinedPluginsClient) List(cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.ShaPreDefinedPluginListResult, error) {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(preDefinedPluginsListInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludeMarkForDeleteObjects", includeMarkForDeleteObjectsParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ShaPreDefinedPluginListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := preDefinedPluginsListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sha.pre_defined_plugins", "list", inputDataValue, executionContext)
	var emptyOutput model.ShaPreDefinedPluginListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), preDefinedPluginsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ShaPreDefinedPluginListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
