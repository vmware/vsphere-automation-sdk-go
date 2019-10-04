/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for package: com.vmware.vcenter.vm.hardware.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package hardware

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
)


// The ``ConnectionState`` enumeration class defines the valid states for a removable device that is configured to be connected.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.

type ConnectionState string

const (
    // The device is connected and working correctly.
     ConnectionState_CONNECTED ConnectionState = "CONNECTED"
    // Device connection failed due to a recoverable error; for example, the virtual device backing is currently in use by another virtual machine.
     ConnectionState_RECOVERABLE_ERROR ConnectionState = "RECOVERABLE_ERROR"
    // Device connection failed due to an unrecoverable error; for example, the virtual device backing does not exist.
     ConnectionState_UNRECOVERABLE_ERROR ConnectionState = "UNRECOVERABLE_ERROR"
    // The device is not connected.
     ConnectionState_NOT_CONNECTED ConnectionState = "NOT_CONNECTED"
    // The device status is unknown.
     ConnectionState_UNKNOWN ConnectionState = "UNKNOWN"
)

func (c ConnectionState) ConnectionState() bool {
    switch c {
        case ConnectionState_CONNECTED:
            return true
        case ConnectionState_RECOVERABLE_ERROR:
            return true
        case ConnectionState_UNRECOVERABLE_ERROR:
            return true
        case ConnectionState_NOT_CONNECTED:
            return true
        case ConnectionState_UNKNOWN:
            return true
        default:
            return false
    }
}





// The ``IdeAddressInfo`` class contains information about the address of a virtual device that is attached to a virtual IDE adapter of a virtual machine.
type IdeAddressInfo struct {
    Primary bool
    Master bool
}






// The ``ScsiAddressInfo`` class contains information about the address of a virtual device that is attached to a virtual SCSI adapter of a virtual machine.
type ScsiAddressInfo struct {
    Bus int64
    Unit int64
}






// The ``SataAddressInfo`` class contains information about the address of a virtual device that is attached to a virtual SATA adapter of a virtual machine.
type SataAddressInfo struct {
    Bus int64
    Unit int64
}






// The ``NvmeAddressInfo`` class contains information about the address of a virtual device that is attached to a virtual NVMe adapter of a virtual machine. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type NvmeAddressInfo struct {
    Bus int64
    Unit int64
}






// The ``IdeAddressSpec`` class contains information for specifying the address of a virtual device that is attached to a virtual IDE adapter of a virtual machine.
type IdeAddressSpec struct {
    Primary *bool
    Master *bool
}






// The ``ScsiAddressSpec`` class contains information for specifying the address of a virtual device that is attached to a virtual SCSI adapter of a virtual machine.
type ScsiAddressSpec struct {
    Bus int64
    Unit *int64
}






// The ``SataAddressSpec`` class contains information for specifying the address of a virtual device that is attached to a virtual SATA adapter of a virtual machine.
type SataAddressSpec struct {
    Bus int64
    Unit *int64
}






// The ``NvmeAddressSpec`` class contains information for specifying the address of a virtual device that is attached to a virtual NVMe adapter of a virtual machine. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type NvmeAddressSpec struct {
    Bus int64
    Unit *int64
}






// The ``ConnectionInfo`` class provides information about the state and configuration of a removable virtual device.
type ConnectionInfo struct {
    State ConnectionState
    StartConnected bool
    AllowGuestControl bool
}






// The ``ConnectionCreateSpec`` class provides a specification for the configuration of a newly-created removable device.
type ConnectionCreateSpec struct {
    StartConnected *bool
    AllowGuestControl *bool
}






// The ``ConnectionUpdateSpec`` class describes the updates to be made to the configuration of a removable virtual device.
type ConnectionUpdateSpec struct {
    StartConnected *bool
    AllowGuestControl *bool
}










func IdeAddressInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["primary"] = bindings.NewBooleanType()
    fieldNameMap["primary"] = "Primary"
    fields["master"] = bindings.NewBooleanType()
    fieldNameMap["master"] = "Master"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.ide_address_info",fields, reflect.TypeOf(IdeAddressInfo{}), fieldNameMap, validators)
}

func ScsiAddressInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["bus"] = bindings.NewIntegerType()
    fieldNameMap["bus"] = "Bus"
    fields["unit"] = bindings.NewIntegerType()
    fieldNameMap["unit"] = "Unit"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.scsi_address_info",fields, reflect.TypeOf(ScsiAddressInfo{}), fieldNameMap, validators)
}

func SataAddressInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["bus"] = bindings.NewIntegerType()
    fieldNameMap["bus"] = "Bus"
    fields["unit"] = bindings.NewIntegerType()
    fieldNameMap["unit"] = "Unit"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.sata_address_info",fields, reflect.TypeOf(SataAddressInfo{}), fieldNameMap, validators)
}

func NvmeAddressInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["bus"] = bindings.NewIntegerType()
    fieldNameMap["bus"] = "Bus"
    fields["unit"] = bindings.NewIntegerType()
    fieldNameMap["unit"] = "Unit"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.nvme_address_info",fields, reflect.TypeOf(NvmeAddressInfo{}), fieldNameMap, validators)
}

func IdeAddressSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["primary"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["primary"] = "Primary"
    fields["master"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["master"] = "Master"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.ide_address_spec",fields, reflect.TypeOf(IdeAddressSpec{}), fieldNameMap, validators)
}

func ScsiAddressSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["bus"] = bindings.NewIntegerType()
    fieldNameMap["bus"] = "Bus"
    fields["unit"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["unit"] = "Unit"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.scsi_address_spec",fields, reflect.TypeOf(ScsiAddressSpec{}), fieldNameMap, validators)
}

func SataAddressSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["bus"] = bindings.NewIntegerType()
    fieldNameMap["bus"] = "Bus"
    fields["unit"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["unit"] = "Unit"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.sata_address_spec",fields, reflect.TypeOf(SataAddressSpec{}), fieldNameMap, validators)
}

func NvmeAddressSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["bus"] = bindings.NewIntegerType()
    fieldNameMap["bus"] = "Bus"
    fields["unit"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["unit"] = "Unit"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.nvme_address_spec",fields, reflect.TypeOf(NvmeAddressSpec{}), fieldNameMap, validators)
}

func ConnectionInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["state"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.connection_state", reflect.TypeOf(ConnectionState(ConnectionState_CONNECTED)))
    fieldNameMap["state"] = "State"
    fields["start_connected"] = bindings.NewBooleanType()
    fieldNameMap["start_connected"] = "StartConnected"
    fields["allow_guest_control"] = bindings.NewBooleanType()
    fieldNameMap["allow_guest_control"] = "AllowGuestControl"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.connection_info",fields, reflect.TypeOf(ConnectionInfo{}), fieldNameMap, validators)
}

func ConnectionCreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["start_connected"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["start_connected"] = "StartConnected"
    fields["allow_guest_control"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["allow_guest_control"] = "AllowGuestControl"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.connection_create_spec",fields, reflect.TypeOf(ConnectionCreateSpec{}), fieldNameMap, validators)
}

func ConnectionUpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["start_connected"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["start_connected"] = "StartConnected"
    fields["allow_guest_control"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["allow_guest_control"] = "AllowGuestControl"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.connection_update_spec",fields, reflect.TypeOf(ConnectionUpdateSpec{}), fieldNameMap, validators)
}


