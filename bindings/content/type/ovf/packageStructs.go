/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for package: com.vmware.content.type.ovf.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package ovf

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/content/type/ovf/policy"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
)



// Provide the CPU information in a template VM.
type Cpu struct {
    NumCpus int64
    Reservation *int64
    Limit *int64
    Shares *int64
}






// Provide the disk information in a template VM.
type Disk struct {
    Name string
    DiskCapacity int64
    StoragePolicy *policy.StoragePolicy
}






// Provide the disk controller information in a template VM.
type DiskController struct {
    Name string
    Type_ *string
    SubType *string
}






// Provide the drive information in a template VM.
type Drive struct {
    Name string
    Type_ *string
    SubType *string
}






// Provide the floppy information in a template VM.
type Floppy struct {
    Name string
    Connected *bool
    Type_ *string
}






// Provide the memory information in a template VM.
type Memory struct {
    Size int64
    Reservation *int64
    Limit *int64
    Shares *int64
}






// Provide network information in a template VM.
type Network struct {
    Name string
    Description *string
}






// Provide NIC information in a VM template
type Nic struct {
    Name string
    NetworkName *string
    MacAddress *string
    StartConnected *bool
}






// Provides extra information about a library item of type "ovf". 
//
//  An OVF library item is the basic building block for instantiating virtual machines from content library. It may contain one or multiple virtual machine templates. This structure provides a rich view of the virtual machines within the ovf container as well as information about to the ovf descriptor associated with the library item 
type OvfTemplate struct {
    Id string
    VmCount int64
    Version string
    LibraryIdParent string
    IsVappTemplate bool
    VmTemplate *VmTemplate
    VappTemplate *VAppTemplate
    Networks []Network
    StoragePolicyGroups []policy.StoragePolicyGroup
}






// Provide USB controller information in a template VM.
type USBController struct {
    Type_ *string
    AutoConnect *bool
    EhciPciSlotNumber *int64
    PciSlotNumber *int64
}






// Provide information for vApp template in an OVF template file.
type VAppTemplate struct {
    VappName *string
    VmTemplates []VmTemplate
    StoragePolicies []policy.StoragePolicy
}






// Provide video card information in a template VM.
type VideoCard struct {
    RenderType *string
    VideoRamSize *int64
    GraphicsMemorySize *int64
    Enable3d *bool
    NumDisplays *int64
    UseAutoDetect *bool
}






// Provide template VM information in an OVF template (see OvfTemplate#type). The template VM provide the information about the operation system, CPU, memory, disks and NICs.
type VmTemplate struct {
    VmName string
    OsType *string
    OsDescription *string
    Cpu *Cpu
    Memory *Memory
    Disks []Disk
    Nics []Nic
    VideoCards []VideoCard
    Drives []Drive
    Floppies []Floppy
    DiskControllers []DiskController
    UsbControllers []USBController
    StoragePolicies []policy.StoragePolicy
}










func CpuBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["num_cpus"] = bindings.NewIntegerType()
    fieldNameMap["num_cpus"] = "NumCpus"
    fields["reservation"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["reservation"] = "Reservation"
    fields["limit"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["limit"] = "Limit"
    fields["shares"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["shares"] = "Shares"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.content.type.ovf.cpu",fields, reflect.TypeOf(Cpu{}), fieldNameMap, validators)
}

func DiskBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["disk_capacity"] = bindings.NewIntegerType()
    fieldNameMap["disk_capacity"] = "DiskCapacity"
    fields["storage_policy"] = bindings.NewOptionalType(bindings.NewReferenceType(policy.StoragePolicyBindingType))
    fieldNameMap["storage_policy"] = "StoragePolicy"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.content.type.ovf.disk",fields, reflect.TypeOf(Disk{}), fieldNameMap, validators)
}

func DiskControllerBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["type"] = "Type_"
    fields["sub_type"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["sub_type"] = "SubType"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.content.type.ovf.disk_controller",fields, reflect.TypeOf(DiskController{}), fieldNameMap, validators)
}

func DriveBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["type"] = "Type_"
    fields["sub_type"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["sub_type"] = "SubType"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.content.type.ovf.drive",fields, reflect.TypeOf(Drive{}), fieldNameMap, validators)
}

func FloppyBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["connected"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["connected"] = "Connected"
    fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["type"] = "Type_"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.content.type.ovf.floppy",fields, reflect.TypeOf(Floppy{}), fieldNameMap, validators)
}

func MemoryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["size"] = bindings.NewIntegerType()
    fieldNameMap["size"] = "Size"
    fields["reservation"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["reservation"] = "Reservation"
    fields["limit"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["limit"] = "Limit"
    fields["shares"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["shares"] = "Shares"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.content.type.ovf.memory",fields, reflect.TypeOf(Memory{}), fieldNameMap, validators)
}

func NetworkBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["description"] = "Description"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.content.type.ovf.network",fields, reflect.TypeOf(Network{}), fieldNameMap, validators)
}

func NicBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["network_name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["network_name"] = "NetworkName"
    fields["mac_address"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["mac_address"] = "MacAddress"
    fields["start_connected"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["start_connected"] = "StartConnected"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.content.type.ovf.nic",fields, reflect.TypeOf(Nic{}), fieldNameMap, validators)
}

func OvfTemplateBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.content.library.Item"}, "")
    fieldNameMap["id"] = "Id"
    fields["vm_count"] = bindings.NewIntegerType()
    fieldNameMap["vm_count"] = "VmCount"
    fields["version"] = bindings.NewStringType()
    fieldNameMap["version"] = "Version"
    fields["library_id_parent"] = bindings.NewIdType([]string {"com.vmware.content.Library"}, "")
    fieldNameMap["library_id_parent"] = "LibraryIdParent"
    fields["is_vapp_template"] = bindings.NewBooleanType()
    fieldNameMap["is_vapp_template"] = "IsVappTemplate"
    fields["vm_template"] = bindings.NewOptionalType(bindings.NewReferenceType(VmTemplateBindingType))
    fieldNameMap["vm_template"] = "VmTemplate"
    fields["vapp_template"] = bindings.NewOptionalType(bindings.NewReferenceType(VAppTemplateBindingType))
    fieldNameMap["vapp_template"] = "VappTemplate"
    fields["networks"] = bindings.NewListType(bindings.NewReferenceType(NetworkBindingType), reflect.TypeOf([]Network{}))
    fieldNameMap["networks"] = "Networks"
    fields["storage_policy_groups"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(policy.StoragePolicyGroupBindingType), reflect.TypeOf([]policy.StoragePolicyGroup{})))
    fieldNameMap["storage_policy_groups"] = "StoragePolicyGroups"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.content.type.ovf.ovf_template",fields, reflect.TypeOf(OvfTemplate{}), fieldNameMap, validators)
}

func USBControllerBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["type"] = "Type_"
    fields["auto_connect"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["auto_connect"] = "AutoConnect"
    fields["ehci_pci_slot_number"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["ehci_pci_slot_number"] = "EhciPciSlotNumber"
    fields["pci_slot_number"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["pci_slot_number"] = "PciSlotNumber"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.content.type.ovf.USB_controller",fields, reflect.TypeOf(USBController{}), fieldNameMap, validators)
}

func VAppTemplateBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vapp_name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["vapp_name"] = "VappName"
    fields["vm_templates"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(VmTemplateBindingType), reflect.TypeOf([]VmTemplate{})))
    fieldNameMap["vm_templates"] = "VmTemplates"
    fields["storage_policies"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(policy.StoragePolicyBindingType), reflect.TypeOf([]policy.StoragePolicy{})))
    fieldNameMap["storage_policies"] = "StoragePolicies"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.content.type.ovf.V_app_template",fields, reflect.TypeOf(VAppTemplate{}), fieldNameMap, validators)
}

func VideoCardBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["render_type"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["render_type"] = "RenderType"
    fields["video_ram_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["video_ram_size"] = "VideoRamSize"
    fields["graphics_memory_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["graphics_memory_size"] = "GraphicsMemorySize"
    fields["enable3d"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["enable3d"] = "Enable3d"
    fields["num_displays"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["num_displays"] = "NumDisplays"
    fields["use_auto_detect"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["use_auto_detect"] = "UseAutoDetect"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.content.type.ovf.video_card",fields, reflect.TypeOf(VideoCard{}), fieldNameMap, validators)
}

func VmTemplateBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm_name"] = bindings.NewStringType()
    fieldNameMap["vm_name"] = "VmName"
    fields["os_type"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["os_type"] = "OsType"
    fields["os_description"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["os_description"] = "OsDescription"
    fields["cpu"] = bindings.NewOptionalType(bindings.NewReferenceType(CpuBindingType))
    fieldNameMap["cpu"] = "Cpu"
    fields["memory"] = bindings.NewOptionalType(bindings.NewReferenceType(MemoryBindingType))
    fieldNameMap["memory"] = "Memory"
    fields["disks"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DiskBindingType), reflect.TypeOf([]Disk{})))
    fieldNameMap["disks"] = "Disks"
    fields["nics"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(NicBindingType), reflect.TypeOf([]Nic{})))
    fieldNameMap["nics"] = "Nics"
    fields["video_cards"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(VideoCardBindingType), reflect.TypeOf([]VideoCard{})))
    fieldNameMap["video_cards"] = "VideoCards"
    fields["drives"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DriveBindingType), reflect.TypeOf([]Drive{})))
    fieldNameMap["drives"] = "Drives"
    fields["floppies"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(FloppyBindingType), reflect.TypeOf([]Floppy{})))
    fieldNameMap["floppies"] = "Floppies"
    fields["disk_controllers"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DiskControllerBindingType), reflect.TypeOf([]DiskController{})))
    fieldNameMap["disk_controllers"] = "DiskControllers"
    fields["usb_controllers"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(USBControllerBindingType), reflect.TypeOf([]USBController{})))
    fieldNameMap["usb_controllers"] = "UsbControllers"
    fields["storage_policies"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(policy.StoragePolicyBindingType), reflect.TypeOf([]policy.StoragePolicy{})))
    fieldNameMap["storage_policies"] = "StoragePolicies"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.content.type.ovf.vm_template",fields, reflect.TypeOf(VmTemplate{}), fieldNameMap, validators)
}


