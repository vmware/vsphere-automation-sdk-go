// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: PluginProfiles
// Used by client-side stubs.

package sha

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = core.SupportedByRuntimeVersion1

type PluginProfilesClient interface {

	// Delete Sha profile.
	//
	// @param shaProfileIdParam Sha profile id (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(shaProfileIdParam string) error

	// API will return Sha profile.
	//
	// @param shaProfileIdParam Sha profile id (required)
	// @return com.vmware.nsx_policy.model.ShaPluginProfile
	// The return value will contain all the properties defined in model.ShaPluginProfile.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(shaProfileIdParam string) (*data.StructValue, error)

	// API will provide list of SHA profile.
	//
	// @param appliedToGroupPathParam String Path of the Policy group path (optional)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param pluginPathParam String Path of the sha plugin (optional)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx_policy.model.ShaPluginProfileListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(appliedToGroupPathParam *string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, pluginPathParam *string, sortAscendingParam *bool, sortByParam *string) (model.ShaPluginProfileListResult, error)

	// Create or Replace Sha profile.
	//
	// @param shaProfileIdParam Sha profile id (required)
	// @param shaPluginProfileParam (required)
	// The parameter must contain all the properties defined in model.ShaPluginProfile.
	// @return com.vmware.nsx_policy.model.ShaPluginProfile
	// The return value will contain all the properties defined in model.ShaPluginProfile.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(shaProfileIdParam string, shaPluginProfileParam *data.StructValue) (*data.StructValue, error)

	// Create or Replace Sha profile.
	//
	// @param shaProfileIdParam Sha profile id (required)
	// @param shaPluginProfileParam (required)
	// The parameter must contain all the properties defined in model.ShaPluginProfile.
	// @return com.vmware.nsx_policy.model.ShaPluginProfile
	// The return value will contain all the properties defined in model.ShaPluginProfile.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(shaProfileIdParam string, shaPluginProfileParam *data.StructValue) (*data.StructValue, error)
}

type pluginProfilesClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewPluginProfilesClient(connector client.Connector) *pluginProfilesClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.sha.plugin_profiles")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"delete": core.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   core.NewMethodIdentifier(interfaceIdentifier, "list"),
		"patch":  core.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": core.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	pIface := pluginProfilesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &pIface
}

func (pIface *pluginProfilesClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := pIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (pIface *pluginProfilesClient) Delete(shaProfileIdParam string) error {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(pluginProfilesDeleteInputType(), typeConverter)
	sv.AddStructField("ShaProfileId", shaProfileIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := pluginProfilesDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sha.plugin_profiles", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (pIface *pluginProfilesClient) Get(shaProfileIdParam string) (*data.StructValue, error) {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(pluginProfilesGetInputType(), typeConverter)
	sv.AddStructField("ShaProfileId", shaProfileIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := pluginProfilesGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sha.plugin_profiles", "get", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), pluginProfilesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(*data.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (pIface *pluginProfilesClient) List(appliedToGroupPathParam *string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, pluginPathParam *string, sortAscendingParam *bool, sortByParam *string) (model.ShaPluginProfileListResult, error) {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(pluginProfilesListInputType(), typeConverter)
	sv.AddStructField("AppliedToGroupPath", appliedToGroupPathParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludeMarkForDeleteObjects", includeMarkForDeleteObjectsParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("PluginPath", pluginPathParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ShaPluginProfileListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := pluginProfilesListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sha.plugin_profiles", "list", inputDataValue, executionContext)
	var emptyOutput model.ShaPluginProfileListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), pluginProfilesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ShaPluginProfileListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (pIface *pluginProfilesClient) Patch(shaProfileIdParam string, shaPluginProfileParam *data.StructValue) (*data.StructValue, error) {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(pluginProfilesPatchInputType(), typeConverter)
	sv.AddStructField("ShaProfileId", shaProfileIdParam)
	sv.AddStructField("ShaPluginProfile", shaPluginProfileParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := pluginProfilesPatchRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sha.plugin_profiles", "patch", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), pluginProfilesPatchOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(*data.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (pIface *pluginProfilesClient) Update(shaProfileIdParam string, shaPluginProfileParam *data.StructValue) (*data.StructValue, error) {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(pluginProfilesUpdateInputType(), typeConverter)
	sv.AddStructField("ShaProfileId", shaProfileIdParam)
	sv.AddStructField("ShaPluginProfile", shaPluginProfileParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := pluginProfilesUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sha.plugin_profiles", "update", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), pluginProfilesUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(*data.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
