/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Boot.
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



// The ``Type`` enumeration class defines the valid firmware types for a virtual machine.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type BootType string

const (
    // Basic Input/Output System (BIOS) firmware.
	BootType_BIOS BootType = "BIOS"
    // Extensible Firmware Interface (EFI) firmware.
	BootType_EFI BootType = "EFI"
)

func (t BootType) BootType() bool {
	switch t {
	case BootType_BIOS:
		return true
	case BootType_EFI:
		return true
	default:
		return false
	}
}


// The ``NetworkProtocol`` enumeration class defines the valid network boot protocols supported when booting a virtual machine with BootType#BootType_EFI firmware over the network.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type BootNetworkProtocol string

const (
    // PXE or Apple NetBoot over IPv4.
	BootNetworkProtocol_IPV4 BootNetworkProtocol = "IPV4"
    // PXE over IPv6.
	BootNetworkProtocol_IPV6 BootNetworkProtocol = "IPV6"
)

func (n BootNetworkProtocol) BootNetworkProtocol() bool {
	switch n {
	case BootNetworkProtocol_IPV4:
		return true
	case BootNetworkProtocol_IPV6:
		return true
	default:
		return false
	}
}


// The ``Info`` class contains information about the virtual machine boot process.
type BootInfo struct {
    // Firmware type used by the virtual machine.
	Type_ BootType
    // Flag indicating whether to use EFI legacy boot mode.
	EfiLegacyBoot *bool
    // Protocol to use when attempting to boot the virtual machine over the network.
	NetworkProtocol *BootNetworkProtocol
    // Delay in milliseconds before beginning the firmware boot process when the virtual machine is powered on. This delay may be used to provide a time window for users to connect to the virtual machine console and enter BIOS setup mode.
	Delay int64
    // Flag indicating whether the virtual machine will automatically retry the boot process after a failure.
	Retry bool
    // Delay in milliseconds before retrying the boot process after a failure; applicable only when BootInfo#retry is true.
	RetryDelay int64
    // Flag indicating whether the firmware boot process will automatically enter setup mode the next time the virtual machine boots. Note that this flag will automatically be reset to false once the virtual machine enters setup mode.
	EnterSetupMode bool
}

// The ``CreateSpec`` class describes settings used when booting a virtual machine.
type BootCreateSpec struct {
    // Firmware type to be used by the virtual machine.
	Type_ *BootType
    // Flag indicating whether to use EFI legacy boot mode.
	EfiLegacyBoot *bool
    // Protocol to use when attempting to boot the virtual machine over the network.
	NetworkProtocol *BootNetworkProtocol
    // Delay in milliseconds before beginning the firmware boot process when the virtual machine is powered on. This delay may be used to provide a time window for users to connect to the virtual machine console and enter BIOS setup mode.
	Delay *int64
    // Flag indicating whether the virtual machine should automatically retry the boot process after a failure.
	Retry *bool
    // Delay in milliseconds before retrying the boot process after a failure; applicable only when BootInfo#retry is true.
	RetryDelay *int64
    // Flag indicating whether the firmware boot process should automatically enter setup mode the next time the virtual machine boots. Note that this flag will automatically be reset to false once the virtual machine enters setup mode.
	EnterSetupMode *bool
}

// The ``UpdateSpec`` class describes the updates to the settings used when booting a virtual machine.
type BootUpdateSpec struct {
    // Firmware type to be used by the virtual machine.
	Type_ *BootType
    // Flag indicating whether to use EFI legacy boot mode.
	EfiLegacyBoot *bool
    // Protocol to use when attempting to boot the virtual machine over the network.
	NetworkProtocol *BootNetworkProtocol
    // Delay in milliseconds before beginning the firmware boot process when the virtual machine is powered on. This delay may be used to provide a time window for users to connect to the virtual machine console and enter BIOS setup mode.
	Delay *int64
    // Flag indicating whether the virtual machine should automatically retry the boot process after a failure.
	Retry *bool
    // Delay in milliseconds before retrying the boot process after a failure; applicable only when BootInfo#retry is true.
	RetryDelay *int64
    // Flag indicating whether the firmware boot process should automatically enter setup mode the next time the virtual machine boots. Note that this flag will automatically be reset to false once the virtual machine enters setup mode.
	EnterSetupMode *bool
}



func bootGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fieldNameMap["vm"] = "Vm"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func bootGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(BootInfoBindingType)
}

func bootGetRestMetadata() protocol.OperationRestMetadata {
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

func bootUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["spec"] = bindings.NewReferenceType(BootUpdateSpecBindingType)
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func bootUpdateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func bootUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["spec"] = bindings.NewReferenceType(BootUpdateSpecBindingType)
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
		map[string]int{"Error": 500,"NotFound": 404,"InvalidArgument": 400,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func BootInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.boot.type", reflect.TypeOf(BootType(BootType_BIOS)))
	fieldNameMap["type"] = "Type_"
	fields["efi_legacy_boot"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["efi_legacy_boot"] = "EfiLegacyBoot"
	fields["network_protocol"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.hardware.boot.network_protocol", reflect.TypeOf(BootNetworkProtocol(BootNetworkProtocol_IPV4))))
	fieldNameMap["network_protocol"] = "NetworkProtocol"
	fields["delay"] = bindings.NewIntegerType()
	fieldNameMap["delay"] = "Delay"
	fields["retry"] = bindings.NewBooleanType()
	fieldNameMap["retry"] = "Retry"
	fields["retry_delay"] = bindings.NewIntegerType()
	fieldNameMap["retry_delay"] = "RetryDelay"
	fields["enter_setup_mode"] = bindings.NewBooleanType()
	fieldNameMap["enter_setup_mode"] = "EnterSetupMode"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
		map[string][]bindings.FieldData{
			"EFI": []bindings.FieldData{
				bindings.NewFieldData("efi_legacy_boot", true),
				bindings.NewFieldData("network_protocol", true),
			},
			"BIOS": []bindings.FieldData{},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.boot.info", fields, reflect.TypeOf(BootInfo{}), fieldNameMap, validators)
}

func BootCreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.hardware.boot.type", reflect.TypeOf(BootType(BootType_BIOS))))
	fieldNameMap["type"] = "Type_"
	fields["efi_legacy_boot"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["efi_legacy_boot"] = "EfiLegacyBoot"
	fields["network_protocol"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.hardware.boot.network_protocol", reflect.TypeOf(BootNetworkProtocol(BootNetworkProtocol_IPV4))))
	fieldNameMap["network_protocol"] = "NetworkProtocol"
	fields["delay"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["delay"] = "Delay"
	fields["retry"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["retry"] = "Retry"
	fields["retry_delay"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["retry_delay"] = "RetryDelay"
	fields["enter_setup_mode"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["enter_setup_mode"] = "EnterSetupMode"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
		map[string][]bindings.FieldData{
			"EFI": []bindings.FieldData{
				bindings.NewFieldData("efi_legacy_boot", false),
				bindings.NewFieldData("network_protocol", false),
			},
			"BIOS": []bindings.FieldData{},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.boot.create_spec", fields, reflect.TypeOf(BootCreateSpec{}), fieldNameMap, validators)
}

func BootUpdateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.hardware.boot.type", reflect.TypeOf(BootType(BootType_BIOS))))
	fieldNameMap["type"] = "Type_"
	fields["efi_legacy_boot"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["efi_legacy_boot"] = "EfiLegacyBoot"
	fields["network_protocol"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.hardware.boot.network_protocol", reflect.TypeOf(BootNetworkProtocol(BootNetworkProtocol_IPV4))))
	fieldNameMap["network_protocol"] = "NetworkProtocol"
	fields["delay"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["delay"] = "Delay"
	fields["retry"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["retry"] = "Retry"
	fields["retry_delay"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["retry_delay"] = "RetryDelay"
	fields["enter_setup_mode"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["enter_setup_mode"] = "EnterSetupMode"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
		map[string][]bindings.FieldData{
			"EFI": []bindings.FieldData{
				bindings.NewFieldData("efi_legacy_boot", false),
				bindings.NewFieldData("network_protocol", false),
			},
			"BIOS": []bindings.FieldData{},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.boot.update_spec", fields, reflect.TypeOf(BootUpdateSpec{}), fieldNameMap, validators)
}

