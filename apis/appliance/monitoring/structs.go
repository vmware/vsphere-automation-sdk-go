/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Monitoring.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package monitoring

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
    "time"
)



// ``FunctionType`` enumeration class Defines aggregation function
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type FunctionType string

const (
    // Aggregation takes count per period (sum)
     FunctionType_COUNT FunctionType = "COUNT"
    // Aggregation takes maximums per period
     FunctionType_MAX FunctionType = "MAX"
    // Aggregation takes average per period
     FunctionType_AVG FunctionType = "AVG"
    // Aggregation takes minimums per period
     FunctionType_MIN FunctionType = "MIN"
)

func (f FunctionType) FunctionType() bool {
    switch f {
        case FunctionType_COUNT:
            return true
        case FunctionType_MAX:
            return true
        case FunctionType_AVG:
            return true
        case FunctionType_MIN:
            return true
        default:
            return false
    }
}




// ``IntervalType`` enumeration class Defines interval between the values in hours and mins, for which aggregation will apply
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type IntervalType string

const (
    // Thirty minutes interval between values. One week is 336 values.
     IntervalType_MINUTES30 IntervalType = "MINUTES30"
    // Two hours interval between values. One month has 360 values.
     IntervalType_HOURS2 IntervalType = "HOURS2"
    // Five minutes interval between values (finest). One day would have 288 values, one week is 2016.
     IntervalType_MINUTES5 IntervalType = "MINUTES5"
    // 24 hours interval between values. One year has 365 values.
     IntervalType_DAY1 IntervalType = "DAY1"
    // Six hour interval between values. One quarter is 360 values.
     IntervalType_HOURS6 IntervalType = "HOURS6"
)

func (i IntervalType) IntervalType() bool {
    switch i {
        case IntervalType_MINUTES30:
            return true
        case IntervalType_HOURS2:
            return true
        case IntervalType_MINUTES5:
            return true
        case IntervalType_DAY1:
            return true
        case IntervalType_HOURS6:
            return true
        default:
            return false
    }
}





// ``MonitoredItemData`` class Structure representing monitored item data.
 type MonitoredItemData struct {
    Name string
    Interval IntervalType
    Function FunctionType
    StartTime time.Time
    EndTime time.Time
    Data []string
}






// ``MonitoredItemDataRequest`` class Structure representing requested monitored item data.
 type MonitoredItemDataRequest struct {
    Names []string
    Interval IntervalType
    Function FunctionType
    StartTime time.Time
    EndTime time.Time
}






// ``MonitoredItem`` class Structure representing requested monitored item data.
 type MonitoredItem struct {
    Id string
    Name string
    Units string
    Category string
    Instance string
    Description string
}









func queryInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["item"] = bindings.NewReferenceType(MonitoredItemDataRequestBindingType)
    fieldNameMap["item"] = "Item"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func queryOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(MonitoredItemDataBindingType), reflect.TypeOf([]MonitoredItemData{}))
}

func queryRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500})
}


func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(MonitoredItemBindingType), reflect.TypeOf([]MonitoredItem{}))
}

func listRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["stat_id"] = bindings.NewIdType([]string {"com.vmware.appliance.monitoring"}, "")
    fieldNameMap["stat_id"] = "StatId"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewReferenceType(MonitoredItemBindingType)
}

func getRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500})
}



func MonitoredItemDataBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewIdType([]string {"com.vmware.appliance.monitoring"}, "")
    fieldNameMap["name"] = "Name"
    fields["interval"] = bindings.NewEnumType("com.vmware.appliance.monitoring.interval_type", reflect.TypeOf(IntervalType(IntervalType_MINUTES30)))
    fieldNameMap["interval"] = "Interval"
    fields["function"] = bindings.NewEnumType("com.vmware.appliance.monitoring.function_type", reflect.TypeOf(FunctionType(FunctionType_COUNT)))
    fieldNameMap["function"] = "Function"
    fields["start_time"] = bindings.NewDateTimeType()
    fieldNameMap["start_time"] = "StartTime"
    fields["end_time"] = bindings.NewDateTimeType()
    fieldNameMap["end_time"] = "EndTime"
    fields["data"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["data"] = "Data"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.monitoring.monitored_item_data",fields, reflect.TypeOf(MonitoredItemData{}), fieldNameMap, validators)
}

func MonitoredItemDataRequestBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["names"] = bindings.NewListType(bindings.NewIdType([]string {"com.vmware.appliance.monitoring"}, ""), reflect.TypeOf([]string{}))
    fieldNameMap["names"] = "Names"
    fields["interval"] = bindings.NewEnumType("com.vmware.appliance.monitoring.interval_type", reflect.TypeOf(IntervalType(IntervalType_MINUTES30)))
    fieldNameMap["interval"] = "Interval"
    fields["function"] = bindings.NewEnumType("com.vmware.appliance.monitoring.function_type", reflect.TypeOf(FunctionType(FunctionType_COUNT)))
    fieldNameMap["function"] = "Function"
    fields["start_time"] = bindings.NewDateTimeType()
    fieldNameMap["start_time"] = "StartTime"
    fields["end_time"] = bindings.NewDateTimeType()
    fieldNameMap["end_time"] = "EndTime"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.monitoring.monitored_item_data_request",fields, reflect.TypeOf(MonitoredItemDataRequest{}), fieldNameMap, validators)
}

func MonitoredItemBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.appliance.monitoring"}, "")
    fieldNameMap["id"] = "Id"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["units"] = bindings.NewStringType()
    fieldNameMap["units"] = "Units"
    fields["category"] = bindings.NewStringType()
    fieldNameMap["category"] = "Category"
    fields["instance"] = bindings.NewStringType()
    fieldNameMap["instance"] = "Instance"
    fields["description"] = bindings.NewStringType()
    fieldNameMap["description"] = "Description"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.monitoring.monitored_item",fields, reflect.TypeOf(MonitoredItem{}), fieldNameMap, validators)
}


