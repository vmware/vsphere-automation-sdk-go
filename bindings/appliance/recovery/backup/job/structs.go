/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Job.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package job

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
    "time"
)



// The ``ReturnStatus`` enumeration class defines the return type for the cancel operation. You specify the return status when you return the result of cancel job. See ReturnResult.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type ReturnStatus string

const (
    // Cancel operation failed.
     ReturnStatus_FAIL ReturnStatus = "FAIL"
    // Cancel operation passed with warnings.
     ReturnStatus_WARNING ReturnStatus = "WARNING"
    // Cancel operation succeeded.
     ReturnStatus_OK ReturnStatus = "OK"
)

func (r ReturnStatus) ReturnStatus() bool {
    switch r {
        case ReturnStatus_FAIL:
            return true
        case ReturnStatus_WARNING:
            return true
        case ReturnStatus_OK:
            return true
        default:
            return false
    }
}




// The ``LocationType`` enumeration class defines the type of destination location for backup/restore. You specify the location type when you create a backup job. See BackupRequest.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type LocationType string

const (
    // Destination is FTP server.
     LocationType_FTP LocationType = "FTP"
    // Destination is HTTP server.
     LocationType_HTTP LocationType = "HTTP"
    // Destination is FTPS server.
     LocationType_FTPS LocationType = "FTPS"
    // Destination is HTTPS server.
     LocationType_HTTPS LocationType = "HTTPS"
    // Destination is SSH server.
     LocationType_SCP LocationType = "SCP"
    // Destination is SFTP server.
     LocationType_SFTP LocationType = "SFTP"
    // Destination is NFS server.
     LocationType_NFS LocationType = "NFS"
    // Destination is SMB server.
     LocationType_SMB LocationType = "SMB"
)

func (l LocationType) LocationType() bool {
    switch l {
        case LocationType_FTP:
            return true
        case LocationType_HTTP:
            return true
        case LocationType_FTPS:
            return true
        case LocationType_HTTPS:
            return true
        case LocationType_SCP:
            return true
        case LocationType_SFTP:
            return true
        case LocationType_NFS:
            return true
        case LocationType_SMB:
            return true
        default:
            return false
    }
}




// The ``BackupRestoreProcessState`` enumeration class defines the possible states of a backup/restore process.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type BackupRestoreProcessState string

const (
    // Backup/Restore job failed.
     BackupRestoreProcessState_FAILED BackupRestoreProcessState = "FAILED"
    // Backup/Restore job is in progress.
     BackupRestoreProcessState_INPROGRESS BackupRestoreProcessState = "INPROGRESS"
    // Backup/Restore job is not started.
     BackupRestoreProcessState_NONE BackupRestoreProcessState = "NONE"
    // Backup/Restore job completed successfully.
     BackupRestoreProcessState_SUCCEEDED BackupRestoreProcessState = "SUCCEEDED"
)

func (b BackupRestoreProcessState) BackupRestoreProcessState() bool {
    switch b {
        case BackupRestoreProcessState_FAILED:
            return true
        case BackupRestoreProcessState_INPROGRESS:
            return true
        case BackupRestoreProcessState_NONE:
            return true
        case BackupRestoreProcessState_SUCCEEDED:
            return true
        default:
            return false
    }
}





// The ``LocalizableMessage`` class represents a localizable message.
 type LocalizableMessage struct {
    Id string
    DefaultMessage string
    Args []string
}






// The ``ReturnResult`` class contains the result information for the cancel operation.
 type ReturnResult struct {
    Status ReturnStatus
    Messages []LocalizableMessage
}






// The ``BackupRequest`` class represents a requested backup piece.
 type BackupRequest struct {
    Parts []string
    BackupPassword *string
    LocationType LocationType
    Location string
    LocationUser *string
    LocationPassword *string
    Comment *string
}






// The ``BackupJobStatus`` class represents the status of a backup/restore job.
 type BackupJobStatus struct {
    Id string
    State BackupRestoreProcessState
    Messages []LocalizableMessage
    Progress int64
    StartTime time.Time
    EndTime *time.Time
}









func cancelInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.appliance.recovery.backup.job"}, "")
    fieldNameMap["id"] = "Id"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cancelOutputType() bindings.BindingType {
    return bindings.NewReferenceType(ReturnResultBindingType)
}

func cancelRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"NotFound": 404,"Error": 500})
}


func createInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["piece"] = bindings.NewReferenceType(BackupRequestBindingType)
    fieldNameMap["piece"] = "Piece"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func createOutputType() bindings.BindingType {
    return bindings.NewReferenceType(BackupJobStatusBindingType)
}

func createRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"FeatureInUse": 400,"Error": 500})
}


func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewIdType([]string {"com.vmware.appliance.recovery.backup.job"}, ""), reflect.TypeOf([]string{}))
}

func listRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.appliance.recovery.backup.job"}, "")
    fieldNameMap["id"] = "Id"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewReferenceType(BackupJobStatusBindingType)
}

func getRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"NotFound": 404,"Error": 500})
}



func LocalizableMessageBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewStringType()
    fieldNameMap["id"] = "Id"
    fields["default_message"] = bindings.NewStringType()
    fieldNameMap["default_message"] = "DefaultMessage"
    fields["args"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["args"] = "Args"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.recovery.backup.job.localizable_message",fields, reflect.TypeOf(LocalizableMessage{}), fieldNameMap, validators)
}

func ReturnResultBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["status"] = bindings.NewEnumType("com.vmware.appliance.recovery.backup.job.return_status", reflect.TypeOf(ReturnStatus(ReturnStatus_FAIL)))
    fieldNameMap["status"] = "Status"
    fields["messages"] = bindings.NewListType(bindings.NewReferenceType(LocalizableMessageBindingType), reflect.TypeOf([]LocalizableMessage{}))
    fieldNameMap["messages"] = "Messages"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.recovery.backup.job.return_result",fields, reflect.TypeOf(ReturnResult{}), fieldNameMap, validators)
}

func BackupRequestBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["parts"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["parts"] = "Parts"
    fields["backup_password"] = bindings.NewOptionalType(bindings.NewSecretType())
    fieldNameMap["backup_password"] = "BackupPassword"
    fields["location_type"] = bindings.NewEnumType("com.vmware.appliance.recovery.backup.job.location_type", reflect.TypeOf(LocationType(LocationType_FTP)))
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
    return bindings.NewStructType("com.vmware.appliance.recovery.backup.job.backup_request",fields, reflect.TypeOf(BackupRequest{}), fieldNameMap, validators)
}

func BackupJobStatusBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewStringType()
    fieldNameMap["id"] = "Id"
    fields["state"] = bindings.NewEnumType("com.vmware.appliance.recovery.backup.job.backup_restore_process_state", reflect.TypeOf(BackupRestoreProcessState(BackupRestoreProcessState_FAILED)))
    fieldNameMap["state"] = "State"
    fields["messages"] = bindings.NewListType(bindings.NewReferenceType(LocalizableMessageBindingType), reflect.TypeOf([]LocalizableMessage{}))
    fieldNameMap["messages"] = "Messages"
    fields["progress"] = bindings.NewIntegerType()
    fieldNameMap["progress"] = "Progress"
    fields["start_time"] = bindings.NewDateTimeType()
    fieldNameMap["start_time"] = "StartTime"
    fields["end_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["end_time"] = "EndTime"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.recovery.backup.job.backup_job_status",fields, reflect.TypeOf(BackupJobStatus{}), fieldNameMap, validators)
}


