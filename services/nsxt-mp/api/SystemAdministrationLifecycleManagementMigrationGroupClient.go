// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: SystemAdministrationLifecycleManagementMigrationGroup
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

type SystemAdministrationLifecycleManagementMigrationGroupClient interface {

	// Add migration units to specified migration unit group. The migration units will be added at the end of the migration unit list.
	//
	// @param groupIdParam (required)
	// @param migrationUnitListParam (required)
	// @return com.vmware.model.MigrationUnitList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	AddMigrationUnitsToGroupAddMigrationUnits(groupIdParam string, migrationUnitListParam model.MigrationUnitList) (model.MigrationUnitList, error)

	// Create a group of migration units.
	//
	// @param migrationUnitGroupParam (required)
	// @return com.vmware.model.MigrationUnitGroup
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	CreateMigrationUnitGroup(migrationUnitGroupParam model.MigrationUnitGroup) (model.MigrationUnitGroup, error)

	// Delete the specified group. NOTE - A group can be deleted only if it is empty. If user tries to delete a group containing one or more migration units, the operation will fail and an error will be returned.
	//
	// @param groupIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteMigrationUnitGroup(groupIdParam string) error

	// Returns information about a specific migration unit group in the migration plan. If request parameter summary is set to true, then only count of migration units will be returned, migration units list will be empty.
	//
	// @param groupIdParam (required)
	// @param summaryParam Flag indicating whether to return the summary (optional, default to false)
	// @return com.vmware.model.MigrationUnitGroup
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetMigrationUnitGroup(groupIdParam string, summaryParam *bool) (model.MigrationUnitGroup, error)

	// Return information of all migration unit groups in the migration plan. If request parameter summary is set to true, then only count of migration units will be returned, migration units list will be empty. If request parameter component type is specified, then all migration unit groups for that component will be returned.
	//
	// @param componentTypeParam Component type based on which migration unit groups to be filtered (optional)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param summaryParam Flag indicating whether to return summary (optional, default to false)
	// @param syncParam Synchronize before returning migration unit groups (optional, default to false)
	// @return com.vmware.model.MigrationUnitGroupAggregateInfoListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetMigrationUnitGroupAggregateInfo(componentTypeParam *string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, summaryParam *bool, syncParam *bool) (model.MigrationUnitGroupAggregateInfoListResult, error)

	// Get migration status for migration units in the specified group. User can specify whether to show only the migration units with errors.
	//
	// @param groupIdParam (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param hasErrorsParam Flag to indicate whether to return only migration units with errors (optional, default to false)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.model.MigrationUnitStatusListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetMigrationUnitGroupStatus(groupIdParam string, cursorParam *string, hasErrorsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.MigrationUnitStatusListResult, error)

	// Return information of all migration unit groups in the migration plan. If request parameter summary is set to true, then only count of migration units will be returned, migration units list will be empty. If request parameter component type is specified, then all migration unit groups for that component will be returned.
	//
	// @param componentTypeParam Component type based on which migration unit groups to be filtered (optional)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param summaryParam Flag indicating whether to return summary (optional, default to false)
	// @param syncParam Synchronize before returning migration unit groups (optional, default to false)
	// @return com.vmware.model.MigrationUnitGroupListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetMigrationUnitGroups(componentTypeParam *string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, summaryParam *bool, syncParam *bool) (model.MigrationUnitGroupListResult, error)

	// Get migration status for migration unit groups
	//
	// @param componentTypeParam Component type based on which migration unit groups to be filtered (optional)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.model.MigrationUnitGroupStatusListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetMigrationUnitGroupsStatus(componentTypeParam *string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.MigrationUnitGroupStatusListResult, error)

	// Reorder an migration unit group by placing it before/after the specified migration unit group.
	//
	// @param groupIdParam (required)
	// @param reorderMigrationRequestParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ReorderMigrationUnitGroupReorder(groupIdParam string, reorderMigrationRequestParam model.ReorderMigrationRequest) error

	// Reorder an migration unit within the migration unit group by placing it before/after the specified migration unit
	//
	// @param groupIdParam (required)
	// @param migrationUnitIdParam (required)
	// @param reorderMigrationRequestParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ReorderMigrationUnitReorder(groupIdParam string, migrationUnitIdParam string, reorderMigrationRequestParam model.ReorderMigrationRequest) error

	// Update the specified migration unit group. Removal of migration units from the group using this is not allowed. An error will be returned in that case.
	//
	// @param groupIdParam (required)
	// @param migrationUnitGroupParam (required)
	// @return com.vmware.model.MigrationUnitGroup
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateMigrationUnitGroup(groupIdParam string, migrationUnitGroupParam model.MigrationUnitGroup) (model.MigrationUnitGroup, error)
}

type systemAdministrationLifecycleManagementMigrationGroupClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSystemAdministrationLifecycleManagementMigrationGroupClient(connector client.Connector) *systemAdministrationLifecycleManagementMigrationGroupClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.system_administration_lifecycle_management_migration_group")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"add_migration_units_to_group_add_migration_units": core.NewMethodIdentifier(interfaceIdentifier, "add_migration_units_to_group_add_migration_units"),
		"create_migration_unit_group":                      core.NewMethodIdentifier(interfaceIdentifier, "create_migration_unit_group"),
		"delete_migration_unit_group":                      core.NewMethodIdentifier(interfaceIdentifier, "delete_migration_unit_group"),
		"get_migration_unit_group":                         core.NewMethodIdentifier(interfaceIdentifier, "get_migration_unit_group"),
		"get_migration_unit_group_aggregate_info":          core.NewMethodIdentifier(interfaceIdentifier, "get_migration_unit_group_aggregate_info"),
		"get_migration_unit_group_status":                  core.NewMethodIdentifier(interfaceIdentifier, "get_migration_unit_group_status"),
		"get_migration_unit_groups":                        core.NewMethodIdentifier(interfaceIdentifier, "get_migration_unit_groups"),
		"get_migration_unit_groups_status":                 core.NewMethodIdentifier(interfaceIdentifier, "get_migration_unit_groups_status"),
		"reorder_migration_unit_group_reorder":             core.NewMethodIdentifier(interfaceIdentifier, "reorder_migration_unit_group_reorder"),
		"reorder_migration_unit_reorder":                   core.NewMethodIdentifier(interfaceIdentifier, "reorder_migration_unit_reorder"),
		"update_migration_unit_group":                      core.NewMethodIdentifier(interfaceIdentifier, "update_migration_unit_group"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := systemAdministrationLifecycleManagementMigrationGroupClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *systemAdministrationLifecycleManagementMigrationGroupClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *systemAdministrationLifecycleManagementMigrationGroupClient) AddMigrationUnitsToGroupAddMigrationUnits(groupIdParam string, migrationUnitListParam model.MigrationUnitList) (model.MigrationUnitList, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationLifecycleManagementMigrationGroupAddMigrationUnitsToGroupAddMigrationUnitsInputType(), typeConverter)
	sv.AddStructField("GroupId", groupIdParam)
	sv.AddStructField("MigrationUnitList", migrationUnitListParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.MigrationUnitList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationLifecycleManagementMigrationGroupAddMigrationUnitsToGroupAddMigrationUnitsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_lifecycle_management_migration_group", "add_migration_units_to_group_add_migration_units", inputDataValue, executionContext)
	var emptyOutput model.MigrationUnitList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationLifecycleManagementMigrationGroupAddMigrationUnitsToGroupAddMigrationUnitsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.MigrationUnitList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationLifecycleManagementMigrationGroupClient) CreateMigrationUnitGroup(migrationUnitGroupParam model.MigrationUnitGroup) (model.MigrationUnitGroup, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationLifecycleManagementMigrationGroupCreateMigrationUnitGroupInputType(), typeConverter)
	sv.AddStructField("MigrationUnitGroup", migrationUnitGroupParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.MigrationUnitGroup
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationLifecycleManagementMigrationGroupCreateMigrationUnitGroupRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_lifecycle_management_migration_group", "create_migration_unit_group", inputDataValue, executionContext)
	var emptyOutput model.MigrationUnitGroup
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationLifecycleManagementMigrationGroupCreateMigrationUnitGroupOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.MigrationUnitGroup), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationLifecycleManagementMigrationGroupClient) DeleteMigrationUnitGroup(groupIdParam string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationLifecycleManagementMigrationGroupDeleteMigrationUnitGroupInputType(), typeConverter)
	sv.AddStructField("GroupId", groupIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationLifecycleManagementMigrationGroupDeleteMigrationUnitGroupRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_lifecycle_management_migration_group", "delete_migration_unit_group", inputDataValue, executionContext)
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

func (sIface *systemAdministrationLifecycleManagementMigrationGroupClient) GetMigrationUnitGroup(groupIdParam string, summaryParam *bool) (model.MigrationUnitGroup, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationLifecycleManagementMigrationGroupGetMigrationUnitGroupInputType(), typeConverter)
	sv.AddStructField("GroupId", groupIdParam)
	sv.AddStructField("Summary", summaryParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.MigrationUnitGroup
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationLifecycleManagementMigrationGroupGetMigrationUnitGroupRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_lifecycle_management_migration_group", "get_migration_unit_group", inputDataValue, executionContext)
	var emptyOutput model.MigrationUnitGroup
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationLifecycleManagementMigrationGroupGetMigrationUnitGroupOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.MigrationUnitGroup), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationLifecycleManagementMigrationGroupClient) GetMigrationUnitGroupAggregateInfo(componentTypeParam *string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, summaryParam *bool, syncParam *bool) (model.MigrationUnitGroupAggregateInfoListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationLifecycleManagementMigrationGroupGetMigrationUnitGroupAggregateInfoInputType(), typeConverter)
	sv.AddStructField("ComponentType", componentTypeParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("Summary", summaryParam)
	sv.AddStructField("Sync", syncParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.MigrationUnitGroupAggregateInfoListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationLifecycleManagementMigrationGroupGetMigrationUnitGroupAggregateInfoRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_lifecycle_management_migration_group", "get_migration_unit_group_aggregate_info", inputDataValue, executionContext)
	var emptyOutput model.MigrationUnitGroupAggregateInfoListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationLifecycleManagementMigrationGroupGetMigrationUnitGroupAggregateInfoOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.MigrationUnitGroupAggregateInfoListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationLifecycleManagementMigrationGroupClient) GetMigrationUnitGroupStatus(groupIdParam string, cursorParam *string, hasErrorsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.MigrationUnitStatusListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationLifecycleManagementMigrationGroupGetMigrationUnitGroupStatusInputType(), typeConverter)
	sv.AddStructField("GroupId", groupIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("HasErrors", hasErrorsParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.MigrationUnitStatusListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationLifecycleManagementMigrationGroupGetMigrationUnitGroupStatusRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_lifecycle_management_migration_group", "get_migration_unit_group_status", inputDataValue, executionContext)
	var emptyOutput model.MigrationUnitStatusListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationLifecycleManagementMigrationGroupGetMigrationUnitGroupStatusOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.MigrationUnitStatusListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationLifecycleManagementMigrationGroupClient) GetMigrationUnitGroups(componentTypeParam *string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, summaryParam *bool, syncParam *bool) (model.MigrationUnitGroupListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationLifecycleManagementMigrationGroupGetMigrationUnitGroupsInputType(), typeConverter)
	sv.AddStructField("ComponentType", componentTypeParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("Summary", summaryParam)
	sv.AddStructField("Sync", syncParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.MigrationUnitGroupListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationLifecycleManagementMigrationGroupGetMigrationUnitGroupsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_lifecycle_management_migration_group", "get_migration_unit_groups", inputDataValue, executionContext)
	var emptyOutput model.MigrationUnitGroupListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationLifecycleManagementMigrationGroupGetMigrationUnitGroupsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.MigrationUnitGroupListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationLifecycleManagementMigrationGroupClient) GetMigrationUnitGroupsStatus(componentTypeParam *string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.MigrationUnitGroupStatusListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationLifecycleManagementMigrationGroupGetMigrationUnitGroupsStatusInputType(), typeConverter)
	sv.AddStructField("ComponentType", componentTypeParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.MigrationUnitGroupStatusListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationLifecycleManagementMigrationGroupGetMigrationUnitGroupsStatusRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_lifecycle_management_migration_group", "get_migration_unit_groups_status", inputDataValue, executionContext)
	var emptyOutput model.MigrationUnitGroupStatusListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationLifecycleManagementMigrationGroupGetMigrationUnitGroupsStatusOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.MigrationUnitGroupStatusListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationLifecycleManagementMigrationGroupClient) ReorderMigrationUnitGroupReorder(groupIdParam string, reorderMigrationRequestParam model.ReorderMigrationRequest) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationLifecycleManagementMigrationGroupReorderMigrationUnitGroupReorderInputType(), typeConverter)
	sv.AddStructField("GroupId", groupIdParam)
	sv.AddStructField("ReorderMigrationRequest", reorderMigrationRequestParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationLifecycleManagementMigrationGroupReorderMigrationUnitGroupReorderRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_lifecycle_management_migration_group", "reorder_migration_unit_group_reorder", inputDataValue, executionContext)
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

func (sIface *systemAdministrationLifecycleManagementMigrationGroupClient) ReorderMigrationUnitReorder(groupIdParam string, migrationUnitIdParam string, reorderMigrationRequestParam model.ReorderMigrationRequest) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationLifecycleManagementMigrationGroupReorderMigrationUnitReorderInputType(), typeConverter)
	sv.AddStructField("GroupId", groupIdParam)
	sv.AddStructField("MigrationUnitId", migrationUnitIdParam)
	sv.AddStructField("ReorderMigrationRequest", reorderMigrationRequestParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationLifecycleManagementMigrationGroupReorderMigrationUnitReorderRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_lifecycle_management_migration_group", "reorder_migration_unit_reorder", inputDataValue, executionContext)
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

func (sIface *systemAdministrationLifecycleManagementMigrationGroupClient) UpdateMigrationUnitGroup(groupIdParam string, migrationUnitGroupParam model.MigrationUnitGroup) (model.MigrationUnitGroup, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationLifecycleManagementMigrationGroupUpdateMigrationUnitGroupInputType(), typeConverter)
	sv.AddStructField("GroupId", groupIdParam)
	sv.AddStructField("MigrationUnitGroup", migrationUnitGroupParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.MigrationUnitGroup
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationLifecycleManagementMigrationGroupUpdateMigrationUnitGroupRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_lifecycle_management_migration_group", "update_migration_unit_group", inputDataValue, executionContext)
	var emptyOutput model.MigrationUnitGroup
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationLifecycleManagementMigrationGroupUpdateMigrationUnitGroupOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.MigrationUnitGroup), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
