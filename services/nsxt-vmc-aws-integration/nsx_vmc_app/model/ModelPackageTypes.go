// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for package: com.vmware.nsx_vmc_app.model.
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

type ActionMessage struct {
	// Optional message for action
	Message *string
}

func (s *ActionMessage) GetType__() vapiBindings_.BindingType {
	return ActionMessageBindingType()
}

func (s *ActionMessage) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ActionMessage._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Advertised BGP route
type AdvertisedRoute struct {
	// Possible values are:
	//
	// * AdvertisedRoute#AdvertisedRoute_ADDRESS_FAMILY_IPV4
	// * AdvertisedRoute#AdvertisedRoute_ADDRESS_FAMILY_IPV6
	//
	//  IP address type
	AddressFamily *string
	// Possible values are:
	//
	// * AdvertisedRoute#AdvertisedRoute_ADVERTISEMENT_STATE_SUCCESS
	// * AdvertisedRoute#AdvertisedRoute_ADVERTISEMENT_STATE_FAILED
	//
	//  State of advertisement
	AdvertisementState *string
	// The route that is advertised to on-premise datacenter via Direct Connect format: ipv4-cidr-block
	Cidr *string
}

const AdvertisedRoute_ADDRESS_FAMILY_IPV4 = "IPv4"
const AdvertisedRoute_ADDRESS_FAMILY_IPV6 = "IPv6"
const AdvertisedRoute_ADVERTISEMENT_STATE_SUCCESS = "SUCCESS"
const AdvertisedRoute_ADVERTISEMENT_STATE_FAILED = "FAILED"

func (s *AdvertisedRoute) GetType__() vapiBindings_.BindingType {
	return AdvertisedRouteBindingType()
}

func (s *AdvertisedRoute) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for AdvertisedRoute._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type AggregatedLogicalRouterPortCounters struct {
	// Timestamp when the data was last updated; unset if data source has never updated the data. format: int64
	LastUpdateTimestamp *int64
	Rx                  *LogicalRouterPortCounters
	Tx                  *LogicalRouterPortCounters
}

func (s *AggregatedLogicalRouterPortCounters) GetType__() vapiBindings_.BindingType {
	return AggregatedLogicalRouterPortCountersBindingType()
}

func (s *AggregatedLogicalRouterPortCounters) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for AggregatedLogicalRouterPortCounters._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// AggregationrouteConfigInfo object.
type AggregationRouteConfigInfo struct {
	// Flag set to true implies that the customer provided CIDRs along with the VMC segments will be advertised irrespective of any match/un-match found
	AdvertiseAlways *bool
	// RouteAggregationConfig path information.
	AggregationPath *string
}

func (s *AggregationRouteConfigInfo) GetType__() vapiBindings_.BindingType {
	return AggregationRouteConfigInfoBindingType()
}

func (s *AggregationRouteConfigInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for AggregationRouteConfigInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Detailed information about an API Error
type ApiError struct {
	// Further details about the error
	Details *string
	// A numeric error code format: int64
	ErrorCode *int64
	// Additional data about the error
	ErrorData *vapiData_.StructValue
	// A description of the error
	ErrorMessage *string
	// The module name where the error occurred
	ModuleName *string
	// Other errors related to this error
	RelatedErrors []RelatedApiError
}

func (s *ApiError) GetType__() vapiBindings_.BindingType {
	return ApiErrorBindingType()
}

func (s *ApiError) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ApiError._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Base abstract associated Group connection infomation for the local SDDC.
type AssociatedBaseGroupConnectionInfo struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	Self   *SelfResourceLink
	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected. format: int32
	Revision *int64
	// Timestamp of resource creation format: int64
	CreateTime *int64
	// ID of the user who created this resource
	CreateUser *string
	// Timestamp of last modification format: int64
	LastModifiedTime *int64
	// ID of the user who last modified this resource
	LastModifiedUser *string
	// Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super user and can modify it, but only when providing the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this entity.
	Protection *string
	// Indicates system owned resource
	SystemOwned *bool
	// SDDC Group description
	Description *string
	// Defaults to ID if not set
	DisplayName *string
	// SDDC Group ID
	Id *string
	// Possible values are:
	//
	// * AssociatedBaseGroupConnectionInfo#AssociatedBaseGroupConnectionInfo_RESOURCE_TYPE_ASSOCIATEDTGWGROUPCONNECTIONINFO
	// * AssociatedBaseGroupConnectionInfo#AssociatedBaseGroupConnectionInfo_RESOURCE_TYPE_ASSOCIATEDTGWGROUPCONNECTIONINFOVERSION2
	//
	//  Group connection type
	ResourceType string
	// Opaque identifiers meaningful to the API user
	Tags []Tag
	// marked for delete identifier
	MarkedForDelete *bool
	// Path of its parent
	ParentPath *string
	// Absolute path of this object
	Path *string
	// Path relative from its parent
	RelativePath *string
	// SDDC Group name
	Name *string
}

// Identifier denoting this class, when it is used in polymorphic context.
//
// This value should be assigned to the property which is used to discriminate the actual type used in the polymorphic context.
const AssociatedBaseGroupConnectionInfo__TYPE_IDENTIFIER = "AssociatedBaseGroupConnectionInfo"
const AssociatedBaseGroupConnectionInfo_RESOURCE_TYPE_ASSOCIATEDTGWGROUPCONNECTIONINFO = "AssociatedTgwGroupConnectionInfo"
const AssociatedBaseGroupConnectionInfo_RESOURCE_TYPE_ASSOCIATEDTGWGROUPCONNECTIONINFOVERSION2 = "AssociatedTgwGroupConnectionInfoVersion2"

func (s *AssociatedBaseGroupConnectionInfo) GetType__() vapiBindings_.BindingType {
	return AssociatedBaseGroupConnectionInfoBindingType()
}

func (s *AssociatedBaseGroupConnectionInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for AssociatedBaseGroupConnectionInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Associated Group connection list result
type AssociatedGroupConnectionInfosListResult struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	Self   *SelfResourceLink
	// Opaque cursor to be used for getting next page of records (supplied by current result page)
	Cursor *string
	// Count of results found (across all pages), set only on first page format: int64
	ResultCount *int64
	// If true, results are sorted in ascending order
	SortAscending *bool
	// Field by which records are sorted
	SortBy  *string
	Results []*vapiData_.StructValue
}

func (s *AssociatedGroupConnectionInfosListResult) GetType__() vapiBindings_.BindingType {
	return AssociatedGroupConnectionInfosListResultBindingType()
}

func (s *AssociatedGroupConnectionInfosListResult) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for AssociatedGroupConnectionInfosListResult._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Associated Group connection infomation for the local SDDC by using AWS TGW as a connector.
type AssociatedTgwGroupConnectionInfo struct {
	// TGW external route table ID used for external customers VPCs association
	ExternalRouteTableId *string
	// TGW SDDCs route table ID used for SDDCs association
	SddcsRouteTableId *string
	// Possible values are:
	//
	// * AssociatedTgwGroupConnectionInfo#AssociatedTgwGroupConnectionInfo_STATE_CONNECTED
	// * AssociatedTgwGroupConnectionInfo#AssociatedTgwGroupConnectionInfo_STATE_DISCONNECTED
	//
	//  The TGW attachment state of the SDDC in the Group
	State *string
	// TGW attachment ID for the local SDDC in the Group
	TgwAttachmentId *string
	// TGW ID for the local SDDC in the Group
	TgwId *string
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	Self   *SelfResourceLink
	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected. format: int32
	Revision *int64
	// Timestamp of resource creation format: int64
	CreateTime *int64
	// ID of the user who created this resource
	CreateUser *string
	// Timestamp of last modification format: int64
	LastModifiedTime *int64
	// ID of the user who last modified this resource
	LastModifiedUser *string
	// Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super user and can modify it, but only when providing the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this entity.
	Protection *string
	// Indicates system owned resource
	SystemOwned *bool
	// SDDC Group description
	Description *string
	// Defaults to ID if not set
	DisplayName *string
	// SDDC Group ID
	Id *string
	// Possible values are:
	//
	// * AssociatedBaseGroupConnectionInfo#AssociatedBaseGroupConnectionInfo_RESOURCE_TYPE_ASSOCIATEDTGWGROUPCONNECTIONINFO
	// * AssociatedBaseGroupConnectionInfo#AssociatedBaseGroupConnectionInfo_RESOURCE_TYPE_ASSOCIATEDTGWGROUPCONNECTIONINFOVERSION2
	//
	//  Group connection type
	ResourceType string
	// Opaque identifiers meaningful to the API user
	Tags []Tag
	// marked for delete identifier
	MarkedForDelete *bool
	// Path of its parent
	ParentPath *string
	// Absolute path of this object
	Path *string
	// Path relative from its parent
	RelativePath *string
	// SDDC Group name
	Name *string
}

// Identifier denoting this class, when it is used in polymorphic context.
//
// This value should be assigned to the property which is used to discriminate the actual type used in the polymorphic context.
const AssociatedTgwGroupConnectionInfo__TYPE_IDENTIFIER = "AssociatedTgwGroupConnectionInfo"
const AssociatedTgwGroupConnectionInfo_STATE_CONNECTED = "CONNECTED"
const AssociatedTgwGroupConnectionInfo_STATE_DISCONNECTED = "DISCONNECTED"

func (s *AssociatedTgwGroupConnectionInfo) GetType__() vapiBindings_.BindingType {
	return AssociatedTgwGroupConnectionInfoBindingType()
}

func (s *AssociatedTgwGroupConnectionInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for AssociatedTgwGroupConnectionInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Associated Group connection infomation for the local SDDC by using AWS TGW as a connector and AWS prefix list as route management.
type AssociatedTgwGroupConnectionInfoVersion2 struct {
	// TGW SDDCs route table ID used for SDDCs association
	AssociatedRouteTable *string
	// AWS region in which SDDC is residing
	Region                   *string
	SddcAdvertiseRouteConfig *SddcAdvertiseRouteConfig
	// Possible values are:
	//
	// * AssociatedTgwGroupConnectionInfoVersion2#AssociatedTgwGroupConnectionInfoVersion2_STATE_CONNECTED
	// * AssociatedTgwGroupConnectionInfoVersion2#AssociatedTgwGroupConnectionInfoVersion2_STATE_DISCONNECTED
	//
	//  The TGW attachment state of the SDDC in the Group
	State *string
	// TGW attachment ID for the local SDDC in the Group
	TgwAttachmentId *string
	// TGW ID for the local SDDC in the Group
	TgwId *string
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	Self   *SelfResourceLink
	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected. format: int32
	Revision *int64
	// Timestamp of resource creation format: int64
	CreateTime *int64
	// ID of the user who created this resource
	CreateUser *string
	// Timestamp of last modification format: int64
	LastModifiedTime *int64
	// ID of the user who last modified this resource
	LastModifiedUser *string
	// Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super user and can modify it, but only when providing the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this entity.
	Protection *string
	// Indicates system owned resource
	SystemOwned *bool
	// SDDC Group description
	Description *string
	// Defaults to ID if not set
	DisplayName *string
	// SDDC Group ID
	Id *string
	// Possible values are:
	//
	// * AssociatedBaseGroupConnectionInfo#AssociatedBaseGroupConnectionInfo_RESOURCE_TYPE_ASSOCIATEDTGWGROUPCONNECTIONINFO
	// * AssociatedBaseGroupConnectionInfo#AssociatedBaseGroupConnectionInfo_RESOURCE_TYPE_ASSOCIATEDTGWGROUPCONNECTIONINFOVERSION2
	//
	//  Group connection type
	ResourceType string
	// Opaque identifiers meaningful to the API user
	Tags []Tag
	// marked for delete identifier
	MarkedForDelete *bool
	// Path of its parent
	ParentPath *string
	// Absolute path of this object
	Path *string
	// Path relative from its parent
	RelativePath *string
	// SDDC Group name
	Name *string
}

// Identifier denoting this class, when it is used in polymorphic context.
//
// This value should be assigned to the property which is used to discriminate the actual type used in the polymorphic context.
const AssociatedTgwGroupConnectionInfoVersion2__TYPE_IDENTIFIER = "AssociatedTgwGroupConnectionInfoVersion2"
const AssociatedTgwGroupConnectionInfoVersion2_STATE_CONNECTED = "CONNECTED"
const AssociatedTgwGroupConnectionInfoVersion2_STATE_DISCONNECTED = "DISCONNECTED"

func (s *AssociatedTgwGroupConnectionInfoVersion2) GetType__() vapiBindings_.BindingType {
	return AssociatedTgwGroupConnectionInfoVersion2BindingType()
}

func (s *AssociatedTgwGroupConnectionInfoVersion2) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for AssociatedTgwGroupConnectionInfoVersion2._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// AWS prefix entry
type AwsPrefixInfo struct {
	// Possible values are:
	//
	// * AwsPrefixInfo#AwsPrefixInfo_ADDRESS_FAMILY_IPV4
	// * AwsPrefixInfo#AwsPrefixInfo_ADDRESS_FAMILY_IPV6
	//
	//  IP address type
	AddressFamily *string
	// Prefix ID
	Id *string
	// Prefix Name
	Name *string
	// AWS region
	Region *string
	// Prefix max entries format: int32
	Size *int64
}

const AwsPrefixInfo_ADDRESS_FAMILY_IPV4 = "IPv4"
const AwsPrefixInfo_ADDRESS_FAMILY_IPV6 = "IPv6"

func (s *AwsPrefixInfo) GetType__() vapiBindings_.BindingType {
	return AwsPrefixInfoBindingType()
}

func (s *AwsPrefixInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for AwsPrefixInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Information related to AWS resource share that is shared with linked VPC account.
type AwsResourceShareInfo struct {
	// Amazon Resource Name (ARN) of AWS resource share.
	AwsResourceShareArn *string
	// Name of AWS resource share.
	AwsResourceShareName *string
	// Possible values are:
	//
	// * AwsResourceShareInfo#AwsResourceShareInfo_AWS_RESOURCE_SHARE_STATE_ACTIVE
	// * AwsResourceShareInfo#AwsResourceShareInfo_AWS_RESOURCE_SHARE_STATE_PENDING
	// * AwsResourceShareInfo#AwsResourceShareInfo_AWS_RESOURCE_SHARE_STATE_FAILED
	//
	//  State of AWS resource share.
	AwsResourceShareState *string
}

const AwsResourceShareInfo_AWS_RESOURCE_SHARE_STATE_ACTIVE = "ACTIVE"
const AwsResourceShareInfo_AWS_RESOURCE_SHARE_STATE_PENDING = "PENDING"
const AwsResourceShareInfo_AWS_RESOURCE_SHARE_STATE_FAILED = "FAILED"

func (s *AwsResourceShareInfo) GetType__() vapiBindings_.BindingType {
	return AwsResourceShareInfoBindingType()
}

func (s *AwsResourceShareInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for AwsResourceShareInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Advertised bgp routes
type BGPAdvertisedRoutes struct {
	// Routes advertised to on-premise datacenter via Direct Connect
	AdvertisedRoutes []AdvertisedRoute
	// Number of routes failed to advertise format: int32
	FailedAdvertisedRoutes *int64
}

func (s *BGPAdvertisedRoutes) GetType__() vapiBindings_.BindingType {
	return BGPAdvertisedRoutesBindingType()
}

func (s *BGPAdvertisedRoutes) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for BGPAdvertisedRoutes._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Learned bgp routes
type BGPLearnedRoutes struct {
	// The route that is learned from BGP via Direct Connect format: ipv4-cidr-block
	Ipv4Cidr []string
	// The route that is learned from BGP via Direct Connect format: ipv6-cidr-block
	Ipv6Cidr []string
}

func (s *BGPLearnedRoutes) GetType__() vapiBindings_.BindingType {
	return BGPLearnedRoutesBindingType()
}

func (s *BGPLearnedRoutes) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for BGPLearnedRoutes._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// CloudProviderInformation
type CloudProviderSddcNetworkingStateInfo struct {
	// Account ID
	AccountId *string
	// Availability Zone
	AvailabilityZone *string
	// Cloud Formation Stack Name
	CloudFormationStackName *string
	// Possible values are:
	//
	// * CloudProviderSddcNetworkingStateInfo#CloudProviderSddcNetworkingStateInfo_CONNECTIVITY_STATUS_CONNECTED
	// * CloudProviderSddcNetworkingStateInfo#CloudProviderSddcNetworkingStateInfo_CONNECTIVITY_STATUS_DISCONNECTED
	//
	//  Connectivity Status
	ConnectivityStatus *string
	// EC2 Status
	Ec2Enabled *bool
	// IAM role names
	IamRoleNames *string
	// S3 Status
	S3Enabled *bool
	// Active network mapping List used for linked VPC traffic.
	TrafficGroupEniMappingInfos []TrafficGroupEniMappingInfo
	// Information related to the subnets where linked ENIs were created.
	Vpc []LinkedSubnetInfo
	// VPC CIDR addresses
	VpcCidr []string
	// VPC id
	VpcId *string
}

const CloudProviderSddcNetworkingStateInfo_CONNECTIVITY_STATUS_CONNECTED = "CONNECTED"
const CloudProviderSddcNetworkingStateInfo_CONNECTIVITY_STATUS_DISCONNECTED = "DISCONNECTED"

func (s *CloudProviderSddcNetworkingStateInfo) GetType__() vapiBindings_.BindingType {
	return CloudProviderSddcNetworkingStateInfoBindingType()
}

func (s *CloudProviderSddcNetworkingStateInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for CloudProviderSddcNetworkingStateInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// This will be used by VMC UI to return the Compute Gateway (MGW) information as part of SddcNetworkingStateInfo API.
type ComputeGatewaySddcNetworkingStateInfo struct {
	// CGW DNS Forwarder Zone Upstream Servers. format: ipv4-cidr-block
	Dns []string
	// Total Firewall Rules format: int32
	TotalFirewallRules *int64
	// Total segments format: int32
	TotalSegments *int64
	// Total Tier1 Gateways format: int32
	TotalTier1Gateways *int64
}

func (s *ComputeGatewaySddcNetworkingStateInfo) GetType__() vapiBindings_.BindingType {
	return ComputeGatewaySddcNetworkingStateInfoBindingType()
}

func (s *ComputeGatewaySddcNetworkingStateInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ComputeGatewaySddcNetworkingStateInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// A list of status of 'Enabled/Disabled' for a service connected to a linked vpc
type ConnectedServiceListResult struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	Self   *SelfResourceLink
	// Opaque cursor to be used for getting next page of records (supplied by current result page)
	Cursor *string
	// Count of results found (across all pages), set only on first page format: int64
	ResultCount *int64
	// If true, results are sorted in ascending order
	SortAscending *bool
	// Field by which records are sorted
	SortBy *string
	// Connected service status list
	Results []ConnectedServiceStatus
}

func (s *ConnectedServiceListResult) GetType__() vapiBindings_.BindingType {
	return ConnectedServiceListResultBindingType()
}

func (s *ConnectedServiceListResult) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ConnectedServiceListResult._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Status of 'Enabled/Disabled' for a service connected to a linked vpc
type ConnectedServiceStatus struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	Self   *SelfResourceLink
	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected. format: int32
	Revision *int64
	// Timestamp of resource creation format: int64
	CreateTime *int64
	// ID of the user who created this resource
	CreateUser *string
	// Timestamp of last modification format: int64
	LastModifiedTime *int64
	// ID of the user who last modified this resource
	LastModifiedUser *string
	// Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super user and can modify it, but only when providing the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this entity.
	Protection *string
	// Indicates system owned resource
	SystemOwned *bool
	// Description of this resource
	Description *string
	// Defaults to ID if not set
	DisplayName *string
	// Unique identifier of this resource
	Id *string
	// The type of this resource.
	ResourceType *string
	// Opaque identifiers meaningful to the API user
	Tags []Tag
	// marked for delete identifier
	MarkedForDelete *bool
	// Path of its parent
	ParentPath *string
	// Absolute path of this object
	Path *string
	// Path relative from its parent
	RelativePath *string
	// status of service
	Enabled *bool
	// service name
	Name *string
}

func (s *ConnectedServiceStatus) GetType__() vapiBindings_.BindingType {
	return ConnectedServiceStatusBindingType()
}

func (s *ConnectedServiceStatus) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ConnectedServiceStatus._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Connectivity Endpoint object.
type ConnectivityEndpointConfig struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	Self   *SelfResourceLink
	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected. format: int32
	Revision *int64
	// Timestamp of resource creation format: int64
	CreateTime *int64
	// ID of the user who created this resource
	CreateUser *string
	// Timestamp of last modification format: int64
	LastModifiedTime *int64
	// ID of the user who last modified this resource
	LastModifiedUser *string
	// Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super user and can modify it, but only when providing the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this entity.
	Protection *string
	// Indicates system owned resource
	SystemOwned *bool
	// Description of the connectivityEndpointConfig.
	Description *string
	// Name of the connectivityEndpointConfig.
	DisplayName *string
	// Id for the connectivityEndpointConfig.
	Id *string
	// The type of this resource.
	ResourceType *string
	// Opaque identifiers meaningful to the API user
	Tags []Tag
	// marked for delete identifier
	MarkedForDelete *bool
	// Path of its parent
	ParentPath *string
	// Absolute path of this object
	Path *string
	// Path relative from its parent
	RelativePath *string
}

func (s *ConnectivityEndpointConfig) GetType__() vapiBindings_.BindingType {
	return ConnectivityEndpointConfigBindingType()
}

func (s *ConnectivityEndpointConfig) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ConnectivityEndpointConfig._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// List of connectivityEndpoint objects.
type ConnectivityEndpointListResult struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	Self   *SelfResourceLink
	// Opaque cursor to be used for getting next page of records (supplied by current result page)
	Cursor *string
	// Count of results found (across all pages), set only on first page format: int64
	ResultCount *int64
	// If true, results are sorted in ascending order
	SortAscending *bool
	// Field by which records are sorted
	SortBy *string
	// Items of connectivityEndpointConfig object
	Results []ConnectivityEndpointConfig
}

func (s *ConnectivityEndpointListResult) GetType__() vapiBindings_.BindingType {
	return ConnectivityEndpointListResultBindingType()
}

func (s *ConnectivityEndpointListResult) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ConnectivityEndpointListResult._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Base type for CSV result.
type CsvListResult struct {
	// File name set by HTTP server if API returns CSV result as a file.
	FileName *string
}

func (s *CsvListResult) GetType__() vapiBindings_.BindingType {
	return CsvListResultBindingType()
}

func (s *CsvListResult) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for CsvListResult._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Base type for CSV records.
type CsvRecord struct {
}

func (s *CsvRecord) GetType__() vapiBindings_.BindingType {
	return CsvRecordBindingType()
}

func (s *CsvRecord) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for CsvRecord._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type CustomerBgpPeerInfo struct {
	// BGP MD5 authentication key
	BgpMd5AuthenticationKey *string
	// Virtual Circuit display name
	DisplayName *string
	// Index of the Virtual Circuit format: int32
	Index *int64
	// SDDC Underlay BGP peer IP/prefix format: ipv4-cidr-block
	LocalPeerIp *string
	// Provider token
	ProviderToken *string
	// Provider token expiration format: int64
	ProviderTokenExpiration *int64
	// Customer BGP peer ASN
	RemotePeerAsn *string
	// Customer BGP peer IP/prefix format: ipv4-cidr-block
	RemotePeerIp *string
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

type CustomerBgpPeerInfoStatus struct {
	// Index of the Virtual Circuit format: int32
	Index *int64
	// Possible values are:
	//
	// * CustomerBgpPeerInfoStatus#CustomerBgpPeerInfoStatus_STATUS_PENDING
	// * CustomerBgpPeerInfoStatus#CustomerBgpPeerInfoStatus_STATUS_CONNECTED
	// * CustomerBgpPeerInfoStatus#CustomerBgpPeerInfoStatus_STATUS_DISCONNECTED
	// * CustomerBgpPeerInfoStatus#CustomerBgpPeerInfoStatus_STATUS_UNKNOWN
	//
	//  Status of the Virtual Circuit
	Status *string
	// Detailed information of the status returned by the provider for the Virtual Circuit
	StatusDetail *string
}

const CustomerBgpPeerInfoStatus_STATUS_PENDING = "PENDING"
const CustomerBgpPeerInfoStatus_STATUS_CONNECTED = "CONNECTED"
const CustomerBgpPeerInfoStatus_STATUS_DISCONNECTED = "DISCONNECTED"
const CustomerBgpPeerInfoStatus_STATUS_UNKNOWN = "UNKNOWN"

func (s *CustomerBgpPeerInfoStatus) GetType__() vapiBindings_.BindingType {
	return CustomerBgpPeerInfoStatusBindingType()
}

func (s *CustomerBgpPeerInfoStatus) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for CustomerBgpPeerInfoStatus._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Direct Connect BGP related information
type DirectConnectBgpInfo struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	Self   *SelfResourceLink
	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected. format: int32
	Revision *int64
	// Timestamp of resource creation format: int64
	CreateTime *int64
	// ID of the user who created this resource
	CreateUser *string
	// Timestamp of last modification format: int64
	LastModifiedTime *int64
	// ID of the user who last modified this resource
	LastModifiedUser *string
	// Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super user and can modify it, but only when providing the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this entity.
	Protection *string
	// Indicates system owned resource
	SystemOwned *bool
	// Description of this resource
	Description *string
	// Defaults to ID if not set
	DisplayName *string
	// Unique identifier of this resource
	Id *string
	// The type of this resource.
	ResourceType *string
	// Opaque identifiers meaningful to the API user
	Tags []Tag
	// marked for delete identifier
	MarkedForDelete *bool
	// Path of its parent
	ParentPath *string
	// Absolute path of this object
	Path *string
	// Path relative from its parent
	RelativePath *string
	// The ASN paired with the VGW attached to the VPC. AWS allowed private BGP ASN range - [64512, 65534] and [4200000000, 4294967294]. If omitted in the payload, BGP ASN will not be modified.
	LocalAsNum *string
	// Maximum transmission unit allowed by the VIF format: int32
	Mtu *int64
	// Possible values are:
	//
	// * DirectConnectBgpInfo#DirectConnectBgpInfo_ROUTE_PREFERENCE_DIRECT_CONNECT_PREFERRED_OVER_VPN
	// * DirectConnectBgpInfo#DirectConnectBgpInfo_ROUTE_PREFERENCE_VPN_PREFERRED_OVER_DIRECT_CONNECT
	//
	//  Direct connect route preference over VPN routes. If omitted in the payload, route preference will not be modified.
	RoutePreference *string
}

const DirectConnectBgpInfo_ROUTE_PREFERENCE_DIRECT_CONNECT_PREFERRED_OVER_VPN = "DIRECT_CONNECT_PREFERRED_OVER_VPN"
const DirectConnectBgpInfo_ROUTE_PREFERENCE_VPN_PREFERRED_OVER_DIRECT_CONNECT = "VPN_PREFERRED_OVER_DIRECT_CONNECT"

func (s *DirectConnectBgpInfo) GetType__() vapiBindings_.BindingType {
	return DirectConnectBgpInfoBindingType()
}

func (s *DirectConnectBgpInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for DirectConnectBgpInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// This will be used by VMC UI to return the Direct Connectivity information as part of SddcNetworkingStateInfo API.
type DirectConnectSddcNetworkingStateInfo struct {
	// The total number of Direct Connect advertised routes. format: int32
	AdvertisedRoutes *int64
	// The total number of attached Direct Connect private VIFs with BGP status as down. format: int32
	ConnectivityDown *int64
	// The total number of attached Direct Connect private VIFs with BGP status as up. format: int32
	ConnectivityUp *int64
	// The list of attached Direct Connect private VIFs with BGP status as down.
	DownVifNames []string
	// The total number of Direct Connect learned routes. format: int32
	LearnedRoutes *int64
	// The total number of attached Direct Connect private VIFs. format: int32
	TotalAttachedVifsCount *int64
}

func (s *DirectConnectSddcNetworkingStateInfo) GetType__() vapiBindings_.BindingType {
	return DirectConnectSddcNetworkingStateInfoBindingType()
}

func (s *DirectConnectSddcNetworkingStateInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for DirectConnectSddcNetworkingStateInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// The list of prefix lists with the corresponding edge index.
type DistributedPrefixList struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	Self   *SelfResourceLink
	// Opaque cursor to be used for getting next page of records (supplied by current result page)
	Cursor *string
	// Count of results found (across all pages), set only on first page format: int64
	ResultCount *int64
	// If true, results are sorted in ascending order
	SortAscending *bool
	// Field by which records are sorted
	SortBy *string
	// Refers to the edge to which the prefix list applies to, index should be either 1 or 2. format: int32
	Index *int64
	// List of prefix lists.
	PrefixLists []string
}

func (s *DistributedPrefixList) GetType__() vapiBindings_.BindingType {
	return DistributedPrefixListBindingType()
}

func (s *DistributedPrefixList) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for DistributedPrefixList._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// External Connectivity configuration for north-south traffic. For eg., in AWS case, this would refer to Direct Connect config. For Dimension, this would refer to the config at Upstream Intranet interface to Tor.
type ExternalConnectivityConfig struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	Self   *SelfResourceLink
	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected. format: int32
	Revision *int64
	// Timestamp of resource creation format: int64
	CreateTime *int64
	// ID of the user who created this resource
	CreateUser *string
	// Timestamp of last modification format: int64
	LastModifiedTime *int64
	// ID of the user who last modified this resource
	LastModifiedUser *string
	// Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super user and can modify it, but only when providing the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this entity.
	Protection *string
	// Indicates system owned resource
	SystemOwned *bool
	// Description of this resource
	Description *string
	// Defaults to ID if not set
	DisplayName *string
	// Unique identifier of this resource
	Id *string
	// The type of this resource.
	ResourceType *string
	// Opaque identifiers meaningful to the API user
	Tags []Tag
	// marked for delete identifier
	MarkedForDelete *bool
	// Path of its parent
	ParentPath *string
	// Absolute path of this object
	Path *string
	// Path relative from its parent
	RelativePath      *string
	InternetInterface *InterfaceConfig
	// Uplink MTU of internet traffic in edge tier-0 router port. format: int32
	InternetMtu       *int64
	IntranetInterface *InterfaceConfig
	// Uplink MTU of direct connect, sddc-grouping and outposts traffic in edge tier-0 router port. format: int32
	IntranetMtu       *int64
	ServicesInterface *InterfaceConfig
	// Uplink MTU of connected VPC traffic in edge tier-0 router port. format: int32
	ServicesMtu *int64
}

func (s *ExternalConnectivityConfig) GetType__() vapiBindings_.BindingType {
	return ExternalConnectivityConfigBindingType()
}

func (s *ExternalConnectivityConfig) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ExternalConnectivityConfig._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// External Connectivity Information
type ExternalConnectivityInfo struct {
	// Object to hold all connectivity endpoints
	Connectivities []string
	// External Endpoint Type
	EndpointType *string
}

func (s *ExternalConnectivityInfo) GetType__() vapiBindings_.BindingType {
	return ExternalConnectivityInfoBindingType()
}

func (s *ExternalConnectivityInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ExternalConnectivityInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type ExternalConnectivityStatus struct {
	// Possible values are:
	//
	// * ExternalConnectivityStatus#ExternalConnectivityStatus_STATUS_CONNECTED
	// * ExternalConnectivityStatus#ExternalConnectivityStatus_STATUS_DISCONNECTED
	//
	//  Indicates whether connectivity is established for a connectivity type.
	Status *string
	// Connectivity endpoint
	Type_ *string
}

const ExternalConnectivityStatus_STATUS_CONNECTED = "CONNECTED"
const ExternalConnectivityStatus_STATUS_DISCONNECTED = "DISCONNECTED"

func (s *ExternalConnectivityStatus) GetType__() vapiBindings_.BindingType {
	return ExternalConnectivityStatusBindingType()
}

func (s *ExternalConnectivityStatus) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ExternalConnectivityStatus._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// External connectivity list
type ExternalConnectivityStatusList struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	Self   *SelfResourceLink
	// Opaque cursor to be used for getting next page of records (supplied by current result page)
	Cursor *string
	// Count of results found (across all pages), set only on first page format: int64
	ResultCount *int64
	// If true, results are sorted in ascending order
	SortAscending *bool
	// Field by which records are sorted
	SortBy *string
	// External connectivity list
	Connectivities []ExternalConnectivityStatus
}

func (s *ExternalConnectivityStatusList) GetType__() vapiBindings_.BindingType {
	return ExternalConnectivityStatusListBindingType()
}

func (s *ExternalConnectivityStatusList) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ExternalConnectivityStatusList._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// External SDDC connectivity
type ExternalSddcConnectivity struct {
	// Possible values are:
	//
	// * ExternalSddcConnectivity#ExternalSddcConnectivity_CONNECTIVITY_TYPE_DIRECT_CONNECT
	// * ExternalSddcConnectivity#ExternalSddcConnectivity_CONNECTIVITY_TYPE_DEPLOYMENT_CONNECTIVITY_GROUP
	// * ExternalSddcConnectivity#ExternalSddcConnectivity_CONNECTIVITY_TYPE_LGW
	//
	//  The external SDDC connectivity type is used by the SDDC for the L3 connectivity. DIRECT_CONNECT means that the external SDDC connectivity is through AWS Direct Connect. DEPLOYMENT_CONNECTIVITY_GROUP means that the external SDDC connectivity is through AWS TGW. LGW means that the external SDDC connectivity is through AWS local gateway.
	ConnectivityType *string
	// The list of regions of the route for the connectivity
	Regions []string
	// The source of the route for the connectivity
	Source *string
	// Possible values are:
	//
	// * ExternalSddcConnectivity#ExternalSddcConnectivity_STATUS_SUCCEEDED
	// * ExternalSddcConnectivity#ExternalSddcConnectivity_STATUS_FAILED
	// * ExternalSddcConnectivity#ExternalSddcConnectivity_STATUS_PENDING
	//
	//  The status of the route for the connectivity
	Status *string
	// The error message if the status is FAILED, the optional warning message if the status is SUCCEEDED.
	StatusMessage *string
}

const ExternalSddcConnectivity_CONNECTIVITY_TYPE_DIRECT_CONNECT = "DIRECT_CONNECT"
const ExternalSddcConnectivity_CONNECTIVITY_TYPE_DEPLOYMENT_CONNECTIVITY_GROUP = "DEPLOYMENT_CONNECTIVITY_GROUP"
const ExternalSddcConnectivity_CONNECTIVITY_TYPE_LGW = "LGW"
const ExternalSddcConnectivity_STATUS_SUCCEEDED = "SUCCEEDED"
const ExternalSddcConnectivity_STATUS_FAILED = "FAILED"
const ExternalSddcConnectivity_STATUS_PENDING = "PENDING"

func (s *ExternalSddcConnectivity) GetType__() vapiBindings_.BindingType {
	return ExternalSddcConnectivityBindingType()
}

func (s *ExternalSddcConnectivity) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ExternalSddcConnectivity._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// External SDDC route
type ExternalSddcRoute struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	Self   *SelfResourceLink
	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected. format: int32
	Revision *int64
	// Timestamp of resource creation format: int64
	CreateTime *int64
	// ID of the user who created this resource
	CreateUser *string
	// Timestamp of last modification format: int64
	LastModifiedTime *int64
	// ID of the user who last modified this resource
	LastModifiedUser *string
	// Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super user and can modify it, but only when providing the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this entity.
	Protection *string
	// Indicates system owned resource
	SystemOwned *bool
	// Description of this resource
	Description *string
	// Defaults to ID if not set
	DisplayName *string
	// Unique identifier of this resource
	Id *string
	// The type of this resource.
	ResourceType *string
	// Opaque identifiers meaningful to the API user
	Tags []Tag
	// marked for delete identifier
	MarkedForDelete *bool
	// Path of its parent
	ParentPath *string
	// Absolute path of this object
	Path *string
	// Path relative from its parent
	RelativePath *string
	// The route used for what kind of connectivities
	Connectivities []ExternalSddcConnectivity
	// Destination IP CIDR Block format: ipv4-cidr-block
	Destination *string
	// Flag to indicate VIF attached status.
	IsVifAttached *bool
	// Source Path of CIDR
	SourcePath *string
	// Path is summarised or not.
	Summarized *bool
}

func (s *ExternalSddcRoute) GetType__() vapiBindings_.BindingType {
	return ExternalSddcRouteBindingType()
}

func (s *ExternalSddcRoute) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ExternalSddcRoute._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// CSV record for External SDDC route
type ExternalSddcRouteCsvRecord struct {
	// The connectivity datails contains status of route, source of the route, connectivity type
	ConnectivityDetails *string
	// Destination IP CIDR Block format: ipv4-cidr-block
	Destination *string
}

func (s *ExternalSddcRouteCsvRecord) GetType__() vapiBindings_.BindingType {
	return ExternalSddcRouteCsvRecordBindingType()
}

func (s *ExternalSddcRouteCsvRecord) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ExternalSddcRouteCsvRecord._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// External SDDC routes list result
type ExternalSddcRoutesListResult struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	Self   *SelfResourceLink
	// Opaque cursor to be used for getting next page of records (supplied by current result page)
	Cursor *string
	// Count of results found (across all pages), set only on first page format: int64
	ResultCount *int64
	// If true, results are sorted in ascending order
	SortAscending *bool
	// Field by which records are sorted
	SortBy *string
	Routes []ExternalSddcRoute
}

func (s *ExternalSddcRoutesListResult) GetType__() vapiBindings_.BindingType {
	return ExternalSddcRoutesListResultBindingType()
}

func (s *ExternalSddcRoutesListResult) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ExternalSddcRoutesListResult._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// External SDDC routes list result in CSV format
type ExternalSddcRoutesListResultInCsvFormat struct {
	// File name set by HTTP server if API returns CSV result as a file.
	FileName *string
	Results  []ExternalSddcRouteCsvRecord
}

func (s *ExternalSddcRoutesListResultInCsvFormat) GetType__() vapiBindings_.BindingType {
	return ExternalSddcRoutesListResultInCsvFormatBindingType()
}

func (s *ExternalSddcRoutesListResultInCsvFormat) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ExternalSddcRoutesListResultInCsvFormat._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Fw rule filtered with Any or/and 0.0.0.0/0
type FilteredMgwGatewayPolicies struct {
	// Provider error message
	ErrorMsg *string
	// Invalid rule details
	InvalidRules []FilteredMgwGatewayPoliciesKeyValuePairs
}

func (s *FilteredMgwGatewayPolicies) GetType__() vapiBindings_.BindingType {
	return FilteredMgwGatewayPoliciesBindingType()
}

func (s *FilteredMgwGatewayPolicies) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for FilteredMgwGatewayPolicies._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Rule details for Any and 0.0.0.0/0.
type FilteredMgwGatewayPoliciesKeyValuePairs struct {
	// Key
	Key *string
	// Value
	Value []FilteredMgwGatewayPoliciesObject
}

func (s *FilteredMgwGatewayPoliciesKeyValuePairs) GetType__() vapiBindings_.BindingType {
	return FilteredMgwGatewayPoliciesKeyValuePairsBindingType()
}

func (s *FilteredMgwGatewayPoliciesKeyValuePairs) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for FilteredMgwGatewayPoliciesKeyValuePairs._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Rule object with id and name.
type FilteredMgwGatewayPoliciesObject struct {
	// Rule id format: int32
	RuleId *int64
	// Rule name
	RuleName *string
}

func (s *FilteredMgwGatewayPoliciesObject) GetType__() vapiBindings_.BindingType {
	return FilteredMgwGatewayPoliciesObjectBindingType()
}

func (s *FilteredMgwGatewayPoliciesObject) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for FilteredMgwGatewayPoliciesObject._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// A list of fields to include in query results
type IncludedFieldsParameters struct {
	// Note - this parameter currently only works when used with the search APIs /policy/api/v1/search/query and /policy/api/v1/search/dsl. It is ignored for other list APIs.
	IncludedFields *string
}

func (s *IncludedFieldsParameters) GetType__() vapiBindings_.BindingType {
	return IncludedFieldsParametersBindingType()
}

func (s *IncludedFieldsParameters) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for IncludedFieldsParameters._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Interface configuration.
type InterfaceConfig struct {
	// Depicts route filtering state corresponding to connectivity endpoint
	RouteFilteringState *bool
	// Possible values are:
	//
	// * InterfaceConfig#InterfaceConfig_URPF_MODE_STRICT
	// * InterfaceConfig#InterfaceConfig_URPF_MODE_NONE
	//
	//  Unicast Reverse Path Forwarding mode of interface.
	UrpfMode *string
}

const InterfaceConfig_URPF_MODE_STRICT = "STRICT"
const InterfaceConfig_URPF_MODE_NONE = "NONE"

func (s *InterfaceConfig) GetType__() vapiBindings_.BindingType {
	return InterfaceConfigBindingType()
}

func (s *InterfaceConfig) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for InterfaceConfig._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Intranet connection configuration parameters.
type IntranetConnectionConfiguration struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	Self   *SelfResourceLink
	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected. format: int32
	Revision *int64
	// Timestamp of resource creation format: int64
	CreateTime *int64
	// ID of the user who created this resource
	CreateUser *string
	// Timestamp of last modification format: int64
	LastModifiedTime *int64
	// ID of the user who last modified this resource
	LastModifiedUser *string
	// Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super user and can modify it, but only when providing the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this entity.
	Protection *string
	// Indicates system owned resource
	SystemOwned *bool
	// Intranet Connection description
	Description *string
	// Intranet Connection name
	DisplayName *string
	// Intranet Connection Identifier
	Id *string
	// The type of this resource.
	ResourceType *string
	// Opaque identifiers meaningful to the API user
	Tags []Tag
	// marked for delete identifier
	MarkedForDelete *bool
	// Path of its parent
	ParentPath *string
	// Absolute path of this object
	Path *string
	// Path relative from its parent
	RelativePath *string
	// Possible values are:
	//
	// * IntranetConnectionConfiguration#IntranetConnectionConfiguration_CONNECTION_REDUNDANCY_SINGLE
	// * IntranetConnectionConfiguration#IntranetConnectionConfiguration_CONNECTION_REDUNDANCY_REDUNDANT
	//
	//  Connection Redundancy type
	ConnectionRedundancy *string
	// Customer BGP peer info list. Must contain 1 entry if connection_redundancy = SINGLE, two entries if connection_redundancy = REDUNDANT
	VirtualCircuits []CustomerBgpPeerInfo
}

const IntranetConnectionConfiguration_CONNECTION_REDUNDANCY_SINGLE = "SINGLE"
const IntranetConnectionConfiguration_CONNECTION_REDUNDANCY_REDUNDANT = "REDUNDANT"

func (s *IntranetConnectionConfiguration) GetType__() vapiBindings_.BindingType {
	return IntranetConnectionConfigurationBindingType()
}

func (s *IntranetConnectionConfiguration) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for IntranetConnectionConfiguration._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// List of Intranet Connection configuration
type IntranetConnectionConfigurationListResult struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	Self   *SelfResourceLink
	// Opaque cursor to be used for getting next page of records (supplied by current result page)
	Cursor *string
	// Count of results found (across all pages), set only on first page format: int64
	ResultCount *int64
	// If true, results are sorted in ascending order
	SortAscending *bool
	// Field by which records are sorted
	SortBy *string
	// List of Intranet Connection configurations
	Results []IntranetConnectionConfiguration
}

func (s *IntranetConnectionConfigurationListResult) GetType__() vapiBindings_.BindingType {
	return IntranetConnectionConfigurationListResultBindingType()
}

func (s *IntranetConnectionConfigurationListResult) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for IntranetConnectionConfigurationListResult._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type IntranetConnectivityStatus struct {
	// Possible values are:
	//
	// * IntranetConnectivityStatus#IntranetConnectivityStatus_AGGREGATE_STATUS_PENDING
	// * IntranetConnectivityStatus#IntranetConnectivityStatus_AGGREGATE_STATUS_CONNECTED
	// * IntranetConnectivityStatus#IntranetConnectivityStatus_AGGREGATE_STATUS_DISCONNECTED
	// * IntranetConnectivityStatus#IntranetConnectivityStatus_AGGREGATE_STATUS_UNKNOWN
	// * IntranetConnectivityStatus#IntranetConnectivityStatus_AGGREGATE_STATUS_DEGRADED
	//
	//  Indicates aggregate status of all Virtual Circuits in an intranet connection.
	AggregateStatus *string
	// Intranet Connection Identifier
	Id              *string
	VirtualCircuits []CustomerBgpPeerInfoStatus
}

const IntranetConnectivityStatus_AGGREGATE_STATUS_PENDING = "PENDING"
const IntranetConnectivityStatus_AGGREGATE_STATUS_CONNECTED = "CONNECTED"
const IntranetConnectivityStatus_AGGREGATE_STATUS_DISCONNECTED = "DISCONNECTED"
const IntranetConnectivityStatus_AGGREGATE_STATUS_UNKNOWN = "UNKNOWN"
const IntranetConnectivityStatus_AGGREGATE_STATUS_DEGRADED = "DEGRADED"

func (s *IntranetConnectivityStatus) GetType__() vapiBindings_.BindingType {
	return IntranetConnectivityStatusBindingType()
}

func (s *IntranetConnectivityStatus) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for IntranetConnectivityStatus._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type IpAttachmentPair struct {
	// Attachment id which maps to management VM IP
	AttachmentId *string
	// Management VM IP Address format: ipv4
	Ip *string
}

func (s *IpAttachmentPair) GetType__() vapiBindings_.BindingType {
	return IpAttachmentPairBindingType()
}

func (s *IpAttachmentPair) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for IpAttachmentPair._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Information related to a subnet where linked ENIs were created.
type LinkedSubnetInfo struct {
	// This describes the availability zone name to which the physical availability zones are randomly mapped which helps in distributing resources across AZs.
	AvailabilityZone *string
	// This describes the AZ ID which is a unique and consistent identifier for an availability zone across all AWS accounts.
	AvailabilityZoneId *string
	// Linked subnet CIDR format: ipv4-cidr-block
	Cidr *string
	// Linked subnet identifier
	Id *string
	// Linked subnet IPv6 CIDRs format: ipv6-cidr-block
	Ipv6Cidr []string
}

func (s *LinkedSubnetInfo) GetType__() vapiBindings_.BindingType {
	return LinkedSubnetInfoBindingType()
}

func (s *LinkedSubnetInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for LinkedSubnetInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Information related to linked VPC ipv6 enablement support including ipv6 enablement mode status etc.
type LinkedVpciPv6ModeInfo struct {
	// Possible values are:
	//
	// * LinkedVpciPv6ModeInfo#LinkedVpciPv6ModeInfo_IPV6MODE_DISABLED
	// * LinkedVpciPv6ModeInfo#LinkedVpciPv6ModeInfo_IPV6MODE_PENDING
	// * LinkedVpciPv6ModeInfo#LinkedVpciPv6ModeInfo_IPV6MODE_ENABLED
	//
	//  Current state of Linked VPC IPv6 enablement
	Ipv6Mode *string
}

const LinkedVpciPv6ModeInfo_IPV6MODE_DISABLED = "DISABLED"
const LinkedVpciPv6ModeInfo_IPV6MODE_PENDING = "PENDING"
const LinkedVpciPv6ModeInfo_IPV6MODE_ENABLED = "ENABLED"

func (s *LinkedVpciPv6ModeInfo) GetType__() vapiBindings_.BindingType {
	return LinkedVpciPv6ModeInfoBindingType()
}

func (s *LinkedVpciPv6ModeInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for LinkedVpciPv6ModeInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Linked VPC info
type LinkedVpcInfo struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	Self   *SelfResourceLink
	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected. format: int32
	Revision *int64
	// Timestamp of resource creation format: int64
	CreateTime *int64
	// ID of the user who created this resource
	CreateUser *string
	// Timestamp of last modification format: int64
	LastModifiedTime *int64
	// ID of the user who last modified this resource
	LastModifiedUser *string
	// Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super user and can modify it, but only when providing the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this entity.
	Protection *string
	// Indicates system owned resource
	SystemOwned *bool
	// Description of this resource
	Description *string
	// Defaults to ID if not set
	DisplayName *string
	// Unique identifier of this resource
	Id *string
	// The type of this resource.
	ResourceType *string
	// Opaque identifiers meaningful to the API user
	Tags []Tag
	// marked for delete identifier
	MarkedForDelete *bool
	// Path of its parent
	ParentPath *string
	// Absolute path of this object
	Path *string
	// Path relative from its parent
	RelativePath *string
	// Active network interface used for linked vpc traffic
	ActiveEni *string
	// ARN role for linked VPC operations
	ArnRole *string
	// External identifier for ARN role
	ExternalId *string
	// Linked VPC account number
	LinkedAccount *string
	// Linked VPC CIDRs format: ipv4-cidr-block
	LinkedVpcAddresses []string
	// Linked VPC identifier
	LinkedVpcId *string
	// Linked VPC IPv6 CIDRs format: ipv6-cidr-block
	LinkedVpcIpv6Addresses         []string
	LinkedVpcIpv6ModeInfo          *LinkedVpciPv6ModeInfo
	LinkedVpcManagedPrefixListInfo *LinkedVpcManagedPrefixListSupportInfo
	// The IPs of linked VPC NAT rule for service access. format: ipv4
	LinkedVpcNatIps []string
	// Infromation related to the subnets where linked ENIs were created.
	LinkedVpcSubnets []LinkedSubnetInfo
	// The identifiers of route tables to be dynamically updated with SDDC networks
	RouteTableIds []string
	// service ARN role
	ServiceArnRole *string
	// Description of mapping between the traffic group name and corresponding Elastic Network Interface (ENI).
	TrafficGroupEniMappings []TrafficGroupEniMappingInfo
}

func (s *LinkedVpcInfo) GetType__() vapiBindings_.BindingType {
	return LinkedVpcInfoBindingType()
}

func (s *LinkedVpcInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for LinkedVpcInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Information related to linked VPC managed prefix list support including prefix list mode state, resource share information and managed prefix list information.
type LinkedVpcManagedPrefixListSupportInfo struct {
	AwsResourceShareInfo *AwsResourceShareInfo
	// Possible values are:
	//
	// * LinkedVpcManagedPrefixListSupportInfo#LinkedVpcManagedPrefixListSupportInfo_MANAGED_PREFIX_LIST_MODE_ENABLED
	// * LinkedVpcManagedPrefixListSupportInfo#LinkedVpcManagedPrefixListSupportInfo_MANAGED_PREFIX_LIST_MODE_DISABLED
	// * LinkedVpcManagedPrefixListSupportInfo#LinkedVpcManagedPrefixListSupportInfo_MANAGED_PREFIX_LIST_MODE_PENDING
	//
	//  Current state of managed prefix list mode
	ManagedPrefixListMode *string
	// Information of managed prefix list shared with current linked VPC
	ManagedPrefixLists []LinkedVpcSharedManagedPrefixListInfo
}

const LinkedVpcManagedPrefixListSupportInfo_MANAGED_PREFIX_LIST_MODE_ENABLED = "ENABLED"
const LinkedVpcManagedPrefixListSupportInfo_MANAGED_PREFIX_LIST_MODE_DISABLED = "DISABLED"
const LinkedVpcManagedPrefixListSupportInfo_MANAGED_PREFIX_LIST_MODE_PENDING = "PENDING"

func (s *LinkedVpcManagedPrefixListSupportInfo) GetType__() vapiBindings_.BindingType {
	return LinkedVpcManagedPrefixListSupportInfoBindingType()
}

func (s *LinkedVpcManagedPrefixListSupportInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for LinkedVpcManagedPrefixListSupportInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Information of managed prefix lists which are shared with linked VPC account.
type LinkedVpcSharedManagedPrefixListInfo struct {
	// Possible values are:
	//
	// * LinkedVpcSharedManagedPrefixListInfo#LinkedVpcSharedManagedPrefixListInfo_ADDRESS_FAMILY_IPV4
	// * LinkedVpcSharedManagedPrefixListInfo#LinkedVpcSharedManagedPrefixListInfo_ADDRESS_FAMILY_IPV6
	//
	//  Managed prefix list's address family
	AddressFamily *string
	// ID of managed prefix list.
	Id *string
	// Indicate that whether managed prefix list is currently in use or not.
	InUse *bool
	// Name of managed prefix list.
	Name            *string
	ProgrammingInfo *ManagedPrefixListProgrammingInfo
}

const LinkedVpcSharedManagedPrefixListInfo_ADDRESS_FAMILY_IPV4 = "IPv4"
const LinkedVpcSharedManagedPrefixListInfo_ADDRESS_FAMILY_IPV6 = "IPv6"

func (s *LinkedVpcSharedManagedPrefixListInfo) GetType__() vapiBindings_.BindingType {
	return LinkedVpcSharedManagedPrefixListInfoBindingType()
}

func (s *LinkedVpcSharedManagedPrefixListInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for LinkedVpcSharedManagedPrefixListInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Linked VPC list query result
type LinkedVpcsListResult struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	Self   *SelfResourceLink
	// Opaque cursor to be used for getting next page of records (supplied by current result page)
	Cursor *string
	// Count of results found (across all pages), set only on first page format: int64
	ResultCount *int64
	// If true, results are sorted in ascending order
	SortAscending *bool
	// Field by which records are sorted
	SortBy *string
	// Linked VPCs list
	Results []LinkedVpcInfo
}

func (s *LinkedVpcsListResult) GetType__() vapiBindings_.BindingType {
	return LinkedVpcsListResultBindingType()
}

func (s *LinkedVpcsListResult) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for LinkedVpcsListResult._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Base class for list results from collections
type ListResult struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	Self   *SelfResourceLink
	// Opaque cursor to be used for getting next page of records (supplied by current result page)
	Cursor *string
	// Count of results found (across all pages), set only on first page format: int64
	ResultCount *int64
	// If true, results are sorted in ascending order
	SortAscending *bool
	// Field by which records are sorted
	SortBy *string
}

func (s *ListResult) GetType__() vapiBindings_.BindingType {
	return ListResultBindingType()
}

func (s *ListResult) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ListResult._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type LogicalRouterPortCounters struct {
	// The total number of packets blocked. format: int64
	BlockedPackets *int64
	// Number of duplicate address detected packets dropped. format: int64
	DadDroppedPackets *int64
	// Number of packtes dropped as destination is not supported. format: int64
	DestinationUnsupportedDroppedPackets *int64
	// The total number of packets dropped. format: int64
	DroppedPackets *int64
	// Number of firewall packets dropped. format: int64
	FirewallDroppedPackets *int64
	// Number of fragmentation needed packets dropped. format: int64
	FragNeededDroppedPackets *int64
	// Number of IPSec packets dropped format: int64
	IpsecDroppedPackets *int64
	// Number of IPSec no security association packets dropped. format: int64
	IpsecNoSaDroppedPackets *int64
	// Number of IPSec packets dropped as no VTI is present. format: int64
	IpsecNoVtiDroppedPackets *int64
	// Number of IPSec policy block packets dropped. format: int64
	IpsecPolBlockDroppedPackets *int64
	// Number of IPSec policy error packets dropped. format: int64
	IpsecPolErrDroppedPackets *int64
	// Number of IPV6 packets dropped. format: int64
	Ipv6DroppedPackets *int64
	// Number of DPDK kernal NIC interface packets dropped. format: int64
	KniDroppedPackets *int64
	// Number of packets dropped due to unsupported L4 port. format: int64
	L4portUnsupportedDroppedPackets *int64
	// Number of packtes dropped as they are malformed. format: int64
	MalformedDroppedPackets *int64
	// Number of no ARP packets dropped. format: int64
	NoArpDroppedPackets *int64
	// Number of packets dropped as no linked ports are present. format: int64
	NoLinkedDroppedPackets *int64
	// Number of packets dropped due to insufficient memory. format: int64
	NoMemDroppedPackets *int64
	// Number of packets dropped due to absence of receiver. format: int64
	NoReceiverDroppedPackets *int64
	// The number of no route packets dropped format: int64
	NoRouteDroppedPackets *int64
	// Number of non IP packets dropped. format: int64
	NonIpDroppedPackets *int64
	// Number of packets dropped as protocol is unsupported. format: int64
	ProtoUnsupportedDroppedPackets *int64
	// Number of redirect packets dropped. format: int64
	RedirectDroppedPackets *int64
	// Number of reverse-path forwarding check packets dropped. format: int64
	RpfCheckDroppedPackets *int64
	// Number of service insert packets dropped. format: int64
	ServiceInsertDroppedPackets *int64
	// The total number of bytes transferred. format: int64
	TotalBytes *int64
	// The total number of packets transferred. format: int64
	TotalPackets *int64
	// Number of time to live exceeded packets dropped. format: int64
	TtlExceededDroppedPackets *int64
}

func (s *LogicalRouterPortCounters) GetType__() vapiBindings_.BindingType {
	return LogicalRouterPortCountersBindingType()
}

func (s *LogicalRouterPortCounters) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for LogicalRouterPortCounters._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type LogicalRouterPortStatisticsSummary struct {
	// Timestamp when the data was last updated; unset if data source has never updated the data. format: int64
	LastUpdateTimestamp *int64
	Rx                  *LogicalRouterPortCounters
	Tx                  *LogicalRouterPortCounters
	// The ID of the logical router port
	LogicalRouterPortId *string
}

func (s *LogicalRouterPortStatisticsSummary) GetType__() vapiBindings_.BindingType {
	return LogicalRouterPortStatisticsSummaryBindingType()
}

func (s *LogicalRouterPortStatisticsSummary) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for LogicalRouterPortStatisticsSummary._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Route programming related information of managed prefix lists which are shared with linked VPC account.
type ManagedPrefixListProgrammingInfo struct {
	// Name of active ENI that current managed prefix list should point to.
	ActiveEni *string
	// Prefixes to be present in AWS managed prefix list.
	ExpectedEntries []string
	// IDs of route tables that current managed prefix list is associated with.
	RouteTableIds []string
}

func (s *ManagedPrefixListProgrammingInfo) GetType__() vapiBindings_.BindingType {
	return ManagedPrefixListProgrammingInfoBindingType()
}

func (s *ManagedPrefixListProgrammingInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ManagedPrefixListProgrammingInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Attributes and state information of a particular managed prefix list.
type ManagedPrefixListRuntimeInfo struct {
	// Possible values are:
	//
	// * ManagedPrefixListRuntimeInfo#ManagedPrefixListRuntimeInfo_ADDRESS_FAMILY_IPV4
	// * ManagedPrefixListRuntimeInfo#ManagedPrefixListRuntimeInfo_ADDRESS_FAMILY_IPV6
	//
	//  Address family of managed prefix list
	AddressFamily *string
	// Amazon Resource Name (ARN) of managed prefix list
	Arn *string
	// Entries of managed prefix list
	Entries []string
	// Known issues found related to prefix list
	Issues []string
	// Maximum number of entries of managed prefix list format: int32
	MaxEntries *int64
	// Name of managed prefix list
	Name *string
	// Current state of managed prefix list
	State *string
	// State message of managed prefix list
	StateMessage *string
	// Tags of managed prefix list
	Tags []ManagedPrefixListTag
	// Current version of managed prefix list format: int64
	Version *int64
	// Unique identifier of current AWS resource.
	Id *string
	// Type identifier of current AWS resource.
	ResourceType string
}

// Identifier denoting this class, when it is used in polymorphic context.
//
// This value should be assigned to the property which is used to discriminate the actual type used in the polymorphic context.
const ManagedPrefixListRuntimeInfo__TYPE_IDENTIFIER = "ManagedPrefixListRuntimeInfo"
const ManagedPrefixListRuntimeInfo_ADDRESS_FAMILY_IPV4 = "IPv4"
const ManagedPrefixListRuntimeInfo_ADDRESS_FAMILY_IPV6 = "IPv6"

func (s *ManagedPrefixListRuntimeInfo) GetType__() vapiBindings_.BindingType {
	return ManagedPrefixListRuntimeInfoBindingType()
}

func (s *ManagedPrefixListRuntimeInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ManagedPrefixListRuntimeInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Tag that is attached with managed prefix list, as key-value pair.
type ManagedPrefixListTag struct {
	// Tag key
	Key *string
	// Tag value
	Value *string
}

func (s *ManagedPrefixListTag) GetType__() vapiBindings_.BindingType {
	return ManagedPrefixListTagBindingType()
}

func (s *ManagedPrefixListTag) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ManagedPrefixListTag._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Base type for resources that are managed by API clients
type ManagedResource struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	Self   *SelfResourceLink
	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected. format: int32
	Revision *int64
	// Timestamp of resource creation format: int64
	CreateTime *int64
	// ID of the user who created this resource
	CreateUser *string
	// Timestamp of last modification format: int64
	LastModifiedTime *int64
	// ID of the user who last modified this resource
	LastModifiedUser *string
	// Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super user and can modify it, but only when providing the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this entity.
	Protection *string
	// Indicates system owned resource
	SystemOwned *bool
	// Description of this resource
	Description *string
	// Defaults to ID if not set
	DisplayName *string
	// Unique identifier of this resource
	Id *string
	// The type of this resource.
	ResourceType *string
	// Opaque identifiers meaningful to the API user
	Tags []Tag
}

func (s *ManagedResource) GetType__() vapiBindings_.BindingType {
	return ManagedResourceBindingType()
}

func (s *ManagedResource) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ManagedResource._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// This will be used by VMC UI to return the Management Gateway (MGW) information as part of SddcNetworkingStateInfo API.
type ManagementGatewaySddcNetworkingStateInfo struct {
	// NSX VMs subnet format: ipv4-cidr-block
	ApplianceSubnet *string
	// MGW DNS Forwarder Zone Upstream Servers. format: ipv4-cidr-block
	Dns []string
	// SDDC infra subnet CIDRs. format: ipv4-cidr-block
	InfraSubnets []string
	// Management subnet CIDRs. format: ipv4-cidr-block
	MgmtSubnets []string
	// Total Firewall Rules format: int32
	TotalFirewallRules *int64
}

func (s *ManagementGatewaySddcNetworkingStateInfo) GetType__() vapiBindings_.BindingType {
	return ManagementGatewaySddcNetworkingStateInfoBindingType()
}

func (s *ManagementGatewaySddcNetworkingStateInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ManagementGatewaySddcNetworkingStateInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// A service entry describes the detail of a network service.
type MgmtServiceEntry struct {
	// Display name for this service
	DisplayName *string
	// Service path should refer to a valid service in the system. Service can be system defined or user defined.
	Path *string
}

func (s *MgmtServiceEntry) GetType__() vapiBindings_.BindingType {
	return MgmtServiceEntryBindingType()
}

func (s *MgmtServiceEntry) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for MgmtServiceEntry._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Management VM access information
type MgmtVmInfo struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	Self   *SelfResourceLink
	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected. format: int32
	Revision *int64
	// Timestamp of resource creation format: int64
	CreateTime *int64
	// ID of the user who created this resource
	CreateUser *string
	// Timestamp of last modification format: int64
	LastModifiedTime *int64
	// ID of the user who last modified this resource
	LastModifiedUser *string
	// Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super user and can modify it, but only when providing the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this entity.
	Protection *string
	// Indicates system owned resource
	SystemOwned *bool
	// Description of this resource
	Description *string
	// Management VM name
	DisplayName *string
	// Management VM identifier
	Id *string
	// The type of this resource.
	ResourceType *string
	// Opaque identifiers meaningful to the API user
	Tags []Tag
	// marked for delete identifier
	MarkedForDelete *bool
	// Path of its parent
	ParentPath *string
	// Absolute path of this object
	Path *string
	// Path relative from its parent
	RelativePath *string
	// For each management VM, a dedicated policy group will be created. This property will reflect its group path.
	GroupPath *string
	// IP address and attachment id pairs for tagging managment VM
	IpAttachmentPairs []IpAttachmentPair
	// Local IPs of a management VM format: address-or-block-or-range
	Ips []string
	// Details services path and display name.
	Services []MgmtServiceEntry
}

func (s *MgmtVmInfo) GetType__() vapiBindings_.BindingType {
	return MgmtVmInfoBindingType()
}

func (s *MgmtVmInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for MgmtVmInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Management VM list query result
type MgmtVmsListResult struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	Self   *SelfResourceLink
	// Opaque cursor to be used for getting next page of records (supplied by current result page)
	Cursor *string
	// Count of results found (across all pages), set only on first page format: int64
	ResultCount *int64
	// If true, results are sorted in ascending order
	SortAscending *bool
	// Field by which records are sorted
	SortBy *string
	// Management VMs list
	Results []MgmtVmInfo
}

func (s *MgmtVmsListResult) GetType__() vapiBindings_.BindingType {
	return MgmtVmsListResultBindingType()
}

func (s *MgmtVmsListResult) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for MgmtVmsListResult._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Interface information (Label)
type ModelInterface struct {
	// Identifier of the Interface label
	Id *string
	// Name of the Interface label
	Name *string
}

func (s *ModelInterface) GetType__() vapiBindings_.BindingType {
	return ModelInterfaceBindingType()
}

func (s *ModelInterface) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ModelInterface._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Other region TGW connection info
type OtherRegionConnectionInfo struct {
	// AWS region
	Region *string
	// TGW ID
	TgwId *string
	// TGW peering id
	TgwPeeringId *string
}

func (s *OtherRegionConnectionInfo) GetType__() vapiBindings_.BindingType {
	return OtherRegionConnectionInfoBindingType()
}

func (s *OtherRegionConnectionInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for OtherRegionConnectionInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Outpost related information.
type OutpostInfo struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	Self   *SelfResourceLink
	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected. format: int32
	Revision *int64
	// Timestamp of resource creation format: int64
	CreateTime *int64
	// ID of the user who created this resource
	CreateUser *string
	// Timestamp of last modification format: int64
	LastModifiedTime *int64
	// ID of the user who last modified this resource
	LastModifiedUser *string
	// Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super user and can modify it, but only when providing the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this entity.
	Protection *string
	// Indicates system owned resource
	SystemOwned *bool
	// Description of this resource
	Description *string
	// Defaults to ID if not set
	DisplayName *string
	// Unique identifier of this resource
	Id *string
	// The type of this resource.
	ResourceType *string
	// Opaque identifiers meaningful to the API user
	Tags []Tag
	// marked for delete identifier
	MarkedForDelete *bool
	// Path of its parent
	ParentPath *string
	// Absolute path of this object
	Path *string
	// Path relative from its parent
	RelativePath *string
	// Local gateway ID
	LgwId *string
	// Local gateway route table ID
	LgwRouteTableId *string
	// Local gateway route table VPC association ID
	LgwRouteTableVpcAssociationId *string
	// State of local gateway route table VPC association
	LgwRouteTableVpcAssociationState *string
}

func (s *OutpostInfo) GetType__() vapiBindings_.BindingType {
	return OutpostInfoBindingType()
}

func (s *OutpostInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for OutpostInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// This will be used by VMC UI to return the Outposts Connect information as part of SddcNetworkingStateInfo API.
type OutpostsConnectSddcNetworkingStateInfo struct {
	// The total number of Outposts Connect advertised routes. format: int32
	AdvertisedRoutes *int64
	// The total number of Outposts Connect learned routes. format: int32
	LearnedRoutes *int64
}

func (s *OutpostsConnectSddcNetworkingStateInfo) GetType__() vapiBindings_.BindingType {
	return OutpostsConnectSddcNetworkingStateInfoBindingType()
}

func (s *OutpostsConnectSddcNetworkingStateInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for OutpostsConnectSddcNetworkingStateInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Tier0 or Tier1 interface statistics on specific Enforcement Point.
type PolicyInterfaceStatisticsSummary struct {
	// Timestamp when the data was last updated; unset if data source has never updated the data. format: int64
	LastUpdateTimestamp *int64
	Rx                  *LogicalRouterPortCounters
	Tx                  *LogicalRouterPortCounters
	// The ID of the logical router port
	LogicalRouterPortId *string
	// Policy path for the interface
	InterfacePolicyPath *string
}

func (s *PolicyInterfaceStatisticsSummary) GetType__() vapiBindings_.BindingType {
	return PolicyInterfaceStatisticsSummaryBindingType()
}

func (s *PolicyInterfaceStatisticsSummary) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for PolicyInterfaceStatisticsSummary._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Prefix lists used by certain features, like Intranet NAT.
type PrefixListInfo struct {
	// Prefix List name
	Name *string
	// Relative Prefix List path
	Path *string
	// Prefix List URL
	Url *string
}

func (s *PrefixListInfo) GetType__() vapiBindings_.BindingType {
	return PrefixListInfoBindingType()
}

func (s *PrefixListInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for PrefixListInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// A list for provider gateways.
type ProviderGatewayKeyValuePairs struct {
	// Key
	Key *string
	// Value
	Value []ProviderObject
}

func (s *ProviderGatewayKeyValuePairs) GetType__() vapiBindings_.BindingType {
	return ProviderGatewayKeyValuePairsBindingType()
}

func (s *ProviderGatewayKeyValuePairs) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ProviderGatewayKeyValuePairs._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Basic properties of SDDC User Config.
type ProviderObject struct {
	// Display name
	DisplayName *string
	// ID
	Id *string
	// Path
	Path *string
	// Optional field, used to identify the object.
	Type_ *string
	// Should start with \"/policy/api/v1/\".
	Url *string
}

func (s *ProviderObject) GetType__() vapiBindings_.BindingType {
	return ProviderObjectBindingType()
}

func (s *ProviderObject) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ProviderObject._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type PublicIp struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	Self   *SelfResourceLink
	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected. format: int32
	Revision *int64
	// Timestamp of resource creation format: int64
	CreateTime *int64
	// ID of the user who created this resource
	CreateUser *string
	// Timestamp of last modification format: int64
	LastModifiedTime *int64
	// ID of the user who last modified this resource
	LastModifiedUser *string
	// Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super user and can modify it, but only when providing the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this entity.
	Protection *string
	// Indicates system owned resource
	SystemOwned *bool
	// Description of this resource
	Description *string
	DisplayName *string
	// Public IP identifier
	Id *string
	// The type of this resource.
	ResourceType *string
	// Opaque identifiers meaningful to the API user
	Tags []Tag
	// marked for delete identifier
	MarkedForDelete *bool
	// Path of its parent
	ParentPath *string
	// Absolute path of this object
	Path *string
	// Path relative from its parent
	RelativePath *string
	// IPv4 address format: ipv4
	Ip *string
}

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

// Public IP list
type PublicIpsListResult struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	Self   *SelfResourceLink
	// Opaque cursor to be used for getting next page of records (supplied by current result page)
	Cursor *string
	// Count of results found (across all pages), set only on first page format: int64
	ResultCount *int64
	// If true, results are sorted in ascending order
	SortAscending *bool
	// Field by which records are sorted
	SortBy *string
	// Public IP list
	Results []PublicIp
}

func (s *PublicIpsListResult) GetType__() vapiBindings_.BindingType {
	return PublicIpsListResultBindingType()
}

func (s *PublicIpsListResult) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for PublicIpsListResult._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Detailed information about a related API error
type RelatedApiError struct {
	// Further details about the error
	Details *string
	// A numeric error code format: int64
	ErrorCode *int64
	// Additional data about the error
	ErrorData *vapiData_.StructValue
	// A description of the error
	ErrorMessage *string
	// The module name where the error occurred
	ModuleName *string
}

func (s *RelatedApiError) GetType__() vapiBindings_.BindingType {
	return RelatedApiErrorBindingType()
}

func (s *RelatedApiError) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for RelatedApiError._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// This model contains CIDR blocks which are reserved for second party services like WCP and HCX. Default CGW segments can not overlap with any reserved CIDR blocks.
type ReservedCIDRBlock struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	Self   *SelfResourceLink
	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected. format: int32
	Revision *int64
	// Timestamp of resource creation format: int64
	CreateTime *int64
	// ID of the user who created this resource
	CreateUser *string
	// Timestamp of last modification format: int64
	LastModifiedTime *int64
	// ID of the user who last modified this resource
	LastModifiedUser *string
	// Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super user and can modify it, but only when providing the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this entity.
	Protection *string
	// Indicates system owned resource
	SystemOwned *bool
	// Description of this resource
	Description *string
	// Defaults to ID if not set
	DisplayName *string
	// Unique identifier of this resource
	Id *string
	// The type of this resource.
	ResourceType *string
	// Opaque identifiers meaningful to the API user
	Tags []Tag
	// marked for delete identifier
	MarkedForDelete *bool
	// Path of its parent
	ParentPath *string
	// Absolute path of this object
	Path *string
	// Path relative from its parent
	RelativePath *string
	// Single CIDR block of the object, will be deprecated soon, please use prefixes going forward. format: ipv4-cidr-block
	//
	// Deprecated: This API element is deprecated.
	Cidr *string
	// List of CIDR blocks
	Prefixes []ReservedCIDRBlockPrefix
}

func (s *ReservedCIDRBlock) GetType__() vapiBindings_.BindingType {
	return ReservedCIDRBlockBindingType()
}

func (s *ReservedCIDRBlock) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ReservedCIDRBlock._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// This model contains CIDR block prefixes which are reserved for second party services like WCP and HCX. Default CGW segments can not overlap with any reserved CIDR blocks.
type ReservedCIDRBlockPrefix struct {
	// Description about the prefix to be reserved
	Description *string
	// Prefix to be reserved format: ipv4-cidr-block
	Prefix *string
}

func (s *ReservedCIDRBlockPrefix) GetType__() vapiBindings_.BindingType {
	return ReservedCIDRBlockPrefixBindingType()
}

func (s *ReservedCIDRBlockPrefix) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ReservedCIDRBlockPrefix._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Reserved CIDR blocks list
type ReservedCIDRBlocksListResult struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	Self   *SelfResourceLink
	// Opaque cursor to be used for getting next page of records (supplied by current result page)
	Cursor *string
	// Count of results found (across all pages), set only on first page format: int64
	ResultCount *int64
	// If true, results are sorted in ascending order
	SortAscending *bool
	// Field by which records are sorted
	SortBy *string
	// Reserved CIDR blocks list
	Results []ReservedCIDRBlock
}

func (s *ReservedCIDRBlocksListResult) GetType__() vapiBindings_.BindingType {
	return ReservedCIDRBlocksListResultBindingType()
}

func (s *ReservedCIDRBlocksListResult) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ReservedCIDRBlocksListResult._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Base class for resources
type Resource struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	Self   *SelfResourceLink
}

func (s *Resource) GetType__() vapiBindings_.BindingType {
	return ResourceBindingType()
}

func (s *Resource) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for Resource._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// A link to a related resource
type ResourceLink struct {
	// Optional action
	Action *string
	// Link to resource
	Href *string
	// Custom relation type (follows RFC 5988 where appropriate definitions exist)
	Rel *string
}

func (s *ResourceLink) GetType__() vapiBindings_.BindingType {
	return ResourceLinkBindingType()
}

func (s *ResourceLink) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ResourceLink._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Base abstract runtime status information for a AWS resource.
type ResourceRuntimeInfo struct {
	// Unique identifier of current AWS resource.
	Id *string
	// Type identifier of current AWS resource.
	ResourceType string
}

// Identifier denoting this class, when it is used in polymorphic context.
//
// This value should be assigned to the property which is used to discriminate the actual type used in the polymorphic context.
const ResourceRuntimeInfo__TYPE_IDENTIFIER = "ResourceRuntimeInfo"

func (s *ResourceRuntimeInfo) GetType__() vapiBindings_.BindingType {
	return ResourceRuntimeInfoBindingType()
}

func (s *ResourceRuntimeInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ResourceRuntimeInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// List result of AWS managed resource(s) runtime information.
type ResourceRuntimeInfoListResult struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	Self   *SelfResourceLink
	// Opaque cursor to be used for getting next page of records (supplied by current result page)
	Cursor *string
	// Count of results found (across all pages), set only on first page format: int64
	ResultCount *int64
	// If true, results are sorted in ascending order
	SortAscending *bool
	// Field by which records are sorted
	SortBy *string
	// AWS managed resources runtime information list
	Results []*vapiData_.StructValue
}

func (s *ResourceRuntimeInfoListResult) GetType__() vapiBindings_.BindingType {
	return ResourceRuntimeInfoListResultBindingType()
}

func (s *ResourceRuntimeInfoListResult) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ResourceRuntimeInfoListResult._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// A base class for types that track revisions
type RevisionedResource struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	Self   *SelfResourceLink
	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected. format: int32
	Revision *int64
}

func (s *RevisionedResource) GetType__() vapiBindings_.BindingType {
	return RevisionedResourceBindingType()
}

func (s *RevisionedResource) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for RevisionedResource._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Route Aggregation object information.
type RouteAggregationConfig struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	Self   *SelfResourceLink
	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected. format: int32
	Revision *int64
	// Timestamp of resource creation format: int64
	CreateTime *int64
	// ID of the user who created this resource
	CreateUser *string
	// Timestamp of last modification format: int64
	LastModifiedTime *int64
	// ID of the user who last modified this resource
	LastModifiedUser *string
	// Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super user and can modify it, but only when providing the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this entity.
	Protection *string
	// Indicates system owned resource
	SystemOwned *bool
	// Description of this resource
	Description *string
	// Route Aggregation object name
	DisplayName *string
	// Unique identifier of this resource
	Id *string
	// The type of this resource.
	ResourceType *string
	// Opaque identifiers meaningful to the API user
	Tags []Tag
	// marked for delete identifier
	MarkedForDelete *bool
	// Path of its parent
	ParentPath *string
	// Absolute path of this object
	Path *string
	// Path relative from its parent
	RelativePath *string
	// Possible values are:
	//
	// * RouteAggregationConfig#RouteAggregationConfig_ADDRESS_FAMILY_IPV4
	// * RouteAggregationConfig#RouteAggregationConfig_ADDRESS_FAMILY_IPV6
	//
	//  IP address type
	AddressFamily *string
	// List of prefix that user wants to aggregate with the VMC segments
	Prefixes []RouteAggregationConfigPrefix
}

const RouteAggregationConfig_ADDRESS_FAMILY_IPV4 = "IPv4"
const RouteAggregationConfig_ADDRESS_FAMILY_IPV6 = "IPv6"

func (s *RouteAggregationConfig) GetType__() vapiBindings_.BindingType {
	return RouteAggregationConfigBindingType()
}

func (s *RouteAggregationConfig) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for RouteAggregationConfig._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Route aggregation object list
type RouteAggregationConfigListResult struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	Self   *SelfResourceLink
	// Opaque cursor to be used for getting next page of records (supplied by current result page)
	Cursor *string
	// Count of results found (across all pages), set only on first page format: int64
	ResultCount *int64
	// If true, results are sorted in ascending order
	SortAscending *bool
	// Field by which records are sorted
	SortBy *string
	// Route aggregation prefixes
	Results []RouteAggregationConfig
}

func (s *RouteAggregationConfigListResult) GetType__() vapiBindings_.BindingType {
	return RouteAggregationConfigListResultBindingType()
}

func (s *RouteAggregationConfigListResult) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for RouteAggregationConfigListResult._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Prefix object that users provide to aggregate.
type RouteAggregationConfigPrefix struct {
	// Description about the prefix that has to be aggregated
	Description *string
	// Prefix that belong to users workload VM's format: ip-cidr-block
	Prefix *string
	// Flag to identify whether to advertise this prefix or not
	SummaryOnly *bool
}

func (s *RouteAggregationConfigPrefix) GetType__() vapiBindings_.BindingType {
	return RouteAggregationConfigPrefixBindingType()
}

func (s *RouteAggregationConfigPrefix) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for RouteAggregationConfigPrefix._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// RouteConfig object
type RouteConfig struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	Self   *SelfResourceLink
	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected. format: int32
	Revision *int64
	// Timestamp of resource creation format: int64
	CreateTime *int64
	// ID of the user who created this resource
	CreateUser *string
	// Timestamp of last modification format: int64
	LastModifiedTime *int64
	// ID of the user who last modified this resource
	LastModifiedUser *string
	// Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super user and can modify it, but only when providing the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this entity.
	Protection *string
	// Indicates system owned resource
	SystemOwned *bool
	// Description of this resource
	Description *string
	// Name of the routeConfig object.
	DisplayName *string
	// Unique identifier of this resource
	Id *string
	// The type of this resource.
	ResourceType *string
	// Opaque identifiers meaningful to the API user
	Tags []Tag
	// marked for delete identifier
	MarkedForDelete *bool
	// Path of its parent
	ParentPath *string
	// Absolute path of this object
	Path *string
	// Path relative from its parent
	RelativePath           *string
	AggregationRouteConfig *AggregationRouteConfigInfo
	// Scope where routeAggregationConfig has to be applied.Existing connectivity endpoint object path to be given as valid values
	ConnectivityEndpointPath *string
	// This flag will be true if filtering is enabled. Similarly, the flag will set to false once filtering gets disabled.
	DisableDefaultCgwNetwork *bool
}

func (s *RouteConfig) GetType__() vapiBindings_.BindingType {
	return RouteConfigBindingType()
}

func (s *RouteConfig) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for RouteConfig._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// List of routeConfig object
type RouteConfigListResult struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	Self   *SelfResourceLink
	// Opaque cursor to be used for getting next page of records (supplied by current result page)
	Cursor *string
	// Count of results found (across all pages), set only on first page format: int64
	ResultCount *int64
	// If true, results are sorted in ascending order
	SortAscending *bool
	// Field by which records are sorted
	SortBy *string
	// Object to hold all the routeConfig objects
	Results []RouteConfig
}

func (s *RouteConfigListResult) GetType__() vapiBindings_.BindingType {
	return RouteConfigListResultBindingType()
}

func (s *RouteConfigListResult) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for RouteConfigListResult._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Filtering object that users provide to filter default cgw segments.
type RouteFiltering struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	Self   *SelfResourceLink
	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected. format: int32
	Revision *int64
	// Timestamp of resource creation format: int64
	CreateTime *int64
	// ID of the user who created this resource
	CreateUser *string
	// Timestamp of last modification format: int64
	LastModifiedTime *int64
	// ID of the user who last modified this resource
	LastModifiedUser *string
	// Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super user and can modify it, but only when providing the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this entity.
	Protection *string
	// Indicates system owned resource
	SystemOwned *bool
	// Description of this resource
	Description *string
	// Defaults to ID if not set
	DisplayName *string
	// Unique identifier of this resource
	Id *string
	// The type of this resource.
	ResourceType *string
	// Opaque identifiers meaningful to the API user
	Tags []Tag
	// marked for delete identifier
	MarkedForDelete *bool
	// Path of its parent
	ParentPath *string
	// Absolute path of this object
	Path *string
	// Path relative from its parent
	RelativePath *string
	// Possible values are:
	//
	// * RouteFiltering#RouteFiltering_ENDPOINTS_INTRANET
	// * RouteFiltering#RouteFiltering_ENDPOINTS_SERVICES
	//
	//  Connectivity endpoints for which user wants to filter the default CGW segments. Valid values INTRANET, SERVICES.
	Endpoints []string
	// Possible values are:
	//
	// * RouteFiltering#RouteFiltering_STATE_ENABLED
	// * RouteFiltering#RouteFiltering_STATE_DISABLED
	//
	//  State of filtering enabled/disabled on a list of endpoints.
	State *string
}

const RouteFiltering_ENDPOINTS_INTRANET = "INTRANET"
const RouteFiltering_ENDPOINTS_SERVICES = "SERVICES"
const RouteFiltering_STATE_ENABLED = "ENABLED"
const RouteFiltering_STATE_DISABLED = "DISABLED"

func (s *RouteFiltering) GetType__() vapiBindings_.BindingType {
	return RouteFilteringBindingType()
}

func (s *RouteFiltering) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for RouteFiltering._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// SDDC advertise route configuration with AWS prefix list and other region TGW connections.
type SddcAdvertiseRouteConfig struct {
	// AWS prefix list for managing TGW routes
	AwsPrefixList []AwsPrefixInfo
	// List of TGW connection information from other regions to local region TGW
	OtherRegionConnections []OtherRegionConnectionInfo
}

func (s *SddcAdvertiseRouteConfig) GetType__() vapiBindings_.BindingType {
	return SddcAdvertiseRouteConfigBindingType()
}

func (s *SddcAdvertiseRouteConfig) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for SddcAdvertiseRouteConfig._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// This will be used by VMC UI to return the network information such as VPN connectivity details, DX connectivity details, Transit Connect details, Management Gateway (MGW), Compute Gateway (default CGW) and Cloud Provider details. This is for read only UI.
type SddcNetworkingStateInfo struct {
	// Cloud Provider
	CloudProvider []CloudProviderSddcNetworkingStateInfo
	// Compute Gateway
	ComputeGateway []ComputeGatewaySddcNetworkingStateInfo
	// Direct Connect
	DirectConnect []DirectConnectSddcNetworkingStateInfo
	// Error Information
	Error_ *string
	// Last updated time
	LastSyncTime *string
	// Management Gateway
	ManagementGateway []ManagementGatewaySddcNetworkingStateInfo
	// Outposts Connect
	OutpostsConnect []OutpostsConnectSddcNetworkingStateInfo
	// Service provider name.
	ProviderName *string
	// SDDC Identifier
	SddcId *string
	// The name of SDDC.
	SddcName *string
	// Region of the SDDC deployment.
	SddcRegion *string
	// Transit Connect
	TransitConnect []TransitConnectSddcNetworkingStateInfo
	// VPN Information
	Vpn []VpnSddcNetworkingStateInfo
}

func (s *SddcNetworkingStateInfo) GetType__() vapiBindings_.BindingType {
	return SddcNetworkingStateInfoBindingType()
}

func (s *SddcNetworkingStateInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for SddcNetworkingStateInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Transport Node details for ESO edge vm deployment.
type SddcTransportNodeInfo struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	Self   *SelfResourceLink
	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected. format: int32
	Revision *int64
	// Timestamp of resource creation format: int64
	CreateTime *int64
	// ID of the user who created this resource
	CreateUser *string
	// Timestamp of last modification format: int64
	LastModifiedTime *int64
	// ID of the user who last modified this resource
	LastModifiedUser *string
	// Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super user and can modify it, but only when providing the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this entity.
	Protection *string
	// Indicates system owned resource
	SystemOwned *bool
	// Description of this resource
	Description *string
	// Defaults to ID if not set
	DisplayName *string
	// Unique identifier of this resource
	Id *string
	// The type of this resource.
	ResourceType *string
	// Opaque identifiers meaningful to the API user
	Tags []Tag
	// marked for delete identifier
	MarkedForDelete *bool
	// Path of its parent
	ParentPath *string
	// Absolute path of this object
	Path *string
	// Path relative from its parent
	RelativePath *string
	// Possible values are:
	//
	// * SddcTransportNodeInfo#SddcTransportNodeInfo_APPLIANCE_CONFIGURATION_LARGE
	// * SddcTransportNodeInfo#SddcTransportNodeInfo_APPLIANCE_CONFIGURATION_MEDIUM
	//
	//  Form factor of the compute server.
	ApplianceConfiguration *string
	// Possible values are:
	//
	// * SddcTransportNodeInfo#SddcTransportNodeInfo_HARDWARE_PLATFORM_TYPE_I3
	// * SddcTransportNodeInfo#SddcTransportNodeInfo_HARDWARE_PLATFORM_TYPE_I3EN
	// * SddcTransportNodeInfo#SddcTransportNodeInfo_HARDWARE_PLATFORM_TYPE_NA
	//
	//  Hardware platform type of the compute server.
	HardwarePlatformType *string
	// No of queues to be set on the scaled out edge nodes. format: int64
	QueueNumPerPort *int64
	ReservationInfo *TransportNodeReservationInfo
}

const SddcTransportNodeInfo_APPLIANCE_CONFIGURATION_LARGE = "large"
const SddcTransportNodeInfo_APPLIANCE_CONFIGURATION_MEDIUM = "medium"
const SddcTransportNodeInfo_HARDWARE_PLATFORM_TYPE_I3 = "i3"
const SddcTransportNodeInfo_HARDWARE_PLATFORM_TYPE_I3EN = "i3en"
const SddcTransportNodeInfo_HARDWARE_PLATFORM_TYPE_NA = "NA"

func (s *SddcTransportNodeInfo) GetType__() vapiBindings_.BindingType {
	return SddcTransportNodeInfoBindingType()
}

func (s *SddcTransportNodeInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for SddcTransportNodeInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// SDDC configuration parameters for users. User-level addresses/CIDRs are provided.
type SddcUserConfiguration struct {
	// All uplink interfaces label name. Only presents in AWS environment now. Deprecated, please use labels.
	//
	// Deprecated: This API element is deprecated.
	AllUplinkInterfaceLabel *string
	// All VPN interfaces label name. Only presents in AWS environment now. Deprecated, please use labels.
	//
	// Deprecated: This API element is deprecated.
	AllVpnInterfaceLabel *string
	// Compute gateway SNAT IP address. format: ipv4
	CgwSnatIp *string
	// Compute domain ID. Only presents in AWS environment now. Deprecated, please use domains.
	//
	// Deprecated: This API element is deprecated.
	ComputeDomain *string
	// Compute gateway name. Only presents in AWS environment now. Deprecated, please use gateways.
	//
	// Deprecated: This API element is deprecated.
	ComputeGateway *string
	// Domain information for management gateway and compute gateway.
	Domains []ProviderObject
	// DirectConnect interface label name. Only presents in AWS environment now. Deprecated, please use labels.
	//
	// Deprecated: This API element is deprecated.
	DxInterfaceLabel *string
	// Policy enforcement point object paths.
	EnforcementPoints []ProviderObject
	// List of applicable SDDC connectivities to external endpoints.
	ExternalConnectivities []ExternalConnectivityInfo
	// Provider gateway list. Including both tier-0 gateways and tier-1 gateways.
	Gateways []ProviderGatewayKeyValuePairs
	// SDDC infra subnet CIDRs. format: ipv4-cidr-block
	InfraSubnets []string
	// Interfaces (labels) including public interface, direct connect interface, linked vpc interface, etc. Only presents in AWS environment now. Deprecated, please use labels.
	//
	// Deprecated: This API element is deprecated.
	Interfaces []ModelInterface
	// Interfaces (labels) including internet, intranet and services.
	Labels []ProviderObject
	// Linked VPC interface label name. Only presents in AWS environment now. Deprecated, please use labels.
	//
	// Deprecated: This API element is deprecated.
	LinkedVpcInterfaceLabel *string
	// Management domain ID. Only presents in AWS environment now. Deprecated, please use domains.
	//
	// Deprecated: This API element is deprecated.
	ManagementDomain *string
	// Management gateway name. Only presents in AWS environment now. Deprecated, please use gateways.
	//
	// Deprecated: This API element is deprecated.
	ManagementGateway               *string
	ManagementGatewayDefaultDnsZone *ProviderObject
	// Management gateway label name. Only presents in AWS environment now. Deprecated, please use labels.
	//
	// Deprecated: This API element is deprecated.
	ManagementGatewayLabel *string
	// Management subnet CIDRs. Only presents in AWS environment now. Deprecated, please use mgmt_subnets. format: ipv4-cidr-block
	//
	// Deprecated: This API element is deprecated.
	MgmtSubnet []string
	// Management subnet CIDRs. format: ipv4-cidr-block
	MgmtSubnets []string
	// Management gateway SNAT IP address. format: ipv4
	MgwSnatIp *string
	// Provider gateway list. Including both tier-0 gateways and tier-1 gateways. Only presents in AWS environment now. Deprecated, please use gateways.
	//
	// Deprecated: This API element is deprecated.
	ProviderGateways []ProviderGatewayKeyValuePairs
	// Service provider name. Only presents in AWS environment now. Deprecated, please use gateways.
	//
	// Deprecated: This API element is deprecated.
	ProviderName *string
	// Prefix lists used by certain features, like Intranet NAT.
	ProviderPrefixLists []PrefixListInfo
	// Public interface label name. Only presents in AWS environment now. Deprecated, please use labels.
	//
	// Deprecated: This API element is deprecated.
	PublicInterfaceLabel *string
	// Flag to indicate if route aggregation is supported on a provider environment. Upstream consumers like HCX, WCP uses this attribute to find out if route aggregation APIs can be used on a provider environment.
	RouteAggregationEnabled *bool
	// SDDC infra subnet CIDRs. Only presents in AWS environment now. Deprecated, please use infra_subnets. format: ipv4-cidr-block
	//
	// Deprecated: This API element is deprecated.
	SddcInfraSubnet []string
	// The name of SDDC.
	SddcName *string
	// Region of the SDDC deployment.
	SddcRegion *string
	// Possible values are:
	//
	// * SddcUserConfiguration#SddcUserConfiguration_SECURITY_CLASSIFICATION_TOP_SECRET
	// * SddcUserConfiguration#SddcUserConfiguration_SECURITY_CLASSIFICATION_SECRET
	// * SddcUserConfiguration#SddcUserConfiguration_SECURITY_CLASSIFICATION_TOP_SECRET_SCI
	// * SddcUserConfiguration#SddcUserConfiguration_SECURITY_CLASSIFICATION_UNCLASSIFIED
	// * SddcUserConfiguration#SddcUserConfiguration_SECURITY_CLASSIFICATION_CONFIDENTIAL
	// * SddcUserConfiguration#SddcUserConfiguration_SECURITY_CLASSIFICATION_UNKNOWN
	//
	//  Security classification
	SecurityClassification *string
	// Local IPs for VPN tunnel over Direct Connect. Only presents in AWS environment now. Deprecated, please use vpn_endpoints instead of vpn_dx_ips. format: ipv4
	//
	// Deprecated: This API element is deprecated.
	VpnDxIps []string
	// VPN tunnel endpoints. Currently containing public IPs for VPN over internet and local IPs for VPN over intranet.
	VpnEndpoints []VpnEndpoint
	// Public IPs for VPN tunnel over internet. Only presents in AWS environment now. Deprecated, please use vpn_endpoints instead of vpn_internet_ips. format: ipv4
	//
	// Deprecated: This API element is deprecated.
	VpnInternetIps []string
}

const SddcUserConfiguration_SECURITY_CLASSIFICATION_TOP_SECRET = "TOP_SECRET"
const SddcUserConfiguration_SECURITY_CLASSIFICATION_SECRET = "SECRET"
const SddcUserConfiguration_SECURITY_CLASSIFICATION_TOP_SECRET_SCI = "TOP_SECRET_SCI"
const SddcUserConfiguration_SECURITY_CLASSIFICATION_UNCLASSIFIED = "UNCLASSIFIED"
const SddcUserConfiguration_SECURITY_CLASSIFICATION_CONFIDENTIAL = "CONFIDENTIAL"
const SddcUserConfiguration_SECURITY_CLASSIFICATION_UNKNOWN = "UNKNOWN"

func (s *SddcUserConfiguration) GetType__() vapiBindings_.BindingType {
	return SddcUserConfigurationBindingType()
}

func (s *SddcUserConfiguration) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for SddcUserConfiguration._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// The server will populate this field when returing the resource. Ignored on PUT and POST.
type SelfResourceLink struct {
	// Optional action
	Action *string
	// Link to resource
	Href *string
	// Custom relation type (follows RFC 5988 where appropriate definitions exist)
	Rel *string
}

func (s *SelfResourceLink) GetType__() vapiBindings_.BindingType {
	return SelfResourceLinkBindingType()
}

func (s *SelfResourceLink) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for SelfResourceLink._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Arbitrary key-value pairs that may be attached to an entity
type Tag struct {
	// Tag searches may optionally be restricted by scope
	Scope *string
	// Identifier meaningful to user with maximum length of 256 characters
	Tag *string
}

func (s *Tag) GetType__() vapiBindings_.BindingType {
	return TagBindingType()
}

func (s *Tag) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for Tag._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Traffic group configuration. A traffic group indicates dedicated bandwidth and computation for a given list of subnets. Creating a traffic group will reserve resources and associating it with desired prefix lists will use the resources for the traffic of the prefix lists. Besides traffic group ID and name, realized state and detailed association maps info of this traffic group are included if verbose info is requested.
type TrafficGroup struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	Self   *SelfResourceLink
	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected. format: int32
	Revision *int64
	// Timestamp of resource creation format: int64
	CreateTime *int64
	// ID of the user who created this resource
	CreateUser *string
	// Timestamp of last modification format: int64
	LastModifiedTime *int64
	// ID of the user who last modified this resource
	LastModifiedUser *string
	// Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super user and can modify it, but only when providing the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this entity.
	Protection *string
	// Indicates system owned resource
	SystemOwned *bool
	// Description of this resource
	Description *string
	// A description field for the traffic group.
	DisplayName *string
	// Unique identifier for a traffic group.
	Id *string
	// The type of this resource.
	ResourceType *string
	// Opaque identifiers meaningful to the API user
	Tags []Tag
	// marked for delete identifier
	MarkedForDelete *bool
	// Path of its parent
	ParentPath *string
	// Absolute path of this object
	Path *string
	// Path relative from its parent
	RelativePath *string
	// Association maps in the traffic group.
	AssociationMaps []TrafficGroupAssociationMap
	// Unique resource identifier for a traffic group. This identifier is used in the traffic group's related resource allocation and naming. format: int32
	ResourceId *int64
	// Possible values are:
	//
	// * TrafficGroup#TrafficGroup_STATE_IN_PROGRESS
	// * TrafficGroup#TrafficGroup_STATE_SUCCESS
	// * TrafficGroup#TrafficGroup_STATE_FAILURE
	// * TrafficGroup#TrafficGroup_STATE_UNAVAILABLE
	// * TrafficGroup#TrafficGroup_STATE_PENDING
	//
	//  Realized state of the traffic group. This matches the realized state (VmcConsolidatedStatus) of the traffic group.
	State *string
	// Information on the current state. Mostly error messages on failure states.
	StateMessage *string
	// Possible values are:
	//
	// * TrafficGroup#TrafficGroup_TYPE_STATIC_HIGHAVAILABILITY
	// * TrafficGroup#TrafficGroup_TYPE_STATIC_HIGHTHROUGHPUT
	// * TrafficGroup#TrafficGroup_TYPE_DYNAMIC_HIGHTHROUGHPUT
	//
	//  Type of traffic group
	Type_ *string
}

const TrafficGroup_STATE_IN_PROGRESS = "IN_PROGRESS"
const TrafficGroup_STATE_SUCCESS = "SUCCESS"
const TrafficGroup_STATE_FAILURE = "FAILURE"
const TrafficGroup_STATE_UNAVAILABLE = "UNAVAILABLE"
const TrafficGroup_STATE_PENDING = "PENDING"
const TrafficGroup_TYPE_STATIC_HIGHAVAILABILITY = "STATIC_HIGHAVAILABILITY"
const TrafficGroup_TYPE_STATIC_HIGHTHROUGHPUT = "STATIC_HIGHTHROUGHPUT"
const TrafficGroup_TYPE_DYNAMIC_HIGHTHROUGHPUT = "DYNAMIC_HIGHTHROUGHPUT"

func (s *TrafficGroup) GetType__() vapiBindings_.BindingType {
	return TrafficGroupBindingType()
}

func (s *TrafficGroup) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for TrafficGroup._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// An association map of a traffic group. It describes the association of prefix lists and logical routers with the traffic group. To make use of dedicated traffic resources through traffic groups, prefix lists need to be linked to the desired traffic group and an association with the target logical router (scope) is required. The scope is for declaring which logical router the prefix lists are associated under a particular traffic group. The scope is not required input.
type TrafficGroupAssociationMap struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	Self   *SelfResourceLink
	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected. format: int32
	Revision *int64
	// Timestamp of resource creation format: int64
	CreateTime *int64
	// ID of the user who created this resource
	CreateUser *string
	// Timestamp of last modification format: int64
	LastModifiedTime *int64
	// ID of the user who last modified this resource
	LastModifiedUser *string
	// Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super user and can modify it, but only when providing the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this entity.
	Protection *string
	// Indicates system owned resource
	SystemOwned *bool
	// Description of this resource
	Description *string
	// A description field for Traffic group association map。
	DisplayName *string
	// The association map id.
	Id *string
	// The type of this resource.
	ResourceType *string
	// Opaque identifiers meaningful to the API user
	Tags []Tag
	// marked for delete identifier
	MarkedForDelete *bool
	// Path of its parent
	ParentPath *string
	// Absolute path of this object
	Path *string
	// Path relative from its parent
	RelativePath *string
	// List of prefix lists to be distributed between the active edges. Valid for traffic groups of type static_distributed.
	DistributedPrefixLists []DistributedPrefixList
	// The list of prefix lists to be associated.
	PrefixLists []string
	// Possible values are:
	//
	// * TrafficGroupAssociationMap#TrafficGroupAssociationMap_SCOPE_1S_CGW
	// * TrafficGroupAssociationMap#TrafficGroupAssociationMap_SCOPE_0S_VMC
	// * TrafficGroupAssociationMap#TrafficGroupAssociationMap_SCOPE_1S_MGW
	//
	//  The targeted logical router (scope) of prefix lists. Non admin users are not allowed to create, update, delete an association map with scope as /infra/tier-0s/vmc.
	Scope *string
	// Possible values are:
	//
	// * TrafficGroupAssociationMap#TrafficGroupAssociationMap_STATE_IN_PROGRESS
	// * TrafficGroupAssociationMap#TrafficGroupAssociationMap_STATE_SUCCESS
	// * TrafficGroupAssociationMap#TrafficGroupAssociationMap_STATE_FAILURE
	// * TrafficGroupAssociationMap#TrafficGroupAssociationMap_STATE_UNAVAILABLE
	// * TrafficGroupAssociationMap#TrafficGroupAssociationMap_STATE_PENDING
	//
	//  Realized state of the traffic group. This matches the realized state (VmcConsolidatedStatus) of the traffic group.
	State *string
	// Information on the current state. Mostly error messages on failure states.
	StateMessage *string
}

const TrafficGroupAssociationMap_SCOPE_1S_CGW = "/infra/tier-1s/cgw"
const TrafficGroupAssociationMap_SCOPE_0S_VMC = "/infra/tier-0s/vmc"
const TrafficGroupAssociationMap_SCOPE_1S_MGW = "/infra/tier-1s/mgw"
const TrafficGroupAssociationMap_STATE_IN_PROGRESS = "IN_PROGRESS"
const TrafficGroupAssociationMap_STATE_SUCCESS = "SUCCESS"
const TrafficGroupAssociationMap_STATE_FAILURE = "FAILURE"
const TrafficGroupAssociationMap_STATE_UNAVAILABLE = "UNAVAILABLE"
const TrafficGroupAssociationMap_STATE_PENDING = "PENDING"

func (s *TrafficGroupAssociationMap) GetType__() vapiBindings_.BindingType {
	return TrafficGroupAssociationMapBindingType()
}

func (s *TrafficGroupAssociationMap) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for TrafficGroupAssociationMap._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// List result for traffic group association maps.
type TrafficGroupAssociationMapsListResult struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	Self   *SelfResourceLink
	// Opaque cursor to be used for getting next page of records (supplied by current result page)
	Cursor *string
	// Count of results found (across all pages), set only on first page format: int64
	ResultCount *int64
	// If true, results are sorted in ascending order
	SortAscending *bool
	// Field by which records are sorted
	SortBy *string
	// Traffic group association map list
	Results []TrafficGroupAssociationMap
}

func (s *TrafficGroupAssociationMapsListResult) GetType__() vapiBindings_.BindingType {
	return TrafficGroupAssociationMapsListResultBindingType()
}

func (s *TrafficGroupAssociationMapsListResult) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for TrafficGroupAssociationMapsListResult._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Information of mapping between traffic group and corresponding active edge's Virtual Distributed Router Elastic Network Interface.
type TrafficGroupEniMappingInfo struct {
	// Elastic Network Interface of Virtual Distributed Router for linked VPC connectivity.
	Eni *string
	// ID of the traffic group
	TrafficGroupId *string
	// Name of the traffic group
	TrafficGroupName *string
}

func (s *TrafficGroupEniMappingInfo) GetType__() vapiBindings_.BindingType {
	return TrafficGroupEniMappingInfoBindingType()
}

func (s *TrafficGroupEniMappingInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for TrafficGroupEniMappingInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// List result for traffic groups.
type TrafficGroupsListResult struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	Self   *SelfResourceLink
	// Opaque cursor to be used for getting next page of records (supplied by current result page)
	Cursor *string
	// Count of results found (across all pages), set only on first page format: int64
	ResultCount *int64
	// If true, results are sorted in ascending order
	SortAscending *bool
	// Field by which records are sorted
	SortBy *string
	// Traffic group list
	Results []TrafficGroup
}

func (s *TrafficGroupsListResult) GetType__() vapiBindings_.BindingType {
	return TrafficGroupsListResultBindingType()
}

func (s *TrafficGroupsListResult) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for TrafficGroupsListResult._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// This will be used by VMC UI to return the Transit Connect information as part of SddcNetworkingStateInfo API.
type TransitConnectSddcNetworkingStateInfo struct {
	// The total number of Transit Connect advertised routes. format: int32
	AdvertisedRoutes *int64
	// Connected SDDC Group Id
	ConnectedSddcGroupId *string
	// Connected SDDC Group Name
	ConnectedSddcGroupName *string
	// The total number of Transit Connect learned routes. format: int32
	LearnedRoutes *int64
}

func (s *TransitConnectSddcNetworkingStateInfo) GetType__() vapiBindings_.BindingType {
	return TransitConnectSddcNetworkingStateInfoBindingType()
}

func (s *TransitConnectSddcNetworkingStateInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for TransitConnectSddcNetworkingStateInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type TransportNodeCpuReservation struct {
	// Reservation in mHz format: int64
	ReservationInMhz *int64
	// Reservation in shares
	ReservationInShares *string
}

func (s *TransportNodeCpuReservation) GetType__() vapiBindings_.BindingType {
	return TransportNodeCpuReservationBindingType()
}

func (s *TransportNodeCpuReservation) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for TransportNodeCpuReservation._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type TransportNodeMemoryReservation struct {
	// Reservation percentage format: int64
	ReservationPercentage *int64
}

func (s *TransportNodeMemoryReservation) GetType__() vapiBindings_.BindingType {
	return TransportNodeMemoryReservationBindingType()
}

func (s *TransportNodeMemoryReservation) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for TransportNodeMemoryReservation._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// CPU and memory reservation info for ESO edge nodes.
type TransportNodeReservationInfo struct {
	CpuReservation    *TransportNodeCpuReservation
	MemoryReservation *TransportNodeMemoryReservation
}

func (s *TransportNodeReservationInfo) GetType__() vapiBindings_.BindingType {
	return TransportNodeReservationInfoBindingType()
}

func (s *TransportNodeReservationInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for TransportNodeReservationInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Shadow account and linked VPC account
type VMCAccounts struct {
	// linked VPC account number
	LinkedVpcAccount *string
	// Shadow VPC account number
	ShadowAccount *string
}

func (s *VMCAccounts) GetType__() vapiBindings_.BindingType {
	return VMCAccountsBindingType()
}

func (s *VMCAccounts) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for VMCAccounts._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Direct Connect VIFs (Virtual Interface) list query result
type VifsListResult struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	Self   *SelfResourceLink
	// Opaque cursor to be used for getting next page of records (supplied by current result page)
	Cursor *string
	// Count of results found (across all pages), set only on first page format: int64
	ResultCount *int64
	// If true, results are sorted in ascending order
	SortAscending *bool
	// Field by which records are sorted
	SortBy *string
	// VIFs list
	Results []VirtualInterface
}

func (s *VifsListResult) GetType__() vapiBindings_.BindingType {
	return VifsListResultBindingType()
}

func (s *VifsListResult) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for VifsListResult._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type VirtualInterface struct {
	// Possible values are:
	//
	// * VirtualInterface#VirtualInterface_BGP_STATUS_UP
	// * VirtualInterface#VirtualInterface_BGP_STATUS_DOWN
	//
	//  BGP status
	BgpStatus *string
	// Identifier for the Direct Connect
	DirectConnectId *string
	// Identifier for the virtual interface
	Id *string
	// amazon side address format: ipv4
	LocalIp *string
	// Maximum transmission unit allowed by the VIF format: int32
	Mtu *int64
	// VIF name
	Name *string
	// Remote autonomous system number of vif
	RemoteAsn *string
	// customer address format: ipv4
	RemoteIp *string
	// Possible values are:
	//
	// * VirtualInterface#VirtualInterface_STATE_CONFIRMING
	// * VirtualInterface#VirtualInterface_STATE_VERIFYING
	// * VirtualInterface#VirtualInterface_STATE_PENDING
	// * VirtualInterface#VirtualInterface_STATE_AVAILABLE
	// * VirtualInterface#VirtualInterface_STATE_DOWN
	// * VirtualInterface#VirtualInterface_STATE_DELETING
	// * VirtualInterface#VirtualInterface_STATE_DELETED
	// * VirtualInterface#VirtualInterface_STATE_REJECTED
	// * VirtualInterface#VirtualInterface_STATE_ATTACHED
	// * VirtualInterface#VirtualInterface_STATE_ATTACHING
	// * VirtualInterface#VirtualInterface_STATE_ERROR
	//
	//  VIF State
	State *string
}

const VirtualInterface_BGP_STATUS_UP = "UP"
const VirtualInterface_BGP_STATUS_DOWN = "DOWN"
const VirtualInterface_STATE_CONFIRMING = "CONFIRMING"
const VirtualInterface_STATE_VERIFYING = "VERIFYING"
const VirtualInterface_STATE_PENDING = "PENDING"
const VirtualInterface_STATE_AVAILABLE = "AVAILABLE"
const VirtualInterface_STATE_DOWN = "DOWN"
const VirtualInterface_STATE_DELETING = "DELETING"
const VirtualInterface_STATE_DELETED = "DELETED"
const VirtualInterface_STATE_REJECTED = "REJECTED"
const VirtualInterface_STATE_ATTACHED = "ATTACHED"
const VirtualInterface_STATE_ATTACHING = "ATTACHING"
const VirtualInterface_STATE_ERROR = "ERROR"

func (s *VirtualInterface) GetType__() vapiBindings_.BindingType {
	return VirtualInterfaceBindingType()
}

func (s *VirtualInterface) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for VirtualInterface._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Abstract base class for all the VmcApp objects.
type VmcAppBaseInfo struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	Self   *SelfResourceLink
	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected. format: int32
	Revision *int64
	// Timestamp of resource creation format: int64
	CreateTime *int64
	// ID of the user who created this resource
	CreateUser *string
	// Timestamp of last modification format: int64
	LastModifiedTime *int64
	// ID of the user who last modified this resource
	LastModifiedUser *string
	// Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super user and can modify it, but only when providing the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this entity.
	Protection *string
	// Indicates system owned resource
	SystemOwned *bool
	// Description of this resource
	Description *string
	// Defaults to ID if not set
	DisplayName *string
	// Unique identifier of this resource
	Id *string
	// The type of this resource.
	ResourceType *string
	// Opaque identifiers meaningful to the API user
	Tags []Tag
	// marked for delete identifier
	MarkedForDelete *bool
	// Path of its parent
	ParentPath *string
	// Absolute path of this object
	Path *string
	// Path relative from its parent
	RelativePath *string
}

func (s *VmcAppBaseInfo) GetType__() vapiBindings_.BindingType {
	return VmcAppBaseInfoBindingType()
}

func (s *VmcAppBaseInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for VmcAppBaseInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Represents aggregated realized status for intent entity across associated realized entities.
type VmcConsolidatedRealizedStatus struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	Self   *SelfResourceLink
	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected. format: int32
	Revision *int64
	// Timestamp of resource creation format: int64
	CreateTime *int64
	// ID of the user who created this resource
	CreateUser *string
	// Timestamp of last modification format: int64
	LastModifiedTime *int64
	// ID of the user who last modified this resource
	LastModifiedUser *string
	// Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super user and can modify it, but only when providing the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this entity.
	Protection *string
	// Indicates system owned resource
	SystemOwned *bool
	// Description of this resource
	Description *string
	// Defaults to ID if not set
	DisplayName *string
	// Unique identifier of this resource
	Id *string
	// The type of this resource.
	ResourceType *string
	// Opaque identifiers meaningful to the API user
	Tags []Tag
	// marked for delete identifier
	MarkedForDelete *bool
	// Path of its parent
	ParentPath *string
	// Absolute path of this object
	Path *string
	// Path relative from its parent
	RelativePath       *string
	ConsolidatedStatus *VmcConsolidatedStatus
	// Aggregated consolidated status by enforcement point.
	ConsolidatedStatusPerObject []VmcConsolidatedStatusPerObject
	// Intent path of the object representing this consolidated state.
	IntentPath *string
}

func (s *VmcConsolidatedRealizedStatus) GetType__() vapiBindings_.BindingType {
	return VmcConsolidatedRealizedStatusBindingType()
}

func (s *VmcConsolidatedRealizedStatus) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for VmcConsolidatedRealizedStatus._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Consolidated status of an object.
type VmcConsolidatedStatus struct {
	// Possible values are:
	//
	// * VmcConsolidatedStatus#VmcConsolidatedStatus_CONSOLIDATED_STATUS_IN_PROGRESS
	// * VmcConsolidatedStatus#VmcConsolidatedStatus_CONSOLIDATED_STATUS_SUCCESS
	// * VmcConsolidatedStatus#VmcConsolidatedStatus_CONSOLIDATED_STATUS_FAILURE
	// * VmcConsolidatedStatus#VmcConsolidatedStatus_CONSOLIDATED_STATUS_UNAVAILABLE
	// * VmcConsolidatedStatus#VmcConsolidatedStatus_CONSOLIDATED_STATUS_PENDING
	//
	//  Realized state of consolidation.
	ConsolidatedStatus *string
	// Help message for the current status regarding an object, providing information for each state.
	StatusMessage *string
}

const VmcConsolidatedStatus_CONSOLIDATED_STATUS_IN_PROGRESS = "IN_PROGRESS"
const VmcConsolidatedStatus_CONSOLIDATED_STATUS_SUCCESS = "SUCCESS"
const VmcConsolidatedStatus_CONSOLIDATED_STATUS_FAILURE = "FAILURE"
const VmcConsolidatedStatus_CONSOLIDATED_STATUS_UNAVAILABLE = "UNAVAILABLE"
const VmcConsolidatedStatus_CONSOLIDATED_STATUS_PENDING = "PENDING"

func (s *VmcConsolidatedStatus) GetType__() vapiBindings_.BindingType {
	return VmcConsolidatedStatusBindingType()
}

func (s *VmcConsolidatedStatus) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for VmcConsolidatedStatus._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Realized status consolidated by individual objects.
type VmcConsolidatedStatusPerObject struct {
	ConsolidatedStatus *VmcConsolidatedStatus
	// Object id used to consolidate state. This can be a particular backend task/job, etc.
	ObjectId *string
}

func (s *VmcConsolidatedStatusPerObject) GetType__() vapiBindings_.BindingType {
	return VmcConsolidatedStatusPerObjectBindingType()
}

func (s *VmcConsolidatedStatusPerObject) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for VmcConsolidatedStatusPerObject._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// VMC Feature Flag
type VmcFeatureFlagInfo struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	Self   *SelfResourceLink
	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected. format: int32
	Revision *int64
	// Timestamp of resource creation format: int64
	CreateTime *int64
	// ID of the user who created this resource
	CreateUser *string
	// Timestamp of last modification format: int64
	LastModifiedTime *int64
	// ID of the user who last modified this resource
	LastModifiedUser *string
	// Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super user and can modify it, but only when providing the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this entity.
	Protection *string
	// Indicates system owned resource
	SystemOwned *bool
	// Description of this resource
	Description *string
	// Display Name
	DisplayName *string
	// Unique identifier of this resource
	Id *string
	// The type of this resource.
	ResourceType *string
	// Opaque identifiers meaningful to the API user
	Tags []Tag
	// marked for delete identifier
	MarkedForDelete *bool
	// Path of its parent
	ParentPath *string
	// Absolute path of this object
	Path *string
	// Path relative from its parent
	RelativePath *string
	// Internal Name
	InternalName *string
	// Message
	Message *string
	// Feature Name
	Name *string
	// Possible values are:
	//
	// * VmcFeatureFlagInfo#VmcFeatureFlagInfo_STATE_ENABLED
	// * VmcFeatureFlagInfo#VmcFeatureFlagInfo_STATE_DISABLED
	// * VmcFeatureFlagInfo#VmcFeatureFlagInfo_STATE_INACTIVE
	//
	//  state
	State *string
	// Unlicensed Message
	UnlicensedMessage *string
}

const VmcFeatureFlagInfo_STATE_ENABLED = "enabled"
const VmcFeatureFlagInfo_STATE_DISABLED = "disabled"
const VmcFeatureFlagInfo_STATE_INACTIVE = "inactive"

func (s *VmcFeatureFlagInfo) GetType__() vapiBindings_.BindingType {
	return VmcFeatureFlagInfoBindingType()
}

func (s *VmcFeatureFlagInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for VmcFeatureFlagInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// VMC Feature flags
type VmcFeatureFlags struct {
	Features []VmcFeatureFlagInfo
}

func (s *VmcFeatureFlags) GetType__() vapiBindings_.BindingType {
	return VmcFeatureFlagsBindingType()
}

func (s *VmcFeatureFlags) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for VmcFeatureFlags._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Tier0 interface statistics for traffic-group edges.
type VmcInterfaceStatistics struct {
	// Timestamp when the data was last updated; unset if data source has never updated the data. format: int64
	LastUpdateTimestamp *int64
	Rx                  *LogicalRouterPortCounters
	Tx                  *LogicalRouterPortCounters
	// The ID of the logical router port
	LogicalRouterPortId *string
	// Policy path for the interface
	InterfacePolicyPath *string
}

func (s *VmcInterfaceStatistics) GetType__() vapiBindings_.BindingType {
	return VmcInterfaceStatisticsBindingType()
}

func (s *VmcInterfaceStatistics) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for VmcInterfaceStatistics._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// A consolidated object of realized entities given an intent path. This accounts for resources / entities which are realized under the intent path.
type VmcRealizedEntities struct {
	// Intent path
	IntentPath *string
	// Detailed realized entities list.
	RealizedEntities []VmcRealizedEntity
	// Realized entities id
	RealizedEntitiesId *string
	// Possible values are:
	//
	// * VmcRealizedEntities#VmcRealizedEntities_REALIZED_STATE_IN_PROGRESS
	// * VmcRealizedEntities#VmcRealizedEntities_REALIZED_STATE_SUCCESS
	// * VmcRealizedEntities#VmcRealizedEntities_REALIZED_STATE_FAILURE
	// * VmcRealizedEntities#VmcRealizedEntities_REALIZED_STATE_UNAVAILABLE
	// * VmcRealizedEntities#VmcRealizedEntities_REALIZED_STATE_PENDING
	//
	//  Realized state
	RealizedState *string
}

const VmcRealizedEntities_REALIZED_STATE_IN_PROGRESS = "IN_PROGRESS"
const VmcRealizedEntities_REALIZED_STATE_SUCCESS = "SUCCESS"
const VmcRealizedEntities_REALIZED_STATE_FAILURE = "FAILURE"
const VmcRealizedEntities_REALIZED_STATE_UNAVAILABLE = "UNAVAILABLE"
const VmcRealizedEntities_REALIZED_STATE_PENDING = "PENDING"

func (s *VmcRealizedEntities) GetType__() vapiBindings_.BindingType {
	return VmcRealizedEntitiesBindingType()
}

func (s *VmcRealizedEntities) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for VmcRealizedEntities._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Realized entity object. This accounts for an resource / entity which is realized within cloud service.
type VmcRealizedEntity struct {
	// Resource realization API path
	RealizationApi *string
	// Resource realization id. This can differ from realized_entity_id as this can be an external id.
	RealizationId *string
	// The path for the realization of an entity. This can be an URI, etc. Some resources are identified by their paths.
	RealizationPath *string
	// Realized entity id
	RealizedEntityId *string
	// Realized entity type
	RealizedEntityType *string
	// Possible values are:
	//
	// * VmcRealizedEntity#VmcRealizedEntity_REALIZED_STATE_IN_PROGRESS
	// * VmcRealizedEntity#VmcRealizedEntity_REALIZED_STATE_SUCCESS
	// * VmcRealizedEntity#VmcRealizedEntity_REALIZED_STATE_FAILURE
	// * VmcRealizedEntity#VmcRealizedEntity_REALIZED_STATE_UNAVAILABLE
	// * VmcRealizedEntity#VmcRealizedEntity_REALIZED_STATE_PENDING
	//
	//  Realized state
	RealizedState *string
}

const VmcRealizedEntity_REALIZED_STATE_IN_PROGRESS = "IN_PROGRESS"
const VmcRealizedEntity_REALIZED_STATE_SUCCESS = "SUCCESS"
const VmcRealizedEntity_REALIZED_STATE_FAILURE = "FAILURE"
const VmcRealizedEntity_REALIZED_STATE_UNAVAILABLE = "UNAVAILABLE"
const VmcRealizedEntity_REALIZED_STATE_PENDING = "PENDING"

func (s *VmcRealizedEntity) GetType__() vapiBindings_.BindingType {
	return VmcRealizedEntityBindingType()
}

func (s *VmcRealizedEntity) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for VmcRealizedEntity._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// VPN endpoint information
type VpnEndpoint struct {
	// Interface label of the VPN endpoint
	InterfaceLabel *string
	// IP address of the VPN endpoint format: ipv4
	Ip *string
	// Name of the VPN endpoint
	Name *string
	// Type of the VPN endpoint
	Type_ *string
}

func (s *VpnEndpoint) GetType__() vapiBindings_.BindingType {
	return VpnEndpointBindingType()
}

func (s *VpnEndpoint) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for VpnEndpoint._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// This will be used by VMC UI to return the VPN information as part of SddcNetworkingStateInfo API.
type VpnSddcNetworkingStateInfo struct {
	// The total number of VPN connections with connectivity status (it's BGP session status) deactivated. format: int32
	ConnectivityDeactivated *int64
	// The total number of VPN connections with connectivity status (it's BGP session status) down. format: int32
	ConnectivityDown *int64
	// The total number of VPN connections with connectivity status (it's BGP session status) up. format: int32
	ConnectivityUp *int64
	// VPN Public IP.
	PublicIp *string
	// The total number of VPN connections. format: int32
	TotalConnections *int64
	// Array of VPN connections with failed status.
	VpnDownConnections []string
}

func (s *VpnSddcNetworkingStateInfo) GetType__() vapiBindings_.BindingType {
	return VpnSddcNetworkingStateInfoBindingType()
}

func (s *VpnSddcNetworkingStateInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for VpnSddcNetworkingStateInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

func ActionMessageBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["message"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["message"] = "Message"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.action_message", fields, reflect.TypeOf(ActionMessage{}), fieldNameMap, validators)
}

func AdvertisedRouteBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["address_family"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["address_family"] = "AddressFamily"
	fields["advertisement_state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["advertisement_state"] = "AdvertisementState"
	fields["cidr"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cidr"] = "Cidr"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.advertised_route", fields, reflect.TypeOf(AdvertisedRoute{}), fieldNameMap, validators)
}

func AggregatedLogicalRouterPortCountersBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["last_update_timestamp"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["last_update_timestamp"] = "LastUpdateTimestamp"
	fields["rx"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(LogicalRouterPortCountersBindingType))
	fieldNameMap["rx"] = "Rx"
	fields["tx"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(LogicalRouterPortCountersBindingType))
	fieldNameMap["tx"] = "Tx"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.aggregated_logical_router_port_counters", fields, reflect.TypeOf(AggregatedLogicalRouterPortCounters{}), fieldNameMap, validators)
}

func AggregationRouteConfigInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["advertise_always"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["advertise_always"] = "AdvertiseAlways"
	fields["aggregation_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["aggregation_path"] = "AggregationPath"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.aggregation_route_config_info", fields, reflect.TypeOf(AggregationRouteConfigInfo{}), fieldNameMap, validators)
}

func ApiErrorBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["details"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["details"] = "Details"
	fields["error_code"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["error_code"] = "ErrorCode"
	fields["error_data"] = vapiBindings_.NewOptionalType(vapiBindings_.NewDynamicStructType(nil))
	fieldNameMap["error_data"] = "ErrorData"
	fields["error_message"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["error_message"] = "ErrorMessage"
	fields["module_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["module_name"] = "ModuleName"
	fields["related_errors"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(RelatedApiErrorBindingType), reflect.TypeOf([]RelatedApiError{})))
	fieldNameMap["related_errors"] = "RelatedErrors"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.api_error", fields, reflect.TypeOf(ApiError{}), fieldNameMap, validators)
}

func AssociatedBaseGroupConnectionInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_revision"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_create_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_create_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_last_modified_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["_protection"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_system_owned"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["description"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["display_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = vapiBindings_.NewStringType()
	fieldNameMap["resource_type"] = "ResourceType"
	fields["tags"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["marked_for_delete"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["parent_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.associated_base_group_connection_info", fields, reflect.TypeOf(AssociatedBaseGroupConnectionInfo{}), fieldNameMap, validators)
}

func AssociatedGroupConnectionInfosListResultBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fields["result_count"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["result_count"] = "ResultCount"
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["sort_ascending"] = "SortAscending"
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sort_by"] = "SortBy"
	fields["results"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewDynamicStructType([]vapiBindings_.ReferenceType{vapiBindings_.NewReferenceType(AssociatedBaseGroupConnectionInfoBindingType)}), reflect.TypeOf([]*vapiData_.StructValue{})))
	fieldNameMap["results"] = "Results"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.associated_group_connection_infos_list_result", fields, reflect.TypeOf(AssociatedGroupConnectionInfosListResult{}), fieldNameMap, validators)
}

func AssociatedTgwGroupConnectionInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["external_route_table_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["external_route_table_id"] = "ExternalRouteTableId"
	fields["sddcs_route_table_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sddcs_route_table_id"] = "SddcsRouteTableId"
	fields["state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["state"] = "State"
	fields["tgw_attachment_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["tgw_attachment_id"] = "TgwAttachmentId"
	fields["tgw_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["tgw_id"] = "TgwId"
	fields["_links"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_revision"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_create_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_create_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_last_modified_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["_protection"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_system_owned"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["description"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["display_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = vapiBindings_.NewStringType()
	fieldNameMap["resource_type"] = "ResourceType"
	fields["tags"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["marked_for_delete"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["parent_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.associated_tgw_group_connection_info", fields, reflect.TypeOf(AssociatedTgwGroupConnectionInfo{}), fieldNameMap, validators)
}

func AssociatedTgwGroupConnectionInfoVersion2BindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["associated_route_table"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["associated_route_table"] = "AssociatedRouteTable"
	fields["region"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["region"] = "Region"
	fields["sddc_advertise_route_config"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SddcAdvertiseRouteConfigBindingType))
	fieldNameMap["sddc_advertise_route_config"] = "SddcAdvertiseRouteConfig"
	fields["state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["state"] = "State"
	fields["tgw_attachment_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["tgw_attachment_id"] = "TgwAttachmentId"
	fields["tgw_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["tgw_id"] = "TgwId"
	fields["_links"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_revision"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_create_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_create_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_last_modified_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["_protection"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_system_owned"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["description"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["display_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = vapiBindings_.NewStringType()
	fieldNameMap["resource_type"] = "ResourceType"
	fields["tags"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["marked_for_delete"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["parent_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.associated_tgw_group_connection_info_version2", fields, reflect.TypeOf(AssociatedTgwGroupConnectionInfoVersion2{}), fieldNameMap, validators)
}

func AwsPrefixInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["address_family"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["address_family"] = "AddressFamily"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["region"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["region"] = "Region"
	fields["size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["size"] = "Size"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.aws_prefix_info", fields, reflect.TypeOf(AwsPrefixInfo{}), fieldNameMap, validators)
}

func AwsResourceShareInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["aws_resource_share_arn"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["aws_resource_share_arn"] = "AwsResourceShareArn"
	fields["aws_resource_share_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["aws_resource_share_name"] = "AwsResourceShareName"
	fields["aws_resource_share_state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["aws_resource_share_state"] = "AwsResourceShareState"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.aws_resource_share_info", fields, reflect.TypeOf(AwsResourceShareInfo{}), fieldNameMap, validators)
}

func BGPAdvertisedRoutesBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["advertised_routes"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(AdvertisedRouteBindingType), reflect.TypeOf([]AdvertisedRoute{})))
	fieldNameMap["advertised_routes"] = "AdvertisedRoutes"
	fields["failed_advertised_routes"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["failed_advertised_routes"] = "FailedAdvertisedRoutes"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.BGP_advertised_routes", fields, reflect.TypeOf(BGPAdvertisedRoutes{}), fieldNameMap, validators)
}

func BGPLearnedRoutesBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipv4_cidr"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["ipv4_cidr"] = "Ipv4Cidr"
	fields["ipv6_cidr"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["ipv6_cidr"] = "Ipv6Cidr"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.BGP_learned_routes", fields, reflect.TypeOf(BGPLearnedRoutes{}), fieldNameMap, validators)
}

func CloudProviderSddcNetworkingStateInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["account_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["account_id"] = "AccountId"
	fields["availability_zone"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["availability_zone"] = "AvailabilityZone"
	fields["cloud_formation_stack_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cloud_formation_stack_name"] = "CloudFormationStackName"
	fields["connectivity_status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["connectivity_status"] = "ConnectivityStatus"
	fields["ec2_enabled"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["ec2_enabled"] = "Ec2Enabled"
	fields["iam_role_names"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["iam_role_names"] = "IamRoleNames"
	fields["s3_enabled"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["s3_enabled"] = "S3Enabled"
	fields["traffic_group_eni_mapping_infos"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(TrafficGroupEniMappingInfoBindingType), reflect.TypeOf([]TrafficGroupEniMappingInfo{})))
	fieldNameMap["traffic_group_eni_mapping_infos"] = "TrafficGroupEniMappingInfos"
	fields["vpc"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(LinkedSubnetInfoBindingType), reflect.TypeOf([]LinkedSubnetInfo{})))
	fieldNameMap["vpc"] = "Vpc"
	fields["vpc_cidr"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["vpc_cidr"] = "VpcCidr"
	fields["vpc_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vpc_id"] = "VpcId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.cloud_provider_sddc_networking_state_info", fields, reflect.TypeOf(CloudProviderSddcNetworkingStateInfo{}), fieldNameMap, validators)
}

func ComputeGatewaySddcNetworkingStateInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["dns"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["dns"] = "Dns"
	fields["total_firewall_rules"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["total_firewall_rules"] = "TotalFirewallRules"
	fields["total_segments"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["total_segments"] = "TotalSegments"
	fields["total_tier1_gateways"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["total_tier1_gateways"] = "TotalTier1Gateways"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.compute_gateway_sddc_networking_state_info", fields, reflect.TypeOf(ComputeGatewaySddcNetworkingStateInfo{}), fieldNameMap, validators)
}

func ConnectedServiceListResultBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fields["result_count"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["result_count"] = "ResultCount"
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["sort_ascending"] = "SortAscending"
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sort_by"] = "SortBy"
	fields["results"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ConnectedServiceStatusBindingType), reflect.TypeOf([]ConnectedServiceStatus{})))
	fieldNameMap["results"] = "Results"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.connected_service_list_result", fields, reflect.TypeOf(ConnectedServiceListResult{}), fieldNameMap, validators)
}

func ConnectedServiceStatusBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_revision"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_create_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_create_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_last_modified_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["_protection"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_system_owned"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["description"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["display_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["tags"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["marked_for_delete"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["parent_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	fields["enabled"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["enabled"] = "Enabled"
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.connected_service_status", fields, reflect.TypeOf(ConnectedServiceStatus{}), fieldNameMap, validators)
}

func ConnectivityEndpointConfigBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_revision"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_create_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_create_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_last_modified_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["_protection"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_system_owned"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["description"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["display_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["tags"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["marked_for_delete"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["parent_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.connectivity_endpoint_config", fields, reflect.TypeOf(ConnectivityEndpointConfig{}), fieldNameMap, validators)
}

func ConnectivityEndpointListResultBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fields["result_count"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["result_count"] = "ResultCount"
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["sort_ascending"] = "SortAscending"
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sort_by"] = "SortBy"
	fields["results"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ConnectivityEndpointConfigBindingType), reflect.TypeOf([]ConnectivityEndpointConfig{})))
	fieldNameMap["results"] = "Results"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.connectivity_endpoint_list_result", fields, reflect.TypeOf(ConnectivityEndpointListResult{}), fieldNameMap, validators)
}

func CsvListResultBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["file_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["file_name"] = "FileName"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.csv_list_result", fields, reflect.TypeOf(CsvListResult{}), fieldNameMap, validators)
}

func CsvRecordBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.csv_record", fields, reflect.TypeOf(CsvRecord{}), fieldNameMap, validators)
}

func CustomerBgpPeerInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["bgp_md5_authentication_key"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["bgp_md5_authentication_key"] = "BgpMd5AuthenticationKey"
	fields["display_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["index"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["index"] = "Index"
	fields["local_peer_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["local_peer_ip"] = "LocalPeerIp"
	fields["provider_token"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["provider_token"] = "ProviderToken"
	fields["provider_token_expiration"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["provider_token_expiration"] = "ProviderTokenExpiration"
	fields["remote_peer_asn"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["remote_peer_asn"] = "RemotePeerAsn"
	fields["remote_peer_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["remote_peer_ip"] = "RemotePeerIp"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.customer_bgp_peer_info", fields, reflect.TypeOf(CustomerBgpPeerInfo{}), fieldNameMap, validators)
}

func CustomerBgpPeerInfoStatusBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["index"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["index"] = "Index"
	fields["status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["status"] = "Status"
	fields["status_detail"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["status_detail"] = "StatusDetail"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.customer_bgp_peer_info_status", fields, reflect.TypeOf(CustomerBgpPeerInfoStatus{}), fieldNameMap, validators)
}

func DirectConnectBgpInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_revision"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_create_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_create_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_last_modified_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["_protection"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_system_owned"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["description"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["display_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["tags"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["marked_for_delete"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["parent_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	fields["local_as_num"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["local_as_num"] = "LocalAsNum"
	fields["mtu"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["mtu"] = "Mtu"
	fields["route_preference"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["route_preference"] = "RoutePreference"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.direct_connect_bgp_info", fields, reflect.TypeOf(DirectConnectBgpInfo{}), fieldNameMap, validators)
}

func DirectConnectSddcNetworkingStateInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["advertised_routes"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["advertised_routes"] = "AdvertisedRoutes"
	fields["connectivity_down"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["connectivity_down"] = "ConnectivityDown"
	fields["connectivity_up"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["connectivity_up"] = "ConnectivityUp"
	fields["down_vif_names"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["down_vif_names"] = "DownVifNames"
	fields["learned_routes"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["learned_routes"] = "LearnedRoutes"
	fields["total_attached_vifs_count"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["total_attached_vifs_count"] = "TotalAttachedVifsCount"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.direct_connect_sddc_networking_state_info", fields, reflect.TypeOf(DirectConnectSddcNetworkingStateInfo{}), fieldNameMap, validators)
}

func DistributedPrefixListBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fields["result_count"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["result_count"] = "ResultCount"
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["sort_ascending"] = "SortAscending"
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sort_by"] = "SortBy"
	fields["index"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["index"] = "Index"
	fields["prefix_lists"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["prefix_lists"] = "PrefixLists"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.distributed_prefix_list", fields, reflect.TypeOf(DistributedPrefixList{}), fieldNameMap, validators)
}

func ExternalConnectivityConfigBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_revision"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_create_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_create_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_last_modified_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["_protection"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_system_owned"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["description"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["display_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["tags"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["marked_for_delete"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["parent_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	fields["internet_interface"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(InterfaceConfigBindingType))
	fieldNameMap["internet_interface"] = "InternetInterface"
	fields["internet_mtu"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["internet_mtu"] = "InternetMtu"
	fields["intranet_interface"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(InterfaceConfigBindingType))
	fieldNameMap["intranet_interface"] = "IntranetInterface"
	fields["intranet_mtu"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["intranet_mtu"] = "IntranetMtu"
	fields["services_interface"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(InterfaceConfigBindingType))
	fieldNameMap["services_interface"] = "ServicesInterface"
	fields["services_mtu"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["services_mtu"] = "ServicesMtu"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.external_connectivity_config", fields, reflect.TypeOf(ExternalConnectivityConfig{}), fieldNameMap, validators)
}

func ExternalConnectivityInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["connectivities"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["connectivities"] = "Connectivities"
	fields["endpoint_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["endpoint_type"] = "EndpointType"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.external_connectivity_info", fields, reflect.TypeOf(ExternalConnectivityInfo{}), fieldNameMap, validators)
}

func ExternalConnectivityStatusBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["status"] = "Status"
	fields["type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["type"] = "Type_"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.external_connectivity_status", fields, reflect.TypeOf(ExternalConnectivityStatus{}), fieldNameMap, validators)
}

func ExternalConnectivityStatusListBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fields["result_count"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["result_count"] = "ResultCount"
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["sort_ascending"] = "SortAscending"
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sort_by"] = "SortBy"
	fields["connectivities"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ExternalConnectivityStatusBindingType), reflect.TypeOf([]ExternalConnectivityStatus{})))
	fieldNameMap["connectivities"] = "Connectivities"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.external_connectivity_status_list", fields, reflect.TypeOf(ExternalConnectivityStatusList{}), fieldNameMap, validators)
}

func ExternalSddcConnectivityBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["connectivity_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["connectivity_type"] = "ConnectivityType"
	fields["regions"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["regions"] = "Regions"
	fields["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["source"] = "Source"
	fields["status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["status"] = "Status"
	fields["status_message"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["status_message"] = "StatusMessage"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.external_sddc_connectivity", fields, reflect.TypeOf(ExternalSddcConnectivity{}), fieldNameMap, validators)
}

func ExternalSddcRouteBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_revision"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_create_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_create_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_last_modified_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["_protection"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_system_owned"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["description"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["display_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["tags"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["marked_for_delete"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["parent_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	fields["connectivities"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ExternalSddcConnectivityBindingType), reflect.TypeOf([]ExternalSddcConnectivity{})))
	fieldNameMap["connectivities"] = "Connectivities"
	fields["destination"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["destination"] = "Destination"
	fields["is_vif_attached"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["is_vif_attached"] = "IsVifAttached"
	fields["source_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["source_path"] = "SourcePath"
	fields["summarized"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["summarized"] = "Summarized"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.external_sddc_route", fields, reflect.TypeOf(ExternalSddcRoute{}), fieldNameMap, validators)
}

func ExternalSddcRouteCsvRecordBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["connectivity_details"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["connectivity_details"] = "ConnectivityDetails"
	fields["destination"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["destination"] = "Destination"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.external_sddc_route_csv_record", fields, reflect.TypeOf(ExternalSddcRouteCsvRecord{}), fieldNameMap, validators)
}

func ExternalSddcRoutesListResultBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fields["result_count"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["result_count"] = "ResultCount"
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["sort_ascending"] = "SortAscending"
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sort_by"] = "SortBy"
	fields["routes"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ExternalSddcRouteBindingType), reflect.TypeOf([]ExternalSddcRoute{})))
	fieldNameMap["routes"] = "Routes"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.external_sddc_routes_list_result", fields, reflect.TypeOf(ExternalSddcRoutesListResult{}), fieldNameMap, validators)
}

func ExternalSddcRoutesListResultInCsvFormatBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["file_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["file_name"] = "FileName"
	fields["results"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ExternalSddcRouteCsvRecordBindingType), reflect.TypeOf([]ExternalSddcRouteCsvRecord{})))
	fieldNameMap["results"] = "Results"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.external_sddc_routes_list_result_in_csv_format", fields, reflect.TypeOf(ExternalSddcRoutesListResultInCsvFormat{}), fieldNameMap, validators)
}

func FilteredMgwGatewayPoliciesBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["error_msg"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["error_msg"] = "ErrorMsg"
	fields["invalid_rules"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(FilteredMgwGatewayPoliciesKeyValuePairsBindingType), reflect.TypeOf([]FilteredMgwGatewayPoliciesKeyValuePairs{})))
	fieldNameMap["invalid_rules"] = "InvalidRules"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.filtered_mgw_gateway_policies", fields, reflect.TypeOf(FilteredMgwGatewayPolicies{}), fieldNameMap, validators)
}

func FilteredMgwGatewayPoliciesKeyValuePairsBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["key"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["key"] = "Key"
	fields["value"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(FilteredMgwGatewayPoliciesObjectBindingType), reflect.TypeOf([]FilteredMgwGatewayPoliciesObject{})))
	fieldNameMap["value"] = "Value"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.filtered_mgw_gateway_policies_key_value_pairs", fields, reflect.TypeOf(FilteredMgwGatewayPoliciesKeyValuePairs{}), fieldNameMap, validators)
}

func FilteredMgwGatewayPoliciesObjectBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["rule_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["rule_id"] = "RuleId"
	fields["rule_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["rule_name"] = "RuleName"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.filtered_mgw_gateway_policies_object", fields, reflect.TypeOf(FilteredMgwGatewayPoliciesObject{}), fieldNameMap, validators)
}

func IncludedFieldsParametersBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["included_fields"] = "IncludedFields"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.included_fields_parameters", fields, reflect.TypeOf(IncludedFieldsParameters{}), fieldNameMap, validators)
}

func InterfaceConfigBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["route_filtering_state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["route_filtering_state"] = "RouteFilteringState"
	fields["urpf_mode"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["urpf_mode"] = "UrpfMode"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.interface_config", fields, reflect.TypeOf(InterfaceConfig{}), fieldNameMap, validators)
}

func IntranetConnectionConfigurationBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_revision"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_create_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_create_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_last_modified_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["_protection"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_system_owned"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["description"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["display_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["tags"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["marked_for_delete"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["parent_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	fields["connection_redundancy"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["connection_redundancy"] = "ConnectionRedundancy"
	fields["virtual_circuits"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(CustomerBgpPeerInfoBindingType), reflect.TypeOf([]CustomerBgpPeerInfo{})))
	fieldNameMap["virtual_circuits"] = "VirtualCircuits"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.intranet_connection_configuration", fields, reflect.TypeOf(IntranetConnectionConfiguration{}), fieldNameMap, validators)
}

func IntranetConnectionConfigurationListResultBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fields["result_count"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["result_count"] = "ResultCount"
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["sort_ascending"] = "SortAscending"
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sort_by"] = "SortBy"
	fields["results"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(IntranetConnectionConfigurationBindingType), reflect.TypeOf([]IntranetConnectionConfiguration{})))
	fieldNameMap["results"] = "Results"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.intranet_connection_configuration_list_result", fields, reflect.TypeOf(IntranetConnectionConfigurationListResult{}), fieldNameMap, validators)
}

func IntranetConnectivityStatusBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["aggregate_status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["aggregate_status"] = "AggregateStatus"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["virtual_circuits"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(CustomerBgpPeerInfoStatusBindingType), reflect.TypeOf([]CustomerBgpPeerInfoStatus{})))
	fieldNameMap["virtual_circuits"] = "VirtualCircuits"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.intranet_connectivity_status", fields, reflect.TypeOf(IntranetConnectivityStatus{}), fieldNameMap, validators)
}

func IpAttachmentPairBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["attachment_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["attachment_id"] = "AttachmentId"
	fields["ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["ip"] = "Ip"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.ip_attachment_pair", fields, reflect.TypeOf(IpAttachmentPair{}), fieldNameMap, validators)
}

func LinkedSubnetInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["availability_zone"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["availability_zone"] = "AvailabilityZone"
	fields["availability_zone_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["availability_zone_id"] = "AvailabilityZoneId"
	fields["cidr"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cidr"] = "Cidr"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["ipv6_cidr"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["ipv6_cidr"] = "Ipv6Cidr"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.linked_subnet_info", fields, reflect.TypeOf(LinkedSubnetInfo{}), fieldNameMap, validators)
}

func LinkedVpciPv6ModeInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipv6_mode"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["ipv6_mode"] = "Ipv6Mode"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.linked_vpci_pv6_mode_info", fields, reflect.TypeOf(LinkedVpciPv6ModeInfo{}), fieldNameMap, validators)
}

func LinkedVpcInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_revision"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_create_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_create_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_last_modified_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["_protection"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_system_owned"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["description"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["display_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["tags"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["marked_for_delete"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["parent_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	fields["active_eni"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["active_eni"] = "ActiveEni"
	fields["arn_role"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["arn_role"] = "ArnRole"
	fields["external_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["external_id"] = "ExternalId"
	fields["linked_account"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["linked_account"] = "LinkedAccount"
	fields["linked_vpc_addresses"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["linked_vpc_addresses"] = "LinkedVpcAddresses"
	fields["linked_vpc_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["linked_vpc_id"] = "LinkedVpcId"
	fields["linked_vpc_ipv6_addresses"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["linked_vpc_ipv6_addresses"] = "LinkedVpcIpv6Addresses"
	fields["linked_vpc_ipv6_mode_info"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(LinkedVpciPv6ModeInfoBindingType))
	fieldNameMap["linked_vpc_ipv6_mode_info"] = "LinkedVpcIpv6ModeInfo"
	fields["linked_vpc_managed_prefix_list_info"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(LinkedVpcManagedPrefixListSupportInfoBindingType))
	fieldNameMap["linked_vpc_managed_prefix_list_info"] = "LinkedVpcManagedPrefixListInfo"
	fields["linked_vpc_nat_ips"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["linked_vpc_nat_ips"] = "LinkedVpcNatIps"
	fields["linked_vpc_subnets"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(LinkedSubnetInfoBindingType), reflect.TypeOf([]LinkedSubnetInfo{})))
	fieldNameMap["linked_vpc_subnets"] = "LinkedVpcSubnets"
	fields["route_table_ids"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["route_table_ids"] = "RouteTableIds"
	fields["service_arn_role"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["service_arn_role"] = "ServiceArnRole"
	fields["traffic_group_eni_mappings"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(TrafficGroupEniMappingInfoBindingType), reflect.TypeOf([]TrafficGroupEniMappingInfo{})))
	fieldNameMap["traffic_group_eni_mappings"] = "TrafficGroupEniMappings"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.linked_vpc_info", fields, reflect.TypeOf(LinkedVpcInfo{}), fieldNameMap, validators)
}

func LinkedVpcManagedPrefixListSupportInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["aws_resource_share_info"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(AwsResourceShareInfoBindingType))
	fieldNameMap["aws_resource_share_info"] = "AwsResourceShareInfo"
	fields["managed_prefix_list_mode"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["managed_prefix_list_mode"] = "ManagedPrefixListMode"
	fields["managed_prefix_lists"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(LinkedVpcSharedManagedPrefixListInfoBindingType), reflect.TypeOf([]LinkedVpcSharedManagedPrefixListInfo{})))
	fieldNameMap["managed_prefix_lists"] = "ManagedPrefixLists"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.linked_vpc_managed_prefix_list_support_info", fields, reflect.TypeOf(LinkedVpcManagedPrefixListSupportInfo{}), fieldNameMap, validators)
}

func LinkedVpcSharedManagedPrefixListInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["address_family"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["address_family"] = "AddressFamily"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["in_use"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["in_use"] = "InUse"
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["programming_info"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(ManagedPrefixListProgrammingInfoBindingType))
	fieldNameMap["programming_info"] = "ProgrammingInfo"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.linked_vpc_shared_managed_prefix_list_info", fields, reflect.TypeOf(LinkedVpcSharedManagedPrefixListInfo{}), fieldNameMap, validators)
}

func LinkedVpcsListResultBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fields["result_count"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["result_count"] = "ResultCount"
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["sort_ascending"] = "SortAscending"
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sort_by"] = "SortBy"
	fields["results"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(LinkedVpcInfoBindingType), reflect.TypeOf([]LinkedVpcInfo{})))
	fieldNameMap["results"] = "Results"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.linked_vpcs_list_result", fields, reflect.TypeOf(LinkedVpcsListResult{}), fieldNameMap, validators)
}

func ListResultBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fields["result_count"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["result_count"] = "ResultCount"
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["sort_ascending"] = "SortAscending"
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.list_result", fields, reflect.TypeOf(ListResult{}), fieldNameMap, validators)
}

func LogicalRouterPortCountersBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["blocked_packets"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["blocked_packets"] = "BlockedPackets"
	fields["dad_dropped_packets"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["dad_dropped_packets"] = "DadDroppedPackets"
	fields["destination_unsupported_dropped_packets"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["destination_unsupported_dropped_packets"] = "DestinationUnsupportedDroppedPackets"
	fields["dropped_packets"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["dropped_packets"] = "DroppedPackets"
	fields["firewall_dropped_packets"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["firewall_dropped_packets"] = "FirewallDroppedPackets"
	fields["frag_needed_dropped_packets"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["frag_needed_dropped_packets"] = "FragNeededDroppedPackets"
	fields["ipsec_dropped_packets"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["ipsec_dropped_packets"] = "IpsecDroppedPackets"
	fields["ipsec_no_sa_dropped_packets"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["ipsec_no_sa_dropped_packets"] = "IpsecNoSaDroppedPackets"
	fields["ipsec_no_vti_dropped_packets"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["ipsec_no_vti_dropped_packets"] = "IpsecNoVtiDroppedPackets"
	fields["ipsec_pol_block_dropped_packets"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["ipsec_pol_block_dropped_packets"] = "IpsecPolBlockDroppedPackets"
	fields["ipsec_pol_err_dropped_packets"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["ipsec_pol_err_dropped_packets"] = "IpsecPolErrDroppedPackets"
	fields["ipv6_dropped_packets"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["ipv6_dropped_packets"] = "Ipv6DroppedPackets"
	fields["kni_dropped_packets"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["kni_dropped_packets"] = "KniDroppedPackets"
	fields["l4port_unsupported_dropped_packets"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["l4port_unsupported_dropped_packets"] = "L4portUnsupportedDroppedPackets"
	fields["malformed_dropped_packets"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["malformed_dropped_packets"] = "MalformedDroppedPackets"
	fields["no_arp_dropped_packets"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["no_arp_dropped_packets"] = "NoArpDroppedPackets"
	fields["no_linked_dropped_packets"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["no_linked_dropped_packets"] = "NoLinkedDroppedPackets"
	fields["no_mem_dropped_packets"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["no_mem_dropped_packets"] = "NoMemDroppedPackets"
	fields["no_receiver_dropped_packets"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["no_receiver_dropped_packets"] = "NoReceiverDroppedPackets"
	fields["no_route_dropped_packets"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["no_route_dropped_packets"] = "NoRouteDroppedPackets"
	fields["non_ip_dropped_packets"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["non_ip_dropped_packets"] = "NonIpDroppedPackets"
	fields["proto_unsupported_dropped_packets"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["proto_unsupported_dropped_packets"] = "ProtoUnsupportedDroppedPackets"
	fields["redirect_dropped_packets"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["redirect_dropped_packets"] = "RedirectDroppedPackets"
	fields["rpf_check_dropped_packets"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["rpf_check_dropped_packets"] = "RpfCheckDroppedPackets"
	fields["service_insert_dropped_packets"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["service_insert_dropped_packets"] = "ServiceInsertDroppedPackets"
	fields["total_bytes"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["total_bytes"] = "TotalBytes"
	fields["total_packets"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["total_packets"] = "TotalPackets"
	fields["ttl_exceeded_dropped_packets"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["ttl_exceeded_dropped_packets"] = "TtlExceededDroppedPackets"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.logical_router_port_counters", fields, reflect.TypeOf(LogicalRouterPortCounters{}), fieldNameMap, validators)
}

func LogicalRouterPortStatisticsSummaryBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["last_update_timestamp"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["last_update_timestamp"] = "LastUpdateTimestamp"
	fields["rx"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(LogicalRouterPortCountersBindingType))
	fieldNameMap["rx"] = "Rx"
	fields["tx"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(LogicalRouterPortCountersBindingType))
	fieldNameMap["tx"] = "Tx"
	fields["logical_router_port_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["logical_router_port_id"] = "LogicalRouterPortId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.logical_router_port_statistics_summary", fields, reflect.TypeOf(LogicalRouterPortStatisticsSummary{}), fieldNameMap, validators)
}

func ManagedPrefixListProgrammingInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["active_eni"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["active_eni"] = "ActiveEni"
	fields["expected_entries"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["expected_entries"] = "ExpectedEntries"
	fields["route_table_ids"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["route_table_ids"] = "RouteTableIds"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.managed_prefix_list_programming_info", fields, reflect.TypeOf(ManagedPrefixListProgrammingInfo{}), fieldNameMap, validators)
}

func ManagedPrefixListRuntimeInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["address_family"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["address_family"] = "AddressFamily"
	fields["arn"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["arn"] = "Arn"
	fields["entries"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["entries"] = "Entries"
	fields["issues"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["issues"] = "Issues"
	fields["max_entries"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["max_entries"] = "MaxEntries"
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["state"] = "State"
	fields["state_message"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["state_message"] = "StateMessage"
	fields["tags"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ManagedPrefixListTagBindingType), reflect.TypeOf([]ManagedPrefixListTag{})))
	fieldNameMap["tags"] = "Tags"
	fields["version"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["version"] = "Version"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = vapiBindings_.NewStringType()
	fieldNameMap["resource_type"] = "ResourceType"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.managed_prefix_list_runtime_info", fields, reflect.TypeOf(ManagedPrefixListRuntimeInfo{}), fieldNameMap, validators)
}

func ManagedPrefixListTagBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["key"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["key"] = "Key"
	fields["value"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["value"] = "Value"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.managed_prefix_list_tag", fields, reflect.TypeOf(ManagedPrefixListTag{}), fieldNameMap, validators)
}

func ManagedResourceBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_revision"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_create_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_create_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_last_modified_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["_protection"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_system_owned"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["description"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["display_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["tags"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.managed_resource", fields, reflect.TypeOf(ManagedResource{}), fieldNameMap, validators)
}

func ManagementGatewaySddcNetworkingStateInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["appliance_subnet"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["appliance_subnet"] = "ApplianceSubnet"
	fields["dns"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["dns"] = "Dns"
	fields["infra_subnets"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["infra_subnets"] = "InfraSubnets"
	fields["mgmt_subnets"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["mgmt_subnets"] = "MgmtSubnets"
	fields["total_firewall_rules"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["total_firewall_rules"] = "TotalFirewallRules"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.management_gateway_sddc_networking_state_info", fields, reflect.TypeOf(ManagementGatewaySddcNetworkingStateInfo{}), fieldNameMap, validators)
}

func MgmtServiceEntryBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["display_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["path"] = "Path"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.mgmt_service_entry", fields, reflect.TypeOf(MgmtServiceEntry{}), fieldNameMap, validators)
}

func MgmtVmInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_revision"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_create_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_create_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_last_modified_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["_protection"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_system_owned"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["description"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["display_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["tags"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["marked_for_delete"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["parent_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	fields["group_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["group_path"] = "GroupPath"
	fields["ip_attachment_pairs"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(IpAttachmentPairBindingType), reflect.TypeOf([]IpAttachmentPair{})))
	fieldNameMap["ip_attachment_pairs"] = "IpAttachmentPairs"
	fields["ips"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["ips"] = "Ips"
	fields["services"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(MgmtServiceEntryBindingType), reflect.TypeOf([]MgmtServiceEntry{})))
	fieldNameMap["services"] = "Services"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.mgmt_vm_info", fields, reflect.TypeOf(MgmtVmInfo{}), fieldNameMap, validators)
}

func MgmtVmsListResultBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fields["result_count"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["result_count"] = "ResultCount"
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["sort_ascending"] = "SortAscending"
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sort_by"] = "SortBy"
	fields["results"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(MgmtVmInfoBindingType), reflect.TypeOf([]MgmtVmInfo{})))
	fieldNameMap["results"] = "Results"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.mgmt_vms_list_result", fields, reflect.TypeOf(MgmtVmsListResult{}), fieldNameMap, validators)
}

func ModelInterfaceBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.model_interface", fields, reflect.TypeOf(ModelInterface{}), fieldNameMap, validators)
}

func OtherRegionConnectionInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["region"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["region"] = "Region"
	fields["tgw_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["tgw_id"] = "TgwId"
	fields["tgw_peering_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["tgw_peering_id"] = "TgwPeeringId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.other_region_connection_info", fields, reflect.TypeOf(OtherRegionConnectionInfo{}), fieldNameMap, validators)
}

func OutpostInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_revision"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_create_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_create_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_last_modified_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["_protection"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_system_owned"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["description"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["display_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["tags"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["marked_for_delete"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["parent_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	fields["lgw_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["lgw_id"] = "LgwId"
	fields["lgw_route_table_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["lgw_route_table_id"] = "LgwRouteTableId"
	fields["lgw_route_table_vpc_association_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["lgw_route_table_vpc_association_id"] = "LgwRouteTableVpcAssociationId"
	fields["lgw_route_table_vpc_association_state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["lgw_route_table_vpc_association_state"] = "LgwRouteTableVpcAssociationState"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.outpost_info", fields, reflect.TypeOf(OutpostInfo{}), fieldNameMap, validators)
}

func OutpostsConnectSddcNetworkingStateInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["advertised_routes"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["advertised_routes"] = "AdvertisedRoutes"
	fields["learned_routes"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["learned_routes"] = "LearnedRoutes"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.outposts_connect_sddc_networking_state_info", fields, reflect.TypeOf(OutpostsConnectSddcNetworkingStateInfo{}), fieldNameMap, validators)
}

func PolicyInterfaceStatisticsSummaryBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["last_update_timestamp"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["last_update_timestamp"] = "LastUpdateTimestamp"
	fields["rx"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(LogicalRouterPortCountersBindingType))
	fieldNameMap["rx"] = "Rx"
	fields["tx"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(LogicalRouterPortCountersBindingType))
	fieldNameMap["tx"] = "Tx"
	fields["logical_router_port_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["logical_router_port_id"] = "LogicalRouterPortId"
	fields["interface_policy_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["interface_policy_path"] = "InterfacePolicyPath"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.policy_interface_statistics_summary", fields, reflect.TypeOf(PolicyInterfaceStatisticsSummary{}), fieldNameMap, validators)
}

func PrefixListInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["url"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["url"] = "Url"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.prefix_list_info", fields, reflect.TypeOf(PrefixListInfo{}), fieldNameMap, validators)
}

func ProviderGatewayKeyValuePairsBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["key"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["key"] = "Key"
	fields["value"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ProviderObjectBindingType), reflect.TypeOf([]ProviderObject{})))
	fieldNameMap["value"] = "Value"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.provider_gateway_key_value_pairs", fields, reflect.TypeOf(ProviderGatewayKeyValuePairs{}), fieldNameMap, validators)
}

func ProviderObjectBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["display_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["type"] = "Type_"
	fields["url"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["url"] = "Url"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.provider_object", fields, reflect.TypeOf(ProviderObject{}), fieldNameMap, validators)
}

func PublicIpBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_revision"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_create_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_create_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_last_modified_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["_protection"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_system_owned"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["description"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["display_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["tags"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["marked_for_delete"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["parent_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	fields["ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["ip"] = "Ip"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.public_ip", fields, reflect.TypeOf(PublicIp{}), fieldNameMap, validators)
}

func PublicIpsListResultBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fields["result_count"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["result_count"] = "ResultCount"
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["sort_ascending"] = "SortAscending"
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sort_by"] = "SortBy"
	fields["results"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(PublicIpBindingType), reflect.TypeOf([]PublicIp{})))
	fieldNameMap["results"] = "Results"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.public_ips_list_result", fields, reflect.TypeOf(PublicIpsListResult{}), fieldNameMap, validators)
}

func RelatedApiErrorBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["details"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["details"] = "Details"
	fields["error_code"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["error_code"] = "ErrorCode"
	fields["error_data"] = vapiBindings_.NewOptionalType(vapiBindings_.NewDynamicStructType(nil))
	fieldNameMap["error_data"] = "ErrorData"
	fields["error_message"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["error_message"] = "ErrorMessage"
	fields["module_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["module_name"] = "ModuleName"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.related_api_error", fields, reflect.TypeOf(RelatedApiError{}), fieldNameMap, validators)
}

func ReservedCIDRBlockBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_revision"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_create_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_create_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_last_modified_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["_protection"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_system_owned"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["description"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["display_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["tags"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["marked_for_delete"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["parent_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	fields["cidr"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cidr"] = "Cidr"
	fields["prefixes"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ReservedCIDRBlockPrefixBindingType), reflect.TypeOf([]ReservedCIDRBlockPrefix{})))
	fieldNameMap["prefixes"] = "Prefixes"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.reserved_CIDR_block", fields, reflect.TypeOf(ReservedCIDRBlock{}), fieldNameMap, validators)
}

func ReservedCIDRBlockPrefixBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["description"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["prefix"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["prefix"] = "Prefix"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.reserved_CIDR_block_prefix", fields, reflect.TypeOf(ReservedCIDRBlockPrefix{}), fieldNameMap, validators)
}

func ReservedCIDRBlocksListResultBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fields["result_count"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["result_count"] = "ResultCount"
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["sort_ascending"] = "SortAscending"
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sort_by"] = "SortBy"
	fields["results"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ReservedCIDRBlockBindingType), reflect.TypeOf([]ReservedCIDRBlock{})))
	fieldNameMap["results"] = "Results"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.reserved_CIDR_blocks_list_result", fields, reflect.TypeOf(ReservedCIDRBlocksListResult{}), fieldNameMap, validators)
}

func ResourceBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.resource", fields, reflect.TypeOf(Resource{}), fieldNameMap, validators)
}

func ResourceLinkBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["action"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["action"] = "Action"
	fields["href"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["href"] = "Href"
	fields["rel"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["rel"] = "Rel"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.resource_link", fields, reflect.TypeOf(ResourceLink{}), fieldNameMap, validators)
}

func ResourceRuntimeInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = vapiBindings_.NewStringType()
	fieldNameMap["resource_type"] = "ResourceType"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.resource_runtime_info", fields, reflect.TypeOf(ResourceRuntimeInfo{}), fieldNameMap, validators)
}

func ResourceRuntimeInfoListResultBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fields["result_count"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["result_count"] = "ResultCount"
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["sort_ascending"] = "SortAscending"
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sort_by"] = "SortBy"
	fields["results"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewDynamicStructType([]vapiBindings_.ReferenceType{vapiBindings_.NewReferenceType(ResourceRuntimeInfoBindingType)}), reflect.TypeOf([]*vapiData_.StructValue{})))
	fieldNameMap["results"] = "Results"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.resource_runtime_info_list_result", fields, reflect.TypeOf(ResourceRuntimeInfoListResult{}), fieldNameMap, validators)
}

func RevisionedResourceBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_revision"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.revisioned_resource", fields, reflect.TypeOf(RevisionedResource{}), fieldNameMap, validators)
}

func RouteAggregationConfigBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_revision"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_create_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_create_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_last_modified_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["_protection"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_system_owned"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["description"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["display_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["tags"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["marked_for_delete"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["parent_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	fields["address_family"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["address_family"] = "AddressFamily"
	fields["prefixes"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(RouteAggregationConfigPrefixBindingType), reflect.TypeOf([]RouteAggregationConfigPrefix{})))
	fieldNameMap["prefixes"] = "Prefixes"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.route_aggregation_config", fields, reflect.TypeOf(RouteAggregationConfig{}), fieldNameMap, validators)
}

func RouteAggregationConfigListResultBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fields["result_count"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["result_count"] = "ResultCount"
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["sort_ascending"] = "SortAscending"
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sort_by"] = "SortBy"
	fields["results"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(RouteAggregationConfigBindingType), reflect.TypeOf([]RouteAggregationConfig{})))
	fieldNameMap["results"] = "Results"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.route_aggregation_config_list_result", fields, reflect.TypeOf(RouteAggregationConfigListResult{}), fieldNameMap, validators)
}

func RouteAggregationConfigPrefixBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["description"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["prefix"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["prefix"] = "Prefix"
	fields["summary_only"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["summary_only"] = "SummaryOnly"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.route_aggregation_config_prefix", fields, reflect.TypeOf(RouteAggregationConfigPrefix{}), fieldNameMap, validators)
}

func RouteConfigBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_revision"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_create_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_create_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_last_modified_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["_protection"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_system_owned"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["description"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["display_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["tags"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["marked_for_delete"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["parent_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	fields["aggregation_route_config"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(AggregationRouteConfigInfoBindingType))
	fieldNameMap["aggregation_route_config"] = "AggregationRouteConfig"
	fields["connectivity_endpoint_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["connectivity_endpoint_path"] = "ConnectivityEndpointPath"
	fields["disable_default_cgw_network"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["disable_default_cgw_network"] = "DisableDefaultCgwNetwork"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.route_config", fields, reflect.TypeOf(RouteConfig{}), fieldNameMap, validators)
}

func RouteConfigListResultBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fields["result_count"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["result_count"] = "ResultCount"
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["sort_ascending"] = "SortAscending"
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sort_by"] = "SortBy"
	fields["results"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(RouteConfigBindingType), reflect.TypeOf([]RouteConfig{})))
	fieldNameMap["results"] = "Results"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.route_config_list_result", fields, reflect.TypeOf(RouteConfigListResult{}), fieldNameMap, validators)
}

func RouteFilteringBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_revision"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_create_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_create_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_last_modified_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["_protection"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_system_owned"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["description"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["display_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["tags"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["marked_for_delete"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["parent_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	fields["endpoints"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["endpoints"] = "Endpoints"
	fields["state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["state"] = "State"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.route_filtering", fields, reflect.TypeOf(RouteFiltering{}), fieldNameMap, validators)
}

func SddcAdvertiseRouteConfigBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["aws_prefix_list"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(AwsPrefixInfoBindingType), reflect.TypeOf([]AwsPrefixInfo{})))
	fieldNameMap["aws_prefix_list"] = "AwsPrefixList"
	fields["other_region_connections"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(OtherRegionConnectionInfoBindingType), reflect.TypeOf([]OtherRegionConnectionInfo{})))
	fieldNameMap["other_region_connections"] = "OtherRegionConnections"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.sddc_advertise_route_config", fields, reflect.TypeOf(SddcAdvertiseRouteConfig{}), fieldNameMap, validators)
}

func SddcNetworkingStateInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cloud_provider"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(CloudProviderSddcNetworkingStateInfoBindingType), reflect.TypeOf([]CloudProviderSddcNetworkingStateInfo{})))
	fieldNameMap["cloud_provider"] = "CloudProvider"
	fields["compute_gateway"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ComputeGatewaySddcNetworkingStateInfoBindingType), reflect.TypeOf([]ComputeGatewaySddcNetworkingStateInfo{})))
	fieldNameMap["compute_gateway"] = "ComputeGateway"
	fields["direct_connect"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(DirectConnectSddcNetworkingStateInfoBindingType), reflect.TypeOf([]DirectConnectSddcNetworkingStateInfo{})))
	fieldNameMap["direct_connect"] = "DirectConnect"
	fields["error"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["error"] = "Error_"
	fields["last_sync_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["last_sync_time"] = "LastSyncTime"
	fields["management_gateway"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ManagementGatewaySddcNetworkingStateInfoBindingType), reflect.TypeOf([]ManagementGatewaySddcNetworkingStateInfo{})))
	fieldNameMap["management_gateway"] = "ManagementGateway"
	fields["outposts_connect"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(OutpostsConnectSddcNetworkingStateInfoBindingType), reflect.TypeOf([]OutpostsConnectSddcNetworkingStateInfo{})))
	fieldNameMap["outposts_connect"] = "OutpostsConnect"
	fields["provider_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["provider_name"] = "ProviderName"
	fields["sddc_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sddc_id"] = "SddcId"
	fields["sddc_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sddc_name"] = "SddcName"
	fields["sddc_region"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sddc_region"] = "SddcRegion"
	fields["transit_connect"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(TransitConnectSddcNetworkingStateInfoBindingType), reflect.TypeOf([]TransitConnectSddcNetworkingStateInfo{})))
	fieldNameMap["transit_connect"] = "TransitConnect"
	fields["vpn"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(VpnSddcNetworkingStateInfoBindingType), reflect.TypeOf([]VpnSddcNetworkingStateInfo{})))
	fieldNameMap["vpn"] = "Vpn"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.sddc_networking_state_info", fields, reflect.TypeOf(SddcNetworkingStateInfo{}), fieldNameMap, validators)
}

func SddcTransportNodeInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_revision"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_create_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_create_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_last_modified_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["_protection"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_system_owned"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["description"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["display_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["tags"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["marked_for_delete"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["parent_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	fields["appliance_configuration"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["appliance_configuration"] = "ApplianceConfiguration"
	fields["hardware_platform_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["hardware_platform_type"] = "HardwarePlatformType"
	fields["queue_num_per_port"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["queue_num_per_port"] = "QueueNumPerPort"
	fields["reservation_info"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(TransportNodeReservationInfoBindingType))
	fieldNameMap["reservation_info"] = "ReservationInfo"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.sddc_transport_node_info", fields, reflect.TypeOf(SddcTransportNodeInfo{}), fieldNameMap, validators)
}

func SddcUserConfigurationBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["all_uplink_interface_label"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["all_uplink_interface_label"] = "AllUplinkInterfaceLabel"
	fields["all_vpn_interface_label"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["all_vpn_interface_label"] = "AllVpnInterfaceLabel"
	fields["cgw_snat_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cgw_snat_ip"] = "CgwSnatIp"
	fields["compute_domain"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["compute_domain"] = "ComputeDomain"
	fields["compute_gateway"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["compute_gateway"] = "ComputeGateway"
	fields["domains"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ProviderObjectBindingType), reflect.TypeOf([]ProviderObject{})))
	fieldNameMap["domains"] = "Domains"
	fields["dx_interface_label"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["dx_interface_label"] = "DxInterfaceLabel"
	fields["enforcement_points"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ProviderObjectBindingType), reflect.TypeOf([]ProviderObject{})))
	fieldNameMap["enforcement_points"] = "EnforcementPoints"
	fields["external_connectivities"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ExternalConnectivityInfoBindingType), reflect.TypeOf([]ExternalConnectivityInfo{})))
	fieldNameMap["external_connectivities"] = "ExternalConnectivities"
	fields["gateways"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ProviderGatewayKeyValuePairsBindingType), reflect.TypeOf([]ProviderGatewayKeyValuePairs{})))
	fieldNameMap["gateways"] = "Gateways"
	fields["infra_subnets"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["infra_subnets"] = "InfraSubnets"
	fields["interfaces"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ModelInterfaceBindingType), reflect.TypeOf([]ModelInterface{})))
	fieldNameMap["interfaces"] = "Interfaces"
	fields["labels"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ProviderObjectBindingType), reflect.TypeOf([]ProviderObject{})))
	fieldNameMap["labels"] = "Labels"
	fields["linked_vpc_interface_label"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["linked_vpc_interface_label"] = "LinkedVpcInterfaceLabel"
	fields["management_domain"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["management_domain"] = "ManagementDomain"
	fields["management_gateway"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["management_gateway"] = "ManagementGateway"
	fields["management_gateway_default_dns_zone"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(ProviderObjectBindingType))
	fieldNameMap["management_gateway_default_dns_zone"] = "ManagementGatewayDefaultDnsZone"
	fields["management_gateway_label"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["management_gateway_label"] = "ManagementGatewayLabel"
	fields["mgmt_subnet"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["mgmt_subnet"] = "MgmtSubnet"
	fields["mgmt_subnets"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["mgmt_subnets"] = "MgmtSubnets"
	fields["mgw_snat_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["mgw_snat_ip"] = "MgwSnatIp"
	fields["provider_gateways"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ProviderGatewayKeyValuePairsBindingType), reflect.TypeOf([]ProviderGatewayKeyValuePairs{})))
	fieldNameMap["provider_gateways"] = "ProviderGateways"
	fields["provider_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["provider_name"] = "ProviderName"
	fields["provider_prefix_lists"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(PrefixListInfoBindingType), reflect.TypeOf([]PrefixListInfo{})))
	fieldNameMap["provider_prefix_lists"] = "ProviderPrefixLists"
	fields["public_interface_label"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["public_interface_label"] = "PublicInterfaceLabel"
	fields["route_aggregation_enabled"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["route_aggregation_enabled"] = "RouteAggregationEnabled"
	fields["sddc_infra_subnet"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["sddc_infra_subnet"] = "SddcInfraSubnet"
	fields["sddc_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sddc_name"] = "SddcName"
	fields["sddc_region"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sddc_region"] = "SddcRegion"
	fields["security_classification"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["security_classification"] = "SecurityClassification"
	fields["vpn_dx_ips"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["vpn_dx_ips"] = "VpnDxIps"
	fields["vpn_endpoints"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(VpnEndpointBindingType), reflect.TypeOf([]VpnEndpoint{})))
	fieldNameMap["vpn_endpoints"] = "VpnEndpoints"
	fields["vpn_internet_ips"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["vpn_internet_ips"] = "VpnInternetIps"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.sddc_user_configuration", fields, reflect.TypeOf(SddcUserConfiguration{}), fieldNameMap, validators)
}

func SelfResourceLinkBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["action"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["action"] = "Action"
	fields["href"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["href"] = "Href"
	fields["rel"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["rel"] = "Rel"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.self_resource_link", fields, reflect.TypeOf(SelfResourceLink{}), fieldNameMap, validators)
}

func TagBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["scope"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["scope"] = "Scope"
	fields["tag"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["tag"] = "Tag"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.tag", fields, reflect.TypeOf(Tag{}), fieldNameMap, validators)
}

func TrafficGroupBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_revision"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_create_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_create_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_last_modified_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["_protection"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_system_owned"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["description"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["display_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["tags"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["marked_for_delete"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["parent_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	fields["association_maps"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(TrafficGroupAssociationMapBindingType), reflect.TypeOf([]TrafficGroupAssociationMap{})))
	fieldNameMap["association_maps"] = "AssociationMaps"
	fields["resource_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["resource_id"] = "ResourceId"
	fields["state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["state"] = "State"
	fields["state_message"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["state_message"] = "StateMessage"
	fields["type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["type"] = "Type_"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.traffic_group", fields, reflect.TypeOf(TrafficGroup{}), fieldNameMap, validators)
}

func TrafficGroupAssociationMapBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_revision"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_create_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_create_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_last_modified_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["_protection"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_system_owned"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["description"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["display_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["tags"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["marked_for_delete"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["parent_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	fields["distributed_prefix_lists"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(DistributedPrefixListBindingType), reflect.TypeOf([]DistributedPrefixList{})))
	fieldNameMap["distributed_prefix_lists"] = "DistributedPrefixLists"
	fields["prefix_lists"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["prefix_lists"] = "PrefixLists"
	fields["scope"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["scope"] = "Scope"
	fields["state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["state"] = "State"
	fields["state_message"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["state_message"] = "StateMessage"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.traffic_group_association_map", fields, reflect.TypeOf(TrafficGroupAssociationMap{}), fieldNameMap, validators)
}

func TrafficGroupAssociationMapsListResultBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fields["result_count"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["result_count"] = "ResultCount"
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["sort_ascending"] = "SortAscending"
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sort_by"] = "SortBy"
	fields["results"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(TrafficGroupAssociationMapBindingType), reflect.TypeOf([]TrafficGroupAssociationMap{})))
	fieldNameMap["results"] = "Results"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.traffic_group_association_maps_list_result", fields, reflect.TypeOf(TrafficGroupAssociationMapsListResult{}), fieldNameMap, validators)
}

func TrafficGroupEniMappingInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["eni"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["eni"] = "Eni"
	fields["traffic_group_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["traffic_group_id"] = "TrafficGroupId"
	fields["traffic_group_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["traffic_group_name"] = "TrafficGroupName"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.traffic_group_eni_mapping_info", fields, reflect.TypeOf(TrafficGroupEniMappingInfo{}), fieldNameMap, validators)
}

func TrafficGroupsListResultBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fields["result_count"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["result_count"] = "ResultCount"
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["sort_ascending"] = "SortAscending"
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sort_by"] = "SortBy"
	fields["results"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(TrafficGroupBindingType), reflect.TypeOf([]TrafficGroup{})))
	fieldNameMap["results"] = "Results"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.traffic_groups_list_result", fields, reflect.TypeOf(TrafficGroupsListResult{}), fieldNameMap, validators)
}

func TransitConnectSddcNetworkingStateInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["advertised_routes"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["advertised_routes"] = "AdvertisedRoutes"
	fields["connected_sddc_group_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["connected_sddc_group_id"] = "ConnectedSddcGroupId"
	fields["connected_sddc_group_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["connected_sddc_group_name"] = "ConnectedSddcGroupName"
	fields["learned_routes"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["learned_routes"] = "LearnedRoutes"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.transit_connect_sddc_networking_state_info", fields, reflect.TypeOf(TransitConnectSddcNetworkingStateInfo{}), fieldNameMap, validators)
}

func TransportNodeCpuReservationBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["reservation_in_mhz"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["reservation_in_mhz"] = "ReservationInMhz"
	fields["reservation_in_shares"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["reservation_in_shares"] = "ReservationInShares"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.transport_node_cpu_reservation", fields, reflect.TypeOf(TransportNodeCpuReservation{}), fieldNameMap, validators)
}

func TransportNodeMemoryReservationBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["reservation_percentage"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["reservation_percentage"] = "ReservationPercentage"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.transport_node_memory_reservation", fields, reflect.TypeOf(TransportNodeMemoryReservation{}), fieldNameMap, validators)
}

func TransportNodeReservationInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cpu_reservation"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(TransportNodeCpuReservationBindingType))
	fieldNameMap["cpu_reservation"] = "CpuReservation"
	fields["memory_reservation"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(TransportNodeMemoryReservationBindingType))
	fieldNameMap["memory_reservation"] = "MemoryReservation"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.transport_node_reservation_info", fields, reflect.TypeOf(TransportNodeReservationInfo{}), fieldNameMap, validators)
}

func VMCAccountsBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["linked_vpc_account"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["linked_vpc_account"] = "LinkedVpcAccount"
	fields["shadow_account"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["shadow_account"] = "ShadowAccount"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.VMC_accounts", fields, reflect.TypeOf(VMCAccounts{}), fieldNameMap, validators)
}

func VifsListResultBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fields["result_count"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["result_count"] = "ResultCount"
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["sort_ascending"] = "SortAscending"
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sort_by"] = "SortBy"
	fields["results"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(VirtualInterfaceBindingType), reflect.TypeOf([]VirtualInterface{})))
	fieldNameMap["results"] = "Results"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.vifs_list_result", fields, reflect.TypeOf(VifsListResult{}), fieldNameMap, validators)
}

func VirtualInterfaceBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["bgp_status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["bgp_status"] = "BgpStatus"
	fields["direct_connect_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["direct_connect_id"] = "DirectConnectId"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["local_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["local_ip"] = "LocalIp"
	fields["mtu"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["mtu"] = "Mtu"
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["remote_asn"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["remote_asn"] = "RemoteAsn"
	fields["remote_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["remote_ip"] = "RemoteIp"
	fields["state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["state"] = "State"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.virtual_interface", fields, reflect.TypeOf(VirtualInterface{}), fieldNameMap, validators)
}

func VmcAppBaseInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_revision"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_create_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_create_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_last_modified_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["_protection"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_system_owned"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["description"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["display_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["tags"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["marked_for_delete"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["parent_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.vmc_app_base_info", fields, reflect.TypeOf(VmcAppBaseInfo{}), fieldNameMap, validators)
}

func VmcConsolidatedRealizedStatusBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_revision"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_create_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_create_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_last_modified_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["_protection"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_system_owned"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["description"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["display_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["tags"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["marked_for_delete"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["parent_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	fields["consolidated_status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(VmcConsolidatedStatusBindingType))
	fieldNameMap["consolidated_status"] = "ConsolidatedStatus"
	fields["consolidated_status_per_object"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(VmcConsolidatedStatusPerObjectBindingType), reflect.TypeOf([]VmcConsolidatedStatusPerObject{})))
	fieldNameMap["consolidated_status_per_object"] = "ConsolidatedStatusPerObject"
	fields["intent_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["intent_path"] = "IntentPath"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.vmc_consolidated_realized_status", fields, reflect.TypeOf(VmcConsolidatedRealizedStatus{}), fieldNameMap, validators)
}

func VmcConsolidatedStatusBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["consolidated_status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["consolidated_status"] = "ConsolidatedStatus"
	fields["status_message"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["status_message"] = "StatusMessage"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.vmc_consolidated_status", fields, reflect.TypeOf(VmcConsolidatedStatus{}), fieldNameMap, validators)
}

func VmcConsolidatedStatusPerObjectBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["consolidated_status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(VmcConsolidatedStatusBindingType))
	fieldNameMap["consolidated_status"] = "ConsolidatedStatus"
	fields["object_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["object_id"] = "ObjectId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.vmc_consolidated_status_per_object", fields, reflect.TypeOf(VmcConsolidatedStatusPerObject{}), fieldNameMap, validators)
}

func VmcFeatureFlagInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_revision"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_create_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_create_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_last_modified_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["_protection"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_system_owned"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["description"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["display_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["tags"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["marked_for_delete"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["parent_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	fields["internal_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["internal_name"] = "InternalName"
	fields["message"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["message"] = "Message"
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["state"] = "State"
	fields["unlicensed_message"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["unlicensed_message"] = "UnlicensedMessage"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.vmc_feature_flag_info", fields, reflect.TypeOf(VmcFeatureFlagInfo{}), fieldNameMap, validators)
}

func VmcFeatureFlagsBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["features"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(VmcFeatureFlagInfoBindingType), reflect.TypeOf([]VmcFeatureFlagInfo{})))
	fieldNameMap["features"] = "Features"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.vmc_feature_flags", fields, reflect.TypeOf(VmcFeatureFlags{}), fieldNameMap, validators)
}

func VmcInterfaceStatisticsBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["last_update_timestamp"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["last_update_timestamp"] = "LastUpdateTimestamp"
	fields["rx"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(LogicalRouterPortCountersBindingType))
	fieldNameMap["rx"] = "Rx"
	fields["tx"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(LogicalRouterPortCountersBindingType))
	fieldNameMap["tx"] = "Tx"
	fields["logical_router_port_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["logical_router_port_id"] = "LogicalRouterPortId"
	fields["interface_policy_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["interface_policy_path"] = "InterfacePolicyPath"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.vmc_interface_statistics", fields, reflect.TypeOf(VmcInterfaceStatistics{}), fieldNameMap, validators)
}

func VmcRealizedEntitiesBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["intent_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["intent_path"] = "IntentPath"
	fields["realized_entities"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(VmcRealizedEntityBindingType), reflect.TypeOf([]VmcRealizedEntity{})))
	fieldNameMap["realized_entities"] = "RealizedEntities"
	fields["realized_entities_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["realized_entities_id"] = "RealizedEntitiesId"
	fields["realized_state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["realized_state"] = "RealizedState"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.vmc_realized_entities", fields, reflect.TypeOf(VmcRealizedEntities{}), fieldNameMap, validators)
}

func VmcRealizedEntityBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["realization_api"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["realization_api"] = "RealizationApi"
	fields["realization_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["realization_id"] = "RealizationId"
	fields["realization_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["realization_path"] = "RealizationPath"
	fields["realized_entity_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["realized_entity_id"] = "RealizedEntityId"
	fields["realized_entity_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["realized_entity_type"] = "RealizedEntityType"
	fields["realized_state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["realized_state"] = "RealizedState"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.vmc_realized_entity", fields, reflect.TypeOf(VmcRealizedEntity{}), fieldNameMap, validators)
}

func VpnEndpointBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["interface_label"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["interface_label"] = "InterfaceLabel"
	fields["ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["ip"] = "Ip"
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["type"] = "Type_"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.vpn_endpoint", fields, reflect.TypeOf(VpnEndpoint{}), fieldNameMap, validators)
}

func VpnSddcNetworkingStateInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["connectivity_deactivated"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["connectivity_deactivated"] = "ConnectivityDeactivated"
	fields["connectivity_down"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["connectivity_down"] = "ConnectivityDown"
	fields["connectivity_up"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["connectivity_up"] = "ConnectivityUp"
	fields["public_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["public_ip"] = "PublicIp"
	fields["total_connections"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["total_connections"] = "TotalConnections"
	fields["vpn_down_connections"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["vpn_down_connections"] = "VpnDownConnections"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.nsx_vmc_app.model.vpn_sddc_networking_state_info", fields, reflect.TypeOf(VpnSddcNetworkingStateInfo{}), fieldNameMap, validators)
}
