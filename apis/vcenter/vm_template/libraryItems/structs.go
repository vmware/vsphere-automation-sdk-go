/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: LibraryItems.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package libraryItems

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vcenter/vm"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)




// The ``CreateSpec`` class defines the information required to create a library item containing a virtual machine template.
 type CreateSpec struct {
    SourceVm string
    Name string
    Description *string
    Library string
    VmHomeStorage *CreateSpecVmHomeStorage
    DiskStorage *CreateSpecDiskStorage
    DiskStorageOverrides map[string]CreateSpecDiskStorage
    Placement *CreatePlacementSpec
}






// The ``CreatePlacementSpec`` class contains information used to place a virtual machine template onto resources within the vCenter inventory.
 type CreatePlacementSpec struct {
    Folder *string
    ResourcePool *string
    Host *string
    Cluster *string
}






// The ``CreateSpecVmHomeStoragePolicy`` class defines the storage policy specification for a virtual machine template's configuration and log files.
 type CreateSpecVmHomeStoragePolicy struct {
    Type_ CreateSpecVmHomeStoragePolicy_Type
    Policy *string
}




    
    // Policy type for the virtual machine template's configuration and log files.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type CreateSpecVmHomeStoragePolicy_Type string

    const (
        // Use the specified policy.
         CreateSpecVmHomeStoragePolicy_Type_USE_SPECIFIED_POLICY CreateSpecVmHomeStoragePolicy_Type = "USE_SPECIFIED_POLICY"
    )

    func (t CreateSpecVmHomeStoragePolicy_Type) CreateSpecVmHomeStoragePolicy_Type() bool {
        switch t {
            case CreateSpecVmHomeStoragePolicy_Type_USE_SPECIFIED_POLICY:
                return true
            default:
                return false
        }
    }



// The ``CreateSpecVmHomeStorage`` class defines the storage specification for a virtual machine template's configuration and log files.
 type CreateSpecVmHomeStorage struct {
    Datastore *string
    StoragePolicy *CreateSpecVmHomeStoragePolicy
}






// The ``CreateSpecDiskStoragePolicy`` class defines the storage policy specification for a virtual machine template's disks.
 type CreateSpecDiskStoragePolicy struct {
    Type_ CreateSpecDiskStoragePolicy_Type
    Policy *string
}




    
    // Policy type for a virtual machine template's disk.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type CreateSpecDiskStoragePolicy_Type string

    const (
        // Use the specified policy.
         CreateSpecDiskStoragePolicy_Type_USE_SPECIFIED_POLICY CreateSpecDiskStoragePolicy_Type = "USE_SPECIFIED_POLICY"
    )

    func (t CreateSpecDiskStoragePolicy_Type) CreateSpecDiskStoragePolicy_Type() bool {
        switch t {
            case CreateSpecDiskStoragePolicy_Type_USE_SPECIFIED_POLICY:
                return true
            default:
                return false
        }
    }



// The ``CreateSpecDiskStorage`` class defines the storage specification for a virtual machine template's disks.
 type CreateSpecDiskStorage struct {
    Datastore *string
    StoragePolicy *CreateSpecDiskStoragePolicy
}






// The ``DeploySpec`` class defines the deployment parameters that can be specified for the ``deploy()`` method.
 type DeploySpec struct {
    Name string
    Description *string
    VmHomeStorage *DeploySpecVmHomeStorage
    DiskStorage *DeploySpecDiskStorage
    DiskStorageOverrides map[string]DeploySpecDiskStorage
    Placement *DeployPlacementSpec
    PoweredOn *bool
    GuestCustomization *GuestCustomizationSpec
    HardwareCustomization *HardwareCustomizationSpec
}






// The ``HardwareCustomizationSpec`` class defines the hardware customization options that are applied to the deployed virtual machine.
 type HardwareCustomizationSpec struct {
    Nics map[string]EthernetUpdateSpec
    DisksToRemove map[string]bool
    DisksToUpdate map[string]DiskUpdateSpec
    CpuUpdate *CpuUpdateSpec
    MemoryUpdate *MemoryUpdateSpec
}






// The ``DiskUpdateSpec`` class describes updates to the configuration of a virtual disk in the deployed virtual machine.
 type DiskUpdateSpec struct {
    Capacity int64
}






// The ``CpuUpdateSpec`` class describes updates to the CPU configuration of the deployed virtual machine.
 type CpuUpdateSpec struct {
    NumCpus *int64
    NumCoresPerSocket *int64
}






// The ``MemoryUpdateSpec`` class describes updates to the memory configuration of the deployed virtual machine.
 type MemoryUpdateSpec struct {
    Memory *int64
}






// The ``EthernetUpdateSpec`` class describes the network that the ethernet adapter of the deployed virtual machine should be connected to.
 type EthernetUpdateSpec struct {
    Network *string
}






// The ``DeployPlacementSpec`` class contains information used to place a virtual machine onto resources within the vCenter inventory.
 type DeployPlacementSpec struct {
    Folder *string
    ResourcePool *string
    Host *string
    Cluster *string
}






// The ``DeploySpecVmHomeStoragePolicy`` class defines the storage policy specification for the deployed virtual machine's configuration and log files.
 type DeploySpecVmHomeStoragePolicy struct {
    Type_ DeploySpecVmHomeStoragePolicy_Type
    Policy *string
}




    
    // Policy type for the deployed virtual machine's configuration and log files.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type DeploySpecVmHomeStoragePolicy_Type string

    const (
        // Use the specified policy.
         DeploySpecVmHomeStoragePolicy_Type_USE_SPECIFIED_POLICY DeploySpecVmHomeStoragePolicy_Type = "USE_SPECIFIED_POLICY"
        // Use the storage policy that is associated with the source virtual machine template's configuration and log files.
         DeploySpecVmHomeStoragePolicy_Type_USE_SOURCE_POLICY DeploySpecVmHomeStoragePolicy_Type = "USE_SOURCE_POLICY"
    )

    func (t DeploySpecVmHomeStoragePolicy_Type) DeploySpecVmHomeStoragePolicy_Type() bool {
        switch t {
            case DeploySpecVmHomeStoragePolicy_Type_USE_SPECIFIED_POLICY:
                return true
            case DeploySpecVmHomeStoragePolicy_Type_USE_SOURCE_POLICY:
                return true
            default:
                return false
        }
    }



// The ``DeploySpecVmHomeStorage`` class defines the storage specification for a deployed virtual machine's configuration and log files.
 type DeploySpecVmHomeStorage struct {
    Datastore *string
    StoragePolicy *DeploySpecVmHomeStoragePolicy
}






// The ``DeploySpecDiskStoragePolicy`` class describes the storage policy specification for the deployed virtual machine's disks.
 type DeploySpecDiskStoragePolicy struct {
    Type_ DeploySpecDiskStoragePolicy_Type
    Policy *string
}




    
    // Policy type for the deployed virtual machine's disk.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type DeploySpecDiskStoragePolicy_Type string

    const (
        // Use the specified policy.
         DeploySpecDiskStoragePolicy_Type_USE_SPECIFIED_POLICY DeploySpecDiskStoragePolicy_Type = "USE_SPECIFIED_POLICY"
        // Use the storage policy that is associated with the corresponding disk in the source virtual machine template.
         DeploySpecDiskStoragePolicy_Type_USE_SOURCE_POLICY DeploySpecDiskStoragePolicy_Type = "USE_SOURCE_POLICY"
    )

    func (t DeploySpecDiskStoragePolicy_Type) DeploySpecDiskStoragePolicy_Type() bool {
        switch t {
            case DeploySpecDiskStoragePolicy_Type_USE_SPECIFIED_POLICY:
                return true
            case DeploySpecDiskStoragePolicy_Type_USE_SOURCE_POLICY:
                return true
            default:
                return false
        }
    }



// The ``DeploySpecDiskStorage`` class contains the storage specification for disks in the virtual machine.
 type DeploySpecDiskStorage struct {
    Datastore *string
    StoragePolicy *DeploySpecDiskStoragePolicy
}






// The ``GuestCustomizationSpec`` class contains information required to customize the deployed virtual machine.
 type GuestCustomizationSpec struct {
    Name *string
}






// The ``Info`` class contains information about a virtual machine template item in content library.
 type Info struct {
    GuestOS vm.GuestOS
    Cpu CpuInfo
    Memory MemoryInfo
    VmHomeStorage VmHomeStorageInfo
    Disks map[string]DiskInfo
    Nics map[string]EthernetInfo
    VmTemplate string
}






// The ``CpuInfo`` class contains CPU related information about the virtual machine template.
 type CpuInfo struct {
    Count int64
    CoresPerSocket int64
}






// The ``MemoryInfo`` class contains memory related information about the virtual machine template.
 type MemoryInfo struct {
    SizeMiB int64
}






// The ``VmHomeStorageInfo`` class contains storage information about the virtual machine template's configuration and log files.
 type VmHomeStorageInfo struct {
    Datastore string
    StoragePolicy *string
}






// The ``DiskInfo`` class contains information about the virtual machine template's virtual disk.
 type DiskInfo struct {
    Capacity *int64
    DiskStorage DiskStorageInfo
}






// The ``DiskStorageInfo`` class contains storage related information about a virtual machine template's virtual disk.
 type DiskStorageInfo struct {
    Datastore string
    StoragePolicy *string
}






// The ``EthernetInfo`` class contains information about a virtual machine template's virtual Ethernet adapter.
 type EthernetInfo struct {
    BackingType EthernetInfo_NetworkBackingType
    MacType EthernetInfo_MacAddressType
    Network *string
}




    
    // The ``NetworkBackingType`` enumeration class defines valid network backing types for a virtual Ethernet adapter.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type EthernetInfo_NetworkBackingType string

    const (
        // vSphere standard portgroup network backing.
         EthernetInfo_NetworkBackingType_STANDARD_PORTGROUP EthernetInfo_NetworkBackingType = "STANDARD_PORTGROUP"
        // Legacy host device network backing. Imported VMs may have virtual Ethernet adapters with this type of backing, but this type of backing cannot be used to create or to update a virtual Ethernet adapter.
         EthernetInfo_NetworkBackingType_HOST_DEVICE EthernetInfo_NetworkBackingType = "HOST_DEVICE"
        // Distributed virtual switch backing.
         EthernetInfo_NetworkBackingType_DISTRIBUTED_PORTGROUP EthernetInfo_NetworkBackingType = "DISTRIBUTED_PORTGROUP"
        // Opaque network backing.
         EthernetInfo_NetworkBackingType_OPAQUE_NETWORK EthernetInfo_NetworkBackingType = "OPAQUE_NETWORK"
    )

    func (n EthernetInfo_NetworkBackingType) EthernetInfo_NetworkBackingType() bool {
        switch n {
            case EthernetInfo_NetworkBackingType_STANDARD_PORTGROUP:
                return true
            case EthernetInfo_NetworkBackingType_HOST_DEVICE:
                return true
            case EthernetInfo_NetworkBackingType_DISTRIBUTED_PORTGROUP:
                return true
            case EthernetInfo_NetworkBackingType_OPAQUE_NETWORK:
                return true
            default:
                return false
        }
    }

    
    // The ``MacAddressType`` enumeration class defines the valid MAC address origins for a virtual Ethernet adapter.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type EthernetInfo_MacAddressType string

    const (
        // MAC address is assigned statically.
         EthernetInfo_MacAddressType_MANUAL EthernetInfo_MacAddressType = "MANUAL"
        // MAC address is generated automatically.
         EthernetInfo_MacAddressType_GENERATED EthernetInfo_MacAddressType = "GENERATED"
        // MAC address is assigned by vCenter Server.
         EthernetInfo_MacAddressType_ASSIGNED EthernetInfo_MacAddressType = "ASSIGNED"
    )

    func (m EthernetInfo_MacAddressType) EthernetInfo_MacAddressType() bool {
        switch m {
            case EthernetInfo_MacAddressType_MANUAL:
                return true
            case EthernetInfo_MacAddressType_GENERATED:
                return true
            case EthernetInfo_MacAddressType_ASSIGNED:
                return true
            default:
                return false
        }
    }






func createInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(CreateSpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func createOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"com.vmware.content.library.Item"}, "")
}

func createRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
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
       map[string]int{"AlreadyExists": 400,"Error": 500,"InvalidArgument": 400,"NotAllowedInCurrentState": 400,"NotFound": 404,"ResourceInaccessible": 500,"ResourceInUse": 400,"ServiceUnavailable": 503,"UnableToAllocateResource": 500,"Unauthenticated": 401,"Unauthorized": 403})
}


func deployInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["template_library_item"] = bindings.NewIdType([]string {"com.vmware.content.library.Item"}, "")
    fields["spec"] = bindings.NewReferenceType(DeploySpecBindingType)
    fieldNameMap["template_library_item"] = "TemplateLibraryItem"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func deployOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"VirtualMachine"}, "")
}

func deployRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
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
       map[string]int{"AlreadyExists": 400,"Error": 500,"InvalidArgument": 400,"NotAllowedInCurrentState": 400,"NotFound": 404,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"UnableToAllocateResource": 500,"Unauthenticated": 401,"Unauthorized": 403})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["template_library_item"] = bindings.NewIdType([]string {"com.vmware.content.library.Item"}, "")
    fieldNameMap["template_library_item"] = "TemplateLibraryItem"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewOptionalType(bindings.NewReferenceType(InfoBindingType))
}

func getRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
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
       map[string]int{"Error": 500,"NotFound": 404,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}



func CreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["source_vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fieldNameMap["source_vm"] = "SourceVm"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["description"] = "Description"
    fields["library"] = bindings.NewIdType([]string {"com.vmware.content.Library"}, "")
    fieldNameMap["library"] = "Library"
    fields["vm_home_storage"] = bindings.NewOptionalType(bindings.NewReferenceType(CreateSpecVmHomeStorageBindingType))
    fieldNameMap["vm_home_storage"] = "VmHomeStorage"
    fields["disk_storage"] = bindings.NewOptionalType(bindings.NewReferenceType(CreateSpecDiskStorageBindingType))
    fieldNameMap["disk_storage"] = "DiskStorage"
    fields["disk_storage_overrides"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Disk"}, ""), bindings.NewReferenceType(CreateSpecDiskStorageBindingType),reflect.TypeOf(map[string]CreateSpecDiskStorage{})))
    fieldNameMap["disk_storage_overrides"] = "DiskStorageOverrides"
    fields["placement"] = bindings.NewOptionalType(bindings.NewReferenceType(CreatePlacementSpecBindingType))
    fieldNameMap["placement"] = "Placement"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.create_spec",fields, reflect.TypeOf(CreateSpec{}), fieldNameMap, validators)
}

func CreatePlacementSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["folder"] = bindings.NewOptionalType(bindings.NewIdType([]string {"Folder"}, ""))
    fieldNameMap["folder"] = "Folder"
    fields["resource_pool"] = bindings.NewOptionalType(bindings.NewIdType([]string {"ResourcePool"}, ""))
    fieldNameMap["resource_pool"] = "ResourcePool"
    fields["host"] = bindings.NewOptionalType(bindings.NewIdType([]string {"HostSystem"}, ""))
    fieldNameMap["host"] = "Host"
    fields["cluster"] = bindings.NewOptionalType(bindings.NewIdType([]string {"ClusterComputeResource"}, ""))
    fieldNameMap["cluster"] = "Cluster"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.create_placement_spec",fields, reflect.TypeOf(CreatePlacementSpec{}), fieldNameMap, validators)
}

func CreateSpecVmHomeStoragePolicyBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm_template.library_items.create_spec_vm_home_storage_policy.type", reflect.TypeOf(CreateSpecVmHomeStoragePolicy_Type(CreateSpecVmHomeStoragePolicy_Type_USE_SPECIFIED_POLICY)))
    fieldNameMap["type"] = "Type_"
    fields["policy"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.spbm.StorageProfile"}, ""))
    fieldNameMap["policy"] = "Policy"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("type",
        map[string][]bindings.FieldData {
            "USE_SPECIFIED_POLICY": []bindings.FieldData {
                 bindings.NewFieldData("policy", true),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.create_spec_vm_home_storage_policy",fields, reflect.TypeOf(CreateSpecVmHomeStoragePolicy{}), fieldNameMap, validators)
}

func CreateSpecVmHomeStorageBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["datastore"] = bindings.NewOptionalType(bindings.NewIdType([]string {"Datastore"}, ""))
    fieldNameMap["datastore"] = "Datastore"
    fields["storage_policy"] = bindings.NewOptionalType(bindings.NewReferenceType(CreateSpecVmHomeStoragePolicyBindingType))
    fieldNameMap["storage_policy"] = "StoragePolicy"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.create_spec_vm_home_storage",fields, reflect.TypeOf(CreateSpecVmHomeStorage{}), fieldNameMap, validators)
}

func CreateSpecDiskStoragePolicyBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm_template.library_items.create_spec_disk_storage_policy.type", reflect.TypeOf(CreateSpecDiskStoragePolicy_Type(CreateSpecDiskStoragePolicy_Type_USE_SPECIFIED_POLICY)))
    fieldNameMap["type"] = "Type_"
    fields["policy"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.spbm.StorageProfile"}, ""))
    fieldNameMap["policy"] = "Policy"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("type",
        map[string][]bindings.FieldData {
            "USE_SPECIFIED_POLICY": []bindings.FieldData {
                 bindings.NewFieldData("policy", true),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.create_spec_disk_storage_policy",fields, reflect.TypeOf(CreateSpecDiskStoragePolicy{}), fieldNameMap, validators)
}

func CreateSpecDiskStorageBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["datastore"] = bindings.NewOptionalType(bindings.NewIdType([]string {"Datastore"}, ""))
    fieldNameMap["datastore"] = "Datastore"
    fields["storage_policy"] = bindings.NewOptionalType(bindings.NewReferenceType(CreateSpecDiskStoragePolicyBindingType))
    fieldNameMap["storage_policy"] = "StoragePolicy"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.create_spec_disk_storage",fields, reflect.TypeOf(CreateSpecDiskStorage{}), fieldNameMap, validators)
}

func DeploySpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["description"] = "Description"
    fields["vm_home_storage"] = bindings.NewOptionalType(bindings.NewReferenceType(DeploySpecVmHomeStorageBindingType))
    fieldNameMap["vm_home_storage"] = "VmHomeStorage"
    fields["disk_storage"] = bindings.NewOptionalType(bindings.NewReferenceType(DeploySpecDiskStorageBindingType))
    fieldNameMap["disk_storage"] = "DiskStorage"
    fields["disk_storage_overrides"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Disk"}, ""), bindings.NewReferenceType(DeploySpecDiskStorageBindingType),reflect.TypeOf(map[string]DeploySpecDiskStorage{})))
    fieldNameMap["disk_storage_overrides"] = "DiskStorageOverrides"
    fields["placement"] = bindings.NewOptionalType(bindings.NewReferenceType(DeployPlacementSpecBindingType))
    fieldNameMap["placement"] = "Placement"
    fields["powered_on"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["powered_on"] = "PoweredOn"
    fields["guest_customization"] = bindings.NewOptionalType(bindings.NewReferenceType(GuestCustomizationSpecBindingType))
    fieldNameMap["guest_customization"] = "GuestCustomization"
    fields["hardware_customization"] = bindings.NewOptionalType(bindings.NewReferenceType(HardwareCustomizationSpecBindingType))
    fieldNameMap["hardware_customization"] = "HardwareCustomization"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.deploy_spec",fields, reflect.TypeOf(DeploySpec{}), fieldNameMap, validators)
}

func HardwareCustomizationSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["nics"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Ethernet"}, ""), bindings.NewReferenceType(EthernetUpdateSpecBindingType),reflect.TypeOf(map[string]EthernetUpdateSpec{})))
    fieldNameMap["nics"] = "Nics"
    fields["disks_to_remove"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Disk"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["disks_to_remove"] = "DisksToRemove"
    fields["disks_to_update"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Disk"}, ""), bindings.NewReferenceType(DiskUpdateSpecBindingType),reflect.TypeOf(map[string]DiskUpdateSpec{})))
    fieldNameMap["disks_to_update"] = "DisksToUpdate"
    fields["cpu_update"] = bindings.NewOptionalType(bindings.NewReferenceType(CpuUpdateSpecBindingType))
    fieldNameMap["cpu_update"] = "CpuUpdate"
    fields["memory_update"] = bindings.NewOptionalType(bindings.NewReferenceType(MemoryUpdateSpecBindingType))
    fieldNameMap["memory_update"] = "MemoryUpdate"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.hardware_customization_spec",fields, reflect.TypeOf(HardwareCustomizationSpec{}), fieldNameMap, validators)
}

func DiskUpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["capacity"] = bindings.NewIntegerType()
    fieldNameMap["capacity"] = "Capacity"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.disk_update_spec",fields, reflect.TypeOf(DiskUpdateSpec{}), fieldNameMap, validators)
}

func CpuUpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["num_cpus"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["num_cpus"] = "NumCpus"
    fields["num_cores_per_socket"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["num_cores_per_socket"] = "NumCoresPerSocket"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.cpu_update_spec",fields, reflect.TypeOf(CpuUpdateSpec{}), fieldNameMap, validators)
}

func MemoryUpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["memory"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["memory"] = "Memory"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.memory_update_spec",fields, reflect.TypeOf(MemoryUpdateSpec{}), fieldNameMap, validators)
}

func EthernetUpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["network"] = bindings.NewOptionalType(bindings.NewIdType([]string {"Network"}, ""))
    fieldNameMap["network"] = "Network"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.ethernet_update_spec",fields, reflect.TypeOf(EthernetUpdateSpec{}), fieldNameMap, validators)
}

func DeployPlacementSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["folder"] = bindings.NewOptionalType(bindings.NewIdType([]string {"Folder"}, ""))
    fieldNameMap["folder"] = "Folder"
    fields["resource_pool"] = bindings.NewOptionalType(bindings.NewIdType([]string {"ResourcePool"}, ""))
    fieldNameMap["resource_pool"] = "ResourcePool"
    fields["host"] = bindings.NewOptionalType(bindings.NewIdType([]string {"HostSystem"}, ""))
    fieldNameMap["host"] = "Host"
    fields["cluster"] = bindings.NewOptionalType(bindings.NewIdType([]string {"ClusterComputeResource"}, ""))
    fieldNameMap["cluster"] = "Cluster"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.deploy_placement_spec",fields, reflect.TypeOf(DeployPlacementSpec{}), fieldNameMap, validators)
}

func DeploySpecVmHomeStoragePolicyBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm_template.library_items.deploy_spec_vm_home_storage_policy.type", reflect.TypeOf(DeploySpecVmHomeStoragePolicy_Type(DeploySpecVmHomeStoragePolicy_Type_USE_SPECIFIED_POLICY)))
    fieldNameMap["type"] = "Type_"
    fields["policy"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.spbm.StorageProfile"}, ""))
    fieldNameMap["policy"] = "Policy"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("type",
        map[string][]bindings.FieldData {
            "USE_SPECIFIED_POLICY": []bindings.FieldData {
                 bindings.NewFieldData("policy", true),
            },
            "USE_SOURCE_POLICY": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.deploy_spec_vm_home_storage_policy",fields, reflect.TypeOf(DeploySpecVmHomeStoragePolicy{}), fieldNameMap, validators)
}

func DeploySpecVmHomeStorageBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["datastore"] = bindings.NewOptionalType(bindings.NewIdType([]string {"Datastore"}, ""))
    fieldNameMap["datastore"] = "Datastore"
    fields["storage_policy"] = bindings.NewOptionalType(bindings.NewReferenceType(DeploySpecVmHomeStoragePolicyBindingType))
    fieldNameMap["storage_policy"] = "StoragePolicy"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.deploy_spec_vm_home_storage",fields, reflect.TypeOf(DeploySpecVmHomeStorage{}), fieldNameMap, validators)
}

func DeploySpecDiskStoragePolicyBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm_template.library_items.deploy_spec_disk_storage_policy.type", reflect.TypeOf(DeploySpecDiskStoragePolicy_Type(DeploySpecDiskStoragePolicy_Type_USE_SPECIFIED_POLICY)))
    fieldNameMap["type"] = "Type_"
    fields["policy"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.spbm.StorageProfile"}, ""))
    fieldNameMap["policy"] = "Policy"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("type",
        map[string][]bindings.FieldData {
            "USE_SPECIFIED_POLICY": []bindings.FieldData {
                 bindings.NewFieldData("policy", true),
            },
            "USE_SOURCE_POLICY": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.deploy_spec_disk_storage_policy",fields, reflect.TypeOf(DeploySpecDiskStoragePolicy{}), fieldNameMap, validators)
}

func DeploySpecDiskStorageBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["datastore"] = bindings.NewOptionalType(bindings.NewIdType([]string {"Datastore"}, ""))
    fieldNameMap["datastore"] = "Datastore"
    fields["storage_policy"] = bindings.NewOptionalType(bindings.NewReferenceType(DeploySpecDiskStoragePolicyBindingType))
    fieldNameMap["storage_policy"] = "StoragePolicy"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.deploy_spec_disk_storage",fields, reflect.TypeOf(DeploySpecDiskStorage{}), fieldNameMap, validators)
}

func GuestCustomizationSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["name"] = "Name"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.guest_customization_spec",fields, reflect.TypeOf(GuestCustomizationSpec{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["guest_OS"] = bindings.NewEnumType("com.vmware.vcenter.vm.guest_OS", reflect.TypeOf(vm.GuestOS(vm.GuestOS_DOS)))
    fieldNameMap["guest_OS"] = "GuestOS"
    fields["cpu"] = bindings.NewReferenceType(CpuInfoBindingType)
    fieldNameMap["cpu"] = "Cpu"
    fields["memory"] = bindings.NewReferenceType(MemoryInfoBindingType)
    fieldNameMap["memory"] = "Memory"
    fields["vm_home_storage"] = bindings.NewReferenceType(VmHomeStorageInfoBindingType)
    fieldNameMap["vm_home_storage"] = "VmHomeStorage"
    fields["disks"] = bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Disk"}, ""), bindings.NewReferenceType(DiskInfoBindingType),reflect.TypeOf(map[string]DiskInfo{}))
    fieldNameMap["disks"] = "Disks"
    fields["nics"] = bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Ethernet"}, ""), bindings.NewReferenceType(EthernetInfoBindingType),reflect.TypeOf(map[string]EthernetInfo{}))
    fieldNameMap["nics"] = "Nics"
    fields["vm_template"] = bindings.NewStringType()
    fieldNameMap["vm_template"] = "VmTemplate"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func CpuInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["count"] = bindings.NewIntegerType()
    fieldNameMap["count"] = "Count"
    fields["cores_per_socket"] = bindings.NewIntegerType()
    fieldNameMap["cores_per_socket"] = "CoresPerSocket"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.cpu_info",fields, reflect.TypeOf(CpuInfo{}), fieldNameMap, validators)
}

func MemoryInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["size_MiB"] = bindings.NewIntegerType()
    fieldNameMap["size_MiB"] = "SizeMiB"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.memory_info",fields, reflect.TypeOf(MemoryInfo{}), fieldNameMap, validators)
}

func VmHomeStorageInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["datastore"] = bindings.NewIdType([]string {"Datastore"}, "")
    fieldNameMap["datastore"] = "Datastore"
    fields["storage_policy"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.spbm.StorageProfile"}, ""))
    fieldNameMap["storage_policy"] = "StoragePolicy"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.vm_home_storage_info",fields, reflect.TypeOf(VmHomeStorageInfo{}), fieldNameMap, validators)
}

func DiskInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["capacity"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["capacity"] = "Capacity"
    fields["disk_storage"] = bindings.NewReferenceType(DiskStorageInfoBindingType)
    fieldNameMap["disk_storage"] = "DiskStorage"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.disk_info",fields, reflect.TypeOf(DiskInfo{}), fieldNameMap, validators)
}

func DiskStorageInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["datastore"] = bindings.NewIdType([]string {"Datastore"}, "")
    fieldNameMap["datastore"] = "Datastore"
    fields["storage_policy"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.spbm.StorageProfile"}, ""))
    fieldNameMap["storage_policy"] = "StoragePolicy"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.disk_storage_info",fields, reflect.TypeOf(DiskStorageInfo{}), fieldNameMap, validators)
}

func EthernetInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["backing_type"] = bindings.NewEnumType("com.vmware.vcenter.vm_template.library_items.ethernet_info.network_backing_type", reflect.TypeOf(EthernetInfo_NetworkBackingType(EthernetInfo_NetworkBackingType_STANDARD_PORTGROUP)))
    fieldNameMap["backing_type"] = "BackingType"
    fields["mac_type"] = bindings.NewEnumType("com.vmware.vcenter.vm_template.library_items.ethernet_info.mac_address_type", reflect.TypeOf(EthernetInfo_MacAddressType(EthernetInfo_MacAddressType_MANUAL)))
    fieldNameMap["mac_type"] = "MacType"
    fields["network"] = bindings.NewOptionalType(bindings.NewIdType([]string {"Network"}, ""))
    fieldNameMap["network"] = "Network"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.ethernet_info",fields, reflect.TypeOf(EthernetInfo{}), fieldNameMap, validators)
}


