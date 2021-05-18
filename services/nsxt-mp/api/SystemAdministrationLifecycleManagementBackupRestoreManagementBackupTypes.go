// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: SystemAdministrationLifecycleManagementBackupRestoreManagementBackup.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package api

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-mp/model"
	"reflect"
)

// Possible value for ``frameType`` of method SystemAdministrationLifecycleManagementBackupRestoreManagementBackup#configureBackupConfig.
const SystemAdministrationLifecycleManagementBackupRestoreManagementBackup_CONFIGURE_BACKUP_CONFIG_FRAME_TYPE_GLOBAL_MANAGER = "GLOBAL_MANAGER"

// Possible value for ``frameType`` of method SystemAdministrationLifecycleManagementBackupRestoreManagementBackup#configureBackupConfig.
const SystemAdministrationLifecycleManagementBackupRestoreManagementBackup_CONFIGURE_BACKUP_CONFIG_FRAME_TYPE_LOCAL_MANAGER = "LOCAL_MANAGER"

// Possible value for ``frameType`` of method SystemAdministrationLifecycleManagementBackupRestoreManagementBackup#configureBackupConfig.
const SystemAdministrationLifecycleManagementBackupRestoreManagementBackup_CONFIGURE_BACKUP_CONFIG_FRAME_TYPE_LOCAL_LOCAL_MANAGER = "LOCAL_LOCAL_MANAGER"

// Possible value for ``frameType`` of method SystemAdministrationLifecycleManagementBackupRestoreManagementBackup#configureBackupConfig.
const SystemAdministrationLifecycleManagementBackupRestoreManagementBackup_CONFIGURE_BACKUP_CONFIG_FRAME_TYPE_NSX_INTELLIGENCE = "NSX_INTELLIGENCE"

// Possible value for ``frameType`` of method SystemAdministrationLifecycleManagementBackupRestoreManagementBackup#getBackupOverview.
const SystemAdministrationLifecycleManagementBackupRestoreManagementBackup_GET_BACKUP_OVERVIEW_FRAME_TYPE_GLOBAL_MANAGER = "GLOBAL_MANAGER"

// Possible value for ``frameType`` of method SystemAdministrationLifecycleManagementBackupRestoreManagementBackup#getBackupOverview.
const SystemAdministrationLifecycleManagementBackupRestoreManagementBackup_GET_BACKUP_OVERVIEW_FRAME_TYPE_LOCAL_MANAGER = "LOCAL_MANAGER"

// Possible value for ``frameType`` of method SystemAdministrationLifecycleManagementBackupRestoreManagementBackup#getBackupOverview.
const SystemAdministrationLifecycleManagementBackupRestoreManagementBackup_GET_BACKUP_OVERVIEW_FRAME_TYPE_LOCAL_LOCAL_MANAGER = "LOCAL_LOCAL_MANAGER"

// Possible value for ``frameType`` of method SystemAdministrationLifecycleManagementBackupRestoreManagementBackup#getBackupOverview.
const SystemAdministrationLifecycleManagementBackupRestoreManagementBackup_GET_BACKUP_OVERVIEW_FRAME_TYPE_NSX_INTELLIGENCE = "NSX_INTELLIGENCE"

// Possible value for ``frameType`` of method SystemAdministrationLifecycleManagementBackupRestoreManagementBackup#requestOnetimeBackupBackupToRemote.
const SystemAdministrationLifecycleManagementBackupRestoreManagementBackup_REQUEST_ONETIME_BACKUP_BACKUP_TO_REMOTE_FRAME_TYPE_GLOBAL_MANAGER = "GLOBAL_MANAGER"

// Possible value for ``frameType`` of method SystemAdministrationLifecycleManagementBackupRestoreManagementBackup#requestOnetimeBackupBackupToRemote.
const SystemAdministrationLifecycleManagementBackupRestoreManagementBackup_REQUEST_ONETIME_BACKUP_BACKUP_TO_REMOTE_FRAME_TYPE_LOCAL_MANAGER = "LOCAL_MANAGER"

// Possible value for ``frameType`` of method SystemAdministrationLifecycleManagementBackupRestoreManagementBackup#requestOnetimeBackupBackupToRemote.
const SystemAdministrationLifecycleManagementBackupRestoreManagementBackup_REQUEST_ONETIME_BACKUP_BACKUP_TO_REMOTE_FRAME_TYPE_LOCAL_LOCAL_MANAGER = "LOCAL_LOCAL_MANAGER"

// Possible value for ``frameType`` of method SystemAdministrationLifecycleManagementBackupRestoreManagementBackup#requestOnetimeBackupBackupToRemote.
const SystemAdministrationLifecycleManagementBackupRestoreManagementBackup_REQUEST_ONETIME_BACKUP_BACKUP_TO_REMOTE_FRAME_TYPE_NSX_INTELLIGENCE = "NSX_INTELLIGENCE"

func systemAdministrationLifecycleManagementBackupRestoreManagementBackupConfigureBackupConfigInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["backup_configuration"] = bindings.NewReferenceType(model.BackupConfigurationBindingType)
	fields["frame_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["site_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["backup_configuration"] = "BackupConfiguration"
	fieldNameMap["frame_type"] = "FrameType"
	fieldNameMap["site_id"] = "SiteId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationLifecycleManagementBackupRestoreManagementBackupConfigureBackupConfigOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.BackupConfigurationBindingType)
}

func systemAdministrationLifecycleManagementBackupRestoreManagementBackupConfigureBackupConfigRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["backup_configuration"] = bindings.NewReferenceType(model.BackupConfigurationBindingType)
	fields["frame_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["site_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["backup_configuration"] = "BackupConfiguration"
	fieldNameMap["frame_type"] = "FrameType"
	fieldNameMap["site_id"] = "SiteId"
	paramsTypeMap["frame_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["site_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["backup_configuration"] = bindings.NewReferenceType(model.BackupConfigurationBindingType)
	queryParams["site_id"] = "site_id"
	queryParams["frame_type"] = "frame_type"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"backup_configuration",
		"PUT",
		"/api/v1/cluster/backups/config",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationLifecycleManagementBackupRestoreManagementBackupGetBackupConfigInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationLifecycleManagementBackupRestoreManagementBackupGetBackupConfigOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.BackupConfigurationBindingType)
}

func systemAdministrationLifecycleManagementBackupRestoreManagementBackupGetBackupConfigRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/api/v1/cluster/backups/config",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationLifecycleManagementBackupRestoreManagementBackupGetBackupHistoryInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationLifecycleManagementBackupRestoreManagementBackupGetBackupHistoryOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.BackupOperationHistoryBindingType)
}

func systemAdministrationLifecycleManagementBackupRestoreManagementBackupGetBackupHistoryRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/api/v1/cluster/backups/history",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationLifecycleManagementBackupRestoreManagementBackupGetBackupOverviewInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["frame_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["show_backups_list"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["site_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["frame_type"] = "FrameType"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["show_backups_list"] = "ShowBackupsList"
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationLifecycleManagementBackupRestoreManagementBackupGetBackupOverviewOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.BackupOverviewBindingType)
}

func systemAdministrationLifecycleManagementBackupRestoreManagementBackupGetBackupOverviewRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["frame_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["show_backups_list"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["site_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["frame_type"] = "FrameType"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["show_backups_list"] = "ShowBackupsList"
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["frame_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["site_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["show_backups_list"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["site_id"] = "site_id"
	queryParams["show_backups_list"] = "show_backups_list"
	queryParams["sort_by"] = "sort_by"
	queryParams["frame_type"] = "frame_type"
	queryParams["page_size"] = "page_size"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/api/v1/cluster/backups/overview",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationLifecycleManagementBackupRestoreManagementBackupGetBackupStatusInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationLifecycleManagementBackupRestoreManagementBackupGetBackupStatusOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.CurrentBackupOperationStatusBindingType)
}

func systemAdministrationLifecycleManagementBackupRestoreManagementBackupGetBackupStatusRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/api/v1/cluster/backups/status",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationLifecycleManagementBackupRestoreManagementBackupGetSshFingerprintOfServerRetrieveSshFingerprintInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["remote_server_fingerprint_request"] = bindings.NewReferenceType(model.RemoteServerFingerprintRequestBindingType)
	fieldNameMap["remote_server_fingerprint_request"] = "RemoteServerFingerprintRequest"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationLifecycleManagementBackupRestoreManagementBackupGetSshFingerprintOfServerRetrieveSshFingerprintOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.RemoteServerFingerprintBindingType)
}

func systemAdministrationLifecycleManagementBackupRestoreManagementBackupGetSshFingerprintOfServerRetrieveSshFingerprintRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["remote_server_fingerprint_request"] = bindings.NewReferenceType(model.RemoteServerFingerprintRequestBindingType)
	fieldNameMap["remote_server_fingerprint_request"] = "RemoteServerFingerprintRequest"
	paramsTypeMap["remote_server_fingerprint_request"] = bindings.NewReferenceType(model.RemoteServerFingerprintRequestBindingType)
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"action=retrieve_ssh_fingerprint",
		"remote_server_fingerprint_request",
		"POST",
		"/api/v1/cluster/backups",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationLifecycleManagementBackupRestoreManagementBackupRequestOnetimeBackupBackupToRemoteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["frame_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["site_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["frame_type"] = "FrameType"
	fieldNameMap["site_id"] = "SiteId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationLifecycleManagementBackupRestoreManagementBackupRequestOnetimeBackupBackupToRemoteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func systemAdministrationLifecycleManagementBackupRestoreManagementBackupRequestOnetimeBackupBackupToRemoteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["frame_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["site_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["frame_type"] = "FrameType"
	fieldNameMap["site_id"] = "SiteId"
	paramsTypeMap["frame_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["site_id"] = bindings.NewOptionalType(bindings.NewStringType())
	queryParams["site_id"] = "site_id"
	queryParams["frame_type"] = "frame_type"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"action=backup_to_remote",
		"",
		"POST",
		"/api/v1/cluster",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationLifecycleManagementBackupRestoreManagementBackupRequestOnetimeInventorySummarySummarizeInventoryToRemoteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationLifecycleManagementBackupRestoreManagementBackupRequestOnetimeInventorySummarySummarizeInventoryToRemoteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func systemAdministrationLifecycleManagementBackupRestoreManagementBackupRequestOnetimeInventorySummarySummarizeInventoryToRemoteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"action=summarize_inventory_to_remote",
		"",
		"POST",
		"/api/v1/cluster",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
