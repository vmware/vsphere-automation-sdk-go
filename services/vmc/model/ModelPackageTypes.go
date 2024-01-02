// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for package: com.vmware.vmc.model.
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
	UpdatedByUserId *string
	Created         time.Time
	// Version of this entity format: int32
	Version int64
	// User name that last updated this record
	UpdatedByUserName *string
	// User name that last updated this record
	UserName *string
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

type AccountLinkConfig struct {
	// Boolean flag identifying whether account linking should be delayed or not for the SDDC.
	DelayAccountLink *bool
}

func (s *AccountLinkConfig) GetType__() vapiBindings_.BindingType {
	return AccountLinkConfigBindingType()
}

func (s *AccountLinkConfig) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for AccountLinkConfig._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type AccountLinkSddcConfig struct {
	CustomerSubnetIds []string
	// The ID of the customer connected account to work with.
	ConnectedAccountId *string
}

func (s *AccountLinkSddcConfig) GetType__() vapiBindings_.BindingType {
	return AccountLinkSddcConfigBindingType()
}

func (s *AccountLinkSddcConfig) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for AccountLinkSddcConfig._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type Agent struct {
	// CSP Oauth client credentials to access information from POP.
	SddcCspOauthClient *AgentCspOauthClient
	// The addresses of the agent including its public IP and DNS names.
	Addresses []string
	// The internal IP address of the agent which is provided by the underlying cloud provider
	InternalIp *string
	// Boolean flag to indicate if the agent is using FQDN in the certificate
	HostnameVerifierEnabled *bool
	// Boolean flag to indicate if the agent is using CA signed certificate
	CertEnabled *bool
	// Network CIDR of the agent
	NetworkCidr *string
	// The agent id
	Id *string
	// Tenant service information that we will save as a part of agent data.
	TenantServiceInfo *TenantServiceInfo
	// The list of white listed domains of the HTTP proxy.
	TinyproxyWhitelist []string
	// Boolean flag to indicate if the agent is healthy.
	Healthy          *bool
	CustomProperties map[string]string
	// The internal management IP address of the agent exposed to the SDDC, which might be different from the internal IP
	ManagementIp *string
	// Boolean flag to indicate if the agent is the master, only the master Agent is accessible
	Master *bool
	// Network netmask of the agent
	NetworkNetmask *string
	// Network gateway of the agent
	NetworkGateway *string
	// The cloud provider
	Provider string
	// The accessible URL of the agent service, it is resolved to public IP address from the Internet and private IP address within SDDC
	AgentUrl *string
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

func (s *Agent) GetType__() vapiBindings_.BindingType {
	return AgentBindingType()
}

func (s *Agent) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for Agent._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type AgentCspOauthClient struct {
	// Agent CSP Oauth client credentials.
	CspOauthClient AgentOauthClient
	// The Oauth client creation time. format: date-time
	CreateAt time.Time
}

func (s *AgentCspOauthClient) GetType__() vapiBindings_.BindingType {
	return AgentCspOauthClientBindingType()
}

func (s *AgentCspOauthClient) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for AgentCspOauthClient._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type AgentOauthClient struct {
	// The Oauth client secret.
	ClientSecret string
	// The Oauth client ID.
	ClientId string
}

func (s *AgentOauthClient) GetType__() vapiBindings_.BindingType {
	return AgentOauthClientBindingType()
}

func (s *AgentOauthClient) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for AgentOauthClient._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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

func (s *AmiInfo) GetType__() vapiBindings_.BindingType {
	return AmiInfoBindingType()
}

func (s *AmiInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for AmiInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type AvailableZoneInfo struct {
	Subnets []Subnet
	// available zone id
	Id *string
}

func (s *AvailableZoneInfo) GetType__() vapiBindings_.BindingType {
	return AvailableZoneInfoBindingType()
}

func (s *AvailableZoneInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for AvailableZoneInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type AwsAgent struct {
	InstanceId             *string
	AvailabilityZoneInfoId *string
	InstanceProfileInfo    *InstanceProfileInfo
	KeyPair                *AwsKeyPair
	DefaultEniId           *string
	// CSP Oauth client credentials to access information from POP.
	SddcCspOauthClient *AgentCspOauthClient
	// The addresses of the agent including its public IP and DNS names.
	Addresses []string
	// The internal IP address of the agent which is provided by the underlying cloud provider
	InternalIp *string
	// Boolean flag to indicate if the agent is using FQDN in the certificate
	HostnameVerifierEnabled *bool
	// Boolean flag to indicate if the agent is using CA signed certificate
	CertEnabled *bool
	// Network CIDR of the agent
	NetworkCidr *string
	// The agent id
	Id *string
	// Tenant service information that we will save as a part of agent data.
	TenantServiceInfo *TenantServiceInfo
	// The list of white listed domains of the HTTP proxy.
	TinyproxyWhitelist []string
	// Boolean flag to indicate if the agent is healthy.
	Healthy          *bool
	CustomProperties map[string]string
	// The internal management IP address of the agent exposed to the SDDC, which might be different from the internal IP
	ManagementIp *string
	// Boolean flag to indicate if the agent is the master, only the master Agent is accessible
	Master *bool
	// Network netmask of the agent
	NetworkNetmask *string
	// Network gateway of the agent
	NetworkGateway *string
	// The cloud provider
	Provider string
	// The accessible URL of the agent service, it is resolved to public IP address from the Internet and private IP address within SDDC
	AgentUrl *string
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

func (s *AwsAgent) GetType__() vapiBindings_.BindingType {
	return AwsAgentBindingType()
}

func (s *AwsAgent) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for AwsAgent._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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

func (s *AwsCloudProvider) GetType__() vapiBindings_.BindingType {
	return AwsCloudProviderBindingType()
}

func (s *AwsCloudProvider) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for AwsCloudProvider._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type AwsCompatibleSubnets struct {
	CustomerAvailableZones []string
	VpcMap                 map[string]VpcInfoSubnets
}

func (s *AwsCompatibleSubnets) GetType__() vapiBindings_.BindingType {
	return AwsCompatibleSubnetsBindingType()
}

func (s *AwsCompatibleSubnets) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for AwsCompatibleSubnets._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type AwsCustomerConnectedAccount struct {
	Updated time.Time
	// User id that last updated this record
	UserId string
	// User id that last updated this record
	UpdatedByUserId *string
	Created         time.Time
	// Version of this entity format: int32
	Version int64
	// User name that last updated this record
	UpdatedByUserName *string
	// User name that last updated this record
	UserName *string
	// Unique ID for this entity
	Id             string
	PolicyPayerArn *string
	// Provides a map of regions to availability zones from the shadow account's perspective
	RegionToAzToShadowMapping map[string]map[string]string
	OrgId                     *string
	CfStackName               *string
	// Possible values are:
	//
	// * AwsCustomerConnectedAccount#AwsCustomerConnectedAccount_STATE_ACTIVE
	// * AwsCustomerConnectedAccount#AwsCustomerConnectedAccount_STATE_BROKEN
	// * AwsCustomerConnectedAccount#AwsCustomerConnectedAccount_STATE_DELETED
	State                *string
	AccountNumber        *string
	PolicyServiceArn     *string
	PolicyExternalId     *string
	PolicyPayerLinkedArn *string
}

const AwsCustomerConnectedAccount_STATE_ACTIVE = "ACTIVE"
const AwsCustomerConnectedAccount_STATE_BROKEN = "BROKEN"
const AwsCustomerConnectedAccount_STATE_DELETED = "DELETED"

func (s *AwsCustomerConnectedAccount) GetType__() vapiBindings_.BindingType {
	return AwsCustomerConnectedAccountBindingType()
}

func (s *AwsCustomerConnectedAccount) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for AwsCustomerConnectedAccount._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type AwsEsxHost struct {
	// name of the selected capacity pool for the instance.
	CapacityPool *string
	EniList      []EniInfo
	XeniInfo     *XEniInfo
	// Partition number alloted to host. format: int32
	PartitionNumber *int64
	// the number of continuous times the host has been in an EC2 RUNNING state. format: int32
	Ec2InstanceRunningStatusCount *int64
	// Nic info in the user data for each host.
	EsxNicInfo           *EsxNicInfo
	InternalPublicIpPool []SddcPublicIp
	// Availability zone where the host is provisioned.
	AvailabilityZone *string
	EsxId            *string
	DurableHostName  *string
	StateLastUpdated *time.Time
	Name             *string
	CustomProperties map[string]string
	Hostname         *string
	Provider         string
	// Backing cloud provider instance type for host.
	InstanceType *string
	MacAddress   *string
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

func (s *AwsEsxHost) GetType__() vapiBindings_.BindingType {
	return AwsEsxHostBindingType()
}

func (s *AwsEsxHost) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for AwsEsxHost._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type AwsKeyPair struct {
	KeyName        *string
	KeyFingerprint *string
	KeyMaterial    *string
	Tags           []Tag
}

func (s *AwsKeyPair) GetType__() vapiBindings_.BindingType {
	return AwsKeyPairBindingType()
}

func (s *AwsKeyPair) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for AwsKeyPair._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type AwsKmsInfo struct {
	// The ARN associated with the customer master key for this cluster.
	AmazonResourceName string
}

func (s *AwsKmsInfo) GetType__() vapiBindings_.BindingType {
	return AwsKmsInfoBindingType()
}

func (s *AwsKmsInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for AwsKmsInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type AwsSddcConfig struct {
	ExistingVpcId *string
	// Indicates the desired licensing support, if any, of Microsoft software.
	MsftLicenseConfig *MsftLicensingConfig
	// The account linking configuration, we will keep this one and remove accountLinkSddcConfig finally.
	AccountLinkConfig *AccountLinkConfig
	// AWS VPC IP range. Only prefix of 16 or 20 is currently supported.
	VpcCidr *string
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
	// Outpost ID of the SDDC. Used only for outpost deployments.
	OutpostId *string
	// The storage capacity value to be requested for the sddc primary cluster, in GiBs. If provided, instead of using the direct-attached storage, a capacity value amount of seperable storage will be used. format: int64
	StorageCapacity *int64
	Name            string
	// A list of the SDDC linking configurations to use.
	AccountLinkSddcConfig []AccountLinkSddcConfig
	// Possible values are:
	//
	// * SddcConfig#SddcConfig_HOST_INSTANCE_TYPE_I3_METAL
	// * SddcConfig#SddcConfig_HOST_INSTANCE_TYPE_R5_METAL
	// * SddcConfig#SddcConfig_HOST_INSTANCE_TYPE_I3EN_METAL
	// * SddcConfig#SddcConfig_HOST_INSTANCE_TYPE_I4I_METAL
	// * SddcConfig#SddcConfig_HOST_INSTANCE_TYPE_C6I_METAL
	// * SddcConfig#SddcConfig_HOST_INSTANCE_TYPE_M7I_METAL_48XL
	// * SddcConfig#SddcConfig_HOST_INSTANCE_TYPE_M7I_METAL_24XL
	//
	//  The instance type for the esx hosts in the primary cluster of the SDDC.
	HostInstanceType *string
	// Identifier of the billing account in subscription.
	BillingAccountId *string
	Region           *string
	// If provided, will be assigned as SDDC id of the provisioned SDDC. format: UUID
	SddcId *string
	// If provided, configuration from the template will applied to the provisioned SDDC. format: UUID
	SddcTemplateId *string
	// Denotes the sddc type , if the value is null or empty, the type is considered as default.
	SddcType *string
	// Possible values are:
	//
	// * SddcConfig#SddcConfig_VSAN_VERSION_VSAN1
	// * SddcConfig#SddcConfig_VSAN_VERSION_VSANESA
	// * SddcConfig#SddcConfig_VSAN_VERSION_NOVSAN
	//
	//  The vSAN version to be used in the SDDC's primary cluster.
	VsanVersion *string
	// Possible values are:
	//
	// * SddcConfig#SddcConfig_PROVIDER_AWS
	// * SddcConfig#SddcConfig_PROVIDER_ZEROCLOUD
	//
	//  Determines what additional properties are available based on cloud provider.
	Provider string
	// The SSO domain name to use for vSphere users. If not specified, vmc.local will be used.
	SsoDomain *string
	NumHosts  int64
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

func (s *AwsSddcConfig) GetType__() vapiBindings_.BindingType {
	return AwsSddcConfigBindingType()
}

func (s *AwsSddcConfig) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for AwsSddcConfig._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type AwsSddcConnection struct {
	Updated time.Time
	// User id that last updated this record
	UserId string
	// User id that last updated this record
	UpdatedByUserId *string
	Created         time.Time
	// Version of this entity format: int32
	Version int64
	// User name that last updated this record
	UpdatedByUserName *string
	// User name that last updated this record
	UserName *string
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

func (s *AwsSddcConnection) GetType__() vapiBindings_.BindingType {
	return AwsSddcConnectionBindingType()
}

func (s *AwsSddcConnection) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for AwsSddcConnection._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type AwsSddcResourceConfig struct {
	BackupRestoreBucket *string
	PublicIpPool        []SddcPublicIp
	VpcInfo             *VpcInfo
	MgwPublicIpPool     []SddcPublicIp
	KmsVpcEndpoint      *KmsVpcEndpoint
	// maximum number of public IP that user can allocate.
	MaxNumPublicIp        *int64
	EsxInstanceProfile    *InstanceProfileInfo
	AccountLinkSddcConfig []SddcLinkConfig
	VsanEncryptionConfig  *VsanEncryptionConfig
	NsxUserClient         *CspOauthClient
	VpcInfoPeeredAgent    *VpcInfo
	NsxServiceClient      *CspOauthClient
	// (deprecated)
	VcIp *string
	// Name for management appliance network.
	MgmtApplianceNetworkName *string
	// URL of the NSX Manager
	NsxMgrUrl *string
	// This flag determines whether vLCM is enabled on this sddc or not.
	VlcmEnabled *bool
	// NSX cloud audit Password
	NsxCloudAuditPassword *string
	// vCenter to csp federation status.
	VcCspLoginStatus *string
	// NSX cloud admin password
	NsxCloudAdminPassword *string
	// The ManagedObjectReference of the management Datastore
	ManagementDs *string
	// nsx api entire base url
	NsxApiPublicEndpointUrl *string
	// Nfs Mode Flag, for nfs mounting.
	NfsMode *bool
	// Password for vCenter SDDC administrator
	CloudPassword *string
	SddcNetworks  []string
	// List of clusters in the SDDC.
	Clusters []Cluster
	// Username for vCenter SDDC administrator
	CloudUsername *string
	// Possible values are:
	//
	// * SddcResourceConfig#SddcResourceConfig_DEPLOYMENT_TYPE_SINGLE_AZ
	// * SddcResourceConfig#SddcResourceConfig_DEPLOYMENT_TYPE_MULTI_AZ
	//
	//  Denotes if this is a SingleAZ SDDC or a MultiAZ SDDC.
	DeploymentType         *string
	PopAgentXeniConnection *PopAgentXeniConnection
	// NSX Manager internal management IP
	NsxMgrManagementIp *string
	// NSX cloud audit user name
	NsxCloudAudit *string
	// Cluster Id to add ESX workflow
	EsxClusterId *string
	// Management Gateway Id
	MgwId *string
	// URL of the vCenter server
	VcUrl    *string
	EsxHosts []AwsEsxHost
	// Group name for vCenter SDDC administrator
	CloudUserGroup *string
	ManagementRp   *string
	// Availability zone where the witness node is provisioned for a MultiAZ SDDC. This is null for a SingleAZ SDDC.
	WitnessAvailabilityZone *string
	// Whether this sddc is maintained by its desired state documents.
	SddcDesiredState *bool
	SddcSize         *SddcSize
	// This flag determines whether CVDS is enabled on this sddc or not.
	CvdsEnabled *bool
	// List of Controller IPs
	NsxControllerIps []string
	// Marks that the SDDC VC should be deployed with two hostnames.
	TwoHostnameVcDeployment *bool
	// ESX host subnet
	EsxHostSubnet *string
	// The SSO domain name to use for vSphere users
	SsoDomain *string
	// region in which sddc is deployed
	Region *string
	// if true, use the private IP addresses to register DNS records for the management VMs
	DnsWithManagementVmPrivateIp *bool
	// vCenter public IP
	VcPublicIp *string
	// (deprecated)
	PscIp *string
	// if true, NSX-T UI is enabled.
	Nsxt *bool
	// Key provider data.
	KeyProvider []KeyProviderData
	// PSC internal management IP
	PscManagementIp *string
	// URL of the PSC server
	PscUrl *string
	Cgws   []string
	// Availability zones over which esx hosts are provisioned. MultiAZ SDDCs will have hosts provisioned over two availability zones while SingleAZ SDDCs will provision over one.
	AvailabilityZones []string
	// Mark if Containerized Permissions has been enabled on vCenter.
	VcContainerizedPermissionsEnabled *bool
	CustomProperties                  map[string]string
	// Possible values are:
	//
	// * SddcResourceConfig#SddcResourceConfig_PROVIDER_AWS
	// * SddcResourceConfig#SddcResourceConfig_PROVIDER_ZEROCLOUD
	//
	//  Discriminator for additional properties
	Provider string
	// vCenter internal management IP
	VcManagementIp *string
	// The Microsoft license status of this SDDC.
	MsftLicenseConfig *MsftLicensingConfig
	// NSX-T Native Oauth client for UI.
	NsxNativeClient *CspOauthClient
	// unique id of the vCenter server
	VcInstanceId *string
	// oAuth client for enabling federation on vCenter.
	VcOauthClient *CspOauthClient
	// skip creating vxlan for compute gateway for SDDC provisioning
	SkipCreatingVxlan *bool
	SddcManifest      *SddcManifest
	// VXLAN IP subnet
	VxlanSubnet  *string
	SddcSecurity *SddcSecurity
	// sddc identifier
	SddcId *string
	// Outpost configuration of this SDDC.
	OutpostConfig *OutpostConfig
	// URL of the NSX Manager UI login for local user access
	NsxMgrLoginUrl *string
	// Break-glass URL for non-federated login.
	VcBreakGlassUrl *string
	// NSX cloud admin user name
	NsxCloudAdmin *string
	NsxtAddons    *NsxtAddons
}

// Identifier denoting this class, when it is used in polymorphic context.
//
// This value should be assigned to the property which is used to discriminate the actual type used in the polymorphic context.
const AwsSddcResourceConfig__TYPE_IDENTIFIER = "AWS"

func (s *AwsSddcResourceConfig) GetType__() vapiBindings_.BindingType {
	return AwsSddcResourceConfigBindingType()
}

func (s *AwsSddcResourceConfig) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for AwsSddcResourceConfig._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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

func (s *AwsSubnet) GetType__() vapiBindings_.BindingType {
	return AwsSubnetBindingType()
}

func (s *AwsSubnet) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for AwsSubnet._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type AwsWitnessEsx struct {
	// Instance id of witness Esx host
	InstanceId *string
	EniInfo    *EniInfo
	// Possible values are:
	//
	// * WitnessEsx#WitnessEsx_ENUM_STATE_DEPLOYING
	// * WitnessEsx#WitnessEsx_ENUM_STATE_PROVISIONED
	// * WitnessEsx#WitnessEsx_ENUM_STATE_DELETING
	// * WitnessEsx#WitnessEsx_ENUM_STATE_READY
	// * WitnessEsx#WitnessEsx_ENUM_STATE_FAILED
	//
	//  The state of the witness Esx Host
	EnumState *string
	// Name of the witness Esx Host
	Name *string
	// Mac Address for witness Esx Host
	MacAddress *string
	// The unique identifier for the witness esx host format: uuid
	EsxId *string
	// Host name
	Hostname *string
	// Possible values are:
	//
	// * WitnessEsx#WitnessEsx_PROVIDER_AWS
	// * WitnessEsx#WitnessEsx_PROVIDER_ZEROCLOUD
	//
	//  The cloud provider for this witness Esx Host
	Provider string
}

// Identifier denoting this class, when it is used in polymorphic context.
//
// This value should be assigned to the property which is used to discriminate the actual type used in the polymorphic context.
const AwsWitnessEsx__TYPE_IDENTIFIER = "AWS"

func (s *AwsWitnessEsx) GetType__() vapiBindings_.BindingType {
	return AwsWitnessEsxBindingType()
}

func (s *AwsWitnessEsx) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for AwsWitnessEsx._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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

type Cluster struct {
	// The capacity of this cluster.
	ClusterCapacity *EntityCapacity
	EsxHostList     []AwsEsxHost
	// WCP details for a given cluster
	WcpDetails *WcpDetails
	// The Microsoft license configuration of this cluster.
	MsftLicenseConfig *MsftLicensingConfig
	// Possible values are:
	//
	// * Cluster#Cluster_CLUSTER_STATE_DEPLOYING
	// * Cluster#Cluster_CLUSTER_STATE_ADDING_HOSTS
	// * Cluster#Cluster_CLUSTER_STATE_READY
	// * Cluster#Cluster_CLUSTER_STATE_FAILED
	ClusterState *string
	// Information of the hosts added to this cluster
	EsxHostInfo *EsxHostInfo
	// Number of cores enabled on ESX hosts added to this cluster format: int32
	HostCpuCoresCount *int64
	// Partition placement group infos
	PartitionPlacementGroupInfo []PartitionPlacementGroupInfo
	ClusterId                   string
	// Possible values are:
	//
	// * Cluster#Cluster_AVAILABILITY_TYPE_SINGLE_AZ
	// * Cluster#Cluster_AVAILABILITY_TYPE_MULTI_AZ
	AvailabilityType *string
	MgmtRpName       *string
	// AWS Key Management Service information associated with this cluster
	AwsKmsInfo        *AwsKmsInfo
	AvailabilityZones []string
	ComputeRpName     *string
	// Witness node
	VsanWitness      *AwsWitnessEsx
	CustomProperties map[string]string
	ClusterName      *string
	// Version of vSAN used in this cluster. Empty means vsan 1.0
	VsanVersion *string
	// Specifies whether hyperThreading is disabled/enabled explicitly
	HyperThreadingEnabled *bool
}

const Cluster_CLUSTER_STATE_DEPLOYING = "DEPLOYING"
const Cluster_CLUSTER_STATE_ADDING_HOSTS = "ADDING_HOSTS"
const Cluster_CLUSTER_STATE_READY = "READY"
const Cluster_CLUSTER_STATE_FAILED = "FAILED"
const Cluster_AVAILABILITY_TYPE_SINGLE_AZ = "SINGLE_AZ"
const Cluster_AVAILABILITY_TYPE_MULTI_AZ = "MULTI_AZ"

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

type ClusterConfig struct {
	// Customize CPU cores on hosts in a cluster. Specify number of cores to be enabled on hosts in a cluster. format: int32
	HostCpuCoresCount *int64
	// Possible values are:
	//
	// * ClusterConfig#ClusterConfig_HOST_INSTANCE_TYPE_I3_METAL
	// * ClusterConfig#ClusterConfig_HOST_INSTANCE_TYPE_R5_METAL
	// * ClusterConfig#ClusterConfig_HOST_INSTANCE_TYPE_I3EN_METAL
	// * ClusterConfig#ClusterConfig_HOST_INSTANCE_TYPE_I4I_METAL
	// * ClusterConfig#ClusterConfig_HOST_INSTANCE_TYPE_C6I_METAL
	// * ClusterConfig#ClusterConfig_HOST_INSTANCE_TYPE_M7I_METAL_48XL
	// * ClusterConfig#ClusterConfig_HOST_INSTANCE_TYPE_M7I_METAL_24XL
	//
	//  The instance type for the esx hosts added to this cluster.
	HostInstanceType *string
	// For EBS-backed instances only, the requested storage capacity in GiB. format: int64
	StorageCapacity *int64
	// Cluster hosts need to be provisioned on these availability zones.
	AvailabilityZones []string
	// The desired Microsoft license status to apply to this cluster.
	MsftLicenseConfig *MsftLicensingConfig
	NumHosts          int64
}

const ClusterConfig_HOST_INSTANCE_TYPE_I3_METAL = "i3.metal"
const ClusterConfig_HOST_INSTANCE_TYPE_R5_METAL = "r5.metal"
const ClusterConfig_HOST_INSTANCE_TYPE_I3EN_METAL = "i3en.metal"
const ClusterConfig_HOST_INSTANCE_TYPE_I4I_METAL = "i4i.metal"
const ClusterConfig_HOST_INSTANCE_TYPE_C6I_METAL = "c6i.metal"
const ClusterConfig_HOST_INSTANCE_TYPE_M7I_METAL_48XL = "m7i.metal-48xl"
const ClusterConfig_HOST_INSTANCE_TYPE_M7I_METAL_24XL = "m7i.metal-24xl"

func (s *ClusterConfig) GetType__() vapiBindings_.BindingType {
	return ClusterConfigBindingType()
}

func (s *ClusterConfig) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ClusterConfig._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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

func (s *ClusterReconfigureParams) GetType__() vapiBindings_.BindingType {
	return ClusterReconfigureParamsBindingType()
}

func (s *ClusterReconfigureParams) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ClusterReconfigureParams._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type ComputeGatewayTemplate struct {
	PublicIp        *SddcPublicIp
	PrimaryDns      *string
	SecondaryDns    *string
	FirewallRules   []FirewallRule
	Vpns            []Vpn
	LogicalNetworks []LogicalNetwork
	NatRules        []NatRule
	L2Vpn           *vapiData_.StructValue
}

func (s *ComputeGatewayTemplate) GetType__() vapiBindings_.BindingType {
	return ComputeGatewayTemplateBindingType()
}

func (s *ComputeGatewayTemplate) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ComputeGatewayTemplate._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Represents a configuration spec for any sddc provision operation.
type ConfigSpec struct {
	// Indicates if this org supports deployments that are powered by VMC on outposts
	PoweredByOutpostAvailable *bool
	// Indicates after how many days the sddc should expire
	ExpiryInDays *int64
	SddcSizes    []string
	// Map of region to instance types available in that region
	Availability map[string][]InstanceTypeConfig
	// Map of region to instance types available for the powered by type.
	PoweredBy map[string][]PoweredByInstanceTypeConfig
}

func (s *ConfigSpec) GetType__() vapiBindings_.BindingType {
	return ConfigSpecBindingType()
}

func (s *ConfigSpec) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ConfigSpec._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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
	// URL path ONLY for CURL tests.
	Path *string
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
	// Package size option used ONLY for PING tests.
	Pktsize *string
}

const ConnectivityAgentValidation_SOURCE_VCENTER = "VCENTER"
const ConnectivityAgentValidation_SOURCE_SRM = "SRM"
const ConnectivityAgentValidation_SOURCE_VR = "VR"
const ConnectivityAgentValidation_TYPE_PING = "PING"
const ConnectivityAgentValidation_TYPE_TRACEROUTE = "TRACEROUTE"
const ConnectivityAgentValidation_TYPE_DNS = "DNS"
const ConnectivityAgentValidation_TYPE_CONNECTIVITY = "CONNECTIVITY"
const ConnectivityAgentValidation_TYPE_CURL = "CURL"

func (s *ConnectivityAgentValidation) GetType__() vapiBindings_.BindingType {
	return ConnectivityAgentValidationBindingType()
}

func (s *ConnectivityAgentValidation) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ConnectivityAgentValidation._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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

func (s *ConnectivityValidationGroup) GetType__() vapiBindings_.BindingType {
	return ConnectivityValidationGroupBindingType()
}

func (s *ConnectivityValidationGroup) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ConnectivityValidationGroup._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type ConnectivityValidationGroups struct {
	// List of groups.
	Groups []ConnectivityValidationGroup
}

func (s *ConnectivityValidationGroups) GetType__() vapiBindings_.BindingType {
	return ConnectivityValidationGroupsBindingType()
}

func (s *ConnectivityValidationGroups) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ConnectivityValidationGroups._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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

func (s *ConnectivityValidationInput) GetType__() vapiBindings_.BindingType {
	return ConnectivityValidationInputBindingType()
}

func (s *ConnectivityValidationInput) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ConnectivityValidationInput._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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

func (s *ConnectivityValidationSubGroup) GetType__() vapiBindings_.BindingType {
	return ConnectivityValidationSubGroupBindingType()
}

func (s *ConnectivityValidationSubGroup) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ConnectivityValidationSubGroup._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type CspOauthClient struct {
	RedirectUris   []string
	AccessTokenTTL *int64
	RedirectUri    *string
	GrantTypes     []string
	OrgId          *string
	// The Oauth client secret.
	Secret          *string
	RefreshTokenTTL *int64
	ResourceLink    *string
	// The Oauth client ID.
	Id string
}

func (s *CspOauthClient) GetType__() vapiBindings_.BindingType {
	return CspOauthClientBindingType()
}

func (s *CspOauthClient) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for CspOauthClient._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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

func (s *CustomerEniInfo) GetType__() vapiBindings_.BindingType {
	return CustomerEniInfoBindingType()
}

func (s *CustomerEniInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for CustomerEniInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// information for EBS-backed VSAN configuration
type EbsBackedVsanConfig struct {
	// instance type for EBS-backed VSAN configuration
	InstanceType *string
}

func (s *EbsBackedVsanConfig) GetType__() vapiBindings_.BindingType {
	return EbsBackedVsanConfigBindingType()
}

func (s *EbsBackedVsanConfig) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for EbsBackedVsanConfig._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type EnablementInfo struct {
	Enabled bool
	// Add-on name, for example \"hcx\", \"drass\", skynet\". If it is \"default\", the value of enabled is for default settings of future add-ons.
	Name string
}

func (s *EnablementInfo) GetType__() vapiBindings_.BindingType {
	return EnablementInfoBindingType()
}

func (s *EnablementInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for EnablementInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Information of the x-eni created.
type EniInfo struct {
	// secondary private ips to be assigned.
	SecondaryIps []string
	// Id of the attachment, it is used to detach eni from instance.
	AttachmentId *string
	// Id of the association for branch ENI, it is used to disassociate branch eni from trunk eni.
	AssociationId *string
	// The vmknic id or null if the ENI does not mapped to a vmknic.
	VmkId *string
	// index of eni device..
	DeviceIndex *int64
	// Security Group of Eni.
	SecurityGroupId *string
	// Id of the instance to be attached.
	InstanceId *string
	// Subnet it belongs to.
	SubnetId *string
	// public ips to be associated with the secondary private ips.
	PublicIps []string
	// Private ip of eni.
	PrivateIp *string
	// Mac address of nic.
	MacAddress *string
	// Whether to set sourceDestCheck as false..
	SourceDestCheckFalse *bool
	// The portgroup name or null if the ENI does not mapped to a portgroup.
	Portgroup *string
	// Interface ID.
	Id *string
}

func (s *EniInfo) GetType__() vapiBindings_.BindingType {
	return EniInfoBindingType()
}

func (s *EniInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for EniInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Decribes the capacity of an entity in a provisioned cluster or to be used in an already-provisioned cluster.
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

func (s *EntityCapacity) GetType__() vapiBindings_.BindingType {
	return EntityCapacityBindingType()
}

func (s *EntityCapacity) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for EntityCapacity._GetDataValue method - %s",
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

type EsxConfig struct {
	// Availability zone where the hosts should be provisioned. (Can be specified only for privileged host operations).
	AvailabilityZone *string
	// An option to indicate if the host needs to be strictly placed in a placement group. Fail the operation otherwise.
	StrictPlacement *bool
	// An optional cluster id if the esxs operation has to be on a specific cluster.
	ClusterId *string
	// An optional list of ESX IDs to remove. format: UUID
	Esxs     []string
	NumHosts int64
}

func (s *EsxConfig) GetType__() vapiBindings_.BindingType {
	return EsxConfigBindingType()
}

func (s *EsxConfig) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for EsxConfig._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type EsxHost struct {
	// Availability zone where the host is provisioned.
	AvailabilityZone *string
	EsxId            *string
	DurableHostName  *string
	StateLastUpdated *time.Time
	Name             *string
	CustomProperties map[string]string
	Hostname         *string
	Provider         string
	// Backing cloud provider instance type for host.
	InstanceType *string
	MacAddress   *string
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

func (s *EsxHost) GetType__() vapiBindings_.BindingType {
	return EsxHostBindingType()
}

func (s *EsxHost) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for EsxHost._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type EsxHostInfo struct {
	// Backing cloud provider instance type for cluster.
	InstanceType *string
}

func (s *EsxHostInfo) GetType__() vapiBindings_.BindingType {
	return EsxHostInfoBindingType()
}

func (s *EsxHostInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for EsxHostInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type EsxNicInfo struct {
	Ip       *string
	Mac      *string
	Vlan     *string
	Gateway  *string
	NetStack *string
}

func (s *EsxNicInfo) GetType__() vapiBindings_.BindingType {
	return EsxNicInfoBindingType()
}

func (s *EsxNicInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for EsxNicInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type FirewallRule struct {
	// Possible values are:
	//
	// * FirewallRule#FirewallRule_RULE_TYPE_USER
	// * FirewallRule#FirewallRule_RULE_TYPE_DEFAULT
	RuleType       *string
	ApplicationIds []string
	Name           *string
	// Deprecated, left for backwards compatibility. Remove once UI stops using it.
	RuleInterface *string
	// Optional. Possible formats are IP, IP1-IPn, CIDR or comma separated list of those entries. If not specified, defaults to 'any'.
	Destination      *string
	Id               *string
	DestinationScope *FirewallRuleScope
	// Optional. Possible formats are IP, IP1-IPn, CIDR or comma separated list of those entries. If not specified, defaults to 'any'.
	Source      *string
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

func (s *FirewallRule) GetType__() vapiBindings_.BindingType {
	return FirewallRuleBindingType()
}

func (s *FirewallRule) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for FirewallRule._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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

func (s *FirewallRuleScope) GetType__() vapiBindings_.BindingType {
	return FirewallRuleScopeBindingType()
}

func (s *FirewallRuleScope) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for FirewallRuleScope._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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

func (s *FirewallService) GetType__() vapiBindings_.BindingType {
	return FirewallServiceBindingType()
}

func (s *FirewallService) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for FirewallService._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Describes common properties for MGW and CGW configuration templates
type GatewayTemplate struct {
	PublicIp      *SddcPublicIp
	PrimaryDns    *string
	SecondaryDns  *string
	FirewallRules []FirewallRule
	Vpns          []Vpn
}

func (s *GatewayTemplate) GetType__() vapiBindings_.BindingType {
	return GatewayTemplateBindingType()
}

func (s *GatewayTemplate) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for GatewayTemplate._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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

func (s *GlcmBundle) GetType__() vapiBindings_.BindingType {
	return GlcmBundleBindingType()
}

func (s *GlcmBundle) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for GlcmBundle._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Decribes the capacity of an instance-type for an unprovisioned sddc or cluster.
type InstanceEntityCapacity struct {
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
	// The vSAN ESA storage capacity for the given entity in GiB.
	EsaStorageCapacityGib *int64
}

func (s *InstanceEntityCapacity) GetType__() vapiBindings_.BindingType {
	return InstanceEntityCapacityBindingType()
}

func (s *InstanceEntityCapacity) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for InstanceEntityCapacity._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type InstanceProfileInfo struct {
	RoleName            *string
	InstanceProfileName *string
	PolicyName          *string
}

func (s *InstanceProfileInfo) GetType__() vapiBindings_.BindingType {
	return InstanceProfileInfoBindingType()
}

func (s *InstanceProfileInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for InstanceProfileInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Represents a structure for basic instance type config.
type InstanceTypeBasicConfig struct {
	// Boolean to indicate whether clusters and SDDCs using this instance type support vSAN ESA.
	VsanEsaSupported *bool
	// Display name of instance_type.
	DisplayName *string
	// Description of the instance_type.
	Description *string
	// Instance type name.
	InstanceType *string
	// The capacity of the given instance type.
	EntityCapacity *InstanceEntityCapacity
	// Label for instance_type.
	Label *string
}

func (s *InstanceTypeBasicConfig) GetType__() vapiBindings_.BindingType {
	return InstanceTypeBasicConfigBindingType()
}

func (s *InstanceTypeBasicConfig) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for InstanceTypeBasicConfig._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Represents a structure for instance type config
type InstanceTypeConfig struct {
	// Boolean to indicate whether clusters and SDDCs using this instance type support vSAN ESA.
	VsanEsaSupported *bool
	// Display name of instance_type.
	DisplayName *string
	// Description of the instance_type.
	Description *string
	// Instance type name.
	InstanceType *string
	// The capacity of the given instance type.
	EntityCapacity *InstanceEntityCapacity
	// Label for instance_type.
	Label *string
	// Error due to which instance provisioning is restricted on the cluster.
	InstanceProvisioningErrorCause *string
	// Map of host count vs valid cpu core values for the given instance type.
	SupportedCpuCores map[string][]int64
	// Array of number of hosts allowed for this operation. Range of hosts user can select for sddc provision
	Hosts []int64
	// Array of valid cpu cores values for the given instance type.
	CpuCores []int64
	// Boolean to indicate whether hyperThreading is supported for an instance type.
	HyperThreadingSupported *bool
}

func (s *InstanceTypeConfig) GetType__() vapiBindings_.BindingType {
	return InstanceTypeConfigBindingType()
}

func (s *InstanceTypeConfig) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for InstanceTypeConfig._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type KeyProviderData struct {
	ProviderType          *string
	NativeSpecKeyMaterial *NativeSpecKeyMaterial
	IsDefaultKeyProvider  *bool
	IsHostKeyProvider     *bool
	Provider              *string
}

func (s *KeyProviderData) GetType__() vapiBindings_.BindingType {
	return KeyProviderDataBindingType()
}

func (s *KeyProviderData) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for KeyProviderData._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type KmsVpcEndpoint struct {
	// The identifier of the VPC endpoint created to AWS Key Management Service
	VpcEndpointId       *string
	NetworkInterfaceIds []string
}

func (s *KmsVpcEndpoint) GetType__() vapiBindings_.BindingType {
	return KmsVpcEndpointBindingType()
}

func (s *KmsVpcEndpoint) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for KmsVpcEndpoint._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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

func (s *L2Vpn) GetType__() vapiBindings_.BindingType {
	return L2VpnBindingType()
}

func (s *L2Vpn) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for L2Vpn._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type LinkRequest struct {
	// The UUID of the task we're using to track the status of this link request. format: uuid
	TrackingTask *string
	// The time at which the template URL expires. Once this has elapsed, the request is no longer invokable. format: date-time
	ExpirationDate *time.Time
	// The constructed template URL. Just shows the contents of the template for viewing/download.
	TemplateUrl *string
	// The final URL to be given to the user as a \"click here to start\" link.
	TemplateExecutionUrl *string
}

func (s *LinkRequest) GetType__() vapiBindings_.BindingType {
	return LinkRequestBindingType()
}

func (s *LinkRequest) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for LinkRequest._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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
	Id       *string
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

func (s *LogicalNetwork) GetType__() vapiBindings_.BindingType {
	return LogicalNetworkBindingType()
}

func (s *LogicalNetwork) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for LogicalNetwork._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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

func (s *MaintenanceWindow) GetType__() vapiBindings_.BindingType {
	return MaintenanceWindowBindingType()
}

func (s *MaintenanceWindow) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for MaintenanceWindow._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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

func (s *MaintenanceWindowEntry) GetType__() vapiBindings_.BindingType {
	return MaintenanceWindowEntryBindingType()
}

func (s *MaintenanceWindowEntry) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for MaintenanceWindowEntry._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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
	DayOfWeek   *string
	HourOfDay   *int64
	DurationMin *int64
	Version     *int64
}

const MaintenanceWindowGet_DAY_OF_WEEK_SUNDAY = "SUNDAY"
const MaintenanceWindowGet_DAY_OF_WEEK_MONDAY = "MONDAY"
const MaintenanceWindowGet_DAY_OF_WEEK_TUESDAY = "TUESDAY"
const MaintenanceWindowGet_DAY_OF_WEEK_WEDNESDAY = "WEDNESDAY"
const MaintenanceWindowGet_DAY_OF_WEEK_THURSDAY = "THURSDAY"
const MaintenanceWindowGet_DAY_OF_WEEK_FRIDAY = "FRIDAY"
const MaintenanceWindowGet_DAY_OF_WEEK_SATURDAY = "SATURDAY"

func (s *MaintenanceWindowGet) GetType__() vapiBindings_.BindingType {
	return MaintenanceWindowGetBindingType()
}

func (s *MaintenanceWindowGet) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for MaintenanceWindowGet._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type ManagementGatewayTemplate struct {
	PublicIp      *SddcPublicIp
	PrimaryDns    *string
	SecondaryDns  *string
	FirewallRules []FirewallRule
	Vpns          []Vpn
	// mgw network subnet cidr
	SubnetCidr *string
}

func (s *ManagementGatewayTemplate) GetType__() vapiBindings_.BindingType {
	return ManagementGatewayTemplateBindingType()
}

func (s *ManagementGatewayTemplate) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ManagementGatewayTemplate._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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

func (s *MapZonesRequest) GetType__() vapiBindings_.BindingType {
	return MapZonesRequestBindingType()
}

func (s *MapZonesRequest) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for MapZonesRequest._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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

func (s *Metadata) GetType__() vapiBindings_.BindingType {
	return MetadataBindingType()
}

func (s *Metadata) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for Metadata._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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
	// Flag to identify if it is Academic Standard or Commercial Standard License.
	AcademicLicense *bool
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

func (s *MsftLicensingConfig) GetType__() vapiBindings_.BindingType {
	return MsftLicensingConfigBindingType()
}

func (s *MsftLicensingConfig) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for MsftLicensingConfig._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type NatRule struct {
	RuleType      *string
	Protocol      *string
	Name          *string
	InternalPorts *string
	PublicPorts   *string
	PublicIp      *string
	InternalIp    *string
	// Possible values are:
	//
	// * NatRule#NatRule_ACTION_DNAT
	// * NatRule#NatRule_ACTION_SNAT
	Action *string
	Id     *string
	// current revision of the list of nat rules, used to protect against concurrent modification (first writer wins) format: int32
	Revision *int64
}

const NatRule_ACTION_DNAT = "dnat"
const NatRule_ACTION_SNAT = "snat"

func (s *NatRule) GetType__() vapiBindings_.BindingType {
	return NatRuleBindingType()
}

func (s *NatRule) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for NatRule._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type NativeSpecKeyMaterial struct {
	KeyId *string
}

func (s *NativeSpecKeyMaterial) GetType__() vapiBindings_.BindingType {
	return NativeSpecKeyMaterialBindingType()
}

func (s *NativeSpecKeyMaterial) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for NativeSpecKeyMaterial._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type NetworkTemplate struct {
	ManagementGatewayTemplates []ManagementGatewayTemplate
	ComputeGatewayTemplates    []ComputeGatewayTemplate
}

func (s *NetworkTemplate) GetType__() vapiBindings_.BindingType {
	return NetworkTemplateBindingType()
}

func (s *NetworkTemplate) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for NetworkTemplate._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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

func (s *NewCredentials) GetType__() vapiBindings_.BindingType {
	return NewCredentialsBindingType()
}

func (s *NewCredentials) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for NewCredentials._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Details the state of different NSX add-ons.
type NsxtAddons struct {
	// Indicates whether NSX Advanced addon is enabled or disabled.
	EnableNsxAdvancedAddon *bool
}

func (s *NsxtAddons) GetType__() vapiBindings_.BindingType {
	return NsxtAddonsBindingType()
}

func (s *NsxtAddons) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for NsxtAddons._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Holder for the offer instances.
type OfferInstancesHolder struct {
	OnDemand OnDemandOfferInstance
	Offers   []TermOfferInstance
}

func (s *OfferInstancesHolder) GetType__() vapiBindings_.BindingType {
	return OfferInstancesHolderBindingType()
}

func (s *OfferInstancesHolder) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for OfferInstancesHolder._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Holder for the on-demand offer instance.
type OnDemandOfferInstance struct {
	Product string
	// Deprecated. Please use product and type fields instead.
	ProductType *string
	Name        string
	Currency    string
	Region      string
	UnitPrice   string
	MonthlyCost string
	Version     string
	Type_       string
	Description string
}

func (s *OnDemandOfferInstance) GetType__() vapiBindings_.BindingType {
	return OnDemandOfferInstanceBindingType()
}

func (s *OnDemandOfferInstance) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for OnDemandOfferInstance._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type OrgProperties struct {
	// A map of string properties to values.
	Values map[string]string
}

func (s *OrgProperties) GetType__() vapiBindings_.BindingType {
	return OrgPropertiesBindingType()
}

func (s *OrgProperties) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for OrgProperties._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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

func (s *OrgSellerInfo) GetType__() vapiBindings_.BindingType {
	return OrgSellerInfoBindingType()
}

func (s *OrgSellerInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for OrgSellerInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type Organization struct {
	Updated time.Time
	// User id that last updated this record
	UserId string
	// User id that last updated this record
	UpdatedByUserId *string
	Created         time.Time
	// Version of this entity format: int32
	Version int64
	// User name that last updated this record
	UpdatedByUserName *string
	// User name that last updated this record
	UserName *string
	// Unique ID for this entity
	Id string
	// ORG_TYPE to be associated with the org
	OrgType       *string
	DisplayName   *string
	Name          *string
	OrgSellerInfo *OrgSellerInfo
	// Possible values are:
	//
	// * Organization#Organization_PROJECT_STATE_CREATED
	// * Organization#Organization_PROJECT_STATE_DELETED
	ProjectState *string
	Properties   *OrgProperties
}

const Organization_PROJECT_STATE_CREATED = "CREATED"
const Organization_PROJECT_STATE_DELETED = "DELETED"

func (s *Organization) GetType__() vapiBindings_.BindingType {
	return OrganizationBindingType()
}

func (s *Organization) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for Organization._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type OutpostConfig struct {
	// Outpost ID
	OutpostId *string
}

func (s *OutpostConfig) GetType__() vapiBindings_.BindingType {
	return OutpostConfigBindingType()
}

func (s *OutpostConfig) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for OutpostConfig._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type PartitionPlacementGroupInfo struct {
	PartitionGroupNames []string
	// The availability zone of the placement group.
	AvailabilityZone *string
}

func (s *PartitionPlacementGroupInfo) GetType__() vapiBindings_.BindingType {
	return PartitionPlacementGroupInfoBindingType()
}

func (s *PartitionPlacementGroupInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for PartitionPlacementGroupInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type PaymentMethodInfo struct {
	Type_           *string
	DefaultFlag     *bool
	PaymentMethodId *string
}

func (s *PaymentMethodInfo) GetType__() vapiBindings_.BindingType {
	return PaymentMethodInfoBindingType()
}

func (s *PaymentMethodInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for PaymentMethodInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type PopAgentXeniConnection struct {
	// The gateway route ip fo the subnet.
	DefaultSubnetRoute *string
	EniInfo            *EniInfo
}

func (s *PopAgentXeniConnection) GetType__() vapiBindings_.BindingType {
	return PopAgentXeniConnectionBindingType()
}

func (s *PopAgentXeniConnection) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for PopAgentXeniConnection._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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

func (s *PopAmiInfo) GetType__() vapiBindings_.BindingType {
	return PopAmiInfoBindingType()
}

func (s *PopAmiInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for PopAmiInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Present a SDDC PoP information.
type PopInfo struct {
	// A map of [region name of PoP / PoP-AMI]:[PopAmiInfo].
	AmiInfos map[string]PopAmiInfo
	// version of the manifest.
	ManifestVersion *string
	// A map of [service type]:[PopServiceInfo]
	ServiceInfos map[string]PopServiceInfo
	// The PopInfo (or PoP AMI) created time. Using ISO 8601 date-time pattern. format: date-time
	CreatedAt *time.Time
	// Possible values are:
	//
	// * PopInfo#PopInfo_OS_TYPE_CENTOS
	// * PopInfo#PopInfo_OS_TYPE_AMAZON
	// * PopInfo#PopInfo_OS_TYPE_PHOTON
	//
	//  Type of OS: CENTOS or AMAZON(Amazon Linux 2) or PHOTON
	OsType *string
	// UUID of the PopInfo format: UUID
	Id *string
}

const PopInfo_OS_TYPE_CENTOS = "CENTOS"
const PopInfo_OS_TYPE_AMAZON = "AMAZON"
const PopInfo_OS_TYPE_PHOTON = "PHOTON"

func (s *PopInfo) GetType__() vapiBindings_.BindingType {
	return PopInfoBindingType()
}

func (s *PopInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for PopInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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
	// PoP related services (including os platform and JRE).
	Service *string
}

func (s *PopServiceInfo) GetType__() vapiBindings_.BindingType {
	return PopServiceInfoBindingType()
}

func (s *PopServiceInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for PopServiceInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Represents a structure for powered by instance type configuration.
type PoweredByInstanceTypeConfig struct {
	// Boolean to indicate whether clusters and SDDCs using this instance type support vSAN ESA.
	VsanEsaSupported *bool
	// Display name of instance_type.
	DisplayName *string
	// Description of the instance_type.
	Description *string
	// Instance type name.
	InstanceType *string
	// The capacity of the given instance type.
	EntityCapacity *InstanceEntityCapacity
	// Label for instance_type.
	Label *string
	// Error due to which instance provisioning is restricted on the cluster.
	InstanceProvisioningErrorCause *string
	// Map of host count vs valid cpu core values for the given instance type.
	SupportedCpuCores map[string][]int64
	// Array of number of hosts allowed for this operation. Range of hosts user can select for sddc provision
	Hosts []int64
	// Array of valid cpu cores values for the given instance type.
	CpuCores []int64
	// Boolean to indicate whether hyperThreading is supported for an instance type.
	HyperThreadingSupported *bool
	// ID of the powered by configuration.
	PoweredById *string
	// Possible values are:
	//
	// * PoweredByInstanceTypeConfig#PoweredByInstanceTypeConfig_POWERED_BY_TYPE_OUTPOST
	//
	//  Represents the type of powered by instance type configuration.
	PoweredByType *string
}

const PoweredByInstanceTypeConfig_POWERED_BY_TYPE_OUTPOST = "OUTPOST"

func (s *PoweredByInstanceTypeConfig) GetType__() vapiBindings_.BindingType {
	return PoweredByInstanceTypeConfigBindingType()
}

func (s *PoweredByInstanceTypeConfig) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for PoweredByInstanceTypeConfig._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Represents a provisioning spec for a sddc
type ProvisionSpec struct {
	// Map of provider to sddc config spec
	Provider map[string]SddcConfigSpec
}

func (s *ProvisionSpec) GetType__() vapiBindings_.BindingType {
	return ProvisionSpecBindingType()
}

func (s *ProvisionSpec) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ProvisionSpec._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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
	ResolvedAt   *time.Time
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

func (s *RequestDetail) GetType__() vapiBindings_.BindingType {
	return RequestDetailBindingType()
}

func (s *RequestDetail) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for RequestDetail._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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

func (s *Reservation) GetType__() vapiBindings_.BindingType {
	return ReservationBindingType()
}

func (s *Reservation) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for Reservation._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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

func (s *ReservationInMw) GetType__() vapiBindings_.BindingType {
	return ReservationInMwBindingType()
}

func (s *ReservationInMw) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ReservationInMw._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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
	DayOfWeek      *string
	HourOfDay      *int64
	DurationMin    *int64
	Version        *int64
	Reservations   []Reservation
	ReservationsMw []ReservationInMw
}

const ReservationSchedule_DAY_OF_WEEK_SUNDAY = "SUNDAY"
const ReservationSchedule_DAY_OF_WEEK_MONDAY = "MONDAY"
const ReservationSchedule_DAY_OF_WEEK_TUESDAY = "TUESDAY"
const ReservationSchedule_DAY_OF_WEEK_WEDNESDAY = "WEDNESDAY"
const ReservationSchedule_DAY_OF_WEEK_THURSDAY = "THURSDAY"
const ReservationSchedule_DAY_OF_WEEK_FRIDAY = "FRIDAY"
const ReservationSchedule_DAY_OF_WEEK_SATURDAY = "SATURDAY"

func (s *ReservationSchedule) GetType__() vapiBindings_.BindingType {
	return ReservationScheduleBindingType()
}

func (s *ReservationSchedule) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ReservationSchedule._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type ReservationWindow struct {
	DurationHours                 *int64
	EstimatedDurationHoursMinimum *int64
	// Possible values are:
	//
	// * ReservationWindow#ReservationWindow_RESERVATION_STATE_SCHEDULED
	// * ReservationWindow#ReservationWindow_RESERVATION_STATE_RUNNING
	// * ReservationWindow#ReservationWindow_RESERVATION_STATE_CANCELED
	// * ReservationWindow#ReservationWindow_RESERVATION_STATE_COMPLETED
	// * ReservationWindow#ReservationWindow_RESERVATION_STATE_TERMINATED
	ReservationState              *string
	Emergency                     *bool
	MaintenanceProperties         *ReservationWindowMaintenanceProperties
	ManifestId                    *string
	StartHour                     *int64
	SddcId                        *string
	StartDate                     *string
	EstimatedDurationHoursMaximum *int64
	ReserveId                     *string
	// Metadata for reservation window, in key-value form
	Metadata map[string]string
}

const ReservationWindow_RESERVATION_STATE_SCHEDULED = "SCHEDULED"
const ReservationWindow_RESERVATION_STATE_RUNNING = "RUNNING"
const ReservationWindow_RESERVATION_STATE_CANCELED = "CANCELED"
const ReservationWindow_RESERVATION_STATE_COMPLETED = "COMPLETED"
const ReservationWindow_RESERVATION_STATE_TERMINATED = "TERMINATED"

func (s *ReservationWindow) GetType__() vapiBindings_.BindingType {
	return ReservationWindowBindingType()
}

func (s *ReservationWindow) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ReservationWindow._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type ReservationWindowMaintenanceProperties struct {
	// Status of upgrade, if any
	Status *string
}

func (s *ReservationWindowMaintenanceProperties) GetType__() vapiBindings_.BindingType {
	return ReservationWindowMaintenancePropertiesBindingType()
}

func (s *ReservationWindowMaintenanceProperties) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ReservationWindowMaintenanceProperties._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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

func (s *RouteTableInfo) GetType__() vapiBindings_.BindingType {
	return RouteTableInfoBindingType()
}

func (s *RouteTableInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for RouteTableInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type Sddc struct {
	Updated time.Time
	// User id that last updated this record
	UserId string
	// User id that last updated this record
	UpdatedByUserId *string
	Created         time.Time
	// Version of this entity format: int32
	Version int64
	// User name that last updated this record
	UpdatedByUserName *string
	// User name that last updated this record
	UserName *string
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
	OrgId          *string
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
	ResourceConfig  *AwsSddcResourceConfig
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

func (s *Sddc) GetType__() vapiBindings_.BindingType {
	return SddcBindingType()
}

func (s *Sddc) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for Sddc._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type SddcChoice struct {
	// SDDC display name
	DisplayText *string
	// SDDC id
	Value *string
}

func (s *SddcChoice) GetType__() vapiBindings_.BindingType {
	return SddcChoiceBindingType()
}

func (s *SddcChoice) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for SddcChoice._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type SddcConfig struct {
	// Indicates the desired licensing support, if any, of Microsoft software.
	MsftLicenseConfig *MsftLicensingConfig
	// The account linking configuration, we will keep this one and remove accountLinkSddcConfig finally.
	AccountLinkConfig *AccountLinkConfig
	// AWS VPC IP range. Only prefix of 16 or 20 is currently supported.
	VpcCidr *string
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
	// Outpost ID of the SDDC. Used only for outpost deployments.
	OutpostId *string
	// The storage capacity value to be requested for the sddc primary cluster, in GiBs. If provided, instead of using the direct-attached storage, a capacity value amount of seperable storage will be used. format: int64
	StorageCapacity *int64
	Name            string
	// A list of the SDDC linking configurations to use.
	AccountLinkSddcConfig []AccountLinkSddcConfig
	// Possible values are:
	//
	// * SddcConfig#SddcConfig_HOST_INSTANCE_TYPE_I3_METAL
	// * SddcConfig#SddcConfig_HOST_INSTANCE_TYPE_R5_METAL
	// * SddcConfig#SddcConfig_HOST_INSTANCE_TYPE_I3EN_METAL
	// * SddcConfig#SddcConfig_HOST_INSTANCE_TYPE_I4I_METAL
	// * SddcConfig#SddcConfig_HOST_INSTANCE_TYPE_C6I_METAL
	// * SddcConfig#SddcConfig_HOST_INSTANCE_TYPE_M7I_METAL_48XL
	// * SddcConfig#SddcConfig_HOST_INSTANCE_TYPE_M7I_METAL_24XL
	//
	//  The instance type for the esx hosts in the primary cluster of the SDDC.
	HostInstanceType *string
	// Identifier of the billing account in subscription.
	BillingAccountId *string
	Region           *string
	// If provided, will be assigned as SDDC id of the provisioned SDDC. format: UUID
	SddcId *string
	// If provided, configuration from the template will applied to the provisioned SDDC. format: UUID
	SddcTemplateId *string
	// Denotes the sddc type , if the value is null or empty, the type is considered as default.
	SddcType *string
	// Possible values are:
	//
	// * SddcConfig#SddcConfig_VSAN_VERSION_VSAN1
	// * SddcConfig#SddcConfig_VSAN_VERSION_VSANESA
	// * SddcConfig#SddcConfig_VSAN_VERSION_NOVSAN
	//
	//  The vSAN version to be used in the SDDC's primary cluster.
	VsanVersion *string
	// Possible values are:
	//
	// * SddcConfig#SddcConfig_PROVIDER_AWS
	// * SddcConfig#SddcConfig_PROVIDER_ZEROCLOUD
	//
	//  Determines what additional properties are available based on cloud provider.
	Provider string
	// The SSO domain name to use for vSphere users. If not specified, vmc.local will be used.
	SsoDomain *string
	NumHosts  int64
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
const SddcConfig_HOST_INSTANCE_TYPE_I3_METAL = "i3.metal"
const SddcConfig_HOST_INSTANCE_TYPE_R5_METAL = "r5.metal"
const SddcConfig_HOST_INSTANCE_TYPE_I3EN_METAL = "i3en.metal"
const SddcConfig_HOST_INSTANCE_TYPE_I4I_METAL = "i4i.metal"
const SddcConfig_HOST_INSTANCE_TYPE_C6I_METAL = "c6i.metal"
const SddcConfig_HOST_INSTANCE_TYPE_M7I_METAL_48XL = "m7i.metal-48xl"
const SddcConfig_HOST_INSTANCE_TYPE_M7I_METAL_24XL = "m7i.metal-24xl"
const SddcConfig_VSAN_VERSION_VSAN1 = "vsan1"
const SddcConfig_VSAN_VERSION_VSANESA = "vsanesa"
const SddcConfig_VSAN_VERSION_NOVSAN = "novsan"
const SddcConfig_PROVIDER_AWS = "AWS"
const SddcConfig_PROVIDER_ZEROCLOUD = "ZEROCLOUD"
const SddcConfig_DEPLOYMENT_TYPE_SINGLEAZ = "SingleAZ"
const SddcConfig_DEPLOYMENT_TYPE_MULTIAZ = "MultiAZ"

func (s *SddcConfig) GetType__() vapiBindings_.BindingType {
	return SddcConfigBindingType()
}

func (s *SddcConfig) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for SddcConfig._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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

func (s *SddcConfigSpec) GetType__() vapiBindings_.BindingType {
	return SddcConfigSpecBindingType()
}

func (s *SddcConfigSpec) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for SddcConfigSpec._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Convert request body for SDDC
type SddcConvertRequest struct {
	// The total number of hosts in the SDDC after conversion.
	NumHosts *int64
}

func (s *SddcConvertRequest) GetType__() vapiBindings_.BindingType {
	return SddcConvertRequestBindingType()
}

func (s *SddcConvertRequest) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for SddcConvertRequest._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type SddcId struct {
	// Sddc ID
	SddcId *string
}

func (s *SddcId) GetType__() vapiBindings_.BindingType {
	return SddcIdBindingType()
}

func (s *SddcId) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for SddcId._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type SddcInput struct {
	Inputs []SddcList
}

func (s *SddcInput) GetType__() vapiBindings_.BindingType {
	return SddcInputBindingType()
}

func (s *SddcInput) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for SddcInput._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type SddcLinkConfig struct {
	CustomerSubnetIds []string
	// Determines which connected customer account to link to
	ConnectedAccountId *string
}

func (s *SddcLinkConfig) GetType__() vapiBindings_.BindingType {
	return SddcLinkConfigBindingType()
}

func (s *SddcLinkConfig) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for SddcLinkConfig._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type SddcList struct {
	Label       *string
	Required    *bool
	Placeholder *string
	Name        *string
	Choice      []SddcChoice
}

func (s *SddcList) GetType__() vapiBindings_.BindingType {
	return SddcListBindingType()
}

func (s *SddcList) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for SddcList._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Describes software components of the installed SDDC
type SddcManifest struct {
	// the vmcVersion of the sddc for display
	VmcVersion *string
	GlcmBundle *GlcmBundle
	PopInfo    *PopInfo
	// the vmcInternalVersion of the sddc for internal use
	VmcInternalVersion  *string
	EbsBackedVsanConfig *EbsBackedVsanConfig
	VsanWitnessAmi      *AmiInfo
	EsxAmi              *AmiInfo
	EsxNsxtAmi          *AmiInfo
	Metadata            *Metadata
}

func (s *SddcManifest) GetType__() vapiBindings_.BindingType {
	return SddcManifestBindingType()
}

func (s *SddcManifest) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for SddcManifest._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Patch request body for SDDC
type SddcPatchRequest struct {
	// The new name of the SDDC to be changed to.
	Name *string
}

func (s *SddcPatchRequest) GetType__() vapiBindings_.BindingType {
	return SddcPatchRequestBindingType()
}

func (s *SddcPatchRequest) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for SddcPatchRequest._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type SddcPublicIp struct {
	PublicIp            string
	Name                *string
	AllocationId        *string
	DnatRuleId          *string
	AssociatedPrivateIp *string
	SnatRuleId          *string
}

func (s *SddcPublicIp) GetType__() vapiBindings_.BindingType {
	return SddcPublicIpBindingType()
}

func (s *SddcPublicIp) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for SddcPublicIp._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type SddcResourceConfig struct {
	// (deprecated)
	VcIp *string
	// Name for management appliance network.
	MgmtApplianceNetworkName *string
	// URL of the NSX Manager
	NsxMgrUrl *string
	// This flag determines whether vLCM is enabled on this sddc or not.
	VlcmEnabled *bool
	// NSX cloud audit Password
	NsxCloudAuditPassword *string
	// vCenter to csp federation status.
	VcCspLoginStatus *string
	// NSX cloud admin password
	NsxCloudAdminPassword *string
	// The ManagedObjectReference of the management Datastore
	ManagementDs *string
	// nsx api entire base url
	NsxApiPublicEndpointUrl *string
	// Nfs Mode Flag, for nfs mounting.
	NfsMode *bool
	// Password for vCenter SDDC administrator
	CloudPassword *string
	SddcNetworks  []string
	// List of clusters in the SDDC.
	Clusters []Cluster
	// Username for vCenter SDDC administrator
	CloudUsername *string
	// Possible values are:
	//
	// * SddcResourceConfig#SddcResourceConfig_DEPLOYMENT_TYPE_SINGLE_AZ
	// * SddcResourceConfig#SddcResourceConfig_DEPLOYMENT_TYPE_MULTI_AZ
	//
	//  Denotes if this is a SingleAZ SDDC or a MultiAZ SDDC.
	DeploymentType         *string
	PopAgentXeniConnection *PopAgentXeniConnection
	// NSX Manager internal management IP
	NsxMgrManagementIp *string
	// NSX cloud audit user name
	NsxCloudAudit *string
	// Cluster Id to add ESX workflow
	EsxClusterId *string
	// Management Gateway Id
	MgwId *string
	// URL of the vCenter server
	VcUrl    *string
	EsxHosts []AwsEsxHost
	// Group name for vCenter SDDC administrator
	CloudUserGroup *string
	ManagementRp   *string
	// Availability zone where the witness node is provisioned for a MultiAZ SDDC. This is null for a SingleAZ SDDC.
	WitnessAvailabilityZone *string
	// Whether this sddc is maintained by its desired state documents.
	SddcDesiredState *bool
	SddcSize         *SddcSize
	// This flag determines whether CVDS is enabled on this sddc or not.
	CvdsEnabled *bool
	// List of Controller IPs
	NsxControllerIps []string
	// Marks that the SDDC VC should be deployed with two hostnames.
	TwoHostnameVcDeployment *bool
	// ESX host subnet
	EsxHostSubnet *string
	// The SSO domain name to use for vSphere users
	SsoDomain *string
	// region in which sddc is deployed
	Region *string
	// if true, use the private IP addresses to register DNS records for the management VMs
	DnsWithManagementVmPrivateIp *bool
	// vCenter public IP
	VcPublicIp *string
	// (deprecated)
	PscIp *string
	// if true, NSX-T UI is enabled.
	Nsxt *bool
	// Key provider data.
	KeyProvider []KeyProviderData
	// PSC internal management IP
	PscManagementIp *string
	// URL of the PSC server
	PscUrl *string
	Cgws   []string
	// Availability zones over which esx hosts are provisioned. MultiAZ SDDCs will have hosts provisioned over two availability zones while SingleAZ SDDCs will provision over one.
	AvailabilityZones []string
	// Mark if Containerized Permissions has been enabled on vCenter.
	VcContainerizedPermissionsEnabled *bool
	CustomProperties                  map[string]string
	// Possible values are:
	//
	// * SddcResourceConfig#SddcResourceConfig_PROVIDER_AWS
	// * SddcResourceConfig#SddcResourceConfig_PROVIDER_ZEROCLOUD
	//
	//  Discriminator for additional properties
	Provider string
	// vCenter internal management IP
	VcManagementIp *string
	// The Microsoft license status of this SDDC.
	MsftLicenseConfig *MsftLicensingConfig
	// NSX-T Native Oauth client for UI.
	NsxNativeClient *CspOauthClient
	// unique id of the vCenter server
	VcInstanceId *string
	// oAuth client for enabling federation on vCenter.
	VcOauthClient *CspOauthClient
	// skip creating vxlan for compute gateway for SDDC provisioning
	SkipCreatingVxlan *bool
	SddcManifest      *SddcManifest
	// VXLAN IP subnet
	VxlanSubnet  *string
	SddcSecurity *SddcSecurity
	// sddc identifier
	SddcId *string
	// Outpost configuration of this SDDC.
	OutpostConfig *OutpostConfig
	// URL of the NSX Manager UI login for local user access
	NsxMgrLoginUrl *string
	// Break-glass URL for non-federated login.
	VcBreakGlassUrl *string
	// NSX cloud admin user name
	NsxCloudAdmin *string
	NsxtAddons    *NsxtAddons
}

// Identifier denoting this class, when it is used in polymorphic context.
//
// This value should be assigned to the property which is used to discriminate the actual type used in the polymorphic context.
const SddcResourceConfig__TYPE_IDENTIFIER = "SddcResourceConfig"
const SddcResourceConfig_DEPLOYMENT_TYPE_SINGLE_AZ = "SINGLE_AZ"
const SddcResourceConfig_DEPLOYMENT_TYPE_MULTI_AZ = "MULTI_AZ"
const SddcResourceConfig_PROVIDER_AWS = "AWS"
const SddcResourceConfig_PROVIDER_ZEROCLOUD = "ZEROCLOUD"

func (s *SddcResourceConfig) GetType__() vapiBindings_.BindingType {
	return SddcResourceConfigBindingType()
}

func (s *SddcResourceConfig) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for SddcResourceConfig._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Details the sddc securtity profiles and Harden state.
type SddcSecurity struct {
	// Possible values are:
	//
	// * SddcSecurity#SddcSecurity_PROFILE_DEFAULT
	// * SddcSecurity#SddcSecurity_PROFILE_PCI_COMMERCIAL
	//
	//  Security Profiles for SDDC Hardening.
	Profile *string
	// Indicates whether SDDC is hardened or not.
	Hardened *bool
}

const SddcSecurity_PROFILE_DEFAULT = "PROFILE_DEFAULT"
const SddcSecurity_PROFILE_PCI_COMMERCIAL = "PROFILE_PCI_COMMERCIAL"

func (s *SddcSecurity) GetType__() vapiBindings_.BindingType {
	return SddcSecurityBindingType()
}

func (s *SddcSecurity) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for SddcSecurity._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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

func (s *SddcSize) GetType__() vapiBindings_.BindingType {
	return SddcSizeBindingType()
}

func (s *SddcSize) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for SddcSize._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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

func (s *SddcStateRequest) GetType__() vapiBindings_.BindingType {
	return SddcStateRequestBindingType()
}

func (s *SddcStateRequest) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for SddcStateRequest._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type SddcTemplate struct {
	Updated time.Time
	// User id that last updated this record
	UserId string
	// User id that last updated this record
	UpdatedByUserId *string
	Created         time.Time
	// Version of this entity format: int32
	Version int64
	// User name that last updated this record
	UpdatedByUserName *string
	// User name that last updated this record
	UserName *string
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
	State           *string
	NetworkTemplate *NetworkTemplate
	// name for SDDC configuration template
	Name         *string
	SourceSddcId *string
	OrgId        *string
	Sddc         *Sddc
}

const SddcTemplate_STATE_INITIALIZATION = "INITIALIZATION"
const SddcTemplate_STATE_AVAILABLE = "AVAILABLE"
const SddcTemplate_STATE_INUSE = "INUSE"
const SddcTemplate_STATE_APPLIED = "APPLIED"
const SddcTemplate_STATE_DELETING = "DELETING"
const SddcTemplate_STATE_DELETED = "DELETED"
const SddcTemplate_STATE_FAILED = "FAILED"

func (s *SddcTemplate) GetType__() vapiBindings_.BindingType {
	return SddcTemplateBindingType()
}

func (s *SddcTemplate) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for SddcTemplate._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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

type ServiceQuotaRequest struct {
	Updated time.Time
	// User id that last updated this record
	UserId string
	// User id that last updated this record
	UpdatedByUserId *string
	Created         time.Time
	// Version of this entity format: int32
	Version int64
	// User name that last updated this record
	UpdatedByUserName *string
	// User name that last updated this record
	UserName *string
	// Unique ID for this entity
	Id             string
	RequesterEmail *string
	// The task for running the service quota request.
	TaskId *string
	// Region for the service quota
	Region           *string
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

func (s *ServiceQuotaRequest) GetType__() vapiBindings_.BindingType {
	return ServiceQuotaRequestBindingType()
}

func (s *ServiceQuotaRequest) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ServiceQuotaRequest._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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

func (s *Site) GetType__() vapiBindings_.BindingType {
	return SiteBindingType()
}

func (s *Site) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for Site._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// (as there's already one SubnetInfo, use Subnet instead)
type Subnet struct {
	// subnet id
	SubnetId *string
	// subnet name
	Name        *string
	RouteTables []SubnetRouteTableInfo
}

func (s *Subnet) GetType__() vapiBindings_.BindingType {
	return SubnetBindingType()
}

func (s *Subnet) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for Subnet._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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

func (s *SubnetInfo) GetType__() vapiBindings_.BindingType {
	return SubnetInfoBindingType()
}

func (s *SubnetInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for SubnetInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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

func (s *SubnetRouteTableInfo) GetType__() vapiBindings_.BindingType {
	return SubnetRouteTableInfoBindingType()
}

func (s *SubnetRouteTableInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for SubnetRouteTableInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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
	Status                 *string
	AnniversaryBillingDate *string
	EndDate                *string
	// The frequency at which the customer is billed. Currently supported values are \"Upfront\" and \"Monthly\"
	BillingFrequency      *string
	AutoRenewedAllowed    *string
	CommitmentTerm        *string
	CspSubscriptionId     *string
	BillingSubscriptionId *string
	OfferVersion          *string
	// Possible values are:
	//
	// * SubscriptionDetails#SubscriptionDetails_OFFER_TYPE_TERM
	// * SubscriptionDetails#SubscriptionDetails_OFFER_TYPE_ON_DEMAND
	OfferType   *string
	Description *string
	ProductId   *string
	Region      *string
	ProductName *string
	OfferName   *string
	// unit of measurment for commitment term
	CommitmentTermUom *string
	StartDate         *string
	Quantity          *string
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
const SubscriptionDetails_OFFER_TYPE_TERM = "TERM"
const SubscriptionDetails_OFFER_TYPE_ON_DEMAND = "ON_DEMAND"

func (s *SubscriptionDetails) GetType__() vapiBindings_.BindingType {
	return SubscriptionDetailsBindingType()
}

func (s *SubscriptionDetails) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for SubscriptionDetails._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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

func (s *SubscriptionProducts) GetType__() vapiBindings_.BindingType {
	return SubscriptionProductsBindingType()
}

func (s *SubscriptionProducts) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for SubscriptionProducts._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type SubscriptionRequest struct {
	// The product for which subscription needs to be created. Refer /vmc/api/orgs/{orgId}/products.
	Product *string
	// Old identifier for product. \*Deprecarted\*. See product and type
	ProductType string
	ProductId   *string
	// Frequency of the billing.
	BillingFrequency *string
	Region           string
	CommitmentTerm   string
	OfferContextId   *string
	OfferVersion     string
	OfferName        string
	Quantity         int64
	// The type of the product for which the subscription needs to be created.
	Type_           *string
	Price           *int64
	ProductChargeId *string
}

func (s *SubscriptionRequest) GetType__() vapiBindings_.BindingType {
	return SubscriptionRequestBindingType()
}

func (s *SubscriptionRequest) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for SubscriptionRequest._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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
	Seats    *int64
	// SDDCs in this window format: UUID
	Sddcs           []string
	DurationHours   *int64
	StartHour       *int64
	SupportWindowId *string
	Metadata        *vapiData_.StructValue
}

const SupportWindow_START_DAY_MONDAY = "MONDAY"
const SupportWindow_START_DAY_TUESDAY = "TUESDAY"
const SupportWindow_START_DAY_WEDNESDAY = "WEDNESDAY"
const SupportWindow_START_DAY_THURSDAY = "THURSDAY"
const SupportWindow_START_DAY_FRIDAY = "FRIDAY"
const SupportWindow_START_DAY_SATURDAY = "SATURDAY"
const SupportWindow_START_DAY_SUNDAY = "SUNDAY"

func (s *SupportWindow) GetType__() vapiBindings_.BindingType {
	return SupportWindowBindingType()
}

func (s *SupportWindow) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for SupportWindow._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type SupportWindowId struct {
	// Support Window ID
	WindowId *string
}

func (s *SupportWindowId) GetType__() vapiBindings_.BindingType {
	return SupportWindowIdBindingType()
}

func (s *SupportWindowId) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for SupportWindowId._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type Tag struct {
	Key   *string
	Value *string
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

type Task struct {
	Updated time.Time
	// User id that last updated this record
	UserId string
	// User id that last updated this record
	UpdatedByUserId *string
	Created         time.Time
	// Version of this entity format: int32
	Version int64
	// User name that last updated this record
	UpdatedByUserName *string
	// User name that last updated this record
	UserName *string
	// Unique ID for this entity
	Id string
	// Possible values are:
	//
	// * Task#Task_STATUS_STARTED
	// * Task#Task_STATUS_CANCELING
	// * Task#Task_STATUS_FINISHED
	// * Task#Task_STATUS_FAILED
	// * Task#Task_STATUS_CANCELED
	Status                *string
	LocalizedErrorMessage *string
	// UUID of the resource the task is acting upon
	ResourceId *string
	// If this task was created by another task - this provides the linkage. Mostly for debugging.
	ParentTaskId *string
	TaskVersion  *string
	// (Optional) Client provided uniqifier to make task creation idempotent. Be aware not all tasks support this. For tasks that do - supplying the same correlation Id, for the same task type, within a predefined window will ensure the operation happens at most once.
	CorrelationId *string
	// Entity version of the resource at the start of the task. This is only set for some task types. format: int32
	StartResourceEntityVersion *int64
	CustomerErrorMessage       *string
	SubStatus                  *string
	TaskType                   *string
	StartTime                  *time.Time
	// Task progress phases involved in current task execution
	TaskProgressPhases []TaskProgressPhase
	ErrorMessage       *string
	OrgId              *string
	// Entity version of the resource at the end of the task. This is only set for some task types. format: int32
	EndResourceEntityVersion *int64
	// Service errors returned from SDDC services.
	ServiceErrors []ServiceError
	OrgType       *string
	// Estimated remaining time in minute of the task execution, < 0 means no estimation for the task. format: int32
	EstimatedRemainingMinutes *int64
	Params                    *vapiData_.StructValue
	// Estimated progress percentage the task executed format: int32
	ProgressPercent *int64
	// The current in progress phase ID in the task execution, if none in progress, empty string returned.
	PhaseInProgress *string
	// Type of resource being acted upon
	ResourceType *string
	EndTime      *time.Time
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

type TenantServiceInfo struct {
	// S3 bucket arn used by log forwarder.
	S3LogBucketArn *string
}

func (s *TenantServiceInfo) GetType__() vapiBindings_.BindingType {
	return TenantServiceInfoBindingType()
}

func (s *TenantServiceInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for TenantServiceInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Holder for term billing options.
type TermBillingOptions struct {
	UnitPrice        *string
	BillingFrequency *string
}

func (s *TermBillingOptions) GetType__() vapiBindings_.BindingType {
	return TermBillingOptionsBindingType()
}

func (s *TermBillingOptions) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for TermBillingOptions._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Holder for the term offer instances.
type TermOfferInstance struct {
	Description string
	Product     string
	// Deprecated. Please use product and type fields instead.
	ProductType    *string
	Name           string
	Currency       string
	Region         string
	CommitmentTerm int64
	// (deprecated. unit_price is moved into TermBillingOptions. For backward compatibility, this field reflect \"Prepaid\" price at the offer level.)
	UnitPrice       string
	BillingOptions  []TermBillingOptions
	Version         string
	OfferContextId  *string
	ProductChargeId *string
	Type_           string
	ProductId       *string
}

func (s *TermOfferInstance) GetType__() vapiBindings_.BindingType {
	return TermOfferInstanceBindingType()
}

func (s *TermOfferInstance) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for TermOfferInstance._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type TermsOfServiceResult struct {
	// The terms of service ID requested.
	TermsId *string
	// Whether or not the terms requested have been signed.
	Signed *bool
}

func (s *TermsOfServiceResult) GetType__() vapiBindings_.BindingType {
	return TermsOfServiceResultBindingType()
}

func (s *TermsOfServiceResult) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for TermsOfServiceResult._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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

func (s *UpdateCredentials) GetType__() vapiBindings_.BindingType {
	return UpdateCredentialsBindingType()
}

func (s *UpdateCredentials) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for UpdateCredentials._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type VmcLocale struct {
	// The locale to be used for translating responses for the session
	Locale *string
}

func (s *VmcLocale) GetType__() vapiBindings_.BindingType {
	return VmcLocaleBindingType()
}

func (s *VmcLocale) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for VmcLocale._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type VpcInfo struct {
	EsxSecurityGroupId       *string
	VpcCidr                  *string
	VgwId                    *string
	EsxPublicSecurityGroupId *string
	// set of virtual interfaces attached to the sddc
	VifIds            []string
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
	// Set of VCDR (VMware Cloud Disaster Recovery) ENIs
	VcdrEnis []EniInfo
	// (deprecated)
	SubnetId          *string
	InternetGatewayId *string
	SecurityGroupId   *string
	// (deprecated)
	AssociationId *string
	// Route table which contains the route to VGW (deprecated)
	VgwRouteTableId *string
	// List of edge vm Ips of traffic gourps added during scale-out
	TrafficGroupEdgeVmIps []string
	// Id of the association between edge subnet and route-table (deprecated)
	EdgeAssociationId *string
	Provider          *string
	// Mapping from AZ to a list of IP addresses assigned to TGW ENI that's connected with Vpc
	TgwIps map[string][]string
	// (deprecated)
	PeeringConnectionId *string
	NetworkType         *string
	AvailableZones      []AvailableZoneInfo
	// map from routeTableName to routeTableInfo
	Routetables map[string]RouteTableInfo
}

func (s *VpcInfo) GetType__() vapiBindings_.BindingType {
	return VpcInfoBindingType()
}

func (s *VpcInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for VpcInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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
	Subnets     []SubnetInfo
}

func (s *VpcInfoSubnets) GetType__() vapiBindings_.BindingType {
	return VpcInfoSubnetsBindingType()
}

func (s *VpcInfoSubnets) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for VpcInfoSubnets._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type Vpn struct {
	Version            *string
	OnPremGatewayIp    *string
	OnPremNetworkCidr  *string
	PfsEnabled         *bool
	Id                 *string
	ChannelStatus      *VpnChannelStatus
	OnPremNatIp        *string
	Name               *string
	InternalNetworkIds []string
	TunnelStatuses     []VpnTunnelStatus
	// Possible values are:
	//
	// * Vpn#Vpn_ENCRYPTION_AES
	// * Vpn#Vpn_ENCRYPTION_AES256
	// * Vpn#Vpn_ENCRYPTION_AES_GCM
	// * Vpn#Vpn_ENCRYPTION_TRIPLE_DES
	// * Vpn#Vpn_ENCRYPTION_UNKNOWN
	Encryption *string
	Enabled    *bool
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
	PreSharedKey   *string
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

type VpnChannelStatus struct {
	// Possible values are:
	//
	// * VpnChannelStatus#VpnChannelStatus_CHANNEL_STATUS_CONNECTED
	// * VpnChannelStatus#VpnChannelStatus_CHANNEL_STATUS_DISCONNECTED
	// * VpnChannelStatus#VpnChannelStatus_CHANNEL_STATUS_UNKNOWN
	ChannelStatus   *string
	ChannelState    *string
	LastInfoMessage *string
	FailureMessage  *string
}

const VpnChannelStatus_CHANNEL_STATUS_CONNECTED = "CONNECTED"
const VpnChannelStatus_CHANNEL_STATUS_DISCONNECTED = "DISCONNECTED"
const VpnChannelStatus_CHANNEL_STATUS_UNKNOWN = "UNKNOWN"

func (s *VpnChannelStatus) GetType__() vapiBindings_.BindingType {
	return VpnChannelStatusBindingType()
}

func (s *VpnChannelStatus) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for VpnChannelStatus._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type VpnTunnelStatus struct {
	OnPremSubnet    *string
	TrafficStats    *VpnTunnelTrafficStats
	LastInfoMessage *string
	LocalSubnet     *string
	TunnelState     *string
	FailureMessage  *string
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

func (s *VpnTunnelStatus) GetType__() vapiBindings_.BindingType {
	return VpnTunnelStatusBindingType()
}

func (s *VpnTunnelStatus) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for VpnTunnelStatus._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type VpnTunnelTrafficStats struct {
	PacketsOut                   *string
	PacketReceivedErrors         *string
	RxBytesOnLocalSubnet         *string
	ReplayErrors                 *string
	SequenceNumberOverFlowErrors *string
	EncryptionFailures           *string
	IntegrityErrors              *string
	PacketSentErrors             *string
	DecryptionFailures           *string
	PacketsIn                    *string
	TxBytesFromLocalSubnet       *string
}

func (s *VpnTunnelTrafficStats) GetType__() vapiBindings_.BindingType {
	return VpnTunnelTrafficStatsBindingType()
}

func (s *VpnTunnelTrafficStats) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for VpnTunnelTrafficStats._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Infomation about an available vSAN capacity in a cluster.
type VsanAvailableCapacity struct {
	Cost    string
	Quality string
	Size    int64
}

func (s *VsanAvailableCapacity) GetType__() vapiBindings_.BindingType {
	return VsanAvailableCapacityBindingType()
}

func (s *VsanAvailableCapacity) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for VsanAvailableCapacity._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Bias for reconfiguring vSAN in a cluster.
type VsanClusterReconfigBias struct {
	ShortDescription string
	FullDescription  string
	Id               string
}

func (s *VsanClusterReconfigBias) GetType__() vapiBindings_.BindingType {
	return VsanClusterReconfigBiasBindingType()
}

func (s *VsanClusterReconfigBias) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for VsanClusterReconfigBias._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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

func (s *VsanClusterReconfigConstraints) GetType__() vapiBindings_.BindingType {
	return VsanClusterReconfigConstraintsBindingType()
}

func (s *VsanClusterReconfigConstraints) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for VsanClusterReconfigConstraints._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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

func (s *VsanConfigConstraints) GetType__() vapiBindings_.BindingType {
	return VsanConfigConstraintsBindingType()
}

func (s *VsanConfigConstraints) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for VsanConfigConstraints._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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

func (s *VsanEncryptionConfig) GetType__() vapiBindings_.BindingType {
	return VsanEncryptionConfigBindingType()
}

func (s *VsanEncryptionConfig) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for VsanEncryptionConfig._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type WcpDetails struct {
	// POD CIDR. optional field.
	PodCidr *string
	// Egress CIDR
	EgressCidr *string
	// Ingress CIDR
	IngressCidr *string
	// Service CIDR
	ServiceCidr *string
	// Possible values are:
	//
	// * WcpDetails#WcpDetails_WCP_STATUS_CERTIFICATE_REVOKE_FAILED
	// * WcpDetails#WcpDetails_WCP_STATUS_CERTIFICATE_REVOKE_IN_PROGRESS
	// * WcpDetails#WcpDetails_WCP_STATUS_CERTIFICATE_UPDATE_FAILED
	// * WcpDetails#WcpDetails_WCP_STATUS_CERTIFICATE_UPDATE_IN_PROGRESS
	// * WcpDetails#WcpDetails_WCP_STATUS_CERTIFICATE_UPDATE_PENDING
	// * WcpDetails#WcpDetails_WCP_STATUS_DISABLED
	// * WcpDetails#WcpDetails_WCP_STATUS_DISABLE_FAILED
	// * WcpDetails#WcpDetails_WCP_STATUS_DISABLE_IN_PROGRESS
	// * WcpDetails#WcpDetails_WCP_STATUS_DISABLE_PENDING
	// * WcpDetails#WcpDetails_WCP_STATUS_ENABLED
	// * WcpDetails#WcpDetails_WCP_STATUS_ENABLE_FAILED
	// * WcpDetails#WcpDetails_WCP_STATUS_ENABLE_IN_PROGRESS
	// * WcpDetails#WcpDetails_WCP_STATUS_ENABLE_PENDING
	WcpStatus *string
}

const WcpDetails_WCP_STATUS_CERTIFICATE_REVOKE_FAILED = "CERTIFICATE_REVOKE_FAILED"
const WcpDetails_WCP_STATUS_CERTIFICATE_REVOKE_IN_PROGRESS = "CERTIFICATE_REVOKE_IN_PROGRESS"
const WcpDetails_WCP_STATUS_CERTIFICATE_UPDATE_FAILED = "CERTIFICATE_UPDATE_FAILED"
const WcpDetails_WCP_STATUS_CERTIFICATE_UPDATE_IN_PROGRESS = "CERTIFICATE_UPDATE_IN_PROGRESS"
const WcpDetails_WCP_STATUS_CERTIFICATE_UPDATE_PENDING = "CERTIFICATE_UPDATE_PENDING"
const WcpDetails_WCP_STATUS_DISABLED = "DISABLED"
const WcpDetails_WCP_STATUS_DISABLE_FAILED = "DISABLE_FAILED"
const WcpDetails_WCP_STATUS_DISABLE_IN_PROGRESS = "DISABLE_IN_PROGRESS"
const WcpDetails_WCP_STATUS_DISABLE_PENDING = "DISABLE_PENDING"
const WcpDetails_WCP_STATUS_ENABLED = "ENABLED"
const WcpDetails_WCP_STATUS_ENABLE_FAILED = "ENABLE_FAILED"
const WcpDetails_WCP_STATUS_ENABLE_IN_PROGRESS = "ENABLE_IN_PROGRESS"
const WcpDetails_WCP_STATUS_ENABLE_PENDING = "ENABLE_PENDING"

func (s *WcpDetails) GetType__() vapiBindings_.BindingType {
	return WcpDetailsBindingType()
}

func (s *WcpDetails) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for WcpDetails._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Infomation about witness ESX host in a cluster.
type WitnessEsx struct {
	// Possible values are:
	//
	// * WitnessEsx#WitnessEsx_ENUM_STATE_DEPLOYING
	// * WitnessEsx#WitnessEsx_ENUM_STATE_PROVISIONED
	// * WitnessEsx#WitnessEsx_ENUM_STATE_DELETING
	// * WitnessEsx#WitnessEsx_ENUM_STATE_READY
	// * WitnessEsx#WitnessEsx_ENUM_STATE_FAILED
	//
	//  The state of the witness Esx Host
	EnumState *string
	// Name of the witness Esx Host
	Name *string
	// Mac Address for witness Esx Host
	MacAddress *string
	// The unique identifier for the witness esx host format: uuid
	EsxId *string
	// Host name
	Hostname *string
	// Possible values are:
	//
	// * WitnessEsx#WitnessEsx_PROVIDER_AWS
	// * WitnessEsx#WitnessEsx_PROVIDER_ZEROCLOUD
	//
	//  The cloud provider for this witness Esx Host
	Provider string
}

// Identifier denoting this class, when it is used in polymorphic context.
//
// This value should be assigned to the property which is used to discriminate the actual type used in the polymorphic context.
const WitnessEsx__TYPE_IDENTIFIER = "WitnessEsx"
const WitnessEsx_ENUM_STATE_DEPLOYING = "DEPLOYING"
const WitnessEsx_ENUM_STATE_PROVISIONED = "PROVISIONED"
const WitnessEsx_ENUM_STATE_DELETING = "DELETING"
const WitnessEsx_ENUM_STATE_READY = "READY"
const WitnessEsx_ENUM_STATE_FAILED = "FAILED"
const WitnessEsx_PROVIDER_AWS = "AWS"
const WitnessEsx_PROVIDER_ZEROCLOUD = "ZEROCLOUD"

func (s *WitnessEsx) GetType__() vapiBindings_.BindingType {
	return WitnessEsxBindingType()
}

func (s *WitnessEsx) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for WitnessEsx._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// This class represent a mapping between eniId created on customer account to the associationId and trunkEniId of ESX.
type XEniInfo struct {
	// Field represent X-ENI associationId (On Shadow account).
	AssociationId *string
	// Field represent one of the eniId of ESX host (used as trunkInterfaceId) during X-ENI association.
	TrunkEniId *string
	// Field represent eniId which is created on customer account (X-ENI-ID).
	XEniId *string
}

func (s *XEniInfo) GetType__() vapiBindings_.BindingType {
	return XEniInfoBindingType()
}

func (s *XEniInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for XEniInfo._GetDataValue method - %s",
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
	fields["updated_by_user_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["updated_by_user_id"] = "UpdatedByUserId"
	fields["created"] = vapiBindings_.NewDateTimeType()
	fieldNameMap["created"] = "Created"
	fields["version"] = vapiBindings_.NewIntegerType()
	fieldNameMap["version"] = "Version"
	fields["updated_by_user_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["updated_by_user_name"] = "UpdatedByUserName"
	fields["user_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["user_name"] = "UserName"
	fields["id"] = vapiBindings_.NewStringType()
	fieldNameMap["id"] = "Id"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.abstract_entity", fields, reflect.TypeOf(AbstractEntity{}), fieldNameMap, validators)
}

func AccountLinkConfigBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["delay_account_link"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["delay_account_link"] = "DelayAccountLink"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.account_link_config", fields, reflect.TypeOf(AccountLinkConfig{}), fieldNameMap, validators)
}

func AccountLinkSddcConfigBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["customer_subnet_ids"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["customer_subnet_ids"] = "CustomerSubnetIds"
	fields["connected_account_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["connected_account_id"] = "ConnectedAccountId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.account_link_sddc_config", fields, reflect.TypeOf(AccountLinkSddcConfig{}), fieldNameMap, validators)
}

func AgentBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["sddc_csp_oauth_client"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(AgentCspOauthClientBindingType))
	fieldNameMap["sddc_csp_oauth_client"] = "SddcCspOauthClient"
	fields["addresses"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["addresses"] = "Addresses"
	fields["internal_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["internal_ip"] = "InternalIp"
	fields["hostname_verifier_enabled"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["hostname_verifier_enabled"] = "HostnameVerifierEnabled"
	fields["cert_enabled"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["cert_enabled"] = "CertEnabled"
	fields["network_cidr"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["network_cidr"] = "NetworkCidr"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["tenant_service_info"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(TenantServiceInfoBindingType))
	fieldNameMap["tenant_service_info"] = "TenantServiceInfo"
	fields["tinyproxy_whitelist"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["tinyproxy_whitelist"] = "TinyproxyWhitelist"
	fields["healthy"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["healthy"] = "Healthy"
	fields["custom_properties"] = vapiBindings_.NewOptionalType(vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewStringType(), reflect.TypeOf(map[string]string{})))
	fieldNameMap["custom_properties"] = "CustomProperties"
	fields["management_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["management_ip"] = "ManagementIp"
	fields["master"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["master"] = "Master"
	fields["network_netmask"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["network_netmask"] = "NetworkNetmask"
	fields["network_gateway"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["network_gateway"] = "NetworkGateway"
	fields["provider"] = vapiBindings_.NewStringType()
	fieldNameMap["provider"] = "Provider"
	fields["agent_url"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["agent_url"] = "AgentUrl"
	fields["agent_state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["agent_state"] = "AgentState"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.agent", fields, reflect.TypeOf(Agent{}), fieldNameMap, validators)
}

func AgentCspOauthClientBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["csp_oauth_client"] = vapiBindings_.NewReferenceType(AgentOauthClientBindingType)
	fieldNameMap["csp_oauth_client"] = "CspOauthClient"
	fields["create_at"] = vapiBindings_.NewDateTimeType()
	fieldNameMap["create_at"] = "CreateAt"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.agent_csp_oauth_client", fields, reflect.TypeOf(AgentCspOauthClient{}), fieldNameMap, validators)
}

func AgentOauthClientBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["client_secret"] = vapiBindings_.NewStringType()
	fieldNameMap["client_secret"] = "ClientSecret"
	fields["client_id"] = vapiBindings_.NewStringType()
	fieldNameMap["client_id"] = "ClientId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.agent_oauth_client", fields, reflect.TypeOf(AgentOauthClient{}), fieldNameMap, validators)
}

func AmiInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["instance_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["instance_type"] = "InstanceType"
	fields["region"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["region"] = "Region"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.ami_info", fields, reflect.TypeOf(AmiInfo{}), fieldNameMap, validators)
}

func AvailableZoneInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["subnets"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(SubnetBindingType), reflect.TypeOf([]Subnet{})))
	fieldNameMap["subnets"] = "Subnets"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.available_zone_info", fields, reflect.TypeOf(AvailableZoneInfo{}), fieldNameMap, validators)
}

func AwsAgentBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["instance_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["instance_id"] = "InstanceId"
	fields["availability_zone_info_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["availability_zone_info_id"] = "AvailabilityZoneInfoId"
	fields["instance_profile_info"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(InstanceProfileInfoBindingType))
	fieldNameMap["instance_profile_info"] = "InstanceProfileInfo"
	fields["key_pair"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(AwsKeyPairBindingType))
	fieldNameMap["key_pair"] = "KeyPair"
	fields["default_eni_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["default_eni_id"] = "DefaultEniId"
	fields["sddc_csp_oauth_client"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(AgentCspOauthClientBindingType))
	fieldNameMap["sddc_csp_oauth_client"] = "SddcCspOauthClient"
	fields["addresses"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["addresses"] = "Addresses"
	fields["internal_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["internal_ip"] = "InternalIp"
	fields["hostname_verifier_enabled"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["hostname_verifier_enabled"] = "HostnameVerifierEnabled"
	fields["cert_enabled"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["cert_enabled"] = "CertEnabled"
	fields["network_cidr"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["network_cidr"] = "NetworkCidr"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["tenant_service_info"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(TenantServiceInfoBindingType))
	fieldNameMap["tenant_service_info"] = "TenantServiceInfo"
	fields["tinyproxy_whitelist"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["tinyproxy_whitelist"] = "TinyproxyWhitelist"
	fields["healthy"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["healthy"] = "Healthy"
	fields["custom_properties"] = vapiBindings_.NewOptionalType(vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewStringType(), reflect.TypeOf(map[string]string{})))
	fieldNameMap["custom_properties"] = "CustomProperties"
	fields["management_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["management_ip"] = "ManagementIp"
	fields["master"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["master"] = "Master"
	fields["network_netmask"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["network_netmask"] = "NetworkNetmask"
	fields["network_gateway"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["network_gateway"] = "NetworkGateway"
	fields["provider"] = vapiBindings_.NewStringType()
	fieldNameMap["provider"] = "Provider"
	fields["agent_url"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["agent_url"] = "AgentUrl"
	fields["agent_state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["agent_state"] = "AgentState"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.aws_agent", fields, reflect.TypeOf(AwsAgent{}), fieldNameMap, validators)
}

func AwsCloudProviderBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["regions"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["regions"] = "Regions"
	fields["provider"] = vapiBindings_.NewStringType()
	fieldNameMap["provider"] = "Provider"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.aws_cloud_provider", fields, reflect.TypeOf(AwsCloudProvider{}), fieldNameMap, validators)
}

func AwsCompatibleSubnetsBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["customer_available_zones"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["customer_available_zones"] = "CustomerAvailableZones"
	fields["vpc_map"] = vapiBindings_.NewOptionalType(vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewReferenceType(VpcInfoSubnetsBindingType), reflect.TypeOf(map[string]VpcInfoSubnets{})))
	fieldNameMap["vpc_map"] = "VpcMap"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.aws_compatible_subnets", fields, reflect.TypeOf(AwsCompatibleSubnets{}), fieldNameMap, validators)
}

func AwsCustomerConnectedAccountBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["updated"] = vapiBindings_.NewDateTimeType()
	fieldNameMap["updated"] = "Updated"
	fields["user_id"] = vapiBindings_.NewStringType()
	fieldNameMap["user_id"] = "UserId"
	fields["updated_by_user_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["updated_by_user_id"] = "UpdatedByUserId"
	fields["created"] = vapiBindings_.NewDateTimeType()
	fieldNameMap["created"] = "Created"
	fields["version"] = vapiBindings_.NewIntegerType()
	fieldNameMap["version"] = "Version"
	fields["updated_by_user_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["updated_by_user_name"] = "UpdatedByUserName"
	fields["user_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["user_name"] = "UserName"
	fields["id"] = vapiBindings_.NewStringType()
	fieldNameMap["id"] = "Id"
	fields["policy_payer_arn"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["policy_payer_arn"] = "PolicyPayerArn"
	fields["region_to_az_to_shadow_mapping"] = vapiBindings_.NewOptionalType(vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewStringType(), reflect.TypeOf(map[string]string{})), reflect.TypeOf(map[string]map[string]string{})))
	fieldNameMap["region_to_az_to_shadow_mapping"] = "RegionToAzToShadowMapping"
	fields["org_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fields["cf_stack_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cf_stack_name"] = "CfStackName"
	fields["state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["state"] = "State"
	fields["account_number"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["account_number"] = "AccountNumber"
	fields["policy_service_arn"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["policy_service_arn"] = "PolicyServiceArn"
	fields["policy_external_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["policy_external_id"] = "PolicyExternalId"
	fields["policy_payer_linked_arn"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["policy_payer_linked_arn"] = "PolicyPayerLinkedArn"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.aws_customer_connected_account", fields, reflect.TypeOf(AwsCustomerConnectedAccount{}), fieldNameMap, validators)
}

func AwsEsxHostBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["capacity_pool"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["capacity_pool"] = "CapacityPool"
	fields["eni_list"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(EniInfoBindingType), reflect.TypeOf([]EniInfo{})))
	fieldNameMap["eni_list"] = "EniList"
	fields["xeni_info"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(XEniInfoBindingType))
	fieldNameMap["xeni_info"] = "XeniInfo"
	fields["partition_number"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["partition_number"] = "PartitionNumber"
	fields["ec2_instance_running_status_count"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["ec2_instance_running_status_count"] = "Ec2InstanceRunningStatusCount"
	fields["esx_nic_info"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(EsxNicInfoBindingType))
	fieldNameMap["esx_nic_info"] = "EsxNicInfo"
	fields["internal_public_ip_pool"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(SddcPublicIpBindingType), reflect.TypeOf([]SddcPublicIp{})))
	fieldNameMap["internal_public_ip_pool"] = "InternalPublicIpPool"
	fields["availability_zone"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["availability_zone"] = "AvailabilityZone"
	fields["esx_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["esx_id"] = "EsxId"
	fields["durable_host_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["durable_host_name"] = "DurableHostName"
	fields["state_last_updated"] = vapiBindings_.NewOptionalType(vapiBindings_.NewDateTimeType())
	fieldNameMap["state_last_updated"] = "StateLastUpdated"
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["custom_properties"] = vapiBindings_.NewOptionalType(vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewStringType(), reflect.TypeOf(map[string]string{})))
	fieldNameMap["custom_properties"] = "CustomProperties"
	fields["hostname"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["hostname"] = "Hostname"
	fields["provider"] = vapiBindings_.NewStringType()
	fieldNameMap["provider"] = "Provider"
	fields["instance_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["instance_type"] = "InstanceType"
	fields["mac_address"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["mac_address"] = "MacAddress"
	fields["esx_state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["esx_state"] = "EsxState"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.aws_esx_host", fields, reflect.TypeOf(AwsEsxHost{}), fieldNameMap, validators)
}

func AwsKeyPairBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["key_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["key_name"] = "KeyName"
	fields["key_fingerprint"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["key_fingerprint"] = "KeyFingerprint"
	fields["key_material"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["key_material"] = "KeyMaterial"
	fields["tags"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(TagBindingType), reflect.TypeOf([]Tag{})))
	fieldNameMap["tags"] = "Tags"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.aws_key_pair", fields, reflect.TypeOf(AwsKeyPair{}), fieldNameMap, validators)
}

func AwsKmsInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["amazon_resource_name"] = vapiBindings_.NewStringType()
	fieldNameMap["amazon_resource_name"] = "AmazonResourceName"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.aws_kms_info", fields, reflect.TypeOf(AwsKmsInfo{}), fieldNameMap, validators)
}

func AwsSddcConfigBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["existing_vpc_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["existing_vpc_id"] = "ExistingVpcId"
	fields["msft_license_config"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(MsftLicensingConfigBindingType))
	fieldNameMap["msft_license_config"] = "MsftLicenseConfig"
	fields["account_link_config"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(AccountLinkConfigBindingType))
	fieldNameMap["account_link_config"] = "AccountLinkConfig"
	fields["vpc_cidr"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vpc_cidr"] = "VpcCidr"
	fields["skip_creating_vxlan"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["skip_creating_vxlan"] = "SkipCreatingVxlan"
	fields["vxlan_subnet"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vxlan_subnet"] = "VxlanSubnet"
	fields["size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["size"] = "Size"
	fields["outpost_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["outpost_id"] = "OutpostId"
	fields["storage_capacity"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["storage_capacity"] = "StorageCapacity"
	fields["name"] = vapiBindings_.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["account_link_sddc_config"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(AccountLinkSddcConfigBindingType), reflect.TypeOf([]AccountLinkSddcConfig{})))
	fieldNameMap["account_link_sddc_config"] = "AccountLinkSddcConfig"
	fields["host_instance_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["host_instance_type"] = "HostInstanceType"
	fields["billing_account_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["billing_account_id"] = "BillingAccountId"
	fields["region"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["region"] = "Region"
	fields["sddc_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sddc_id"] = "SddcId"
	fields["sddc_template_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sddc_template_id"] = "SddcTemplateId"
	fields["sddc_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sddc_type"] = "SddcType"
	fields["vsan_version"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vsan_version"] = "VsanVersion"
	fields["provider"] = vapiBindings_.NewStringType()
	fieldNameMap["provider"] = "Provider"
	fields["sso_domain"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sso_domain"] = "SsoDomain"
	fields["num_hosts"] = vapiBindings_.NewIntegerType()
	fieldNameMap["num_hosts"] = "NumHosts"
	fields["deployment_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["deployment_type"] = "DeploymentType"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.aws_sddc_config", fields, reflect.TypeOf(AwsSddcConfig{}), fieldNameMap, validators)
}

func AwsSddcConnectionBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["updated"] = vapiBindings_.NewDateTimeType()
	fieldNameMap["updated"] = "Updated"
	fields["user_id"] = vapiBindings_.NewStringType()
	fieldNameMap["user_id"] = "UserId"
	fields["updated_by_user_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["updated_by_user_id"] = "UpdatedByUserId"
	fields["created"] = vapiBindings_.NewDateTimeType()
	fieldNameMap["created"] = "Created"
	fields["version"] = vapiBindings_.NewIntegerType()
	fieldNameMap["version"] = "Version"
	fields["updated_by_user_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["updated_by_user_name"] = "UpdatedByUserName"
	fields["user_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["user_name"] = "UserName"
	fields["id"] = vapiBindings_.NewStringType()
	fieldNameMap["id"] = "Id"
	fields["cidr_block_subnet"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cidr_block_subnet"] = "CidrBlockSubnet"
	fields["connected_account_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["connected_account_id"] = "ConnectedAccountId"
	fields["eni_group"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["eni_group"] = "EniGroup"
	fields["subnet_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["subnet_id"] = "SubnetId"
	fields["cgw_present"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["cgw_present"] = "CgwPresent"
	fields["org_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fields["sddc_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sddc_id"] = "SddcId"
	fields["cidr_block_vpc"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cidr_block_vpc"] = "CidrBlockVpc"
	fields["connection_order"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["connection_order"] = "ConnectionOrder"
	fields["state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["state"] = "State"
	fields["subnet_availability_zone"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["subnet_availability_zone"] = "SubnetAvailabilityZone"
	fields["vpc_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vpc_id"] = "VpcId"
	fields["customer_eni_infos"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(CustomerEniInfoBindingType), reflect.TypeOf([]CustomerEniInfo{})))
	fieldNameMap["customer_eni_infos"] = "CustomerEniInfos"
	fields["default_route_table"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["default_route_table"] = "DefaultRouteTable"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.aws_sddc_connection", fields, reflect.TypeOf(AwsSddcConnection{}), fieldNameMap, validators)
}

func AwsSddcResourceConfigBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["backup_restore_bucket"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["backup_restore_bucket"] = "BackupRestoreBucket"
	fields["public_ip_pool"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(SddcPublicIpBindingType), reflect.TypeOf([]SddcPublicIp{})))
	fieldNameMap["public_ip_pool"] = "PublicIpPool"
	fields["vpc_info"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(VpcInfoBindingType))
	fieldNameMap["vpc_info"] = "VpcInfo"
	fields["mgw_public_ip_pool"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(SddcPublicIpBindingType), reflect.TypeOf([]SddcPublicIp{})))
	fieldNameMap["mgw_public_ip_pool"] = "MgwPublicIpPool"
	fields["kms_vpc_endpoint"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(KmsVpcEndpointBindingType))
	fieldNameMap["kms_vpc_endpoint"] = "KmsVpcEndpoint"
	fields["max_num_public_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["max_num_public_ip"] = "MaxNumPublicIp"
	fields["esx_instance_profile"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(InstanceProfileInfoBindingType))
	fieldNameMap["esx_instance_profile"] = "EsxInstanceProfile"
	fields["account_link_sddc_config"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(SddcLinkConfigBindingType), reflect.TypeOf([]SddcLinkConfig{})))
	fieldNameMap["account_link_sddc_config"] = "AccountLinkSddcConfig"
	fields["vsan_encryption_config"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(VsanEncryptionConfigBindingType))
	fieldNameMap["vsan_encryption_config"] = "VsanEncryptionConfig"
	fields["nsx_user_client"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(CspOauthClientBindingType))
	fieldNameMap["nsx_user_client"] = "NsxUserClient"
	fields["vpc_info_peered_agent"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(VpcInfoBindingType))
	fieldNameMap["vpc_info_peered_agent"] = "VpcInfoPeeredAgent"
	fields["nsx_service_client"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(CspOauthClientBindingType))
	fieldNameMap["nsx_service_client"] = "NsxServiceClient"
	fields["vc_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vc_ip"] = "VcIp"
	fields["mgmt_appliance_network_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["mgmt_appliance_network_name"] = "MgmtApplianceNetworkName"
	fields["nsx_mgr_url"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["nsx_mgr_url"] = "NsxMgrUrl"
	fields["vlcm_enabled"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["vlcm_enabled"] = "VlcmEnabled"
	fields["nsx_cloud_audit_password"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["nsx_cloud_audit_password"] = "NsxCloudAuditPassword"
	fields["vc_csp_login_status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vc_csp_login_status"] = "VcCspLoginStatus"
	fields["nsx_cloud_admin_password"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["nsx_cloud_admin_password"] = "NsxCloudAdminPassword"
	fields["management_ds"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["management_ds"] = "ManagementDs"
	fields["nsx_api_public_endpoint_url"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["nsx_api_public_endpoint_url"] = "NsxApiPublicEndpointUrl"
	fields["nfs_mode"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["nfs_mode"] = "NfsMode"
	fields["cloud_password"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cloud_password"] = "CloudPassword"
	fields["sddc_networks"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["sddc_networks"] = "SddcNetworks"
	fields["clusters"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ClusterBindingType), reflect.TypeOf([]Cluster{})))
	fieldNameMap["clusters"] = "Clusters"
	fields["cloud_username"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cloud_username"] = "CloudUsername"
	fields["deployment_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["deployment_type"] = "DeploymentType"
	fields["pop_agent_xeni_connection"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(PopAgentXeniConnectionBindingType))
	fieldNameMap["pop_agent_xeni_connection"] = "PopAgentXeniConnection"
	fields["nsx_mgr_management_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["nsx_mgr_management_ip"] = "NsxMgrManagementIp"
	fields["nsx_cloud_audit"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["nsx_cloud_audit"] = "NsxCloudAudit"
	fields["esx_cluster_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["esx_cluster_id"] = "EsxClusterId"
	fields["mgw_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["mgw_id"] = "MgwId"
	fields["vc_url"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vc_url"] = "VcUrl"
	fields["esx_hosts"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(AwsEsxHostBindingType), reflect.TypeOf([]AwsEsxHost{})))
	fieldNameMap["esx_hosts"] = "EsxHosts"
	fields["cloud_user_group"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cloud_user_group"] = "CloudUserGroup"
	fields["management_rp"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["management_rp"] = "ManagementRp"
	fields["witness_availability_zone"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["witness_availability_zone"] = "WitnessAvailabilityZone"
	fields["sddc_desired_state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["sddc_desired_state"] = "SddcDesiredState"
	fields["sddc_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SddcSizeBindingType))
	fieldNameMap["sddc_size"] = "SddcSize"
	fields["cvds_enabled"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["cvds_enabled"] = "CvdsEnabled"
	fields["nsx_controller_ips"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["nsx_controller_ips"] = "NsxControllerIps"
	fields["two_hostname_vc_deployment"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["two_hostname_vc_deployment"] = "TwoHostnameVcDeployment"
	fields["esx_host_subnet"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["esx_host_subnet"] = "EsxHostSubnet"
	fields["sso_domain"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sso_domain"] = "SsoDomain"
	fields["region"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["region"] = "Region"
	fields["dns_with_management_vm_private_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["dns_with_management_vm_private_ip"] = "DnsWithManagementVmPrivateIp"
	fields["vc_public_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vc_public_ip"] = "VcPublicIp"
	fields["psc_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["psc_ip"] = "PscIp"
	fields["nsxt"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["nsxt"] = "Nsxt"
	fields["key_provider"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(KeyProviderDataBindingType), reflect.TypeOf([]KeyProviderData{})))
	fieldNameMap["key_provider"] = "KeyProvider"
	fields["psc_management_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["psc_management_ip"] = "PscManagementIp"
	fields["psc_url"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["psc_url"] = "PscUrl"
	fields["cgws"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["cgws"] = "Cgws"
	fields["availability_zones"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["availability_zones"] = "AvailabilityZones"
	fields["vc_containerized_permissions_enabled"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["vc_containerized_permissions_enabled"] = "VcContainerizedPermissionsEnabled"
	fields["custom_properties"] = vapiBindings_.NewOptionalType(vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewStringType(), reflect.TypeOf(map[string]string{})))
	fieldNameMap["custom_properties"] = "CustomProperties"
	fields["provider"] = vapiBindings_.NewStringType()
	fieldNameMap["provider"] = "Provider"
	fields["vc_management_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vc_management_ip"] = "VcManagementIp"
	fields["msft_license_config"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(MsftLicensingConfigBindingType))
	fieldNameMap["msft_license_config"] = "MsftLicenseConfig"
	fields["nsx_native_client"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(CspOauthClientBindingType))
	fieldNameMap["nsx_native_client"] = "NsxNativeClient"
	fields["vc_instance_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vc_instance_id"] = "VcInstanceId"
	fields["vc_oauth_client"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(CspOauthClientBindingType))
	fieldNameMap["vc_oauth_client"] = "VcOauthClient"
	fields["skip_creating_vxlan"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["skip_creating_vxlan"] = "SkipCreatingVxlan"
	fields["sddc_manifest"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SddcManifestBindingType))
	fieldNameMap["sddc_manifest"] = "SddcManifest"
	fields["vxlan_subnet"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vxlan_subnet"] = "VxlanSubnet"
	fields["sddc_security"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SddcSecurityBindingType))
	fieldNameMap["sddc_security"] = "SddcSecurity"
	fields["sddc_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sddc_id"] = "SddcId"
	fields["outpost_config"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(OutpostConfigBindingType))
	fieldNameMap["outpost_config"] = "OutpostConfig"
	fields["nsx_mgr_login_url"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["nsx_mgr_login_url"] = "NsxMgrLoginUrl"
	fields["vc_break_glass_url"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vc_break_glass_url"] = "VcBreakGlassUrl"
	fields["nsx_cloud_admin"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["nsx_cloud_admin"] = "NsxCloudAdmin"
	fields["nsxt_addons"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(NsxtAddonsBindingType))
	fieldNameMap["nsxt_addons"] = "NsxtAddons"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.aws_sddc_resource_config", fields, reflect.TypeOf(AwsSddcResourceConfig{}), fieldNameMap, validators)
}

func AwsSubnetBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["connected_account_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["connected_account_id"] = "ConnectedAccountId"
	fields["region_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["region_name"] = "RegionName"
	fields["availability_zone"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["availability_zone"] = "AvailabilityZone"
	fields["subnet_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["subnet_id"] = "SubnetId"
	fields["subnet_cidr_block"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["subnet_cidr_block"] = "SubnetCidrBlock"
	fields["is_compatible"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["is_compatible"] = "IsCompatible"
	fields["vpc_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vpc_id"] = "VpcId"
	fields["vpc_cidr_block"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vpc_cidr_block"] = "VpcCidrBlock"
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.aws_subnet", fields, reflect.TypeOf(AwsSubnet{}), fieldNameMap, validators)
}

func AwsWitnessEsxBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["instance_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["instance_id"] = "InstanceId"
	fields["eni_info"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(EniInfoBindingType))
	fieldNameMap["eni_info"] = "EniInfo"
	fields["enum_state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["enum_state"] = "EnumState"
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["mac_address"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["mac_address"] = "MacAddress"
	fields["esx_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["esx_id"] = "EsxId"
	fields["hostname"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["hostname"] = "Hostname"
	fields["provider"] = vapiBindings_.NewStringType()
	fieldNameMap["provider"] = "Provider"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.aws_witness_esx", fields, reflect.TypeOf(AwsWitnessEsx{}), fieldNameMap, validators)
}

func CloudProviderBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["provider"] = vapiBindings_.NewStringType()
	fieldNameMap["provider"] = "Provider"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.cloud_provider", fields, reflect.TypeOf(CloudProvider{}), fieldNameMap, validators)
}

func ClusterBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster_capacity"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(EntityCapacityBindingType))
	fieldNameMap["cluster_capacity"] = "ClusterCapacity"
	fields["esx_host_list"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(AwsEsxHostBindingType), reflect.TypeOf([]AwsEsxHost{})))
	fieldNameMap["esx_host_list"] = "EsxHostList"
	fields["wcp_details"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(WcpDetailsBindingType))
	fieldNameMap["wcp_details"] = "WcpDetails"
	fields["msft_license_config"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(MsftLicensingConfigBindingType))
	fieldNameMap["msft_license_config"] = "MsftLicenseConfig"
	fields["cluster_state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cluster_state"] = "ClusterState"
	fields["esx_host_info"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(EsxHostInfoBindingType))
	fieldNameMap["esx_host_info"] = "EsxHostInfo"
	fields["host_cpu_cores_count"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["host_cpu_cores_count"] = "HostCpuCoresCount"
	fields["partition_placement_group_info"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(PartitionPlacementGroupInfoBindingType), reflect.TypeOf([]PartitionPlacementGroupInfo{})))
	fieldNameMap["partition_placement_group_info"] = "PartitionPlacementGroupInfo"
	fields["cluster_id"] = vapiBindings_.NewStringType()
	fieldNameMap["cluster_id"] = "ClusterId"
	fields["availability_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["availability_type"] = "AvailabilityType"
	fields["mgmt_rp_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["mgmt_rp_name"] = "MgmtRpName"
	fields["aws_kms_info"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(AwsKmsInfoBindingType))
	fieldNameMap["aws_kms_info"] = "AwsKmsInfo"
	fields["availability_zones"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["availability_zones"] = "AvailabilityZones"
	fields["compute_rp_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["compute_rp_name"] = "ComputeRpName"
	fields["vsan_witness"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(AwsWitnessEsxBindingType))
	fieldNameMap["vsan_witness"] = "VsanWitness"
	fields["custom_properties"] = vapiBindings_.NewOptionalType(vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewStringType(), reflect.TypeOf(map[string]string{})))
	fieldNameMap["custom_properties"] = "CustomProperties"
	fields["cluster_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cluster_name"] = "ClusterName"
	fields["vsan_version"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vsan_version"] = "VsanVersion"
	fields["hyper_threading_enabled"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["hyper_threading_enabled"] = "HyperThreadingEnabled"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.cluster", fields, reflect.TypeOf(Cluster{}), fieldNameMap, validators)
}

func ClusterConfigBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["host_cpu_cores_count"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["host_cpu_cores_count"] = "HostCpuCoresCount"
	fields["host_instance_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["host_instance_type"] = "HostInstanceType"
	fields["storage_capacity"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["storage_capacity"] = "StorageCapacity"
	fields["availability_zones"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["availability_zones"] = "AvailabilityZones"
	fields["msft_license_config"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(MsftLicensingConfigBindingType))
	fieldNameMap["msft_license_config"] = "MsftLicenseConfig"
	fields["num_hosts"] = vapiBindings_.NewIntegerType()
	fieldNameMap["num_hosts"] = "NumHosts"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.cluster_config", fields, reflect.TypeOf(ClusterConfig{}), fieldNameMap, validators)
}

func ClusterReconfigureParamsBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["storage_capacity"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["storage_capacity"] = "StorageCapacity"
	fields["bias"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["bias"] = "Bias"
	fields["num_hosts"] = vapiBindings_.NewIntegerType()
	fieldNameMap["num_hosts"] = "NumHosts"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.cluster_reconfigure_params", fields, reflect.TypeOf(ClusterReconfigureParams{}), fieldNameMap, validators)
}

func ComputeGatewayTemplateBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["public_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SddcPublicIpBindingType))
	fieldNameMap["public_ip"] = "PublicIp"
	fields["primary_dns"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["primary_dns"] = "PrimaryDns"
	fields["secondary_dns"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["secondary_dns"] = "SecondaryDns"
	fields["firewall_rules"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(FirewallRuleBindingType), reflect.TypeOf([]FirewallRule{})))
	fieldNameMap["firewall_rules"] = "FirewallRules"
	fields["vpns"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(VpnBindingType), reflect.TypeOf([]Vpn{})))
	fieldNameMap["vpns"] = "Vpns"
	fields["logical_networks"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(LogicalNetworkBindingType), reflect.TypeOf([]LogicalNetwork{})))
	fieldNameMap["logical_networks"] = "LogicalNetworks"
	fields["nat_rules"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(NatRuleBindingType), reflect.TypeOf([]NatRule{})))
	fieldNameMap["nat_rules"] = "NatRules"
	fields["l2_vpn"] = vapiBindings_.NewOptionalType(vapiBindings_.NewDynamicStructType(nil))
	fieldNameMap["l2_vpn"] = "L2Vpn"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.compute_gateway_template", fields, reflect.TypeOf(ComputeGatewayTemplate{}), fieldNameMap, validators)
}

func ConfigSpecBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["powered_by_outpost_available"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["powered_by_outpost_available"] = "PoweredByOutpostAvailable"
	fields["expiry_in_days"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["expiry_in_days"] = "ExpiryInDays"
	fields["sddc_sizes"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["sddc_sizes"] = "SddcSizes"
	fields["availability"] = vapiBindings_.NewOptionalType(vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewListType(vapiBindings_.NewReferenceType(InstanceTypeConfigBindingType), reflect.TypeOf([]InstanceTypeConfig{})), reflect.TypeOf(map[string][]InstanceTypeConfig{})))
	fieldNameMap["availability"] = "Availability"
	fields["powered_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewListType(vapiBindings_.NewReferenceType(PoweredByInstanceTypeConfigBindingType), reflect.TypeOf([]PoweredByInstanceTypeConfig{})), reflect.TypeOf(map[string][]PoweredByInstanceTypeConfig{})))
	fieldNameMap["powered_by"] = "PoweredBy"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.config_spec", fields, reflect.TypeOf(ConfigSpec{}), fieldNameMap, validators)
}

func ConnectivityAgentValidationBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["source"] = "Source"
	fields["path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["type"] = "Type_"
	fields["ports"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["ports"] = "Ports"
	fields["pktsize"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["pktsize"] = "Pktsize"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.connectivity_agent_validation", fields, reflect.TypeOf(ConnectivityAgentValidation{}), fieldNameMap, validators)
}

func ConnectivityValidationGroupBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["sub_groups"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ConnectivityValidationSubGroupBindingType), reflect.TypeOf([]ConnectivityValidationSubGroup{})))
	fieldNameMap["sub_groups"] = "SubGroups"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.connectivity_validation_group", fields, reflect.TypeOf(ConnectivityValidationGroup{}), fieldNameMap, validators)
}

func ConnectivityValidationGroupsBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["groups"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ConnectivityValidationGroupBindingType), reflect.TypeOf([]ConnectivityValidationGroup{})))
	fieldNameMap["groups"] = "Groups"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.connectivity_validation_groups", fields, reflect.TypeOf(ConnectivityValidationGroups{}), fieldNameMap, validators)
}

func ConnectivityValidationInputBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["value"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["value"] = "Value"
	fields["label"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["label"] = "Label"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.connectivity_validation_input", fields, reflect.TypeOf(ConnectivityValidationInput{}), fieldNameMap, validators)
}

func ConnectivityValidationSubGroupBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["inputs"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ConnectivityValidationInputBindingType), reflect.TypeOf([]ConnectivityValidationInput{})))
	fieldNameMap["inputs"] = "Inputs"
	fields["tests"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ConnectivityAgentValidationBindingType), reflect.TypeOf([]ConnectivityAgentValidation{})))
	fieldNameMap["tests"] = "Tests"
	fields["label"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["label"] = "Label"
	fields["help"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["help"] = "Help"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.connectivity_validation_sub_group", fields, reflect.TypeOf(ConnectivityValidationSubGroup{}), fieldNameMap, validators)
}

func CspOauthClientBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["redirect_uris"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["redirect_uris"] = "RedirectUris"
	fields["accessTokenTTL"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["accessTokenTTL"] = "AccessTokenTTL"
	fields["redirect_uri"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["redirect_uri"] = "RedirectUri"
	fields["grant_types"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["grant_types"] = "GrantTypes"
	fields["org_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fields["secret"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["secret"] = "Secret"
	fields["refreshTokenTTL"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["refreshTokenTTL"] = "RefreshTokenTTL"
	fields["resource_link"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["resource_link"] = "ResourceLink"
	fields["id"] = vapiBindings_.NewStringType()
	fieldNameMap["id"] = "Id"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.csp_oauth_client", fields, reflect.TypeOf(CspOauthClient{}), fieldNameMap, validators)
}

func CustomerEniInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["secondary_ip_addresses"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["secondary_ip_addresses"] = "SecondaryIpAddresses"
	fields["eni_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["eni_id"] = "EniId"
	fields["primary_ip_address"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["primary_ip_address"] = "PrimaryIpAddress"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.customer_eni_info", fields, reflect.TypeOf(CustomerEniInfo{}), fieldNameMap, validators)
}

func EbsBackedVsanConfigBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["instance_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["instance_type"] = "InstanceType"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.ebs_backed_vsan_config", fields, reflect.TypeOf(EbsBackedVsanConfig{}), fieldNameMap, validators)
}

func EnablementInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["enabled"] = vapiBindings_.NewBooleanType()
	fieldNameMap["enabled"] = "Enabled"
	fields["name"] = vapiBindings_.NewStringType()
	fieldNameMap["name"] = "Name"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.enablement_info", fields, reflect.TypeOf(EnablementInfo{}), fieldNameMap, validators)
}

func EniInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["secondary_ips"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["secondary_ips"] = "SecondaryIps"
	fields["attachment_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["attachment_id"] = "AttachmentId"
	fields["association_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["association_id"] = "AssociationId"
	fields["vmk_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vmk_id"] = "VmkId"
	fields["device_index"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["device_index"] = "DeviceIndex"
	fields["security_group_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["security_group_id"] = "SecurityGroupId"
	fields["instance_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["instance_id"] = "InstanceId"
	fields["subnet_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["subnet_id"] = "SubnetId"
	fields["public_ips"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["public_ips"] = "PublicIps"
	fields["private_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["private_ip"] = "PrivateIp"
	fields["mac_address"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["mac_address"] = "MacAddress"
	fields["source_dest_check_false"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["source_dest_check_false"] = "SourceDestCheckFalse"
	fields["portgroup"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["portgroup"] = "Portgroup"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.eni_info", fields, reflect.TypeOf(EniInfo{}), fieldNameMap, validators)
}

func EntityCapacityBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["storage_capacity_gib"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["storage_capacity_gib"] = "StorageCapacityGib"
	fields["memory_capacity_gib"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["memory_capacity_gib"] = "MemoryCapacityGib"
	fields["total_number_of_cores"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["total_number_of_cores"] = "TotalNumberOfCores"
	fields["number_of_ssds"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["number_of_ssds"] = "NumberOfSsds"
	fields["cpu_capacity_ghz"] = vapiBindings_.NewOptionalType(vapiBindings_.NewDoubleType())
	fieldNameMap["cpu_capacity_ghz"] = "CpuCapacityGhz"
	fields["number_of_sockets"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["number_of_sockets"] = "NumberOfSockets"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.entity_capacity", fields, reflect.TypeOf(EntityCapacity{}), fieldNameMap, validators)
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
	return vapiBindings_.NewStructType("com.vmware.vmc.model.error_response", fields, reflect.TypeOf(ErrorResponse{}), fieldNameMap, validators)
}

func EsxConfigBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["availability_zone"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["availability_zone"] = "AvailabilityZone"
	fields["strict_placement"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["strict_placement"] = "StrictPlacement"
	fields["cluster_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cluster_id"] = "ClusterId"
	fields["esxs"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["esxs"] = "Esxs"
	fields["num_hosts"] = vapiBindings_.NewIntegerType()
	fieldNameMap["num_hosts"] = "NumHosts"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.esx_config", fields, reflect.TypeOf(EsxConfig{}), fieldNameMap, validators)
}

func EsxHostBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["availability_zone"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["availability_zone"] = "AvailabilityZone"
	fields["esx_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["esx_id"] = "EsxId"
	fields["durable_host_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["durable_host_name"] = "DurableHostName"
	fields["state_last_updated"] = vapiBindings_.NewOptionalType(vapiBindings_.NewDateTimeType())
	fieldNameMap["state_last_updated"] = "StateLastUpdated"
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["custom_properties"] = vapiBindings_.NewOptionalType(vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewStringType(), reflect.TypeOf(map[string]string{})))
	fieldNameMap["custom_properties"] = "CustomProperties"
	fields["hostname"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["hostname"] = "Hostname"
	fields["provider"] = vapiBindings_.NewStringType()
	fieldNameMap["provider"] = "Provider"
	fields["instance_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["instance_type"] = "InstanceType"
	fields["mac_address"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["mac_address"] = "MacAddress"
	fields["esx_state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["esx_state"] = "EsxState"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.esx_host", fields, reflect.TypeOf(EsxHost{}), fieldNameMap, validators)
}

func EsxHostInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["instance_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["instance_type"] = "InstanceType"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.esx_host_info", fields, reflect.TypeOf(EsxHostInfo{}), fieldNameMap, validators)
}

func EsxNicInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["ip"] = "Ip"
	fields["mac"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["mac"] = "Mac"
	fields["vlan"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vlan"] = "Vlan"
	fields["gateway"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["gateway"] = "Gateway"
	fields["net-stack"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["net-stack"] = "NetStack"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.esx_nic_info", fields, reflect.TypeOf(EsxNicInfo{}), fieldNameMap, validators)
}

func FirewallRuleBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["rule_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["rule_type"] = "RuleType"
	fields["application_ids"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["application_ids"] = "ApplicationIds"
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["rule_interface"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["rule_interface"] = "RuleInterface"
	fields["destination"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["destination"] = "Destination"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["destination_scope"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(FirewallRuleScopeBindingType))
	fieldNameMap["destination_scope"] = "DestinationScope"
	fields["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["source"] = "Source"
	fields["source_scope"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(FirewallRuleScopeBindingType))
	fieldNameMap["source_scope"] = "SourceScope"
	fields["services"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(FirewallServiceBindingType), reflect.TypeOf([]FirewallService{})))
	fieldNameMap["services"] = "Services"
	fields["action"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["action"] = "Action"
	fields["revision"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["revision"] = "Revision"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.firewall_rule", fields, reflect.TypeOf(FirewallRule{}), fieldNameMap, validators)
}

func FirewallRuleScopeBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["grouping_object_ids"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["grouping_object_ids"] = "GroupingObjectIds"
	fields["vnic_group_ids"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["vnic_group_ids"] = "VnicGroupIds"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.firewall_rule_scope", fields, reflect.TypeOf(FirewallRuleScope{}), fieldNameMap, validators)
}

func FirewallServiceBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["protocol"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["protocol"] = "Protocol"
	fields["ports"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["ports"] = "Ports"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.firewall_service", fields, reflect.TypeOf(FirewallService{}), fieldNameMap, validators)
}

func GatewayTemplateBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["public_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SddcPublicIpBindingType))
	fieldNameMap["public_ip"] = "PublicIp"
	fields["primary_dns"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["primary_dns"] = "PrimaryDns"
	fields["secondary_dns"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["secondary_dns"] = "SecondaryDns"
	fields["firewall_rules"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(FirewallRuleBindingType), reflect.TypeOf([]FirewallRule{})))
	fieldNameMap["firewall_rules"] = "FirewallRules"
	fields["vpns"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(VpnBindingType), reflect.TypeOf([]Vpn{})))
	fieldNameMap["vpns"] = "Vpns"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.gateway_template", fields, reflect.TypeOf(GatewayTemplate{}), fieldNameMap, validators)
}

func GlcmBundleBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["s3Bucket"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["s3Bucket"] = "S3Bucket"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.glcm_bundle", fields, reflect.TypeOf(GlcmBundle{}), fieldNameMap, validators)
}

func InstanceEntityCapacityBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["storage_capacity_gib"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["storage_capacity_gib"] = "StorageCapacityGib"
	fields["memory_capacity_gib"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["memory_capacity_gib"] = "MemoryCapacityGib"
	fields["total_number_of_cores"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["total_number_of_cores"] = "TotalNumberOfCores"
	fields["number_of_ssds"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["number_of_ssds"] = "NumberOfSsds"
	fields["cpu_capacity_ghz"] = vapiBindings_.NewOptionalType(vapiBindings_.NewDoubleType())
	fieldNameMap["cpu_capacity_ghz"] = "CpuCapacityGhz"
	fields["number_of_sockets"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["number_of_sockets"] = "NumberOfSockets"
	fields["esa_storage_capacity_gib"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["esa_storage_capacity_gib"] = "EsaStorageCapacityGib"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.instance_entity_capacity", fields, reflect.TypeOf(InstanceEntityCapacity{}), fieldNameMap, validators)
}

func InstanceProfileInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["role_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["role_name"] = "RoleName"
	fields["instance_profile_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["instance_profile_name"] = "InstanceProfileName"
	fields["policy_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["policy_name"] = "PolicyName"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.instance_profile_info", fields, reflect.TypeOf(InstanceProfileInfo{}), fieldNameMap, validators)
}

func InstanceTypeBasicConfigBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vsan_esa_supported"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["vsan_esa_supported"] = "VsanEsaSupported"
	fields["display_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["description"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["instance_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["instance_type"] = "InstanceType"
	fields["entity_capacity"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(InstanceEntityCapacityBindingType))
	fieldNameMap["entity_capacity"] = "EntityCapacity"
	fields["label"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["label"] = "Label"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.instance_type_basic_config", fields, reflect.TypeOf(InstanceTypeBasicConfig{}), fieldNameMap, validators)
}

func InstanceTypeConfigBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vsan_esa_supported"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["vsan_esa_supported"] = "VsanEsaSupported"
	fields["display_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["description"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["instance_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["instance_type"] = "InstanceType"
	fields["entity_capacity"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(InstanceEntityCapacityBindingType))
	fieldNameMap["entity_capacity"] = "EntityCapacity"
	fields["label"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["label"] = "Label"
	fields["instance_provisioning_error_cause"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["instance_provisioning_error_cause"] = "InstanceProvisioningErrorCause"
	fields["supported_cpu_cores"] = vapiBindings_.NewOptionalType(vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewListType(vapiBindings_.NewIntegerType(), reflect.TypeOf([]int64{})), reflect.TypeOf(map[string][]int64{})))
	fieldNameMap["supported_cpu_cores"] = "SupportedCpuCores"
	fields["hosts"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewIntegerType(), reflect.TypeOf([]int64{})))
	fieldNameMap["hosts"] = "Hosts"
	fields["cpu_cores"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewIntegerType(), reflect.TypeOf([]int64{})))
	fieldNameMap["cpu_cores"] = "CpuCores"
	fields["hyper_threading_supported"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["hyper_threading_supported"] = "HyperThreadingSupported"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.instance_type_config", fields, reflect.TypeOf(InstanceTypeConfig{}), fieldNameMap, validators)
}

func KeyProviderDataBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["provider_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["provider_type"] = "ProviderType"
	fields["native_spec_key_material"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(NativeSpecKeyMaterialBindingType))
	fieldNameMap["native_spec_key_material"] = "NativeSpecKeyMaterial"
	fields["is_default_key_provider"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["is_default_key_provider"] = "IsDefaultKeyProvider"
	fields["is_host_key_provider"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["is_host_key_provider"] = "IsHostKeyProvider"
	fields["provider"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["provider"] = "Provider"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.key_provider_data", fields, reflect.TypeOf(KeyProviderData{}), fieldNameMap, validators)
}

func KmsVpcEndpointBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vpc_endpoint_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vpc_endpoint_id"] = "VpcEndpointId"
	fields["network_interface_ids"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["network_interface_ids"] = "NetworkInterfaceIds"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.kms_vpc_endpoint", fields, reflect.TypeOf(KmsVpcEndpoint{}), fieldNameMap, validators)
}

func L2VpnBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["enabled"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["enabled"] = "Enabled"
	fields["sites"] = vapiBindings_.NewListType(vapiBindings_.NewReferenceType(SiteBindingType), reflect.TypeOf([]Site{}))
	fieldNameMap["sites"] = "Sites"
	fields["listener_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["listener_ip"] = "ListenerIp"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.l2_vpn", fields, reflect.TypeOf(L2Vpn{}), fieldNameMap, validators)
}

func LinkRequestBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tracking_task"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["tracking_task"] = "TrackingTask"
	fields["expiration_date"] = vapiBindings_.NewOptionalType(vapiBindings_.NewDateTimeType())
	fieldNameMap["expiration_date"] = "ExpirationDate"
	fields["template_url"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["template_url"] = "TemplateUrl"
	fields["template_execution_url"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["template_execution_url"] = "TemplateExecutionUrl"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.link_request", fields, reflect.TypeOf(LinkRequest{}), fieldNameMap, validators)
}

func LogicalNetworkBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["subnet_cidr"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["subnet_cidr"] = "SubnetCidr"
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["gatewayIp"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["gatewayIp"] = "GatewayIp"
	fields["dhcp_enabled"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["dhcp_enabled"] = "DhcpEnabled"
	fields["dhcp_ip_range"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["dhcp_ip_range"] = "DhcpIpRange"
	fields["tunnel_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["tunnel_id"] = "TunnelId"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["network_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["network_type"] = "NetworkType"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.logical_network", fields, reflect.TypeOf(LogicalNetwork{}), fieldNameMap, validators)
}

func MaintenanceWindowBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["day_of_week"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["day_of_week"] = "DayOfWeek"
	fields["hour_of_day"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["hour_of_day"] = "HourOfDay"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.maintenance_window", fields, reflect.TypeOf(MaintenanceWindow{}), fieldNameMap, validators)
}

func MaintenanceWindowEntryBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["in_maintenance_window"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["in_maintenance_window"] = "InMaintenanceWindow"
	fields["reservation_schedule"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(ReservationScheduleBindingType))
	fieldNameMap["reservation_schedule"] = "ReservationSchedule"
	fields["reservation_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["reservation_id"] = "ReservationId"
	fields["in_maintenance_mode"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["in_maintenance_mode"] = "InMaintenanceMode"
	fields["sddc_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sddc_id"] = "SddcId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.maintenance_window_entry", fields, reflect.TypeOf(MaintenanceWindowEntry{}), fieldNameMap, validators)
}

func MaintenanceWindowGetBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["day_of_week"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["day_of_week"] = "DayOfWeek"
	fields["hour_of_day"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["hour_of_day"] = "HourOfDay"
	fields["duration_min"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["duration_min"] = "DurationMin"
	fields["version"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["version"] = "Version"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.maintenance_window_get", fields, reflect.TypeOf(MaintenanceWindowGet{}), fieldNameMap, validators)
}

func ManagementGatewayTemplateBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["public_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SddcPublicIpBindingType))
	fieldNameMap["public_ip"] = "PublicIp"
	fields["primary_dns"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["primary_dns"] = "PrimaryDns"
	fields["secondary_dns"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["secondary_dns"] = "SecondaryDns"
	fields["firewall_rules"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(FirewallRuleBindingType), reflect.TypeOf([]FirewallRule{})))
	fieldNameMap["firewall_rules"] = "FirewallRules"
	fields["vpns"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(VpnBindingType), reflect.TypeOf([]Vpn{})))
	fieldNameMap["vpns"] = "Vpns"
	fields["subnet_cidr"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["subnet_cidr"] = "SubnetCidr"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.management_gateway_template", fields, reflect.TypeOf(ManagementGatewayTemplate{}), fieldNameMap, validators)
}

func MapZonesRequestBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["connected_account_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["connected_account_id"] = "ConnectedAccountId"
	fields["org_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fields["petronas_regions_to_map"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["petronas_regions_to_map"] = "PetronasRegionsToMap"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.map_zones_request", fields, reflect.TypeOf(MapZonesRequest{}), fieldNameMap, validators)
}

func MetadataBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["timestamp"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["timestamp"] = "Timestamp"
	fields["cycle_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cycle_id"] = "CycleId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.metadata", fields, reflect.TypeOf(Metadata{}), fieldNameMap, validators)
}

func MsftLicensingConfigBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["mssql_licensing"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["mssql_licensing"] = "MssqlLicensing"
	fields["academic_license"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["academic_license"] = "AcademicLicense"
	fields["windows_licensing"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["windows_licensing"] = "WindowsLicensing"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.msft_licensing_config", fields, reflect.TypeOf(MsftLicensingConfig{}), fieldNameMap, validators)
}

func NatRuleBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["rule_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["rule_type"] = "RuleType"
	fields["protocol"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["protocol"] = "Protocol"
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["internal_ports"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["internal_ports"] = "InternalPorts"
	fields["public_ports"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["public_ports"] = "PublicPorts"
	fields["public_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["public_ip"] = "PublicIp"
	fields["internal_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["internal_ip"] = "InternalIp"
	fields["action"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["action"] = "Action"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["revision"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["revision"] = "Revision"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.nat_rule", fields, reflect.TypeOf(NatRule{}), fieldNameMap, validators)
}

func NativeSpecKeyMaterialBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["key_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["key_id"] = "KeyId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.native_spec_key_material", fields, reflect.TypeOf(NativeSpecKeyMaterial{}), fieldNameMap, validators)
}

func NetworkTemplateBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["management_gateway_templates"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ManagementGatewayTemplateBindingType), reflect.TypeOf([]ManagementGatewayTemplate{})))
	fieldNameMap["management_gateway_templates"] = "ManagementGatewayTemplates"
	fields["compute_gateway_templates"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ComputeGatewayTemplateBindingType), reflect.TypeOf([]ComputeGatewayTemplate{})))
	fieldNameMap["compute_gateway_templates"] = "ComputeGatewayTemplates"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.network_template", fields, reflect.TypeOf(NetworkTemplate{}), fieldNameMap, validators)
}

func NewCredentialsBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["username"] = vapiBindings_.NewStringType()
	fieldNameMap["username"] = "Username"
	fields["password"] = vapiBindings_.NewStringType()
	fieldNameMap["password"] = "Password"
	fields["name"] = vapiBindings_.NewStringType()
	fieldNameMap["name"] = "Name"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.new_credentials", fields, reflect.TypeOf(NewCredentials{}), fieldNameMap, validators)
}

func NsxtAddonsBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["enable_nsx_advanced_addon"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["enable_nsx_advanced_addon"] = "EnableNsxAdvancedAddon"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.nsxt_addons", fields, reflect.TypeOf(NsxtAddons{}), fieldNameMap, validators)
}

func OfferInstancesHolderBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["on_demand"] = vapiBindings_.NewReferenceType(OnDemandOfferInstanceBindingType)
	fieldNameMap["on_demand"] = "OnDemand"
	fields["offers"] = vapiBindings_.NewListType(vapiBindings_.NewReferenceType(TermOfferInstanceBindingType), reflect.TypeOf([]TermOfferInstance{}))
	fieldNameMap["offers"] = "Offers"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.offer_instances_holder", fields, reflect.TypeOf(OfferInstancesHolder{}), fieldNameMap, validators)
}

func OnDemandOfferInstanceBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["product"] = vapiBindings_.NewStringType()
	fieldNameMap["product"] = "Product"
	fields["product_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["product_type"] = "ProductType"
	fields["name"] = vapiBindings_.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["currency"] = vapiBindings_.NewStringType()
	fieldNameMap["currency"] = "Currency"
	fields["region"] = vapiBindings_.NewStringType()
	fieldNameMap["region"] = "Region"
	fields["unit_price"] = vapiBindings_.NewStringType()
	fieldNameMap["unit_price"] = "UnitPrice"
	fields["monthly_cost"] = vapiBindings_.NewStringType()
	fieldNameMap["monthly_cost"] = "MonthlyCost"
	fields["version"] = vapiBindings_.NewStringType()
	fieldNameMap["version"] = "Version"
	fields["type"] = vapiBindings_.NewStringType()
	fieldNameMap["type"] = "Type_"
	fields["description"] = vapiBindings_.NewStringType()
	fieldNameMap["description"] = "Description"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.on_demand_offer_instance", fields, reflect.TypeOf(OnDemandOfferInstance{}), fieldNameMap, validators)
}

func OrgPropertiesBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["values"] = vapiBindings_.NewOptionalType(vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewStringType(), reflect.TypeOf(map[string]string{})))
	fieldNameMap["values"] = "Values"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.org_properties", fields, reflect.TypeOf(OrgProperties{}), fieldNameMap, validators)
}

func OrgSellerInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["seller_account_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["seller_account_id"] = "SellerAccountId"
	fields["seller"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["seller"] = "Seller"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.org_seller_info", fields, reflect.TypeOf(OrgSellerInfo{}), fieldNameMap, validators)
}

func OrganizationBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["updated"] = vapiBindings_.NewDateTimeType()
	fieldNameMap["updated"] = "Updated"
	fields["user_id"] = vapiBindings_.NewStringType()
	fieldNameMap["user_id"] = "UserId"
	fields["updated_by_user_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["updated_by_user_id"] = "UpdatedByUserId"
	fields["created"] = vapiBindings_.NewDateTimeType()
	fieldNameMap["created"] = "Created"
	fields["version"] = vapiBindings_.NewIntegerType()
	fieldNameMap["version"] = "Version"
	fields["updated_by_user_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["updated_by_user_name"] = "UpdatedByUserName"
	fields["user_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["user_name"] = "UserName"
	fields["id"] = vapiBindings_.NewStringType()
	fieldNameMap["id"] = "Id"
	fields["org_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_type"] = "OrgType"
	fields["display_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["org_seller_info"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(OrgSellerInfoBindingType))
	fieldNameMap["org_seller_info"] = "OrgSellerInfo"
	fields["project_state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["project_state"] = "ProjectState"
	fields["properties"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(OrgPropertiesBindingType))
	fieldNameMap["properties"] = "Properties"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.organization", fields, reflect.TypeOf(Organization{}), fieldNameMap, validators)
}

func OutpostConfigBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["outpost_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["outpost_id"] = "OutpostId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.outpost_config", fields, reflect.TypeOf(OutpostConfig{}), fieldNameMap, validators)
}

func PartitionPlacementGroupInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["partition_group_names"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["partition_group_names"] = "PartitionGroupNames"
	fields["availability_zone"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["availability_zone"] = "AvailabilityZone"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.partition_placement_group_info", fields, reflect.TypeOf(PartitionPlacementGroupInfo{}), fieldNameMap, validators)
}

func PaymentMethodInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["type"] = "Type_"
	fields["default_flag"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["default_flag"] = "DefaultFlag"
	fields["payment_method_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["payment_method_id"] = "PaymentMethodId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.payment_method_info", fields, reflect.TypeOf(PaymentMethodInfo{}), fieldNameMap, validators)
}

func PopAgentXeniConnectionBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["default_subnet_route"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["default_subnet_route"] = "DefaultSubnetRoute"
	fields["eni_info"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(EniInfoBindingType))
	fieldNameMap["eni_info"] = "EniInfo"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.pop_agent_xeni_connection", fields, reflect.TypeOf(PopAgentXeniConnection{}), fieldNameMap, validators)
}

func PopAmiInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["instance_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["instance_type"] = "InstanceType"
	fields["region"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["region"] = "Region"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["type"] = "Type_"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.pop_ami_info", fields, reflect.TypeOf(PopAmiInfo{}), fieldNameMap, validators)
}

func PopInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ami_infos"] = vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewReferenceType(PopAmiInfoBindingType), reflect.TypeOf(map[string]PopAmiInfo{}))
	fieldNameMap["ami_infos"] = "AmiInfos"
	fields["manifest_version"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["manifest_version"] = "ManifestVersion"
	fields["service_infos"] = vapiBindings_.NewOptionalType(vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewReferenceType(PopServiceInfoBindingType), reflect.TypeOf(map[string]PopServiceInfo{})))
	fieldNameMap["service_infos"] = "ServiceInfos"
	fields["created_at"] = vapiBindings_.NewOptionalType(vapiBindings_.NewDateTimeType())
	fieldNameMap["created_at"] = "CreatedAt"
	fields["os_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["os_type"] = "OsType"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.pop_info", fields, reflect.TypeOf(PopInfo{}), fieldNameMap, validators)
}

func PopServiceInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cln"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cln"] = "Cln"
	fields["version"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["version"] = "Version"
	fields["build"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["build"] = "Build"
	fields["service"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["service"] = "Service"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.pop_service_info", fields, reflect.TypeOf(PopServiceInfo{}), fieldNameMap, validators)
}

func PoweredByInstanceTypeConfigBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vsan_esa_supported"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["vsan_esa_supported"] = "VsanEsaSupported"
	fields["display_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["description"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["instance_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["instance_type"] = "InstanceType"
	fields["entity_capacity"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(InstanceEntityCapacityBindingType))
	fieldNameMap["entity_capacity"] = "EntityCapacity"
	fields["label"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["label"] = "Label"
	fields["instance_provisioning_error_cause"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["instance_provisioning_error_cause"] = "InstanceProvisioningErrorCause"
	fields["supported_cpu_cores"] = vapiBindings_.NewOptionalType(vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewListType(vapiBindings_.NewIntegerType(), reflect.TypeOf([]int64{})), reflect.TypeOf(map[string][]int64{})))
	fieldNameMap["supported_cpu_cores"] = "SupportedCpuCores"
	fields["hosts"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewIntegerType(), reflect.TypeOf([]int64{})))
	fieldNameMap["hosts"] = "Hosts"
	fields["cpu_cores"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewIntegerType(), reflect.TypeOf([]int64{})))
	fieldNameMap["cpu_cores"] = "CpuCores"
	fields["hyper_threading_supported"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["hyper_threading_supported"] = "HyperThreadingSupported"
	fields["powered_by_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["powered_by_id"] = "PoweredById"
	fields["powered_by_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["powered_by_type"] = "PoweredByType"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.powered_by_instance_type_config", fields, reflect.TypeOf(PoweredByInstanceTypeConfig{}), fieldNameMap, validators)
}

func ProvisionSpecBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["provider"] = vapiBindings_.NewOptionalType(vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewReferenceType(SddcConfigSpecBindingType), reflect.TypeOf(map[string]SddcConfigSpec{})))
	fieldNameMap["provider"] = "Provider"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.provision_spec", fields, reflect.TypeOf(ProvisionSpec{}), fieldNameMap, validators)
}

func RequestDetailBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["aws_quota_request_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["aws_quota_request_id"] = "AwsQuotaRequestId"
	fields["detail_status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["detail_status"] = "DetailStatus"
	fields["resolved_at"] = vapiBindings_.NewOptionalType(vapiBindings_.NewDateTimeType())
	fieldNameMap["resolved_at"] = "ResolvedAt"
	fields["desired_value"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["desired_value"] = "DesiredValue"
	fields["aws_support_case_status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["aws_support_case_status"] = "AwsSupportCaseStatus"
	fields["aws_support_case_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["aws_support_case_id"] = "AwsSupportCaseId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.request_detail", fields, reflect.TypeOf(RequestDetail{}), fieldNameMap, validators)
}

func ReservationBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["duration"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["duration"] = "Duration"
	fields["rid"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["rid"] = "Rid"
	fields["create_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["create_time"] = "CreateTime"
	fields["start_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewDateTimeType())
	fieldNameMap["start_time"] = "StartTime"
	fields["metadata"] = vapiBindings_.NewOptionalType(vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewStringType(), reflect.TypeOf(map[string]string{})))
	fieldNameMap["metadata"] = "Metadata"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.reservation", fields, reflect.TypeOf(Reservation{}), fieldNameMap, validators)
}

func ReservationInMwBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["rid"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["rid"] = "Rid"
	fields["week_of"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["week_of"] = "WeekOf"
	fields["create_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewDateTimeType())
	fieldNameMap["create_time"] = "CreateTime"
	fields["metadata"] = vapiBindings_.NewOptionalType(vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewStringType(), reflect.TypeOf(map[string]string{})))
	fieldNameMap["metadata"] = "Metadata"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.reservation_in_mw", fields, reflect.TypeOf(ReservationInMw{}), fieldNameMap, validators)
}

func ReservationScheduleBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["day_of_week"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["day_of_week"] = "DayOfWeek"
	fields["hour_of_day"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["hour_of_day"] = "HourOfDay"
	fields["duration_min"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["duration_min"] = "DurationMin"
	fields["version"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["version"] = "Version"
	fields["reservations"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ReservationBindingType), reflect.TypeOf([]Reservation{})))
	fieldNameMap["reservations"] = "Reservations"
	fields["reservations_mw"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ReservationInMwBindingType), reflect.TypeOf([]ReservationInMw{})))
	fieldNameMap["reservations_mw"] = "ReservationsMw"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.reservation_schedule", fields, reflect.TypeOf(ReservationSchedule{}), fieldNameMap, validators)
}

func ReservationWindowBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["duration_hours"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["duration_hours"] = "DurationHours"
	fields["estimated_duration_hours_minimum"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["estimated_duration_hours_minimum"] = "EstimatedDurationHoursMinimum"
	fields["reservation_state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["reservation_state"] = "ReservationState"
	fields["emergency"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["emergency"] = "Emergency"
	fields["maintenance_properties"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(ReservationWindowMaintenancePropertiesBindingType))
	fieldNameMap["maintenance_properties"] = "MaintenanceProperties"
	fields["manifest_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["manifest_id"] = "ManifestId"
	fields["start_hour"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["start_hour"] = "StartHour"
	fields["sddc_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sddc_id"] = "SddcId"
	fields["start_date"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["start_date"] = "StartDate"
	fields["estimated_duration_hours_maximum"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["estimated_duration_hours_maximum"] = "EstimatedDurationHoursMaximum"
	fields["reserve_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["reserve_id"] = "ReserveId"
	fields["metadata"] = vapiBindings_.NewOptionalType(vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewStringType(), reflect.TypeOf(map[string]string{})))
	fieldNameMap["metadata"] = "Metadata"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.reservation_window", fields, reflect.TypeOf(ReservationWindow{}), fieldNameMap, validators)
}

func ReservationWindowMaintenancePropertiesBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["status"] = "Status"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.reservation_window_maintenance_properties", fields, reflect.TypeOf(ReservationWindowMaintenanceProperties{}), fieldNameMap, validators)
}

func RouteTableInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.route_table_info", fields, reflect.TypeOf(RouteTableInfo{}), fieldNameMap, validators)
}

func SddcBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["updated"] = vapiBindings_.NewDateTimeType()
	fieldNameMap["updated"] = "Updated"
	fields["user_id"] = vapiBindings_.NewStringType()
	fieldNameMap["user_id"] = "UserId"
	fields["updated_by_user_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["updated_by_user_id"] = "UpdatedByUserId"
	fields["created"] = vapiBindings_.NewDateTimeType()
	fieldNameMap["created"] = "Created"
	fields["version"] = vapiBindings_.NewIntegerType()
	fieldNameMap["version"] = "Version"
	fields["updated_by_user_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["updated_by_user_name"] = "UpdatedByUserName"
	fields["user_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["user_name"] = "UserName"
	fields["id"] = vapiBindings_.NewStringType()
	fieldNameMap["id"] = "Id"
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["sddc_state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sddc_state"] = "SddcState"
	fields["expiration_date"] = vapiBindings_.NewOptionalType(vapiBindings_.NewDateTimeType())
	fieldNameMap["expiration_date"] = "ExpirationDate"
	fields["org_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fields["sddc_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sddc_type"] = "SddcType"
	fields["provider"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["provider"] = "Provider"
	fields["account_link_state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["account_link_state"] = "AccountLinkState"
	fields["sddc_access_state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sddc_access_state"] = "SddcAccessState"
	fields["resource_config"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(AwsSddcResourceConfigBindingType))
	fieldNameMap["resource_config"] = "ResourceConfig"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.sddc", fields, reflect.TypeOf(Sddc{}), fieldNameMap, validators)
}

func SddcChoiceBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["displayText"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["displayText"] = "DisplayText"
	fields["value"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["value"] = "Value"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.sddc_choice", fields, reflect.TypeOf(SddcChoice{}), fieldNameMap, validators)
}

func SddcConfigBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["msft_license_config"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(MsftLicensingConfigBindingType))
	fieldNameMap["msft_license_config"] = "MsftLicenseConfig"
	fields["account_link_config"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(AccountLinkConfigBindingType))
	fieldNameMap["account_link_config"] = "AccountLinkConfig"
	fields["vpc_cidr"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vpc_cidr"] = "VpcCidr"
	fields["skip_creating_vxlan"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["skip_creating_vxlan"] = "SkipCreatingVxlan"
	fields["vxlan_subnet"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vxlan_subnet"] = "VxlanSubnet"
	fields["size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["size"] = "Size"
	fields["outpost_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["outpost_id"] = "OutpostId"
	fields["storage_capacity"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["storage_capacity"] = "StorageCapacity"
	fields["name"] = vapiBindings_.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["account_link_sddc_config"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(AccountLinkSddcConfigBindingType), reflect.TypeOf([]AccountLinkSddcConfig{})))
	fieldNameMap["account_link_sddc_config"] = "AccountLinkSddcConfig"
	fields["host_instance_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["host_instance_type"] = "HostInstanceType"
	fields["billing_account_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["billing_account_id"] = "BillingAccountId"
	fields["region"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["region"] = "Region"
	fields["sddc_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sddc_id"] = "SddcId"
	fields["sddc_template_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sddc_template_id"] = "SddcTemplateId"
	fields["sddc_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sddc_type"] = "SddcType"
	fields["vsan_version"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vsan_version"] = "VsanVersion"
	fields["provider"] = vapiBindings_.NewStringType()
	fieldNameMap["provider"] = "Provider"
	fields["sso_domain"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sso_domain"] = "SsoDomain"
	fields["num_hosts"] = vapiBindings_.NewIntegerType()
	fieldNameMap["num_hosts"] = "NumHosts"
	fields["deployment_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["deployment_type"] = "DeploymentType"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.sddc_config", fields, reflect.TypeOf(SddcConfig{}), fieldNameMap, validators)
}

func SddcConfigSpecBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["sddc_type_config_spec"] = vapiBindings_.NewOptionalType(vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewReferenceType(ConfigSpecBindingType), reflect.TypeOf(map[string]ConfigSpec{})))
	fieldNameMap["sddc_type_config_spec"] = "SddcTypeConfigSpec"
	fields["region_display_names"] = vapiBindings_.NewOptionalType(vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewStringType(), reflect.TypeOf(map[string]string{})))
	fieldNameMap["region_display_names"] = "RegionDisplayNames"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.sddc_config_spec", fields, reflect.TypeOf(SddcConfigSpec{}), fieldNameMap, validators)
}

func SddcConvertRequestBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["num_hosts"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["num_hosts"] = "NumHosts"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.sddc_convert_request", fields, reflect.TypeOf(SddcConvertRequest{}), fieldNameMap, validators)
}

func SddcIdBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["sddc_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sddc_id"] = "SddcId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.sddc_id", fields, reflect.TypeOf(SddcId{}), fieldNameMap, validators)
}

func SddcInputBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["inputs"] = vapiBindings_.NewListType(vapiBindings_.NewReferenceType(SddcListBindingType), reflect.TypeOf([]SddcList{}))
	fieldNameMap["inputs"] = "Inputs"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.sddc_input", fields, reflect.TypeOf(SddcInput{}), fieldNameMap, validators)
}

func SddcLinkConfigBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["customer_subnet_ids"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["customer_subnet_ids"] = "CustomerSubnetIds"
	fields["connected_account_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["connected_account_id"] = "ConnectedAccountId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.sddc_link_config", fields, reflect.TypeOf(SddcLinkConfig{}), fieldNameMap, validators)
}

func SddcListBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["label"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["label"] = "Label"
	fields["required"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["required"] = "Required"
	fields["placeholder"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["placeholder"] = "Placeholder"
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["choice"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(SddcChoiceBindingType), reflect.TypeOf([]SddcChoice{})))
	fieldNameMap["choice"] = "Choice"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.sddc_list", fields, reflect.TypeOf(SddcList{}), fieldNameMap, validators)
}

func SddcManifestBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vmc_version"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vmc_version"] = "VmcVersion"
	fields["glcm_bundle"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(GlcmBundleBindingType))
	fieldNameMap["glcm_bundle"] = "GlcmBundle"
	fields["pop_info"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(PopInfoBindingType))
	fieldNameMap["pop_info"] = "PopInfo"
	fields["vmc_internal_version"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vmc_internal_version"] = "VmcInternalVersion"
	fields["ebs_backed_vsan_config"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(EbsBackedVsanConfigBindingType))
	fieldNameMap["ebs_backed_vsan_config"] = "EbsBackedVsanConfig"
	fields["vsan_witness_ami"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(AmiInfoBindingType))
	fieldNameMap["vsan_witness_ami"] = "VsanWitnessAmi"
	fields["esx_ami"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(AmiInfoBindingType))
	fieldNameMap["esx_ami"] = "EsxAmi"
	fields["esx_nsxt_ami"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(AmiInfoBindingType))
	fieldNameMap["esx_nsxt_ami"] = "EsxNsxtAmi"
	fields["metadata"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(MetadataBindingType))
	fieldNameMap["metadata"] = "Metadata"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.sddc_manifest", fields, reflect.TypeOf(SddcManifest{}), fieldNameMap, validators)
}

func SddcPatchRequestBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.sddc_patch_request", fields, reflect.TypeOf(SddcPatchRequest{}), fieldNameMap, validators)
}

func SddcPublicIpBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["public_ip"] = vapiBindings_.NewStringType()
	fieldNameMap["public_ip"] = "PublicIp"
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["allocation_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["allocation_id"] = "AllocationId"
	fields["dnat_rule_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["dnat_rule_id"] = "DnatRuleId"
	fields["associated_private_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["associated_private_ip"] = "AssociatedPrivateIp"
	fields["snat_rule_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["snat_rule_id"] = "SnatRuleId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.sddc_public_ip", fields, reflect.TypeOf(SddcPublicIp{}), fieldNameMap, validators)
}

func SddcResourceConfigBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vc_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vc_ip"] = "VcIp"
	fields["mgmt_appliance_network_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["mgmt_appliance_network_name"] = "MgmtApplianceNetworkName"
	fields["nsx_mgr_url"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["nsx_mgr_url"] = "NsxMgrUrl"
	fields["vlcm_enabled"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["vlcm_enabled"] = "VlcmEnabled"
	fields["nsx_cloud_audit_password"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["nsx_cloud_audit_password"] = "NsxCloudAuditPassword"
	fields["vc_csp_login_status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vc_csp_login_status"] = "VcCspLoginStatus"
	fields["nsx_cloud_admin_password"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["nsx_cloud_admin_password"] = "NsxCloudAdminPassword"
	fields["management_ds"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["management_ds"] = "ManagementDs"
	fields["nsx_api_public_endpoint_url"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["nsx_api_public_endpoint_url"] = "NsxApiPublicEndpointUrl"
	fields["nfs_mode"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["nfs_mode"] = "NfsMode"
	fields["cloud_password"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cloud_password"] = "CloudPassword"
	fields["sddc_networks"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["sddc_networks"] = "SddcNetworks"
	fields["clusters"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ClusterBindingType), reflect.TypeOf([]Cluster{})))
	fieldNameMap["clusters"] = "Clusters"
	fields["cloud_username"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cloud_username"] = "CloudUsername"
	fields["deployment_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["deployment_type"] = "DeploymentType"
	fields["pop_agent_xeni_connection"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(PopAgentXeniConnectionBindingType))
	fieldNameMap["pop_agent_xeni_connection"] = "PopAgentXeniConnection"
	fields["nsx_mgr_management_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["nsx_mgr_management_ip"] = "NsxMgrManagementIp"
	fields["nsx_cloud_audit"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["nsx_cloud_audit"] = "NsxCloudAudit"
	fields["esx_cluster_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["esx_cluster_id"] = "EsxClusterId"
	fields["mgw_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["mgw_id"] = "MgwId"
	fields["vc_url"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vc_url"] = "VcUrl"
	fields["esx_hosts"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(AwsEsxHostBindingType), reflect.TypeOf([]AwsEsxHost{})))
	fieldNameMap["esx_hosts"] = "EsxHosts"
	fields["cloud_user_group"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cloud_user_group"] = "CloudUserGroup"
	fields["management_rp"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["management_rp"] = "ManagementRp"
	fields["witness_availability_zone"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["witness_availability_zone"] = "WitnessAvailabilityZone"
	fields["sddc_desired_state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["sddc_desired_state"] = "SddcDesiredState"
	fields["sddc_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SddcSizeBindingType))
	fieldNameMap["sddc_size"] = "SddcSize"
	fields["cvds_enabled"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["cvds_enabled"] = "CvdsEnabled"
	fields["nsx_controller_ips"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["nsx_controller_ips"] = "NsxControllerIps"
	fields["two_hostname_vc_deployment"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["two_hostname_vc_deployment"] = "TwoHostnameVcDeployment"
	fields["esx_host_subnet"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["esx_host_subnet"] = "EsxHostSubnet"
	fields["sso_domain"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sso_domain"] = "SsoDomain"
	fields["region"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["region"] = "Region"
	fields["dns_with_management_vm_private_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["dns_with_management_vm_private_ip"] = "DnsWithManagementVmPrivateIp"
	fields["vc_public_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vc_public_ip"] = "VcPublicIp"
	fields["psc_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["psc_ip"] = "PscIp"
	fields["nsxt"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["nsxt"] = "Nsxt"
	fields["key_provider"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(KeyProviderDataBindingType), reflect.TypeOf([]KeyProviderData{})))
	fieldNameMap["key_provider"] = "KeyProvider"
	fields["psc_management_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["psc_management_ip"] = "PscManagementIp"
	fields["psc_url"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["psc_url"] = "PscUrl"
	fields["cgws"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["cgws"] = "Cgws"
	fields["availability_zones"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["availability_zones"] = "AvailabilityZones"
	fields["vc_containerized_permissions_enabled"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["vc_containerized_permissions_enabled"] = "VcContainerizedPermissionsEnabled"
	fields["custom_properties"] = vapiBindings_.NewOptionalType(vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewStringType(), reflect.TypeOf(map[string]string{})))
	fieldNameMap["custom_properties"] = "CustomProperties"
	fields["provider"] = vapiBindings_.NewStringType()
	fieldNameMap["provider"] = "Provider"
	fields["vc_management_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vc_management_ip"] = "VcManagementIp"
	fields["msft_license_config"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(MsftLicensingConfigBindingType))
	fieldNameMap["msft_license_config"] = "MsftLicenseConfig"
	fields["nsx_native_client"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(CspOauthClientBindingType))
	fieldNameMap["nsx_native_client"] = "NsxNativeClient"
	fields["vc_instance_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vc_instance_id"] = "VcInstanceId"
	fields["vc_oauth_client"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(CspOauthClientBindingType))
	fieldNameMap["vc_oauth_client"] = "VcOauthClient"
	fields["skip_creating_vxlan"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["skip_creating_vxlan"] = "SkipCreatingVxlan"
	fields["sddc_manifest"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SddcManifestBindingType))
	fieldNameMap["sddc_manifest"] = "SddcManifest"
	fields["vxlan_subnet"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vxlan_subnet"] = "VxlanSubnet"
	fields["sddc_security"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SddcSecurityBindingType))
	fieldNameMap["sddc_security"] = "SddcSecurity"
	fields["sddc_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sddc_id"] = "SddcId"
	fields["outpost_config"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(OutpostConfigBindingType))
	fieldNameMap["outpost_config"] = "OutpostConfig"
	fields["nsx_mgr_login_url"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["nsx_mgr_login_url"] = "NsxMgrLoginUrl"
	fields["vc_break_glass_url"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vc_break_glass_url"] = "VcBreakGlassUrl"
	fields["nsx_cloud_admin"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["nsx_cloud_admin"] = "NsxCloudAdmin"
	fields["nsxt_addons"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(NsxtAddonsBindingType))
	fieldNameMap["nsxt_addons"] = "NsxtAddons"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.sddc_resource_config", fields, reflect.TypeOf(SddcResourceConfig{}), fieldNameMap, validators)
}

func SddcSecurityBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["profile"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["profile"] = "Profile"
	fields["hardened"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["hardened"] = "Hardened"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.sddc_security", fields, reflect.TypeOf(SddcSecurity{}), fieldNameMap, validators)
}

func SddcSizeBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vc_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vc_size"] = "VcSize"
	fields["nsx_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["nsx_size"] = "NsxSize"
	fields["size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["size"] = "Size"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.sddc_size", fields, reflect.TypeOf(SddcSize{}), fieldNameMap, validators)
}

func SddcStateRequestBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["sddcs"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["sddcs"] = "Sddcs"
	fields["states"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["states"] = "States"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.sddc_state_request", fields, reflect.TypeOf(SddcStateRequest{}), fieldNameMap, validators)
}

func SddcTemplateBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["updated"] = vapiBindings_.NewDateTimeType()
	fieldNameMap["updated"] = "Updated"
	fields["user_id"] = vapiBindings_.NewStringType()
	fieldNameMap["user_id"] = "UserId"
	fields["updated_by_user_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["updated_by_user_id"] = "UpdatedByUserId"
	fields["created"] = vapiBindings_.NewDateTimeType()
	fieldNameMap["created"] = "Created"
	fields["version"] = vapiBindings_.NewIntegerType()
	fieldNameMap["version"] = "Version"
	fields["updated_by_user_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["updated_by_user_name"] = "UpdatedByUserName"
	fields["user_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["user_name"] = "UserName"
	fields["id"] = vapiBindings_.NewStringType()
	fieldNameMap["id"] = "Id"
	fields["account_link_sddc_configs"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(AccountLinkSddcConfigBindingType), reflect.TypeOf([]AccountLinkSddcConfig{})))
	fieldNameMap["account_link_sddc_configs"] = "AccountLinkSddcConfigs"
	fields["state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["state"] = "State"
	fields["network_template"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(NetworkTemplateBindingType))
	fieldNameMap["network_template"] = "NetworkTemplate"
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["source_sddc_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["source_sddc_id"] = "SourceSddcId"
	fields["org_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fields["sddc"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SddcBindingType))
	fieldNameMap["sddc"] = "Sddc"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.sddc_template", fields, reflect.TypeOf(SddcTemplate{}), fieldNameMap, validators)
}

func ServiceErrorBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["default_message"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["default_message"] = "DefaultMessage"
	fields["original_service"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["original_service"] = "OriginalService"
	fields["localized_message"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["localized_message"] = "LocalizedMessage"
	fields["original_service_error_code"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["original_service_error_code"] = "OriginalServiceErrorCode"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.service_error", fields, reflect.TypeOf(ServiceError{}), fieldNameMap, validators)
}

func ServiceQuotaRequestBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["updated"] = vapiBindings_.NewDateTimeType()
	fieldNameMap["updated"] = "Updated"
	fields["user_id"] = vapiBindings_.NewStringType()
	fieldNameMap["user_id"] = "UserId"
	fields["updated_by_user_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["updated_by_user_id"] = "UpdatedByUserId"
	fields["created"] = vapiBindings_.NewDateTimeType()
	fieldNameMap["created"] = "Created"
	fields["version"] = vapiBindings_.NewIntegerType()
	fieldNameMap["version"] = "Version"
	fields["updated_by_user_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["updated_by_user_name"] = "UpdatedByUserName"
	fields["user_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["user_name"] = "UserName"
	fields["id"] = vapiBindings_.NewStringType()
	fieldNameMap["id"] = "Id"
	fields["requester_email"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["requester_email"] = "RequesterEmail"
	fields["task_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["task_id"] = "TaskId"
	fields["region"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["region"] = "Region"
	fields["aws_account_number"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["aws_account_number"] = "AwsAccountNumber"
	fields["org_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fields["reason"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["reason"] = "Reason"
	fields["request_status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["request_status"] = "RequestStatus"
	fields["request_details"] = vapiBindings_.NewOptionalType(vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewReferenceType(RequestDetailBindingType), reflect.TypeOf(map[string]RequestDetail{})))
	fieldNameMap["request_details"] = "RequestDetails"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.service_quota_request", fields, reflect.TypeOf(ServiceQuotaRequest{}), fieldNameMap, validators)
}

func SiteBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["password"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["password"] = "Password"
	fields["user_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["user_id"] = "UserId"
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["rx_bytes_on_local_subnet"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["rx_bytes_on_local_subnet"] = "RxBytesOnLocalSubnet"
	fields["secure_traffic"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["secure_traffic"] = "SecureTraffic"
	fields["established_date"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["established_date"] = "EstablishedDate"
	fields["failure_message"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["failure_message"] = "FailureMessage"
	fields["dropped_tx_packets"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["dropped_tx_packets"] = "DroppedTxPackets"
	fields["dropped_rx_packets"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["dropped_rx_packets"] = "DroppedRxPackets"
	fields["tunnel_status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["tunnel_status"] = "TunnelStatus"
	fields["tx_bytes_from_local_subnet"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["tx_bytes_from_local_subnet"] = "TxBytesFromLocalSubnet"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.site", fields, reflect.TypeOf(Site{}), fieldNameMap, validators)
}

func SubnetBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["subnet_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["subnet_id"] = "SubnetId"
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["route_tables"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(SubnetRouteTableInfoBindingType), reflect.TypeOf([]SubnetRouteTableInfo{})))
	fieldNameMap["route_tables"] = "RouteTables"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.subnet", fields, reflect.TypeOf(Subnet{}), fieldNameMap, validators)
}

func SubnetInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["compatible"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["compatible"] = "Compatible"
	fields["connected_account_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["connected_account_id"] = "ConnectedAccountId"
	fields["region_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["region_name"] = "RegionName"
	fields["availability_zone"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["availability_zone"] = "AvailabilityZone"
	fields["subnet_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["subnet_id"] = "SubnetId"
	fields["availability_zone_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["availability_zone_id"] = "AvailabilityZoneId"
	fields["subnet_cidr_block"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["subnet_cidr_block"] = "SubnetCidrBlock"
	fields["note"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["note"] = "Note"
	fields["vpc_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vpc_id"] = "VpcId"
	fields["vpc_cidr_block"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vpc_cidr_block"] = "VpcCidrBlock"
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.subnet_info", fields, reflect.TypeOf(SubnetInfo{}), fieldNameMap, validators)
}

func SubnetRouteTableInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["subnet_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["subnet_id"] = "SubnetId"
	fields["association_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["association_id"] = "AssociationId"
	fields["routetable_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["routetable_id"] = "RoutetableId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.subnet_route_table_info", fields, reflect.TypeOf(SubnetRouteTableInfo{}), fieldNameMap, validators)
}

func SubscriptionDetailsBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["status"] = "Status"
	fields["anniversary_billing_date"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["anniversary_billing_date"] = "AnniversaryBillingDate"
	fields["end_date"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["end_date"] = "EndDate"
	fields["billing_frequency"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["billing_frequency"] = "BillingFrequency"
	fields["auto_renewed_allowed"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["auto_renewed_allowed"] = "AutoRenewedAllowed"
	fields["commitment_term"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["commitment_term"] = "CommitmentTerm"
	fields["csp_subscription_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["csp_subscription_id"] = "CspSubscriptionId"
	fields["billing_subscription_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["billing_subscription_id"] = "BillingSubscriptionId"
	fields["offer_version"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["offer_version"] = "OfferVersion"
	fields["offer_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["offer_type"] = "OfferType"
	fields["description"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["product_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["product_id"] = "ProductId"
	fields["region"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["region"] = "Region"
	fields["product_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["product_name"] = "ProductName"
	fields["offer_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["offer_name"] = "OfferName"
	fields["commitment_term_uom"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["commitment_term_uom"] = "CommitmentTermUom"
	fields["start_date"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["start_date"] = "StartDate"
	fields["quantity"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["quantity"] = "Quantity"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.subscription_details", fields, reflect.TypeOf(SubscriptionDetails{}), fieldNameMap, validators)
}

func SubscriptionProductsBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["product"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["product"] = "Product"
	fields["types"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["types"] = "Types"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.subscription_products", fields, reflect.TypeOf(SubscriptionProducts{}), fieldNameMap, validators)
}

func SubscriptionRequestBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["product"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["product"] = "Product"
	fields["product_type"] = vapiBindings_.NewStringType()
	fieldNameMap["product_type"] = "ProductType"
	fields["product_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["product_id"] = "ProductId"
	fields["billing_frequency"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["billing_frequency"] = "BillingFrequency"
	fields["region"] = vapiBindings_.NewStringType()
	fieldNameMap["region"] = "Region"
	fields["commitment_term"] = vapiBindings_.NewStringType()
	fieldNameMap["commitment_term"] = "CommitmentTerm"
	fields["offer_context_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["offer_context_id"] = "OfferContextId"
	fields["offer_version"] = vapiBindings_.NewStringType()
	fieldNameMap["offer_version"] = "OfferVersion"
	fields["offer_name"] = vapiBindings_.NewStringType()
	fieldNameMap["offer_name"] = "OfferName"
	fields["quantity"] = vapiBindings_.NewIntegerType()
	fieldNameMap["quantity"] = "Quantity"
	fields["type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["type"] = "Type_"
	fields["price"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["price"] = "Price"
	fields["product_charge_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["product_charge_id"] = "ProductChargeId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.subscription_request", fields, reflect.TypeOf(SubscriptionRequest{}), fieldNameMap, validators)
}

func SupportWindowBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["start_day"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["start_day"] = "StartDay"
	fields["seats"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["seats"] = "Seats"
	fields["sddcs"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["sddcs"] = "Sddcs"
	fields["duration_hours"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["duration_hours"] = "DurationHours"
	fields["start_hour"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["start_hour"] = "StartHour"
	fields["support_window_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["support_window_id"] = "SupportWindowId"
	fields["metadata"] = vapiBindings_.NewOptionalType(vapiBindings_.NewDynamicStructType(nil))
	fieldNameMap["metadata"] = "Metadata"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.support_window", fields, reflect.TypeOf(SupportWindow{}), fieldNameMap, validators)
}

func SupportWindowIdBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["window_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["window_id"] = "WindowId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.support_window_id", fields, reflect.TypeOf(SupportWindowId{}), fieldNameMap, validators)
}

func TagBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["key"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["key"] = "Key"
	fields["value"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["value"] = "Value"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.tag", fields, reflect.TypeOf(Tag{}), fieldNameMap, validators)
}

func TaskBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["updated"] = vapiBindings_.NewDateTimeType()
	fieldNameMap["updated"] = "Updated"
	fields["user_id"] = vapiBindings_.NewStringType()
	fieldNameMap["user_id"] = "UserId"
	fields["updated_by_user_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["updated_by_user_id"] = "UpdatedByUserId"
	fields["created"] = vapiBindings_.NewDateTimeType()
	fieldNameMap["created"] = "Created"
	fields["version"] = vapiBindings_.NewIntegerType()
	fieldNameMap["version"] = "Version"
	fields["updated_by_user_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["updated_by_user_name"] = "UpdatedByUserName"
	fields["user_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["user_name"] = "UserName"
	fields["id"] = vapiBindings_.NewStringType()
	fieldNameMap["id"] = "Id"
	fields["status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["status"] = "Status"
	fields["localized_error_message"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["localized_error_message"] = "LocalizedErrorMessage"
	fields["resource_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["resource_id"] = "ResourceId"
	fields["parent_task_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["parent_task_id"] = "ParentTaskId"
	fields["task_version"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["task_version"] = "TaskVersion"
	fields["correlation_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["correlation_id"] = "CorrelationId"
	fields["start_resource_entity_version"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["start_resource_entity_version"] = "StartResourceEntityVersion"
	fields["customer_error_message"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["customer_error_message"] = "CustomerErrorMessage"
	fields["sub_status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sub_status"] = "SubStatus"
	fields["task_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["task_type"] = "TaskType"
	fields["start_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewDateTimeType())
	fieldNameMap["start_time"] = "StartTime"
	fields["task_progress_phases"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(TaskProgressPhaseBindingType), reflect.TypeOf([]TaskProgressPhase{})))
	fieldNameMap["task_progress_phases"] = "TaskProgressPhases"
	fields["error_message"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["error_message"] = "ErrorMessage"
	fields["org_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fields["end_resource_entity_version"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["end_resource_entity_version"] = "EndResourceEntityVersion"
	fields["service_errors"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ServiceErrorBindingType), reflect.TypeOf([]ServiceError{})))
	fieldNameMap["service_errors"] = "ServiceErrors"
	fields["org_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_type"] = "OrgType"
	fields["estimated_remaining_minutes"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["estimated_remaining_minutes"] = "EstimatedRemainingMinutes"
	fields["params"] = vapiBindings_.NewOptionalType(vapiBindings_.NewDynamicStructType(nil))
	fieldNameMap["params"] = "Params"
	fields["progress_percent"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["progress_percent"] = "ProgressPercent"
	fields["phase_in_progress"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["phase_in_progress"] = "PhaseInProgress"
	fields["resource_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fields["end_time"] = vapiBindings_.NewOptionalType(vapiBindings_.NewDateTimeType())
	fieldNameMap["end_time"] = "EndTime"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.task", fields, reflect.TypeOf(Task{}), fieldNameMap, validators)
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
	return vapiBindings_.NewStructType("com.vmware.vmc.model.task_progress_phase", fields, reflect.TypeOf(TaskProgressPhase{}), fieldNameMap, validators)
}

func TenantServiceInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["s3_log_bucket_arn"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["s3_log_bucket_arn"] = "S3LogBucketArn"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.tenant_service_info", fields, reflect.TypeOf(TenantServiceInfo{}), fieldNameMap, validators)
}

func TermBillingOptionsBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["unit_price"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["unit_price"] = "UnitPrice"
	fields["billing_frequency"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["billing_frequency"] = "BillingFrequency"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.term_billing_options", fields, reflect.TypeOf(TermBillingOptions{}), fieldNameMap, validators)
}

func TermOfferInstanceBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["description"] = vapiBindings_.NewStringType()
	fieldNameMap["description"] = "Description"
	fields["product"] = vapiBindings_.NewStringType()
	fieldNameMap["product"] = "Product"
	fields["product_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["product_type"] = "ProductType"
	fields["name"] = vapiBindings_.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["currency"] = vapiBindings_.NewStringType()
	fieldNameMap["currency"] = "Currency"
	fields["region"] = vapiBindings_.NewStringType()
	fieldNameMap["region"] = "Region"
	fields["commitment_term"] = vapiBindings_.NewIntegerType()
	fieldNameMap["commitment_term"] = "CommitmentTerm"
	fields["unit_price"] = vapiBindings_.NewStringType()
	fieldNameMap["unit_price"] = "UnitPrice"
	fields["billing_options"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(TermBillingOptionsBindingType), reflect.TypeOf([]TermBillingOptions{})))
	fieldNameMap["billing_options"] = "BillingOptions"
	fields["version"] = vapiBindings_.NewStringType()
	fieldNameMap["version"] = "Version"
	fields["offer_context_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["offer_context_id"] = "OfferContextId"
	fields["product_charge_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["product_charge_id"] = "ProductChargeId"
	fields["type"] = vapiBindings_.NewStringType()
	fieldNameMap["type"] = "Type_"
	fields["product_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["product_id"] = "ProductId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.term_offer_instance", fields, reflect.TypeOf(TermOfferInstance{}), fieldNameMap, validators)
}

func TermsOfServiceResultBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["terms_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["terms_id"] = "TermsId"
	fields["signed"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["signed"] = "Signed"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.terms_of_service_result", fields, reflect.TypeOf(TermsOfServiceResult{}), fieldNameMap, validators)
}

func UpdateCredentialsBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["username"] = vapiBindings_.NewStringType()
	fieldNameMap["username"] = "Username"
	fields["password"] = vapiBindings_.NewStringType()
	fieldNameMap["password"] = "Password"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.update_credentials", fields, reflect.TypeOf(UpdateCredentials{}), fieldNameMap, validators)
}

func VmcLocaleBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["locale"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["locale"] = "Locale"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.vmc_locale", fields, reflect.TypeOf(VmcLocale{}), fieldNameMap, validators)
}

func VpcInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["esx_security_group_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["esx_security_group_id"] = "EsxSecurityGroupId"
	fields["vpc_cidr"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vpc_cidr"] = "VpcCidr"
	fields["vgw_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vgw_id"] = "VgwId"
	fields["esx_public_security_group_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["esx_public_security_group_id"] = "EsxPublicSecurityGroupId"
	fields["vif_ids"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["vif_ids"] = "VifIds"
	fields["vm_security_group_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vm_security_group_id"] = "VmSecurityGroupId"
	fields["route_table_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["route_table_id"] = "RouteTableId"
	fields["edge_subnet_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["edge_subnet_id"] = "EdgeSubnetId"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["api_association_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["api_association_id"] = "ApiAssociationId"
	fields["api_subnet_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["api_subnet_id"] = "ApiSubnetId"
	fields["private_subnet_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["private_subnet_id"] = "PrivateSubnetId"
	fields["private_association_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["private_association_id"] = "PrivateAssociationId"
	fields["vcdr_enis"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(EniInfoBindingType), reflect.TypeOf([]EniInfo{})))
	fieldNameMap["vcdr_enis"] = "VcdrEnis"
	fields["subnet_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["subnet_id"] = "SubnetId"
	fields["internet_gateway_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["internet_gateway_id"] = "InternetGatewayId"
	fields["security_group_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["security_group_id"] = "SecurityGroupId"
	fields["association_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["association_id"] = "AssociationId"
	fields["vgw_route_table_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vgw_route_table_id"] = "VgwRouteTableId"
	fields["traffic_group_edge_vm_ips"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["traffic_group_edge_vm_ips"] = "TrafficGroupEdgeVmIps"
	fields["edge_association_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["edge_association_id"] = "EdgeAssociationId"
	fields["provider"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["provider"] = "Provider"
	fields["tgw_ips"] = vapiBindings_.NewOptionalType(vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})), reflect.TypeOf(map[string][]string{})))
	fieldNameMap["tgw_ips"] = "TgwIps"
	fields["peering_connection_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["peering_connection_id"] = "PeeringConnectionId"
	fields["network_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["network_type"] = "NetworkType"
	fields["available_zones"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(AvailableZoneInfoBindingType), reflect.TypeOf([]AvailableZoneInfo{})))
	fieldNameMap["available_zones"] = "AvailableZones"
	fields["routetables"] = vapiBindings_.NewOptionalType(vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewReferenceType(RouteTableInfoBindingType), reflect.TypeOf(map[string]RouteTableInfo{})))
	fieldNameMap["routetables"] = "Routetables"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.vpc_info", fields, reflect.TypeOf(VpcInfo{}), fieldNameMap, validators)
}

func VpcInfoSubnetsBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vpc_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vpc_id"] = "VpcId"
	fields["cidr_block"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cidr_block"] = "CidrBlock"
	fields["description"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["subnets"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(SubnetInfoBindingType), reflect.TypeOf([]SubnetInfo{})))
	fieldNameMap["subnets"] = "Subnets"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.vpc_info_subnets", fields, reflect.TypeOf(VpcInfoSubnets{}), fieldNameMap, validators)
}

func VpnBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["version"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["version"] = "Version"
	fields["on_prem_gateway_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["on_prem_gateway_ip"] = "OnPremGatewayIp"
	fields["on_prem_network_cidr"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["on_prem_network_cidr"] = "OnPremNetworkCidr"
	fields["pfs_enabled"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["pfs_enabled"] = "PfsEnabled"
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["channel_status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(VpnChannelStatusBindingType))
	fieldNameMap["channel_status"] = "ChannelStatus"
	fields["on_prem_nat_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["on_prem_nat_ip"] = "OnPremNatIp"
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["internal_network_ids"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["internal_network_ids"] = "InternalNetworkIds"
	fields["tunnel_statuses"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(VpnTunnelStatusBindingType), reflect.TypeOf([]VpnTunnelStatus{})))
	fieldNameMap["tunnel_statuses"] = "TunnelStatuses"
	fields["encryption"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["encryption"] = "Encryption"
	fields["enabled"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["enabled"] = "Enabled"
	fields["state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["state"] = "State"
	fields["dh_group"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["dh_group"] = "DhGroup"
	fields["authentication"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["authentication"] = "Authentication"
	fields["pre_shared_key"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["pre_shared_key"] = "PreSharedKey"
	fields["ike_option"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["ike_option"] = "IkeOption"
	fields["digest_algorithm"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["digest_algorithm"] = "DigestAlgorithm"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.vpn", fields, reflect.TypeOf(Vpn{}), fieldNameMap, validators)
}

func VpnChannelStatusBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["channel_status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["channel_status"] = "ChannelStatus"
	fields["channel_state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["channel_state"] = "ChannelState"
	fields["last_info_message"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["last_info_message"] = "LastInfoMessage"
	fields["failure_message"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["failure_message"] = "FailureMessage"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.vpn_channel_status", fields, reflect.TypeOf(VpnChannelStatus{}), fieldNameMap, validators)
}

func VpnTunnelStatusBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["on_prem_subnet"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["on_prem_subnet"] = "OnPremSubnet"
	fields["traffic_stats"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(VpnTunnelTrafficStatsBindingType))
	fieldNameMap["traffic_stats"] = "TrafficStats"
	fields["last_info_message"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["last_info_message"] = "LastInfoMessage"
	fields["local_subnet"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["local_subnet"] = "LocalSubnet"
	fields["tunnel_state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["tunnel_state"] = "TunnelState"
	fields["failure_message"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["failure_message"] = "FailureMessage"
	fields["tunnel_status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["tunnel_status"] = "TunnelStatus"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.vpn_tunnel_status", fields, reflect.TypeOf(VpnTunnelStatus{}), fieldNameMap, validators)
}

func VpnTunnelTrafficStatsBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["packets_out"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["packets_out"] = "PacketsOut"
	fields["packet_received_errors"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["packet_received_errors"] = "PacketReceivedErrors"
	fields["rx_bytes_on_local_subnet"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["rx_bytes_on_local_subnet"] = "RxBytesOnLocalSubnet"
	fields["replay_errors"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["replay_errors"] = "ReplayErrors"
	fields["sequence_number_over_flow_errors"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["sequence_number_over_flow_errors"] = "SequenceNumberOverFlowErrors"
	fields["encryption_failures"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["encryption_failures"] = "EncryptionFailures"
	fields["integrity_errors"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["integrity_errors"] = "IntegrityErrors"
	fields["packet_sent_errors"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["packet_sent_errors"] = "PacketSentErrors"
	fields["decryption_failures"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["decryption_failures"] = "DecryptionFailures"
	fields["packets_in"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["packets_in"] = "PacketsIn"
	fields["tx_bytes_from_local_subnet"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["tx_bytes_from_local_subnet"] = "TxBytesFromLocalSubnet"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.vpn_tunnel_traffic_stats", fields, reflect.TypeOf(VpnTunnelTrafficStats{}), fieldNameMap, validators)
}

func VsanAvailableCapacityBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cost"] = vapiBindings_.NewStringType()
	fieldNameMap["cost"] = "Cost"
	fields["quality"] = vapiBindings_.NewStringType()
	fieldNameMap["quality"] = "Quality"
	fields["size"] = vapiBindings_.NewIntegerType()
	fieldNameMap["size"] = "Size"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.vsan_available_capacity", fields, reflect.TypeOf(VsanAvailableCapacity{}), fieldNameMap, validators)
}

func VsanClusterReconfigBiasBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["short_description"] = vapiBindings_.NewStringType()
	fieldNameMap["short_description"] = "ShortDescription"
	fields["full_description"] = vapiBindings_.NewStringType()
	fieldNameMap["full_description"] = "FullDescription"
	fields["id"] = vapiBindings_.NewStringType()
	fieldNameMap["id"] = "Id"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.vsan_cluster_reconfig_bias", fields, reflect.TypeOf(VsanClusterReconfigBias{}), fieldNameMap, validators)
}

func VsanClusterReconfigConstraintsBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["reconfig_biases"] = vapiBindings_.NewListType(vapiBindings_.NewReferenceType(VsanClusterReconfigBiasBindingType), reflect.TypeOf([]VsanClusterReconfigBias{}))
	fieldNameMap["reconfig_biases"] = "ReconfigBiases"
	fields["available_capacities"] = vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewListType(vapiBindings_.NewReferenceType(VsanAvailableCapacityBindingType), reflect.TypeOf([]VsanAvailableCapacity{})), reflect.TypeOf(map[string][]VsanAvailableCapacity{}))
	fieldNameMap["available_capacities"] = "AvailableCapacities"
	fields["default_capacities"] = vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewReferenceType(VsanAvailableCapacityBindingType), reflect.TypeOf(map[string]VsanAvailableCapacity{}))
	fieldNameMap["default_capacities"] = "DefaultCapacities"
	fields["hosts"] = vapiBindings_.NewIntegerType()
	fieldNameMap["hosts"] = "Hosts"
	fields["default_reconfig_bias_id"] = vapiBindings_.NewStringType()
	fieldNameMap["default_reconfig_bias_id"] = "DefaultReconfigBiasId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.vsan_cluster_reconfig_constraints", fields, reflect.TypeOf(VsanClusterReconfigConstraints{}), fieldNameMap, validators)
}

func VsanConfigConstraintsBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["max_capacity"] = vapiBindings_.NewIntegerType()
	fieldNameMap["max_capacity"] = "MaxCapacity"
	fields["recommended_capacities"] = vapiBindings_.NewListType(vapiBindings_.NewIntegerType(), reflect.TypeOf([]int64{}))
	fieldNameMap["recommended_capacities"] = "RecommendedCapacities"
	fields["supported_capacity_increment"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["supported_capacity_increment"] = "SupportedCapacityIncrement"
	fields["min_capacity"] = vapiBindings_.NewIntegerType()
	fieldNameMap["min_capacity"] = "MinCapacity"
	fields["num_hosts"] = vapiBindings_.NewIntegerType()
	fieldNameMap["num_hosts"] = "NumHosts"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.vsan_config_constraints", fields, reflect.TypeOf(VsanConfigConstraints{}), fieldNameMap, validators)
}

func VsanEncryptionConfigBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["port"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["port"] = "Port"
	fields["certificate"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["certificate"] = "Certificate"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.vsan_encryption_config", fields, reflect.TypeOf(VsanEncryptionConfig{}), fieldNameMap, validators)
}

func WcpDetailsBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["pod_cidr"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["pod_cidr"] = "PodCidr"
	fields["egress_cidr"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["egress_cidr"] = "EgressCidr"
	fields["ingress_cidr"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["ingress_cidr"] = "IngressCidr"
	fields["service_cidr"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["service_cidr"] = "ServiceCidr"
	fields["wcp_status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["wcp_status"] = "WcpStatus"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.wcp_details", fields, reflect.TypeOf(WcpDetails{}), fieldNameMap, validators)
}

func WitnessEsxBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["enum_state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["enum_state"] = "EnumState"
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["mac_address"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["mac_address"] = "MacAddress"
	fields["esx_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["esx_id"] = "EsxId"
	fields["hostname"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["hostname"] = "Hostname"
	fields["provider"] = vapiBindings_.NewStringType()
	fieldNameMap["provider"] = "Provider"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.witness_esx", fields, reflect.TypeOf(WitnessEsx{}), fieldNameMap, validators)
}

func XEniInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["association_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["association_id"] = "AssociationId"
	fields["trunk_eni_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["trunk_eni_id"] = "TrunkEniId"
	fields["x_eni_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["x_eni_id"] = "XEniId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmc.model.X_eni_info", fields, reflect.TypeOf(XEniInfo{}), fieldNameMap, validators)
}
