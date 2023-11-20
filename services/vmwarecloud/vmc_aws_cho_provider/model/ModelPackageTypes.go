// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for package: com.vmware.vmwarecloud.vmc_aws_cho_provider.model.
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

// Response of Activity API
type ActivityResponse struct {
	// ID of the create activity.
	Id *string
	// Org id associated with the activity.
	OrgId *string
	// Resource ID on which activity is performed.
	ResourceId *string
}

func (s *ActivityResponse) GetType__() vapiBindings_.BindingType {
	return ActivityResponseBindingType()
}

func (s *ActivityResponse) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ActivityResponse._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type DeleteHostIntent struct {
	// The number of hosts to delete.
	NumHosts int64
	// The cluster id from which the host(s) shall be deleted.
	ClusterId string
	// A list of Host IDs to be removed from cluster
	HostIds []string
}

func (s *DeleteHostIntent) GetType__() vapiBindings_.BindingType {
	return DeleteHostIntentBindingType()
}

func (s *DeleteHostIntent) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for DeleteHostIntent._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type DeleteHostInternalIntent struct {
	// The availability zone in which the host needs to be deleted.
	AvailabilityZone *string
	// To indicate context of a host provisioning. If internal, then host is not billable and if user, the host is billable to customer
	HostOperationContext *string
	DeleteHostReason     *DeleteHostReason
	RemoveType           string
	// The number of hosts to delete.
	NumHosts int64
	// The cluster id from which the host(s) shall be deleted.
	ClusterId string
	// A list of Host IDs to be removed from cluster
	HostIds []string
}

// To indicate context of a host provisioning. If internal, then host is not billable and if user, the host is billable to customer
const DeleteHostInternalIntent_HOST_OPERATION_CONTEXT_INTERNAL = "INTERNAL"

// To indicate context of a host provisioning. If internal, then host is not billable and if user, the host is billable to customer
const DeleteHostInternalIntent_HOST_OPERATION_CONTEXT_USER = "USER"

// The remove type for the esx hosts
const DeleteHostInternalIntent_REMOVE_TYPE_REMOVE_TYPE_REMOVE = "remove"

// The remove type for the esx hosts
const DeleteHostInternalIntent_REMOVE_TYPE_REMOVE_TYPE_DISCONNECTED_REMOVE = "disconnected-remove"

// The remove type for the esx hosts
const DeleteHostInternalIntent_REMOVE_TYPE_REMOVE_TYPE_FORCE_REMOVE = "force-remove"

func (s *DeleteHostInternalIntent) GetType__() vapiBindings_.BindingType {
	return DeleteHostInternalIntentBindingType()
}

func (s *DeleteHostInternalIntent) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for DeleteHostInternalIntent._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type AddClusterIntent struct {
	// Number of hosts
	NumHosts int64
	// Customize CPU cores on hosts in a cluster. Specify number of cores to be enabled on hosts in a cluster.
	HostCpuCoresCount *int64
	HostInstanceType  VmcAwsHostInstanceType
	// A list of availability zones for cluster
	AvailabilityZones []string
	// Details about the vsan_version
	VsanVersion *string
}

// Details about the vsan_version
const AddClusterIntent_VSAN_VERSION_VSAN1 = "vsan1"

// Details about the vsan_version
const AddClusterIntent_VSAN_VERSION_VSANESA = "vsanesa"

// Details about the vsan_version
const AddClusterIntent_VSAN_VERSION_NOVSAN = "novsan"

func (s *AddClusterIntent) GetType__() vapiBindings_.BindingType {
	return AddClusterIntentBindingType()
}

func (s *AddClusterIntent) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for AddClusterIntent._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type PublishMsftLicenseIntent struct {
	// The cluster id from which the host(s) shall be deleted.
	ClusterId       string
	AwsLicensesMsft AwsLicensesMsftConfig
}

func (s *PublishMsftLicenseIntent) GetType__() vapiBindings_.BindingType {
	return PublishMsftLicenseIntentBindingType()
}

func (s *PublishMsftLicenseIntent) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for PublishMsftLicenseIntent._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type RestError struct {
	// Unique error occurrence identifier
	Instance string
	// HTTP status code
	Status int64
	// Unique identifier for the type of error
	Type_ string
	// Array of error messages
	Messages []ErrorMessage
}

func (s *RestError) GetType__() vapiBindings_.BindingType {
	return RestErrorBindingType()
}

func (s *RestError) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for RestError._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type ModifyInternalHostCountIntent struct {
	// The cluster id from which the internal host count shall be modified.
	ClusterId string
	// increment for increase internal host count, decrement for decrease internal host count
	Action string
}

func (s *ModifyInternalHostCountIntent) GetType__() vapiBindings_.BindingType {
	return ModifyInternalHostCountIntentBindingType()
}

func (s *ModifyInternalHostCountIntent) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ModifyInternalHostCountIntent._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Maintenance host definition
type MaintenanceHost struct {
	// ID of MaintenanceHost.
	Id *string
	// ID of the durable host.
	DurableHostId *string
	// Org id associated with the MaintenanceHost.
	OrgId *string
	// ID of the deployment.
	DeploymentId *string
	// Id of the cluster
	ClusterId *string
	// AWS region of SPLA
	AwsRegion *string
}

func (s *MaintenanceHost) GetType__() vapiBindings_.BindingType {
	return MaintenanceHostBindingType()
}

func (s *MaintenanceHost) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for MaintenanceHost._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// SPLA maintenance hosts info
type SplaMaintenanceHostsInfo struct {
	// AWS region of SPLA
	AwsRegion *string
	// A list of MaintenanceHost object
	DurableIdentifiers []MaintenanceHost
}

func (s *SplaMaintenanceHostsInfo) GetType__() vapiBindings_.BindingType {
	return SplaMaintenanceHostsInfoBindingType()
}

func (s *SplaMaintenanceHostsInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for SplaMaintenanceHostsInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type DeleteClusterIntent struct {
	// The cluster id that should be deleted.
	ClusterId string
}

func (s *DeleteClusterIntent) GetType__() vapiBindings_.BindingType {
	return DeleteClusterIntentBindingType()
}

func (s *DeleteClusterIntent) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for DeleteClusterIntent._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type WorkflowRoutingPayload struct {
	ClusterType   string
	FeatureType   *FeatureType
	InstanceType  *WorkflowInstanceType
	WorkflowName  string
	SourceOfTruth string
}

// Cluster Type
const WorkflowRoutingPayload_CLUSTER_TYPE_CLUSTER_TYPE_1NODE = "1NODE"

// Cluster Type
const WorkflowRoutingPayload_CLUSTER_TYPE_CLUSTER_TYPE_2NODE = "2NODE"

// Cluster Type
const WorkflowRoutingPayload_CLUSTER_TYPE_CLUSTER_TYPE_STANDARD = "STANDARD"

// Cluster Type
const WorkflowRoutingPayload_CLUSTER_TYPE_CLUSTER_TYPE_MULTI_AZ = "MULTI_AZ"

// Type of workflow
const WorkflowRoutingPayload_WORKFLOW_NAME_WORKFLOW_NAME_SDDC_PROVISION = "SDDC_PROVISION"

// Type of workflow
const WorkflowRoutingPayload_WORKFLOW_NAME_WORKFLOW_NAME_SDDC_DELETE = "SDDC_DELETE"

// Type of workflow
const WorkflowRoutingPayload_WORKFLOW_NAME_WORKFLOW_NAME_CLUSTER_CREATE = "CLUSTER_CREATE"

// Type of workflow
const WorkflowRoutingPayload_WORKFLOW_NAME_WORKFLOW_NAME_TWO_NODE_CLUSTER_CREATE = "TWO_NODE_CLUSTER_CREATE"

// Type of workflow
const WorkflowRoutingPayload_WORKFLOW_NAME_WORKFLOW_NAME_ONE_NODE_CLUSTER_CREATE = "ONE_NODE_CLUSTER_CREATE"

// Type of workflow
const WorkflowRoutingPayload_WORKFLOW_NAME_WORKFLOW_NAME_ADD_HOST = "ADD_HOST"

// Type of workflow
const WorkflowRoutingPayload_WORKFLOW_NAME_WORKFLOW_NAME_REMOVE_HOST = "REMOVE_HOST"

// Type of workflow
const WorkflowRoutingPayload_WORKFLOW_NAME_WORKFLOW_NAME_DELETE_CLUSTER = "DELETE_CLUSTER"

// Type of workflow
const WorkflowRoutingPayload_WORKFLOW_NAME_WORKFLOW_NAME_TWO_NODE_DELETE_CLUSTER = "TWO_NODE_DELETE_CLUSTER"

// Type of workflow
const WorkflowRoutingPayload_WORKFLOW_NAME_WORKFLOW_NAME_ONE_NODE_DELETE_CLUSTER = "ONE_NODE_DELETE_CLUSTER"

// Type of workflow
const WorkflowRoutingPayload_WORKFLOW_NAME_WORKFLOW_NAME_SCALE_UP = "SCALE_UP"

// Type of workflow
const WorkflowRoutingPayload_WORKFLOW_NAME_WORKFLOW_NAME_SCALE_DOWN = "SCALE_DOWN"

// Source Of Truth
const WorkflowRoutingPayload_SOURCE_OF_TRUTH_SOURCE_OF_TRUTH_V1 = "V1"

// Source Of Truth
const WorkflowRoutingPayload_SOURCE_OF_TRUTH_SOURCE_OF_TRUTH_V2 = "V2"

// Source Of Truth
const WorkflowRoutingPayload_SOURCE_OF_TRUTH_SOURCE_OF_TRUTH_NA = "NA"

func (s *WorkflowRoutingPayload) GetType__() vapiBindings_.BindingType {
	return WorkflowRoutingPayloadBindingType()
}

func (s *WorkflowRoutingPayload) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for WorkflowRoutingPayload._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type WorkflowDestinationResponse struct {
	Destination *string
}

// Workflow route type
const WorkflowDestinationResponse_DESTINATION_WORKFLOW_ROUTE_TYPE_V1 = "V1"

// Workflow route type
const WorkflowDestinationResponse_DESTINATION_WORKFLOW_ROUTE_TYPE_V2 = "V2"

func (s *WorkflowDestinationResponse) GetType__() vapiBindings_.BindingType {
	return WorkflowDestinationResponseBindingType()
}

func (s *WorkflowDestinationResponse) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for WorkflowDestinationResponse._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type RuleVersionPromotionPayload struct {
	RuleType string
	// Rule version
	Version int64
}

// Version Type
const RuleVersionPromotionPayload_RULE_TYPE_RULE_VERSION_TYPE_SNAPSHOT = "SNAPSHOT"

// Version Type
const RuleVersionPromotionPayload_RULE_TYPE_RULE_VERSION_TYPE_RELEASE = "RELEASE"

func (s *RuleVersionPromotionPayload) GetType__() vapiBindings_.BindingType {
	return RuleVersionPromotionPayloadBindingType()
}

func (s *RuleVersionPromotionPayload) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for RuleVersionPromotionPayload._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type DeleteHostReason struct {
}

func (s *DeleteHostReason) GetType__() vapiBindings_.BindingType {
	return DeleteHostReasonBindingType()
}

func (s *DeleteHostReason) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for DeleteHostReason._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type VmcAwsHostInstanceType struct {
}

func (s *VmcAwsHostInstanceType) GetType__() vapiBindings_.BindingType {
	return VmcAwsHostInstanceTypeBindingType()
}

func (s *VmcAwsHostInstanceType) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for VmcAwsHostInstanceType._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// The SPLA licensing configuration for the cluster
type AwsLicensesMsftConfig struct {
	// Windows Server license config
	WindowsLicensing *string
	// SQL Server license config
	MssqlLicensing *string
	// Academic institution recognized by Microsoft
	Academic bool
}

// Windows Server license config
const AwsLicensesMsftConfig_WINDOWS_LICENSING_DISABLED = "DISABLED"

// Windows Server license config
const AwsLicensesMsftConfig_WINDOWS_LICENSING_ENABLED = "ENABLED"

// Windows Server license config
const AwsLicensesMsftConfig_WINDOWS_LICENSING_CUSTOMER_SUPPLIED = "CUSTOMER_SUPPLIED"

// SQL Server license config
const AwsLicensesMsftConfig_MSSQL_LICENSING_DISABLED = "DISABLED"

// SQL Server license config
const AwsLicensesMsftConfig_MSSQL_LICENSING_ENABLED = "ENABLED"

// SQL Server license config
const AwsLicensesMsftConfig_MSSQL_LICENSING_CUSTOMER_SUPPLIED = "CUSTOMER_SUPPLIED"

func (s *AwsLicensesMsftConfig) GetType__() vapiBindings_.BindingType {
	return AwsLicensesMsftConfigBindingType()
}

func (s *AwsLicensesMsftConfig) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for AwsLicensesMsftConfig._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type ErrorMessage struct {
	// Globally resolvable bundle key.
	Id string
	// Error message.
	Message string
	// Additional information to further contextualize the error.
	Source *string
	Args   []string
}

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

type FeatureType struct {
}

func (s *FeatureType) GetType__() vapiBindings_.BindingType {
	return FeatureTypeBindingType()
}

func (s *FeatureType) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for FeatureType._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type WorkflowInstanceType struct {
}

func (s *WorkflowInstanceType) GetType__() vapiBindings_.BindingType {
	return WorkflowInstanceTypeBindingType()
}

func (s *WorkflowInstanceType) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for WorkflowInstanceType._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

func ActivityResponseBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["org_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fields["resource_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["resource_id"] = "ResourceId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_aws_cho_provider.model.activity_response", fields, reflect.TypeOf(ActivityResponse{}), fieldNameMap, validators)
}

func DeleteHostIntentBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["num_hosts"] = vapiBindings_.NewIntegerType()
	fieldNameMap["num_hosts"] = "NumHosts"
	fields["cluster_id"] = vapiBindings_.NewStringType()
	fieldNameMap["cluster_id"] = "ClusterId"
	fields["host_ids"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["host_ids"] = "HostIds"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_aws_cho_provider.model.delete_host_intent", fields, reflect.TypeOf(DeleteHostIntent{}), fieldNameMap, validators)
}

func DeleteHostInternalIntentBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["availability_zone"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["availability_zone"] = "AvailabilityZone"
	fields["host_operation_context"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["host_operation_context"] = "HostOperationContext"
	fields["delete_host_reason"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(DeleteHostReasonBindingType))
	fieldNameMap["delete_host_reason"] = "DeleteHostReason"
	fields["remove_type"] = vapiBindings_.NewStringType()
	fieldNameMap["remove_type"] = "RemoveType"
	fields["num_hosts"] = vapiBindings_.NewIntegerType()
	fieldNameMap["num_hosts"] = "NumHosts"
	fields["cluster_id"] = vapiBindings_.NewStringType()
	fieldNameMap["cluster_id"] = "ClusterId"
	fields["host_ids"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["host_ids"] = "HostIds"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_aws_cho_provider.model.delete_host_internal_intent", fields, reflect.TypeOf(DeleteHostInternalIntent{}), fieldNameMap, validators)
}

func AddClusterIntentBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["num_hosts"] = vapiBindings_.NewIntegerType()
	fieldNameMap["num_hosts"] = "NumHosts"
	fields["host_cpu_cores_count"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["host_cpu_cores_count"] = "HostCpuCoresCount"
	fields["host_instance_type"] = vapiBindings_.NewReferenceType(VmcAwsHostInstanceTypeBindingType)
	fieldNameMap["host_instance_type"] = "HostInstanceType"
	fields["availability_zones"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["availability_zones"] = "AvailabilityZones"
	fields["vsan_version"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vsan_version"] = "VsanVersion"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_aws_cho_provider.model.add_cluster_intent", fields, reflect.TypeOf(AddClusterIntent{}), fieldNameMap, validators)
}

func PublishMsftLicenseIntentBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster_id"] = vapiBindings_.NewStringType()
	fieldNameMap["cluster_id"] = "ClusterId"
	fields["aws_licenses_msft"] = vapiBindings_.NewReferenceType(AwsLicensesMsftConfigBindingType)
	fieldNameMap["aws_licenses_msft"] = "AwsLicensesMsft"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_aws_cho_provider.model.publish_msft_license_intent", fields, reflect.TypeOf(PublishMsftLicenseIntent{}), fieldNameMap, validators)
}

func RestErrorBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["instance"] = vapiBindings_.NewStringType()
	fieldNameMap["instance"] = "Instance"
	fields["status"] = vapiBindings_.NewIntegerType()
	fieldNameMap["status"] = "Status"
	fields["type"] = vapiBindings_.NewStringType()
	fieldNameMap["type"] = "Type_"
	fields["messages"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(ErrorMessageBindingType), reflect.TypeOf([]ErrorMessage{})))
	fieldNameMap["messages"] = "Messages"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_aws_cho_provider.model.rest_error", fields, reflect.TypeOf(RestError{}), fieldNameMap, validators)
}

func ModifyInternalHostCountIntentBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster_id"] = vapiBindings_.NewStringType()
	fieldNameMap["cluster_id"] = "ClusterId"
	fields["action"] = vapiBindings_.NewStringType()
	fieldNameMap["action"] = "Action"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_aws_cho_provider.model.modify_internal_host_count_intent", fields, reflect.TypeOf(ModifyInternalHostCountIntent{}), fieldNameMap, validators)
}

func MaintenanceHostBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["durable_host_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["durable_host_id"] = "DurableHostId"
	fields["org_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fields["deployment_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["deployment_id"] = "DeploymentId"
	fields["cluster_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cluster_id"] = "ClusterId"
	fields["aws_region"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["aws_region"] = "AwsRegion"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_aws_cho_provider.model.maintenance_host", fields, reflect.TypeOf(MaintenanceHost{}), fieldNameMap, validators)
}

func SplaMaintenanceHostsInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["aws_region"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["aws_region"] = "AwsRegion"
	fields["durable_identifiers"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(MaintenanceHostBindingType), reflect.TypeOf([]MaintenanceHost{})))
	fieldNameMap["durable_identifiers"] = "DurableIdentifiers"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_aws_cho_provider.model.spla_maintenance_hosts_info", fields, reflect.TypeOf(SplaMaintenanceHostsInfo{}), fieldNameMap, validators)
}

func DeleteClusterIntentBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster_id"] = vapiBindings_.NewStringType()
	fieldNameMap["cluster_id"] = "ClusterId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_aws_cho_provider.model.delete_cluster_intent", fields, reflect.TypeOf(DeleteClusterIntent{}), fieldNameMap, validators)
}

func WorkflowRoutingPayloadBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster_type"] = vapiBindings_.NewStringType()
	fieldNameMap["cluster_type"] = "ClusterType"
	fields["feature_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(FeatureTypeBindingType))
	fieldNameMap["feature_type"] = "FeatureType"
	fields["instance_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(WorkflowInstanceTypeBindingType))
	fieldNameMap["instance_type"] = "InstanceType"
	fields["workflow_name"] = vapiBindings_.NewStringType()
	fieldNameMap["workflow_name"] = "WorkflowName"
	fields["source_of_truth"] = vapiBindings_.NewStringType()
	fieldNameMap["source_of_truth"] = "SourceOfTruth"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_aws_cho_provider.model.workflow_routing_payload", fields, reflect.TypeOf(WorkflowRoutingPayload{}), fieldNameMap, validators)
}

func WorkflowDestinationResponseBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["destination"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["destination"] = "Destination"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_aws_cho_provider.model.workflow_destination_response", fields, reflect.TypeOf(WorkflowDestinationResponse{}), fieldNameMap, validators)
}

func RuleVersionPromotionPayloadBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["rule_type"] = vapiBindings_.NewStringType()
	fieldNameMap["rule_type"] = "RuleType"
	fields["version"] = vapiBindings_.NewIntegerType()
	fieldNameMap["version"] = "Version"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_aws_cho_provider.model.rule_version_promotion_payload", fields, reflect.TypeOf(RuleVersionPromotionPayload{}), fieldNameMap, validators)
}

func DeleteHostReasonBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_aws_cho_provider.model.delete_host_reason", fields, reflect.TypeOf(DeleteHostReason{}), fieldNameMap, validators)
}

func VmcAwsHostInstanceTypeBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_aws_cho_provider.model.vmc_aws_host_instance_type", fields, reflect.TypeOf(VmcAwsHostInstanceType{}), fieldNameMap, validators)
}

func AwsLicensesMsftConfigBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["windows_licensing"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["windows_licensing"] = "WindowsLicensing"
	fields["mssql_licensing"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["mssql_licensing"] = "MssqlLicensing"
	fields["academic"] = vapiBindings_.NewBooleanType()
	fieldNameMap["academic"] = "Academic"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_aws_cho_provider.model.aws_licenses_msft_config", fields, reflect.TypeOf(AwsLicensesMsftConfig{}), fieldNameMap, validators)
}

func ErrorMessageBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewStringType()
	fieldNameMap["id"] = "Id"
	fields["message"] = vapiBindings_.NewStringType()
	fieldNameMap["message"] = "Message"
	fields["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["source"] = "Source"
	fields["args"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["args"] = "Args"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_aws_cho_provider.model.error_message", fields, reflect.TypeOf(ErrorMessage{}), fieldNameMap, validators)
}

func FeatureTypeBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_aws_cho_provider.model.feature_type", fields, reflect.TypeOf(FeatureType{}), fieldNameMap, validators)
}

func WorkflowInstanceTypeBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_aws_cho_provider.model.workflow_instance_type", fields, reflect.TypeOf(WorkflowInstanceType{}), fieldNameMap, validators)
}
