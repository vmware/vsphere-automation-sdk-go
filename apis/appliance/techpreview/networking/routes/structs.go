/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Routes.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package routes

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// ``TestStatus`` enumeration class Health indicator
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type TestStatus string

const (
    // In case data has more than one test, this indicates not all tests were successful
     TestStatus_orange TestStatus = "orange"
    // All tests were successful for given data
     TestStatus_green TestStatus = "green"
    // All tests failed for given data
     TestStatus_red TestStatus = "red"
)

func (t TestStatus) TestStatus() bool {
    switch t {
        case TestStatus_orange:
            return true
        case TestStatus_green:
            return true
        case TestStatus_red:
            return true
        default:
            return false
    }
}




// ``MessageStatus`` enumeration class Individual test result
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type MessageStatus string

const (
    // message indicates the test failed.
     MessageStatus_failure MessageStatus = "failure"
    // message indicates that the test was successful.
     MessageStatus_success MessageStatus = "success"
)

func (m MessageStatus) MessageStatus() bool {
    switch m {
        case MessageStatus_failure:
            return true
        case MessageStatus_success:
            return true
        default:
            return false
    }
}





// ``RouteReadOnly`` class Structure that describes how routing is performed for a particular destination and prefix. A destination/prefix of 0.0.0.0/0 ( for IPv4) or ::/0 (for IPv6) refers to the default gateway.
 type RouteReadOnly struct {
    Destination string
    Prefix int64
    Gateway string
    InterfaceName string
    IsStatic bool
}






// ``Route`` class Structure that describes how routing is performed for a particular destination and prefix. A destination/prefix of 0.0.0.0/0 ( for IPv4) or ::/0 (for IPv6) refers to the default gateway.
 type Route struct {
    Destination string
    Prefix int64
    Gateway string
    InterfaceName string
}






// ``Message`` class Test result and message
 type Message struct {
    Message string
    Result MessageStatus
}






// ``TestStatusInfo`` class Overall test result
 type TestStatusInfo struct {
    Status TestStatus
    Messages []Message
}









func testInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["gateways"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["gateways"] = "Gateways"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func testOutputType() bindings.BindingType {
    return bindings.NewReferenceType(TestStatusInfoBindingType)
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


func addInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["route"] = bindings.NewReferenceType(RouteBindingType)
    fieldNameMap["route"] = "Route"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func addOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func addRestMetadata() protocol.OperationRestMetadata {
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
    fields["routes"] = bindings.NewListType(bindings.NewReferenceType(RouteBindingType), reflect.TypeOf([]Route{}))
    fieldNameMap["routes"] = "Routes"
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


func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(RouteReadOnlyBindingType), reflect.TypeOf([]RouteReadOnly{}))
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


func deleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["route"] = bindings.NewReferenceType(RouteBindingType)
    fieldNameMap["route"] = "Route"
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



func RouteReadOnlyBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["destination"] = bindings.NewStringType()
    fieldNameMap["destination"] = "Destination"
    fields["prefix"] = bindings.NewIntegerType()
    fieldNameMap["prefix"] = "Prefix"
    fields["gateway"] = bindings.NewStringType()
    fieldNameMap["gateway"] = "Gateway"
    fields["interface_name"] = bindings.NewStringType()
    fieldNameMap["interface_name"] = "InterfaceName"
    fields["is_static"] = bindings.NewBooleanType()
    fieldNameMap["is_static"] = "IsStatic"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.networking.routes.route_read_only",fields, reflect.TypeOf(RouteReadOnly{}), fieldNameMap, validators)
}

func RouteBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["destination"] = bindings.NewStringType()
    fieldNameMap["destination"] = "Destination"
    fields["prefix"] = bindings.NewIntegerType()
    fieldNameMap["prefix"] = "Prefix"
    fields["gateway"] = bindings.NewStringType()
    fieldNameMap["gateway"] = "Gateway"
    fields["interface_name"] = bindings.NewStringType()
    fieldNameMap["interface_name"] = "InterfaceName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.networking.routes.route",fields, reflect.TypeOf(Route{}), fieldNameMap, validators)
}

func MessageBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["message"] = bindings.NewStringType()
    fieldNameMap["message"] = "Message"
    fields["result"] = bindings.NewEnumType("com.vmware.appliance.techpreview.networking.routes.message_status", reflect.TypeOf(MessageStatus(MessageStatus_failure)))
    fieldNameMap["result"] = "Result"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.networking.routes.message",fields, reflect.TypeOf(Message{}), fieldNameMap, validators)
}

func TestStatusInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["status"] = bindings.NewEnumType("com.vmware.appliance.techpreview.networking.routes.test_status", reflect.TypeOf(TestStatus(TestStatus_orange)))
    fieldNameMap["status"] = "Status"
    fields["messages"] = bindings.NewListType(bindings.NewReferenceType(MessageBindingType), reflect.TypeOf([]Message{}))
    fieldNameMap["messages"] = "Messages"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.networking.routes.test_status_info",fields, reflect.TypeOf(TestStatusInfo{}), fieldNameMap, validators)
}


