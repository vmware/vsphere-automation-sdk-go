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
)


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


// Direct Connect BGP related information
type DirectConnectBgpInfo struct {
    // The ASN paired with the VGW attached to the VPC. AWS allowed private BGP ASN range - [64512, 65534] and [4200000000, 4294967294]. If omitted in the payload, BGP ASN will not be modified.
	LocalAsNum *string
    // Possible values are: 
    //
    // * DirectConnectBgpInfo#DirectConnectBgpInfo_ROUTE_PREFERENCE_DX_PREFERED_OVER_VPN
    // * DirectConnectBgpInfo#DirectConnectBgpInfo_ROUTE_PREFERENCE_VPN_PREFERED_OVER_DX
    //
    //  Direct connect route preference over VPN routes. If omitted in the payload, route preference will not be modified.
	RoutePreference *string
    // Maximum transmission unit allowed by the VIF format: int32
	Mtu *int64
}
const DirectConnectBgpInfo_ROUTE_PREFERENCE_DX_PREFERED_OVER_VPN = "DX_PREFERED_OVER_VPN"
const DirectConnectBgpInfo_ROUTE_PREFERENCE_VPN_PREFERED_OVER_DX = "VPN_PREFERED_OVER_DX"

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


// Base class for resources that are discovered and automatically updated
type DiscoveredResource struct {
    // Link to this resource
	Self *SelfResourceLink
    // The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
    // Schema for this resource
	Schema *string
    // Timestamp of last modification format: int64
	LastSyncTime *int64
    // Defaults to ID if not set
	DisplayName *string
    // Description of this resource
	Description *string
    // The type of this resource.
	ResourceType *string
    // Opaque identifiers meaningful to the API user
	Tags []Tag
}

func (s DiscoveredResource) GetType__() bindings.BindingType {
	return DiscoveredResourceBindingType()
}

func (s DiscoveredResource) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for DiscoveredResource._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Base class for resources that are embedded in other resources
type EmbeddedResource struct {
    // Link to this resource
	Self *SelfResourceLink
    // The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
    // Schema for this resource
	Schema *string
    // The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected. format: int32
	Revision *int64
    // Owner of this resource
	Owner *OwnerResourceLink
    // Defaults to ID if not set
	DisplayName *string
    // Identifier of the resource
	Id *string
    // The type of this resource.
	ResourceType *string
    // Description of this resource
	Description *string
}

func (s EmbeddedResource) GetType__() bindings.BindingType {
	return EmbeddedResourceBindingType()
}

func (s EmbeddedResource) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for EmbeddedResource._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Host elastic network interface (ENI)
type HostEni struct {
    // Interface mac
	InterfaceMac *string
    // List of associated public IPs. format: ipv4
	AssociatedPublicIps []string
    // Primary IP format: ipv4
	PrimaryIp *string
    // Description
	Description *string
    // Subnet ID
	SubnetId *string
    // Virtual distributed router (VDR) type
	VdrType *string
    // Interface id
	InterfaceId *string
}

func (s HostEni) GetType__() bindings.BindingType {
	return HostEniBindingType()
}

func (s HostEni) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for HostEni._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Host status
type HostStatus struct {
    // List of VDRs
	Vdr []HostVdr
    // Host UUID
	HostId *string
    // Possible values are: 
    //
    // * HostStatus#HostStatus_VMCD_STATUS_GREEN
    // * HostStatus#HostStatus_VMCD_STATUS_RED
    // * HostStatus#HostStatus_VMCD_STATUS_UNKNOWN
    //
    //  Status of vmcd, a service running on host
	VmcdStatus *string
    // List of ENIs
	Eni []HostEni
    // Host ip format: ipv4
	HostIp *string
    // issues with the host
	Issues []string
}
const HostStatus_VMCD_STATUS_GREEN = "GREEN"
const HostStatus_VMCD_STATUS_RED = "RED"
const HostStatus_VMCD_STATUS_UNKNOWN = "UNKNOWN"

func (s HostStatus) GetType__() bindings.BindingType {
	return HostStatusBindingType()
}

func (s HostStatus) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for HostStatus._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Host status list result
type HostStatusListResult struct {
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
	Results []HostStatus
}

func (s HostStatusListResult) GetType__() bindings.BindingType {
	return HostStatusListResultBindingType()
}

func (s HostStatusListResult) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for HostStatusListResult._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Host virtual distributed router (VDR)
type HostVdr struct {
    // List of VDR routes
	Routes []VdrRoute
    // List of VDR lifs
	Lifs []VdrLif
    // VDR type
	Type_ *string
}

func (s HostVdr) GetType__() bindings.BindingType {
	return HostVdrBindingType()
}

func (s HostVdr) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for HostVdr._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Statistics for a network interface
type InterfaceStatistics struct {
    // Link to this resource
	Self *SelfResourceLink
    // The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
    // Schema for this resource
	Schema *string
    // Count of packets received on this port format: int64
	RxPackets *int64
    // Count of receive errors occurring on this port format: int64
	RxErrors *int64
    // Count of bytes received on this port format: int64
	RxBytes *int64
    // Count of transmit errors occurring on this port format: int64
	TxErrors *int64
    // Count of bytes transmitted on this port format: int64
	TxBytes *int64
    // Count of packets transmitted on this port format: int64
	TxPackets *int64
}

func (s InterfaceStatistics) GetType__() bindings.BindingType {
	return InterfaceStatisticsBindingType()
}

func (s InterfaceStatistics) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for InterfaceStatistics._GetDataValue method - %s",
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
    // Local IPs of a management VM format: address-or-block-or-range
	Ips []string
    // Management VM name
	DisplayName *string
    // Details services path and display name.
	Services []MgmtServiceEntry
    // For each management VM, a dedicated policy group will be created. This property will reflect its group path.
	GroupPath *string
    // IP address and attachment id pairs for tagging managment VM
	IpAttachmentPairs []IpAttachmentPair
    // Management VM identifier
	Id *string
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


// Network status entry
type NetworkStatusEntry struct {
    // Indicate whether issues is a non-empty array
	IssuesFound *bool
    // IP address programmed with the entry format: ipv4
	IpAddress *string
    // IPs of hosts that store this entry format: ipv4
	HostIps []string
    // Known issues detected with the entry
	Issues []string
}

func (s NetworkStatusEntry) GetType__() bindings.BindingType {
	return NetworkStatusEntryBindingType()
}

func (s NetworkStatusEntry) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for NetworkStatusEntry._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Key used to group network status inquiry results.
type NetworkStatusKey struct {
    // Possible values are: 
    //
    // * NetworkStatusKey#NetworkStatusKey_NETWORK_TYPE_DNS
    // * NetworkStatusKey#NetworkStatusKey_NETWORK_TYPE_VPN
    // * NetworkStatusKey#NetworkStatusKey_NETWORK_TYPE_MANAGEMENT
    // * NetworkStatusKey#NetworkStatusKey_NETWORK_TYPE_NAT
    // * NetworkStatusKey#NetworkStatusKey_NETWORK_TYPE_LOGICAL
    // * NetworkStatusKey#NetworkStatusKey_NETWORK_TYPE_INFRA
    //
    //  Network type in the network status pair.
	NetworkType string
    // Possible values are: 
    //
    // * NetworkStatusKey#NetworkStatusKey_CONTEXT_INVALID_NETWORK
    // * NetworkStatusKey#NetworkStatusKey_CONTEXT_MANAGEMENT
    // * NetworkStatusKey#NetworkStatusKey_CONTEXT_PUBLIC
    // * NetworkStatusKey#NetworkStatusKey_CONTEXT_CONNECTED_VPC
    // * NetworkStatusKey#NetworkStatusKey_CONTEXT_DIRECT_CONNECT
    //
    //  The context that the entry is used in
	Context string
}
const NetworkStatusKey_NETWORK_TYPE_DNS = "DNS"
const NetworkStatusKey_NETWORK_TYPE_VPN = "VPN"
const NetworkStatusKey_NETWORK_TYPE_MANAGEMENT = "MANAGEMENT"
const NetworkStatusKey_NETWORK_TYPE_NAT = "NAT"
const NetworkStatusKey_NETWORK_TYPE_LOGICAL = "LOGICAL"
const NetworkStatusKey_NETWORK_TYPE_INFRA = "INFRA"
const NetworkStatusKey_CONTEXT_INVALID_NETWORK = "INVALID_NETWORK"
const NetworkStatusKey_CONTEXT_MANAGEMENT = "MANAGEMENT"
const NetworkStatusKey_CONTEXT_PUBLIC = "PUBLIC"
const NetworkStatusKey_CONTEXT_CONNECTED_VPC = "CONNECTED_VPC"
const NetworkStatusKey_CONTEXT_DIRECT_CONNECT = "DIRECT_CONNECT"

func (s NetworkStatusKey) GetType__() bindings.BindingType {
	return NetworkStatusKeyBindingType()
}

func (s NetworkStatusKey) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for NetworkStatusKey._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// List of network status
type NetworkStatusKeyValuePair struct {
    // Network status value
	Values []NetworkStatusEntry
    // Network status key
	Key *NetworkStatusKey
}

func (s NetworkStatusKeyValuePair) GetType__() bindings.BindingType {
	return NetworkStatusKeyValuePairBindingType()
}

func (s NetworkStatusKeyValuePair) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for NetworkStatusKeyValuePair._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// List of network status
type NetworkStatusListResult struct {
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
    // List of network status key value pairs.
	Results []NetworkStatusKeyValuePair
    // List of overall issues encountered during the inquiry
	Issues []string
}

func (s NetworkStatusListResult) GetType__() bindings.BindingType {
	return NetworkStatusListResultBindingType()
}

func (s NetworkStatusListResult) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for NetworkStatusListResult._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// The server will populate this field when returing the resource. Ignored on PUT and POST.
type OwnerResourceLink struct {
    // Optional action
	Action *string
    // Link to resource
	Href *string
    // Custom relation type (follows RFC 5988 where appropriate definitions exist)
	Rel *string
}

func (s OwnerResourceLink) GetType__() bindings.BindingType {
	return OwnerResourceLinkBindingType()
}

func (s OwnerResourceLink) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for OwnerResourceLink._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Service IP prefixes information
type PrefixInfo struct {
    // Service IP prefixes format: ipv4-cidr-block
	Prefixes []string
    // Display name
	DisplayName *string
}

func (s PrefixInfo) GetType__() bindings.BindingType {
	return PrefixInfoBindingType()
}

func (s PrefixInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for PrefixInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Service Prefix list query result
type PrefixesListResult struct {
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
    // Service Prefixes list
	Results []PrefixInfo
}

func (s PrefixesListResult) GetType__() bindings.BindingType {
	return PrefixesListResultBindingType()
}

func (s PrefixesListResult) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for PrefixesListResult._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type PublicIp struct {
    // IPv4 address format: ipv4
	Ip *string
	DisplayName *string
    // Public IP identifier
	Id *string
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
    // Public IPs for VPN tunnel over internet format: ipv4
	VpnInternetIps []string
    // MGW SNAT ip address format: ipv4
	MgwSnatIp *string
    // Compute gateway name
	ComputeGateway string
    // All VPN interfaces label name
	AllVpnInterfaceLabel string
    // Local IPs for VPN tunnel over Direct Connect format: ipv4
	VpnDxIps []string
    // All uplink interfaces label name
	AllUplinkInterfaceLabel string
    // DirectConnect interface label name
	DxInterfaceLabel string
    // SDDC Infra CIDRs format: ipv4-cidr-block
	SddcInfraSubnet []string
    // Management gateway name
	ManagementGateway string
    // Linked VPC interface label name
	LinkedVpcInterfaceLabel string
    // Public interface label name
	PublicInterfaceLabel string
    // CGW SNAT ip address format: ipv4
	CgwSnatIp *string
    // Provider Name
	ProviderName string
    // Management VMs CIDRs format: ipv4-cidr-block
	MgmtSubnet []string
    // Management gateway label name
	ManagementGatewayLabel string
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


// Arbitrary key-value pairs that may be attached to an entity
type Tag struct {
    // Tag searches may optionally be restricted by scope
	Scope *string
    // Identifier meaningful to user
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


// Task properties
type TaskProperties struct {
    // Link to this resource
	Self *SelfResourceLink
    // The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink
    // Schema for this resource
	Schema *string
    // Possible values are: 
    //
    // * TaskProperties#TaskProperties_STATUS_RUNNING
    // * TaskProperties#TaskProperties_STATUS_ERROR
    // * TaskProperties#TaskProperties_STATUS_SUCCESS
    // * TaskProperties#TaskProperties_STATUS_CANCELING
    // * TaskProperties#TaskProperties_STATUS_CANCELED
    // * TaskProperties#TaskProperties_STATUS_KILLED
    //
    //  Current status of the task
	Status *string
    // True if response for asynchronous request is available
	AsyncResponseAvailable *bool
    // Description of the task
	Description *string
    // The start time of the task in epoch milliseconds format: int64
	StartTime *int64
    // True if this task can be canceled
	Cancelable *bool
    // HTTP request method
	RequestMethod *string
    // Name of the user who created this task
	User *string
    // Task progress if known, from 0 to 100 format: int64
	Progress *int64
    // A message describing the disposition of the task
	Message *string
    // URI of the method invocation that spawned this task
	RequestUri *string
    // Identifier for this task
	Id *string
    // The end time of the task in epoch milliseconds format: int64
	EndTime *int64
}
const TaskProperties_STATUS_RUNNING = "running"
const TaskProperties_STATUS_ERROR = "error"
const TaskProperties_STATUS_SUCCESS = "success"
const TaskProperties_STATUS_CANCELING = "canceling"
const TaskProperties_STATUS_CANCELED = "canceled"
const TaskProperties_STATUS_KILLED = "killed"

func (s TaskProperties) GetType__() bindings.BindingType {
	return TaskPropertiesBindingType()
}

func (s TaskProperties) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for TaskProperties._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// A component that processed the packet injected by traceflow
type TraceflowAction struct {
    // Reason to drop or reject packet if it is not forwarded
	Reason *string
    // Name of a component instance
	ComponentName *string
    // Action taken by the component to process the packet
	ResourceType *string
    // Possible values are: 
    //
    // * TraceflowAction#TraceflowAction_COMPONENT_TYPE_VMC
    //
    //  Type of component
	ComponentType *string
    // Possible values are: 
    //
    // * TraceflowAction#TraceflowAction_COMPONENT_SUB_TYPE_EDGE_UPLINK
    // * TraceflowAction#TraceflowAction_COMPONENT_SUB_TYPE_VDR
    // * TraceflowAction#TraceflowAction_COMPONENT_SUB_TYPE_ENI
    // * TraceflowAction#TraceflowAction_COMPONENT_SUB_TYPE_AWS_GATEWAY
    //
    //  Subtype of component
	ComponentSubType *string
}
const TraceflowAction_COMPONENT_TYPE_VMC = "VMC"
const TraceflowAction_COMPONENT_SUB_TYPE_EDGE_UPLINK = "EDGE_UPLINK"
const TraceflowAction_COMPONENT_SUB_TYPE_VDR = "VDR"
const TraceflowAction_COMPONENT_SUB_TYPE_ENI = "ENI"
const TraceflowAction_COMPONENT_SUB_TYPE_AWS_GATEWAY = "AWS_GATEWAY"

func (s TraceflowAction) GetType__() bindings.BindingType {
	return TraceflowActionBindingType()
}

func (s TraceflowAction) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for TraceflowAction._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// A list containing all traceflow actions that have been taken to process the packet
type TraceflowActionListResults struct {
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
    // Result containing all traceflow actions that have processed the packet
	Results []TraceflowAction
}

func (s TraceflowActionListResults) GetType__() bindings.BindingType {
	return TraceflowActionListResultsBindingType()
}

func (s TraceflowActionListResults) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for TraceflowActionListResults._GetDataValue method - %s",
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


// Virtual distributed router (VDR) LIF
type VdrLif struct {
    // VDR LIF IP format: ipv4
	Ip *string
    // VDR LIF subnet mask format: ipv4
	Netmask *string
    // VDR LIF ID
	Id *string
    // VDR LIF VLAN ID format: int64
	VlanId *int64
}

func (s VdrLif) GetType__() bindings.BindingType {
	return VdrLifBindingType()
}

func (s VdrLif) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for VdrLif._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Virtual Distributed Router (VDR) route entry
type VdrRoute struct {
    // Destination IP CIDR Block format: ipv4-cidr-block
	Destination *string
    // Outgoing gateway
	Gateway *string
    // Outgoing Lif ID
	LifId *string
}

func (s VdrRoute) GetType__() bindings.BindingType {
	return VdrRouteBindingType()
}

func (s VdrRoute) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for VdrRoute._GetDataValue method - %s",
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
    // * VmcConsolidatedStatus#VmcConsolidatedStatus_CONSOLIDATED_STATUS_ERROR
    // * VmcConsolidatedStatus#VmcConsolidatedStatus_CONSOLIDATED_STATUS_UNAVAILABLE
    //
    //  Possible values could be IN_PROGRESS, SUCCESS, ERROR, UNAVAILABLE. IN_PROGRESS - The object realization is in progress. ERROR - The object realization fails or is caught in an error. SUCCESS - The realization succeeds. UNAVAILABLE - The object realization status is unavailable.
	ConsolidatedStatus *string
}
const VmcConsolidatedStatus_CONSOLIDATED_STATUS_IN_PROGRESS = "IN_PROGRESS"
const VmcConsolidatedStatus_CONSOLIDATED_STATUS_SUCCESS = "SUCCESS"
const VmcConsolidatedStatus_CONSOLIDATED_STATUS_ERROR = "ERROR"
const VmcConsolidatedStatus_CONSOLIDATED_STATUS_UNAVAILABLE = "UNAVAILABLE"

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
	fields["enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["enabled"] = "Enabled"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.connected_service_status", fields, reflect.TypeOf(ConnectedServiceStatus{}), fieldNameMap, validators)
}

func DirectConnectBgpInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["local_as_num"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["local_as_num"] = "LocalAsNum"
	fields["route_preference"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["route_preference"] = "RoutePreference"
	fields["mtu"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["mtu"] = "Mtu"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.direct_connect_bgp_info", fields, reflect.TypeOf(DirectConnectBgpInfo{}), fieldNameMap, validators)
}

func DiscoveredResourceBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["_last_sync_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["_last_sync_time"] = "LastSyncTime"
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["tags"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.discovered_resource", fields, reflect.TypeOf(DiscoveredResource{}), fieldNameMap, validators)
}

func EmbeddedResourceBindingType() bindings.BindingType {
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
	fields["_owner"] = bindings.NewOptionalType(bindings.NewReferenceType(OwnerResourceLinkBindingType))
	fieldNameMap["_owner"] = "Owner"
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.embedded_resource", fields, reflect.TypeOf(EmbeddedResource{}), fieldNameMap, validators)
}

func HostEniBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["interface_mac"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["interface_mac"] = "InterfaceMac"
	fields["associated_public_ips"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["associated_public_ips"] = "AssociatedPublicIps"
	fields["primary_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["primary_ip"] = "PrimaryIp"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["subnet_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["subnet_id"] = "SubnetId"
	fields["vdr_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vdr_type"] = "VdrType"
	fields["interface_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["interface_id"] = "InterfaceId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.host_eni", fields, reflect.TypeOf(HostEni{}), fieldNameMap, validators)
}

func HostStatusBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vdr"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(HostVdrBindingType), reflect.TypeOf([]HostVdr{})))
	fieldNameMap["vdr"] = "Vdr"
	fields["host_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["host_id"] = "HostId"
	fields["vmcd_status"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vmcd_status"] = "VmcdStatus"
	fields["eni"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(HostEniBindingType), reflect.TypeOf([]HostEni{})))
	fieldNameMap["eni"] = "Eni"
	fields["host_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["host_ip"] = "HostIp"
	fields["issues"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["issues"] = "Issues"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.host_status", fields, reflect.TypeOf(HostStatus{}), fieldNameMap, validators)
}

func HostStatusListResultBindingType() bindings.BindingType {
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
	fields["results"] = bindings.NewListType(bindings.NewReferenceType(HostStatusBindingType), reflect.TypeOf([]HostStatus{}))
	fieldNameMap["results"] = "Results"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.host_status_list_result", fields, reflect.TypeOf(HostStatusListResult{}), fieldNameMap, validators)
}

func HostVdrBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["routes"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(VdrRouteBindingType), reflect.TypeOf([]VdrRoute{})))
	fieldNameMap["routes"] = "Routes"
	fields["lifs"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(VdrLifBindingType), reflect.TypeOf([]VdrLif{})))
	fieldNameMap["lifs"] = "Lifs"
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["type"] = "Type_"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.host_vdr", fields, reflect.TypeOf(HostVdr{}), fieldNameMap, validators)
}

func InterfaceStatisticsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["rx_packets"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["rx_packets"] = "RxPackets"
	fields["rx_errors"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["rx_errors"] = "RxErrors"
	fields["rx_bytes"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["rx_bytes"] = "RxBytes"
	fields["tx_errors"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["tx_errors"] = "TxErrors"
	fields["tx_bytes"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["tx_bytes"] = "TxBytes"
	fields["tx_packets"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["tx_packets"] = "TxPackets"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.interface_statistics", fields, reflect.TypeOf(InterfaceStatistics{}), fieldNameMap, validators)
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
	fields["ips"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["ips"] = "Ips"
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["services"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(MgmtServiceEntryBindingType), reflect.TypeOf([]MgmtServiceEntry{})))
	fieldNameMap["services"] = "Services"
	fields["group_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["group_path"] = "GroupPath"
	fields["ip_attachment_pairs"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(IpAttachmentPairBindingType), reflect.TypeOf([]IpAttachmentPair{})))
	fieldNameMap["ip_attachment_pairs"] = "IpAttachmentPairs"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
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

func NetworkStatusEntryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["issues_found"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["issues_found"] = "IssuesFound"
	fields["ip_address"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["ip_address"] = "IpAddress"
	fields["host_ips"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["host_ips"] = "HostIps"
	fields["issues"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["issues"] = "Issues"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.network_status_entry", fields, reflect.TypeOf(NetworkStatusEntry{}), fieldNameMap, validators)
}

func NetworkStatusKeyBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["network_type"] = bindings.NewStringType()
	fieldNameMap["network_type"] = "NetworkType"
	fields["context"] = bindings.NewStringType()
	fieldNameMap["context"] = "Context"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.network_status_key", fields, reflect.TypeOf(NetworkStatusKey{}), fieldNameMap, validators)
}

func NetworkStatusKeyValuePairBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["values"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(NetworkStatusEntryBindingType), reflect.TypeOf([]NetworkStatusEntry{})))
	fieldNameMap["values"] = "Values"
	fields["key"] = bindings.NewOptionalType(bindings.NewReferenceType(NetworkStatusKeyBindingType))
	fieldNameMap["key"] = "Key"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.network_status_key_value_pair", fields, reflect.TypeOf(NetworkStatusKeyValuePair{}), fieldNameMap, validators)
}

func NetworkStatusListResultBindingType() bindings.BindingType {
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
	fields["results"] = bindings.NewListType(bindings.NewReferenceType(NetworkStatusKeyValuePairBindingType), reflect.TypeOf([]NetworkStatusKeyValuePair{}))
	fieldNameMap["results"] = "Results"
	fields["issues"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["issues"] = "Issues"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.network_status_list_result", fields, reflect.TypeOf(NetworkStatusListResult{}), fieldNameMap, validators)
}

func OwnerResourceLinkBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["action"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["action"] = "Action"
	fields["href"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["href"] = "Href"
	fields["rel"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["rel"] = "Rel"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.owner_resource_link", fields, reflect.TypeOf(OwnerResourceLink{}), fieldNameMap, validators)
}

func PrefixInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["prefixes"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["prefixes"] = "Prefixes"
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.prefix_info", fields, reflect.TypeOf(PrefixInfo{}), fieldNameMap, validators)
}

func PrefixesListResultBindingType() bindings.BindingType {
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
	fields["results"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(PrefixInfoBindingType), reflect.TypeOf([]PrefixInfo{})))
	fieldNameMap["results"] = "Results"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.prefixes_list_result", fields, reflect.TypeOf(PrefixesListResult{}), fieldNameMap, validators)
}

func PublicIpBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["ip"] = "Ip"
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
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
	fields["vpn_internet_ips"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["vpn_internet_ips"] = "VpnInternetIps"
	fields["mgw_snat_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["mgw_snat_ip"] = "MgwSnatIp"
	fields["compute_gateway"] = bindings.NewStringType()
	fieldNameMap["compute_gateway"] = "ComputeGateway"
	fields["all_vpn_interface_label"] = bindings.NewStringType()
	fieldNameMap["all_vpn_interface_label"] = "AllVpnInterfaceLabel"
	fields["vpn_dx_ips"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["vpn_dx_ips"] = "VpnDxIps"
	fields["all_uplink_interface_label"] = bindings.NewStringType()
	fieldNameMap["all_uplink_interface_label"] = "AllUplinkInterfaceLabel"
	fields["dx_interface_label"] = bindings.NewStringType()
	fieldNameMap["dx_interface_label"] = "DxInterfaceLabel"
	fields["sddc_infra_subnet"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["sddc_infra_subnet"] = "SddcInfraSubnet"
	fields["management_gateway"] = bindings.NewStringType()
	fieldNameMap["management_gateway"] = "ManagementGateway"
	fields["linked_vpc_interface_label"] = bindings.NewStringType()
	fieldNameMap["linked_vpc_interface_label"] = "LinkedVpcInterfaceLabel"
	fields["public_interface_label"] = bindings.NewStringType()
	fieldNameMap["public_interface_label"] = "PublicInterfaceLabel"
	fields["cgw_snat_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cgw_snat_ip"] = "CgwSnatIp"
	fields["provider_name"] = bindings.NewStringType()
	fieldNameMap["provider_name"] = "ProviderName"
	fields["mgmt_subnet"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["mgmt_subnet"] = "MgmtSubnet"
	fields["management_gateway_label"] = bindings.NewStringType()
	fieldNameMap["management_gateway_label"] = "ManagementGatewayLabel"
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

func TaskPropertiesBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["_self"] = bindings.NewOptionalType(bindings.NewReferenceType(SelfResourceLinkBindingType))
	fieldNameMap["_self"] = "Self"
	fields["_links"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ResourceLinkBindingType), reflect.TypeOf([]ResourceLink{})))
	fieldNameMap["_links"] = "Links"
	fields["_schema"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["_schema"] = "Schema"
	fields["status"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["status"] = "Status"
	fields["async_response_available"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["async_response_available"] = "AsyncResponseAvailable"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["start_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["start_time"] = "StartTime"
	fields["cancelable"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["cancelable"] = "Cancelable"
	fields["request_method"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["request_method"] = "RequestMethod"
	fields["user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["user"] = "User"
	fields["progress"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["progress"] = "Progress"
	fields["message"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["message"] = "Message"
	fields["request_uri"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["request_uri"] = "RequestUri"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["end_time"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["end_time"] = "EndTime"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.task_properties", fields, reflect.TypeOf(TaskProperties{}), fieldNameMap, validators)
}

func TraceflowActionBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["reason"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["reason"] = "Reason"
	fields["component_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["component_name"] = "ComponentName"
	fields["resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["component_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["component_type"] = "ComponentType"
	fields["component_sub_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["component_sub_type"] = "ComponentSubType"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.traceflow_action", fields, reflect.TypeOf(TraceflowAction{}), fieldNameMap, validators)
}

func TraceflowActionListResultsBindingType() bindings.BindingType {
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
	fields["results"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(TraceflowActionBindingType), reflect.TypeOf([]TraceflowAction{})))
	fieldNameMap["results"] = "Results"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.traceflow_action_list_results", fields, reflect.TypeOf(TraceflowActionListResults{}), fieldNameMap, validators)
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

func VdrLifBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["ip"] = "Ip"
	fields["netmask"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["netmask"] = "Netmask"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["vlan_id"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["vlan_id"] = "VlanId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.vdr_lif", fields, reflect.TypeOf(VdrLif{}), fieldNameMap, validators)
}

func VdrRouteBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["destination"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["destination"] = "Destination"
	fields["gateway"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["gateway"] = "Gateway"
	fields["lif_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["lif_id"] = "LifId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.model.vdr_route", fields, reflect.TypeOf(VdrRoute{}), fieldNameMap, validators)
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


