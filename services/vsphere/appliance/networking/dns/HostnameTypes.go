/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Hostname.
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



// ``TestStatus`` enumeration class Health indicator
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type HostnameTestStatus string

const (
    // In case data has more than one test, this indicates not all tests were successful
	HostnameTestStatus_orange HostnameTestStatus = "orange"
    // All tests were successful for given data
	HostnameTestStatus_green HostnameTestStatus = "green"
    // All tests failed for given data
	HostnameTestStatus_red HostnameTestStatus = "red"
)

func (t HostnameTestStatus) HostnameTestStatus() bool {
	switch t {
	case HostnameTestStatus_orange:
		return true
	case HostnameTestStatus_green:
		return true
	case HostnameTestStatus_red:
		return true
	default:
		return false
	}
}


// ``MessageStatus`` enumeration class Individual test result
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type HostnameMessageStatus string

const (
    // message indicates the test failed.
	HostnameMessageStatus_failure HostnameMessageStatus = "failure"
    // message indicates that the test was successful.
	HostnameMessageStatus_success HostnameMessageStatus = "success"
)

func (m HostnameMessageStatus) HostnameMessageStatus() bool {
	switch m {
	case HostnameMessageStatus_failure:
		return true
	case HostnameMessageStatus_success:
		return true
	default:
		return false
	}
}


// ``Message`` class Test result and message
type HostnameMessage struct {
    // message
	Message string
    // result of the test
	Result HostnameMessageStatus
}

// ``TestStatusInfo`` class Overall test result
type HostnameTestStatusInfo struct {
    // Overall status of tests run.
	Status HostnameTestStatus
    // messages
	Messages []HostnameMessage
}



func hostnameTestInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func hostnameTestOutputType() bindings.BindingType {
	return bindings.NewReferenceType(HostnameTestStatusInfoBindingType)
}

func hostnameTestRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
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

func hostnameSetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func hostnameSetOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func hostnameSetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
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

func hostnameGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func hostnameGetOutputType() bindings.BindingType {
	return bindings.NewStringType()
}

func hostnameGetRestMetadata() protocol.OperationRestMetadata {
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


func HostnameMessageBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["message"] = bindings.NewStringType()
	fieldNameMap["message"] = "Message"
	fields["result"] = bindings.NewEnumType("com.vmware.appliance.networking.dns.hostname.message_status", reflect.TypeOf(HostnameMessageStatus(HostnameMessageStatus_failure)))
	fieldNameMap["result"] = "Result"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.networking.dns.hostname.message", fields, reflect.TypeOf(HostnameMessage{}), fieldNameMap, validators)
}

func HostnameTestStatusInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["status"] = bindings.NewEnumType("com.vmware.appliance.networking.dns.hostname.test_status", reflect.TypeOf(HostnameTestStatus(HostnameTestStatus_orange)))
	fieldNameMap["status"] = "Status"
	fields["messages"] = bindings.NewListType(bindings.NewReferenceType(HostnameMessageBindingType), reflect.TypeOf([]HostnameMessage{}))
	fieldNameMap["messages"] = "Messages"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.networking.dns.hostname.test_status_info", fields, reflect.TypeOf(HostnameTestStatusInfo{}), fieldNameMap, validators)
}

