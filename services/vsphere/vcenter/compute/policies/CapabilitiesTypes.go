/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Capabilities.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package policies

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// The resource type for the compute policy capability. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
const Capabilities_RESOURCE_TYPE = "com.vmware.vcenter.compute.policies.Capability"


// The ``Summary`` class contains commonly used information about a compute policy capability. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type CapabilitiesSummary struct {
    // Identifier of the capability. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Capability string
    // Name of the capability. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Name string
    // Description of the capability. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Description string
}

// The ``Info`` class contains information about a compute policy capability. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type CapabilitiesInfo struct {
    // Name of the capability. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Name string
    // Description of the capability. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Description string
    // Identifier of the class used to create a policy based on this capability. See Policies#create. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CreateSpecType string
    // Identifier of the class returned when retrieving information about a policy based on this capability. See Policies#get. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	InfoType string
}



func capabilitiesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func capabilitiesListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(CapabilitiesSummaryBindingType), reflect.TypeOf([]CapabilitiesSummary{}))
}

func capabilitiesListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
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
		"/vcenter/compute/policies/capabilities",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Unauthorized": 403})
}

func capabilitiesGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["capability"] = bindings.NewIdType([]string{"com.vmware.vcenter.compute.policies.Capability"}, "")
	fieldNameMap["capability"] = "Capability"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func capabilitiesGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(CapabilitiesInfoBindingType)
}

func capabilitiesGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["capability"] = bindings.NewIdType([]string{"com.vmware.vcenter.compute.policies.Capability"}, "")
	fieldNameMap["capability"] = "Capability"
	paramsTypeMap["capability"] = bindings.NewIdType([]string{"com.vmware.vcenter.compute.policies.Capability"}, "")
	paramsTypeMap["capability"] = bindings.NewIdType([]string{"com.vmware.vcenter.compute.policies.Capability"}, "")
	pathParams["capability"] = "capability"
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
		"/vcenter/compute/policies/capabilities/{capability}",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"NotFound": 404,"Unauthorized": 403})
}


func CapabilitiesSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["capability"] = bindings.NewIdType([]string{"com.vmware.vcenter.compute.policies.Capability"}, "")
	fieldNameMap["capability"] = "Capability"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["description"] = bindings.NewStringType()
	fieldNameMap["description"] = "Description"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.compute.policies.capabilities.summary", fields, reflect.TypeOf(CapabilitiesSummary{}), fieldNameMap, validators)
}

func CapabilitiesInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["description"] = bindings.NewStringType()
	fieldNameMap["description"] = "Description"
	fields["create_spec_type"] = bindings.NewIdType([]string{"com.vmware.vapi.structure"}, "")
	fieldNameMap["create_spec_type"] = "CreateSpecType"
	fields["info_type"] = bindings.NewIdType([]string{"com.vmware.vapi.structure"}, "")
	fieldNameMap["info_type"] = "InfoType"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.compute.policies.capabilities.info", fields, reflect.TypeOf(CapabilitiesInfo{}), fieldNameMap, validators)
}

