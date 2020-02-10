/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Policies.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package compute

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/vsphere/vcenter/compute/policies"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``Info`` class contains information about the compliance of a virtual machine with a compute policy. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type PoliciesInfo struct {
    // The compliance status of the policy on a specified object.**Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Status policies.ObjectCompliance
}



func policiesGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["policy"] = bindings.NewIdType([]string{"com.vmware.vcenter.compute.Resources.COMPUTE_POLICY"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["policy"] = "Policy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func policiesGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(PoliciesInfoBindingType)
}

func policiesGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["policy"] = bindings.NewIdType([]string{"com.vmware.vcenter.compute.Resources.COMPUTE_POLICY"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["policy"] = "Policy"
	paramsTypeMap["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	paramsTypeMap["policy"] = bindings.NewIdType([]string{"com.vmware.vcenter.compute.Resources.COMPUTE_POLICY"}, "")
	paramsTypeMap["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	paramsTypeMap["policy"] = bindings.NewIdType([]string{"com.vmware.vcenter.compute.Resources.COMPUTE_POLICY"}, "")
	pathParams["vm"] = "vm"
	pathParams["policy"] = "policy"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"GET",
		"/vcenter/vm/{vm}/compute/policies/{policy}",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"NotFound": 404,"Unauthorized": 403})
}


func PoliciesInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["status"] = bindings.NewEnumType("com.vmware.vcenter.compute.policies.object_compliance", reflect.TypeOf(policies.ObjectCompliance(policies.ObjectCompliance_UNKNOWN)))
	fieldNameMap["status"] = "Status"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm.compute.policies.info", fields, reflect.TypeOf(PoliciesInfo{}), fieldNameMap, validators)
}

