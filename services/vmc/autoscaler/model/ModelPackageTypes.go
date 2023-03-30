// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for package: com.vmware.model.
// Includes binding types of a top level structures and enumerations.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package model

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiLog_ "github.com/vmware/vsphere-automation-sdk-go/runtime/log"
	"reflect"
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
	Created           time.Time
	// User name that last updated this record
	UserName string
	// Unique ID for this entity
	Id string
}

func (s *AbstractEntity) GetType__() vapiBindings_.BindingType {
	return AbstractEntityBindingType()
}

func (s *AbstractEntity) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for AbstractEntity._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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

func (s *AwsEvent) GetType__() vapiBindings_.BindingType {
	return AwsEventBindingType()
}

func (s *AwsEvent) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for AwsEvent._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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
	EdrsPolicy    EdrsPolicy
}

const EdrsClusterInfo_STATUS_KEY_MIN_HOSTS = "skyscraper.autoscaler.elastic.drs.min.hosts"
const EdrsClusterInfo_STATUS_KEY_MAX_HOSTS = "skyscraper.autoscaler.elastic.drs.max.hosts"
const EdrsClusterInfo_STATUS_KEY_FAILED_HOSTS = "skyscraper.autoscaler.elastic.drs.failed.hosts"

func (s *EdrsClusterInfo) GetType__() vapiBindings_.BindingType {
	return EdrsClusterInfoBindingType()
}

func (s *EdrsClusterInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for EdrsClusterInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type EdrsPolicy struct {
	// True if EDRS is enabled
	EnableEdrs                 bool
	EdrsPolicyOptionsOverrides *EdrsPolicyOptionsOverrides
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

func (s *EdrsPolicy) GetType__() vapiBindings_.BindingType {
	return EdrsPolicyBindingType()
}

func (s *EdrsPolicy) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for EdrsPolicy._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type EdrsPolicyOptionsOverrides struct {
	// The number of hosts which will be added to the cluster during the scale out operation.
	ScaleUpHostIncrement *int64
}

func (s *EdrsPolicyOptionsOverrides) GetType__() vapiBindings_.BindingType {
	return EdrsPolicyOptionsOverridesBindingType()
}

func (s *EdrsPolicyOptionsOverrides) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for EdrsPolicyOptionsOverrides._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type EdrsPolicySpec struct {
	// Possible values are:
	//
	// * EdrsPolicySpec#EdrsPolicySpec_POLICY_TYPE_COST
	// * EdrsPolicySpec#EdrsPolicySpec_POLICY_TYPE_PERFORMANCE
	// * EdrsPolicySpec#EdrsPolicySpec_POLICY_TYPE_STORAGE_SCALEUP
	// * EdrsPolicySpec#EdrsPolicySpec_POLICY_TYPE_RAPID_SCALEUP
	//
	//  EDRS policy type.
	PolicyType *string
	// True means that cluster is eligible for the edrs policy.
	ClusterEligibleForPolicy *bool
	// True means that rapid scale-up host increment is configurable.
	ConfigurableScaleupIncrement *bool
	// List of supported number of hosts bounded by the min and max host allowed by the cluster.
	MinMaxHostRange []int64
	// List of supported rapid scale-up host increment values.
	ScaleupHostIncrementRange []int64
	// True means that policy is configurable.
	Configurable *bool
}

const EdrsPolicySpec_POLICY_TYPE_COST = "cost"
const EdrsPolicySpec_POLICY_TYPE_PERFORMANCE = "performance"
const EdrsPolicySpec_POLICY_TYPE_STORAGE_SCALEUP = "storage-scaleup"
const EdrsPolicySpec_POLICY_TYPE_RAPID_SCALEUP = "rapid-scaleup"

func (s *EdrsPolicySpec) GetType__() vapiBindings_.BindingType {
	return EdrsPolicySpecBindingType()
}

func (s *EdrsPolicySpec) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for EdrsPolicySpec._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type EdrsProvisioningSpec struct {
	EdrsClusterInfo *EdrsClusterInfo
	// List of spec for all EDRS policies.
	Policies []EdrsPolicySpec
}

func (s *EdrsProvisioningSpec) GetType__() vapiBindings_.BindingType {
	return EdrsProvisioningSpecBindingType()
}

func (s *EdrsProvisioningSpec) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for EdrsProvisioningSpec._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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

func (s *ErrorResponse) GetType__() vapiBindings_.BindingType {
	return ErrorResponseBindingType()
}

func (s *ErrorResponse) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ErrorResponse._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type RemediationClusterInfo struct {
	// The cluster name identifier
	ClusterName string
	// The cluster identifier
	ClusterId         string
	RemediationPolicy RemediationPolicy
}

func (s *RemediationClusterInfo) GetType__() vapiBindings_.BindingType {
	return RemediationClusterInfoBindingType()
}

func (s *RemediationClusterInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for RemediationClusterInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type RemediationPolicy struct {
	// True if Remediation is enabled
	AutoRemediationEnabled bool
}

func (s *RemediationPolicy) GetType__() vapiBindings_.BindingType {
	return RemediationPolicyBindingType()
}

func (s *RemediationPolicy) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for RemediationPolicy._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type RemediationPolicyPatch struct {
	// True if Remediation is enabled
	AutoRemediationEnabled bool
}

func (s *RemediationPolicyPatch) GetType__() vapiBindings_.BindingType {
	return RemediationPolicyPatchBindingType()
}

func (s *RemediationPolicyPatch) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for RemediationPolicyPatch._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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

func (s *ServiceError) GetType__() vapiBindings_.BindingType {
	return ServiceErrorBindingType()
}

func (s *ServiceError) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ServiceError._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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
	Created           time.Time
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
	StartTime  *time.Time
	// Service errors returned from SDDC services.
	ServiceErrors []ServiceError
	SubStatus     *string
	TaskType      *string
	// Task progress phases involved in current task execution
	TaskProgressPhases []TaskProgressPhase
	ErrorMessage       *string
	OrgId              *string
	// Estimated progress percentage the task executed format: int32
	ProgressPercent *int64
	// Estimated remaining time in minute of the task execution, < 0 means no estimation for the task. format: int32
	EstimatedRemainingMinutes *int64
	Params                    *vapiData_.StructValue
	EndTime                   *time.Time
	// The current in progress phase ID in the task execution, if none in progress, empty string returned.
	PhaseInProgress *string
	TaskVersion     *string
	// Type of resource being acted upon
	ResourceType *string
}

const Task_STATUS_STARTED = "STARTED"
const Task_STATUS_CANCELING = "CANCELING"
const Task_STATUS_FINISHED = "FINISHED"
const Task_STATUS_FAILED = "FAILED"
const Task_STATUS_CANCELED = "CANCELED"

func (s *Task) GetType__() vapiBindings_.BindingType {
	return TaskBindingType()
}

func (s *Task) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for Task._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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

func (s *TaskProgressPhase) GetType__() vapiBindings_.BindingType {
	return TaskProgressPhaseBindingType()
}

func (s *TaskProgressPhase) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for TaskProgressPhase._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

func AbstractEntityBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["updated"] = vapiBindings_.NewDateTimeType()
	fieldNameMap["updated"] = "Updated"
	fields["user_id"] = vapiBindings_.NewStringType()
	fieldNameMap["user_id"] = "UserId"
	fields["updated_by_user_id"] = vapiBindings_.NewStringType()
	fieldNameMap["updated_by_user_id"] = "UpdatedByUserId"
	fields["updated_by_user_name"] = vapiBindings_.NewStringType()
	fieldNameMap["updated_by_user_name"] = "UpdatedByUserName"
	fields["created"] = vapiBindings_.NewDateTimeType()
	fieldNameMap["created"] = "Created"
	fields["user_name"] = vapiBindings_.NewStringType()
	fieldNameMap["user_name"] = "UserName"
	fields["id"] = vapiBindings_.NewStringType()
	fieldNameMap["id"] = "Id"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.model.abstract_entity", fields, reflect.TypeOf(AbstractEntity{}), fieldNameMap, validators)
}

func AwsEventBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["instance_id"] = vapiBindings_.NewStringType()
	fieldNameMap["instance_id"] = "InstanceId"
	fields["start_time"] = vapiBindings_.NewDateTimeType()
	fieldNameMap["start_time"] = "StartTime"
	fields["type"] = vapiBindings_.NewStringType()
	fieldNameMap["type"] = "Type_"
	fields["account_id"] = vapiBindings_.NewStringType()
	fieldNameMap["account_id"] = "AccountId"
	fields["description"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["description"] = "Description"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.model.aws_event", fields, reflect.TypeOf(AwsEvent{}), fieldNameMap, validators)
}

func EdrsClusterInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["status_key"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["status_key"] = "StatusKey"
	fields["cluster_id"] = vapiBindings_.NewStringType()
	fieldNameMap["cluster_id"] = "ClusterId"
	fields["status_message"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["status_message"] = "StatusMessage"
	fields["edrs_policy"] = vapiBindings_.NewReferenceType(EdrsPolicyBindingType)
	fieldNameMap["edrs_policy"] = "EdrsPolicy"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.model.edrs_cluster_info", fields, reflect.TypeOf(EdrsClusterInfo{}), fieldNameMap, validators)
}

func EdrsPolicyBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["enable_edrs"] = vapiBindings_.NewBooleanType()
	fieldNameMap["enable_edrs"] = "EnableEdrs"
	fields["edrs_policy_options_overrides"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(EdrsPolicyOptionsOverridesBindingType))
	fieldNameMap["edrs_policy_options_overrides"] = "EdrsPolicyOptionsOverrides"
	fields["min_hosts"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["min_hosts"] = "MinHosts"
	fields["policy_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["policy_type"] = "PolicyType"
	fields["max_hosts"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["max_hosts"] = "MaxHosts"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.model.edrs_policy", fields, reflect.TypeOf(EdrsPolicy{}), fieldNameMap, validators)
}

func EdrsPolicyOptionsOverridesBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["scale_up_host_increment"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["scale_up_host_increment"] = "ScaleUpHostIncrement"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.model.edrs_policy_options_overrides", fields, reflect.TypeOf(EdrsPolicyOptionsOverrides{}), fieldNameMap, validators)
}

func EdrsPolicySpecBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["policy_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["policy_type"] = "PolicyType"
	fields["cluster_eligible_for_policy"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["cluster_eligible_for_policy"] = "ClusterEligibleForPolicy"
	fields["configurable_scaleup_increment"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["configurable_scaleup_increment"] = "ConfigurableScaleupIncrement"
	fields["min_max_host_range"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewIntegerType(), reflect.TypeOf([]int64{})))
	fieldNameMap["min_max_host_range"] = "MinMaxHostRange"
	fields["scaleup_host_increment_range"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewIntegerType(), reflect.TypeOf([]int64{})))
	fieldNameMap["scaleup_host_increment_range"] = "ScaleupHostIncrementRange"
	fields["configurable"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["configurable"] = "Configurable"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.model.edrs_policy_spec", fields, reflect.TypeOf(EdrsPolicySpec{}), fieldNameMap, validators)
}

func EdrsProvisioningSpecBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["edrs_cluster_info"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(EdrsClusterInfoBindingType))
	fieldNameMap["edrs_cluster_info"] = "EdrsClusterInfo"
	fields["policies"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(EdrsPolicySpecBindingType), reflect.TypeOf([]EdrsPolicySpec{})))
	fieldNameMap["policies"] = "Policies"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.model.edrs_provisioning_spec", fields, reflect.TypeOf(EdrsProvisioningSpec{}), fieldNameMap, validators)
}

func ErrorResponseBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["status"] = vapiBindings_.NewIntegerType()
	fieldNameMap["status"] = "Status"
	fields["path"] = vapiBindings_.NewStringType()
	fieldNameMap["path"] = "Path"
	fields["retryable"] = vapiBindings_.NewBooleanType()
	fieldNameMap["retryable"] = "Retryable"
	fields["error_code"] = vapiBindings_.NewStringType()
	fieldNameMap["error_code"] = "ErrorCode"
	fields["error_messages"] = vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["error_messages"] = "ErrorMessages"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.model.error_response", fields, reflect.TypeOf(ErrorResponse{}), fieldNameMap, validators)
}

func RemediationClusterInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster_name"] = vapiBindings_.NewStringType()
	fieldNameMap["cluster_name"] = "ClusterName"
	fields["cluster_id"] = vapiBindings_.NewStringType()
	fieldNameMap["cluster_id"] = "ClusterId"
	fields["remediation_policy"] = vapiBindings_.NewReferenceType(RemediationPolicyBindingType)
	fieldNameMap["remediation_policy"] = "RemediationPolicy"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.model.remediation_cluster_info", fields, reflect.TypeOf(RemediationClusterInfo{}), fieldNameMap, validators)
}

func RemediationPolicyBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["auto_remediation_enabled"] = vapiBindings_.NewBooleanType()
	fieldNameMap["auto_remediation_enabled"] = "AutoRemediationEnabled"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.model.remediation_policy", fields, reflect.TypeOf(RemediationPolicy{}), fieldNameMap, validators)
}

func RemediationPolicyPatchBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["auto_remediation_enabled"] = vapiBindings_.NewBooleanType()
	fieldNameMap["auto_remediation_enabled"] = "AutoRemediationEnabled"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.model.remediation_policy_patch", fields, reflect.TypeOf(RemediationPolicyPatch{}), fieldNameMap, validators)
}

func ServiceErrorBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["original_service"] = vapiBindings_.NewStringType()
	fieldNameMap["original_service"] = "OriginalService"
	fields["params"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["params"] = "Params"
	fields["default_message"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["default_message"] = "DefaultMessage"
	fields["original_service_error_code"] = vapiBindings_.NewStringType()
	fieldNameMap["original_service_error_code"] = "OriginalServiceErrorCode"
	fields["error_code"] = vapiBindings_.NewStringType()
	fieldNameMap["error_code"] = "ErrorCode"
	fields["localized_message"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["localized_message"] = "LocalizedMessage"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.model.service_error", fields, reflect.TypeOf(ServiceError{}), fieldNameMap, validators)
}

func TaskBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["updated"] = vapiBindings_.NewDateTimeType()
	fieldNameMap["updated"] = "Updated"
	fields["user_id"] = vapiBindings_.NewStringType()
	fieldNameMap["user_id"] = "UserId"
	fields["updated_by_user_id"] = vapiBindings_.NewStringType()
	fieldNameMap["updated_by_user_id"] = "UpdatedByUserId"
	fields["updated_by_user_name"] = vapiBindings_.NewStringType()
	fieldNameMap["updated_by_user_name"] = "UpdatedByUserName"
	fields["created"] = vapiBindings_.NewDateTimeType()
	fieldNameMap["created"] = "Created"
	fields["user_name"] = vapiBindings_.NewStringType()
	fieldNameMap["user_name"] = "UserName"
	fields["id"] = vapiBindings_.NewStringType()
	fieldNameMap["id"] = "Id"
	fields["status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["status"] = "Status"
	fields["resource_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["resource_id"] = "ResourceId"
	fields["start_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewDateTimeType())
	fieldNameMap["start_time"] = "StartTime"
	fields["service_errors"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ServiceErrorBindingType), reflect.TypeOf([]ServiceError{})))
	fieldNameMap["service_errors"] = "ServiceErrors"
	fields["sub_status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sub_status"] = "SubStatus"
	fields["task_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["task_type"] = "TaskType"
	fields["task_progress_phases"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(TaskProgressPhaseBindingType), reflect.TypeOf([]TaskProgressPhase{})))
	fieldNameMap["task_progress_phases"] = "TaskProgressPhases"
	fields["error_message"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["error_message"] = "ErrorMessage"
	fields["org_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fields["progress_percent"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["progress_percent"] = "ProgressPercent"
	fields["estimated_remaining_minutes"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["estimated_remaining_minutes"] = "EstimatedRemainingMinutes"
	fields["params"] = vapiBindings_.NewOptionalType(vapiBindings_.NewDynamicStructType(nil))
	fieldNameMap["params"] = "Params"
	fields["end_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewDateTimeType())
	fieldNameMap["end_time"] = "EndTime"
	fields["phase_in_progress"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["phase_in_progress"] = "PhaseInProgress"
	fields["task_version"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["task_version"] = "TaskVersion"
	fields["resource_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.model.task", fields, reflect.TypeOf(Task{}), fieldNameMap, validators)
}

func TaskProgressPhaseBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewStringType()
	fieldNameMap["id"] = "Id"
	fields["name"] = vapiBindings_.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["progress_percent"] = vapiBindings_.NewIntegerType()
	fieldNameMap["progress_percent"] = "ProgressPercent"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.model.task_progress_phase", fields, reflect.TypeOf(TaskProgressPhase{}), fieldNameMap, validators)
}
