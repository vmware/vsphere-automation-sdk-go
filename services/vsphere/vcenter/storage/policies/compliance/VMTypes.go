/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: VM.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package compliance

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The {\\\\@Status} enumeration class defines he valid compliance status values for a virtual machine or virtual disk. This enumeration was added in vSphere API 6.7.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type VMStatus string

const (
    // The virtual machine or virtual disk is in compliance. This constant field was added in vSphere API 6.7.
	VMStatus_COMPLIANT VMStatus = "COMPLIANT"
    // The virtual machine or virtual disk is in not in compliance. This constant field was added in vSphere API 6.7.
	VMStatus_NON_COMPLIANT VMStatus = "NON_COMPLIANT"
    // Compliance status of the virtual machine or virtual disk is not known. This constant field was added in vSphere API 6.7.
	VMStatus_UNKNOWN_COMPLIANCE VMStatus = "UNKNOWN_COMPLIANCE"
    // Compliance computation is not applicable for this virtual machine or disk because it does not have any storage requirement that apply to the object-based datastore on which the entity is placed. This constant field was added in vSphere API 6.7.
	VMStatus_NOT_APPLICABLE VMStatus = "NOT_APPLICABLE"
    // Compliance status becomes out of date when the profile associated with the virtual machine or disk is edited and not applied. The compliance status will remain out of date until the latest policy is applied. This constant field was added in vSphere API 6.7.
	VMStatus_OUT_OF_DATE VMStatus = "OUT_OF_DATE"
)

func (s VMStatus) VMStatus() bool {
	switch s {
	case VMStatus_COMPLIANT:
		return true
	case VMStatus_NON_COMPLIANT:
		return true
	case VMStatus_UNKNOWN_COMPLIANCE:
		return true
	case VMStatus_NOT_APPLICABLE:
		return true
	case VMStatus_OUT_OF_DATE:
		return true
	default:
		return false
	}
}


// Provides the compliance details of a virtual machine and its associated entities which match the given compliance statuses. This class was added in vSphere API 6.7.
type VMInfo struct {
    // Compliance status of the virtual machine home. This property was added in vSphere API 6.7.
	VmHome *VMStatus
    // A Map of virtual disks and their compliance status If empty, the virtual machine does not have any disks or its disks are not associated with a storage policy. This property was added in vSphere API 6.7.
	Disks map[string]VMStatus
}

// The ``FilterSpec`` class contains Status used to filter the results when listing virtual machines (see VM#list). This class was added in vSphere API 6.7.
type VMFilterSpec struct {
    // Compliance Status that a virtual machine must have to match the filter. Atleast one status must be specified. This property was added in vSphere API 6.7.
	Status map[VMStatus]bool
    // Identifiers of virtual machines that can match the filter. This property was added in vSphere API 6.7.
	Vms map[string]bool
}



func vMListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["filter"] = bindings.NewReferenceType(VMFilterSpecBindingType)
	fieldNameMap["filter"] = "Filter"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func vMListOutputType() bindings.BindingType {
	return bindings.NewMapType(bindings.NewIdType([]string{"VirtualMachine"}, ""), bindings.NewReferenceType(VMInfoBindingType),reflect.TypeOf(map[string]VMInfo{}))
}

func vMListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["filter"] = bindings.NewReferenceType(VMFilterSpecBindingType)
	fieldNameMap["filter"] = "Filter"
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
		map[string]int{"Error": 500,"InvalidArgument": 400,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403,"Unsupported": 400,"UnableToAllocateResource": 500})
}


func VMInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm_home"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.storage.policies.compliance.VM.status", reflect.TypeOf(VMStatus(VMStatus_COMPLIANT))))
	fieldNameMap["vm_home"] = "VmHome"
	fields["disks"] = bindings.NewMapType(bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Disk"}, ""), bindings.NewEnumType("com.vmware.vcenter.storage.policies.compliance.VM.status", reflect.TypeOf(VMStatus(VMStatus_COMPLIANT))),reflect.TypeOf(map[string]VMStatus{}))
	fieldNameMap["disks"] = "Disks"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.storage.policies.compliance.VM.info", fields, reflect.TypeOf(VMInfo{}), fieldNameMap, validators)
}

func VMFilterSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["status"] = bindings.NewSetType(bindings.NewEnumType("com.vmware.vcenter.storage.policies.compliance.VM.status", reflect.TypeOf(VMStatus(VMStatus_COMPLIANT))), reflect.TypeOf(map[VMStatus]bool{}))
	fieldNameMap["status"] = "Status"
	fields["vms"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string{"VirtualMachine"}, ""), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["vms"] = "Vms"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.storage.policies.compliance.VM.filter_spec", fields, reflect.TypeOf(VMFilterSpec{}), fieldNameMap, validators)
}

