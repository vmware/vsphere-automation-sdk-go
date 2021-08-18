// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: InterfaceGroups
// Used by client-side stubs.

package locale_services

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = core.SupportedByRuntimeVersion1

type InterfaceGroupsClient interface {

	// Delete Tier-0 Interface group
	//
	// @param tier0IdParam (required)
	// @param localeServiceIdParam (required)
	// @param labelIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(tier0IdParam string, localeServiceIdParam string, labelIdParam string) error

	// Read Tier-0 Interface group
	//
	// @param tier0IdParam (required)
	// @param localeServiceIdParam (required)
	// @param labelIdParam (required)
	// @return com.vmware.nsx_policy.model.Tier0InterfaceGroup
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(tier0IdParam string, localeServiceIdParam string, labelIdParam string) (model.Tier0InterfaceGroup, error)

	// Paginated list of all Tier-0 Interface groups under locale service.
	//
	// @param tier0IdParam (required)
	// @param localeServiceIdParam (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx_policy.model.Tier0InterfaceGroupListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(tier0IdParam string, localeServiceIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.Tier0InterfaceGroupListResult, error)

	// If an Interface group with the label-id is not already present, create a new Interface group. If it already exists, update the Interface group for specified attributes.
	//
	// @param tier0IdParam (required)
	// @param localeServiceIdParam (required)
	// @param labelIdParam (required)
	// @param tier0InterfaceGroupParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(tier0IdParam string, localeServiceIdParam string, labelIdParam string, tier0InterfaceGroupParam model.Tier0InterfaceGroup) error

	// Update the Interface group for specified attributes.
	//
	// @param tier0IdParam (required)
	// @param localeServiceIdParam (required)
	// @param labelIdParam (required)
	// @param tier0InterfaceGroupParam (required)
	// @return com.vmware.nsx_policy.model.Tier0InterfaceGroup
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(tier0IdParam string, localeServiceIdParam string, labelIdParam string, tier0InterfaceGroupParam model.Tier0InterfaceGroup) (model.Tier0InterfaceGroup, error)
}

type interfaceGroupsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewInterfaceGroupsClient(connector client.Connector) *interfaceGroupsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.tier_0s.locale_services.interface_groups")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"delete": core.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   core.NewMethodIdentifier(interfaceIdentifier, "list"),
		"patch":  core.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": core.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	iIface := interfaceGroupsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &iIface
}

func (iIface *interfaceGroupsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := iIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (iIface *interfaceGroupsClient) Delete(tier0IdParam string, localeServiceIdParam string, labelIdParam string) error {
	typeConverter := iIface.connector.TypeConverter()
	executionContext := iIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(interfaceGroupsDeleteInputType(), typeConverter)
	sv.AddStructField("Tier0Id", tier0IdParam)
	sv.AddStructField("LocaleServiceId", localeServiceIdParam)
	sv.AddStructField("LabelId", labelIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := interfaceGroupsDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	iIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := iIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.tier_0s.locale_services.interface_groups", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (iIface *interfaceGroupsClient) Get(tier0IdParam string, localeServiceIdParam string, labelIdParam string) (model.Tier0InterfaceGroup, error) {
	typeConverter := iIface.connector.TypeConverter()
	executionContext := iIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(interfaceGroupsGetInputType(), typeConverter)
	sv.AddStructField("Tier0Id", tier0IdParam)
	sv.AddStructField("LocaleServiceId", localeServiceIdParam)
	sv.AddStructField("LabelId", labelIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.Tier0InterfaceGroup
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := interfaceGroupsGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	iIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := iIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.tier_0s.locale_services.interface_groups", "get", inputDataValue, executionContext)
	var emptyOutput model.Tier0InterfaceGroup
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), interfaceGroupsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.Tier0InterfaceGroup), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (iIface *interfaceGroupsClient) List(tier0IdParam string, localeServiceIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.Tier0InterfaceGroupListResult, error) {
	typeConverter := iIface.connector.TypeConverter()
	executionContext := iIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(interfaceGroupsListInputType(), typeConverter)
	sv.AddStructField("Tier0Id", tier0IdParam)
	sv.AddStructField("LocaleServiceId", localeServiceIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludeMarkForDeleteObjects", includeMarkForDeleteObjectsParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.Tier0InterfaceGroupListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := interfaceGroupsListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	iIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := iIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.tier_0s.locale_services.interface_groups", "list", inputDataValue, executionContext)
	var emptyOutput model.Tier0InterfaceGroupListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), interfaceGroupsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.Tier0InterfaceGroupListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (iIface *interfaceGroupsClient) Patch(tier0IdParam string, localeServiceIdParam string, labelIdParam string, tier0InterfaceGroupParam model.Tier0InterfaceGroup) error {
	typeConverter := iIface.connector.TypeConverter()
	executionContext := iIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(interfaceGroupsPatchInputType(), typeConverter)
	sv.AddStructField("Tier0Id", tier0IdParam)
	sv.AddStructField("LocaleServiceId", localeServiceIdParam)
	sv.AddStructField("LabelId", labelIdParam)
	sv.AddStructField("Tier0InterfaceGroup", tier0InterfaceGroupParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := interfaceGroupsPatchRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	iIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := iIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.tier_0s.locale_services.interface_groups", "patch", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (iIface *interfaceGroupsClient) Update(tier0IdParam string, localeServiceIdParam string, labelIdParam string, tier0InterfaceGroupParam model.Tier0InterfaceGroup) (model.Tier0InterfaceGroup, error) {
	typeConverter := iIface.connector.TypeConverter()
	executionContext := iIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(interfaceGroupsUpdateInputType(), typeConverter)
	sv.AddStructField("Tier0Id", tier0IdParam)
	sv.AddStructField("LocaleServiceId", localeServiceIdParam)
	sv.AddStructField("LabelId", labelIdParam)
	sv.AddStructField("Tier0InterfaceGroup", tier0InterfaceGroupParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.Tier0InterfaceGroup
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := interfaceGroupsUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	iIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := iIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.tier_0s.locale_services.interface_groups", "update", inputDataValue, executionContext)
	var emptyOutput model.Tier0InterfaceGroup
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), interfaceGroupsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.Tier0InterfaceGroup), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
