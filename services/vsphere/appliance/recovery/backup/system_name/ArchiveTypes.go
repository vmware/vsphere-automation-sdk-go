/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Archive.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package system_name

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/vsphere/appliance/recovery/backup"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
	"net/url"
	"time"
)



// The ``Info`` class represents backup archive information. This class was added in vSphere API 6.7.
type ArchiveInfo struct {
    // Time when this backup was completed. This property was added in vSphere API 6.7.
	Timestamp time.Time
    // Backup location URL. This property was added in vSphere API 6.7.
	Location url.URL
    // List of parts included in the backup. This property was added in vSphere API 6.7.
	Parts []string
    // The version of the appliance represented by the backup. This property was added in vSphere API 6.7.
	Version string
    // The system name identifier of the appliance represented by the backup. This property was added in vSphere API 6.7.
	SystemName string
    // Custom comment added by the user for this backup. This property was added in vSphere API 6.7.
	Comment string
}

// The ``Summary`` class contains commonly used information about a backup archive. This class was added in vSphere API 6.7.
type ArchiveSummary struct {
    // Backup archive identifier. This property was added in vSphere API 6.7.
	Archive string
    // Time when this backup was started. This property was added in vSphere API 6.7.
	Timestamp time.Time
    // The version of the appliance represented by the backup archive. This property was added in vSphere API 6.7.
	Version string
    // Custom comment added by the user for this backup. This property was added in vSphere API 6.7.
	Comment string
}

// The ``FilterSpec`` class contains properties used to filter the results when listing backup archives (see Archive#list). If multiple properties are specified, only backup archives matching all of the properties match the filter. This class was added in vSphere API 6.7.
type ArchiveFilterSpec struct {
    // Backup must have been taken on or after this time to match the filter. This property was added in vSphere API 6.7.
	StartTimestamp *time.Time
    // Backup must have been taken on or before this time to match the filter. This property was added in vSphere API 6.7.
	EndTimestamp *time.Time
    // Backup comment must contain this string to match the filter. This property was added in vSphere API 6.7.
	CommentSubstring *string
    // Limit result to a max count of most recent backups. This property was added in vSphere API 6.7.
	MaxResults *int64
}



func archiveGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewReferenceType(backup.LocationSpecBindingType)
	fields["system_name"] = bindings.NewIdType([]string{"com.vmware.appliance.recovery.backup.system_name"}, "")
	fields["archive"] = bindings.NewIdType([]string{"com.vmware.appliance.recovery.backup.system_name.archive"}, "")
	fieldNameMap["spec"] = "Spec"
	fieldNameMap["system_name"] = "SystemName"
	fieldNameMap["archive"] = "Archive"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func archiveGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(ArchiveInfoBindingType)
}

func archiveGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["spec"] = bindings.NewReferenceType(backup.LocationSpecBindingType)
	fields["system_name"] = bindings.NewIdType([]string{"com.vmware.appliance.recovery.backup.system_name"}, "")
	fields["archive"] = bindings.NewIdType([]string{"com.vmware.appliance.recovery.backup.system_name.archive"}, "")
	fieldNameMap["spec"] = "Spec"
	fieldNameMap["system_name"] = "SystemName"
	fieldNameMap["archive"] = "Archive"
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

func archiveListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["loc_spec"] = bindings.NewReferenceType(backup.LocationSpecBindingType)
	fields["system_name"] = bindings.NewIdType([]string{"com.vmware.appliance.recovery.backup.system_name"}, "")
	fields["filter_spec"] = bindings.NewReferenceType(ArchiveFilterSpecBindingType)
	fieldNameMap["loc_spec"] = "LocSpec"
	fieldNameMap["system_name"] = "SystemName"
	fieldNameMap["filter_spec"] = "FilterSpec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func archiveListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(ArchiveSummaryBindingType), reflect.TypeOf([]ArchiveSummary{}))
}

func archiveListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["loc_spec"] = bindings.NewReferenceType(backup.LocationSpecBindingType)
	fields["system_name"] = bindings.NewIdType([]string{"com.vmware.appliance.recovery.backup.system_name"}, "")
	fields["filter_spec"] = bindings.NewReferenceType(ArchiveFilterSpecBindingType)
	fieldNameMap["loc_spec"] = "LocSpec"
	fieldNameMap["system_name"] = "SystemName"
	fieldNameMap["filter_spec"] = "FilterSpec"
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


func ArchiveInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["timestamp"] = bindings.NewDateTimeType()
	fieldNameMap["timestamp"] = "Timestamp"
	fields["location"] = bindings.NewUriType()
	fieldNameMap["location"] = "Location"
	fields["parts"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["parts"] = "Parts"
	fields["version"] = bindings.NewStringType()
	fieldNameMap["version"] = "Version"
	fields["system_name"] = bindings.NewIdType([]string{"com.vmware.appliance.recovery.backup.system_name"}, "")
	fieldNameMap["system_name"] = "SystemName"
	fields["comment"] = bindings.NewStringType()
	fieldNameMap["comment"] = "Comment"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.recovery.backup.system_name.archive.info", fields, reflect.TypeOf(ArchiveInfo{}), fieldNameMap, validators)
}

func ArchiveSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["archive"] = bindings.NewIdType([]string{"com.vmware.appliance.recovery.backup.system_name.archive"}, "")
	fieldNameMap["archive"] = "Archive"
	fields["timestamp"] = bindings.NewDateTimeType()
	fieldNameMap["timestamp"] = "Timestamp"
	fields["version"] = bindings.NewStringType()
	fieldNameMap["version"] = "Version"
	fields["comment"] = bindings.NewStringType()
	fieldNameMap["comment"] = "Comment"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.recovery.backup.system_name.archive.summary", fields, reflect.TypeOf(ArchiveSummary{}), fieldNameMap, validators)
}

func ArchiveFilterSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["start_timestamp"] = bindings.NewOptionalType(bindings.NewDateTimeType())
	fieldNameMap["start_timestamp"] = "StartTimestamp"
	fields["end_timestamp"] = bindings.NewOptionalType(bindings.NewDateTimeType())
	fieldNameMap["end_timestamp"] = "EndTimestamp"
	fields["comment_substring"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["comment_substring"] = "CommentSubstring"
	fields["max_results"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["max_results"] = "MaxResults"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.recovery.backup.system_name.archive.filter_spec", fields, reflect.TypeOf(ArchiveFilterSpec{}), fieldNameMap, validators)
}

