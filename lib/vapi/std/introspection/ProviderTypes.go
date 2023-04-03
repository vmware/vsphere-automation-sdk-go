// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Provider.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package introspection

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiLog_ "github.com/vmware/vsphere-automation-sdk-go/runtime/log"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
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

func (s *ProviderInfo) GetType__() vapiBindings_.BindingType {
	return ProviderInfoBindingType()
}

func (s *ProviderInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ProviderInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

func providerGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ProviderGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(ProviderInfoBindingType)
}

func providerGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return vapiProtocol_.NewOperationRestMetadata(
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

func ProviderInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewIdType([]string{"com.vmware.vapi.provider"}, "")
	fieldNameMap["id"] = "Id"
	fields["checksum"] = vapiBindings_.NewStringType()
	fieldNameMap["checksum"] = "Checksum"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vapi.std.introspection.provider.info", fields, reflect.TypeOf(ProviderInfo{}), fieldNameMap, validators)
}
