/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Details.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package job

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/cis/task"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
	"net/url"
	"time"
)



// The ``Type`` enumeration class defines the type of backup job. This enumeration was added in vSphere API 6.7.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type DetailsType string

const (
    // Job type is Scheduled. This constant field was added in vSphere API 6.7.
	DetailsType_SCHEDULED DetailsType = "SCHEDULED"
    // Job type is Manual. This constant field was added in vSphere API 6.7.
	DetailsType_MANUAL DetailsType = "MANUAL"
)

func (t DetailsType) DetailsType() bool {
	switch t {
	case DetailsType_SCHEDULED:
		return true
	case DetailsType_MANUAL:
		return true
	default:
		return false
	}
}


// The ``Info`` class contains information about a backup job. This class was added in vSphere API 6.7.
type DetailsInfo struct {
    // URL of the backup location. This property was added in vSphere API 6.7.
	Location url.URL
    // Time in seconds since the backup job was started or the time it took to complete the backup job. This property was added in vSphere API 6.7.
	Duration *int64
    // Size of the backup data transferred to remote location. This property was added in vSphere API 6.7.
	Size *int64
    // Progress of the job. This property was added in vSphere API 6.7.
	Progress *task.Progress
    // The username for the remote backup location. This property was added in vSphere API 6.7.
	LocationUser string
    // Type of the backup job. Indicates whether the backup was started manually or as a scheduled backup. This property was added in vSphere API 6.7.
	Type_ DetailsType
    // List of any info/warning/error messages returned by the backup job. This property was added in vSphere API 6.7.
	Messages []std.LocalizableMessage
    // Information about the build of the appliance. This property was added in vSphere API 6.7.2.
	Build *DetailsBuildInfo
    // Description of the operation associated with the task.
	Description std.LocalizableMessage
    // Identifier of the service containing the operation.
	Service string
    // Identifier of the operation associated with the task.
	Operation string
    // Parent of the current task.
	Parent *string
    // Identifier of the target created by the operation or an existing one the operation performed on.
	Target *std.DynamicID
    // Status of the operation associated with the task.
	Status task.Status
    // Flag to indicate whether or not the operation can be cancelled. The value may change as the operation progresses.
	Cancelable bool
    // Description of the error if the operation status is "FAILED".
	Error_ *data.ErrorValue
    // Time when the operation is started.
	StartTime *time.Time
    // Time when the operation is completed.
	EndTime *time.Time
    // Name of the user who performed the operation.
	User *string
}

// The ``FilterSpec`` class contains properties used to filter the results when listing backup jobs details (see Details#list). This class was added in vSphere API 6.7.
type DetailsFilterSpec struct {
    // Identifiers of backup jobs that can match the filter. This property was added in vSphere API 6.7.
	Jobs map[string]bool
}

// The ``BuildInfo`` class contains information about the build of the appliance. This class was added in vSphere API 6.7.2.
type DetailsBuildInfo struct {
    // Appliance product type, for example 6.8.2 GA. This property was added in vSphere API 6.7.2.
	VersionName string
    // Appliance version, for example 6.8.2.10000. This property was added in vSphere API 6.7.2.
	Version string
    // Build Number of the appliance. This property was added in vSphere API 6.7.2.
	BuildNumber string
}



func detailsListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(DetailsFilterSpecBindingType))
	fieldNameMap["filter"] = "Filter"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func detailsListOutputType() bindings.BindingType {
	return bindings.NewMapType(bindings.NewIdType([]string{"com.vmware.appliance.recovery.backup.job"}, ""), bindings.NewReferenceType(DetailsInfoBindingType),reflect.TypeOf(map[string]DetailsInfo{}))
}

func detailsListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(DetailsFilterSpecBindingType))
	fieldNameMap["filter"] = "Filter"
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


func DetailsInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["location"] = bindings.NewUriType()
	fieldNameMap["location"] = "Location"
	fields["duration"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["duration"] = "Duration"
	fields["size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["size"] = "Size"
	fields["progress"] = bindings.NewOptionalType(bindings.NewReferenceType(task.ProgressBindingType))
	fieldNameMap["progress"] = "Progress"
	fields["location_user"] = bindings.NewStringType()
	fieldNameMap["location_user"] = "LocationUser"
	fields["type"] = bindings.NewEnumType("com.vmware.appliance.recovery.backup.job.details.type", reflect.TypeOf(DetailsType(DetailsType_SCHEDULED)))
	fieldNameMap["type"] = "Type_"
	fields["messages"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
	fieldNameMap["messages"] = "Messages"
	fields["build"] = bindings.NewOptionalType(bindings.NewReferenceType(DetailsBuildInfoBindingType))
	fieldNameMap["build"] = "Build"
	fields["description"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
	fieldNameMap["description"] = "Description"
	fields["service"] = bindings.NewIdType([]string{"com.vmware.vapi.service"}, "")
	fieldNameMap["service"] = "Service"
	fields["operation"] = bindings.NewIdType([]string{"com.vmware.vapi.operation"}, "")
	fieldNameMap["operation"] = "Operation"
	fields["parent"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.cis.task"}, ""))
	fieldNameMap["parent"] = "Parent"
	fields["target"] = bindings.NewOptionalType(bindings.NewReferenceType(std.DynamicIDBindingType))
	fieldNameMap["target"] = "Target"
	fields["status"] = bindings.NewEnumType("com.vmware.cis.task.status", reflect.TypeOf(task.Status(task.Status_PENDING)))
	fieldNameMap["status"] = "Status"
	fields["cancelable"] = bindings.NewBooleanType()
	fieldNameMap["cancelable"] = "Cancelable"
	fields["error"] = bindings.NewOptionalType(bindings.NewAnyErrorType())
	fieldNameMap["error"] = "Error_"
	fields["start_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
	fieldNameMap["start_time"] = "StartTime"
	fields["end_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
	fieldNameMap["end_time"] = "EndTime"
	fields["user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["user"] = "User"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("status",
		map[string][]bindings.FieldData{
			"SUCCEEDED": []bindings.FieldData{
				bindings.NewFieldData("duration", true),
				bindings.NewFieldData("size", true),
				bindings.NewFieldData("progress", true),
				bindings.NewFieldData("start_time", true),
				bindings.NewFieldData("end_time", true),
			},
			"FAILED": []bindings.FieldData{
				bindings.NewFieldData("duration", true),
				bindings.NewFieldData("size", true),
				bindings.NewFieldData("progress", true),
				bindings.NewFieldData("error", false),
				bindings.NewFieldData("start_time", true),
				bindings.NewFieldData("end_time", true),
			},
			"RUNNING": []bindings.FieldData{
				bindings.NewFieldData("duration", true),
				bindings.NewFieldData("size", true),
				bindings.NewFieldData("progress", true),
				bindings.NewFieldData("start_time", true),
			},
			"BLOCKED": []bindings.FieldData{
				bindings.NewFieldData("start_time", true),
			},
			"PENDING": []bindings.FieldData{},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.appliance.recovery.backup.job.details.info", fields, reflect.TypeOf(DetailsInfo{}), fieldNameMap, validators)
}

func DetailsFilterSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["jobs"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string{"com.vmware.appliance.recovery.backup.job"}, ""), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["jobs"] = "Jobs"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.recovery.backup.job.details.filter_spec", fields, reflect.TypeOf(DetailsFilterSpec{}), fieldNameMap, validators)
}

func DetailsBuildInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["version_name"] = bindings.NewStringType()
	fieldNameMap["version_name"] = "VersionName"
	fields["version"] = bindings.NewStringType()
	fieldNameMap["version"] = "Version"
	fields["build_number"] = bindings.NewStringType()
	fieldNameMap["build_number"] = "BuildNumber"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.recovery.backup.job.details.build_info", fields, reflect.TypeOf(DetailsBuildInfo{}), fieldNameMap, validators)
}

