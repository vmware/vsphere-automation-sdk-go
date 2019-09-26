/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: LibraryItem.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package libraryItem

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/vcenter/ovf"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// The types of resources that can be created by deploying an OVF package and can be captured to create a library item using the ``LibraryItem`` interface.
var LibraryItem_DEPLOYABLE = []string{"VirtualMachine", "VirtualApp"}



// The ``DeployableIdentity`` class describes the resource created by a deployment, or the source resource from which library item can be created, by specifying its resource type and resource identifier.
 type DeployableIdentity struct {
    Type_ string
    Id string
}






// The ``ResourcePoolDeploymentSpec`` class defines the deployment parameters that can be specified for the ``deploy()`` method where the deployment target is a resource pool. See LibraryItem#deploy.
 type ResourcePoolDeploymentSpec struct {
    Name *string
    Annotation *string
    AcceptAllEULA bool
    NetworkMappings map[string]string
    StorageMappings map[string]StorageGroupMapping
    StorageProvisioning *ovf.DiskProvisioningType
    StorageProfileId *string
    Locale *string
    Flags []string
    AdditionalParameters []*data.StructValue
    DefaultDatastoreId *string
}






// The ``StorageGroupMapping`` class defines the storage deployment target and storage provisioning type for a section of type vmw:StorageGroupSection in the OVF descriptor.
 type StorageGroupMapping struct {
    Type_ StorageGroupMapping_Type
    DatastoreId *string
    StorageProfileId *string
    Provisioning *ovf.DiskProvisioningType
}




    
    // The ``Type`` enumeration class defines the supported types of storage targets for sections of type vmw:StorageGroupSection in the OVF descriptor.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type StorageGroupMapping_Type string

    const (
        // Storage deployment target is a datastore.
         StorageGroupMapping_Type_DATASTORE StorageGroupMapping_Type = "DATASTORE"
        // Storage deployment target is a storage profile.
         StorageGroupMapping_Type_STORAGE_PROFILE StorageGroupMapping_Type = "STORAGE_PROFILE"
    )

    func (t StorageGroupMapping_Type) StorageGroupMapping_Type() bool {
        switch t {
            case StorageGroupMapping_Type_DATASTORE:
                return true
            case StorageGroupMapping_Type_STORAGE_PROFILE:
                return true
            default:
                return false
        }
    }



// The ``ResultInfo`` class defines the information returned along with the result of a ``create()`` or ``deploy()`` method to describe errors, warnings, and informational messages produced by the server.
 type ResultInfo struct {
    Errors []ovf.OvfError
    Warnings []ovf.OvfWarning
    Information []ovf.OvfInfo
}






// The ``DeploymentResult`` class defines the result of the ``deploy()`` method. See LibraryItem#deploy.
 type DeploymentResult struct {
    Succeeded bool
    ResourceId *DeployableIdentity
    Error *ResultInfo
}






// The ``DeploymentTarget`` class describes the location (target) where a virtual machine or virtual appliance should be deployed. It is used in the ``deploy()`` and ``filter()`` methods. See LibraryItem#deploy and LibraryItem#filter.
 type DeploymentTarget struct {
    ResourcePoolId string
    HostId *string
    FolderId *string
}






// The ``OvfSummary`` class defines the result of the ``filter()`` method. See LibraryItem#filter. The properties in the class describe parameterizable information in the OVF descriptor, with respect to a deployment target, for the ``deploy()`` method. See LibraryItem#deploy.
 type OvfSummary struct {
    Name *string
    Annotation *string
    EULAs []string
    Networks []string
    StorageGroups []string
    AdditionalParams []*data.StructValue
}






// The ``CreateTarget`` class specifies the target library item when capturing a virtual machine or virtual appliance as an OVF package in a library item in a content library. The target can be an existing library item, which will be updated, creating a new version, or it can be a newly created library item in a specified library. See LibraryItem#create.
 type CreateTarget struct {
    LibraryId *string
    LibraryItemId *string
}






// The ``CreateSpec`` class defines the information used to create or update a library item containing an OVF package.
 type CreateSpec struct {
    Name *string
    Description *string
    Flags []string
}






// The ``CreateResult`` class defines the result of the ``create()`` method. See LibraryItem#create.
 type CreateResult struct {
    Succeeded bool
    OvfLibraryItemId *string
    Error *ResultInfo
}









func deployInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["client_token"] = bindings.NewOptionalType(bindings.NewStringType())
    fields["ovf_library_item_id"] = bindings.NewIdType([]string {"com.vmware.content.library.Item"}, "")
    fields["target"] = bindings.NewReferenceType(DeploymentTargetBindingType)
    fields["deployment_spec"] = bindings.NewReferenceType(ResourcePoolDeploymentSpecBindingType)
    fieldNameMap["client_token"] = "ClientToken"
    fieldNameMap["ovf_library_item_id"] = "OvfLibraryItemId"
    fieldNameMap["target"] = "Target"
    fieldNameMap["deployment_spec"] = "DeploymentSpec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func deployOutputType() bindings.BindingType {
    return bindings.NewReferenceType(DeploymentResultBindingType)
}

func deployRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"InvalidArgument": 400,"NotFound": 404,"ResourceInaccessible": 500,"Unauthorized": 403})
}


func filterInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["ovf_library_item_id"] = bindings.NewIdType([]string {"com.vmware.content.library.Item"}, "")
    fields["target"] = bindings.NewReferenceType(DeploymentTargetBindingType)
    fieldNameMap["ovf_library_item_id"] = "OvfLibraryItemId"
    fieldNameMap["target"] = "Target"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func filterOutputType() bindings.BindingType {
    return bindings.NewReferenceType(OvfSummaryBindingType)
}

func filterRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"InvalidArgument": 400,"NotFound": 404,"ResourceInaccessible": 500})
}


func createInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["client_token"] = bindings.NewOptionalType(bindings.NewStringType())
    fields["source"] = bindings.NewReferenceType(DeployableIdentityBindingType)
    fields["target"] = bindings.NewReferenceType(CreateTargetBindingType)
    fields["create_spec"] = bindings.NewReferenceType(CreateSpecBindingType)
    fieldNameMap["client_token"] = "ClientToken"
    fieldNameMap["source"] = "Source"
    fieldNameMap["target"] = "Target"
    fieldNameMap["create_spec"] = "CreateSpec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func createOutputType() bindings.BindingType {
    return bindings.NewReferenceType(CreateResultBindingType)
}

func createRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"InvalidArgument": 400,"NotFound": 404,"NotAllowedInCurrentState": 400,"ResourceInaccessible": 500,"ResourceBusy": 500})
}



func DeployableIdentityBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewStringType()
    fieldNameMap["type"] = "Type_"
    fields["id"] = bindings.NewIdType([]string {"null", "null"}, "type")
    fieldNameMap["id"] = "Id"
    var validators = []bindings.Validator{}
    isv1 := bindings.NewIsOneOfValidator(
        "type",
        []string {
             "VirtualMachine",
             "VirtualApp",
        },
    )
    validators = append(validators, isv1)
    return bindings.NewStructType("com.vmware.vcenter.ovf.library_item.deployable_identity",fields, reflect.TypeOf(DeployableIdentity{}), fieldNameMap, validators)
}

func ResourcePoolDeploymentSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["name"] = "Name"
    fields["annotation"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["annotation"] = "Annotation"
    fields["accept_all_EULA"] = bindings.NewBooleanType()
    fieldNameMap["accept_all_EULA"] = "AcceptAllEULA"
    fields["network_mappings"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewIdType([]string {"Network"}, ""),reflect.TypeOf(map[string]string{})))
    fieldNameMap["network_mappings"] = "NetworkMappings"
    fields["storage_mappings"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewReferenceType(StorageGroupMappingBindingType),reflect.TypeOf(map[string]StorageGroupMapping{})))
    fieldNameMap["storage_mappings"] = "StorageMappings"
    fields["storage_provisioning"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.ovf.disk_provisioning_type", reflect.TypeOf(ovf.DiskProvisioningType(ovf.DiskProvisioningType_thin))))
    fieldNameMap["storage_provisioning"] = "StorageProvisioning"
    fields["storage_profile_id"] = bindings.NewOptionalType(bindings.NewIdType([]string {"StorageProfile"}, ""))
    fieldNameMap["storage_profile_id"] = "StorageProfileId"
    fields["locale"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["locale"] = "Locale"
    fields["flags"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["flags"] = "Flags"
    fields["additional_parameters"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewDynamicStructType([]bindings.ReferenceType {bindings.NewReferenceType(ovf.OvfParamsBindingType),}, bindings.JSONRPC), reflect.TypeOf([]*data.StructValue{})))
    fieldNameMap["additional_parameters"] = "AdditionalParameters"
    fields["default_datastore_id"] = bindings.NewOptionalType(bindings.NewIdType([]string {"Datastore"}, ""))
    fieldNameMap["default_datastore_id"] = "DefaultDatastoreId"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.library_item.resource_pool_deployment_spec",fields, reflect.TypeOf(ResourcePoolDeploymentSpec{}), fieldNameMap, validators)
}

func StorageGroupMappingBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewEnumType("com.vmware.vcenter.ovf.library_item.storage_group_mapping.type", reflect.TypeOf(StorageGroupMapping_Type(StorageGroupMapping_Type_DATASTORE)))
    fieldNameMap["type"] = "Type_"
    fields["datastore_id"] = bindings.NewOptionalType(bindings.NewIdType([]string {"Datastore"}, ""))
    fieldNameMap["datastore_id"] = "DatastoreId"
    fields["storage_profile_id"] = bindings.NewOptionalType(bindings.NewIdType([]string {"StorageProfile"}, ""))
    fieldNameMap["storage_profile_id"] = "StorageProfileId"
    fields["provisioning"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.ovf.disk_provisioning_type", reflect.TypeOf(ovf.DiskProvisioningType(ovf.DiskProvisioningType_thin))))
    fieldNameMap["provisioning"] = "Provisioning"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("type",
        map[string][]bindings.FieldData {
            "DATASTORE": []bindings.FieldData {
                 bindings.NewFieldData("datastore_id", true),
            },
            "STORAGE_PROFILE": []bindings.FieldData {
                 bindings.NewFieldData("storage_profile_id", true),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.ovf.library_item.storage_group_mapping",fields, reflect.TypeOf(StorageGroupMapping{}), fieldNameMap, validators)
}

func ResultInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["errors"] = bindings.NewListType(bindings.NewReferenceType(ovf.OvfErrorBindingType), reflect.TypeOf([]ovf.OvfError{}))
    fieldNameMap["errors"] = "Errors"
    fields["warnings"] = bindings.NewListType(bindings.NewReferenceType(ovf.OvfWarningBindingType), reflect.TypeOf([]ovf.OvfWarning{}))
    fieldNameMap["warnings"] = "Warnings"
    fields["information"] = bindings.NewListType(bindings.NewReferenceType(ovf.OvfInfoBindingType), reflect.TypeOf([]ovf.OvfInfo{}))
    fieldNameMap["information"] = "Information"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.library_item.result_info",fields, reflect.TypeOf(ResultInfo{}), fieldNameMap, validators)
}

func DeploymentResultBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["succeeded"] = bindings.NewBooleanType()
    fieldNameMap["succeeded"] = "Succeeded"
    fields["resource_id"] = bindings.NewOptionalType(bindings.NewReferenceType(DeployableIdentityBindingType))
    fieldNameMap["resource_id"] = "ResourceId"
    fields["error"] = bindings.NewOptionalType(bindings.NewReferenceType(ResultInfoBindingType))
    fieldNameMap["error"] = "Error"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.library_item.deployment_result",fields, reflect.TypeOf(DeploymentResult{}), fieldNameMap, validators)
}

func DeploymentTargetBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["resource_pool_id"] = bindings.NewIdType([]string {"ResourcePool"}, "")
    fieldNameMap["resource_pool_id"] = "ResourcePoolId"
    fields["host_id"] = bindings.NewOptionalType(bindings.NewIdType([]string {"HostSystem"}, ""))
    fieldNameMap["host_id"] = "HostId"
    fields["folder_id"] = bindings.NewOptionalType(bindings.NewIdType([]string {"Folder"}, ""))
    fieldNameMap["folder_id"] = "FolderId"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.library_item.deployment_target",fields, reflect.TypeOf(DeploymentTarget{}), fieldNameMap, validators)
}

func OvfSummaryBindingType() bindings.BindingType {
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
    fields["additional_params"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewDynamicStructType([]bindings.ReferenceType {bindings.NewReferenceType(ovf.OvfParamsBindingType),}, bindings.JSONRPC), reflect.TypeOf([]*data.StructValue{})))
    fieldNameMap["additional_params"] = "AdditionalParams"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.library_item.ovf_summary",fields, reflect.TypeOf(OvfSummary{}), fieldNameMap, validators)
}

func CreateTargetBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["library_id"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.content.Library"}, ""))
    fieldNameMap["library_id"] = "LibraryId"
    fields["library_item_id"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.content.library.Item"}, ""))
    fieldNameMap["library_item_id"] = "LibraryItemId"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.library_item.create_target",fields, reflect.TypeOf(CreateTarget{}), fieldNameMap, validators)
}

func CreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["name"] = "Name"
    fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["description"] = "Description"
    fields["flags"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["flags"] = "Flags"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.library_item.create_spec",fields, reflect.TypeOf(CreateSpec{}), fieldNameMap, validators)
}

func CreateResultBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["succeeded"] = bindings.NewBooleanType()
    fieldNameMap["succeeded"] = "Succeeded"
    fields["ovf_library_item_id"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.content.library.Item"}, ""))
    fieldNameMap["ovf_library_item_id"] = "OvfLibraryItemId"
    fields["error"] = bindings.NewOptionalType(bindings.NewReferenceType(ResultInfoBindingType))
    fieldNameMap["error"] = "Error"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.library_item.create_result",fields, reflect.TypeOf(CreateResult{}), fieldNameMap, validators)
}


