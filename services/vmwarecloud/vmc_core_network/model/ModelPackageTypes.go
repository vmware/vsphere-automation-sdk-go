// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for package: com.vmware.vmwarecloud.vmc_core_network.model.
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

// Activity response
type Activity struct {
	Id                    *string
	OrgId                 *string
	State                 *string
	ActivityTypeKey       *string
	ActivityTypeName      *string
	ActivityTypeNamespace *string
	Progress              *float64
	Resources             []AffectedResource
	CreatedBy             *string
	CreatedTimestamp      *string
	LastUpdatedBy         *string
	LastUpdatedTimestamp  *string
}

const Activity_STATE_INITIALIZED = "INITIALIZED"
const Activity_STATE_SUBMITTED = "SUBMITTED"
const Activity_STATE_RUNNING = "RUNNING"
const Activity_STATE_COMPLETED = "COMPLETED"
const Activity_STATE_CANCELED = "CANCELED"
const Activity_STATE_FAILED = "FAILED"
const Activity_STATE_SUSPENDED = "SUSPENDED"
const Activity_STATE_UNRECOGNIZED = "UNRECOGNIZED"
const Activity_STATE_RESTARTING = "RESTARTING"
const Activity_STATE_PENDING_PAUSE = "PENDING_PAUSE"
const Activity_STATE_PENDING_RETRY = "PENDING_RETRY"
const Activity_STATE_PENDING_RESUME = "PENDING_RESUME"
const Activity_STATE_PENDING_CANCEL = "PENDING_CANCEL"

func (s *Activity) GetType__() vapiBindings_.BindingType {
	return ActivityBindingType()
}

func (s *Activity) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for Activity._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type RouteEntry struct {
	// Identifier of the route entry
	Id *string
	// Identifier of the route table for this route entry
	RouteTableId *string
	// Destination CIDR Block for this route entry.
	Destination *string
	// CIDR Block IP Address format
	AddressFamily *string
	// Variable to indicate if Route Table is duplicated
	IsDuplicated    *bool
	OverlapCidrs    []string
	Target          *RouteEntryTarget
	Creator         *Creator
	InternalCreator *Creator
	Updater         *Updater
	// Customer defined resource tags
	Tags map[string]string
	// Internal resource properties
	Properties map[string]string
	// Error message specific to this route entry
	ErrorMsgs map[string]string
}

// CIDR Block IP Address format
const RouteEntry_ADDRESS_FAMILY_IPV4 = "IPv4"

// CIDR Block IP Address format
const RouteEntry_ADDRESS_FAMILY_IPV6 = "IPv6"

func (s *RouteEntry) GetType__() vapiBindings_.BindingType {
	return RouteEntryBindingType()
}

func (s *RouteEntry) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for RouteEntry._GetDataValue method - %s",
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

type NsxAuthenticationEndpoint struct {
	Idp          *string
	PreferredUrl *string
	OtherUrls    []string
	Desc         *string
}

func (s *NsxAuthenticationEndpoint) GetType__() vapiBindings_.BindingType {
	return NsxAuthenticationEndpointBindingType()
}

func (s *NsxAuthenticationEndpoint) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for NsxAuthenticationEndpoint._GetDataValue method - %s",
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

// the response body of a Bad Request, which due to validation errors
type ValidationErrorResponse struct {
	Details []ValidationErrorDetail
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

func (s *ValidationErrorResponse) GetType__() vapiBindings_.BindingType {
	return ValidationErrorResponseBindingType()
}

func (s *ValidationErrorResponse) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ValidationErrorResponse._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type SharedPrefixListResourceShare struct {
	Name  *string
	Arn   *string
	State *string
}

const SharedPrefixListResourceShare_STATE_ASSOCIATING = "ASSOCIATING"
const SharedPrefixListResourceShare_STATE_ASSOCIATED = "ASSOCIATED"
const SharedPrefixListResourceShare_STATE_FAILED = "FAILED"
const SharedPrefixListResourceShare_STATE_DISASSOCIATING = "DISASSOCIATING"
const SharedPrefixListResourceShare_STATE_DISASSOCIATED = "DISASSOCIATED"
const SharedPrefixListResourceShare_STATE_UNKNOWN = "UNKNOWN"
const SharedPrefixListResourceShare_STATE_PARTIALLY_ASSOCIATED = "PARTIALLY_ASSOCIATED"

func (s *SharedPrefixListResourceShare) GetType__() vapiBindings_.BindingType {
	return SharedPrefixListResourceShareBindingType()
}

func (s *SharedPrefixListResourceShare) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for SharedPrefixListResourceShare._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type SharedPrefixListAccount struct {
	AccountId   *string
	ResourceIds []string
}

func (s *SharedPrefixListAccount) GetType__() vapiBindings_.BindingType {
	return SharedPrefixListAccountBindingType()
}

func (s *SharedPrefixListAccount) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for SharedPrefixListAccount._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type GroupToGroupConnectivityPeering struct {
	// Region ID for Source Group
	Region *string
	// Region name for Source Group
	RegionName *string
	// TGW ID for Region of the Source Group
	TgwId *string
	// Region ID of Peering (target) Group
	PeeringRegion *string
	// Region name of Peering (target) Group
	PeeringRegionName *string
	// TGW ID for Region of the Peering (target) Group
	PeeringTgwId *string
	// Attachment ID of TGW Peering for group-to-group connectivity
	AttachmentId *string
	// Status of a TGW group-to-group connectivity peering
	Status *string
	// Status of a TGW group-to-group connectivity peering for Localization support
	StatusText *string
}

// Status of a TGW group-to-group connectivity peering
const GroupToGroupConnectivityPeering_STATUS_CONNECTED = "CONNECTED"

// Status of a TGW group-to-group connectivity peering
const GroupToGroupConnectivityPeering_STATUS_PENDING = "PENDING"

// Status of a TGW group-to-group connectivity peering
const GroupToGroupConnectivityPeering_STATUS_DELETING = "DELETING"

// Status of a TGW group-to-group connectivity peering
const GroupToGroupConnectivityPeering_STATUS_FAILED = "FAILED"

func (s *GroupToGroupConnectivityPeering) GetType__() vapiBindings_.BindingType {
	return GroupToGroupConnectivityPeeringBindingType()
}

func (s *GroupToGroupConnectivityPeering) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for GroupToGroupConnectivityPeering._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type GroupToGroupConnectivityInfo struct {
	// Identifier of peering organization
	PeeringOrgId *string
	// Identifier of the peering network connectivity config for group-to-group connectivity
	PeeringNetworkConnectivityConfigId *string
	// Identifier of the peering group, this is deployment group id of the peering group
	PeeringGroupId *string
	// Name of the peering group
	PeeringGroupName *string
	// Status of group-to-group Peering Connectivity
	Status *string
	// Status of group-to-group Peering Connectivity for Localization support
	StatusText *string
	// Peering information for group-to-group connectivity
	Peerings []GroupToGroupConnectivityPeering
}

// Status of group-to-group Peering Connectivity
const GroupToGroupConnectivityInfo_STATUS_CONNECTED = "CONNECTED"

// Status of group-to-group Peering Connectivity
const GroupToGroupConnectivityInfo_STATUS_IN_PROGRESS = "IN PROGRESS"

// Status of group-to-group Peering Connectivity
const GroupToGroupConnectivityInfo_STATUS_FAILED = "FAILED"

func (s *GroupToGroupConnectivityInfo) GetType__() vapiBindings_.BindingType {
	return GroupToGroupConnectivityInfoBindingType()
}

func (s *GroupToGroupConnectivityInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for GroupToGroupConnectivityInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type GroupToGroupConnectivity struct {
	// Atlas generated identifier for group-to-group connectivity
	Id *string
	// Identifier of the source group organization
	OrgId *string
	// Name of the group-to-group connectivity
	Name *string
	// Description of the group-to-group connectivity
	Description *string
	// Identifier of the network connectivity config for group-to-group connectivity
	NetworkConnectivityConfigId *string
	// Identifier of the source group, this is deployment group id of the source group
	GroupId *string
	// Name of the source group
	GroupName *string
	// Timestamp of last group-to-group connectivity sync
	LastSyncTime *string
	// Version of group-to-group connectivity
	Version *int64
	// Status of group-to-group connectivity
	State *string
	// Status of group-to-group connectivity for Localization support
	StateText       *string
	Creator         *Creator
	Updater         *Updater
	InternalCreator *Creator
	// Error message in case any failure
	ErrorMsg *string
	// Information for group-to-group connectivity
	ConnectivityInfos []GroupToGroupConnectivityInfo
}

// Status of group-to-group connectivity
const GroupToGroupConnectivity_STATE_CONNECTING = "CONNECTING"

// Status of group-to-group connectivity
const GroupToGroupConnectivity_STATE_CONNECTED = "CONNECTED"

// Status of group-to-group connectivity
const GroupToGroupConnectivity_STATE_CONNECTION_FAILED = "CONNECTION_FAILED"

// Status of group-to-group connectivity
const GroupToGroupConnectivity_STATE_DELETING = "DELETING"

// Status of group-to-group connectivity
const GroupToGroupConnectivity_STATE_DELETED = "DELETED"

// Status of group-to-group connectivity
const GroupToGroupConnectivity_STATE_DELETION_FAILED = "DELETION_FAILED"

func (s *GroupToGroupConnectivity) GetType__() vapiBindings_.BindingType {
	return GroupToGroupConnectivityBindingType()
}

func (s *GroupToGroupConnectivity) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for GroupToGroupConnectivity._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type AvailableGroupForConnectivity struct {
	OrgId *string
	// Identifier of the network connectivity config for group-to-group connectivity
	NetworkConnectivityConfigId *string
	// Identifier of the source group, this is deployment group id of the source group
	GroupId *string
	// Name of the available group(s) for group-to-group connectivity.
	GroupName *string
}

func (s *AvailableGroupForConnectivity) GetType__() vapiBindings_.BindingType {
	return AvailableGroupForConnectivityBindingType()
}

func (s *AvailableGroupForConnectivity) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for AvailableGroupForConnectivity._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type Sddcs struct {
	SddcId   *string
	SddcName *string
	Region   *string
}

func (s *Sddcs) GetType__() vapiBindings_.BindingType {
	return SddcsBindingType()
}

func (s *Sddcs) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for Sddcs._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type DeploymentGroupMember struct {
	// ID of deployment group member
	Id *string
}

func (s *DeploymentGroupMember) GetType__() vapiBindings_.BindingType {
	return DeploymentGroupMemberBindingType()
}

func (s *DeploymentGroupMember) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for DeploymentGroupMember._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type RouteTable struct {
	// Identifier of the route table
	Id *string
	// DisplayName of the route table
	Name     *string
	Location *Location
	// region name (deprecated in favor of ``location``)
	Region *string
	// Description of the route table
	Description *string
	// Identifier of the route table assigned by skynet
	ProviderAssignedId *string
	// Identifier of the route table assigned by cloud provider
	ProviderAssignedRouteTableId *string
	// Identifier of the route table organization
	OrgId *string
	// Identifier of the network connectivity config for this route-table
	NetworkConnectivityConfigId *string
	// Timestamp of last routes sync
	RoutesLastSyncTime *string
	State              *RouteTableSyncState
	Creator            *Creator
	InternalCreator    *Creator
	Updater            *Updater
	// Customer defined resource tags
	Tags map[string]string
	// Internal resource properties
	Properties map[string]string
}

func (s *RouteTable) GetType__() vapiBindings_.BindingType {
	return RouteTableBindingType()
}

func (s *RouteTable) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for RouteTable._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type CreateGroupNetworkConnectivityRequest struct {
	// Name of the deployment group
	Name string
	// Description of the deployment group
	Description *string
	// ID of deployment group members
	Members []DeploymentGroupMember
}

func (s *CreateGroupNetworkConnectivityRequest) GetType__() vapiBindings_.BindingType {
	return CreateGroupNetworkConnectivityRequestBindingType()
}

func (s *CreateGroupNetworkConnectivityRequest) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for CreateGroupNetworkConnectivityRequest._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type NsxDetails struct {
	NsxPrivateIp     *string
	NsxPublicIp      *string
	NsxDefaultAccess *string
	NsxDefaultAuth   *string
	NsxUsers         []NsxCredentialInfo
	NsxPrivateFqdn   *string
	NsxPublicFqdn    *string
	LoginUrls        []NsxLoginInfo
}

func (s *NsxDetails) GetType__() vapiBindings_.BindingType {
	return NsxDetailsBindingType()
}

func (s *NsxDetails) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for NsxDetails._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type State struct {
	// Name of the state name
	Name string
	// Error message on failed states
	ErrorMsg *string
	// Unique code identifying the error
	ErrorCode *string
	// Internal error message on failed states
	InternalErrorMsg *string
	// Internal unique code identifying the error
	InternalErrorCode *string
}

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

type NetworkConnectivityConfig struct {
	// Identifier of the NetworkConnectivityConfig
	Id *string
	// Identifier of deployment group
	GroupId *string
	// Name of the NetworkConnectivityConfig
	Name string
	// Identifier of the NetworkConnectivityConfig organization
	OrgId *string
	// Identifier of a Group to Group Connectivity
	GroupToGroupConnectivityId *string
	Version                    *NetworkConnectivityConfigVersion
	Type_                      *PolicyType
	State                      *ResourceState
	Creator                    *Creator
	Updater                    *Updater
	InternalCreator            *Creator
	// Customer defined resource tags
	Tags map[string]string
	// Internal resource properties
	Properties map[string]string
	// Extended resource traits
	Traits map[string]vapiData_.DataValue
	Trait  []Trait
}

func (s *NetworkConnectivityConfig) GetType__() vapiBindings_.BindingType {
	return NetworkConnectivityConfigBindingType()
}

func (s *NetworkConnectivityConfig) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for NetworkConnectivityConfig._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type Trait struct {
	// ID of trait
	Id *string
	// ID of organization
	OrgId *string
	// Name of trait (deprecated by type)
	Name string
	// Type of trait
	Type_ *string
	// Identifier of the owning resource
	ResourceId *string
	// version of trait
	Version         *int64
	State           *State
	Creator         *Creator
	Updater         *Updater
	InternalCreator *Creator
	// (internal) Property bag used to serialize the derived type
	UnmappedProperties map[string]*vapiData_.StructValue
}

func (s *Trait) GetType__() vapiBindings_.BindingType {
	return TraitBindingType()
}

func (s *Trait) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for Trait._GetDataValue method - %s",
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
}

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

// (deprecated)
type NetworkConnectivityConfigVersion struct {
	// (deprecated) Name of the NetworkConnectivityConfig version
	Name string
	// (deprecated) Unique code identifying the NetworkConnectivityConfig version
	Code string
}

func (s *NetworkConnectivityConfigVersion) GetType__() vapiBindings_.BindingType {
	return NetworkConnectivityConfigVersionBindingType()
}

func (s *NetworkConnectivityConfigVersion) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for NetworkConnectivityConfigVersion._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type RouteTableSyncState struct {
	// State name
	Name string
	// Error message specific to operation
	ErrorMsgs map[string]string
}

// State name
const RouteTableSyncState_NAME_READY = "READY"

// State name
const RouteTableSyncState_NAME_SYNCING = "SYNCING"

// State name
const RouteTableSyncState_NAME_FAILED = "FAILED"

// State name
const RouteTableSyncState_NAME_DELETED = "DELETED"

func (s *RouteTableSyncState) GetType__() vapiBindings_.BindingType {
	return RouteTableSyncStateBindingType()
}

func (s *RouteTableSyncState) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for RouteTableSyncState._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type RouteTableResponse struct {
	// List of RouteTables
	Content  []RouteTable
	Pageable *Pageable
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

func (s *RouteTableResponse) GetType__() vapiBindings_.BindingType {
	return RouteTableResponseBindingType()
}

func (s *RouteTableResponse) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for RouteTableResponse._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type ResourceState struct {
	// Name of the state
	Name string
	// Error message specific to operation
	ErrorMsgs map[string]string
}

// Name of the state
const ResourceState_NAME_PENDING = "PENDING"

// Name of the state
const ResourceState_NAME_UPDATING = "UPDATING"

// Name of the state
const ResourceState_NAME_CONNECTED = "CONNECTED"

// Name of the state
const ResourceState_NAME_FAILED = "FAILED"

// Name of the state
const ResourceState_NAME_DELETING = "DELETING"

// Name of the state
const ResourceState_NAME_DELETED = "DELETED"

// Name of the state
const ResourceState_NAME_DELETION_FAILED = "DELETION_FAILED"

func (s *ResourceState) GetType__() vapiBindings_.BindingType {
	return ResourceStateBindingType()
}

func (s *ResourceState) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ResourceState._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type NsxCredentialInfo struct {
	UserName *string
	Password *string
	// user type
	UserType *string
}

// user type
const NsxCredentialInfo_USER_TYPE_CLOUD_ADMIN = "CLOUD_ADMIN"

// user type
const NsxCredentialInfo_USER_TYPE_CLOUD_AUDIT = "CLOUD_AUDIT"

func (s *NsxCredentialInfo) GetType__() vapiBindings_.BindingType {
	return NsxCredentialInfoBindingType()
}

func (s *NsxCredentialInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for NsxCredentialInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type PolicyType struct {
	// Type name of NetworkConnectivityConfig
	Name string
	// Unique code identifying NetworkConnectivityConfig
	Code *string
}

func (s *PolicyType) GetType__() vapiBindings_.BindingType {
	return PolicyTypeBindingType()
}

func (s *PolicyType) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for PolicyType._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type RouteEntryTarget struct {
	// ID of the target
	Id *string
	// Name of the target
	Name *string
	// Type of the target
	Type_    *string
	Location *Location
	// region name (deprecated in favor of ``location``)
	Region *string
}

// Type of the target
const RouteEntryTarget_TYPE_SDDC = "SDDC"

// Type of the target
const RouteEntryTarget_TYPE_VPC = "VPC"

// Type of the target
const RouteEntryTarget_TYPE_DGW = "DGW"

// Type of the target
const RouteEntryTarget_TYPE_TGW = "TGW"

// Type of the target
const RouteEntryTarget_TYPE_CONNECTED_SDDC_GROUP = "Connected SDDC Group"

func (s *RouteEntryTarget) GetType__() vapiBindings_.BindingType {
	return RouteEntryTargetBindingType()
}

func (s *RouteEntryTarget) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for RouteEntryTarget._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type CreateGroupNetworkConnectivityResponse struct {
	// ID of deployment group
	GroupId *string
	// ID of NetworkConnectivityConfig
	ConfigId *string
	// ID of the applyNetworkConnectivityConfig operation
	OperationId *string
}

func (s *CreateGroupNetworkConnectivityResponse) GetType__() vapiBindings_.BindingType {
	return CreateGroupNetworkConnectivityResponseBindingType()
}

func (s *CreateGroupNetworkConnectivityResponse) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for CreateGroupNetworkConnectivityResponse._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Request payload for validate-members API
type ValidationPayload struct {
	Members []DeploymentGroupMember
	// DeploymentGroupId (optional). To be specified when adding new member(s) to an existing deployentGroup.
	DeploymentGroupId *string
}

func (s *ValidationPayload) GetType__() vapiBindings_.BindingType {
	return ValidationPayloadBindingType()
}

func (s *ValidationPayload) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ValidationPayload._GetDataValue method - %s",
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

type RouteEntryResponse struct {
	// List of RouteEntries
	Content  []RouteEntry
	Pageable *Pageable
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

func (s *RouteEntryResponse) GetType__() vapiBindings_.BindingType {
	return RouteEntryResponseBindingType()
}

func (s *RouteEntryResponse) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for RouteEntryResponse._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type NsxLoginInfo struct {
	// access type of the url
	AccessType *string
	// auth type used for nsx login
	AuthType     *string
	PreferredUrl *string
	OtherUrls    []string
}

// access type of the url
const NsxLoginInfo_ACCESS_TYPE_PRIVATE = "PRIVATE"

// access type of the url
const NsxLoginInfo_ACCESS_TYPE_PUBLIC = "PUBLIC"

// auth type used for nsx login
const NsxLoginInfo_AUTH_TYPE_CSP = "CSP"

// auth type used for nsx login
const NsxLoginInfo_AUTH_TYPE_LOCAL = "LOCAL"

// auth type used for nsx login
const NsxLoginInfo_AUTH_TYPE_DEFAULT = "DEFAULT"

func (s *NsxLoginInfo) GetType__() vapiBindings_.BindingType {
	return NsxLoginInfoBindingType()
}

func (s *NsxLoginInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for NsxLoginInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type ValidationErrorDetail struct {
	// Detailed message for validation error
	ValidationErrorMessage string
	// Localized detailed message for validation error
	LocalizedValidationErrorMessage *string
	Members                         []string
}

func (s *ValidationErrorDetail) GetType__() vapiBindings_.BindingType {
	return ValidationErrorDetailBindingType()
}

func (s *ValidationErrorDetail) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ValidationErrorDetail._GetDataValue method - %s",
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

type AwsSharedPrefixList struct {
	Id                          *string
	OrgId                       *string
	NetworkConnectivityConfigId *string
	// Name of aws prefix list
	Name *string
	// Prefix List IP address type
	AddressFamily *string
	// ID of aws prefix list
	CloudProviderAssignedId      *string
	CloudProviderAssignedVersion *string
	Location                     *Location
	// status of the Shared Prefix List
	State           *string
	Creator         *Creator
	Updater         *Updater
	InternalCreator *Creator
	// prefix list size
	Size *int64
	// total routes from selected SDDC members
	TotalEntries   *int64
	MembershipType *string
	Tags           []map[string]string
	SddcLocation   *Location
	Sddcs          []Sddcs
	Accounts       []SharedPrefixListAccount
	ResourceShare  *SharedPrefixListResourceShare
	Version        *int64
	// Error message specific to this shared prefix list object
	ErrorMsgs *string
}

// Prefix List IP address type
const AwsSharedPrefixList_ADDRESS_FAMILY_IPV4 = "IPv4"

// Prefix List IP address type
const AwsSharedPrefixList_ADDRESS_FAMILY_IPV6 = "IPv6"

// status of the Shared Prefix List
const AwsSharedPrefixList_STATE_CREATE_IN_PROGRESS = "CREATE_IN_PROGRESS"

// status of the Shared Prefix List
const AwsSharedPrefixList_STATE_CREATED = "CREATED"

// status of the Shared Prefix List
const AwsSharedPrefixList_STATE_CREATE_FAILED = "CREATE_FAILED"

// status of the Shared Prefix List
const AwsSharedPrefixList_STATE_MODIFY_IN_PROGRESS = "MODIFY_IN_PROGRESS"

// status of the Shared Prefix List
const AwsSharedPrefixList_STATE_MODIFIED = "MODIFIED"

// status of the Shared Prefix List
const AwsSharedPrefixList_STATE_MODIFY_FAILED = "MODIFY_FAILED"

// status of the Shared Prefix List
const AwsSharedPrefixList_STATE_DELETE_IN_PROGRESS = "DELETE_IN_PROGRESS"

// status of the Shared Prefix List
const AwsSharedPrefixList_STATE_DELETE_FAILED = "DELETE_FAILED"

// status of the Shared Prefix List
const AwsSharedPrefixList_STATE_RESIZE_LIMIT_EXCEEDED = "RESIZE_LIMIT_EXCEEDED"
const AwsSharedPrefixList_MEMBERSHIP_TYPE_STATIC = "STATIC"
const AwsSharedPrefixList_MEMBERSHIP_TYPE_DYNAMIC = "DYNAMIC"

func (s *AwsSharedPrefixList) GetType__() vapiBindings_.BindingType {
	return AwsSharedPrefixListBindingType()
}

func (s *AwsSharedPrefixList) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for AwsSharedPrefixList._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type SharedPrefixListEntryResponse struct {
	// List of Shared Prefix List Route Entries
	Content  []SharedPrefixListEntry
	Pageable *Pageable
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

func (s *SharedPrefixListEntryResponse) GetType__() vapiBindings_.BindingType {
	return SharedPrefixListEntryResponseBindingType()
}

func (s *SharedPrefixListEntryResponse) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for SharedPrefixListEntryResponse._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type SharedPrefixListEntry struct {
	Id           *string
	OrgId        *string
	SharedPlId   string
	Cidr         *string
	IsDuplicated *bool
	Target       *RouteEntryTarget
	Creator      *Creator
	// Error message specific to this shared prefix list entry
	ErrorMsg *string
}

func (s *SharedPrefixListEntry) GetType__() vapiBindings_.BindingType {
	return SharedPrefixListEntryBindingType()
}

func (s *SharedPrefixListEntry) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for SharedPrefixListEntry._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type CustomerNetworkPeeringInfo struct {
	CustomerNetworkPeerings []MetalConnection
}

func (s *CustomerNetworkPeeringInfo) GetType__() vapiBindings_.BindingType {
	return CustomerNetworkPeeringInfoBindingType()
}

func (s *CustomerNetworkPeeringInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for CustomerNetworkPeeringInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Network Activity response
type NetworkActivity struct {
	Id                   *string
	OrgId                *string
	DeploymentId         *string
	ResourceId           *string
	ResourceType         *string
	State                *string
	Error_               *ExecutionError
	ActivityTypeKey      *string
	CreatedTimestamp     *string
	LastUpdatedTimestamp *string
}

const NetworkActivity_RESOURCE_TYPE_CUSTOMER_NETWORK_PEERING = "CUSTOMER_NETWORK_PEERING"
const NetworkActivity_RESOURCE_TYPE_PUBLIC_IP = "PUBLIC_IP"
const NetworkActivity_STATE_INITIALIZED = "INITIALIZED"
const NetworkActivity_STATE_SUBMITTED = "SUBMITTED"
const NetworkActivity_STATE_RUNNING = "RUNNING"
const NetworkActivity_STATE_COMPLETED = "COMPLETED"
const NetworkActivity_STATE_CANCELED = "CANCELED"
const NetworkActivity_STATE_FAILED = "FAILED"
const NetworkActivity_STATE_SUSPENDED = "SUSPENDED"
const NetworkActivity_STATE_UNRECOGNIZED = "UNRECOGNIZED"
const NetworkActivity_STATE_RESTARTING = "RESTARTING"
const NetworkActivity_STATE_PENDING_PAUSE = "PENDING_PAUSE"
const NetworkActivity_STATE_PENDING_RETRY = "PENDING_RETRY"
const NetworkActivity_STATE_PENDING_RESUME = "PENDING_RESUME"
const NetworkActivity_STATE_PENDING_CANCEL = "PENDING_CANCEL"

func (s *NetworkActivity) GetType__() vapiBindings_.BindingType {
	return NetworkActivityBindingType()
}

func (s *NetworkActivity) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for NetworkActivity._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Object to store details about the activity error
type ExecutionError struct {
	// Error code
	Code *string
	// Error status for the failed states
	Status *string
	// Error message for the failed states
	Message *string
}

func (s *ExecutionError) GetType__() vapiBindings_.BindingType {
	return ExecutionErrorBindingType()
}

func (s *ExecutionError) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ExecutionError._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type NsxSettings struct {
	NsxManagerAccessMode *string
}

const NsxSettings_NSX_MANAGER_ACCESS_MODE_PUBLIC = "PUBLIC"
const NsxSettings_NSX_MANAGER_ACCESS_MODE_PRIVATE = "PRIVATE"

func (s *NsxSettings) GetType__() vapiBindings_.BindingType {
	return NsxSettingsBindingType()
}

func (s *NsxSettings) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for NsxSettings._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type NsxAddonResponse struct {
	Name     *string
	Action   *string
	ErrorMsg *string
}

func (s *NsxAddonResponse) GetType__() vapiBindings_.BindingType {
	return NsxAddonResponseBindingType()
}

func (s *NsxAddonResponse) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for NsxAddonResponse._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type Vpn struct {
	ConnectivityDown        *int64
	ConnectivityUp          *int64
	ConnectivityDeactivated *int64
	TotalConnections        *int64
	PublicIp                *string
	VpnDownConnections      []string
}

func (s *Vpn) GetType__() vapiBindings_.BindingType {
	return VpnBindingType()
}

func (s *Vpn) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for Vpn._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type DirectConnect struct {
	TotalAttachedVifsCount *int64
	ConnectivityDown       *int64
	ConnectivityUp         *int64
	LearnedRoutes          *int64
	AdvertisedRoutes       *int64
	DownVifNames           []string
}

func (s *DirectConnect) GetType__() vapiBindings_.BindingType {
	return DirectConnectBindingType()
}

func (s *DirectConnect) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for DirectConnect._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type TransitConnect struct {
	ConnectedSddcGroupId   *string
	ConnectedSddcGroupName *string
	LearnedRoutes          *int64
	AdvertisedRoutes       *int64
}

func (s *TransitConnect) GetType__() vapiBindings_.BindingType {
	return TransitConnectBindingType()
}

func (s *TransitConnect) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for TransitConnect._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type ManagementGateway struct {
	TotalFirewallRules *int64
	MgmtSubnets        []string
	ApplianceSubnet    *string
	InfraSubnets       []string
	Dns                []string
}

func (s *ManagementGateway) GetType__() vapiBindings_.BindingType {
	return ManagementGatewayBindingType()
}

func (s *ManagementGateway) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ManagementGateway._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type ComputeGateway struct {
	TotalFirewallRules *int64
	TotalSegments      *int64
	Dns                []string
	TotalTier1Gateways *int64
}

func (s *ComputeGateway) GetType__() vapiBindings_.BindingType {
	return ComputeGatewayBindingType()
}

func (s *ComputeGateway) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ComputeGateway._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type Vpc struct {
	Id                 *string
	AvailabilityZone   *string
	AvailabilityZoneId *string
	Cidr               *string
}

func (s *Vpc) GetType__() vapiBindings_.BindingType {
	return VpcBindingType()
}

func (s *Vpc) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for Vpc._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type CloudProvider struct {
	AccountId *string
	// Name of the state
	ConnectivityStatus          *string
	TrafficGroupEniMappingInfos []TrafficGroupEniMappingInfos
	Vpc                         []Vpc
	CloudFormationStackName     *string
	S3Enabled                   *bool
	Ec2Enabled                  *bool
	IamRoleNames                *string
	VpcId                       *string
	VpcCidr                     []string
}

// Name of the state
const CloudProvider_CONNECTIVITY_STATUS_CONNECTED = "CONNECTED"

// Name of the state
const CloudProvider_CONNECTIVITY_STATUS_DISCONNECTED = "DISCONNECTED"

func (s *CloudProvider) GetType__() vapiBindings_.BindingType {
	return CloudProviderBindingType()
}

func (s *CloudProvider) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for CloudProvider._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type DeploymentNetworkingStateInfo struct {
	Id                *string
	Name              *string
	Location          *Location
	SddcId            *string
	OrgId             *string
	ProviderCode      *string
	ProviderType      *string
	LastSyncTime      *string
	Vpn               []Vpn
	DirectConnect     []DirectConnect
	OutpostsConnect   []OutpostConnect
	TransitConnect    []TransitConnect
	ManagementGateway []ManagementGateway
	ComputeGateway    []ComputeGateway
	CloudProvider     []CloudProvider
	Error_            *ErrorMessage
}

func (s *DeploymentNetworkingStateInfo) GetType__() vapiBindings_.BindingType {
	return DeploymentNetworkingStateInfoBindingType()
}

func (s *DeploymentNetworkingStateInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for DeploymentNetworkingStateInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type ErrorMessage struct {
	// Name of the state
	Code       *string
	RetryCount *int64
	Message    *string
}

// Name of the state
const ErrorMessage_CODE_VERSION_NOT_SUPPORTED = "VERSION_NOT_SUPPORTED"

// Name of the state
const ErrorMessage_CODE_OUT_OF_SYNC = "OUT_OF_SYNC"

func (s *ErrorMessage) GetType__() vapiBindings_.BindingType {
	return ErrorMessageBindingType()
}

func (s *ErrorMessage) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ErrorMessage._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type OutpostConnect struct {
	LearnedRoutes    *int64
	AdvertisedRoutes *int64
}

func (s *OutpostConnect) GetType__() vapiBindings_.BindingType {
	return OutpostConnectBindingType()
}

func (s *OutpostConnect) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for OutpostConnect._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type MetalServiceToken struct {
	Id *string
	// Token type - a_side or z_side
	ServiceTokenType *string
	// Maximum speed for the interconnection
	MaxAllowedSpeed *int64
	// Token expiry date/time
	ExpiresAt *string
	// Token role ex: primary
	Role *string
}

// Token type - a_side or z_side
const MetalServiceToken_SERVICE_TOKEN_TYPE_A_SIDE = "A_SIDE"

// Token type - a_side or z_side
const MetalServiceToken_SERVICE_TOKEN_TYPE_Z_SIDE = "Z_SIDE"

func (s *MetalServiceToken) GetType__() vapiBindings_.BindingType {
	return MetalServiceTokenBindingType()
}

func (s *MetalServiceToken) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for MetalServiceToken._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type MetalVirtualCircuit struct {
	Id                *string
	Name              *string
	Vrf               *string
	Status            *string
	ConnectionAuthKey *string
	BgpPeerIp         *string
	BgpPeerNetMask    *int64
	MetalBgpPeerIp    *string
	BgpPeerAsn        *int64
}

func (s *MetalVirtualCircuit) GetType__() vapiBindings_.BindingType {
	return MetalVirtualCircuitBindingType()
}

func (s *MetalVirtualCircuit) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for MetalVirtualCircuit._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type MetalPort struct {
	Id             *string
	Name           *string
	Role           *string
	Speed          *int64
	VirtualCircuit *MetalVirtualCircuit
	ServiceToken   *MetalServiceToken
}

func (s *MetalPort) GetType__() vapiBindings_.BindingType {
	return MetalPortBindingType()
}

func (s *MetalPort) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for MetalPort._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type MetalConnection struct {
	Id           *string
	ConnectionId *string
	Name         *string
	Description  *string
	// Connection type in VMC context
	VmcConnectionType *string
	// Connection redundancy type
	Redundancy *string
	// Connection type in Equinix context (shared or dedicated)
	Type_ *string
	Metro *string
	Ports []MetalPort
}

// Connection type in VMC context
const MetalConnection_VMC_CONNECTION_TYPE_INTERNAL_POP = "INTERNAL_POP"

// Connection type in VMC context
const MetalConnection_VMC_CONNECTION_TYPE_CUSTOMER = "CUSTOMER"

// Connection redundancy type
const MetalConnection_REDUNDANCY_PRIMARY = "PRIMARY"

// Connection redundancy type
const MetalConnection_REDUNDANCY_REDUNDANT = "REDUNDANT"

// Connection type in Equinix context (shared or dedicated)
const MetalConnection_TYPE_SHARED = "SHARED"

// Connection type in Equinix context (shared or dedicated)
const MetalConnection_TYPE_DEDICATED = "DEDICATED"

func (s *MetalConnection) GetType__() vapiBindings_.BindingType {
	return MetalConnectionBindingType()
}

func (s *MetalConnection) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for MetalConnection._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type PublicIp struct {
	// Atlas generated resource id
	Id *string
	// EIP reservation id
	ReservationId *string
	// EIP attachment id
	AttachmentId *string
	// state of ip assignment
	State *string
	// CIDR for ip assignment
	Cidr *int64
	// IP Address
	Address *string
	// IP Address
	NextHop *string
	// Gateway Address
	Gateway *string
	// Network Address
	Network *string
	// description of IP
	Description *string
}

// state of ip assignment
const PublicIp_STATE_PENDING = "PENDING"

// state of ip assignment
const PublicIp_STATE_ACTIVE = "ACTIVE"

// state of ip assignment
const PublicIp_STATE_DELETING = "DELETING"

func (s *PublicIp) GetType__() vapiBindings_.BindingType {
	return PublicIpBindingType()
}

func (s *PublicIp) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for PublicIp._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type AssignElasticIpRequest struct {
	// description of ElasticIp
	Description *string
}

func (s *AssignElasticIpRequest) GetType__() vapiBindings_.BindingType {
	return AssignElasticIpRequestBindingType()
}

func (s *AssignElasticIpRequest) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for AssignElasticIpRequest._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type PrivateIp struct {
	// Atlas generated resource id
	Id *string
	// privateIp current state
	State *string
	// privateIp address
	Address *string
	// subnet leveraged to assign privateIp
	SubnetName *string
	// ip address family categorization
	AddressFamily *string
	// source name requesting the privateIp
	RequestSource *string
	// description of privateIp
	Description *string
}

// privateIp current state
const PrivateIp_STATE_ASSIGNED = "ASSIGNED"

// privateIp current state
const PrivateIp_STATE_RESERVED = "RESERVED"

// privateIp current state
const PrivateIp_STATE_DELETED = "DELETED"

// subnet leveraged to assign privateIp
const PrivateIp_SUBNET_NAME_SECOND_THIRD_PARTY_APPLIANCES_CIDR = "SECOND_THIRD_PARTY_APPLIANCES_CIDR"

// ip address family categorization
const PrivateIp_ADDRESS_FAMILY_IPV4 = "IPv4"

// ip address family categorization
const PrivateIp_ADDRESS_FAMILY_IPV6 = "IPv6"

// source name requesting the privateIp
const PrivateIp_REQUEST_SOURCE_VSR = "VSR"

func (s *PrivateIp) GetType__() vapiBindings_.BindingType {
	return PrivateIpBindingType()
}

func (s *PrivateIp) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for PrivateIp._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type AssignPrivateIpRequest struct {
	// subnet leveraged to assign privateIp
	SubnetName *string
	// source requesting a privateIp
	RequestSource *string
	// description of the privateIp
	Description *string
}

// subnet leveraged to assign privateIp
const AssignPrivateIpRequest_SUBNET_NAME_SECOND_THIRD_PARTY_APPLIANCES_CIDR = "SECOND_THIRD_PARTY_APPLIANCES_CIDR"

// source requesting a privateIp
const AssignPrivateIpRequest_REQUEST_SOURCE_VSR = "VSR"

func (s *AssignPrivateIpRequest) GetType__() vapiBindings_.BindingType {
	return AssignPrivateIpRequestBindingType()
}

func (s *AssignPrivateIpRequest) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for AssignPrivateIpRequest._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// This object is used for validating Equinix Network Properties
type NetworkProperties struct {
	// management subnet
	ManagementSubnet string
	// vxlan subnet
	VxlanSubnet *string
	// equinix metal router asn
	EquinixMetalRouterAsn string
	// nsx edge uplink router asn
	NsxEdgeUplinkRouterAsn string
}

func (s *NetworkProperties) GetType__() vapiBindings_.BindingType {
	return NetworkPropertiesBindingType()
}

func (s *NetworkProperties) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for NetworkProperties._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type CreateNsxTraceflowRequest struct {
	// ID of the network connectivity config
	NetworkConnectivityConfigId *string
	// sddc traceflow name
	Name        *string
	Source      *NsxTraceflowEndpoint
	Destination *NsxTraceflowEndpoint
	Packet      *PacketData
	SourceId    *NsxTraceflowSourceId
	// Timeout for traceflow observation results. Maximum time in seconds the management plane will wait for observation result to be generated. The default, minimum and maximum timeout values, in seconds, for: Single site environment - minimum 5, default 10, maximum 15. Federated enviroment - minimum 15, default 30, maximum 60. These values are validated by the system based on type of environment.
	Timeout *int64
}

func (s *CreateNsxTraceflowRequest) GetType__() vapiBindings_.BindingType {
	return CreateNsxTraceflowRequestBindingType()
}

func (s *CreateNsxTraceflowRequest) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for CreateNsxTraceflowRequest._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Indicates the type of individual traceflow and provides an ID for the source/destination of the traceflow
type NsxTraceflowEndpoint struct {
	// ID of the traceflow endpoint
	Id *string
	// Type indicator of the traceflow endpoint
	TraceflowType *string
}

// Type indicator of the traceflow endpoint
const NsxTraceflowEndpoint_TRACEFLOW_TYPE_SDDC = "SDDC"

func (s *NsxTraceflowEndpoint) GetType__() vapiBindings_.BindingType {
	return NsxTraceflowEndpointBindingType()
}

func (s *NsxTraceflowEndpoint) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for NsxTraceflowEndpoint._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Contains a pre-populated source port path or indicates that it must be populated at the time of executing the individual traceflow
type NsxTraceflowSourceId struct {
	// Source ID
	Id *string
	// Indicates if source Dneeds to be populated. If type is ACTIVE_EDGE_UPLINK, then "id" field needs to be populated
	SourceIdType *string
}

// Indicates if source Dneeds to be populated. If type is ACTIVE_EDGE_UPLINK, then "id" field needs to be populated
const NsxTraceflowSourceId_SOURCE_ID_TYPE_VALUE = "VALUE"

// Indicates if source Dneeds to be populated. If type is ACTIVE_EDGE_UPLINK, then "id" field needs to be populated
const NsxTraceflowSourceId_SOURCE_ID_TYPE_ACTIVE_EDGE_UPLINK = "ACTIVE_EDGE_UPLINK"

func (s *NsxTraceflowSourceId) GetType__() vapiBindings_.BindingType {
	return NsxTraceflowSourceIdBindingType()
}

func (s *NsxTraceflowSourceId) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for NsxTraceflowSourceId._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type SddcTraceflow struct {
	// ID of the traceflow
	Id *string
	// Human readable ID of the traceflow
	Name *string
	// ID of the org
	OrgId *string
	// ID of the network connectivity config
	NetworkConnectivityConfigId *string
	SourceId                    *NsxTraceflowSourceId
	// Status of the traceflow
	State        *string
	Packet       *PacketData
	Source       *NsxTraceflowEndpoint
	Destination  *NsxTraceflowEndpoint
	Observations *TraceflowObservationListResult
	// Timeout for traceflow observation results. Maximum time in seconds the management plane will wait for observation result to be generated. The default, minimum and maximum timeout values, in seconds, for: Single site environment - minimum 5, default 10, maximum 15. Federated enviroment - minimum 15, default 30, maximum 60. These values are validated by the system based on type of environment.
	Timeout *int64
	Creator *Creator
	Updater *Updater
	// Error messages specific to this traceflow
	ErrorMsgs *string
	// Error message from worker service traceflow realization to be returned to the user
	RealizationErrorMsg *string
}

// Status of the traceflow
const SddcTraceflow_STATE_CREATE_IN_PROGRESS = "CREATE_IN_PROGRESS"

// Status of the traceflow
const SddcTraceflow_STATE_CREATED = "CREATED"

// Status of the traceflow
const SddcTraceflow_STATE_CREATE_FAILED = "CREATE_FAILED"

// Status of the traceflow
const SddcTraceflow_STATE_DELETE_FAILED = "DELETE_FAILED"

func (s *SddcTraceflow) GetType__() vapiBindings_.BindingType {
	return SddcTraceflowBindingType()
}

func (s *SddcTraceflow) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for SddcTraceflow._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type PacketData struct {
	// Packet configuration
	ResourceType *string
	ArpHeader    *ArpHeader
	EthHeader    *EthernetHeader
	IpHeader     *Ipv4Header
	Ipv6Header   *Ipv6Header
	// RFC3548 compatible base64-encoded payload. Up to 1000 bytes of payload may be supplied (with a base64-encoded length of 1336 bytes.) Additional bytes of traceflow metadata will be appended to the payload. The payload contains any data the user wants to put after the transport header.
	Payload         *string
	TransportHeader *TransportProtocolHeader
	// Requested total size of the (logical) packet in bytes. If the requested frame_size is too small (given the payload and traceflow metadata requirement of 16 bytes), the traceflow request will fail with an appropriate message. The frame will be zero padded to the requested size.
	FrameSize *int64
	// A flag, when set true, indicates that the traceflow packet is of L3 routing.
	Routed *bool
	// transport type of the traceflow packet
	TransportType *string
}

// Packet configuration
const PacketData_RESOURCE_TYPE_BINARY_PACKET_DATA = "BinaryPacketData"

// Packet configuration
const PacketData_RESOURCE_TYPE_FIELDS_PACKET_DATA = "FieldsPacketData"

// transport type of the traceflow packet
const PacketData_TRANSPORT_TYPE_UNICAST = "UNICAST"

func (s *PacketData) GetType__() vapiBindings_.BindingType {
	return PacketDataBindingType()
}

func (s *PacketData) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for PacketData._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type Ipv4Header struct {
	// The source IPv4 address
	SrcIp *string
	// The destination IPv4 address
	DstIp *string
	// This is used together with src_ip to calculate dst_ip for broadcast when dst_ip is not given; not used in all other cases.
	SrcSubnetPrefixLen *int64
	// IP protocol - defaults to ICMP
	Protocol *int64
	// Time to live (ttl)
	Ttl *int64
	// IP flags
	Flags *int64
}

func (s *Ipv4Header) GetType__() vapiBindings_.BindingType {
	return Ipv4HeaderBindingType()
}

func (s *Ipv4Header) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for Ipv4Header._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type Ipv6Header struct {
	// The source IPv6 address
	SrcIp *string
	// The destination IPv6 address
	DstIp *string
	// Decremented by 1 by each node that forwards the packets. The packet is discarded if Hop Limit is decremented to zero.
	HopLimit *int64
	// Identifies the type of header immediately following the IPv6 header.
	NextHeader *int64
}

func (s *Ipv6Header) GetType__() vapiBindings_.BindingType {
	return Ipv6HeaderBindingType()
}

func (s *Ipv6Header) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for Ipv6Header._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type ArpHeader struct {
	DstIp string
	// This field specifies the nature of the Arp message being sent.
	OpCode string
	// This field specifies the IP address of the sender. If omitted, the src_ip is set to 0.0.0.0.
	SrcIp *string
}

// This field specifies the nature of the Arp message being sent.
const ArpHeader_OP_CODE_ARP_REQUEST = "ARP_REQUEST"

// This field specifies the nature of the Arp message being sent.
const ArpHeader_OP_CODE_ARP_REPLY = "ARP_REPLY"

func (s *ArpHeader) GetType__() vapiBindings_.BindingType {
	return ArpHeaderBindingType()
}

func (s *ArpHeader) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ArpHeader._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type EthernetHeader struct {
	// Source MAC address of the Ethernet header
	SrcMac *string
	// Destination MAC address of the Ethernet header
	DstMac *string
	// The value of the type field to be put into the Ethernet header. This field defaults to IPv4.
	EthType *int64
}

func (s *EthernetHeader) GetType__() vapiBindings_.BindingType {
	return EthernetHeaderBindingType()
}

func (s *EthernetHeader) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for EthernetHeader._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type TransportProtocolHeader struct {
	IcmpEchoRequestHeader *IcmpEchoRequestHeader
	UdpHeader             *UdpHeader
	TcpHeader             *TcpHeader
	DhcpHeader            *DhcpHeader
	Dhcpv6Header          *Dhcpv6Header
	DnsHeader             *DnsHeader
	NdpHeader             *NdpHeader
}

func (s *TransportProtocolHeader) GetType__() vapiBindings_.BindingType {
	return TransportProtocolHeaderBindingType()
}

func (s *TransportProtocolHeader) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for TransportProtocolHeader._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type IcmpEchoRequestHeader struct {
	// ICMP Id
	Id *int64
	// ICMP sequence number
	Sequence *int64
}

func (s *IcmpEchoRequestHeader) GetType__() vapiBindings_.BindingType {
	return IcmpEchoRequestHeaderBindingType()
}

func (s *IcmpEchoRequestHeader) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for IcmpEchoRequestHeader._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type UdpHeader struct {
	// Source port of udp header
	SrcPort *int64
	// Destination port of udp header
	DstPort *int64
}

func (s *UdpHeader) GetType__() vapiBindings_.BindingType {
	return UdpHeaderBindingType()
}

func (s *UdpHeader) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for UdpHeader._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type TcpHeader struct {
	// Source port of tcp header
	SrcPort *int64
	// Destination port of tcp header
	DstPort *int64
	// TCP flags
	TcpFlags *int64
}

func (s *TcpHeader) GetType__() vapiBindings_.BindingType {
	return TcpHeaderBindingType()
}

func (s *TcpHeader) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for TcpHeader._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type DhcpHeader struct {
	// Message op code / message type
	OpCode *string
}

// Message op code / message type
const DhcpHeader_OP_CODE_BOOTREQUEST = "BOOTREQUEST"

// Message op code / message type
const DhcpHeader_OP_CODE_BOOTREPLY = "BOOTREPLY"

func (s *DhcpHeader) GetType__() vapiBindings_.BindingType {
	return DhcpHeaderBindingType()
}

func (s *DhcpHeader) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for DhcpHeader._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type Dhcpv6Header struct {
	// DHCP message type
	MsgType *string
}

// DHCP message type
const Dhcpv6Header_MSG_TYPE_SOLICIT = "SOLICIT"

// DHCP message type
const Dhcpv6Header_MSG_TYPE_ADVERTISE = "ADVERTISE"

// DHCP message type
const Dhcpv6Header_MSG_TYPE_REQUEST = "REQUEST"

// DHCP message type
const Dhcpv6Header_MSG_TYPE_REPLY = "REPLY"

func (s *Dhcpv6Header) GetType__() vapiBindings_.BindingType {
	return Dhcpv6HeaderBindingType()
}

func (s *Dhcpv6Header) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for Dhcpv6Header._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type DnsHeader struct {
	// Domain name/IP to query/response
	Address *string
	// This is used to specify the type of the address
	AddressType *string
	// Specifies the message type whether it is a query or a response.
	MessageType *string
}

// This is used to specify the type of the address
const DnsHeader_ADDRESS_TYPE_V4 = "V4"

// This is used to specify the type of the address
const DnsHeader_ADDRESS_TYPE_V6 = "V6"

// Specifies the message type whether it is a query or a response.
const DnsHeader_MESSAGE_TYPE_QUERY = "QUERY"

// Specifies the message type whether it is a query or a response.
const DnsHeader_MESSAGE_TYPE_RESPONSE = "RESPONSE"

func (s *DnsHeader) GetType__() vapiBindings_.BindingType {
	return DnsHeaderBindingType()
}

func (s *DnsHeader) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for DnsHeader._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type NdpHeader struct {
	// The destination IP address
	DstIp *string
	// NDP message type
	MsgType *string
}

// NDP message type
const NdpHeader_MSG_TYPE_NEIGHBOR_SOLICITATION = "NEIGHBOR_SOLICITATION"

// NDP message type
const NdpHeader_MSG_TYPE_NEIGHBOR_ADVERTISEMENT = "NEIGHBOR_ADVERTISEMENT"

func (s *NdpHeader) GetType__() vapiBindings_.BindingType {
	return NdpHeaderBindingType()
}

func (s *NdpHeader) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for NdpHeader._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type NsxVmList struct {
	// List of virtual machines
	VmList []NsxVm
}

func (s *NsxVmList) GetType__() vapiBindings_.BindingType {
	return NsxVmListBindingType()
}

func (s *NsxVmList) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for NsxVmList._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type NsxVm struct {
	// Current external id of this virtual machine in the system
	VmId *string
	// Identifier to use when displaying entity in logs or GUI
	DisplayName *string
	// Id of the host in which this virtual machine exists
	HostId *string
	// List of virtual interfaces associated with this virtual machine
	VirtualInterfaces []NsxVirtualInterface
}

func (s *NsxVm) GetType__() vapiBindings_.BindingType {
	return NsxVmBindingType()
}

func (s *NsxVm) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for NsxVm._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type NsxVirtualInterface struct {
	// External Id of the virtual network inferface
	VifId *string
	// Identifier to use when displaying entity in logs or GUI
	DisplayName *string
	// IP Addresses of the the virtual network interface, from source VM_TOOLS
	IpAddresses []NsxVifIpAddress
	// MAC address of the virtual network interface
	MacAddress *string
	// LPort Attachment Id of the virtual network interface
	AttachmentId *string
	// Id of the segment port associated with to the virtual network interface
	SegmentPortId *string
	// Path of the segment port associated with the virtual network interface
	SegmentPortPath *string
}

func (s *NsxVirtualInterface) GetType__() vapiBindings_.BindingType {
	return NsxVirtualInterfaceBindingType()
}

func (s *NsxVirtualInterface) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for NsxVirtualInterface._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type TraceflowObservationListResult struct {
	// TraceflowObservation list results
	ResultCount *int64
	Results     []TraceflowObservation
	// Error message from traceflow realization
	RealizationErrorMsg *string
}

func (s *TraceflowObservationListResult) GetType__() vapiBindings_.BindingType {
	return TraceflowObservationListResultBindingType()
}

func (s *TraceflowObservationListResult) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for TraceflowObservationListResult._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type TraceflowObservation struct {
	// The name of the component that issued the observation.
	ComponentName *string
	// The sub type of the component that issued the observation.
	ComponentSubType *string
	ComponentType    *string
	ComponentPath    *string
	InterfacePath    *string
	ResourceType     string
	// the sequence number is the traceflow observation hop count
	SequenceNo int64
	// the site path where this observation was generated
	SitePath *string
	// Timestamp when the observation was created by the transport node (milliseconds epoch)
	Timestamp *int64
	// Timestamp when the observation was created by the transport node (microseconds epoch)
	TimestampMicro *int64
	// id of the transport node that observed a traceflow packet
	TransportNodeId *string
	// name of the transport node that observed a traceflow packet
	TransportNodeName *string
	// type of the transport node that observed a traceflow packet
	TransportNodeType *string
}

// The sub type of the component that issued the observation.
const TraceflowObservation_COMPONENT_SUB_TYPE_LR_TIER0 = "LR_TIER0"

// The sub type of the component that issued the observation.
const TraceflowObservation_COMPONENT_SUB_TYPE_LR_TIER1 = "LR_TIER1"

// The sub type of the component that issued the observation.
const TraceflowObservation_COMPONENT_SUB_TYPE_LR_VRF_TIER0 = "LR_VRF_TIER0"

// The sub type of the component that issued the observation.
const TraceflowObservation_COMPONENT_SUB_TYPE_LS_TRANSIT = "LS_TRANSIT"

// The sub type of the component that issued the observation.
const TraceflowObservation_COMPONENT_SUB_TYPE_SI_CLASSIFIER = "SI_CLASSIFIER"

// The sub type of the component that issued the observation.
const TraceflowObservation_COMPONENT_SUB_TYPE_SI_PROXY = "SI_PROXY"

// The sub type of the component that issued the observation.
const TraceflowObservation_COMPONENT_SUB_TYPE_VDR = "VDR"

// The sub type of the component that issued the observation.
const TraceflowObservation_COMPONENT_SUB_TYPE_ENI = "ENI"

// The sub type of the component that issued the observation.
const TraceflowObservation_COMPONENT_SUB_TYPE_AWS_GATEWAY = "AWS_GATEWAY"

// The sub type of the component that issued the observation.
const TraceflowObservation_COMPONENT_SUB_TYPE_TGW_ROUTE = "TGW_ROUTE"

// The sub type of the component that issued the observation.
const TraceflowObservation_COMPONENT_SUB_TYPE_EDGE_UPLINK = "EDGE_UPLINK"

// The sub type of the component that issued the observation.
const TraceflowObservation_COMPONENT_SUB_TYPE_DELL_GATEWAY = "DELL_GATEWAY"

// The sub type of the component that issued the observation.
const TraceflowObservation_COMPONENT_SUB_TYPE_UNKNOWN = "UNKNOWN"
const TraceflowObservation_COMPONENT_TYPE_TRACEFLOW_COMPONENT_TYPE_PHYSICAL = "PHYSICAL"
const TraceflowObservation_COMPONENT_TYPE_TRACEFLOW_COMPONENT_TYPE_LR = "LR"
const TraceflowObservation_COMPONENT_TYPE_TRACEFLOW_COMPONENT_TYPE_LS = "LS"
const TraceflowObservation_COMPONENT_TYPE_TRACEFLOW_COMPONENT_TYPE_DFW = "DFW"
const TraceflowObservation_COMPONENT_TYPE_TRACEFLOW_COMPONENT_TYPE_BRIDGE = "BRIDGE"
const TraceflowObservation_COMPONENT_TYPE_TRACEFLOW_COMPONENT_TYPE_EDGE_TUNNEL = "EDGE_TUNNEL"
const TraceflowObservation_COMPONENT_TYPE_TRACEFLOW_COMPONENT_TYPE_EDGE_HOSTSWITCH = "EDGE_HOSTSWITCH"
const TraceflowObservation_COMPONENT_TYPE_TRACEFLOW_COMPONENT_TYPE_FW_BRIDGE = "FW_BRIDGE"
const TraceflowObservation_COMPONENT_TYPE_TRACEFLOW_COMPONENT_TYPE_LOAD_BALANCER = "LOAD_BALANCER"
const TraceflowObservation_COMPONENT_TYPE_TRACEFLOW_COMPONENT_TYPE_NAT = "NAT"
const TraceflowObservation_COMPONENT_TYPE_TRACEFLOW_COMPONENT_TYPE_IPSEC = "IPSEC"
const TraceflowObservation_COMPONENT_TYPE_TRACEFLOW_COMPONENT_TYPE_SERVICE_INSERTION = "SERVICE_INSERTION"
const TraceflowObservation_COMPONENT_TYPE_TRACEFLOW_COMPONENT_TYPE_VMC = "VMC"
const TraceflowObservation_COMPONENT_TYPE_TRACEFLOW_COMPONENT_TYPE_SPOOFGUARD = "SPOOFGUARD"
const TraceflowObservation_COMPONENT_TYPE_TRACEFLOW_COMPONENT_TYPE_EDGE_FW = "EDGE_FW"
const TraceflowObservation_COMPONENT_TYPE_TRACEFLOW_COMPONENT_TYPE_UNKNOWN = "UNKNOWN"
const TraceflowObservation_RESOURCE_TYPE_TRACEFLOW_OBSERVATION_FORWARDED = "TraceflowObservationForwarded"
const TraceflowObservation_RESOURCE_TYPE_TRACEFLOW_OBSERVATION_DROPPED = "TraceflowObservationDropped"
const TraceflowObservation_RESOURCE_TYPE_TRACEFLOW_OBSERVATION_DELIVERED = "TraceflowObservationDelivered"
const TraceflowObservation_RESOURCE_TYPE_TRACEFLOW_OBSERVATION_RECEIVED = "TraceflowObservationReceived"
const TraceflowObservation_RESOURCE_TYPE_TRACEFLOW_OBSERVATION_FORWARDED_LOGICAL = "TraceflowObservationForwardedLogical"
const TraceflowObservation_RESOURCE_TYPE_TRACEFLOW_OBSERVATION_DROPPED_LOGICAL = "TraceflowObservationDroppedLogical"
const TraceflowObservation_RESOURCE_TYPE_TRACEFLOW_OBSERVATION_RECEIVED_LOGICAL = "TraceflowObservationReceivedLogical"
const TraceflowObservation_RESOURCE_TYPE_TRACEFLOW_OBSERVATION_REPLICATION_LOGICAL = "TraceflowObservationReplicationLogical"
const TraceflowObservation_RESOURCE_TYPE_TRACEFLOW_OBSERVATION_RELAYED_LOGICAL = "TraceflowObservationRelayedLogical"

// type of the transport node that observed a traceflow packet
const TraceflowObservation_TRANSPORT_NODE_TYPE_ESX = "ESX"

// type of the transport node that observed a traceflow packet
const TraceflowObservation_TRANSPORT_NODE_TYPE_RHELKVM = "RHELKVM"

// type of the transport node that observed a traceflow packet
const TraceflowObservation_TRANSPORT_NODE_TYPE_UBUNTUKVM = "UBUNTUKVM"

// type of the transport node that observed a traceflow packet
const TraceflowObservation_TRANSPORT_NODE_TYPE_EDGE = "EDGE"

// type of the transport node that observed a traceflow packet
const TraceflowObservation_TRANSPORT_NODE_TYPE_PUBLIC_CLOUD_GATEWAY_NODE = "PUBLIC_CLOUD_GATEWAY_NODE"

// type of the transport node that observed a traceflow packet
const TraceflowObservation_TRANSPORT_NODE_TYPE_OTHERS = "OTHERS"

// type of the transport node that observed a traceflow packet
const TraceflowObservation_TRANSPORT_NODE_TYPE_HYPERV = "HYPERV"

func (s *TraceflowObservation) GetType__() vapiBindings_.BindingType {
	return TraceflowObservationBindingType()
}

func (s *TraceflowObservation) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for TraceflowObservation._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type AddCustomerNetworkPeeringRequest struct {
	// Metal Connection name
	Name *string
	// Metal Connection description
	Description *string
	// Redundancy
	FabricVcRedundancy      *string
	CustomerBgpPeerInfoList []CustomerBgpPeerInfo
}

// Redundancy
const AddCustomerNetworkPeeringRequest_FABRIC_VC_REDUNDANCY_REDUNDANT = "REDUNDANT"

// Redundancy
const AddCustomerNetworkPeeringRequest_FABRIC_VC_REDUNDANCY_PRIMARY = "PRIMARY"

func (s *AddCustomerNetworkPeeringRequest) GetType__() vapiBindings_.BindingType {
	return AddCustomerNetworkPeeringRequestBindingType()
}

func (s *AddCustomerNetworkPeeringRequest) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for AddCustomerNetworkPeeringRequest._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type UpdateCustomerNetworkPeeringRequest struct {
	// Metal Connection name
	Name *string
	// Metal Connection description
	Description             *string
	CustomerBgpPeerInfoList []CustomerBgpPeerInfo
}

func (s *UpdateCustomerNetworkPeeringRequest) GetType__() vapiBindings_.BindingType {
	return UpdateCustomerNetworkPeeringRequestBindingType()
}

func (s *UpdateCustomerNetworkPeeringRequest) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for UpdateCustomerNetworkPeeringRequest._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// This is Response Object for VPC Peering Public APIs.
type VpcPeeringInfo struct {
	// Identifier of the vpc peering
	Id *string
	// VPC Id of Deployment
	SddcVpcId *string
	// Cidrs list of Deployment
	SddcCidr     []string
	SddcLocation *Location
	// AWS Account Number of Peering VPC
	CustomerAccountId *string
	// Identifier of the customer vpc
	CustomerVpcId       *string
	CustomerVpcLocation *Location
	// Cidrs list of customer vpc
	CustomerVpcCidrs []string
	// State of Vpc peering
	State *string
	// Error description
	ErrorMsg *string
	// Identifier of cloud provider
	CloudProviderAssignedId *string
}

// State of Vpc peering
const VpcPeeringInfo_STATE_INITIATING = "INITIATING"

// State of Vpc peering
const VpcPeeringInfo_STATE_PENDING_ACCEPTANCE = "PENDING_ACCEPTANCE"

// State of Vpc peering
const VpcPeeringInfo_STATE_PROVISIONING = "PROVISIONING"

// State of Vpc peering
const VpcPeeringInfo_STATE_NACL_ASSOCIATION_FAILED = "NACL_ASSOCIATION_FAILED"

// State of Vpc peering
const VpcPeeringInfo_STATE_ROUTES_UPDATE_FAILED = "ROUTES_UPDATE_FAILED"

// State of Vpc peering
const VpcPeeringInfo_STATE_ROUTE_TABLE_SIZE_FAILED = "ROUTE_TABLE_SIZE_FAILED"

// State of Vpc peering
const VpcPeeringInfo_STATE_ACTIVE = "ACTIVE"

// State of Vpc peering
const VpcPeeringInfo_STATE_FAILED = "FAILED"

// State of Vpc peering
const VpcPeeringInfo_STATE_FAILING = "FAILING"

// State of Vpc peering
const VpcPeeringInfo_STATE_EXPIRED = "EXPIRED"

// State of Vpc peering
const VpcPeeringInfo_STATE_REJECTED = "REJECTED"

// State of Vpc peering
const VpcPeeringInfo_STATE_REJECTING = "REJECTING"

// State of Vpc peering
const VpcPeeringInfo_STATE_DELETED = "DELETED"

// State of Vpc peering
const VpcPeeringInfo_STATE_DELETING = "DELETING"

// State of Vpc peering
const VpcPeeringInfo_STATE_MISSING_FROM_CLOUD_PROVIDER = "MISSING_FROM_CLOUD_PROVIDER"

func (s *VpcPeeringInfo) GetType__() vapiBindings_.BindingType {
	return VpcPeeringInfoBindingType()
}

func (s *VpcPeeringInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for VpcPeeringInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// This DB Entity Object for VPC Peering.
type VpcPeering struct {
	// Identifier of the vpc peering
	Id *string
	// Identifier of the Organization Id.
	OrgId *string
	// Identifier of the Sddc / Deployment Id.
	SddcId *string
	// VPC Id of SDDC
	SddcVpcId *string
	// Region of SDDC
	SddcRegion *string
	// CIDRs list of SDDC / deployment
	SddcCidr []string
	// AWS Account Number of Peering VPC
	CustomerAccountId *string
	// Identifier of the customer vpc
	CustomerVpcId *string
	// Region the customer vpc
	CustomerVpcRegion *string
	// CIDRs list of customer vpc
	CustomerVpcCidrs []string
	// State of Vpc peering
	State *string
	// Error description
	ErrorMsg *string
	// Identifier of cloud provider VPC Peering
	CloudProviderAssignedId *string
	// List of Custom Network ACLs
	NetworkAcls []NetworkAcl
	// List of Reserved Network CIDR used to restrict Overlapping.
	ReservedNetworks []ReservedNetwork
	Creator          *Creator
	Updater          *Updater
}

// State of Vpc peering
const VpcPeering_STATE_INITIATING = "INITIATING"

// State of Vpc peering
const VpcPeering_STATE_PENDING_ACCEPTANCE = "PENDING_ACCEPTANCE"

// State of Vpc peering
const VpcPeering_STATE_PROVISIONING = "PROVISIONING"

// State of Vpc peering
const VpcPeering_STATE_NACL_ASSOCIATION_FAILED = "NACL_ASSOCIATION_FAILED"

// State of Vpc peering
const VpcPeering_STATE_ROUTES_UPDATE_FAILED = "ROUTES_UPDATE_FAILED"

// State of Vpc peering
const VpcPeering_STATE_ROUTE_TABLE_SIZE_FAILED = "ROUTE_TABLE_SIZE_FAILED"

// State of Vpc peering
const VpcPeering_STATE_ACTIVE = "ACTIVE"

// State of Vpc peering
const VpcPeering_STATE_FAILED = "FAILED"

// State of Vpc peering
const VpcPeering_STATE_FAILING = "FAILING"

// State of Vpc peering
const VpcPeering_STATE_EXPIRED = "EXPIRED"

// State of Vpc peering
const VpcPeering_STATE_REJECTED = "REJECTED"

// State of Vpc peering
const VpcPeering_STATE_REJECTING = "REJECTING"

// State of Vpc peering
const VpcPeering_STATE_DELETED = "DELETED"

// State of Vpc peering
const VpcPeering_STATE_DELETING = "DELETING"

// State of Vpc peering
const VpcPeering_STATE_MISSING_FROM_CLOUD_PROVIDER = "MISSING_FROM_CLOUD_PROVIDER"

func (s *VpcPeering) GetType__() vapiBindings_.BindingType {
	return VpcPeeringBindingType()
}

func (s *VpcPeering) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for VpcPeering._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type VpcPeeringRequest struct {
	// AWS Account Number of Peering VPC
	AccountId string
	// SDDC's VPC Id from Shadow Account
	VpcId string
}

func (s *VpcPeeringRequest) GetType__() vapiBindings_.BindingType {
	return VpcPeeringRequestBindingType()
}

func (s *VpcPeeringRequest) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for VpcPeeringRequest._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Activity affected resource
type AffectedResource struct {
	Type_      *string
	ResourceId *string
}

func (s *AffectedResource) GetType__() vapiBindings_.BindingType {
	return AffectedResourceBindingType()
}

func (s *AffectedResource) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for AffectedResource._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type TrafficGroupEniMappingInfos struct {
	TrafficGroupName *string
	Eni              *string
}

func (s *TrafficGroupEniMappingInfos) GetType__() vapiBindings_.BindingType {
	return TrafficGroupEniMappingInfosBindingType()
}

func (s *TrafficGroupEniMappingInfos) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for TrafficGroupEniMappingInfos._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type NsxVifIpAddress struct {
	// IP address text
	IpAddress *string
	// Indicates address is type IPv4 or IPv6
	IpType *string
}

// Indicates address is type IPv4 or IPv6
const NsxVifIpAddress_IP_TYPE_IPV4 = "IPV4"

// Indicates address is type IPv4 or IPv6
const NsxVifIpAddress_IP_TYPE_IPV6 = "IPV6"

func (s *NsxVifIpAddress) GetType__() vapiBindings_.BindingType {
	return NsxVifIpAddressBindingType()
}

func (s *NsxVifIpAddress) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for NsxVifIpAddress._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type CustomerBgpPeerInfo struct {
	// Customer BGP peer IP/ prefix
	CustomerBgpPeerIp *string
	// Metal BGP peer IP/prefix
	MetalBgpPeerIp *string
	// Customer BGP peer ASN
	CustomerBgpPeerAsn *int64
	// Optional BGP MD5 authentication key
	BgpMd5AuthenticationKey *string
	// Provided by VMC App to determine the order of processing virtual circuit
	Index *int64
}

func (s *CustomerBgpPeerInfo) GetType__() vapiBindings_.BindingType {
	return CustomerBgpPeerInfoBindingType()
}

func (s *CustomerBgpPeerInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for CustomerBgpPeerInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type NetworkAcl struct {
	// Identifier of Network ACL
	Id *string
	// Name of Network ACL
	Name *string
}

func (s *NetworkAcl) GetType__() vapiBindings_.BindingType {
	return NetworkAclBindingType()
}

func (s *NetworkAcl) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for NetworkAcl._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type ReservedNetwork struct {
	// Identifier of Reserved Network.
	Id *string
	// CIDR of Reserved Network
	Cidr *string
}

func (s *ReservedNetwork) GetType__() vapiBindings_.BindingType {
	return ReservedNetworkBindingType()
}

func (s *ReservedNetwork) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ReservedNetwork._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

func ActivityBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["org_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fields["state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["state"] = "State"
	fields["activity_type_key"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["activity_type_key"] = "ActivityTypeKey"
	fields["activity_type_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["activity_type_name"] = "ActivityTypeName"
	fields["activity_type_namespace"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["activity_type_namespace"] = "ActivityTypeNamespace"
	fields["progress"] = vapiBindings_.NewOptionalType(vapiBindings_.NewDoubleType())
	fieldNameMap["progress"] = "Progress"
	fields["resources"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(AffectedResourceBindingType), reflect.TypeOf([]AffectedResource{})))
	fieldNameMap["resources"] = "Resources"
	fields["created_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["created_by"] = "CreatedBy"
	fields["created_timestamp"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["created_timestamp"] = "CreatedTimestamp"
	fields["last_updated_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["last_updated_by"] = "LastUpdatedBy"
	fields["last_updated_timestamp"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["last_updated_timestamp"] = "LastUpdatedTimestamp"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.activity", fields, reflect.TypeOf(Activity{}), fieldNameMap, validators)
}

func RouteEntryBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["route_table_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["route_table_id"] = "RouteTableId"
	fields["destination"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["destination"] = "Destination"
	fields["address_family"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["address_family"] = "AddressFamily"
	fields["is_duplicated"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["is_duplicated"] = "IsDuplicated"
	fields["overlap_cidrs"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["overlap_cidrs"] = "OverlapCidrs"
	fields["target"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(RouteEntryTargetBindingType))
	fieldNameMap["target"] = "Target"
	fields["creator"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(CreatorBindingType))
	fieldNameMap["creator"] = "Creator"
	fields["internal_creator"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(CreatorBindingType))
	fieldNameMap["internal_creator"] = "InternalCreator"
	fields["updater"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(UpdaterBindingType))
	fieldNameMap["updater"] = "Updater"
	fields["tags"] = vapiBindings_.NewOptionalType(vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewStringType(), reflect.TypeOf(map[string]string{})))
	fieldNameMap["tags"] = "Tags"
	fields["properties"] = vapiBindings_.NewOptionalType(vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewStringType(), reflect.TypeOf(map[string]string{})))
	fieldNameMap["properties"] = "Properties"
	fields["error_msgs"] = vapiBindings_.NewOptionalType(vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewStringType(), reflect.TypeOf(map[string]string{})))
	fieldNameMap["error_msgs"] = "ErrorMsgs"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.route_entry", fields, reflect.TypeOf(RouteEntry{}), fieldNameMap, validators)
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
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.creator", fields, reflect.TypeOf(Creator{}), fieldNameMap, validators)
}

func NsxAuthenticationEndpointBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["idp"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["idp"] = "Idp"
	fields["preferred_url"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["preferred_url"] = "PreferredUrl"
	fields["other_urls"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["other_urls"] = "OtherUrls"
	fields["desc"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["desc"] = "Desc"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.nsx_authentication_endpoint", fields, reflect.TypeOf(NsxAuthenticationEndpoint{}), fieldNameMap, validators)
}

func PageableBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["page_size"] = "PageSize"
	fields["page_number"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["page_number"] = "PageNumber"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.pageable", fields, reflect.TypeOf(Pageable{}), fieldNameMap, validators)
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
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.updater", fields, reflect.TypeOf(Updater{}), fieldNameMap, validators)
}

func ValidationErrorResponseBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["details"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ValidationErrorDetailBindingType), reflect.TypeOf([]ValidationErrorDetail{})))
	fieldNameMap["details"] = "Details"
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
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.validation_error_response", fields, reflect.TypeOf(ValidationErrorResponse{}), fieldNameMap, validators)
}

func SharedPrefixListResourceShareBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["arn"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["arn"] = "Arn"
	fields["state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["state"] = "State"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.shared_prefix_list_resource_share", fields, reflect.TypeOf(SharedPrefixListResourceShare{}), fieldNameMap, validators)
}

func SharedPrefixListAccountBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["account_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["account_id"] = "AccountId"
	fields["resource_ids"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["resource_ids"] = "ResourceIds"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.shared_prefix_list_account", fields, reflect.TypeOf(SharedPrefixListAccount{}), fieldNameMap, validators)
}

func GroupToGroupConnectivityPeeringBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["region"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["region"] = "Region"
	fields["region_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["region_name"] = "RegionName"
	fields["tgw_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["tgw_id"] = "TgwId"
	fields["peering_region"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["peering_region"] = "PeeringRegion"
	fields["peering_region_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["peering_region_name"] = "PeeringRegionName"
	fields["peering_tgw_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["peering_tgw_id"] = "PeeringTgwId"
	fields["attachment_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["attachment_id"] = "AttachmentId"
	fields["status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["status"] = "Status"
	fields["status_text"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["status_text"] = "StatusText"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.group_to_group_connectivity_peering", fields, reflect.TypeOf(GroupToGroupConnectivityPeering{}), fieldNameMap, validators)
}

func GroupToGroupConnectivityInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["peering_org_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["peering_org_id"] = "PeeringOrgId"
	fields["peering_network_connectivity_config_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["peering_network_connectivity_config_id"] = "PeeringNetworkConnectivityConfigId"
	fields["peering_group_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["peering_group_id"] = "PeeringGroupId"
	fields["peering_group_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["peering_group_name"] = "PeeringGroupName"
	fields["status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["status"] = "Status"
	fields["status_text"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["status_text"] = "StatusText"
	fields["peerings"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(GroupToGroupConnectivityPeeringBindingType), reflect.TypeOf([]GroupToGroupConnectivityPeering{})))
	fieldNameMap["peerings"] = "Peerings"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.group_to_group_connectivity_info", fields, reflect.TypeOf(GroupToGroupConnectivityInfo{}), fieldNameMap, validators)
}

func GroupToGroupConnectivityBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["org_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["description"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["network_connectivity_config_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["network_connectivity_config_id"] = "NetworkConnectivityConfigId"
	fields["group_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["group_id"] = "GroupId"
	fields["group_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["group_name"] = "GroupName"
	fields["last_sync_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["last_sync_time"] = "LastSyncTime"
	fields["version"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["version"] = "Version"
	fields["state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["state"] = "State"
	fields["state_text"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["state_text"] = "StateText"
	fields["creator"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(CreatorBindingType))
	fieldNameMap["creator"] = "Creator"
	fields["updater"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(UpdaterBindingType))
	fieldNameMap["updater"] = "Updater"
	fields["internal_creator"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(CreatorBindingType))
	fieldNameMap["internal_creator"] = "InternalCreator"
	fields["error_msg"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["error_msg"] = "ErrorMsg"
	fields["connectivity_infos"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(GroupToGroupConnectivityInfoBindingType), reflect.TypeOf([]GroupToGroupConnectivityInfo{})))
	fieldNameMap["connectivity_infos"] = "ConnectivityInfos"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.group_to_group_connectivity", fields, reflect.TypeOf(GroupToGroupConnectivity{}), fieldNameMap, validators)
}

func AvailableGroupForConnectivityBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fields["network_connectivity_config_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["network_connectivity_config_id"] = "NetworkConnectivityConfigId"
	fields["group_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["group_id"] = "GroupId"
	fields["group_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["group_name"] = "GroupName"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.available_group_for_connectivity", fields, reflect.TypeOf(AvailableGroupForConnectivity{}), fieldNameMap, validators)
}

func SddcsBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["sddc_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sddc_id"] = "SddcId"
	fields["sddc_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sddc_name"] = "SddcName"
	fields["region"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["region"] = "Region"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.sddcs", fields, reflect.TypeOf(Sddcs{}), fieldNameMap, validators)
}

func DeploymentGroupMemberBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.deployment_group_member", fields, reflect.TypeOf(DeploymentGroupMember{}), fieldNameMap, validators)
}

func RouteTableBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["location"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(LocationBindingType))
	fieldNameMap["location"] = "Location"
	fields["region"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["region"] = "Region"
	fields["description"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["provider_assigned_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["provider_assigned_id"] = "ProviderAssignedId"
	fields["provider_assigned_route_table_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["provider_assigned_route_table_id"] = "ProviderAssignedRouteTableId"
	fields["org_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fields["network_connectivity_config_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["network_connectivity_config_id"] = "NetworkConnectivityConfigId"
	fields["routes_last_sync_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["routes_last_sync_time"] = "RoutesLastSyncTime"
	fields["state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(RouteTableSyncStateBindingType))
	fieldNameMap["state"] = "State"
	fields["creator"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(CreatorBindingType))
	fieldNameMap["creator"] = "Creator"
	fields["internal_creator"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(CreatorBindingType))
	fieldNameMap["internal_creator"] = "InternalCreator"
	fields["updater"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(UpdaterBindingType))
	fieldNameMap["updater"] = "Updater"
	fields["tags"] = vapiBindings_.NewOptionalType(vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewStringType(), reflect.TypeOf(map[string]string{})))
	fieldNameMap["tags"] = "Tags"
	fields["properties"] = vapiBindings_.NewOptionalType(vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewStringType(), reflect.TypeOf(map[string]string{})))
	fieldNameMap["properties"] = "Properties"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.route_table", fields, reflect.TypeOf(RouteTable{}), fieldNameMap, validators)
}

func CreateGroupNetworkConnectivityRequestBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = vapiBindings_.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["description"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["members"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(DeploymentGroupMemberBindingType), reflect.TypeOf([]DeploymentGroupMember{})))
	fieldNameMap["members"] = "Members"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.create_group_network_connectivity_request", fields, reflect.TypeOf(CreateGroupNetworkConnectivityRequest{}), fieldNameMap, validators)
}

func NsxDetailsBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["nsx_private_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["nsx_private_ip"] = "NsxPrivateIp"
	fields["nsx_public_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["nsx_public_ip"] = "NsxPublicIp"
	fields["nsx_default_access"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["nsx_default_access"] = "NsxDefaultAccess"
	fields["nsx_default_auth"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["nsx_default_auth"] = "NsxDefaultAuth"
	fields["nsx_users"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(NsxCredentialInfoBindingType), reflect.TypeOf([]NsxCredentialInfo{})))
	fieldNameMap["nsx_users"] = "NsxUsers"
	fields["nsx_private_fqdn"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["nsx_private_fqdn"] = "NsxPrivateFqdn"
	fields["nsx_public_fqdn"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["nsx_public_fqdn"] = "NsxPublicFqdn"
	fields["login_urls"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(NsxLoginInfoBindingType), reflect.TypeOf([]NsxLoginInfo{})))
	fieldNameMap["login_urls"] = "LoginUrls"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.nsx_details", fields, reflect.TypeOf(NsxDetails{}), fieldNameMap, validators)
}

func StateBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = vapiBindings_.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["error_msg"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["error_msg"] = "ErrorMsg"
	fields["error_code"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["error_code"] = "ErrorCode"
	fields["internal_error_msg"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["internal_error_msg"] = "InternalErrorMsg"
	fields["internal_error_code"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["internal_error_code"] = "InternalErrorCode"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.state", fields, reflect.TypeOf(State{}), fieldNameMap, validators)
}

func NetworkConnectivityConfigBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["group_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["group_id"] = "GroupId"
	fields["name"] = vapiBindings_.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["org_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fields["group_to_group_connectivity_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["group_to_group_connectivity_id"] = "GroupToGroupConnectivityId"
	fields["version"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(NetworkConnectivityConfigVersionBindingType))
	fieldNameMap["version"] = "Version"
	fields["type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(PolicyTypeBindingType))
	fieldNameMap["type"] = "Type_"
	fields["state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(ResourceStateBindingType))
	fieldNameMap["state"] = "State"
	fields["creator"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(CreatorBindingType))
	fieldNameMap["creator"] = "Creator"
	fields["updater"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(UpdaterBindingType))
	fieldNameMap["updater"] = "Updater"
	fields["internal_creator"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(CreatorBindingType))
	fieldNameMap["internal_creator"] = "InternalCreator"
	fields["tags"] = vapiBindings_.NewOptionalType(vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewStringType(), reflect.TypeOf(map[string]string{})))
	fieldNameMap["tags"] = "Tags"
	fields["properties"] = vapiBindings_.NewOptionalType(vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewStringType(), reflect.TypeOf(map[string]string{})))
	fieldNameMap["properties"] = "Properties"
	fields["traits"] = vapiBindings_.NewOptionalType(vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewOpaqueType(), reflect.TypeOf(map[string]vapiData_.DataValue{})))
	fieldNameMap["traits"] = "Traits"
	fields["trait"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(TraitBindingType), reflect.TypeOf([]Trait{})))
	fieldNameMap["trait"] = "Trait"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.network_connectivity_config", fields, reflect.TypeOf(NetworkConnectivityConfig{}), fieldNameMap, validators)
}

func TraitBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["org_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fields["name"] = vapiBindings_.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["type"] = "Type_"
	fields["resource_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["resource_id"] = "ResourceId"
	fields["version"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["version"] = "Version"
	fields["state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(StateBindingType))
	fieldNameMap["state"] = "State"
	fields["creator"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(CreatorBindingType))
	fieldNameMap["creator"] = "Creator"
	fields["updater"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(UpdaterBindingType))
	fieldNameMap["updater"] = "Updater"
	fields["internal_creator"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(CreatorBindingType))
	fieldNameMap["internal_creator"] = "InternalCreator"
	fields["unmapped_properties"] = vapiBindings_.NewOptionalType(vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewDynamicStructType(nil), reflect.TypeOf(map[string]*vapiData_.StructValue{})))
	fieldNameMap["unmapped_properties"] = "UnmappedProperties"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.trait", fields, reflect.TypeOf(Trait{}), fieldNameMap, validators)
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
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.config", fields, reflect.TypeOf(Config{}), fieldNameMap, validators)
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
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.operation", fields, reflect.TypeOf(Operation{}), fieldNameMap, validators)
}

func NetworkConnectivityConfigVersionBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = vapiBindings_.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["code"] = vapiBindings_.NewStringType()
	fieldNameMap["code"] = "Code"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.network_connectivity_config_version", fields, reflect.TypeOf(NetworkConnectivityConfigVersion{}), fieldNameMap, validators)
}

func RouteTableSyncStateBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = vapiBindings_.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["error_msgs"] = vapiBindings_.NewOptionalType(vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewStringType(), reflect.TypeOf(map[string]string{})))
	fieldNameMap["error_msgs"] = "ErrorMsgs"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.route_table_sync_state", fields, reflect.TypeOf(RouteTableSyncState{}), fieldNameMap, validators)
}

func RouteTableResponseBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["content"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(RouteTableBindingType), reflect.TypeOf([]RouteTable{})))
	fieldNameMap["content"] = "Content"
	fields["pageable"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(PageableBindingType))
	fieldNameMap["pageable"] = "Pageable"
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
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.route_table_response", fields, reflect.TypeOf(RouteTableResponse{}), fieldNameMap, validators)
}

func ResourceStateBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = vapiBindings_.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["error_msgs"] = vapiBindings_.NewOptionalType(vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewStringType(), reflect.TypeOf(map[string]string{})))
	fieldNameMap["error_msgs"] = "ErrorMsgs"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.resource_state", fields, reflect.TypeOf(ResourceState{}), fieldNameMap, validators)
}

func NsxCredentialInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["user_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["user_name"] = "UserName"
	fields["password"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["password"] = "Password"
	fields["user_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["user_type"] = "UserType"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.nsx_credential_info", fields, reflect.TypeOf(NsxCredentialInfo{}), fieldNameMap, validators)
}

func PolicyTypeBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = vapiBindings_.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["code"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["code"] = "Code"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.policy_type", fields, reflect.TypeOf(PolicyType{}), fieldNameMap, validators)
}

func RouteEntryTargetBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["type"] = "Type_"
	fields["location"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(LocationBindingType))
	fieldNameMap["location"] = "Location"
	fields["region"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["region"] = "Region"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.route_entry_target", fields, reflect.TypeOf(RouteEntryTarget{}), fieldNameMap, validators)
}

func CreateGroupNetworkConnectivityResponseBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["group_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["group_id"] = "GroupId"
	fields["config_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["config_id"] = "ConfigId"
	fields["operation_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["operation_id"] = "OperationId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.create_group_network_connectivity_response", fields, reflect.TypeOf(CreateGroupNetworkConnectivityResponse{}), fieldNameMap, validators)
}

func ValidationPayloadBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["members"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(DeploymentGroupMemberBindingType), reflect.TypeOf([]DeploymentGroupMember{})))
	fieldNameMap["members"] = "Members"
	fields["deployment_group_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["deployment_group_id"] = "DeploymentGroupId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.validation_payload", fields, reflect.TypeOf(ValidationPayload{}), fieldNameMap, validators)
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
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.error_response", fields, reflect.TypeOf(ErrorResponse{}), fieldNameMap, validators)
}

func RouteEntryResponseBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["content"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(RouteEntryBindingType), reflect.TypeOf([]RouteEntry{})))
	fieldNameMap["content"] = "Content"
	fields["pageable"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(PageableBindingType))
	fieldNameMap["pageable"] = "Pageable"
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
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.route_entry_response", fields, reflect.TypeOf(RouteEntryResponse{}), fieldNameMap, validators)
}

func NsxLoginInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["access_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["access_type"] = "AccessType"
	fields["auth_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["auth_type"] = "AuthType"
	fields["preferred_url"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["preferred_url"] = "PreferredUrl"
	fields["other_urls"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["other_urls"] = "OtherUrls"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.nsx_login_info", fields, reflect.TypeOf(NsxLoginInfo{}), fieldNameMap, validators)
}

func ValidationErrorDetailBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["validation_error_message"] = vapiBindings_.NewStringType()
	fieldNameMap["validation_error_message"] = "ValidationErrorMessage"
	fields["localized_validation_error_message"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["localized_validation_error_message"] = "LocalizedValidationErrorMessage"
	fields["members"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["members"] = "Members"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.validation_error_detail", fields, reflect.TypeOf(ValidationErrorDetail{}), fieldNameMap, validators)
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
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.location", fields, reflect.TypeOf(Location{}), fieldNameMap, validators)
}

func AwsSharedPrefixListBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["org_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fields["network_connectivity_config_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["network_connectivity_config_id"] = "NetworkConnectivityConfigId"
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["address_family"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["address_family"] = "AddressFamily"
	fields["cloud_provider_assigned_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cloud_provider_assigned_id"] = "CloudProviderAssignedId"
	fields["cloud_provider_assigned_version"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cloud_provider_assigned_version"] = "CloudProviderAssignedVersion"
	fields["location"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(LocationBindingType))
	fieldNameMap["location"] = "Location"
	fields["state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["state"] = "State"
	fields["creator"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(CreatorBindingType))
	fieldNameMap["creator"] = "Creator"
	fields["updater"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(UpdaterBindingType))
	fieldNameMap["updater"] = "Updater"
	fields["internal_creator"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(CreatorBindingType))
	fieldNameMap["internal_creator"] = "InternalCreator"
	fields["size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["size"] = "Size"
	fields["total_entries"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["total_entries"] = "TotalEntries"
	fields["membership_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["membership_type"] = "MembershipType"
	fields["tags"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewStringType(), reflect.TypeOf(map[string]string{})), reflect.TypeOf([]map[string]string{})))
	fieldNameMap["tags"] = "Tags"
	fields["sddc_location"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(LocationBindingType))
	fieldNameMap["sddc_location"] = "SddcLocation"
	fields["sddcs"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(SddcsBindingType), reflect.TypeOf([]Sddcs{})))
	fieldNameMap["sddcs"] = "Sddcs"
	fields["accounts"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(SharedPrefixListAccountBindingType), reflect.TypeOf([]SharedPrefixListAccount{})))
	fieldNameMap["accounts"] = "Accounts"
	fields["resource_share"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SharedPrefixListResourceShareBindingType))
	fieldNameMap["resource_share"] = "ResourceShare"
	fields["version"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["version"] = "Version"
	fields["error_msgs"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["error_msgs"] = "ErrorMsgs"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.aws_shared_prefix_list", fields, reflect.TypeOf(AwsSharedPrefixList{}), fieldNameMap, validators)
}

func SharedPrefixListEntryResponseBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["content"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(SharedPrefixListEntryBindingType), reflect.TypeOf([]SharedPrefixListEntry{})))
	fieldNameMap["content"] = "Content"
	fields["pageable"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(PageableBindingType))
	fieldNameMap["pageable"] = "Pageable"
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
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.shared_prefix_list_entry_response", fields, reflect.TypeOf(SharedPrefixListEntryResponse{}), fieldNameMap, validators)
}

func SharedPrefixListEntryBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["org_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fields["shared_pl_id"] = vapiBindings_.NewStringType()
	fieldNameMap["shared_pl_id"] = "SharedPlId"
	fields["cidr"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cidr"] = "Cidr"
	fields["is_duplicated"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["is_duplicated"] = "IsDuplicated"
	fields["target"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(RouteEntryTargetBindingType))
	fieldNameMap["target"] = "Target"
	fields["creator"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(CreatorBindingType))
	fieldNameMap["creator"] = "Creator"
	fields["error_msg"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["error_msg"] = "ErrorMsg"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.shared_prefix_list_entry", fields, reflect.TypeOf(SharedPrefixListEntry{}), fieldNameMap, validators)
}

func CustomerNetworkPeeringInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["customer_network_peerings"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(MetalConnectionBindingType), reflect.TypeOf([]MetalConnection{})))
	fieldNameMap["customer_network_peerings"] = "CustomerNetworkPeerings"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.customer_network_peering_info", fields, reflect.TypeOf(CustomerNetworkPeeringInfo{}), fieldNameMap, validators)
}

func NetworkActivityBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["org_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fields["deployment_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["deployment_id"] = "DeploymentId"
	fields["resource_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["resource_id"] = "ResourceId"
	fields["resource_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["state"] = "State"
	fields["error"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(ExecutionErrorBindingType))
	fieldNameMap["error"] = "Error_"
	fields["activity_type_key"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["activity_type_key"] = "ActivityTypeKey"
	fields["created_timestamp"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["created_timestamp"] = "CreatedTimestamp"
	fields["last_updated_timestamp"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["last_updated_timestamp"] = "LastUpdatedTimestamp"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.network_activity", fields, reflect.TypeOf(NetworkActivity{}), fieldNameMap, validators)
}

func ExecutionErrorBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["code"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["code"] = "Code"
	fields["status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["status"] = "Status"
	fields["message"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["message"] = "Message"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.execution_error", fields, reflect.TypeOf(ExecutionError{}), fieldNameMap, validators)
}

func NsxSettingsBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["nsx_manager_access_mode"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["nsx_manager_access_mode"] = "NsxManagerAccessMode"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.nsx_settings", fields, reflect.TypeOf(NsxSettings{}), fieldNameMap, validators)
}

func NsxAddonResponseBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["action"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["action"] = "Action"
	fields["error_msg"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["error_msg"] = "ErrorMsg"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.nsx_addon_response", fields, reflect.TypeOf(NsxAddonResponse{}), fieldNameMap, validators)
}

func VpnBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["connectivity_down"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["connectivity_down"] = "ConnectivityDown"
	fields["connectivity_up"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["connectivity_up"] = "ConnectivityUp"
	fields["connectivity_deactivated"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["connectivity_deactivated"] = "ConnectivityDeactivated"
	fields["total_connections"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["total_connections"] = "TotalConnections"
	fields["public_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["public_ip"] = "PublicIp"
	fields["vpn_down_connections"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["vpn_down_connections"] = "VpnDownConnections"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.vpn", fields, reflect.TypeOf(Vpn{}), fieldNameMap, validators)
}

func DirectConnectBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["total_attached_vifs_count"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["total_attached_vifs_count"] = "TotalAttachedVifsCount"
	fields["connectivity_down"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["connectivity_down"] = "ConnectivityDown"
	fields["connectivity_up"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["connectivity_up"] = "ConnectivityUp"
	fields["learned_routes"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["learned_routes"] = "LearnedRoutes"
	fields["advertised_routes"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["advertised_routes"] = "AdvertisedRoutes"
	fields["down_vif_names"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["down_vif_names"] = "DownVifNames"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.direct_connect", fields, reflect.TypeOf(DirectConnect{}), fieldNameMap, validators)
}

func TransitConnectBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["connected_sddc_group_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["connected_sddc_group_id"] = "ConnectedSddcGroupId"
	fields["connected_sddc_group_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["connected_sddc_group_name"] = "ConnectedSddcGroupName"
	fields["learned_routes"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["learned_routes"] = "LearnedRoutes"
	fields["advertised_routes"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["advertised_routes"] = "AdvertisedRoutes"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.transit_connect", fields, reflect.TypeOf(TransitConnect{}), fieldNameMap, validators)
}

func ManagementGatewayBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["total_firewall_rules"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["total_firewall_rules"] = "TotalFirewallRules"
	fields["mgmt_subnets"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["mgmt_subnets"] = "MgmtSubnets"
	fields["appliance_subnet"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["appliance_subnet"] = "ApplianceSubnet"
	fields["infra_subnets"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["infra_subnets"] = "InfraSubnets"
	fields["dns"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["dns"] = "Dns"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.management_gateway", fields, reflect.TypeOf(ManagementGateway{}), fieldNameMap, validators)
}

func ComputeGatewayBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["total_firewall_rules"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["total_firewall_rules"] = "TotalFirewallRules"
	fields["total_segments"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["total_segments"] = "TotalSegments"
	fields["dns"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["dns"] = "Dns"
	fields["total_tier1_gateways"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["total_tier1_gateways"] = "TotalTier1Gateways"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.compute_gateway", fields, reflect.TypeOf(ComputeGateway{}), fieldNameMap, validators)
}

func VpcBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["availability_zone"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["availability_zone"] = "AvailabilityZone"
	fields["availability_zone_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["availability_zone_id"] = "AvailabilityZoneId"
	fields["cidr"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cidr"] = "Cidr"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.vpc", fields, reflect.TypeOf(Vpc{}), fieldNameMap, validators)
}

func CloudProviderBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["account_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["account_id"] = "AccountId"
	fields["connectivity_status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["connectivity_status"] = "ConnectivityStatus"
	fields["traffic_group_eni_mapping_infos"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(TrafficGroupEniMappingInfosBindingType), reflect.TypeOf([]TrafficGroupEniMappingInfos{})))
	fieldNameMap["traffic_group_eni_mapping_infos"] = "TrafficGroupEniMappingInfos"
	fields["vpc"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(VpcBindingType), reflect.TypeOf([]Vpc{})))
	fieldNameMap["vpc"] = "Vpc"
	fields["cloud_formation_stack_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cloud_formation_stack_name"] = "CloudFormationStackName"
	fields["s3_enabled"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["s3_enabled"] = "S3Enabled"
	fields["ec2_enabled"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["ec2_enabled"] = "Ec2Enabled"
	fields["iam_role_names"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["iam_role_names"] = "IamRoleNames"
	fields["vpc_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vpc_id"] = "VpcId"
	fields["vpc_cidr"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["vpc_cidr"] = "VpcCidr"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.cloud_provider", fields, reflect.TypeOf(CloudProvider{}), fieldNameMap, validators)
}

func DeploymentNetworkingStateInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["location"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(LocationBindingType))
	fieldNameMap["location"] = "Location"
	fields["sddc_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sddc_id"] = "SddcId"
	fields["org_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fields["provider_code"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["provider_code"] = "ProviderCode"
	fields["provider_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["provider_type"] = "ProviderType"
	fields["last_sync_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["last_sync_time"] = "LastSyncTime"
	fields["vpn"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(VpnBindingType), reflect.TypeOf([]Vpn{})))
	fieldNameMap["vpn"] = "Vpn"
	fields["direct_connect"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(DirectConnectBindingType), reflect.TypeOf([]DirectConnect{})))
	fieldNameMap["direct_connect"] = "DirectConnect"
	fields["outposts_connect"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(OutpostConnectBindingType), reflect.TypeOf([]OutpostConnect{})))
	fieldNameMap["outposts_connect"] = "OutpostsConnect"
	fields["transit_connect"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(TransitConnectBindingType), reflect.TypeOf([]TransitConnect{})))
	fieldNameMap["transit_connect"] = "TransitConnect"
	fields["management_gateway"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ManagementGatewayBindingType), reflect.TypeOf([]ManagementGateway{})))
	fieldNameMap["management_gateway"] = "ManagementGateway"
	fields["compute_gateway"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ComputeGatewayBindingType), reflect.TypeOf([]ComputeGateway{})))
	fieldNameMap["compute_gateway"] = "ComputeGateway"
	fields["cloud_provider"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(CloudProviderBindingType), reflect.TypeOf([]CloudProvider{})))
	fieldNameMap["cloud_provider"] = "CloudProvider"
	fields["error"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(ErrorMessageBindingType))
	fieldNameMap["error"] = "Error_"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.deployment_networking_state_info", fields, reflect.TypeOf(DeploymentNetworkingStateInfo{}), fieldNameMap, validators)
}

func ErrorMessageBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["code"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["code"] = "Code"
	fields["retry_count"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["retry_count"] = "RetryCount"
	fields["message"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["message"] = "Message"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.error_message", fields, reflect.TypeOf(ErrorMessage{}), fieldNameMap, validators)
}

func OutpostConnectBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["learned_routes"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["learned_routes"] = "LearnedRoutes"
	fields["advertised_routes"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["advertised_routes"] = "AdvertisedRoutes"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.outpost_connect", fields, reflect.TypeOf(OutpostConnect{}), fieldNameMap, validators)
}

func MetalServiceTokenBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["service_token_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["service_token_type"] = "ServiceTokenType"
	fields["max_allowed_speed"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["max_allowed_speed"] = "MaxAllowedSpeed"
	fields["expires_at"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["expires_at"] = "ExpiresAt"
	fields["role"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["role"] = "Role"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.metal_service_token", fields, reflect.TypeOf(MetalServiceToken{}), fieldNameMap, validators)
}

func MetalVirtualCircuitBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["vrf"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vrf"] = "Vrf"
	fields["status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["status"] = "Status"
	fields["connection_auth_key"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["connection_auth_key"] = "ConnectionAuthKey"
	fields["bgp_peer_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["bgp_peer_ip"] = "BgpPeerIp"
	fields["bgp_peer_net_mask"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["bgp_peer_net_mask"] = "BgpPeerNetMask"
	fields["metal_bgp_peer_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["metal_bgp_peer_ip"] = "MetalBgpPeerIp"
	fields["bgp_peer_asn"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["bgp_peer_asn"] = "BgpPeerAsn"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.metal_virtual_circuit", fields, reflect.TypeOf(MetalVirtualCircuit{}), fieldNameMap, validators)
}

func MetalPortBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["role"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["role"] = "Role"
	fields["speed"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["speed"] = "Speed"
	fields["virtual_circuit"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(MetalVirtualCircuitBindingType))
	fieldNameMap["virtual_circuit"] = "VirtualCircuit"
	fields["service_token"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(MetalServiceTokenBindingType))
	fieldNameMap["service_token"] = "ServiceToken"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.metal_port", fields, reflect.TypeOf(MetalPort{}), fieldNameMap, validators)
}

func MetalConnectionBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["connection_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["connection_id"] = "ConnectionId"
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["description"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["vmc_connection_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vmc_connection_type"] = "VmcConnectionType"
	fields["redundancy"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["redundancy"] = "Redundancy"
	fields["type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["type"] = "Type_"
	fields["metro"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["metro"] = "Metro"
	fields["ports"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(MetalPortBindingType), reflect.TypeOf([]MetalPort{})))
	fieldNameMap["ports"] = "Ports"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.metal_connection", fields, reflect.TypeOf(MetalConnection{}), fieldNameMap, validators)
}

func PublicIpBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["reservation_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["reservation_id"] = "ReservationId"
	fields["attachment_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["attachment_id"] = "AttachmentId"
	fields["state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["state"] = "State"
	fields["cidr"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["cidr"] = "Cidr"
	fields["address"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["address"] = "Address"
	fields["nextHop"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["nextHop"] = "NextHop"
	fields["gateway"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["gateway"] = "Gateway"
	fields["network"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["network"] = "Network"
	fields["description"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["description"] = "Description"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.public_ip", fields, reflect.TypeOf(PublicIp{}), fieldNameMap, validators)
}

func AssignElasticIpRequestBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["description"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["description"] = "Description"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.assign_elastic_ip_request", fields, reflect.TypeOf(AssignElasticIpRequest{}), fieldNameMap, validators)
}

func PrivateIpBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["state"] = "State"
	fields["address"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["address"] = "Address"
	fields["subnet_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["subnet_name"] = "SubnetName"
	fields["address_family"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["address_family"] = "AddressFamily"
	fields["request_source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["request_source"] = "RequestSource"
	fields["description"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["description"] = "Description"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.private_ip", fields, reflect.TypeOf(PrivateIp{}), fieldNameMap, validators)
}

func AssignPrivateIpRequestBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["subnet_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["subnet_name"] = "SubnetName"
	fields["request_source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["request_source"] = "RequestSource"
	fields["description"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["description"] = "Description"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.assign_private_ip_request", fields, reflect.TypeOf(AssignPrivateIpRequest{}), fieldNameMap, validators)
}

func NetworkPropertiesBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["management_subnet"] = vapiBindings_.NewStringType()
	fieldNameMap["management_subnet"] = "ManagementSubnet"
	fields["vxlan_subnet"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vxlan_subnet"] = "VxlanSubnet"
	fields["equinix_metal_router_asn"] = vapiBindings_.NewStringType()
	fieldNameMap["equinix_metal_router_asn"] = "EquinixMetalRouterAsn"
	fields["nsx_edge_uplink_router_asn"] = vapiBindings_.NewStringType()
	fieldNameMap["nsx_edge_uplink_router_asn"] = "NsxEdgeUplinkRouterAsn"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.network_properties", fields, reflect.TypeOf(NetworkProperties{}), fieldNameMap, validators)
}

func CreateNsxTraceflowRequestBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["network_connectivity_config_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["network_connectivity_config_id"] = "NetworkConnectivityConfigId"
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(NsxTraceflowEndpointBindingType))
	fieldNameMap["source"] = "Source"
	fields["destination"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(NsxTraceflowEndpointBindingType))
	fieldNameMap["destination"] = "Destination"
	fields["packet"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(PacketDataBindingType))
	fieldNameMap["packet"] = "Packet"
	fields["source_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(NsxTraceflowSourceIdBindingType))
	fieldNameMap["source_id"] = "SourceId"
	fields["timeout"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["timeout"] = "Timeout"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.create_nsx_traceflow_request", fields, reflect.TypeOf(CreateNsxTraceflowRequest{}), fieldNameMap, validators)
}

func NsxTraceflowEndpointBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["traceflow_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["traceflow_type"] = "TraceflowType"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.nsx_traceflow_endpoint", fields, reflect.TypeOf(NsxTraceflowEndpoint{}), fieldNameMap, validators)
}

func NsxTraceflowSourceIdBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["source_id_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["source_id_type"] = "SourceIdType"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.nsx_traceflow_source_id", fields, reflect.TypeOf(NsxTraceflowSourceId{}), fieldNameMap, validators)
}

func SddcTraceflowBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["org_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fields["network_connectivity_config_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["network_connectivity_config_id"] = "NetworkConnectivityConfigId"
	fields["source_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(NsxTraceflowSourceIdBindingType))
	fieldNameMap["source_id"] = "SourceId"
	fields["state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["state"] = "State"
	fields["packet"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(PacketDataBindingType))
	fieldNameMap["packet"] = "Packet"
	fields["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(NsxTraceflowEndpointBindingType))
	fieldNameMap["source"] = "Source"
	fields["destination"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(NsxTraceflowEndpointBindingType))
	fieldNameMap["destination"] = "Destination"
	fields["observations"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(TraceflowObservationListResultBindingType))
	fieldNameMap["observations"] = "Observations"
	fields["timeout"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["timeout"] = "Timeout"
	fields["creator"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(CreatorBindingType))
	fieldNameMap["creator"] = "Creator"
	fields["updater"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(UpdaterBindingType))
	fieldNameMap["updater"] = "Updater"
	fields["error_msgs"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["error_msgs"] = "ErrorMsgs"
	fields["realization_error_msg"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["realization_error_msg"] = "RealizationErrorMsg"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.sddc_traceflow", fields, reflect.TypeOf(SddcTraceflow{}), fieldNameMap, validators)
}

func PacketDataBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["resource_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["arp_header"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(ArpHeaderBindingType))
	fieldNameMap["arp_header"] = "ArpHeader"
	fields["eth_header"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(EthernetHeaderBindingType))
	fieldNameMap["eth_header"] = "EthHeader"
	fields["ip_header"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(Ipv4HeaderBindingType))
	fieldNameMap["ip_header"] = "IpHeader"
	fields["ipv6_header"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(Ipv6HeaderBindingType))
	fieldNameMap["ipv6_header"] = "Ipv6Header"
	fields["payload"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["payload"] = "Payload"
	fields["transport_header"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(TransportProtocolHeaderBindingType))
	fieldNameMap["transport_header"] = "TransportHeader"
	fields["frame_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["frame_size"] = "FrameSize"
	fields["routed"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["routed"] = "Routed"
	fields["transport_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["transport_type"] = "TransportType"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.packet_data", fields, reflect.TypeOf(PacketData{}), fieldNameMap, validators)
}

func Ipv4HeaderBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["src_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["src_ip"] = "SrcIp"
	fields["dst_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["dst_ip"] = "DstIp"
	fields["src_subnet_prefix_len"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["src_subnet_prefix_len"] = "SrcSubnetPrefixLen"
	fields["protocol"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["protocol"] = "Protocol"
	fields["ttl"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["ttl"] = "Ttl"
	fields["flags"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["flags"] = "Flags"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.ipv4_header", fields, reflect.TypeOf(Ipv4Header{}), fieldNameMap, validators)
}

func Ipv6HeaderBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["src_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["src_ip"] = "SrcIp"
	fields["dst_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["dst_ip"] = "DstIp"
	fields["hop_limit"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["hop_limit"] = "HopLimit"
	fields["next_header"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["next_header"] = "NextHeader"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.ipv6_header", fields, reflect.TypeOf(Ipv6Header{}), fieldNameMap, validators)
}

func ArpHeaderBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["dst_ip"] = vapiBindings_.NewStringType()
	fieldNameMap["dst_ip"] = "DstIp"
	fields["op_code"] = vapiBindings_.NewStringType()
	fieldNameMap["op_code"] = "OpCode"
	fields["src_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["src_ip"] = "SrcIp"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.arp_header", fields, reflect.TypeOf(ArpHeader{}), fieldNameMap, validators)
}

func EthernetHeaderBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["src_mac"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["src_mac"] = "SrcMac"
	fields["dst_mac"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["dst_mac"] = "DstMac"
	fields["eth_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["eth_type"] = "EthType"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.ethernet_header", fields, reflect.TypeOf(EthernetHeader{}), fieldNameMap, validators)
}

func TransportProtocolHeaderBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["icmp_echo_request_header"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(IcmpEchoRequestHeaderBindingType))
	fieldNameMap["icmp_echo_request_header"] = "IcmpEchoRequestHeader"
	fields["udp_header"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(UdpHeaderBindingType))
	fieldNameMap["udp_header"] = "UdpHeader"
	fields["tcp_header"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(TcpHeaderBindingType))
	fieldNameMap["tcp_header"] = "TcpHeader"
	fields["dhcp_header"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(DhcpHeaderBindingType))
	fieldNameMap["dhcp_header"] = "DhcpHeader"
	fields["dhcpv6_header"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(Dhcpv6HeaderBindingType))
	fieldNameMap["dhcpv6_header"] = "Dhcpv6Header"
	fields["dns_header"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(DnsHeaderBindingType))
	fieldNameMap["dns_header"] = "DnsHeader"
	fields["ndp_header"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(NdpHeaderBindingType))
	fieldNameMap["ndp_header"] = "NdpHeader"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.transport_protocol_header", fields, reflect.TypeOf(TransportProtocolHeader{}), fieldNameMap, validators)
}

func IcmpEchoRequestHeaderBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["id"] = "Id"
	fields["sequence"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["sequence"] = "Sequence"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.icmp_echo_request_header", fields, reflect.TypeOf(IcmpEchoRequestHeader{}), fieldNameMap, validators)
}

func UdpHeaderBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["src_port"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["src_port"] = "SrcPort"
	fields["dst_port"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["dst_port"] = "DstPort"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.udp_header", fields, reflect.TypeOf(UdpHeader{}), fieldNameMap, validators)
}

func TcpHeaderBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["src_port"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["src_port"] = "SrcPort"
	fields["dst_port"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["dst_port"] = "DstPort"
	fields["tcp_flags"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["tcp_flags"] = "TcpFlags"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.tcp_header", fields, reflect.TypeOf(TcpHeader{}), fieldNameMap, validators)
}

func DhcpHeaderBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["op_code"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["op_code"] = "OpCode"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.dhcp_header", fields, reflect.TypeOf(DhcpHeader{}), fieldNameMap, validators)
}

func Dhcpv6HeaderBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["msg_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["msg_type"] = "MsgType"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.dhcpv6_header", fields, reflect.TypeOf(Dhcpv6Header{}), fieldNameMap, validators)
}

func DnsHeaderBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["address"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["address"] = "Address"
	fields["address_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["address_type"] = "AddressType"
	fields["message_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["message_type"] = "MessageType"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.dns_header", fields, reflect.TypeOf(DnsHeader{}), fieldNameMap, validators)
}

func NdpHeaderBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["dst_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["dst_ip"] = "DstIp"
	fields["msg_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["msg_type"] = "MsgType"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.ndp_header", fields, reflect.TypeOf(NdpHeader{}), fieldNameMap, validators)
}

func NsxVmListBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm_list"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(NsxVmBindingType), reflect.TypeOf([]NsxVm{})))
	fieldNameMap["vm_list"] = "VmList"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.nsx_vm_list", fields, reflect.TypeOf(NsxVmList{}), fieldNameMap, validators)
}

func NsxVmBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vm_id"] = "VmId"
	fields["display_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["host_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["host_id"] = "HostId"
	fields["virtual_interfaces"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(NsxVirtualInterfaceBindingType), reflect.TypeOf([]NsxVirtualInterface{})))
	fieldNameMap["virtual_interfaces"] = "VirtualInterfaces"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.nsx_vm", fields, reflect.TypeOf(NsxVm{}), fieldNameMap, validators)
}

func NsxVirtualInterfaceBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vif_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vif_id"] = "VifId"
	fields["display_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["ip_addresses"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(NsxVifIpAddressBindingType), reflect.TypeOf([]NsxVifIpAddress{})))
	fieldNameMap["ip_addresses"] = "IpAddresses"
	fields["mac_address"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["mac_address"] = "MacAddress"
	fields["attachment_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["attachment_id"] = "AttachmentId"
	fields["segment_port_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["segment_port_id"] = "SegmentPortId"
	fields["segment_port_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["segment_port_path"] = "SegmentPortPath"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.nsx_virtual_interface", fields, reflect.TypeOf(NsxVirtualInterface{}), fieldNameMap, validators)
}

func TraceflowObservationListResultBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["result_count"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["result_count"] = "ResultCount"
	fields["results"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(TraceflowObservationBindingType), reflect.TypeOf([]TraceflowObservation{})))
	fieldNameMap["results"] = "Results"
	fields["realization_error_msg"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["realization_error_msg"] = "RealizationErrorMsg"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.traceflow_observation_list_result", fields, reflect.TypeOf(TraceflowObservationListResult{}), fieldNameMap, validators)
}

func TraceflowObservationBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["component_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["component_name"] = "ComponentName"
	fields["component_sub_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["component_sub_type"] = "ComponentSubType"
	fields["component_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["component_type"] = "ComponentType"
	fields["component_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["component_path"] = "ComponentPath"
	fields["interface_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["interface_path"] = "InterfacePath"
	fields["resource_type"] = vapiBindings_.NewStringType()
	fieldNameMap["resource_type"] = "ResourceType"
	fields["sequence_no"] = vapiBindings_.NewIntegerType()
	fieldNameMap["sequence_no"] = "SequenceNo"
	fields["site_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["site_path"] = "SitePath"
	fields["timestamp"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["timestamp"] = "Timestamp"
	fields["timestamp_micro"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["timestamp_micro"] = "TimestampMicro"
	fields["transport_node_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	fields["transport_node_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["transport_node_name"] = "TransportNodeName"
	fields["transport_node_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["transport_node_type"] = "TransportNodeType"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.traceflow_observation", fields, reflect.TypeOf(TraceflowObservation{}), fieldNameMap, validators)
}

func AddCustomerNetworkPeeringRequestBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["description"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["fabric_vc_redundancy"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["fabric_vc_redundancy"] = "FabricVcRedundancy"
	fields["customer_bgp_peer_info_list"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(CustomerBgpPeerInfoBindingType), reflect.TypeOf([]CustomerBgpPeerInfo{})))
	fieldNameMap["customer_bgp_peer_info_list"] = "CustomerBgpPeerInfoList"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.add_customer_network_peering_request", fields, reflect.TypeOf(AddCustomerNetworkPeeringRequest{}), fieldNameMap, validators)
}

func UpdateCustomerNetworkPeeringRequestBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["description"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["customer_bgp_peer_info_list"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(CustomerBgpPeerInfoBindingType), reflect.TypeOf([]CustomerBgpPeerInfo{})))
	fieldNameMap["customer_bgp_peer_info_list"] = "CustomerBgpPeerInfoList"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.update_customer_network_peering_request", fields, reflect.TypeOf(UpdateCustomerNetworkPeeringRequest{}), fieldNameMap, validators)
}

func VpcPeeringInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["sddc_vpc_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sddc_vpc_id"] = "SddcVpcId"
	fields["sddc_cidr"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["sddc_cidr"] = "SddcCidr"
	fields["sddc_location"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(LocationBindingType))
	fieldNameMap["sddc_location"] = "SddcLocation"
	fields["customer_account_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["customer_account_id"] = "CustomerAccountId"
	fields["customer_vpc_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["customer_vpc_id"] = "CustomerVpcId"
	fields["customer_vpc_location"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(LocationBindingType))
	fieldNameMap["customer_vpc_location"] = "CustomerVpcLocation"
	fields["customer_vpc_cidrs"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["customer_vpc_cidrs"] = "CustomerVpcCidrs"
	fields["state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["state"] = "State"
	fields["error_msg"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["error_msg"] = "ErrorMsg"
	fields["cloud_provider_assigned_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cloud_provider_assigned_id"] = "CloudProviderAssignedId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.vpc_peering_info", fields, reflect.TypeOf(VpcPeeringInfo{}), fieldNameMap, validators)
}

func VpcPeeringBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["org_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fields["sddc_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sddc_id"] = "SddcId"
	fields["sddc_vpc_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sddc_vpc_id"] = "SddcVpcId"
	fields["sddc_region"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sddc_region"] = "SddcRegion"
	fields["sddc_cidr"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["sddc_cidr"] = "SddcCidr"
	fields["customer_account_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["customer_account_id"] = "CustomerAccountId"
	fields["customer_vpc_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["customer_vpc_id"] = "CustomerVpcId"
	fields["customer_vpc_region"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["customer_vpc_region"] = "CustomerVpcRegion"
	fields["customer_vpc_cidrs"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["customer_vpc_cidrs"] = "CustomerVpcCidrs"
	fields["state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["state"] = "State"
	fields["error_msg"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["error_msg"] = "ErrorMsg"
	fields["cloud_provider_assigned_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cloud_provider_assigned_id"] = "CloudProviderAssignedId"
	fields["network_acls"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(NetworkAclBindingType), reflect.TypeOf([]NetworkAcl{})))
	fieldNameMap["network_acls"] = "NetworkAcls"
	fields["reserved_networks"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ReservedNetworkBindingType), reflect.TypeOf([]ReservedNetwork{})))
	fieldNameMap["reserved_networks"] = "ReservedNetworks"
	fields["creator"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(CreatorBindingType))
	fieldNameMap["creator"] = "Creator"
	fields["updater"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(UpdaterBindingType))
	fieldNameMap["updater"] = "Updater"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.vpc_peering", fields, reflect.TypeOf(VpcPeering{}), fieldNameMap, validators)
}

func VpcPeeringRequestBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["account_id"] = vapiBindings_.NewStringType()
	fieldNameMap["account_id"] = "AccountId"
	fields["vpc_id"] = vapiBindings_.NewStringType()
	fieldNameMap["vpc_id"] = "VpcId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.vpc_peering_request", fields, reflect.TypeOf(VpcPeeringRequest{}), fieldNameMap, validators)
}

func AffectedResourceBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["type"] = "Type_"
	fields["resource_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["resource_id"] = "ResourceId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.affected_resource", fields, reflect.TypeOf(AffectedResource{}), fieldNameMap, validators)
}

func TrafficGroupEniMappingInfosBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["traffic_group_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["traffic_group_name"] = "TrafficGroupName"
	fields["eni"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["eni"] = "Eni"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.traffic_group_eni_mapping_infos", fields, reflect.TypeOf(TrafficGroupEniMappingInfos{}), fieldNameMap, validators)
}

func NsxVifIpAddressBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ip_address"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["ip_address"] = "IpAddress"
	fields["ip_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["ip_type"] = "IpType"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.nsx_vif_ip_address", fields, reflect.TypeOf(NsxVifIpAddress{}), fieldNameMap, validators)
}

func CustomerBgpPeerInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["customer_bgp_peer_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["customer_bgp_peer_ip"] = "CustomerBgpPeerIp"
	fields["metal_bgp_peer_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["metal_bgp_peer_ip"] = "MetalBgpPeerIp"
	fields["customer_bgp_peer_asn"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["customer_bgp_peer_asn"] = "CustomerBgpPeerAsn"
	fields["bgp_md5_authentication_key"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["bgp_md5_authentication_key"] = "BgpMd5AuthenticationKey"
	fields["index"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["index"] = "Index"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.customer_bgp_peer_info", fields, reflect.TypeOf(CustomerBgpPeerInfo{}), fieldNameMap, validators)
}

func NetworkAclBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.network_acl", fields, reflect.TypeOf(NetworkAcl{}), fieldNameMap, validators)
}

func ReservedNetworkBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["cidr"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cidr"] = "Cidr"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_network.model.reserved_network", fields, reflect.TypeOf(ReservedNetwork{}), fieldNameMap, validators)
}
