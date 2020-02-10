/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: TagUsage.
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



// The ``Summary`` class contains common information about a tag used by a policy. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type TagUsageSummary struct {
    // Identifier of the policy that uses the tag indicated by TagUsageSummary#tag. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Policy string
    // Identifier of the tag type used by the policy indicated by TagUsageSummary#policy. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	TagType string
    // Identifier of the tag used by the policy indicated by TagUsageSummary#policy. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Tag string
    // Name of the tag used by the policy indicated by TagUsageSummary#policy. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	TagName string
    // Name of the category that has TagUsageSummary#tag. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CategoryName string
}

// The ``FilterSpec`` class contains properties used to filter the results when listing the tags used by policies as available in this vCenter server (see TagUsage#list). If multiple properties are specified, only the tags used by policies that match an element of each property match the filter. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type TagUsageFilterSpec struct {
    // Identifiers that compute policies must have to match the filter. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Policies map[string]bool
    // Identifiers that tags must have to match the filter. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Tags map[string]bool
    // Identifiers that tag types must have to match the filter. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	TagTypes map[string]bool
}



func tagUsageListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(TagUsageFilterSpecBindingType))
	fieldNameMap["filter"] = "Filter"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tagUsageListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(TagUsageSummaryBindingType), reflect.TypeOf([]TagUsageSummary{}))
}

func tagUsageListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(TagUsageFilterSpecBindingType))
	fieldNameMap["filter"] = "Filter"
	paramsTypeMap["filter.policies"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string{"com.vmware.vcenter.compute.Policy"}, ""), reflect.TypeOf(map[string]bool{})))
	paramsTypeMap["filter.tags"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag"}, ""), reflect.TypeOf(map[string]bool{})))
	paramsTypeMap["filter.tag_types"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string{"com.vmware.vapi.resource"}, ""), reflect.TypeOf(map[string]bool{})))
	queryParams["filter.policies"] = "policies"
	queryParams["filter.tag_types"] = "tag_types"
	queryParams["filter.tags"] = "tags"
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
		"/vcenter/compute/policies/tag-usage",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Unauthorized": 403})
}


func TagUsageSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["policy"] = bindings.NewIdType([]string{"com.vmware.vcenter.compute.Policy"}, "")
	fieldNameMap["policy"] = "Policy"
	fields["tag_type"] = bindings.NewIdType([]string{"com.vmware.vapi.resource"}, "")
	fieldNameMap["tag_type"] = "TagType"
	fields["tag"] = bindings.NewIdType(nil, "tag_type")
	fieldNameMap["tag"] = "Tag"
	fields["tag_name"] = bindings.NewStringType()
	fieldNameMap["tag_name"] = "TagName"
	fields["category_name"] = bindings.NewStringType()
	fieldNameMap["category_name"] = "CategoryName"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.compute.policies.tag_usage.summary", fields, reflect.TypeOf(TagUsageSummary{}), fieldNameMap, validators)
}

func TagUsageFilterSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["policies"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string{"com.vmware.vcenter.compute.Policy"}, ""), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["policies"] = "Policies"
	fields["tags"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag"}, ""), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["tags"] = "Tags"
	fields["tag_types"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string{"com.vmware.vapi.resource"}, ""), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["tag_types"] = "TagTypes"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.compute.policies.tag_usage.filter_spec", fields, reflect.TypeOf(TagUsageFilterSpec{}), fieldNameMap, validators)
}

