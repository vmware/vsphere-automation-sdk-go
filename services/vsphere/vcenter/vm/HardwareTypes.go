/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Hardware.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package vm

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/vsphere/vcenter/vm/Hardware"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``Version`` enumeration class defines the valid virtual hardware versions for a virtual machine. See https://kb.vmware.com/s/article/1003746 (Virtual machine hardware versions (1003746)).
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type Hardware.HardwareVersion string

const (
    // Hardware version 3, first supported in ESXi 2.5.
	HardwareVersion_VMX_03 Hardware.HardwareVersion = "VMX_03"
    // Hardware version 4, first supported in ESXi 3.0.
	HardwareVersion_VMX_04 Hardware.HardwareVersion = "VMX_04"
    // Hardware version 6, first supported in WS 6.0.
	HardwareVersion_VMX_06 Hardware.HardwareVersion = "VMX_06"
    // Hardware version 7, first supported in ESXi 4.0.
	HardwareVersion_VMX_07 Hardware.HardwareVersion = "VMX_07"
    // Hardware version 8, first supported in ESXi 5.0.
	HardwareVersion_VMX_08 Hardware.HardwareVersion = "VMX_08"
    // Hardware version 9, first supported in ESXi 5.1.
	HardwareVersion_VMX_09 Hardware.HardwareVersion = "VMX_09"
    // Hardware version 10, first supported in ESXi 5.5.
	HardwareVersion_VMX_10 Hardware.HardwareVersion = "VMX_10"
    // Hardware version 11, first supported in ESXi 6.0.
	HardwareVersion_VMX_11 Hardware.HardwareVersion = "VMX_11"
    // Hardware version 12, first supported in Workstation 12.0.
	HardwareVersion_VMX_12 Hardware.HardwareVersion = "VMX_12"
    // Hardware version 13, first supported in ESXi 6.5.
	HardwareVersion_VMX_13 Hardware.HardwareVersion = "VMX_13"
    // Hardware version 14, first supported in ESXi 6.7. This constant field was added in vSphere API 6.7.
	HardwareVersion_VMX_14 Hardware.HardwareVersion = "VMX_14"
    // Hardware version 15, first supported in ESXi 6.7.0 Update 2. This constant field was added in vSphere API 6.7.2.
	HardwareVersion_VMX_15 Hardware.HardwareVersion = "VMX_15"
    // Hardware version 16, first supported in Workstation 15.0. This constant field was added in vSphere API 7.0.0.
	HardwareVersion_VMX_16 Hardware.HardwareVersion = "VMX_16"
    // Hardware version 17, first supported in ESX 7.0. This constant field was added in vSphere API 7.0.0.
	HardwareVersion_VMX_17 Hardware.HardwareVersion = "VMX_17"
)

func (v Hardware.HardwareVersion) Hardware.HardwareVersion() bool {
	switch v {
	case HardwareVersion_VMX_03:
		return true
	case HardwareVersion_VMX_04:
		return true
	case HardwareVersion_VMX_06:
		return true
	case HardwareVersion_VMX_07:
		return true
	case HardwareVersion_VMX_08:
		return true
	case HardwareVersion_VMX_09:
		return true
	case HardwareVersion_VMX_10:
		return true
	case HardwareVersion_VMX_11:
		return true
	case HardwareVersion_VMX_12:
		return true
	case HardwareVersion_VMX_13:
		return true
	case HardwareVersion_VMX_14:
		return true
	case HardwareVersion_VMX_15:
		return true
	case HardwareVersion_VMX_16:
		return true
	case HardwareVersion_VMX_17:
		return true
	default:
		return false
	}
}


// The ``UpgradePolicy`` enumeration class defines the valid virtual hardware upgrade policies for a virtual machine.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type HardwareUpgradePolicy string

const (
    // Do not upgrade the virtual machine when it is powered on.
	HardwareUpgradePolicy_NEVER HardwareUpgradePolicy = "NEVER"
    // Run scheduled upgrade when the virtual machine is powered on after a clean shutdown of the guest operating system.
	HardwareUpgradePolicy_AFTER_CLEAN_SHUTDOWN HardwareUpgradePolicy = "AFTER_CLEAN_SHUTDOWN"
    // Run scheduled upgrade when the virtual machine is powered on.
	HardwareUpgradePolicy_ALWAYS HardwareUpgradePolicy = "ALWAYS"
)

func (u HardwareUpgradePolicy) HardwareUpgradePolicy() bool {
	switch u {
	case HardwareUpgradePolicy_NEVER:
		return true
	case HardwareUpgradePolicy_AFTER_CLEAN_SHUTDOWN:
		return true
	case HardwareUpgradePolicy_ALWAYS:
		return true
	default:
		return false
	}
}


// The ``UpgradeStatus`` enumeration class defines the valid virtual hardware upgrade statuses for a virtual machine.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type HardwareUpgradeStatus string

const (
    // No scheduled upgrade has been attempted.
	HardwareUpgradeStatus_NONE HardwareUpgradeStatus = "NONE"
    // Upgrade is scheduled but has not yet been run.
	HardwareUpgradeStatus_PENDING HardwareUpgradeStatus = "PENDING"
    // The most recent scheduled upgrade was successful.
	HardwareUpgradeStatus_SUCCESS HardwareUpgradeStatus = "SUCCESS"
    // The most recent scheduled upgrade was not successful.
	HardwareUpgradeStatus_FAILED HardwareUpgradeStatus = "FAILED"
)

func (u HardwareUpgradeStatus) HardwareUpgradeStatus() bool {
	switch u {
	case HardwareUpgradeStatus_NONE:
		return true
	case HardwareUpgradeStatus_PENDING:
		return true
	case HardwareUpgradeStatus_SUCCESS:
		return true
	case HardwareUpgradeStatus_FAILED:
		return true
	default:
		return false
	}
}


// The ``Info`` class contains information related to the virtual hardware of a virtual machine.
type HardwareInfo struct {
    // Virtual hardware version.
	Version Hardware.HardwareVersion
    // Scheduled upgrade policy.
	UpgradePolicy HardwareUpgradePolicy
    // Target hardware version to be used on the next scheduled virtual hardware upgrade.
	UpgradeVersion *Hardware.HardwareVersion
    // Scheduled upgrade status.
	UpgradeStatus HardwareUpgradeStatus
    // Reason for the scheduled upgrade failure.
	UpgradeError *data.ErrorValue
}

// The ``UpdateSpec`` class describes the updates to virtual hardware settings of a virtual machine.
type HardwareUpdateSpec struct {
    // Scheduled upgrade policy. 
    //
    //  If set to HardwareUpgradePolicy#HardwareUpgradePolicy_NEVER, the HardwareInfo#upgradeVersion property will be reset to null.
	UpgradePolicy *HardwareUpgradePolicy
    // Target hardware version to be used on the next scheduled virtual hardware upgrade. 
    //
    //  If specified, this property must represent a newer virtual hardware version than the current virtual hardware version reported in HardwareInfo#version.
	UpgradeVersion *Hardware.HardwareVersion
}



func hardwareGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fieldNameMap["vm"] = "Vm"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func hardwareGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(HardwareInfoBindingType)
}

func hardwareGetRestMetadata() protocol.OperationRestMetadata {
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

func hardwareUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["spec"] = bindings.NewReferenceType(HardwareUpdateSpecBindingType)
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func hardwareUpdateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func hardwareUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["spec"] = bindings.NewReferenceType(HardwareUpdateSpecBindingType)
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
		map[string]int{"Error": 500,"NotFound": 404,"AlreadyInDesiredState": 400,"InvalidArgument": 400,"Unsupported": 400,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func hardwareUpgradeInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["version"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.hardware.version", reflect.TypeOf(Hardware.HardwareVersion(Hardware.HardwareVersion_VMX_03))))
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["version"] = "Version"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func hardwareUpgradeOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func hardwareUpgradeRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["version"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.hardware.version", reflect.TypeOf(Hardware.HardwareVersion(Hardware.HardwareVersion_VMX_03))))
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["version"] = "Version"
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
		map[string]int{"Error": 500,"NotFound": 404,"NotAllowedInCurrentState": 400,"AlreadyInDesiredState": 400,"InvalidArgument": 400,"Unsupported": 400,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func HardwareInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["version"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.version", reflect.TypeOf(Hardware.HardwareVersion(Hardware.HardwareVersion_VMX_03)))
	fieldNameMap["version"] = "Version"
	fields["upgrade_policy"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.upgrade_policy", reflect.TypeOf(HardwareUpgradePolicy(HardwareUpgradePolicy_NEVER)))
	fieldNameMap["upgrade_policy"] = "UpgradePolicy"
	fields["upgrade_version"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.hardware.version", reflect.TypeOf(Hardware.HardwareVersion(Hardware.HardwareVersion_VMX_03))))
	fieldNameMap["upgrade_version"] = "UpgradeVersion"
	fields["upgrade_status"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.upgrade_status", reflect.TypeOf(HardwareUpgradeStatus(HardwareUpgradeStatus_NONE)))
	fieldNameMap["upgrade_status"] = "UpgradeStatus"
	fields["upgrade_error"] = bindings.NewOptionalType(bindings.NewAnyErrorType())
	fieldNameMap["upgrade_error"] = "UpgradeError"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("upgrade_policy",
		map[string][]bindings.FieldData{
			"AFTER_CLEAN_SHUTDOWN": []bindings.FieldData{
				bindings.NewFieldData("upgrade_version", true),
			},
			"ALWAYS": []bindings.FieldData{
				bindings.NewFieldData("upgrade_version", true),
			},
			"NEVER": []bindings.FieldData{},
		},
	)
	validators = append(validators, uv1)
	uv2 := bindings.NewUnionValidator("upgrade_status",
		map[string][]bindings.FieldData{
			"FAILED": []bindings.FieldData{
				bindings.NewFieldData("upgrade_error", true),
			},
			"NONE": []bindings.FieldData{},
			"PENDING": []bindings.FieldData{},
			"SUCCESS": []bindings.FieldData{},
		},
	)
	validators = append(validators, uv2)
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.info", fields, reflect.TypeOf(HardwareInfo{}), fieldNameMap, validators)
}

func HardwareUpdateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["upgrade_policy"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.hardware.upgrade_policy", reflect.TypeOf(HardwareUpgradePolicy(HardwareUpgradePolicy_NEVER))))
	fieldNameMap["upgrade_policy"] = "UpgradePolicy"
	fields["upgrade_version"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.hardware.version", reflect.TypeOf(Hardware.HardwareVersion(Hardware.HardwareVersion_VMX_03))))
	fieldNameMap["upgrade_version"] = "UpgradeVersion"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("upgrade_policy",
		map[string][]bindings.FieldData{
			"AFTER_CLEAN_SHUTDOWN": []bindings.FieldData{
				bindings.NewFieldData("upgrade_version", false),
			},
			"ALWAYS": []bindings.FieldData{
				bindings.NewFieldData("upgrade_version", false),
			},
			"NEVER": []bindings.FieldData{},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.update_spec", fields, reflect.TypeOf(HardwareUpdateSpec{}), fieldNameMap, validators)
}

