/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Job.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package backup

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
	"time"
)



// The ``ReturnStatus`` enumeration class defines the return type for the cancel operation. You specify the return status when you return the result of cancel job. See JobReturnResult.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type JobReturnStatus string

const (
    // Cancel operation failed.
	JobReturnStatus_FAIL JobReturnStatus = "FAIL"
    // Cancel operation passed with warnings.
	JobReturnStatus_WARNING JobReturnStatus = "WARNING"
    // Cancel operation succeeded.
	JobReturnStatus_OK JobReturnStatus = "OK"
)

func (r JobReturnStatus) JobReturnStatus() bool {
	switch r {
	case JobReturnStatus_FAIL:
		return true
	case JobReturnStatus_WARNING:
		return true
	case JobReturnStatus_OK:
		return true
	default:
		return false
	}
}


// The ``LocationType`` enumeration class defines the type of destination location for backup/restore. You specify the location type when you create a backup job. See JobBackupRequest.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type JobLocationType string

const (
    // Destination is FTP server.
	JobLocationType_FTP JobLocationType = "FTP"
    // Destination is HTTP server.
	JobLocationType_HTTP JobLocationType = "HTTP"
    // Destination is FTPS server.
	JobLocationType_FTPS JobLocationType = "FTPS"
    // Destination is HTTPS server.
	JobLocationType_HTTPS JobLocationType = "HTTPS"
    // Destination is SSH server.
	JobLocationType_SCP JobLocationType = "SCP"
    // Destination is SFTP server.
	JobLocationType_SFTP JobLocationType = "SFTP"
    // Destination is NFS server. This constant field was added in vSphere API 6.7.2.
	JobLocationType_NFS JobLocationType = "NFS"
    // Destination is SMB server. This constant field was added in vSphere API 6.7.2.
	JobLocationType_SMB JobLocationType = "SMB"
)

func (l JobLocationType) JobLocationType() bool {
	switch l {
	case JobLocationType_FTP:
		return true
	case JobLocationType_HTTP:
		return true
	case JobLocationType_FTPS:
		return true
	case JobLocationType_HTTPS:
		return true
	case JobLocationType_SCP:
		return true
	case JobLocationType_SFTP:
		return true
	case JobLocationType_NFS:
		return true
	case JobLocationType_SMB:
		return true
	default:
		return false
	}
}


// The ``BackupRestoreProcessState`` enumeration class defines the possible states of a backup/restore process.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type JobBackupRestoreProcessState string

const (
    // Backup/Restore job failed.
	JobBackupRestoreProcessState_FAILED JobBackupRestoreProcessState = "FAILED"
    // Backup/Restore job is in progress.
	JobBackupRestoreProcessState_INPROGRESS JobBackupRestoreProcessState = "INPROGRESS"
    // Backup/Restore job is not started.
	JobBackupRestoreProcessState_NONE JobBackupRestoreProcessState = "NONE"
    // Backup/Restore job completed successfully.
	JobBackupRestoreProcessState_SUCCEEDED JobBackupRestoreProcessState = "SUCCEEDED"
)

func (b JobBackupRestoreProcessState) JobBackupRestoreProcessState() bool {
	switch b {
	case JobBackupRestoreProcessState_FAILED:
		return true
	case JobBackupRestoreProcessState_INPROGRESS:
		return true
	case JobBackupRestoreProcessState_NONE:
		return true
	case JobBackupRestoreProcessState_SUCCEEDED:
		return true
	default:
		return false
	}
}


// The ``LocalizableMessage`` class represents a localizable message.
type JobLocalizableMessage struct {
    // Id in message bundle.
	Id string
    // Text in english.
	DefaultMessage string
    // Nested data.
	Args []string
}

// The ``ReturnResult`` class contains the result information for the cancel operation.
type JobReturnResult struct {
    // Status of the cancel operation.
	Status JobReturnStatus
    // List of messages.
	Messages []JobLocalizableMessage
}

// The ``BackupRequest`` class represents a requested backup piece.
type JobBackupRequest struct {
    // List of optional parts that will be included in the backup. Use the Parts#list method to get information about the supported parts.
	Parts []string
    // Password for a backup piece. The backupPassword must adhere to the following password requirements: At least 8 characters, cannot be more than 20 characters in length. At least 1 uppercase letter. At least 1 lowercase letter. At least 1 numeric digit. At least 1 special character (i.e. any character not in [0-9,a-z,A-Z]). Only visible ASCII characters (for example, no space).
	BackupPassword *string
    // Type of backup location.
	LocationType JobLocationType
    // Path or URL of the backup location.
	Location string
    // Username for the given location.
	LocationUser *string
    // Password for the given location.
	LocationPassword *string
    // Custom comment provided by the user.
	Comment *string
}

// The ``BackupJobStatus`` class represents the status of a backup/restore job.
type JobBackupJobStatus struct {
    // TimeStamp based ID.
	Id string
    // The state of the backup job.
	State JobBackupRestoreProcessState
    // List of messages.
	Messages []JobLocalizableMessage
    // Progress of the job in percentage.
	Progress int64
    // Time when the backup was started.
	StartTime time.Time
    // Time when the backup was finished.
	EndTime *time.Time
}



func jobCancelInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = bindings.NewIdType([]string{"com.vmware.appliance.recovery.backup.job"}, "")
	fieldNameMap["id"] = "Id"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func jobCancelOutputType() bindings.BindingType {
	return bindings.NewReferenceType(JobReturnResultBindingType)
}

func jobCancelRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["id"] = bindings.NewIdType([]string{"com.vmware.appliance.recovery.backup.job"}, "")
	fieldNameMap["id"] = "Id"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"NotFound": 404,"Error": 500})
}

func jobCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["piece"] = bindings.NewReferenceType(JobBackupRequestBindingType)
	fieldNameMap["piece"] = "Piece"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func jobCreateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(JobBackupJobStatusBindingType)
}

func jobCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["piece"] = bindings.NewReferenceType(JobBackupRequestBindingType)
	fieldNameMap["piece"] = "Piece"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"FeatureInUse": 400,"Error": 500})
}

func jobListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func jobListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewIdType([]string{"com.vmware.appliance.recovery.backup.job"}, ""), reflect.TypeOf([]string{}))
}

func jobListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"Error": 500})
}

func jobGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = bindings.NewIdType([]string{"com.vmware.appliance.recovery.backup.job"}, "")
	fieldNameMap["id"] = "Id"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func jobGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(JobBackupJobStatusBindingType)
}

func jobGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["id"] = bindings.NewIdType([]string{"com.vmware.appliance.recovery.backup.job"}, "")
	fieldNameMap["id"] = "Id"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"NotFound": 404,"Error": 500})
}


func JobLocalizableMessageBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = bindings.NewStringType()
	fieldNameMap["id"] = "Id"
	fields["default_message"] = bindings.NewStringType()
	fieldNameMap["default_message"] = "DefaultMessage"
	fields["args"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["args"] = "Args"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.recovery.backup.job.localizable_message", fields, reflect.TypeOf(JobLocalizableMessage{}), fieldNameMap, validators)
}

func JobReturnResultBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["status"] = bindings.NewEnumType("com.vmware.appliance.recovery.backup.job.return_status", reflect.TypeOf(JobReturnStatus(JobReturnStatus_FAIL)))
	fieldNameMap["status"] = "Status"
	fields["messages"] = bindings.NewListType(bindings.NewReferenceType(JobLocalizableMessageBindingType), reflect.TypeOf([]JobLocalizableMessage{}))
	fieldNameMap["messages"] = "Messages"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.recovery.backup.job.return_result", fields, reflect.TypeOf(JobReturnResult{}), fieldNameMap, validators)
}

func JobBackupRequestBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["parts"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["parts"] = "Parts"
	fields["backup_password"] = bindings.NewOptionalType(bindings.NewSecretType())
	fieldNameMap["backup_password"] = "BackupPassword"
	fields["location_type"] = bindings.NewEnumType("com.vmware.appliance.recovery.backup.job.location_type", reflect.TypeOf(JobLocationType(JobLocationType_FTP)))
	fieldNameMap["location_type"] = "LocationType"
	fields["location"] = bindings.NewStringType()
	fieldNameMap["location"] = "Location"
	fields["location_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["location_user"] = "LocationUser"
	fields["location_password"] = bindings.NewOptionalType(bindings.NewSecretType())
	fieldNameMap["location_password"] = "LocationPassword"
	fields["comment"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["comment"] = "Comment"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.recovery.backup.job.backup_request", fields, reflect.TypeOf(JobBackupRequest{}), fieldNameMap, validators)
}

func JobBackupJobStatusBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = bindings.NewStringType()
	fieldNameMap["id"] = "Id"
	fields["state"] = bindings.NewEnumType("com.vmware.appliance.recovery.backup.job.backup_restore_process_state", reflect.TypeOf(JobBackupRestoreProcessState(JobBackupRestoreProcessState_FAILED)))
	fieldNameMap["state"] = "State"
	fields["messages"] = bindings.NewListType(bindings.NewReferenceType(JobLocalizableMessageBindingType), reflect.TypeOf([]JobLocalizableMessage{}))
	fieldNameMap["messages"] = "Messages"
	fields["progress"] = bindings.NewIntegerType()
	fieldNameMap["progress"] = "Progress"
	fields["start_time"] = bindings.NewDateTimeType()
	fieldNameMap["start_time"] = "StartTime"
	fields["end_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
	fieldNameMap["end_time"] = "EndTime"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.recovery.backup.job.backup_job_status", fields, reflect.TypeOf(JobBackupJobStatus{}), fieldNameMap, validators)
}

