/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Servers.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package servers

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// ``DNSServerMode`` enumeration class Describes DNS Server source (DHCP,static)
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type DNSServerMode string

const (
    // DNS address is automatically assigned by a DHCP server.
     DNSServerMode_dhcp DNSServerMode = "dhcp"
    // DNS address is static.
     DNSServerMode_is_static DNSServerMode = "is_static"
)

func (d DNSServerMode) DNSServerMode() bool {
    switch d {
        case DNSServerMode_dhcp:
            return true
        case DNSServerMode_is_static:
            return true
        default:
            return false
    }
}




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





// ``DNSServerConfig`` class This structure represents the configuration state used to determine DNS servers.
 type DNSServerConfig struct {
    Mode DNSServerMode
    Servers []string
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
    fields["servers"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["servers"] = "Servers"
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
    fields["server"] = bindings.NewStringType()
    fieldNameMap["server"] = "Server"
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
    fields["config"] = bindings.NewReferenceType(DNSServerConfigBindingType)
    fieldNameMap["config"] = "Config"
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
    return bindings.NewReferenceType(DNSServerConfigBindingType)
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



func DNSServerConfigBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["mode"] = bindings.NewEnumType("com.vmware.appliance.networking.dns.servers.DNS_server_mode", reflect.TypeOf(DNSServerMode(DNSServerMode_dhcp)))
    fieldNameMap["mode"] = "Mode"
    fields["servers"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["servers"] = "Servers"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.networking.dns.servers.DNS_server_config",fields, reflect.TypeOf(DNSServerConfig{}), fieldNameMap, validators)
}

func MessageBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["message"] = bindings.NewStringType()
    fieldNameMap["message"] = "Message"
    fields["result"] = bindings.NewEnumType("com.vmware.appliance.networking.dns.servers.message_status", reflect.TypeOf(MessageStatus(MessageStatus_failure)))
    fieldNameMap["result"] = "Result"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.networking.dns.servers.message",fields, reflect.TypeOf(Message{}), fieldNameMap, validators)
}

func TestStatusInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["status"] = bindings.NewEnumType("com.vmware.appliance.networking.dns.servers.test_status", reflect.TypeOf(TestStatus(TestStatus_orange)))
    fieldNameMap["status"] = "Status"
    fields["messages"] = bindings.NewListType(bindings.NewReferenceType(MessageBindingType), reflect.TypeOf([]Message{}))
    fieldNameMap["messages"] = "Messages"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.networking.dns.servers.test_status_info",fields, reflect.TypeOf(TestStatusInfo{}), fieldNameMap, validators)
}


