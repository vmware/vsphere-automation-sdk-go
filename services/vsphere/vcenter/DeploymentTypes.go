/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Deployment.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package vcenter

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/cis/task"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/vsphere/vcenter/deployment"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
	"time"
)



// The ``Task`` class contains properties to describe a particular deployment task. This class was added in vSphere API 6.7.
type DeploymentTask struct {
    // The progress info of this deployment task. This property was added in vSphere API 6.7.
	Progress *task.Progress
    // Result of the task. This property was added in vSphere API 6.7.
	Result *deployment.Notifications
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

// The ``Info`` class contains properties to describe the state of the appliance. This class was added in vSphere API 6.7.
type DeploymentInfo struct {
    // State of the vCenter Server Appliance. This property was added in vSphere API 6.7.
	State deployment.ApplianceState
    // The progress info of the current appliance status. This property was added in vSphere API 6.7.
	Progress *task.Progress
    // The ordered list of subtasks for this deployment operation. This property was added in vSphere API 6.7.
	SubtaskOrder []string
    // The map of the deployment subtasks and their status infomation. This property was added in vSphere API 6.7.
	Subtasks map[string]DeploymentTask
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



func deploymentGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func deploymentGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(DeploymentInfoBindingType)
}

func deploymentGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
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
		map[string]int{"Unauthenticated": 401,"NotFound": 404})
}

func deploymentRollbackInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func deploymentRollbackOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func deploymentRollbackRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
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
		map[string]int{"Unsupported": 400,"Unauthenticated": 401})
}


func DeploymentTaskBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["progress"] = bindings.NewOptionalType(bindings.NewReferenceType(task.ProgressBindingType))
	fieldNameMap["progress"] = "Progress"
	fields["result"] = bindings.NewOptionalType(bindings.NewReferenceType(deployment.NotificationsBindingType))
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
	return bindings.NewStructType("com.vmware.vcenter.deployment.task", fields, reflect.TypeOf(DeploymentTask{}), fieldNameMap, validators)
}

func DeploymentInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["state"] = bindings.NewEnumType("com.vmware.vcenter.deployment.appliance_state", reflect.TypeOf(deployment.ApplianceState(deployment.ApplianceState_NOT_INITIALIZED)))
	fieldNameMap["state"] = "State"
	fields["progress"] = bindings.NewOptionalType(bindings.NewReferenceType(task.ProgressBindingType))
	fieldNameMap["progress"] = "Progress"
	fields["subtask_order"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["subtask_order"] = "SubtaskOrder"
	fields["subtasks"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewReferenceType(DeploymentTaskBindingType),reflect.TypeOf(map[string]DeploymentTask{})))
	fieldNameMap["subtasks"] = "Subtasks"
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
			"RUNNING": []bindings.FieldData{
				bindings.NewFieldData("progress", true),
				bindings.NewFieldData("start_time", true),
			},
			"BLOCKED": []bindings.FieldData{
				bindings.NewFieldData("progress", true),
				bindings.NewFieldData("start_time", true),
			},
			"SUCCEEDED": []bindings.FieldData{
				bindings.NewFieldData("progress", true),
				bindings.NewFieldData("start_time", true),
				bindings.NewFieldData("end_time", true),
			},
			"FAILED": []bindings.FieldData{
				bindings.NewFieldData("progress", true),
				bindings.NewFieldData("error", false),
				bindings.NewFieldData("start_time", true),
				bindings.NewFieldData("end_time", true),
			},
			"PENDING": []bindings.FieldData{},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.deployment.info", fields, reflect.TypeOf(DeploymentInfo{}), fieldNameMap, validators)
}

