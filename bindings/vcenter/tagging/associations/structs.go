/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Associations.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package associations

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The last status for the iterator. A field of this type is returned as part of the result and indicates to the caller of the API whether it can continue to make requests for more data. 
//
//  The last status only reports on the state of the iteration at the time data was last returned. As a result, it not does guarantee if the next call will succeed in getting more data or not. 
//
//  Failures to retrieve results will be returned as Error responses. These last statuses are only returned when the iterator is operating as expected.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type LastIterationStatus string

const (
    // Iterator has more data pending and is ready to provide it. The caller can request the next page of data at any time. 
    //
    //  The number of results returned may be less than the usual size. In other words, the iterator may not fill the page. The iterator has returned at least 1 result.
     LastIterationStatus_READY LastIterationStatus = "READY"
    // Iterator has finished iterating through its inventory. There are currently no more entities to return and the caller can terminate iteration. If the iterator returned some data, the marker may be set to allow the iterator to continue from where it left off when additional data does become available. This value is used to indicate that all available data has been returned by the iterator.
     LastIterationStatus_END_OF_DATA LastIterationStatus = "END_OF_DATA"
)

func (l LastIterationStatus) LastIterationStatus() bool {
    switch l {
        case LastIterationStatus_READY:
            return true
        case LastIterationStatus_END_OF_DATA:
            return true
        default:
            return false
    }
}





// The ``IterationSpec`` class contains properties used to break results into pages when listing tags associated to objects see Associations#list).
 type IterationSpec struct {
    Marker *string
}






// The ``Summary`` describes a tag association.
 type Summary struct {
    Tag string
    Object std.DynamicID
}






// The ``ListResult`` class contains the list of tag associations in a page, as well as related metadata fields.
 type ListResult struct {
    Associations []Summary
    Marker *string
    Status LastIterationStatus
}









func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["iterate"] = bindings.NewOptionalType(bindings.NewReferenceType(IterationSpecBindingType))
    fieldNameMap["iterate"] = "Iterate"
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
    paramsTypeMap["iterate.marker"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.vcenter.tagging.associations.Marker"}, ""))
    queryParams["iterate.marker"] = "marker"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "GET",
      "/vcenter/tagging/associations",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"InvalidArgument": 400,"Unauthorized": 403})
}



func IterationSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["marker"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.vcenter.tagging.associations.Marker"}, ""))
    fieldNameMap["marker"] = "Marker"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.tagging.associations.iteration_spec",fields, reflect.TypeOf(IterationSpec{}), fieldNameMap, validators)
}

func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["tag"] = bindings.NewIdType([]string {"com.vmware.cis.tagging.Tag"}, "")
    fieldNameMap["tag"] = "Tag"
    fields["object"] = bindings.NewReferenceType(std.DynamicIDBindingType)
    fieldNameMap["object"] = "Object"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.tagging.associations.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}

func ListResultBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["associations"] = bindings.NewListType(bindings.NewReferenceType(SummaryBindingType), reflect.TypeOf([]Summary{}))
    fieldNameMap["associations"] = "Associations"
    fields["marker"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.vcenter.tagging.associations.Marker"}, ""))
    fieldNameMap["marker"] = "Marker"
    fields["status"] = bindings.NewEnumType("com.vmware.vcenter.tagging.associations.last_iteration_status", reflect.TypeOf(LastIterationStatus(LastIterationStatus_READY)))
    fieldNameMap["status"] = "Status"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.tagging.associations.list_result",fields, reflect.TypeOf(ListResult{}), fieldNameMap, validators)
}


