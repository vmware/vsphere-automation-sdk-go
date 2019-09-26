/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Time.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package time

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)




// ``SystemTimeStruct`` class Structure representing the system time.
 type SystemTimeStruct struct {
    SecondsSinceEpoch float64
    Date string
    Time string
    Timezone string
}









func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewReferenceType(SystemTimeStructBindingType)
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



func SystemTimeStructBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["seconds_since_epoch"] = bindings.NewDoubleType()
    fieldNameMap["seconds_since_epoch"] = "SecondsSinceEpoch"
    fields["date"] = bindings.NewStringType()
    fieldNameMap["date"] = "Date"
    fields["time"] = bindings.NewStringType()
    fieldNameMap["time"] = "Time"
    fields["timezone"] = bindings.NewStringType()
    fieldNameMap["timezone"] = "Timezone"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.system.time.system_time_struct",fields, reflect.TypeOf(SystemTimeStruct{}), fieldNameMap, validators)
}


