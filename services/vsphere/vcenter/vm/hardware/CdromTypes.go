/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Cdrom.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package hardware

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// Resource type for the virtual CD-ROM device.
const Cdrom_RESOURCE_TYPE = "com.vmware.vcenter.vm.hardware.Cdrom"


// The ``HostBusAdapterType`` enumeration class defines the valid types of host bus adapters that may be used for attaching a Cdrom to a virtual machine.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type CdromHostBusAdapterType string

const (
    // Cdrom is attached to an IDE adapter.
	CdromHostBusAdapterType_IDE CdromHostBusAdapterType = "IDE"
    // Cdrom is attached to a SATA adapter.
	CdromHostBusAdapterType_SATA CdromHostBusAdapterType = "SATA"
)

func (h CdromHostBusAdapterType) CdromHostBusAdapterType() bool {
	switch h {
	case CdromHostBusAdapterType_IDE:
		return true
	case CdromHostBusAdapterType_SATA:
		return true
	default:
		return false
	}
}


// The ``BackingType`` enumeration class defines the valid backing types for a virtual CD-ROM device.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type CdromBackingType string

const (
    // Virtual CD-ROM device is backed by an ISO file.
	CdromBackingType_ISO_FILE CdromBackingType = "ISO_FILE"
    // Virtual CD-ROM device is backed by a device on the host where the virtual machine is running.
	CdromBackingType_HOST_DEVICE CdromBackingType = "HOST_DEVICE"
    // Virtual CD-ROM device is backed by a device on the client that is connected to the virtual machine console.
	CdromBackingType_CLIENT_DEVICE CdromBackingType = "CLIENT_DEVICE"
)

func (b CdromBackingType) CdromBackingType() bool {
	switch b {
	case CdromBackingType_ISO_FILE:
		return true
	case CdromBackingType_HOST_DEVICE:
		return true
	case CdromBackingType_CLIENT_DEVICE:
		return true
	default:
		return false
	}
}


// The ``DeviceAccessType`` enumeration class defines the valid device access types for a physical device packing of a virtual CD-ROM device.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type CdromDeviceAccessType string

const (
    // ATAPI or SCSI device emulation.
	CdromDeviceAccessType_EMULATION CdromDeviceAccessType = "EMULATION"
    // Raw passthru device access.
	CdromDeviceAccessType_PASSTHRU CdromDeviceAccessType = "PASSTHRU"
    // Raw passthru device access, with exclusive access to the device.
	CdromDeviceAccessType_PASSTHRU_EXCLUSIVE CdromDeviceAccessType = "PASSTHRU_EXCLUSIVE"
)

func (d CdromDeviceAccessType) CdromDeviceAccessType() bool {
	switch d {
	case CdromDeviceAccessType_EMULATION:
		return true
	case CdromDeviceAccessType_PASSTHRU:
		return true
	case CdromDeviceAccessType_PASSTHRU_EXCLUSIVE:
		return true
	default:
		return false
	}
}


// The ``BackingInfo`` class contains information about the physical resource backing a virtual CD-ROM device.
type CdromBackingInfo struct {
    // Backing type for the virtual CD-ROM device.
	Type_ CdromBackingType
    // Path of the image file backing the virtual CD-ROM device.
	IsoFile *string
    // Name of the host device backing the virtual CD-ROM device. 
	HostDevice *string
    // Flag indicating whether the virtual CD-ROM device is configured to automatically detect a suitable host device.
	AutoDetect *bool
    // Access type for the device backing.
	DeviceAccessType *CdromDeviceAccessType
}

// The ``BackingSpec`` class provides a specification of the physical resource backing a virtual CD-ROM device.
type CdromBackingSpec struct {
    // Backing type for the virtual CD-ROM device.
	Type_ CdromBackingType
    // Path of the image file that should be used as the virtual CD-ROM device backing.
	IsoFile *string
    // Name of the device that should be used as the virtual CD-ROM device backing.
	HostDevice *string
    // Access type for the device backing.
	DeviceAccessType *CdromDeviceAccessType
}

// The ``Info`` class contains information about a virtual CD-ROM device.
type CdromInfo struct {
    // Type of host bus adapter to which the device is attached.
	Type_ CdromHostBusAdapterType
    // Device label.
	Label string
    // Address of device attached to a virtual IDE adapter.
	Ide *IdeAddressInfo
    // Address of device attached to a virtual SATA adapter.
	Sata *SataAddressInfo
    // Physical resource backing for the virtual CD-ROM device.
	Backing CdromBackingInfo
    // Connection status of the virtual device.
	State ConnectionState
    // Flag indicating whether the virtual device should be connected whenever the virtual machine is powered on.
	StartConnected bool
    // Flag indicating whether the guest can connect and disconnect the device.
	AllowGuestControl bool
}

// The ``CreateSpec`` class provides a specification for the configuration of a newly-created virtual CD-ROM device.
type CdromCreateSpec struct {
    // Type of host bus adapter to which the device should be attached.
	Type_ *CdromHostBusAdapterType
    // Address for attaching the device to a virtual IDE adapter.
	Ide *IdeAddressSpec
    // Address for attaching the device to a virtual SATA adapter.
	Sata *SataAddressSpec
    // Physical resource backing for the virtual CD-ROM device.
	Backing *CdromBackingSpec
    // Flag indicating whether the virtual device should be connected whenever the virtual machine is powered on.
	StartConnected *bool
    // Flag indicating whether the guest can connect and disconnect the device.
	AllowGuestControl *bool
}

// The ``UpdateSpec`` class describes the updates to be made to the configuration of a virtual CD-ROM device.
type CdromUpdateSpec struct {
    // Physical resource backing for the virtual CD-ROM device. 
    //
    //  This property may only be modified if the virtual machine is not powered on or the virtual CD-ROM device is not connected.
	Backing *CdromBackingSpec
    // Flag indicating whether the virtual device should be connected whenever the virtual machine is powered on.
	StartConnected *bool
    // Flag indicating whether the guest can connect and disconnect the device.
	AllowGuestControl *bool
}

// The ``Summary`` class contains commonly used information about a virtual CD-ROM device.
type CdromSummary struct {
    // Identifier of the virtual CD-ROM device.
	Cdrom string
}



func cdromListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fieldNameMap["vm"] = "Vm"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cdromListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(CdromSummaryBindingType), reflect.TypeOf([]CdromSummary{}))
}

func cdromListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fieldNameMap["vm"] = "Vm"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
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
		map[string]int{"Error": 500,"NotFound": 404,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func cdromGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["cdrom"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Cdrom"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["cdrom"] = "Cdrom"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cdromGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(CdromInfoBindingType)
}

func cdromGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["cdrom"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Cdrom"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["cdrom"] = "Cdrom"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
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
		map[string]int{"Error": 500,"NotFound": 404,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func cdromCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["spec"] = bindings.NewReferenceType(CdromCreateSpecBindingType)
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cdromCreateOutputType() bindings.BindingType {
	return bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Cdrom"}, "")
}

func cdromCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["spec"] = bindings.NewReferenceType(CdromCreateSpecBindingType)
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["spec"] = "Spec"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
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
		map[string]int{"Error": 500,"NotFound": 404,"NotAllowedInCurrentState": 400,"UnableToAllocateResource": 500,"ResourceInUse": 400,"InvalidArgument": 400,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403,"Unsupported": 400})
}

func cdromUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["cdrom"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Cdrom"}, "")
	fields["spec"] = bindings.NewReferenceType(CdromUpdateSpecBindingType)
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["cdrom"] = "Cdrom"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cdromUpdateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func cdromUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["cdrom"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Cdrom"}, "")
	fields["spec"] = bindings.NewReferenceType(CdromUpdateSpecBindingType)
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["cdrom"] = "Cdrom"
	fieldNameMap["spec"] = "Spec"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
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
		map[string]int{"Error": 500,"NotFound": 404,"NotAllowedInCurrentState": 400,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func cdromDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["cdrom"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Cdrom"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["cdrom"] = "Cdrom"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cdromDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func cdromDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["cdrom"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Cdrom"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["cdrom"] = "Cdrom"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
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
		map[string]int{"Error": 500,"NotFound": 404,"NotAllowedInCurrentState": 400,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func cdromConnectInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["cdrom"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Cdrom"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["cdrom"] = "Cdrom"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cdromConnectOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func cdromConnectRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["cdrom"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Cdrom"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["cdrom"] = "Cdrom"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
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
		map[string]int{"Error": 500,"NotFound": 404,"AlreadyInDesiredState": 400,"NotAllowedInCurrentState": 400,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func cdromDisconnectInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["cdrom"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Cdrom"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["cdrom"] = "Cdrom"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cdromDisconnectOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func cdromDisconnectRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["cdrom"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Cdrom"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["cdrom"] = "Cdrom"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
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
		map[string]int{"Error": 500,"NotFound": 404,"AlreadyInDesiredState": 400,"NotAllowedInCurrentState": 400,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func CdromBackingInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.cdrom.backing_type", reflect.TypeOf(CdromBackingType(CdromBackingType_ISO_FILE)))
	fieldNameMap["type"] = "Type_"
	fields["iso_file"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["iso_file"] = "IsoFile"
	fields["host_device"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["host_device"] = "HostDevice"
	fields["auto_detect"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["auto_detect"] = "AutoDetect"
	fields["device_access_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.hardware.cdrom.device_access_type", reflect.TypeOf(CdromDeviceAccessType(CdromDeviceAccessType_EMULATION))))
	fieldNameMap["device_access_type"] = "DeviceAccessType"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
		map[string][]bindings.FieldData{
			"ISO_FILE": []bindings.FieldData{
				bindings.NewFieldData("iso_file", true),
			},
			"HOST_DEVICE": []bindings.FieldData{
				bindings.NewFieldData("host_device", false),
				bindings.NewFieldData("auto_detect", true),
				bindings.NewFieldData("device_access_type", true),
			},
			"CLIENT_DEVICE": []bindings.FieldData{
				bindings.NewFieldData("device_access_type", true),
			},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.cdrom.backing_info", fields, reflect.TypeOf(CdromBackingInfo{}), fieldNameMap, validators)
}

func CdromBackingSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.cdrom.backing_type", reflect.TypeOf(CdromBackingType(CdromBackingType_ISO_FILE)))
	fieldNameMap["type"] = "Type_"
	fields["iso_file"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["iso_file"] = "IsoFile"
	fields["host_device"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["host_device"] = "HostDevice"
	fields["device_access_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.hardware.cdrom.device_access_type", reflect.TypeOf(CdromDeviceAccessType(CdromDeviceAccessType_EMULATION))))
	fieldNameMap["device_access_type"] = "DeviceAccessType"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
		map[string][]bindings.FieldData{
			"ISO_FILE": []bindings.FieldData{
				bindings.NewFieldData("iso_file", true),
			},
			"HOST_DEVICE": []bindings.FieldData{
				bindings.NewFieldData("host_device", false),
				bindings.NewFieldData("device_access_type", false),
			},
			"CLIENT_DEVICE": []bindings.FieldData{
				bindings.NewFieldData("device_access_type", false),
			},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.cdrom.backing_spec", fields, reflect.TypeOf(CdromBackingSpec{}), fieldNameMap, validators)
}

func CdromInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.cdrom.host_bus_adapter_type", reflect.TypeOf(CdromHostBusAdapterType(CdromHostBusAdapterType_IDE)))
	fieldNameMap["type"] = "Type_"
	fields["label"] = bindings.NewStringType()
	fieldNameMap["label"] = "Label"
	fields["ide"] = bindings.NewOptionalType(bindings.NewReferenceType(IdeAddressInfoBindingType))
	fieldNameMap["ide"] = "Ide"
	fields["sata"] = bindings.NewOptionalType(bindings.NewReferenceType(SataAddressInfoBindingType))
	fieldNameMap["sata"] = "Sata"
	fields["backing"] = bindings.NewReferenceType(CdromBackingInfoBindingType)
	fieldNameMap["backing"] = "Backing"
	fields["state"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.connection_state", reflect.TypeOf(ConnectionState(ConnectionState_CONNECTED)))
	fieldNameMap["state"] = "State"
	fields["start_connected"] = bindings.NewBooleanType()
	fieldNameMap["start_connected"] = "StartConnected"
	fields["allow_guest_control"] = bindings.NewBooleanType()
	fieldNameMap["allow_guest_control"] = "AllowGuestControl"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
		map[string][]bindings.FieldData{
			"IDE": []bindings.FieldData{
				bindings.NewFieldData("ide", true),
			},
			"SATA": []bindings.FieldData{
				bindings.NewFieldData("sata", true),
			},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.cdrom.info", fields, reflect.TypeOf(CdromInfo{}), fieldNameMap, validators)
}

func CdromCreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.hardware.cdrom.host_bus_adapter_type", reflect.TypeOf(CdromHostBusAdapterType(CdromHostBusAdapterType_IDE))))
	fieldNameMap["type"] = "Type_"
	fields["ide"] = bindings.NewOptionalType(bindings.NewReferenceType(IdeAddressSpecBindingType))
	fieldNameMap["ide"] = "Ide"
	fields["sata"] = bindings.NewOptionalType(bindings.NewReferenceType(SataAddressSpecBindingType))
	fieldNameMap["sata"] = "Sata"
	fields["backing"] = bindings.NewOptionalType(bindings.NewReferenceType(CdromBackingSpecBindingType))
	fieldNameMap["backing"] = "Backing"
	fields["start_connected"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["start_connected"] = "StartConnected"
	fields["allow_guest_control"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["allow_guest_control"] = "AllowGuestControl"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
		map[string][]bindings.FieldData{
			"IDE": []bindings.FieldData{
				bindings.NewFieldData("ide", false),
			},
			"SATA": []bindings.FieldData{
				bindings.NewFieldData("sata", false),
			},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.cdrom.create_spec", fields, reflect.TypeOf(CdromCreateSpec{}), fieldNameMap, validators)
}

func CdromUpdateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["backing"] = bindings.NewOptionalType(bindings.NewReferenceType(CdromBackingSpecBindingType))
	fieldNameMap["backing"] = "Backing"
	fields["start_connected"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["start_connected"] = "StartConnected"
	fields["allow_guest_control"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["allow_guest_control"] = "AllowGuestControl"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.cdrom.update_spec", fields, reflect.TypeOf(CdromUpdateSpec{}), fieldNameMap, validators)
}

func CdromSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cdrom"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Cdrom"}, "")
	fieldNameMap["cdrom"] = "Cdrom"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.cdrom.summary", fields, reflect.TypeOf(CdromSummary{}), fieldNameMap, validators)
}

