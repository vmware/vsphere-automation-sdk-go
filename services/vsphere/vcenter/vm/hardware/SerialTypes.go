/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Serial.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package hardware

import (
	"reflect"
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
type SerialBackingType string

const (
    // Virtual serial port is backed by a file.
	SerialBackingType_FILE SerialBackingType = "FILE"
    // Virtual serial port is backed by a device on the host where the virtual machine is running.
	SerialBackingType_HOST_DEVICE SerialBackingType = "HOST_DEVICE"
    // Virtual serial port is backed by a named pipe server. The virtual machine will accept a connection from a host application or another virtual machine on the same host. This is useful for capturing debugging information sent through the virtual serial port.
	SerialBackingType_PIPE_SERVER SerialBackingType = "PIPE_SERVER"
    // Virtual serial port is backed by a named pipe client. The virtual machine will connect to the named pipe provided by a host application or another virtual machine on the same host. This is useful for capturing debugging information sent through the virtual serial port.
	SerialBackingType_PIPE_CLIENT SerialBackingType = "PIPE_CLIENT"
    // Virtual serial port is backed by a network server. This backing may be used to create a network-accessible serial port on the virtual machine, accepting a connection from a remote system.
	SerialBackingType_NETWORK_SERVER SerialBackingType = "NETWORK_SERVER"
    // Virtual serial port is backed by a network client. This backing may be used to create a network-accessible serial port on the virtual machine, initiating a connection to a remote system.
	SerialBackingType_NETWORK_CLIENT SerialBackingType = "NETWORK_CLIENT"
)

func (b SerialBackingType) SerialBackingType() bool {
	switch b {
	case SerialBackingType_FILE:
		return true
	case SerialBackingType_HOST_DEVICE:
		return true
	case SerialBackingType_PIPE_SERVER:
		return true
	case SerialBackingType_PIPE_CLIENT:
		return true
	case SerialBackingType_NETWORK_SERVER:
		return true
	case SerialBackingType_NETWORK_CLIENT:
		return true
	default:
		return false
	}
}


// The ``BackingInfo`` class contains information about the physical resource backing a virtual serial port.
type SerialBackingInfo struct {
    // Backing type for the virtual serial port.
	Type_ SerialBackingType
    // Path of the file backing the virtual serial port.
	File *string
    // Name of the device backing the virtual serial port. 
	HostDevice *string
    // Flag indicating whether the virtual serial port is configured to automatically detect a suitable host device.
	AutoDetect *bool
    // Name of the pipe backing the virtual serial port.
	Pipe *string
    // Flag that enables optimized data transfer over the pipe. When the value is true, the host buffers data to prevent data overrun. This allows the virtual machine to read all of the data transferred over the pipe with no data loss.
	NoRxLoss *bool
    // URI specifying the location of the network service backing the virtual serial port. 
    //
    // * If SerialBackingInfo#type is SerialBackingType#SerialBackingType_NETWORK_SERVER, this property is the location used by clients to connect to this server. The hostname part of the URI should either be empty or should specify the address of the host on which the virtual machine is running.
    // * If SerialBackingInfo#type is SerialBackingType#SerialBackingType_NETWORK_CLIENT, this property is the location used by the virtual machine to connect to the remote server.
	NetworkLocation *url.URL
    // Proxy service that provides network access to the network backing. If set, the virtual machine initiates a connection with the proxy service and forwards the traffic to the proxy.
	Proxy *url.URL
}

// The ``BackingSpec`` class provides a specification of the physical resource backing a virtual serial port.
type SerialBackingSpec struct {
    // Backing type for the virtual serial port.
	Type_ SerialBackingType
    // Path of the file backing the virtual serial port.
	File *string
    // Name of the device backing the virtual serial port. 
	HostDevice *string
    // Name of the pipe backing the virtual serial port.
	Pipe *string
    // Flag that enables optimized data transfer over the pipe. When the value is true, the host buffers data to prevent data overrun. This allows the virtual machine to read all of the data transferred over the pipe with no data loss.
	NoRxLoss *bool
    // URI specifying the location of the network service backing the virtual serial port. 
    //
    // * If SerialBackingSpec#type is SerialBackingType#SerialBackingType_NETWORK_SERVER, this property is the location used by clients to connect to this server. The hostname part of the URI should either be empty or should specify the address of the host on which the virtual machine is running.
    // * If SerialBackingSpec#type is SerialBackingType#SerialBackingType_NETWORK_CLIENT, this property is the location used by the virtual machine to connect to the remote server.
	NetworkLocation *url.URL
    // Proxy service that provides network access to the network backing. If set, the virtual machine initiates a connection with the proxy service and forwards the traffic to the proxy.
	Proxy *url.URL
}

// The ``Info`` class contains information about a virtual serial port.
type SerialInfo struct {
    // Device label.
	Label string
    // CPU yield behavior. If set to true, the virtual machine will periodically relinquish the processor if its sole task is polling the virtual serial port. The amount of time it takes to regain the processor will depend on the degree of other virtual machine activity on the host.
	YieldOnPoll bool
    // Physical resource backing for the virtual serial port.
	Backing SerialBackingInfo
    // Connection status of the virtual device.
	State ConnectionState
    // Flag indicating whether the virtual device should be connected whenever the virtual machine is powered on.
	StartConnected bool
    // Flag indicating whether the guest can connect and disconnect the device.
	AllowGuestControl bool
}

// The ``CreateSpec`` class provides a specification for the configuration of a newly-created virtual serial port.
type SerialCreateSpec struct {
    // CPU yield behavior. If set to true, the virtual machine will periodically relinquish the processor if its sole task is polling the virtual serial port. The amount of time it takes to regain the processor will depend on the degree of other virtual machine activity on the host.
	YieldOnPoll *bool
    // Physical resource backing for the virtual serial port.
	Backing *SerialBackingSpec
    // Flag indicating whether the virtual device should be connected whenever the virtual machine is powered on.
	StartConnected *bool
    // Flag indicating whether the guest can connect and disconnect the device.
	AllowGuestControl *bool
}

// The ``UpdateSpec`` class describes the updates to be made to the configuration of a virtual serial port.
type SerialUpdateSpec struct {
    // CPU yield behavior. If set to true, the virtual machine will periodically relinquish the processor if its sole task is polling the virtual serial port. The amount of time it takes to regain the processor will depend on the degree of other virtual machine activity on the host. 
    //
    //  This property may be modified at any time, and changes applied to a connected virtual serial port take effect immediately.
	YieldOnPoll *bool
    // Physical resource backing for the virtual serial port. 
    //
    //  This property may only be modified if the virtual machine is not powered on or the virtual serial port is not connected.
	Backing *SerialBackingSpec
    // Flag indicating whether the virtual device should be connected whenever the virtual machine is powered on.
	StartConnected *bool
    // Flag indicating whether the guest can connect and disconnect the device.
	AllowGuestControl *bool
}

// The ``Summary`` class contains commonly used information about a virtual serial port.
type SerialSummary struct {
    // Identifier of the virtual serial port.
	Port string
}



func serialListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fieldNameMap["vm"] = "Vm"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serialListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(SerialSummaryBindingType), reflect.TypeOf([]SerialSummary{}))
}

func serialListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fieldNameMap["vm"] = "Vm"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"Error": 500,"NotFound": 404,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func serialGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["port"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.SerialPort"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["port"] = "Port"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serialGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(SerialInfoBindingType)
}

func serialGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["port"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.SerialPort"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["port"] = "Port"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"Error": 500,"NotFound": 404,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func serialCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["spec"] = bindings.NewReferenceType(SerialCreateSpecBindingType)
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serialCreateOutputType() bindings.BindingType {
	return bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.SerialPort"}, "")
}

func serialCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["spec"] = bindings.NewReferenceType(SerialCreateSpecBindingType)
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["spec"] = "Spec"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"Error": 500,"NotFound": 404,"NotAllowedInCurrentState": 400,"UnableToAllocateResource": 500,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func serialUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["port"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.SerialPort"}, "")
	fields["spec"] = bindings.NewReferenceType(SerialUpdateSpecBindingType)
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["port"] = "Port"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serialUpdateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func serialUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["port"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.SerialPort"}, "")
	fields["spec"] = bindings.NewReferenceType(SerialUpdateSpecBindingType)
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["port"] = "Port"
	fieldNameMap["spec"] = "Spec"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"Error": 500,"NotFound": 404,"NotAllowedInCurrentState": 400,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func serialDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["port"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.SerialPort"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["port"] = "Port"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serialDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func serialDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["port"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.SerialPort"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["port"] = "Port"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"Error": 500,"NotFound": 404,"NotAllowedInCurrentState": 400,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func serialConnectInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["port"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.SerialPort"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["port"] = "Port"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serialConnectOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func serialConnectRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["port"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.SerialPort"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["port"] = "Port"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"Error": 500,"NotFound": 404,"AlreadyInDesiredState": 400,"NotAllowedInCurrentState": 400,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func serialDisconnectInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["port"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.SerialPort"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["port"] = "Port"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serialDisconnectOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func serialDisconnectRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["port"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.SerialPort"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["port"] = "Port"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"Error": 500,"NotFound": 404,"AlreadyInDesiredState": 400,"NotAllowedInCurrentState": 400,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func SerialBackingInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.serial.backing_type", reflect.TypeOf(SerialBackingType(SerialBackingType_FILE)))
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
		map[string][]bindings.FieldData{
			"FILE": []bindings.FieldData{
				bindings.NewFieldData("file", true),
			},
			"HOST_DEVICE": []bindings.FieldData{
				bindings.NewFieldData("host_device", false),
				bindings.NewFieldData("auto_detect", true),
			},
			"PIPE_SERVER": []bindings.FieldData{
				bindings.NewFieldData("pipe", true),
				bindings.NewFieldData("no_rx_loss", true),
			},
			"PIPE_CLIENT": []bindings.FieldData{
				bindings.NewFieldData("pipe", true),
				bindings.NewFieldData("no_rx_loss", true),
			},
			"NETWORK_SERVER": []bindings.FieldData{
				bindings.NewFieldData("network_location", true),
				bindings.NewFieldData("proxy", false),
			},
			"NETWORK_CLIENT": []bindings.FieldData{
				bindings.NewFieldData("network_location", true),
				bindings.NewFieldData("proxy", false),
			},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.serial.backing_info", fields, reflect.TypeOf(SerialBackingInfo{}), fieldNameMap, validators)
}

func SerialBackingSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.serial.backing_type", reflect.TypeOf(SerialBackingType(SerialBackingType_FILE)))
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
		map[string][]bindings.FieldData{
			"FILE": []bindings.FieldData{
				bindings.NewFieldData("file", true),
			},
			"HOST_DEVICE": []bindings.FieldData{
				bindings.NewFieldData("host_device", false),
			},
			"PIPE_SERVER": []bindings.FieldData{
				bindings.NewFieldData("pipe", true),
				bindings.NewFieldData("no_rx_loss", false),
			},
			"PIPE_CLIENT": []bindings.FieldData{
				bindings.NewFieldData("pipe", true),
				bindings.NewFieldData("no_rx_loss", false),
			},
			"NETWORK_SERVER": []bindings.FieldData{
				bindings.NewFieldData("network_location", true),
				bindings.NewFieldData("proxy", false),
			},
			"NETWORK_CLIENT": []bindings.FieldData{
				bindings.NewFieldData("network_location", true),
				bindings.NewFieldData("proxy", false),
			},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.serial.backing_spec", fields, reflect.TypeOf(SerialBackingSpec{}), fieldNameMap, validators)
}

func SerialInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["label"] = bindings.NewStringType()
	fieldNameMap["label"] = "Label"
	fields["yield_on_poll"] = bindings.NewBooleanType()
	fieldNameMap["yield_on_poll"] = "YieldOnPoll"
	fields["backing"] = bindings.NewReferenceType(SerialBackingInfoBindingType)
	fieldNameMap["backing"] = "Backing"
	fields["state"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.connection_state", reflect.TypeOf(ConnectionState(ConnectionState_CONNECTED)))
	fieldNameMap["state"] = "State"
	fields["start_connected"] = bindings.NewBooleanType()
	fieldNameMap["start_connected"] = "StartConnected"
	fields["allow_guest_control"] = bindings.NewBooleanType()
	fieldNameMap["allow_guest_control"] = "AllowGuestControl"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.serial.info", fields, reflect.TypeOf(SerialInfo{}), fieldNameMap, validators)
}

func SerialCreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["yield_on_poll"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["yield_on_poll"] = "YieldOnPoll"
	fields["backing"] = bindings.NewOptionalType(bindings.NewReferenceType(SerialBackingSpecBindingType))
	fieldNameMap["backing"] = "Backing"
	fields["start_connected"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["start_connected"] = "StartConnected"
	fields["allow_guest_control"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["allow_guest_control"] = "AllowGuestControl"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.serial.create_spec", fields, reflect.TypeOf(SerialCreateSpec{}), fieldNameMap, validators)
}

func SerialUpdateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["yield_on_poll"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["yield_on_poll"] = "YieldOnPoll"
	fields["backing"] = bindings.NewOptionalType(bindings.NewReferenceType(SerialBackingSpecBindingType))
	fieldNameMap["backing"] = "Backing"
	fields["start_connected"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["start_connected"] = "StartConnected"
	fields["allow_guest_control"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["allow_guest_control"] = "AllowGuestControl"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.serial.update_spec", fields, reflect.TypeOf(SerialUpdateSpec{}), fieldNameMap, validators)
}

func SerialSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["port"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.SerialPort"}, "")
	fieldNameMap["port"] = "Port"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.serial.summary", fields, reflect.TypeOf(SerialSummary{}), fieldNameMap, validators)
}

