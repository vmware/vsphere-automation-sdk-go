/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: VM.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package vm

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vcenter/vm/hardware"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vcenter/vm/hardware/adapter/nvme"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vcenter/vm/hardware/adapter/sata"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vcenter/vm/hardware/adapter/scsi"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vcenter/vm/hardware/boot"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vcenter/vm/hardware/boot/device"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vcenter/vm/hardware/cdrom"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vcenter/vm/hardware/cpu"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vcenter/vm/hardware/disk"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vcenter/vm/hardware/ethernet"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vcenter/vm/hardware/floppy"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vcenter/vm/hardware/memory"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vcenter/vm/hardware/parallel"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vcenter/vm/hardware/serial"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vcenter/vm/identity"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vcenter/vm/power"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// Resource type for virtual machines.
const VM_RESOURCE_TYPE = "VirtualMachine"



// The ``InventoryPlacementSpec`` class contains information used to place a virtual machine in the vCenter inventory.
 type InventoryPlacementSpec struct {
    Folder *string
}






// The ``ComputePlacementSpec`` class contains information used to place a virtual machine on compute resources.
 type ComputePlacementSpec struct {
    ResourcePool *string
    Host *string
    Cluster *string
}






// The ``StoragePlacementSpec`` class contains information used to store a virtual machine's files.
 type StoragePlacementSpec struct {
    Datastore *string
}






// The ``PlacementSpec`` class contains information used to place a virtual machine onto resources within the vCenter inventory.
 type PlacementSpec struct {
    Folder *string
    ResourcePool *string
    Host *string
    Cluster *string
    Datastore *string
}






// The ``StoragePolicySpec`` class contains information about the storage policy to be associated with a virtual machine object.
 type StoragePolicySpec struct {
    Policy string
}






// Document-based creation spec.
 type CreateSpec struct {
    GuestOS GuestOS
    Name *string
    Placement *PlacementSpec
    HardwareVersion *hardware.Version
    Boot *boot.CreateSpec
    BootDevices []device.EntryCreateSpec
    Cpu *cpu.UpdateSpec
    Memory *memory.UpdateSpec
    Disks []disk.CreateSpec
    Nics []ethernet.CreateSpec
    Cdroms []cdrom.CreateSpec
    Floppies []floppy.CreateSpec
    ParallelPorts []parallel.CreateSpec
    SerialPorts []serial.CreateSpec
    SataAdapters []sata.CreateSpec
    ScsiAdapters []scsi.CreateSpec
    NvmeAdapters []nvme.CreateSpec
    StoragePolicy *StoragePolicySpec
}






// Document-based info.
 type Info struct {
    GuestOS GuestOS
    Name string
    Identity *identity.Info
    PowerState power.State
    InstantCloneFrozen *bool
    Hardware hardware.Info
    Boot boot.Info
    BootDevices []device.Entry
    Cpu cpu.Info
    Memory memory.Info
    Disks map[string]disk.Info
    Nics map[string]ethernet.Info
    Cdroms map[string]cdrom.Info
    Floppies map[string]floppy.Info
    ParallelPorts map[string]parallel.Info
    SerialPorts map[string]serial.Info
    SataAdapters map[string]sata.Info
    ScsiAdapters map[string]scsi.Info
    NvmeAdapters map[string]nvme.Info
}






// The ``GuestCustomizationSpec`` class contains information required to customize a virtual machine when deploying it. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type GuestCustomizationSpec struct {
    Name *string
}






// Document-based disk clone spec. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type DiskCloneSpec struct {
    Datastore *string
    StoragePolicy *disk.StoragePolicySpec
}






// The ``ClonePlacementSpec`` class contains information used to place a clone of a virtual machine onto resources within the vCenter inventory. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ClonePlacementSpec struct {
    Folder *string
    ResourcePool *string
    Host *string
    Cluster *string
    Datastore *string
}






// Document-based clone spec. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type CloneSpec struct {
    Source string
    Name string
    Placement *ClonePlacementSpec
    DisksToRemove map[string]bool
    DisksToUpdate map[string]DiskCloneSpec
    VmHomeStoragePolicy *StoragePolicySpec
    DiskStoragePolicy *StoragePolicySpec
    FallbackToDatastoreDefaultPolicy *bool
    PowerOn *bool
    GuestCustomizationSpec *GuestCustomizationSpec
}






// Document-based disk relocate spec. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type DiskRelocateSpec struct {
    Datastore *string
    StoragePolicy *disk.StoragePolicySpec
}






// The ``RelocatePlacementSpec`` class contains information used to change the placement of an existing virtual machine within the vCenter inventory. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type RelocatePlacementSpec struct {
    Folder *string
    ResourcePool *string
    Host *string
    Cluster *string
    Datastore *string
}






// Document-based relocate spec. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type RelocateSpec struct {
    Placement *RelocatePlacementSpec
    Disks map[string]DiskRelocateSpec
    VmHomeStoragePolicy *StoragePolicySpec
    DiskStoragePolicy *StoragePolicySpec
    FallbackToDatastoreDefaultPolicy *bool
}






// The ``InstantClonePlacementSpec`` class contains information used to place an InstantClone of a virtual machine onto resources within the vCenter inventory.
 type InstantClonePlacementSpec struct {
    Folder *string
    ResourcePool *string
    Datastore *string
}






// Document-based InstantClone spec.
 type InstantCloneSpec struct {
    Source string
    Name string
    Placement *InstantClonePlacementSpec
    NicsToUpdate map[string]ethernet.UpdateSpec
    DisconnectAllNics *bool
    ParallelPortsToUpdate map[string]parallel.UpdateSpec
    SerialPortsToUpdate map[string]serial.UpdateSpec
    BiosUuid *string
}






// The ``FilterSpec`` class contains properties used to filter the results when listing virtual machines (see VM#list). If multiple properties are specified, only virtual machines matching all of the properties match the filter.
 type FilterSpec struct {
    Vms map[string]bool
    Names map[string]bool
    Folders map[string]bool
    Datacenters map[string]bool
    Hosts map[string]bool
    Clusters map[string]bool
    ResourcePools map[string]bool
    PowerStates map[power.State]bool
}






// The ``Summary`` class contains commonly used information about a virtual machine.
 type Summary struct {
    Vm string
    Name string
    PowerState power.State
    CpuCount *int64
    MemorySizeMiB *int64
}






// The ``RegisterPlacementSpec`` class contains information used to place a virtual machine, created from existing virtual machine files on storage, onto resources within the vCenter inventory.
 type RegisterPlacementSpec struct {
    Folder *string
    ResourcePool *string
    Host *string
    Cluster *string
}






// The ``RegisterSpec`` class contains information used to create a virtual machine from existing virtual machine files on storage. 
//
//  The location of the virtual machine files on storage must be specified by providing either RegisterSpec#datastore and RegisterSpec#path or by providing RegisterSpec#datastorePath. If RegisterSpec#datastore and RegisterSpec#path are map with bool value, RegisterSpec#datastorePath must be null, and if RegisterSpec#datastorePath is map with bool value, RegisterSpec#datastore and RegisterSpec#path must be null.
 type RegisterSpec struct {
    Datastore *string
    Path *string
    DatastorePath *string
    Name *string
    Placement *RegisterPlacementSpec
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
    return bindings.NewIdType([]string {"VirtualMachine"}, "")
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
       map[string]int{"AlreadyExists": 400,"Error": 500,"InvalidArgument": 400,"NotFound": 404,"ResourceInaccessible": 500,"ResourceInUse": 400,"ServiceUnavailable": 503,"UnableToAllocateResource": 500,"Unauthenticated": 401,"Unauthorized": 403,"Unsupported": 400})
}


func cloneInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(CloneSpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cloneOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"VirtualMachine"}, "")
}

func cloneRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"AlreadyExists": 400,"Error": 500,"InvalidArgument": 400,"NotFound": 404,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"UnableToAllocateResource": 500,"Unauthenticated": 401,"Unauthorized": 403})
}


func relocateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fields["spec"] = bindings.NewReferenceType(RelocateSpecBindingType)
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func relocateOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func relocateRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func instantCloneInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(InstantCloneSpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func instantCloneOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"VirtualMachine"}, "")
}

func instantCloneRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"AlreadyExists": 400,"Error": 500,"InvalidArgument": 400,"NotFound": 404,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"UnableToAllocateResource": 500,"Unauthenticated": 401,"Unauthorized": 403})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fieldNameMap["vm"] = "Vm"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewReferenceType(InfoBindingType)
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


func deleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fieldNameMap["vm"] = "Vm"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func deleteOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func deleteRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"NotAllowedInCurrentState": 400,"NotFound": 404,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(FilterSpecBindingType))
    fieldNameMap["filter"] = "Filter"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(SummaryBindingType), reflect.TypeOf([]Summary{}))
}

func listRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"InvalidArgument": 400,"UnableToAllocateResource": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func registerInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(RegisterSpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func registerOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"VirtualMachine"}, "")
}

func registerRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"AlreadyExists": 400,"Error": 500,"InvalidArgument": 400,"NotFound": 404,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"UnableToAllocateResource": 500,"Unauthenticated": 401,"Unauthorized": 403})
}


func unregisterInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fieldNameMap["vm"] = "Vm"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func unregisterOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func unregisterRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotAllowedInCurrentState": 400,"NotFound": 404,"ResourceBusy": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}



func InventoryPlacementSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["folder"] = bindings.NewOptionalType(bindings.NewIdType([]string {"Folder"}, ""))
    fieldNameMap["folder"] = "Folder"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.VM.inventory_placement_spec",fields, reflect.TypeOf(InventoryPlacementSpec{}), fieldNameMap, validators)
}

func ComputePlacementSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["resource_pool"] = bindings.NewOptionalType(bindings.NewIdType([]string {"ResourcePool"}, ""))
    fieldNameMap["resource_pool"] = "ResourcePool"
    fields["host"] = bindings.NewOptionalType(bindings.NewIdType([]string {"HostSystem"}, ""))
    fieldNameMap["host"] = "Host"
    fields["cluster"] = bindings.NewOptionalType(bindings.NewIdType([]string {"ClusterComputeResource"}, ""))
    fieldNameMap["cluster"] = "Cluster"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.VM.compute_placement_spec",fields, reflect.TypeOf(ComputePlacementSpec{}), fieldNameMap, validators)
}

func StoragePlacementSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["datastore"] = bindings.NewOptionalType(bindings.NewIdType([]string {"Datastore"}, ""))
    fieldNameMap["datastore"] = "Datastore"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.VM.storage_placement_spec",fields, reflect.TypeOf(StoragePlacementSpec{}), fieldNameMap, validators)
}

func PlacementSpecBindingType() bindings.BindingType {
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
    fields["datastore"] = bindings.NewOptionalType(bindings.NewIdType([]string {"Datastore"}, ""))
    fieldNameMap["datastore"] = "Datastore"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.VM.placement_spec",fields, reflect.TypeOf(PlacementSpec{}), fieldNameMap, validators)
}

func StoragePolicySpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["policy"] = bindings.NewIdType([]string {"com.vmware.vcenter.StoragePolicy"}, "")
    fieldNameMap["policy"] = "Policy"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.VM.storage_policy_spec",fields, reflect.TypeOf(StoragePolicySpec{}), fieldNameMap, validators)
}

func CreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["guest_OS"] = bindings.NewEnumType("com.vmware.vcenter.vm.guest_OS", reflect.TypeOf(GuestOS(GuestOS_DOS)))
    fieldNameMap["guest_OS"] = "GuestOS"
    fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["name"] = "Name"
    fields["placement"] = bindings.NewOptionalType(bindings.NewReferenceType(PlacementSpecBindingType))
    fieldNameMap["placement"] = "Placement"
    fields["hardware_version"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.hardware.version", reflect.TypeOf(hardware.Version(hardware.Version_VMX_03))))
    fieldNameMap["hardware_version"] = "HardwareVersion"
    fields["boot"] = bindings.NewOptionalType(bindings.NewReferenceType(boot.CreateSpecBindingType))
    fieldNameMap["boot"] = "Boot"
    fields["boot_devices"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(device.EntryCreateSpecBindingType), reflect.TypeOf([]device.EntryCreateSpec{})))
    fieldNameMap["boot_devices"] = "BootDevices"
    fields["cpu"] = bindings.NewOptionalType(bindings.NewReferenceType(cpu.UpdateSpecBindingType))
    fieldNameMap["cpu"] = "Cpu"
    fields["memory"] = bindings.NewOptionalType(bindings.NewReferenceType(memory.UpdateSpecBindingType))
    fieldNameMap["memory"] = "Memory"
    fields["disks"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(disk.CreateSpecBindingType), reflect.TypeOf([]disk.CreateSpec{})))
    fieldNameMap["disks"] = "Disks"
    fields["nics"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ethernet.CreateSpecBindingType), reflect.TypeOf([]ethernet.CreateSpec{})))
    fieldNameMap["nics"] = "Nics"
    fields["cdroms"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(cdrom.CreateSpecBindingType), reflect.TypeOf([]cdrom.CreateSpec{})))
    fieldNameMap["cdroms"] = "Cdroms"
    fields["floppies"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(floppy.CreateSpecBindingType), reflect.TypeOf([]floppy.CreateSpec{})))
    fieldNameMap["floppies"] = "Floppies"
    fields["parallel_ports"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(parallel.CreateSpecBindingType), reflect.TypeOf([]parallel.CreateSpec{})))
    fieldNameMap["parallel_ports"] = "ParallelPorts"
    fields["serial_ports"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(serial.CreateSpecBindingType), reflect.TypeOf([]serial.CreateSpec{})))
    fieldNameMap["serial_ports"] = "SerialPorts"
    fields["sata_adapters"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(sata.CreateSpecBindingType), reflect.TypeOf([]sata.CreateSpec{})))
    fieldNameMap["sata_adapters"] = "SataAdapters"
    fields["scsi_adapters"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(scsi.CreateSpecBindingType), reflect.TypeOf([]scsi.CreateSpec{})))
    fieldNameMap["scsi_adapters"] = "ScsiAdapters"
    fields["nvme_adapters"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(nvme.CreateSpecBindingType), reflect.TypeOf([]nvme.CreateSpec{})))
    fieldNameMap["nvme_adapters"] = "NvmeAdapters"
    fields["storage_policy"] = bindings.NewOptionalType(bindings.NewReferenceType(StoragePolicySpecBindingType))
    fieldNameMap["storage_policy"] = "StoragePolicy"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.VM.create_spec",fields, reflect.TypeOf(CreateSpec{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["guest_OS"] = bindings.NewEnumType("com.vmware.vcenter.vm.guest_OS", reflect.TypeOf(GuestOS(GuestOS_DOS)))
    fieldNameMap["guest_OS"] = "GuestOS"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["identity"] = bindings.NewOptionalType(bindings.NewReferenceType(identity.InfoBindingType))
    fieldNameMap["identity"] = "Identity"
    fields["power_state"] = bindings.NewEnumType("com.vmware.vcenter.vm.power.state", reflect.TypeOf(power.State(power.State_POWERED_OFF)))
    fieldNameMap["power_state"] = "PowerState"
    fields["instant_clone_frozen"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["instant_clone_frozen"] = "InstantCloneFrozen"
    fields["hardware"] = bindings.NewReferenceType(hardware.InfoBindingType)
    fieldNameMap["hardware"] = "Hardware"
    fields["boot"] = bindings.NewReferenceType(boot.InfoBindingType)
    fieldNameMap["boot"] = "Boot"
    fields["boot_devices"] = bindings.NewListType(bindings.NewReferenceType(device.EntryBindingType), reflect.TypeOf([]device.Entry{}))
    fieldNameMap["boot_devices"] = "BootDevices"
    fields["cpu"] = bindings.NewReferenceType(cpu.InfoBindingType)
    fieldNameMap["cpu"] = "Cpu"
    fields["memory"] = bindings.NewReferenceType(memory.InfoBindingType)
    fieldNameMap["memory"] = "Memory"
    fields["disks"] = bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Disk"}, ""), bindings.NewReferenceType(disk.InfoBindingType),reflect.TypeOf(map[string]disk.Info{}))
    fieldNameMap["disks"] = "Disks"
    fields["nics"] = bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Ethernet"}, ""), bindings.NewReferenceType(ethernet.InfoBindingType),reflect.TypeOf(map[string]ethernet.Info{}))
    fieldNameMap["nics"] = "Nics"
    fields["cdroms"] = bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Cdrom"}, ""), bindings.NewReferenceType(cdrom.InfoBindingType),reflect.TypeOf(map[string]cdrom.Info{}))
    fieldNameMap["cdroms"] = "Cdroms"
    fields["floppies"] = bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Floppy"}, ""), bindings.NewReferenceType(floppy.InfoBindingType),reflect.TypeOf(map[string]floppy.Info{}))
    fieldNameMap["floppies"] = "Floppies"
    fields["parallel_ports"] = bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.ParallelPort"}, ""), bindings.NewReferenceType(parallel.InfoBindingType),reflect.TypeOf(map[string]parallel.Info{}))
    fieldNameMap["parallel_ports"] = "ParallelPorts"
    fields["serial_ports"] = bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.SerialPort"}, ""), bindings.NewReferenceType(serial.InfoBindingType),reflect.TypeOf(map[string]serial.Info{}))
    fieldNameMap["serial_ports"] = "SerialPorts"
    fields["sata_adapters"] = bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.SataAdapter"}, ""), bindings.NewReferenceType(sata.InfoBindingType),reflect.TypeOf(map[string]sata.Info{}))
    fieldNameMap["sata_adapters"] = "SataAdapters"
    fields["scsi_adapters"] = bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.ScsiAdapter"}, ""), bindings.NewReferenceType(scsi.InfoBindingType),reflect.TypeOf(map[string]scsi.Info{}))
    fieldNameMap["scsi_adapters"] = "ScsiAdapters"
    fields["nvme_adapters"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.NvmeAdapter"}, ""), bindings.NewReferenceType(nvme.InfoBindingType),reflect.TypeOf(map[string]nvme.Info{})))
    fieldNameMap["nvme_adapters"] = "NvmeAdapters"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.VM.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func GuestCustomizationSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["name"] = "Name"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.VM.guest_customization_spec",fields, reflect.TypeOf(GuestCustomizationSpec{}), fieldNameMap, validators)
}

func DiskCloneSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["datastore"] = bindings.NewOptionalType(bindings.NewIdType([]string {"Datastore"}, ""))
    fieldNameMap["datastore"] = "Datastore"
    fields["storage_policy"] = bindings.NewOptionalType(bindings.NewReferenceType(disk.StoragePolicySpecBindingType))
    fieldNameMap["storage_policy"] = "StoragePolicy"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.VM.disk_clone_spec",fields, reflect.TypeOf(DiskCloneSpec{}), fieldNameMap, validators)
}

func ClonePlacementSpecBindingType() bindings.BindingType {
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
    fields["datastore"] = bindings.NewOptionalType(bindings.NewIdType([]string {"Datastore"}, ""))
    fieldNameMap["datastore"] = "Datastore"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.VM.clone_placement_spec",fields, reflect.TypeOf(ClonePlacementSpec{}), fieldNameMap, validators)
}

func CloneSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["source"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fieldNameMap["source"] = "Source"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["placement"] = bindings.NewOptionalType(bindings.NewReferenceType(ClonePlacementSpecBindingType))
    fieldNameMap["placement"] = "Placement"
    fields["disks_to_remove"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Disk"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["disks_to_remove"] = "DisksToRemove"
    fields["disks_to_update"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Disk"}, ""), bindings.NewReferenceType(DiskCloneSpecBindingType),reflect.TypeOf(map[string]DiskCloneSpec{})))
    fieldNameMap["disks_to_update"] = "DisksToUpdate"
    fields["vm_home_storage_policy"] = bindings.NewOptionalType(bindings.NewReferenceType(StoragePolicySpecBindingType))
    fieldNameMap["vm_home_storage_policy"] = "VmHomeStoragePolicy"
    fields["disk_storage_policy"] = bindings.NewOptionalType(bindings.NewReferenceType(StoragePolicySpecBindingType))
    fieldNameMap["disk_storage_policy"] = "DiskStoragePolicy"
    fields["fallback_to_datastore_default_policy"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["fallback_to_datastore_default_policy"] = "FallbackToDatastoreDefaultPolicy"
    fields["power_on"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["power_on"] = "PowerOn"
    fields["guest_customization_spec"] = bindings.NewOptionalType(bindings.NewReferenceType(GuestCustomizationSpecBindingType))
    fieldNameMap["guest_customization_spec"] = "GuestCustomizationSpec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.VM.clone_spec",fields, reflect.TypeOf(CloneSpec{}), fieldNameMap, validators)
}

func DiskRelocateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["datastore"] = bindings.NewOptionalType(bindings.NewIdType([]string {"Datastore"}, ""))
    fieldNameMap["datastore"] = "Datastore"
    fields["storage_policy"] = bindings.NewOptionalType(bindings.NewReferenceType(disk.StoragePolicySpecBindingType))
    fieldNameMap["storage_policy"] = "StoragePolicy"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.VM.disk_relocate_spec",fields, reflect.TypeOf(DiskRelocateSpec{}), fieldNameMap, validators)
}

func RelocatePlacementSpecBindingType() bindings.BindingType {
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
    fields["datastore"] = bindings.NewOptionalType(bindings.NewIdType([]string {"Datastore"}, ""))
    fieldNameMap["datastore"] = "Datastore"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.VM.relocate_placement_spec",fields, reflect.TypeOf(RelocatePlacementSpec{}), fieldNameMap, validators)
}

func RelocateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["placement"] = bindings.NewOptionalType(bindings.NewReferenceType(RelocatePlacementSpecBindingType))
    fieldNameMap["placement"] = "Placement"
    fields["disks"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Disk"}, ""), bindings.NewReferenceType(DiskRelocateSpecBindingType),reflect.TypeOf(map[string]DiskRelocateSpec{})))
    fieldNameMap["disks"] = "Disks"
    fields["vm_home_storage_policy"] = bindings.NewOptionalType(bindings.NewReferenceType(StoragePolicySpecBindingType))
    fieldNameMap["vm_home_storage_policy"] = "VmHomeStoragePolicy"
    fields["disk_storage_policy"] = bindings.NewOptionalType(bindings.NewReferenceType(StoragePolicySpecBindingType))
    fieldNameMap["disk_storage_policy"] = "DiskStoragePolicy"
    fields["fallback_to_datastore_default_policy"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["fallback_to_datastore_default_policy"] = "FallbackToDatastoreDefaultPolicy"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.VM.relocate_spec",fields, reflect.TypeOf(RelocateSpec{}), fieldNameMap, validators)
}

func InstantClonePlacementSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["folder"] = bindings.NewOptionalType(bindings.NewIdType([]string {"Folder"}, ""))
    fieldNameMap["folder"] = "Folder"
    fields["resource_pool"] = bindings.NewOptionalType(bindings.NewIdType([]string {"ResourcePool"}, ""))
    fieldNameMap["resource_pool"] = "ResourcePool"
    fields["datastore"] = bindings.NewOptionalType(bindings.NewIdType([]string {"Datastore"}, ""))
    fieldNameMap["datastore"] = "Datastore"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.VM.instant_clone_placement_spec",fields, reflect.TypeOf(InstantClonePlacementSpec{}), fieldNameMap, validators)
}

func InstantCloneSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["source"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fieldNameMap["source"] = "Source"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["placement"] = bindings.NewOptionalType(bindings.NewReferenceType(InstantClonePlacementSpecBindingType))
    fieldNameMap["placement"] = "Placement"
    fields["nics_to_update"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Ethernet"}, ""), bindings.NewReferenceType(ethernet.UpdateSpecBindingType),reflect.TypeOf(map[string]ethernet.UpdateSpec{})))
    fieldNameMap["nics_to_update"] = "NicsToUpdate"
    fields["disconnect_all_nics"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["disconnect_all_nics"] = "DisconnectAllNics"
    fields["parallel_ports_to_update"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.ParallelPort"}, ""), bindings.NewReferenceType(parallel.UpdateSpecBindingType),reflect.TypeOf(map[string]parallel.UpdateSpec{})))
    fieldNameMap["parallel_ports_to_update"] = "ParallelPortsToUpdate"
    fields["serial_ports_to_update"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.SerialPort"}, ""), bindings.NewReferenceType(serial.UpdateSpecBindingType),reflect.TypeOf(map[string]serial.UpdateSpec{})))
    fieldNameMap["serial_ports_to_update"] = "SerialPortsToUpdate"
    fields["bios_uuid"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["bios_uuid"] = "BiosUuid"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.VM.instant_clone_spec",fields, reflect.TypeOf(InstantCloneSpec{}), fieldNameMap, validators)
}

func FilterSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vms"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"VirtualMachine"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["vms"] = "Vms"
    fields["names"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["names"] = "Names"
    fields["folders"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"Folder"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["folders"] = "Folders"
    fields["datacenters"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"Datacenter"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["datacenters"] = "Datacenters"
    fields["hosts"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"HostSystem"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["hosts"] = "Hosts"
    fields["clusters"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"ClusterComputeResource"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["clusters"] = "Clusters"
    fields["resource_pools"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"ResourcePool"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["resource_pools"] = "ResourcePools"
    fields["power_states"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewEnumType("com.vmware.vcenter.vm.power.state", reflect.TypeOf(power.State(power.State_POWERED_OFF))), reflect.TypeOf(map[power.State]bool{})))
    fieldNameMap["power_states"] = "PowerStates"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.VM.filter_spec",fields, reflect.TypeOf(FilterSpec{}), fieldNameMap, validators)
}

func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fieldNameMap["vm"] = "Vm"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["power_state"] = bindings.NewEnumType("com.vmware.vcenter.vm.power.state", reflect.TypeOf(power.State(power.State_POWERED_OFF)))
    fieldNameMap["power_state"] = "PowerState"
    fields["cpu_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["cpu_count"] = "CpuCount"
    fields["memory_size_MiB"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["memory_size_MiB"] = "MemorySizeMiB"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.VM.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}

func RegisterPlacementSpecBindingType() bindings.BindingType {
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
    return bindings.NewStructType("com.vmware.vcenter.VM.register_placement_spec",fields, reflect.TypeOf(RegisterPlacementSpec{}), fieldNameMap, validators)
}

func RegisterSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["datastore"] = bindings.NewOptionalType(bindings.NewIdType([]string {"Datastore"}, ""))
    fieldNameMap["datastore"] = "Datastore"
    fields["path"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["path"] = "Path"
    fields["datastore_path"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["datastore_path"] = "DatastorePath"
    fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["name"] = "Name"
    fields["placement"] = bindings.NewOptionalType(bindings.NewReferenceType(RegisterPlacementSpecBindingType))
    fieldNameMap["placement"] = "Placement"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.VM.register_spec",fields, reflect.TypeOf(RegisterSpec{}), fieldNameMap, validators)
}


