/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Device.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package boot

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``Type`` enumeration class defines the valid device types that may be used as bootable devices.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type DeviceType string

const (
    // Virtual CD-ROM device.
	DeviceType_CDROM DeviceType = "CDROM"
    // Virtual disk device.
	DeviceType_DISK DeviceType = "DISK"
    // Virtual Ethernet adapter.
	DeviceType_ETHERNET DeviceType = "ETHERNET"
    // Virtual floppy drive.
	DeviceType_FLOPPY DeviceType = "FLOPPY"
)

func (t DeviceType) DeviceType() bool {
	switch t {
	case DeviceType_CDROM:
		return true
	case DeviceType_DISK:
		return true
	case DeviceType_ETHERNET:
		return true
	case DeviceType_FLOPPY:
		return true
	default:
		return false
	}
}


// The class ``EntryCreateSpec`` specifies a list of bootable virtual device classes. When a VM is being created and a array of ``EntryCreateSpec`` is specified, the boot order of the specific device instances are not specified in this class. The boot order of the specific device instance will be the order in which the Ethernet and Disk devices appear in the ``nics`` and ``disks`` respectively.
type DeviceEntryCreateSpec struct {
    // Virtual Boot device type.
	Type_ DeviceType
}

// The ``Entry`` class specifies a bootable virtual device class or specific bootable virtual device(s).
type DeviceEntry struct {
    // Virtual device type.
	Type_ DeviceType
    // Virtual Ethernet device. Ethernet device to use as boot device for this entry.
	Nic *string
    // Virtual disk device. List of virtual disks in boot order.
	Disks []string
}



func deviceGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fieldNameMap["vm"] = "Vm"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func deviceGetOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(DeviceEntryBindingType), reflect.TypeOf([]DeviceEntry{}))
}

func deviceGetRestMetadata() protocol.OperationRestMetadata {
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

func deviceSetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["devices"] = bindings.NewListType(bindings.NewReferenceType(DeviceEntryBindingType), reflect.TypeOf([]DeviceEntry{}))
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["devices"] = "Devices"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func deviceSetOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func deviceSetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["devices"] = bindings.NewListType(bindings.NewReferenceType(DeviceEntryBindingType), reflect.TypeOf([]DeviceEntry{}))
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["devices"] = "Devices"
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
		map[string]int{"Error": 500,"NotFound": 404,"InvalidArgument": 400,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func DeviceEntryCreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.boot.device.type", reflect.TypeOf(DeviceType(DeviceType_CDROM)))
	fieldNameMap["type"] = "Type_"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.boot.device.entry_create_spec", fields, reflect.TypeOf(DeviceEntryCreateSpec{}), fieldNameMap, validators)
}

func DeviceEntryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.boot.device.type", reflect.TypeOf(DeviceType(DeviceType_CDROM)))
	fieldNameMap["type"] = "Type_"
	fields["nic"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Ethernet"}, ""))
	fieldNameMap["nic"] = "Nic"
	fields["disks"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Disk"}, ""), reflect.TypeOf([]string{})))
	fieldNameMap["disks"] = "Disks"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
		map[string][]bindings.FieldData{
			"ETHERNET": []bindings.FieldData{
				bindings.NewFieldData("nic", true),
			},
			"DISK": []bindings.FieldData{
				bindings.NewFieldData("disks", true),
			},
			"CDROM": []bindings.FieldData{},
			"FLOPPY": []bindings.FieldData{},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.boot.device.entry", fields, reflect.TypeOf(DeviceEntry{}), fieldNameMap, validators)
}

