// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for package: com.vmware.cis.task.
// Includes binding types of a top level structures and enumerations.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package task

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/log"
	"reflect"
	"time"
)

// The ``Status`` enumeration class defines the status values that can be reported for an operation.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type StatusEnum string

const (
	// The operation is in pending state.
	Status_PENDING StatusEnum = "PENDING"
	// The operation is in progress.
	Status_RUNNING StatusEnum = "RUNNING"
	// The operation is blocked.
	Status_BLOCKED StatusEnum = "BLOCKED"
	// The operation completed successfully.
	Status_SUCCEEDED StatusEnum = "SUCCEEDED"
	// The operation failed.
	Status_FAILED StatusEnum = "FAILED"
)

func (s StatusEnum) StatusEnum() bool {
	switch s {
	case Status_PENDING:
		return true
	case Status_RUNNING:
		return true
	case Status_BLOCKED:
		return true
	case Status_SUCCEEDED:
		return true
	case Status_FAILED:
		return true
	default:
		return false
	}
}

// The ``Progress`` class contains information describe the progress of an operation.
type Progress struct {
	// Total amount of the work for the operation.
	Total int64
	// The amount of work completed for the operation. The value can only be incremented.
	Completed int64
	// Message about the work progress.
	Message std.LocalizableMessage
}

func (s *Progress) GetType__() bindings.BindingType {
	return ProgressBindingType()
}

func (s *Progress) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for Progress._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// The ``CommonInfo`` class contains information common to all tasks.
type CommonInfo struct {
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
	Status StatusEnum
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

func (s *CommonInfo) GetType__() bindings.BindingType {
	return CommonInfoBindingType()
}

func (s *CommonInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for CommonInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// The ``Info`` class contains information about a task.
type Info struct {
	// Progress of the operation.
	Progress *Progress
	// Result of the operation. If an operation reports partial results before it completes, this property could be map with bool value before the CommonInfo#status has the value StatusEnum#Status_SUCCEEDED. The value could change as the operation progresses.
	Result data.DataValue
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
	Status StatusEnum
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

func (s *Info) GetType__() bindings.BindingType {
	return InfoBindingType()
}

func (s *Info) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for Info._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

func ProgressBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["total"] = bindings.NewIntegerType()
	fieldNameMap["total"] = "Total"
	fields["completed"] = bindings.NewIntegerType()
	fieldNameMap["completed"] = "Completed"
	fields["message"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
	fieldNameMap["message"] = "Message"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.cis.task.progress", fields, reflect.TypeOf(Progress{}), fieldNameMap, validators)
}

func CommonInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
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
	fields["status"] = bindings.NewEnumType("com.vmware.cis.task.status", reflect.TypeOf(StatusEnum(Status_PENDING)))
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
			"FAILED": []bindings.FieldData{
				bindings.NewFieldData("error", false),
				bindings.NewFieldData("start_time", true),
				bindings.NewFieldData("end_time", true),
			},
			"RUNNING": []bindings.FieldData{
				bindings.NewFieldData("start_time", true),
			},
			"BLOCKED": []bindings.FieldData{
				bindings.NewFieldData("start_time", true),
			},
			"SUCCEEDED": []bindings.FieldData{
				bindings.NewFieldData("start_time", true),
				bindings.NewFieldData("end_time", true),
			},
			"PENDING": []bindings.FieldData{},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.cis.task.common_info", fields, reflect.TypeOf(CommonInfo{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["progress"] = bindings.NewOptionalType(bindings.NewReferenceType(ProgressBindingType))
	fieldNameMap["progress"] = "Progress"
	fields["result"] = bindings.NewOptionalType(bindings.NewOpaqueType())
	fieldNameMap["result"] = "Result"
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
	fields["status"] = bindings.NewEnumType("com.vmware.cis.task.status", reflect.TypeOf(StatusEnum(Status_PENDING)))
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
			"RUNNING": []bindings.FieldData{
				bindings.NewFieldData("progress", true),
				bindings.NewFieldData("result", false),
				bindings.NewFieldData("start_time", true),
			},
			"BLOCKED": []bindings.FieldData{
				bindings.NewFieldData("progress", true),
				bindings.NewFieldData("result", false),
				bindings.NewFieldData("start_time", true),
			},
			"SUCCEEDED": []bindings.FieldData{
				bindings.NewFieldData("progress", true),
				bindings.NewFieldData("result", false),
				bindings.NewFieldData("start_time", true),
				bindings.NewFieldData("end_time", true),
			},
			"FAILED": []bindings.FieldData{
				bindings.NewFieldData("progress", true),
				bindings.NewFieldData("result", false),
				bindings.NewFieldData("error", false),
				bindings.NewFieldData("start_time", true),
				bindings.NewFieldData("end_time", true),
			},
			"PENDING": []bindings.FieldData{},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.cis.task.info", fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}
