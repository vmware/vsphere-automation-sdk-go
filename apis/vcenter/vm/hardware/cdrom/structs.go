/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Cdrom.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package cdrom

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vcenter/vm/hardware"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// Resource type for the virtual CD-ROM device.
const Cdrom_RESOURCE_TYPE = "com.vmware.vcenter.vm.hardware.Cdrom"


// The ``HostBusAdapterType`` enumeration class defines the valid types of host bus adapters that may be used for attaching a Cdrom to a virtual machine.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type HostBusAdapterType string

const (
    // Cdrom is attached to an IDE adapter.
     HostBusAdapterType_IDE HostBusAdapterType = "IDE"
    // Cdrom is attached to a SATA adapter.
     HostBusAdapterType_SATA HostBusAdapterType = "SATA"
)

func (h HostBusAdapterType) HostBusAdapterType() bool {
    switch h {
        case HostBusAdapterType_IDE:
            return true
        case HostBusAdapterType_SATA:
            return true
        default:
            return false
    }
}




// The ``BackingType`` enumeration class defines the valid backing types for a virtual CD-ROM device.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type BackingType string

const (
    // Virtual CD-ROM device is backed by an ISO file.
     BackingType_ISO_FILE BackingType = "ISO_FILE"
    // Virtual CD-ROM device is backed by a device on the host where the virtual machine is running.
     BackingType_HOST_DEVICE BackingType = "HOST_DEVICE"
    // Virtual CD-ROM device is backed by a device on the client that is connected to the virtual machine console.
     BackingType_CLIENT_DEVICE BackingType = "CLIENT_DEVICE"
)

func (b BackingType) BackingType() bool {
    switch b {
        case BackingType_ISO_FILE:
            return true
        case BackingType_HOST_DEVICE:
            return true
        case BackingType_CLIENT_DEVICE:
            return true
        default:
            return false
    }
}




// The ``DeviceAccessType`` enumeration class defines the valid device access types for a physical device packing of a virtual CD-ROM device.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type DeviceAccessType string

const (
    // ATAPI or SCSI device emulation.
     DeviceAccessType_EMULATION DeviceAccessType = "EMULATION"
    // Raw passthru device access.
     DeviceAccessType_PASSTHRU DeviceAccessType = "PASSTHRU"
    // Raw passthru device access, with exclusive access to the device.
     DeviceAccessType_PASSTHRU_EXCLUSIVE DeviceAccessType = "PASSTHRU_EXCLUSIVE"
)

func (d DeviceAccessType) DeviceAccessType() bool {
    switch d {
        case DeviceAccessType_EMULATION:
            return true
        case DeviceAccessType_PASSTHRU:
            return true
        case DeviceAccessType_PASSTHRU_EXCLUSIVE:
            return true
        default:
            return false
    }
}





// The ``BackingInfo`` class contains information about the physical resource backing a virtual CD-ROM device.
 type BackingInfo struct {
    Type_ BackingType
    IsoFile *string
    HostDevice *string
    AutoDetect *bool
    DeviceAccessType *DeviceAccessType
}






// The ``BackingSpec`` class provides a specification of the physical resource backing a virtual CD-ROM device.
 type BackingSpec struct {
    Type_ BackingType
    IsoFile *string
    HostDevice *string
    DeviceAccessType *DeviceAccessType
}






// The ``Info`` class contains information about a virtual CD-ROM device.
 type Info struct {
    Type_ HostBusAdapterType
    Label string
    Ide *hardware.IdeAddressInfo
    Sata *hardware.SataAddressInfo
    Backing BackingInfo
    State hardware.ConnectionState
    StartConnected bool
    AllowGuestControl bool
}






// The ``CreateSpec`` class provides a specification for the configuration of a newly-created virtual CD-ROM device.
 type CreateSpec struct {
    Type_ *HostBusAdapterType
    Ide *hardware.IdeAddressSpec
    Sata *hardware.SataAddressSpec
    Backing *BackingSpec
    StartConnected *bool
    AllowGuestControl *bool
}






// The ``UpdateSpec`` class describes the updates to be made to the configuration of a virtual CD-ROM device.
 type UpdateSpec struct {
    Backing *BackingSpec
    StartConnected *bool
    AllowGuestControl *bool
}






// The ``Summary`` class contains commonly used information about a virtual CD-ROM device.
 type Summary struct {
    Cdrom string
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
    fields["cdrom"] = bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Cdrom"}, "")
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["cdrom"] = "Cdrom"
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
    return bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Cdrom"}, "")
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
    fields["cdrom"] = bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Cdrom"}, "")
    fields["spec"] = bindings.NewReferenceType(UpdateSpecBindingType)
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["cdrom"] = "Cdrom"
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
    fields["cdrom"] = bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Cdrom"}, "")
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["cdrom"] = "Cdrom"
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


func connectInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fields["cdrom"] = bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Cdrom"}, "")
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["cdrom"] = "Cdrom"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func connectOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func connectRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"NotFound": 404,"AlreadyInDesiredState": 400,"NotAllowedInCurrentState": 400,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func disconnectInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fields["cdrom"] = bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Cdrom"}, "")
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["cdrom"] = "Cdrom"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func disconnectOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func disconnectRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"NotFound": 404,"AlreadyInDesiredState": 400,"NotAllowedInCurrentState": 400,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}



func BackingInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.cdrom.backing_type", reflect.TypeOf(BackingType(BackingType_ISO_FILE)))
    fieldNameMap["type"] = "Type_"
    fields["iso_file"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["iso_file"] = "IsoFile"
    fields["host_device"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["host_device"] = "HostDevice"
    fields["auto_detect"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["auto_detect"] = "AutoDetect"
    fields["device_access_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.hardware.cdrom.device_access_type", reflect.TypeOf(DeviceAccessType(DeviceAccessType_EMULATION))))
    fieldNameMap["device_access_type"] = "DeviceAccessType"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("type",
        map[string][]bindings.FieldData {
            "ISO_FILE": []bindings.FieldData {
                 bindings.NewFieldData("iso_file", true),
            },
            "HOST_DEVICE": []bindings.FieldData {
                 bindings.NewFieldData("host_device", false),
                 bindings.NewFieldData("auto_detect", true),
                 bindings.NewFieldData("device_access_type", true),
            },
            "CLIENT_DEVICE": []bindings.FieldData {
                 bindings.NewFieldData("device_access_type", true),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.cdrom.backing_info",fields, reflect.TypeOf(BackingInfo{}), fieldNameMap, validators)
}

func BackingSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.cdrom.backing_type", reflect.TypeOf(BackingType(BackingType_ISO_FILE)))
    fieldNameMap["type"] = "Type_"
    fields["iso_file"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["iso_file"] = "IsoFile"
    fields["host_device"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["host_device"] = "HostDevice"
    fields["device_access_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.hardware.cdrom.device_access_type", reflect.TypeOf(DeviceAccessType(DeviceAccessType_EMULATION))))
    fieldNameMap["device_access_type"] = "DeviceAccessType"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("type",
        map[string][]bindings.FieldData {
            "ISO_FILE": []bindings.FieldData {
                 bindings.NewFieldData("iso_file", true),
            },
            "HOST_DEVICE": []bindings.FieldData {
                 bindings.NewFieldData("host_device", false),
                 bindings.NewFieldData("device_access_type", false),
            },
            "CLIENT_DEVICE": []bindings.FieldData {
                 bindings.NewFieldData("device_access_type", false),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.cdrom.backing_spec",fields, reflect.TypeOf(BackingSpec{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.cdrom.host_bus_adapter_type", reflect.TypeOf(HostBusAdapterType(HostBusAdapterType_IDE)))
    fieldNameMap["type"] = "Type_"
    fields["label"] = bindings.NewStringType()
    fieldNameMap["label"] = "Label"
    fields["ide"] = bindings.NewOptionalType(bindings.NewReferenceType(hardware.IdeAddressInfoBindingType))
    fieldNameMap["ide"] = "Ide"
    fields["sata"] = bindings.NewOptionalType(bindings.NewReferenceType(hardware.SataAddressInfoBindingType))
    fieldNameMap["sata"] = "Sata"
    fields["backing"] = bindings.NewReferenceType(BackingInfoBindingType)
    fieldNameMap["backing"] = "Backing"
    fields["state"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.connection_state", reflect.TypeOf(hardware.ConnectionState(hardware.ConnectionState_CONNECTED)))
    fieldNameMap["state"] = "State"
    fields["start_connected"] = bindings.NewBooleanType()
    fieldNameMap["start_connected"] = "StartConnected"
    fields["allow_guest_control"] = bindings.NewBooleanType()
    fieldNameMap["allow_guest_control"] = "AllowGuestControl"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("type",
        map[string][]bindings.FieldData {
            "IDE": []bindings.FieldData {
                 bindings.NewFieldData("ide", true),
            },
            "SATA": []bindings.FieldData {
                 bindings.NewFieldData("sata", true),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.cdrom.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func CreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.hardware.cdrom.host_bus_adapter_type", reflect.TypeOf(HostBusAdapterType(HostBusAdapterType_IDE))))
    fieldNameMap["type"] = "Type_"
    fields["ide"] = bindings.NewOptionalType(bindings.NewReferenceType(hardware.IdeAddressSpecBindingType))
    fieldNameMap["ide"] = "Ide"
    fields["sata"] = bindings.NewOptionalType(bindings.NewReferenceType(hardware.SataAddressSpecBindingType))
    fieldNameMap["sata"] = "Sata"
    fields["backing"] = bindings.NewOptionalType(bindings.NewReferenceType(BackingSpecBindingType))
    fieldNameMap["backing"] = "Backing"
    fields["start_connected"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["start_connected"] = "StartConnected"
    fields["allow_guest_control"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["allow_guest_control"] = "AllowGuestControl"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("type",
        map[string][]bindings.FieldData {
            "IDE": []bindings.FieldData {
                 bindings.NewFieldData("ide", false),
            },
            "SATA": []bindings.FieldData {
                 bindings.NewFieldData("sata", false),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.cdrom.create_spec",fields, reflect.TypeOf(CreateSpec{}), fieldNameMap, validators)
}

func UpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["backing"] = bindings.NewOptionalType(bindings.NewReferenceType(BackingSpecBindingType))
    fieldNameMap["backing"] = "Backing"
    fields["start_connected"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["start_connected"] = "StartConnected"
    fields["allow_guest_control"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["allow_guest_control"] = "AllowGuestControl"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.cdrom.update_spec",fields, reflect.TypeOf(UpdateSpec{}), fieldNameMap, validators)
}

func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cdrom"] = bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Cdrom"}, "")
    fieldNameMap["cdrom"] = "Cdrom"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.cdrom.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}


