// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Service.
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

// Information about a vAPI service
type ServiceInfo struct {
	// Set of identifiers of operations present in the service
	Operations map[string]bool
}

func (s *ServiceInfo) GetType__() bindings.BindingType {
	return ServiceInfoBindingType()
}

func (s *ServiceInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ServiceInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

func serviceListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serviceListOutputType() bindings.BindingType {
	return bindings.NewSetType(bindings.NewIdType([]string{"com.vmware.vapi.service"}, ""), reflect.TypeOf(map[string]bool{}))
}

func serviceListRestMetadata() protocol.OperationRestMetadata {
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
		"/vapi/std/introspection/service",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{})
}

func serviceGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = bindings.NewIdType([]string{"com.vmware.vapi.service"}, "")
	fieldNameMap["id"] = "Id"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serviceGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(ServiceInfoBindingType)
}

func serviceGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["id"] = bindings.NewIdType([]string{"com.vmware.vapi.service"}, "")
	fieldNameMap["id"] = "Id"
	paramsTypeMap["id"] = bindings.NewIdType([]string{"com.vmware.vapi.service"}, "")
	paramsTypeMap["id"] = bindings.NewIdType([]string{"com.vmware.vapi.service"}, "")
	pathParams["id"] = "id"
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
		"/vapi/std/introspection/service/{id}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.not_found": 404})
}

func ServiceInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["operations"] = bindings.NewSetType(bindings.NewIdType([]string{"com.vmware.vapi.operation"}, ""), reflect.TypeOf(map[string]bool{}))
	fieldNameMap["operations"] = "Operations"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vapi.std.introspection.service.info", fields, reflect.TypeOf(ServiceInfo{}), fieldNameMap, validators)
}
