/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Disk.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package hardware

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/vsphere/vcenter/vm/hardware/Disk"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// Resource type for the virtual disk.
const Disk_RESOURCE_TYPE = "com.vmware.vcenter.vm.hardware.Disk"


// The ``HostBusAdapterType`` enumeration class defines the valid types of host bus adapters that may be used for attaching a virtual storage device to a virtual machine.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type Disk.DiskHostBusAdapterType string

const (
    // Disk is attached to an IDE adapter.
	DiskHostBusAdapterType_IDE Disk.DiskHostBusAdapterType = "IDE"
    // Disk is attached to a SCSI adapter.
	DiskHostBusAdapterType_SCSI Disk.DiskHostBusAdapterType = "SCSI"
    // Disk is attached to a SATA adapter.
	DiskHostBusAdapterType_SATA Disk.DiskHostBusAdapterType = "SATA"
)

func (h Disk.DiskHostBusAdapterType) Disk.DiskHostBusAdapterType() bool {
	switch h {
	case DiskHostBusAdapterType_IDE:
		return true
	case DiskHostBusAdapterType_SCSI:
		return true
	case DiskHostBusAdapterType_SATA:
		return true
	default:
		return false
	}
}


// The ``BackingType`` enumeration class defines the valid backing types for a virtual disk.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type DiskBackingType string

const (
    // Virtual disk is backed by a VMDK file.
	DiskBackingType_VMDK_FILE DiskBackingType = "VMDK_FILE"
)

func (b DiskBackingType) DiskBackingType() bool {
	switch b {
	case DiskBackingType_VMDK_FILE:
		return true
	default:
		return false
	}
}


// The ``BackingInfo`` class contains information about the physical resource backing a virtual disk.
type DiskBackingInfo struct {
    // Backing type for the virtual disk.
	Type_ DiskBackingType
    // Path of the VMDK file backing the virtual disk.
	VmdkFile *string
}

// The ``BackingSpec`` class provides a specification of the physical resource backing a virtual disk.
type DiskBackingSpec struct {
    // Backing type for the virtual disk.
	Type_ DiskBackingType
    // Path of the VMDK file backing the virtual disk.
	VmdkFile *string
}

// The ``VmdkCreateSpec`` class provides a specification for creating a new VMDK file to be used as a backing for a virtual disk. The virtual disk will be stored in the same directory as the virtual machine's configuration file.
type DiskVmdkCreateSpec struct {
    // Base name of the VMDK file. The name should not include the '.vmdk' file extension.
	Name *string
    // Capacity of the virtual disk backing in bytes.
	Capacity *int64
    // The ``StoragePolicySpec`` class contains information about the storage policy that is to be associated the with VMDK file. This property was added in vSphere API 6.7.
	StoragePolicy *DiskStoragePolicySpec
}

// The ``Info`` class contains information about a virtual disk.
type DiskInfo struct {
    // Device label.
	Label string
    // Type of host bus adapter to which the device is attached.
	Type_ Disk.DiskHostBusAdapterType
    // Address of device attached to a virtual IDE adapter.
	Ide *IdeAddressInfo
    // Address of device attached to a virtual SCSI adapter.
	Scsi *ScsiAddressInfo
    // Address of device attached to a virtual SATA adapter.
	Sata *SataAddressInfo
    // Physical resource backing for the virtual disk.
	Backing DiskBackingInfo
    // Capacity of the virtual disk in bytes.
	Capacity *int64
}

// The ``StoragePolicySpec`` class contains information about the storage policy be associated with a VMDK file. This class was added in vSphere API 6.7.
type DiskStoragePolicySpec struct {
    // Identifier of the storage policy which should be associated with the VMDK file. This property was added in vSphere API 6.7.
	Policy string
}

// The ``CreateSpec`` class provides a specification for the configuration of a newly-created virtual disk.
type DiskCreateSpec struct {
    // Type of host bus adapter to which the device should be attached.
	Type_ *Disk.DiskHostBusAdapterType
    // Address for attaching the device to a virtual IDE adapter.
	Ide *IdeAddressSpec
    // Address for attaching the device to a virtual SCSI adapter.
	Scsi *ScsiAddressSpec
    // Address for attaching the device to a virtual SATA adapter.
	Sata *SataAddressSpec
    // Existing physical resource backing for the virtual disk. Exactly one of ``backing`` or ``newVmdk`` must be specified.
	Backing *DiskBackingSpec
    // Specification for creating a new VMDK backing for the virtual disk. Exactly one of ``backing`` or ``newVmdk`` must be specified.
	NewVmdk *DiskVmdkCreateSpec
}

// The ``UpdateSpec`` class describes the updates to be made to the configuration of a virtual disk.
type DiskUpdateSpec struct {
    // Physical resource backing for the virtual disk. 
    //
    //  This property may only be modified if the virtual machine is not powered on.
	Backing *DiskBackingSpec
}

// The ``Summary`` class contains commonly used information about a virtual disk.
type DiskSummary struct {
    // Identifier of the virtual Disk.
	Disk string
}



func diskListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fieldNameMap["vm"] = "Vm"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func diskListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(DiskSummaryBindingType), reflect.TypeOf([]DiskSummary{}))
}

func diskListRestMetadata() protocol.OperationRestMetadata {
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

func diskGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["disk"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Disk"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["disk"] = "Disk"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func diskGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(Disk.DiskInfoBindingType)
}

func diskGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["disk"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Disk"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["disk"] = "Disk"
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

func diskCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["spec"] = bindings.NewReferenceType(Disk.DiskCreateSpecBindingType)
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func diskCreateOutputType() bindings.BindingType {
	return bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Disk"}, "")
}

func diskCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["spec"] = bindings.NewReferenceType(Disk.DiskCreateSpecBindingType)
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

func diskUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["disk"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Disk"}, "")
	fields["spec"] = bindings.NewReferenceType(DiskUpdateSpecBindingType)
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["disk"] = "Disk"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func diskUpdateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func diskUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["disk"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Disk"}, "")
	fields["spec"] = bindings.NewReferenceType(DiskUpdateSpecBindingType)
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["disk"] = "Disk"
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

func diskDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["disk"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Disk"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["disk"] = "Disk"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func diskDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func diskDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["disk"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Disk"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["disk"] = "Disk"
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


func DiskBackingInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.disk.backing_type", reflect.TypeOf(DiskBackingType(DiskBackingType_VMDK_FILE)))
	fieldNameMap["type"] = "Type_"
	fields["vmdk_file"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vmdk_file"] = "VmdkFile"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
		map[string][]bindings.FieldData{
			"VMDK_FILE": []bindings.FieldData{
				bindings.NewFieldData("vmdk_file", true),
			},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.disk.backing_info", fields, reflect.TypeOf(DiskBackingInfo{}), fieldNameMap, validators)
}

func DiskBackingSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.disk.backing_type", reflect.TypeOf(DiskBackingType(DiskBackingType_VMDK_FILE)))
	fieldNameMap["type"] = "Type_"
	fields["vmdk_file"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["vmdk_file"] = "VmdkFile"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
		map[string][]bindings.FieldData{
			"VMDK_FILE": []bindings.FieldData{
				bindings.NewFieldData("vmdk_file", true),
			},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.disk.backing_spec", fields, reflect.TypeOf(DiskBackingSpec{}), fieldNameMap, validators)
}

func DiskVmdkCreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["capacity"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["capacity"] = "Capacity"
	fields["storage_policy"] = bindings.NewOptionalType(bindings.NewReferenceType(DiskStoragePolicySpecBindingType))
	fieldNameMap["storage_policy"] = "StoragePolicy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.disk.vmdk_create_spec", fields, reflect.TypeOf(DiskVmdkCreateSpec{}), fieldNameMap, validators)
}

func DiskInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["label"] = bindings.NewStringType()
	fieldNameMap["label"] = "Label"
	fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.disk.host_bus_adapter_type", reflect.TypeOf(Disk.DiskHostBusAdapterType(Disk.DiskHostBusAdapterType_IDE)))
	fieldNameMap["type"] = "Type_"
	fields["ide"] = bindings.NewOptionalType(bindings.NewReferenceType(IdeAddressInfoBindingType))
	fieldNameMap["ide"] = "Ide"
	fields["scsi"] = bindings.NewOptionalType(bindings.NewReferenceType(ScsiAddressInfoBindingType))
	fieldNameMap["scsi"] = "Scsi"
	fields["sata"] = bindings.NewOptionalType(bindings.NewReferenceType(SataAddressInfoBindingType))
	fieldNameMap["sata"] = "Sata"
	fields["backing"] = bindings.NewReferenceType(DiskBackingInfoBindingType)
	fieldNameMap["backing"] = "Backing"
	fields["capacity"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["capacity"] = "Capacity"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
		map[string][]bindings.FieldData{
			"IDE": []bindings.FieldData{
				bindings.NewFieldData("ide", true),
			},
			"SCSI": []bindings.FieldData{
				bindings.NewFieldData("scsi", true),
			},
			"SATA": []bindings.FieldData{
				bindings.NewFieldData("sata", true),
			},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.disk.info", fields, reflect.TypeOf(Disk.DiskInfo{}), fieldNameMap, validators)
}

func DiskStoragePolicySpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["policy"] = bindings.NewIdType([]string{"com.vmware.vcenter.StoragePolicy"}, "")
	fieldNameMap["policy"] = "Policy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.disk.storage_policy_spec", fields, reflect.TypeOf(DiskStoragePolicySpec{}), fieldNameMap, validators)
}

func DiskCreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.hardware.disk.host_bus_adapter_type", reflect.TypeOf(Disk.DiskHostBusAdapterType(Disk.DiskHostBusAdapterType_IDE))))
	fieldNameMap["type"] = "Type_"
	fields["ide"] = bindings.NewOptionalType(bindings.NewReferenceType(IdeAddressSpecBindingType))
	fieldNameMap["ide"] = "Ide"
	fields["scsi"] = bindings.NewOptionalType(bindings.NewReferenceType(ScsiAddressSpecBindingType))
	fieldNameMap["scsi"] = "Scsi"
	fields["sata"] = bindings.NewOptionalType(bindings.NewReferenceType(SataAddressSpecBindingType))
	fieldNameMap["sata"] = "Sata"
	fields["backing"] = bindings.NewOptionalType(bindings.NewReferenceType(DiskBackingSpecBindingType))
	fieldNameMap["backing"] = "Backing"
	fields["new_vmdk"] = bindings.NewOptionalType(bindings.NewReferenceType(DiskVmdkCreateSpecBindingType))
	fieldNameMap["new_vmdk"] = "NewVmdk"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
		map[string][]bindings.FieldData{
			"IDE": []bindings.FieldData{
				bindings.NewFieldData("ide", false),
			},
			"SCSI": []bindings.FieldData{
				bindings.NewFieldData("scsi", false),
			},
			"SATA": []bindings.FieldData{
				bindings.NewFieldData("sata", false),
			},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.disk.create_spec", fields, reflect.TypeOf(Disk.DiskCreateSpec{}), fieldNameMap, validators)
}

func DiskUpdateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["backing"] = bindings.NewOptionalType(bindings.NewReferenceType(DiskBackingSpecBindingType))
	fieldNameMap["backing"] = "Backing"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.disk.update_spec", fields, reflect.TypeOf(DiskUpdateSpec{}), fieldNameMap, validators)
}

func DiskSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["disk"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Disk"}, "")
	fieldNameMap["disk"] = "Disk"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.disk.summary", fields, reflect.TypeOf(DiskSummary{}), fieldNameMap, validators)
}

