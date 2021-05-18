// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Source.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package cli

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/metadata"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/log"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"net/url"
	"reflect"
)

// The ``Info`` class contains the metadata source information.
type SourceInfo struct {
	// English language human readable description of the source.
	Description string
	// The type (FILE, REMOTE) of the metadata source.
	Type_ metadata.SourceTypeEnum
	// Absolute file path of the CLI metadata file that has the CLI information about one component. The ``filepath`` is the path to the file in the server's filesystem.
	Filepath *string
	// Connection information for the remote server. This should be in the format http(s)://IP:port/namespace.
	//
	//  The remote server must contain the interfaces in the com.vmware.vapi.metadata.cli package. It must expose CLI information of one or more components.
	Address *url.URL
}

func (s *SourceInfo) GetType__() bindings.BindingType {
	return SourceInfoBindingType()
}

func (s *SourceInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SourceInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// The ``CreateSpec`` class contains the registration information of a CLI source.
type SourceCreateSpec struct {
	// English language human readable description of the source.
	Description string
	// Type of the metadata source.
	Type_ metadata.SourceTypeEnum
	// Absolute file path of the metamodel metadata file that has the metamodel information about one component element.
	Filepath *string
	// Connection information of the remote server. This should be of the format http(s)://IP:port/namespace.
	//
	//  The remote server should contain the interfaces in com.vmware.vapi.metadata.metamodel package. It could expose metamodel information of one or more components.
	Address *url.URL
}

func (s *SourceCreateSpec) GetType__() bindings.BindingType {
	return SourceCreateSpecBindingType()
}

func (s *SourceCreateSpec) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for SourceCreateSpec._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

func sourceCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["source_id"] = bindings.NewIdType([]string{"com.vmware.vapi.metadata.source"}, "")
	fields["spec"] = bindings.NewReferenceType(SourceCreateSpecBindingType)
	fieldNameMap["source_id"] = "SourceId"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func sourceCreateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func sourceCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["source_id"] = bindings.NewIdType([]string{"com.vmware.vapi.metadata.source"}, "")
	fields["spec"] = bindings.NewReferenceType(SourceCreateSpecBindingType)
	fieldNameMap["source_id"] = "SourceId"
	fieldNameMap["spec"] = "Spec"
	bodyFieldsMap["source_id"] = "source_id"
	bodyFieldsMap["spec"] = "spec"
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
		"POST",
		"/vapi/metadata/cli/source",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.already_exists": 400, "com.vmware.vapi.std.errors.invalid_argument": 400, "com.vmware.vapi.std.errors.not_found": 404})
}

func sourceDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["source_id"] = bindings.NewIdType([]string{"com.vmware.vapi.metadata.source"}, "")
	fieldNameMap["source_id"] = "SourceId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func sourceDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func sourceDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["source_id"] = bindings.NewIdType([]string{"com.vmware.vapi.metadata.source"}, "")
	fieldNameMap["source_id"] = "SourceId"
	paramsTypeMap["source_id"] = bindings.NewIdType([]string{"com.vmware.vapi.metadata.source"}, "")
	paramsTypeMap["sourceId"] = bindings.NewIdType([]string{"com.vmware.vapi.metadata.source"}, "")
	pathParams["source_id"] = "sourceId"
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
		"DELETE",
		"/vapi/metadata/cli/source/{sourceId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.not_found": 404})
}

func sourceGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["source_id"] = bindings.NewIdType([]string{"com.vmware.vapi.metadata.source"}, "")
	fieldNameMap["source_id"] = "SourceId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func sourceGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(SourceInfoBindingType)
}

func sourceGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["source_id"] = bindings.NewIdType([]string{"com.vmware.vapi.metadata.source"}, "")
	fieldNameMap["source_id"] = "SourceId"
	paramsTypeMap["source_id"] = bindings.NewIdType([]string{"com.vmware.vapi.metadata.source"}, "")
	paramsTypeMap["sourceId"] = bindings.NewIdType([]string{"com.vmware.vapi.metadata.source"}, "")
	pathParams["source_id"] = "sourceId"
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
		"/vapi/metadata/cli/source/{sourceId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.not_found": 404})
}

func sourceListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func sourceListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewIdType([]string{"com.vmware.vapi.metadata.source"}, ""), reflect.TypeOf([]string{}))
}

func sourceListRestMetadata() protocol.OperationRestMetadata {
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
		"/vapi/metadata/cli/source",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{})
}

func sourceReloadInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["source_id"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.vapi.metadata.source"}, ""))
	fieldNameMap["source_id"] = "SourceId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func sourceReloadOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func sourceReloadRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["source_id"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.vapi.metadata.source"}, ""))
	fieldNameMap["source_id"] = "SourceId"
	bodyFieldsMap["source_id"] = "source_id"
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
		"action=reload",
		"",
		"POST",
		"/vapi/metadata/cli/source",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.not_found": 404})
}

func sourceFingerprintInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["source_id"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.vapi.metadata.source"}, ""))
	fieldNameMap["source_id"] = "SourceId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func sourceFingerprintOutputType() bindings.BindingType {
	return bindings.NewStringType()
}

func sourceFingerprintRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["source_id"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.vapi.metadata.source"}, ""))
	fieldNameMap["source_id"] = "SourceId"
	bodyFieldsMap["source_id"] = "source_id"
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
		"action=fingerprint",
		"",
		"POST",
		"/vapi/metadata/cli/source",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.not_found": 404})
}

func SourceInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["description"] = bindings.NewStringType()
	fieldNameMap["description"] = "Description"
	fields["type"] = bindings.NewEnumType("com.vmware.vapi.metadata.source_type", reflect.TypeOf(metadata.SourceTypeEnum(metadata.SourceType_FILE)))
	fieldNameMap["type"] = "Type_"
	fields["filepath"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["filepath"] = "Filepath"
	fields["address"] = bindings.NewOptionalType(bindings.NewUriType())
	fieldNameMap["address"] = "Address"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
		map[string][]bindings.FieldData{
			"FILE": []bindings.FieldData{
				bindings.NewFieldData("filepath", true),
			},
			"REMOTE": []bindings.FieldData{
				bindings.NewFieldData("address", true),
			},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vapi.metadata.cli.source.info", fields, reflect.TypeOf(SourceInfo{}), fieldNameMap, validators)
}

func SourceCreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["description"] = bindings.NewStringType()
	fieldNameMap["description"] = "Description"
	fields["type"] = bindings.NewEnumType("com.vmware.vapi.metadata.source_type", reflect.TypeOf(metadata.SourceTypeEnum(metadata.SourceType_FILE)))
	fieldNameMap["type"] = "Type_"
	fields["filepath"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["filepath"] = "Filepath"
	fields["address"] = bindings.NewOptionalType(bindings.NewUriType())
	fieldNameMap["address"] = "Address"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
		map[string][]bindings.FieldData{
			"FILE": []bindings.FieldData{
				bindings.NewFieldData("filepath", true),
			},
			"REMOTE": []bindings.FieldData{
				bindings.NewFieldData("address", true),
			},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vapi.metadata.cli.source.create_spec", fields, reflect.TypeOf(SourceCreateSpec{}), fieldNameMap, validators)
}
