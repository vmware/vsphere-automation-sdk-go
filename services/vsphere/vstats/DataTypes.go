/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Data.
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



// The ``DataPoint`` class is an instance of a measurement or stat. A data point is comprised of a Counter, CounterMetadata, Resource, timestamp and value. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type DataDataPoint struct {
    // Counter Id. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Cid string
    // CounterMetadata Id. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Mid string
    // Resource Id. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Rid string
    // Timestamp for the data point. format: 64-bit integer. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Ts int64
    // Stat value. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Val float64
}

// The ``FilterSpec`` class contains properties used to filter the results when listing DataPoint. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type DataFilterSpec struct {
    // Start of a time window (included), timestamp in seconds UTC. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Start *int64
    // End of a time window (excluded), timestamp in seconds UTC. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	End *int64
    // Counter ID. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Cid *string
    // Metric name. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Metric *string
    // List of Resource types. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Types []string
    // Resources to include in the query. Each resource is specified through a composite string that follows the following format. 
    //
    //  ``type.<resource type>[.<scheme>]=<resource id>`` 
    //
    //  **resource type** specifies the type of resource for example ``VM``, ``VCPU`` etc. 
    //
    //  **scheme** is an optional element to disambiguate the resource as needed for example to differentiate managed object id from ``uuid``. 
    //
    //  **resource id** is the unique resource identifier value for example: ``vm-41`` 
    //
    //  Example values include: ``type.VM=vm-41``, ``type.VCPU=1``, ``type.VM.moid=vm-41``. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Resources []string
    // Directs the server to order the returned data. Passing a value of ``DEFAULT`` will apply default ordering of the results that makes them easier for consumption. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Order *string
    // Used to retrieve paged data for larger result sets. The value of this token is generated by server and returned as ``next`` property in the result of ``queryDataPoints()`` methods. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Page *string
}

// The ``DataPointsResult`` class contains properties used to return DataPoints. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type DataDataPointsResult struct {
    // List of DataPoints received. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	DataPoints []DataDataPoint
    // The ``next`` property is a token used to retrieve paged data for larger result sets. This is opaque token generated by the server. It is to be sent in the DataFilterSpec#page property to issue a subsequent call to the query method for retrieving results that did not fit the current page. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Next *string
}



func dataQueryDataPointsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(DataFilterSpecBindingType))
	fieldNameMap["filter"] = "Filter"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func dataQueryDataPointsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(DataDataPointsResultBindingType)
}

func dataQueryDataPointsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(DataFilterSpecBindingType))
	fieldNameMap["filter"] = "Filter"
	paramsTypeMap["filter.cid"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.vstats.model.Counter"}, ""))
	paramsTypeMap["filter.types"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewIdType([]string{"com.vmware.vstats.model.RsrcType"}, ""), reflect.TypeOf([]string{})))
	paramsTypeMap["filter.metric"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.vstats.model.Metric"}, ""))
	paramsTypeMap["filter.end"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["filter.page"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["filter.start"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["filter.resources"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	paramsTypeMap["filter.order"] = bindings.NewOptionalType(bindings.NewStringType())
	queryParams["filter.types"] = "types"
	queryParams["filter.metric"] = "metric"
	queryParams["filter.start"] = "start"
	queryParams["filter.end"] = "end"
	queryParams["filter.page"] = "page"
	queryParams["filter.resources"] = "rsrcs"
	queryParams["filter.cid"] = "cid"
	queryParams["filter.order"] = "order"
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
		"/stats/data/dp",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Error": 500,"InvalidArgument": 400,"Unauthenticated": 401,"Unauthorized": 403})
}


func DataDataPointBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cid"] = bindings.NewIdType([]string{"com.vmware.vstats.model.Counter"}, "")
	fieldNameMap["cid"] = "Cid"
	fields["mid"] = bindings.NewIdType([]string{"com.vmware.vstats.model.CounterMetadata"}, "")
	fieldNameMap["mid"] = "Mid"
	fields["rid"] = bindings.NewStringType()
	fieldNameMap["rid"] = "Rid"
	fields["ts"] = bindings.NewIntegerType()
	fieldNameMap["ts"] = "Ts"
	fields["val"] = bindings.NewDoubleType()
	fieldNameMap["val"] = "Val"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vstats.data.data_point", fields, reflect.TypeOf(DataDataPoint{}), fieldNameMap, validators)
}

func DataFilterSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["start"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["start"] = "Start"
	fields["end"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["end"] = "End"
	fields["cid"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.vstats.model.Counter"}, ""))
	fieldNameMap["cid"] = "Cid"
	fields["metric"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.vstats.model.Metric"}, ""))
	fieldNameMap["metric"] = "Metric"
	fields["types"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewIdType([]string{"com.vmware.vstats.model.RsrcType"}, ""), reflect.TypeOf([]string{})))
	fieldNameMap["types"] = "Types"
	fields["resources"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["resources"] = "Resources"
	fields["order"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["order"] = "Order"
	fields["page"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["page"] = "Page"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vstats.data.filter_spec", fields, reflect.TypeOf(DataFilterSpec{}), fieldNameMap, validators)
}

func DataDataPointsResultBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["data_points"] = bindings.NewListType(bindings.NewReferenceType(DataDataPointBindingType), reflect.TypeOf([]DataDataPoint{}))
	fieldNameMap["data_points"] = "DataPoints"
	fields["next"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["next"] = "Next"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vstats.data.data_points_result", fields, reflect.TypeOf(DataDataPointsResult{}), fieldNameMap, validators)
}

