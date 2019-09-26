/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Ntp.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package ntp

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// ``ServerStatus`` enumeration class Status of server during test
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type ServerStatus string

const (
    // Server is reachable.
     ServerStatus_SERVER_REACHABLE ServerStatus = "SERVER_REACHABLE"
    // Server is unreachable.
     ServerStatus_SERVER_UNREACHABLE ServerStatus = "SERVER_UNREACHABLE"
)

func (s ServerStatus) ServerStatus() bool {
    switch s {
        case ServerStatus_SERVER_REACHABLE:
            return true
        case ServerStatus_SERVER_UNREACHABLE:
            return true
        default:
            return false
    }
}





// ``LocalizableMessage`` class Structure representing message
 type LocalizableMessage struct {
    Id string
    DefaultMessage string
    Args []string
}






// ``TestRunStatus`` class Status of the test
 type TestRunStatus struct {
    Server string
    Status ServerStatus
    Message LocalizableMessage
}









func testInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["servers"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["servers"] = "Servers"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func testOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(TestRunStatusBindingType), reflect.TypeOf([]TestRunStatus{}))
}

func testRestMetadata() protocol.OperationRestMetadata {
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


func setInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["servers"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["servers"] = "Servers"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func setOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func setRestMetadata() protocol.OperationRestMetadata {
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
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
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



func LocalizableMessageBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewStringType()
    fieldNameMap["id"] = "Id"
    fields["default_message"] = bindings.NewStringType()
    fieldNameMap["default_message"] = "DefaultMessage"
    fields["args"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["args"] = "Args"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.ntp.localizable_message",fields, reflect.TypeOf(LocalizableMessage{}), fieldNameMap, validators)
}

func TestRunStatusBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["server"] = bindings.NewStringType()
    fieldNameMap["server"] = "Server"
    fields["status"] = bindings.NewEnumType("com.vmware.appliance.ntp.server_status", reflect.TypeOf(ServerStatus(ServerStatus_SERVER_REACHABLE)))
    fieldNameMap["status"] = "Status"
    fields["message"] = bindings.NewReferenceType(LocalizableMessageBindingType)
    fieldNameMap["message"] = "Message"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.ntp.test_run_status",fields, reflect.TypeOf(TestRunStatus{}), fieldNameMap, validators)
}


