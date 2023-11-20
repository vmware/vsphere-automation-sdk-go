// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for package: com.vmware.vmwarecloud.vmc_aws_provider.model.
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

type ClusterRenameIntent struct {
	// New name for the cluster
	ClusterName *string
	// Current name of the cluster
	CurrentClusterName *string
}

func (s *ClusterRenameIntent) GetType__() vapiBindings_.BindingType {
	return ClusterRenameIntentBindingType()
}

func (s *ClusterRenameIntent) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ClusterRenameIntent._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type AddHostIntent struct {
	// Number of hosts
	NumHosts int64
	// The cluster id to which the host(s) shall be added.
	ClusterId string
}

func (s *AddHostIntent) GetType__() vapiBindings_.BindingType {
	return AddHostIntentBindingType()
}

func (s *AddHostIntent) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for AddHostIntent._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type MountNfsDatastoreIntent struct {
	MountInfo []VmcNfsMountInfo
}

func (s *MountNfsDatastoreIntent) GetType__() vapiBindings_.BindingType {
	return MountNfsDatastoreIntentBindingType()
}

func (s *MountNfsDatastoreIntent) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for MountNfsDatastoreIntent._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type MountNfsDatastorePrecheckIntent struct {
	// NFS server IP addresses.
	StorageEndpoint []string
	// Storage vendor for datastore.
	StorageVendor *string
}

func (s *MountNfsDatastorePrecheckIntent) GetType__() vapiBindings_.BindingType {
	return MountNfsDatastorePrecheckIntentBindingType()
}

func (s *MountNfsDatastorePrecheckIntent) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for MountNfsDatastorePrecheckIntent._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type UnmountNfsDatastoreIntent struct {
	DatastoreId   []string
	DatastoreName []string
}

func (s *UnmountNfsDatastoreIntent) GetType__() vapiBindings_.BindingType {
	return UnmountNfsDatastoreIntentBindingType()
}

func (s *UnmountNfsDatastoreIntent) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for UnmountNfsDatastoreIntent._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type RestErrorV2 struct {
	// Unique error occurrence identifier
	Instance string
	// HTTP status code
	Status int64
	// Unique identifier for the type of error
	Type_ string
	// Array of error messages
	Messages []ErrorMessage
}

func (s *RestErrorV2) GetType__() vapiBindings_.BindingType {
	return RestErrorV2BindingType()
}

func (s *RestErrorV2) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for RestErrorV2._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type VmcNfsMountInfo struct {
	// Name of datastore to be attached.
	DatastoreName string
	// Datastores IP addresses.
	StorageEndpoint []string
	// Volume path to be mounted
	VolumePath string
	AccessMode *string
	Protocol   *string
	// Storage vendor for datastore.
	StorageVendor string
	// Maximum request queue depth for datastore
	MaxqDepth *int64
	// Number of TCP connections for the particular NFSv3 Server during NAS volume mount operation
	Connections *int64
}

// Supported datastore access modes.
const VmcNfsMountInfo_ACCESS_MODE_NFS_ACCESS_MODE_READONLY = "READONLY"

// Supported datastore access modes.
const VmcNfsMountInfo_ACCESS_MODE_NFS_ACCESS_MODE_READWRITE = "READWRITE"

// Supported datastore versions.
const VmcNfsMountInfo_PROTOCOL_NFS_VERSION_NFS = "NFS"

func (s *VmcNfsMountInfo) GetType__() vapiBindings_.BindingType {
	return VmcNfsMountInfoBindingType()
}

func (s *VmcNfsMountInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for VmcNfsMountInfo._GetDataValue method - %s",
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
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_aws_provider.model.activity_response", fields, reflect.TypeOf(ActivityResponse{}), fieldNameMap, validators)
}

func ClusterRenameIntentBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cluster_name"] = "ClusterName"
	fields["current_cluster_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["current_cluster_name"] = "CurrentClusterName"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_aws_provider.model.cluster_rename_intent", fields, reflect.TypeOf(ClusterRenameIntent{}), fieldNameMap, validators)
}

func AddHostIntentBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["num_hosts"] = vapiBindings_.NewIntegerType()
	fieldNameMap["num_hosts"] = "NumHosts"
	fields["cluster_id"] = vapiBindings_.NewStringType()
	fieldNameMap["cluster_id"] = "ClusterId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_aws_provider.model.add_host_intent", fields, reflect.TypeOf(AddHostIntent{}), fieldNameMap, validators)
}

func MountNfsDatastoreIntentBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["mount_info"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(VmcNfsMountInfoBindingType), reflect.TypeOf([]VmcNfsMountInfo{})))
	fieldNameMap["mount_info"] = "MountInfo"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_aws_provider.model.mount_nfs_datastore_intent", fields, reflect.TypeOf(MountNfsDatastoreIntent{}), fieldNameMap, validators)
}

func MountNfsDatastorePrecheckIntentBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["storage_endpoint"] = vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["storage_endpoint"] = "StorageEndpoint"
	fields["storage_vendor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["storage_vendor"] = "StorageVendor"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_aws_provider.model.mount_nfs_datastore_precheck_intent", fields, reflect.TypeOf(MountNfsDatastorePrecheckIntent{}), fieldNameMap, validators)
}

func UnmountNfsDatastoreIntentBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["datastore_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["datastore_id"] = "DatastoreId"
	fields["datastore_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["datastore_name"] = "DatastoreName"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_aws_provider.model.unmount_nfs_datastore_intent", fields, reflect.TypeOf(UnmountNfsDatastoreIntent{}), fieldNameMap, validators)
}

func RestErrorV2BindingType() vapiBindings_.BindingType {
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
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_aws_provider.model.rest_error_v2", fields, reflect.TypeOf(RestErrorV2{}), fieldNameMap, validators)
}

func VmcNfsMountInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["datastore_name"] = vapiBindings_.NewStringType()
	fieldNameMap["datastore_name"] = "DatastoreName"
	fields["storage_endpoint"] = vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["storage_endpoint"] = "StorageEndpoint"
	fields["volume_path"] = vapiBindings_.NewStringType()
	fieldNameMap["volume_path"] = "VolumePath"
	fields["access_mode"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["access_mode"] = "AccessMode"
	fields["protocol"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["protocol"] = "Protocol"
	fields["storage_vendor"] = vapiBindings_.NewStringType()
	fieldNameMap["storage_vendor"] = "StorageVendor"
	fields["max_q_depth"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["max_q_depth"] = "MaxqDepth"
	fields["connections"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["connections"] = "Connections"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_aws_provider.model.vmc_nfs_mount_info", fields, reflect.TypeOf(VmcNfsMountInfo{}), fieldNameMap, validators)
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
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_aws_provider.model.error_message", fields, reflect.TypeOf(ErrorMessage{}), fieldNameMap, validators)
}
