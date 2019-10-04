/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Backup.
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






// ``BackupRequest`` class Structure representing requested backup piece
 type BackupRequest struct {
    Parts []string
    BackupPassword *string
    LocationType LocationType
    Location string
    LocationUser *string
    LocationPassword *string
    Comment *string
}









func validateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["piece"] = bindings.NewReferenceType(BackupRequestBindingType)
    fieldNameMap["piece"] = "Piece"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func validateOutputType() bindings.BindingType {
    return bindings.NewReferenceType(ReturnResultBindingType)
}

func validateRestMetadata() protocol.OperationRestMetadata {
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
    return bindings.NewStructType("com.vmware.appliance.recovery.backup.localizable_message",fields, reflect.TypeOf(LocalizableMessage{}), fieldNameMap, validators)
}

func ReturnResultBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["status"] = bindings.NewEnumType("com.vmware.appliance.recovery.backup.return_status", reflect.TypeOf(ReturnStatus(ReturnStatus_FAIL)))
    fieldNameMap["status"] = "Status"
    fields["messages"] = bindings.NewListType(bindings.NewReferenceType(LocalizableMessageBindingType), reflect.TypeOf([]LocalizableMessage{}))
    fieldNameMap["messages"] = "Messages"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.recovery.backup.return_result",fields, reflect.TypeOf(ReturnResult{}), fieldNameMap, validators)
}

func BackupRequestBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["parts"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["parts"] = "Parts"
    fields["backup_password"] = bindings.NewOptionalType(bindings.NewSecretType())
    fieldNameMap["backup_password"] = "BackupPassword"
    fields["location_type"] = bindings.NewEnumType("com.vmware.appliance.recovery.backup.location_type", reflect.TypeOf(LocationType(LocationType_FTP)))
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
    return bindings.NewStructType("com.vmware.appliance.recovery.backup.backup_request",fields, reflect.TypeOf(BackupRequest{}), fieldNameMap, validators)
}


