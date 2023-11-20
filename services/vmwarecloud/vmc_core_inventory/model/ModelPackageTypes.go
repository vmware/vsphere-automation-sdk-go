// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for package: com.vmware.vmwarecloud.vmc_core_inventory.model.
// Includes binding types of a top level structures and enumerations.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package model

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiLog_ "github.com/vmware/vsphere-automation-sdk-go/runtime/log"
	"reflect"
)

type Deployment struct {
	// Identifier of the deployment
	Id *string
	// Name of the deployment
	Name string
	// Identifier of the deployment organization
	OrgId           *string
	Type_           *DeploymentType
	State           *State
	Location        *Location
	Version         *DeploymentVersion
	Creator         *Creator
	Updater         *Updater
	InternalCreator *Creator
	// Internal resource properties
	Properties map[string]string
	Summary    *DeploymentSummary
	// Indicates if the current deployment is multi tenant
	MultiTenant *bool
	// Indicates if the current deployment is deleted
	Deleted *bool
	// Version of the record content to handle optimistic control verification
	ResourceVersion *int64
}

func (s *Deployment) GetType__() vapiBindings_.BindingType {
	return DeploymentBindingType()
}

func (s *Deployment) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for Deployment._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type Creator struct {
	// Name of user that created the deployment
	UserName *string
	// Identifier of the user that created the deployment
	UserId *string
	// Identifier of the client that created the deployment
	ClientId *string
	// Timestamp of deployment creation
	Timestamp *string
}

func (s *Creator) GetType__() vapiBindings_.BindingType {
	return CreatorBindingType()
}

func (s *Creator) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for Creator._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type Pageable struct {
	// Number of elements allowed in each page
	PageSize *int64
	// Page number/offset (0-indexed)
	PageNumber *int64
}

func (s *Pageable) GetType__() vapiBindings_.BindingType {
	return PageableBindingType()
}

func (s *Pageable) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for Pageable._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type Capacity struct {
	// Configured cpu capacity for the host (mhz)
	CpuMhz *int64
	// Configured memory capacity for the host (bytes)
	MemoryBytes *int64
	// Configured storage capacity for the host (bytes)
	StorageBytes *int64
}

func (s *Capacity) GetType__() vapiBindings_.BindingType {
	return CapacityBindingType()
}

func (s *Capacity) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for Capacity._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type Updater struct {
	// Name of user that updated the deployment
	UserName *string
	// Identifier of the user that updated the deployment
	UserId *string
	// Identifier of the client that updated the deployment
	ClientId *string
	// Timestamp of deployment update
	Timestamp *string
}

func (s *Updater) GetType__() vapiBindings_.BindingType {
	return UpdaterBindingType()
}

func (s *Updater) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for Updater._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type HostResponse struct {
	// List of hosts
	Content  []Host
	Pageable *Pageable
	First    *bool
	Last     *bool
	// Total number of pages
	TotalPages *int64
	// Total number of elements
	TotalElements *int64
	// Number of elements returned in current page
	NumberOfElements *int64
	// Number of elements allowed in each page
	Size *int64
	// Page number/offset (0-indexed)
	Number *int64
}

func (s *HostResponse) GetType__() vapiBindings_.BindingType {
	return HostResponseBindingType()
}

func (s *HostResponse) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for HostResponse._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type Membership struct {
	// Explicitly included group members
	Included []GroupMember
	// Explicitly excludes group members (from automatic filter)
	Excluded []GroupMember
}

func (s *Membership) GetType__() vapiBindings_.BindingType {
	return MembershipBindingType()
}

func (s *Membership) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for Membership._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Summary properties related to a deployment overview (deprecated, using DeploymentVcenterInfo)
type DeploymentVcenterSummary struct {
	// URL of the vCenter UI
	VcUrl *string
	// FQDC of the vCenter server
	VcFqdn *string
	// URL of the PowerCLI connect endpoint
	VcPublicIp *string
	// Identifier of the summary
	Id *string
	// Name of the summary
	Name string
	// Identifier of the parent resource
	ResourceId string
	// Identifier of the deployment organization
	OrgId *string
	// Version of the record content to handle optimistic control verification
	ResourceVersion *int64
}

func (s *DeploymentVcenterSummary) GetType__() vapiBindings_.BindingType {
	return DeploymentVcenterSummaryBindingType()
}

func (s *DeploymentVcenterSummary) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for DeploymentVcenterSummary._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Summary properties related to network overview
type DeploymentNetworkSummary struct {
	// NSX manager IP address
	NsxManagerIp *string
	// NSX manager public URL
	PublicNsxManagerUrl *string
	// NSX manager private URL
	PrivateNsxManagerUrl *string
	// Identifier of the summary
	Id *string
	// Name of the summary
	Name string
	// Identifier of the parent resource
	ResourceId string
	// Identifier of the deployment organization
	OrgId *string
	// Version of the record content to handle optimistic control verification
	ResourceVersion *int64
}

func (s *DeploymentNetworkSummary) GetType__() vapiBindings_.BindingType {
	return DeploymentNetworkSummaryBindingType()
}

func (s *DeploymentNetworkSummary) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for DeploymentNetworkSummary._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type ResourceSummary struct {
	// Identifier of the summary
	Id *string
	// Name of the summary
	Name string
	// Identifier of the parent resource
	ResourceId string
	// Identifier of the deployment organization
	OrgId *string
	// Version of the record content to handle optimistic control verification
	ResourceVersion *int64
}

func (s *ResourceSummary) GetType__() vapiBindings_.BindingType {
	return ResourceSummaryBindingType()
}

func (s *ResourceSummary) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ResourceSummary._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type DeploymentGroupResponse struct {
	// List of deployment groups
	Content  []DeploymentGroup
	Pageable *Pageable
	First    *bool
	Last     *bool
	// Total number of pages
	TotalPages *int64
	// Total number of elements
	TotalElements *int64
	// Number of elements returned in current page
	NumberOfElements *int64
	// Number of elements allowed in each page
	Size *int64
	// Page number/offset (0-indexed)
	Number *int64
}

func (s *DeploymentGroupResponse) GetType__() vapiBindings_.BindingType {
	return DeploymentGroupResponseBindingType()
}

func (s *DeploymentGroupResponse) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for DeploymentGroupResponse._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type State struct {
	Phase string
	// sub phase to the high level phase
	SubPhase *string
	// Internationalized state name
	DisplayName *string
	// Error message on failed states
	ErrorMsg *string
	// Unique code identifying the error
	ErrorCode *string
	// Custom provider information
	Provider *string
	// Internal error message on failed states
	InternalErrorMsg *string
	// Internal unique code identifying the error
	InternalErrorCode *string
}

// Resource Lifecycle high level STATE
const State_PHASE_PHASE_ENUM_PENDING = "PENDING"

// Resource Lifecycle high level STATE
const State_PHASE_PHASE_ENUM_READY = "READY"

// Resource Lifecycle high level STATE
const State_PHASE_PHASE_ENUM_DELETED = "DELETED"

// Resource Lifecycle high level STATE
const State_PHASE_PHASE_ENUM_FAILED = "FAILED"

// Resource Lifecycle high level STATE
const State_PHASE_PHASE_ENUM_CANCELED = "CANCELED"

func (s *State) GetType__() vapiBindings_.BindingType {
	return StateBindingType()
}

func (s *State) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for State._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type GroupMember struct {
	// ID of the deployment
	DeploymentId *string
}

func (s *GroupMember) GetType__() vapiBindings_.BindingType {
	return GroupMemberBindingType()
}

func (s *GroupMember) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for GroupMember._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type Config struct {
	// ID of config
	Id *string
	// ID of organization
	OrgId *string
	// Type of config
	Type_ string
	// ID of the owning operation
	OperationId     *string
	Creator         *Creator
	Updater         *Updater
	InternalCreator *Creator
	// (internal) Property bag used to serialize the derived type
	UnmappedProperties map[string]*vapiData_.StructValue
	// The vSAN version to be used in the SDDC's primary cluster
	VsanVersion *string
}

// The vSAN version to be used in the SDDC's primary cluster
const Config_VSAN_VERSION_VSAN1 = "VSAN1"

// The vSAN version to be used in the SDDC's primary cluster
const Config_VSAN_VERSION_VSANESA = "VSANESA"

// The vSAN version to be used in the SDDC's primary cluster
const Config_VSAN_VERSION_NOVSAN = "NOVSAN"

func (s *Config) GetType__() vapiBindings_.BindingType {
	return ConfigBindingType()
}

func (s *Config) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for Config._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type DeploymentVersion struct {
	// Name of the SDDC version bundle
	Name string
	// Unique code identifying SDDC version
	Code *string
}

func (s *DeploymentVersion) GetType__() vapiBindings_.BindingType {
	return DeploymentVersionBindingType()
}

func (s *DeploymentVersion) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for DeploymentVersion._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Summary properties related to a deployment overview
type DeploymentOverviewSummary struct {
	// Access state of the SDDC (ENABLED, DISABLED)
	SddcAccessState *string
	// Total number of clusters in the deployment
	ClusterCount *int64
	// Total number of hosts in the deployment
	HostCount *int64
	// Size of the SDDC appliance (medium, large, ...)
	SddcSize           *string
	ConfiguredCapacity *Capacity
	// Identifier of the summary
	Id *string
	// Name of the summary
	Name string
	// Identifier of the parent resource
	ResourceId string
	// Identifier of the deployment organization
	OrgId *string
	// Version of the record content to handle optimistic control verification
	ResourceVersion *int64
}

func (s *DeploymentOverviewSummary) GetType__() vapiBindings_.BindingType {
	return DeploymentOverviewSummaryBindingType()
}

func (s *DeploymentOverviewSummary) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for DeploymentOverviewSummary._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type DeploymentSummary struct {
	DeploymentOverviewSummary *DeploymentOverviewSummary
	DeploymentVcenterSummary  *DeploymentVcenterSummary
	DeploymentNetworkSummary  *DeploymentNetworkSummary
}

func (s *DeploymentSummary) GetType__() vapiBindings_.BindingType {
	return DeploymentSummaryBindingType()
}

func (s *DeploymentSummary) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for DeploymentSummary._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type DeploymentType struct {
	Code *string
	// Name of the provider name
	ProviderName *string
	ProviderCode *string
	// Type of configuration (deprecated by Cluster.availability_type)
	ConfigurationType *string
	// Provider type of the deployment
	ProviderInternalType *string
	// Provider type of the deployment in case of multi-provider enviornment
	MultiProviderInternalType *string
	// Provider subtype of the deployment in case of multi-provider enviornment
	MultiProviderInternalSubType *string
}

// provider code
const DeploymentType_CODE_PROVIDER_CODE_ENUM_VMC_AWS = "vmc-aws"

// provider code
const DeploymentType_CODE_PROVIDER_CODE_ENUM_ONPREM_VC = "onprem-vc"

// provider code
const DeploymentType_CODE_PROVIDER_CODE_ENUM_VMC_DIMENSION = "vmc-dimension"

// provider code
const DeploymentType_CODE_PROVIDER_CODE_ENUM_VMC_AWS_OUTPOSTS = "vmc-aws-outposts"

// provider code
const DeploymentType_CODE_PROVIDER_CODE_ENUM_VCFPLUS_WLD = "vcfplus-wld"

// provider code
const DeploymentType_CODE_PROVIDER_CODE_ENUM_VCFPLUS_MGMT = "vcfplus-mgmt"

// provider code
const DeploymentType_CODE_PROVIDER_CODE_ENUM_VCF_WLD = "vcf-wld"

// provider code
const DeploymentType_CODE_PROVIDER_CODE_ENUM_VCF_MGMT = "vcf-mgmt"

// provider code
const DeploymentType_CODE_PROVIDER_CODE_ENUM_EQUINIX = "equinix"

// provider code
const DeploymentType_PROVIDER_CODE_PROVIDER_CODE_ENUM_VMC_AWS = "vmc-aws"

// provider code
const DeploymentType_PROVIDER_CODE_PROVIDER_CODE_ENUM_ONPREM_VC = "onprem-vc"

// provider code
const DeploymentType_PROVIDER_CODE_PROVIDER_CODE_ENUM_VMC_DIMENSION = "vmc-dimension"

// provider code
const DeploymentType_PROVIDER_CODE_PROVIDER_CODE_ENUM_VMC_AWS_OUTPOSTS = "vmc-aws-outposts"

// provider code
const DeploymentType_PROVIDER_CODE_PROVIDER_CODE_ENUM_VCFPLUS_WLD = "vcfplus-wld"

// provider code
const DeploymentType_PROVIDER_CODE_PROVIDER_CODE_ENUM_VCFPLUS_MGMT = "vcfplus-mgmt"

// provider code
const DeploymentType_PROVIDER_CODE_PROVIDER_CODE_ENUM_VCF_WLD = "vcf-wld"

// provider code
const DeploymentType_PROVIDER_CODE_PROVIDER_CODE_ENUM_VCF_MGMT = "vcf-mgmt"

// provider code
const DeploymentType_PROVIDER_CODE_PROVIDER_CODE_ENUM_EQUINIX = "equinix"

func (s *DeploymentType) GetType__() vapiBindings_.BindingType {
	return DeploymentTypeBindingType()
}

func (s *DeploymentType) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for DeploymentType._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type Operation struct {
	// Identifier of the operation
	Id *string
	// Type of the operation
	Type_ string
	// Identifier of the resource being operated on
	ResourceId string
	// Type of the resource being operated on
	ResourceType string
	// Identifier of operation organization
	OrgId *string
	// Identifier assigned by the provider
	ProviderAssignedId *string
	// Identifier of scheduler task created for operation
	TaskId          *string
	State           *State
	Creator         *Creator
	Updater         *Updater
	InternalCreator *Creator
	Config          *Config
	// Indicates if the current operation is deleted
	Deleted *bool
}

func (s *Operation) GetType__() vapiBindings_.BindingType {
	return OperationBindingType()
}

func (s *Operation) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for Operation._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type ClusterResponse struct {
	// List of clusters
	Content  []Cluster
	Pageable *Pageable
	First    *bool
	Last     *bool
	// Total number of pages
	TotalPages *int64
	// Total number of elements
	TotalElements *int64
	// Number of elements returned in current page
	NumberOfElements *int64
	// Number of elements allowed in each page
	Size *int64
	// Page number/offset (0-indexed)
	Number *int64
}

func (s *ClusterResponse) GetType__() vapiBindings_.BindingType {
	return ClusterResponseBindingType()
}

func (s *ClusterResponse) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ClusterResponse._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type Cluster struct {
	// Identifier of the cluster
	Id *string
	// Name of the cluster
	Name string
	// Identifier of the parent deployment
	DeploymentId string
	// Identifier of the organization
	OrgId *string
	// Checks if the current cluster is a primary cluster of the Deployment
	PrimaryCluster  *bool
	State           *State
	Creator         *Creator
	Updater         *Updater
	InternalCreator *Creator
	// Indicates if the current cluster is deleted
	Deleted *bool
	// Version of the record content to handle optimistic control verification
	ResourceVersion *int64
}

func (s *Cluster) GetType__() vapiBindings_.BindingType {
	return ClusterBindingType()
}

func (s *Cluster) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for Cluster._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type Host struct {
	// Identifier of the deployment host
	Id *string
	// Name of host
	Name string
	// Identifier of the parent deployment
	DeploymentId string
	// Identifier of the parent cluster
	ClusterId string
	// Identifier of the organization
	OrgId           *string
	State           *State
	Capacity        *Capacity
	Creator         *Creator
	Updater         *Updater
	InternalCreator *Creator
	// Indicates if the current host is deleted
	Deleted *bool
	// Version of the record content to handle optimistic control verification
	ResourceVersion *int64
}

func (s *Host) GetType__() vapiBindings_.BindingType {
	return HostBindingType()
}

func (s *Host) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for Host._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type ErrorResponse struct {
	// Time error was created
	Timestamp string
	// HTTP status code
	Status int64
	// Translation of status
	Error_ string
	// Error message
	Message string
	// Originating request URI
	Path string
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

type DeploymentResponse struct {
	// List of deployments
	Content  []Deployment
	Pageable *Pageable
	First    *bool
	Last     *bool
	// Total number of pages
	TotalPages *int64
	// Total number of elements
	TotalElements *int64
	// Number of elements returned in current page
	NumberOfElements *int64
	// Number of elements allowed in each page
	Size *int64
	// Page number/offset (0-indexed)
	Number *int64
}

func (s *DeploymentResponse) GetType__() vapiBindings_.BindingType {
	return DeploymentResponseBindingType()
}

func (s *DeploymentResponse) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for DeploymentResponse._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type OperationResponse struct {
	// List of operations
	Content  []Operation
	Pageable *Pageable
	First    *bool
	Last     *bool
	// Total number of pages
	TotalPages *int64
	// Total number of elements
	TotalElements *int64
	// Number of elements returned in current page
	NumberOfElements *int64
	// Number of elements allowed in each page
	Size *int64
	// Page number/offset (0-indexed)
	Number *int64
}

func (s *OperationResponse) GetType__() vapiBindings_.BindingType {
	return OperationResponseBindingType()
}

func (s *OperationResponse) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for OperationResponse._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type DeploymentGroup struct {
	// Identifier of the deployment group
	Id *string
	// Name of the deployment group
	Name string
	// Identifier of the deployment group organization
	OrgId *string
	// Describes the deployment group
	Description     *string
	DeploymentType  *DeploymentType
	State           *State
	Creator         *Creator
	Updater         *Updater
	InternalCreator *Creator
	Membership      *Membership
	// Customer defined resource tags
	Tags map[string]string
	// Internal resource properties
	Properties map[string]string
	// Indicates if the current deployment group is deleted
	Deleted *bool
	// Version of the record content to handle optimistic control verification
	ResourceVersion *int64
}

func (s *DeploymentGroup) GetType__() vapiBindings_.BindingType {
	return DeploymentGroupBindingType()
}

func (s *DeploymentGroup) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for DeploymentGroup._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type Location struct {
	// Name of the location
	Name string
	// Unique code identifying location
	Code *string
	// Provider specific deployment name
	Address map[string]string
}

func (s *Location) GetType__() vapiBindings_.BindingType {
	return LocationBindingType()
}

func (s *Location) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for Location._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

func DeploymentBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["name"] = vapiBindings_.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["org_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fields["type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(DeploymentTypeBindingType))
	fieldNameMap["type"] = "Type_"
	fields["state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(StateBindingType))
	fieldNameMap["state"] = "State"
	fields["location"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(LocationBindingType))
	fieldNameMap["location"] = "Location"
	fields["version"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(DeploymentVersionBindingType))
	fieldNameMap["version"] = "Version"
	fields["creator"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(CreatorBindingType))
	fieldNameMap["creator"] = "Creator"
	fields["updater"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(UpdaterBindingType))
	fieldNameMap["updater"] = "Updater"
	fields["internal_creator"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(CreatorBindingType))
	fieldNameMap["internal_creator"] = "InternalCreator"
	fields["properties"] = vapiBindings_.NewOptionalType(vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewStringType(), reflect.TypeOf(map[string]string{})))
	fieldNameMap["properties"] = "Properties"
	fields["summary"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(DeploymentSummaryBindingType))
	fieldNameMap["summary"] = "Summary"
	fields["multi_tenant"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["multi_tenant"] = "MultiTenant"
	fields["deleted"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["deleted"] = "Deleted"
	fields["resource_version"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["resource_version"] = "ResourceVersion"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_inventory.model.deployment", fields, reflect.TypeOf(Deployment{}), fieldNameMap, validators)
}

func CreatorBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["user_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["user_name"] = "UserName"
	fields["user_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["user_id"] = "UserId"
	fields["client_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["client_id"] = "ClientId"
	fields["timestamp"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["timestamp"] = "Timestamp"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_inventory.model.creator", fields, reflect.TypeOf(Creator{}), fieldNameMap, validators)
}

func PageableBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["page_size"] = "PageSize"
	fields["page_number"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["page_number"] = "PageNumber"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_inventory.model.pageable", fields, reflect.TypeOf(Pageable{}), fieldNameMap, validators)
}

func CapacityBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cpu_mhz"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["cpu_mhz"] = "CpuMhz"
	fields["memory_bytes"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["memory_bytes"] = "MemoryBytes"
	fields["storage_bytes"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["storage_bytes"] = "StorageBytes"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_inventory.model.capacity", fields, reflect.TypeOf(Capacity{}), fieldNameMap, validators)
}

func UpdaterBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["user_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["user_name"] = "UserName"
	fields["user_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["user_id"] = "UserId"
	fields["client_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["client_id"] = "ClientId"
	fields["timestamp"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["timestamp"] = "Timestamp"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_inventory.model.updater", fields, reflect.TypeOf(Updater{}), fieldNameMap, validators)
}

func HostResponseBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["content"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(HostBindingType), reflect.TypeOf([]Host{})))
	fieldNameMap["content"] = "Content"
	fields["pageable"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(PageableBindingType))
	fieldNameMap["pageable"] = "Pageable"
	fields["first"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["first"] = "First"
	fields["last"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["last"] = "Last"
	fields["total_pages"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["total_pages"] = "TotalPages"
	fields["total_elements"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["total_elements"] = "TotalElements"
	fields["number_of_elements"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["number_of_elements"] = "NumberOfElements"
	fields["size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["size"] = "Size"
	fields["number"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["number"] = "Number"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_inventory.model.host_response", fields, reflect.TypeOf(HostResponse{}), fieldNameMap, validators)
}

func MembershipBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["included"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(GroupMemberBindingType), reflect.TypeOf([]GroupMember{})))
	fieldNameMap["included"] = "Included"
	fields["excluded"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(GroupMemberBindingType), reflect.TypeOf([]GroupMember{})))
	fieldNameMap["excluded"] = "Excluded"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_inventory.model.membership", fields, reflect.TypeOf(Membership{}), fieldNameMap, validators)
}

func DeploymentVcenterSummaryBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vc_url"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vc_url"] = "VcUrl"
	fields["vc_fqdn"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vc_fqdn"] = "VcFqdn"
	fields["vc_public_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vc_public_ip"] = "VcPublicIp"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["name"] = vapiBindings_.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["resource_id"] = vapiBindings_.NewStringType()
	fieldNameMap["resource_id"] = "ResourceId"
	fields["org_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fields["resource_version"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["resource_version"] = "ResourceVersion"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_inventory.model.deployment_vcenter_summary", fields, reflect.TypeOf(DeploymentVcenterSummary{}), fieldNameMap, validators)
}

func DeploymentNetworkSummaryBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["nsx_manager_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["nsx_manager_ip"] = "NsxManagerIp"
	fields["public_nsx_manager_url"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["public_nsx_manager_url"] = "PublicNsxManagerUrl"
	fields["private_nsx_manager_url"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["private_nsx_manager_url"] = "PrivateNsxManagerUrl"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["name"] = vapiBindings_.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["resource_id"] = vapiBindings_.NewStringType()
	fieldNameMap["resource_id"] = "ResourceId"
	fields["org_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fields["resource_version"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["resource_version"] = "ResourceVersion"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_inventory.model.deployment_network_summary", fields, reflect.TypeOf(DeploymentNetworkSummary{}), fieldNameMap, validators)
}

func ResourceSummaryBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["name"] = vapiBindings_.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["resource_id"] = vapiBindings_.NewStringType()
	fieldNameMap["resource_id"] = "ResourceId"
	fields["org_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fields["resource_version"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["resource_version"] = "ResourceVersion"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_inventory.model.resource_summary", fields, reflect.TypeOf(ResourceSummary{}), fieldNameMap, validators)
}

func DeploymentGroupResponseBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["content"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(DeploymentGroupBindingType), reflect.TypeOf([]DeploymentGroup{})))
	fieldNameMap["content"] = "Content"
	fields["pageable"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(PageableBindingType))
	fieldNameMap["pageable"] = "Pageable"
	fields["first"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["first"] = "First"
	fields["last"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["last"] = "Last"
	fields["total_pages"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["total_pages"] = "TotalPages"
	fields["total_elements"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["total_elements"] = "TotalElements"
	fields["number_of_elements"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["number_of_elements"] = "NumberOfElements"
	fields["size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["size"] = "Size"
	fields["number"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["number"] = "Number"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_inventory.model.deployment_group_response", fields, reflect.TypeOf(DeploymentGroupResponse{}), fieldNameMap, validators)
}

func StateBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["phase"] = vapiBindings_.NewStringType()
	fieldNameMap["phase"] = "Phase"
	fields["sub_phase"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sub_phase"] = "SubPhase"
	fields["display_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["error_msg"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["error_msg"] = "ErrorMsg"
	fields["error_code"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["error_code"] = "ErrorCode"
	fields["provider"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["provider"] = "Provider"
	fields["internal_error_msg"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["internal_error_msg"] = "InternalErrorMsg"
	fields["internal_error_code"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["internal_error_code"] = "InternalErrorCode"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_inventory.model.state", fields, reflect.TypeOf(State{}), fieldNameMap, validators)
}

func GroupMemberBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["deployment_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["deployment_id"] = "DeploymentId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_inventory.model.group_member", fields, reflect.TypeOf(GroupMember{}), fieldNameMap, validators)
}

func ConfigBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["org_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fields["type"] = vapiBindings_.NewStringType()
	fieldNameMap["type"] = "Type_"
	fields["operation_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["operation_id"] = "OperationId"
	fields["creator"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(CreatorBindingType))
	fieldNameMap["creator"] = "Creator"
	fields["updater"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(UpdaterBindingType))
	fieldNameMap["updater"] = "Updater"
	fields["internal_creator"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(CreatorBindingType))
	fieldNameMap["internal_creator"] = "InternalCreator"
	fields["unmapped_properties"] = vapiBindings_.NewOptionalType(vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewDynamicStructType(nil), reflect.TypeOf(map[string]*vapiData_.StructValue{})))
	fieldNameMap["unmapped_properties"] = "UnmappedProperties"
	fields["vsan_version"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vsan_version"] = "VsanVersion"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_inventory.model.config", fields, reflect.TypeOf(Config{}), fieldNameMap, validators)
}

func DeploymentVersionBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = vapiBindings_.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["code"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["code"] = "Code"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_inventory.model.deployment_version", fields, reflect.TypeOf(DeploymentVersion{}), fieldNameMap, validators)
}

func DeploymentOverviewSummaryBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["sddc_access_state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sddc_access_state"] = "SddcAccessState"
	fields["cluster_count"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["cluster_count"] = "ClusterCount"
	fields["host_count"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["host_count"] = "HostCount"
	fields["sddc_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sddc_size"] = "SddcSize"
	fields["configured_capacity"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(CapacityBindingType))
	fieldNameMap["configured_capacity"] = "ConfiguredCapacity"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["name"] = vapiBindings_.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["resource_id"] = vapiBindings_.NewStringType()
	fieldNameMap["resource_id"] = "ResourceId"
	fields["org_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fields["resource_version"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["resource_version"] = "ResourceVersion"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_inventory.model.deployment_overview_summary", fields, reflect.TypeOf(DeploymentOverviewSummary{}), fieldNameMap, validators)
}

func DeploymentSummaryBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["DeploymentOverviewSummary"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(DeploymentOverviewSummaryBindingType))
	fieldNameMap["DeploymentOverviewSummary"] = "DeploymentOverviewSummary"
	fields["DeploymentVcenterSummary"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(DeploymentVcenterSummaryBindingType))
	fieldNameMap["DeploymentVcenterSummary"] = "DeploymentVcenterSummary"
	fields["DeploymentNetworkSummary"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(DeploymentNetworkSummaryBindingType))
	fieldNameMap["DeploymentNetworkSummary"] = "DeploymentNetworkSummary"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_inventory.model.deployment_summary", fields, reflect.TypeOf(DeploymentSummary{}), fieldNameMap, validators)
}

func DeploymentTypeBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["code"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["code"] = "Code"
	fields["provider_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["provider_name"] = "ProviderName"
	fields["provider_code"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["provider_code"] = "ProviderCode"
	fields["configuration_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["configuration_type"] = "ConfigurationType"
	fields["provider_internal_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["provider_internal_type"] = "ProviderInternalType"
	fields["multi_provider_internal_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["multi_provider_internal_type"] = "MultiProviderInternalType"
	fields["multi_provider_internal_sub_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["multi_provider_internal_sub_type"] = "MultiProviderInternalSubType"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_inventory.model.deployment_type", fields, reflect.TypeOf(DeploymentType{}), fieldNameMap, validators)
}

func OperationBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["type"] = vapiBindings_.NewStringType()
	fieldNameMap["type"] = "Type_"
	fields["resource_id"] = vapiBindings_.NewStringType()
	fieldNameMap["resource_id"] = "ResourceId"
	fields["resource_type"] = vapiBindings_.NewStringType()
	fieldNameMap["resource_type"] = "ResourceType"
	fields["org_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fields["provider_assigned_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["provider_assigned_id"] = "ProviderAssignedId"
	fields["task_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["task_id"] = "TaskId"
	fields["state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(StateBindingType))
	fieldNameMap["state"] = "State"
	fields["creator"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(CreatorBindingType))
	fieldNameMap["creator"] = "Creator"
	fields["updater"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(UpdaterBindingType))
	fieldNameMap["updater"] = "Updater"
	fields["internal_creator"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(CreatorBindingType))
	fieldNameMap["internal_creator"] = "InternalCreator"
	fields["config"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(ConfigBindingType))
	fieldNameMap["config"] = "Config"
	fields["deleted"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["deleted"] = "Deleted"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_inventory.model.operation", fields, reflect.TypeOf(Operation{}), fieldNameMap, validators)
}

func ClusterResponseBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["content"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ClusterBindingType), reflect.TypeOf([]Cluster{})))
	fieldNameMap["content"] = "Content"
	fields["pageable"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(PageableBindingType))
	fieldNameMap["pageable"] = "Pageable"
	fields["first"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["first"] = "First"
	fields["last"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["last"] = "Last"
	fields["total_pages"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["total_pages"] = "TotalPages"
	fields["total_elements"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["total_elements"] = "TotalElements"
	fields["number_of_elements"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["number_of_elements"] = "NumberOfElements"
	fields["size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["size"] = "Size"
	fields["number"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["number"] = "Number"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_inventory.model.cluster_response", fields, reflect.TypeOf(ClusterResponse{}), fieldNameMap, validators)
}

func ClusterBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["name"] = vapiBindings_.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["deployment_id"] = vapiBindings_.NewStringType()
	fieldNameMap["deployment_id"] = "DeploymentId"
	fields["org_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fields["primary_cluster"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["primary_cluster"] = "PrimaryCluster"
	fields["state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(StateBindingType))
	fieldNameMap["state"] = "State"
	fields["creator"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(CreatorBindingType))
	fieldNameMap["creator"] = "Creator"
	fields["updater"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(UpdaterBindingType))
	fieldNameMap["updater"] = "Updater"
	fields["internal_creator"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(CreatorBindingType))
	fieldNameMap["internal_creator"] = "InternalCreator"
	fields["deleted"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["deleted"] = "Deleted"
	fields["resource_version"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["resource_version"] = "ResourceVersion"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_inventory.model.cluster", fields, reflect.TypeOf(Cluster{}), fieldNameMap, validators)
}

func HostBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["name"] = vapiBindings_.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["deployment_id"] = vapiBindings_.NewStringType()
	fieldNameMap["deployment_id"] = "DeploymentId"
	fields["cluster_id"] = vapiBindings_.NewStringType()
	fieldNameMap["cluster_id"] = "ClusterId"
	fields["org_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fields["state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(StateBindingType))
	fieldNameMap["state"] = "State"
	fields["capacity"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(CapacityBindingType))
	fieldNameMap["capacity"] = "Capacity"
	fields["creator"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(CreatorBindingType))
	fieldNameMap["creator"] = "Creator"
	fields["updater"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(UpdaterBindingType))
	fieldNameMap["updater"] = "Updater"
	fields["internal_creator"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(CreatorBindingType))
	fieldNameMap["internal_creator"] = "InternalCreator"
	fields["deleted"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["deleted"] = "Deleted"
	fields["resource_version"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["resource_version"] = "ResourceVersion"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_inventory.model.host", fields, reflect.TypeOf(Host{}), fieldNameMap, validators)
}

func ErrorResponseBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["timestamp"] = vapiBindings_.NewStringType()
	fieldNameMap["timestamp"] = "Timestamp"
	fields["status"] = vapiBindings_.NewIntegerType()
	fieldNameMap["status"] = "Status"
	fields["error"] = vapiBindings_.NewStringType()
	fieldNameMap["error"] = "Error_"
	fields["message"] = vapiBindings_.NewStringType()
	fieldNameMap["message"] = "Message"
	fields["path"] = vapiBindings_.NewStringType()
	fieldNameMap["path"] = "Path"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_inventory.model.error_response", fields, reflect.TypeOf(ErrorResponse{}), fieldNameMap, validators)
}

func DeploymentResponseBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["content"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(DeploymentBindingType), reflect.TypeOf([]Deployment{})))
	fieldNameMap["content"] = "Content"
	fields["pageable"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(PageableBindingType))
	fieldNameMap["pageable"] = "Pageable"
	fields["first"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["first"] = "First"
	fields["last"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["last"] = "Last"
	fields["total_pages"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["total_pages"] = "TotalPages"
	fields["total_elements"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["total_elements"] = "TotalElements"
	fields["number_of_elements"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["number_of_elements"] = "NumberOfElements"
	fields["size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["size"] = "Size"
	fields["number"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["number"] = "Number"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_inventory.model.deployment_response", fields, reflect.TypeOf(DeploymentResponse{}), fieldNameMap, validators)
}

func OperationResponseBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["content"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(OperationBindingType), reflect.TypeOf([]Operation{})))
	fieldNameMap["content"] = "Content"
	fields["pageable"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(PageableBindingType))
	fieldNameMap["pageable"] = "Pageable"
	fields["first"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["first"] = "First"
	fields["last"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["last"] = "Last"
	fields["total_pages"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["total_pages"] = "TotalPages"
	fields["total_elements"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["total_elements"] = "TotalElements"
	fields["number_of_elements"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["number_of_elements"] = "NumberOfElements"
	fields["size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["size"] = "Size"
	fields["number"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["number"] = "Number"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_inventory.model.operation_response", fields, reflect.TypeOf(OperationResponse{}), fieldNameMap, validators)
}

func DeploymentGroupBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["name"] = vapiBindings_.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["org_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fields["description"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["deployment_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(DeploymentTypeBindingType))
	fieldNameMap["deployment_type"] = "DeploymentType"
	fields["state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(StateBindingType))
	fieldNameMap["state"] = "State"
	fields["creator"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(CreatorBindingType))
	fieldNameMap["creator"] = "Creator"
	fields["updater"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(UpdaterBindingType))
	fieldNameMap["updater"] = "Updater"
	fields["internal_creator"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(CreatorBindingType))
	fieldNameMap["internal_creator"] = "InternalCreator"
	fields["membership"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(MembershipBindingType))
	fieldNameMap["membership"] = "Membership"
	fields["tags"] = vapiBindings_.NewOptionalType(vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewStringType(), reflect.TypeOf(map[string]string{})))
	fieldNameMap["tags"] = "Tags"
	fields["properties"] = vapiBindings_.NewOptionalType(vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewStringType(), reflect.TypeOf(map[string]string{})))
	fieldNameMap["properties"] = "Properties"
	fields["deleted"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["deleted"] = "Deleted"
	fields["resource_version"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["resource_version"] = "ResourceVersion"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_inventory.model.deployment_group", fields, reflect.TypeOf(DeploymentGroup{}), fieldNameMap, validators)
}

func LocationBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = vapiBindings_.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["code"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["code"] = "Code"
	fields["address"] = vapiBindings_.NewOptionalType(vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewStringType(), reflect.TypeOf(map[string]string{})))
	fieldNameMap["address"] = "Address"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_inventory.model.location", fields, reflect.TypeOf(Location{}), fieldNameMap, validators)
}
