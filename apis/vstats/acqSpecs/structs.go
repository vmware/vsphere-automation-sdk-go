/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: AcqSpecs.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package acqSpecs

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vstats"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// Resource type for acquisition specifications.
const AcqSpecs_RESOURCE_TYPE = "com.vmware.vstats.model.AcqSpec"


// Describes the status of an Acquisition Specification.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Status string

const (
    // The acquisition specification is enabled.
     Status_ENABLED Status = "ENABLED"
    // The acquisition specification is disabled.
     Status_DISABLED Status = "DISABLED"
    // The acquisition specification is expired.
     Status_EXPIRED Status = "EXPIRED"
)

func (s Status) Status() bool {
    switch s {
        case Status_ENABLED:
            return true
        case Status_DISABLED:
            return true
        case Status_EXPIRED:
            return true
        default:
            return false
    }
}





// The ``CounterSpec`` class designates a counter or counter set in an acquisition specification.
 type CounterSpec struct {
    CidMid *vstats.CidMid
    SetId *string
}






// The ``CreateSpec`` class contains information for a new data acquisition specification.
 type CreateSpec struct {
    Counters CounterSpec
    Resources []vstats.RsrcId
    Interval *int64
    Expiration *int64
    Memo_ *string
}






// The ``Info`` class is the information about an acquisition specification. It specifies the statistical data that should be collected at desired sampling rates. It designates the resources and their counters which should be sampled, and a desired sampling rate.
 type Info struct {
    Id string
    Counters CounterSpec
    Resources []vstats.RsrcId
    Interval *int64
    Expiration *int64
    Memo_ string
    Status Status
}






// The ``ListResult`` class contains properties used to return the acquisition specifications.
 type ListResult struct {
    AcqSpecs []Info
    Next *string
}






// The ``FilterSpec`` class contains properties used to filter the results when listing acquisition specifications.
 type FilterSpec struct {
    Page *string
}






// The ``UpdateSpec`` class contains properties that can be updated in an acquisition specification.
 type UpdateSpec struct {
    Counters *CounterSpec
    Resources []vstats.RsrcId
    Interval *int64
    Expiration *int64
    Memo_ *string
}









func createInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["acq_spec"] = bindings.NewReferenceType(CreateSpecBindingType)
    fieldNameMap["acq_spec"] = "AcqSpec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func createOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"com.vmware.vstats.model.AcqSpec"}, "")
}

func createRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["acq_spec"] = bindings.NewReferenceType(CreateSpecBindingType)
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "acq_spec",
      "POST",
      "/stats/acq-specs",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"Unauthenticated": 401,"Unauthorized": 403})
}


func deleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.vstats.model.AcqSpec"}, "")
    fieldNameMap["id"] = "Id"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func deleteOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func deleteRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["id"] = bindings.NewIdType([]string {"com.vmware.vstats.model.AcqSpec"}, "")
    paramsTypeMap["id"] = bindings.NewIdType([]string {"com.vmware.vstats.model.AcqSpec"}, "")
    pathParams["id"] = "id"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "DELETE",
      "/stats/acq-specs/{id}",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
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
    return bindings.NewReferenceType(ListResultBindingType)
}

func listRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["filter.page"] = bindings.NewOptionalType(bindings.NewStringType())
    queryParams["filter.page"] = "page"
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
      "/stats/acq-specs",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"Unauthenticated": 401,"Unauthorized": 403})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.vstats.model.AcqSpec"}, "")
    fieldNameMap["id"] = "Id"
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
    paramsTypeMap["id"] = bindings.NewIdType([]string {"com.vmware.vstats.model.AcqSpec"}, "")
    paramsTypeMap["id"] = bindings.NewIdType([]string {"com.vmware.vstats.model.AcqSpec"}, "")
    pathParams["id"] = "id"
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
      "/stats/acq-specs/{id}",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}


func updateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.vstats.model.AcqSpec"}, "")
    fields["acq_spec"] = bindings.NewReferenceType(UpdateSpecBindingType)
    fieldNameMap["id"] = "Id"
    fieldNameMap["acq_spec"] = "AcqSpec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func updateOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func updateRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["id"] = bindings.NewIdType([]string {"com.vmware.vstats.model.AcqSpec"}, "")
    paramsTypeMap["acq_spec"] = bindings.NewReferenceType(UpdateSpecBindingType)
    paramsTypeMap["id"] = bindings.NewIdType([]string {"com.vmware.vstats.model.AcqSpec"}, "")
    pathParams["id"] = "id"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "acq_spec",
      "PATCH",
      "/stats/acq-specs/{id}",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}



func CounterSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cid_mid"] = bindings.NewOptionalType(bindings.NewReferenceType(vstats.CidMidBindingType))
    fieldNameMap["cid_mid"] = "CidMid"
    fields["set_id"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.vstats.model.CounterSet"}, ""))
    fieldNameMap["set_id"] = "SetId"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.acq_specs.counter_spec",fields, reflect.TypeOf(CounterSpec{}), fieldNameMap, validators)
}

func CreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["counters"] = bindings.NewReferenceType(CounterSpecBindingType)
    fieldNameMap["counters"] = "Counters"
    fields["resources"] = bindings.NewListType(bindings.NewReferenceType(vstats.RsrcIdBindingType), reflect.TypeOf([]vstats.RsrcId{}))
    fieldNameMap["resources"] = "Resources"
    fields["interval"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["interval"] = "Interval"
    fields["expiration"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["expiration"] = "Expiration"
    fields["memo_"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["memo_"] = "Memo_"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.acq_specs.create_spec",fields, reflect.TypeOf(CreateSpec{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.vstats.model.AcqSpec"}, "")
    fieldNameMap["id"] = "Id"
    fields["counters"] = bindings.NewReferenceType(CounterSpecBindingType)
    fieldNameMap["counters"] = "Counters"
    fields["resources"] = bindings.NewListType(bindings.NewReferenceType(vstats.RsrcIdBindingType), reflect.TypeOf([]vstats.RsrcId{}))
    fieldNameMap["resources"] = "Resources"
    fields["interval"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["interval"] = "Interval"
    fields["expiration"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["expiration"] = "Expiration"
    fields["memo_"] = bindings.NewStringType()
    fieldNameMap["memo_"] = "Memo_"
    fields["status"] = bindings.NewEnumType("com.vmware.vstats.acq_specs.status", reflect.TypeOf(Status(Status_ENABLED)))
    fieldNameMap["status"] = "Status"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.acq_specs.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func ListResultBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["acq_specs"] = bindings.NewListType(bindings.NewReferenceType(InfoBindingType), reflect.TypeOf([]Info{}))
    fieldNameMap["acq_specs"] = "AcqSpecs"
    fields["next"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["next"] = "Next"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.acq_specs.list_result",fields, reflect.TypeOf(ListResult{}), fieldNameMap, validators)
}

func FilterSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["page"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["page"] = "Page"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.acq_specs.filter_spec",fields, reflect.TypeOf(FilterSpec{}), fieldNameMap, validators)
}

func UpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["counters"] = bindings.NewOptionalType(bindings.NewReferenceType(CounterSpecBindingType))
    fieldNameMap["counters"] = "Counters"
    fields["resources"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(vstats.RsrcIdBindingType), reflect.TypeOf([]vstats.RsrcId{})))
    fieldNameMap["resources"] = "Resources"
    fields["interval"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["interval"] = "Interval"
    fields["expiration"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["expiration"] = "Expiration"
    fields["memo_"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["memo_"] = "Memo_"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.acq_specs.update_spec",fields, reflect.TypeOf(UpdateSpec{}), fieldNameMap, validators)
}


