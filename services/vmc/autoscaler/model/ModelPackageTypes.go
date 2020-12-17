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


type ActionMessage struct {
    // Optional message for action
	Message *string
}

func (s ActionMessage) GetType__() bindings.BindingType {
	return ActionMessageBindingType()
}

func (s ActionMessage) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ActionMessage._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Advertised BGP route
type AdvertisedRoute struct {
    // The route that is advertised to on-premise datacenter via Direct Connect format: ipv4-cidr-block
	Ipv4Cidr string
    // Possible values are: 
    //
    // * AdvertisedRoute#AdvertisedRoute_ADVERTISEMENT_STATE_SUCCESS
    // * AdvertisedRoute#AdvertisedRoute_ADVERTISEMENT_STATE_FAILED
    //
    //  State of advertisement
	AdvertisementState string
}
const AdvertisedRoute_ADVERTISEMENT_STATE_SUCCESS = "SUCCESS"
const AdvertisedRoute_ADVERTISEMENT_STATE_FAILED = "FAILED"

func (s AdvertisedRoute) GetType__() bindings.BindingType {
	return AdvertisedRouteBindingType()
}

func (s AdvertisedRoute) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for AdvertisedRoute._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Detailed information about an API Error
type ApiError struct {
    // The module name where the error occurred
	ModuleName *string
    // A description of the error
	ErrorMessage *string
    // A numeric error code format: int64
	ErrorCode *int64
    // Further details about the error
	Details *string
    // Additional data about the error
	ErrorData *data.StructValue
    // Other errors related to this error
	RelatedErrors []RelatedApiError
}

func (s ApiError) GetType__() bindings.BindingType {
	return ApiErrorBindingType()
}

func (s ApiError) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ApiError._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Base abstract associated Group connection infomation for the local SDDC.
type AssociatedBaseGroupConnectionInfo struct {
    // Link to this resource
	Self *SelfResourceLink
    // The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
    // Schema for this resource
	Schema *string
    // The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected. format: int32
	Revision *int64
    // Indicates system owned resource
	SystemOwned *bool
    // Defaults to ID if not set
	DisplayName *string
    // SDDC Group description
	Description *string
    // Opaque identifiers meaningful to the API user
	Tags []Tag
    // ID of the user who created this resource
	CreateUser *string
    // Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super user and can modify it, but only when providing the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this entity.
	Protection *string
    // Timestamp of resource creation format: int64
	CreateTime *int64
    // Timestamp of last modification format: int64
	LastModifiedTime *int64
    // ID of the user who last modified this resource
	LastModifiedUser *string
    // SDDC Group ID
	Id string
    // Possible values are: 
    //
    // * AssociatedBaseGroupConnectionInfo#AssociatedBaseGroupConnectionInfo_RESOURCE_TYPE_ASSOCIATEDTGWGROUPCONNECTIONINFO
    //
    //  Group connection type
	ResourceType string
    // Absolute path of this object
	Path *string
    // Path relative from its parent
	RelativePath *string
    // Path of its parent
	ParentPath *string
    // marked for delete identifier
	MarkedForDelete *bool
    // SDDC Group name
	Name *string
}
// Identifier denoting this class, when it is used in polymorphic context. 
//
// This value should be assigned to the property which is used to discriminate the actual type used in the polymorphic context.
const AssociatedBaseGroupConnectionInfo__TYPE_IDENTIFIER = "AssociatedBaseGroupConnectionInfo"
const AssociatedBaseGroupConnectionInfo_RESOURCE_TYPE_ASSOCIATEDTGWGROUPCONNECTIONINFO = "AssociatedTgwGroupConnectionInfo"

func (s AssociatedBaseGroupConnectionInfo) GetType__() bindings.BindingType {
	return AssociatedBaseGroupConnectionInfoBindingType()
}

func (s AssociatedBaseGroupConnectionInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for AssociatedBaseGroupConnectionInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Associated Group connection list result
type AssociatedGroupConnectionInfosListResult struct {
    // Link to this resource
	Self *SelfResourceLink
    // The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
    // Schema for this resource
	Schema *string
    // Opaque cursor to be used for getting next page of records (supplied by current result page)
	Cursor *string
    // If true, results are sorted in ascending order
	SortAscending *bool
    // Field by which records are sorted
	SortBy *string
    // Count of results found (across all pages), set only on first page format: int64
	ResultCount *int64
	Results []*data.StructValue
}

func (s AssociatedGroupConnectionInfosListResult) GetType__() bindings.BindingType {
	return AssociatedGroupConnectionInfosListResultBindingType()
}

func (s AssociatedGroupConnectionInfosListResult) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for AssociatedGroupConnectionInfosListResult._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Associated Group connection infomation for the local SDDC by using AWS TGW as a connector.
type AssociatedTgwGroupConnectionInfo struct {
    // TGW ID for the local SDDC in the Group
	TgwId string
    // TGW external route table ID used for external customers VPCs association
	ExternalRouteTableId string
    // Possible values are: 
    //
    // * AssociatedTgwGroupConnectionInfo#AssociatedTgwGroupConnectionInfo_STATE_CONNECTED
    // * AssociatedTgwGroupConnectionInfo#AssociatedTgwGroupConnectionInfo_STATE_DISCONNECTED
    //
    //  The TGW attachment state of the SDDC in the Group
	State *string
    // TGW SDDCs route table ID used for SDDCs association
	SddcsRouteTableId string
    // TGW attachment ID for the local SDDC in the Group
	TgwAttachmentId string
    // Link to this resource
	Self *SelfResourceLink
    // The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
    // Schema for this resource
	Schema *string
    // The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected. format: int32
	Revision *int64
    // Indicates system owned resource
	SystemOwned *bool
    // Defaults to ID if not set
	DisplayName *string
    // SDDC Group description
	Description *string
    // Opaque identifiers meaningful to the API user
	Tags []Tag
    // ID of the user who created this resource
	CreateUser *string
    // Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super user and can modify it, but only when providing the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this entity.
	Protection *string
    // Timestamp of resource creation format: int64
	CreateTime *int64
    // Timestamp of last modification format: int64
	LastModifiedTime *int64
    // ID of the user who last modified this resource
	LastModifiedUser *string
    // SDDC Group ID
	Id string
    // Possible values are: 
    //
    // * AssociatedBaseGroupConnectionInfo#AssociatedBaseGroupConnectionInfo_RESOURCE_TYPE_ASSOCIATEDTGWGROUPCONNECTIONINFO
    //
    //  Group connection type
	ResourceType string
    // Absolute path of this object
	Path *string
    // Path relative from its parent
	RelativePath *string
    // Path of its parent
	ParentPath *string
    // marked for delete identifier
	MarkedForDelete *bool
    // SDDC Group name
	Name *string
}
// Identifier denoting this class, when it is used in polymorphic context. 
//
// This value should be assigned to the property which is used to discriminate the actual type used in the polymorphic context.
const AssociatedTgwGroupConnectionInfo__TYPE_IDENTIFIER = "AssociatedTgwGroupConnectionInfo"
const AssociatedTgwGroupConnectionInfo_STATE_CONNECTED = "CONNECTED"
const AssociatedTgwGroupConnectionInfo_STATE_DISCONNECTED = "DISCONNECTED"

func (s AssociatedTgwGroupConnectionInfo) GetType__() bindings.BindingType {
	return AssociatedTgwGroupConnectionInfoBindingType()
}

func (s AssociatedTgwGroupConnectionInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for AssociatedTgwGroupConnectionInfo._GetDataValue method - %s",
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


// Advertised bgp routes
type BGPAdvertisedRoutes struct {
    // Number of routes failed to advertise format: int32
	FailedAdvertisedRoutes *int64
    // Routes advertised to on-premise datacenter via Direct Connect
	AdvertisedRoutes []AdvertisedRoute
}

func (s BGPAdvertisedRoutes) GetType__() bindings.BindingType {
	return BGPAdvertisedRoutesBindingType()
}

func (s BGPAdvertisedRoutes) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for BGPAdvertisedRoutes._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Learned bgp routes
type BGPLearnedRoutes struct {
    // The route that is learned from BGP via Direct Connect format: ipv4-cidr-block
	Ipv4Cidr []string
}

func (s BGPLearnedRoutes) GetType__() bindings.BindingType {
	return BGPLearnedRoutesBindingType()
}

func (s BGPLearnedRoutes) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for BGPLearnedRoutes._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// A list of status of 'Enabled/Disabled' for a service connected to a linked vpc
type ConnectedServiceListResult struct {
    // Link to this resource
	Self *SelfResourceLink
    // The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
    // Schema for this resource
	Schema *string
    // Opaque cursor to be used for getting next page of records (supplied by current result page)
	Cursor *string
    // If true, results are sorted in ascending order
	SortAscending *bool
    // Field by which records are sorted
	SortBy *string
    // Count of results found (across all pages), set only on first page format: int64
	ResultCount *int64
    // Connected service status list
	Results []ConnectedServiceStatus
}

func (s ConnectedServiceListResult) GetType__() bindings.BindingType {
	return ConnectedServiceListResultBindingType()
}

func (s ConnectedServiceListResult) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ConnectedServiceListResult._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Status of 'Enabled/Disabled' for a service connected to a linked vpc
type ConnectedServiceStatus struct {
    // Link to this resource
	Self *SelfResourceLink
    // The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
    // Schema for this resource
	Schema *string
    // The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected. format: int32
	Revision *int64
    // Indicates system owned resource
	SystemOwned *bool
    // Defaults to ID if not set
	DisplayName *string
    // Description of this resource
	Description *string
    // Opaque identifiers meaningful to the API user
	Tags []Tag
    // ID of the user who created this resource
	CreateUser *string
    // Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super user and can modify it, but only when providing the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this entity.
	Protection *string
    // Timestamp of resource creation format: int64
	CreateTime *int64
    // Timestamp of last modification format: int64
	LastModifiedTime *int64
    // ID of the user who last modified this resource
	LastModifiedUser *string
    // Unique identifier of this resource
	Id *string
    // The type of this resource.
	ResourceType *string
    // Absolute path of this object
	Path *string
    // Path relative from its parent
	RelativePath *string
    // Path of its parent
	ParentPath *string
    // marked for delete identifier
	MarkedForDelete *bool
    // status of service
	Enabled *bool
    // service name
	Name *string
}

func (s ConnectedServiceStatus) GetType__() bindings.BindingType {
	return ConnectedServiceStatusBindingType()
}

func (s ConnectedServiceStatus) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ConnectedServiceStatus._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Base type for CSV result.
type CsvListResult struct {
    // File name set by HTTP server if API returns CSV result as a file.
	FileName *string
}

func (s CsvListResult) GetType__() bindings.BindingType {
	return CsvListResultBindingType()
}

func (s CsvListResult) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for CsvListResult._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Base type for CSV records.
type CsvRecord struct {
}

func (s CsvRecord) GetType__() bindings.BindingType {
	return CsvRecordBindingType()
}

func (s CsvRecord) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for CsvRecord._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Direct Connect BGP related information
type DirectConnectBgpInfo struct {
    // Link to this resource
	Self *SelfResourceLink
    // The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
    // Schema for this resource
	Schema *string
    // The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected. format: int32
	Revision *int64
    // Indicates system owned resource
	SystemOwned *bool
    // Defaults to ID if not set
	DisplayName *string
    // Description of this resource
	Description *string
    // Opaque identifiers meaningful to the API user
	Tags []Tag
    // ID of the user who created this resource
	CreateUser *string
    // Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super user and can modify it, but only when providing the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this entity.
	Protection *string
    // Timestamp of resource creation format: int64
	CreateTime *int64
    // Timestamp of last modification format: int64
	LastModifiedTime *int64
    // ID of the user who last modified this resource
	LastModifiedUser *string
    // Unique identifier of this resource
	Id *string
    // The type of this resource.
	ResourceType *string
    // Absolute path of this object
	Path *string
    // Path relative from its parent
	RelativePath *string
    // Path of its parent
	ParentPath *string
    // marked for delete identifier
	MarkedForDelete *bool
    // The ASN paired with the VGW attached to the VPC. AWS allowed private BGP ASN range - [64512, 65534] and [4200000000, 4294967294]. If omitted in the payload, BGP ASN will not be modified.
	LocalAsNum *string
    // Possible values are: 
    //
    // * DirectConnectBgpInfo#DirectConnectBgpInfo_ROUTE_PREFERENCE_DIRECT_CONNECT_PREFERRED_OVER_VPN
    // * DirectConnectBgpInfo#DirectConnectBgpInfo_ROUTE_PREFERENCE_VPN_PREFERRED_OVER_DIRECT_CONNECT
    //
    //  Direct connect route preference over VPN routes. If omitted in the payload, route preference will not be modified.
	RoutePreference *string
    // Maximum transmission unit allowed by the VIF format: int32
	Mtu *int64
}
const DirectConnectBgpInfo_ROUTE_PREFERENCE_DIRECT_CONNECT_PREFERRED_OVER_VPN = "DIRECT_CONNECT_PREFERRED_OVER_VPN"
const DirectConnectBgpInfo_ROUTE_PREFERENCE_VPN_PREFERRED_OVER_DIRECT_CONNECT = "VPN_PREFERRED_OVER_DIRECT_CONNECT"

func (s DirectConnectBgpInfo) GetType__() bindings.BindingType {
	return DirectConnectBgpInfoBindingType()
}

func (s DirectConnectBgpInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for DirectConnectBgpInfo._GetDataValue method - %s",
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


type EdrsPolicyOptionsOverrides struct {
    // The number of hosts which will be added to the cluster during the scale out operation.
	ScaleUpHostIncrement *int64
}

func (s EdrsPolicyOptionsOverrides) GetType__() bindings.BindingType {
	return EdrsPolicyOptionsOverridesBindingType()
}

func (s EdrsPolicyOptionsOverrides) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for EdrsPolicyOptionsOverrides._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
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

func (s EdrsPolicySpec) GetType__() bindings.BindingType {
	return EdrsPolicySpecBindingType()
}

func (s EdrsPolicySpec) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for EdrsPolicySpec._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type EdrsProvisioningSpec struct {
	CurrentEdrsPolicy *EdrsPolicy
    // List of spec for all EDRS policies.
	Policies []EdrsPolicySpec
}

func (s EdrsProvisioningSpec) GetType__() bindings.BindingType {
	return EdrsProvisioningSpecBindingType()
}

func (s EdrsProvisioningSpec) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for EdrsProvisioningSpec._GetDataValue method - %s",
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


// External Connectivity configuration for north-south traffic. For eg., in AWS case, this would refer to Direct Connect config. For Dimension, this would refer to the config at Upstream Intranet interface to Tor.
type ExternalConnectivityConfig struct {
    // Link to this resource
	Self *SelfResourceLink
    // The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
    // Schema for this resource
	Schema *string
    // The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected. format: int32
	Revision *int64
    // Indicates system owned resource
	SystemOwned *bool
    // Defaults to ID if not set
	DisplayName *string
    // Description of this resource
	Description *string
    // Opaque identifiers meaningful to the API user
	Tags []Tag
    // ID of the user who created this resource
	CreateUser *string
    // Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super user and can modify it, but only when providing the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this entity.
	Protection *string
    // Timestamp of resource creation format: int64
	CreateTime *int64
    // Timestamp of last modification format: int64
	LastModifiedTime *int64
    // ID of the user who last modified this resource
	LastModifiedUser *string
    // Unique identifier of this resource
	Id *string
    // The type of this resource.
	ResourceType *string
    // Absolute path of this object
	Path *string
    // Path relative from its parent
	RelativePath *string
    // Path of its parent
	ParentPath *string
    // marked for delete identifier
	MarkedForDelete *bool
    // Uplink MTU of direct connect, sddc-grouping and outposts traffic in edge tier-0 router port. format: int32
	IntranetMtu int64
    // Uplink MTU of connected VPC traffic in edge tier-0 router port. format: int32
	ServicesMtu *int64
    // Uplink MTU of internet traffic in edge tier-0 router port. format: int32
	InternetMtu *int64
}

func (s ExternalConnectivityConfig) GetType__() bindings.BindingType {
	return ExternalConnectivityConfigBindingType()
}

func (s ExternalConnectivityConfig) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ExternalConnectivityConfig._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// External SDDC connectivity
type ExternalSddcConnectivity struct {
    // Possible values are: 
    //
    // * ExternalSddcConnectivity#ExternalSddcConnectivity_STATUS_SUCCEEDED
    // * ExternalSddcConnectivity#ExternalSddcConnectivity_STATUS_FAILED
    //
    //  The status of the route for the connectivity
	Status *string
    // Possible values are: 
    //
    // * ExternalSddcConnectivity#ExternalSddcConnectivity_CONNECTIVITY_TYPE_DIRECT_CONNECT
    // * ExternalSddcConnectivity#ExternalSddcConnectivity_CONNECTIVITY_TYPE_DEPLOYMENT_CONNECTIVITY_GROUP
    //
    //  The external SDDC connectivity type is used by the SDDC for the L3 connectivity. DIRECT_CONNECT means that the external SDDC connectivity is through AWS Direct Connect. DEPLOYMENT_CONNECTIVITY_GROUP means that the external SDDC connectivity is through AWS TGW.
	ConnectivityType string
    // Possible values are: 
    //
    // * ExternalSddcConnectivity#ExternalSddcConnectivity_ROUTE_TYPE_PROPAGATED
    // * ExternalSddcConnectivity#ExternalSddcConnectivity_ROUTE_TYPE_STATIC
    //
    //  The type of the route for the connectivity
	RouteType *string
    // The error message if the status is FAILED, the optional warning message if the status is SUCCEEDED.
	StatusMessage *string
    // The source of the route for the connectivity
	Source *string
}
const ExternalSddcConnectivity_STATUS_SUCCEEDED = "SUCCEEDED"
const ExternalSddcConnectivity_STATUS_FAILED = "FAILED"
const ExternalSddcConnectivity_CONNECTIVITY_TYPE_DIRECT_CONNECT = "DIRECT_CONNECT"
const ExternalSddcConnectivity_CONNECTIVITY_TYPE_DEPLOYMENT_CONNECTIVITY_GROUP = "DEPLOYMENT_CONNECTIVITY_GROUP"
const ExternalSddcConnectivity_ROUTE_TYPE_PROPAGATED = "PROPAGATED"
const ExternalSddcConnectivity_ROUTE_TYPE_STATIC = "STATIC"

func (s ExternalSddcConnectivity) GetType__() bindings.BindingType {
	return ExternalSddcConnectivityBindingType()
}

func (s ExternalSddcConnectivity) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ExternalSddcConnectivity._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// External SDDC route
type ExternalSddcRoute struct {
    // Link to this resource
	Self *SelfResourceLink
    // The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
    // Schema for this resource
	Schema *string
    // The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected. format: int32
	Revision *int64
    // Indicates system owned resource
	SystemOwned *bool
    // Defaults to ID if not set
	DisplayName *string
    // Description of this resource
	Description *string
    // Opaque identifiers meaningful to the API user
	Tags []Tag
    // ID of the user who created this resource
	CreateUser *string
    // Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super user and can modify it, but only when providing the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this entity.
	Protection *string
    // Timestamp of resource creation format: int64
	CreateTime *int64
    // Timestamp of last modification format: int64
	LastModifiedTime *int64
    // ID of the user who last modified this resource
	LastModifiedUser *string
    // Unique identifier of this resource
	Id *string
    // The type of this resource.
	ResourceType *string
    // Absolute path of this object
	Path *string
    // Path relative from its parent
	RelativePath *string
    // Path of its parent
	ParentPath *string
    // marked for delete identifier
	MarkedForDelete *bool
    // Destination IP CIDR Block format: ipv4-cidr-block
	Destination string
    // The route used for what kind of connectivities
	Connectivities []ExternalSddcConnectivity
}

func (s ExternalSddcRoute) GetType__() bindings.BindingType {
	return ExternalSddcRouteBindingType()
}

func (s ExternalSddcRoute) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ExternalSddcRoute._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// CSV record for External SDDC route
type ExternalSddcRouteCsvRecord struct {
    // Destination IP CIDR Block format: ipv4-cidr-block
	Destination string
    // The connectivity datails contains status of route, source of the route, connectivity type
	ConnectivityDetails string
}

func (s ExternalSddcRouteCsvRecord) GetType__() bindings.BindingType {
	return ExternalSddcRouteCsvRecordBindingType()
}

func (s ExternalSddcRouteCsvRecord) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ExternalSddcRouteCsvRecord._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// External SDDC routes list result
type ExternalSddcRoutesListResult struct {
    // Link to this resource
	Self *SelfResourceLink
    // The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
    // Schema for this resource
	Schema *string
    // Opaque cursor to be used for getting next page of records (supplied by current result page)
	Cursor *string
    // If true, results are sorted in ascending order
	SortAscending *bool
    // Field by which records are sorted
	SortBy *string
    // Count of results found (across all pages), set only on first page format: int64
	ResultCount *int64
	Routes []ExternalSddcRoute
}

func (s ExternalSddcRoutesListResult) GetType__() bindings.BindingType {
	return ExternalSddcRoutesListResultBindingType()
}

func (s ExternalSddcRoutesListResult) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ExternalSddcRoutesListResult._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// External SDDC routes list result in CSV format
type ExternalSddcRoutesListResultInCsvFormat struct {
    // File name set by HTTP server if API returns CSV result as a file.
	FileName *string
	Results []ExternalSddcRouteCsvRecord
}

func (s ExternalSddcRoutesListResultInCsvFormat) GetType__() bindings.BindingType {
	return ExternalSddcRoutesListResultInCsvFormatBindingType()
}

func (s ExternalSddcRoutesListResultInCsvFormat) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ExternalSddcRoutesListResultInCsvFormat._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// A list of fields to include in query results
type IncludedFieldsParameters struct {
    // Comma separated list of fields that should be included in query result
	IncludedFields *string
}

func (s IncludedFieldsParameters) GetType__() bindings.BindingType {
	return IncludedFieldsParametersBindingType()
}

func (s IncludedFieldsParameters) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for IncludedFieldsParameters._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type IpAttachmentPair struct {
    // Management VM IP Address format: ipv4
	Ip string
    // Attachment id which maps to management VM IP
	AttachmentId string
}

func (s IpAttachmentPair) GetType__() bindings.BindingType {
	return IpAttachmentPairBindingType()
}

func (s IpAttachmentPair) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for IpAttachmentPair._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Infromation related to a subnet where linked ENIs were created.
type LinkedSubnetInfo struct {
    // Linked subnet CIDR format: ipv4-cidr-block
	Cidr string
    // Linked subnet identifier
	Id string
    // Linked subnet availability zone
	AvailabilityZone string
}

func (s LinkedSubnetInfo) GetType__() bindings.BindingType {
	return LinkedSubnetInfoBindingType()
}

func (s LinkedSubnetInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for LinkedSubnetInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Linked VPC info
type LinkedVpcInfo struct {
    // Link to this resource
	Self *SelfResourceLink
    // The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
    // Schema for this resource
	Schema *string
    // The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected. format: int32
	Revision *int64
    // Indicates system owned resource
	SystemOwned *bool
    // Defaults to ID if not set
	DisplayName *string
    // Description of this resource
	Description *string
    // Opaque identifiers meaningful to the API user
	Tags []Tag
    // ID of the user who created this resource
	CreateUser *string
    // Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super user and can modify it, but only when providing the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this entity.
	Protection *string
    // Timestamp of resource creation format: int64
	CreateTime *int64
    // Timestamp of last modification format: int64
	LastModifiedTime *int64
    // ID of the user who last modified this resource
	LastModifiedUser *string
    // Unique identifier of this resource
	Id *string
    // The type of this resource.
	ResourceType *string
    // Absolute path of this object
	Path *string
    // Path relative from its parent
	RelativePath *string
    // Path of its parent
	ParentPath *string
    // marked for delete identifier
	MarkedForDelete *bool
    // ARN role for linked VPC operations
	ArnRole string
    // Active network interface used for linked vpc traffic
	ActiveEni *string
    // Linked VPC identifier
	LinkedVpcId *string
    // Linked VPC CIDRs format: ipv4-cidr-block
	LinkedVpcAddresses []string
    // The IPs of linked VPC NAT rule for service access. format: ipv4
	LinkedVpcNatIps []string
    // Linked VPC account number
	LinkedAccount string
    // The identifiers of route tables to be dynamically updated with SDDC networks
	RouteTableIds []string
    // service ARN role
	ServiceArnRole *string
    // Infromation related to the subnets where linked ENIs were created.
	LinkedVpcSubnets []LinkedSubnetInfo
    // External identifier for ARN role
	ExternalId string
}

func (s LinkedVpcInfo) GetType__() bindings.BindingType {
	return LinkedVpcInfoBindingType()
}

func (s LinkedVpcInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for LinkedVpcInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Linked VPC list query result
type LinkedVpcsListResult struct {
    // Link to this resource
	Self *SelfResourceLink
    // The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
    // Schema for this resource
	Schema *string
    // Opaque cursor to be used for getting next page of records (supplied by current result page)
	Cursor *string
    // If true, results are sorted in ascending order
	SortAscending *bool
    // Field by which records are sorted
	SortBy *string
    // Count of results found (across all pages), set only on first page format: int64
	ResultCount *int64
    // Linked VPCs list
	Results []LinkedVpcInfo
}

func (s LinkedVpcsListResult) GetType__() bindings.BindingType {
	return LinkedVpcsListResultBindingType()
}

func (s LinkedVpcsListResult) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for LinkedVpcsListResult._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Base class for list results from collections
type ListResult struct {
    // Link to this resource
	Self *SelfResourceLink
    // The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
    // Schema for this resource
	Schema *string
    // Opaque cursor to be used for getting next page of records (supplied by current result page)
	Cursor *string
    // If true, results are sorted in ascending order
	SortAscending *bool
    // Field by which records are sorted
	SortBy *string
    // Count of results found (across all pages), set only on first page format: int64
	ResultCount *int64
}

func (s ListResult) GetType__() bindings.BindingType {
	return ListResultBindingType()
}

func (s ListResult) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ListResult._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Base type for resources that are managed by API clients
type ManagedResource struct {
    // Link to this resource
	Self *SelfResourceLink
    // The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
    // Schema for this resource
	Schema *string
    // The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected. format: int32
	Revision *int64
    // Indicates system owned resource
	SystemOwned *bool
    // Defaults to ID if not set
	DisplayName *string
    // Description of this resource
	Description *string
    // Opaque identifiers meaningful to the API user
	Tags []Tag
    // ID of the user who created this resource
	CreateUser *string
    // Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super user and can modify it, but only when providing the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this entity.
	Protection *string
    // Timestamp of resource creation format: int64
	CreateTime *int64
    // Timestamp of last modification format: int64
	LastModifiedTime *int64
    // ID of the user who last modified this resource
	LastModifiedUser *string
    // Unique identifier of this resource
	Id *string
    // The type of this resource.
	ResourceType *string
}

func (s ManagedResource) GetType__() bindings.BindingType {
	return ManagedResourceBindingType()
}

func (s ManagedResource) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ManagedResource._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// A service entry describes the detail of a network service.
type MgmtServiceEntry struct {
    // Service path should refer to a valid service in the system. Service can be system defined or user defined.
	Path *string
    // Display name for this service
	DisplayName *string
}

func (s MgmtServiceEntry) GetType__() bindings.BindingType {
	return MgmtServiceEntryBindingType()
}

func (s MgmtServiceEntry) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for MgmtServiceEntry._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Management VM access information
type MgmtVmInfo struct {
    // Link to this resource
	Self *SelfResourceLink
    // The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
    // Schema for this resource
	Schema *string
    // The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected. format: int32
	Revision *int64
    // Indicates system owned resource
	SystemOwned *bool
    // Management VM name
	DisplayName *string
    // Description of this resource
	Description *string
    // Opaque identifiers meaningful to the API user
	Tags []Tag
    // ID of the user who created this resource
	CreateUser *string
    // Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super user and can modify it, but only when providing the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this entity.
	Protection *string
    // Timestamp of resource creation format: int64
	CreateTime *int64
    // Timestamp of last modification format: int64
	LastModifiedTime *int64
    // ID of the user who last modified this resource
	LastModifiedUser *string
    // Management VM identifier
	Id *string
    // The type of this resource.
	ResourceType *string
    // Absolute path of this object
	Path *string
    // Path relative from its parent
	RelativePath *string
    // Path of its parent
	ParentPath *string
    // marked for delete identifier
	MarkedForDelete *bool
    // Local IPs of a management VM format: address-or-block-or-range
	Ips []string
    // Details services path and display name.
	Services []MgmtServiceEntry
    // For each management VM, a dedicated policy group will be created. This property will reflect its group path.
	GroupPath *string
    // IP address and attachment id pairs for tagging managment VM
	IpAttachmentPairs []IpAttachmentPair
}

func (s MgmtVmInfo) GetType__() bindings.BindingType {
	return MgmtVmInfoBindingType()
}

func (s MgmtVmInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for MgmtVmInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Management VM list query result
type MgmtVmsListResult struct {
    // Link to this resource
	Self *SelfResourceLink
    // The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
    // Schema for this resource
	Schema *string
    // Opaque cursor to be used for getting next page of records (supplied by current result page)
	Cursor *string
    // If true, results are sorted in ascending order
	SortAscending *bool
    // Field by which records are sorted
	SortBy *string
    // Count of results found (across all pages), set only on first page format: int64
	ResultCount *int64
    // Management VMs list
	Results []MgmtVmInfo
}

func (s MgmtVmsListResult) GetType__() bindings.BindingType {
	return MgmtVmsListResultBindingType()
}

func (s MgmtVmsListResult) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for MgmtVmsListResult._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Interface information (Label)
type ModelInterface struct {
    // Name of the Interface label
	Name *string
    // Identifier of the Interface label
	Id string
}

func (s ModelInterface) GetType__() bindings.BindingType {
	return ModelInterfaceBindingType()
}

func (s ModelInterface) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ModelInterface._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Prefix lists used by certain features, like Intranet NAT.
type PrefixListInfo struct {
    // Prefix List URL
	Url string
    // Relative Prefix List path
	Path string
    // Prefix List name
	Name string
}

func (s PrefixListInfo) GetType__() bindings.BindingType {
	return PrefixListInfoBindingType()
}

func (s PrefixListInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for PrefixListInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// A list for provider gateways
type ProviderGatewayKeyValuePairs struct {
    // Value
	Value []ProviderObject
    // Key
	Key string
}

func (s ProviderGatewayKeyValuePairs) GetType__() bindings.BindingType {
	return ProviderGatewayKeyValuePairsBindingType()
}

func (s ProviderGatewayKeyValuePairs) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ProviderGatewayKeyValuePairs._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Basic properties of SDDC User Config.
type ProviderObject struct {
    // Should start with \"/policy/api/v1/\".
	Url *string
    // Path
	Path *string
    // Display name
	DisplayName *string
    // ID
	Id *string
    // Optional field, used to identify the object.
	Type_ *string
}

func (s ProviderObject) GetType__() bindings.BindingType {
	return ProviderObjectBindingType()
}

func (s ProviderObject) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ProviderObject._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type PublicIp struct {
    // Link to this resource
	Self *SelfResourceLink
    // The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
    // Schema for this resource
	Schema *string
    // The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected. format: int32
	Revision *int64
    // Indicates system owned resource
	SystemOwned *bool
	DisplayName *string
    // Description of this resource
	Description *string
    // Opaque identifiers meaningful to the API user
	Tags []Tag
    // ID of the user who created this resource
	CreateUser *string
    // Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super user and can modify it, but only when providing the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this entity.
	Protection *string
    // Timestamp of resource creation format: int64
	CreateTime *int64
    // Timestamp of last modification format: int64
	LastModifiedTime *int64
    // ID of the user who last modified this resource
	LastModifiedUser *string
    // Public IP identifier
	Id *string
    // The type of this resource.
	ResourceType *string
    // Absolute path of this object
	Path *string
    // Path relative from its parent
	RelativePath *string
    // Path of its parent
	ParentPath *string
    // marked for delete identifier
	MarkedForDelete *bool
    // IPv4 address format: ipv4
	Ip *string
}

func (s PublicIp) GetType__() bindings.BindingType {
	return PublicIpBindingType()
}

func (s PublicIp) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for PublicIp._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Public IP list
type PublicIpsListResult struct {
    // Link to this resource
	Self *SelfResourceLink
    // The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
    // Schema for this resource
	Schema *string
    // Opaque cursor to be used for getting next page of records (supplied by current result page)
	Cursor *string
    // If true, results are sorted in ascending order
	SortAscending *bool
    // Field by which records are sorted
	SortBy *string
    // Count of results found (across all pages), set only on first page format: int64
	ResultCount *int64
    // Public IP list
	Results []PublicIp
}

func (s PublicIpsListResult) GetType__() bindings.BindingType {
	return PublicIpsListResultBindingType()
}

func (s PublicIpsListResult) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for PublicIpsListResult._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Detailed information about a related API error
type RelatedApiError struct {
    // The module name where the error occurred
	ModuleName *string
    // A description of the error
	ErrorMessage *string
    // A numeric error code format: int64
	ErrorCode *int64
    // Further details about the error
	Details *string
    // Additional data about the error
	ErrorData *data.StructValue
}

func (s RelatedApiError) GetType__() bindings.BindingType {
	return RelatedApiErrorBindingType()
}

func (s RelatedApiError) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for RelatedApiError._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Base class for resources
type Resource struct {
    // Link to this resource
	Self *SelfResourceLink
    // The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
    // Schema for this resource
	Schema *string
}

func (s Resource) GetType__() bindings.BindingType {
	return ResourceBindingType()
}

func (s Resource) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for Resource._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
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

func (s ResourceLink) GetType__() bindings.BindingType {
	return ResourceLinkBindingType()
}

func (s ResourceLink) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ResourceLink._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// A base class for types that track revisions
type RevisionedResource struct {
    // Link to this resource
	Self *SelfResourceLink
    // The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
    // Schema for this resource
	Schema *string
    // The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected. format: int32
	Revision *int64
}

func (s RevisionedResource) GetType__() bindings.BindingType {
	return RevisionedResourceBindingType()
}

func (s RevisionedResource) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for RevisionedResource._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// SDDC configuration parameters for users. User-level addresses/CIDRs are provided.
type SddcUserConfiguration struct {
    // Prefix lists used by certain features, like Intranet NAT.
	ProviderPrefixLists []PrefixListInfo
    // Management gateway SNAT IP address. format: ipv4
	MgwSnatIp *string
    // Compute domain ID. Deprecated, please use domains.
	ComputeDomain *string
    // Interfaces (labels) including public interface, direct connect interface, linked vpc interface, etc. Deprecated, please use labels.
	Interfaces []ModelInterface
    // Interfaces (labels) including internet, intranet and services.
	Labels []ProviderObject
    // All VPN interfaces label name. Deprecated, please use labels.
	AllVpnInterfaceLabel *string
    // All uplink interfaces label name. Deprecated, please use labels.
	AllUplinkInterfaceLabel *string
    // Domain information for management gateway and compute gateway.
	Domains []ProviderObject
    // SDDC infra subnet CIDRs. Deprecated, please use infra_subnets. format: ipv4-cidr-block
	SddcInfraSubnet []string
    // SDDC infra subnet CIDRs. format: ipv4-cidr-block
	InfraSubnets []string
    // Management gateway label name. Deprecated, please use labels.
	ManagementGatewayLabel *string
    // Public IPs for VPN tunnel over internet. Deprecated. Please use vpn_endpoints instead of vpn_internet_ips. format: ipv4
	VpnInternetIps []string
    // Management subnet CIDRs. format: ipv4-cidr-block
	MgmtSubnets []string
    // Policy enforcement point object paths.
	EnforcementPoints []ProviderObject
    // Compute gateway name. Deprecated, please use gateways.
	ComputeGateway *string
    // Local IPs for VPN tunnel over Direct Connect. Deprecated. Please use vpn_endpoints instead of vpn_dx_ips. format: ipv4
	VpnDxIps []string
    // Linked VPC interface label name. Deprecated, please use labels.
	LinkedVpcInterfaceLabel *string
    // DirectConnect interface label name. Deprecated, please use labels.
	DxInterfaceLabel *string
    // Management subnet CIDRs. Deprecated, please use mgmt_subnets. format: ipv4-cidr-block
	MgmtSubnet []string
    // Management gateway name. Deprecated, please use gateways.
	ManagementGateway *string
    // Provider gateway list. Including both tier-0 gateways and tier-1 gateways.
	Gateways []ProviderGatewayKeyValuePairs
    // Public interface label name. Deprecated, please use labels.
	PublicInterfaceLabel *string
    // Compute gateway SNAT IP address. format: ipv4
	CgwSnatIp *string
    // VPN tunnel endpoints. Currently containing public IPs for VPN over internet and local IPs for VPN over intranet.
	VpnEndpoints []VpnEndpoint
    // Management domain ID. Deprecated, please use domains.
	ManagementDomain *string
    // Service provider name. Deprecated, please use gateways.
	ProviderName *string
    // Provider gateway list. Including both tier-0 gateways and tier-1 gateways. Deprecated, please use gateways.
	ProviderGateways []ProviderGatewayKeyValuePairs
    // Management gateway default DNS zone
	ManagementGatewayDefaultDnsZone ProviderObject
}

func (s SddcUserConfiguration) GetType__() bindings.BindingType {
	return SddcUserConfigurationBindingType()
}

func (s SddcUserConfiguration) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SddcUserConfiguration._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
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

func (s SelfResourceLink) GetType__() bindings.BindingType {
	return SelfResourceLinkBindingType()
}

func (s SelfResourceLink) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SelfResourceLink._GetDataValue method - %s",
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


// Arbitrary key-value pairs that may be attached to an entity
type Tag struct {
    // Tag searches may optionally be restricted by scope
	Scope *string
    // Identifier meaningful to user with maximum length of 256 characters
	Tag *string
}

func (s Tag) GetType__() bindings.BindingType {
	return TagBindingType()
}

func (s Tag) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for Tag._GetDataValue method - %s",
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


// Traffic group configuration. A traffic group indicates dedicated bandwidth and computation for a given list of subnets. Creating a traffic group will reserve resources and associating it with desired prefix lists will use the resources for the traffic of the prefix lists. Besides traffic group ID and name, realized state and detailed association maps info of this traffic group are included if verbose info is requested.
type TrafficGroup struct {
    // Link to this resource
	Self *SelfResourceLink
    // The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
    // Schema for this resource
	Schema *string
    // The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected. format: int32
	Revision *int64
    // Indicates system owned resource
	SystemOwned *bool
    // A description field for the traffic group.
	DisplayName *string
    // Description of this resource
	Description *string
    // Opaque identifiers meaningful to the API user
	Tags []Tag
    // ID of the user who created this resource
	CreateUser *string
    // Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super user and can modify it, but only when providing the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this entity.
	Protection *string
    // Timestamp of resource creation format: int64
	CreateTime *int64
    // Timestamp of last modification format: int64
	LastModifiedTime *int64
    // ID of the user who last modified this resource
	LastModifiedUser *string
    // Unique identifier for a traffic group.
	Id *string
    // The type of this resource.
	ResourceType *string
    // Absolute path of this object
	Path *string
    // Path relative from its parent
	RelativePath *string
    // Path of its parent
	ParentPath *string
    // marked for delete identifier
	MarkedForDelete *bool
    // Information on the current state. Mostly error messages on failure states.
	StateMessage *string
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
}
const TrafficGroup_STATE_IN_PROGRESS = "IN_PROGRESS"
const TrafficGroup_STATE_SUCCESS = "SUCCESS"
const TrafficGroup_STATE_FAILURE = "FAILURE"
const TrafficGroup_STATE_UNAVAILABLE = "UNAVAILABLE"
const TrafficGroup_STATE_PENDING = "PENDING"

func (s TrafficGroup) GetType__() bindings.BindingType {
	return TrafficGroupBindingType()
}

func (s TrafficGroup) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for TrafficGroup._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// An association map of a traffic group. It describes the association of prefix lists and logical routers with the traffic group. To make use of dedicated traffic resources through traffic groups, prefix lists need to be linked to the desired traffic group and an association with the target logical router (scope) is required. The scope is for declaring which logical router the prefix lists are associated under a particular traffic group. The scope is not required input.
type TrafficGroupAssociationMap struct {
    // Link to this resource
	Self *SelfResourceLink
    // The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
    // Schema for this resource
	Schema *string
    // The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected. format: int32
	Revision *int64
    // Indicates system owned resource
	SystemOwned *bool
    // A description field for Traffic group association map。
	DisplayName *string
    // Description of this resource
	Description *string
    // Opaque identifiers meaningful to the API user
	Tags []Tag
    // ID of the user who created this resource
	CreateUser *string
    // Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super user and can modify it, but only when providing the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this entity.
	Protection *string
    // Timestamp of resource creation format: int64
	CreateTime *int64
    // Timestamp of last modification format: int64
	LastModifiedTime *int64
    // ID of the user who last modified this resource
	LastModifiedUser *string
    // The association map id.
	Id *string
    // The type of this resource.
	ResourceType *string
    // Absolute path of this object
	Path *string
    // Path relative from its parent
	RelativePath *string
    // Path of its parent
	ParentPath *string
    // marked for delete identifier
	MarkedForDelete *bool
    // Information on the current state. Mostly error messages on failure states.
	StateMessage *string
    // Possible values are: 
    //
    // * TrafficGroupAssociationMap#TrafficGroupAssociationMap_SCOPE_1S_CGW
    // * TrafficGroupAssociationMap#TrafficGroupAssociationMap_SCOPE_0S_VMC
    //
    //  The targeted logical router (scope) of prefix lists. Non admin users are not allowed to create, update, delete an association map with scope as /infra/tier-0s/vmc.
	Scope *string
    // The list of prefix lists to be associated.
	PrefixLists []string
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
}
const TrafficGroupAssociationMap_SCOPE_1S_CGW = "/infra/tier-1s/cgw"
const TrafficGroupAssociationMap_SCOPE_0S_VMC = "/infra/tier-0s/vmc"
const TrafficGroupAssociationMap_STATE_IN_PROGRESS = "IN_PROGRESS"
const TrafficGroupAssociationMap_STATE_SUCCESS = "SUCCESS"
const TrafficGroupAssociationMap_STATE_FAILURE = "FAILURE"
const TrafficGroupAssociationMap_STATE_UNAVAILABLE = "UNAVAILABLE"
const TrafficGroupAssociationMap_STATE_PENDING = "PENDING"

func (s TrafficGroupAssociationMap) GetType__() bindings.BindingType {
	return TrafficGroupAssociationMapBindingType()
}

func (s TrafficGroupAssociationMap) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for TrafficGroupAssociationMap._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// List result for traffic group association maps.
type TrafficGroupAssociationMapsListResult struct {
    // Link to this resource
	Self *SelfResourceLink
    // The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
    // Schema for this resource
	Schema *string
    // Opaque cursor to be used for getting next page of records (supplied by current result page)
	Cursor *string
    // If true, results are sorted in ascending order
	SortAscending *bool
    // Field by which records are sorted
	SortBy *string
    // Count of results found (across all pages), set only on first page format: int64
	ResultCount *int64
    // Traffic group association map list
	Results []TrafficGroupAssociationMap
}

func (s TrafficGroupAssociationMapsListResult) GetType__() bindings.BindingType {
	return TrafficGroupAssociationMapsListResultBindingType()
}

func (s TrafficGroupAssociationMapsListResult) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for TrafficGroupAssociationMapsListResult._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// List result for traffic groups.
type TrafficGroupsListResult struct {
    // Link to this resource
	Self *SelfResourceLink
    // The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
    // Schema for this resource
	Schema *string
    // Opaque cursor to be used for getting next page of records (supplied by current result page)
	Cursor *string
    // If true, results are sorted in ascending order
	SortAscending *bool
    // Field by which records are sorted
	SortBy *string
    // Count of results found (across all pages), set only on first page format: int64
	ResultCount *int64
    // Traffic group list
	Results []TrafficGroup
}

func (s TrafficGroupsListResult) GetType__() bindings.BindingType {
	return TrafficGroupsListResultBindingType()
}

func (s TrafficGroupsListResult) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for TrafficGroupsListResult._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Shadow account and linked VPC account
type VMCAccounts struct {
    // linked VPC account number
	LinkedVpcAccount *string
    // Shadow VPC account number
	ShadowAccount string
}

func (s VMCAccounts) GetType__() bindings.BindingType {
	return VMCAccountsBindingType()
}

func (s VMCAccounts) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for VMCAccounts._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Direct Connect VIFs (Virtual Interface) list query result
type VifsListResult struct {
    // Link to this resource
	Self *SelfResourceLink
    // The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
    // Schema for this resource
	Schema *string
    // Opaque cursor to be used for getting next page of records (supplied by current result page)
	Cursor *string
    // If true, results are sorted in ascending order
	SortAscending *bool
    // Field by which records are sorted
	SortBy *string
    // Count of results found (across all pages), set only on first page format: int64
	ResultCount *int64
    // VIFs list
	Results []VirtualInterface
}

func (s VifsListResult) GetType__() bindings.BindingType {
	return VifsListResultBindingType()
}

func (s VifsListResult) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for VifsListResult._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type VirtualInterface struct {
    // Identifier for the Direct Connect
	DirectConnectId string
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
	State string
    // Remote autonomous system number of vif
	RemoteAsn *string
    // VIF name
	Name string
    // amazon side address format: ipv4
	LocalIp *string
    // customer address format: ipv4
	RemoteIp *string
    // Identifier for the virtual interface
	Id string
    // Possible values are: 
    //
    // * VirtualInterface#VirtualInterface_BGP_STATUS_UP
    // * VirtualInterface#VirtualInterface_BGP_STATUS_DOWN
    //
    //  BGP status
	BgpStatus string
    // Maximum transmission unit allowed by the VIF format: int32
	Mtu *int64
}
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
const VirtualInterface_BGP_STATUS_UP = "UP"
const VirtualInterface_BGP_STATUS_DOWN = "DOWN"

func (s VirtualInterface) GetType__() bindings.BindingType {
	return VirtualInterfaceBindingType()
}

func (s VirtualInterface) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for VirtualInterface._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Abstract base class for all the VmcApp objects.
type VmcAppBaseInfo struct {
    // Link to this resource
	Self *SelfResourceLink
    // The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
    // Schema for this resource
	Schema *string
    // The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected. format: int32
	Revision *int64
    // Indicates system owned resource
	SystemOwned *bool
    // Defaults to ID if not set
	DisplayName *string
    // Description of this resource
	Description *string
    // Opaque identifiers meaningful to the API user
	Tags []Tag
    // ID of the user who created this resource
	CreateUser *string
    // Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super user and can modify it, but only when providing the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this entity.
	Protection *string
    // Timestamp of resource creation format: int64
	CreateTime *int64
    // Timestamp of last modification format: int64
	LastModifiedTime *int64
    // ID of the user who last modified this resource
	LastModifiedUser *string
    // Unique identifier of this resource
	Id *string
    // The type of this resource.
	ResourceType *string
    // Absolute path of this object
	Path *string
    // Path relative from its parent
	RelativePath *string
    // Path of its parent
	ParentPath *string
    // marked for delete identifier
	MarkedForDelete *bool
}

func (s VmcAppBaseInfo) GetType__() bindings.BindingType {
	return VmcAppBaseInfoBindingType()
}

func (s VmcAppBaseInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for VmcAppBaseInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Represents aggregated realized status for intent entity across associated realized entities.
type VmcConsolidatedRealizedStatus struct {
    // Consolidated state of objects for a given intent entity.
	ConsolidatedStatus *VmcConsolidatedStatus
    // Aggregated consolidated status by enforcement point.
	ConsolidatedStatusPerObject []VmcConsolidatedStatusPerObject
    // Intent path of the object representing this consolidated state.
	IntentPath *string
}

func (s VmcConsolidatedRealizedStatus) GetType__() bindings.BindingType {
	return VmcConsolidatedRealizedStatusBindingType()
}

func (s VmcConsolidatedRealizedStatus) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for VmcConsolidatedRealizedStatus._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Consolidated status of an object.
type VmcConsolidatedStatus struct {
    // Help message for the current status regarding an object, providing information for each state.
	StatusMessage *string
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
}
const VmcConsolidatedStatus_CONSOLIDATED_STATUS_IN_PROGRESS = "IN_PROGRESS"
const VmcConsolidatedStatus_CONSOLIDATED_STATUS_SUCCESS = "SUCCESS"
const VmcConsolidatedStatus_CONSOLIDATED_STATUS_FAILURE = "FAILURE"
const VmcConsolidatedStatus_CONSOLIDATED_STATUS_UNAVAILABLE = "UNAVAILABLE"
const VmcConsolidatedStatus_CONSOLIDATED_STATUS_PENDING = "PENDING"

func (s VmcConsolidatedStatus) GetType__() bindings.BindingType {
	return VmcConsolidatedStatusBindingType()
}

func (s VmcConsolidatedStatus) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for VmcConsolidatedStatus._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Realized status consolidated by individual objects.
type VmcConsolidatedStatusPerObject struct {
    // Detailed consolidated realized status for an intent object.
	ConsolidatedStatus *VmcConsolidatedStatus
    // Object id used to consolidate state. This can be a particular backend task/job, etc.
	ObjectId string
}

func (s VmcConsolidatedStatusPerObject) GetType__() bindings.BindingType {
	return VmcConsolidatedStatusPerObjectBindingType()
}

func (s VmcConsolidatedStatusPerObject) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for VmcConsolidatedStatusPerObject._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// VMC Feature Flag
type VmcFeatureFlagInfo struct {
    // Link to this resource
	Self *SelfResourceLink
    // The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
    // Schema for this resource
	Schema *string
    // The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected. format: int32
	Revision *int64
    // Indicates system owned resource
	SystemOwned *bool
    // Defaults to ID if not set
	DisplayName *string
    // Description of this resource
	Description *string
    // Opaque identifiers meaningful to the API user
	Tags []Tag
    // ID of the user who created this resource
	CreateUser *string
    // Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super user and can modify it, but only when providing the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this entity.
	Protection *string
    // Timestamp of resource creation format: int64
	CreateTime *int64
    // Timestamp of last modification format: int64
	LastModifiedTime *int64
    // ID of the user who last modified this resource
	LastModifiedUser *string
    // Unique identifier of this resource
	Id *string
    // The type of this resource.
	ResourceType *string
    // Absolute path of this object
	Path *string
    // Path relative from its parent
	RelativePath *string
    // Path of its parent
	ParentPath *string
    // marked for delete identifier
	MarkedForDelete *bool
    // Possible values are: 
    //
    // * VmcFeatureFlagInfo#VmcFeatureFlagInfo_STATE_ENABLED
    // * VmcFeatureFlagInfo#VmcFeatureFlagInfo_STATE_DISABLED
    // * VmcFeatureFlagInfo#VmcFeatureFlagInfo_STATE_INACTIVE
    //
    //  state
	State string
    // Message
	Message *string
    // Internal Name
	InternalName *string
    // Feature Name
	Name string
}
const VmcFeatureFlagInfo_STATE_ENABLED = "enabled"
const VmcFeatureFlagInfo_STATE_DISABLED = "disabled"
const VmcFeatureFlagInfo_STATE_INACTIVE = "inactive"

func (s VmcFeatureFlagInfo) GetType__() bindings.BindingType {
	return VmcFeatureFlagInfoBindingType()
}

func (s VmcFeatureFlagInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for VmcFeatureFlagInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// VMC Feature flags
type VmcFeatureFlags struct {
	Features []VmcFeatureFlagInfo
}

func (s VmcFeatureFlags) GetType__() bindings.BindingType {
	return VmcFeatureFlagsBindingType()
}

func (s VmcFeatureFlags) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for VmcFeatureFlags._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// A consolidated object of realized entities given an intent path. This accounts for resources / entities which are realized under the intent path.
type VmcRealizedEntities struct {
    // Detailed realized entities list.
	RealizedEntities []VmcRealizedEntity
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
    // Realized entities id
	RealizedEntitiesId *string
    // Intent path
	IntentPath *string
}
const VmcRealizedEntities_REALIZED_STATE_IN_PROGRESS = "IN_PROGRESS"
const VmcRealizedEntities_REALIZED_STATE_SUCCESS = "SUCCESS"
const VmcRealizedEntities_REALIZED_STATE_FAILURE = "FAILURE"
const VmcRealizedEntities_REALIZED_STATE_UNAVAILABLE = "UNAVAILABLE"
const VmcRealizedEntities_REALIZED_STATE_PENDING = "PENDING"

func (s VmcRealizedEntities) GetType__() bindings.BindingType {
	return VmcRealizedEntitiesBindingType()
}

func (s VmcRealizedEntities) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for VmcRealizedEntities._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Realized entity object. This accounts for an resource / entity which is realized within cloud service.
type VmcRealizedEntity struct {
    // Resource realization id. This can differ from realized_entity_id as this can be an external id.
	RealizationId *string
    // Resource realization API path
	RealizationApi *string
    // Realized entity id
	RealizedEntityId *string
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
    // The path for the realization of an entity. This can be an URI, etc. Some resources are identified by their paths.
	RealizationPath *string
    // Realized entity type
	RealizedEntityType *string
}
const VmcRealizedEntity_REALIZED_STATE_IN_PROGRESS = "IN_PROGRESS"
const VmcRealizedEntity_REALIZED_STATE_SUCCESS = "SUCCESS"
const VmcRealizedEntity_REALIZED_STATE_FAILURE = "FAILURE"
const VmcRealizedEntity_REALIZED_STATE_UNAVAILABLE = "UNAVAILABLE"
const VmcRealizedEntity_REALIZED_STATE_PENDING = "PENDING"

func (s VmcRealizedEntity) GetType__() bindings.BindingType {
	return VmcRealizedEntityBindingType()
}

func (s VmcRealizedEntity) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for VmcRealizedEntity._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// VPN endpoint information
type VpnEndpoint struct {
    // IP address of the VPN endpoint format: ipv4
	Ip string
    // Type of the VPN endpoint
	Type_ string
    // Name of the VPN endpoint
	Name string
    // Interface label of the VPN endpoint
	InterfaceLabel string
}

func (s VpnEndpoint) GetType__() bindings.BindingType {
	return VpnEndpointBindingType()
}

func (s VpnEndpoint) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for VpnEndpoint._GetDataValue method - %s",
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

func ActionMessageBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["message"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["message"] = "Message"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.action_message", fields, reflect.TypeOf(ActionMessage{}), fieldNameMap, validators)
}

func AdvertisedRouteBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipv4_cidr"] = bindings.NewStringType()
	fieldNameMap["ipv4_cidr"] = "Ipv4Cidr"
	fields["advertisement_state"] = bindings.NewStringType()
	fieldNameMap["advertisement_state"] = "AdvertisementState"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.advertised_route", fields, reflect.TypeOf(AdvertisedRoute{}), fieldNameMap, validators)
}

func ApiErrorBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["module_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["module_name"] = "ModuleName"
	fields["error_message"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["error_message"] = "ErrorMessage"
	fields["error_code"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["error_code"] = "ErrorCode"
	fields["details"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["details"] = "Details"
	fields["error_data"] = bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
	fieldNameMap["error_data"] = "ErrorData"
	fields["related_errors"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(RelatedApiErrorBindingType), reflect.TypeOf([]RelatedApiError{})))
	fieldNameMap["related_errors"] = "RelatedErrors"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.api_error", fields, reflect.TypeOf(ApiError{}), fieldNameMap, validators)
}

func AssociatedBaseGroupConnectionInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_revision"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_system_owned"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["tags"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["_create_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_protection"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_create_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_last_modified_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["id"] = bindings.NewStringType()
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = bindings.NewStringType()
	fieldNameMap["resource_type"] = "ResourceType"
	fields["path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	fields["parent_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["marked_for_delete"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.associated_base_group_connection_info", fields, reflect.TypeOf(AssociatedBaseGroupConnectionInfo{}), fieldNameMap, validators)
}

func AssociatedGroupConnectionInfosListResultBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["sort_ascending"] = "SortAscending"
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sort_by"] = "SortBy"
	fields["result_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["result_count"] = "ResultCount"
	fields["results"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(AssociatedBaseGroupConnectionInfoBindingType),}, bindings.JSONRPC), reflect.TypeOf([]*data.StructValue{})))
	fieldNameMap["results"] = "Results"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.associated_group_connection_infos_list_result", fields, reflect.TypeOf(AssociatedGroupConnectionInfosListResult{}), fieldNameMap, validators)
}

func AssociatedTgwGroupConnectionInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tgw_id"] = bindings.NewStringType()
	fieldNameMap["tgw_id"] = "TgwId"
	fields["external_route_table_id"] = bindings.NewStringType()
	fieldNameMap["external_route_table_id"] = "ExternalRouteTableId"
	fields["state"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["state"] = "State"
	fields["sddcs_route_table_id"] = bindings.NewStringType()
	fieldNameMap["sddcs_route_table_id"] = "SddcsRouteTableId"
	fields["tgw_attachment_id"] = bindings.NewStringType()
	fieldNameMap["tgw_attachment_id"] = "TgwAttachmentId"
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_revision"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_system_owned"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["tags"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["_create_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_protection"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_create_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_last_modified_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["id"] = bindings.NewStringType()
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = bindings.NewStringType()
	fieldNameMap["resource_type"] = "ResourceType"
	fields["path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	fields["parent_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["marked_for_delete"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.associated_tgw_group_connection_info", fields, reflect.TypeOf(AssociatedTgwGroupConnectionInfo{}), fieldNameMap, validators)
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

func BGPAdvertisedRoutesBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["failed_advertised_routes"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["failed_advertised_routes"] = "FailedAdvertisedRoutes"
	fields["advertised_routes"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(AdvertisedRouteBindingType), reflect.TypeOf([]AdvertisedRoute{})))
	fieldNameMap["advertised_routes"] = "AdvertisedRoutes"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.BGP_advertised_routes", fields, reflect.TypeOf(BGPAdvertisedRoutes{}), fieldNameMap, validators)
}

func BGPLearnedRoutesBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipv4_cidr"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["ipv4_cidr"] = "Ipv4Cidr"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.BGP_learned_routes", fields, reflect.TypeOf(BGPLearnedRoutes{}), fieldNameMap, validators)
}

func ConnectedServiceListResultBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["sort_ascending"] = "SortAscending"
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sort_by"] = "SortBy"
	fields["result_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["result_count"] = "ResultCount"
	fields["results"] = bindings.NewListType(bindings.NewReferenceType(ConnectedServiceStatusBindingType), reflect.TypeOf([]ConnectedServiceStatus{}))
	fieldNameMap["results"] = "Results"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.connected_service_list_result", fields, reflect.TypeOf(ConnectedServiceListResult{}), fieldNameMap, validators)
}

func ConnectedServiceStatusBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_revision"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_system_owned"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["tags"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["_create_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_protection"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_create_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_last_modified_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	fields["parent_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["marked_for_delete"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["enabled"] = "Enabled"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.connected_service_status", fields, reflect.TypeOf(ConnectedServiceStatus{}), fieldNameMap, validators)
}

func CsvListResultBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["file_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["file_name"] = "FileName"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.csv_list_result", fields, reflect.TypeOf(CsvListResult{}), fieldNameMap, validators)
}

func CsvRecordBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.csv_record", fields, reflect.TypeOf(CsvRecord{}), fieldNameMap, validators)
}

func DirectConnectBgpInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_revision"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_system_owned"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["tags"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["_create_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_protection"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_create_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_last_modified_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	fields["parent_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["marked_for_delete"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["local_as_num"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["local_as_num"] = "LocalAsNum"
	fields["route_preference"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["route_preference"] = "RoutePreference"
	fields["mtu"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["mtu"] = "Mtu"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.direct_connect_bgp_info", fields, reflect.TypeOf(DirectConnectBgpInfo{}), fieldNameMap, validators)
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
	fields["edrs_policy_options_overrides"] = bindings.NewOptionalType(bindings.NewReferenceType(EdrsPolicyOptionsOverridesBindingType))
	fieldNameMap["edrs_policy_options_overrides"] = "EdrsPolicyOptionsOverrides"
	fields["min_hosts"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["min_hosts"] = "MinHosts"
	fields["policy_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["policy_type"] = "PolicyType"
	fields["max_hosts"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["max_hosts"] = "MaxHosts"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.edrs_policy", fields, reflect.TypeOf(EdrsPolicy{}), fieldNameMap, validators)
}

func EdrsPolicyOptionsOverridesBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["scale_up_host_increment"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["scale_up_host_increment"] = "ScaleUpHostIncrement"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.edrs_policy_options_overrides", fields, reflect.TypeOf(EdrsPolicyOptionsOverrides{}), fieldNameMap, validators)
}

func EdrsPolicySpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["policy_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["policy_type"] = "PolicyType"
	fields["cluster_eligible_for_policy"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["cluster_eligible_for_policy"] = "ClusterEligibleForPolicy"
	fields["configurable_scaleup_increment"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["configurable_scaleup_increment"] = "ConfigurableScaleupIncrement"
	fields["min_max_host_range"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewIntegerType(), reflect.TypeOf([]int64{})))
	fieldNameMap["min_max_host_range"] = "MinMaxHostRange"
	fields["scaleup_host_increment_range"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewIntegerType(), reflect.TypeOf([]int64{})))
	fieldNameMap["scaleup_host_increment_range"] = "ScaleupHostIncrementRange"
	fields["configurable"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["configurable"] = "Configurable"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.edrs_policy_spec", fields, reflect.TypeOf(EdrsPolicySpec{}), fieldNameMap, validators)
}

func EdrsProvisioningSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["current_edrs_policy"] = bindings.NewOptionalType(bindings.NewReferenceType(EdrsPolicyBindingType))
	fieldNameMap["current_edrs_policy"] = "CurrentEdrsPolicy"
	fields["policies"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(EdrsPolicySpecBindingType), reflect.TypeOf([]EdrsPolicySpec{})))
	fieldNameMap["policies"] = "Policies"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.edrs_provisioning_spec", fields, reflect.TypeOf(EdrsProvisioningSpec{}), fieldNameMap, validators)
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

func ExternalConnectivityConfigBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_revision"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_system_owned"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["tags"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["_create_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_protection"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_create_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_last_modified_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	fields["parent_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["marked_for_delete"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["intranet_mtu"] = bindings.NewIntegerType()
	fieldNameMap["intranet_mtu"] = "IntranetMtu"
	fields["services_mtu"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["services_mtu"] = "ServicesMtu"
	fields["internet_mtu"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["internet_mtu"] = "InternetMtu"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.external_connectivity_config", fields, reflect.TypeOf(ExternalConnectivityConfig{}), fieldNameMap, validators)
}

func ExternalSddcConnectivityBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["status"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["status"] = "Status"
	fields["connectivity_type"] = bindings.NewStringType()
	fieldNameMap["connectivity_type"] = "ConnectivityType"
	fields["route_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["route_type"] = "RouteType"
	fields["status_message"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["status_message"] = "StatusMessage"
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["source"] = "Source"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.external_sddc_connectivity", fields, reflect.TypeOf(ExternalSddcConnectivity{}), fieldNameMap, validators)
}

func ExternalSddcRouteBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_revision"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_system_owned"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["tags"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["_create_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_protection"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_create_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_last_modified_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	fields["parent_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["marked_for_delete"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["destination"] = bindings.NewStringType()
	fieldNameMap["destination"] = "Destination"
	fields["connectivities"] = bindings.NewListType(bindings.NewReferenceType(ExternalSddcConnectivityBindingType), reflect.TypeOf([]ExternalSddcConnectivity{}))
	fieldNameMap["connectivities"] = "Connectivities"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.external_sddc_route", fields, reflect.TypeOf(ExternalSddcRoute{}), fieldNameMap, validators)
}

func ExternalSddcRouteCsvRecordBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["destination"] = bindings.NewStringType()
	fieldNameMap["destination"] = "Destination"
	fields["connectivity_details"] = bindings.NewStringType()
	fieldNameMap["connectivity_details"] = "ConnectivityDetails"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.external_sddc_route_csv_record", fields, reflect.TypeOf(ExternalSddcRouteCsvRecord{}), fieldNameMap, validators)
}

func ExternalSddcRoutesListResultBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["sort_ascending"] = "SortAscending"
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sort_by"] = "SortBy"
	fields["result_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["result_count"] = "ResultCount"
	fields["routes"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ExternalSddcRouteBindingType), reflect.TypeOf([]ExternalSddcRoute{})))
	fieldNameMap["routes"] = "Routes"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.external_sddc_routes_list_result", fields, reflect.TypeOf(ExternalSddcRoutesListResult{}), fieldNameMap, validators)
}

func ExternalSddcRoutesListResultInCsvFormatBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["file_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["file_name"] = "FileName"
	fields["results"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ExternalSddcRouteCsvRecordBindingType), reflect.TypeOf([]ExternalSddcRouteCsvRecord{})))
	fieldNameMap["results"] = "Results"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.external_sddc_routes_list_result_in_csv_format", fields, reflect.TypeOf(ExternalSddcRoutesListResultInCsvFormat{}), fieldNameMap, validators)
}

func IncludedFieldsParametersBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["included_fields"] = "IncludedFields"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.included_fields_parameters", fields, reflect.TypeOf(IncludedFieldsParameters{}), fieldNameMap, validators)
}

func IpAttachmentPairBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ip"] = bindings.NewStringType()
	fieldNameMap["ip"] = "Ip"
	fields["attachment_id"] = bindings.NewStringType()
	fieldNameMap["attachment_id"] = "AttachmentId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.ip_attachment_pair", fields, reflect.TypeOf(IpAttachmentPair{}), fieldNameMap, validators)
}

func LinkedSubnetInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cidr"] = bindings.NewStringType()
	fieldNameMap["cidr"] = "Cidr"
	fields["id"] = bindings.NewStringType()
	fieldNameMap["id"] = "Id"
	fields["availability_zone"] = bindings.NewStringType()
	fieldNameMap["availability_zone"] = "AvailabilityZone"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.linked_subnet_info", fields, reflect.TypeOf(LinkedSubnetInfo{}), fieldNameMap, validators)
}

func LinkedVpcInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_revision"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_system_owned"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["tags"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["_create_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_protection"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_create_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_last_modified_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	fields["parent_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["marked_for_delete"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["arn_role"] = bindings.NewStringType()
	fieldNameMap["arn_role"] = "ArnRole"
	fields["active_eni"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["active_eni"] = "ActiveEni"
	fields["linked_vpc_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["linked_vpc_id"] = "LinkedVpcId"
	fields["linked_vpc_addresses"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["linked_vpc_addresses"] = "LinkedVpcAddresses"
	fields["linked_vpc_nat_ips"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["linked_vpc_nat_ips"] = "LinkedVpcNatIps"
	fields["linked_account"] = bindings.NewStringType()
	fieldNameMap["linked_account"] = "LinkedAccount"
	fields["route_table_ids"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["route_table_ids"] = "RouteTableIds"
	fields["service_arn_role"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["service_arn_role"] = "ServiceArnRole"
	fields["linked_vpc_subnets"] = bindings.NewListType(bindings.NewReferenceType(LinkedSubnetInfoBindingType), reflect.TypeOf([]LinkedSubnetInfo{}))
	fieldNameMap["linked_vpc_subnets"] = "LinkedVpcSubnets"
	fields["external_id"] = bindings.NewStringType()
	fieldNameMap["external_id"] = "ExternalId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.linked_vpc_info", fields, reflect.TypeOf(LinkedVpcInfo{}), fieldNameMap, validators)
}

func LinkedVpcsListResultBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["sort_ascending"] = "SortAscending"
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sort_by"] = "SortBy"
	fields["result_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["result_count"] = "ResultCount"
	fields["results"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(LinkedVpcInfoBindingType), reflect.TypeOf([]LinkedVpcInfo{})))
	fieldNameMap["results"] = "Results"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.linked_vpcs_list_result", fields, reflect.TypeOf(LinkedVpcsListResult{}), fieldNameMap, validators)
}

func ListResultBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["sort_ascending"] = "SortAscending"
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sort_by"] = "SortBy"
	fields["result_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["result_count"] = "ResultCount"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.list_result", fields, reflect.TypeOf(ListResult{}), fieldNameMap, validators)
}

func ManagedResourceBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_revision"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_system_owned"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["tags"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["_create_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_protection"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_create_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_last_modified_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.managed_resource", fields, reflect.TypeOf(ManagedResource{}), fieldNameMap, validators)
}

func MgmtServiceEntryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.mgmt_service_entry", fields, reflect.TypeOf(MgmtServiceEntry{}), fieldNameMap, validators)
}

func MgmtVmInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_revision"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_system_owned"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["tags"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["_create_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_protection"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_create_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_last_modified_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	fields["parent_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["marked_for_delete"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["ips"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["ips"] = "Ips"
	fields["services"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(MgmtServiceEntryBindingType), reflect.TypeOf([]MgmtServiceEntry{})))
	fieldNameMap["services"] = "Services"
	fields["group_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["group_path"] = "GroupPath"
	fields["ip_attachment_pairs"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(IpAttachmentPairBindingType), reflect.TypeOf([]IpAttachmentPair{})))
	fieldNameMap["ip_attachment_pairs"] = "IpAttachmentPairs"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.mgmt_vm_info", fields, reflect.TypeOf(MgmtVmInfo{}), fieldNameMap, validators)
}

func MgmtVmsListResultBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["sort_ascending"] = "SortAscending"
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sort_by"] = "SortBy"
	fields["result_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["result_count"] = "ResultCount"
	fields["results"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(MgmtVmInfoBindingType), reflect.TypeOf([]MgmtVmInfo{})))
	fieldNameMap["results"] = "Results"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.mgmt_vms_list_result", fields, reflect.TypeOf(MgmtVmsListResult{}), fieldNameMap, validators)
}

func ModelInterfaceBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["id"] = bindings.NewStringType()
	fieldNameMap["id"] = "Id"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.model_interface", fields, reflect.TypeOf(ModelInterface{}), fieldNameMap, validators)
}

func PrefixListInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["url"] = bindings.NewStringType()
	fieldNameMap["url"] = "Url"
	fields["path"] = bindings.NewStringType()
	fieldNameMap["path"] = "Path"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.prefix_list_info", fields, reflect.TypeOf(PrefixListInfo{}), fieldNameMap, validators)
}

func ProviderGatewayKeyValuePairsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["value"] = bindings.NewListType(bindings.NewReferenceType(ProviderObjectBindingType), reflect.TypeOf([]ProviderObject{}))
	fieldNameMap["value"] = "Value"
	fields["key"] = bindings.NewStringType()
	fieldNameMap["key"] = "Key"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.provider_gateway_key_value_pairs", fields, reflect.TypeOf(ProviderGatewayKeyValuePairs{}), fieldNameMap, validators)
}

func ProviderObjectBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["url"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["url"] = "Url"
	fields["path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["type"] = "Type_"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.provider_object", fields, reflect.TypeOf(ProviderObject{}), fieldNameMap, validators)
}

func PublicIpBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_revision"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_system_owned"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["tags"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["_create_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_protection"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_create_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_last_modified_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	fields["parent_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["marked_for_delete"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["ip"] = "Ip"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.public_ip", fields, reflect.TypeOf(PublicIp{}), fieldNameMap, validators)
}

func PublicIpsListResultBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["sort_ascending"] = "SortAscending"
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sort_by"] = "SortBy"
	fields["result_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["result_count"] = "ResultCount"
	fields["results"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(PublicIpBindingType), reflect.TypeOf([]PublicIp{})))
	fieldNameMap["results"] = "Results"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.public_ips_list_result", fields, reflect.TypeOf(PublicIpsListResult{}), fieldNameMap, validators)
}

func RelatedApiErrorBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["module_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["module_name"] = "ModuleName"
	fields["error_message"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["error_message"] = "ErrorMessage"
	fields["error_code"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["error_code"] = "ErrorCode"
	fields["details"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["details"] = "Details"
	fields["error_data"] = bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
	fieldNameMap["error_data"] = "ErrorData"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.related_api_error", fields, reflect.TypeOf(RelatedApiError{}), fieldNameMap, validators)
}

func ResourceBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.resource", fields, reflect.TypeOf(Resource{}), fieldNameMap, validators)
}

func ResourceLinkBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["action"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["action"] = "Action"
	fields["href"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["href"] = "Href"
	fields["rel"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["rel"] = "Rel"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.resource_link", fields, reflect.TypeOf(ResourceLink{}), fieldNameMap, validators)
}

func RevisionedResourceBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_revision"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.revisioned_resource", fields, reflect.TypeOf(RevisionedResource{}), fieldNameMap, validators)
}

func SddcUserConfigurationBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["provider_prefix_lists"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(PrefixListInfoBindingType), reflect.TypeOf([]PrefixListInfo{})))
	fieldNameMap["provider_prefix_lists"] = "ProviderPrefixLists"
	fields["mgw_snat_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["mgw_snat_ip"] = "MgwSnatIp"
	fields["compute_domain"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["compute_domain"] = "ComputeDomain"
	fields["interfaces"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ModelInterfaceBindingType), reflect.TypeOf([]ModelInterface{})))
	fieldNameMap["interfaces"] = "Interfaces"
	fields["labels"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ProviderObjectBindingType), reflect.TypeOf([]ProviderObject{})))
	fieldNameMap["labels"] = "Labels"
	fields["all_vpn_interface_label"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["all_vpn_interface_label"] = "AllVpnInterfaceLabel"
	fields["all_uplink_interface_label"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["all_uplink_interface_label"] = "AllUplinkInterfaceLabel"
	fields["domains"] = bindings.NewListType(bindings.NewReferenceType(ProviderObjectBindingType), reflect.TypeOf([]ProviderObject{}))
	fieldNameMap["domains"] = "Domains"
	fields["sddc_infra_subnet"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["sddc_infra_subnet"] = "SddcInfraSubnet"
	fields["infra_subnets"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["infra_subnets"] = "InfraSubnets"
	fields["management_gateway_label"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["management_gateway_label"] = "ManagementGatewayLabel"
	fields["vpn_internet_ips"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["vpn_internet_ips"] = "VpnInternetIps"
	fields["mgmt_subnets"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["mgmt_subnets"] = "MgmtSubnets"
	fields["enforcement_points"] = bindings.NewListType(bindings.NewReferenceType(ProviderObjectBindingType), reflect.TypeOf([]ProviderObject{}))
	fieldNameMap["enforcement_points"] = "EnforcementPoints"
	fields["compute_gateway"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["compute_gateway"] = "ComputeGateway"
	fields["vpn_dx_ips"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["vpn_dx_ips"] = "VpnDxIps"
	fields["linked_vpc_interface_label"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["linked_vpc_interface_label"] = "LinkedVpcInterfaceLabel"
	fields["dx_interface_label"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["dx_interface_label"] = "DxInterfaceLabel"
	fields["mgmt_subnet"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["mgmt_subnet"] = "MgmtSubnet"
	fields["management_gateway"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["management_gateway"] = "ManagementGateway"
	fields["gateways"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ProviderGatewayKeyValuePairsBindingType), reflect.TypeOf([]ProviderGatewayKeyValuePairs{})))
	fieldNameMap["gateways"] = "Gateways"
	fields["public_interface_label"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["public_interface_label"] = "PublicInterfaceLabel"
	fields["cgw_snat_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cgw_snat_ip"] = "CgwSnatIp"
	fields["vpn_endpoints"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(VpnEndpointBindingType), reflect.TypeOf([]VpnEndpoint{})))
	fieldNameMap["vpn_endpoints"] = "VpnEndpoints"
	fields["management_domain"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["management_domain"] = "ManagementDomain"
	fields["provider_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["provider_name"] = "ProviderName"
	fields["provider_gateways"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ProviderGatewayKeyValuePairsBindingType), reflect.TypeOf([]ProviderGatewayKeyValuePairs{})))
	fieldNameMap["provider_gateways"] = "ProviderGateways"
	fields["management_gateway_default_dns_zone"] = bindings.NewReferenceType(ProviderObjectBindingType)
	fieldNameMap["management_gateway_default_dns_zone"] = "ManagementGatewayDefaultDnsZone"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.sddc_user_configuration", fields, reflect.TypeOf(SddcUserConfiguration{}), fieldNameMap, validators)
}

func SelfResourceLinkBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["action"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["action"] = "Action"
	fields["href"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["href"] = "Href"
	fields["rel"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["rel"] = "Rel"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.self_resource_link", fields, reflect.TypeOf(SelfResourceLink{}), fieldNameMap, validators)
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

func TagBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["scope"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["scope"] = "Scope"
	fields["tag"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["tag"] = "Tag"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.tag", fields, reflect.TypeOf(Tag{}), fieldNameMap, validators)
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

func TrafficGroupBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_revision"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_system_owned"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["tags"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["_create_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_protection"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_create_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_last_modified_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	fields["parent_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["marked_for_delete"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["state_message"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["state_message"] = "StateMessage"
	fields["association_maps"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(TrafficGroupAssociationMapBindingType), reflect.TypeOf([]TrafficGroupAssociationMap{})))
	fieldNameMap["association_maps"] = "AssociationMaps"
	fields["resource_id"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["resource_id"] = "ResourceId"
	fields["state"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["state"] = "State"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.traffic_group", fields, reflect.TypeOf(TrafficGroup{}), fieldNameMap, validators)
}

func TrafficGroupAssociationMapBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_revision"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_system_owned"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["tags"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["_create_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_protection"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_create_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_last_modified_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	fields["parent_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["marked_for_delete"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["state_message"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["state_message"] = "StateMessage"
	fields["scope"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["scope"] = "Scope"
	fields["prefix_lists"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["prefix_lists"] = "PrefixLists"
	fields["state"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["state"] = "State"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.traffic_group_association_map", fields, reflect.TypeOf(TrafficGroupAssociationMap{}), fieldNameMap, validators)
}

func TrafficGroupAssociationMapsListResultBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["sort_ascending"] = "SortAscending"
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sort_by"] = "SortBy"
	fields["result_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["result_count"] = "ResultCount"
	fields["results"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(TrafficGroupAssociationMapBindingType), reflect.TypeOf([]TrafficGroupAssociationMap{})))
	fieldNameMap["results"] = "Results"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.traffic_group_association_maps_list_result", fields, reflect.TypeOf(TrafficGroupAssociationMapsListResult{}), fieldNameMap, validators)
}

func TrafficGroupsListResultBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["sort_ascending"] = "SortAscending"
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sort_by"] = "SortBy"
	fields["result_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["result_count"] = "ResultCount"
	fields["results"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(TrafficGroupBindingType), reflect.TypeOf([]TrafficGroup{})))
	fieldNameMap["results"] = "Results"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.traffic_groups_list_result", fields, reflect.TypeOf(TrafficGroupsListResult{}), fieldNameMap, validators)
}

func VMCAccountsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["linked_vpc_account"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["linked_vpc_account"] = "LinkedVpcAccount"
	fields["shadow_account"] = bindings.NewStringType()
	fieldNameMap["shadow_account"] = "ShadowAccount"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.VMC_accounts", fields, reflect.TypeOf(VMCAccounts{}), fieldNameMap, validators)
}

func VifsListResultBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["sort_ascending"] = "SortAscending"
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sort_by"] = "SortBy"
	fields["result_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["result_count"] = "ResultCount"
	fields["results"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(VirtualInterfaceBindingType), reflect.TypeOf([]VirtualInterface{})))
	fieldNameMap["results"] = "Results"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.vifs_list_result", fields, reflect.TypeOf(VifsListResult{}), fieldNameMap, validators)
}

func VirtualInterfaceBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["direct_connect_id"] = bindings.NewStringType()
	fieldNameMap["direct_connect_id"] = "DirectConnectId"
	fields["state"] = bindings.NewStringType()
	fieldNameMap["state"] = "State"
	fields["remote_asn"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["remote_asn"] = "RemoteAsn"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["local_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["local_ip"] = "LocalIp"
	fields["remote_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["remote_ip"] = "RemoteIp"
	fields["id"] = bindings.NewStringType()
	fieldNameMap["id"] = "Id"
	fields["bgp_status"] = bindings.NewStringType()
	fieldNameMap["bgp_status"] = "BgpStatus"
	fields["mtu"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["mtu"] = "Mtu"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.virtual_interface", fields, reflect.TypeOf(VirtualInterface{}), fieldNameMap, validators)
}

func VmcAppBaseInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_revision"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_system_owned"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["tags"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["_create_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_protection"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_create_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_last_modified_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	fields["parent_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["marked_for_delete"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.vmc_app_base_info", fields, reflect.TypeOf(VmcAppBaseInfo{}), fieldNameMap, validators)
}

func VmcConsolidatedRealizedStatusBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["consolidated_status"] = bindings.NewOptionalType(bindings.NewReferenceType(VmcConsolidatedStatusBindingType))
	fieldNameMap["consolidated_status"] = "ConsolidatedStatus"
	fields["consolidated_status_per_object"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(VmcConsolidatedStatusPerObjectBindingType), reflect.TypeOf([]VmcConsolidatedStatusPerObject{})))
	fieldNameMap["consolidated_status_per_object"] = "ConsolidatedStatusPerObject"
	fields["intent_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["intent_path"] = "IntentPath"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.vmc_consolidated_realized_status", fields, reflect.TypeOf(VmcConsolidatedRealizedStatus{}), fieldNameMap, validators)
}

func VmcConsolidatedStatusBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["status_message"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["status_message"] = "StatusMessage"
	fields["consolidated_status"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["consolidated_status"] = "ConsolidatedStatus"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.vmc_consolidated_status", fields, reflect.TypeOf(VmcConsolidatedStatus{}), fieldNameMap, validators)
}

func VmcConsolidatedStatusPerObjectBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["consolidated_status"] = bindings.NewOptionalType(bindings.NewReferenceType(VmcConsolidatedStatusBindingType))
	fieldNameMap["consolidated_status"] = "ConsolidatedStatus"
	fields["object_id"] = bindings.NewStringType()
	fieldNameMap["object_id"] = "ObjectId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.vmc_consolidated_status_per_object", fields, reflect.TypeOf(VmcConsolidatedStatusPerObject{}), fieldNameMap, validators)
}

func VmcFeatureFlagInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_revision"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_system_owned"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["tags"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["_create_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_protection"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_create_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_last_modified_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	fields["parent_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["marked_for_delete"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["state"] = bindings.NewStringType()
	fieldNameMap["state"] = "State"
	fields["message"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["message"] = "Message"
	fields["internal_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["internal_name"] = "InternalName"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.vmc_feature_flag_info", fields, reflect.TypeOf(VmcFeatureFlagInfo{}), fieldNameMap, validators)
}

func VmcFeatureFlagsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["features"] = bindings.NewListType(bindings.NewReferenceType(VmcFeatureFlagInfoBindingType), reflect.TypeOf([]VmcFeatureFlagInfo{}))
	fieldNameMap["features"] = "Features"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.vmc_feature_flags", fields, reflect.TypeOf(VmcFeatureFlags{}), fieldNameMap, validators)
}

func VmcRealizedEntitiesBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["realized_entities"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(VmcRealizedEntityBindingType), reflect.TypeOf([]VmcRealizedEntity{})))
	fieldNameMap["realized_entities"] = "RealizedEntities"
	fields["realized_state"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["realized_state"] = "RealizedState"
	fields["realized_entities_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["realized_entities_id"] = "RealizedEntitiesId"
	fields["intent_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["intent_path"] = "IntentPath"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.vmc_realized_entities", fields, reflect.TypeOf(VmcRealizedEntities{}), fieldNameMap, validators)
}

func VmcRealizedEntityBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["realization_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["realization_id"] = "RealizationId"
	fields["realization_api"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["realization_api"] = "RealizationApi"
	fields["realized_entity_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["realized_entity_id"] = "RealizedEntityId"
	fields["realized_state"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["realized_state"] = "RealizedState"
	fields["realization_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["realization_path"] = "RealizationPath"
	fields["realized_entity_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["realized_entity_type"] = "RealizedEntityType"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.vmc_realized_entity", fields, reflect.TypeOf(VmcRealizedEntity{}), fieldNameMap, validators)
}

func VpnEndpointBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ip"] = bindings.NewStringType()
	fieldNameMap["ip"] = "Ip"
	fields["type"] = bindings.NewStringType()
	fieldNameMap["type"] = "Type_"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["interface_label"] = bindings.NewStringType()
	fieldNameMap["interface_label"] = "InterfaceLabel"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.vpn_endpoint", fields, reflect.TypeOf(VpnEndpoint{}), fieldNameMap, validators)
}


