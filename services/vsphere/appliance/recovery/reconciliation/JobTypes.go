/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Job.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package reconciliation

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
	"time"
)



// The ``Status`` enumeration class defines the status values that can be reported for an operation. This enumeration was added in vSphere API 6.7.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type JobStatus string

const (
    // The operation is not running. This constant field was added in vSphere API 6.7.
	JobStatus_NONE JobStatus = "NONE"
    // The operation is in progress. This constant field was added in vSphere API 6.7.
	JobStatus_RUNNING JobStatus = "RUNNING"
    // The operation completed successfully. This constant field was added in vSphere API 6.7.
	JobStatus_SUCCEEDED JobStatus = "SUCCEEDED"
    // The operation failed. This constant field was added in vSphere API 6.7.
	JobStatus_FAILED JobStatus = "FAILED"
)

func (s JobStatus) JobStatus() bool {
	switch s {
	case JobStatus_NONE:
		return true
	case JobStatus_RUNNING:
		return true
	case JobStatus_SUCCEEDED:
		return true
	case JobStatus_FAILED:
		return true
	default:
		return false
	}
}


// The ``CreateSpec`` class has the fields to request the start of reconciliation job. This class was added in vSphere API 6.7.
type JobCreateSpec struct {
    // Administrators username for SSO. This property was added in vSphere API 6.7.
	SsoAdminUserName *string
    // Password for SSO admin user. This property was added in vSphere API 6.7.
	SsoAdminUserPassword *string
    // Flag indicating whether warnings should be ignored during reconciliation. This property was added in vSphere API 6.7.
	IgnoreWarnings *bool
}

// The ``Info`` class represents the reconciliation job information. It contains information related to current Status, any associated messages and progress as percentage. This class was added in vSphere API 6.7.
type JobInfo struct {
    // Description of the operation associated with the task. This property was added in vSphere API 6.7.
	Description std.LocalizableMessage
    // Name of the service containing the operation. This property was added in vSphere API 6.7.
	Service string
    // Name of the operation associated with the task. This property was added in vSphere API 6.7.
	Operation string
    // Parent of the current task. This property was added in vSphere API 6.7.
	Parent *string
    // Identifier of the target resource the operation modifies. This property was added in vSphere API 6.7.
	Target *std.DynamicID
    // Status of the operation associated with the task. This property was added in vSphere API 6.7.
	Status JobStatus
    // Flag to indicate whether or not the operation can be cancelled. The value may change as the operation progresses. This property was added in vSphere API 6.7.
	Cancelable *bool
    // Description of the error if the operation status is "FAILED". This property was added in vSphere API 6.7.
	Error_ *data.ErrorValue
    // Time when the operation is started. This property was added in vSphere API 6.7.
	StartTime *time.Time
    // Time when the operation is completed. This property was added in vSphere API 6.7.
	EndTime *time.Time
    // A list of localized messages. This property was added in vSphere API 6.7.
	Messages []std.LocalizableMessage
    // The progress of the job as a percentage. This property was added in vSphere API 6.7.
	Progress int64
}



func jobCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewReferenceType(JobCreateSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func jobCreateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(JobInfoBindingType)
}

func jobCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["spec"] = bindings.NewReferenceType(JobCreateSpecBindingType)
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
		map[string]int{"FeatureInUse": 400,"NotAllowedInCurrentState": 400,"InvalidArgument": 400,"Error": 500})
}

func jobGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func jobGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(JobInfoBindingType)
}

func jobGetRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"NotFound": 404,"Error": 500})
}


func JobCreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["sso_admin_user_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sso_admin_user_name"] = "SsoAdminUserName"
	fields["sso_admin_user_password"] = bindings.NewOptionalType(bindings.NewSecretType())
	fieldNameMap["sso_admin_user_password"] = "SsoAdminUserPassword"
	fields["ignore_warnings"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["ignore_warnings"] = "IgnoreWarnings"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.recovery.reconciliation.job.create_spec", fields, reflect.TypeOf(JobCreateSpec{}), fieldNameMap, validators)
}

func JobInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["description"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
	fieldNameMap["description"] = "Description"
	fields["service"] = bindings.NewIdType([]string{"com.vmware.appliance.recovery.reconciliation.job"}, "")
	fieldNameMap["service"] = "Service"
	fields["operation"] = bindings.NewStringType()
	fieldNameMap["operation"] = "Operation"
	fields["parent"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.appliance.recovery.reconciliation.job"}, ""))
	fieldNameMap["parent"] = "Parent"
	fields["target"] = bindings.NewOptionalType(bindings.NewReferenceType(std.DynamicIDBindingType))
	fieldNameMap["target"] = "Target"
	fields["status"] = bindings.NewEnumType("com.vmware.appliance.recovery.reconciliation.job.status", reflect.TypeOf(JobStatus(JobStatus_NONE)))
	fieldNameMap["status"] = "Status"
	fields["cancelable"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["cancelable"] = "Cancelable"
	fields["error"] = bindings.NewOptionalType(bindings.NewAnyErrorType())
	fieldNameMap["error"] = "Error_"
	fields["start_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
	fieldNameMap["start_time"] = "StartTime"
	fields["end_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
	fieldNameMap["end_time"] = "EndTime"
	fields["messages"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
	fieldNameMap["messages"] = "Messages"
	fields["progress"] = bindings.NewIntegerType()
	fieldNameMap["progress"] = "Progress"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("status",
		map[string][]bindings.FieldData{
			"FAILED": []bindings.FieldData{
				bindings.NewFieldData("error", false),
				bindings.NewFieldData("start_time", true),
				bindings.NewFieldData("end_time", true),
			},
			"RUNNING": []bindings.FieldData{
				bindings.NewFieldData("start_time", true),
			},
			"SUCCEEDED": []bindings.FieldData{
				bindings.NewFieldData("start_time", true),
				bindings.NewFieldData("end_time", true),
			},
			"NONE": []bindings.FieldData{},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.appliance.recovery.reconciliation.job.info", fields, reflect.TypeOf(JobInfo{}), fieldNameMap, validators)
}

