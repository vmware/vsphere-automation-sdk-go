/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Policy.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package storage

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``VmHomePolicySpec`` class provides a specification for the storage policy to be associated with the virtual machine home's directory. This class was added in vSphere API 6.7.
type PolicyVmHomePolicySpec struct {
    // Policy type to be used while performing update operation on the virtual machine home's directory. This property was added in vSphere API 6.7.
	Type_ PolicyVmHomePolicySpecPolicyType
    // Storage Policy identification. This property was added in vSphere API 6.7.
	Policy *string
}

// The ``PolicyType`` enumeration class defines the choices for how to specify the policy to be associated with the virtual machine home's directory. This enumeration was added in vSphere API 6.7.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type PolicyVmHomePolicySpecPolicyType string

const (
    // Use the specified policy (see PolicyVmHomePolicySpec#policy). This constant field was added in vSphere API 6.7.
	PolicyVmHomePolicySpecPolicyType_USE_SPECIFIED_POLICY PolicyVmHomePolicySpecPolicyType = "USE_SPECIFIED_POLICY"
    // Use the default storage policy of the datastore. This constant field was added in vSphere API 6.7.
	PolicyVmHomePolicySpecPolicyType_USE_DEFAULT_POLICY PolicyVmHomePolicySpecPolicyType = "USE_DEFAULT_POLICY"
)

func (p PolicyVmHomePolicySpecPolicyType) PolicyVmHomePolicySpecPolicyType() bool {
	switch p {
	case PolicyVmHomePolicySpecPolicyType_USE_SPECIFIED_POLICY:
		return true
	case PolicyVmHomePolicySpecPolicyType_USE_DEFAULT_POLICY:
		return true
	default:
		return false
	}
}


// The ``DiskPolicySpec`` class provides a specification for the storage policy to be associated with the virtual disks. This class was added in vSphere API 6.7.
type PolicyDiskPolicySpec struct {
    // Policy type to be used while performing update operation on the virtual disks. This property was added in vSphere API 6.7.
	Type_ PolicyDiskPolicySpecPolicyType
    // Storage Policy identification. This property was added in vSphere API 6.7.
	Policy *string
}

// The ``DiskPolicySpec`` enumeration class defines the choices for how to specify the policy to be associated with a virtual disk. This enumeration was added in vSphere API 6.7.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type PolicyDiskPolicySpecPolicyType string

const (
    // Use the specified policy (see PolicyDiskPolicySpec#policy). This constant field was added in vSphere API 6.7.
	PolicyDiskPolicySpecPolicyType_USE_SPECIFIED_POLICY PolicyDiskPolicySpecPolicyType = "USE_SPECIFIED_POLICY"
    // Use the default storage policy of the datastore. This constant field was added in vSphere API 6.7.
	PolicyDiskPolicySpecPolicyType_USE_DEFAULT_POLICY PolicyDiskPolicySpecPolicyType = "USE_DEFAULT_POLICY"
)

func (p PolicyDiskPolicySpecPolicyType) PolicyDiskPolicySpecPolicyType() bool {
	switch p {
	case PolicyDiskPolicySpecPolicyType_USE_SPECIFIED_POLICY:
		return true
	case PolicyDiskPolicySpecPolicyType_USE_DEFAULT_POLICY:
		return true
	default:
		return false
	}
}


// The ``UpdateSpec`` class describes the updates to be made to the storage policies associated with the virtual machine home and/or its virtual disks. This class was added in vSphere API 6.7.
type PolicyUpdateSpec struct {
    // Storage policy to be used when reconfiguring the virtual machine home. This property was added in vSphere API 6.7.
	VmHome *PolicyVmHomePolicySpec
    // Storage policy or policies to be used when reconfiguring virtual machine diks. This property was added in vSphere API 6.7.
	Disks map[string]PolicyDiskPolicySpec
}

// The ``Info`` class contains information about the storage policies associated with virtual machine's home directory and virtual hard disks. This class was added in vSphere API 6.7.
type PolicyInfo struct {
    // Storage Policy associated with virtual machine home. This property was added in vSphere API 6.7.
	VmHome *string
    // Storage policies associated with virtual disks. The values in this map are storage policy identifiers. They will be identifiers for the resource type:com.vmware.vcenter.StoragePolicy If the map is empty, the virtual machine does not have any disks or its disks are not associated with a storage policy. This property was added in vSphere API 6.7.
	Disks map[string]string
}



func policyUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["spec"] = bindings.NewReferenceType(PolicyUpdateSpecBindingType)
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func policyUpdateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func policyUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["spec"] = bindings.NewReferenceType(PolicyUpdateSpecBindingType)
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
		map[string]int{"Error": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403,"InvalidArgument": 400,"ResourceBusy": 500,"ResourceInaccessible": 500})
}

func policyGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fieldNameMap["vm"] = "Vm"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func policyGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(PolicyInfoBindingType)
}

func policyGetRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"Error": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func PolicyVmHomePolicySpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm.storage.policy.vm_home_policy_spec.policy_type", reflect.TypeOf(PolicyVmHomePolicySpecPolicyType(PolicyVmHomePolicySpecPolicyType_USE_SPECIFIED_POLICY)))
	fieldNameMap["type"] = "Type_"
	fields["policy"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.vcenter.StoragePolicy"}, ""))
	fieldNameMap["policy"] = "Policy"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
		map[string][]bindings.FieldData{
			"USE_SPECIFIED_POLICY": []bindings.FieldData{
				bindings.NewFieldData("policy", true),
			},
			"USE_DEFAULT_POLICY": []bindings.FieldData{},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.vm.storage.policy.vm_home_policy_spec", fields, reflect.TypeOf(PolicyVmHomePolicySpec{}), fieldNameMap, validators)
}

func PolicyDiskPolicySpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm.storage.policy.disk_policy_spec.policy_type", reflect.TypeOf(PolicyDiskPolicySpecPolicyType(PolicyDiskPolicySpecPolicyType_USE_SPECIFIED_POLICY)))
	fieldNameMap["type"] = "Type_"
	fields["policy"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.vcenter.StoragePolicy"}, ""))
	fieldNameMap["policy"] = "Policy"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
		map[string][]bindings.FieldData{
			"USE_SPECIFIED_POLICY": []bindings.FieldData{
				bindings.NewFieldData("policy", true),
			},
			"USE_DEFAULT_POLICY": []bindings.FieldData{},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.vm.storage.policy.disk_policy_spec", fields, reflect.TypeOf(PolicyDiskPolicySpec{}), fieldNameMap, validators)
}

func PolicyUpdateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm_home"] = bindings.NewOptionalType(bindings.NewReferenceType(PolicyVmHomePolicySpecBindingType))
	fieldNameMap["vm_home"] = "VmHome"
	fields["disks"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Disk"}, ""), bindings.NewReferenceType(PolicyDiskPolicySpecBindingType),reflect.TypeOf(map[string]PolicyDiskPolicySpec{})))
	fieldNameMap["disks"] = "Disks"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm.storage.policy.update_spec", fields, reflect.TypeOf(PolicyUpdateSpec{}), fieldNameMap, validators)
}

func PolicyInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm_home"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.vcenter.StoragePolicy"}, ""))
	fieldNameMap["vm_home"] = "VmHome"
	fields["disks"] = bindings.NewMapType(bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Disk"}, ""), bindings.NewStringType(),reflect.TypeOf(map[string]string{}))
	fieldNameMap["disks"] = "Disks"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm.storage.policy.info", fields, reflect.TypeOf(PolicyInfo{}), fieldNameMap, validators)
}

