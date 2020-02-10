/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Restore.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package recovery

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
type RestoreLocationType string

const (
    // Destination is FTP server
	RestoreLocationType_FTP RestoreLocationType = "FTP"
    // Destination is HTTP server
	RestoreLocationType_HTTP RestoreLocationType = "HTTP"
    // Destination is FTPS server
	RestoreLocationType_FTPS RestoreLocationType = "FTPS"
    // Destination is HTTPS server
	RestoreLocationType_HTTPS RestoreLocationType = "HTTPS"
    // Destination is SSH server
	RestoreLocationType_SCP RestoreLocationType = "SCP"
    // Destination is SFTP server
	RestoreLocationType_SFTP RestoreLocationType = "SFTP"
    // Destination is NFS server. This constant field was added in vSphere API 6.7.2.
	RestoreLocationType_NFS RestoreLocationType = "NFS"
    // Destination is SMB server. This constant field was added in vSphere API 6.7.2.
	RestoreLocationType_SMB RestoreLocationType = "SMB"
)

func (l RestoreLocationType) RestoreLocationType() bool {
	switch l {
	case RestoreLocationType_FTP:
		return true
	case RestoreLocationType_HTTP:
		return true
	case RestoreLocationType_FTPS:
		return true
	case RestoreLocationType_HTTPS:
		return true
	case RestoreLocationType_SCP:
		return true
	case RestoreLocationType_SFTP:
		return true
	case RestoreLocationType_NFS:
		return true
	case RestoreLocationType_SMB:
		return true
	default:
		return false
	}
}


// ``RestoreRequest`` class Structure representing requested restore piece
type RestoreRestoreRequest struct {
    // a password for a backup piece
	BackupPassword *string
    // a type of location
	LocationType RestoreLocationType
    // path or url
	Location string
    // username for location
	LocationUser *string
    // password for location
	LocationPassword *string
    // Administrators username for SSO. This property was added in vSphere API 6.7.
	SsoAdminUserName *string
    // The password for SSO admin user. This property was added in vSphere API 6.7.
	SsoAdminUserPassword *string
}

// ``LocalizableMessage`` class Structure representing message
type RestoreLocalizableMessage struct {
    // id in message bundle
	Id string
    // text in english
	DefaultMessage string
    // nested data
	Args []string
}

// ``Metadata`` class Structure representing metadata
type RestoreMetadata struct {
    // Time when this backup was completed.
	Timestamp time.Time
    // List of parts included in the backup.
	Parts []string
    // VCSA version
	Version string
    // Box name is PNID/ FQDN etc
	Boxname string
    // Is SSO login required for the vCenter server. This property was added in vSphere API 6.7.
	SsoLoginRequired *bool
    // Custom comment
	Comment string
    // Does the VCSA match the deployment type, network properties and version of backed up VC
	Applicable bool
    // Any messages if the backup is not aplicable
	Messages []RestoreLocalizableMessage
}



func restoreValidateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["piece"] = bindings.NewReferenceType(RestoreRestoreRequestBindingType)
	fieldNameMap["piece"] = "Piece"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func restoreValidateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(RestoreMetadataBindingType)
}

func restoreValidateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["piece"] = bindings.NewReferenceType(RestoreRestoreRequestBindingType)
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
		map[string]int{"Error": 500})
}


func RestoreRestoreRequestBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["backup_password"] = bindings.NewOptionalType(bindings.NewSecretType())
	fieldNameMap["backup_password"] = "BackupPassword"
	fields["location_type"] = bindings.NewEnumType("com.vmware.appliance.recovery.restore.location_type", reflect.TypeOf(RestoreLocationType(RestoreLocationType_FTP)))
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
	return bindings.NewStructType("com.vmware.appliance.recovery.restore.restore_request", fields, reflect.TypeOf(RestoreRestoreRequest{}), fieldNameMap, validators)
}

func RestoreLocalizableMessageBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = bindings.NewStringType()
	fieldNameMap["id"] = "Id"
	fields["default_message"] = bindings.NewStringType()
	fieldNameMap["default_message"] = "DefaultMessage"
	fields["args"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["args"] = "Args"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.recovery.restore.localizable_message", fields, reflect.TypeOf(RestoreLocalizableMessage{}), fieldNameMap, validators)
}

func RestoreMetadataBindingType() bindings.BindingType {
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
	fields["messages"] = bindings.NewListType(bindings.NewReferenceType(RestoreLocalizableMessageBindingType), reflect.TypeOf([]RestoreLocalizableMessage{}))
	fieldNameMap["messages"] = "Messages"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.recovery.restore.metadata", fields, reflect.TypeOf(RestoreMetadata{}), fieldNameMap, validators)
}

