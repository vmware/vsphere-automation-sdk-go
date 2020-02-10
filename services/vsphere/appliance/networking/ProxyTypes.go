/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Proxy.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package networking

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// ``Protocol`` enumeration class defines the protocols for which proxying is supported. This enumeration was added in vSphere API 6.7.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type ProxyProtocol string

const (
    // Proxy configuration for http. This constant field was added in vSphere API 6.7.
	ProxyProtocol_HTTP ProxyProtocol = "HTTP"
    // Proxy configuration for https. This constant field was added in vSphere API 6.7.
	ProxyProtocol_HTTPS ProxyProtocol = "HTTPS"
    // Proxy configuration for ftp. This constant field was added in vSphere API 6.7.
	ProxyProtocol_FTP ProxyProtocol = "FTP"
)

func (p ProxyProtocol) ProxyProtocol() bool {
	switch p {
	case ProxyProtocol_HTTP:
		return true
	case ProxyProtocol_HTTPS:
		return true
	case ProxyProtocol_FTP:
		return true
	default:
		return false
	}
}


// ``ServerStatus`` enumeration class defines the status of the server associated with the test run. This enumeration was added in vSphere API 6.7.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type ProxyServerStatus string

const (
    // Server is reachable. This constant field was added in vSphere API 6.7.
	ProxyServerStatus_SERVER_REACHABLE ProxyServerStatus = "SERVER_REACHABLE"
    // Server is unreachable. This constant field was added in vSphere API 6.7.
	ProxyServerStatus_SERVER_UNREACHABLE ProxyServerStatus = "SERVER_UNREACHABLE"
)

func (s ProxyServerStatus) ProxyServerStatus() bool {
	switch s {
	case ProxyServerStatus_SERVER_REACHABLE:
		return true
	case ProxyServerStatus_SERVER_UNREACHABLE:
		return true
	default:
		return false
	}
}


// The ``Config`` class defines proxy configuration. This class was added in vSphere API 6.7.
type ProxyConfig struct {
    // Hostname or IP address of the proxy server. This property was added in vSphere API 6.7.
	Server string
    // Port to connect to the proxy server. In a 'get' call, indicates the port connected to the proxy server. In a 'set' call, specifies the port to connect to the proxy server. A value of -1 indicates the default port. This property was added in vSphere API 6.7.
	Port int64
    // Username for proxy server. This property was added in vSphere API 6.7.
	Username *string
    // Password for proxy server. This property was added in vSphere API 6.7.
	Password *string
    // In the result of the ``#get`` and ``#list`` methods this property indicates whether proxying is enabled for a particular protocol. In the input to the ``test`` and ``set`` methods this property specifies whether proxying should be enabled for a particular protocol. This property was added in vSphere API 6.7.
	Enabled bool
}

// The ``TestResult`` class contains information about the test operation done on a proxy server. This class was added in vSphere API 6.7.
type ProxyTestResult struct {
    // Status of the proxy server indicating whether the proxy server is reachable. This property was added in vSphere API 6.7.
	Status ProxyServerStatus
    // Message associated with status. This property was added in vSphere API 6.7.
	Message std.LocalizableMessage
}



func proxyTestInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["host"] = bindings.NewStringType()
	fields["protocol"] = bindings.NewStringType()
	fields["config"] = bindings.NewReferenceType(ProxyConfigBindingType)
	fieldNameMap["host"] = "Host"
	fieldNameMap["protocol"] = "Protocol"
	fieldNameMap["config"] = "Config"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func proxyTestOutputType() bindings.BindingType {
	return bindings.NewReferenceType(ProxyTestResultBindingType)
}

func proxyTestRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["host"] = bindings.NewStringType()
	fields["protocol"] = bindings.NewStringType()
	fields["config"] = bindings.NewReferenceType(ProxyConfigBindingType)
	fieldNameMap["host"] = "Host"
	fieldNameMap["protocol"] = "Protocol"
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

func proxySetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["protocol"] = bindings.NewStringType()
	fields["config"] = bindings.NewReferenceType(ProxyConfigBindingType)
	fieldNameMap["protocol"] = "Protocol"
	fieldNameMap["config"] = "Config"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func proxySetOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func proxySetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["protocol"] = bindings.NewStringType()
	fields["config"] = bindings.NewReferenceType(ProxyConfigBindingType)
	fieldNameMap["protocol"] = "Protocol"
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

func proxyDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["protocol"] = bindings.NewStringType()
	fieldNameMap["protocol"] = "Protocol"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func proxyDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func proxyDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["protocol"] = bindings.NewStringType()
	fieldNameMap["protocol"] = "Protocol"
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

func proxyListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func proxyListOutputType() bindings.BindingType {
	return bindings.NewMapType(bindings.NewEnumType("com.vmware.appliance.networking.proxy.protocol", reflect.TypeOf(ProxyProtocol(ProxyProtocol_HTTP))), bindings.NewReferenceType(ProxyConfigBindingType),reflect.TypeOf(map[ProxyProtocol]ProxyConfig{}))
}

func proxyListRestMetadata() protocol.OperationRestMetadata {
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

func proxyGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["protocol"] = bindings.NewStringType()
	fieldNameMap["protocol"] = "Protocol"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func proxyGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(ProxyConfigBindingType)
}

func proxyGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["protocol"] = bindings.NewStringType()
	fieldNameMap["protocol"] = "Protocol"
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


func ProxyConfigBindingType() bindings.BindingType {
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
	return bindings.NewStructType("com.vmware.appliance.networking.proxy.config", fields, reflect.TypeOf(ProxyConfig{}), fieldNameMap, validators)
}

func ProxyTestResultBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["status"] = bindings.NewEnumType("com.vmware.appliance.networking.proxy.server_status", reflect.TypeOf(ProxyServerStatus(ProxyServerStatus_SERVER_REACHABLE)))
	fieldNameMap["status"] = "Status"
	fields["message"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
	fieldNameMap["message"] = "Message"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.networking.proxy.test_result", fields, reflect.TypeOf(ProxyTestResult{}), fieldNameMap, validators)
}

