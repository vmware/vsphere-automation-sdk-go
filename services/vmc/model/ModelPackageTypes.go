/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for package: com.vmware.vmc.model.
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


type HostInstanceTypes string

const (
	HostInstanceTypes_I3_METAL HostInstanceTypes = "I3_METAL"
	HostInstanceTypes_R5_METAL HostInstanceTypes = "R5_METAL"
	HostInstanceTypes_I3EN_METAL HostInstanceTypes = "I3EN_METAL"
)

func (h HostInstanceTypes) HostInstanceTypes() bool {
	switch h {
	case HostInstanceTypes_I3_METAL:
		return true
	case HostInstanceTypes_R5_METAL:
		return true
	case HostInstanceTypes_I3EN_METAL:
		return true
	default:
		return false
	}
}


type OfferType string

const (
	OfferType_TERM OfferType = "TERM"
	OfferType_ON_DEMAND OfferType = "ON_DEMAND"
)

func (o OfferType) OfferType() bool {
	switch o {
	case OfferType_TERM:
		return true
	case OfferType_ON_DEMAND:
		return true
	default:
		return false
	}
}


type AbstractEntity struct {
	Updated time.Time
    // User id that last updated this record
	UserId string
    // User id that last updated this record
	UpdatedByUserId string
	Created time.Time
    // Version of this entity format: int32
	Version int64
    // User name that last updated this record
	UpdatedByUserName *string
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


type AccountLinkConfig struct {
    // Boolean flag identifying whether account linking should be delayed or not for the SDDC.
	DelayAccountLink *bool
}

func (s AccountLinkConfig) GetType__() bindings.BindingType {
	return AccountLinkConfigBindingType()
}

func (s AccountLinkConfig) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for AccountLinkConfig._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type AccountLinkSddcConfig struct {
	CustomerSubnetIds []string
    // The ID of the customer connected account to work with.
	ConnectedAccountId *string
}

func (s AccountLinkSddcConfig) GetType__() bindings.BindingType {
	return AccountLinkSddcConfigBindingType()
}

func (s AccountLinkSddcConfig) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for AccountLinkSddcConfig._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type Agent struct {
    // The internal IP address of the agent which is provided by the underlying cloud provider
	InternalIp *string
    // The accessible URL of the agent service, it is resolved to public IP address from the Internet and private IP address within SDDC
	AgentUrl *string
    // The internal management IP address of the agent exposed to the SDDC, which might be different from the internal IP
	ManagementIp *string
    // Boolean flag to indicate if the agent is using FQDN in the certificate
	HostnameVerifierEnabled *bool
    // Boolean flag to indicate if the agent is the master, only the master Agent is accessible
	Master *bool
    // Network netmask of the agent
	NetworkNetmask *string
    // Network gateway of the agent
	NetworkGateway *string
    // The cloud provider
	Provider string
    // Boolean flag to indicate if the agent is using CA signed certificate
	CertEnabled *bool
    // Possible values are: 
    //
    // * Agent#Agent_AGENT_STATE_NOT_READY
    // * Agent#Agent_AGENT_STATE_DEPLOYING
    // * Agent#Agent_AGENT_STATE_CUSTOMIZING
    // * Agent#Agent_AGENT_STATE_READY
    // * Agent#Agent_AGENT_STATE_DELETING
    // * Agent#Agent_AGENT_STATE_DELETED
    // * Agent#Agent_AGENT_STATE_FAILED
    //
    //  Agent state
	AgentState *string
}
// Identifier denoting this class, when it is used in polymorphic context. 
//
// This value should be assigned to the property which is used to discriminate the actual type used in the polymorphic context.
const Agent__TYPE_IDENTIFIER = "Agent"
const Agent_AGENT_STATE_NOT_READY = "NOT_READY"
const Agent_AGENT_STATE_DEPLOYING = "DEPLOYING"
const Agent_AGENT_STATE_CUSTOMIZING = "CUSTOMIZING"
const Agent_AGENT_STATE_READY = "READY"
const Agent_AGENT_STATE_DELETING = "DELETING"
const Agent_AGENT_STATE_DELETED = "DELETED"
const Agent_AGENT_STATE_FAILED = "FAILED"

func (s Agent) GetType__() bindings.BindingType {
	return AgentBindingType()
}

func (s Agent) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for Agent._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// the AmiInfo used for deploying esx of the sddc
type AmiInfo struct {
    // instance type of the esx ami
	InstanceType *string
    // the region of the esx ami
	Region *string
    // the ami id for the esx
	Id *string
    // the name of the esx ami
	Name *string
}

func (s AmiInfo) GetType__() bindings.BindingType {
	return AmiInfoBindingType()
}

func (s AmiInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for AmiInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type AvailableZoneInfo struct {
	Subnets []Subnet
    // available zone id
	Id *string
}

func (s AvailableZoneInfo) GetType__() bindings.BindingType {
	return AvailableZoneInfoBindingType()
}

func (s AvailableZoneInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for AvailableZoneInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type AwsAgent struct {
	InstanceId *string
	KeyPair *AwsKeyPair
    // The internal IP address of the agent which is provided by the underlying cloud provider
	InternalIp *string
    // The accessible URL of the agent service, it is resolved to public IP address from the Internet and private IP address within SDDC
	AgentUrl *string
    // The internal management IP address of the agent exposed to the SDDC, which might be different from the internal IP
	ManagementIp *string
    // Boolean flag to indicate if the agent is using FQDN in the certificate
	HostnameVerifierEnabled *bool
    // Boolean flag to indicate if the agent is the master, only the master Agent is accessible
	Master *bool
    // Network netmask of the agent
	NetworkNetmask *string
    // Network gateway of the agent
	NetworkGateway *string
    // The cloud provider
	Provider string
    // Boolean flag to indicate if the agent is using CA signed certificate
	CertEnabled *bool
    // Possible values are: 
    //
    // * Agent#Agent_AGENT_STATE_NOT_READY
    // * Agent#Agent_AGENT_STATE_DEPLOYING
    // * Agent#Agent_AGENT_STATE_CUSTOMIZING
    // * Agent#Agent_AGENT_STATE_READY
    // * Agent#Agent_AGENT_STATE_DELETING
    // * Agent#Agent_AGENT_STATE_DELETED
    // * Agent#Agent_AGENT_STATE_FAILED
    //
    //  Agent state
	AgentState *string
}
// Identifier denoting this class, when it is used in polymorphic context. 
//
// This value should be assigned to the property which is used to discriminate the actual type used in the polymorphic context.
const AwsAgent__TYPE_IDENTIFIER = "AWS"

func (s AwsAgent) GetType__() bindings.BindingType {
	return AwsAgentBindingType()
}

func (s AwsAgent) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for AwsAgent._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type AwsCloudProvider struct {
	Regions []string
    // Name of the Cloud Provider
	Provider string
}
// Identifier denoting this class, when it is used in polymorphic context. 
//
// This value should be assigned to the property which is used to discriminate the actual type used in the polymorphic context.
const AwsCloudProvider__TYPE_IDENTIFIER = "AWS"

func (s AwsCloudProvider) GetType__() bindings.BindingType {
	return AwsCloudProviderBindingType()
}

func (s AwsCloudProvider) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for AwsCloudProvider._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type AwsCompatibleSubnets struct {
	CustomerAvailableZones []string
	VpcMap map[string]VpcInfoSubnets
}

func (s AwsCompatibleSubnets) GetType__() bindings.BindingType {
	return AwsCompatibleSubnetsBindingType()
}

func (s AwsCompatibleSubnets) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for AwsCompatibleSubnets._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type AwsCustomerConnectedAccount struct {
	Updated time.Time
    // User id that last updated this record
	UserId string
    // User id that last updated this record
	UpdatedByUserId string
	Created time.Time
    // Version of this entity format: int32
	Version int64
    // User name that last updated this record
	UpdatedByUserName *string
    // User name that last updated this record
	UserName string
    // Unique ID for this entity
	Id string
	PolicyPayerArn *string
    // Provides a map of regions to availability zones from the shadow account's perspective
	RegionToAzToShadowMapping map[string]map[string]string
	OrgId *string
	CfStackName *string
    // Possible values are: 
    //
    // * AwsCustomerConnectedAccount#AwsCustomerConnectedAccount_STATE_ACTIVE
    // * AwsCustomerConnectedAccount#AwsCustomerConnectedAccount_STATE_BROKEN
    // * AwsCustomerConnectedAccount#AwsCustomerConnectedAccount_STATE_DELETED
	State *string
	AccountNumber *string
	PolicyServiceArn *string
	PolicyExternalId *string
	PolicyPayerLinkedArn *string
}
const AwsCustomerConnectedAccount_STATE_ACTIVE = "ACTIVE"
const AwsCustomerConnectedAccount_STATE_BROKEN = "BROKEN"
const AwsCustomerConnectedAccount_STATE_DELETED = "DELETED"

func (s AwsCustomerConnectedAccount) GetType__() bindings.BindingType {
	return AwsCustomerConnectedAccountBindingType()
}

func (s AwsCustomerConnectedAccount) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for AwsCustomerConnectedAccount._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type AwsEsxHost struct {
	InternalPublicIpPool []SddcPublicIp
	Name *string
    // Availability zone where the host is provisioned.
	AvailabilityZone *string
	EsxId *string
	Hostname *string
	Provider string
    // Backing cloud provider instance type for host.
	InstanceType *string
	MacAddress *string
	CustomProperties map[string]string
    // Possible values are: 
    //
    // * EsxHost#EsxHost_ESX_STATE_DEPLOYING
    // * EsxHost#EsxHost_ESX_STATE_INITIALIZING
    // * EsxHost#EsxHost_ESX_STATE_PROVISIONED
    // * EsxHost#EsxHost_ESX_STATE_READY
    // * EsxHost#EsxHost_ESX_STATE_DELETING
    // * EsxHost#EsxHost_ESX_STATE_DELETED
    // * EsxHost#EsxHost_ESX_STATE_FAILED
    // * EsxHost#EsxHost_ESX_STATE_ADDING_TO_VCENTER
    // * EsxHost#EsxHost_ESX_STATE_DELETING_FROM_VCENTER
    // * EsxHost#EsxHost_ESX_STATE_PENDING_CLOUD_DELETION
	EsxState *string
}
// Identifier denoting this class, when it is used in polymorphic context. 
//
// This value should be assigned to the property which is used to discriminate the actual type used in the polymorphic context.
const AwsEsxHost__TYPE_IDENTIFIER = "AWS"

func (s AwsEsxHost) GetType__() bindings.BindingType {
	return AwsEsxHostBindingType()
}

func (s AwsEsxHost) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for AwsEsxHost._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type AwsKeyPair struct {
	KeyName *string
	KeyFingerprint *string
	KeyMaterial *string
}

func (s AwsKeyPair) GetType__() bindings.BindingType {
	return AwsKeyPairBindingType()
}

func (s AwsKeyPair) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for AwsKeyPair._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type AwsKmsInfo struct {
    // The ARN associated with the customer master key for this cluster.
	AmazonResourceName string
}

func (s AwsKmsInfo) GetType__() bindings.BindingType {
	return AwsKmsInfoBindingType()
}

func (s AwsKmsInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for AwsKmsInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type AwsSddcConfig struct {
	Region string
    // Indicates the desired licensing support, if any, of Microsoft software.
	MsftLicenseConfig *MsftLicensingConfig
    // AWS VPC IP range. Only prefix of 16 or 20 is currently supported.
	VpcCidr *string
    // The instance type for the esx hosts in the primary cluster of the SDDC.
	HostInstanceType *HostInstanceTypes
    // skip creating vxlan for compute gateway for SDDC provisioning
	SkipCreatingVxlan *bool
    // VXLAN IP subnet in CIDR for compute gateway
	VxlanSubnet *string
    // Possible values are: 
    //
    // * SddcConfig#SddcConfig_SIZE_NSX_SMALL
    // * SddcConfig#SddcConfig_SIZE_MEDIUM
    // * SddcConfig#SddcConfig_SIZE_LARGE
    // * SddcConfig#SddcConfig_SIZE_NSX_LARGE
    //
    //  The size of the vCenter and NSX appliances. \"large\" sddcSize corresponds to a 'large' vCenter appliance and 'large' NSX appliance. 'medium' sddcSize corresponds to 'medium' vCenter appliance and 'medium' NSX appliance. Value defaults to 'medium'.
	Size *string
    // The storage capacity value to be requested for the sddc primary cluster, in GiBs. If provided, instead of using the direct-attached storage, a capacity value amount of seperable storage will be used. format: int64
	StorageCapacity *int64
	Name string
    // A list of the SDDC linking configurations to use.
	AccountLinkSddcConfig []AccountLinkSddcConfig
    // If provided, will be assigned as SDDC id of the provisioned SDDC. format: UUID
	SddcId *string
	NumHosts int64
    // Denotes the sddc type , if the value is null or empty, the type is considered as default.
	SddcType *string
    // The account linking configuration, we will keep this one and remove accountLinkSddcConfig finally.
	AccountLinkConfig *AccountLinkConfig
    // Possible values are: 
    //
    // * SddcConfig#SddcConfig_PROVIDER_AWS
    // * SddcConfig#SddcConfig_PROVIDER_ZEROCLOUD
    //
    //  Determines what additional properties are available based on cloud provider.
	Provider string
    // The SSO domain name to use for vSphere users. If not specified, vmc.local will be used.
	SsoDomain *string
    // If provided, configuration from the template will applied to the provisioned SDDC. format: UUID
	SddcTemplateId *string
    // Possible values are: 
    //
    // * SddcConfig#SddcConfig_DEPLOYMENT_TYPE_SINGLEAZ
    // * SddcConfig#SddcConfig_DEPLOYMENT_TYPE_MULTIAZ
    //
    //  Denotes if request is for a SingleAZ or a MultiAZ SDDC. Default is SingleAZ.
	DeploymentType *string
}
// Identifier denoting this class, when it is used in polymorphic context. 
//
// This value should be assigned to the property which is used to discriminate the actual type used in the polymorphic context.
const AwsSddcConfig__TYPE_IDENTIFIER = "AWS"

func (s AwsSddcConfig) GetType__() bindings.BindingType {
	return AwsSddcConfigBindingType()
}

func (s AwsSddcConfig) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for AwsSddcConfig._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type AwsSddcConnection struct {
	Updated time.Time
    // User id that last updated this record
	UserId string
    // User id that last updated this record
	UpdatedByUserId string
	Created time.Time
    // Version of this entity format: int32
	Version int64
    // User name that last updated this record
	UpdatedByUserName *string
    // User name that last updated this record
	UserName string
    // Unique ID for this entity
	Id string
    // The CIDR block of the customer's subnet this link is in.
	CidrBlockSubnet *string
    // The corresponding connected (customer) account UUID this connection is attached to.
	ConnectedAccountId *string
    // Which group the ENIs belongs to. (deprecated)
	EniGroup *string
    // The ID of the subnet this link is to.
	SubnetId *string
    // Determines whether the CGW is present in this connection set or not. Used for multi-az deployments.
	CgwPresent *bool
    // The org this link belongs to.
	OrgId *string
    // The SDDC this link is used for.
	SddcId *string
    // The CIDR block of the customer's VPC.
	CidrBlockVpc *string
    // The order of the connection
	ConnectionOrder *int64
    // Possible values are: 
    //
    // * AwsSddcConnection#AwsSddcConnection_STATE_ACTIVE
    // * AwsSddcConnection#AwsSddcConnection_STATE_BROKEN
    // * AwsSddcConnection#AwsSddcConnection_STATE_DELETED
    //
    //  The state of the connection.
	State *string
    // Which availability zone is this connection in?
	SubnetAvailabilityZone *string
    // The VPC ID of the subnet this link is to.
	VpcId *string
    // A list of all ENIs used for this connection.
	CustomerEniInfos []CustomerEniInfo
    // The default routing table in the customer's VPC.
	DefaultRouteTable *string
}
const AwsSddcConnection_STATE_ACTIVE = "ACTIVE"
const AwsSddcConnection_STATE_BROKEN = "BROKEN"
const AwsSddcConnection_STATE_DELETED = "DELETED"

func (s AwsSddcConnection) GetType__() bindings.BindingType {
	return AwsSddcConnectionBindingType()
}

func (s AwsSddcConnection) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for AwsSddcConnection._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type AwsSddcResourceConfig struct {
	BackupRestoreBucket *string
	PublicIpPool []SddcPublicIp
	VpcInfo *VpcInfo
	KmsVpcEndpoint *KmsVpcEndpoint
    // maximum number of public IP that user can allocate.
	MaxNumPublicIp *int64
	AccountLinkSddcConfig []SddcLinkConfig
	VsanEncryptionConfig *VsanEncryptionConfig
	VpcInfoPeeredAgent *VpcInfo
    // Name for management appliance network.
	MgmtApplianceNetworkName *string
    // if true, NSX-T UI is enabled.
	Nsxt *bool
    // Management Gateway Id
	MgwId *string
    // URL of the NSX Manager
	NsxMgrUrl *string
    // PSC internal management IP
	PscManagementIp *string
    // URL of the PSC server
	PscUrl *string
	Cgws []string
    // Availability zones over which esx hosts are provisioned. MultiAZ SDDCs will have hosts provisioned over two availability zones while SingleAZ SDDCs will provision over one.
	AvailabilityZones []string
    // The ManagedObjectReference of the management Datastore
	ManagementDs *string
    // nsx api entire base url
	NsxApiPublicEndpointUrl *string
	CustomProperties map[string]string
    // Password for vCenter SDDC administrator
	CloudPassword *string
    // Possible values are: 
    //
    // * SddcResourceConfig#SddcResourceConfig_PROVIDER_AWS
    // * SddcResourceConfig#SddcResourceConfig_PROVIDER_ZEROCLOUD
    //
    //  Discriminator for additional properties
	Provider string
    // List of clusters in the SDDC.
	Clusters []Cluster
    // vCenter internal management IP
	VcManagementIp *string
	SddcNetworks []string
    // Username for vCenter SDDC administrator
	CloudUsername *string
	EsxHosts []AwsEsxHost
    // NSX Manager internal management IP
	NsxMgrManagementIp *string
    // unique id of the vCenter server
	VcInstanceId *string
    // Cluster Id to add ESX workflow
	EsxClusterId *string
    // vCenter public IP
	VcPublicIp *string
    // skip creating vxlan for compute gateway for SDDC provisioning
	SkipCreatingVxlan *bool
    // URL of the vCenter server
	VcUrl *string
	SddcManifest *SddcManifest
    // VXLAN IP subnet
	VxlanSubnet *string
    // Group name for vCenter SDDC administrator
	CloudUserGroup *string
	ManagementRp *string
    // region in which sddc is deployed
	Region *string
    // Availability zone where the witness node is provisioned for a MultiAZ SDDC. This is null for a SingleAZ SDDC.
	WitnessAvailabilityZone *string
    // sddc identifier
	SddcId *string
	PopAgentXeniConnection *PopAgentXeniConnection
	SddcSize *SddcSize
    // List of Controller IPs
	NsxControllerIps []string
    // ESX host subnet
	EsxHostSubnet *string
    // The SSO domain name to use for vSphere users
	SsoDomain *string
    // Possible values are: 
    //
    // * SddcResourceConfig#SddcResourceConfig_DEPLOYMENT_TYPE_SINGLE_AZ
    // * SddcResourceConfig#SddcResourceConfig_DEPLOYMENT_TYPE_MULTI_AZ
    //
    //  Denotes if this is a SingleAZ SDDC or a MultiAZ SDDC.
	DeploymentType *string
    // The Microsoft license status of this SDDC.
	MsftLicenseConfig *MsftLicensingConfig
	NsxtAddons *NsxtAddons
    // if true, use the private IP addresses to register DNS records for the management VMs
	DnsWithManagementVmPrivateIp *bool
}
// Identifier denoting this class, when it is used in polymorphic context. 
//
// This value should be assigned to the property which is used to discriminate the actual type used in the polymorphic context.
const AwsSddcResourceConfig__TYPE_IDENTIFIER = "AWS"

func (s AwsSddcResourceConfig) GetType__() bindings.BindingType {
	return AwsSddcResourceConfigBindingType()
}

func (s AwsSddcResourceConfig) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for AwsSddcResourceConfig._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type AwsSubnet struct {
    // The connected account ID this subnet is accessible through. This is an internal ID formatted as a UUID specific to Skyscraper.
	ConnectedAccountId *string
    // The region this subnet is in, usually in the form of country code, general location, and a number (ex. us-west-2).
	RegionName *string
    // The availability zone this subnet is in, which should be the region name plus one extra letter (ex. us-west-2a).
	AvailabilityZone *string
    // The subnet ID in AWS, provided in the form 'subnet-######'.
	SubnetId *string
    // The CIDR block of the subnet, in the form of '#.#.#.#/#'.
	SubnetCidrBlock *string
    // Flag indicating whether this subnet is compatible. If true, this is a valid choice for the customer to deploy a SDDC in.
	IsCompatible *bool
    // The VPC ID the subnet resides in within AWS. Tends to be 'vpc-#######'.
	VpcId *string
    // The CIDR block of the VPC, in the form of '#.#.#.#/#'.
	VpcCidrBlock *string
    // Optional field (may not be provided by AWS), indicates the found name tag for the subnet.
	Name *string
}

func (s AwsSubnet) GetType__() bindings.BindingType {
	return AwsSubnetBindingType()
}

func (s AwsSubnet) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for AwsSubnet._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type CloudProvider struct {
    // Name of the Cloud Provider
	Provider string
}
// Identifier denoting this class, when it is used in polymorphic context. 
//
// This value should be assigned to the property which is used to discriminate the actual type used in the polymorphic context.
const CloudProvider__TYPE_IDENTIFIER = "CloudProvider"

func (s CloudProvider) GetType__() bindings.BindingType {
	return CloudProviderBindingType()
}

func (s CloudProvider) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for CloudProvider._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type Cluster struct {
	EsxHostList []AwsEsxHost
    // The Microsoft license configuration of this cluster.
	MsftLicenseConfig *MsftLicensingConfig
    // Possible values are: 
    //
    // * Cluster#Cluster_CLUSTER_STATE_DEPLOYING
    // * Cluster#Cluster_CLUSTER_STATE_ADDING_HOSTS
    // * Cluster#Cluster_CLUSTER_STATE_READY
    // * Cluster#Cluster_CLUSTER_STATE_FAILED
	ClusterState *string
    // AWS Key Management Service information associated with this cluster
	AwsKmsInfo *AwsKmsInfo
    // The capacity of this cluster.
	ClusterCapacity *EntityCapacity
    // Information of the hosts added to this cluster
	EsxHostInfo *EsxHostInfo
    // Number of cores enabled on ESX hosts added to this cluster format: int32
	HostCpuCoresCount *int64
	ClusterId string
	ClusterName *string
}
const Cluster_CLUSTER_STATE_DEPLOYING = "DEPLOYING"
const Cluster_CLUSTER_STATE_ADDING_HOSTS = "ADDING_HOSTS"
const Cluster_CLUSTER_STATE_READY = "READY"
const Cluster_CLUSTER_STATE_FAILED = "FAILED"

func (s Cluster) GetType__() bindings.BindingType {
	return ClusterBindingType()
}

func (s Cluster) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for Cluster._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type ClusterConfig struct {
    // Customize CPU cores on hosts in a cluster. Specify number of cores to be enabled on hosts in a cluster. format: int32
	HostCpuCoresCount *int64
    // The instance type for the esx hosts added to this cluster.
	HostInstanceType *HostInstanceTypes
    // For EBS-backed instances only, the requested storage capacity in GiB. format: int64
	StorageCapacity *int64
    // The desired Microsoft license status to apply to this cluster.
	MsftLicenseConfig *MsftLicensingConfig
	NumHosts int64
}

func (s ClusterConfig) GetType__() bindings.BindingType {
	return ClusterConfigBindingType()
}

func (s ClusterConfig) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ClusterConfig._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type ClusterReconfigureParams struct {
    // The final desired storage capacity after reconfiguring the cluster in GiB. format: int64
	StorageCapacity *int64
    // Bias value as obtained from the storage constraints call.
	Bias *string
    // Number of hosts in the cluster after reconfiguring. format: int32
	NumHosts int64
}

func (s ClusterReconfigureParams) GetType__() bindings.BindingType {
	return ClusterReconfigureParamsBindingType()
}

func (s ClusterReconfigureParams) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ClusterReconfigureParams._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type ComputeGatewayTemplate struct {
	PublicIp *SddcPublicIp
	PrimaryDns *string
	SecondaryDns *string
	FirewallRules []FirewallRule
	Vpns []Vpn
	LogicalNetworks []LogicalNetwork
	NatRules []NatRule
	L2Vpn *data.StructValue
}

func (s ComputeGatewayTemplate) GetType__() bindings.BindingType {
	return ComputeGatewayTemplateBindingType()
}

func (s ComputeGatewayTemplate) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ComputeGatewayTemplate._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Represents a configuration spec for any sddc provision operation.
type ConfigSpec struct {
    // Indicates after how many days the sddc should expire
	ExpiryInDays *int64
    // Map of region to instance types available in that region
	Availability map[string][]InstanceTypeConfig
	SddcSizes []string
}

func (s ConfigSpec) GetType__() bindings.BindingType {
	return ConfigSpecBindingType()
}

func (s ConfigSpec) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ConfigSpec._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type ConnectivityAgentValidation struct {
    // Possible values are: 
    //
    // * ConnectivityAgentValidation#ConnectivityAgentValidation_SOURCE_VCENTER
    // * ConnectivityAgentValidation#ConnectivityAgentValidation_SOURCE_SRM
    // * ConnectivityAgentValidation#ConnectivityAgentValidation_SOURCE_VR
    //
    //  source appliance of connectivity test, i.e. VCENTER, SRM, VR.
	Source *string
    // Possible values are: 
    //
    // * ConnectivityAgentValidation#ConnectivityAgentValidation_TYPE_PING
    // * ConnectivityAgentValidation#ConnectivityAgentValidation_TYPE_TRACEROUTE
    // * ConnectivityAgentValidation#ConnectivityAgentValidation_TYPE_DNS
    // * ConnectivityAgentValidation#ConnectivityAgentValidation_TYPE_CONNECTIVITY
    // * ConnectivityAgentValidation#ConnectivityAgentValidation_TYPE_CURL
    //
    //  type of connectivity test, i.e. PING, TRACEROUTE, DNS, CONNECTIVITY, CURL. For CONNECTIVITY and CURL tests only, please specify the ports to be tested against.
	Type_ *string
    // TCP ports ONLY for CONNECTIVITY and CURL tests.
	Ports []string
    // URL path ONLY for CURL tests.
	Path *string
}
const ConnectivityAgentValidation_SOURCE_VCENTER = "VCENTER"
const ConnectivityAgentValidation_SOURCE_SRM = "SRM"
const ConnectivityAgentValidation_SOURCE_VR = "VR"
const ConnectivityAgentValidation_TYPE_PING = "PING"
const ConnectivityAgentValidation_TYPE_TRACEROUTE = "TRACEROUTE"
const ConnectivityAgentValidation_TYPE_DNS = "DNS"
const ConnectivityAgentValidation_TYPE_CONNECTIVITY = "CONNECTIVITY"
const ConnectivityAgentValidation_TYPE_CURL = "CURL"

func (s ConnectivityAgentValidation) GetType__() bindings.BindingType {
	return ConnectivityAgentValidationBindingType()
}

func (s ConnectivityAgentValidation) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ConnectivityAgentValidation._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type ConnectivityValidationGroup struct {
    // Possible values are: 
    //
    // * ConnectivityValidationGroup#ConnectivityValidationGroup_ID_HLM
    // * ConnectivityValidationGroup#ConnectivityValidationGroup_ID_DRAAS
    //
    //  test group id, currently, only HLM.
	Id *string
    // Name of the test group.
	Name *string
    // List of sub groups.
	SubGroups []ConnectivityValidationSubGroup
}
const ConnectivityValidationGroup_ID_HLM = "HLM"
const ConnectivityValidationGroup_ID_DRAAS = "DRAAS"

func (s ConnectivityValidationGroup) GetType__() bindings.BindingType {
	return ConnectivityValidationGroupBindingType()
}

func (s ConnectivityValidationGroup) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ConnectivityValidationGroup._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type ConnectivityValidationGroups struct {
    // List of groups.
	Groups []ConnectivityValidationGroup
}

func (s ConnectivityValidationGroups) GetType__() bindings.BindingType {
	return ConnectivityValidationGroupsBindingType()
}

func (s ConnectivityValidationGroups) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ConnectivityValidationGroups._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type ConnectivityValidationInput struct {
    // Possible values are: 
    //
    // * ConnectivityValidationInput#ConnectivityValidationInput_ID_HOSTNAME
    // * ConnectivityValidationInput#ConnectivityValidationInput_ID_HOST_IP
    // * ConnectivityValidationInput#ConnectivityValidationInput_ID_HOSTNAME_OR_IP
    //
    //  input value type, i.e. HOSTNAME_OR_IP, HOST_IP, HOSTNAME. Accept FQDN or IP address as input value when id = HOSTNAME_OR_IP, accept FQDN ONLY when id = HOSTNAME, accept IP address ONLY when id = HOST_IP.
	Id *string
    // the FQDN or IP address to run the test against, use \\#primary-dns or \\#secondary-dns as the on-prem primary/secondary DNS server IP.
	Value *string
    // (Optional, for UI display only) input value label.
	Label *string
}
const ConnectivityValidationInput_ID_HOSTNAME = "HOSTNAME"
const ConnectivityValidationInput_ID_HOST_IP = "HOST_IP"
const ConnectivityValidationInput_ID_HOSTNAME_OR_IP = "HOSTNAME_OR_IP"

func (s ConnectivityValidationInput) GetType__() bindings.BindingType {
	return ConnectivityValidationInputBindingType()
}

func (s ConnectivityValidationInput) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ConnectivityValidationInput._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type ConnectivityValidationSubGroup struct {
    // List of user inputs for the sub group.
	Inputs []ConnectivityValidationInput
    // List of connectivity tests.
	Tests []ConnectivityAgentValidation
    // Name of the sub-group.
	Label *string
    // Help text.
	Help *string
    // Possible values are: 
    //
    // * ConnectivityValidationSubGroup#ConnectivityValidationSubGroup_ID_PRIMARY_DNS
    // * ConnectivityValidationSubGroup#ConnectivityValidationSubGroup_ID_SECONDARY_DNS
    // * ConnectivityValidationSubGroup#ConnectivityValidationSubGroup_ID_ONPREM_VCENTER
    // * ConnectivityValidationSubGroup#ConnectivityValidationSubGroup_ID_ONPREM_PSC
    // * ConnectivityValidationSubGroup#ConnectivityValidationSubGroup_ID_ACTIVE_DIRECTORY
    // * ConnectivityValidationSubGroup#ConnectivityValidationSubGroup_ID_ONPREM_ESX
    // * ConnectivityValidationSubGroup#ConnectivityValidationSubGroup_ID_DRAAS_ONPREM_VCENTER
    // * ConnectivityValidationSubGroup#ConnectivityValidationSubGroup_ID_DRAAS_ONPREM_PSC
    // * ConnectivityValidationSubGroup#ConnectivityValidationSubGroup_ID_DRAAS_ONPREM_SRM
    // * ConnectivityValidationSubGroup#ConnectivityValidationSubGroup_ID_DRAAS_ONPREM_VR
    //
    //  subGroup id, i.e. PRIMARY_DNS, SECONDARY_DNS, ONPREM_VCENTER, ONPREM_PSC, ACTIVE_DIRECTORY, ONPREM_ESX, DRAAS_ONPREM_VCENTER, DRAAS_ONPREM_PSC, DRAAS_ONPREM_SRM and DRAAS_ONPREM_VR.
	Id *string
}
const ConnectivityValidationSubGroup_ID_PRIMARY_DNS = "PRIMARY_DNS"
const ConnectivityValidationSubGroup_ID_SECONDARY_DNS = "SECONDARY_DNS"
const ConnectivityValidationSubGroup_ID_ONPREM_VCENTER = "ONPREM_VCENTER"
const ConnectivityValidationSubGroup_ID_ONPREM_PSC = "ONPREM_PSC"
const ConnectivityValidationSubGroup_ID_ACTIVE_DIRECTORY = "ACTIVE_DIRECTORY"
const ConnectivityValidationSubGroup_ID_ONPREM_ESX = "ONPREM_ESX"
const ConnectivityValidationSubGroup_ID_DRAAS_ONPREM_VCENTER = "DRAAS_ONPREM_VCENTER"
const ConnectivityValidationSubGroup_ID_DRAAS_ONPREM_PSC = "DRAAS_ONPREM_PSC"
const ConnectivityValidationSubGroup_ID_DRAAS_ONPREM_SRM = "DRAAS_ONPREM_SRM"
const ConnectivityValidationSubGroup_ID_DRAAS_ONPREM_VR = "DRAAS_ONPREM_VR"

func (s ConnectivityValidationSubGroup) GetType__() bindings.BindingType {
	return ConnectivityValidationSubGroupBindingType()
}

func (s ConnectivityValidationSubGroup) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ConnectivityValidationSubGroup._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Indicates a single cross-account ENI and its characteristics.
type CustomerEniInfo struct {
    // Indicates list of secondary IP created for this ENI.
	SecondaryIpAddresses []string
    // Interface ID on customer account.
	EniId *string
    // Indicates primary address of the ENI.
	PrimaryIpAddress *string
}

func (s CustomerEniInfo) GetType__() bindings.BindingType {
	return CustomerEniInfoBindingType()
}

func (s CustomerEniInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for CustomerEniInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// information for EBS-backed VSAN configuration
type EbsBackedVsanConfig struct {
    // instance type for EBS-backed VSAN configuration
	InstanceType *string
}

func (s EbsBackedVsanConfig) GetType__() bindings.BindingType {
	return EbsBackedVsanConfigBindingType()
}

func (s EbsBackedVsanConfig) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for EbsBackedVsanConfig._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Information of the x-eni created.
type EniInfo struct {
    // Subnet it belongs to.
	SubnetId *string
    // Interface ID.
	Id *string
    // Security Group of Eni.
	SecurityGroupId *string
    // Private ip of eni.
	PrivateIp *string
    // Mac address of nic.
	MacAddress *string
}

func (s EniInfo) GetType__() bindings.BindingType {
	return EniInfoBindingType()
}

func (s EniInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for EniInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Decribes the capacity of a given entity.
type EntityCapacity struct {
    // The storage capacity for the given entity in GiB.
	StorageCapacityGib *int64
    // The memory capacity for the given entity in GiB.
	MemoryCapacityGib *int64
    // The number of CPU cores for the given entity.
	TotalNumberOfCores *int64
    // The number of SSDs for the given entity.
	NumberOfSsds *int64
    // The CPU capacity for the given entity in Ghz.
	CpuCapacityGhz *float64
    // The number of sockets for the given entity.
	NumberOfSockets *int64
}

func (s EntityCapacity) GetType__() bindings.BindingType {
	return EntityCapacityBindingType()
}

func (s EntityCapacity) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for EntityCapacity._GetDataValue method - %s",
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


type EsxConfig struct {
    // Availability zone where the hosts should be provisioned. (Can be specified only for privileged host operations).
	AvailabilityZone *string
    // An option to indicate if the host needs to be strictly placed in a placement group. Fail the operation otherwise.
	StrictPlacement *bool
    // An optional cluster id if the esxs operation has to be on a specific cluster.
	ClusterId *string
    // An optional list of ESX IDs to remove. format: UUID
	Esxs []string
	NumHosts int64
}

func (s EsxConfig) GetType__() bindings.BindingType {
	return EsxConfigBindingType()
}

func (s EsxConfig) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for EsxConfig._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type EsxHost struct {
	Name *string
    // Availability zone where the host is provisioned.
	AvailabilityZone *string
	EsxId *string
	Hostname *string
	Provider string
    // Backing cloud provider instance type for host.
	InstanceType *string
	MacAddress *string
	CustomProperties map[string]string
    // Possible values are: 
    //
    // * EsxHost#EsxHost_ESX_STATE_DEPLOYING
    // * EsxHost#EsxHost_ESX_STATE_INITIALIZING
    // * EsxHost#EsxHost_ESX_STATE_PROVISIONED
    // * EsxHost#EsxHost_ESX_STATE_READY
    // * EsxHost#EsxHost_ESX_STATE_DELETING
    // * EsxHost#EsxHost_ESX_STATE_DELETED
    // * EsxHost#EsxHost_ESX_STATE_FAILED
    // * EsxHost#EsxHost_ESX_STATE_ADDING_TO_VCENTER
    // * EsxHost#EsxHost_ESX_STATE_DELETING_FROM_VCENTER
    // * EsxHost#EsxHost_ESX_STATE_PENDING_CLOUD_DELETION
	EsxState *string
}
// Identifier denoting this class, when it is used in polymorphic context. 
//
// This value should be assigned to the property which is used to discriminate the actual type used in the polymorphic context.
const EsxHost__TYPE_IDENTIFIER = "EsxHost"
const EsxHost_ESX_STATE_DEPLOYING = "DEPLOYING"
const EsxHost_ESX_STATE_INITIALIZING = "INITIALIZING"
const EsxHost_ESX_STATE_PROVISIONED = "PROVISIONED"
const EsxHost_ESX_STATE_READY = "READY"
const EsxHost_ESX_STATE_DELETING = "DELETING"
const EsxHost_ESX_STATE_DELETED = "DELETED"
const EsxHost_ESX_STATE_FAILED = "FAILED"
const EsxHost_ESX_STATE_ADDING_TO_VCENTER = "ADDING_TO_VCENTER"
const EsxHost_ESX_STATE_DELETING_FROM_VCENTER = "DELETING_FROM_VCENTER"
const EsxHost_ESX_STATE_PENDING_CLOUD_DELETION = "PENDING_CLOUD_DELETION"

func (s EsxHost) GetType__() bindings.BindingType {
	return EsxHostBindingType()
}

func (s EsxHost) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for EsxHost._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type EsxHostInfo struct {
    // Backing cloud provider instance type for cluster.
	InstanceType *string
}

func (s EsxHostInfo) GetType__() bindings.BindingType {
	return EsxHostInfoBindingType()
}

func (s EsxHostInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for EsxHostInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type FirewallRule struct {
    // Possible values are: 
    //
    // * FirewallRule#FirewallRule_RULE_TYPE_USER
    // * FirewallRule#FirewallRule_RULE_TYPE_DEFAULT
	RuleType *string
	ApplicationIds []string
	Name *string
    // Deprecated, left for backwards compatibility. Remove once UI stops using it.
	RuleInterface *string
    // Optional. Possible formats are IP, IP1-IPn, CIDR or comma separated list of those entries. If not specified, defaults to 'any'.
	Destination *string
	Id *string
	DestinationScope *FirewallRuleScope
    // Optional. Possible formats are IP, IP1-IPn, CIDR or comma separated list of those entries. If not specified, defaults to 'any'.
	Source *string
	SourceScope *FirewallRuleScope
    // list of protocols and ports for this firewall rule
	Services []FirewallService
    // Possible values are: 
    //
    // * FirewallRule#FirewallRule_ACTION_ALLOW
    // * FirewallRule#FirewallRule_ACTION_DENY
	Action *string
    // current revision of the list of firewall rules, used to protect against concurrent modification (first writer wins) format: int32
	Revision *int64
}
const FirewallRule_RULE_TYPE_USER = "USER"
const FirewallRule_RULE_TYPE_DEFAULT = "DEFAULT"
const FirewallRule_ACTION_ALLOW = "ALLOW"
const FirewallRule_ACTION_DENY = "DENY"

func (s FirewallRule) GetType__() bindings.BindingType {
	return FirewallRuleBindingType()
}

func (s FirewallRule) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for FirewallRule._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Optional for FirewallRule. If not specified, defaults to 'any'.
type FirewallRuleScope struct {
	GroupingObjectIds []string
    // Possible values are: 
    //
    // * FirewallRuleScope#FirewallRuleScope_VNIC_GROUP_IDS_VSE
    // * FirewallRuleScope#FirewallRuleScope_VNIC_GROUP_IDS_INTERNAL
    // * FirewallRuleScope#FirewallRuleScope_VNIC_GROUP_IDS_EXTERNAL
    // * FirewallRuleScope#FirewallRuleScope_VNIC_GROUP_IDS_VNIC_INDEX_0
    // * FirewallRuleScope#FirewallRuleScope_VNIC_GROUP_IDS_VNIC_INDEX_1
    // * FirewallRuleScope#FirewallRuleScope_VNIC_GROUP_IDS_VNIC_INDEX_2
    // * FirewallRuleScope#FirewallRuleScope_VNIC_GROUP_IDS_VNIC_INDEX_3
    // * FirewallRuleScope#FirewallRuleScope_VNIC_GROUP_IDS_VNIC_INDEX_4
    // * FirewallRuleScope#FirewallRuleScope_VNIC_GROUP_IDS_VNIC_INDEX_5
    // * FirewallRuleScope#FirewallRuleScope_VNIC_GROUP_IDS_VNIC_INDEX_6
    // * FirewallRuleScope#FirewallRuleScope_VNIC_GROUP_IDS_VNIC_INDEX_7
    // * FirewallRuleScope#FirewallRuleScope_VNIC_GROUP_IDS_VNIC_INDEX_8
    // * FirewallRuleScope#FirewallRuleScope_VNIC_GROUP_IDS_VNIC_INDEX_9
	VnicGroupIds []string
}
const FirewallRuleScope_VNIC_GROUP_IDS_VSE = "vse"
const FirewallRuleScope_VNIC_GROUP_IDS_INTERNAL = "internal"
const FirewallRuleScope_VNIC_GROUP_IDS_EXTERNAL = "external"
const FirewallRuleScope_VNIC_GROUP_IDS_VNIC_INDEX_0 = "vnic-index-0"
const FirewallRuleScope_VNIC_GROUP_IDS_VNIC_INDEX_1 = "vnic-index-1"
const FirewallRuleScope_VNIC_GROUP_IDS_VNIC_INDEX_2 = "vnic-index-2"
const FirewallRuleScope_VNIC_GROUP_IDS_VNIC_INDEX_3 = "vnic-index-3"
const FirewallRuleScope_VNIC_GROUP_IDS_VNIC_INDEX_4 = "vnic-index-4"
const FirewallRuleScope_VNIC_GROUP_IDS_VNIC_INDEX_5 = "vnic-index-5"
const FirewallRuleScope_VNIC_GROUP_IDS_VNIC_INDEX_6 = "vnic-index-6"
const FirewallRuleScope_VNIC_GROUP_IDS_VNIC_INDEX_7 = "vnic-index-7"
const FirewallRuleScope_VNIC_GROUP_IDS_VNIC_INDEX_8 = "vnic-index-8"
const FirewallRuleScope_VNIC_GROUP_IDS_VNIC_INDEX_9 = "vnic-index-9"

func (s FirewallRuleScope) GetType__() bindings.BindingType {
	return FirewallRuleScopeBindingType()
}

func (s FirewallRuleScope) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for FirewallRuleScope._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type FirewallService struct {
    // protocol name, such as 'tcp', 'udp' etc.
	Protocol *string
    // a list of port numbers and port ranges, such as {80, 91-95, 99}. If not specified, defaults to 'any'.
	Ports []string
}

func (s FirewallService) GetType__() bindings.BindingType {
	return FirewallServiceBindingType()
}

func (s FirewallService) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for FirewallService._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Describes common properties for MGW and CGW configuration templates
type GatewayTemplate struct {
	PublicIp *SddcPublicIp
	PrimaryDns *string
	SecondaryDns *string
	FirewallRules []FirewallRule
	Vpns []Vpn
}

func (s GatewayTemplate) GetType__() bindings.BindingType {
	return GatewayTemplateBindingType()
}

func (s GatewayTemplate) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for GatewayTemplate._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// the GlcmBundle used for deploying the sddc
type GlcmBundle struct {
    // the glcmbundle's s3 bucket
	S3Bucket *string
    // the glcmbundle's id
	Id *string
}

func (s GlcmBundle) GetType__() bindings.BindingType {
	return GlcmBundleBindingType()
}

func (s GlcmBundle) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for GlcmBundle._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Represents a structure for instance type config
type InstanceTypeConfig struct {
    // Instance type name.
	InstanceType *string
    // Array of number of hosts allowed for this operation. Range of hosts user can select for sddc provision
	Hosts []int64
    // Display name of instance_type.
	DisplayName *string
    // The capacity of the given instance type.
	EntityCapacity *EntityCapacity
}

func (s InstanceTypeConfig) GetType__() bindings.BindingType {
	return InstanceTypeConfigBindingType()
}

func (s InstanceTypeConfig) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for InstanceTypeConfig._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type KmsVpcEndpoint struct {
    // The identifier of the VPC endpoint created to AWS Key Management Service
	VpcEndpointId *string
	NetworkInterfaceIds []string
}

func (s KmsVpcEndpoint) GetType__() bindings.BindingType {
	return KmsVpcEndpointBindingType()
}

func (s KmsVpcEndpoint) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for KmsVpcEndpoint._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type L2Vpn struct {
    // Enable (true) or disable (false) L2 VPN.
	Enabled *bool
    // Array of L2 vpn site config.
	Sites []Site
    // Public uplink ip address. IP of external interface on which L2VPN service listens to.
	ListenerIp *string
}

func (s L2Vpn) GetType__() bindings.BindingType {
	return L2VpnBindingType()
}

func (s L2Vpn) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for L2Vpn._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type LogicalNetwork struct {
    // the subnet cidr
	SubnetCidr *string
    // name of the network
	Name *string
    // gateway ip of the logical network
	GatewayIp *string
    // if 'true' - enabled; if 'false' - disabled
	DhcpEnabled *string
    // ip range within the subnet mask, range delimiter is '-' (example 10.118.10.130-10.118.10.140)
	DhcpIpRange *string
    // tunnel id of extended network format: int32
	TunnelId *int64
	Id *string
    // Possible values are: 
    //
    // * LogicalNetwork#LogicalNetwork_NETWORK_TYPE_HOSTED
    // * LogicalNetwork#LogicalNetwork_NETWORK_TYPE_ROUTED
    // * LogicalNetwork#LogicalNetwork_NETWORK_TYPE_EXTENDED
	NetworkType *string
}
const LogicalNetwork_NETWORK_TYPE_HOSTED = "HOSTED"
const LogicalNetwork_NETWORK_TYPE_ROUTED = "ROUTED"
const LogicalNetwork_NETWORK_TYPE_EXTENDED = "EXTENDED"

func (s LogicalNetwork) GetType__() bindings.BindingType {
	return LogicalNetworkBindingType()
}

func (s LogicalNetwork) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for LogicalNetwork._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type MaintenanceWindow struct {
    // Possible values are: 
    //
    // * MaintenanceWindow#MaintenanceWindow_DAY_OF_WEEK_SUNDAY
    // * MaintenanceWindow#MaintenanceWindow_DAY_OF_WEEK_MONDAY
    // * MaintenanceWindow#MaintenanceWindow_DAY_OF_WEEK_TUESDAY
    // * MaintenanceWindow#MaintenanceWindow_DAY_OF_WEEK_WEDNESDAY
    // * MaintenanceWindow#MaintenanceWindow_DAY_OF_WEEK_THURSDAY
    // * MaintenanceWindow#MaintenanceWindow_DAY_OF_WEEK_FRIDAY
    // * MaintenanceWindow#MaintenanceWindow_DAY_OF_WEEK_SATURDAY
	DayOfWeek *string
	HourOfDay *int64
}
const MaintenanceWindow_DAY_OF_WEEK_SUNDAY = "SUNDAY"
const MaintenanceWindow_DAY_OF_WEEK_MONDAY = "MONDAY"
const MaintenanceWindow_DAY_OF_WEEK_TUESDAY = "TUESDAY"
const MaintenanceWindow_DAY_OF_WEEK_WEDNESDAY = "WEDNESDAY"
const MaintenanceWindow_DAY_OF_WEEK_THURSDAY = "THURSDAY"
const MaintenanceWindow_DAY_OF_WEEK_FRIDAY = "FRIDAY"
const MaintenanceWindow_DAY_OF_WEEK_SATURDAY = "SATURDAY"

func (s MaintenanceWindow) GetType__() bindings.BindingType {
	return MaintenanceWindowBindingType()
}

func (s MaintenanceWindow) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for MaintenanceWindow._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type MaintenanceWindowEntry struct {
    // true if the SDDC is in the defined Mainentance Window
	InMaintenanceWindow *bool
	ReservationSchedule *ReservationSchedule
    // ID for reservation format: uuid
	ReservationId *string
    // true if the SDDC is currently undergoing maintenance
	InMaintenanceMode *bool
    // SDDC ID for this reservation format: uuid
	SddcId *string
}

func (s MaintenanceWindowEntry) GetType__() bindings.BindingType {
	return MaintenanceWindowEntryBindingType()
}

func (s MaintenanceWindowEntry) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for MaintenanceWindowEntry._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type MaintenanceWindowGet struct {
    // Possible values are: 
    //
    // * MaintenanceWindowGet#MaintenanceWindowGet_DAY_OF_WEEK_SUNDAY
    // * MaintenanceWindowGet#MaintenanceWindowGet_DAY_OF_WEEK_MONDAY
    // * MaintenanceWindowGet#MaintenanceWindowGet_DAY_OF_WEEK_TUESDAY
    // * MaintenanceWindowGet#MaintenanceWindowGet_DAY_OF_WEEK_WEDNESDAY
    // * MaintenanceWindowGet#MaintenanceWindowGet_DAY_OF_WEEK_THURSDAY
    // * MaintenanceWindowGet#MaintenanceWindowGet_DAY_OF_WEEK_FRIDAY
    // * MaintenanceWindowGet#MaintenanceWindowGet_DAY_OF_WEEK_SATURDAY
	DayOfWeek *string
	HourOfDay *int64
	DurationMin *int64
	Version *int64
}
const MaintenanceWindowGet_DAY_OF_WEEK_SUNDAY = "SUNDAY"
const MaintenanceWindowGet_DAY_OF_WEEK_MONDAY = "MONDAY"
const MaintenanceWindowGet_DAY_OF_WEEK_TUESDAY = "TUESDAY"
const MaintenanceWindowGet_DAY_OF_WEEK_WEDNESDAY = "WEDNESDAY"
const MaintenanceWindowGet_DAY_OF_WEEK_THURSDAY = "THURSDAY"
const MaintenanceWindowGet_DAY_OF_WEEK_FRIDAY = "FRIDAY"
const MaintenanceWindowGet_DAY_OF_WEEK_SATURDAY = "SATURDAY"

func (s MaintenanceWindowGet) GetType__() bindings.BindingType {
	return MaintenanceWindowGetBindingType()
}

func (s MaintenanceWindowGet) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for MaintenanceWindowGet._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type ManagementGatewayTemplate struct {
	PublicIp *SddcPublicIp
	PrimaryDns *string
	SecondaryDns *string
	FirewallRules []FirewallRule
	Vpns []Vpn
    // mgw network subnet cidr
	SubnetCidr *string
}

func (s ManagementGatewayTemplate) GetType__() bindings.BindingType {
	return ManagementGatewayTemplateBindingType()
}

func (s ManagementGatewayTemplate) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ManagementGatewayTemplate._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type MapZonesRequest struct {
    // The connected account ID to remap. This is a standard UUID.
	ConnectedAccountId *string
    // The org ID to remap in. This is a standard UUID.
	OrgId *string
    // A list of Petronas regions to map.
	PetronasRegionsToMap []string
}

func (s MapZonesRequest) GetType__() bindings.BindingType {
	return MapZonesRequestBindingType()
}

func (s MapZonesRequest) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for MapZonesRequest._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// metadata of the sddc manifest
type Metadata struct {
    // the timestamp for the bundle
	Timestamp *string
    // the cycle id
	CycleId *string
}

func (s Metadata) GetType__() bindings.BindingType {
	return MetadataBindingType()
}

func (s Metadata) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for Metadata._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type MsftLicensingConfig struct {
    // Possible values are: 
    //
    // * MsftLicensingConfig#MsftLicensingConfig_MSSQL_LICENSING_DISABLED
    // * MsftLicensingConfig#MsftLicensingConfig_MSSQL_LICENSING_CUSTOMER_SUPPLIED
    // * MsftLicensingConfig#MsftLicensingConfig_MSSQL_LICENSING_ENABLED
    //
    //  The status MSSQL licensing for this SDDC's clusters.
	MssqlLicensing *string
    // Possible values are: 
    //
    // * MsftLicensingConfig#MsftLicensingConfig_WINDOWS_LICENSING_DISABLED
    // * MsftLicensingConfig#MsftLicensingConfig_WINDOWS_LICENSING_CUSTOMER_SUPPLIED
    // * MsftLicensingConfig#MsftLicensingConfig_WINDOWS_LICENSING_ENABLED
    //
    //  The status of Windows licensing for this SDDC's clusters. Can be enabled, disabled, or customer's.
	WindowsLicensing *string
}
const MsftLicensingConfig_MSSQL_LICENSING_DISABLED = "DISABLED"
const MsftLicensingConfig_MSSQL_LICENSING_CUSTOMER_SUPPLIED = "CUSTOMER_SUPPLIED"
const MsftLicensingConfig_MSSQL_LICENSING_ENABLED = "ENABLED"
const MsftLicensingConfig_WINDOWS_LICENSING_DISABLED = "DISABLED"
const MsftLicensingConfig_WINDOWS_LICENSING_CUSTOMER_SUPPLIED = "CUSTOMER_SUPPLIED"
const MsftLicensingConfig_WINDOWS_LICENSING_ENABLED = "ENABLED"

func (s MsftLicensingConfig) GetType__() bindings.BindingType {
	return MsftLicensingConfigBindingType()
}

func (s MsftLicensingConfig) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for MsftLicensingConfig._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type NatRule struct {
	RuleType *string
	Protocol *string
	Name *string
	InternalPorts *string
	PublicPorts *string
	PublicIp *string
	InternalIp *string
    // Possible values are: 
    //
    // * NatRule#NatRule_ACTION_DNAT
    // * NatRule#NatRule_ACTION_SNAT
	Action *string
	Id *string
    // current revision of the list of nat rules, used to protect against concurrent modification (first writer wins) format: int32
	Revision *int64
}
const NatRule_ACTION_DNAT = "dnat"
const NatRule_ACTION_SNAT = "snat"

func (s NatRule) GetType__() bindings.BindingType {
	return NatRuleBindingType()
}

func (s NatRule) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for NatRule._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type NetworkTemplate struct {
	ManagementGatewayTemplates []ManagementGatewayTemplate
	ComputeGatewayTemplates []ComputeGatewayTemplate
}

func (s NetworkTemplate) GetType__() bindings.BindingType {
	return NetworkTemplateBindingType()
}

func (s NetworkTemplate) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for NetworkTemplate._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type NewCredentials struct {
    // Username of the credentials
	Username string
    // Password associated with the credentials
	Password string
    // Name of the credentials
	Name string
}

func (s NewCredentials) GetType__() bindings.BindingType {
	return NewCredentialsBindingType()
}

func (s NewCredentials) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for NewCredentials._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Details the state of different NSX add-ons.
type NsxtAddons struct {
    // Indicates whether NSX Advanced addon is enabled or disabled.
	EnableNsxAdvancedAddon *bool
}

func (s NsxtAddons) GetType__() bindings.BindingType {
	return NsxtAddonsBindingType()
}

func (s NsxtAddons) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for NsxtAddons._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Holder for the offer instances.
type OfferInstancesHolder struct {
	OnDemand OnDemandOfferInstance
	Offers []TermOfferInstance
}

func (s OfferInstancesHolder) GetType__() bindings.BindingType {
	return OfferInstancesHolderBindingType()
}

func (s OfferInstancesHolder) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for OfferInstancesHolder._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Holder for the on-demand offer instance.
type OnDemandOfferInstance struct {
	Product string
    // Deprecated. Please use product and type fields instead.
	ProductType *string
	Name string
	Currency string
	Region string
	UnitPrice string
	MonthlyCost string
	Version string
	Type_ string
	Description string
}

func (s OnDemandOfferInstance) GetType__() bindings.BindingType {
	return OnDemandOfferInstanceBindingType()
}

func (s OnDemandOfferInstance) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for OnDemandOfferInstance._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type OrgProperties struct {
    // A map of string properties to values.
	Values map[string]string
}

func (s OrgProperties) GetType__() bindings.BindingType {
	return OrgPropertiesBindingType()
}

func (s OrgProperties) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for OrgProperties._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type OrgSellerInfo struct {
    // The accountid for this org for the seller-of-record. NILLABLE.
	SellerAccountId *string
    // The seller-of-record for the current organization. For example AWS or VMWARE
	Seller *string
}

func (s OrgSellerInfo) GetType__() bindings.BindingType {
	return OrgSellerInfoBindingType()
}

func (s OrgSellerInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for OrgSellerInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type Organization struct {
	Updated time.Time
    // User id that last updated this record
	UserId string
    // User id that last updated this record
	UpdatedByUserId string
	Created time.Time
    // Version of this entity format: int32
	Version int64
    // User name that last updated this record
	UpdatedByUserName *string
    // User name that last updated this record
	UserName string
    // Unique ID for this entity
	Id string
    // ORG_TYPE to be associated with the org
	OrgType *string
	DisplayName *string
	Name *string
	OrgSellerInfo *OrgSellerInfo
    // Possible values are: 
    //
    // * Organization#Organization_PROJECT_STATE_CREATED
    // * Organization#Organization_PROJECT_STATE_DELETED
	ProjectState *string
	Properties *OrgProperties
}
const Organization_PROJECT_STATE_CREATED = "CREATED"
const Organization_PROJECT_STATE_DELETED = "DELETED"

func (s Organization) GetType__() bindings.BindingType {
	return OrganizationBindingType()
}

func (s Organization) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for Organization._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type PaymentMethodInfo struct {
	Type_ *string
	DefaultFlag *bool
	PaymentMethodId *string
}

func (s PaymentMethodInfo) GetType__() bindings.BindingType {
	return PaymentMethodInfoBindingType()
}

func (s PaymentMethodInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for PaymentMethodInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type PopAgentXeniConnection struct {
    // The gateway route ip fo the subnet.
	DefaultSubnetRoute *string
	EniInfo *EniInfo
}

func (s PopAgentXeniConnection) GetType__() bindings.BindingType {
	return PopAgentXeniConnectionBindingType()
}

func (s PopAgentXeniConnection) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for PopAgentXeniConnection._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type PopAmiInfo struct {
    // instance type of the esx ami
	InstanceType *string
    // the region of the esx ami
	Region *string
    // the ami id for the esx
	Id *string
    // the name of the esx ami
	Name *string
    // Possible values are: 
    //
    // * PopAmiInfo#PopAmiInfo_TYPE_CENTOS
    // * PopAmiInfo#PopAmiInfo_TYPE_POP
    //
    //  PoP AMI type. CENTOS: a Centos AMI; POP: a PoP AMI.
	Type_ *string
}
const PopAmiInfo_TYPE_CENTOS = "CENTOS"
const PopAmiInfo_TYPE_POP = "POP"

func (s PopAmiInfo) GetType__() bindings.BindingType {
	return PopAmiInfoBindingType()
}

func (s PopAmiInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for PopAmiInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Present a SDDC PoP information.
type PopInfo struct {
    // A map of [region name of PoP / PoP-AMI]:[PopAmiInfo].
	AmiInfos map[string]PopAmiInfo
    // The PopInfo (or PoP AMI) created time. Using ISO 8601 date-time pattern. format: date-time
	CreatedAt *time.Time
    // A map of [service type]:[PopServiceInfo]
	ServiceInfos map[string]PopServiceInfo
    // UUID of the PopInfo format: UUID
	Id *string
    // version of the manifest.
	ManifestVersion *string
}

func (s PopInfo) GetType__() bindings.BindingType {
	return PopInfoBindingType()
}

func (s PopInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for PopInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type PopServiceInfo struct {
    // The service change set number.
	Cln *string
    // The service API version.
	Version *string
    // The service build number.
	Build *string
    // Possible values are: 
    //
    // * PopServiceInfo#PopServiceInfo_SERVICE_OS
    // * PopServiceInfo#PopServiceInfo_SERVICE_AGENT
    // * PopServiceInfo#PopServiceInfo_SERVICE_GLCM
    // * PopServiceInfo#PopServiceInfo_SERVICE_S3_ADAPTER
    // * PopServiceInfo#PopServiceInfo_SERVICE_JRE
    // * PopServiceInfo#PopServiceInfo_SERVICE_DOCKER
    // * PopServiceInfo#PopServiceInfo_SERVICE_AIDE
    // * PopServiceInfo#PopServiceInfo_SERVICE_RTS
    // * PopServiceInfo#PopServiceInfo_SERVICE_FM_LOG_COLLECTOR
    // * PopServiceInfo#PopServiceInfo_SERVICE_FM_METRICS_COLLECTOR
    // * PopServiceInfo#PopServiceInfo_SERVICE_BRE
    // * PopServiceInfo#PopServiceInfo_SERVICE_BRF
    // * PopServiceInfo#PopServiceInfo_SERVICE_REVERSE_PROXY
    // * PopServiceInfo#PopServiceInfo_SERVICE_FORWARD_PROXY
    // * PopServiceInfo#PopServiceInfo_SERVICE_DNS
    // * PopServiceInfo#PopServiceInfo_SERVICE_NTP
    // * PopServiceInfo#PopServiceInfo_SERVICE_LOGZ_LOG_COLLECTOR
    //
    //  An enum of PoP related services (including os platform and JRE).
	Service *string
}
const PopServiceInfo_SERVICE_OS = "OS"
const PopServiceInfo_SERVICE_AGENT = "AGENT"
const PopServiceInfo_SERVICE_GLCM = "GLCM"
const PopServiceInfo_SERVICE_S3_ADAPTER = "S3_ADAPTER"
const PopServiceInfo_SERVICE_JRE = "JRE"
const PopServiceInfo_SERVICE_DOCKER = "DOCKER"
const PopServiceInfo_SERVICE_AIDE = "AIDE"
const PopServiceInfo_SERVICE_RTS = "RTS"
const PopServiceInfo_SERVICE_FM_LOG_COLLECTOR = "FM_LOG_COLLECTOR"
const PopServiceInfo_SERVICE_FM_METRICS_COLLECTOR = "FM_METRICS_COLLECTOR"
const PopServiceInfo_SERVICE_BRE = "BRE"
const PopServiceInfo_SERVICE_BRF = "BRF"
const PopServiceInfo_SERVICE_REVERSE_PROXY = "REVERSE_PROXY"
const PopServiceInfo_SERVICE_FORWARD_PROXY = "FORWARD_PROXY"
const PopServiceInfo_SERVICE_DNS = "DNS"
const PopServiceInfo_SERVICE_NTP = "NTP"
const PopServiceInfo_SERVICE_LOGZ_LOG_COLLECTOR = "LOGZ_LOG_COLLECTOR"

func (s PopServiceInfo) GetType__() bindings.BindingType {
	return PopServiceInfoBindingType()
}

func (s PopServiceInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for PopServiceInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Represents a provisioning spec for a sddc
type ProvisionSpec struct {
    // Map of provider to sddc config spec
	Provider map[string]SddcConfigSpec
}

func (s ProvisionSpec) GetType__() bindings.BindingType {
	return ProvisionSpecBindingType()
}

func (s ProvisionSpec) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ProvisionSpec._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type RequestDetail struct {
    // AWS quota increase request id
	AwsQuotaRequestId *string
    // Possible values are: 
    //
    // * RequestDetail#RequestDetail_DETAIL_STATUS_NEW
    // * RequestDetail#RequestDetail_DETAIL_STATUS_PENDINGSUBMIT
    // * RequestDetail#RequestDetail_DETAIL_STATUS_SUBMITTED
    // * RequestDetail#RequestDetail_DETAIL_STATUS_PENDINGVERFICATION
    // * RequestDetail#RequestDetail_DETAIL_STATUS_RESOLVED
    // * RequestDetail#RequestDetail_DETAIL_STATUS_DENIED
    // * RequestDetail#RequestDetail_DETAIL_STATUS_ERROR
	DetailStatus *string
	ResolvedAt *time.Time
    // desired value for the quota increase request
	DesiredValue *int64
    // AWS support case status
	AwsSupportCaseStatus *string
    // AWS support caes id
	AwsSupportCaseId *string
}
const RequestDetail_DETAIL_STATUS_NEW = "NEW"
const RequestDetail_DETAIL_STATUS_PENDINGSUBMIT = "PENDINGSUBMIT"
const RequestDetail_DETAIL_STATUS_SUBMITTED = "SUBMITTED"
const RequestDetail_DETAIL_STATUS_PENDINGVERFICATION = "PENDINGVERFICATION"
const RequestDetail_DETAIL_STATUS_RESOLVED = "RESOLVED"
const RequestDetail_DETAIL_STATUS_DENIED = "DENIED"
const RequestDetail_DETAIL_STATUS_ERROR = "ERROR"

func (s RequestDetail) GetType__() bindings.BindingType {
	return RequestDetailBindingType()
}

func (s RequestDetail) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for RequestDetail._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type Reservation struct {
    // Duration - required for reservation in maintenance window format: int64
	Duration *int64
    // Reservation ID format: uuid
	Rid *string
    // Optional
	CreateTime *string
    // Start time of a reservation format: date-time
	StartTime *time.Time
    // Optional
	Metadata map[string]string
}

func (s Reservation) GetType__() bindings.BindingType {
	return ReservationBindingType()
}

func (s Reservation) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for Reservation._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type ReservationInMw struct {
    // Reservation ID format: uuid
	Rid *string
    // SUNDAY of the week that maintenance is scheduled, ISO format date
	WeekOf *string
    // Optional format: date-time
	CreateTime *time.Time
    // Optional
	Metadata map[string]string
}

func (s ReservationInMw) GetType__() bindings.BindingType {
	return ReservationInMwBindingType()
}

func (s ReservationInMw) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ReservationInMw._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type ReservationSchedule struct {
    // Possible values are: 
    //
    // * ReservationSchedule#ReservationSchedule_DAY_OF_WEEK_SUNDAY
    // * ReservationSchedule#ReservationSchedule_DAY_OF_WEEK_MONDAY
    // * ReservationSchedule#ReservationSchedule_DAY_OF_WEEK_TUESDAY
    // * ReservationSchedule#ReservationSchedule_DAY_OF_WEEK_WEDNESDAY
    // * ReservationSchedule#ReservationSchedule_DAY_OF_WEEK_THURSDAY
    // * ReservationSchedule#ReservationSchedule_DAY_OF_WEEK_FRIDAY
    // * ReservationSchedule#ReservationSchedule_DAY_OF_WEEK_SATURDAY
	DayOfWeek *string
	HourOfDay *int64
	DurationMin *int64
	Version *int64
	Reservations []Reservation
	ReservationsMw []ReservationInMw
}
const ReservationSchedule_DAY_OF_WEEK_SUNDAY = "SUNDAY"
const ReservationSchedule_DAY_OF_WEEK_MONDAY = "MONDAY"
const ReservationSchedule_DAY_OF_WEEK_TUESDAY = "TUESDAY"
const ReservationSchedule_DAY_OF_WEEK_WEDNESDAY = "WEDNESDAY"
const ReservationSchedule_DAY_OF_WEEK_THURSDAY = "THURSDAY"
const ReservationSchedule_DAY_OF_WEEK_FRIDAY = "FRIDAY"
const ReservationSchedule_DAY_OF_WEEK_SATURDAY = "SATURDAY"

func (s ReservationSchedule) GetType__() bindings.BindingType {
	return ReservationScheduleBindingType()
}

func (s ReservationSchedule) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ReservationSchedule._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type ReservationWindow struct {
    // Possible values are: 
    //
    // * ReservationWindow#ReservationWindow_RESERVATION_STATE_SCHEDULED
    // * ReservationWindow#ReservationWindow_RESERVATION_STATE_RUNNING
    // * ReservationWindow#ReservationWindow_RESERVATION_STATE_CANCELED
    // * ReservationWindow#ReservationWindow_RESERVATION_STATE_COMPLETED
	ReservationState *string
	Emergency *bool
	MaintenanceProperties *ReservationWindowMaintenanceProperties
	ReserveId *string
	StartHour *int64
	SddcId *string
	ManifestId *string
	DurationHours *int64
	StartDate *string
    // Metadata for reservation window, in key-value form
	Metadata map[string]string
}
const ReservationWindow_RESERVATION_STATE_SCHEDULED = "SCHEDULED"
const ReservationWindow_RESERVATION_STATE_RUNNING = "RUNNING"
const ReservationWindow_RESERVATION_STATE_CANCELED = "CANCELED"
const ReservationWindow_RESERVATION_STATE_COMPLETED = "COMPLETED"

func (s ReservationWindow) GetType__() bindings.BindingType {
	return ReservationWindowBindingType()
}

func (s ReservationWindow) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ReservationWindow._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type ReservationWindowMaintenanceProperties struct {
    // Status of upgrade, if any
	Status *string
}

func (s ReservationWindowMaintenanceProperties) GetType__() bindings.BindingType {
	return ReservationWindowMaintenancePropertiesBindingType()
}

func (s ReservationWindowMaintenanceProperties) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ReservationWindowMaintenanceProperties._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type RouteTableInfo struct {
    // route table name
	Name *string
    // route table id
	Id *string
}

func (s RouteTableInfo) GetType__() bindings.BindingType {
	return RouteTableInfoBindingType()
}

func (s RouteTableInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for RouteTableInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type Sddc struct {
	Updated time.Time
    // User id that last updated this record
	UserId string
    // User id that last updated this record
	UpdatedByUserId string
	Created time.Time
    // Version of this entity format: int32
	Version int64
    // User name that last updated this record
	UpdatedByUserName *string
    // User name that last updated this record
	UserName string
    // Unique ID for this entity
	Id string
    // name for SDDC
	Name *string
    // Possible values are: 
    //
    // * Sddc#Sddc_SDDC_STATE_DEPLOYING
    // * Sddc#Sddc_SDDC_STATE_READY
    // * Sddc#Sddc_SDDC_STATE_DELETING
    // * Sddc#Sddc_SDDC_STATE_DELETED
    // * Sddc#Sddc_SDDC_STATE_FAILED
    // * Sddc#Sddc_SDDC_STATE_CANCELED
    // * Sddc#Sddc_SDDC_STATE_READY_FOR_GLCM_BRINGUP
	SddcState *string
    // Expiration date of a sddc in UTC (will be set if its applicable) format: date-time
	ExpirationDate *time.Time
	OrgId *string
    // Type of the sddc
	SddcType *string
    // Possible values are: 
    //
    // * Sddc#Sddc_PROVIDER_AWS
    // * Sddc#Sddc_PROVIDER_ZEROCLOUD
	Provider *string
    // Possible values are: 
    //
    // * Sddc#Sddc_ACCOUNT_LINK_STATE_DELAYED
    // * Sddc#Sddc_ACCOUNT_LINK_STATE_LINKED
    // * Sddc#Sddc_ACCOUNT_LINK_STATE_UNLINKED
    //
    //  Account linking state of the sddc
	AccountLinkState *string
    // Describes the access state of sddc, valid state is DISABLED or ENABLED
	SddcAccessState *string
	ResourceConfig *AwsSddcResourceConfig
}
const Sddc_SDDC_STATE_DEPLOYING = "DEPLOYING"
const Sddc_SDDC_STATE_READY = "READY"
const Sddc_SDDC_STATE_DELETING = "DELETING"
const Sddc_SDDC_STATE_DELETED = "DELETED"
const Sddc_SDDC_STATE_FAILED = "FAILED"
const Sddc_SDDC_STATE_CANCELED = "CANCELED"
const Sddc_SDDC_STATE_READY_FOR_GLCM_BRINGUP = "READY_FOR_GLCM_BRINGUP"
const Sddc_PROVIDER_AWS = "AWS"
const Sddc_PROVIDER_ZEROCLOUD = "ZEROCLOUD"
const Sddc_ACCOUNT_LINK_STATE_DELAYED = "DELAYED"
const Sddc_ACCOUNT_LINK_STATE_LINKED = "LINKED"
const Sddc_ACCOUNT_LINK_STATE_UNLINKED = "UNLINKED"

func (s Sddc) GetType__() bindings.BindingType {
	return SddcBindingType()
}

func (s Sddc) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for Sddc._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type SddcConfig struct {
    // Indicates the desired licensing support, if any, of Microsoft software.
	MsftLicenseConfig *MsftLicensingConfig
    // AWS VPC IP range. Only prefix of 16 or 20 is currently supported.
	VpcCidr *string
    // The instance type for the esx hosts in the primary cluster of the SDDC.
	HostInstanceType *HostInstanceTypes
    // skip creating vxlan for compute gateway for SDDC provisioning
	SkipCreatingVxlan *bool
    // VXLAN IP subnet in CIDR for compute gateway
	VxlanSubnet *string
    // Possible values are: 
    //
    // * SddcConfig#SddcConfig_SIZE_NSX_SMALL
    // * SddcConfig#SddcConfig_SIZE_MEDIUM
    // * SddcConfig#SddcConfig_SIZE_LARGE
    // * SddcConfig#SddcConfig_SIZE_NSX_LARGE
    //
    //  The size of the vCenter and NSX appliances. \"large\" sddcSize corresponds to a 'large' vCenter appliance and 'large' NSX appliance. 'medium' sddcSize corresponds to 'medium' vCenter appliance and 'medium' NSX appliance. Value defaults to 'medium'.
	Size *string
    // The storage capacity value to be requested for the sddc primary cluster, in GiBs. If provided, instead of using the direct-attached storage, a capacity value amount of seperable storage will be used. format: int64
	StorageCapacity *int64
	Name string
    // A list of the SDDC linking configurations to use.
	AccountLinkSddcConfig []AccountLinkSddcConfig
    // If provided, will be assigned as SDDC id of the provisioned SDDC. format: UUID
	SddcId *string
	NumHosts int64
    // Denotes the sddc type , if the value is null or empty, the type is considered as default.
	SddcType *string
    // The account linking configuration, we will keep this one and remove accountLinkSddcConfig finally.
	AccountLinkConfig *AccountLinkConfig
    // Possible values are: 
    //
    // * SddcConfig#SddcConfig_PROVIDER_AWS
    // * SddcConfig#SddcConfig_PROVIDER_ZEROCLOUD
    //
    //  Determines what additional properties are available based on cloud provider.
	Provider string
    // The SSO domain name to use for vSphere users. If not specified, vmc.local will be used.
	SsoDomain *string
    // If provided, configuration from the template will applied to the provisioned SDDC. format: UUID
	SddcTemplateId *string
    // Possible values are: 
    //
    // * SddcConfig#SddcConfig_DEPLOYMENT_TYPE_SINGLEAZ
    // * SddcConfig#SddcConfig_DEPLOYMENT_TYPE_MULTIAZ
    //
    //  Denotes if request is for a SingleAZ or a MultiAZ SDDC. Default is SingleAZ.
	DeploymentType *string
}
// Identifier denoting this class, when it is used in polymorphic context. 
//
// This value should be assigned to the property which is used to discriminate the actual type used in the polymorphic context.
const SddcConfig__TYPE_IDENTIFIER = "SddcConfig"
const SddcConfig_SIZE_NSX_SMALL = "nsx_small"
const SddcConfig_SIZE_MEDIUM = "medium"
const SddcConfig_SIZE_LARGE = "large"
const SddcConfig_SIZE_NSX_LARGE = "nsx_large"
const SddcConfig_PROVIDER_AWS = "AWS"
const SddcConfig_PROVIDER_ZEROCLOUD = "ZEROCLOUD"
const SddcConfig_DEPLOYMENT_TYPE_SINGLEAZ = "SingleAZ"
const SddcConfig_DEPLOYMENT_TYPE_MULTIAZ = "MultiAZ"

func (s SddcConfig) GetType__() bindings.BindingType {
	return SddcConfigBindingType()
}

func (s SddcConfig) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SddcConfig._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Represents a configuration spec for a sddc
type SddcConfigSpec struct {
    // Map of sddc type to config spec
	SddcTypeConfigSpec map[string]ConfigSpec
    // The region name to display names mapping
	RegionDisplayNames map[string]string
}

func (s SddcConfigSpec) GetType__() bindings.BindingType {
	return SddcConfigSpecBindingType()
}

func (s SddcConfigSpec) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SddcConfigSpec._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type SddcId struct {
    // Sddc ID
	SddcId *string
}

func (s SddcId) GetType__() bindings.BindingType {
	return SddcIdBindingType()
}

func (s SddcId) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SddcId._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type SddcLinkConfig struct {
	CustomerSubnetIds []string
    // Determines which connected customer account to link to
	ConnectedAccountId *string
}

func (s SddcLinkConfig) GetType__() bindings.BindingType {
	return SddcLinkConfigBindingType()
}

func (s SddcLinkConfig) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SddcLinkConfig._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Describes software components of the installed SDDC
type SddcManifest struct {
    // the vmcVersion of the sddc for display
	VmcVersion *string
	GlcmBundle *GlcmBundle
	PopInfo *PopInfo
    // the vmcInternalVersion of the sddc for internal use
	VmcInternalVersion *string
	EbsBackedVsanConfig *EbsBackedVsanConfig
	VsanWitnessAmi *AmiInfo
	EsxAmi *AmiInfo
	EsxNsxtAmi *AmiInfo
	Metadata *Metadata
}

func (s SddcManifest) GetType__() bindings.BindingType {
	return SddcManifestBindingType()
}

func (s SddcManifest) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SddcManifest._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Patch request body for SDDC
type SddcPatchRequest struct {
    // The new name of the SDDC to be changed to.
	Name *string
}

func (s SddcPatchRequest) GetType__() bindings.BindingType {
	return SddcPatchRequestBindingType()
}

func (s SddcPatchRequest) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SddcPatchRequest._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type SddcPublicIp struct {
	PublicIp string
	Name *string
	AllocationId *string
	DnatRuleId *string
	AssociatedPrivateIp *string
	SnatRuleId *string
}

func (s SddcPublicIp) GetType__() bindings.BindingType {
	return SddcPublicIpBindingType()
}

func (s SddcPublicIp) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SddcPublicIp._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type SddcResourceConfig struct {
    // Name for management appliance network.
	MgmtApplianceNetworkName *string
    // if true, NSX-T UI is enabled.
	Nsxt *bool
    // Management Gateway Id
	MgwId *string
    // URL of the NSX Manager
	NsxMgrUrl *string
    // PSC internal management IP
	PscManagementIp *string
    // URL of the PSC server
	PscUrl *string
	Cgws []string
    // Availability zones over which esx hosts are provisioned. MultiAZ SDDCs will have hosts provisioned over two availability zones while SingleAZ SDDCs will provision over one.
	AvailabilityZones []string
    // The ManagedObjectReference of the management Datastore
	ManagementDs *string
    // nsx api entire base url
	NsxApiPublicEndpointUrl *string
	CustomProperties map[string]string
    // Password for vCenter SDDC administrator
	CloudPassword *string
    // Possible values are: 
    //
    // * SddcResourceConfig#SddcResourceConfig_PROVIDER_AWS
    // * SddcResourceConfig#SddcResourceConfig_PROVIDER_ZEROCLOUD
    //
    //  Discriminator for additional properties
	Provider string
    // List of clusters in the SDDC.
	Clusters []Cluster
    // vCenter internal management IP
	VcManagementIp *string
	SddcNetworks []string
    // Username for vCenter SDDC administrator
	CloudUsername *string
	EsxHosts []AwsEsxHost
    // NSX Manager internal management IP
	NsxMgrManagementIp *string
    // unique id of the vCenter server
	VcInstanceId *string
    // Cluster Id to add ESX workflow
	EsxClusterId *string
    // vCenter public IP
	VcPublicIp *string
    // skip creating vxlan for compute gateway for SDDC provisioning
	SkipCreatingVxlan *bool
    // URL of the vCenter server
	VcUrl *string
	SddcManifest *SddcManifest
    // VXLAN IP subnet
	VxlanSubnet *string
    // Group name for vCenter SDDC administrator
	CloudUserGroup *string
	ManagementRp *string
    // region in which sddc is deployed
	Region *string
    // Availability zone where the witness node is provisioned for a MultiAZ SDDC. This is null for a SingleAZ SDDC.
	WitnessAvailabilityZone *string
    // sddc identifier
	SddcId *string
	PopAgentXeniConnection *PopAgentXeniConnection
	SddcSize *SddcSize
    // List of Controller IPs
	NsxControllerIps []string
    // ESX host subnet
	EsxHostSubnet *string
    // The SSO domain name to use for vSphere users
	SsoDomain *string
    // Possible values are: 
    //
    // * SddcResourceConfig#SddcResourceConfig_DEPLOYMENT_TYPE_SINGLE_AZ
    // * SddcResourceConfig#SddcResourceConfig_DEPLOYMENT_TYPE_MULTI_AZ
    //
    //  Denotes if this is a SingleAZ SDDC or a MultiAZ SDDC.
	DeploymentType *string
    // The Microsoft license status of this SDDC.
	MsftLicenseConfig *MsftLicensingConfig
	NsxtAddons *NsxtAddons
    // if true, use the private IP addresses to register DNS records for the management VMs
	DnsWithManagementVmPrivateIp *bool
}
// Identifier denoting this class, when it is used in polymorphic context. 
//
// This value should be assigned to the property which is used to discriminate the actual type used in the polymorphic context.
const SddcResourceConfig__TYPE_IDENTIFIER = "SddcResourceConfig"
const SddcResourceConfig_PROVIDER_AWS = "AWS"
const SddcResourceConfig_PROVIDER_ZEROCLOUD = "ZEROCLOUD"
const SddcResourceConfig_DEPLOYMENT_TYPE_SINGLE_AZ = "SINGLE_AZ"
const SddcResourceConfig_DEPLOYMENT_TYPE_MULTI_AZ = "MULTI_AZ"

func (s SddcResourceConfig) GetType__() bindings.BindingType {
	return SddcResourceConfigBindingType()
}

func (s SddcResourceConfig) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SddcResourceConfig._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Size of the SDDC
type SddcSize struct {
    // Possible values are: 
    //
    // * SddcSize#SddcSize_VC_SIZE_MEDIUM
    // * SddcSize#SddcSize_VC_SIZE_LARGE
	VcSize *string
    // Possible values are: 
    //
    // * SddcSize#SddcSize_NSX_SIZE_SMALL
    // * SddcSize#SddcSize_NSX_SIZE_MEDIUM
    // * SddcSize#SddcSize_NSX_SIZE_LARGE
	NsxSize *string
    // Possible values are: 
    //
    // * SddcSize#SddcSize_SIZE_NSX_SMALL
    // * SddcSize#SddcSize_SIZE_MEDIUM
    // * SddcSize#SddcSize_SIZE_LARGE
    // * SddcSize#SddcSize_SIZE_NSX_LARGE
	Size *string
}
const SddcSize_VC_SIZE_MEDIUM = "medium"
const SddcSize_VC_SIZE_LARGE = "large"
const SddcSize_NSX_SIZE_SMALL = "small"
const SddcSize_NSX_SIZE_MEDIUM = "medium"
const SddcSize_NSX_SIZE_LARGE = "large"
const SddcSize_SIZE_NSX_SMALL = "NSX_SMALL"
const SddcSize_SIZE_MEDIUM = "MEDIUM"
const SddcSize_SIZE_LARGE = "LARGE"
const SddcSize_SIZE_NSX_LARGE = "NSX_LARGE"

func (s SddcSize) GetType__() bindings.BindingType {
	return SddcSizeBindingType()
}

func (s SddcSize) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SddcSize._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type SddcStateRequest struct {
	Sddcs []string
    // Possible values are: 
    //
    // * SddcStateRequest#SddcStateRequest_STATES_SCHEDULED
    // * SddcStateRequest#SddcStateRequest_STATES_RUNNING
    // * SddcStateRequest#SddcStateRequest_STATES_CANCELED
    // * SddcStateRequest#SddcStateRequest_STATES_COMPLETED
	States []string
}
const SddcStateRequest_STATES_SCHEDULED = "SCHEDULED"
const SddcStateRequest_STATES_RUNNING = "RUNNING"
const SddcStateRequest_STATES_CANCELED = "CANCELED"
const SddcStateRequest_STATES_COMPLETED = "COMPLETED"

func (s SddcStateRequest) GetType__() bindings.BindingType {
	return SddcStateRequestBindingType()
}

func (s SddcStateRequest) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SddcStateRequest._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type SddcTemplate struct {
	Updated time.Time
    // User id that last updated this record
	UserId string
    // User id that last updated this record
	UpdatedByUserId string
	Created time.Time
    // Version of this entity format: int32
	Version int64
    // User name that last updated this record
	UpdatedByUserName *string
    // User name that last updated this record
	UserName string
    // Unique ID for this entity
	Id string
    // A list of the SDDC linking configurations to use.
	AccountLinkSddcConfigs []AccountLinkSddcConfig
    // Possible values are: 
    //
    // * SddcTemplate#SddcTemplate_STATE_INITIALIZATION
    // * SddcTemplate#SddcTemplate_STATE_AVAILABLE
    // * SddcTemplate#SddcTemplate_STATE_INUSE
    // * SddcTemplate#SddcTemplate_STATE_APPLIED
    // * SddcTemplate#SddcTemplate_STATE_DELETING
    // * SddcTemplate#SddcTemplate_STATE_DELETED
    // * SddcTemplate#SddcTemplate_STATE_FAILED
	State *string
	NetworkTemplate *NetworkTemplate
    // name for SDDC configuration template
	Name *string
	SourceSddcId *string
	OrgId *string
	Sddc *Sddc
}
const SddcTemplate_STATE_INITIALIZATION = "INITIALIZATION"
const SddcTemplate_STATE_AVAILABLE = "AVAILABLE"
const SddcTemplate_STATE_INUSE = "INUSE"
const SddcTemplate_STATE_APPLIED = "APPLIED"
const SddcTemplate_STATE_DELETING = "DELETING"
const SddcTemplate_STATE_DELETED = "DELETED"
const SddcTemplate_STATE_FAILED = "FAILED"

func (s SddcTemplate) GetType__() bindings.BindingType {
	return SddcTemplateBindingType()
}

func (s SddcTemplate) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SddcTemplate._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Detailed service errors associated with a task.
type ServiceError struct {
    // Error message in English.
	DefaultMessage *string
    // The original service name of the error.
	OriginalService *string
    // The localized message.
	LocalizedMessage *string
    // The original error code of the service.
	OriginalServiceErrorCode *string
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


type ServiceQuotaRequest struct {
	Updated time.Time
    // User id that last updated this record
	UserId string
    // User id that last updated this record
	UpdatedByUserId string
	Created time.Time
    // Version of this entity format: int32
	Version int64
    // User name that last updated this record
	UpdatedByUserName *string
    // User name that last updated this record
	UserName string
    // Unique ID for this entity
	Id string
	RequesterEmail *string
    // The task for running the service quota request.
	TaskId *string
    // Region for the service quota
	Region *string
	AwsAccountNumber *string
    // The org ID for this request. This is a standard UUID.
	OrgId *string
    // Reason for this quota increase
	Reason *string
    // Possible values are: 
    //
    // * ServiceQuotaRequest#ServiceQuotaRequest_REQUEST_STATUS_NEW
    // * ServiceQuotaRequest#ServiceQuotaRequest_REQUEST_STATUS_PENDING
    // * ServiceQuotaRequest#ServiceQuotaRequest_REQUEST_STATUS_RESOLVED
    // * ServiceQuotaRequest#ServiceQuotaRequest_REQUEST_STATUS_DENIED
    // * ServiceQuotaRequest#ServiceQuotaRequest_REQUEST_STATUS_ERROR
	RequestStatus *string
    // service quota request item details
	RequestDetails map[string]RequestDetail
}
const ServiceQuotaRequest_REQUEST_STATUS_NEW = "NEW"
const ServiceQuotaRequest_REQUEST_STATUS_PENDING = "PENDING"
const ServiceQuotaRequest_REQUEST_STATUS_RESOLVED = "RESOLVED"
const ServiceQuotaRequest_REQUEST_STATUS_DENIED = "DENIED"
const ServiceQuotaRequest_REQUEST_STATUS_ERROR = "ERROR"

func (s ServiceQuotaRequest) GetType__() bindings.BindingType {
	return ServiceQuotaRequestBindingType()
}

func (s ServiceQuotaRequest) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ServiceQuotaRequest._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type Site struct {
    // Site password.
	Password *string
    // Site user id.
	UserId *string
    // Unique name for the site getting configured.
	Name *string
    // Bytes received on local network. format: int64
	RxBytesOnLocalSubnet *int64
    // Enable/disable encription.
	SecureTraffic *bool
    // Date tunnel was established.
	EstablishedDate *string
    // failure message.
	FailureMessage *string
    // Number of transmitted packets dropped.
	DroppedTxPackets *string
    // Number of received packets dropped.
	DroppedRxPackets *string
    // Possible values are: 
    //
    // * Site#Site_TUNNEL_STATUS_CONNECTED
    // * Site#Site_TUNNEL_STATUS_DISCONNECTED
    // * Site#Site_TUNNEL_STATUS_UNKNOWN
    //
    //  Site tunnel status.
	TunnelStatus *string
    // Bytes transmitted from local subnet. format: int64
	TxBytesFromLocalSubnet *int64
}
const Site_TUNNEL_STATUS_CONNECTED = "CONNECTED"
const Site_TUNNEL_STATUS_DISCONNECTED = "DISCONNECTED"
const Site_TUNNEL_STATUS_UNKNOWN = "UNKNOWN"

func (s Site) GetType__() bindings.BindingType {
	return SiteBindingType()
}

func (s Site) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for Site._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// (as there's already one SubnetInfo, use Subnet instead)
type Subnet struct {
    // subnet id
	SubnetId *string
    // subnet name
	Name *string
	RouteTables []SubnetRouteTableInfo
}

func (s Subnet) GetType__() bindings.BindingType {
	return SubnetBindingType()
}

func (s Subnet) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for Subnet._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type SubnetInfo struct {
    // Is this customer subnet compatible with the SDDC?
	Compatible *bool
    // The ID of the connected account this subnet is from.
	ConnectedAccountId *string
    // The region this subnet is from.
	RegionName *string
    // The availability zone (customer-centric) this subnet is in.
	AvailabilityZone *string
    // The ID of the subnet.
	SubnetId *string
    // The availability zone id (customer-centric) this subnet is in.
	AvailabilityZoneId *string
    // The CIDR block of the subnet.
	SubnetCidrBlock *string
    // Why a subnet is marked as not compatible. May be blank if compatible.
	Note *string
    // The ID of the VPC this subnet resides in.
	VpcId *string
    // The CIDR block of the VPC containing this subnet.
	VpcCidrBlock *string
    // The name of the subnet. This is either the tagged name or the default AWS id it was given.
	Name *string
}

func (s SubnetInfo) GetType__() bindings.BindingType {
	return SubnetInfoBindingType()
}

func (s SubnetInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SubnetInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type SubnetRouteTableInfo struct {
    // subnet id
	SubnetId *string
    // subnet - route table association id
	AssociationId *string
    // route table id
	RoutetableId *string
}

func (s SubnetRouteTableInfo) GetType__() bindings.BindingType {
	return SubnetRouteTableInfoBindingType()
}

func (s SubnetRouteTableInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SubnetRouteTableInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// details of a subscription
type SubscriptionDetails struct {
    // Possible values are: 
    //
    // * SubscriptionDetails#SubscriptionDetails_STATUS_CREATED
    // * SubscriptionDetails#SubscriptionDetails_STATUS_ACTIVATED
    // * SubscriptionDetails#SubscriptionDetails_STATUS_FAILED
    // * SubscriptionDetails#SubscriptionDetails_STATUS_CANCELLED
    // * SubscriptionDetails#SubscriptionDetails_STATUS_EXPIRED
    // * SubscriptionDetails#SubscriptionDetails_STATUS_PENDING_PROVISIONING
    // * SubscriptionDetails#SubscriptionDetails_STATUS_ORDER_SUBMITTED
    // * SubscriptionDetails#SubscriptionDetails_STATUS_SUSPENDED
    // * SubscriptionDetails#SubscriptionDetails_STATUS_TERMINATED
    // * SubscriptionDetails#SubscriptionDetails_STATUS_UKNOWN
	Status *string
	AnniversaryBillingDate *string
	EndDate *string
    // The frequency at which the customer is billed. Currently supported values are \"Upfront\" and \"Monthly\"
	BillingFrequency *string
	AutoRenewedAllowed *string
	CommitmentTerm *string
	CspSubscriptionId *string
	BillingSubscriptionId *string
	OfferVersion *string
	OfferType *OfferType
	Description *string
	ProductId *string
	Region *string
	ProductName *string
	OfferName *string
    // unit of measurment for commitment term
	CommitmentTermUom *string
	StartDate *string
	Quantity *string
}
const SubscriptionDetails_STATUS_CREATED = "CREATED"
const SubscriptionDetails_STATUS_ACTIVATED = "ACTIVATED"
const SubscriptionDetails_STATUS_FAILED = "FAILED"
const SubscriptionDetails_STATUS_CANCELLED = "CANCELLED"
const SubscriptionDetails_STATUS_EXPIRED = "EXPIRED"
const SubscriptionDetails_STATUS_PENDING_PROVISIONING = "PENDING_PROVISIONING"
const SubscriptionDetails_STATUS_ORDER_SUBMITTED = "ORDER_SUBMITTED"
const SubscriptionDetails_STATUS_SUSPENDED = "SUSPENDED"
const SubscriptionDetails_STATUS_TERMINATED = "TERMINATED"
const SubscriptionDetails_STATUS_UKNOWN = "UKNOWN"

func (s SubscriptionDetails) GetType__() bindings.BindingType {
	return SubscriptionDetailsBindingType()
}

func (s SubscriptionDetails) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SubscriptionDetails._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Details of products that are available for purchase.
type SubscriptionProducts struct {
    // The name of the product
	Product *string
    // A list of different types/version for the product.
	Types []string
}

func (s SubscriptionProducts) GetType__() bindings.BindingType {
	return SubscriptionProductsBindingType()
}

func (s SubscriptionProducts) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SubscriptionProducts._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type SubscriptionRequest struct {
    // The product for which subscription needs to be created. Refer /vmc/api/orgs/{orgId}/products.
	Product *string
    // Old identifier for product. \*Deprecarted\*. See product and type
	ProductType string
	ProductId *string
    // Frequency of the billing.
	BillingFrequency *string
	Region string
	CommitmentTerm string
	OfferContextId *string
	OfferVersion string
	OfferName string
	Quantity int64
    // The type of the product for which the subscription needs to be created.
	Type_ *string
	Price *int64
	ProductChargeId *string
}

func (s SubscriptionRequest) GetType__() bindings.BindingType {
	return SubscriptionRequestBindingType()
}

func (s SubscriptionRequest) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SubscriptionRequest._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type SupportWindow struct {
    // Possible values are: 
    //
    // * SupportWindow#SupportWindow_START_DAY_MONDAY
    // * SupportWindow#SupportWindow_START_DAY_TUESDAY
    // * SupportWindow#SupportWindow_START_DAY_WEDNESDAY
    // * SupportWindow#SupportWindow_START_DAY_THURSDAY
    // * SupportWindow#SupportWindow_START_DAY_FRIDAY
    // * SupportWindow#SupportWindow_START_DAY_SATURDAY
    // * SupportWindow#SupportWindow_START_DAY_SUNDAY
	StartDay *string
	Seats *int64
    // SDDCs in this window format: UUID
	Sddcs []string
	DurationHours *int64
	StartHour *int64
	SupportWindowId *string
	Metadata *data.StructValue
}
const SupportWindow_START_DAY_MONDAY = "MONDAY"
const SupportWindow_START_DAY_TUESDAY = "TUESDAY"
const SupportWindow_START_DAY_WEDNESDAY = "WEDNESDAY"
const SupportWindow_START_DAY_THURSDAY = "THURSDAY"
const SupportWindow_START_DAY_FRIDAY = "FRIDAY"
const SupportWindow_START_DAY_SATURDAY = "SATURDAY"
const SupportWindow_START_DAY_SUNDAY = "SUNDAY"

func (s SupportWindow) GetType__() bindings.BindingType {
	return SupportWindowBindingType()
}

func (s SupportWindow) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SupportWindow._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type SupportWindowId struct {
    // Support Window ID
	WindowId *string
}

func (s SupportWindowId) GetType__() bindings.BindingType {
	return SupportWindowIdBindingType()
}

func (s SupportWindowId) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SupportWindowId._GetDataValue method - %s",
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
	Created time.Time
    // Version of this entity format: int32
	Version int64
    // User name that last updated this record
	UpdatedByUserName *string
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
	LocalizedErrorMessage *string
    // UUID of the resource the task is acting upon
	ResourceId *string
    // If this task was created by another task - this provides the linkage. Mostly for debugging.
	ParentTaskId *string
	TaskVersion *string
    // (Optional) Client provided uniqifier to make task creation idempotent. Be aware not all tasks support this. For tasks that do - supplying the same correlation Id, for the same task type, within a predefined window will ensure the operation happens at most once.
	CorrelationId *string
    // Entity version of the resource at the start of the task. This is only set for some task types. format: int32
	StartResourceEntityVersion *int64
	SubStatus *string
	TaskType *string
	StartTime *time.Time
    // Task progress phases involved in current task execution
	TaskProgressPhases []TaskProgressPhase
	ErrorMessage *string
	OrgId *string
    // Entity version of the resource at the end of the task. This is only set for some task types. format: int32
	EndResourceEntityVersion *int64
    // Service errors returned from SDDC services.
	ServiceErrors []ServiceError
	OrgType *string
    // Estimated remaining time in minute of the task execution, < 0 means no estimation for the task. format: int32
	EstimatedRemainingMinutes *int64
	Params *data.StructValue
    // Estimated progress percentage the task executed format: int32
	ProgressPercent *int64
    // The current in progress phase ID in the task execution, if none in progress, empty string returned.
	PhaseInProgress *string
    // Type of resource being acted upon
	ResourceType *string
	EndTime *time.Time
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


// Holder for term billing options.
type TermBillingOptions struct {
	UnitPrice *string
	BillingFrequency *string
}

func (s TermBillingOptions) GetType__() bindings.BindingType {
	return TermBillingOptionsBindingType()
}

func (s TermBillingOptions) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for TermBillingOptions._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Holder for the term offer instances.
type TermOfferInstance struct {
	Description string
	Product string
    // Deprecated. Please use product and type fields instead.
	ProductType *string
	Name string
	Currency string
	Region string
	CommitmentTerm int64
    // (deprecated. unit_price is moved into TermBillingOptions. For backward compatibility, this field reflect \"Prepaid\" price at the offer level.)
	UnitPrice string
	BillingOptions []TermBillingOptions
	Version string
	OfferContextId *string
	ProductChargeId *string
	Type_ string
	ProductId *string
}

func (s TermOfferInstance) GetType__() bindings.BindingType {
	return TermOfferInstanceBindingType()
}

func (s TermOfferInstance) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for TermOfferInstance._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type TermsOfServiceResult struct {
    // The terms of service ID requested.
	TermsId *string
    // Wehther or not the terms requested have been signed.
	Signed *bool
}

func (s TermsOfServiceResult) GetType__() bindings.BindingType {
	return TermsOfServiceResultBindingType()
}

func (s TermsOfServiceResult) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for TermsOfServiceResult._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type UpdateCredentials struct {
    // Username of the credentials
	Username string
    // Password associated with the credentials
	Password string
}

func (s UpdateCredentials) GetType__() bindings.BindingType {
	return UpdateCredentialsBindingType()
}

func (s UpdateCredentials) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for UpdateCredentials._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type VmcLocale struct {
    // The locale to be used for translating responses for the session
	Locale *string
}

func (s VmcLocale) GetType__() bindings.BindingType {
	return VmcLocaleBindingType()
}

func (s VmcLocale) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for VmcLocale._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type VpcInfo struct {
	VpcCidr *string
	VgwId *string
	EsxPublicSecurityGroupId *string
    // set of virtual interfaces attached to the sddc
	VifIds []string
	VmSecurityGroupId *string
    // (deprecated)
	RouteTableId *string
    // Id of the NSX edge associated with this VPC (deprecated)
	EdgeSubnetId *string
    // vpc id
	Id *string
    // Id of the association between subnet and route-table (deprecated)
	ApiAssociationId *string
    // Id associated with this VPC (deprecated)
	ApiSubnetId *string
    // (deprecated)
	PrivateSubnetId *string
    // (deprecated)
	PrivateAssociationId *string
	EsxSecurityGroupId *string
    // (deprecated)
	SubnetId *string
	InternetGatewayId *string
	SecurityGroupId *string
    // (deprecated)
	AssociationId *string
    // Route table which contains the route to VGW (deprecated)
	VgwRouteTableId *string
    // List of edge vm Ips of traffic gourps added during scale-out
	TrafficGroupEdgeVmIps []string
    // Id of the association between edge subnet and route-table (deprecated)
	EdgeAssociationId *string
	Provider *string
    // Mapping from AZ to a list of IP addresses assigned to TGW ENI that's connected with Vpc
	TgwIps map[string][]string
    // (deprecated)
	PeeringConnectionId *string
	NetworkType *string
	AvailableZones []AvailableZoneInfo
    // map from routeTableName to routeTableInfo
	Routetables map[string]RouteTableInfo
}

func (s VpcInfo) GetType__() bindings.BindingType {
	return VpcInfoBindingType()
}

func (s VpcInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for VpcInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type VpcInfoSubnets struct {
    // The ID of the VPC these subnets belong to.
	VpcId *string
    // The overall CIDR block of the VPC. This is the AWS primary CIDR block.
	CidrBlock *string
    // The description of the VPC; usually it's name or id.
	Description *string
	Subnets []SubnetInfo
}

func (s VpcInfoSubnets) GetType__() bindings.BindingType {
	return VpcInfoSubnetsBindingType()
}

func (s VpcInfoSubnets) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for VpcInfoSubnets._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type Vpn struct {
	Version *string
	OnPremGatewayIp *string
	OnPremNetworkCidr *string
	PfsEnabled *bool
	Id *string
	ChannelStatus *VpnChannelStatus
	OnPremNatIp *string
	Name *string
	InternalNetworkIds []string
	TunnelStatuses []VpnTunnelStatus
    // Possible values are: 
    //
    // * Vpn#Vpn_ENCRYPTION_AES
    // * Vpn#Vpn_ENCRYPTION_AES256
    // * Vpn#Vpn_ENCRYPTION_AES_GCM
    // * Vpn#Vpn_ENCRYPTION_TRIPLE_DES
    // * Vpn#Vpn_ENCRYPTION_UNKNOWN
	Encryption *string
	Enabled *bool
    // Possible values are: 
    //
    // * Vpn#Vpn_STATE_CONNECTED
    // * Vpn#Vpn_STATE_DISCONNECTED
    // * Vpn#Vpn_STATE_PARTIALLY_CONNECTED
    // * Vpn#Vpn_STATE_UNKNOWN
	State *string
    // Possible values are: 
    //
    // * Vpn#Vpn_DH_GROUP_DH2
    // * Vpn#Vpn_DH_GROUP_DH5
    // * Vpn#Vpn_DH_GROUP_DH14
    // * Vpn#Vpn_DH_GROUP_DH15
    // * Vpn#Vpn_DH_GROUP_DH16
    // * Vpn#Vpn_DH_GROUP_UNKNOWN
	DhGroup *string
    // Possible values are: 
    //
    // * Vpn#Vpn_AUTHENTICATION_PSK
    // * Vpn#Vpn_AUTHENTICATION_UNKNOWN
	Authentication *string
	PreSharedKey *string
    // Possible values are: 
    //
    // * Vpn#Vpn_IKE_OPTION_IKEV1
    // * Vpn#Vpn_IKE_OPTION_IKEV2
	IkeOption *string
    // Possible values are: 
    //
    // * Vpn#Vpn_DIGEST_ALGORITHM_SHA1
    // * Vpn#Vpn_DIGEST_ALGORITHM_SHA_256
	DigestAlgorithm *string
}
const Vpn_ENCRYPTION_AES = "AES"
const Vpn_ENCRYPTION_AES256 = "AES256"
const Vpn_ENCRYPTION_AES_GCM = "AES_GCM"
const Vpn_ENCRYPTION_TRIPLE_DES = "TRIPLE_DES"
const Vpn_ENCRYPTION_UNKNOWN = "UNKNOWN"
const Vpn_STATE_CONNECTED = "CONNECTED"
const Vpn_STATE_DISCONNECTED = "DISCONNECTED"
const Vpn_STATE_PARTIALLY_CONNECTED = "PARTIALLY_CONNECTED"
const Vpn_STATE_UNKNOWN = "UNKNOWN"
const Vpn_DH_GROUP_DH2 = "DH2"
const Vpn_DH_GROUP_DH5 = "DH5"
const Vpn_DH_GROUP_DH14 = "DH14"
const Vpn_DH_GROUP_DH15 = "DH15"
const Vpn_DH_GROUP_DH16 = "DH16"
const Vpn_DH_GROUP_UNKNOWN = "UNKNOWN"
const Vpn_AUTHENTICATION_PSK = "PSK"
const Vpn_AUTHENTICATION_UNKNOWN = "UNKNOWN"
const Vpn_IKE_OPTION_IKEV1 = "IKEV1"
const Vpn_IKE_OPTION_IKEV2 = "IKEV2"
const Vpn_DIGEST_ALGORITHM_SHA1 = "SHA1"
const Vpn_DIGEST_ALGORITHM_SHA_256 = "SHA_256"

func (s Vpn) GetType__() bindings.BindingType {
	return VpnBindingType()
}

func (s Vpn) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for Vpn._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type VpnChannelStatus struct {
    // Possible values are: 
    //
    // * VpnChannelStatus#VpnChannelStatus_CHANNEL_STATUS_CONNECTED
    // * VpnChannelStatus#VpnChannelStatus_CHANNEL_STATUS_DISCONNECTED
    // * VpnChannelStatus#VpnChannelStatus_CHANNEL_STATUS_UNKNOWN
	ChannelStatus *string
	ChannelState *string
	LastInfoMessage *string
	FailureMessage *string
}
const VpnChannelStatus_CHANNEL_STATUS_CONNECTED = "CONNECTED"
const VpnChannelStatus_CHANNEL_STATUS_DISCONNECTED = "DISCONNECTED"
const VpnChannelStatus_CHANNEL_STATUS_UNKNOWN = "UNKNOWN"

func (s VpnChannelStatus) GetType__() bindings.BindingType {
	return VpnChannelStatusBindingType()
}

func (s VpnChannelStatus) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for VpnChannelStatus._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type VpnTunnelStatus struct {
	OnPremSubnet *string
	TrafficStats *VpnTunnelTrafficStats
	LastInfoMessage *string
	LocalSubnet *string
	TunnelState *string
	FailureMessage *string
    // Possible values are: 
    //
    // * VpnTunnelStatus#VpnTunnelStatus_TUNNEL_STATUS_CONNECTED
    // * VpnTunnelStatus#VpnTunnelStatus_TUNNEL_STATUS_DISCONNECTED
    // * VpnTunnelStatus#VpnTunnelStatus_TUNNEL_STATUS_UNKNOWN
	TunnelStatus *string
}
const VpnTunnelStatus_TUNNEL_STATUS_CONNECTED = "CONNECTED"
const VpnTunnelStatus_TUNNEL_STATUS_DISCONNECTED = "DISCONNECTED"
const VpnTunnelStatus_TUNNEL_STATUS_UNKNOWN = "UNKNOWN"

func (s VpnTunnelStatus) GetType__() bindings.BindingType {
	return VpnTunnelStatusBindingType()
}

func (s VpnTunnelStatus) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for VpnTunnelStatus._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type VpnTunnelTrafficStats struct {
	PacketsOut *string
	PacketReceivedErrors *string
	RxBytesOnLocalSubnet *string
	ReplayErrors *string
	SequenceNumberOverFlowErrors *string
	EncryptionFailures *string
	IntegrityErrors *string
	PacketSentErrors *string
	DecryptionFailures *string
	PacketsIn *string
	TxBytesFromLocalSubnet *string
}

func (s VpnTunnelTrafficStats) GetType__() bindings.BindingType {
	return VpnTunnelTrafficStatsBindingType()
}

func (s VpnTunnelTrafficStats) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for VpnTunnelTrafficStats._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Infomation about an available vSAN capacity in a cluster.
type VsanAvailableCapacity struct {
	Cost string
	Quality string
	Size int64
}

func (s VsanAvailableCapacity) GetType__() bindings.BindingType {
	return VsanAvailableCapacityBindingType()
}

func (s VsanAvailableCapacity) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for VsanAvailableCapacity._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Bias for reconfiguring vSAN in a cluster.
type VsanClusterReconfigBias struct {
	ShortDescription string
	FullDescription string
	Id string
}

func (s VsanClusterReconfigBias) GetType__() bindings.BindingType {
	return VsanClusterReconfigBiasBindingType()
}

func (s VsanClusterReconfigBias) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for VsanClusterReconfigBias._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// Storage constraint information for reconfiguring vSAN in existing cluster.
type VsanClusterReconfigConstraints struct {
    // Biases to reconfigure vSAN in an existing cluster.
	ReconfigBiases []VsanClusterReconfigBias
    // A map of VsanClusterReconfigBias id to the list of VsanAvailableCapacity. It gives all of available vSAN capacities for each of reconfiguration biases.
	AvailableCapacities map[string][]VsanAvailableCapacity
    // A map of VsanClusterReconfigBias id to a VsanAvailableCapacity. It gives the default VsanAvailableCapacity for each of reconfiguration biases.
	DefaultCapacities map[string]VsanAvailableCapacity
    // The number of hosts in a cluster for the constraints. format: int32
	Hosts int64
    // The id of default VsanClusterReconfigBias for this constraints.
	DefaultReconfigBiasId string
}

func (s VsanClusterReconfigConstraints) GetType__() bindings.BindingType {
	return VsanClusterReconfigConstraintsBindingType()
}

func (s VsanClusterReconfigConstraints) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for VsanClusterReconfigConstraints._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


// This describes the possible physical storage capacity choices for use with a given VsanStorageDesigner implementation. These choices are specific to a customer-defined number of hosts per cluster.
type VsanConfigConstraints struct {
    // Maximum capacity supported for cluster (GiB). format: int64
	MaxCapacity int64
    // List of supported capacities which may offer preferable performance (GiB). format: int64
	RecommendedCapacities []int64
    // Increment to be added to min_capacity to result in a supported capacity (GiB). format: int64
	SupportedCapacityIncrement *int64
    // Minimum capacity supported for cluster (GiB). format: int64
	MinCapacity int64
    // Number of hosts in cluster. format: int64
	NumHosts int64
}

func (s VsanConfigConstraints) GetType__() bindings.BindingType {
	return VsanConfigConstraintsBindingType()
}

func (s VsanConfigConstraints) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for VsanConfigConstraints._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}


type VsanEncryptionConfig struct {
    // Port to connect to AWS Key Management Service
	Port *int64
    // Public certificate used to connect to AWS Key Management Service
	Certificate *string
}

func (s VsanEncryptionConfig) GetType__() bindings.BindingType {
	return VsanEncryptionConfigBindingType()
}

func (s VsanEncryptionConfig) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for VsanEncryptionConfig._GetDataValue method - %s",
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
	fields["created"] = bindings.NewDateTimeType()
	fieldNameMap["created"] = "Created"
	fields["version"] = bindings.NewIntegerType()
	fieldNameMap["version"] = "Version"
	fields["updated_by_user_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["updated_by_user_name"] = "UpdatedByUserName"
	fields["user_name"] = bindings.NewStringType()
	fieldNameMap["user_name"] = "UserName"
	fields["id"] = bindings.NewStringType()
	fieldNameMap["id"] = "Id"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.abstract_entity", fields, reflect.TypeOf(AbstractEntity{}), fieldNameMap, validators)
}

func AccountLinkConfigBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["delay_account_link"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["delay_account_link"] = "DelayAccountLink"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.account_link_config", fields, reflect.TypeOf(AccountLinkConfig{}), fieldNameMap, validators)
}

func AccountLinkSddcConfigBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["customer_subnet_ids"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["customer_subnet_ids"] = "CustomerSubnetIds"
	fields["connected_account_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["connected_account_id"] = "ConnectedAccountId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.account_link_sddc_config", fields, reflect.TypeOf(AccountLinkSddcConfig{}), fieldNameMap, validators)
}

func AgentBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["internal_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["internal_ip"] = "InternalIp"
	fields["agent_url"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["agent_url"] = "AgentUrl"
	fields["management_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["management_ip"] = "ManagementIp"
	fields["hostname_verifier_enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["hostname_verifier_enabled"] = "HostnameVerifierEnabled"
	fields["master"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["master"] = "Master"
	fields["network_netmask"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["network_netmask"] = "NetworkNetmask"
	fields["network_gateway"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["network_gateway"] = "NetworkGateway"
	fields["provider"] = bindings.NewStringType()
	fieldNameMap["provider"] = "Provider"
	fields["cert_enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["cert_enabled"] = "CertEnabled"
	fields["agent_state"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["agent_state"] = "AgentState"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.agent", fields, reflect.TypeOf(Agent{}), fieldNameMap, validators)
}

func AmiInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["instance_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["instance_type"] = "InstanceType"
	fields["region"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["region"] = "Region"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.ami_info", fields, reflect.TypeOf(AmiInfo{}), fieldNameMap, validators)
}

func AvailableZoneInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["subnets"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(SubnetBindingType), reflect.TypeOf([]Subnet{})))
	fieldNameMap["subnets"] = "Subnets"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.available_zone_info", fields, reflect.TypeOf(AvailableZoneInfo{}), fieldNameMap, validators)
}

func AwsAgentBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["instance_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["instance_id"] = "InstanceId"
	fields["key_pair"] = bindings.NewOptionalType(bindings.NewReferenceType(AwsKeyPairBindingType))
	fieldNameMap["key_pair"] = "KeyPair"
	fields["internal_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["internal_ip"] = "InternalIp"
	fields["agent_url"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["agent_url"] = "AgentUrl"
	fields["management_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["management_ip"] = "ManagementIp"
	fields["hostname_verifier_enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["hostname_verifier_enabled"] = "HostnameVerifierEnabled"
	fields["master"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["master"] = "Master"
	fields["network_netmask"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["network_netmask"] = "NetworkNetmask"
	fields["network_gateway"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["network_gateway"] = "NetworkGateway"
	fields["provider"] = bindings.NewStringType()
	fieldNameMap["provider"] = "Provider"
	fields["cert_enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["cert_enabled"] = "CertEnabled"
	fields["agent_state"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["agent_state"] = "AgentState"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.aws_agent", fields, reflect.TypeOf(AwsAgent{}), fieldNameMap, validators)
}

func AwsCloudProviderBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["regions"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["regions"] = "Regions"
	fields["provider"] = bindings.NewStringType()
	fieldNameMap["provider"] = "Provider"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.aws_cloud_provider", fields, reflect.TypeOf(AwsCloudProvider{}), fieldNameMap, validators)
}

func AwsCompatibleSubnetsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["customer_available_zones"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["customer_available_zones"] = "CustomerAvailableZones"
	fields["vpc_map"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewReferenceType(VpcInfoSubnetsBindingType),reflect.TypeOf(map[string]VpcInfoSubnets{})))
	fieldNameMap["vpc_map"] = "VpcMap"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.aws_compatible_subnets", fields, reflect.TypeOf(AwsCompatibleSubnets{}), fieldNameMap, validators)
}

func AwsCustomerConnectedAccountBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["updated"] = bindings.NewDateTimeType()
	fieldNameMap["updated"] = "Updated"
	fields["user_id"] = bindings.NewStringType()
	fieldNameMap["user_id"] = "UserId"
	fields["updated_by_user_id"] = bindings.NewStringType()
	fieldNameMap["updated_by_user_id"] = "UpdatedByUserId"
	fields["created"] = bindings.NewDateTimeType()
	fieldNameMap["created"] = "Created"
	fields["version"] = bindings.NewIntegerType()
	fieldNameMap["version"] = "Version"
	fields["updated_by_user_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["updated_by_user_name"] = "UpdatedByUserName"
	fields["user_name"] = bindings.NewStringType()
	fieldNameMap["user_name"] = "UserName"
	fields["id"] = bindings.NewStringType()
	fieldNameMap["id"] = "Id"
	fields["policy_payer_arn"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["policy_payer_arn"] = "PolicyPayerArn"
	fields["region_to_az_to_shadow_mapping"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewMapType(bindings.NewStringType(), bindings.NewStringType(),reflect.TypeOf(map[string]string{})),reflect.TypeOf(map[string]map[string]string{})))
	fieldNameMap["region_to_az_to_shadow_mapping"] = "RegionToAzToShadowMapping"
	fields["org_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fields["cf_stack_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cf_stack_name"] = "CfStackName"
	fields["state"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["state"] = "State"
	fields["account_number"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["account_number"] = "AccountNumber"
	fields["policy_service_arn"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["policy_service_arn"] = "PolicyServiceArn"
	fields["policy_external_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["policy_external_id"] = "PolicyExternalId"
	fields["policy_payer_linked_arn"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["policy_payer_linked_arn"] = "PolicyPayerLinkedArn"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.aws_customer_connected_account", fields, reflect.TypeOf(AwsCustomerConnectedAccount{}), fieldNameMap, validators)
}

func AwsEsxHostBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["internal_public_ip_pool"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(SddcPublicIpBindingType), reflect.TypeOf([]SddcPublicIp{})))
	fieldNameMap["internal_public_ip_pool"] = "InternalPublicIpPool"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["availability_zone"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["availability_zone"] = "AvailabilityZone"
	fields["esx_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["esx_id"] = "EsxId"
	fields["hostname"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["hostname"] = "Hostname"
	fields["provider"] = bindings.NewStringType()
	fieldNameMap["provider"] = "Provider"
	fields["instance_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["instance_type"] = "InstanceType"
	fields["mac_address"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["mac_address"] = "MacAddress"
	fields["custom_properties"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewStringType(),reflect.TypeOf(map[string]string{})))
	fieldNameMap["custom_properties"] = "CustomProperties"
	fields["esx_state"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["esx_state"] = "EsxState"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.aws_esx_host", fields, reflect.TypeOf(AwsEsxHost{}), fieldNameMap, validators)
}

func AwsKeyPairBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["key_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["key_name"] = "KeyName"
	fields["key_fingerprint"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["key_fingerprint"] = "KeyFingerprint"
	fields["key_material"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["key_material"] = "KeyMaterial"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.aws_key_pair", fields, reflect.TypeOf(AwsKeyPair{}), fieldNameMap, validators)
}

func AwsKmsInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["amazon_resource_name"] = bindings.NewStringType()
	fieldNameMap["amazon_resource_name"] = "AmazonResourceName"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.aws_kms_info", fields, reflect.TypeOf(AwsKmsInfo{}), fieldNameMap, validators)
}

func AwsSddcConfigBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["region"] = bindings.NewStringType()
	fieldNameMap["region"] = "Region"
	fields["msft_license_config"] = bindings.NewOptionalType(bindings.NewReferenceType(MsftLicensingConfigBindingType))
	fieldNameMap["msft_license_config"] = "MsftLicenseConfig"
	fields["vpc_cidr"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vpc_cidr"] = "VpcCidr"
	fields["host_instance_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vmc.model.host_instance_types", reflect.TypeOf(HostInstanceTypes(HostInstanceTypes_I3_METAL))))
	fieldNameMap["host_instance_type"] = "HostInstanceType"
	fields["skip_creating_vxlan"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["skip_creating_vxlan"] = "SkipCreatingVxlan"
	fields["vxlan_subnet"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vxlan_subnet"] = "VxlanSubnet"
	fields["size"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["size"] = "Size"
	fields["storage_capacity"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["storage_capacity"] = "StorageCapacity"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["account_link_sddc_config"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(AccountLinkSddcConfigBindingType), reflect.TypeOf([]AccountLinkSddcConfig{})))
	fieldNameMap["account_link_sddc_config"] = "AccountLinkSddcConfig"
	fields["sddc_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sddc_id"] = "SddcId"
	fields["num_hosts"] = bindings.NewIntegerType()
	fieldNameMap["num_hosts"] = "NumHosts"
	fields["sddc_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sddc_type"] = "SddcType"
	fields["account_link_config"] = bindings.NewOptionalType(bindings.NewReferenceType(AccountLinkConfigBindingType))
	fieldNameMap["account_link_config"] = "AccountLinkConfig"
	fields["provider"] = bindings.NewStringType()
	fieldNameMap["provider"] = "Provider"
	fields["sso_domain"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sso_domain"] = "SsoDomain"
	fields["sddc_template_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sddc_template_id"] = "SddcTemplateId"
	fields["deployment_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["deployment_type"] = "DeploymentType"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.aws_sddc_config", fields, reflect.TypeOf(AwsSddcConfig{}), fieldNameMap, validators)
}

func AwsSddcConnectionBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["updated"] = bindings.NewDateTimeType()
	fieldNameMap["updated"] = "Updated"
	fields["user_id"] = bindings.NewStringType()
	fieldNameMap["user_id"] = "UserId"
	fields["updated_by_user_id"] = bindings.NewStringType()
	fieldNameMap["updated_by_user_id"] = "UpdatedByUserId"
	fields["created"] = bindings.NewDateTimeType()
	fieldNameMap["created"] = "Created"
	fields["version"] = bindings.NewIntegerType()
	fieldNameMap["version"] = "Version"
	fields["updated_by_user_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["updated_by_user_name"] = "UpdatedByUserName"
	fields["user_name"] = bindings.NewStringType()
	fieldNameMap["user_name"] = "UserName"
	fields["id"] = bindings.NewStringType()
	fieldNameMap["id"] = "Id"
	fields["cidr_block_subnet"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cidr_block_subnet"] = "CidrBlockSubnet"
	fields["connected_account_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["connected_account_id"] = "ConnectedAccountId"
	fields["eni_group"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["eni_group"] = "EniGroup"
	fields["subnet_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["subnet_id"] = "SubnetId"
	fields["cgw_present"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["cgw_present"] = "CgwPresent"
	fields["org_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fields["sddc_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sddc_id"] = "SddcId"
	fields["cidr_block_vpc"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cidr_block_vpc"] = "CidrBlockVpc"
	fields["connection_order"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["connection_order"] = "ConnectionOrder"
	fields["state"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["state"] = "State"
	fields["subnet_availability_zone"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["subnet_availability_zone"] = "SubnetAvailabilityZone"
	fields["vpc_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vpc_id"] = "VpcId"
	fields["customer_eni_infos"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(CustomerEniInfoBindingType), reflect.TypeOf([]CustomerEniInfo{})))
	fieldNameMap["customer_eni_infos"] = "CustomerEniInfos"
	fields["default_route_table"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["default_route_table"] = "DefaultRouteTable"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.aws_sddc_connection", fields, reflect.TypeOf(AwsSddcConnection{}), fieldNameMap, validators)
}

func AwsSddcResourceConfigBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["backup_restore_bucket"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["backup_restore_bucket"] = "BackupRestoreBucket"
	fields["public_ip_pool"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(SddcPublicIpBindingType), reflect.TypeOf([]SddcPublicIp{})))
	fieldNameMap["public_ip_pool"] = "PublicIpPool"
	fields["vpc_info"] = bindings.NewOptionalType(bindings.NewReferenceType(VpcInfoBindingType))
	fieldNameMap["vpc_info"] = "VpcInfo"
	fields["kms_vpc_endpoint"] = bindings.NewOptionalType(bindings.NewReferenceType(KmsVpcEndpointBindingType))
	fieldNameMap["kms_vpc_endpoint"] = "KmsVpcEndpoint"
	fields["max_num_public_ip"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["max_num_public_ip"] = "MaxNumPublicIp"
	fields["account_link_sddc_config"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(SddcLinkConfigBindingType), reflect.TypeOf([]SddcLinkConfig{})))
	fieldNameMap["account_link_sddc_config"] = "AccountLinkSddcConfig"
	fields["vsan_encryption_config"] = bindings.NewOptionalType(bindings.NewReferenceType(VsanEncryptionConfigBindingType))
	fieldNameMap["vsan_encryption_config"] = "VsanEncryptionConfig"
	fields["vpc_info_peered_agent"] = bindings.NewOptionalType(bindings.NewReferenceType(VpcInfoBindingType))
	fieldNameMap["vpc_info_peered_agent"] = "VpcInfoPeeredAgent"
	fields["mgmt_appliance_network_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["mgmt_appliance_network_name"] = "MgmtApplianceNetworkName"
	fields["nsxt"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["nsxt"] = "Nsxt"
	fields["mgw_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["mgw_id"] = "MgwId"
	fields["nsx_mgr_url"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["nsx_mgr_url"] = "NsxMgrUrl"
	fields["psc_management_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["psc_management_ip"] = "PscManagementIp"
	fields["psc_url"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["psc_url"] = "PscUrl"
	fields["cgws"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["cgws"] = "Cgws"
	fields["availability_zones"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["availability_zones"] = "AvailabilityZones"
	fields["management_ds"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["management_ds"] = "ManagementDs"
	fields["nsx_api_public_endpoint_url"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["nsx_api_public_endpoint_url"] = "NsxApiPublicEndpointUrl"
	fields["custom_properties"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewStringType(),reflect.TypeOf(map[string]string{})))
	fieldNameMap["custom_properties"] = "CustomProperties"
	fields["cloud_password"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cloud_password"] = "CloudPassword"
	fields["provider"] = bindings.NewStringType()
	fieldNameMap["provider"] = "Provider"
	fields["clusters"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ClusterBindingType), reflect.TypeOf([]Cluster{})))
	fieldNameMap["clusters"] = "Clusters"
	fields["vc_management_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vc_management_ip"] = "VcManagementIp"
	fields["sddc_networks"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["sddc_networks"] = "SddcNetworks"
	fields["cloud_username"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cloud_username"] = "CloudUsername"
	fields["esx_hosts"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(AwsEsxHostBindingType), reflect.TypeOf([]AwsEsxHost{})))
	fieldNameMap["esx_hosts"] = "EsxHosts"
	fields["nsx_mgr_management_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["nsx_mgr_management_ip"] = "NsxMgrManagementIp"
	fields["vc_instance_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vc_instance_id"] = "VcInstanceId"
	fields["esx_cluster_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["esx_cluster_id"] = "EsxClusterId"
	fields["vc_public_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vc_public_ip"] = "VcPublicIp"
	fields["skip_creating_vxlan"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["skip_creating_vxlan"] = "SkipCreatingVxlan"
	fields["vc_url"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vc_url"] = "VcUrl"
	fields["sddc_manifest"] = bindings.NewOptionalType(bindings.NewReferenceType(SddcManifestBindingType))
	fieldNameMap["sddc_manifest"] = "SddcManifest"
	fields["vxlan_subnet"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vxlan_subnet"] = "VxlanSubnet"
	fields["cloud_user_group"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cloud_user_group"] = "CloudUserGroup"
	fields["management_rp"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["management_rp"] = "ManagementRp"
	fields["region"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["region"] = "Region"
	fields["witness_availability_zone"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["witness_availability_zone"] = "WitnessAvailabilityZone"
	fields["sddc_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sddc_id"] = "SddcId"
	fields["pop_agent_xeni_connection"] = bindings.NewOptionalType(bindings.NewReferenceType(PopAgentXeniConnectionBindingType))
	fieldNameMap["pop_agent_xeni_connection"] = "PopAgentXeniConnection"
	fields["sddc_size"] = bindings.NewOptionalType(bindings.NewReferenceType(SddcSizeBindingType))
	fieldNameMap["sddc_size"] = "SddcSize"
	fields["nsx_controller_ips"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["nsx_controller_ips"] = "NsxControllerIps"
	fields["esx_host_subnet"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["esx_host_subnet"] = "EsxHostSubnet"
	fields["sso_domain"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sso_domain"] = "SsoDomain"
	fields["deployment_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["deployment_type"] = "DeploymentType"
	fields["msft_license_config"] = bindings.NewOptionalType(bindings.NewReferenceType(MsftLicensingConfigBindingType))
	fieldNameMap["msft_license_config"] = "MsftLicenseConfig"
	fields["nsxt_addons"] = bindings.NewOptionalType(bindings.NewReferenceType(NsxtAddonsBindingType))
	fieldNameMap["nsxt_addons"] = "NsxtAddons"
	fields["dns_with_management_vm_private_ip"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["dns_with_management_vm_private_ip"] = "DnsWithManagementVmPrivateIp"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.aws_sddc_resource_config", fields, reflect.TypeOf(AwsSddcResourceConfig{}), fieldNameMap, validators)
}

func AwsSubnetBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["connected_account_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["connected_account_id"] = "ConnectedAccountId"
	fields["region_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["region_name"] = "RegionName"
	fields["availability_zone"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["availability_zone"] = "AvailabilityZone"
	fields["subnet_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["subnet_id"] = "SubnetId"
	fields["subnet_cidr_block"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["subnet_cidr_block"] = "SubnetCidrBlock"
	fields["is_compatible"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["is_compatible"] = "IsCompatible"
	fields["vpc_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vpc_id"] = "VpcId"
	fields["vpc_cidr_block"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vpc_cidr_block"] = "VpcCidrBlock"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.aws_subnet", fields, reflect.TypeOf(AwsSubnet{}), fieldNameMap, validators)
}

func CloudProviderBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["provider"] = bindings.NewStringType()
	fieldNameMap["provider"] = "Provider"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.cloud_provider", fields, reflect.TypeOf(CloudProvider{}), fieldNameMap, validators)
}

func ClusterBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["esx_host_list"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(AwsEsxHostBindingType), reflect.TypeOf([]AwsEsxHost{})))
	fieldNameMap["esx_host_list"] = "EsxHostList"
	fields["msft_license_config"] = bindings.NewOptionalType(bindings.NewReferenceType(MsftLicensingConfigBindingType))
	fieldNameMap["msft_license_config"] = "MsftLicenseConfig"
	fields["cluster_state"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cluster_state"] = "ClusterState"
	fields["aws_kms_info"] = bindings.NewOptionalType(bindings.NewReferenceType(AwsKmsInfoBindingType))
	fieldNameMap["aws_kms_info"] = "AwsKmsInfo"
	fields["cluster_capacity"] = bindings.NewOptionalType(bindings.NewReferenceType(EntityCapacityBindingType))
	fieldNameMap["cluster_capacity"] = "ClusterCapacity"
	fields["esx_host_info"] = bindings.NewOptionalType(bindings.NewReferenceType(EsxHostInfoBindingType))
	fieldNameMap["esx_host_info"] = "EsxHostInfo"
	fields["host_cpu_cores_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["host_cpu_cores_count"] = "HostCpuCoresCount"
	fields["cluster_id"] = bindings.NewStringType()
	fieldNameMap["cluster_id"] = "ClusterId"
	fields["cluster_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cluster_name"] = "ClusterName"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.cluster", fields, reflect.TypeOf(Cluster{}), fieldNameMap, validators)
}

func ClusterConfigBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["host_cpu_cores_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["host_cpu_cores_count"] = "HostCpuCoresCount"
	fields["host_instance_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vmc.model.host_instance_types", reflect.TypeOf(HostInstanceTypes(HostInstanceTypes_I3_METAL))))
	fieldNameMap["host_instance_type"] = "HostInstanceType"
	fields["storage_capacity"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["storage_capacity"] = "StorageCapacity"
	fields["msft_license_config"] = bindings.NewOptionalType(bindings.NewReferenceType(MsftLicensingConfigBindingType))
	fieldNameMap["msft_license_config"] = "MsftLicenseConfig"
	fields["num_hosts"] = bindings.NewIntegerType()
	fieldNameMap["num_hosts"] = "NumHosts"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.cluster_config", fields, reflect.TypeOf(ClusterConfig{}), fieldNameMap, validators)
}

func ClusterReconfigureParamsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["storage_capacity"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["storage_capacity"] = "StorageCapacity"
	fields["bias"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["bias"] = "Bias"
	fields["num_hosts"] = bindings.NewIntegerType()
	fieldNameMap["num_hosts"] = "NumHosts"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.cluster_reconfigure_params", fields, reflect.TypeOf(ClusterReconfigureParams{}), fieldNameMap, validators)
}

func ComputeGatewayTemplateBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["public_ip"] = bindings.NewOptionalType(bindings.NewReferenceType(SddcPublicIpBindingType))
	fieldNameMap["public_ip"] = "PublicIp"
	fields["primary_dns"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["primary_dns"] = "PrimaryDns"
	fields["secondary_dns"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["secondary_dns"] = "SecondaryDns"
	fields["firewall_rules"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(FirewallRuleBindingType), reflect.TypeOf([]FirewallRule{})))
	fieldNameMap["firewall_rules"] = "FirewallRules"
	fields["vpns"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(VpnBindingType), reflect.TypeOf([]Vpn{})))
	fieldNameMap["vpns"] = "Vpns"
	fields["logical_networks"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(LogicalNetworkBindingType), reflect.TypeOf([]LogicalNetwork{})))
	fieldNameMap["logical_networks"] = "LogicalNetworks"
	fields["nat_rules"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(NatRuleBindingType), reflect.TypeOf([]NatRule{})))
	fieldNameMap["nat_rules"] = "NatRules"
	fields["l2_vpn"] = bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.REST))
	fieldNameMap["l2_vpn"] = "L2Vpn"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.compute_gateway_template", fields, reflect.TypeOf(ComputeGatewayTemplate{}), fieldNameMap, validators)
}

func ConfigSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["expiry_in_days"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["expiry_in_days"] = "ExpiryInDays"
	fields["availability"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewListType(bindings.NewReferenceType(InstanceTypeConfigBindingType), reflect.TypeOf([]InstanceTypeConfig{})),reflect.TypeOf(map[string][]InstanceTypeConfig{})))
	fieldNameMap["availability"] = "Availability"
	fields["sddc_sizes"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["sddc_sizes"] = "SddcSizes"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.config_spec", fields, reflect.TypeOf(ConfigSpec{}), fieldNameMap, validators)
}

func ConnectivityAgentValidationBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["source"] = "Source"
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["type"] = "Type_"
	fields["ports"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["ports"] = "Ports"
	fields["path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["path"] = "Path"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.connectivity_agent_validation", fields, reflect.TypeOf(ConnectivityAgentValidation{}), fieldNameMap, validators)
}

func ConnectivityValidationGroupBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["sub_groups"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ConnectivityValidationSubGroupBindingType), reflect.TypeOf([]ConnectivityValidationSubGroup{})))
	fieldNameMap["sub_groups"] = "SubGroups"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.connectivity_validation_group", fields, reflect.TypeOf(ConnectivityValidationGroup{}), fieldNameMap, validators)
}

func ConnectivityValidationGroupsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["groups"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ConnectivityValidationGroupBindingType), reflect.TypeOf([]ConnectivityValidationGroup{})))
	fieldNameMap["groups"] = "Groups"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.connectivity_validation_groups", fields, reflect.TypeOf(ConnectivityValidationGroups{}), fieldNameMap, validators)
}

func ConnectivityValidationInputBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["value"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["value"] = "Value"
	fields["label"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["label"] = "Label"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.connectivity_validation_input", fields, reflect.TypeOf(ConnectivityValidationInput{}), fieldNameMap, validators)
}

func ConnectivityValidationSubGroupBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["inputs"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ConnectivityValidationInputBindingType), reflect.TypeOf([]ConnectivityValidationInput{})))
	fieldNameMap["inputs"] = "Inputs"
	fields["tests"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ConnectivityAgentValidationBindingType), reflect.TypeOf([]ConnectivityAgentValidation{})))
	fieldNameMap["tests"] = "Tests"
	fields["label"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["label"] = "Label"
	fields["help"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["help"] = "Help"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.connectivity_validation_sub_group", fields, reflect.TypeOf(ConnectivityValidationSubGroup{}), fieldNameMap, validators)
}

func CustomerEniInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["secondary_ip_addresses"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["secondary_ip_addresses"] = "SecondaryIpAddresses"
	fields["eni_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["eni_id"] = "EniId"
	fields["primary_ip_address"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["primary_ip_address"] = "PrimaryIpAddress"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.customer_eni_info", fields, reflect.TypeOf(CustomerEniInfo{}), fieldNameMap, validators)
}

func EbsBackedVsanConfigBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["instance_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["instance_type"] = "InstanceType"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.ebs_backed_vsan_config", fields, reflect.TypeOf(EbsBackedVsanConfig{}), fieldNameMap, validators)
}

func EniInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["subnet_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["subnet_id"] = "SubnetId"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["security_group_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["security_group_id"] = "SecurityGroupId"
	fields["private_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["private_ip"] = "PrivateIp"
	fields["mac_address"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["mac_address"] = "MacAddress"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.eni_info", fields, reflect.TypeOf(EniInfo{}), fieldNameMap, validators)
}

func EntityCapacityBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["storage_capacity_gib"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["storage_capacity_gib"] = "StorageCapacityGib"
	fields["memory_capacity_gib"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["memory_capacity_gib"] = "MemoryCapacityGib"
	fields["total_number_of_cores"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["total_number_of_cores"] = "TotalNumberOfCores"
	fields["number_of_ssds"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["number_of_ssds"] = "NumberOfSsds"
	fields["cpu_capacity_ghz"] = bindings.NewOptionalType(bindings.NewDoubleType())
	fieldNameMap["cpu_capacity_ghz"] = "CpuCapacityGhz"
	fields["number_of_sockets"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["number_of_sockets"] = "NumberOfSockets"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.entity_capacity", fields, reflect.TypeOf(EntityCapacity{}), fieldNameMap, validators)
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
	return bindings.NewStructType("com.vmware.vmc.model.error_response", fields, reflect.TypeOf(ErrorResponse{}), fieldNameMap, validators)
}

func EsxConfigBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["availability_zone"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["availability_zone"] = "AvailabilityZone"
	fields["strict_placement"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["strict_placement"] = "StrictPlacement"
	fields["cluster_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cluster_id"] = "ClusterId"
	fields["esxs"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["esxs"] = "Esxs"
	fields["num_hosts"] = bindings.NewIntegerType()
	fieldNameMap["num_hosts"] = "NumHosts"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.esx_config", fields, reflect.TypeOf(EsxConfig{}), fieldNameMap, validators)
}

func EsxHostBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["availability_zone"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["availability_zone"] = "AvailabilityZone"
	fields["esx_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["esx_id"] = "EsxId"
	fields["hostname"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["hostname"] = "Hostname"
	fields["provider"] = bindings.NewStringType()
	fieldNameMap["provider"] = "Provider"
	fields["instance_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["instance_type"] = "InstanceType"
	fields["mac_address"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["mac_address"] = "MacAddress"
	fields["custom_properties"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewStringType(),reflect.TypeOf(map[string]string{})))
	fieldNameMap["custom_properties"] = "CustomProperties"
	fields["esx_state"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["esx_state"] = "EsxState"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.esx_host", fields, reflect.TypeOf(EsxHost{}), fieldNameMap, validators)
}

func EsxHostInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["instance_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["instance_type"] = "InstanceType"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.esx_host_info", fields, reflect.TypeOf(EsxHostInfo{}), fieldNameMap, validators)
}

func FirewallRuleBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["rule_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["rule_type"] = "RuleType"
	fields["application_ids"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["application_ids"] = "ApplicationIds"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["rule_interface"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["rule_interface"] = "RuleInterface"
	fields["destination"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["destination"] = "Destination"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["destination_scope"] = bindings.NewOptionalType(bindings.NewReferenceType(FirewallRuleScopeBindingType))
	fieldNameMap["destination_scope"] = "DestinationScope"
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["source"] = "Source"
	fields["source_scope"] = bindings.NewOptionalType(bindings.NewReferenceType(FirewallRuleScopeBindingType))
	fieldNameMap["source_scope"] = "SourceScope"
	fields["services"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(FirewallServiceBindingType), reflect.TypeOf([]FirewallService{})))
	fieldNameMap["services"] = "Services"
	fields["action"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["action"] = "Action"
	fields["revision"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["revision"] = "Revision"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.firewall_rule", fields, reflect.TypeOf(FirewallRule{}), fieldNameMap, validators)
}

func FirewallRuleScopeBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["grouping_object_ids"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["grouping_object_ids"] = "GroupingObjectIds"
	fields["vnic_group_ids"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["vnic_group_ids"] = "VnicGroupIds"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.firewall_rule_scope", fields, reflect.TypeOf(FirewallRuleScope{}), fieldNameMap, validators)
}

func FirewallServiceBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["protocol"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["protocol"] = "Protocol"
	fields["ports"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["ports"] = "Ports"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.firewall_service", fields, reflect.TypeOf(FirewallService{}), fieldNameMap, validators)
}

func GatewayTemplateBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["public_ip"] = bindings.NewOptionalType(bindings.NewReferenceType(SddcPublicIpBindingType))
	fieldNameMap["public_ip"] = "PublicIp"
	fields["primary_dns"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["primary_dns"] = "PrimaryDns"
	fields["secondary_dns"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["secondary_dns"] = "SecondaryDns"
	fields["firewall_rules"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(FirewallRuleBindingType), reflect.TypeOf([]FirewallRule{})))
	fieldNameMap["firewall_rules"] = "FirewallRules"
	fields["vpns"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(VpnBindingType), reflect.TypeOf([]Vpn{})))
	fieldNameMap["vpns"] = "Vpns"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.gateway_template", fields, reflect.TypeOf(GatewayTemplate{}), fieldNameMap, validators)
}

func GlcmBundleBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["s3Bucket"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["s3Bucket"] = "S3Bucket"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.glcm_bundle", fields, reflect.TypeOf(GlcmBundle{}), fieldNameMap, validators)
}

func InstanceTypeConfigBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["instance_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["instance_type"] = "InstanceType"
	fields["hosts"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewIntegerType(), reflect.TypeOf([]int64{})))
	fieldNameMap["hosts"] = "Hosts"
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["entity_capacity"] = bindings.NewOptionalType(bindings.NewReferenceType(EntityCapacityBindingType))
	fieldNameMap["entity_capacity"] = "EntityCapacity"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.instance_type_config", fields, reflect.TypeOf(InstanceTypeConfig{}), fieldNameMap, validators)
}

func KmsVpcEndpointBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vpc_endpoint_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vpc_endpoint_id"] = "VpcEndpointId"
	fields["network_interface_ids"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["network_interface_ids"] = "NetworkInterfaceIds"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.kms_vpc_endpoint", fields, reflect.TypeOf(KmsVpcEndpoint{}), fieldNameMap, validators)
}

func L2VpnBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["enabled"] = "Enabled"
	fields["sites"] = bindings.NewListType(bindings.NewReferenceType(SiteBindingType), reflect.TypeOf([]Site{}))
	fieldNameMap["sites"] = "Sites"
	fields["listener_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["listener_ip"] = "ListenerIp"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.l2_vpn", fields, reflect.TypeOf(L2Vpn{}), fieldNameMap, validators)
}

func LogicalNetworkBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["subnet_cidr"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["subnet_cidr"] = "SubnetCidr"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["gatewayIp"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["gatewayIp"] = "GatewayIp"
	fields["dhcp_enabled"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["dhcp_enabled"] = "DhcpEnabled"
	fields["dhcp_ip_range"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["dhcp_ip_range"] = "DhcpIpRange"
	fields["tunnel_id"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["tunnel_id"] = "TunnelId"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["network_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["network_type"] = "NetworkType"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.logical_network", fields, reflect.TypeOf(LogicalNetwork{}), fieldNameMap, validators)
}

func MaintenanceWindowBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["day_of_week"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["day_of_week"] = "DayOfWeek"
	fields["hour_of_day"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["hour_of_day"] = "HourOfDay"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.maintenance_window", fields, reflect.TypeOf(MaintenanceWindow{}), fieldNameMap, validators)
}

func MaintenanceWindowEntryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["in_maintenance_window"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["in_maintenance_window"] = "InMaintenanceWindow"
	fields["reservation_schedule"] = bindings.NewOptionalType(bindings.NewReferenceType(ReservationScheduleBindingType))
	fieldNameMap["reservation_schedule"] = "ReservationSchedule"
	fields["reservation_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["reservation_id"] = "ReservationId"
	fields["in_maintenance_mode"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["in_maintenance_mode"] = "InMaintenanceMode"
	fields["sddc_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sddc_id"] = "SddcId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.maintenance_window_entry", fields, reflect.TypeOf(MaintenanceWindowEntry{}), fieldNameMap, validators)
}

func MaintenanceWindowGetBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["day_of_week"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["day_of_week"] = "DayOfWeek"
	fields["hour_of_day"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["hour_of_day"] = "HourOfDay"
	fields["duration_min"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["duration_min"] = "DurationMin"
	fields["version"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["version"] = "Version"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.maintenance_window_get", fields, reflect.TypeOf(MaintenanceWindowGet{}), fieldNameMap, validators)
}

func ManagementGatewayTemplateBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["public_ip"] = bindings.NewOptionalType(bindings.NewReferenceType(SddcPublicIpBindingType))
	fieldNameMap["public_ip"] = "PublicIp"
	fields["primary_dns"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["primary_dns"] = "PrimaryDns"
	fields["secondary_dns"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["secondary_dns"] = "SecondaryDns"
	fields["firewall_rules"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(FirewallRuleBindingType), reflect.TypeOf([]FirewallRule{})))
	fieldNameMap["firewall_rules"] = "FirewallRules"
	fields["vpns"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(VpnBindingType), reflect.TypeOf([]Vpn{})))
	fieldNameMap["vpns"] = "Vpns"
	fields["subnet_cidr"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["subnet_cidr"] = "SubnetCidr"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.management_gateway_template", fields, reflect.TypeOf(ManagementGatewayTemplate{}), fieldNameMap, validators)
}

func MapZonesRequestBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["connected_account_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["connected_account_id"] = "ConnectedAccountId"
	fields["org_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fields["petronas_regions_to_map"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["petronas_regions_to_map"] = "PetronasRegionsToMap"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.map_zones_request", fields, reflect.TypeOf(MapZonesRequest{}), fieldNameMap, validators)
}

func MetadataBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["timestamp"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["timestamp"] = "Timestamp"
	fields["cycle_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cycle_id"] = "CycleId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.metadata", fields, reflect.TypeOf(Metadata{}), fieldNameMap, validators)
}

func MsftLicensingConfigBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["mssql_licensing"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["mssql_licensing"] = "MssqlLicensing"
	fields["windows_licensing"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["windows_licensing"] = "WindowsLicensing"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.msft_licensing_config", fields, reflect.TypeOf(MsftLicensingConfig{}), fieldNameMap, validators)
}

func NatRuleBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["rule_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["rule_type"] = "RuleType"
	fields["protocol"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["protocol"] = "Protocol"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["internal_ports"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["internal_ports"] = "InternalPorts"
	fields["public_ports"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["public_ports"] = "PublicPorts"
	fields["public_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["public_ip"] = "PublicIp"
	fields["internal_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["internal_ip"] = "InternalIp"
	fields["action"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["action"] = "Action"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["revision"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["revision"] = "Revision"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.nat_rule", fields, reflect.TypeOf(NatRule{}), fieldNameMap, validators)
}

func NetworkTemplateBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["management_gateway_templates"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ManagementGatewayTemplateBindingType), reflect.TypeOf([]ManagementGatewayTemplate{})))
	fieldNameMap["management_gateway_templates"] = "ManagementGatewayTemplates"
	fields["compute_gateway_templates"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ComputeGatewayTemplateBindingType), reflect.TypeOf([]ComputeGatewayTemplate{})))
	fieldNameMap["compute_gateway_templates"] = "ComputeGatewayTemplates"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.network_template", fields, reflect.TypeOf(NetworkTemplate{}), fieldNameMap, validators)
}

func NewCredentialsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["username"] = bindings.NewStringType()
	fieldNameMap["username"] = "Username"
	fields["password"] = bindings.NewStringType()
	fieldNameMap["password"] = "Password"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.new_credentials", fields, reflect.TypeOf(NewCredentials{}), fieldNameMap, validators)
}

func NsxtAddonsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["enable_nsx_advanced_addon"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["enable_nsx_advanced_addon"] = "EnableNsxAdvancedAddon"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.nsxt_addons", fields, reflect.TypeOf(NsxtAddons{}), fieldNameMap, validators)
}

func OfferInstancesHolderBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["on_demand"] = bindings.NewReferenceType(OnDemandOfferInstanceBindingType)
	fieldNameMap["on_demand"] = "OnDemand"
	fields["offers"] = bindings.NewListType(bindings.NewReferenceType(TermOfferInstanceBindingType), reflect.TypeOf([]TermOfferInstance{}))
	fieldNameMap["offers"] = "Offers"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.offer_instances_holder", fields, reflect.TypeOf(OfferInstancesHolder{}), fieldNameMap, validators)
}

func OnDemandOfferInstanceBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["product"] = bindings.NewStringType()
	fieldNameMap["product"] = "Product"
	fields["product_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["product_type"] = "ProductType"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["currency"] = bindings.NewStringType()
	fieldNameMap["currency"] = "Currency"
	fields["region"] = bindings.NewStringType()
	fieldNameMap["region"] = "Region"
	fields["unit_price"] = bindings.NewStringType()
	fieldNameMap["unit_price"] = "UnitPrice"
	fields["monthly_cost"] = bindings.NewStringType()
	fieldNameMap["monthly_cost"] = "MonthlyCost"
	fields["version"] = bindings.NewStringType()
	fieldNameMap["version"] = "Version"
	fields["type"] = bindings.NewStringType()
	fieldNameMap["type"] = "Type_"
	fields["description"] = bindings.NewStringType()
	fieldNameMap["description"] = "Description"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.on_demand_offer_instance", fields, reflect.TypeOf(OnDemandOfferInstance{}), fieldNameMap, validators)
}

func OrgPropertiesBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["values"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewStringType(),reflect.TypeOf(map[string]string{})))
	fieldNameMap["values"] = "Values"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.org_properties", fields, reflect.TypeOf(OrgProperties{}), fieldNameMap, validators)
}

func OrgSellerInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["seller_account_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["seller_account_id"] = "SellerAccountId"
	fields["seller"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["seller"] = "Seller"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.org_seller_info", fields, reflect.TypeOf(OrgSellerInfo{}), fieldNameMap, validators)
}

func OrganizationBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["updated"] = bindings.NewDateTimeType()
	fieldNameMap["updated"] = "Updated"
	fields["user_id"] = bindings.NewStringType()
	fieldNameMap["user_id"] = "UserId"
	fields["updated_by_user_id"] = bindings.NewStringType()
	fieldNameMap["updated_by_user_id"] = "UpdatedByUserId"
	fields["created"] = bindings.NewDateTimeType()
	fieldNameMap["created"] = "Created"
	fields["version"] = bindings.NewIntegerType()
	fieldNameMap["version"] = "Version"
	fields["updated_by_user_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["updated_by_user_name"] = "UpdatedByUserName"
	fields["user_name"] = bindings.NewStringType()
	fieldNameMap["user_name"] = "UserName"
	fields["id"] = bindings.NewStringType()
	fieldNameMap["id"] = "Id"
	fields["org_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org_type"] = "OrgType"
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["org_seller_info"] = bindings.NewOptionalType(bindings.NewReferenceType(OrgSellerInfoBindingType))
	fieldNameMap["org_seller_info"] = "OrgSellerInfo"
	fields["project_state"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["project_state"] = "ProjectState"
	fields["properties"] = bindings.NewOptionalType(bindings.NewReferenceType(OrgPropertiesBindingType))
	fieldNameMap["properties"] = "Properties"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.organization", fields, reflect.TypeOf(Organization{}), fieldNameMap, validators)
}

func PaymentMethodInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["type"] = "Type_"
	fields["default_flag"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["default_flag"] = "DefaultFlag"
	fields["payment_method_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["payment_method_id"] = "PaymentMethodId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.payment_method_info", fields, reflect.TypeOf(PaymentMethodInfo{}), fieldNameMap, validators)
}

func PopAgentXeniConnectionBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["default_subnet_route"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["default_subnet_route"] = "DefaultSubnetRoute"
	fields["eni_info"] = bindings.NewOptionalType(bindings.NewReferenceType(EniInfoBindingType))
	fieldNameMap["eni_info"] = "EniInfo"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.pop_agent_xeni_connection", fields, reflect.TypeOf(PopAgentXeniConnection{}), fieldNameMap, validators)
}

func PopAmiInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["instance_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["instance_type"] = "InstanceType"
	fields["region"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["region"] = "Region"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["type"] = "Type_"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.pop_ami_info", fields, reflect.TypeOf(PopAmiInfo{}), fieldNameMap, validators)
}

func PopInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ami_infos"] = bindings.NewMapType(bindings.NewStringType(), bindings.NewReferenceType(PopAmiInfoBindingType),reflect.TypeOf(map[string]PopAmiInfo{}))
	fieldNameMap["ami_infos"] = "AmiInfos"
	fields["created_at"] = bindings.NewOptionalType(bindings.NewDateTimeType())
	fieldNameMap["created_at"] = "CreatedAt"
	fields["service_infos"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewReferenceType(PopServiceInfoBindingType),reflect.TypeOf(map[string]PopServiceInfo{})))
	fieldNameMap["service_infos"] = "ServiceInfos"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["manifest_version"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["manifest_version"] = "ManifestVersion"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.pop_info", fields, reflect.TypeOf(PopInfo{}), fieldNameMap, validators)
}

func PopServiceInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cln"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cln"] = "Cln"
	fields["version"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["version"] = "Version"
	fields["build"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["build"] = "Build"
	fields["service"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["service"] = "Service"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.pop_service_info", fields, reflect.TypeOf(PopServiceInfo{}), fieldNameMap, validators)
}

func ProvisionSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["provider"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewReferenceType(SddcConfigSpecBindingType),reflect.TypeOf(map[string]SddcConfigSpec{})))
	fieldNameMap["provider"] = "Provider"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.provision_spec", fields, reflect.TypeOf(ProvisionSpec{}), fieldNameMap, validators)
}

func RequestDetailBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["aws_quota_request_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["aws_quota_request_id"] = "AwsQuotaRequestId"
	fields["detail_status"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["detail_status"] = "DetailStatus"
	fields["resolved_at"] = bindings.NewOptionalType(bindings.NewDateTimeType())
	fieldNameMap["resolved_at"] = "ResolvedAt"
	fields["desired_value"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["desired_value"] = "DesiredValue"
	fields["aws_support_case_status"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["aws_support_case_status"] = "AwsSupportCaseStatus"
	fields["aws_support_case_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["aws_support_case_id"] = "AwsSupportCaseId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.request_detail", fields, reflect.TypeOf(RequestDetail{}), fieldNameMap, validators)
}

func ReservationBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["duration"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["duration"] = "Duration"
	fields["rid"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["rid"] = "Rid"
	fields["create_time"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["create_time"] = "CreateTime"
	fields["start_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
	fieldNameMap["start_time"] = "StartTime"
	fields["metadata"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewStringType(),reflect.TypeOf(map[string]string{})))
	fieldNameMap["metadata"] = "Metadata"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.reservation", fields, reflect.TypeOf(Reservation{}), fieldNameMap, validators)
}

func ReservationInMwBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["rid"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["rid"] = "Rid"
	fields["week_of"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["week_of"] = "WeekOf"
	fields["create_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
	fieldNameMap["create_time"] = "CreateTime"
	fields["metadata"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewStringType(),reflect.TypeOf(map[string]string{})))
	fieldNameMap["metadata"] = "Metadata"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.reservation_in_mw", fields, reflect.TypeOf(ReservationInMw{}), fieldNameMap, validators)
}

func ReservationScheduleBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["day_of_week"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["day_of_week"] = "DayOfWeek"
	fields["hour_of_day"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["hour_of_day"] = "HourOfDay"
	fields["duration_min"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["duration_min"] = "DurationMin"
	fields["version"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["version"] = "Version"
	fields["reservations"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ReservationBindingType), reflect.TypeOf([]Reservation{})))
	fieldNameMap["reservations"] = "Reservations"
	fields["reservations_mw"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ReservationInMwBindingType), reflect.TypeOf([]ReservationInMw{})))
	fieldNameMap["reservations_mw"] = "ReservationsMw"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.reservation_schedule", fields, reflect.TypeOf(ReservationSchedule{}), fieldNameMap, validators)
}

func ReservationWindowBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["reservation_state"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["reservation_state"] = "ReservationState"
	fields["emergency"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["emergency"] = "Emergency"
	fields["maintenance_properties"] = bindings.NewOptionalType(bindings.NewReferenceType(ReservationWindowMaintenancePropertiesBindingType))
	fieldNameMap["maintenance_properties"] = "MaintenanceProperties"
	fields["reserve_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["reserve_id"] = "ReserveId"
	fields["start_hour"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["start_hour"] = "StartHour"
	fields["sddc_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sddc_id"] = "SddcId"
	fields["manifest_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["manifest_id"] = "ManifestId"
	fields["duration_hours"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["duration_hours"] = "DurationHours"
	fields["start_date"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["start_date"] = "StartDate"
	fields["metadata"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewStringType(),reflect.TypeOf(map[string]string{})))
	fieldNameMap["metadata"] = "Metadata"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.reservation_window", fields, reflect.TypeOf(ReservationWindow{}), fieldNameMap, validators)
}

func ReservationWindowMaintenancePropertiesBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["status"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["status"] = "Status"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.reservation_window_maintenance_properties", fields, reflect.TypeOf(ReservationWindowMaintenanceProperties{}), fieldNameMap, validators)
}

func RouteTableInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.route_table_info", fields, reflect.TypeOf(RouteTableInfo{}), fieldNameMap, validators)
}

func SddcBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["updated"] = bindings.NewDateTimeType()
	fieldNameMap["updated"] = "Updated"
	fields["user_id"] = bindings.NewStringType()
	fieldNameMap["user_id"] = "UserId"
	fields["updated_by_user_id"] = bindings.NewStringType()
	fieldNameMap["updated_by_user_id"] = "UpdatedByUserId"
	fields["created"] = bindings.NewDateTimeType()
	fieldNameMap["created"] = "Created"
	fields["version"] = bindings.NewIntegerType()
	fieldNameMap["version"] = "Version"
	fields["updated_by_user_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["updated_by_user_name"] = "UpdatedByUserName"
	fields["user_name"] = bindings.NewStringType()
	fieldNameMap["user_name"] = "UserName"
	fields["id"] = bindings.NewStringType()
	fieldNameMap["id"] = "Id"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["sddc_state"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sddc_state"] = "SddcState"
	fields["expiration_date"] = bindings.NewOptionalType(bindings.NewDateTimeType())
	fieldNameMap["expiration_date"] = "ExpirationDate"
	fields["org_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fields["sddc_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sddc_type"] = "SddcType"
	fields["provider"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["provider"] = "Provider"
	fields["account_link_state"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["account_link_state"] = "AccountLinkState"
	fields["sddc_access_state"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sddc_access_state"] = "SddcAccessState"
	fields["resource_config"] = bindings.NewOptionalType(bindings.NewReferenceType(AwsSddcResourceConfigBindingType))
	fieldNameMap["resource_config"] = "ResourceConfig"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.sddc", fields, reflect.TypeOf(Sddc{}), fieldNameMap, validators)
}

func SddcConfigBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["msft_license_config"] = bindings.NewOptionalType(bindings.NewReferenceType(MsftLicensingConfigBindingType))
	fieldNameMap["msft_license_config"] = "MsftLicenseConfig"
	fields["vpc_cidr"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vpc_cidr"] = "VpcCidr"
	fields["host_instance_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vmc.model.host_instance_types", reflect.TypeOf(HostInstanceTypes(HostInstanceTypes_I3_METAL))))
	fieldNameMap["host_instance_type"] = "HostInstanceType"
	fields["skip_creating_vxlan"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["skip_creating_vxlan"] = "SkipCreatingVxlan"
	fields["vxlan_subnet"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vxlan_subnet"] = "VxlanSubnet"
	fields["size"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["size"] = "Size"
	fields["storage_capacity"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["storage_capacity"] = "StorageCapacity"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["account_link_sddc_config"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(AccountLinkSddcConfigBindingType), reflect.TypeOf([]AccountLinkSddcConfig{})))
	fieldNameMap["account_link_sddc_config"] = "AccountLinkSddcConfig"
	fields["sddc_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sddc_id"] = "SddcId"
	fields["num_hosts"] = bindings.NewIntegerType()
	fieldNameMap["num_hosts"] = "NumHosts"
	fields["sddc_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sddc_type"] = "SddcType"
	fields["account_link_config"] = bindings.NewOptionalType(bindings.NewReferenceType(AccountLinkConfigBindingType))
	fieldNameMap["account_link_config"] = "AccountLinkConfig"
	fields["provider"] = bindings.NewStringType()
	fieldNameMap["provider"] = "Provider"
	fields["sso_domain"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sso_domain"] = "SsoDomain"
	fields["sddc_template_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sddc_template_id"] = "SddcTemplateId"
	fields["deployment_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["deployment_type"] = "DeploymentType"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.sddc_config", fields, reflect.TypeOf(SddcConfig{}), fieldNameMap, validators)
}

func SddcConfigSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["sddc_type_config_spec"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewReferenceType(ConfigSpecBindingType),reflect.TypeOf(map[string]ConfigSpec{})))
	fieldNameMap["sddc_type_config_spec"] = "SddcTypeConfigSpec"
	fields["region_display_names"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewStringType(),reflect.TypeOf(map[string]string{})))
	fieldNameMap["region_display_names"] = "RegionDisplayNames"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.sddc_config_spec", fields, reflect.TypeOf(SddcConfigSpec{}), fieldNameMap, validators)
}

func SddcIdBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["sddc_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sddc_id"] = "SddcId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.sddc_id", fields, reflect.TypeOf(SddcId{}), fieldNameMap, validators)
}

func SddcLinkConfigBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["customer_subnet_ids"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["customer_subnet_ids"] = "CustomerSubnetIds"
	fields["connected_account_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["connected_account_id"] = "ConnectedAccountId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.sddc_link_config", fields, reflect.TypeOf(SddcLinkConfig{}), fieldNameMap, validators)
}

func SddcManifestBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vmc_version"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vmc_version"] = "VmcVersion"
	fields["glcm_bundle"] = bindings.NewOptionalType(bindings.NewReferenceType(GlcmBundleBindingType))
	fieldNameMap["glcm_bundle"] = "GlcmBundle"
	fields["pop_info"] = bindings.NewOptionalType(bindings.NewReferenceType(PopInfoBindingType))
	fieldNameMap["pop_info"] = "PopInfo"
	fields["vmc_internal_version"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vmc_internal_version"] = "VmcInternalVersion"
	fields["ebs_backed_vsan_config"] = bindings.NewOptionalType(bindings.NewReferenceType(EbsBackedVsanConfigBindingType))
	fieldNameMap["ebs_backed_vsan_config"] = "EbsBackedVsanConfig"
	fields["vsan_witness_ami"] = bindings.NewOptionalType(bindings.NewReferenceType(AmiInfoBindingType))
	fieldNameMap["vsan_witness_ami"] = "VsanWitnessAmi"
	fields["esx_ami"] = bindings.NewOptionalType(bindings.NewReferenceType(AmiInfoBindingType))
	fieldNameMap["esx_ami"] = "EsxAmi"
	fields["esx_nsxt_ami"] = bindings.NewOptionalType(bindings.NewReferenceType(AmiInfoBindingType))
	fieldNameMap["esx_nsxt_ami"] = "EsxNsxtAmi"
	fields["metadata"] = bindings.NewOptionalType(bindings.NewReferenceType(MetadataBindingType))
	fieldNameMap["metadata"] = "Metadata"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.sddc_manifest", fields, reflect.TypeOf(SddcManifest{}), fieldNameMap, validators)
}

func SddcPatchRequestBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.sddc_patch_request", fields, reflect.TypeOf(SddcPatchRequest{}), fieldNameMap, validators)
}

func SddcPublicIpBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["public_ip"] = bindings.NewStringType()
	fieldNameMap["public_ip"] = "PublicIp"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["allocation_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["allocation_id"] = "AllocationId"
	fields["dnat_rule_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["dnat_rule_id"] = "DnatRuleId"
	fields["associated_private_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["associated_private_ip"] = "AssociatedPrivateIp"
	fields["snat_rule_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["snat_rule_id"] = "SnatRuleId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.sddc_public_ip", fields, reflect.TypeOf(SddcPublicIp{}), fieldNameMap, validators)
}

func SddcResourceConfigBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["mgmt_appliance_network_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["mgmt_appliance_network_name"] = "MgmtApplianceNetworkName"
	fields["nsxt"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["nsxt"] = "Nsxt"
	fields["mgw_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["mgw_id"] = "MgwId"
	fields["nsx_mgr_url"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["nsx_mgr_url"] = "NsxMgrUrl"
	fields["psc_management_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["psc_management_ip"] = "PscManagementIp"
	fields["psc_url"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["psc_url"] = "PscUrl"
	fields["cgws"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["cgws"] = "Cgws"
	fields["availability_zones"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["availability_zones"] = "AvailabilityZones"
	fields["management_ds"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["management_ds"] = "ManagementDs"
	fields["nsx_api_public_endpoint_url"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["nsx_api_public_endpoint_url"] = "NsxApiPublicEndpointUrl"
	fields["custom_properties"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewStringType(),reflect.TypeOf(map[string]string{})))
	fieldNameMap["custom_properties"] = "CustomProperties"
	fields["cloud_password"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cloud_password"] = "CloudPassword"
	fields["provider"] = bindings.NewStringType()
	fieldNameMap["provider"] = "Provider"
	fields["clusters"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ClusterBindingType), reflect.TypeOf([]Cluster{})))
	fieldNameMap["clusters"] = "Clusters"
	fields["vc_management_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vc_management_ip"] = "VcManagementIp"
	fields["sddc_networks"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["sddc_networks"] = "SddcNetworks"
	fields["cloud_username"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cloud_username"] = "CloudUsername"
	fields["esx_hosts"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(AwsEsxHostBindingType), reflect.TypeOf([]AwsEsxHost{})))
	fieldNameMap["esx_hosts"] = "EsxHosts"
	fields["nsx_mgr_management_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["nsx_mgr_management_ip"] = "NsxMgrManagementIp"
	fields["vc_instance_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vc_instance_id"] = "VcInstanceId"
	fields["esx_cluster_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["esx_cluster_id"] = "EsxClusterId"
	fields["vc_public_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vc_public_ip"] = "VcPublicIp"
	fields["skip_creating_vxlan"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["skip_creating_vxlan"] = "SkipCreatingVxlan"
	fields["vc_url"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vc_url"] = "VcUrl"
	fields["sddc_manifest"] = bindings.NewOptionalType(bindings.NewReferenceType(SddcManifestBindingType))
	fieldNameMap["sddc_manifest"] = "SddcManifest"
	fields["vxlan_subnet"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vxlan_subnet"] = "VxlanSubnet"
	fields["cloud_user_group"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cloud_user_group"] = "CloudUserGroup"
	fields["management_rp"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["management_rp"] = "ManagementRp"
	fields["region"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["region"] = "Region"
	fields["witness_availability_zone"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["witness_availability_zone"] = "WitnessAvailabilityZone"
	fields["sddc_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sddc_id"] = "SddcId"
	fields["pop_agent_xeni_connection"] = bindings.NewOptionalType(bindings.NewReferenceType(PopAgentXeniConnectionBindingType))
	fieldNameMap["pop_agent_xeni_connection"] = "PopAgentXeniConnection"
	fields["sddc_size"] = bindings.NewOptionalType(bindings.NewReferenceType(SddcSizeBindingType))
	fieldNameMap["sddc_size"] = "SddcSize"
	fields["nsx_controller_ips"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["nsx_controller_ips"] = "NsxControllerIps"
	fields["esx_host_subnet"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["esx_host_subnet"] = "EsxHostSubnet"
	fields["sso_domain"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sso_domain"] = "SsoDomain"
	fields["deployment_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["deployment_type"] = "DeploymentType"
	fields["msft_license_config"] = bindings.NewOptionalType(bindings.NewReferenceType(MsftLicensingConfigBindingType))
	fieldNameMap["msft_license_config"] = "MsftLicenseConfig"
	fields["nsxt_addons"] = bindings.NewOptionalType(bindings.NewReferenceType(NsxtAddonsBindingType))
	fieldNameMap["nsxt_addons"] = "NsxtAddons"
	fields["dns_with_management_vm_private_ip"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["dns_with_management_vm_private_ip"] = "DnsWithManagementVmPrivateIp"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.sddc_resource_config", fields, reflect.TypeOf(SddcResourceConfig{}), fieldNameMap, validators)
}

func SddcSizeBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vc_size"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vc_size"] = "VcSize"
	fields["nsx_size"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["nsx_size"] = "NsxSize"
	fields["size"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["size"] = "Size"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.sddc_size", fields, reflect.TypeOf(SddcSize{}), fieldNameMap, validators)
}

func SddcStateRequestBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["sddcs"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["sddcs"] = "Sddcs"
	fields["states"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["states"] = "States"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.sddc_state_request", fields, reflect.TypeOf(SddcStateRequest{}), fieldNameMap, validators)
}

func SddcTemplateBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["updated"] = bindings.NewDateTimeType()
	fieldNameMap["updated"] = "Updated"
	fields["user_id"] = bindings.NewStringType()
	fieldNameMap["user_id"] = "UserId"
	fields["updated_by_user_id"] = bindings.NewStringType()
	fieldNameMap["updated_by_user_id"] = "UpdatedByUserId"
	fields["created"] = bindings.NewDateTimeType()
	fieldNameMap["created"] = "Created"
	fields["version"] = bindings.NewIntegerType()
	fieldNameMap["version"] = "Version"
	fields["updated_by_user_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["updated_by_user_name"] = "UpdatedByUserName"
	fields["user_name"] = bindings.NewStringType()
	fieldNameMap["user_name"] = "UserName"
	fields["id"] = bindings.NewStringType()
	fieldNameMap["id"] = "Id"
	fields["account_link_sddc_configs"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(AccountLinkSddcConfigBindingType), reflect.TypeOf([]AccountLinkSddcConfig{})))
	fieldNameMap["account_link_sddc_configs"] = "AccountLinkSddcConfigs"
	fields["state"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["state"] = "State"
	fields["network_template"] = bindings.NewOptionalType(bindings.NewReferenceType(NetworkTemplateBindingType))
	fieldNameMap["network_template"] = "NetworkTemplate"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["source_sddc_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["source_sddc_id"] = "SourceSddcId"
	fields["org_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fields["sddc"] = bindings.NewOptionalType(bindings.NewReferenceType(SddcBindingType))
	fieldNameMap["sddc"] = "Sddc"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.sddc_template", fields, reflect.TypeOf(SddcTemplate{}), fieldNameMap, validators)
}

func ServiceErrorBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["default_message"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["default_message"] = "DefaultMessage"
	fields["original_service"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["original_service"] = "OriginalService"
	fields["localized_message"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["localized_message"] = "LocalizedMessage"
	fields["original_service_error_code"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["original_service_error_code"] = "OriginalServiceErrorCode"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.service_error", fields, reflect.TypeOf(ServiceError{}), fieldNameMap, validators)
}

func ServiceQuotaRequestBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["updated"] = bindings.NewDateTimeType()
	fieldNameMap["updated"] = "Updated"
	fields["user_id"] = bindings.NewStringType()
	fieldNameMap["user_id"] = "UserId"
	fields["updated_by_user_id"] = bindings.NewStringType()
	fieldNameMap["updated_by_user_id"] = "UpdatedByUserId"
	fields["created"] = bindings.NewDateTimeType()
	fieldNameMap["created"] = "Created"
	fields["version"] = bindings.NewIntegerType()
	fieldNameMap["version"] = "Version"
	fields["updated_by_user_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["updated_by_user_name"] = "UpdatedByUserName"
	fields["user_name"] = bindings.NewStringType()
	fieldNameMap["user_name"] = "UserName"
	fields["id"] = bindings.NewStringType()
	fieldNameMap["id"] = "Id"
	fields["requester_email"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["requester_email"] = "RequesterEmail"
	fields["task_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["task_id"] = "TaskId"
	fields["region"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["region"] = "Region"
	fields["aws_account_number"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["aws_account_number"] = "AwsAccountNumber"
	fields["org_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fields["reason"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["reason"] = "Reason"
	fields["request_status"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["request_status"] = "RequestStatus"
	fields["request_details"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewReferenceType(RequestDetailBindingType),reflect.TypeOf(map[string]RequestDetail{})))
	fieldNameMap["request_details"] = "RequestDetails"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.service_quota_request", fields, reflect.TypeOf(ServiceQuotaRequest{}), fieldNameMap, validators)
}

func SiteBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["password"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["password"] = "Password"
	fields["user_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["user_id"] = "UserId"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["rx_bytes_on_local_subnet"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["rx_bytes_on_local_subnet"] = "RxBytesOnLocalSubnet"
	fields["secure_traffic"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["secure_traffic"] = "SecureTraffic"
	fields["established_date"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["established_date"] = "EstablishedDate"
	fields["failure_message"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["failure_message"] = "FailureMessage"
	fields["dropped_tx_packets"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["dropped_tx_packets"] = "DroppedTxPackets"
	fields["dropped_rx_packets"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["dropped_rx_packets"] = "DroppedRxPackets"
	fields["tunnel_status"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["tunnel_status"] = "TunnelStatus"
	fields["tx_bytes_from_local_subnet"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["tx_bytes_from_local_subnet"] = "TxBytesFromLocalSubnet"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.site", fields, reflect.TypeOf(Site{}), fieldNameMap, validators)
}

func SubnetBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["subnet_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["subnet_id"] = "SubnetId"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["route_tables"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(SubnetRouteTableInfoBindingType), reflect.TypeOf([]SubnetRouteTableInfo{})))
	fieldNameMap["route_tables"] = "RouteTables"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.subnet", fields, reflect.TypeOf(Subnet{}), fieldNameMap, validators)
}

func SubnetInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["compatible"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["compatible"] = "Compatible"
	fields["connected_account_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["connected_account_id"] = "ConnectedAccountId"
	fields["region_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["region_name"] = "RegionName"
	fields["availability_zone"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["availability_zone"] = "AvailabilityZone"
	fields["subnet_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["subnet_id"] = "SubnetId"
	fields["availability_zone_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["availability_zone_id"] = "AvailabilityZoneId"
	fields["subnet_cidr_block"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["subnet_cidr_block"] = "SubnetCidrBlock"
	fields["note"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["note"] = "Note"
	fields["vpc_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vpc_id"] = "VpcId"
	fields["vpc_cidr_block"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vpc_cidr_block"] = "VpcCidrBlock"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.subnet_info", fields, reflect.TypeOf(SubnetInfo{}), fieldNameMap, validators)
}

func SubnetRouteTableInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["subnet_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["subnet_id"] = "SubnetId"
	fields["association_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["association_id"] = "AssociationId"
	fields["routetable_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["routetable_id"] = "RoutetableId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.subnet_route_table_info", fields, reflect.TypeOf(SubnetRouteTableInfo{}), fieldNameMap, validators)
}

func SubscriptionDetailsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["status"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["status"] = "Status"
	fields["anniversary_billing_date"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["anniversary_billing_date"] = "AnniversaryBillingDate"
	fields["end_date"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["end_date"] = "EndDate"
	fields["billing_frequency"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["billing_frequency"] = "BillingFrequency"
	fields["auto_renewed_allowed"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["auto_renewed_allowed"] = "AutoRenewedAllowed"
	fields["commitment_term"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["commitment_term"] = "CommitmentTerm"
	fields["csp_subscription_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["csp_subscription_id"] = "CspSubscriptionId"
	fields["billing_subscription_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["billing_subscription_id"] = "BillingSubscriptionId"
	fields["offer_version"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["offer_version"] = "OfferVersion"
	fields["offer_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vmc.model.offer_type", reflect.TypeOf(OfferType(OfferType_TERM))))
	fieldNameMap["offer_type"] = "OfferType"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["product_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["product_id"] = "ProductId"
	fields["region"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["region"] = "Region"
	fields["product_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["product_name"] = "ProductName"
	fields["offer_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["offer_name"] = "OfferName"
	fields["commitment_term_uom"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["commitment_term_uom"] = "CommitmentTermUom"
	fields["start_date"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["start_date"] = "StartDate"
	fields["quantity"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["quantity"] = "Quantity"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.subscription_details", fields, reflect.TypeOf(SubscriptionDetails{}), fieldNameMap, validators)
}

func SubscriptionProductsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["product"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["product"] = "Product"
	fields["types"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["types"] = "Types"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.subscription_products", fields, reflect.TypeOf(SubscriptionProducts{}), fieldNameMap, validators)
}

func SubscriptionRequestBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["product"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["product"] = "Product"
	fields["product_type"] = bindings.NewStringType()
	fieldNameMap["product_type"] = "ProductType"
	fields["product_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["product_id"] = "ProductId"
	fields["billing_frequency"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["billing_frequency"] = "BillingFrequency"
	fields["region"] = bindings.NewStringType()
	fieldNameMap["region"] = "Region"
	fields["commitment_term"] = bindings.NewStringType()
	fieldNameMap["commitment_term"] = "CommitmentTerm"
	fields["offer_context_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["offer_context_id"] = "OfferContextId"
	fields["offer_version"] = bindings.NewStringType()
	fieldNameMap["offer_version"] = "OfferVersion"
	fields["offer_name"] = bindings.NewStringType()
	fieldNameMap["offer_name"] = "OfferName"
	fields["quantity"] = bindings.NewIntegerType()
	fieldNameMap["quantity"] = "Quantity"
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["type"] = "Type_"
	fields["price"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["price"] = "Price"
	fields["product_charge_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["product_charge_id"] = "ProductChargeId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.subscription_request", fields, reflect.TypeOf(SubscriptionRequest{}), fieldNameMap, validators)
}

func SupportWindowBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["start_day"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["start_day"] = "StartDay"
	fields["seats"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["seats"] = "Seats"
	fields["sddcs"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["sddcs"] = "Sddcs"
	fields["duration_hours"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["duration_hours"] = "DurationHours"
	fields["start_hour"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["start_hour"] = "StartHour"
	fields["support_window_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["support_window_id"] = "SupportWindowId"
	fields["metadata"] = bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.REST))
	fieldNameMap["metadata"] = "Metadata"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.support_window", fields, reflect.TypeOf(SupportWindow{}), fieldNameMap, validators)
}

func SupportWindowIdBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["window_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["window_id"] = "WindowId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.support_window_id", fields, reflect.TypeOf(SupportWindowId{}), fieldNameMap, validators)
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
	fields["created"] = bindings.NewDateTimeType()
	fieldNameMap["created"] = "Created"
	fields["version"] = bindings.NewIntegerType()
	fieldNameMap["version"] = "Version"
	fields["updated_by_user_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["updated_by_user_name"] = "UpdatedByUserName"
	fields["user_name"] = bindings.NewStringType()
	fieldNameMap["user_name"] = "UserName"
	fields["id"] = bindings.NewStringType()
	fieldNameMap["id"] = "Id"
	fields["status"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["status"] = "Status"
	fields["localized_error_message"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["localized_error_message"] = "LocalizedErrorMessage"
	fields["resource_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["resource_id"] = "ResourceId"
	fields["parent_task_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["parent_task_id"] = "ParentTaskId"
	fields["task_version"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["task_version"] = "TaskVersion"
	fields["correlation_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["correlation_id"] = "CorrelationId"
	fields["start_resource_entity_version"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["start_resource_entity_version"] = "StartResourceEntityVersion"
	fields["sub_status"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sub_status"] = "SubStatus"
	fields["task_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["task_type"] = "TaskType"
	fields["start_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
	fieldNameMap["start_time"] = "StartTime"
	fields["task_progress_phases"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(TaskProgressPhaseBindingType), reflect.TypeOf([]TaskProgressPhase{})))
	fieldNameMap["task_progress_phases"] = "TaskProgressPhases"
	fields["error_message"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["error_message"] = "ErrorMessage"
	fields["org_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fields["end_resource_entity_version"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["end_resource_entity_version"] = "EndResourceEntityVersion"
	fields["service_errors"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ServiceErrorBindingType), reflect.TypeOf([]ServiceError{})))
	fieldNameMap["service_errors"] = "ServiceErrors"
	fields["org_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org_type"] = "OrgType"
	fields["estimated_remaining_minutes"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["estimated_remaining_minutes"] = "EstimatedRemainingMinutes"
	fields["params"] = bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.REST))
	fieldNameMap["params"] = "Params"
	fields["progress_percent"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["progress_percent"] = "ProgressPercent"
	fields["phase_in_progress"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["phase_in_progress"] = "PhaseInProgress"
	fields["resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["end_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
	fieldNameMap["end_time"] = "EndTime"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.task", fields, reflect.TypeOf(Task{}), fieldNameMap, validators)
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
	return bindings.NewStructType("com.vmware.vmc.model.task_progress_phase", fields, reflect.TypeOf(TaskProgressPhase{}), fieldNameMap, validators)
}

func TermBillingOptionsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["unit_price"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["unit_price"] = "UnitPrice"
	fields["billing_frequency"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["billing_frequency"] = "BillingFrequency"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.term_billing_options", fields, reflect.TypeOf(TermBillingOptions{}), fieldNameMap, validators)
}

func TermOfferInstanceBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["description"] = bindings.NewStringType()
	fieldNameMap["description"] = "Description"
	fields["product"] = bindings.NewStringType()
	fieldNameMap["product"] = "Product"
	fields["product_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["product_type"] = "ProductType"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["currency"] = bindings.NewStringType()
	fieldNameMap["currency"] = "Currency"
	fields["region"] = bindings.NewStringType()
	fieldNameMap["region"] = "Region"
	fields["commitment_term"] = bindings.NewIntegerType()
	fieldNameMap["commitment_term"] = "CommitmentTerm"
	fields["unit_price"] = bindings.NewStringType()
	fieldNameMap["unit_price"] = "UnitPrice"
	fields["billing_options"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(TermBillingOptionsBindingType), reflect.TypeOf([]TermBillingOptions{})))
	fieldNameMap["billing_options"] = "BillingOptions"
	fields["version"] = bindings.NewStringType()
	fieldNameMap["version"] = "Version"
	fields["offer_context_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["offer_context_id"] = "OfferContextId"
	fields["product_charge_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["product_charge_id"] = "ProductChargeId"
	fields["type"] = bindings.NewStringType()
	fieldNameMap["type"] = "Type_"
	fields["product_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["product_id"] = "ProductId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.term_offer_instance", fields, reflect.TypeOf(TermOfferInstance{}), fieldNameMap, validators)
}

func TermsOfServiceResultBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["terms_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["terms_id"] = "TermsId"
	fields["signed"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["signed"] = "Signed"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.terms_of_service_result", fields, reflect.TypeOf(TermsOfServiceResult{}), fieldNameMap, validators)
}

func UpdateCredentialsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["username"] = bindings.NewStringType()
	fieldNameMap["username"] = "Username"
	fields["password"] = bindings.NewStringType()
	fieldNameMap["password"] = "Password"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.update_credentials", fields, reflect.TypeOf(UpdateCredentials{}), fieldNameMap, validators)
}

func VmcLocaleBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["locale"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["locale"] = "Locale"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.vmc_locale", fields, reflect.TypeOf(VmcLocale{}), fieldNameMap, validators)
}

func VpcInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vpc_cidr"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vpc_cidr"] = "VpcCidr"
	fields["vgw_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vgw_id"] = "VgwId"
	fields["esx_public_security_group_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["esx_public_security_group_id"] = "EsxPublicSecurityGroupId"
	fields["vif_ids"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["vif_ids"] = "VifIds"
	fields["vm_security_group_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vm_security_group_id"] = "VmSecurityGroupId"
	fields["route_table_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["route_table_id"] = "RouteTableId"
	fields["edge_subnet_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["edge_subnet_id"] = "EdgeSubnetId"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["api_association_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["api_association_id"] = "ApiAssociationId"
	fields["api_subnet_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["api_subnet_id"] = "ApiSubnetId"
	fields["private_subnet_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["private_subnet_id"] = "PrivateSubnetId"
	fields["private_association_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["private_association_id"] = "PrivateAssociationId"
	fields["esx_security_group_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["esx_security_group_id"] = "EsxSecurityGroupId"
	fields["subnet_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["subnet_id"] = "SubnetId"
	fields["internet_gateway_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["internet_gateway_id"] = "InternetGatewayId"
	fields["security_group_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["security_group_id"] = "SecurityGroupId"
	fields["association_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["association_id"] = "AssociationId"
	fields["vgw_route_table_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vgw_route_table_id"] = "VgwRouteTableId"
	fields["traffic_group_edge_vm_ips"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["traffic_group_edge_vm_ips"] = "TrafficGroupEdgeVmIps"
	fields["edge_association_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["edge_association_id"] = "EdgeAssociationId"
	fields["provider"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["provider"] = "Provider"
	fields["tgw_ips"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})),reflect.TypeOf(map[string][]string{})))
	fieldNameMap["tgw_ips"] = "TgwIps"
	fields["peering_connection_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["peering_connection_id"] = "PeeringConnectionId"
	fields["network_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["network_type"] = "NetworkType"
	fields["available_zones"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(AvailableZoneInfoBindingType), reflect.TypeOf([]AvailableZoneInfo{})))
	fieldNameMap["available_zones"] = "AvailableZones"
	fields["routetables"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewReferenceType(RouteTableInfoBindingType),reflect.TypeOf(map[string]RouteTableInfo{})))
	fieldNameMap["routetables"] = "Routetables"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.vpc_info", fields, reflect.TypeOf(VpcInfo{}), fieldNameMap, validators)
}

func VpcInfoSubnetsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vpc_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vpc_id"] = "VpcId"
	fields["cidr_block"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cidr_block"] = "CidrBlock"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["subnets"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(SubnetInfoBindingType), reflect.TypeOf([]SubnetInfo{})))
	fieldNameMap["subnets"] = "Subnets"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.vpc_info_subnets", fields, reflect.TypeOf(VpcInfoSubnets{}), fieldNameMap, validators)
}

func VpnBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["version"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["version"] = "Version"
	fields["on_prem_gateway_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["on_prem_gateway_ip"] = "OnPremGatewayIp"
	fields["on_prem_network_cidr"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["on_prem_network_cidr"] = "OnPremNetworkCidr"
	fields["pfs_enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["pfs_enabled"] = "PfsEnabled"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["channel_status"] = bindings.NewOptionalType(bindings.NewReferenceType(VpnChannelStatusBindingType))
	fieldNameMap["channel_status"] = "ChannelStatus"
	fields["on_prem_nat_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["on_prem_nat_ip"] = "OnPremNatIp"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["internal_network_ids"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["internal_network_ids"] = "InternalNetworkIds"
	fields["tunnel_statuses"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(VpnTunnelStatusBindingType), reflect.TypeOf([]VpnTunnelStatus{})))
	fieldNameMap["tunnel_statuses"] = "TunnelStatuses"
	fields["encryption"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["encryption"] = "Encryption"
	fields["enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["enabled"] = "Enabled"
	fields["state"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["state"] = "State"
	fields["dh_group"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["dh_group"] = "DhGroup"
	fields["authentication"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["authentication"] = "Authentication"
	fields["pre_shared_key"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["pre_shared_key"] = "PreSharedKey"
	fields["ike_option"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["ike_option"] = "IkeOption"
	fields["digest_algorithm"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["digest_algorithm"] = "DigestAlgorithm"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.vpn", fields, reflect.TypeOf(Vpn{}), fieldNameMap, validators)
}

func VpnChannelStatusBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["channel_status"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["channel_status"] = "ChannelStatus"
	fields["channel_state"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["channel_state"] = "ChannelState"
	fields["last_info_message"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["last_info_message"] = "LastInfoMessage"
	fields["failure_message"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["failure_message"] = "FailureMessage"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.vpn_channel_status", fields, reflect.TypeOf(VpnChannelStatus{}), fieldNameMap, validators)
}

func VpnTunnelStatusBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["on_prem_subnet"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["on_prem_subnet"] = "OnPremSubnet"
	fields["traffic_stats"] = bindings.NewOptionalType(bindings.NewReferenceType(VpnTunnelTrafficStatsBindingType))
	fieldNameMap["traffic_stats"] = "TrafficStats"
	fields["last_info_message"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["last_info_message"] = "LastInfoMessage"
	fields["local_subnet"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["local_subnet"] = "LocalSubnet"
	fields["tunnel_state"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["tunnel_state"] = "TunnelState"
	fields["failure_message"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["failure_message"] = "FailureMessage"
	fields["tunnel_status"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["tunnel_status"] = "TunnelStatus"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.vpn_tunnel_status", fields, reflect.TypeOf(VpnTunnelStatus{}), fieldNameMap, validators)
}

func VpnTunnelTrafficStatsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["packets_out"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["packets_out"] = "PacketsOut"
	fields["packet_received_errors"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["packet_received_errors"] = "PacketReceivedErrors"
	fields["rx_bytes_on_local_subnet"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["rx_bytes_on_local_subnet"] = "RxBytesOnLocalSubnet"
	fields["replay_errors"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["replay_errors"] = "ReplayErrors"
	fields["sequence_number_over_flow_errors"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sequence_number_over_flow_errors"] = "SequenceNumberOverFlowErrors"
	fields["encryption_failures"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["encryption_failures"] = "EncryptionFailures"
	fields["integrity_errors"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["integrity_errors"] = "IntegrityErrors"
	fields["packet_sent_errors"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["packet_sent_errors"] = "PacketSentErrors"
	fields["decryption_failures"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["decryption_failures"] = "DecryptionFailures"
	fields["packets_in"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["packets_in"] = "PacketsIn"
	fields["tx_bytes_from_local_subnet"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["tx_bytes_from_local_subnet"] = "TxBytesFromLocalSubnet"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.vpn_tunnel_traffic_stats", fields, reflect.TypeOf(VpnTunnelTrafficStats{}), fieldNameMap, validators)
}

func VsanAvailableCapacityBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cost"] = bindings.NewStringType()
	fieldNameMap["cost"] = "Cost"
	fields["quality"] = bindings.NewStringType()
	fieldNameMap["quality"] = "Quality"
	fields["size"] = bindings.NewIntegerType()
	fieldNameMap["size"] = "Size"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.vsan_available_capacity", fields, reflect.TypeOf(VsanAvailableCapacity{}), fieldNameMap, validators)
}

func VsanClusterReconfigBiasBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["short_description"] = bindings.NewStringType()
	fieldNameMap["short_description"] = "ShortDescription"
	fields["full_description"] = bindings.NewStringType()
	fieldNameMap["full_description"] = "FullDescription"
	fields["id"] = bindings.NewStringType()
	fieldNameMap["id"] = "Id"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.vsan_cluster_reconfig_bias", fields, reflect.TypeOf(VsanClusterReconfigBias{}), fieldNameMap, validators)
}

func VsanClusterReconfigConstraintsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["reconfig_biases"] = bindings.NewListType(bindings.NewReferenceType(VsanClusterReconfigBiasBindingType), reflect.TypeOf([]VsanClusterReconfigBias{}))
	fieldNameMap["reconfig_biases"] = "ReconfigBiases"
	fields["available_capacities"] = bindings.NewMapType(bindings.NewStringType(), bindings.NewListType(bindings.NewReferenceType(VsanAvailableCapacityBindingType), reflect.TypeOf([]VsanAvailableCapacity{})),reflect.TypeOf(map[string][]VsanAvailableCapacity{}))
	fieldNameMap["available_capacities"] = "AvailableCapacities"
	fields["default_capacities"] = bindings.NewMapType(bindings.NewStringType(), bindings.NewReferenceType(VsanAvailableCapacityBindingType),reflect.TypeOf(map[string]VsanAvailableCapacity{}))
	fieldNameMap["default_capacities"] = "DefaultCapacities"
	fields["hosts"] = bindings.NewIntegerType()
	fieldNameMap["hosts"] = "Hosts"
	fields["default_reconfig_bias_id"] = bindings.NewStringType()
	fieldNameMap["default_reconfig_bias_id"] = "DefaultReconfigBiasId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.vsan_cluster_reconfig_constraints", fields, reflect.TypeOf(VsanClusterReconfigConstraints{}), fieldNameMap, validators)
}

func VsanConfigConstraintsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["max_capacity"] = bindings.NewIntegerType()
	fieldNameMap["max_capacity"] = "MaxCapacity"
	fields["recommended_capacities"] = bindings.NewListType(bindings.NewIntegerType(), reflect.TypeOf([]int64{}))
	fieldNameMap["recommended_capacities"] = "RecommendedCapacities"
	fields["supported_capacity_increment"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["supported_capacity_increment"] = "SupportedCapacityIncrement"
	fields["min_capacity"] = bindings.NewIntegerType()
	fieldNameMap["min_capacity"] = "MinCapacity"
	fields["num_hosts"] = bindings.NewIntegerType()
	fieldNameMap["num_hosts"] = "NumHosts"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.vsan_config_constraints", fields, reflect.TypeOf(VsanConfigConstraints{}), fieldNameMap, validators)
}

func VsanEncryptionConfigBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["port"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["port"] = "Port"
	fields["certificate"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["certificate"] = "Certificate"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vmc.model.vsan_encryption_config", fields, reflect.TypeOf(VsanEncryptionConfig{}), fieldNameMap, validators)
}


