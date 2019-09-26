/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Proxy.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package proxy

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// ``Protocol`` enumeration class defines the protocols for which proxying is supported.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Protocol string

const (
    // Proxy configuration for http.
     Protocol_HTTP Protocol = "HTTP"
    // Proxy configuration for https.
     Protocol_HTTPS Protocol = "HTTPS"
    // Proxy configuration for ftp.
     Protocol_FTP Protocol = "FTP"
)

func (p Protocol) Protocol() bool {
    switch p {
        case Protocol_HTTP:
            return true
        case Protocol_HTTPS:
            return true
        case Protocol_FTP:
            return true
        default:
            return false
    }
}




// ``ServerStatus`` enumeration class defines the status of the server associated with the test run.
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





// The ``Config`` class defines proxy configuration.
 type Config struct {
    Server string
    Port int64
    Username *string
    Password *string
    Enabled bool
}






// The ``TestResult`` class contains information about the test operation done on a proxy server.
 type TestResult struct {
    Status ServerStatus
    Message std.LocalizableMessage
}









func testInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["host"] = bindings.NewStringType()
    fields["protocol"] = bindings.NewStringType()
    fields["config"] = bindings.NewReferenceType(ConfigBindingType)
    fieldNameMap["host"] = "Host"
    fieldNameMap["protocol"] = "Protocol"
    fieldNameMap["config"] = "Config"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func testOutputType() bindings.BindingType {
    return bindings.NewReferenceType(TestResultBindingType)
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
    fields["protocol"] = bindings.NewStringType()
    fields["config"] = bindings.NewReferenceType(ConfigBindingType)
    fieldNameMap["protocol"] = "Protocol"
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


func deleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["protocol"] = bindings.NewStringType()
    fieldNameMap["protocol"] = "Protocol"
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


func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listOutputType() bindings.BindingType {
    return bindings.NewMapType(bindings.NewEnumType("com.vmware.appliance.networking.proxy.protocol", reflect.TypeOf(Protocol(Protocol_HTTP))), bindings.NewReferenceType(ConfigBindingType),reflect.TypeOf(map[Protocol]Config{}))
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


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["protocol"] = bindings.NewStringType()
    fieldNameMap["protocol"] = "Protocol"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewReferenceType(ConfigBindingType)
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
    fields["server"] = bindings.NewStringType()
    fieldNameMap["server"] = "Server"
    fields["port"] = bindings.NewIntegerType()
    fieldNameMap["port"] = "Port"
    fields["username"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["username"] = "Username"
    fields["password"] = bindings.NewOptionalType(bindings.NewSecretType())
    fieldNameMap["password"] = "Password"
    fields["enabled"] = bindings.NewBooleanType()
    fieldNameMap["enabled"] = "Enabled"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.networking.proxy.config",fields, reflect.TypeOf(Config{}), fieldNameMap, validators)
}

func TestResultBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["status"] = bindings.NewEnumType("com.vmware.appliance.networking.proxy.server_status", reflect.TypeOf(ServerStatus(ServerStatus_SERVER_REACHABLE)))
    fieldNameMap["status"] = "Status"
    fields["message"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
    fieldNameMap["message"] = "Message"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.networking.proxy.test_result",fields, reflect.TypeOf(TestResult{}), fieldNameMap, validators)
}


