// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for package: com.vmware.model.
// Includes binding types of a top level structures and enumerations.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package model

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/log"
	"reflect"
)

type ActionMessage struct {
	// Optional message for action
	Message *string
}

func (s *ActionMessage) GetType__() bindings.BindingType {
	return ActionMessageBindingType()
}

func (s *ActionMessage) GetDataValue__() (data.DataValue, []error) {
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
	// Possible values are:
	//
	// * AdvertisedRoute#AdvertisedRoute_ADVERTISEMENT_STATE_SUCCESS
	// * AdvertisedRoute#AdvertisedRoute_ADVERTISEMENT_STATE_FAILED
	//
	//  State of advertisement
	AdvertisementState string
	// The route that is advertised to on-premise datacenter via Direct Connect format: ipv4-cidr-block
	Ipv4Cidr string
}

const AdvertisedRoute_ADVERTISEMENT_STATE_SUCCESS = "SUCCESS"
const AdvertisedRoute_ADVERTISEMENT_STATE_FAILED = "FAILED"

func (s *AdvertisedRoute) GetType__() bindings.BindingType {
	return AdvertisedRouteBindingType()
}

func (s *AdvertisedRoute) GetDataValue__() (data.DataValue, []error) {
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
	// Further details about the error
	Details *string
	// A numeric error code format: int64
	ErrorCode *int64
	// Additional data about the error
	ErrorData *data.StructValue
	// A description of the error
	ErrorMessage *string
	// The module name where the error occurred
	ModuleName *string
	// Other errors related to this error
	RelatedErrors []RelatedApiError
}

func (s *ApiError) GetType__() bindings.BindingType {
	return ApiErrorBindingType()
}

func (s *ApiError) GetDataValue__() (data.DataValue, []error) {
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
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	// Link to this resource
	Self *SelfResourceLink
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
	Id string
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

func (s *AssociatedBaseGroupConnectionInfo) GetType__() bindings.BindingType {
	return AssociatedBaseGroupConnectionInfoBindingType()
}

func (s *AssociatedBaseGroupConnectionInfo) GetDataValue__() (data.DataValue, []error) {
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
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	// Link to this resource
	Self *SelfResourceLink
	// Opaque cursor to be used for getting next page of records (supplied by current result page)
	Cursor *string
	// Count of results found (across all pages), set only on first page format: int64
	ResultCount *int64
	// If true, results are sorted in ascending order
	SortAscending *bool
	// Field by which records are sorted
	SortBy  *string
	Results []*data.StructValue
}

func (s *AssociatedGroupConnectionInfosListResult) GetType__() bindings.BindingType {
	return AssociatedGroupConnectionInfosListResultBindingType()
}

func (s *AssociatedGroupConnectionInfosListResult) GetDataValue__() (data.DataValue, []error) {
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
	// TGW external route table ID used for external customers VPCs association
	ExternalRouteTableId string
	// TGW SDDCs route table ID used for SDDCs association
	SddcsRouteTableId string
	// Possible values are:
	//
	// * AssociatedTgwGroupConnectionInfo#AssociatedTgwGroupConnectionInfo_STATE_CONNECTED
	// * AssociatedTgwGroupConnectionInfo#AssociatedTgwGroupConnectionInfo_STATE_DISCONNECTED
	//
	//  The TGW attachment state of the SDDC in the Group
	State *string
	// TGW attachment ID for the local SDDC in the Group
	TgwAttachmentId string
	// TGW ID for the local SDDC in the Group
	TgwId string
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	// Link to this resource
	Self *SelfResourceLink
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
	Id string
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

func (s *AssociatedTgwGroupConnectionInfo) GetType__() bindings.BindingType {
	return AssociatedTgwGroupConnectionInfoBindingType()
}

func (s *AssociatedTgwGroupConnectionInfo) GetDataValue__() (data.DataValue, []error) {
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

// Associated Group connection infomation for the local SDDC by using AWS TGW as a connector and AWS prefix list as route management.
type AssociatedTgwGroupConnectionInfoVersion2 struct {
	// TGW SDDCs route table ID used for SDDCs association
	AssociatedRouteTable string
	// AWS region in which SDDC is residing
	Region string
	// SDDC advertise route configuration with AWS prefix list and other region TGW connections.
	SddcAdvertiseRouteConfig SddcAdvertiseRouteConfig
	// Possible values are:
	//
	// * AssociatedTgwGroupConnectionInfoVersion2#AssociatedTgwGroupConnectionInfoVersion2_STATE_CONNECTED
	// * AssociatedTgwGroupConnectionInfoVersion2#AssociatedTgwGroupConnectionInfoVersion2_STATE_DISCONNECTED
	//
	//  The TGW attachment state of the SDDC in the Group
	State *string
	// TGW attachment ID for the local SDDC in the Group
	TgwAttachmentId string
	// TGW ID for the local SDDC in the Group
	TgwId string
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	// Link to this resource
	Self *SelfResourceLink
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
	Id string
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

func (s *AssociatedTgwGroupConnectionInfoVersion2) GetType__() bindings.BindingType {
	return AssociatedTgwGroupConnectionInfoVersion2BindingType()
}

func (s *AssociatedTgwGroupConnectionInfoVersion2) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for AssociatedTgwGroupConnectionInfoVersion2._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// AWS prefix entry
type AwsPrefixInfo struct {
	// Prefix ID
	Id *string
	// Prefix Name
	Name *string
	// AWS region
	Region *string
	// Prefix max entries format: int32
	Size *int64
}

func (s *AwsPrefixInfo) GetType__() bindings.BindingType {
	return AwsPrefixInfoBindingType()
}

func (s *AwsPrefixInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for AwsPrefixInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
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

func (s *BGPAdvertisedRoutes) GetType__() bindings.BindingType {
	return BGPAdvertisedRoutesBindingType()
}

func (s *BGPAdvertisedRoutes) GetDataValue__() (data.DataValue, []error) {
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

func (s *BGPLearnedRoutes) GetType__() bindings.BindingType {
	return BGPLearnedRoutesBindingType()
}

func (s *BGPLearnedRoutes) GetDataValue__() (data.DataValue, []error) {
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
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	// Link to this resource
	Self *SelfResourceLink
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

func (s *ConnectedServiceListResult) GetType__() bindings.BindingType {
	return ConnectedServiceListResultBindingType()
}

func (s *ConnectedServiceListResult) GetDataValue__() (data.DataValue, []error) {
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
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	// Link to this resource
	Self *SelfResourceLink
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

func (s *ConnectedServiceStatus) GetType__() bindings.BindingType {
	return ConnectedServiceStatusBindingType()
}

func (s *ConnectedServiceStatus) GetDataValue__() (data.DataValue, []error) {
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

func (s *CsvListResult) GetType__() bindings.BindingType {
	return CsvListResultBindingType()
}

func (s *CsvListResult) GetDataValue__() (data.DataValue, []error) {
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

func (s *CsvRecord) GetType__() bindings.BindingType {
	return CsvRecordBindingType()
}

func (s *CsvRecord) GetDataValue__() (data.DataValue, []error) {
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
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	// Link to this resource
	Self *SelfResourceLink
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

func (s *DirectConnectBgpInfo) GetType__() bindings.BindingType {
	return DirectConnectBgpInfoBindingType()
}

func (s *DirectConnectBgpInfo) GetDataValue__() (data.DataValue, []error) {
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

// External Connectivity configuration for north-south traffic. For eg., in AWS case, this would refer to Direct Connect config. For Dimension, this would refer to the config at Upstream Intranet interface to Tor.
type ExternalConnectivityConfig struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	// Link to this resource
	Self *SelfResourceLink
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
	// Configuration for Internet Interface.
	InternetInterface *InterfaceConfig
	// Uplink MTU of internet traffic in edge tier-0 router port. format: int32
	InternetMtu *int64
	// Configuration for Intranet Interface.
	IntranetInterface *InterfaceConfig
	// Uplink MTU of direct connect, sddc-grouping and outposts traffic in edge tier-0 router port. format: int32
	IntranetMtu int64
	// Configuration for Services Interface.
	ServicesInterface *InterfaceConfig
	// Uplink MTU of connected VPC traffic in edge tier-0 router port. format: int32
	ServicesMtu *int64
}

func (s *ExternalConnectivityConfig) GetType__() bindings.BindingType {
	return ExternalConnectivityConfigBindingType()
}

func (s *ExternalConnectivityConfig) GetDataValue__() (data.DataValue, []error) {
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
	// * ExternalSddcConnectivity#ExternalSddcConnectivity_CONNECTIVITY_TYPE_DIRECT_CONNECT
	// * ExternalSddcConnectivity#ExternalSddcConnectivity_CONNECTIVITY_TYPE_DEPLOYMENT_CONNECTIVITY_GROUP
	//
	//  The external SDDC connectivity type is used by the SDDC for the L3 connectivity. DIRECT_CONNECT means that the external SDDC connectivity is through AWS Direct Connect. DEPLOYMENT_CONNECTIVITY_GROUP means that the external SDDC connectivity is through AWS TGW.
	ConnectivityType string
	// The list of regions of the route for the connectivity
	Regions []string
	// The source of the route for the connectivity
	Source *string
	// Possible values are:
	//
	// * ExternalSddcConnectivity#ExternalSddcConnectivity_STATUS_SUCCEEDED
	// * ExternalSddcConnectivity#ExternalSddcConnectivity_STATUS_FAILED
	//
	//  The status of the route for the connectivity
	Status *string
	// The error message if the status is FAILED, the optional warning message if the status is SUCCEEDED.
	StatusMessage *string
}

const ExternalSddcConnectivity_CONNECTIVITY_TYPE_DIRECT_CONNECT = "DIRECT_CONNECT"
const ExternalSddcConnectivity_CONNECTIVITY_TYPE_DEPLOYMENT_CONNECTIVITY_GROUP = "DEPLOYMENT_CONNECTIVITY_GROUP"
const ExternalSddcConnectivity_STATUS_SUCCEEDED = "SUCCEEDED"
const ExternalSddcConnectivity_STATUS_FAILED = "FAILED"

func (s *ExternalSddcConnectivity) GetType__() bindings.BindingType {
	return ExternalSddcConnectivityBindingType()
}

func (s *ExternalSddcConnectivity) GetDataValue__() (data.DataValue, []error) {
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
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	// Link to this resource
	Self *SelfResourceLink
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
	Destination string
}

func (s *ExternalSddcRoute) GetType__() bindings.BindingType {
	return ExternalSddcRouteBindingType()
}

func (s *ExternalSddcRoute) GetDataValue__() (data.DataValue, []error) {
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
	// The connectivity datails contains status of route, source of the route, connectivity type
	ConnectivityDetails string
	// Destination IP CIDR Block format: ipv4-cidr-block
	Destination string
}

func (s *ExternalSddcRouteCsvRecord) GetType__() bindings.BindingType {
	return ExternalSddcRouteCsvRecordBindingType()
}

func (s *ExternalSddcRouteCsvRecord) GetDataValue__() (data.DataValue, []error) {
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
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	// Link to this resource
	Self *SelfResourceLink
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

func (s *ExternalSddcRoutesListResult) GetType__() bindings.BindingType {
	return ExternalSddcRoutesListResultBindingType()
}

func (s *ExternalSddcRoutesListResult) GetDataValue__() (data.DataValue, []error) {
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
	Results  []ExternalSddcRouteCsvRecord
}

func (s *ExternalSddcRoutesListResultInCsvFormat) GetType__() bindings.BindingType {
	return ExternalSddcRoutesListResultInCsvFormatBindingType()
}

func (s *ExternalSddcRoutesListResultInCsvFormat) GetDataValue__() (data.DataValue, []error) {
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

func (s *IncludedFieldsParameters) GetType__() bindings.BindingType {
	return IncludedFieldsParametersBindingType()
}

func (s *IncludedFieldsParameters) GetDataValue__() (data.DataValue, []error) {
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

// Interface configuration.
type InterfaceConfig struct {
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

func (s *InterfaceConfig) GetType__() bindings.BindingType {
	return InterfaceConfigBindingType()
}

func (s *InterfaceConfig) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for InterfaceConfig._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type IpAttachmentPair struct {
	// Attachment id which maps to management VM IP
	AttachmentId string
	// Management VM IP Address format: ipv4
	Ip string
}

func (s *IpAttachmentPair) GetType__() bindings.BindingType {
	return IpAttachmentPairBindingType()
}

func (s *IpAttachmentPair) GetDataValue__() (data.DataValue, []error) {
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
	// Linked subnet availability zone
	AvailabilityZone string
	// Linked subnet CIDR format: ipv4-cidr-block
	Cidr string
	// Linked subnet identifier
	Id string
}

func (s *LinkedSubnetInfo) GetType__() bindings.BindingType {
	return LinkedSubnetInfoBindingType()
}

func (s *LinkedSubnetInfo) GetDataValue__() (data.DataValue, []error) {
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
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	// Link to this resource
	Self *SelfResourceLink
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
	ArnRole string
	// External identifier for ARN role
	ExternalId string
	// Linked VPC account number
	LinkedAccount string
	// Linked VPC CIDRs format: ipv4-cidr-block
	LinkedVpcAddresses []string
	// Linked VPC identifier
	LinkedVpcId *string
	// The IPs of linked VPC NAT rule for service access. format: ipv4
	LinkedVpcNatIps []string
	// Infromation related to the subnets where linked ENIs were created.
	LinkedVpcSubnets []LinkedSubnetInfo
	// The identifiers of route tables to be dynamically updated with SDDC networks
	RouteTableIds []string
	// service ARN role
	ServiceArnRole *string
}

func (s *LinkedVpcInfo) GetType__() bindings.BindingType {
	return LinkedVpcInfoBindingType()
}

func (s *LinkedVpcInfo) GetDataValue__() (data.DataValue, []error) {
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
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	// Link to this resource
	Self *SelfResourceLink
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

func (s *LinkedVpcsListResult) GetType__() bindings.BindingType {
	return LinkedVpcsListResultBindingType()
}

func (s *LinkedVpcsListResult) GetDataValue__() (data.DataValue, []error) {
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
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	// Link to this resource
	Self *SelfResourceLink
	// Opaque cursor to be used for getting next page of records (supplied by current result page)
	Cursor *string
	// Count of results found (across all pages), set only on first page format: int64
	ResultCount *int64
	// If true, results are sorted in ascending order
	SortAscending *bool
	// Field by which records are sorted
	SortBy *string
}

func (s *ListResult) GetType__() bindings.BindingType {
	return ListResultBindingType()
}

func (s *ListResult) GetDataValue__() (data.DataValue, []error) {
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
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	// Link to this resource
	Self *SelfResourceLink
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

func (s *ManagedResource) GetType__() bindings.BindingType {
	return ManagedResourceBindingType()
}

func (s *ManagedResource) GetDataValue__() (data.DataValue, []error) {
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
	// Display name for this service
	DisplayName *string
	// Service path should refer to a valid service in the system. Service can be system defined or user defined.
	Path *string
}

func (s *MgmtServiceEntry) GetType__() bindings.BindingType {
	return MgmtServiceEntryBindingType()
}

func (s *MgmtServiceEntry) GetDataValue__() (data.DataValue, []error) {
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
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	// Link to this resource
	Self *SelfResourceLink
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

func (s *MgmtVmInfo) GetType__() bindings.BindingType {
	return MgmtVmInfoBindingType()
}

func (s *MgmtVmInfo) GetDataValue__() (data.DataValue, []error) {
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
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	// Link to this resource
	Self *SelfResourceLink
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

func (s *MgmtVmsListResult) GetType__() bindings.BindingType {
	return MgmtVmsListResultBindingType()
}

func (s *MgmtVmsListResult) GetDataValue__() (data.DataValue, []error) {
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
	// Identifier of the Interface label
	Id string
	// Name of the Interface label
	Name *string
}

func (s *ModelInterface) GetType__() bindings.BindingType {
	return ModelInterfaceBindingType()
}

func (s *ModelInterface) GetDataValue__() (data.DataValue, []error) {
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

// Other region TGW connection info
type OtherRegionConnectionInfo struct {
	// AWS region
	Region string
	// TGW ID
	TgwId string
	// TGW peering id
	TgwPeeringId string
}

func (s *OtherRegionConnectionInfo) GetType__() bindings.BindingType {
	return OtherRegionConnectionInfoBindingType()
}

func (s *OtherRegionConnectionInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for OtherRegionConnectionInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Prefix lists used by certain features, like Intranet NAT.
type PrefixListInfo struct {
	// Prefix List name
	Name string
	// Relative Prefix List path
	Path string
	// Prefix List URL
	Url string
}

func (s *PrefixListInfo) GetType__() bindings.BindingType {
	return PrefixListInfoBindingType()
}

func (s *PrefixListInfo) GetDataValue__() (data.DataValue, []error) {
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
	// Key
	Key string
	// Value
	Value []ProviderObject
}

func (s *ProviderGatewayKeyValuePairs) GetType__() bindings.BindingType {
	return ProviderGatewayKeyValuePairsBindingType()
}

func (s *ProviderGatewayKeyValuePairs) GetDataValue__() (data.DataValue, []error) {
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

func (s *ProviderObject) GetType__() bindings.BindingType {
	return ProviderObjectBindingType()
}

func (s *ProviderObject) GetDataValue__() (data.DataValue, []error) {
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
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	// Link to this resource
	Self *SelfResourceLink
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

func (s *PublicIp) GetType__() bindings.BindingType {
	return PublicIpBindingType()
}

func (s *PublicIp) GetDataValue__() (data.DataValue, []error) {
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
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	// Link to this resource
	Self *SelfResourceLink
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

func (s *PublicIpsListResult) GetType__() bindings.BindingType {
	return PublicIpsListResultBindingType()
}

func (s *PublicIpsListResult) GetDataValue__() (data.DataValue, []error) {
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
	// Further details about the error
	Details *string
	// A numeric error code format: int64
	ErrorCode *int64
	// Additional data about the error
	ErrorData *data.StructValue
	// A description of the error
	ErrorMessage *string
	// The module name where the error occurred
	ModuleName *string
}

func (s *RelatedApiError) GetType__() bindings.BindingType {
	return RelatedApiErrorBindingType()
}

func (s *RelatedApiError) GetDataValue__() (data.DataValue, []error) {
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
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	// Link to this resource
	Self *SelfResourceLink
}

func (s *Resource) GetType__() bindings.BindingType {
	return ResourceBindingType()
}

func (s *Resource) GetDataValue__() (data.DataValue, []error) {
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

func (s *ResourceLink) GetType__() bindings.BindingType {
	return ResourceLinkBindingType()
}

func (s *ResourceLink) GetDataValue__() (data.DataValue, []error) {
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
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	// Link to this resource
	Self *SelfResourceLink
	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected. format: int32
	Revision *int64
}

func (s *RevisionedResource) GetType__() bindings.BindingType {
	return RevisionedResourceBindingType()
}

func (s *RevisionedResource) GetDataValue__() (data.DataValue, []error) {
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

// SDDC advertise route configuration with AWS prefix list and other region TGW connections.
type SddcAdvertiseRouteConfig struct {
	// AWS prefix list for managing TGW routes
	AwsPrefixList []AwsPrefixInfo
	// List of TGW connection information from other regions to local region TGW
	OtherRegionConnections []OtherRegionConnectionInfo
}

func (s *SddcAdvertiseRouteConfig) GetType__() bindings.BindingType {
	return SddcAdvertiseRouteConfigBindingType()
}

func (s *SddcAdvertiseRouteConfig) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SddcAdvertiseRouteConfig._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// SDDC configuration parameters for users. User-level addresses/CIDRs are provided.
type SddcUserConfiguration struct {
	// All uplink interfaces label name. Only presents in AWS environment now. Deprecated, please use labels.
	AllUplinkInterfaceLabel *string
	// All VPN interfaces label name. Only presents in AWS environment now. Deprecated, please use labels.
	AllVpnInterfaceLabel *string
	// Compute gateway SNAT IP address. format: ipv4
	CgwSnatIp *string
	// Compute domain ID. Only presents in AWS environment now. Deprecated, please use domains.
	ComputeDomain *string
	// Compute gateway name. Only presents in AWS environment now. Deprecated, please use gateways.
	ComputeGateway *string
	// Domain information for management gateway and compute gateway.
	Domains []ProviderObject
	// DirectConnect interface label name. Only presents in AWS environment now. Deprecated, please use labels.
	DxInterfaceLabel *string
	// Policy enforcement point object paths.
	EnforcementPoints []ProviderObject
	// Provider gateway list. Including both tier-0 gateways and tier-1 gateways.
	Gateways []ProviderGatewayKeyValuePairs
	// SDDC infra subnet CIDRs. format: ipv4-cidr-block
	InfraSubnets []string
	// Interfaces (labels) including public interface, direct connect interface, linked vpc interface, etc. Only presents in AWS environment now. Deprecated, please use labels.
	Interfaces []ModelInterface
	// Interfaces (labels) including internet, intranet and services.
	Labels []ProviderObject
	// Linked VPC interface label name. Only presents in AWS environment now. Deprecated, please use labels.
	LinkedVpcInterfaceLabel *string
	// Management domain ID. Only presents in AWS environment now. Deprecated, please use domains.
	ManagementDomain *string
	// Management gateway name. Only presents in AWS environment now. Deprecated, please use gateways.
	ManagementGateway *string
	// Management gateway default DNS zone
	ManagementGatewayDefaultDnsZone ProviderObject
	// Management gateway label name. Only presents in AWS environment now. Deprecated, please use labels.
	ManagementGatewayLabel *string
	// Management subnet CIDRs. Only presents in AWS environment now. Deprecated, please use mgmt_subnets. format: ipv4-cidr-block
	MgmtSubnet []string
	// Management subnet CIDRs. format: ipv4-cidr-block
	MgmtSubnets []string
	// Management gateway SNAT IP address. format: ipv4
	MgwSnatIp *string
	// Provider gateway list. Including both tier-0 gateways and tier-1 gateways. Only presents in AWS environment now. Deprecated, please use gateways.
	ProviderGateways []ProviderGatewayKeyValuePairs
	// Service provider name. Only presents in AWS environment now. Deprecated, please use gateways.
	ProviderName *string
	// Prefix lists used by certain features, like Intranet NAT.
	ProviderPrefixLists []PrefixListInfo
	// Public interface label name. Only presents in AWS environment now. Deprecated, please use labels.
	PublicInterfaceLabel *string
	// SDDC infra subnet CIDRs. Only presents in AWS environment now. Deprecated, please use infra_subnets. format: ipv4-cidr-block
	SddcInfraSubnet []string
	// Local IPs for VPN tunnel over Direct Connect. Only presents in AWS environment now. Deprecated, please use vpn_endpoints instead of vpn_dx_ips. format: ipv4
	VpnDxIps []string
	// VPN tunnel endpoints. Currently containing public IPs for VPN over internet and local IPs for VPN over intranet.
	VpnEndpoints []VpnEndpoint
	// Public IPs for VPN tunnel over internet. Only presents in AWS environment now. Deprecated, please use vpn_endpoints instead of vpn_internet_ips. format: ipv4
	VpnInternetIps []string
}

func (s *SddcUserConfiguration) GetType__() bindings.BindingType {
	return SddcUserConfigurationBindingType()
}

func (s *SddcUserConfiguration) GetDataValue__() (data.DataValue, []error) {
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

func (s *SelfResourceLink) GetType__() bindings.BindingType {
	return SelfResourceLinkBindingType()
}

func (s *SelfResourceLink) GetDataValue__() (data.DataValue, []error) {
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

// Arbitrary key-value pairs that may be attached to an entity
type Tag struct {
	// Tag searches may optionally be restricted by scope
	Scope *string
	// Identifier meaningful to user with maximum length of 256 characters
	Tag *string
}

func (s *Tag) GetType__() bindings.BindingType {
	return TagBindingType()
}

func (s *Tag) GetDataValue__() (data.DataValue, []error) {
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

// Traffic group configuration. A traffic group indicates dedicated bandwidth and computation for a given list of subnets. Creating a traffic group will reserve resources and associating it with desired prefix lists will use the resources for the traffic of the prefix lists. Besides traffic group ID and name, realized state and detailed association maps info of this traffic group are included if verbose info is requested.
type TrafficGroup struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	// Link to this resource
	Self *SelfResourceLink
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
}

const TrafficGroup_STATE_IN_PROGRESS = "IN_PROGRESS"
const TrafficGroup_STATE_SUCCESS = "SUCCESS"
const TrafficGroup_STATE_FAILURE = "FAILURE"
const TrafficGroup_STATE_UNAVAILABLE = "UNAVAILABLE"
const TrafficGroup_STATE_PENDING = "PENDING"

func (s *TrafficGroup) GetType__() bindings.BindingType {
	return TrafficGroupBindingType()
}

func (s *TrafficGroup) GetDataValue__() (data.DataValue, []error) {
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
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	// Link to this resource
	Self *SelfResourceLink
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
	// The list of prefix lists to be associated.
	PrefixLists []string
	// Possible values are:
	//
	// * TrafficGroupAssociationMap#TrafficGroupAssociationMap_SCOPE_1S_CGW
	// * TrafficGroupAssociationMap#TrafficGroupAssociationMap_SCOPE_0S_VMC
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
const TrafficGroupAssociationMap_STATE_IN_PROGRESS = "IN_PROGRESS"
const TrafficGroupAssociationMap_STATE_SUCCESS = "SUCCESS"
const TrafficGroupAssociationMap_STATE_FAILURE = "FAILURE"
const TrafficGroupAssociationMap_STATE_UNAVAILABLE = "UNAVAILABLE"
const TrafficGroupAssociationMap_STATE_PENDING = "PENDING"

func (s *TrafficGroupAssociationMap) GetType__() bindings.BindingType {
	return TrafficGroupAssociationMapBindingType()
}

func (s *TrafficGroupAssociationMap) GetDataValue__() (data.DataValue, []error) {
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
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	// Link to this resource
	Self *SelfResourceLink
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

func (s *TrafficGroupAssociationMapsListResult) GetType__() bindings.BindingType {
	return TrafficGroupAssociationMapsListResultBindingType()
}

func (s *TrafficGroupAssociationMapsListResult) GetDataValue__() (data.DataValue, []error) {
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
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	// Link to this resource
	Self *SelfResourceLink
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

func (s *TrafficGroupsListResult) GetType__() bindings.BindingType {
	return TrafficGroupsListResultBindingType()
}

func (s *TrafficGroupsListResult) GetDataValue__() (data.DataValue, []error) {
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

// Upgrade versions information
type UpgradeVersionInfo struct {
	// The version of SDDC on which the upgrade is started.
	FromSddcVersion *string
	// The version to which the SDDC is upgraded. port.
	ToSddcVersion *string
}

func (s *UpgradeVersionInfo) GetType__() bindings.BindingType {
	return UpgradeVersionInfoBindingType()
}

func (s *UpgradeVersionInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for UpgradeVersionInfo._GetDataValue method - %s",
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

func (s *VMCAccounts) GetType__() bindings.BindingType {
	return VMCAccountsBindingType()
}

func (s *VMCAccounts) GetDataValue__() (data.DataValue, []error) {
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
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	// Link to this resource
	Self *SelfResourceLink
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

func (s *VifsListResult) GetType__() bindings.BindingType {
	return VifsListResultBindingType()
}

func (s *VifsListResult) GetDataValue__() (data.DataValue, []error) {
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
	// Possible values are:
	//
	// * VirtualInterface#VirtualInterface_BGP_STATUS_UP
	// * VirtualInterface#VirtualInterface_BGP_STATUS_DOWN
	//
	//  BGP status
	BgpStatus string
	// Identifier for the Direct Connect
	DirectConnectId string
	// Identifier for the virtual interface
	Id string
	// amazon side address format: ipv4
	LocalIp *string
	// Maximum transmission unit allowed by the VIF format: int32
	Mtu *int64
	// VIF name
	Name string
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
	State string
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

func (s *VirtualInterface) GetType__() bindings.BindingType {
	return VirtualInterfaceBindingType()
}

func (s *VirtualInterface) GetDataValue__() (data.DataValue, []error) {
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
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	// Link to this resource
	Self *SelfResourceLink
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

func (s *VmcAppBaseInfo) GetType__() bindings.BindingType {
	return VmcAppBaseInfoBindingType()
}

func (s *VmcAppBaseInfo) GetDataValue__() (data.DataValue, []error) {
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
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	// Link to this resource
	Self *SelfResourceLink
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
	// Consolidated state of objects for a given intent entity.
	ConsolidatedStatus *VmcConsolidatedStatus
	// Aggregated consolidated status by enforcement point.
	ConsolidatedStatusPerObject []VmcConsolidatedStatusPerObject
	// Intent path of the object representing this consolidated state.
	IntentPath *string
}

func (s *VmcConsolidatedRealizedStatus) GetType__() bindings.BindingType {
	return VmcConsolidatedRealizedStatusBindingType()
}

func (s *VmcConsolidatedRealizedStatus) GetDataValue__() (data.DataValue, []error) {
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

func (s *VmcConsolidatedStatus) GetType__() bindings.BindingType {
	return VmcConsolidatedStatusBindingType()
}

func (s *VmcConsolidatedStatus) GetDataValue__() (data.DataValue, []error) {
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

func (s *VmcConsolidatedStatusPerObject) GetType__() bindings.BindingType {
	return VmcConsolidatedStatusPerObjectBindingType()
}

func (s *VmcConsolidatedStatusPerObject) GetDataValue__() (data.DataValue, []error) {
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
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
	// Schema for this resource
	Schema *string
	// Link to this resource
	Self *SelfResourceLink
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
	// Internal Name
	InternalName *string
	// Message
	Message *string
	// Feature Name
	Name string
	// Possible values are:
	//
	// * VmcFeatureFlagInfo#VmcFeatureFlagInfo_STATE_ENABLED
	// * VmcFeatureFlagInfo#VmcFeatureFlagInfo_STATE_DISABLED
	// * VmcFeatureFlagInfo#VmcFeatureFlagInfo_STATE_INACTIVE
	//
	//  state
	State string
	// Unlicensed Message
	UnlicensedMessage *string
}

const VmcFeatureFlagInfo_STATE_ENABLED = "enabled"
const VmcFeatureFlagInfo_STATE_DISABLED = "disabled"
const VmcFeatureFlagInfo_STATE_INACTIVE = "inactive"

func (s *VmcFeatureFlagInfo) GetType__() bindings.BindingType {
	return VmcFeatureFlagInfoBindingType()
}

func (s *VmcFeatureFlagInfo) GetDataValue__() (data.DataValue, []error) {
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

func (s *VmcFeatureFlags) GetType__() bindings.BindingType {
	return VmcFeatureFlagsBindingType()
}

func (s *VmcFeatureFlags) GetDataValue__() (data.DataValue, []error) {
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

func (s *VmcRealizedEntities) GetType__() bindings.BindingType {
	return VmcRealizedEntitiesBindingType()
}

func (s *VmcRealizedEntities) GetDataValue__() (data.DataValue, []error) {
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

func (s *VmcRealizedEntity) GetType__() bindings.BindingType {
	return VmcRealizedEntityBindingType()
}

func (s *VmcRealizedEntity) GetDataValue__() (data.DataValue, []error) {
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
	// Interface label of the VPN endpoint
	InterfaceLabel string
	// IP address of the VPN endpoint format: ipv4
	Ip string
	// Name of the VPN endpoint
	Name string
	// Type of the VPN endpoint
	Type_ string
}

func (s *VpnEndpoint) GetType__() bindings.BindingType {
	return VpnEndpointBindingType()
}

func (s *VpnEndpoint) GetDataValue__() (data.DataValue, []error) {
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
	fields["advertisement_state"] = bindings.NewStringType()
	fieldNameMap["advertisement_state"] = "AdvertisementState"
	fields["ipv4_cidr"] = bindings.NewStringType()
	fieldNameMap["ipv4_cidr"] = "Ipv4Cidr"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.advertised_route", fields, reflect.TypeOf(AdvertisedRoute{}), fieldNameMap, validators)
}

func ApiErrorBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["details"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["details"] = "Details"
	fields["error_code"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["error_code"] = "ErrorCode"
	fields["error_data"] = bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
	fieldNameMap["error_data"] = "ErrorData"
	fields["error_message"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["error_message"] = "ErrorMessage"
	fields["module_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["module_name"] = "ModuleName"
	fields["related_errors"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(RelatedApiErrorBindingType), reflect.TypeOf([]RelatedApiError{})))
	fieldNameMap["related_errors"] = "RelatedErrors"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.api_error", fields, reflect.TypeOf(ApiError{}), fieldNameMap, validators)
}

func AssociatedBaseGroupConnectionInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_revision"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_create_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_create_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_last_modified_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["_protection"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_system_owned"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["id"] = bindings.NewStringType()
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = bindings.NewStringType()
	fieldNameMap["resource_type"] = "ResourceType"
	fields["tags"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["marked_for_delete"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["parent_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.associated_base_group_connection_info", fields, reflect.TypeOf(AssociatedBaseGroupConnectionInfo{}), fieldNameMap, validators)
}

func AssociatedGroupConnectionInfosListResultBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fields["result_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["result_count"] = "ResultCount"
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["sort_ascending"] = "SortAscending"
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sort_by"] = "SortBy"
	fields["results"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(AssociatedBaseGroupConnectionInfoBindingType)}, bindings.JSONRPC), reflect.TypeOf([]*data.StructValue{})))
	fieldNameMap["results"] = "Results"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.associated_group_connection_infos_list_result", fields, reflect.TypeOf(AssociatedGroupConnectionInfosListResult{}), fieldNameMap, validators)
}

func AssociatedTgwGroupConnectionInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["external_route_table_id"] = bindings.NewStringType()
	fieldNameMap["external_route_table_id"] = "ExternalRouteTableId"
	fields["sddcs_route_table_id"] = bindings.NewStringType()
	fieldNameMap["sddcs_route_table_id"] = "SddcsRouteTableId"
	fields["state"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["state"] = "State"
	fields["tgw_attachment_id"] = bindings.NewStringType()
	fieldNameMap["tgw_attachment_id"] = "TgwAttachmentId"
	fields["tgw_id"] = bindings.NewStringType()
	fieldNameMap["tgw_id"] = "TgwId"
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_revision"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_create_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_create_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_last_modified_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["_protection"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_system_owned"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["id"] = bindings.NewStringType()
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = bindings.NewStringType()
	fieldNameMap["resource_type"] = "ResourceType"
	fields["tags"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["marked_for_delete"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["parent_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.associated_tgw_group_connection_info", fields, reflect.TypeOf(AssociatedTgwGroupConnectionInfo{}), fieldNameMap, validators)
}

func AssociatedTgwGroupConnectionInfoVersion2BindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["associated_route_table"] = bindings.NewStringType()
	fieldNameMap["associated_route_table"] = "AssociatedRouteTable"
	fields["region"] = bindings.NewStringType()
	fieldNameMap["region"] = "Region"
	fields["sddc_advertise_route_config"] = bindings.NewReferenceType(SddcAdvertiseRouteConfigBindingType)
	fieldNameMap["sddc_advertise_route_config"] = "SddcAdvertiseRouteConfig"
	fields["state"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["state"] = "State"
	fields["tgw_attachment_id"] = bindings.NewStringType()
	fieldNameMap["tgw_attachment_id"] = "TgwAttachmentId"
	fields["tgw_id"] = bindings.NewStringType()
	fieldNameMap["tgw_id"] = "TgwId"
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_revision"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_create_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_create_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_last_modified_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["_protection"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_system_owned"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["id"] = bindings.NewStringType()
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = bindings.NewStringType()
	fieldNameMap["resource_type"] = "ResourceType"
	fields["tags"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["marked_for_delete"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["parent_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.associated_tgw_group_connection_info_version2", fields, reflect.TypeOf(AssociatedTgwGroupConnectionInfoVersion2{}), fieldNameMap, validators)
}

func AwsPrefixInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["region"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["region"] = "Region"
	fields["size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["size"] = "Size"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.aws_prefix_info", fields, reflect.TypeOf(AwsPrefixInfo{}), fieldNameMap, validators)
}

func BGPAdvertisedRoutesBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["advertised_routes"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(AdvertisedRouteBindingType), reflect.TypeOf([]AdvertisedRoute{})))
	fieldNameMap["advertised_routes"] = "AdvertisedRoutes"
	fields["failed_advertised_routes"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["failed_advertised_routes"] = "FailedAdvertisedRoutes"
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
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fields["result_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["result_count"] = "ResultCount"
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["sort_ascending"] = "SortAscending"
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sort_by"] = "SortBy"
	fields["results"] = bindings.NewListType(bindings.NewReferenceType(ConnectedServiceStatusBindingType), reflect.TypeOf([]ConnectedServiceStatus{}))
	fieldNameMap["results"] = "Results"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.connected_service_list_result", fields, reflect.TypeOf(ConnectedServiceListResult{}), fieldNameMap, validators)
}

func ConnectedServiceStatusBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_revision"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_create_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_create_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_last_modified_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["_protection"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_system_owned"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["tags"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["marked_for_delete"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["parent_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
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
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_revision"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_create_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_create_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_last_modified_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["_protection"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_system_owned"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["tags"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["marked_for_delete"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["parent_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	fields["local_as_num"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["local_as_num"] = "LocalAsNum"
	fields["mtu"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["mtu"] = "Mtu"
	fields["route_preference"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["route_preference"] = "RoutePreference"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.direct_connect_bgp_info", fields, reflect.TypeOf(DirectConnectBgpInfo{}), fieldNameMap, validators)
}

func ExternalConnectivityConfigBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_revision"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_create_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_create_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_last_modified_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["_protection"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_system_owned"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["tags"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["marked_for_delete"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["parent_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	fields["internet_interface"] = bindings.NewOptionalType(bindings.NewReferenceType(InterfaceConfigBindingType))
	fieldNameMap["internet_interface"] = "InternetInterface"
	fields["internet_mtu"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["internet_mtu"] = "InternetMtu"
	fields["intranet_interface"] = bindings.NewOptionalType(bindings.NewReferenceType(InterfaceConfigBindingType))
	fieldNameMap["intranet_interface"] = "IntranetInterface"
	fields["intranet_mtu"] = bindings.NewIntegerType()
	fieldNameMap["intranet_mtu"] = "IntranetMtu"
	fields["services_interface"] = bindings.NewOptionalType(bindings.NewReferenceType(InterfaceConfigBindingType))
	fieldNameMap["services_interface"] = "ServicesInterface"
	fields["services_mtu"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["services_mtu"] = "ServicesMtu"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.external_connectivity_config", fields, reflect.TypeOf(ExternalConnectivityConfig{}), fieldNameMap, validators)
}

func ExternalSddcConnectivityBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["connectivity_type"] = bindings.NewStringType()
	fieldNameMap["connectivity_type"] = "ConnectivityType"
	fields["regions"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["regions"] = "Regions"
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["source"] = "Source"
	fields["status"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["status"] = "Status"
	fields["status_message"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["status_message"] = "StatusMessage"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.external_sddc_connectivity", fields, reflect.TypeOf(ExternalSddcConnectivity{}), fieldNameMap, validators)
}

func ExternalSddcRouteBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_revision"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_create_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_create_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_last_modified_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["_protection"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_system_owned"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["tags"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["marked_for_delete"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["parent_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	fields["connectivities"] = bindings.NewListType(bindings.NewReferenceType(ExternalSddcConnectivityBindingType), reflect.TypeOf([]ExternalSddcConnectivity{}))
	fieldNameMap["connectivities"] = "Connectivities"
	fields["destination"] = bindings.NewStringType()
	fieldNameMap["destination"] = "Destination"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.external_sddc_route", fields, reflect.TypeOf(ExternalSddcRoute{}), fieldNameMap, validators)
}

func ExternalSddcRouteCsvRecordBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["connectivity_details"] = bindings.NewStringType()
	fieldNameMap["connectivity_details"] = "ConnectivityDetails"
	fields["destination"] = bindings.NewStringType()
	fieldNameMap["destination"] = "Destination"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.external_sddc_route_csv_record", fields, reflect.TypeOf(ExternalSddcRouteCsvRecord{}), fieldNameMap, validators)
}

func ExternalSddcRoutesListResultBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fields["result_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["result_count"] = "ResultCount"
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["sort_ascending"] = "SortAscending"
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sort_by"] = "SortBy"
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

func InterfaceConfigBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["urpf_mode"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["urpf_mode"] = "UrpfMode"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.interface_config", fields, reflect.TypeOf(InterfaceConfig{}), fieldNameMap, validators)
}

func IpAttachmentPairBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["attachment_id"] = bindings.NewStringType()
	fieldNameMap["attachment_id"] = "AttachmentId"
	fields["ip"] = bindings.NewStringType()
	fieldNameMap["ip"] = "Ip"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.ip_attachment_pair", fields, reflect.TypeOf(IpAttachmentPair{}), fieldNameMap, validators)
}

func LinkedSubnetInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["availability_zone"] = bindings.NewStringType()
	fieldNameMap["availability_zone"] = "AvailabilityZone"
	fields["cidr"] = bindings.NewStringType()
	fieldNameMap["cidr"] = "Cidr"
	fields["id"] = bindings.NewStringType()
	fieldNameMap["id"] = "Id"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.linked_subnet_info", fields, reflect.TypeOf(LinkedSubnetInfo{}), fieldNameMap, validators)
}

func LinkedVpcInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_revision"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_create_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_create_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_last_modified_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["_protection"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_system_owned"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["tags"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["marked_for_delete"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["parent_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	fields["active_eni"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["active_eni"] = "ActiveEni"
	fields["arn_role"] = bindings.NewStringType()
	fieldNameMap["arn_role"] = "ArnRole"
	fields["external_id"] = bindings.NewStringType()
	fieldNameMap["external_id"] = "ExternalId"
	fields["linked_account"] = bindings.NewStringType()
	fieldNameMap["linked_account"] = "LinkedAccount"
	fields["linked_vpc_addresses"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["linked_vpc_addresses"] = "LinkedVpcAddresses"
	fields["linked_vpc_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["linked_vpc_id"] = "LinkedVpcId"
	fields["linked_vpc_nat_ips"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["linked_vpc_nat_ips"] = "LinkedVpcNatIps"
	fields["linked_vpc_subnets"] = bindings.NewListType(bindings.NewReferenceType(LinkedSubnetInfoBindingType), reflect.TypeOf([]LinkedSubnetInfo{}))
	fieldNameMap["linked_vpc_subnets"] = "LinkedVpcSubnets"
	fields["route_table_ids"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["route_table_ids"] = "RouteTableIds"
	fields["service_arn_role"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["service_arn_role"] = "ServiceArnRole"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.linked_vpc_info", fields, reflect.TypeOf(LinkedVpcInfo{}), fieldNameMap, validators)
}

func LinkedVpcsListResultBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fields["result_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["result_count"] = "ResultCount"
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["sort_ascending"] = "SortAscending"
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sort_by"] = "SortBy"
	fields["results"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(LinkedVpcInfoBindingType), reflect.TypeOf([]LinkedVpcInfo{})))
	fieldNameMap["results"] = "Results"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.linked_vpcs_list_result", fields, reflect.TypeOf(LinkedVpcsListResult{}), fieldNameMap, validators)
}

func ListResultBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fields["result_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["result_count"] = "ResultCount"
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["sort_ascending"] = "SortAscending"
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.list_result", fields, reflect.TypeOf(ListResult{}), fieldNameMap, validators)
}

func ManagedResourceBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_revision"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_create_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_create_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_last_modified_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["_protection"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_system_owned"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["tags"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.managed_resource", fields, reflect.TypeOf(ManagedResource{}), fieldNameMap, validators)
}

func MgmtServiceEntryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["path"] = "Path"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.mgmt_service_entry", fields, reflect.TypeOf(MgmtServiceEntry{}), fieldNameMap, validators)
}

func MgmtVmInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_revision"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_create_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_create_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_last_modified_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["_protection"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_system_owned"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["tags"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["marked_for_delete"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["parent_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	fields["group_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["group_path"] = "GroupPath"
	fields["ip_attachment_pairs"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(IpAttachmentPairBindingType), reflect.TypeOf([]IpAttachmentPair{})))
	fieldNameMap["ip_attachment_pairs"] = "IpAttachmentPairs"
	fields["ips"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["ips"] = "Ips"
	fields["services"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(MgmtServiceEntryBindingType), reflect.TypeOf([]MgmtServiceEntry{})))
	fieldNameMap["services"] = "Services"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.mgmt_vm_info", fields, reflect.TypeOf(MgmtVmInfo{}), fieldNameMap, validators)
}

func MgmtVmsListResultBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fields["result_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["result_count"] = "ResultCount"
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["sort_ascending"] = "SortAscending"
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sort_by"] = "SortBy"
	fields["results"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(MgmtVmInfoBindingType), reflect.TypeOf([]MgmtVmInfo{})))
	fieldNameMap["results"] = "Results"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.mgmt_vms_list_result", fields, reflect.TypeOf(MgmtVmsListResult{}), fieldNameMap, validators)
}

func ModelInterfaceBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = bindings.NewStringType()
	fieldNameMap["id"] = "Id"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.model_interface", fields, reflect.TypeOf(ModelInterface{}), fieldNameMap, validators)
}

func OtherRegionConnectionInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["region"] = bindings.NewStringType()
	fieldNameMap["region"] = "Region"
	fields["tgw_id"] = bindings.NewStringType()
	fieldNameMap["tgw_id"] = "TgwId"
	fields["tgw_peering_id"] = bindings.NewStringType()
	fieldNameMap["tgw_peering_id"] = "TgwPeeringId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.other_region_connection_info", fields, reflect.TypeOf(OtherRegionConnectionInfo{}), fieldNameMap, validators)
}

func PrefixListInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["path"] = bindings.NewStringType()
	fieldNameMap["path"] = "Path"
	fields["url"] = bindings.NewStringType()
	fieldNameMap["url"] = "Url"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.prefix_list_info", fields, reflect.TypeOf(PrefixListInfo{}), fieldNameMap, validators)
}

func ProviderGatewayKeyValuePairsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["key"] = bindings.NewStringType()
	fieldNameMap["key"] = "Key"
	fields["value"] = bindings.NewListType(bindings.NewReferenceType(ProviderObjectBindingType), reflect.TypeOf([]ProviderObject{}))
	fieldNameMap["value"] = "Value"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.provider_gateway_key_value_pairs", fields, reflect.TypeOf(ProviderGatewayKeyValuePairs{}), fieldNameMap, validators)
}

func ProviderObjectBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["type"] = "Type_"
	fields["url"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["url"] = "Url"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.provider_object", fields, reflect.TypeOf(ProviderObject{}), fieldNameMap, validators)
}

func PublicIpBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_revision"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_create_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_create_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_last_modified_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["_protection"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_system_owned"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["tags"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["marked_for_delete"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["parent_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	fields["ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["ip"] = "Ip"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.public_ip", fields, reflect.TypeOf(PublicIp{}), fieldNameMap, validators)
}

func PublicIpsListResultBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fields["result_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["result_count"] = "ResultCount"
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["sort_ascending"] = "SortAscending"
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sort_by"] = "SortBy"
	fields["results"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(PublicIpBindingType), reflect.TypeOf([]PublicIp{})))
	fieldNameMap["results"] = "Results"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.public_ips_list_result", fields, reflect.TypeOf(PublicIpsListResult{}), fieldNameMap, validators)
}

func RelatedApiErrorBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["details"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["details"] = "Details"
	fields["error_code"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["error_code"] = "ErrorCode"
	fields["error_data"] = bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
	fieldNameMap["error_data"] = "ErrorData"
	fields["error_message"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["error_message"] = "ErrorMessage"
	fields["module_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["module_name"] = "ModuleName"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.related_api_error", fields, reflect.TypeOf(RelatedApiError{}), fieldNameMap, validators)
}

func ResourceBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
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
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_revision"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.revisioned_resource", fields, reflect.TypeOf(RevisionedResource{}), fieldNameMap, validators)
}

func SddcAdvertiseRouteConfigBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["aws_prefix_list"] = bindings.NewListType(bindings.NewReferenceType(AwsPrefixInfoBindingType), reflect.TypeOf([]AwsPrefixInfo{}))
	fieldNameMap["aws_prefix_list"] = "AwsPrefixList"
	fields["other_region_connections"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(OtherRegionConnectionInfoBindingType), reflect.TypeOf([]OtherRegionConnectionInfo{})))
	fieldNameMap["other_region_connections"] = "OtherRegionConnections"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.sddc_advertise_route_config", fields, reflect.TypeOf(SddcAdvertiseRouteConfig{}), fieldNameMap, validators)
}

func SddcUserConfigurationBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["all_uplink_interface_label"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["all_uplink_interface_label"] = "AllUplinkInterfaceLabel"
	fields["all_vpn_interface_label"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["all_vpn_interface_label"] = "AllVpnInterfaceLabel"
	fields["cgw_snat_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cgw_snat_ip"] = "CgwSnatIp"
	fields["compute_domain"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["compute_domain"] = "ComputeDomain"
	fields["compute_gateway"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["compute_gateway"] = "ComputeGateway"
	fields["domains"] = bindings.NewListType(bindings.NewReferenceType(ProviderObjectBindingType), reflect.TypeOf([]ProviderObject{}))
	fieldNameMap["domains"] = "Domains"
	fields["dx_interface_label"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["dx_interface_label"] = "DxInterfaceLabel"
	fields["enforcement_points"] = bindings.NewListType(bindings.NewReferenceType(ProviderObjectBindingType), reflect.TypeOf([]ProviderObject{}))
	fieldNameMap["enforcement_points"] = "EnforcementPoints"
	fields["gateways"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ProviderGatewayKeyValuePairsBindingType), reflect.TypeOf([]ProviderGatewayKeyValuePairs{})))
	fieldNameMap["gateways"] = "Gateways"
	fields["infra_subnets"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["infra_subnets"] = "InfraSubnets"
	fields["interfaces"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ModelInterfaceBindingType), reflect.TypeOf([]ModelInterface{})))
	fieldNameMap["interfaces"] = "Interfaces"
	fields["labels"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ProviderObjectBindingType), reflect.TypeOf([]ProviderObject{})))
	fieldNameMap["labels"] = "Labels"
	fields["linked_vpc_interface_label"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["linked_vpc_interface_label"] = "LinkedVpcInterfaceLabel"
	fields["management_domain"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["management_domain"] = "ManagementDomain"
	fields["management_gateway"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["management_gateway"] = "ManagementGateway"
	fields["management_gateway_default_dns_zone"] = bindings.NewReferenceType(ProviderObjectBindingType)
	fieldNameMap["management_gateway_default_dns_zone"] = "ManagementGatewayDefaultDnsZone"
	fields["management_gateway_label"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["management_gateway_label"] = "ManagementGatewayLabel"
	fields["mgmt_subnet"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["mgmt_subnet"] = "MgmtSubnet"
	fields["mgmt_subnets"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["mgmt_subnets"] = "MgmtSubnets"
	fields["mgw_snat_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["mgw_snat_ip"] = "MgwSnatIp"
	fields["provider_gateways"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ProviderGatewayKeyValuePairsBindingType), reflect.TypeOf([]ProviderGatewayKeyValuePairs{})))
	fieldNameMap["provider_gateways"] = "ProviderGateways"
	fields["provider_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["provider_name"] = "ProviderName"
	fields["provider_prefix_lists"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(PrefixListInfoBindingType), reflect.TypeOf([]PrefixListInfo{})))
	fieldNameMap["provider_prefix_lists"] = "ProviderPrefixLists"
	fields["public_interface_label"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["public_interface_label"] = "PublicInterfaceLabel"
	fields["sddc_infra_subnet"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["sddc_infra_subnet"] = "SddcInfraSubnet"
	fields["vpn_dx_ips"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["vpn_dx_ips"] = "VpnDxIps"
	fields["vpn_endpoints"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(VpnEndpointBindingType), reflect.TypeOf([]VpnEndpoint{})))
	fieldNameMap["vpn_endpoints"] = "VpnEndpoints"
	fields["vpn_internet_ips"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["vpn_internet_ips"] = "VpnInternetIps"
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

func TrafficGroupBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_revision"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_create_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_create_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_last_modified_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["_protection"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_system_owned"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["tags"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["marked_for_delete"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["parent_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	fields["association_maps"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(TrafficGroupAssociationMapBindingType), reflect.TypeOf([]TrafficGroupAssociationMap{})))
	fieldNameMap["association_maps"] = "AssociationMaps"
	fields["resource_id"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["resource_id"] = "ResourceId"
	fields["state"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["state"] = "State"
	fields["state_message"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["state_message"] = "StateMessage"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.traffic_group", fields, reflect.TypeOf(TrafficGroup{}), fieldNameMap, validators)
}

func TrafficGroupAssociationMapBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_revision"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_create_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_create_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_last_modified_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["_protection"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_system_owned"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["tags"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["marked_for_delete"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["parent_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	fields["prefix_lists"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["prefix_lists"] = "PrefixLists"
	fields["scope"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["scope"] = "Scope"
	fields["state"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["state"] = "State"
	fields["state_message"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["state_message"] = "StateMessage"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.traffic_group_association_map", fields, reflect.TypeOf(TrafficGroupAssociationMap{}), fieldNameMap, validators)
}

func TrafficGroupAssociationMapsListResultBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fields["result_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["result_count"] = "ResultCount"
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["sort_ascending"] = "SortAscending"
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sort_by"] = "SortBy"
	fields["results"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(TrafficGroupAssociationMapBindingType), reflect.TypeOf([]TrafficGroupAssociationMap{})))
	fieldNameMap["results"] = "Results"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.traffic_group_association_maps_list_result", fields, reflect.TypeOf(TrafficGroupAssociationMapsListResult{}), fieldNameMap, validators)
}

func TrafficGroupsListResultBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fields["result_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["result_count"] = "ResultCount"
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["sort_ascending"] = "SortAscending"
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sort_by"] = "SortBy"
	fields["results"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(TrafficGroupBindingType), reflect.TypeOf([]TrafficGroup{})))
	fieldNameMap["results"] = "Results"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.traffic_groups_list_result", fields, reflect.TypeOf(TrafficGroupsListResult{}), fieldNameMap, validators)
}

func UpgradeVersionInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["from_sddc_version"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["from_sddc_version"] = "FromSddcVersion"
	fields["to_sddc_version"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["to_sddc_version"] = "ToSddcVersion"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.upgrade_version_info", fields, reflect.TypeOf(UpgradeVersionInfo{}), fieldNameMap, validators)
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
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fields["result_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["result_count"] = "ResultCount"
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["sort_ascending"] = "SortAscending"
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sort_by"] = "SortBy"
	fields["results"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(VirtualInterfaceBindingType), reflect.TypeOf([]VirtualInterface{})))
	fieldNameMap["results"] = "Results"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.vifs_list_result", fields, reflect.TypeOf(VifsListResult{}), fieldNameMap, validators)
}

func VirtualInterfaceBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["bgp_status"] = bindings.NewStringType()
	fieldNameMap["bgp_status"] = "BgpStatus"
	fields["direct_connect_id"] = bindings.NewStringType()
	fieldNameMap["direct_connect_id"] = "DirectConnectId"
	fields["id"] = bindings.NewStringType()
	fieldNameMap["id"] = "Id"
	fields["local_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["local_ip"] = "LocalIp"
	fields["mtu"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["mtu"] = "Mtu"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["remote_asn"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["remote_asn"] = "RemoteAsn"
	fields["remote_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["remote_ip"] = "RemoteIp"
	fields["state"] = bindings.NewStringType()
	fieldNameMap["state"] = "State"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.virtual_interface", fields, reflect.TypeOf(VirtualInterface{}), fieldNameMap, validators)
}

func VmcAppBaseInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_revision"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_create_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_create_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_last_modified_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["_protection"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_system_owned"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["tags"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["marked_for_delete"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["parent_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.vmc_app_base_info", fields, reflect.TypeOf(VmcAppBaseInfo{}), fieldNameMap, validators)
}

func VmcConsolidatedRealizedStatusBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_revision"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_create_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_create_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_last_modified_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["_protection"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_system_owned"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["tags"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["marked_for_delete"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["parent_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
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
	fields["consolidated_status"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["consolidated_status"] = "ConsolidatedStatus"
	fields["status_message"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["status_message"] = "StatusMessage"
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
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_revision"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_revision"] = "Revision"
	fields["_create_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_create_time"] = "CreateTime"
	fields["_create_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_create_user"] = "CreateUser"
	fields["_last_modified_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_last_modified_time"] = "LastModifiedTime"
	fields["_last_modified_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_last_modified_user"] = "LastModifiedUser"
	fields["_protection"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_protection"] = "Protection"
	fields["_system_owned"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["_system_owned"] = "SystemOwned"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["tags"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	fields["marked_for_delete"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["marked_for_delete"] = "MarkedForDelete"
	fields["parent_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["parent_path"] = "ParentPath"
	fields["path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["relative_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["relative_path"] = "RelativePath"
	fields["internal_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["internal_name"] = "InternalName"
	fields["message"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["message"] = "Message"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["state"] = bindings.NewStringType()
	fieldNameMap["state"] = "State"
	fields["unlicensed_message"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["unlicensed_message"] = "UnlicensedMessage"
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
	fields["intent_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["intent_path"] = "IntentPath"
	fields["realized_entities"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(VmcRealizedEntityBindingType), reflect.TypeOf([]VmcRealizedEntity{})))
	fieldNameMap["realized_entities"] = "RealizedEntities"
	fields["realized_entities_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["realized_entities_id"] = "RealizedEntitiesId"
	fields["realized_state"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["realized_state"] = "RealizedState"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.vmc_realized_entities", fields, reflect.TypeOf(VmcRealizedEntities{}), fieldNameMap, validators)
}

func VmcRealizedEntityBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["realization_api"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["realization_api"] = "RealizationApi"
	fields["realization_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["realization_id"] = "RealizationId"
	fields["realization_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["realization_path"] = "RealizationPath"
	fields["realized_entity_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["realized_entity_id"] = "RealizedEntityId"
	fields["realized_entity_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["realized_entity_type"] = "RealizedEntityType"
	fields["realized_state"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["realized_state"] = "RealizedState"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.vmc_realized_entity", fields, reflect.TypeOf(VmcRealizedEntity{}), fieldNameMap, validators)
}

func VpnEndpointBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["interface_label"] = bindings.NewStringType()
	fieldNameMap["interface_label"] = "InterfaceLabel"
	fields["ip"] = bindings.NewStringType()
	fieldNameMap["ip"] = "Ip"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["type"] = bindings.NewStringType()
	fieldNameMap["type"] = "Type_"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.vpn_endpoint", fields, reflect.TypeOf(VpnEndpoint{}), fieldNameMap, validators)
}
