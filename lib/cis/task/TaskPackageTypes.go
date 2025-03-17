// Copyright (c) 2019-2025 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for package: com.vmware.cis.task.
// Includes binding types of a top level structures and enumerations.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package task

import (
	vapiStd_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiLog_ "github.com/vmware/vsphere-automation-sdk-go/runtime/log"
	"reflect"
	"time"
)

// The “Status“ enumeration class defines the status values that can be reported for an operation.
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

// The “Progress“ class contains information describe the progress of an operation.
type Progress struct {
	// Total amount of the work for the operation.
	Total int64
	// The amount of work completed for the operation. The value can only be incremented.
	Completed int64
	// Message about the work progress.
	Message vapiStd_.LocalizableMessage
}

func (s *Progress) GetType__() vapiBindings_.BindingType {
	return ProgressBindingType()
}

func (s *Progress) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for Progress._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// The “CommonInfo“ class contains information common to all tasks.
type CommonInfo struct {
	// Description of the operation associated with the task.
	Description vapiStd_.LocalizableMessage
	// Identifier of the service containing the operation.
	Service string
	// Identifier of the operation associated with the task.
	Operation string
	// Parent of the current task.
	Parent *string
	// Identifier of the target created by the operation or an existing one the operation performed on.
	Target *vapiStd_.DynamicID
	// Status of the operation associated with the task.
	Status StatusEnum
	// Flag to indicate whether or not the operation can be cancelled. The value may change as the operation progresses.
	Cancelable bool
	// Description of the error if the operation status is "FAILED".
	Error_ *vapiData_.ErrorValue
	// Time when the operation is started.
	StartTime *time.Time
	// Time when the operation is completed.
	EndTime *time.Time
	// Name of the user who performed the operation.
	User *string
}

func (s *CommonInfo) GetType__() vapiBindings_.BindingType {
	return CommonInfoBindingType()
}

func (s *CommonInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for CommonInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// The “Info“ class contains information about a task.
type Info struct {
	// Progress of the operation.
	Progress *Progress
	// Result of the operation.
	//
	//  If an operation reports partial results before it completes, this property could be map with bool value before the CommonInfo#status has the value StatusEnum#Status_SUCCEEDED. The value could change as the operation progresses.
	Result vapiData_.DataValue
	// Description of the operation associated with the task.
	Description vapiStd_.LocalizableMessage
	// Identifier of the service containing the operation.
	Service string
	// Identifier of the operation associated with the task.
	Operation string
	// Parent of the current task.
	Parent *string
	// Identifier of the target created by the operation or an existing one the operation performed on.
	Target *vapiStd_.DynamicID
	// Status of the operation associated with the task.
	Status StatusEnum
	// Flag to indicate whether or not the operation can be cancelled. The value may change as the operation progresses.
	Cancelable bool
	// Description of the error if the operation status is "FAILED".
	Error_ *vapiData_.ErrorValue
	// Time when the operation is started.
	StartTime *time.Time
	// Time when the operation is completed.
	EndTime *time.Time
	// Name of the user who performed the operation.
	User *string
}

func (s *Info) GetType__() vapiBindings_.BindingType {
	return InfoBindingType()
}

func (s *Info) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for Info._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

func ProgressBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["total"] = vapiBindings_.NewIntegerType()
	fieldNameMap["total"] = "Total"
	fields["completed"] = vapiBindings_.NewIntegerType()
	fieldNameMap["completed"] = "Completed"
	fields["message"] = vapiBindings_.NewReferenceType(vapiStd_.LocalizableMessageBindingType)
	fieldNameMap["message"] = "Message"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.cis.task.progress", fields, reflect.TypeOf(Progress{}), fieldNameMap, validators)
}

func CommonInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["description"] = vapiBindings_.NewReferenceType(vapiStd_.LocalizableMessageBindingType)
	fieldNameMap["description"] = "Description"
	fields["service"] = vapiBindings_.NewIdType([]string{"com.vmware.vapi.service"}, "")
	fieldNameMap["service"] = "Service"
	fields["operation"] = vapiBindings_.NewIdType([]string{"com.vmware.vapi.operation"}, "")
	fieldNameMap["operation"] = "Operation"
	fields["parent"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIdType([]string{"com.vmware.cis.task"}, ""))
	fieldNameMap["parent"] = "Parent"
	fields["target"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(vapiStd_.DynamicIDBindingType))
	fieldNameMap["target"] = "Target"
	fields["status"] = vapiBindings_.NewEnumType("com.vmware.cis.task.status", reflect.TypeOf(StatusEnum(Status_PENDING)))
	fieldNameMap["status"] = "Status"
	fields["cancelable"] = vapiBindings_.NewBooleanType()
	fieldNameMap["cancelable"] = "Cancelable"
	fields["error"] = vapiBindings_.NewOptionalType(vapiBindings_.NewAnyErrorType())
	fieldNameMap["error"] = "Error_"
	fields["start_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewDateTimeType())
	fieldNameMap["start_time"] = "StartTime"
	fields["end_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewDateTimeType())
	fieldNameMap["end_time"] = "EndTime"
	fields["user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["user"] = "User"
	var validators = []vapiBindings_.Validator{}
	uv1 := vapiBindings_.NewUnionValidator("status",
		map[string][]vapiBindings_.FieldData{
			"FAILED": []vapiBindings_.FieldData{
				vapiBindings_.NewFieldData("error", false),
				vapiBindings_.NewFieldData("start_time", true),
				vapiBindings_.NewFieldData("end_time", true),
			},
			"RUNNING": []vapiBindings_.FieldData{
				vapiBindings_.NewFieldData("start_time", true),
			},
			"BLOCKED": []vapiBindings_.FieldData{
				vapiBindings_.NewFieldData("start_time", true),
			},
			"SUCCEEDED": []vapiBindings_.FieldData{
				vapiBindings_.NewFieldData("start_time", true),
				vapiBindings_.NewFieldData("end_time", true),
			},
			"PENDING": []vapiBindings_.FieldData{},
		},
	)
	validators = append(validators, uv1)
	return vapiBindings_.NewStructType("com.vmware.cis.task.common_info", fields, reflect.TypeOf(CommonInfo{}), fieldNameMap, validators)
}

func InfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["progress"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(ProgressBindingType))
	fieldNameMap["progress"] = "Progress"
	fields["result"] = vapiBindings_.NewOptionalType(vapiBindings_.NewOpaqueType())
	fieldNameMap["result"] = "Result"
	fields["description"] = vapiBindings_.NewReferenceType(vapiStd_.LocalizableMessageBindingType)
	fieldNameMap["description"] = "Description"
	fields["service"] = vapiBindings_.NewIdType([]string{"com.vmware.vapi.service"}, "")
	fieldNameMap["service"] = "Service"
	fields["operation"] = vapiBindings_.NewIdType([]string{"com.vmware.vapi.operation"}, "")
	fieldNameMap["operation"] = "Operation"
	fields["parent"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIdType([]string{"com.vmware.cis.task"}, ""))
	fieldNameMap["parent"] = "Parent"
	fields["target"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(vapiStd_.DynamicIDBindingType))
	fieldNameMap["target"] = "Target"
	fields["status"] = vapiBindings_.NewEnumType("com.vmware.cis.task.status", reflect.TypeOf(StatusEnum(Status_PENDING)))
	fieldNameMap["status"] = "Status"
	fields["cancelable"] = vapiBindings_.NewBooleanType()
	fieldNameMap["cancelable"] = "Cancelable"
	fields["error"] = vapiBindings_.NewOptionalType(vapiBindings_.NewAnyErrorType())
	fieldNameMap["error"] = "Error_"
	fields["start_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewDateTimeType())
	fieldNameMap["start_time"] = "StartTime"
	fields["end_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewDateTimeType())
	fieldNameMap["end_time"] = "EndTime"
	fields["user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["user"] = "User"
	var validators = []vapiBindings_.Validator{}
	uv1 := vapiBindings_.NewUnionValidator("status",
		map[string][]vapiBindings_.FieldData{
			"RUNNING": []vapiBindings_.FieldData{
				vapiBindings_.NewFieldData("progress", true),
				vapiBindings_.NewFieldData("result", false),
				vapiBindings_.NewFieldData("start_time", true),
			},
			"BLOCKED": []vapiBindings_.FieldData{
				vapiBindings_.NewFieldData("progress", true),
				vapiBindings_.NewFieldData("result", false),
				vapiBindings_.NewFieldData("start_time", true),
			},
			"SUCCEEDED": []vapiBindings_.FieldData{
				vapiBindings_.NewFieldData("progress", true),
				vapiBindings_.NewFieldData("result", false),
				vapiBindings_.NewFieldData("start_time", true),
				vapiBindings_.NewFieldData("end_time", true),
			},
			"FAILED": []vapiBindings_.FieldData{
				vapiBindings_.NewFieldData("progress", true),
				vapiBindings_.NewFieldData("result", false),
				vapiBindings_.NewFieldData("error", false),
				vapiBindings_.NewFieldData("start_time", true),
				vapiBindings_.NewFieldData("end_time", true),
			},
			"PENDING": []vapiBindings_.FieldData{},
		},
	)
	validators = append(validators, uv1)
	return vapiBindings_.NewStructType("com.vmware.cis.task.info", fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}
