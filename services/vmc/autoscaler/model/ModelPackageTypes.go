/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for package: com.vmware.model.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package model

import (
	"reflect"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/log"
	"time"
)


type AbstractEntity struct {
	Updated time.Time
    // User id that last updated this record
	UserId string
    // User id that last updated this record
	UpdatedByUserId string
    // User name that last updated this record
	UpdatedByUserName string
	Created time.Time
    // User name that last updated this record
	UserName string
    // Unique ID for this entity
	Id string
}

func (s AbstractEntity) GetType__() bindings.BindingType {
	return AbstractEntityBindingType()
}

func (s AbstractEntity) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for AbstractEntity._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type AwsEvent struct {
    // AWS instance id of the host.
	InstanceId string
    // The date & time when the AWS event for the host is scheduled. format: date-time
	StartTime time.Time
    // Type of the scheduled event (retirement, reboot, ...)
	Type_ string
    // Customer account id the instance belongs to.
	AccountId string
    // Description of the AWS scheduled event.
	Description *string
}
// Identifier denoting this class, when it is used in polymorphic context. 
//
// This value should be assigned to the property which is used to discriminate the actual type used in the polymorphic context.
const AwsEvent__TYPE_IDENTIFIER = "AwsEvent"

func (s AwsEvent) GetType__() bindings.BindingType {
	return AwsEventBindingType()
}

func (s AwsEvent) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for AwsEvent._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type EdrsClusterInfo struct {
    // Possible values are: 
    //
    // * EdrsClusterInfo#EdrsClusterInfo_STATUS_KEY_MIN_HOSTS
    // * EdrsClusterInfo#EdrsClusterInfo_STATUS_KEY_MAX_HOSTS
    // * EdrsClusterInfo#EdrsClusterInfo_STATUS_KEY_FAILED_HOSTS
    //
    //  Key identifying the status type
	StatusKey *string
    // The cluster identifier
	ClusterId string
    // The status description
	StatusMessage *string
	EdrsPolicy EdrsPolicy
}
const EdrsClusterInfo_STATUS_KEY_MIN_HOSTS = "skyscraper.autoscaler.elastic.drs.min.hosts"
const EdrsClusterInfo_STATUS_KEY_MAX_HOSTS = "skyscraper.autoscaler.elastic.drs.max.hosts"
const EdrsClusterInfo_STATUS_KEY_FAILED_HOSTS = "skyscraper.autoscaler.elastic.drs.failed.hosts"

func (s EdrsClusterInfo) GetType__() bindings.BindingType {
	return EdrsClusterInfoBindingType()
}

func (s EdrsClusterInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for EdrsClusterInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type EdrsPolicy struct {
    // True if EDRS is enabled
	EnableEdrs bool
    // The minimum number of hosts that the cluster can scale in to.
	MinHosts *int64
    // Possible values are: 
    //
    // * EdrsPolicy#EdrsPolicy_POLICY_TYPE_COST
    // * EdrsPolicy#EdrsPolicy_POLICY_TYPE_PERFORMANCE
    // * EdrsPolicy#EdrsPolicy_POLICY_TYPE_STORAGE_SCALEUP
    // * EdrsPolicy#EdrsPolicy_POLICY_TYPE_RAPID_SCALEUP
    //
    //  The EDRS policy type. This can either be 'cost', 'performance', 'storage-scaleup' or 'rapid-scaleup'.
	PolicyType *string
    // The maximum number of hosts that the cluster can scale out to.
	MaxHosts *int64
}
const EdrsPolicy_POLICY_TYPE_COST = "cost"
const EdrsPolicy_POLICY_TYPE_PERFORMANCE = "performance"
const EdrsPolicy_POLICY_TYPE_STORAGE_SCALEUP = "storage-scaleup"
const EdrsPolicy_POLICY_TYPE_RAPID_SCALEUP = "rapid-scaleup"

func (s EdrsPolicy) GetType__() bindings.BindingType {
	return EdrsPolicyBindingType()
}

func (s EdrsPolicy) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for EdrsPolicy._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type ErrorResponse struct {
    // HTTP status code
	Status int64
    // Originating request URI
	Path string
    // If true, client should retry operation
	Retryable bool
    // unique error code
	ErrorCode string
    // localized error messages
	ErrorMessages []string
}

func (s ErrorResponse) GetType__() bindings.BindingType {
	return ErrorResponseBindingType()
}

func (s ErrorResponse) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ErrorResponse._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Detailed service errors associated with a task.
type ServiceError struct {
    // The original service name of the error.
	OriginalService string
    // The parameters of the service error.
	Params []string
    // Error message in English.
	DefaultMessage *string
    // The original error code of the service.
	OriginalServiceErrorCode string
    // Localizable error code.
	ErrorCode string
    // The localized message.
	LocalizedMessage *string
}

func (s ServiceError) GetType__() bindings.BindingType {
	return ServiceErrorBindingType()
}

func (s ServiceError) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ServiceError._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type Task struct {
	Updated time.Time
    // User id that last updated this record
	UserId string
    // User id that last updated this record
	UpdatedByUserId string
    // User name that last updated this record
	UpdatedByUserName string
	Created time.Time
    // User name that last updated this record
	UserName string
    // Unique ID for this entity
	Id string
    // Possible values are: 
    //
    // * Task#Task_STATUS_STARTED
    // * Task#Task_STATUS_CANCELING
    // * Task#Task_STATUS_FINISHED
    // * Task#Task_STATUS_FAILED
    // * Task#Task_STATUS_CANCELED
	Status *string
    // UUID of resources task is acting upon
	ResourceId *string
	StartTime *time.Time
    // Service errors returned from SDDC services.
	ServiceErrors []ServiceError
	SubStatus *string
	TaskType *string
    // Task progress phases involved in current task execution
	TaskProgressPhases []TaskProgressPhase
	ErrorMessage *string
	OrgId *string
    // Estimated progress percentage the task executed format: int32
	ProgressPercent *int64
    // Estimated remaining time in minute of the task execution, < 0 means no estimation for the task. format: int32
	EstimatedRemainingMinutes *int64
	Params *data.StructValue
	EndTime *time.Time
    // The current in progress phase ID in the task execution, if none in progress, empty string returned.
	PhaseInProgress *string
	TaskVersion *string
    // Type of resource being acted upon
	ResourceType *string
}
const Task_STATUS_STARTED = "STARTED"
const Task_STATUS_CANCELING = "CANCELING"
const Task_STATUS_FINISHED = "FINISHED"
const Task_STATUS_FAILED = "FAILED"
const Task_STATUS_CANCELED = "CANCELED"

func (s Task) GetType__() bindings.BindingType {
	return TaskBindingType()
}

func (s Task) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for Task._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// A task progress can be (but does NOT have to be) divided to more meaningful progress phases.
type TaskProgressPhase struct {
    // The identifier of the task progress phase
	Id string
    // The display name of the task progress phase
	Name string
    // The percentage of the phase that has completed format: int32
	ProgressPercent int64
}

func (s TaskProgressPhase) GetType__() bindings.BindingType {
	return TaskProgressPhaseBindingType()
}

func (s TaskProgressPhase) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for TaskProgressPhase._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}





func AbstractEntityBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["updated"] = bindings.NewDateTimeType()
	fieldNameMap["updated"] = "Updated"
	fields["user_id"] = bindings.NewStringType()
	fieldNameMap["user_id"] = "UserId"
	fields["updated_by_user_id"] = bindings.NewStringType()
	fieldNameMap["updated_by_user_id"] = "UpdatedByUserId"
	fields["updated_by_user_name"] = bindings.NewStringType()
	fieldNameMap["updated_by_user_name"] = "UpdatedByUserName"
	fields["created"] = bindings.NewDateTimeType()
	fieldNameMap["created"] = "Created"
	fields["user_name"] = bindings.NewStringType()
	fieldNameMap["user_name"] = "UserName"
	fields["id"] = bindings.NewStringType()
	fieldNameMap["id"] = "Id"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.abstract_entity", fields, reflect.TypeOf(AbstractEntity{}), fieldNameMap, validators)
}

func AwsEventBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["instance_id"] = bindings.NewStringType()
	fieldNameMap["instance_id"] = "InstanceId"
	fields["start_time"] = bindings.NewDateTimeType()
	fieldNameMap["start_time"] = "StartTime"
	fields["type"] = bindings.NewStringType()
	fieldNameMap["type"] = "Type_"
	fields["account_id"] = bindings.NewStringType()
	fieldNameMap["account_id"] = "AccountId"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.aws_event", fields, reflect.TypeOf(AwsEvent{}), fieldNameMap, validators)
}

func EdrsClusterInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["status_key"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["status_key"] = "StatusKey"
	fields["cluster_id"] = bindings.NewStringType()
	fieldNameMap["cluster_id"] = "ClusterId"
	fields["status_message"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["status_message"] = "StatusMessage"
	fields["edrs_policy"] = bindings.NewReferenceType(EdrsPolicyBindingType)
	fieldNameMap["edrs_policy"] = "EdrsPolicy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.edrs_cluster_info", fields, reflect.TypeOf(EdrsClusterInfo{}), fieldNameMap, validators)
}

func EdrsPolicyBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["enable_edrs"] = bindings.NewBooleanType()
	fieldNameMap["enable_edrs"] = "EnableEdrs"
	fields["min_hosts"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["min_hosts"] = "MinHosts"
	fields["policy_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["policy_type"] = "PolicyType"
	fields["max_hosts"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["max_hosts"] = "MaxHosts"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.edrs_policy", fields, reflect.TypeOf(EdrsPolicy{}), fieldNameMap, validators)
}

func ErrorResponseBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["status"] = bindings.NewIntegerType()
	fieldNameMap["status"] = "Status"
	fields["path"] = bindings.NewStringType()
	fieldNameMap["path"] = "Path"
	fields["retryable"] = bindings.NewBooleanType()
	fieldNameMap["retryable"] = "Retryable"
	fields["error_code"] = bindings.NewStringType()
	fieldNameMap["error_code"] = "ErrorCode"
	fields["error_messages"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["error_messages"] = "ErrorMessages"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.error_response", fields, reflect.TypeOf(ErrorResponse{}), fieldNameMap, validators)
}

func ServiceErrorBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["original_service"] = bindings.NewStringType()
	fieldNameMap["original_service"] = "OriginalService"
	fields["params"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["params"] = "Params"
	fields["default_message"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["default_message"] = "DefaultMessage"
	fields["original_service_error_code"] = bindings.NewStringType()
	fieldNameMap["original_service_error_code"] = "OriginalServiceErrorCode"
	fields["error_code"] = bindings.NewStringType()
	fieldNameMap["error_code"] = "ErrorCode"
	fields["localized_message"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["localized_message"] = "LocalizedMessage"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.service_error", fields, reflect.TypeOf(ServiceError{}), fieldNameMap, validators)
}

func TaskBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["updated"] = bindings.NewDateTimeType()
	fieldNameMap["updated"] = "Updated"
	fields["user_id"] = bindings.NewStringType()
	fieldNameMap["user_id"] = "UserId"
	fields["updated_by_user_id"] = bindings.NewStringType()
	fieldNameMap["updated_by_user_id"] = "UpdatedByUserId"
	fields["updated_by_user_name"] = bindings.NewStringType()
	fieldNameMap["updated_by_user_name"] = "UpdatedByUserName"
	fields["created"] = bindings.NewDateTimeType()
	fieldNameMap["created"] = "Created"
	fields["user_name"] = bindings.NewStringType()
	fieldNameMap["user_name"] = "UserName"
	fields["id"] = bindings.NewStringType()
	fieldNameMap["id"] = "Id"
	fields["status"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["status"] = "Status"
	fields["resource_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["resource_id"] = "ResourceId"
	fields["start_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
	fieldNameMap["start_time"] = "StartTime"
	fields["service_errors"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ServiceErrorBindingType), reflect.TypeOf([]ServiceError{})))
	fieldNameMap["service_errors"] = "ServiceErrors"
	fields["sub_status"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sub_status"] = "SubStatus"
	fields["task_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["task_type"] = "TaskType"
	fields["task_progress_phases"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(TaskProgressPhaseBindingType), reflect.TypeOf([]TaskProgressPhase{})))
	fieldNameMap["task_progress_phases"] = "TaskProgressPhases"
	fields["error_message"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["error_message"] = "ErrorMessage"
	fields["org_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fields["progress_percent"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["progress_percent"] = "ProgressPercent"
	fields["estimated_remaining_minutes"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["estimated_remaining_minutes"] = "EstimatedRemainingMinutes"
	fields["params"] = bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
	fieldNameMap["params"] = "Params"
	fields["end_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
	fieldNameMap["end_time"] = "EndTime"
	fields["phase_in_progress"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["phase_in_progress"] = "PhaseInProgress"
	fields["task_version"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["task_version"] = "TaskVersion"
	fields["resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.task", fields, reflect.TypeOf(Task{}), fieldNameMap, validators)
}

func TaskProgressPhaseBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = bindings.NewStringType()
	fieldNameMap["id"] = "Id"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["progress_percent"] = bindings.NewIntegerType()
	fieldNameMap["progress_percent"] = "ProgressPercent"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.task_progress_phase", fields, reflect.TypeOf(TaskProgressPhase{}), fieldNameMap, validators)
}


