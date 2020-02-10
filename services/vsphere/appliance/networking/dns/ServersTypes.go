/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Servers.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package dns

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// ``DNSServerMode`` enumeration class Describes DNS Server source (DHCP,static)
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type ServersDNSServerMode string

const (
    // DNS address is automatically assigned by a DHCP server.
	ServersDNSServerMode_dhcp ServersDNSServerMode = "dhcp"
    // DNS address is static.
	ServersDNSServerMode_is_static ServersDNSServerMode = "is_static"
)

func (d ServersDNSServerMode) ServersDNSServerMode() bool {
	switch d {
	case ServersDNSServerMode_dhcp:
		return true
	case ServersDNSServerMode_is_static:
		return true
	default:
		return false
	}
}


// ``TestStatus`` enumeration class Health indicator
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type ServersTestStatus string

const (
    // In case data has more than one test, this indicates not all tests were successful
	ServersTestStatus_orange ServersTestStatus = "orange"
    // All tests were successful for given data
	ServersTestStatus_green ServersTestStatus = "green"
    // All tests failed for given data
	ServersTestStatus_red ServersTestStatus = "red"
)

func (t ServersTestStatus) ServersTestStatus() bool {
	switch t {
	case ServersTestStatus_orange:
		return true
	case ServersTestStatus_green:
		return true
	case ServersTestStatus_red:
		return true
	default:
		return false
	}
}


// ``MessageStatus`` enumeration class Individual test result
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type ServersMessageStatus string

const (
    // message indicates the test failed.
	ServersMessageStatus_failure ServersMessageStatus = "failure"
    // message indicates that the test was successful.
	ServersMessageStatus_success ServersMessageStatus = "success"
)

func (m ServersMessageStatus) ServersMessageStatus() bool {
	switch m {
	case ServersMessageStatus_failure:
		return true
	case ServersMessageStatus_success:
		return true
	default:
		return false
	}
}


// ``DNSServerConfig`` class This structure represents the configuration state used to determine DNS servers.
type ServersDNSServerConfig struct {
    // Define how to determine the DNS servers. Leave the servers argument empty if the mode argument is "DHCP". Set the servers argument to a comma-separated list of DNS servers if the mode argument is "static". The DNS server are assigned from the specified list.
	Mode ServersDNSServerMode
    // List of the currently used DNS servers.
	Servers []string
}

// ``Message`` class Test result and message
type ServersMessage struct {
    // message
	Message string
    // result of the test
	Result ServersMessageStatus
}

// ``TestStatusInfo`` class Overall test result
type ServersTestStatusInfo struct {
    // Overall status of tests run.
	Status ServersTestStatus
    // messages
	Messages []ServersMessage
}



func serversTestInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["servers"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["servers"] = "Servers"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serversTestOutputType() bindings.BindingType {
	return bindings.NewReferenceType(ServersTestStatusInfoBindingType)
}

func serversTestRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["servers"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["servers"] = "Servers"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
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
		map[string]int{"Error": 500})
}

func serversAddInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["server"] = bindings.NewStringType()
	fieldNameMap["server"] = "Server"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serversAddOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func serversAddRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["server"] = bindings.NewStringType()
	fieldNameMap["server"] = "Server"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
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
		map[string]int{"Error": 500})
}

func serversSetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["config"] = bindings.NewReferenceType(ServersDNSServerConfigBindingType)
	fieldNameMap["config"] = "Config"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serversSetOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func serversSetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["config"] = bindings.NewReferenceType(ServersDNSServerConfigBindingType)
	fieldNameMap["config"] = "Config"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
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
		map[string]int{"Error": 500})
}

func serversGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serversGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(ServersDNSServerConfigBindingType)
}

func serversGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
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
		map[string]int{"Error": 500})
}


func ServersDNSServerConfigBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["mode"] = bindings.NewEnumType("com.vmware.appliance.networking.dns.servers.DNS_server_mode", reflect.TypeOf(ServersDNSServerMode(ServersDNSServerMode_dhcp)))
	fieldNameMap["mode"] = "Mode"
	fields["servers"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["servers"] = "Servers"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.networking.dns.servers.DNS_server_config", fields, reflect.TypeOf(ServersDNSServerConfig{}), fieldNameMap, validators)
}

func ServersMessageBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["message"] = bindings.NewStringType()
	fieldNameMap["message"] = "Message"
	fields["result"] = bindings.NewEnumType("com.vmware.appliance.networking.dns.servers.message_status", reflect.TypeOf(ServersMessageStatus(ServersMessageStatus_failure)))
	fieldNameMap["result"] = "Result"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.networking.dns.servers.message", fields, reflect.TypeOf(ServersMessage{}), fieldNameMap, validators)
}

func ServersTestStatusInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["status"] = bindings.NewEnumType("com.vmware.appliance.networking.dns.servers.test_status", reflect.TypeOf(ServersTestStatus(ServersTestStatus_orange)))
	fieldNameMap["status"] = "Status"
	fields["messages"] = bindings.NewListType(bindings.NewReferenceType(ServersMessageBindingType), reflect.TypeOf([]ServersMessage{}))
	fieldNameMap["messages"] = "Messages"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.networking.dns.servers.test_status_info", fields, reflect.TypeOf(ServersTestStatusInfo{}), fieldNameMap, validators)
}

