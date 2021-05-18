// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: SystemAdministrationLifecycleManagementBackupRestoreManagementBackup
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

type SystemAdministrationLifecycleManagementBackupRestoreManagementBackupClient interface {

	// Configure file server and timers for automated backup. If secret fields are omitted (password, passphrase) then use the previously set value.
	//
	// @param backupConfigurationParam (required)
	// @param frameTypeParam Frame type (optional, default to LOCAL_LOCAL_MANAGER)
	// @param siteIdParam Site ID (optional, default to localhost)
	// @return com.vmware.model.BackupConfiguration
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ConfigureBackupConfig(backupConfigurationParam model.BackupConfiguration, frameTypeParam *string, siteIdParam *string) (model.BackupConfiguration, error)

	// Get a configuration of a file server and timers for automated backup. Fields that contain secrets (password, passphrase) are not returned.
	// @return com.vmware.model.BackupConfiguration
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetBackupConfig() (model.BackupConfiguration, error)

	// Get history of previous backup operations
	// @return com.vmware.model.BackupOperationHistory
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetBackupHistory() (model.BackupOperationHistory, error)

	// Get a configuration of a file server, timers for automated backup, latest backup status, backups list for a site. Fields that contain secrets (password, passphrase) are not returned.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param frameTypeParam Frame type (optional, default to LOCAL_LOCAL_MANAGER)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param showBackupsListParam Need a list of backups (optional, default to true)
	// @param siteIdParam UUID of the site (optional, default to localhost)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.model.BackupOverview
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetBackupOverview(cursorParam *string, frameTypeParam *string, includedFieldsParam *string, pageSizeParam *int64, showBackupsListParam *bool, siteIdParam *string, sortAscendingParam *bool, sortByParam *string) (model.BackupOverview, error)

	// Get status of active backup operations
	// @return com.vmware.model.CurrentBackupOperationStatus
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetBackupStatus() (model.CurrentBackupOperationStatus, error)

	// Get SHA256 fingerprint of ECDSA key of remote server. The caller should independently verify that the key is trusted.
	//
	// @param remoteServerFingerprintRequestParam (required)
	// @return com.vmware.model.RemoteServerFingerprint
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetSshFingerprintOfServerRetrieveSshFingerprint(remoteServerFingerprintRequestParam model.RemoteServerFingerprintRequest) (model.RemoteServerFingerprint, error)

	// Request one-time backup. The backup will be uploaded using the same server configuration as for automatic backup.
	//
	// @param frameTypeParam Frame type (optional, default to LOCAL_LOCAL_MANAGER)
	// @param siteIdParam Site ID (optional, default to localhost)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	RequestOnetimeBackupBackupToRemote(frameTypeParam *string, siteIdParam *string) error

	// Request one-time inventory summary. The backup will be uploaded using the same server configuration as for an automatic backup.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	RequestOnetimeInventorySummarySummarizeInventoryToRemote() error
}

type systemAdministrationLifecycleManagementBackupRestoreManagementBackupClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSystemAdministrationLifecycleManagementBackupRestoreManagementBackupClient(connector client.Connector) *systemAdministrationLifecycleManagementBackupRestoreManagementBackupClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.system_administration_lifecycle_management_backup_restore_management_backup")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"configure_backup_config": core.NewMethodIdentifier(interfaceIdentifier, "configure_backup_config"),
		"get_backup_config":       core.NewMethodIdentifier(interfaceIdentifier, "get_backup_config"),
		"get_backup_history":      core.NewMethodIdentifier(interfaceIdentifier, "get_backup_history"),
		"get_backup_overview":     core.NewMethodIdentifier(interfaceIdentifier, "get_backup_overview"),
		"get_backup_status":       core.NewMethodIdentifier(interfaceIdentifier, "get_backup_status"),
		"get_ssh_fingerprint_of_server_retrieve_ssh_fingerprint":          core.NewMethodIdentifier(interfaceIdentifier, "get_ssh_fingerprint_of_server_retrieve_ssh_fingerprint"),
		"request_onetime_backup_backup_to_remote":                         core.NewMethodIdentifier(interfaceIdentifier, "request_onetime_backup_backup_to_remote"),
		"request_onetime_inventory_summary_summarize_inventory_to_remote": core.NewMethodIdentifier(interfaceIdentifier, "request_onetime_inventory_summary_summarize_inventory_to_remote"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := systemAdministrationLifecycleManagementBackupRestoreManagementBackupClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *systemAdministrationLifecycleManagementBackupRestoreManagementBackupClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *systemAdministrationLifecycleManagementBackupRestoreManagementBackupClient) ConfigureBackupConfig(backupConfigurationParam model.BackupConfiguration, frameTypeParam *string, siteIdParam *string) (model.BackupConfiguration, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationLifecycleManagementBackupRestoreManagementBackupConfigureBackupConfigInputType(), typeConverter)
	sv.AddStructField("BackupConfiguration", backupConfigurationParam)
	sv.AddStructField("FrameType", frameTypeParam)
	sv.AddStructField("SiteId", siteIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.BackupConfiguration
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationLifecycleManagementBackupRestoreManagementBackupConfigureBackupConfigRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_lifecycle_management_backup_restore_management_backup", "configure_backup_config", inputDataValue, executionContext)
	var emptyOutput model.BackupConfiguration
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationLifecycleManagementBackupRestoreManagementBackupConfigureBackupConfigOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BackupConfiguration), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationLifecycleManagementBackupRestoreManagementBackupClient) GetBackupConfig() (model.BackupConfiguration, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationLifecycleManagementBackupRestoreManagementBackupGetBackupConfigInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.BackupConfiguration
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationLifecycleManagementBackupRestoreManagementBackupGetBackupConfigRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_lifecycle_management_backup_restore_management_backup", "get_backup_config", inputDataValue, executionContext)
	var emptyOutput model.BackupConfiguration
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationLifecycleManagementBackupRestoreManagementBackupGetBackupConfigOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BackupConfiguration), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationLifecycleManagementBackupRestoreManagementBackupClient) GetBackupHistory() (model.BackupOperationHistory, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationLifecycleManagementBackupRestoreManagementBackupGetBackupHistoryInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.BackupOperationHistory
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationLifecycleManagementBackupRestoreManagementBackupGetBackupHistoryRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_lifecycle_management_backup_restore_management_backup", "get_backup_history", inputDataValue, executionContext)
	var emptyOutput model.BackupOperationHistory
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationLifecycleManagementBackupRestoreManagementBackupGetBackupHistoryOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BackupOperationHistory), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationLifecycleManagementBackupRestoreManagementBackupClient) GetBackupOverview(cursorParam *string, frameTypeParam *string, includedFieldsParam *string, pageSizeParam *int64, showBackupsListParam *bool, siteIdParam *string, sortAscendingParam *bool, sortByParam *string) (model.BackupOverview, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationLifecycleManagementBackupRestoreManagementBackupGetBackupOverviewInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("FrameType", frameTypeParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("ShowBackupsList", showBackupsListParam)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.BackupOverview
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationLifecycleManagementBackupRestoreManagementBackupGetBackupOverviewRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_lifecycle_management_backup_restore_management_backup", "get_backup_overview", inputDataValue, executionContext)
	var emptyOutput model.BackupOverview
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationLifecycleManagementBackupRestoreManagementBackupGetBackupOverviewOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BackupOverview), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationLifecycleManagementBackupRestoreManagementBackupClient) GetBackupStatus() (model.CurrentBackupOperationStatus, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationLifecycleManagementBackupRestoreManagementBackupGetBackupStatusInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.CurrentBackupOperationStatus
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationLifecycleManagementBackupRestoreManagementBackupGetBackupStatusRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_lifecycle_management_backup_restore_management_backup", "get_backup_status", inputDataValue, executionContext)
	var emptyOutput model.CurrentBackupOperationStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationLifecycleManagementBackupRestoreManagementBackupGetBackupStatusOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.CurrentBackupOperationStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationLifecycleManagementBackupRestoreManagementBackupClient) GetSshFingerprintOfServerRetrieveSshFingerprint(remoteServerFingerprintRequestParam model.RemoteServerFingerprintRequest) (model.RemoteServerFingerprint, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationLifecycleManagementBackupRestoreManagementBackupGetSshFingerprintOfServerRetrieveSshFingerprintInputType(), typeConverter)
	sv.AddStructField("RemoteServerFingerprintRequest", remoteServerFingerprintRequestParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.RemoteServerFingerprint
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationLifecycleManagementBackupRestoreManagementBackupGetSshFingerprintOfServerRetrieveSshFingerprintRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_lifecycle_management_backup_restore_management_backup", "get_ssh_fingerprint_of_server_retrieve_ssh_fingerprint", inputDataValue, executionContext)
	var emptyOutput model.RemoteServerFingerprint
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationLifecycleManagementBackupRestoreManagementBackupGetSshFingerprintOfServerRetrieveSshFingerprintOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.RemoteServerFingerprint), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationLifecycleManagementBackupRestoreManagementBackupClient) RequestOnetimeBackupBackupToRemote(frameTypeParam *string, siteIdParam *string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationLifecycleManagementBackupRestoreManagementBackupRequestOnetimeBackupBackupToRemoteInputType(), typeConverter)
	sv.AddStructField("FrameType", frameTypeParam)
	sv.AddStructField("SiteId", siteIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationLifecycleManagementBackupRestoreManagementBackupRequestOnetimeBackupBackupToRemoteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_lifecycle_management_backup_restore_management_backup", "request_onetime_backup_backup_to_remote", inputDataValue, executionContext)
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

func (sIface *systemAdministrationLifecycleManagementBackupRestoreManagementBackupClient) RequestOnetimeInventorySummarySummarizeInventoryToRemote() error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationLifecycleManagementBackupRestoreManagementBackupRequestOnetimeInventorySummarySummarizeInventoryToRemoteInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationLifecycleManagementBackupRestoreManagementBackupRequestOnetimeInventorySummarySummarizeInventoryToRemoteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_lifecycle_management_backup_restore_management_backup", "request_onetime_inventory_summary_summarize_inventory_to_remote", inputDataValue, executionContext)
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
