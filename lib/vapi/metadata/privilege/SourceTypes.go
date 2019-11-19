/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Source.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package privilege

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/metadata"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
	"net/url"
)

// Resource type for metadata source.
const Source_RESOURCE_TYPE = "com.vmware.vapi.metadata.privilege.source"


// The ``Info`` class contains the metadata source information.
type SourceInfo struct {
    // English language human readable description of the source.
	Description string
    // Type of the metadata source.
	Type_ metadata.SourceType
    // Absolute file path of the privilege metadata file that has the privilege information about one component element. The ``filePath`` is the path to the file in the server's filesystem.
	Filepath *string
    // Connection information for the remote server. This must be in the format http(s)://IP:port/namespace. 
    //
    //  The remote server must support the interfaces in the com.vmware.vapi.metadata.privilege package. It must expose privilege information of one or more components.
	Address *url.URL
}

// The ``CreateSpec`` class contains the registration information of a privilege source.
type SourceCreateSpec struct {
    // English language human readable description of the source.
	Description string
    // Type of the metadata source.
	Type_ metadata.SourceType
    // Absolute file path of the metamodel metadata file that has the metamodel information about one component element.
	Filepath *string
    // Connection information of the remote server. This should be of the format http(s)://IP:port/namespace. 
    //
    //  The remote server should contain the interfaces in com.vmware.vapi.metadata.metamodel package. It could expose metamodel information of one or more components.
	Address *url.URL
}



func sourceCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["source_id"] = bindings.NewIdType([]string{"com.vmware.vapi.metadata.privilege.source"}, "")
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
	fields["source_id"] = bindings.NewIdType([]string{"com.vmware.vapi.metadata.privilege.source"}, "")
	fields["spec"] = bindings.NewReferenceType(SourceCreateSpecBindingType)
	fieldNameMap["source_id"] = "SourceId"
	fieldNameMap["spec"] = "Spec"
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
		map[string]int{"AlreadyExists": 400,"InvalidArgument": 400,"NotFound": 404})
}

func sourceDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["source_id"] = bindings.NewIdType([]string{"com.vmware.vapi.metadata.privilege.source"}, "")
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
	fields["source_id"] = bindings.NewIdType([]string{"com.vmware.vapi.metadata.privilege.source"}, "")
	fieldNameMap["source_id"] = "SourceId"
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
		map[string]int{"NotFound": 404})
}

func sourceGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["source_id"] = bindings.NewIdType([]string{"com.vmware.vapi.metadata.privilege.source"}, "")
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
	fields["source_id"] = bindings.NewIdType([]string{"com.vmware.vapi.metadata.privilege.source"}, "")
	fieldNameMap["source_id"] = "SourceId"
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
		map[string]int{"NotFound": 404})
}

func sourceListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func sourceListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewIdType([]string{"com.vmware.vapi.metadata.privilege.source"}, ""), reflect.TypeOf([]string{}))
}

func sourceListRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{})
}

func sourceReloadInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["source_id"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.vapi.metadata.privilege.source"}, ""))
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
	fields["source_id"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.vapi.metadata.privilege.source"}, ""))
	fieldNameMap["source_id"] = "SourceId"
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
		map[string]int{"NotFound": 404})
}

func sourceFingerprintInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["source_id"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.vapi.metadata.privilege.source"}, ""))
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
	fields["source_id"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.vapi.metadata.privilege.source"}, ""))
	fieldNameMap["source_id"] = "SourceId"
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
		map[string]int{"NotFound": 404})
}


func SourceInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["description"] = bindings.NewStringType()
	fieldNameMap["description"] = "Description"
	fields["type"] = bindings.NewEnumType("com.vmware.vapi.metadata.source_type", reflect.TypeOf(metadata.SourceType(metadata.SourceType_FILE)))
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
	return bindings.NewStructType("com.vmware.vapi.metadata.privilege.source.info", fields, reflect.TypeOf(SourceInfo{}), fieldNameMap, validators)
}

func SourceCreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["description"] = bindings.NewStringType()
	fieldNameMap["description"] = "Description"
	fields["type"] = bindings.NewEnumType("com.vmware.vapi.metadata.source_type", reflect.TypeOf(metadata.SourceType(metadata.SourceType_FILE)))
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
	return bindings.NewStructType("com.vmware.vapi.metadata.privilege.source.create_spec", fields, reflect.TypeOf(SourceCreateSpec{}), fieldNameMap, validators)
}

