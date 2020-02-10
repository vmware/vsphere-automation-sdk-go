/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Schedules.
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
	"net/url"
)



// The ``DayOfWeek`` enumeration class defines the set of days when backup can be scheduled. The days can be specified as a list of individual days. You specify the days when you set the recurrence for a schedule. See SchedulesRecurrenceInfo#days. This enumeration was added in vSphere API 6.7.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type SchedulesDayOfWeek string

const (
    // Monday. This constant field was added in vSphere API 6.7.
	SchedulesDayOfWeek_MONDAY SchedulesDayOfWeek = "MONDAY"
    // Tuesday. This constant field was added in vSphere API 6.7.
	SchedulesDayOfWeek_TUESDAY SchedulesDayOfWeek = "TUESDAY"
    // Wednesday. This constant field was added in vSphere API 6.7.
	SchedulesDayOfWeek_WEDNESDAY SchedulesDayOfWeek = "WEDNESDAY"
    // Thursday. This constant field was added in vSphere API 6.7.
	SchedulesDayOfWeek_THURSDAY SchedulesDayOfWeek = "THURSDAY"
    // Friday. This constant field was added in vSphere API 6.7.
	SchedulesDayOfWeek_FRIDAY SchedulesDayOfWeek = "FRIDAY"
    // Saturday. This constant field was added in vSphere API 6.7.
	SchedulesDayOfWeek_SATURDAY SchedulesDayOfWeek = "SATURDAY"
    // Sunday. This constant field was added in vSphere API 6.7.
	SchedulesDayOfWeek_SUNDAY SchedulesDayOfWeek = "SUNDAY"
)

func (d SchedulesDayOfWeek) SchedulesDayOfWeek() bool {
	switch d {
	case SchedulesDayOfWeek_MONDAY:
		return true
	case SchedulesDayOfWeek_TUESDAY:
		return true
	case SchedulesDayOfWeek_WEDNESDAY:
		return true
	case SchedulesDayOfWeek_THURSDAY:
		return true
	case SchedulesDayOfWeek_FRIDAY:
		return true
	case SchedulesDayOfWeek_SATURDAY:
		return true
	case SchedulesDayOfWeek_SUNDAY:
		return true
	default:
		return false
	}
}


// The ``RetentionInfo`` class contains retention information associated with a schedule. This class was added in vSphere API 6.7.
type SchedulesRetentionInfo struct {
    // Number of backups which should be retained. If retention is not set, all the backups will be retained forever. This property was added in vSphere API 6.7.
	MaxCount int64
}

// The ``RecurrenceInfo`` class contains the recurrence information associated with a schedule. This class was added in vSphere API 6.7.
type SchedulesRecurrenceInfo struct {
    // Minute when backup should run. This property was added in vSphere API 6.7.
	Minute int64
    // Hour when backup should run. The hour should be specified in 24-hour clock format. This property was added in vSphere API 6.7.
	Hour int64
    // Day of week when the backup should be run. Days can be specified as list of days. This property was added in vSphere API 6.7.
	Days []SchedulesDayOfWeek
}

// The ``CreateSpec`` class contains fields to be specified for creating a new schedule. The structure includes parts, location information, encryption password and enable flag. This class was added in vSphere API 6.7.
type SchedulesCreateSpec struct {
    // List of optional parts to be backed up. Use the Parts#list method to get information about the supported parts. This property was added in vSphere API 6.7.
	Parts []string
    // Password for a backup piece. The backupPassword must adhere to the following password requirements: At least 8 characters, cannot be more than 20 characters in length. At least 1 uppercase letter. At least 1 lowercase letter. At least 1 numeric digit. At least 1 special character (i.e. any character not in [0-9,a-z,A-Z]). Only visible ASCII characters (for example, no space). This property was added in vSphere API 6.7.
	BackupPassword *string
    // URL of the backup location. This property was added in vSphere API 6.7.
	Location url.URL
    // Username for the given location. This property was added in vSphere API 6.7.
	LocationUser *string
    // Password for the given location. This property was added in vSphere API 6.7.
	LocationPassword *string
    // Enable or disable a schedule. This property was added in vSphere API 6.7.
	Enable *bool
    // Recurrence information for the schedule. This property was added in vSphere API 6.7.
	RecurrenceInfo *SchedulesRecurrenceInfo
    // Retention information for the schedule. This property was added in vSphere API 6.7.
	RetentionInfo *SchedulesRetentionInfo
}

// The ``Info`` class contains information about an existing schedule. The structure includes Schedule ID, parts, location information, encryption password, enable flag, recurrence and retention information. This class was added in vSphere API 6.7.
type SchedulesInfo struct {
    // List of optional parts that will be included in backups based on this schedule details. Use the Parts#list method to get information about the supported parts. This property was added in vSphere API 6.7.
	Parts []string
    // URL of the backup location. This property was added in vSphere API 6.7.
	Location url.URL
    // Username for the given location. This property was added in vSphere API 6.7.
	LocationUser *string
    // Enable or disable a schedule, by default when created a schedule will be enabled. This property was added in vSphere API 6.7.
	Enable bool
    // Recurrence information for the schedule. This property was added in vSphere API 6.7.
	RecurrenceInfo *SchedulesRecurrenceInfo
    // Retention information for the schedule. This property was added in vSphere API 6.7.
	RetentionInfo *SchedulesRetentionInfo
}

// The ``UpdateSpec`` class contains the fields of the existing schedule which can be updated. This class was added in vSphere API 6.7.
type SchedulesUpdateSpec struct {
    // List of optional parts. Use the Parts#list method to get information about the supported parts. This property was added in vSphere API 6.7.
	Parts []string
    // Password for a backup piece. The backupPassword must adhere to the following password requirements: At least 8 characters, cannot be more than 20 characters in length. At least 1 uppercase letter. At least 1 lowercase letter. At least 1 numeric digit. At least 1 special character (i.e. any character not in [0-9,a-z,A-Z]). Only visible ASCII characters (for example, no space). This property was added in vSphere API 6.7.
	BackupPassword *string
    // URL of the backup location. This property was added in vSphere API 6.7.
	Location *url.URL
    // Username for the given location. This property was added in vSphere API 6.7.
	LocationUser *string
    // Password for the given location. This property was added in vSphere API 6.7.
	LocationPassword *string
    // Enable or disable a schedule. This property was added in vSphere API 6.7.
	Enable *bool
    // Recurrence information for the schedule. This property was added in vSphere API 6.7.
	RecurrenceInfo *SchedulesRecurrenceInfo
    // Retention information for the schedule. This property was added in vSphere API 6.7.
	RetentionInfo *SchedulesRetentionInfo
}



func schedulesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func schedulesListOutputType() bindings.BindingType {
	return bindings.NewMapType(bindings.NewIdType([]string{"com.vmware.appliance.recovery.backup.schedule"}, ""), bindings.NewReferenceType(SchedulesInfoBindingType),reflect.TypeOf(map[string]SchedulesInfo{}))
}

func schedulesListRestMetadata() protocol.OperationRestMetadata {
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

func schedulesRunInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["schedule"] = bindings.NewIdType([]string{"com.vmware.appliance.recovery.backup.schedule"}, "")
	fields["comment"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["schedule"] = "Schedule"
	fieldNameMap["comment"] = "Comment"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func schedulesRunOutputType() bindings.BindingType {
	return bindings.NewReferenceType(JobBackupJobStatusBindingType)
}

func schedulesRunRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["schedule"] = bindings.NewIdType([]string{"com.vmware.appliance.recovery.backup.schedule"}, "")
	fields["comment"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["schedule"] = "Schedule"
	fieldNameMap["comment"] = "Comment"
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
		map[string]int{"FeatureInUse": 400,"NotFound": 404,"Error": 500})
}

func schedulesGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["schedule"] = bindings.NewIdType([]string{"com.vmware.appliance.recovery.backup.schedule"}, "")
	fieldNameMap["schedule"] = "Schedule"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func schedulesGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(SchedulesInfoBindingType)
}

func schedulesGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["schedule"] = bindings.NewIdType([]string{"com.vmware.appliance.recovery.backup.schedule"}, "")
	fieldNameMap["schedule"] = "Schedule"
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

func schedulesCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["schedule"] = bindings.NewIdType([]string{"com.vmware.appliance.recovery.backup.schedule"}, "")
	fields["spec"] = bindings.NewReferenceType(SchedulesCreateSpecBindingType)
	fieldNameMap["schedule"] = "Schedule"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func schedulesCreateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func schedulesCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["schedule"] = bindings.NewIdType([]string{"com.vmware.appliance.recovery.backup.schedule"}, "")
	fields["spec"] = bindings.NewReferenceType(SchedulesCreateSpecBindingType)
	fieldNameMap["schedule"] = "Schedule"
	fieldNameMap["spec"] = "Spec"
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
		map[string]int{"InvalidArgument": 400,"AlreadyExists": 400,"Error": 500})
}

func schedulesUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["schedule"] = bindings.NewIdType([]string{"com.vmware.appliance.recovery.backup.schedule"}, "")
	fields["spec"] = bindings.NewReferenceType(SchedulesUpdateSpecBindingType)
	fieldNameMap["schedule"] = "Schedule"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func schedulesUpdateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func schedulesUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["schedule"] = bindings.NewIdType([]string{"com.vmware.appliance.recovery.backup.schedule"}, "")
	fields["spec"] = bindings.NewReferenceType(SchedulesUpdateSpecBindingType)
	fieldNameMap["schedule"] = "Schedule"
	fieldNameMap["spec"] = "Spec"
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
		map[string]int{"InvalidArgument": 400,"NotFound": 404,"Error": 500})
}

func schedulesDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["schedule"] = bindings.NewIdType([]string{"com.vmware.appliance.recovery.backup.schedule"}, "")
	fieldNameMap["schedule"] = "Schedule"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func schedulesDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func schedulesDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["schedule"] = bindings.NewIdType([]string{"com.vmware.appliance.recovery.backup.schedule"}, "")
	fieldNameMap["schedule"] = "Schedule"
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


func SchedulesRetentionInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["max_count"] = bindings.NewIntegerType()
	fieldNameMap["max_count"] = "MaxCount"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.recovery.backup.schedules.retention_info", fields, reflect.TypeOf(SchedulesRetentionInfo{}), fieldNameMap, validators)
}

func SchedulesRecurrenceInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["minute"] = bindings.NewIntegerType()
	fieldNameMap["minute"] = "Minute"
	fields["hour"] = bindings.NewIntegerType()
	fieldNameMap["hour"] = "Hour"
	fields["days"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewEnumType("com.vmware.appliance.recovery.backup.schedules.day_of_week", reflect.TypeOf(SchedulesDayOfWeek(SchedulesDayOfWeek_MONDAY))), reflect.TypeOf([]SchedulesDayOfWeek{})))
	fieldNameMap["days"] = "Days"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.recovery.backup.schedules.recurrence_info", fields, reflect.TypeOf(SchedulesRecurrenceInfo{}), fieldNameMap, validators)
}

func SchedulesCreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["parts"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["parts"] = "Parts"
	fields["backup_password"] = bindings.NewOptionalType(bindings.NewSecretType())
	fieldNameMap["backup_password"] = "BackupPassword"
	fields["location"] = bindings.NewUriType()
	fieldNameMap["location"] = "Location"
	fields["location_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["location_user"] = "LocationUser"
	fields["location_password"] = bindings.NewOptionalType(bindings.NewSecretType())
	fieldNameMap["location_password"] = "LocationPassword"
	fields["enable"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["enable"] = "Enable"
	fields["recurrence_info"] = bindings.NewOptionalType(bindings.NewReferenceType(SchedulesRecurrenceInfoBindingType))
	fieldNameMap["recurrence_info"] = "RecurrenceInfo"
	fields["retention_info"] = bindings.NewOptionalType(bindings.NewReferenceType(SchedulesRetentionInfoBindingType))
	fieldNameMap["retention_info"] = "RetentionInfo"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.recovery.backup.schedules.create_spec", fields, reflect.TypeOf(SchedulesCreateSpec{}), fieldNameMap, validators)
}

func SchedulesInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["parts"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["parts"] = "Parts"
	fields["location"] = bindings.NewUriType()
	fieldNameMap["location"] = "Location"
	fields["location_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["location_user"] = "LocationUser"
	fields["enable"] = bindings.NewBooleanType()
	fieldNameMap["enable"] = "Enable"
	fields["recurrence_info"] = bindings.NewOptionalType(bindings.NewReferenceType(SchedulesRecurrenceInfoBindingType))
	fieldNameMap["recurrence_info"] = "RecurrenceInfo"
	fields["retention_info"] = bindings.NewOptionalType(bindings.NewReferenceType(SchedulesRetentionInfoBindingType))
	fieldNameMap["retention_info"] = "RetentionInfo"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.recovery.backup.schedules.info", fields, reflect.TypeOf(SchedulesInfo{}), fieldNameMap, validators)
}

func SchedulesUpdateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["parts"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["parts"] = "Parts"
	fields["backup_password"] = bindings.NewOptionalType(bindings.NewSecretType())
	fieldNameMap["backup_password"] = "BackupPassword"
	fields["location"] = bindings.NewOptionalType(bindings.NewUriType())
	fieldNameMap["location"] = "Location"
	fields["location_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["location_user"] = "LocationUser"
	fields["location_password"] = bindings.NewOptionalType(bindings.NewSecretType())
	fieldNameMap["location_password"] = "LocationPassword"
	fields["enable"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["enable"] = "Enable"
	fields["recurrence_info"] = bindings.NewOptionalType(bindings.NewReferenceType(SchedulesRecurrenceInfoBindingType))
	fieldNameMap["recurrence_info"] = "RecurrenceInfo"
	fields["retention_info"] = bindings.NewOptionalType(bindings.NewReferenceType(SchedulesRetentionInfoBindingType))
	fieldNameMap["retention_info"] = "RetentionInfo"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.recovery.backup.schedules.update_spec", fields, reflect.TypeOf(SchedulesUpdateSpec{}), fieldNameMap, validators)
}

