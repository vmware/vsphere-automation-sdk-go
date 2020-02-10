/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Counters.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package vstats

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// Resource type for counters. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
const Counters_RESOURCE_TYPE = "com.vmware.vstats.model.Counter"


// The ``Info`` class contains the counter information. Counter is derived from metric. It applies the metric to a particular class of a resource. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type CountersInfo struct {
    // Counter Id. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Cid string
    // A metric is typically human-legible ASCII/English name of a measurable aspect, for example "length" and "temperature". It is not internationalizable. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Metric string
    // Describes formally how to address (or identify) the resources the counter could be bound to (or instantiated for). **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	ResourceAddressSchema string
}

// The ``FilterSpec`` class contains fields that can be used to filter list of counters. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type CountersFilterSpec struct {
    // Counter identifier that will be listed. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Cid *string
    // Resource type filter. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Types []string
    // Metric for which counters will be listed. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Metric *string
}



func countersListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(CountersFilterSpecBindingType))
	fieldNameMap["filter"] = "Filter"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func countersListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(CountersInfoBindingType), reflect.TypeOf([]CountersInfo{}))
}

func countersListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(CountersFilterSpecBindingType))
	fieldNameMap["filter"] = "Filter"
	paramsTypeMap["filter.cid"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.vstats.model.Counter"}, ""))
	paramsTypeMap["filter.types"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewIdType([]string{"com.vmware.vstats.model.RsrcType"}, ""), reflect.TypeOf([]string{})))
	paramsTypeMap["filter.metric"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.vstats.model.Metric"}, ""))
	queryParams["filter.types"] = "types"
	queryParams["filter.metric"] = "metric"
	queryParams["filter.cid"] = "cid"
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
		"GET",
		"/stats/counters",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Error": 500,"InvalidArgument": 400,"Unauthenticated": 401,"Unauthorized": 403})
}

func countersGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cid"] = bindings.NewIdType([]string{"com.vmware.vstats.model.Counter"}, "")
	fieldNameMap["cid"] = "Cid"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func countersGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(CountersInfoBindingType)
}

func countersGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["cid"] = bindings.NewIdType([]string{"com.vmware.vstats.model.Counter"}, "")
	fieldNameMap["cid"] = "Cid"
	paramsTypeMap["cid"] = bindings.NewIdType([]string{"com.vmware.vstats.model.Counter"}, "")
	paramsTypeMap["cid"] = bindings.NewIdType([]string{"com.vmware.vstats.model.Counter"}, "")
	pathParams["cid"] = "cid"
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
		"GET",
		"/stats/counters/{cid}",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}


func CountersInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cid"] = bindings.NewIdType([]string{"com.vmware.vstats.model.Counter"}, "")
	fieldNameMap["cid"] = "Cid"
	fields["metric"] = bindings.NewIdType([]string{"com.vmware.vstats.model.Metric"}, "")
	fieldNameMap["metric"] = "Metric"
	fields["resource_address_schema"] = bindings.NewIdType([]string{"com.vmware.vstats.model.RsrcAddrSchema"}, "")
	fieldNameMap["resource_address_schema"] = "ResourceAddressSchema"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vstats.counters.info", fields, reflect.TypeOf(CountersInfo{}), fieldNameMap, validators)
}

func CountersFilterSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cid"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.vstats.model.Counter"}, ""))
	fieldNameMap["cid"] = "Cid"
	fields["types"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewIdType([]string{"com.vmware.vstats.model.RsrcType"}, ""), reflect.TypeOf([]string{})))
	fieldNameMap["types"] = "Types"
	fields["metric"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.vstats.model.Metric"}, ""))
	fieldNameMap["metric"] = "Metric"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vstats.counters.filter_spec", fields, reflect.TypeOf(CountersFilterSpec{}), fieldNameMap, validators)
}

