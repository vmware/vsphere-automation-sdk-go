/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: LibraryItem.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package ovf

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// The types of resources that can be created by deploying an OVF package and can be captured to create a library item using the ``LibraryItem`` interface.
var LibraryItem_DEPLOYABLE = []string{"VirtualMachine", "VirtualApp"}


// The ``DeployableIdentity`` class describes the resource created by a deployment, or the source resource from which library item can be created, by specifying its resource type and resource identifier.
type LibraryItemDeployableIdentity struct {
    // Type of the deployable resource.
	Type_ string
    // Identifier of the deployable resource.
	Id string
}

// The ``ResourcePoolDeploymentSpec`` class defines the deployment parameters that can be specified for the ``deploy()`` method where the deployment target is a resource pool. See LibraryItem#deploy.
type LibraryItemResourcePoolDeploymentSpec struct {
    // Name assigned to the deployed target virtual machine or virtual appliance.
	Name *string
    // Annotation assigned to the deployed target virtual machine or virtual appliance.
	Annotation *string
    // Whether to accept all End User License Agreements. See LibraryItemOvfSummary#eulas.
	AcceptAllEULA bool
    // Specification of the target network to use for sections of type ovf:NetworkSection in the OVF descriptor. The key in the map is the section identifier of the ovf:NetworkSection section in the OVF descriptor and the value is the target network to be used for deployment.
	NetworkMappings map[string]string
    // Specification of the target storage to use for sections of type vmw:StorageGroupSection in the OVF descriptor. The key in the map is the section identifier of the ovf:StorageGroupSection section in the OVF descriptor and the value is the target storage specification to be used for deployment. See LibraryItemStorageGroupMapping.
	StorageMappings map[string]LibraryItemStorageGroupMapping
    // Default storage provisioning type to use for all sections of type vmw:StorageSection in the OVF descriptor.
	StorageProvisioning *DiskProvisioningType
    // Default storage profile to use for all sections of type vmw:StorageSection in the OVF descriptor.
	StorageProfileId *string
    // The locale to use for parsing the OVF descriptor.
	Locale *string
    // Flags to be use for deployment. The supported flag values can be obtained using ImportFlag#list.
	Flags []string
    // Additional OVF parameters that may be needed for the deployment. Additional OVF parameters may be required by the OVF descriptor of the OVF package in the library item. Examples of OVF parameters that can be specified through this property include, but are not limited to: 
    //
    // * DeploymentOptionParams
    // * ExtraConfigParams
    // * IpAllocationParams
    // * PropertyParams
    // * ScaleOutParams
    // * VcenterExtensionParams
	AdditionalParameters []*data.StructValue
    // Default datastore to use for all sections of type vmw:StorageSection in the OVF descriptor.
	DefaultDatastoreId *string
}

// The ``StorageGroupMapping`` class defines the storage deployment target and storage provisioning type for a section of type vmw:StorageGroupSection in the OVF descriptor.
type LibraryItemStorageGroupMapping struct {
    // Type of storage deployment target to use for the vmw:StorageGroupSection section. The specified value must be LibraryItemStorageGroupMappingType#Type_DATASTORE or LibraryItemStorageGroupMappingType#Type_STORAGE_PROFILE.
	Type_ LibraryItemStorageGroupMappingType
    // Target datastore to be used for the storage group.
	DatastoreId *string
    // Target storage profile to be used for the storage group.
	StorageProfileId *string
    // Target provisioning type to use for the storage group.
	Provisioning *DiskProvisioningType
}

// The ``Type`` enumeration class defines the supported types of storage targets for sections of type vmw:StorageGroupSection in the OVF descriptor.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type LibraryItemStorageGroupMappingType string

const (
    // Storage deployment target is a datastore.
	LibraryItemStorageGroupMappingType_DATASTORE LibraryItemStorageGroupMappingType = "DATASTORE"
    // Storage deployment target is a storage profile.
	LibraryItemStorageGroupMappingType_STORAGE_PROFILE LibraryItemStorageGroupMappingType = "STORAGE_PROFILE"
)

func (t LibraryItemStorageGroupMappingType) LibraryItemStorageGroupMappingType() bool {
	switch t {
	case LibraryItemStorageGroupMappingType_DATASTORE:
		return true
	case LibraryItemStorageGroupMappingType_STORAGE_PROFILE:
		return true
	default:
		return false
	}
}


// The ``ResultInfo`` class defines the information returned along with the result of a ``create()`` or ``deploy()`` method to describe errors, warnings, and informational messages produced by the server.
type LibraryItemResultInfo struct {
    // Errors reported by the ``create()`` or ``deploy()`` method. These errors would have prevented the ``create()`` or ``deploy()`` method from completing successfully.
	Errors []OvfError
    // Warnings reported by the ``create()`` or ``deploy()`` method. These warnings would not have prevented the ``create()`` or ``deploy()`` method from completing successfully, but there might be issues that warrant attention.
	Warnings []OvfWarning
    // Information messages reported by the ``create()`` or ``deploy()`` method. For example, a non-required parameter was ignored.
	Information []OvfInfo
}

// The ``DeploymentResult`` class defines the result of the ``deploy()`` method. See LibraryItem#deploy.
type LibraryItemDeploymentResult struct {
    // Whether the ``deploy()`` method completed successfully.
	Succeeded bool
    // Identifier of the deployed resource entity.
	ResourceId *LibraryItemDeployableIdentity
    // Errors, warnings, and informational messages produced by the ``deploy()`` method.
	Error_ *LibraryItemResultInfo
}

// The ``DeploymentTarget`` class describes the location (target) where a virtual machine or virtual appliance should be deployed. It is used in the ``deploy()`` and ``filter()`` methods. See LibraryItem#deploy and LibraryItem#filter.
type LibraryItemDeploymentTarget struct {
    // Identifier of the resource pool to which the virtual machine or virtual appliance should be attached.
	ResourcePoolId string
    // Identifier of the target host on which the virtual machine or virtual appliance will run. The target host must be a member of the cluster that contains the resource pool identified by LibraryItemDeploymentTarget#resourcePoolId.
	HostId *string
    // Identifier of the vCenter folder that should contain the virtual machine or virtual appliance. The folder must be virtual machine folder.
	FolderId *string
}

// The ``OvfSummary`` class defines the result of the ``filter()`` method. See LibraryItem#filter. The properties in the class describe parameterizable information in the OVF descriptor, with respect to a deployment target, for the ``deploy()`` method. See LibraryItem#deploy.
type LibraryItemOvfSummary struct {
    // Default name for the virtual machine or virtual appliance.
	Name *string
    // Default annotation for the virtual machine or virtual appliance.
	Annotation *string
    // End User License Agreements specified in the OVF descriptor. All end user license agreements must be accepted in order for the ``deploy()`` method to succeed. See LibraryItemResourcePoolDeploymentSpec#acceptAllEula.
	EULAs []string
    // Section identifiers for sections of type ovf:NetworkSection in the OVF descriptor. These identifiers can be used as keys in LibraryItemResourcePoolDeploymentSpec#networkMappings.
	Networks []string
    // Section identifiers for sections of type vmw:StorageGroupSection in the OVF descriptor. These identifiers can be used as keys in LibraryItemResourcePoolDeploymentSpec#storageMappings.
	StorageGroups []string
    // Additional OVF parameters which can be specified for the deployment target. These OVF parameters can be inspected, optionally modified, and used as values in LibraryItemResourcePoolDeploymentSpec#additionalParameters for the ``deploy()`` method.
	AdditionalParams []*data.StructValue
}

// The ``CreateTarget`` class specifies the target library item when capturing a virtual machine or virtual appliance as an OVF package in a library item in a content library. The target can be an existing library item, which will be updated, creating a new version, or it can be a newly created library item in a specified library. See LibraryItem#create.
type LibraryItemCreateTarget struct {
    // Identifier of the library in which a new library item should be created. This property is not used if the ``libraryItemId`` property is specified.
	LibraryId *string
    // Identifier of the library item that should be should be updated.
	LibraryItemId *string
}

// The ``CreateSpec`` class defines the information used to create or update a library item containing an OVF package.
type LibraryItemCreateSpec struct {
    // Name to use in the OVF descriptor stored in the library item.
	Name *string
    // Description to use in the OVF descriptor stored in the library item.
	Description *string
    // Flags to use for OVF package creation. The supported flags can be obtained using ExportFlag#list.
	Flags []string
}

// The ``CreateResult`` class defines the result of the ``create()`` method. See LibraryItem#create.
type LibraryItemCreateResult struct {
    // Whether the ``create()`` method completed successfully.
	Succeeded bool
    // Identifier of the created or updated library item.
	OvfLibraryItemId *string
    // Errors, warnings, and informational messages produced by the ``create()`` method.
	Error_ *LibraryItemResultInfo
}



func libraryItemDeployInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["client_token"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["ovf_library_item_id"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fields["target"] = bindings.NewReferenceType(LibraryItemDeploymentTargetBindingType)
	fields["deployment_spec"] = bindings.NewReferenceType(LibraryItemResourcePoolDeploymentSpecBindingType)
	fieldNameMap["client_token"] = "ClientToken"
	fieldNameMap["ovf_library_item_id"] = "OvfLibraryItemId"
	fieldNameMap["target"] = "Target"
	fieldNameMap["deployment_spec"] = "DeploymentSpec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func libraryItemDeployOutputType() bindings.BindingType {
	return bindings.NewReferenceType(LibraryItemDeploymentResultBindingType)
}

func libraryItemDeployRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["client_token"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["ovf_library_item_id"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fields["target"] = bindings.NewReferenceType(LibraryItemDeploymentTargetBindingType)
	fields["deployment_spec"] = bindings.NewReferenceType(LibraryItemResourcePoolDeploymentSpecBindingType)
	fieldNameMap["client_token"] = "ClientToken"
	fieldNameMap["ovf_library_item_id"] = "OvfLibraryItemId"
	fieldNameMap["target"] = "Target"
	fieldNameMap["deployment_spec"] = "DeploymentSpec"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"InvalidArgument": 400,"NotFound": 404,"ResourceInaccessible": 500,"Unauthorized": 403})
}

func libraryItemFilterInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ovf_library_item_id"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fields["target"] = bindings.NewReferenceType(LibraryItemDeploymentTargetBindingType)
	fieldNameMap["ovf_library_item_id"] = "OvfLibraryItemId"
	fieldNameMap["target"] = "Target"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func libraryItemFilterOutputType() bindings.BindingType {
	return bindings.NewReferenceType(LibraryItemOvfSummaryBindingType)
}

func libraryItemFilterRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["ovf_library_item_id"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fields["target"] = bindings.NewReferenceType(LibraryItemDeploymentTargetBindingType)
	fieldNameMap["ovf_library_item_id"] = "OvfLibraryItemId"
	fieldNameMap["target"] = "Target"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"InvalidArgument": 400,"NotFound": 404,"ResourceInaccessible": 500})
}

func libraryItemCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["client_token"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["source"] = bindings.NewReferenceType(LibraryItemDeployableIdentityBindingType)
	fields["target"] = bindings.NewReferenceType(LibraryItemCreateTargetBindingType)
	fields["create_spec"] = bindings.NewReferenceType(LibraryItemCreateSpecBindingType)
	fieldNameMap["client_token"] = "ClientToken"
	fieldNameMap["source"] = "Source"
	fieldNameMap["target"] = "Target"
	fieldNameMap["create_spec"] = "CreateSpec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func libraryItemCreateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(LibraryItemCreateResultBindingType)
}

func libraryItemCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["client_token"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["source"] = bindings.NewReferenceType(LibraryItemDeployableIdentityBindingType)
	fields["target"] = bindings.NewReferenceType(LibraryItemCreateTargetBindingType)
	fields["create_spec"] = bindings.NewReferenceType(LibraryItemCreateSpecBindingType)
	fieldNameMap["client_token"] = "ClientToken"
	fieldNameMap["source"] = "Source"
	fieldNameMap["target"] = "Target"
	fieldNameMap["create_spec"] = "CreateSpec"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"InvalidArgument": 400,"NotFound": 404,"NotAllowedInCurrentState": 400,"ResourceInaccessible": 500,"ResourceBusy": 500})
}


func LibraryItemDeployableIdentityBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewStringType()
	fieldNameMap["type"] = "Type_"
	fields["id"] = bindings.NewIdType([]string{"null", "null"}, "type")
	fieldNameMap["id"] = "Id"
	var validators = []bindings.Validator{}
	isv1 := bindings.NewIsOneOfValidator(
		"type",
		[]string{
			"VirtualMachine",
			"VirtualApp",
		},
	)
	validators = append(validators, isv1)
	return bindings.NewStructType("com.vmware.vcenter.ovf.library_item.deployable_identity", fields, reflect.TypeOf(LibraryItemDeployableIdentity{}), fieldNameMap, validators)
}

func LibraryItemResourcePoolDeploymentSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["annotation"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["annotation"] = "Annotation"
	fields["accept_all_EULA"] = bindings.NewBooleanType()
	fieldNameMap["accept_all_EULA"] = "AcceptAllEULA"
	fields["network_mappings"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewIdType([]string{"Network"}, ""),reflect.TypeOf(map[string]string{})))
	fieldNameMap["network_mappings"] = "NetworkMappings"
	fields["storage_mappings"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewReferenceType(LibraryItemStorageGroupMappingBindingType),reflect.TypeOf(map[string]LibraryItemStorageGroupMapping{})))
	fieldNameMap["storage_mappings"] = "StorageMappings"
	fields["storage_provisioning"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.ovf.disk_provisioning_type", reflect.TypeOf(DiskProvisioningType(DiskProvisioningType_thin))))
	fieldNameMap["storage_provisioning"] = "StorageProvisioning"
	fields["storage_profile_id"] = bindings.NewOptionalType(bindings.NewIdType([]string{"StorageProfile"}, ""))
	fieldNameMap["storage_profile_id"] = "StorageProfileId"
	fields["locale"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["locale"] = "Locale"
	fields["flags"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["flags"] = "Flags"
	fields["additional_parameters"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(OvfParamsBindingType),}, bindings.JSONRPC), reflect.TypeOf([]*data.StructValue{})))
	fieldNameMap["additional_parameters"] = "AdditionalParameters"
	fields["default_datastore_id"] = bindings.NewOptionalType(bindings.NewIdType([]string{"Datastore"}, ""))
	fieldNameMap["default_datastore_id"] = "DefaultDatastoreId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.ovf.library_item.resource_pool_deployment_spec", fields, reflect.TypeOf(LibraryItemResourcePoolDeploymentSpec{}), fieldNameMap, validators)
}

func LibraryItemStorageGroupMappingBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewEnumType("com.vmware.vcenter.ovf.library_item.storage_group_mapping.type", reflect.TypeOf(LibraryItemStorageGroupMappingType(LibraryItemStorageGroupMappingType_DATASTORE)))
	fieldNameMap["type"] = "Type_"
	fields["datastore_id"] = bindings.NewOptionalType(bindings.NewIdType([]string{"Datastore"}, ""))
	fieldNameMap["datastore_id"] = "DatastoreId"
	fields["storage_profile_id"] = bindings.NewOptionalType(bindings.NewIdType([]string{"StorageProfile"}, ""))
	fieldNameMap["storage_profile_id"] = "StorageProfileId"
	fields["provisioning"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.ovf.disk_provisioning_type", reflect.TypeOf(DiskProvisioningType(DiskProvisioningType_thin))))
	fieldNameMap["provisioning"] = "Provisioning"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
		map[string][]bindings.FieldData{
			"DATASTORE": []bindings.FieldData{
				bindings.NewFieldData("datastore_id", true),
			},
			"STORAGE_PROFILE": []bindings.FieldData{
				bindings.NewFieldData("storage_profile_id", true),
			},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.ovf.library_item.storage_group_mapping", fields, reflect.TypeOf(LibraryItemStorageGroupMapping{}), fieldNameMap, validators)
}

func LibraryItemResultInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["errors"] = bindings.NewListType(bindings.NewReferenceType(OvfErrorBindingType), reflect.TypeOf([]OvfError{}))
	fieldNameMap["errors"] = "Errors"
	fields["warnings"] = bindings.NewListType(bindings.NewReferenceType(OvfWarningBindingType), reflect.TypeOf([]OvfWarning{}))
	fieldNameMap["warnings"] = "Warnings"
	fields["information"] = bindings.NewListType(bindings.NewReferenceType(OvfInfoBindingType), reflect.TypeOf([]OvfInfo{}))
	fieldNameMap["information"] = "Information"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.ovf.library_item.result_info", fields, reflect.TypeOf(LibraryItemResultInfo{}), fieldNameMap, validators)
}

func LibraryItemDeploymentResultBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["succeeded"] = bindings.NewBooleanType()
	fieldNameMap["succeeded"] = "Succeeded"
	fields["resource_id"] = bindings.NewOptionalType(bindings.NewReferenceType(LibraryItemDeployableIdentityBindingType))
	fieldNameMap["resource_id"] = "ResourceId"
	fields["error"] = bindings.NewOptionalType(bindings.NewReferenceType(LibraryItemResultInfoBindingType))
	fieldNameMap["error"] = "Error_"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.ovf.library_item.deployment_result", fields, reflect.TypeOf(LibraryItemDeploymentResult{}), fieldNameMap, validators)
}

func LibraryItemDeploymentTargetBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["resource_pool_id"] = bindings.NewIdType([]string{"ResourcePool"}, "")
	fieldNameMap["resource_pool_id"] = "ResourcePoolId"
	fields["host_id"] = bindings.NewOptionalType(bindings.NewIdType([]string{"HostSystem"}, ""))
	fieldNameMap["host_id"] = "HostId"
	fields["folder_id"] = bindings.NewOptionalType(bindings.NewIdType([]string{"Folder"}, ""))
	fieldNameMap["folder_id"] = "FolderId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.ovf.library_item.deployment_target", fields, reflect.TypeOf(LibraryItemDeploymentTarget{}), fieldNameMap, validators)
}

func LibraryItemOvfSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["annotation"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["annotation"] = "Annotation"
	fields["EULAs"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["EULAs"] = "EULAs"
	fields["networks"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["networks"] = "Networks"
	fields["storage_groups"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["storage_groups"] = "StorageGroups"
	fields["additional_params"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(OvfParamsBindingType),}, bindings.JSONRPC), reflect.TypeOf([]*data.StructValue{})))
	fieldNameMap["additional_params"] = "AdditionalParams"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.ovf.library_item.ovf_summary", fields, reflect.TypeOf(LibraryItemOvfSummary{}), fieldNameMap, validators)
}

func LibraryItemCreateTargetBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["library_id"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.content.Library"}, ""))
	fieldNameMap["library_id"] = "LibraryId"
	fields["library_item_id"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.content.library.Item"}, ""))
	fieldNameMap["library_item_id"] = "LibraryItemId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.ovf.library_item.create_target", fields, reflect.TypeOf(LibraryItemCreateTarget{}), fieldNameMap, validators)
}

func LibraryItemCreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["flags"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["flags"] = "Flags"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.ovf.library_item.create_spec", fields, reflect.TypeOf(LibraryItemCreateSpec{}), fieldNameMap, validators)
}

func LibraryItemCreateResultBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["succeeded"] = bindings.NewBooleanType()
	fieldNameMap["succeeded"] = "Succeeded"
	fields["ovf_library_item_id"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.content.library.Item"}, ""))
	fieldNameMap["ovf_library_item_id"] = "OvfLibraryItemId"
	fields["error"] = bindings.NewOptionalType(bindings.NewReferenceType(LibraryItemResultInfoBindingType))
	fieldNameMap["error"] = "Error_"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.ovf.library_item.create_result", fields, reflect.TypeOf(LibraryItemCreateResult{}), fieldNameMap, validators)
}

