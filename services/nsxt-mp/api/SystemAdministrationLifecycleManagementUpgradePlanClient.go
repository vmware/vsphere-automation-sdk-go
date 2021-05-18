// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: SystemAdministrationLifecycleManagementUpgradePlan
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

type SystemAdministrationLifecycleManagementUpgradePlanClient interface {

	// Aborts execution of pre-upgrade checks if already in progress. Halts the execution of checks awaiting execution at this point and makes best-effort attempts to stop checks already in execution. Returns without action if execution of pre-upgrade checks is not in progress.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	AbortPreUpgradeChecksAbortPreUpgradeChecks() error

	// Continue the upgrade. Resumes the upgrade from the point where it was paused.
	//
	// @param componentTypeParam Component to upgrade. (optional)
	// @param skipParam Skip to upgrade of next component. (optional, default to false)
	// @throws ConcurrentChange  Conflict
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ContinueUpgradeContinue(componentTypeParam *string, skipParam *bool) error

	// Run pre-defined checks to identify issues after upgrade of a component. The results of the checks are added to the respective upgrade units aggregate-info. The progress and status of post-upgrade checks is part of aggregate-info of individual upgrade unit groups. Returns HTTP status 500 with error code 30953 if execution of post-upgrade checks is already in progress.
	//
	// @param componentTypeParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ExecutePostUpgradeChecksExecutePostUpgradeChecks(componentTypeParam string) error

	// Run pre-defined checks to identify potential issues which can be encountered during an upgrade or can cause an upgrade to fail. The results of the checks are added to the respective upgrade units aggregate-info. The progress and status of operation is part of upgrade status summary of individual components. Returns HTTP status 500 with error code 30953 if execution of pre-upgrade checks is already in progress.
	//
	// @param componentTypeParam Component type on which the action is performed or on which the results are filtered (optional)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ExecutePreUpgradeChecksExecutePreUpgradeChecks(componentTypeParam *string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) error

	// Returns pre-upgrade checks in csv format
	// @return com.vmware.model.UpgradeCheckCsvListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetAllPreUpgradeChecksInCsvFormatCsv() (model.UpgradeCheckCsvListResult, error)

	// Get failures resulting from the last execution of pre-upgrade checks. If the execution of checks is in progress, the response has the list of failures observed so far.
	//
	// @param componentTypeParam Component type (optional)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param filterTextParam Filter text (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param originTypeParam Type of origin of failure (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param type_Param Status of the upgrade check (optional)
	// @return com.vmware.model.UpgradeCheckFailureListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetPreUpgradeCheckFailures(componentTypeParam *string, cursorParam *string, filterTextParam *string, includedFieldsParam *string, originTypeParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, type_Param *string) (model.UpgradeCheckFailureListResult, error)

	// Returns information of pre-upgrade and post-upgrade checks. If request parameter component type is specified, then returns information about all pre-upgrade and post-upgrade for the component. Otherwise returns information of checks across all component types.
	//
	// @param componentTypeParam Component type based on which upgrade checks are to be filtered (optional)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.model.ComponentUpgradeChecksInfoListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetUpgradeChecksInfo(componentTypeParam *string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.ComponentUpgradeChecksInfoListResult, error)

	// Get the upgrade plan settings for the component.
	//
	// @param componentTypeParam (required)
	// @return com.vmware.model.UpgradePlanSettings
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetUpgradePlanSettings(componentTypeParam string) (model.UpgradePlanSettings, error)

	// Pause the upgrade. Upgrade will be paused after upgrade of all the nodes currently in progress is completed either successfully or with failure. User can make changes in the upgrade plan when the upgrade is paused.
	// @throws ConcurrentChange  Conflict
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	PauseUpgradePause() error

	// Reset the upgrade plan to default plan. User has an option to change the default plan. But if after making changes, user wants to go back to the default plan, this is the way to do so.
	//
	// @param componentTypeParam Component type (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ResetUpgradePlanReset(componentTypeParam string) error

	// This will stage the upgrade bits so that upgrade can be driven from outside NSX upgrade-coordinator. If component type is not provided as a request parameter then upgrade staging will be performed on all the components.
	//
	// @param componentTypeParam Component type (optional)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	StageUpgradeStageUpgrade(componentTypeParam *string) error

	// Start the upgrade. Upgrade will start as per the upgrade plan.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	StartUpgradeStart() error

	// Update the upgrade plan settings for the component.
	//
	// @param componentTypeParam (required)
	// @param upgradePlanSettingsParam (required)
	// @return com.vmware.model.UpgradePlanSettings
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateUpgradePlanSettings(componentTypeParam string, upgradePlanSettingsParam model.UpgradePlanSettings) (model.UpgradePlanSettings, error)

	// Upgrades, Resumes the upgrade of a selected set of units.
	//
	// @param upgradeUnitListParam (required)
	// @throws ConcurrentChange  Conflict
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpgradeSelectedUnitsUpgradeSelectedUnits(upgradeUnitListParam model.UpgradeUnitList) error
}

type systemAdministrationLifecycleManagementUpgradePlanClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSystemAdministrationLifecycleManagementUpgradePlanClient(connector client.Connector) *systemAdministrationLifecycleManagementUpgradePlanClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.system_administration_lifecycle_management_upgrade_plan")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"abort_pre_upgrade_checks_abort_pre_upgrade_checks":       core.NewMethodIdentifier(interfaceIdentifier, "abort_pre_upgrade_checks_abort_pre_upgrade_checks"),
		"continue_upgrade_continue":                               core.NewMethodIdentifier(interfaceIdentifier, "continue_upgrade_continue"),
		"execute_post_upgrade_checks_execute_post_upgrade_checks": core.NewMethodIdentifier(interfaceIdentifier, "execute_post_upgrade_checks_execute_post_upgrade_checks"),
		"execute_pre_upgrade_checks_execute_pre_upgrade_checks":   core.NewMethodIdentifier(interfaceIdentifier, "execute_pre_upgrade_checks_execute_pre_upgrade_checks"),
		"get_all_pre_upgrade_checks_in_csv_format_csv":            core.NewMethodIdentifier(interfaceIdentifier, "get_all_pre_upgrade_checks_in_csv_format_csv"),
		"get_pre_upgrade_check_failures":                          core.NewMethodIdentifier(interfaceIdentifier, "get_pre_upgrade_check_failures"),
		"get_upgrade_checks_info":                                 core.NewMethodIdentifier(interfaceIdentifier, "get_upgrade_checks_info"),
		"get_upgrade_plan_settings":                               core.NewMethodIdentifier(interfaceIdentifier, "get_upgrade_plan_settings"),
		"pause_upgrade_pause":                                     core.NewMethodIdentifier(interfaceIdentifier, "pause_upgrade_pause"),
		"reset_upgrade_plan_reset":                                core.NewMethodIdentifier(interfaceIdentifier, "reset_upgrade_plan_reset"),
		"stage_upgrade_stage_upgrade":                             core.NewMethodIdentifier(interfaceIdentifier, "stage_upgrade_stage_upgrade"),
		"start_upgrade_start":                                     core.NewMethodIdentifier(interfaceIdentifier, "start_upgrade_start"),
		"update_upgrade_plan_settings":                            core.NewMethodIdentifier(interfaceIdentifier, "update_upgrade_plan_settings"),
		"upgrade_selected_units_upgrade_selected_units":           core.NewMethodIdentifier(interfaceIdentifier, "upgrade_selected_units_upgrade_selected_units"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := systemAdministrationLifecycleManagementUpgradePlanClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *systemAdministrationLifecycleManagementUpgradePlanClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *systemAdministrationLifecycleManagementUpgradePlanClient) AbortPreUpgradeChecksAbortPreUpgradeChecks() error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationLifecycleManagementUpgradePlanAbortPreUpgradeChecksAbortPreUpgradeChecksInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationLifecycleManagementUpgradePlanAbortPreUpgradeChecksAbortPreUpgradeChecksRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_lifecycle_management_upgrade_plan", "abort_pre_upgrade_checks_abort_pre_upgrade_checks", inputDataValue, executionContext)
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

func (sIface *systemAdministrationLifecycleManagementUpgradePlanClient) ContinueUpgradeContinue(componentTypeParam *string, skipParam *bool) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationLifecycleManagementUpgradePlanContinueUpgradeContinueInputType(), typeConverter)
	sv.AddStructField("ComponentType", componentTypeParam)
	sv.AddStructField("Skip", skipParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationLifecycleManagementUpgradePlanContinueUpgradeContinueRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_lifecycle_management_upgrade_plan", "continue_upgrade_continue", inputDataValue, executionContext)
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

func (sIface *systemAdministrationLifecycleManagementUpgradePlanClient) ExecutePostUpgradeChecksExecutePostUpgradeChecks(componentTypeParam string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationLifecycleManagementUpgradePlanExecutePostUpgradeChecksExecutePostUpgradeChecksInputType(), typeConverter)
	sv.AddStructField("ComponentType", componentTypeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationLifecycleManagementUpgradePlanExecutePostUpgradeChecksExecutePostUpgradeChecksRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_lifecycle_management_upgrade_plan", "execute_post_upgrade_checks_execute_post_upgrade_checks", inputDataValue, executionContext)
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

func (sIface *systemAdministrationLifecycleManagementUpgradePlanClient) ExecutePreUpgradeChecksExecutePreUpgradeChecks(componentTypeParam *string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationLifecycleManagementUpgradePlanExecutePreUpgradeChecksExecutePreUpgradeChecksInputType(), typeConverter)
	sv.AddStructField("ComponentType", componentTypeParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationLifecycleManagementUpgradePlanExecutePreUpgradeChecksExecutePreUpgradeChecksRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_lifecycle_management_upgrade_plan", "execute_pre_upgrade_checks_execute_pre_upgrade_checks", inputDataValue, executionContext)
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

func (sIface *systemAdministrationLifecycleManagementUpgradePlanClient) GetAllPreUpgradeChecksInCsvFormatCsv() (model.UpgradeCheckCsvListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationLifecycleManagementUpgradePlanGetAllPreUpgradeChecksInCsvFormatCsvInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.UpgradeCheckCsvListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationLifecycleManagementUpgradePlanGetAllPreUpgradeChecksInCsvFormatCsvRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_lifecycle_management_upgrade_plan", "get_all_pre_upgrade_checks_in_csv_format_csv", inputDataValue, executionContext)
	var emptyOutput model.UpgradeCheckCsvListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationLifecycleManagementUpgradePlanGetAllPreUpgradeChecksInCsvFormatCsvOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.UpgradeCheckCsvListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationLifecycleManagementUpgradePlanClient) GetPreUpgradeCheckFailures(componentTypeParam *string, cursorParam *string, filterTextParam *string, includedFieldsParam *string, originTypeParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, type_Param *string) (model.UpgradeCheckFailureListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationLifecycleManagementUpgradePlanGetPreUpgradeCheckFailuresInputType(), typeConverter)
	sv.AddStructField("ComponentType", componentTypeParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("FilterText", filterTextParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("OriginType", originTypeParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("Type_", type_Param)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.UpgradeCheckFailureListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationLifecycleManagementUpgradePlanGetPreUpgradeCheckFailuresRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_lifecycle_management_upgrade_plan", "get_pre_upgrade_check_failures", inputDataValue, executionContext)
	var emptyOutput model.UpgradeCheckFailureListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationLifecycleManagementUpgradePlanGetPreUpgradeCheckFailuresOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.UpgradeCheckFailureListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationLifecycleManagementUpgradePlanClient) GetUpgradeChecksInfo(componentTypeParam *string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.ComponentUpgradeChecksInfoListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationLifecycleManagementUpgradePlanGetUpgradeChecksInfoInputType(), typeConverter)
	sv.AddStructField("ComponentType", componentTypeParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ComponentUpgradeChecksInfoListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationLifecycleManagementUpgradePlanGetUpgradeChecksInfoRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_lifecycle_management_upgrade_plan", "get_upgrade_checks_info", inputDataValue, executionContext)
	var emptyOutput model.ComponentUpgradeChecksInfoListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationLifecycleManagementUpgradePlanGetUpgradeChecksInfoOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ComponentUpgradeChecksInfoListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationLifecycleManagementUpgradePlanClient) GetUpgradePlanSettings(componentTypeParam string) (model.UpgradePlanSettings, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationLifecycleManagementUpgradePlanGetUpgradePlanSettingsInputType(), typeConverter)
	sv.AddStructField("ComponentType", componentTypeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.UpgradePlanSettings
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationLifecycleManagementUpgradePlanGetUpgradePlanSettingsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_lifecycle_management_upgrade_plan", "get_upgrade_plan_settings", inputDataValue, executionContext)
	var emptyOutput model.UpgradePlanSettings
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationLifecycleManagementUpgradePlanGetUpgradePlanSettingsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.UpgradePlanSettings), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationLifecycleManagementUpgradePlanClient) PauseUpgradePause() error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationLifecycleManagementUpgradePlanPauseUpgradePauseInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationLifecycleManagementUpgradePlanPauseUpgradePauseRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_lifecycle_management_upgrade_plan", "pause_upgrade_pause", inputDataValue, executionContext)
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

func (sIface *systemAdministrationLifecycleManagementUpgradePlanClient) ResetUpgradePlanReset(componentTypeParam string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationLifecycleManagementUpgradePlanResetUpgradePlanResetInputType(), typeConverter)
	sv.AddStructField("ComponentType", componentTypeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationLifecycleManagementUpgradePlanResetUpgradePlanResetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_lifecycle_management_upgrade_plan", "reset_upgrade_plan_reset", inputDataValue, executionContext)
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

func (sIface *systemAdministrationLifecycleManagementUpgradePlanClient) StageUpgradeStageUpgrade(componentTypeParam *string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationLifecycleManagementUpgradePlanStageUpgradeStageUpgradeInputType(), typeConverter)
	sv.AddStructField("ComponentType", componentTypeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationLifecycleManagementUpgradePlanStageUpgradeStageUpgradeRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_lifecycle_management_upgrade_plan", "stage_upgrade_stage_upgrade", inputDataValue, executionContext)
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

func (sIface *systemAdministrationLifecycleManagementUpgradePlanClient) StartUpgradeStart() error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationLifecycleManagementUpgradePlanStartUpgradeStartInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationLifecycleManagementUpgradePlanStartUpgradeStartRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_lifecycle_management_upgrade_plan", "start_upgrade_start", inputDataValue, executionContext)
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

func (sIface *systemAdministrationLifecycleManagementUpgradePlanClient) UpdateUpgradePlanSettings(componentTypeParam string, upgradePlanSettingsParam model.UpgradePlanSettings) (model.UpgradePlanSettings, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationLifecycleManagementUpgradePlanUpdateUpgradePlanSettingsInputType(), typeConverter)
	sv.AddStructField("ComponentType", componentTypeParam)
	sv.AddStructField("UpgradePlanSettings", upgradePlanSettingsParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.UpgradePlanSettings
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationLifecycleManagementUpgradePlanUpdateUpgradePlanSettingsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_lifecycle_management_upgrade_plan", "update_upgrade_plan_settings", inputDataValue, executionContext)
	var emptyOutput model.UpgradePlanSettings
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationLifecycleManagementUpgradePlanUpdateUpgradePlanSettingsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.UpgradePlanSettings), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationLifecycleManagementUpgradePlanClient) UpgradeSelectedUnitsUpgradeSelectedUnits(upgradeUnitListParam model.UpgradeUnitList) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationLifecycleManagementUpgradePlanUpgradeSelectedUnitsUpgradeSelectedUnitsInputType(), typeConverter)
	sv.AddStructField("UpgradeUnitList", upgradeUnitListParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationLifecycleManagementUpgradePlanUpgradeSelectedUnitsUpgradeSelectedUnitsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_lifecycle_management_upgrade_plan", "upgrade_selected_units_upgrade_selected_units", inputDataValue, executionContext)
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
