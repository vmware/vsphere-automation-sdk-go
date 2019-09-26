/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Disk.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package disk

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/vcenter/vm/hardware"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// Resource type for the virtual disk.
const Disk_RESOURCE_TYPE = "com.vmware.vcenter.vm.hardware.Disk"


// The ``HostBusAdapterType`` enumeration class defines the valid types of host bus adapters that may be used for attaching a virtual storage device to a virtual machine.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type HostBusAdapterType string

const (
    // Disk is attached to an IDE adapter.
     HostBusAdapterType_IDE HostBusAdapterType = "IDE"
    // Disk is attached to a SCSI adapter.
     HostBusAdapterType_SCSI HostBusAdapterType = "SCSI"
    // Disk is attached to a SATA adapter.
     HostBusAdapterType_SATA HostBusAdapterType = "SATA"
    // Disk is attached to a NVMe adapter. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     HostBusAdapterType_NVME HostBusAdapterType = "NVME"
)

func (h HostBusAdapterType) HostBusAdapterType() bool {
    switch h {
        case HostBusAdapterType_IDE:
            return true
        case HostBusAdapterType_SCSI:
            return true
        case HostBusAdapterType_SATA:
            return true
        case HostBusAdapterType_NVME:
            return true
        default:
            return false
    }
}




// The ``BackingType`` enumeration class defines the valid backing types for a virtual disk.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type BackingType string

const (
    // Virtual disk is backed by a VMDK file.
     BackingType_VMDK_FILE BackingType = "VMDK_FILE"
)

func (b BackingType) BackingType() bool {
    switch b {
        case BackingType_VMDK_FILE:
            return true
        default:
            return false
    }
}





// The ``BackingInfo`` class contains information about the physical resource backing a virtual disk.
 type BackingInfo struct {
    Type_ BackingType
    VmdkFile *string
}






// The ``BackingSpec`` class provides a specification of the physical resource backing a virtual disk.
 type BackingSpec struct {
    Type_ BackingType
    VmdkFile *string
}






// The ``VmdkCreateSpec`` class provides a specification for creating a new VMDK file to be used as a backing for a virtual disk. The virtual disk will be stored in the same directory as the virtual machine's configuration file.
 type VmdkCreateSpec struct {
    Name *string
    Capacity *int64
    StoragePolicy *StoragePolicySpec
}






// The ``Info`` class contains information about a virtual disk.
 type Info struct {
    Label string
    Type_ HostBusAdapterType
    Ide *hardware.IdeAddressInfo
    Scsi *hardware.ScsiAddressInfo
    Sata *hardware.SataAddressInfo
    Nvme *hardware.NvmeAddressInfo
    Backing BackingInfo
    Capacity *int64
}






// The ``StoragePolicySpec`` class contains information about the storage policy be associated with a VMDK file.
 type StoragePolicySpec struct {
    Policy string
}






// The ``CreateSpec`` class provides a specification for the configuration of a newly-created virtual disk.
 type CreateSpec struct {
    Type_ *HostBusAdapterType
    Ide *hardware.IdeAddressSpec
    Scsi *hardware.ScsiAddressSpec
    Sata *hardware.SataAddressSpec
    Nvme *hardware.NvmeAddressSpec
    Backing *BackingSpec
    NewVmdk *VmdkCreateSpec
}






// The ``UpdateSpec`` class describes the updates to be made to the configuration of a virtual disk.
 type UpdateSpec struct {
    Backing *BackingSpec
}






// The ``Summary`` class contains commonly used information about a virtual disk.
 type Summary struct {
    Disk string
}









func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fieldNameMap["vm"] = "Vm"
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
       map[string]int{"Error": 500,"NotFound": 404,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fields["disk"] = bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Disk"}, "")
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["disk"] = "Disk"
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


func createInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fields["spec"] = bindings.NewReferenceType(CreateSpecBindingType)
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func createOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Disk"}, "")
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
       map[string]int{"Error": 500,"NotFound": 404,"NotAllowedInCurrentState": 400,"UnableToAllocateResource": 500,"ResourceInUse": 400,"InvalidArgument": 400,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403,"Unsupported": 400})
}


func updateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fields["disk"] = bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Disk"}, "")
    fields["spec"] = bindings.NewReferenceType(UpdateSpecBindingType)
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["disk"] = "Disk"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func updateOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func updateRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"NotFound": 404,"NotAllowedInCurrentState": 400,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func deleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fields["disk"] = bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Disk"}, "")
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["disk"] = "Disk"
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
       map[string]int{"Error": 500,"NotFound": 404,"NotAllowedInCurrentState": 400,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}



func BackingInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.disk.backing_type", reflect.TypeOf(BackingType(BackingType_VMDK_FILE)))
    fieldNameMap["type"] = "Type_"
    fields["vmdk_file"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["vmdk_file"] = "VmdkFile"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("type",
        map[string][]bindings.FieldData {
            "VMDK_FILE": []bindings.FieldData {
                 bindings.NewFieldData("vmdk_file", true),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.disk.backing_info",fields, reflect.TypeOf(BackingInfo{}), fieldNameMap, validators)
}

func BackingSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.disk.backing_type", reflect.TypeOf(BackingType(BackingType_VMDK_FILE)))
    fieldNameMap["type"] = "Type_"
    fields["vmdk_file"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["vmdk_file"] = "VmdkFile"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("type",
        map[string][]bindings.FieldData {
            "VMDK_FILE": []bindings.FieldData {
                 bindings.NewFieldData("vmdk_file", true),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.disk.backing_spec",fields, reflect.TypeOf(BackingSpec{}), fieldNameMap, validators)
}

func VmdkCreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["name"] = "Name"
    fields["capacity"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["capacity"] = "Capacity"
    fields["storage_policy"] = bindings.NewOptionalType(bindings.NewReferenceType(StoragePolicySpecBindingType))
    fieldNameMap["storage_policy"] = "StoragePolicy"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.disk.vmdk_create_spec",fields, reflect.TypeOf(VmdkCreateSpec{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["label"] = bindings.NewStringType()
    fieldNameMap["label"] = "Label"
    fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.disk.host_bus_adapter_type", reflect.TypeOf(HostBusAdapterType(HostBusAdapterType_IDE)))
    fieldNameMap["type"] = "Type_"
    fields["ide"] = bindings.NewOptionalType(bindings.NewReferenceType(hardware.IdeAddressInfoBindingType))
    fieldNameMap["ide"] = "Ide"
    fields["scsi"] = bindings.NewOptionalType(bindings.NewReferenceType(hardware.ScsiAddressInfoBindingType))
    fieldNameMap["scsi"] = "Scsi"
    fields["sata"] = bindings.NewOptionalType(bindings.NewReferenceType(hardware.SataAddressInfoBindingType))
    fieldNameMap["sata"] = "Sata"
    fields["nvme"] = bindings.NewOptionalType(bindings.NewReferenceType(hardware.NvmeAddressInfoBindingType))
    fieldNameMap["nvme"] = "Nvme"
    fields["backing"] = bindings.NewReferenceType(BackingInfoBindingType)
    fieldNameMap["backing"] = "Backing"
    fields["capacity"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["capacity"] = "Capacity"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("type",
        map[string][]bindings.FieldData {
            "IDE": []bindings.FieldData {
                 bindings.NewFieldData("ide", true),
            },
            "SCSI": []bindings.FieldData {
                 bindings.NewFieldData("scsi", true),
            },
            "SATA": []bindings.FieldData {
                 bindings.NewFieldData("sata", true),
            },
            "NVME": []bindings.FieldData {
                 bindings.NewFieldData("nvme", true),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.disk.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func StoragePolicySpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["policy"] = bindings.NewIdType([]string {"com.vmware.vcenter.StoragePolicy"}, "")
    fieldNameMap["policy"] = "Policy"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.disk.storage_policy_spec",fields, reflect.TypeOf(StoragePolicySpec{}), fieldNameMap, validators)
}

func CreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.hardware.disk.host_bus_adapter_type", reflect.TypeOf(HostBusAdapterType(HostBusAdapterType_IDE))))
    fieldNameMap["type"] = "Type_"
    fields["ide"] = bindings.NewOptionalType(bindings.NewReferenceType(hardware.IdeAddressSpecBindingType))
    fieldNameMap["ide"] = "Ide"
    fields["scsi"] = bindings.NewOptionalType(bindings.NewReferenceType(hardware.ScsiAddressSpecBindingType))
    fieldNameMap["scsi"] = "Scsi"
    fields["sata"] = bindings.NewOptionalType(bindings.NewReferenceType(hardware.SataAddressSpecBindingType))
    fieldNameMap["sata"] = "Sata"
    fields["nvme"] = bindings.NewOptionalType(bindings.NewReferenceType(hardware.NvmeAddressSpecBindingType))
    fieldNameMap["nvme"] = "Nvme"
    fields["backing"] = bindings.NewOptionalType(bindings.NewReferenceType(BackingSpecBindingType))
    fieldNameMap["backing"] = "Backing"
    fields["new_vmdk"] = bindings.NewOptionalType(bindings.NewReferenceType(VmdkCreateSpecBindingType))
    fieldNameMap["new_vmdk"] = "NewVmdk"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("type",
        map[string][]bindings.FieldData {
            "IDE": []bindings.FieldData {
                 bindings.NewFieldData("ide", false),
            },
            "SCSI": []bindings.FieldData {
                 bindings.NewFieldData("scsi", false),
            },
            "SATA": []bindings.FieldData {
                 bindings.NewFieldData("sata", false),
            },
            "NVME": []bindings.FieldData {
                 bindings.NewFieldData("nvme", false),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.disk.create_spec",fields, reflect.TypeOf(CreateSpec{}), fieldNameMap, validators)
}

func UpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["backing"] = bindings.NewOptionalType(bindings.NewReferenceType(BackingSpecBindingType))
    fieldNameMap["backing"] = "Backing"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.disk.update_spec",fields, reflect.TypeOf(UpdateSpec{}), fieldNameMap, validators)
}

func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["disk"] = bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Disk"}, "")
    fieldNameMap["disk"] = "Disk"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.disk.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}


