/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Forwarding.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package forwarding

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``Protocol`` enumeration class defines transport protocols for outbound log messages.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Protocol string

const (
    // Log messages will be forwarded to the remote host by using the TLS protocol.
     Protocol_TLS Protocol = "TLS"
    // Log messages will be forwarded to the remote host using the UDP protocol.
     Protocol_UDP Protocol = "UDP"
    // Log messages will be forwarded to the remote host using the TCP protocol.
     Protocol_TCP Protocol = "TCP"
)

func (p Protocol) Protocol() bool {
    switch p {
        case Protocol_TLS:
            return true
        case Protocol_UDP:
            return true
        case Protocol_TCP:
            return true
        default:
            return false
    }
}





// The ``Config`` class defines the configuration for log message forwarding to remote logging servers.
 type Config struct {
    Hostname string
    Port int64
    Protocol Protocol
}






 type ConnectionStatus struct {
    Hostname string
    State ConnectionStatus_State
    Message *std.LocalizableMessage
}




    
    // The ``State`` enumeration class defines the state values that a remote logging server can be in.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type ConnectionStatus_State string

    const (
        // The remote logging server is reachable.
         ConnectionStatus_State_UP ConnectionStatus_State = "UP"
        // The remote logging server is not reachable.
         ConnectionStatus_State_DOWN ConnectionStatus_State = "DOWN"
        // The status of remote logging server is unknown.
         ConnectionStatus_State_UNKNOWN ConnectionStatus_State = "UNKNOWN"
    )

    func (s ConnectionStatus_State) ConnectionStatus_State() bool {
        switch s {
            case ConnectionStatus_State_UP:
                return true
            case ConnectionStatus_State_DOWN:
                return true
            case ConnectionStatus_State_UNKNOWN:
                return true
            default:
                return false
        }
    }






func testInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["send_test_message"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["send_test_message"] = "SendTestMessage"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func testOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(ConnectionStatusBindingType), reflect.TypeOf([]ConnectionStatus{}))
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
       map[string]int{})
}


func setInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cfg_list"] = bindings.NewListType(bindings.NewReferenceType(ConfigBindingType), reflect.TypeOf([]Config{}))
    fieldNameMap["cfg_list"] = "CfgList"
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
       map[string]int{"InvalidArgument": 400,"UnableToAllocateResource": 500,"Error": 500})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(ConfigBindingType), reflect.TypeOf([]Config{}))
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
       map[string]int{})
}



func ConfigBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["hostname"] = bindings.NewStringType()
    fieldNameMap["hostname"] = "Hostname"
    fields["port"] = bindings.NewIntegerType()
    fieldNameMap["port"] = "Port"
    fields["protocol"] = bindings.NewEnumType("com.vmware.appliance.logging.forwarding.protocol", reflect.TypeOf(Protocol(Protocol_TLS)))
    fieldNameMap["protocol"] = "Protocol"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.logging.forwarding.config",fields, reflect.TypeOf(Config{}), fieldNameMap, validators)
}

func ConnectionStatusBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["hostname"] = bindings.NewStringType()
    fieldNameMap["hostname"] = "Hostname"
    fields["state"] = bindings.NewEnumType("com.vmware.appliance.logging.forwarding.connection_status.state", reflect.TypeOf(ConnectionStatus_State(ConnectionStatus_State_UP)))
    fieldNameMap["state"] = "State"
    fields["message"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
    fieldNameMap["message"] = "Message"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("state",
        map[string][]bindings.FieldData {
            "DOWN": []bindings.FieldData {
                 bindings.NewFieldData("message", false),
            },
            "UP": []bindings.FieldData {},
            "UNKNOWN": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.appliance.logging.forwarding.connection_status",fields, reflect.TypeOf(ConnectionStatus{}), fieldNameMap, validators)
}


