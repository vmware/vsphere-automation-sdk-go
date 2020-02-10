/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Floppy.
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

// Resource type for the virtual floppy drive device.
const Floppy_RESOURCE_TYPE = "com.vmware.vcenter.vm.hardware.Floppy"


// The ``BackingType`` enumeration class defines the valid backing types for a virtual floppy drive.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type FloppyBackingType string

const (
    // Virtual floppy drive is backed by an image file.
	FloppyBackingType_IMAGE_FILE FloppyBackingType = "IMAGE_FILE"
    // Virtual floppy drive is backed by a device on the host where the virtual machine is running.
	FloppyBackingType_HOST_DEVICE FloppyBackingType = "HOST_DEVICE"
    // Virtual floppy drive is backed by a device on the client that is connected to the virtual machine console.
	FloppyBackingType_CLIENT_DEVICE FloppyBackingType = "CLIENT_DEVICE"
)

func (b FloppyBackingType) FloppyBackingType() bool {
	switch b {
	case FloppyBackingType_IMAGE_FILE:
		return true
	case FloppyBackingType_HOST_DEVICE:
		return true
	case FloppyBackingType_CLIENT_DEVICE:
		return true
	default:
		return false
	}
}


// The ``BackingInfo`` class contains information about the physical resource backing a virtual floppy drive.
type FloppyBackingInfo struct {
    // Backing type for the virtual floppy drive.
	Type_ FloppyBackingType
    // Path of the image file backing the virtual floppy drive.
	ImageFile *string
    // Name of the host device backing the virtual floppy drive. 
	HostDevice *string
    // Flag indicating whether the virtual floppy drive is configured to automatically detect a suitable host device.
	AutoDetect *bool
}

// The ``BackingSpec`` class provides a specification of the physical resource backing a virtual floppy drive.
type FloppyBackingSpec struct {
    // Backing type for the virtual floppy drive.
	Type_ FloppyBackingType
    // Path of the image file that should be used as the virtual floppy drive backing.
	ImageFile *string
    // Name of the device that should be used as the virtual floppy drive backing.
	HostDevice *string
}

// The ``Info`` class contains information about a virtual floppy drive.
type FloppyInfo struct {
    // Device label.
	Label string
    // Physical resource backing for the virtual floppy drive.
	Backing FloppyBackingInfo
    // Connection status of the virtual device.
	State ConnectionState
    // Flag indicating whether the virtual device should be connected whenever the virtual machine is powered on.
	StartConnected bool
    // Flag indicating whether the guest can connect and disconnect the device.
	AllowGuestControl bool
}

// The ``CreateSpec`` class provides a specification for the configuration of a newly-created virtual floppy drive.
type FloppyCreateSpec struct {
    // Physical resource backing for the virtual floppy drive.
	Backing *FloppyBackingSpec
    // Flag indicating whether the virtual device should be connected whenever the virtual machine is powered on.
	StartConnected *bool
    // Flag indicating whether the guest can connect and disconnect the device.
	AllowGuestControl *bool
}

// The ``UpdateSpec`` class describes the updates to be made to the configuration of a virtual floppy drive.
type FloppyUpdateSpec struct {
    // Physical resource backing for the virtual floppy drive. 
    //
    //  This property may only be modified if the virtual machine is not powered on or the virtual floppy drive is not connected.
	Backing *FloppyBackingSpec
    // Flag indicating whether the virtual device should be connected whenever the virtual machine is powered on.
	StartConnected *bool
    // Flag indicating whether the guest can connect and disconnect the device.
	AllowGuestControl *bool
}

// The ``Summary`` class contains commonly used information about a virtual floppy drive.
type FloppySummary struct {
    // Identifier of the virtual floppy drive.
	Floppy string
}



func floppyListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fieldNameMap["vm"] = "Vm"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func floppyListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(FloppySummaryBindingType), reflect.TypeOf([]FloppySummary{}))
}

func floppyListRestMetadata() protocol.OperationRestMetadata {
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

func floppyGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["floppy"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Floppy"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["floppy"] = "Floppy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func floppyGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(FloppyInfoBindingType)
}

func floppyGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["floppy"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Floppy"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["floppy"] = "Floppy"
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

func floppyCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["spec"] = bindings.NewReferenceType(FloppyCreateSpecBindingType)
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func floppyCreateOutputType() bindings.BindingType {
	return bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Floppy"}, "")
}

func floppyCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["spec"] = bindings.NewReferenceType(FloppyCreateSpecBindingType)
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
		map[string]int{"Error": 500,"NotFound": 404,"NotAllowedInCurrentState": 400,"UnableToAllocateResource": 500,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func floppyUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["floppy"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Floppy"}, "")
	fields["spec"] = bindings.NewReferenceType(FloppyUpdateSpecBindingType)
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["floppy"] = "Floppy"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func floppyUpdateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func floppyUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["floppy"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Floppy"}, "")
	fields["spec"] = bindings.NewReferenceType(FloppyUpdateSpecBindingType)
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["floppy"] = "Floppy"
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

func floppyDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["floppy"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Floppy"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["floppy"] = "Floppy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func floppyDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func floppyDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["floppy"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Floppy"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["floppy"] = "Floppy"
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

func floppyConnectInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["floppy"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Floppy"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["floppy"] = "Floppy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func floppyConnectOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func floppyConnectRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["floppy"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Floppy"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["floppy"] = "Floppy"
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

func floppyDisconnectInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["floppy"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Floppy"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["floppy"] = "Floppy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func floppyDisconnectOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func floppyDisconnectRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["floppy"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Floppy"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["floppy"] = "Floppy"
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


func FloppyBackingInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.floppy.backing_type", reflect.TypeOf(FloppyBackingType(FloppyBackingType_IMAGE_FILE)))
	fieldNameMap["type"] = "Type_"
	fields["image_file"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["image_file"] = "ImageFile"
	fields["host_device"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["host_device"] = "HostDevice"
	fields["auto_detect"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["auto_detect"] = "AutoDetect"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
		map[string][]bindings.FieldData{
			"IMAGE_FILE": []bindings.FieldData{
				bindings.NewFieldData("image_file", true),
			},
			"HOST_DEVICE": []bindings.FieldData{
				bindings.NewFieldData("host_device", false),
				bindings.NewFieldData("auto_detect", true),
			},
			"CLIENT_DEVICE": []bindings.FieldData{},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.floppy.backing_info", fields, reflect.TypeOf(FloppyBackingInfo{}), fieldNameMap, validators)
}

func FloppyBackingSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.floppy.backing_type", reflect.TypeOf(FloppyBackingType(FloppyBackingType_IMAGE_FILE)))
	fieldNameMap["type"] = "Type_"
	fields["image_file"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["image_file"] = "ImageFile"
	fields["host_device"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["host_device"] = "HostDevice"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
		map[string][]bindings.FieldData{
			"IMAGE_FILE": []bindings.FieldData{
				bindings.NewFieldData("image_file", true),
			},
			"HOST_DEVICE": []bindings.FieldData{
				bindings.NewFieldData("host_device", false),
			},
			"CLIENT_DEVICE": []bindings.FieldData{},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.floppy.backing_spec", fields, reflect.TypeOf(FloppyBackingSpec{}), fieldNameMap, validators)
}

func FloppyInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["label"] = bindings.NewStringType()
	fieldNameMap["label"] = "Label"
	fields["backing"] = bindings.NewReferenceType(FloppyBackingInfoBindingType)
	fieldNameMap["backing"] = "Backing"
	fields["state"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.connection_state", reflect.TypeOf(ConnectionState(ConnectionState_CONNECTED)))
	fieldNameMap["state"] = "State"
	fields["start_connected"] = bindings.NewBooleanType()
	fieldNameMap["start_connected"] = "StartConnected"
	fields["allow_guest_control"] = bindings.NewBooleanType()
	fieldNameMap["allow_guest_control"] = "AllowGuestControl"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.floppy.info", fields, reflect.TypeOf(FloppyInfo{}), fieldNameMap, validators)
}

func FloppyCreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["backing"] = bindings.NewOptionalType(bindings.NewReferenceType(FloppyBackingSpecBindingType))
	fieldNameMap["backing"] = "Backing"
	fields["start_connected"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["start_connected"] = "StartConnected"
	fields["allow_guest_control"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["allow_guest_control"] = "AllowGuestControl"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.floppy.create_spec", fields, reflect.TypeOf(FloppyCreateSpec{}), fieldNameMap, validators)
}

func FloppyUpdateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["backing"] = bindings.NewOptionalType(bindings.NewReferenceType(FloppyBackingSpecBindingType))
	fieldNameMap["backing"] = "Backing"
	fields["start_connected"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["start_connected"] = "StartConnected"
	fields["allow_guest_control"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["allow_guest_control"] = "AllowGuestControl"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.floppy.update_spec", fields, reflect.TypeOf(FloppyUpdateSpec{}), fieldNameMap, validators)
}

func FloppySummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["floppy"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Floppy"}, "")
	fieldNameMap["floppy"] = "Floppy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.floppy.summary", fields, reflect.TypeOf(FloppySummary{}), fieldNameMap, validators)
}

