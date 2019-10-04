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
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// ``Mode`` enumeration class Mode of the routes static or autoconfigured. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Mode string

const (
    // Route installed by administrator. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Mode_STATIC Mode = "STATIC"
    // Autoconfigured route. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Mode_AUTO Mode = "AUTO"
)

func (m Mode) Mode() bool {
    switch m {
        case Mode_STATIC:
            return true
        case Mode_AUTO:
            return true
        default:
            return false
    }
}




// ``OverallStatus`` enumeration class Health indicator. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type OverallStatus string

const (
    // In case data has more than one test, this indicates not all tests were successful. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     OverallStatus_PARTIAL_FAILED OverallStatus = "PARTIAL_FAILED"
    // All tests were successful for given data. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     OverallStatus_SUCCESSFUL OverallStatus = "SUCCESSFUL"
    // All tests failed for given data. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     OverallStatus_FAILED OverallStatus = "FAILED"
)

func (o OverallStatus) OverallStatus() bool {
    switch o {
        case OverallStatus_PARTIAL_FAILED:
            return true
        case OverallStatus_SUCCESSFUL:
            return true
        case OverallStatus_FAILED:
            return true
        default:
            return false
    }
}




// ``Status`` enumeration class Individual test result. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Status string

const (
    // message indicates the test failed. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Status_FAILURE Status = "FAILURE"
    // message indicates that the test was successful. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Status_SUCCESS Status = "SUCCESS"
)

func (s Status) Status() bool {
    switch s {
        case Status_FAILURE:
            return true
        case Status_SUCCESS:
            return true
        default:
            return false
    }
}





// ``Config`` class Structure that describes how routing is performed for a particular destination and prefix. A destination/prefix of 0.0.0.0/0 ( for IPv4) or ::/0 (for IPv6) refers to the default gateway. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Config struct {
    Destination *string
    Prefix int64
    Gateway *string
    InterfaceName *string
}






// ``Info`` class Structure that describes how routing is performed for a particular destination and prefix. A destination/prefix of 0.0.0.0/0 ( for IPv4) or ::/0 (for IPv6) refers to the default gateway. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Info struct {
    Mode Mode
    Destination *string
    Prefix int64
    Gateway *string
    InterfaceName *string
}






// ``RouteStatus`` class Test result and message. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type RouteStatus struct {
    Message std.LocalizableMessage
    Status Status
}






// ``TestStatus`` class Overall test result. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type TestStatus struct {
    OverallStatus OverallStatus
    RouteStatus []RouteStatus
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
    return bindings.NewReferenceType(TestStatusBindingType)
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
    fields["routes"] = bindings.NewListType(bindings.NewReferenceType(ConfigBindingType), reflect.TypeOf([]Config{}))
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


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(InfoBindingType), reflect.TypeOf([]Info{}))
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



func ConfigBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["destination"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["destination"] = "Destination"
    fields["prefix"] = bindings.NewIntegerType()
    fieldNameMap["prefix"] = "Prefix"
    fields["gateway"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["gateway"] = "Gateway"
    fields["interface_name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["interface_name"] = "InterfaceName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.networking.routes.config",fields, reflect.TypeOf(Config{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["mode"] = bindings.NewEnumType("com.vmware.appliance.networking.routes.mode", reflect.TypeOf(Mode(Mode_STATIC)))
    fieldNameMap["mode"] = "Mode"
    fields["destination"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["destination"] = "Destination"
    fields["prefix"] = bindings.NewIntegerType()
    fieldNameMap["prefix"] = "Prefix"
    fields["gateway"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["gateway"] = "Gateway"
    fields["interface_name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["interface_name"] = "InterfaceName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.networking.routes.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func RouteStatusBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["message"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
    fieldNameMap["message"] = "Message"
    fields["status"] = bindings.NewEnumType("com.vmware.appliance.networking.routes.status", reflect.TypeOf(Status(Status_FAILURE)))
    fieldNameMap["status"] = "Status"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.networking.routes.route_status",fields, reflect.TypeOf(RouteStatus{}), fieldNameMap, validators)
}

func TestStatusBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["overall_status"] = bindings.NewEnumType("com.vmware.appliance.networking.routes.overall_status", reflect.TypeOf(OverallStatus(OverallStatus_PARTIAL_FAILED)))
    fieldNameMap["overall_status"] = "OverallStatus"
    fields["route_status"] = bindings.NewListType(bindings.NewReferenceType(RouteStatusBindingType), reflect.TypeOf([]RouteStatus{}))
    fieldNameMap["route_status"] = "RouteStatus"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.networking.routes.test_status",fields, reflect.TypeOf(TestStatus{}), fieldNameMap, validators)
}


