/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Tasks.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package tasks

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/cis/task"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// Resource type for task.
const Tasks_RESOURCE_TYPE = "com.vmware.cis.task"



// The ``GetSpec`` class describes what data should be included when retrieving information about a task.
 type GetSpec struct {
    ReturnAll *bool
    ExcludeResult *bool
}






// The ``FilterSpec`` class contains properties used to filter the results when listing tasks (see Tasks#list). If multiple properties are specified, only tasks matching all of the properties match the filter. 
//
//  Currently at least one of FilterSpec#tasks or FilterSpec#services must be specified and not empty.
 type FilterSpec struct {
    Tasks map[string]bool
    Services map[string]bool
    Operations map[string]bool
    Status map[task.Status]bool
    Targets []std.DynamicID
    Users map[string]bool
}









func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["task"] = bindings.NewIdType([]string {"com.vmware.cis.task"}, "")
    fields["spec"] = bindings.NewOptionalType(bindings.NewReferenceType(GetSpecBindingType))
    fieldNameMap["task"] = "Task"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewReferenceType(task.InfoBindingType)
}

func getRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
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
       map[string]int{"Error": 500,"NotFound": 404,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["filter_spec"] = bindings.NewOptionalType(bindings.NewReferenceType(FilterSpecBindingType))
    fields["result_spec"] = bindings.NewOptionalType(bindings.NewReferenceType(GetSpecBindingType))
    fieldNameMap["filter_spec"] = "FilterSpec"
    fieldNameMap["result_spec"] = "ResultSpec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listOutputType() bindings.BindingType {
    return bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.cis.task"}, ""), bindings.NewReferenceType(task.InfoBindingType),reflect.TypeOf(map[string]task.Info{}))
}

func listRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
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
       map[string]int{"InvalidArgument": 400,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func cancelInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["task"] = bindings.NewIdType([]string {"com.vmware.cis.task"}, "")
    fieldNameMap["task"] = "Task"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cancelOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func cancelRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
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
       map[string]int{"Error": 500,"NotAllowedInCurrentState": 400,"NotFound": 404,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403,"Unsupported": 400})
}



func GetSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["return_all"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["return_all"] = "ReturnAll"
    fields["exclude_result"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["exclude_result"] = "ExcludeResult"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.cis.tasks.get_spec",fields, reflect.TypeOf(GetSpec{}), fieldNameMap, validators)
}

func FilterSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["tasks"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"com.vmware.cis.task"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["tasks"] = "Tasks"
    fields["services"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"com.vmware.vapi.service"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["services"] = "Services"
    fields["operations"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"com.vmware.vapi.operation"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["operations"] = "Operations"
    fields["status"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewEnumType("com.vmware.cis.task.status", reflect.TypeOf(task.Status(task.Status_PENDING))), reflect.TypeOf(map[task.Status]bool{})))
    fieldNameMap["status"] = "Status"
    fields["targets"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(std.DynamicIDBindingType), reflect.TypeOf([]std.DynamicID{})))
    fieldNameMap["targets"] = "Targets"
    fields["users"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["users"] = "Users"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.cis.tasks.filter_spec",fields, reflect.TypeOf(FilterSpec{}), fieldNameMap, validators)
}


