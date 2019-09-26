/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Ethernet.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package ethernet

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/vcenter/vm/hardware"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// Resource type for the virtual Ethernet adapter.
const Ethernet_RESOURCE_TYPE = "com.vmware.vcenter.vm.hardware.Ethernet"


// The ``EmulationType`` enumeration class defines the valid emulation types for a virtual Ethernet adapter.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type EmulationType string

const (
    // E1000 ethernet adapter.
     EmulationType_E1000 EmulationType = "E1000"
    // E1000e ethernet adapter.
     EmulationType_E1000E EmulationType = "E1000E"
    // AMD Lance PCNet32 Ethernet adapter.
     EmulationType_PCNET32 EmulationType = "PCNET32"
    // VMware Vmxnet virtual Ethernet adapter.
     EmulationType_VMXNET EmulationType = "VMXNET"
    // VMware Vmxnet2 virtual Ethernet adapter.
     EmulationType_VMXNET2 EmulationType = "VMXNET2"
    // VMware Vmxnet3 virtual Ethernet adapter.
     EmulationType_VMXNET3 EmulationType = "VMXNET3"
)

func (e EmulationType) EmulationType() bool {
    switch e {
        case EmulationType_E1000:
            return true
        case EmulationType_E1000E:
            return true
        case EmulationType_PCNET32:
            return true
        case EmulationType_VMXNET:
            return true
        case EmulationType_VMXNET2:
            return true
        case EmulationType_VMXNET3:
            return true
        default:
            return false
    }
}




// The ``MacAddressType`` enumeration class defines the valid MAC address origins for a virtual Ethernet adapter.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type MacAddressType string

const (
    // MAC address is assigned statically.
     MacAddressType_MANUAL MacAddressType = "MANUAL"
    // MAC address is generated automatically.
     MacAddressType_GENERATED MacAddressType = "GENERATED"
    // MAC address is assigned by vCenter Server.
     MacAddressType_ASSIGNED MacAddressType = "ASSIGNED"
)

func (m MacAddressType) MacAddressType() bool {
    switch m {
        case MacAddressType_MANUAL:
            return true
        case MacAddressType_GENERATED:
            return true
        case MacAddressType_ASSIGNED:
            return true
        default:
            return false
    }
}




// The ``BackingType`` enumeration class defines the valid backing types for a virtual Ethernet adapter.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type BackingType string

const (
    // vSphere standard portgroup network backing.
     BackingType_STANDARD_PORTGROUP BackingType = "STANDARD_PORTGROUP"
    // Legacy host device network backing. Imported VMs may have virtual Ethernet adapters with this type of backing, but this type of backing cannot be used to create or to update a virtual Ethernet adapter.
     BackingType_HOST_DEVICE BackingType = "HOST_DEVICE"
    // Distributed virtual switch backing.
     BackingType_DISTRIBUTED_PORTGROUP BackingType = "DISTRIBUTED_PORTGROUP"
    // Opaque network backing.
     BackingType_OPAQUE_NETWORK BackingType = "OPAQUE_NETWORK"
)

func (b BackingType) BackingType() bool {
    switch b {
        case BackingType_STANDARD_PORTGROUP:
            return true
        case BackingType_HOST_DEVICE:
            return true
        case BackingType_DISTRIBUTED_PORTGROUP:
            return true
        case BackingType_OPAQUE_NETWORK:
            return true
        default:
            return false
    }
}





// The ``BackingInfo`` class contains information about the physical resource backing a virtual Ethernet adapter.
 type BackingInfo struct {
    Type_ BackingType
    Network *string
    NetworkName *string
    HostDevice *string
    DistributedSwitchUuid *string
    DistributedPort *string
    ConnectionCookie *int64
    OpaqueNetworkType *string
    OpaqueNetworkId *string
}






// The ``BackingSpec`` class provides a specification of the physical resource that backs a virtual Ethernet adapter.
 type BackingSpec struct {
    Type_ BackingType
    Network *string
    DistributedPort *string
}






// The ``Info`` class contains information about a virtual Ethernet adapter.
 type Info struct {
    Label string
    Type_ EmulationType
    UptCompatibilityEnabled *bool
    MacType MacAddressType
    MacAddress *string
    PciSlotNumber *int64
    WakeOnLanEnabled bool
    Backing BackingInfo
    State hardware.ConnectionState
    StartConnected bool
    AllowGuestControl bool
}






// The ``CreateSpec`` class provides a specification for the configuration of a newly-created virtual Ethernet adapter.
 type CreateSpec struct {
    Type_ *EmulationType
    UptCompatibilityEnabled *bool
    MacType *MacAddressType
    MacAddress *string
    PciSlotNumber *int64
    WakeOnLanEnabled *bool
    Backing *BackingSpec
    StartConnected *bool
    AllowGuestControl *bool
}






// The ``UpdateSpec`` class describes the updates to be made to the configuration of a virtual Ethernet adapter.
 type UpdateSpec struct {
    UptCompatibilityEnabled *bool
    MacType *MacAddressType
    MacAddress *string
    WakeOnLanEnabled *bool
    Backing *BackingSpec
    StartConnected *bool
    AllowGuestControl *bool
}






// The ``Summary`` class contains commonly used information about a virtual Ethernet adapter.
 type Summary struct {
    Nic string
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
    fields["nic"] = bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Ethernet"}, "")
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["nic"] = "Nic"
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
    return bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Ethernet"}, "")
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
       map[string]int{"Error": 500,"NotFound": 404,"UnableToAllocateResource": 500,"InvalidArgument": 400,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403,"Unsupported": 400})
}


func updateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fields["nic"] = bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Ethernet"}, "")
    fields["spec"] = bindings.NewReferenceType(UpdateSpecBindingType)
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["nic"] = "Nic"
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
       map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func deleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fields["nic"] = bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Ethernet"}, "")
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["nic"] = "Nic"
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
       map[string]int{"Error": 500,"NotFound": 404,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func connectInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fields["nic"] = bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Ethernet"}, "")
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["nic"] = "Nic"
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
    fields["nic"] = bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Ethernet"}, "")
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["nic"] = "Nic"
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
    fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.ethernet.backing_type", reflect.TypeOf(BackingType(BackingType_STANDARD_PORTGROUP)))
    fieldNameMap["type"] = "Type_"
    fields["network"] = bindings.NewOptionalType(bindings.NewIdType([]string {"Network"}, ""))
    fieldNameMap["network"] = "Network"
    fields["network_name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["network_name"] = "NetworkName"
    fields["host_device"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["host_device"] = "HostDevice"
    fields["distributed_switch_uuid"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["distributed_switch_uuid"] = "DistributedSwitchUuid"
    fields["distributed_port"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["distributed_port"] = "DistributedPort"
    fields["connection_cookie"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["connection_cookie"] = "ConnectionCookie"
    fields["opaque_network_type"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["opaque_network_type"] = "OpaqueNetworkType"
    fields["opaque_network_id"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["opaque_network_id"] = "OpaqueNetworkId"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("type",
        map[string][]bindings.FieldData {
            "STANDARD_PORTGROUP": []bindings.FieldData {
                 bindings.NewFieldData("network", false),
                 bindings.NewFieldData("network_name", true),
            },
            "DISTRIBUTED_PORTGROUP": []bindings.FieldData {
                 bindings.NewFieldData("network", false),
                 bindings.NewFieldData("distributed_switch_uuid", true),
                 bindings.NewFieldData("distributed_port", false),
                 bindings.NewFieldData("connection_cookie", false),
            },
            "OPAQUE_NETWORK": []bindings.FieldData {
                 bindings.NewFieldData("network", false),
                 bindings.NewFieldData("opaque_network_type", true),
                 bindings.NewFieldData("opaque_network_id", true),
            },
            "HOST_DEVICE": []bindings.FieldData {
                 bindings.NewFieldData("host_device", true),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.ethernet.backing_info",fields, reflect.TypeOf(BackingInfo{}), fieldNameMap, validators)
}

func BackingSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.ethernet.backing_type", reflect.TypeOf(BackingType(BackingType_STANDARD_PORTGROUP)))
    fieldNameMap["type"] = "Type_"
    fields["network"] = bindings.NewOptionalType(bindings.NewIdType([]string {"Network"}, ""))
    fieldNameMap["network"] = "Network"
    fields["distributed_port"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["distributed_port"] = "DistributedPort"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("type",
        map[string][]bindings.FieldData {
            "STANDARD_PORTGROUP": []bindings.FieldData {
                 bindings.NewFieldData("network", true),
            },
            "DISTRIBUTED_PORTGROUP": []bindings.FieldData {
                 bindings.NewFieldData("network", true),
                 bindings.NewFieldData("distributed_port", false),
            },
            "OPAQUE_NETWORK": []bindings.FieldData {
                 bindings.NewFieldData("network", true),
            },
            "HOST_DEVICE": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.ethernet.backing_spec",fields, reflect.TypeOf(BackingSpec{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["label"] = bindings.NewStringType()
    fieldNameMap["label"] = "Label"
    fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.ethernet.emulation_type", reflect.TypeOf(EmulationType(EmulationType_E1000)))
    fieldNameMap["type"] = "Type_"
    fields["upt_compatibility_enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["upt_compatibility_enabled"] = "UptCompatibilityEnabled"
    fields["mac_type"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.ethernet.mac_address_type", reflect.TypeOf(MacAddressType(MacAddressType_MANUAL)))
    fieldNameMap["mac_type"] = "MacType"
    fields["mac_address"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["mac_address"] = "MacAddress"
    fields["pci_slot_number"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["pci_slot_number"] = "PciSlotNumber"
    fields["wake_on_lan_enabled"] = bindings.NewBooleanType()
    fieldNameMap["wake_on_lan_enabled"] = "WakeOnLanEnabled"
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
            "VMXNET3": []bindings.FieldData {
                 bindings.NewFieldData("upt_compatibility_enabled", true),
            },
            "E1000": []bindings.FieldData {},
            "E1000E": []bindings.FieldData {},
            "PCNET32": []bindings.FieldData {},
            "VMXNET": []bindings.FieldData {},
            "VMXNET2": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.ethernet.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func CreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.hardware.ethernet.emulation_type", reflect.TypeOf(EmulationType(EmulationType_E1000))))
    fieldNameMap["type"] = "Type_"
    fields["upt_compatibility_enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["upt_compatibility_enabled"] = "UptCompatibilityEnabled"
    fields["mac_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.hardware.ethernet.mac_address_type", reflect.TypeOf(MacAddressType(MacAddressType_MANUAL))))
    fieldNameMap["mac_type"] = "MacType"
    fields["mac_address"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["mac_address"] = "MacAddress"
    fields["pci_slot_number"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["pci_slot_number"] = "PciSlotNumber"
    fields["wake_on_lan_enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["wake_on_lan_enabled"] = "WakeOnLanEnabled"
    fields["backing"] = bindings.NewOptionalType(bindings.NewReferenceType(BackingSpecBindingType))
    fieldNameMap["backing"] = "Backing"
    fields["start_connected"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["start_connected"] = "StartConnected"
    fields["allow_guest_control"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["allow_guest_control"] = "AllowGuestControl"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("type",
        map[string][]bindings.FieldData {
            "VMXNET3": []bindings.FieldData {
                 bindings.NewFieldData("upt_compatibility_enabled", false),
            },
            "E1000": []bindings.FieldData {},
            "E1000E": []bindings.FieldData {},
            "PCNET32": []bindings.FieldData {},
            "VMXNET": []bindings.FieldData {},
            "VMXNET2": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    uv2 := bindings.NewUnionValidator("mac_type",
        map[string][]bindings.FieldData {
            "MANUAL": []bindings.FieldData {
                 bindings.NewFieldData("mac_address", true),
            },
            "GENERATED": []bindings.FieldData {},
            "ASSIGNED": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv2)
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.ethernet.create_spec",fields, reflect.TypeOf(CreateSpec{}), fieldNameMap, validators)
}

func UpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["upt_compatibility_enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["upt_compatibility_enabled"] = "UptCompatibilityEnabled"
    fields["mac_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.hardware.ethernet.mac_address_type", reflect.TypeOf(MacAddressType(MacAddressType_MANUAL))))
    fieldNameMap["mac_type"] = "MacType"
    fields["mac_address"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["mac_address"] = "MacAddress"
    fields["wake_on_lan_enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["wake_on_lan_enabled"] = "WakeOnLanEnabled"
    fields["backing"] = bindings.NewOptionalType(bindings.NewReferenceType(BackingSpecBindingType))
    fieldNameMap["backing"] = "Backing"
    fields["start_connected"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["start_connected"] = "StartConnected"
    fields["allow_guest_control"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["allow_guest_control"] = "AllowGuestControl"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.ethernet.update_spec",fields, reflect.TypeOf(UpdateSpec{}), fieldNameMap, validators)
}

func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["nic"] = bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Ethernet"}, "")
    fieldNameMap["nic"] = "Nic"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.ethernet.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}


