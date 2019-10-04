/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Validation.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package validation

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vcenter/lcm"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)




 type ApplianceNameRequest struct {
    DestinationLocation ApplianceNameDestinationLocation
    DestinationAppliance ApplianceNameDestinationAppliance
}






// Data container for appliance name information used in validation of appliance name request. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ApplianceNameDestinationAppliance struct {
    ApplianceName string
}






 type ApplianceNameDestinationLocation struct {
    Esx *ApplianceNameEsx
    Vcenter *ApplianceNameVc
}






// This section describes the ESX host on which to deploy the appliance. Required if you are deploying the appliance directly on an ESX host. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ApplianceNameEsx struct {
    Connection lcm.Connection
}






// The configuration of ESX inventory. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ApplianceNameEsxInventory struct {
    ResourcePoolPath *string
}






// This subsection describes the vCenter on which to deploy the appliance. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ApplianceNameVc struct {
    Connection lcm.Connection
    Inventory ApplianceNameVcInventory
}






// All names are case-sensitive. you can install the appliance to one of the following destinations: 1. A resource pool in a cluster, use 'cluster_path'. 2. A specific ESX host in a cluster, use 'host_path'. 3. A resource pool in a specific ESX host being managed by the current vCenter, use 'resource_pool_path'. You must always provide the 'network_name' key. To install a new appliance to a specific ESX host in a cluster, provide the 'host_path' key, and the 'datastore_name', e.g. 'host_path': '/MyDataCenter/host/MyCluster/10.20.30.40', 'datastore_name': 'Your Datastore'. To install a new appliance to a specific resource pool, provide the 'resource_pool_path', and the 'datastore_name', e.g. 'resource_pool_path': '/Your Datacenter Folder/Your Datacenter/host/Your Cluster/Resources/Your Resource Pool', 'datastore_name': 'Your Datastore'. To place a new appliance to a virtual machine Folder, provide the 'vm_folder_path', e.g. 'vm_folder_path': 'VM Folder 0/VM Folder1'. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ApplianceNameVcInventory struct {
    VmFolderPath *string
    ResourcePoolPath *string
    ClusterPath *string
    HostPath *string
}









func checkApplianceNameInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(ApplianceNameRequestBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func checkApplianceNameOutputType() bindings.BindingType {
    return bindings.NewBooleanType()
}

func checkApplianceNameRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{})
}



func ApplianceNameRequestBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["destination_location"] = bindings.NewReferenceType(ApplianceNameDestinationLocationBindingType)
    fieldNameMap["destination_location"] = "DestinationLocation"
    fields["destination_appliance"] = bindings.NewReferenceType(ApplianceNameDestinationApplianceBindingType)
    fieldNameMap["destination_appliance"] = "DestinationAppliance"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.validation.appliance_name_request",fields, reflect.TypeOf(ApplianceNameRequest{}), fieldNameMap, validators)
}

func ApplianceNameDestinationApplianceBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["appliance_name"] = bindings.NewStringType()
    fieldNameMap["appliance_name"] = "ApplianceName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.validation.appliance_name_destination_appliance",fields, reflect.TypeOf(ApplianceNameDestinationAppliance{}), fieldNameMap, validators)
}

func ApplianceNameDestinationLocationBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["esx"] = bindings.NewOptionalType(bindings.NewReferenceType(ApplianceNameEsxBindingType))
    fieldNameMap["esx"] = "Esx"
    fields["vcenter"] = bindings.NewOptionalType(bindings.NewReferenceType(ApplianceNameVcBindingType))
    fieldNameMap["vcenter"] = "Vcenter"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.validation.appliance_name_destination_location",fields, reflect.TypeOf(ApplianceNameDestinationLocation{}), fieldNameMap, validators)
}

func ApplianceNameEsxBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["connection"] = bindings.NewReferenceType(lcm.ConnectionBindingType)
    fieldNameMap["connection"] = "Connection"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.validation.appliance_name_esx",fields, reflect.TypeOf(ApplianceNameEsx{}), fieldNameMap, validators)
}

func ApplianceNameEsxInventoryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["resource_pool_path"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["resource_pool_path"] = "ResourcePoolPath"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.validation.appliance_name_esx_inventory",fields, reflect.TypeOf(ApplianceNameEsxInventory{}), fieldNameMap, validators)
}

func ApplianceNameVcBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["connection"] = bindings.NewReferenceType(lcm.ConnectionBindingType)
    fieldNameMap["connection"] = "Connection"
    fields["inventory"] = bindings.NewReferenceType(ApplianceNameVcInventoryBindingType)
    fieldNameMap["inventory"] = "Inventory"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.validation.appliance_name_vc",fields, reflect.TypeOf(ApplianceNameVc{}), fieldNameMap, validators)
}

func ApplianceNameVcInventoryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm_folder_path"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["vm_folder_path"] = "VmFolderPath"
    fields["resource_pool_path"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["resource_pool_path"] = "ResourcePoolPath"
    fields["cluster_path"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["cluster_path"] = "ClusterPath"
    fields["host_path"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["host_path"] = "HostPath"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.validation.appliance_name_vc_inventory",fields, reflect.TypeOf(ApplianceNameVcInventory{}), fieldNameMap, validators)
}


