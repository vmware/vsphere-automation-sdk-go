/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: CounterMetadata.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package counterMetadata

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vstats"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// Counter metadata status.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type CounterEditionStatus string

const (
    // The counter edition is current and is the default.
     CounterEditionStatus_DEFAULT CounterEditionStatus = "DEFAULT"
    // The counter edition is current. This implies a support commitment.
     CounterEditionStatus_CURRENT CounterEditionStatus = "CURRENT"
    // The counter edition is deprecated. It will be decommissioned rather soon.
     CounterEditionStatus_DEPRECATED CounterEditionStatus = "DEPRECATED"
    // The counter edition is experimental. Consumers shouldn't rely on it for the long haul.
     CounterEditionStatus_EXPERIMENTAL CounterEditionStatus = "EXPERIMENTAL"
    // The counter edition was removed.
     CounterEditionStatus_REMOVED CounterEditionStatus = "REMOVED"
)

func (c CounterEditionStatus) CounterEditionStatus() bool {
    switch c {
        case CounterEditionStatus_DEFAULT:
            return true
        case CounterEditionStatus_CURRENT:
            return true
        case CounterEditionStatus_DEPRECATED:
            return true
        case CounterEditionStatus_EXPERIMENTAL:
            return true
        case CounterEditionStatus_REMOVED:
            return true
        default:
            return false
    }
}




// Type of the sampled data.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type SampleType string

const (
    // Raw samples. The value unprocessed as-is sampled.
     SampleType_RAW SampleType = "RAW"
    // Absolute value samples. Represents an actual value of the counter.
     SampleType_ABSOLUTE SampleType = "ABSOLUTE"
    // Fraction samples. Implies range from 0.00 to 1.00.
     SampleType_FRACTION SampleType = "FRACTION"
    // Rate samples. Represents a value that has been normalized over the time period.
     SampleType_RATE SampleType = "RATE"
    // Delta samples. Represents an amount of change for the counter between the current time-stamp and the last time-stamp when the counter was sampled.
     SampleType_DELTA SampleType = "DELTA"
    // Log(n) samples. A natural logarithm of the value.
     SampleType_LOGN SampleType = "LOGN"
)

func (s SampleType) SampleType() bool {
    switch s {
        case SampleType_RAW:
            return true
        case SampleType_ABSOLUTE:
            return true
        case SampleType_FRACTION:
            return true
        case SampleType_RATE:
            return true
        case SampleType_DELTA:
            return true
        case SampleType_LOGN:
            return true
        default:
            return false
    }
}




// Unit used by a metric.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type MetricUnits string

const (
    // Percent
     MetricUnits_PERCENT MetricUnits = "PERCENT"
    // Number
     MetricUnits_NUMBER MetricUnits = "NUMBER"
    // Second
     MetricUnits_SECOND MetricUnits = "SECOND"
    // Hertz
     MetricUnits_HERTZ MetricUnits = "HERTZ"
    // Meter
     MetricUnits_METER MetricUnits = "METER"
    // Meters per second
     MetricUnits_METERSPERSECOND MetricUnits = "METERSPERSECOND"
    // Meters per second squared
     MetricUnits_METERSPERSECONDSQUARED MetricUnits = "METERSPERSECONDSQUARED"
    // Byte
     MetricUnits_BYTE MetricUnits = "BYTE"
    // Bit
     MetricUnits_BIT MetricUnits = "BIT"
    // Bytes per second
     MetricUnits_BYTESPERSECOND MetricUnits = "BYTESPERSECOND"
    // Bits per second
     MetricUnits_BITSPERSECOND MetricUnits = "BITSPERSECOND"
    // Kilogram
     MetricUnits_KILOGRAM MetricUnits = "KILOGRAM"
    // Gram
     MetricUnits_GRAM MetricUnits = "GRAM"
    // Celsius
     MetricUnits_CELSIUS MetricUnits = "CELSIUS"
    // Kelvin
     MetricUnits_KELVIN MetricUnits = "KELVIN"
    // Joule
     MetricUnits_JOULE MetricUnits = "JOULE"
    // Watt
     MetricUnits_WATT MetricUnits = "WATT"
    // Volt
     MetricUnits_VOLT MetricUnits = "VOLT"
    // Ampere
     MetricUnits_AMPERE MetricUnits = "AMPERE"
    // Volt Ampere
     MetricUnits_VOLTAMPERE MetricUnits = "VOLTAMPERE"
    // Candela
     MetricUnits_CANDELA MetricUnits = "CANDELA"
    // Mole
     MetricUnits_MOLE MetricUnits = "MOLE"
)

func (m MetricUnits) MetricUnits() bool {
    switch m {
        case MetricUnits_PERCENT:
            return true
        case MetricUnits_NUMBER:
            return true
        case MetricUnits_SECOND:
            return true
        case MetricUnits_HERTZ:
            return true
        case MetricUnits_METER:
            return true
        case MetricUnits_METERSPERSECOND:
            return true
        case MetricUnits_METERSPERSECONDSQUARED:
            return true
        case MetricUnits_BYTE:
            return true
        case MetricUnits_BIT:
            return true
        case MetricUnits_BYTESPERSECOND:
            return true
        case MetricUnits_BITSPERSECOND:
            return true
        case MetricUnits_KILOGRAM:
            return true
        case MetricUnits_GRAM:
            return true
        case MetricUnits_CELSIUS:
            return true
        case MetricUnits_KELVIN:
            return true
        case MetricUnits_JOULE:
            return true
        case MetricUnits_WATT:
            return true
        case MetricUnits_VOLT:
            return true
        case MetricUnits_AMPERE:
            return true
        case MetricUnits_VOLTAMPERE:
            return true
        case MetricUnits_CANDELA:
            return true
        case MetricUnits_MOLE:
            return true
        default:
            return false
    }
}




// Unit factor of measurement.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type UnitsFactor string

const (
    // Yotta 10^24
     UnitsFactor_YOTTA UnitsFactor = "YOTTA"
    // Zetta 10^21
     UnitsFactor_ZETTA UnitsFactor = "ZETTA"
    // Exa 10^18
     UnitsFactor_EXA UnitsFactor = "EXA"
    // Peta 10^15
     UnitsFactor_PETA UnitsFactor = "PETA"
    // Tera 10^12
     UnitsFactor_TERA UnitsFactor = "TERA"
    // Giga 10^9
     UnitsFactor_GIGA UnitsFactor = "GIGA"
    // Mega 10^6
     UnitsFactor_MEGA UnitsFactor = "MEGA"
    // Kilo 10^3
     UnitsFactor_KILO UnitsFactor = "KILO"
    // Hecto 10^2
     UnitsFactor_HECTO UnitsFactor = "HECTO"
    // Deca 10
     UnitsFactor_DECA UnitsFactor = "DECA"
    // One
     UnitsFactor_ONE UnitsFactor = "ONE"
    // Deci 10^-1
     UnitsFactor_DECI UnitsFactor = "DECI"
    // Centi 10^-2
     UnitsFactor_CENTI UnitsFactor = "CENTI"
    // Milli 10^-3
     UnitsFactor_MILLI UnitsFactor = "MILLI"
    // Micro 10^-6
     UnitsFactor_MICRO UnitsFactor = "MICRO"
    // Nano 10^-9
     UnitsFactor_NANO UnitsFactor = "NANO"
    // Pico 10^-12
     UnitsFactor_PIKO UnitsFactor = "PIKO"
    // Femto 10^-15
     UnitsFactor_FEMTO UnitsFactor = "FEMTO"
    // Atto 10^-18
     UnitsFactor_ATTO UnitsFactor = "ATTO"
    // Zepto 10^-21
     UnitsFactor_ZEPTO UnitsFactor = "ZEPTO"
    // Yocto 10^-24
     UnitsFactor_YOCTO UnitsFactor = "YOCTO"
    // Yobi 2^80, 1024^8
     UnitsFactor_YOBI UnitsFactor = "YOBI"
    // Zebi 2^70, 1024^7
     UnitsFactor_ZEBI UnitsFactor = "ZEBI"
    // Exbi 2^60, 1024^6
     UnitsFactor_EXBI UnitsFactor = "EXBI"
    // Pebi 2^50, 1024^5
     UnitsFactor_PEBI UnitsFactor = "PEBI"
    // Tebi 2^40, 1024^4
     UnitsFactor_TEBI UnitsFactor = "TEBI"
    // Gibi 2^30, 1024^3
     UnitsFactor_GIBI UnitsFactor = "GIBI"
    // Mebi 2^20, 1024^2
     UnitsFactor_MEBI UnitsFactor = "MEBI"
    // Kibi 2^10, 1024
     UnitsFactor_KIBI UnitsFactor = "KIBI"
)

func (u UnitsFactor) UnitsFactor() bool {
    switch u {
        case UnitsFactor_YOTTA:
            return true
        case UnitsFactor_ZETTA:
            return true
        case UnitsFactor_EXA:
            return true
        case UnitsFactor_PETA:
            return true
        case UnitsFactor_TERA:
            return true
        case UnitsFactor_GIGA:
            return true
        case UnitsFactor_MEGA:
            return true
        case UnitsFactor_KILO:
            return true
        case UnitsFactor_HECTO:
            return true
        case UnitsFactor_DECA:
            return true
        case UnitsFactor_ONE:
            return true
        case UnitsFactor_DECI:
            return true
        case UnitsFactor_CENTI:
            return true
        case UnitsFactor_MILLI:
            return true
        case UnitsFactor_MICRO:
            return true
        case UnitsFactor_NANO:
            return true
        case UnitsFactor_PIKO:
            return true
        case UnitsFactor_FEMTO:
            return true
        case UnitsFactor_ATTO:
            return true
        case UnitsFactor_ZEPTO:
            return true
        case UnitsFactor_YOCTO:
            return true
        case UnitsFactor_YOBI:
            return true
        case UnitsFactor_ZEBI:
            return true
        case UnitsFactor_EXBI:
            return true
        case UnitsFactor_PEBI:
            return true
        case UnitsFactor_TEBI:
            return true
        case UnitsFactor_GIBI:
            return true
        case UnitsFactor_MEBI:
            return true
        case UnitsFactor_KIBI:
            return true
        default:
            return false
    }
}





// The ``Info`` class contains information about CounterMetadata. It represents edition of the Counter.
 type Info struct {
    Cid string
    Mid string
    Status CounterEditionStatus
    Type_ SampleType
    Units MetricUnits
    Scale *UnitsFactor
    UserInfo *vstats.UserInfo
    Pid *string
}






// The ``FilterSpec`` class is used to filter the counter metadata list.
 type FilterSpec struct {
    Status *CounterEditionStatus
}









func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cid"] = bindings.NewIdType([]string {"com.vmware.vstats.model.Counter"}, "")
    fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(FilterSpecBindingType))
    fieldNameMap["cid"] = "Cid"
    fieldNameMap["filter"] = "Filter"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(InfoBindingType), reflect.TypeOf([]Info{}))
}

func listRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["filter.status"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vstats.counter_metadata.counter_edition_status", reflect.TypeOf(CounterEditionStatus(CounterEditionStatus_DEFAULT))))
    paramsTypeMap["cid"] = bindings.NewIdType([]string {"com.vmware.vstats.model.Counter"}, "")
    paramsTypeMap["cid"] = bindings.NewIdType([]string {"com.vmware.vstats.model.Counter"}, "")
    pathParams["cid"] = "cid"
    queryParams["filter.status"] = "status"
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
      "/stats/counters/{cid}/metadata",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}


func getDefaultInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cid"] = bindings.NewIdType([]string {"com.vmware.vstats.model.Counter"}, "")
    fieldNameMap["cid"] = "Cid"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getDefaultOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(InfoBindingType), reflect.TypeOf([]Info{}))
}

func getDefaultRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["cid"] = bindings.NewIdType([]string {"com.vmware.vstats.model.Counter"}, "")
    paramsTypeMap["cid"] = bindings.NewIdType([]string {"com.vmware.vstats.model.Counter"}, "")
    pathParams["cid"] = "cid"
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
      "/stats/counters/{cid}/metadata/default",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cid"] = bindings.NewIdType([]string {"com.vmware.vstats.model.Counter"}, "")
    fields["mid"] = bindings.NewIdType([]string {"com.vmware.vstats.model.CounterMetadata"}, "")
    fieldNameMap["cid"] = "Cid"
    fieldNameMap["mid"] = "Mid"
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
    paramsTypeMap["mid"] = bindings.NewIdType([]string {"com.vmware.vstats.model.CounterMetadata"}, "")
    paramsTypeMap["cid"] = bindings.NewIdType([]string {"com.vmware.vstats.model.Counter"}, "")
    paramsTypeMap["cid"] = bindings.NewIdType([]string {"com.vmware.vstats.model.Counter"}, "")
    paramsTypeMap["mid"] = bindings.NewIdType([]string {"com.vmware.vstats.model.CounterMetadata"}, "")
    pathParams["mid"] = "mid"
    pathParams["cid"] = "cid"
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
      "/stats/counters/{cid}/metadata/{mid}",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}



func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cid"] = bindings.NewIdType([]string {"com.vmware.vstats.model.Counter"}, "")
    fieldNameMap["cid"] = "Cid"
    fields["mid"] = bindings.NewIdType([]string {"com.vmware.vstats.model.CounterMetadata"}, "")
    fieldNameMap["mid"] = "Mid"
    fields["status"] = bindings.NewEnumType("com.vmware.vstats.counter_metadata.counter_edition_status", reflect.TypeOf(CounterEditionStatus(CounterEditionStatus_DEFAULT)))
    fieldNameMap["status"] = "Status"
    fields["type"] = bindings.NewEnumType("com.vmware.vstats.counter_metadata.sample_type", reflect.TypeOf(SampleType(SampleType_RAW)))
    fieldNameMap["type"] = "Type_"
    fields["units"] = bindings.NewEnumType("com.vmware.vstats.counter_metadata.metric_units", reflect.TypeOf(MetricUnits(MetricUnits_PERCENT)))
    fieldNameMap["units"] = "Units"
    fields["scale"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vstats.counter_metadata.units_factor", reflect.TypeOf(UnitsFactor(UnitsFactor_YOTTA))))
    fieldNameMap["scale"] = "Scale"
    fields["user_info"] = bindings.NewOptionalType(bindings.NewReferenceType(vstats.UserInfoBindingType))
    fieldNameMap["user_info"] = "UserInfo"
    fields["pid"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.vstats.model.Provider"}, ""))
    fieldNameMap["pid"] = "Pid"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.counter_metadata.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func FilterSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["status"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vstats.counter_metadata.counter_edition_status", reflect.TypeOf(CounterEditionStatus(CounterEditionStatus_DEFAULT))))
    fieldNameMap["status"] = "Status"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.counter_metadata.filter_spec",fields, reflect.TypeOf(FilterSpec{}), fieldNameMap, validators)
}


