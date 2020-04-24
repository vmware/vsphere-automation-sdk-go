/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for package: com.vmware.vmc.draas.model.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package model

import (
	"reflect"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"time"
)


type AbstractEntity struct {
	Updated time.Time
    // User id that last updated this record
	UserId string
	Created time.Time
    // User id that last updated this record
	UpdatedByUserId string
    // Version of this entity format: int32
	Version int64
    // User name that last updated this record
	UpdatedByUserName string
    // User name that last updated this record
	UserName string
	Id string
}

type ActivateSiteRecoveryConfig struct {
    // Optional custom extension key suffix for SRM. If not specified, default extension key will be used. The custom extension suffix must contain 13 characters or less, be composed of letters, numbers, ., -, and _ characters. The extension suffix must begin and end with a letter or number. The suffix is appended to com.vmware.vcDr- to form the full extension key.
	SrmExtensionKeySuffix *string
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

type ProvisionSrmConfig struct {
    // Optional custom extension key suffix for SRM. If not specified, default extension key will be used.
	SrmExtensionKeySuffix *string
}

type ReplicaDisk struct {
	SpaceRequirement *float64
	Name *string
	CollectionId *string
	DatastoresForSingleHostMove []*data.StructValue
	Movable *bool
	DiskId *string
	DatastoreMoId *string
}

type ReplicaDiskCollection struct {
	CollectionId *string
	Generated *time.Time
	Disks []ReplicaDisk
	PlaceholderVmMoId *string
	Name *string
}

type SiteRecovery struct {
	Updated time.Time
    // User id that last updated this record
	UserId string
	Created time.Time
    // User id that last updated this record
	UpdatedByUserId string
    // Version of this entity format: int32
	Version int64
    // User name that last updated this record
	UpdatedByUserName string
    // User name that last updated this record
	UserName string
	Id string
    // Possible values are: 
    //
    // * SiteRecovery#SiteRecovery_SITE_RECOVERY_STATE_ACTIVATING
    // * SiteRecovery#SiteRecovery_SITE_RECOVERY_STATE_ACTIVATED
    // * SiteRecovery#SiteRecovery_SITE_RECOVERY_STATE_DEACTIVATING
    // * SiteRecovery#SiteRecovery_SITE_RECOVERY_STATE_DEACTIVATED
    // * SiteRecovery#SiteRecovery_SITE_RECOVERY_STATE_FAILED
    // * SiteRecovery#SiteRecovery_SITE_RECOVERY_STATE_CANCELED
    // * SiteRecovery#SiteRecovery_SITE_RECOVERY_STATE_DELETED
	SiteRecoveryState *string
	VrNode *SiteRecoveryNode
	SrmNodes []SiteRecoveryNode
	SddcId *string
	DraasH5Url *string
}
const SiteRecovery_SITE_RECOVERY_STATE_ACTIVATING = "ACTIVATING"
const SiteRecovery_SITE_RECOVERY_STATE_ACTIVATED = "ACTIVATED"
const SiteRecovery_SITE_RECOVERY_STATE_DEACTIVATING = "DEACTIVATING"
const SiteRecovery_SITE_RECOVERY_STATE_DEACTIVATED = "DEACTIVATED"
const SiteRecovery_SITE_RECOVERY_STATE_FAILED = "FAILED"
const SiteRecovery_SITE_RECOVERY_STATE_CANCELED = "CANCELED"
const SiteRecovery_SITE_RECOVERY_STATE_DELETED = "DELETED"

type SiteRecoveryNode struct {
	VmMorefId *string
	IpAddress *string
	Hostname *string
	Id *string
    // Possible values are: 
    //
    // * SiteRecoveryNode#SiteRecoveryNode_STATE_DEPLOYING
    // * SiteRecoveryNode#SiteRecoveryNode_STATE_PROVISIONED
    // * SiteRecoveryNode#SiteRecoveryNode_STATE_READY
    // * SiteRecoveryNode#SiteRecoveryNode_STATE_DELETING
    // * SiteRecoveryNode#SiteRecoveryNode_STATE_FAILED
    // * SiteRecoveryNode#SiteRecoveryNode_STATE_CANCELED
	State *string
    // Possible values are: 
    //
    // * SiteRecoveryNode#SiteRecoveryNode_TYPE_VRMS
    // * SiteRecoveryNode#SiteRecoveryNode_TYPE_SRM
    // * SiteRecoveryNode#SiteRecoveryNode_TYPE_VRS
	Type_ *string
}
const SiteRecoveryNode_STATE_DEPLOYING = "DEPLOYING"
const SiteRecoveryNode_STATE_PROVISIONED = "PROVISIONED"
const SiteRecoveryNode_STATE_READY = "READY"
const SiteRecoveryNode_STATE_DELETING = "DELETING"
const SiteRecoveryNode_STATE_FAILED = "FAILED"
const SiteRecoveryNode_STATE_CANCELED = "CANCELED"
const SiteRecoveryNode_TYPE_VRMS = "VRMS"
const SiteRecoveryNode_TYPE_SRM = "SRM"
const SiteRecoveryNode_TYPE_VRS = "VRS"

type SiteRecoveryNodeVersion struct {
    // Possible values are: 
    //
    // * SiteRecoveryNodeVersion#SiteRecoveryNodeVersion_VERSION_SOURCE_VAMICLI
    // * SiteRecoveryNodeVersion#SiteRecoveryNodeVersion_VERSION_SOURCE_LS
	VersionSource *string
	NodeId *string
	BuildNumber *string
	ApplianceVersion *string
	NodeIp *string
	FullVersion *string
    // Possible values are: 
    //
    // * SiteRecoveryNodeVersion#SiteRecoveryNodeVersion_NODE_TYPE_VRMS
    // * SiteRecoveryNodeVersion#SiteRecoveryNodeVersion_NODE_TYPE_SRM
    // * SiteRecoveryNodeVersion#SiteRecoveryNodeVersion_NODE_TYPE_VRS
	NodeType *string
}
const SiteRecoveryNodeVersion_VERSION_SOURCE_VAMICLI = "vamicli"
const SiteRecoveryNodeVersion_VERSION_SOURCE_LS = "ls"
const SiteRecoveryNodeVersion_NODE_TYPE_VRMS = "VRMS"
const SiteRecoveryNodeVersion_NODE_TYPE_SRM = "SRM"
const SiteRecoveryNodeVersion_NODE_TYPE_VRS = "VRS"

type SiteRecoveryVersions struct {
	Generated *time.Time
	SddcId *string
    // list of site recovery node version
	NodeVersions []SiteRecoveryNodeVersion
}

type Task struct {
	Updated time.Time
    // User id that last updated this record
	UserId string
	Created time.Time
    // User id that last updated this record
	UpdatedByUserId string
    // Version of this entity format: int32
	Version int64
    // User name that last updated this record
	UpdatedByUserName string
    // User name that last updated this record
	UserName string
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
	StartTime *string
	Retries *int64
	TaskType *string
    // Task progress phases involved in current task execution
	TaskProgressPhases []TaskProgressPhase
	TenantId *string
	ErrorMessage *string
	ParentTaskId *string
    // Estimated progress percentage the task executed format: int32
	ProgressPercent *int64
    // Estimated remaining time in minute of the task execution, < 0 means no estimation for the task. format: int32
	EstimatedRemainingMinutes *int64
	Params *data.StructValue
	EndTime *time.Time
	TaskVersion *string
    // Type of resource being acted upon
	ResourceType *string
	SubStatus *string
}
const Task_STATUS_STARTED = "STARTED"
const Task_STATUS_CANCELING = "CANCELING"
const Task_STATUS_FINISHED = "FINISHED"
const Task_STATUS_FAILED = "FAILED"
const Task_STATUS_CANCELED = "CANCELED"

// A task progress can be (but does NOT have to be) divided to more meaningful progress phases.
type TaskProgressPhase struct {
    // The identifier of the task progress phase
	Id string
    // The display name of the task progress phase
	Name string
    // The percentage of the phase that has completed format: int32
	ProgressPercent int64
}




func AbstractEntityBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["updated"] = bindings.NewDateTimeType()
	fieldNameMap["updated"] = "Updated"
	fields["user_id"] = bindings.NewStringType()
	fieldNameMap["user_id"] = "UserId"
	fields["created"] = bindings.NewDateTimeType()
	fieldNameMap["created"] = "Created"
	fields["updated_by_user_id"] = bindings.NewStringType()
	fieldNameMap["updated_by_user_id"] = "UpdatedByUserId"
	fields["version"] = bindings.NewIntegerType()
	fieldNameMap["version"] = "Version"
	fields["updated_by_user_name"] = bindings.NewStringType()
	fieldNameMap["updated_by_user_name"] = "UpdatedByUserName"
	fields["user_name"] = bindings.NewStringType()
	fieldNameMap["user_name"] = "UserName"
	fields["id"] = bindings.NewStringType()
	fieldNameMap["id"] = "Id"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.draas.model.abstract_entity", fields, reflect.TypeOf(AbstractEntity{}), fieldNameMap, validators)
}

func ActivateSiteRecoveryConfigBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["srm_extension_key_suffix"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["srm_extension_key_suffix"] = "SrmExtensionKeySuffix"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.draas.model.activate_site_recovery_config", fields, reflect.TypeOf(ActivateSiteRecoveryConfig{}), fieldNameMap, validators)
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
	return bindings.NewStructType("com.vmware.vmc.draas.model.error_response", fields, reflect.TypeOf(ErrorResponse{}), fieldNameMap, validators)
}

func ProvisionSrmConfigBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["srm_extension_key_suffix"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["srm_extension_key_suffix"] = "SrmExtensionKeySuffix"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.draas.model.provision_srm_config", fields, reflect.TypeOf(ProvisionSrmConfig{}), fieldNameMap, validators)
}

func ReplicaDiskBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["space_requirement"] = bindings.NewOptionalType(bindings.NewDoubleType())
	fieldNameMap["space_requirement"] = "SpaceRequirement"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["collection_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["collection_id"] = "CollectionId"
	fields["datastores_for_single_host_move"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewDynamicStructType(nil, bindings.JSONRPC), reflect.TypeOf([]*data.StructValue{})))
	fieldNameMap["datastores_for_single_host_move"] = "DatastoresForSingleHostMove"
	fields["movable"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["movable"] = "Movable"
	fields["disk_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["disk_id"] = "DiskId"
	fields["datastore_mo_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["datastore_mo_id"] = "DatastoreMoId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.draas.model.replica_disk", fields, reflect.TypeOf(ReplicaDisk{}), fieldNameMap, validators)
}

func ReplicaDiskCollectionBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["collection_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["collection_id"] = "CollectionId"
	fields["generated"] = bindings.NewOptionalType(bindings.NewDateTimeType())
	fieldNameMap["generated"] = "Generated"
	fields["disks"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ReplicaDiskBindingType), reflect.TypeOf([]ReplicaDisk{})))
	fieldNameMap["disks"] = "Disks"
	fields["placeholder_vm_mo_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["placeholder_vm_mo_id"] = "PlaceholderVmMoId"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.draas.model.replica_disk_collection", fields, reflect.TypeOf(ReplicaDiskCollection{}), fieldNameMap, validators)
}

func SiteRecoveryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["updated"] = bindings.NewDateTimeType()
	fieldNameMap["updated"] = "Updated"
	fields["user_id"] = bindings.NewStringType()
	fieldNameMap["user_id"] = "UserId"
	fields["created"] = bindings.NewDateTimeType()
	fieldNameMap["created"] = "Created"
	fields["updated_by_user_id"] = bindings.NewStringType()
	fieldNameMap["updated_by_user_id"] = "UpdatedByUserId"
	fields["version"] = bindings.NewIntegerType()
	fieldNameMap["version"] = "Version"
	fields["updated_by_user_name"] = bindings.NewStringType()
	fieldNameMap["updated_by_user_name"] = "UpdatedByUserName"
	fields["user_name"] = bindings.NewStringType()
	fieldNameMap["user_name"] = "UserName"
	fields["id"] = bindings.NewStringType()
	fieldNameMap["id"] = "Id"
	fields["site_recovery_state"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["site_recovery_state"] = "SiteRecoveryState"
	fields["vr_node"] = bindings.NewOptionalType(bindings.NewReferenceType(SiteRecoveryNodeBindingType))
	fieldNameMap["vr_node"] = "VrNode"
	fields["srm_nodes"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(SiteRecoveryNodeBindingType), reflect.TypeOf([]SiteRecoveryNode{})))
	fieldNameMap["srm_nodes"] = "SrmNodes"
	fields["sddc_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sddc_id"] = "SddcId"
	fields["draas_h5_url"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["draas_h5_url"] = "DraasH5Url"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.draas.model.site_recovery", fields, reflect.TypeOf(SiteRecovery{}), fieldNameMap, validators)
}

func SiteRecoveryNodeBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm_moref_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vm_moref_id"] = "VmMorefId"
	fields["ip_address"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["ip_address"] = "IpAddress"
	fields["hostname"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["hostname"] = "Hostname"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["state"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["state"] = "State"
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["type"] = "Type_"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.draas.model.site_recovery_node", fields, reflect.TypeOf(SiteRecoveryNode{}), fieldNameMap, validators)
}

func SiteRecoveryNodeVersionBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["version_source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["version_source"] = "VersionSource"
	fields["node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["node_id"] = "NodeId"
	fields["build_number"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["build_number"] = "BuildNumber"
	fields["appliance_version"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["appliance_version"] = "ApplianceVersion"
	fields["node_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["node_ip"] = "NodeIp"
	fields["full_version"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["full_version"] = "FullVersion"
	fields["node_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["node_type"] = "NodeType"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.draas.model.site_recovery_node_version", fields, reflect.TypeOf(SiteRecoveryNodeVersion{}), fieldNameMap, validators)
}

func SiteRecoveryVersionsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["generated"] = bindings.NewOptionalType(bindings.NewDateTimeType())
	fieldNameMap["generated"] = "Generated"
	fields["sddc_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sddc_id"] = "SddcId"
	fields["node_versions"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(SiteRecoveryNodeVersionBindingType), reflect.TypeOf([]SiteRecoveryNodeVersion{})))
	fieldNameMap["node_versions"] = "NodeVersions"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.draas.model.site_recovery_versions", fields, reflect.TypeOf(SiteRecoveryVersions{}), fieldNameMap, validators)
}

func TaskBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["updated"] = bindings.NewDateTimeType()
	fieldNameMap["updated"] = "Updated"
	fields["user_id"] = bindings.NewStringType()
	fieldNameMap["user_id"] = "UserId"
	fields["created"] = bindings.NewDateTimeType()
	fieldNameMap["created"] = "Created"
	fields["updated_by_user_id"] = bindings.NewStringType()
	fieldNameMap["updated_by_user_id"] = "UpdatedByUserId"
	fields["version"] = bindings.NewIntegerType()
	fieldNameMap["version"] = "Version"
	fields["updated_by_user_name"] = bindings.NewStringType()
	fieldNameMap["updated_by_user_name"] = "UpdatedByUserName"
	fields["user_name"] = bindings.NewStringType()
	fieldNameMap["user_name"] = "UserName"
	fields["id"] = bindings.NewStringType()
	fieldNameMap["id"] = "Id"
	fields["status"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["status"] = "Status"
	fields["resource_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["resource_id"] = "ResourceId"
	fields["start_time"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["start_time"] = "StartTime"
	fields["retries"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["retries"] = "Retries"
	fields["task_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["task_type"] = "TaskType"
	fields["task_progress_phases"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(TaskProgressPhaseBindingType), reflect.TypeOf([]TaskProgressPhase{})))
	fieldNameMap["task_progress_phases"] = "TaskProgressPhases"
	fields["tenant_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["tenant_id"] = "TenantId"
	fields["error_message"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["error_message"] = "ErrorMessage"
	fields["parent_task_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["parent_task_id"] = "ParentTaskId"
	fields["progress_percent"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["progress_percent"] = "ProgressPercent"
	fields["estimated_remaining_minutes"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["estimated_remaining_minutes"] = "EstimatedRemainingMinutes"
	fields["params"] = bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
	fieldNameMap["params"] = "Params"
	fields["end_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
	fieldNameMap["end_time"] = "EndTime"
	fields["task_version"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["task_version"] = "TaskVersion"
	fields["resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["sub_status"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sub_status"] = "SubStatus"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.draas.model.task", fields, reflect.TypeOf(Task{}), fieldNameMap, validators)
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
	return bindings.NewStructType("com.vmware.vmc.draas.model.task_progress_phase", fields, reflect.TypeOf(TaskProgressPhase{}), fieldNameMap, validators)
}


