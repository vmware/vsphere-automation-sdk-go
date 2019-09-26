/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: TagUsage.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package tagUsage

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)




// The ``Summary`` class contains common information about a tag used by a policy. **Warning:** This class is available as technical preview. It may be changed in a future release.
 type Summary struct {
    Policy string
    TagType string
    Tag string
    TagName string
    CategoryName string
}






// The ``FilterSpec`` class contains properties used to filter the results when listing the tags used by policies as available in this vCenter server (see TagUsage#list). If multiple properties are specified, only the tags used by policies that match an element of each property match the filter. **Warning:** This class is available as technical preview. It may be changed in a future release.
 type FilterSpec struct {
    Policies map[string]bool
    Tags map[string]bool
    TagTypes map[string]bool
}









func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(FilterSpecBindingType))
    fieldNameMap["filter"] = "Filter"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(SummaryBindingType), reflect.TypeOf([]Summary{}))
}

func listRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["filter.policies"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"com.vmware.vcenter.compute.Policy"}, ""), reflect.TypeOf(map[string]bool{})))
    paramsTypeMap["filter.tags"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"com.vmware.cis.tagging.Tag"}, ""), reflect.TypeOf(map[string]bool{})))
    paramsTypeMap["filter.tag_types"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"com.vmware.vapi.resource"}, ""), reflect.TypeOf(map[string]bool{})))
    queryParams["filter.policies"] = "policies"
    queryParams["filter.tag_types"] = "tag_types"
    queryParams["filter.tags"] = "tags"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "GET",
      "/vcenter/compute/policies/tag-usage",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Unauthorized": 403})
}



func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["policy"] = bindings.NewIdType([]string {"com.vmware.vcenter.compute.Policy"}, "")
    fieldNameMap["policy"] = "Policy"
    fields["tag_type"] = bindings.NewIdType([]string {"com.vmware.vapi.resource"}, "")
    fieldNameMap["tag_type"] = "TagType"
    fields["tag"] = bindings.NewIdType(nil, "tag_type")
    fieldNameMap["tag"] = "Tag"
    fields["tag_name"] = bindings.NewStringType()
    fieldNameMap["tag_name"] = "TagName"
    fields["category_name"] = bindings.NewStringType()
    fieldNameMap["category_name"] = "CategoryName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.compute.policies.tag_usage.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}

func FilterSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["policies"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"com.vmware.vcenter.compute.Policy"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["policies"] = "Policies"
    fields["tags"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"com.vmware.cis.tagging.Tag"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["tags"] = "Tags"
    fields["tag_types"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"com.vmware.vapi.resource"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["tag_types"] = "TagTypes"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.compute.policies.tag_usage.filter_spec",fields, reflect.TypeOf(FilterSpec{}), fieldNameMap, validators)
}


