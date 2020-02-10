/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Backup.
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
)



// ``ReturnStatus`` enumeration class Defines the state of precheck
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type BackupReturnStatus string

const (
    // Check failed
	BackupReturnStatus_FAIL BackupReturnStatus = "FAIL"
    // Passed with warnings
	BackupReturnStatus_WARNING BackupReturnStatus = "WARNING"
    // Check passed
	BackupReturnStatus_OK BackupReturnStatus = "OK"
)

func (r BackupReturnStatus) BackupReturnStatus() bool {
	switch r {
	case BackupReturnStatus_FAIL:
		return true
	case BackupReturnStatus_WARNING:
		return true
	case BackupReturnStatus_OK:
		return true
	default:
		return false
	}
}


// ``LocationType`` enumeration class Defines type of all locations for backup/restore
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type BackupLocationType string

const (
    // Destination is FTP server
	BackupLocationType_FTP BackupLocationType = "FTP"
    // Destination is HTTP server
	BackupLocationType_HTTP BackupLocationType = "HTTP"
    // Destination is FTPS server
	BackupLocationType_FTPS BackupLocationType = "FTPS"
    // Destination is HTTPS server
	BackupLocationType_HTTPS BackupLocationType = "HTTPS"
    // Destination is SSH server
	BackupLocationType_SCP BackupLocationType = "SCP"
    // Destination is SFTP server
	BackupLocationType_SFTP BackupLocationType = "SFTP"
    // Destination is NFS server. This constant field was added in vSphere API 6.7.2.
	BackupLocationType_NFS BackupLocationType = "NFS"
    // Destination is SMB server. This constant field was added in vSphere API 6.7.2.
	BackupLocationType_SMB BackupLocationType = "SMB"
)

func (l BackupLocationType) BackupLocationType() bool {
	switch l {
	case BackupLocationType_FTP:
		return true
	case BackupLocationType_HTTP:
		return true
	case BackupLocationType_FTPS:
		return true
	case BackupLocationType_HTTPS:
		return true
	case BackupLocationType_SCP:
		return true
	case BackupLocationType_SFTP:
		return true
	case BackupLocationType_NFS:
		return true
	case BackupLocationType_SMB:
		return true
	default:
		return false
	}
}


// ``LocalizableMessage`` class Structure representing message
type BackupLocalizableMessage struct {
    // id in message bundle
	Id string
    // text in english
	DefaultMessage string
    // nested data
	Args []string
}

// ``ReturnResult`` class Structure representing precheck result
type BackupReturnResult struct {
    // Check status
	Status BackupReturnStatus
    // List of messages
	Messages []BackupLocalizableMessage
}

// ``BackupRequest`` class Structure representing requested backup piece
type BackupBackupRequest struct {
    // a list of optional parts. Run backup parts APIs to get list of optional parts and description
	Parts []string
    // a password for a backup piece The backupPassword must adhere to the following password requirements: At least 8 characters, cannot be more than 20 characters in length. At least 1 uppercase letter. At least 1 lowercase letter. At least 1 numeric digit. At least 1 special character (i.e. any character not in [0-9,a-z,A-Z]). Only visible ASCII characters (for example, no space).
	BackupPassword *string
    // a type of location
	LocationType BackupLocationType
    // path or url
	Location string
    // username for location
	LocationUser *string
    // password for location
	LocationPassword *string
    // Custom comment
	Comment *string
}



func backupValidateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["piece"] = bindings.NewReferenceType(BackupBackupRequestBindingType)
	fieldNameMap["piece"] = "Piece"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func backupValidateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(BackupReturnResultBindingType)
}

func backupValidateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["piece"] = bindings.NewReferenceType(BackupBackupRequestBindingType)
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


func BackupLocalizableMessageBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = bindings.NewStringType()
	fieldNameMap["id"] = "Id"
	fields["default_message"] = bindings.NewStringType()
	fieldNameMap["default_message"] = "DefaultMessage"
	fields["args"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["args"] = "Args"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.recovery.backup.localizable_message", fields, reflect.TypeOf(BackupLocalizableMessage{}), fieldNameMap, validators)
}

func BackupReturnResultBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["status"] = bindings.NewEnumType("com.vmware.appliance.recovery.backup.return_status", reflect.TypeOf(BackupReturnStatus(BackupReturnStatus_FAIL)))
	fieldNameMap["status"] = "Status"
	fields["messages"] = bindings.NewListType(bindings.NewReferenceType(BackupLocalizableMessageBindingType), reflect.TypeOf([]BackupLocalizableMessage{}))
	fieldNameMap["messages"] = "Messages"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.recovery.backup.return_result", fields, reflect.TypeOf(BackupReturnResult{}), fieldNameMap, validators)
}

func BackupBackupRequestBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["parts"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["parts"] = "Parts"
	fields["backup_password"] = bindings.NewOptionalType(bindings.NewSecretType())
	fieldNameMap["backup_password"] = "BackupPassword"
	fields["location_type"] = bindings.NewEnumType("com.vmware.appliance.recovery.backup.location_type", reflect.TypeOf(BackupLocationType(BackupLocationType_FTP)))
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
	return bindings.NewStructType("com.vmware.appliance.recovery.backup.backup_request", fields, reflect.TypeOf(BackupBackupRequest{}), fieldNameMap, validators)
}

