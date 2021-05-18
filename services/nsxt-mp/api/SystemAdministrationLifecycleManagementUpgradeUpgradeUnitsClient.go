// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: SystemAdministrationLifecycleManagementUpgradeUpgradeUnits
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

type SystemAdministrationLifecycleManagementUpgradeUpgradeUnitsClient interface {

	// Get a specific upgrade unit
	//
	// @param upgradeUnitIdParam (required)
	// @return com.vmware.model.UpgradeUnit
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetUpgradeUnit(upgradeUnitIdParam string) (model.UpgradeUnit, error)

	// Get upgrade units aggregate-info
	//
	// @param componentTypeParam Component type based on which upgrade units to be filtered (optional)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param groupIdParam Identifier of group based on which upgrade units to be filtered (optional)
	// @param hasErrorsParam Flag to indicate whether to return only upgrade units with errors (optional, default to false)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param metadataParam Metadata about upgrade unit to filter on (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param selectionStatusParam Flag to indicate whether to return only selected, only deselected or both type of upgrade units (optional, default to ALL)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param upgradeUnitDisplayNameParam Display name of upgrade unit (optional)
	// @return com.vmware.model.UpgradeUnitAggregateInfoListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetUpgradeUnitAggregateInfo(componentTypeParam *string, cursorParam *string, groupIdParam *string, hasErrorsParam *bool, includedFieldsParam *string, metadataParam *string, pageSizeParam *int64, selectionStatusParam *string, sortAscendingParam *bool, sortByParam *string, upgradeUnitDisplayNameParam *string) (model.UpgradeUnitAggregateInfoListResult, error)

	// Get upgrade units
	//
	// @param componentTypeParam Component type based on which upgrade units to be filtered (optional)
	// @param currentVersionParam Current version of upgrade unit based on which upgrade units to be filtered (optional)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param groupIdParam UUID of group based on which upgrade units to be filtered (optional)
	// @param hasWarningsParam Flag to indicate whether to return only upgrade units with warnings (optional, default to false)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param metadataParam Metadata about upgrade unit to filter on (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param upgradeUnitTypeParam Upgrade unit type based on which upgrade units to be filtered (optional)
	// @return com.vmware.model.UpgradeUnitListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetUpgradeUnits(componentTypeParam *string, currentVersionParam *string, cursorParam *string, groupIdParam *string, hasWarningsParam *bool, includedFieldsParam *string, metadataParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, upgradeUnitTypeParam *string) (model.UpgradeUnitListResult, error)

	// Get upgrade units stats
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param syncParam Synchronize before returning upgrade unit stats (optional, default to false)
	// @return com.vmware.model.UpgradeUnitTypeStatsList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetUpgradeUnitsStats(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, syncParam *bool) (model.UpgradeUnitTypeStatsList, error)
}

type systemAdministrationLifecycleManagementUpgradeUpgradeUnitsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSystemAdministrationLifecycleManagementUpgradeUpgradeUnitsClient(connector client.Connector) *systemAdministrationLifecycleManagementUpgradeUpgradeUnitsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.system_administration_lifecycle_management_upgrade_upgrade_units")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get_upgrade_unit":                core.NewMethodIdentifier(interfaceIdentifier, "get_upgrade_unit"),
		"get_upgrade_unit_aggregate_info": core.NewMethodIdentifier(interfaceIdentifier, "get_upgrade_unit_aggregate_info"),
		"get_upgrade_units":               core.NewMethodIdentifier(interfaceIdentifier, "get_upgrade_units"),
		"get_upgrade_units_stats":         core.NewMethodIdentifier(interfaceIdentifier, "get_upgrade_units_stats"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := systemAdministrationLifecycleManagementUpgradeUpgradeUnitsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *systemAdministrationLifecycleManagementUpgradeUpgradeUnitsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *systemAdministrationLifecycleManagementUpgradeUpgradeUnitsClient) GetUpgradeUnit(upgradeUnitIdParam string) (model.UpgradeUnit, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationLifecycleManagementUpgradeUpgradeUnitsGetUpgradeUnitInputType(), typeConverter)
	sv.AddStructField("UpgradeUnitId", upgradeUnitIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.UpgradeUnit
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationLifecycleManagementUpgradeUpgradeUnitsGetUpgradeUnitRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_lifecycle_management_upgrade_upgrade_units", "get_upgrade_unit", inputDataValue, executionContext)
	var emptyOutput model.UpgradeUnit
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationLifecycleManagementUpgradeUpgradeUnitsGetUpgradeUnitOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.UpgradeUnit), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationLifecycleManagementUpgradeUpgradeUnitsClient) GetUpgradeUnitAggregateInfo(componentTypeParam *string, cursorParam *string, groupIdParam *string, hasErrorsParam *bool, includedFieldsParam *string, metadataParam *string, pageSizeParam *int64, selectionStatusParam *string, sortAscendingParam *bool, sortByParam *string, upgradeUnitDisplayNameParam *string) (model.UpgradeUnitAggregateInfoListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationLifecycleManagementUpgradeUpgradeUnitsGetUpgradeUnitAggregateInfoInputType(), typeConverter)
	sv.AddStructField("ComponentType", componentTypeParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("GroupId", groupIdParam)
	sv.AddStructField("HasErrors", hasErrorsParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("Metadata", metadataParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SelectionStatus", selectionStatusParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("UpgradeUnitDisplayName", upgradeUnitDisplayNameParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.UpgradeUnitAggregateInfoListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationLifecycleManagementUpgradeUpgradeUnitsGetUpgradeUnitAggregateInfoRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_lifecycle_management_upgrade_upgrade_units", "get_upgrade_unit_aggregate_info", inputDataValue, executionContext)
	var emptyOutput model.UpgradeUnitAggregateInfoListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationLifecycleManagementUpgradeUpgradeUnitsGetUpgradeUnitAggregateInfoOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.UpgradeUnitAggregateInfoListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationLifecycleManagementUpgradeUpgradeUnitsClient) GetUpgradeUnits(componentTypeParam *string, currentVersionParam *string, cursorParam *string, groupIdParam *string, hasWarningsParam *bool, includedFieldsParam *string, metadataParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, upgradeUnitTypeParam *string) (model.UpgradeUnitListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationLifecycleManagementUpgradeUpgradeUnitsGetUpgradeUnitsInputType(), typeConverter)
	sv.AddStructField("ComponentType", componentTypeParam)
	sv.AddStructField("CurrentVersion", currentVersionParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("GroupId", groupIdParam)
	sv.AddStructField("HasWarnings", hasWarningsParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("Metadata", metadataParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("UpgradeUnitType", upgradeUnitTypeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.UpgradeUnitListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationLifecycleManagementUpgradeUpgradeUnitsGetUpgradeUnitsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_lifecycle_management_upgrade_upgrade_units", "get_upgrade_units", inputDataValue, executionContext)
	var emptyOutput model.UpgradeUnitListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationLifecycleManagementUpgradeUpgradeUnitsGetUpgradeUnitsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.UpgradeUnitListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationLifecycleManagementUpgradeUpgradeUnitsClient) GetUpgradeUnitsStats(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, syncParam *bool) (model.UpgradeUnitTypeStatsList, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationLifecycleManagementUpgradeUpgradeUnitsGetUpgradeUnitsStatsInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("Sync", syncParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.UpgradeUnitTypeStatsList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationLifecycleManagementUpgradeUpgradeUnitsGetUpgradeUnitsStatsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_lifecycle_management_upgrade_upgrade_units", "get_upgrade_units_stats", inputDataValue, executionContext)
	var emptyOutput model.UpgradeUnitTypeStatsList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationLifecycleManagementUpgradeUpgradeUnitsGetUpgradeUnitsStatsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.UpgradeUnitTypeStatsList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
