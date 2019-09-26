/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Events.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package events

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)




// The ``Event`` contains selected fields from the corresponding Kubernetes event. Please refer Kubernetes Events API for more details. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Event struct {
    Name string
    Kind string
    Type_ string
    Message string
    Reason string
    Component string
    Count int64
    LastTimeStamp int64
}









func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["namespace"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, "")
    fieldNameMap["namespace"] = "Namespace"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(EventBindingType), reflect.TypeOf([]Event{}))
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
       map[string]int{"NotFound": 404,"Unsupported": 400,"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}



func EventBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["kind"] = bindings.NewStringType()
    fieldNameMap["kind"] = "Kind"
    fields["type"] = bindings.NewStringType()
    fieldNameMap["type"] = "Type_"
    fields["message"] = bindings.NewStringType()
    fieldNameMap["message"] = "Message"
    fields["reason"] = bindings.NewStringType()
    fieldNameMap["reason"] = "Reason"
    fields["component"] = bindings.NewStringType()
    fieldNameMap["component"] = "Component"
    fields["count"] = bindings.NewIntegerType()
    fieldNameMap["count"] = "Count"
    fields["last_time_stamp"] = bindings.NewIntegerType()
    fieldNameMap["last_time_stamp"] = "LastTimeStamp"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespaces.events.events.event",fields, reflect.TypeOf(Event{}), fieldNameMap, validators)
}


