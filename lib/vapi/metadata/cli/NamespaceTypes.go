// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Namespace.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package cli

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/log"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
	"reflect"
)

// The ``Identity`` class uniquely identifies a namespace in the CLI namespace tree.
type NamespaceIdentity struct {
	// The dot-separated path of the namespace containing the namespace in the CLI node tree. For top-level namespace this will be empty.
	Path string
	// The name displayed to the user for this namespace.
	Name string
}

func (s *NamespaceIdentity) GetType__() bindings.BindingType {
	return NamespaceIdentityBindingType()
}

func (s *NamespaceIdentity) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for NamespaceIdentity._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// The ``Info`` class contains information about a namespace. It includes the identity of the namespace, a description, information children namespaces.
type NamespaceInfo struct {
	// Basic namespace identity.
	Identity NamespaceIdentity
	// The text description displayed to the user in help output.
	Description string
	// The children of this namespace in the tree of CLI namespaces.
	Children []NamespaceIdentity
}

func (s *NamespaceInfo) GetType__() bindings.BindingType {
	return NamespaceInfoBindingType()
}

func (s *NamespaceInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for NamespaceInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

func namespaceListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func namespaceListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(NamespaceIdentityBindingType), reflect.TypeOf([]NamespaceIdentity{}))
}

func namespaceListRestMetadata() protocol.OperationRestMetadata {
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
		"/vapi/metadata/cli/namespace",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{})
}

func namespaceGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["identity"] = bindings.NewReferenceType(NamespaceIdentityBindingType)
	fieldNameMap["identity"] = "Identity"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func namespaceGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(NamespaceInfoBindingType)
}

func namespaceGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["identity"] = bindings.NewReferenceType(NamespaceIdentityBindingType)
	fieldNameMap["identity"] = "Identity"
	bodyFieldsMap["identity"] = "identity"
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
		"action=get",
		"",
		"POST",
		"/vapi/metadata/cli/namespace",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.not_found": 404})
}

func namespaceFingerprintInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func namespaceFingerprintOutputType() bindings.BindingType {
	return bindings.NewStringType()
}

func namespaceFingerprintRestMetadata() protocol.OperationRestMetadata {
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
		"/vapi/metadata/cli/namespace/fingerprint",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{})
}

func NamespaceIdentityBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["path"] = bindings.NewStringType()
	fieldNameMap["path"] = "Path"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vapi.metadata.cli.namespace.identity", fields, reflect.TypeOf(NamespaceIdentity{}), fieldNameMap, validators)
}

func NamespaceInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["identity"] = bindings.NewReferenceType(NamespaceIdentityBindingType)
	fieldNameMap["identity"] = "Identity"
	fields["description"] = bindings.NewStringType()
	fieldNameMap["description"] = "Description"
	fields["children"] = bindings.NewListType(bindings.NewReferenceType(NamespaceIdentityBindingType), reflect.TypeOf([]NamespaceIdentity{}))
	fieldNameMap["children"] = "Children"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vapi.metadata.cli.namespace.info", fields, reflect.TypeOf(NamespaceInfo{}), fieldNameMap, validators)
}
