/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Data.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package data

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)




// The ``DataPoint`` class is an instance of a measurement or stat. A data point is comprised of a Counter, CounterMetadata, Resource, timestamp and value.
 type DataPoint struct {
    Cid string
    Mid string
    Rid string
    Ts int64
    Val float64
}






// The ``FilterSpec`` class contains properties used to filter the results when listing DataPoint.
 type FilterSpec struct {
    Start *int64
    End *int64
    Cid *string
    Metric *string
    Types []string
    Resources []string
    Order *string
    Page *string
}






// The ``DataPointsResult`` class contains properties used to return DataPoints.
 type DataPointsResult struct {
    DataPoints []DataPoint
    Next *string
}









func queryDataPointsInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(FilterSpecBindingType))
    fieldNameMap["filter"] = "Filter"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func queryDataPointsOutputType() bindings.BindingType {
    return bindings.NewReferenceType(DataPointsResultBindingType)
}

func queryDataPointsRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["filter.cid"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.vstats.model.Counter"}, ""))
    paramsTypeMap["filter.types"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewIdType([]string {"com.vmware.vstats.model.RsrcType"}, ""), reflect.TypeOf([]string{})))
    paramsTypeMap["filter.metric"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.vstats.model.Metric"}, ""))
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
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "GET",
      "/stats/data/dp",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"Unauthenticated": 401,"Unauthorized": 403})
}



func DataPointBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cid"] = bindings.NewIdType([]string {"com.vmware.vstats.model.Counter"}, "")
    fieldNameMap["cid"] = "Cid"
    fields["mid"] = bindings.NewIdType([]string {"com.vmware.vstats.model.CounterMetadata"}, "")
    fieldNameMap["mid"] = "Mid"
    fields["rid"] = bindings.NewStringType()
    fieldNameMap["rid"] = "Rid"
    fields["ts"] = bindings.NewIntegerType()
    fieldNameMap["ts"] = "Ts"
    fields["val"] = bindings.NewDoubleType()
    fieldNameMap["val"] = "Val"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.data.data_point",fields, reflect.TypeOf(DataPoint{}), fieldNameMap, validators)
}

func FilterSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["start"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["start"] = "Start"
    fields["end"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["end"] = "End"
    fields["cid"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.vstats.model.Counter"}, ""))
    fieldNameMap["cid"] = "Cid"
    fields["metric"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.vstats.model.Metric"}, ""))
    fieldNameMap["metric"] = "Metric"
    fields["types"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewIdType([]string {"com.vmware.vstats.model.RsrcType"}, ""), reflect.TypeOf([]string{})))
    fieldNameMap["types"] = "Types"
    fields["resources"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["resources"] = "Resources"
    fields["order"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["order"] = "Order"
    fields["page"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["page"] = "Page"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.data.filter_spec",fields, reflect.TypeOf(FilterSpec{}), fieldNameMap, validators)
}

func DataPointsResultBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["data_points"] = bindings.NewListType(bindings.NewReferenceType(DataPointBindingType), reflect.TypeOf([]DataPoint{}))
    fieldNameMap["data_points"] = "DataPoints"
    fields["next"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["next"] = "Next"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.data.data_points_result",fields, reflect.TypeOf(DataPointsResult{}), fieldNameMap, validators)
}


