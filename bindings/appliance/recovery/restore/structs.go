/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Restore.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package restore

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
    "time"
)



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





// ``RestoreRequest`` class Structure representing requested restore piece
 type RestoreRequest struct {
    BackupPassword *string
    LocationType LocationType
    Location string
    LocationUser *string
    LocationPassword *string
    SsoAdminUserName *string
    SsoAdminUserPassword *string
}






// ``LocalizableMessage`` class Structure representing message
 type LocalizableMessage struct {
    Id string
    DefaultMessage string
    Args []string
}






// ``Metadata`` class Structure representing metadata
 type Metadata struct {
    Timestamp time.Time
    Parts []string
    Version string
    Boxname string
    SsoLoginRequired *bool
    Comment string
    Applicable bool
    Messages []LocalizableMessage
}









func validateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["piece"] = bindings.NewReferenceType(RestoreRequestBindingType)
    fieldNameMap["piece"] = "Piece"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func validateOutputType() bindings.BindingType {
    return bindings.NewReferenceType(MetadataBindingType)
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
       map[string]int{"Error": 500})
}



func RestoreRequestBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["backup_password"] = bindings.NewOptionalType(bindings.NewSecretType())
    fieldNameMap["backup_password"] = "BackupPassword"
    fields["location_type"] = bindings.NewEnumType("com.vmware.appliance.recovery.restore.location_type", reflect.TypeOf(LocationType(LocationType_FTP)))
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
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.recovery.restore.restore_request",fields, reflect.TypeOf(RestoreRequest{}), fieldNameMap, validators)
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
    return bindings.NewStructType("com.vmware.appliance.recovery.restore.localizable_message",fields, reflect.TypeOf(LocalizableMessage{}), fieldNameMap, validators)
}

func MetadataBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["timestamp"] = bindings.NewDateTimeType()
    fieldNameMap["timestamp"] = "Timestamp"
    fields["parts"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["parts"] = "Parts"
    fields["version"] = bindings.NewStringType()
    fieldNameMap["version"] = "Version"
    fields["boxname"] = bindings.NewStringType()
    fieldNameMap["boxname"] = "Boxname"
    fields["sso_login_required"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["sso_login_required"] = "SsoLoginRequired"
    fields["comment"] = bindings.NewStringType()
    fieldNameMap["comment"] = "Comment"
    fields["applicable"] = bindings.NewBooleanType()
    fieldNameMap["applicable"] = "Applicable"
    fields["messages"] = bindings.NewListType(bindings.NewReferenceType(LocalizableMessageBindingType), reflect.TypeOf([]LocalizableMessage{}))
    fieldNameMap["messages"] = "Messages"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.recovery.restore.metadata",fields, reflect.TypeOf(Metadata{}), fieldNameMap, validators)
}


