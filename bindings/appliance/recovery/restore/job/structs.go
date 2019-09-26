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
)



// ``ReturnStatus`` enumeration class Defines the state of precheck
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type ReturnStatus string

const (
    // Check failed
     ReturnStatus_FAIL ReturnStatus = "FAIL"
    // Passed with warnings
     ReturnStatus_WARNING ReturnStatus = "WARNING"
    // Check passed
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




// ``LocationType`` enumeration class Defines type of all locations for backup/restore
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type LocationType string

const (
    // Destination is FTP server
     LocationType_FTP LocationType = "FTP"
    // Destination is HTTP server
     LocationType_HTTP LocationType = "HTTP"
    // Destination is FTPS server
     LocationType_FTPS LocationType = "FTPS"
    // Destination is HTTPS server
     LocationType_HTTPS LocationType = "HTTPS"
    // Destination is SSH server
     LocationType_SCP LocationType = "SCP"
    // Destination is SFTP server
     LocationType_SFTP LocationType = "SFTP"
    // Destination is NFS server
     LocationType_NFS LocationType = "NFS"
    // Destination is SMB server
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




// ``BackupRestoreProcessState`` enumeration class Defines state of backup/restore process
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type BackupRestoreProcessState string

const (
    // Failed
     BackupRestoreProcessState_FAILED BackupRestoreProcessState = "FAILED"
    // In progress
     BackupRestoreProcessState_INPROGRESS BackupRestoreProcessState = "INPROGRESS"
    // Not started
     BackupRestoreProcessState_NONE BackupRestoreProcessState = "NONE"
    // Completed successfully
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





// ``LocalizableMessage`` class Structure representing message
 type LocalizableMessage struct {
    Id string
    DefaultMessage string
    Args []string
}






// ``ReturnResult`` class Structure representing precheck result
 type ReturnResult struct {
    Status ReturnStatus
    Messages []LocalizableMessage
}






// ``RestoreRequest`` class Structure representing requested restore piece
 type RestoreRequest struct {
    BackupPassword *string
    LocationType LocationType
    Location string
    LocationUser *string
    LocationPassword *string
    SsoAdminUserName *string
    SsoAdminUserPassword *string
    IgnoreWarnings *bool
}






// ``RestoreJobStatus`` class Structure representing backup restore status
 type RestoreJobStatus struct {
    State BackupRestoreProcessState
    Messages []LocalizableMessage
    Progress int64
}









func cancelInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
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
       map[string]int{"Error": 500})
}


func createInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["piece"] = bindings.NewReferenceType(RestoreRequestBindingType)
    fieldNameMap["piece"] = "Piece"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func createOutputType() bindings.BindingType {
    return bindings.NewReferenceType(RestoreJobStatusBindingType)
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
       map[string]int{"FeatureInUse": 400,"NotAllowedInCurrentState": 400,"Error": 500})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewReferenceType(RestoreJobStatusBindingType)
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
       map[string]int{"Error": 500})
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
    return bindings.NewStructType("com.vmware.appliance.recovery.restore.job.localizable_message",fields, reflect.TypeOf(LocalizableMessage{}), fieldNameMap, validators)
}

func ReturnResultBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["status"] = bindings.NewEnumType("com.vmware.appliance.recovery.restore.job.return_status", reflect.TypeOf(ReturnStatus(ReturnStatus_FAIL)))
    fieldNameMap["status"] = "Status"
    fields["messages"] = bindings.NewListType(bindings.NewReferenceType(LocalizableMessageBindingType), reflect.TypeOf([]LocalizableMessage{}))
    fieldNameMap["messages"] = "Messages"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.recovery.restore.job.return_result",fields, reflect.TypeOf(ReturnResult{}), fieldNameMap, validators)
}

func RestoreRequestBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["backup_password"] = bindings.NewOptionalType(bindings.NewSecretType())
    fieldNameMap["backup_password"] = "BackupPassword"
    fields["location_type"] = bindings.NewEnumType("com.vmware.appliance.recovery.restore.job.location_type", reflect.TypeOf(LocationType(LocationType_FTP)))
    fieldNameMap["location_type"] = "LocationType"
    fields["location"] = bindings.NewStringType()
    fieldNameMap["location"] = "Location"
    fields["location_user"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["location_user"] = "LocationUser"
    fields["location_password"] = bindings.NewOptionalType(bindings.NewSecretType())
    fieldNameMap["location_password"] = "LocationPassword"
    fields["sso_admin_user_name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["sso_admin_user_name"] = "SsoAdminUserName"
    fields["sso_admin_user_password"] = bindings.NewOptionalType(bindings.NewSecretType())
    fieldNameMap["sso_admin_user_password"] = "SsoAdminUserPassword"
    fields["ignore_warnings"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["ignore_warnings"] = "IgnoreWarnings"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.recovery.restore.job.restore_request",fields, reflect.TypeOf(RestoreRequest{}), fieldNameMap, validators)
}

func RestoreJobStatusBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["state"] = bindings.NewEnumType("com.vmware.appliance.recovery.restore.job.backup_restore_process_state", reflect.TypeOf(BackupRestoreProcessState(BackupRestoreProcessState_FAILED)))
    fieldNameMap["state"] = "State"
    fields["messages"] = bindings.NewListType(bindings.NewReferenceType(LocalizableMessageBindingType), reflect.TypeOf([]LocalizableMessage{}))
    fieldNameMap["messages"] = "Messages"
    fields["progress"] = bindings.NewIntegerType()
    fieldNameMap["progress"] = "Progress"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.recovery.restore.job.restore_job_status",fields, reflect.TypeOf(RestoreJobStatus{}), fieldNameMap, validators)
}


