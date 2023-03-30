// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Namespace.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package cli

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiLog_ "github.com/vmware/vsphere-automation-sdk-go/runtime/log"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"reflect"
)

// The ``Identity`` class uniquely identifies a namespace in the CLI namespace tree.
type NamespaceIdentity struct {
	// The dot-separated path of the namespace containing the namespace in the CLI node tree. For top-level namespace this will be empty.
	Path string
	// The name displayed to the user for this namespace.
	Name string
}

func (s *NamespaceIdentity) GetType__() vapiBindings_.BindingType {
	return NamespaceIdentityBindingType()
}

func (s *NamespaceIdentity) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for NamespaceIdentity._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
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

func (s *NamespaceInfo) GetType__() vapiBindings_.BindingType {
	return NamespaceInfoBindingType()
}

func (s *NamespaceInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for NamespaceInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

func namespaceListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func NamespaceListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewListType(vapiBindings_.NewReferenceType(NamespaceIdentityBindingType), reflect.TypeOf([]NamespaceIdentity{}))
}

func namespaceListRestMetadata() vapiProtocol_.OperationRestMetadata {
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
		"/vapi/metadata/cli/namespace",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{})
}

func namespaceGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["identity"] = vapiBindings_.NewReferenceType(NamespaceIdentityBindingType)
	fieldNameMap["identity"] = "Identity"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func NamespaceGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(NamespaceInfoBindingType)
}

func namespaceGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["identity"] = vapiBindings_.NewReferenceType(NamespaceIdentityBindingType)
	fieldNameMap["identity"] = "Identity"
	bodyFieldsMap["identity"] = "identity"
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

func namespaceFingerprintInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func NamespaceFingerprintOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewStringType()
}

func namespaceFingerprintRestMetadata() vapiProtocol_.OperationRestMetadata {
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
		"/vapi/metadata/cli/namespace/fingerprint",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{})
}

func NamespaceIdentityBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["path"] = vapiBindings_.NewStringType()
	fieldNameMap["path"] = "Path"
	fields["name"] = vapiBindings_.NewStringType()
	fieldNameMap["name"] = "Name"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vapi.metadata.cli.namespace.identity", fields, reflect.TypeOf(NamespaceIdentity{}), fieldNameMap, validators)
}

func NamespaceInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["identity"] = vapiBindings_.NewReferenceType(NamespaceIdentityBindingType)
	fieldNameMap["identity"] = "Identity"
	fields["description"] = vapiBindings_.NewStringType()
	fieldNameMap["description"] = "Description"
	fields["children"] = vapiBindings_.NewListType(vapiBindings_.NewReferenceType(NamespaceIdentityBindingType), reflect.TypeOf([]NamespaceIdentity{}))
	fieldNameMap["children"] = "Children"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vapi.metadata.cli.namespace.info", fields, reflect.TypeOf(NamespaceInfo{}), fieldNameMap, validators)
}
