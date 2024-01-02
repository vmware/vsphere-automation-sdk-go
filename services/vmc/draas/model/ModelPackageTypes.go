// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for package: com.vmware.vmc.draas.model.
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
	Id string
	// Version of this entity format: int32
	Version int64
	Created time.Time
	// User id that last updated this record
	UserId string
	// User name that last updated this record
	UserName string
	Updated  time.Time
	// User id that last updated this record
	UpdatedByUserId string
	// User name that last updated this record
	UpdatedByUserName string
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

type ActivateSiteRecoveryConfig struct {
	// Optional custom extension key suffix for SRM. If not specified, default extension key will be used. The custom extension suffix must contain 13 characters or less, be composed of letters, numbers, ., -, and _ characters. The extension suffix must begin and end with a letter or number. The suffix is appended to com.vmware.vcDr- to form the full extension key.
	SrmExtensionKeySuffix *string
}

func (s *ActivateSiteRecoveryConfig) GetType__() vapiBindings_.BindingType {
	return ActivateSiteRecoveryConfigBindingType()
}

func (s *ActivateSiteRecoveryConfig) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ActivateSiteRecoveryConfig._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// The build version.
type BuildVersion struct {
	Version     *string
	BuildNumber *string
}

func (s *BuildVersion) GetType__() vapiBindings_.BindingType {
	return BuildVersionBindingType()
}

func (s *BuildVersion) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for BuildVersion._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type DatastoreEntity struct {
	Id                 *string
	Name               *string
	DatastoreType      *string
	InstanceType       *string
	NumberOfHosts      *float64
	ReplicationsNumber *float64
	TotalDatastoreSize *float64
	ReplicaDisks       []ReplicaDisk
}

func (s *DatastoreEntity) GetType__() vapiBindings_.BindingType {
	return DatastoreEntityBindingType()
}

func (s *DatastoreEntity) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for DatastoreEntity._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Deactivate Site Recovery or delete Site Recovery node configuration
type DeleteConfigInternal struct {
	// If = 'true', will deactivate site recovery forcefully.
	Force *bool
	// CSSD/CSCM ticket number, such as CSSD-xxxx/CSCM-xxxx
	Ticket *string
	// The confirmation code generated in prepare delete operation.
	ConfirmationCode *string
}

func (s *DeleteConfigInternal) GetType__() vapiBindings_.BindingType {
	return DeleteConfigInternalBindingType()
}

func (s *DeleteConfigInternal) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for DeleteConfigInternal._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type ErrorResponse struct {
	// unique error code
	ErrorCode string
	// localized error messages
	ErrorMessages []string
	// HTTP status code
	Status int64
	// Originating request URI
	Path string
	// If true, client should retry operation
	Retryable bool
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

type Fault struct {
	Message *string
	Cause   *vapiData_.StructValue
}

func (s *Fault) GetType__() vapiBindings_.BindingType {
	return FaultBindingType()
}

func (s *Fault) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for Fault._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type HmsIssueInfo struct {
	TargetObjectName *string
	TargetObjectMoId string
	TriggeredTime    time.Time
	IssueType        string
	Fault            *Fault
	Severity         string
}

func (s *HmsIssueInfo) GetType__() vapiBindings_.BindingType {
	return HmsIssueInfoBindingType()
}

func (s *HmsIssueInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for HmsIssueInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type HmsReplicationIssueInfo struct {
	TargetObjectName    *string
	TargetObjectMoId    string
	TriggeredTime       time.Time
	IssueType           string
	Fault               *Fault
	Severity            string
	SourceSiteUuid      *string
	DestinationSiteUuid *string
	Direction           *string
}

func (s *HmsReplicationIssueInfo) GetType__() vapiBindings_.BindingType {
	return HmsReplicationIssueInfoBindingType()
}

func (s *HmsReplicationIssueInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for HmsReplicationIssueInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type HmsSiteIssueInfo struct {
	TargetObjectName *string
	TargetObjectMoId string
	TriggeredTime    time.Time
	IssueType        string
	Fault            *Fault
	Severity         string
}

func (s *HmsSiteIssueInfo) GetType__() vapiBindings_.BindingType {
	return HmsSiteIssueInfoBindingType()
}

func (s *HmsSiteIssueInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for HmsSiteIssueInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type ProvisionSrmConfig struct {
	// Optional custom extension key suffix for SRM. If not specified, default extension key will be used.
	SrmExtensionKeySuffix *string
}

func (s *ProvisionSrmConfig) GetType__() vapiBindings_.BindingType {
	return ProvisionSrmConfigBindingType()
}

func (s *ProvisionSrmConfig) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ProvisionSrmConfig._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type ReplicaData struct {
	// DatastoreEntity
	DatastoreEntities []DatastoreEntity
}

func (s *ReplicaData) GetType__() vapiBindings_.BindingType {
	return ReplicaDataBindingType()
}

func (s *ReplicaData) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ReplicaData._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type ReplicaDisk struct {
	DiskId                      *string
	CollectionId                *string
	Name                        *string
	DatastoreMoId               *string
	SpaceRequirement            *float64
	Movable                     *bool
	DatastoresForSingleHostMove []*vapiData_.StructValue
}

func (s *ReplicaDisk) GetType__() vapiBindings_.BindingType {
	return ReplicaDiskBindingType()
}

func (s *ReplicaDisk) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ReplicaDisk._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type ReplicaDiskCollection struct {
	CollectionId      *string
	Name              *string
	PlaceholderVmMoId *string
	Disks             []ReplicaDisk
	Generated         *time.Time
}

func (s *ReplicaDiskCollection) GetType__() vapiBindings_.BindingType {
	return ReplicaDiskCollectionBindingType()
}

func (s *ReplicaDiskCollection) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ReplicaDiskCollection._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type ReplicationData struct {
	ReplicationName *string
	ReplicationUuid *string
	HbrServerUuid   *string
	HbrServerName   *string
	ScaleOut        *bool
}

func (s *ReplicationData) GetType__() vapiBindings_.BindingType {
	return ReplicationDataBindingType()
}

func (s *ReplicationData) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ReplicationData._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type ReplicationsCount struct {
	Incoming        *string
	ReplicationUuid *int64
	Outgoing        *int64
}

func (s *ReplicationsCount) GetType__() vapiBindings_.BindingType {
	return ReplicationsCountBindingType()
}

func (s *ReplicationsCount) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ReplicationsCount._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type SiteRecovery struct {
	Id string
	// Version of this entity format: int32
	Version int64
	Created time.Time
	// User id that last updated this record
	UserId string
	// User name that last updated this record
	UserName string
	Updated  time.Time
	// User id that last updated this record
	UpdatedByUserId string
	// User name that last updated this record
	UpdatedByUserName string
	SddcId            *string
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
	DraasH5Url        *string
	VrNode            *SiteRecoveryNode
	SrmNodes          []SrmNode
	VrsNodes          []SiteRecoveryNode
}

const SiteRecovery_SITE_RECOVERY_STATE_ACTIVATING = "ACTIVATING"
const SiteRecovery_SITE_RECOVERY_STATE_ACTIVATED = "ACTIVATED"
const SiteRecovery_SITE_RECOVERY_STATE_DEACTIVATING = "DEACTIVATING"
const SiteRecovery_SITE_RECOVERY_STATE_DEACTIVATED = "DEACTIVATED"
const SiteRecovery_SITE_RECOVERY_STATE_FAILED = "FAILED"
const SiteRecovery_SITE_RECOVERY_STATE_CANCELED = "CANCELED"
const SiteRecovery_SITE_RECOVERY_STATE_DELETED = "DELETED"

func (s *SiteRecovery) GetType__() vapiBindings_.BindingType {
	return SiteRecoveryBindingType()
}

func (s *SiteRecovery) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for SiteRecovery._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type SiteRecoveryNode struct {
	Id *string
	// Possible values are:
	//
	// * SiteRecoveryNode#SiteRecoveryNode_TYPE_VRMS
	// * SiteRecoveryNode#SiteRecoveryNode_TYPE_SRM
	// * SiteRecoveryNode#SiteRecoveryNode_TYPE_VRS
	Type_ *string
	// Possible values are:
	//
	// * SiteRecoveryNode#SiteRecoveryNode_STATE_DEPLOYING
	// * SiteRecoveryNode#SiteRecoveryNode_STATE_PROVISIONED
	// * SiteRecoveryNode#SiteRecoveryNode_STATE_READY
	// * SiteRecoveryNode#SiteRecoveryNode_STATE_DELETING
	// * SiteRecoveryNode#SiteRecoveryNode_STATE_FAILED
	// * SiteRecoveryNode#SiteRecoveryNode_STATE_CANCELED
	State     *string
	VmMorefId *string
	IpAddress *string
	Hostname  *string
}

const SiteRecoveryNode_TYPE_VRMS = "VRMS"
const SiteRecoveryNode_TYPE_SRM = "SRM"
const SiteRecoveryNode_TYPE_VRS = "VRS"
const SiteRecoveryNode_STATE_DEPLOYING = "DEPLOYING"
const SiteRecoveryNode_STATE_PROVISIONED = "PROVISIONED"
const SiteRecoveryNode_STATE_READY = "READY"
const SiteRecoveryNode_STATE_DELETING = "DELETING"
const SiteRecoveryNode_STATE_FAILED = "FAILED"
const SiteRecoveryNode_STATE_CANCELED = "CANCELED"

func (s *SiteRecoveryNode) GetType__() vapiBindings_.BindingType {
	return SiteRecoveryNodeBindingType()
}

func (s *SiteRecoveryNode) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for SiteRecoveryNode._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type SiteRecoveryNodeVersion struct {
	NodeId *string
	// Possible values are:
	//
	// * SiteRecoveryNodeVersion#SiteRecoveryNodeVersion_NODE_TYPE_VRMS
	// * SiteRecoveryNodeVersion#SiteRecoveryNodeVersion_NODE_TYPE_SRM
	// * SiteRecoveryNodeVersion#SiteRecoveryNodeVersion_NODE_TYPE_VRS
	NodeType     *string
	NodeIp       *string
	BuildVersion *BuildVersion
	FullVersion  *string
}

const SiteRecoveryNodeVersion_NODE_TYPE_VRMS = "VRMS"
const SiteRecoveryNodeVersion_NODE_TYPE_SRM = "SRM"
const SiteRecoveryNodeVersion_NODE_TYPE_VRS = "VRS"

func (s *SiteRecoveryNodeVersion) GetType__() vapiBindings_.BindingType {
	return SiteRecoveryNodeVersionBindingType()
}

func (s *SiteRecoveryNodeVersion) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for SiteRecoveryNodeVersion._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type SiteRecoveryVersions struct {
	SddcId *string
	// list of site recovery node version
	NodeVersions []SiteRecoveryNodeVersion
	Generated    *time.Time
}

func (s *SiteRecoveryVersions) GetType__() vapiBindings_.BindingType {
	return SiteRecoveryVersionsBindingType()
}

func (s *SiteRecoveryVersions) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for SiteRecoveryVersions._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type SrmNode struct {
	Id *string
	// Possible values are:
	//
	// * SrmNode#SrmNode_TYPE_VRMS
	// * SrmNode#SrmNode_TYPE_SRM
	// * SrmNode#SrmNode_TYPE_VRS
	Type_ *string
	// Possible values are:
	//
	// * SrmNode#SrmNode_STATE_DEPLOYING
	// * SrmNode#SrmNode_STATE_PROVISIONED
	// * SrmNode#SrmNode_STATE_READY
	// * SrmNode#SrmNode_STATE_DELETING
	// * SrmNode#SrmNode_STATE_FAILED
	// * SrmNode#SrmNode_STATE_CANCELED
	State                 *string
	VmMorefId             *string
	IpAddress             *string
	Hostname              *string
	SrmExtensionKeySuffix *string
	SrmExtensionKey       *string
}

const SrmNode_TYPE_VRMS = "VRMS"
const SrmNode_TYPE_SRM = "SRM"
const SrmNode_TYPE_VRS = "VRS"
const SrmNode_STATE_DEPLOYING = "DEPLOYING"
const SrmNode_STATE_PROVISIONED = "PROVISIONED"
const SrmNode_STATE_READY = "READY"
const SrmNode_STATE_DELETING = "DELETING"
const SrmNode_STATE_FAILED = "FAILED"
const SrmNode_STATE_CANCELED = "CANCELED"

func (s *SrmNode) GetType__() vapiBindings_.BindingType {
	return SrmNodeBindingType()
}

func (s *SrmNode) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for SrmNode._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type Task struct {
	Id string
	// Version of this entity format: int32
	Version int64
	Created time.Time
	// User id that last updated this record
	UserId string
	// User name that last updated this record
	UserName string
	Updated  time.Time
	// User id that last updated this record
	UpdatedByUserId string
	// User name that last updated this record
	UpdatedByUserName string
	EndTime           *time.Time
	ErrorMessage      *string
	TenantId          *string
	Params            *vapiData_.StructValue
	// UUID of resources task is acting upon
	ResourceId *string
	// Type of resource being acted upon
	ResourceType *string
	ParentTaskId *string
	Retries      *int64
	StartTime    *time.Time
	// Possible values are:
	//
	// * Task#Task_STATUS_STARTED
	// * Task#Task_STATUS_CANCELING
	// * Task#Task_STATUS_FINISHED
	// * Task#Task_STATUS_FAILED
	// * Task#Task_STATUS_CANCELED
	Status      *string
	SubStatus   *string
	TaskType    *string
	TaskVersion *string
	// Task progress phases involved in current task execution
	TaskProgressPhases []TaskProgressPhase
	// Estimated progress percentage the task executed format: int32
	ProgressPercent *int64
	// Estimated remaining time in minute of the task execution, < 0 means no estimation for the task. format: int32
	EstimatedRemainingMinutes *int64
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
	fields["id"] = vapiBindings_.NewStringType()
	fieldNameMap["id"] = "Id"
	fields["version"] = vapiBindings_.NewIntegerType()
	fieldNameMap["version"] = "Version"
	fields["created"] = vapiBindings_.NewDateTimeType()
	fieldNameMap["created"] = "Created"
	fields["user_id"] = vapiBindings_.NewStringType()
	fieldNameMap["user_id"] = "UserId"
	fields["user_name"] = vapiBindings_.NewStringType()
	fieldNameMap["user_name"] = "UserName"
	fields["updated"] = vapiBindings_.NewDateTimeType()
	fieldNameMap["updated"] = "Updated"
	fields["updated_by_user_id"] = vapiBindings_.NewStringType()
	fieldNameMap["updated_by_user_id"] = "UpdatedByUserId"
	fields["updated_by_user_name"] = vapiBindings_.NewStringType()
	fieldNameMap["updated_by_user_name"] = "UpdatedByUserName"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.draas.model.abstract_entity", fields, reflect.TypeOf(AbstractEntity{}), fieldNameMap, validators)
}

func ActivateSiteRecoveryConfigBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["srm_extension_key_suffix"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["srm_extension_key_suffix"] = "SrmExtensionKeySuffix"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.draas.model.activate_site_recovery_config", fields, reflect.TypeOf(ActivateSiteRecoveryConfig{}), fieldNameMap, validators)
}

func BuildVersionBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["version"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["version"] = "Version"
	fields["build_number"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["build_number"] = "BuildNumber"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.draas.model.build_version", fields, reflect.TypeOf(BuildVersion{}), fieldNameMap, validators)
}

func DatastoreEntityBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["datastore_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["datastore_type"] = "DatastoreType"
	fields["instance_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["instance_type"] = "InstanceType"
	fields["number_of_hosts"] = vapiBindings_.NewOptionalType(vapiBindings_.NewDoubleType())
	fieldNameMap["number_of_hosts"] = "NumberOfHosts"
	fields["replications_number"] = vapiBindings_.NewOptionalType(vapiBindings_.NewDoubleType())
	fieldNameMap["replications_number"] = "ReplicationsNumber"
	fields["total_datastore_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewDoubleType())
	fieldNameMap["total_datastore_size"] = "TotalDatastoreSize"
	fields["replica_disks"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ReplicaDiskBindingType), reflect.TypeOf([]ReplicaDisk{})))
	fieldNameMap["replica_disks"] = "ReplicaDisks"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.draas.model.datastore_entity", fields, reflect.TypeOf(DatastoreEntity{}), fieldNameMap, validators)
}

func DeleteConfigInternalBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["force"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["force"] = "Force"
	fields["ticket"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["ticket"] = "Ticket"
	fields["confirmation_code"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["confirmation_code"] = "ConfirmationCode"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.draas.model.delete_config_internal", fields, reflect.TypeOf(DeleteConfigInternal{}), fieldNameMap, validators)
}

func ErrorResponseBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["error_code"] = vapiBindings_.NewStringType()
	fieldNameMap["error_code"] = "ErrorCode"
	fields["error_messages"] = vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["error_messages"] = "ErrorMessages"
	fields["status"] = vapiBindings_.NewIntegerType()
	fieldNameMap["status"] = "Status"
	fields["path"] = vapiBindings_.NewStringType()
	fieldNameMap["path"] = "Path"
	fields["retryable"] = vapiBindings_.NewBooleanType()
	fieldNameMap["retryable"] = "Retryable"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.draas.model.error_response", fields, reflect.TypeOf(ErrorResponse{}), fieldNameMap, validators)
}

func FaultBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["message"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["message"] = "Message"
	fields["cause"] = vapiBindings_.NewOptionalType(vapiBindings_.NewDynamicStructType(nil))
	fieldNameMap["cause"] = "Cause"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.draas.model.fault", fields, reflect.TypeOf(Fault{}), fieldNameMap, validators)
}

func HmsIssueInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["target_object_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["target_object_name"] = "TargetObjectName"
	fields["target_object_mo_id"] = vapiBindings_.NewStringType()
	fieldNameMap["target_object_mo_id"] = "TargetObjectMoId"
	fields["triggered_time"] = vapiBindings_.NewDateTimeType()
	fieldNameMap["triggered_time"] = "TriggeredTime"
	fields["issue_type"] = vapiBindings_.NewStringType()
	fieldNameMap["issue_type"] = "IssueType"
	fields["fault"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(FaultBindingType))
	fieldNameMap["fault"] = "Fault"
	fields["severity"] = vapiBindings_.NewStringType()
	fieldNameMap["severity"] = "Severity"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.draas.model.hms_issue_info", fields, reflect.TypeOf(HmsIssueInfo{}), fieldNameMap, validators)
}

func HmsReplicationIssueInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["target_object_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["target_object_name"] = "TargetObjectName"
	fields["target_object_mo_id"] = vapiBindings_.NewStringType()
	fieldNameMap["target_object_mo_id"] = "TargetObjectMoId"
	fields["triggered_time"] = vapiBindings_.NewDateTimeType()
	fieldNameMap["triggered_time"] = "TriggeredTime"
	fields["issue_type"] = vapiBindings_.NewStringType()
	fieldNameMap["issue_type"] = "IssueType"
	fields["fault"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(FaultBindingType))
	fieldNameMap["fault"] = "Fault"
	fields["severity"] = vapiBindings_.NewStringType()
	fieldNameMap["severity"] = "Severity"
	fields["source_site_uuid"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["source_site_uuid"] = "SourceSiteUuid"
	fields["destination_site_uuid"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["destination_site_uuid"] = "DestinationSiteUuid"
	fields["direction"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["direction"] = "Direction"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.draas.model.hms_replication_issue_info", fields, reflect.TypeOf(HmsReplicationIssueInfo{}), fieldNameMap, validators)
}

func HmsSiteIssueInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["target_object_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["target_object_name"] = "TargetObjectName"
	fields["target_object_mo_id"] = vapiBindings_.NewStringType()
	fieldNameMap["target_object_mo_id"] = "TargetObjectMoId"
	fields["triggered_time"] = vapiBindings_.NewDateTimeType()
	fieldNameMap["triggered_time"] = "TriggeredTime"
	fields["issue_type"] = vapiBindings_.NewStringType()
	fieldNameMap["issue_type"] = "IssueType"
	fields["fault"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(FaultBindingType))
	fieldNameMap["fault"] = "Fault"
	fields["severity"] = vapiBindings_.NewStringType()
	fieldNameMap["severity"] = "Severity"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.draas.model.hms_site_issue_info", fields, reflect.TypeOf(HmsSiteIssueInfo{}), fieldNameMap, validators)
}

func ProvisionSrmConfigBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["srm_extension_key_suffix"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["srm_extension_key_suffix"] = "SrmExtensionKeySuffix"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.draas.model.provision_srm_config", fields, reflect.TypeOf(ProvisionSrmConfig{}), fieldNameMap, validators)
}

func ReplicaDataBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["datastore_entities"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(DatastoreEntityBindingType), reflect.TypeOf([]DatastoreEntity{})))
	fieldNameMap["datastore_entities"] = "DatastoreEntities"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.draas.model.replica_data", fields, reflect.TypeOf(ReplicaData{}), fieldNameMap, validators)
}

func ReplicaDiskBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["disk_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["disk_id"] = "DiskId"
	fields["collection_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["collection_id"] = "CollectionId"
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["datastore_mo_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["datastore_mo_id"] = "DatastoreMoId"
	fields["space_requirement"] = vapiBindings_.NewOptionalType(vapiBindings_.NewDoubleType())
	fieldNameMap["space_requirement"] = "SpaceRequirement"
	fields["movable"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["movable"] = "Movable"
	fields["datastores_for_single_host_move"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewDynamicStructType(nil), reflect.TypeOf([]*vapiData_.StructValue{})))
	fieldNameMap["datastores_for_single_host_move"] = "DatastoresForSingleHostMove"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.draas.model.replica_disk", fields, reflect.TypeOf(ReplicaDisk{}), fieldNameMap, validators)
}

func ReplicaDiskCollectionBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["collection_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["collection_id"] = "CollectionId"
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["placeholder_vm_mo_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["placeholder_vm_mo_id"] = "PlaceholderVmMoId"
	fields["disks"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ReplicaDiskBindingType), reflect.TypeOf([]ReplicaDisk{})))
	fieldNameMap["disks"] = "Disks"
	fields["generated"] = vapiBindings_.NewOptionalType(vapiBindings_.NewDateTimeType())
	fieldNameMap["generated"] = "Generated"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.draas.model.replica_disk_collection", fields, reflect.TypeOf(ReplicaDiskCollection{}), fieldNameMap, validators)
}

func ReplicationDataBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["replication_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["replication_name"] = "ReplicationName"
	fields["replication_uuid"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["replication_uuid"] = "ReplicationUuid"
	fields["hbr_server_uuid"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["hbr_server_uuid"] = "HbrServerUuid"
	fields["hbr_server_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["hbr_server_name"] = "HbrServerName"
	fields["scale_out"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["scale_out"] = "ScaleOut"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.draas.model.replication_data", fields, reflect.TypeOf(ReplicationData{}), fieldNameMap, validators)
}

func ReplicationsCountBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["incoming"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["incoming"] = "Incoming"
	fields["replication_uuid"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["replication_uuid"] = "ReplicationUuid"
	fields["outgoing"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["outgoing"] = "Outgoing"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.draas.model.replications_count", fields, reflect.TypeOf(ReplicationsCount{}), fieldNameMap, validators)
}

func SiteRecoveryBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewStringType()
	fieldNameMap["id"] = "Id"
	fields["version"] = vapiBindings_.NewIntegerType()
	fieldNameMap["version"] = "Version"
	fields["created"] = vapiBindings_.NewDateTimeType()
	fieldNameMap["created"] = "Created"
	fields["user_id"] = vapiBindings_.NewStringType()
	fieldNameMap["user_id"] = "UserId"
	fields["user_name"] = vapiBindings_.NewStringType()
	fieldNameMap["user_name"] = "UserName"
	fields["updated"] = vapiBindings_.NewDateTimeType()
	fieldNameMap["updated"] = "Updated"
	fields["updated_by_user_id"] = vapiBindings_.NewStringType()
	fieldNameMap["updated_by_user_id"] = "UpdatedByUserId"
	fields["updated_by_user_name"] = vapiBindings_.NewStringType()
	fieldNameMap["updated_by_user_name"] = "UpdatedByUserName"
	fields["sddc_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sddc_id"] = "SddcId"
	fields["site_recovery_state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["site_recovery_state"] = "SiteRecoveryState"
	fields["draas_h5_url"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["draas_h5_url"] = "DraasH5Url"
	fields["vr_node"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SiteRecoveryNodeBindingType))
	fieldNameMap["vr_node"] = "VrNode"
	fields["srm_nodes"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(SrmNodeBindingType), reflect.TypeOf([]SrmNode{})))
	fieldNameMap["srm_nodes"] = "SrmNodes"
	fields["vrs_nodes"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(SiteRecoveryNodeBindingType), reflect.TypeOf([]SiteRecoveryNode{})))
	fieldNameMap["vrs_nodes"] = "VrsNodes"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.draas.model.site_recovery", fields, reflect.TypeOf(SiteRecovery{}), fieldNameMap, validators)
}

func SiteRecoveryNodeBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["type"] = "Type_"
	fields["state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["state"] = "State"
	fields["vm_moref_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vm_moref_id"] = "VmMorefId"
	fields["ip_address"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["ip_address"] = "IpAddress"
	fields["hostname"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["hostname"] = "Hostname"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.draas.model.site_recovery_node", fields, reflect.TypeOf(SiteRecoveryNode{}), fieldNameMap, validators)
}

func SiteRecoveryNodeVersionBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["node_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["node_id"] = "NodeId"
	fields["node_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["node_type"] = "NodeType"
	fields["node_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["node_ip"] = "NodeIp"
	fields["build_version"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(BuildVersionBindingType))
	fieldNameMap["build_version"] = "BuildVersion"
	fields["full_version"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["full_version"] = "FullVersion"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.draas.model.site_recovery_node_version", fields, reflect.TypeOf(SiteRecoveryNodeVersion{}), fieldNameMap, validators)
}

func SiteRecoveryVersionsBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["sddc_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sddc_id"] = "SddcId"
	fields["node_versions"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(SiteRecoveryNodeVersionBindingType), reflect.TypeOf([]SiteRecoveryNodeVersion{})))
	fieldNameMap["node_versions"] = "NodeVersions"
	fields["generated"] = vapiBindings_.NewOptionalType(vapiBindings_.NewDateTimeType())
	fieldNameMap["generated"] = "Generated"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.draas.model.site_recovery_versions", fields, reflect.TypeOf(SiteRecoveryVersions{}), fieldNameMap, validators)
}

func SrmNodeBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["type"] = "Type_"
	fields["state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["state"] = "State"
	fields["vm_moref_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vm_moref_id"] = "VmMorefId"
	fields["ip_address"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["ip_address"] = "IpAddress"
	fields["hostname"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["hostname"] = "Hostname"
	fields["srm_extension_key_suffix"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["srm_extension_key_suffix"] = "SrmExtensionKeySuffix"
	fields["srm_extension_key"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["srm_extension_key"] = "SrmExtensionKey"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.draas.model.srm_node", fields, reflect.TypeOf(SrmNode{}), fieldNameMap, validators)
}

func TaskBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewStringType()
	fieldNameMap["id"] = "Id"
	fields["version"] = vapiBindings_.NewIntegerType()
	fieldNameMap["version"] = "Version"
	fields["created"] = vapiBindings_.NewDateTimeType()
	fieldNameMap["created"] = "Created"
	fields["user_id"] = vapiBindings_.NewStringType()
	fieldNameMap["user_id"] = "UserId"
	fields["user_name"] = vapiBindings_.NewStringType()
	fieldNameMap["user_name"] = "UserName"
	fields["updated"] = vapiBindings_.NewDateTimeType()
	fieldNameMap["updated"] = "Updated"
	fields["updated_by_user_id"] = vapiBindings_.NewStringType()
	fieldNameMap["updated_by_user_id"] = "UpdatedByUserId"
	fields["updated_by_user_name"] = vapiBindings_.NewStringType()
	fieldNameMap["updated_by_user_name"] = "UpdatedByUserName"
	fields["end_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewDateTimeType())
	fieldNameMap["end_time"] = "EndTime"
	fields["error_message"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["error_message"] = "ErrorMessage"
	fields["tenant_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["tenant_id"] = "TenantId"
	fields["params"] = vapiBindings_.NewOptionalType(vapiBindings_.NewDynamicStructType(nil))
	fieldNameMap["params"] = "Params"
	fields["resource_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["resource_id"] = "ResourceId"
	fields["resource_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["parent_task_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["parent_task_id"] = "ParentTaskId"
	fields["retries"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["retries"] = "Retries"
	fields["start_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewDateTimeType())
	fieldNameMap["start_time"] = "StartTime"
	fields["status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["status"] = "Status"
	fields["sub_status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sub_status"] = "SubStatus"
	fields["task_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["task_type"] = "TaskType"
	fields["task_version"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["task_version"] = "TaskVersion"
	fields["task_progress_phases"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(TaskProgressPhaseBindingType), reflect.TypeOf([]TaskProgressPhase{})))
	fieldNameMap["task_progress_phases"] = "TaskProgressPhases"
	fields["progress_percent"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["progress_percent"] = "ProgressPercent"
	fields["estimated_remaining_minutes"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["estimated_remaining_minutes"] = "EstimatedRemainingMinutes"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.draas.model.task", fields, reflect.TypeOf(Task{}), fieldNameMap, validators)
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
	return vapiBindings_.NewStructType("com.vmware.vmc.draas.model.task_progress_phase", fields, reflect.TypeOf(TaskProgressPhase{}), fieldNameMap, validators)
}
