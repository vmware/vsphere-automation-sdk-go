/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Serial.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package serial

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/vcenter/vm/hardware"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
    "net/url"
)

// Resource type for the virtual serial port device.
const Serial_RESOURCE_TYPE = "com.vmware.vcenter.vm.hardware.SerialPort"


// The ``BackingType`` enumeration class defines the valid backing types for a virtual serial port.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type BackingType string

const (
    // Virtual serial port is backed by a file.
     BackingType_FILE BackingType = "FILE"
    // Virtual serial port is backed by a device on the host where the virtual machine is running.
     BackingType_HOST_DEVICE BackingType = "HOST_DEVICE"
    // Virtual serial port is backed by a named pipe server. The virtual machine will accept a connection from a host application or another virtual machine on the same host. This is useful for capturing debugging information sent through the virtual serial port.
     BackingType_PIPE_SERVER BackingType = "PIPE_SERVER"
    // Virtual serial port is backed by a named pipe client. The virtual machine will connect to the named pipe provided by a host application or another virtual machine on the same host. This is useful for capturing debugging information sent through the virtual serial port.
     BackingType_PIPE_CLIENT BackingType = "PIPE_CLIENT"
    // Virtual serial port is backed by a network server. This backing may be used to create a network-accessible serial port on the virtual machine, accepting a connection from a remote system.
     BackingType_NETWORK_SERVER BackingType = "NETWORK_SERVER"
    // Virtual serial port is backed by a network client. This backing may be used to create a network-accessible serial port on the virtual machine, initiating a connection to a remote system.
     BackingType_NETWORK_CLIENT BackingType = "NETWORK_CLIENT"
)

func (b BackingType) BackingType() bool {
    switch b {
        case BackingType_FILE:
            return true
        case BackingType_HOST_DEVICE:
            return true
        case BackingType_PIPE_SERVER:
            return true
        case BackingType_PIPE_CLIENT:
            return true
        case BackingType_NETWORK_SERVER:
            return true
        case BackingType_NETWORK_CLIENT:
            return true
        default:
            return false
    }
}





// The ``BackingInfo`` class contains information about the physical resource backing a virtual serial port.
 type BackingInfo struct {
    Type_ BackingType
    File *string
    HostDevice *string
    AutoDetect *bool
    Pipe *string
    NoRxLoss *bool
    NetworkLocation *url.URL
    Proxy *url.URL
}






// The ``BackingSpec`` class provides a specification of the physical resource backing a virtual serial port.
 type BackingSpec struct {
    Type_ BackingType
    File *string
    HostDevice *string
    Pipe *string
    NoRxLoss *bool
    NetworkLocation *url.URL
    Proxy *url.URL
}






// The ``Info`` class contains information about a virtual serial port.
 type Info struct {
    Label string
    YieldOnPoll bool
    Backing BackingInfo
    State hardware.ConnectionState
    StartConnected bool
    AllowGuestControl bool
}






// The ``CreateSpec`` class provides a specification for the configuration of a newly-created virtual serial port.
 type CreateSpec struct {
    YieldOnPoll *bool
    Backing *BackingSpec
    StartConnected *bool
    AllowGuestControl *bool
}






// The ``UpdateSpec`` class describes the updates to be made to the configuration of a virtual serial port.
 type UpdateSpec struct {
    YieldOnPoll *bool
    Backing *BackingSpec
    StartConnected *bool
    AllowGuestControl *bool
}






// The ``Summary`` class contains commonly used information about a virtual serial port.
 type Summary struct {
    Port string
}









func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fieldNameMap["vm"] = "Vm"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(SummaryBindingType), reflect.TypeOf([]Summary{}))
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
       map[string]int{"Error": 500,"NotFound": 404,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fields["port"] = bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.SerialPort"}, "")
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["port"] = "Port"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewReferenceType(InfoBindingType)
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


func createInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fields["spec"] = bindings.NewReferenceType(CreateSpecBindingType)
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func createOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.SerialPort"}, "")
}

func createRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"NotFound": 404,"NotAllowedInCurrentState": 400,"UnableToAllocateResource": 500,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func updateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fields["port"] = bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.SerialPort"}, "")
    fields["spec"] = bindings.NewReferenceType(UpdateSpecBindingType)
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["port"] = "Port"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func updateOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func updateRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"NotFound": 404,"NotAllowedInCurrentState": 400,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func deleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fields["port"] = bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.SerialPort"}, "")
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["port"] = "Port"
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
       map[string]int{"Error": 500,"NotFound": 404,"NotAllowedInCurrentState": 400,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func connectInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fields["port"] = bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.SerialPort"}, "")
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["port"] = "Port"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func connectOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func connectRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"NotFound": 404,"AlreadyInDesiredState": 400,"NotAllowedInCurrentState": 400,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func disconnectInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fields["port"] = bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.SerialPort"}, "")
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["port"] = "Port"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func disconnectOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func disconnectRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"NotFound": 404,"AlreadyInDesiredState": 400,"NotAllowedInCurrentState": 400,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}



func BackingInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.serial.backing_type", reflect.TypeOf(BackingType(BackingType_FILE)))
    fieldNameMap["type"] = "Type_"
    fields["file"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["file"] = "File"
    fields["host_device"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["host_device"] = "HostDevice"
    fields["auto_detect"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["auto_detect"] = "AutoDetect"
    fields["pipe"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["pipe"] = "Pipe"
    fields["no_rx_loss"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["no_rx_loss"] = "NoRxLoss"
    fields["network_location"] = bindings.NewOptionalType(bindings.NewUriType())
    fieldNameMap["network_location"] = "NetworkLocation"
    fields["proxy"] = bindings.NewOptionalType(bindings.NewUriType())
    fieldNameMap["proxy"] = "Proxy"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("type",
        map[string][]bindings.FieldData {
            "FILE": []bindings.FieldData {
                 bindings.NewFieldData("file", true),
            },
            "HOST_DEVICE": []bindings.FieldData {
                 bindings.NewFieldData("host_device", false),
                 bindings.NewFieldData("auto_detect", true),
            },
            "PIPE_SERVER": []bindings.FieldData {
                 bindings.NewFieldData("pipe", true),
                 bindings.NewFieldData("no_rx_loss", true),
            },
            "PIPE_CLIENT": []bindings.FieldData {
                 bindings.NewFieldData("pipe", true),
                 bindings.NewFieldData("no_rx_loss", true),
            },
            "NETWORK_SERVER": []bindings.FieldData {
                 bindings.NewFieldData("network_location", true),
                 bindings.NewFieldData("proxy", false),
            },
            "NETWORK_CLIENT": []bindings.FieldData {
                 bindings.NewFieldData("network_location", true),
                 bindings.NewFieldData("proxy", false),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.serial.backing_info",fields, reflect.TypeOf(BackingInfo{}), fieldNameMap, validators)
}

func BackingSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.serial.backing_type", reflect.TypeOf(BackingType(BackingType_FILE)))
    fieldNameMap["type"] = "Type_"
    fields["file"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["file"] = "File"
    fields["host_device"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["host_device"] = "HostDevice"
    fields["pipe"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["pipe"] = "Pipe"
    fields["no_rx_loss"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["no_rx_loss"] = "NoRxLoss"
    fields["network_location"] = bindings.NewOptionalType(bindings.NewUriType())
    fieldNameMap["network_location"] = "NetworkLocation"
    fields["proxy"] = bindings.NewOptionalType(bindings.NewUriType())
    fieldNameMap["proxy"] = "Proxy"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("type",
        map[string][]bindings.FieldData {
            "FILE": []bindings.FieldData {
                 bindings.NewFieldData("file", true),
            },
            "HOST_DEVICE": []bindings.FieldData {
                 bindings.NewFieldData("host_device", false),
            },
            "PIPE_SERVER": []bindings.FieldData {
                 bindings.NewFieldData("pipe", true),
                 bindings.NewFieldData("no_rx_loss", false),
            },
            "PIPE_CLIENT": []bindings.FieldData {
                 bindings.NewFieldData("pipe", true),
                 bindings.NewFieldData("no_rx_loss", false),
            },
            "NETWORK_SERVER": []bindings.FieldData {
                 bindings.NewFieldData("network_location", true),
                 bindings.NewFieldData("proxy", false),
            },
            "NETWORK_CLIENT": []bindings.FieldData {
                 bindings.NewFieldData("network_location", true),
                 bindings.NewFieldData("proxy", false),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.serial.backing_spec",fields, reflect.TypeOf(BackingSpec{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["label"] = bindings.NewStringType()
    fieldNameMap["label"] = "Label"
    fields["yield_on_poll"] = bindings.NewBooleanType()
    fieldNameMap["yield_on_poll"] = "YieldOnPoll"
    fields["backing"] = bindings.NewReferenceType(BackingInfoBindingType)
    fieldNameMap["backing"] = "Backing"
    fields["state"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.connection_state", reflect.TypeOf(hardware.ConnectionState(hardware.ConnectionState_CONNECTED)))
    fieldNameMap["state"] = "State"
    fields["start_connected"] = bindings.NewBooleanType()
    fieldNameMap["start_connected"] = "StartConnected"
    fields["allow_guest_control"] = bindings.NewBooleanType()
    fieldNameMap["allow_guest_control"] = "AllowGuestControl"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.serial.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func CreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["yield_on_poll"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["yield_on_poll"] = "YieldOnPoll"
    fields["backing"] = bindings.NewOptionalType(bindings.NewReferenceType(BackingSpecBindingType))
    fieldNameMap["backing"] = "Backing"
    fields["start_connected"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["start_connected"] = "StartConnected"
    fields["allow_guest_control"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["allow_guest_control"] = "AllowGuestControl"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.serial.create_spec",fields, reflect.TypeOf(CreateSpec{}), fieldNameMap, validators)
}

func UpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["yield_on_poll"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["yield_on_poll"] = "YieldOnPoll"
    fields["backing"] = bindings.NewOptionalType(bindings.NewReferenceType(BackingSpecBindingType))
    fieldNameMap["backing"] = "Backing"
    fields["start_connected"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["start_connected"] = "StartConnected"
    fields["allow_guest_control"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["allow_guest_control"] = "AllowGuestControl"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.serial.update_spec",fields, reflect.TypeOf(UpdateSpec{}), fieldNameMap, validators)
}

func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["port"] = bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.SerialPort"}, "")
    fieldNameMap["port"] = "Port"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.serial.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}


