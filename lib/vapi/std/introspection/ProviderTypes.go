// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Provider.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package introspection

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/log"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
	"reflect"
)

// Information about a vAPI provider
type ProviderInfo struct {
	// Identifier of the provider
	Id string
	// Checksum of the information present in the provider.
	//
	//  Clients can use this information to check if the service information has changed. When a new service is added or removed (or) one of the existing service information is modified, the value of the checksum changes.
	//
	//  The information used to calculate the checksum includes:
	//
	// * service identifiers
	// * operation identifiers inside the service
	// * input, output and error definitions of an operation
	Checksum string
}

func (s *ProviderInfo) GetType__() bindings.BindingType {
	return ProviderInfoBindingType()
}

func (s *ProviderInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ProviderInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

func providerGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func providerGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(ProviderInfoBindingType)
}

func providerGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/vapi/std/introspection/provider",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{})
}

func ProviderInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = bindings.NewIdType([]string{"com.vmware.vapi.provider"}, "")
	fieldNameMap["id"] = "Id"
	fields["checksum"] = bindings.NewStringType()
	fieldNameMap["checksum"] = "Checksum"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vapi.std.introspection.provider.info", fields, reflect.TypeOf(ProviderInfo{}), fieldNameMap, validators)
}
