/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Telemetry.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package telemetry

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// Type of telemetry object.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type TelemetryType string

const (
    // Filter for telemetry counters.
     TelemetryType_COUNTERS TelemetryType = "COUNTERS"
    // Filter for telemetry timers.
     TelemetryType_TIMERS TelemetryType = "TIMERS"
    // Filter for telemetry gauges.
     TelemetryType_GAUGES TelemetryType = "GAUGES"
    // Filter for telemetry meters.
     TelemetryType_METERS TelemetryType = "METERS"
    // Filter for telemetry counters.
     TelemetryType_HISTOGRAMS TelemetryType = "HISTOGRAMS"
)

func (t TelemetryType) TelemetryType() bool {
    switch t {
        case TelemetryType_COUNTERS:
            return true
        case TelemetryType_TIMERS:
            return true
        case TelemetryType_GAUGES:
            return true
        case TelemetryType_METERS:
            return true
        case TelemetryType_HISTOGRAMS:
            return true
        default:
            return false
    }
}





// The ``TelemetryCounter`` class specifies vStats monotonically increasing integer.
 type TelemetryCounter struct {
    Name string
    Count int64
}






// The ``TelemetryGauge`` class specifies vStats telemetry-gauge, both integer and floating point are presented here.
 type TelemetryGauge struct {
    Name string
    Value float64
}






// The ``TelemetryHistogram`` class specifies vStats telemetry-histogram.
 type TelemetryHistogram struct {
    Name string
    Count int64
    Min int64
    Max int64
    Mean float64
    Stddev float64
    Median float64
    P75 float64
    P95 float64
    P99 float64
    P999 float64
}






// The ``TelemetryMeter`` class specifies vStats telemetry-meter.
 type TelemetryMeter struct {
    Name string
    Count int64
    Min1Rate float64
    Min5Rate float64
    Min15Rate float64
    MeanRate float64
}






// The ``TelemetryTimer`` class specifies vStats telemetry-timer. Units are in nanoseconds.
 type TelemetryTimer struct {
    Name string
    Count int64
    Min float64
    Max float64
    Mean float64
    Stddev float64
    Median float64
    P75 float64
    P95 float64
    P99 float64
    P999 float64
    Min1Rate float64
    Min5Rate float64
    Min15Rate float64
    MeanRate float64
}






// The ``Info`` class contains overall telemetry Information.
 type Info struct {
    Counters []TelemetryCounter
    Timers []TelemetryTimer
    Gauges []TelemetryGauge
    Meters []TelemetryMeter
    Histograms []TelemetryHistogram
}






// ``FilterSpec`` class describes filter criteria for telemetry.
 type FilterSpec struct {
    Prefix *string
    Types map[TelemetryType]bool
}









func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(FilterSpecBindingType))
    fieldNameMap["filter"] = "Filter"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewReferenceType(InfoBindingType)
}

func getRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["filter.prefix"] = bindings.NewOptionalType(bindings.NewStringType())
    paramsTypeMap["filter.types"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewEnumType("com.vmware.vstats.telemetry.telemetry_type", reflect.TypeOf(TelemetryType(TelemetryType_COUNTERS))), reflect.TypeOf(map[TelemetryType]bool{})))
    queryParams["filter.types"] = "types"
    queryParams["filter.prefix"] = "prefix"
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
      "/stats/telemetry",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}



func TelemetryCounterBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["count"] = bindings.NewIntegerType()
    fieldNameMap["count"] = "Count"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.telemetry.telemetry_counter",fields, reflect.TypeOf(TelemetryCounter{}), fieldNameMap, validators)
}

func TelemetryGaugeBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["value"] = bindings.NewDoubleType()
    fieldNameMap["value"] = "Value"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.telemetry.telemetry_gauge",fields, reflect.TypeOf(TelemetryGauge{}), fieldNameMap, validators)
}

func TelemetryHistogramBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["count"] = bindings.NewIntegerType()
    fieldNameMap["count"] = "Count"
    fields["min"] = bindings.NewIntegerType()
    fieldNameMap["min"] = "Min"
    fields["max"] = bindings.NewIntegerType()
    fieldNameMap["max"] = "Max"
    fields["mean"] = bindings.NewDoubleType()
    fieldNameMap["mean"] = "Mean"
    fields["stddev"] = bindings.NewDoubleType()
    fieldNameMap["stddev"] = "Stddev"
    fields["median"] = bindings.NewDoubleType()
    fieldNameMap["median"] = "Median"
    fields["p75"] = bindings.NewDoubleType()
    fieldNameMap["p75"] = "P75"
    fields["p95"] = bindings.NewDoubleType()
    fieldNameMap["p95"] = "P95"
    fields["p99"] = bindings.NewDoubleType()
    fieldNameMap["p99"] = "P99"
    fields["p999"] = bindings.NewDoubleType()
    fieldNameMap["p999"] = "P999"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.telemetry.telemetry_histogram",fields, reflect.TypeOf(TelemetryHistogram{}), fieldNameMap, validators)
}

func TelemetryMeterBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["count"] = bindings.NewIntegerType()
    fieldNameMap["count"] = "Count"
    fields["min1_rate"] = bindings.NewDoubleType()
    fieldNameMap["min1_rate"] = "Min1Rate"
    fields["min5_rate"] = bindings.NewDoubleType()
    fieldNameMap["min5_rate"] = "Min5Rate"
    fields["min15_rate"] = bindings.NewDoubleType()
    fieldNameMap["min15_rate"] = "Min15Rate"
    fields["mean_rate"] = bindings.NewDoubleType()
    fieldNameMap["mean_rate"] = "MeanRate"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.telemetry.telemetry_meter",fields, reflect.TypeOf(TelemetryMeter{}), fieldNameMap, validators)
}

func TelemetryTimerBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["count"] = bindings.NewIntegerType()
    fieldNameMap["count"] = "Count"
    fields["min"] = bindings.NewDoubleType()
    fieldNameMap["min"] = "Min"
    fields["max"] = bindings.NewDoubleType()
    fieldNameMap["max"] = "Max"
    fields["mean"] = bindings.NewDoubleType()
    fieldNameMap["mean"] = "Mean"
    fields["stddev"] = bindings.NewDoubleType()
    fieldNameMap["stddev"] = "Stddev"
    fields["median"] = bindings.NewDoubleType()
    fieldNameMap["median"] = "Median"
    fields["p75"] = bindings.NewDoubleType()
    fieldNameMap["p75"] = "P75"
    fields["p95"] = bindings.NewDoubleType()
    fieldNameMap["p95"] = "P95"
    fields["p99"] = bindings.NewDoubleType()
    fieldNameMap["p99"] = "P99"
    fields["p999"] = bindings.NewDoubleType()
    fieldNameMap["p999"] = "P999"
    fields["min1_rate"] = bindings.NewDoubleType()
    fieldNameMap["min1_rate"] = "Min1Rate"
    fields["min5_rate"] = bindings.NewDoubleType()
    fieldNameMap["min5_rate"] = "Min5Rate"
    fields["min15_rate"] = bindings.NewDoubleType()
    fieldNameMap["min15_rate"] = "Min15Rate"
    fields["mean_rate"] = bindings.NewDoubleType()
    fieldNameMap["mean_rate"] = "MeanRate"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.telemetry.telemetry_timer",fields, reflect.TypeOf(TelemetryTimer{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["counters"] = bindings.NewListType(bindings.NewReferenceType(TelemetryCounterBindingType), reflect.TypeOf([]TelemetryCounter{}))
    fieldNameMap["counters"] = "Counters"
    fields["timers"] = bindings.NewListType(bindings.NewReferenceType(TelemetryTimerBindingType), reflect.TypeOf([]TelemetryTimer{}))
    fieldNameMap["timers"] = "Timers"
    fields["gauges"] = bindings.NewListType(bindings.NewReferenceType(TelemetryGaugeBindingType), reflect.TypeOf([]TelemetryGauge{}))
    fieldNameMap["gauges"] = "Gauges"
    fields["meters"] = bindings.NewListType(bindings.NewReferenceType(TelemetryMeterBindingType), reflect.TypeOf([]TelemetryMeter{}))
    fieldNameMap["meters"] = "Meters"
    fields["histograms"] = bindings.NewListType(bindings.NewReferenceType(TelemetryHistogramBindingType), reflect.TypeOf([]TelemetryHistogram{}))
    fieldNameMap["histograms"] = "Histograms"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.telemetry.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func FilterSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["prefix"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["prefix"] = "Prefix"
    fields["types"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewEnumType("com.vmware.vstats.telemetry.telemetry_type", reflect.TypeOf(TelemetryType(TelemetryType_COUNTERS))), reflect.TypeOf(map[TelemetryType]bool{})))
    fieldNameMap["types"] = "Types"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.telemetry.filter_spec",fields, reflect.TypeOf(FilterSpec{}), fieldNameMap, validators)
}


