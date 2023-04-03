// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Source.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package routing

import (
	vapiMetadata_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/metadata"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiLog_ "github.com/vmware/vsphere-automation-sdk-go/runtime/log"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"net/url"
	"reflect"
)

// Resource type for vAPI metadata source.
const Source_RESOURCE_TYPE = "com.vmware.vapi.metadata.source"

// Metadata source info.
type SourceInfo struct {
	// Description of the source.
	Description string
	// Type of the metadata source.
	Type_ vapiMetadata_.SourceTypeEnum
	// Absolute file path of the file that has the metadata information.
	Filepath *string
	// URI of the remote vAPI endpoint. This should be of the format http(s):IP:port/namespace.
	Address *url.URL
}

func (s *SourceInfo) GetType__() vapiBindings_.BindingType {
	return SourceInfoBindingType()
}

func (s *SourceInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for SourceInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Metadata source create spec.
type SourceCreateSpec struct {
	// English language human readable description of the source.
	Description string
	// Type of the metadata source.
	Type_ vapiMetadata_.SourceTypeEnum
	// Absolute file path of the metamodel metadata file that has the metamodel information about one component element.
	Filepath *string
	// Connection information of the remote server. This should be of the format http(s)://IP:port/namespace.
	//
	//  The remote server should contain the interfaces in com.vmware.vapi.metadata.metamodel package. It could expose metamodel information of one or more components.
	Address *url.URL
}

func (s *SourceCreateSpec) GetType__() vapiBindings_.BindingType {
	return SourceCreateSpecBindingType()
}

func (s *SourceCreateSpec) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for SourceCreateSpec._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

func sourceCreateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["source_id"] = vapiBindings_.NewIdType([]string{"com.vmware.vapi.metadata.source"}, "")
	fields["spec"] = vapiBindings_.NewReferenceType(SourceCreateSpecBindingType)
	fieldNameMap["source_id"] = "SourceId"
	fieldNameMap["spec"] = "Spec"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SourceCreateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func sourceCreateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["source_id"] = vapiBindings_.NewIdType([]string{"com.vmware.vapi.metadata.source"}, "")
	fields["spec"] = vapiBindings_.NewReferenceType(SourceCreateSpecBindingType)
	fieldNameMap["source_id"] = "SourceId"
	fieldNameMap["spec"] = "Spec"
	bodyFieldsMap["source_id"] = "source_id"
	bodyFieldsMap["spec"] = "spec"
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
		"POST",
		"/vapi/metadata/routing/source",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.already_exists": 400, "com.vmware.vapi.std.errors.invalid_argument": 400, "com.vmware.vapi.std.errors.not_found": 404})
}

func sourceDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["source_id"] = vapiBindings_.NewIdType([]string{"com.vmware.vapi.metadata.source"}, "")
	fieldNameMap["source_id"] = "SourceId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SourceDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func sourceDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["source_id"] = vapiBindings_.NewIdType([]string{"com.vmware.vapi.metadata.source"}, "")
	fieldNameMap["source_id"] = "SourceId"
	paramsTypeMap["source_id"] = vapiBindings_.NewIdType([]string{"com.vmware.vapi.metadata.source"}, "")
	paramsTypeMap["sourceId"] = vapiBindings_.NewIdType([]string{"com.vmware.vapi.metadata.source"}, "")
	pathParams["source_id"] = "sourceId"
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
		"DELETE",
		"/vapi/metadata/routing/source/{sourceId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.not_found": 404})
}

func sourceGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["source_id"] = vapiBindings_.NewIdType([]string{"com.vmware.vapi.metadata.source"}, "")
	fieldNameMap["source_id"] = "SourceId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SourceGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(SourceInfoBindingType)
}

func sourceGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["source_id"] = vapiBindings_.NewIdType([]string{"com.vmware.vapi.metadata.source"}, "")
	fieldNameMap["source_id"] = "SourceId"
	paramsTypeMap["source_id"] = vapiBindings_.NewIdType([]string{"com.vmware.vapi.metadata.source"}, "")
	paramsTypeMap["sourceId"] = vapiBindings_.NewIdType([]string{"com.vmware.vapi.metadata.source"}, "")
	pathParams["source_id"] = "sourceId"
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
		"/vapi/metadata/routing/source/{sourceId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.not_found": 404})
}

func sourceListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SourceListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewListType(vapiBindings_.NewIdType([]string{"com.vmware.vapi.metadata.source"}, ""), reflect.TypeOf([]string{}))
}

func sourceListRestMetadata() vapiProtocol_.OperationRestMetadata {
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
		"/vapi/metadata/routing/source",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{})
}

func sourceReloadInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["source_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIdType([]string{"com.vmware.vapi.metadata.source"}, ""))
	fieldNameMap["source_id"] = "SourceId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SourceReloadOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func sourceReloadRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["source_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIdType([]string{"com.vmware.vapi.metadata.source"}, ""))
	fieldNameMap["source_id"] = "SourceId"
	bodyFieldsMap["source_id"] = "source_id"
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
		"action=reload",
		"",
		"POST",
		"/vapi/metadata/routing/source",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.not_found": 404})
}

func sourceFingerprintInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["source_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIdType([]string{"com.vmware.vapi.metadata.source"}, ""))
	fieldNameMap["source_id"] = "SourceId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SourceFingerprintOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewStringType()
}

func sourceFingerprintRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["source_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIdType([]string{"com.vmware.vapi.metadata.source"}, ""))
	fieldNameMap["source_id"] = "SourceId"
	bodyFieldsMap["source_id"] = "source_id"
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
		"action=fingerprint",
		"",
		"POST",
		"/vapi/metadata/routing/source",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.not_found": 404})
}

func SourceInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["description"] = vapiBindings_.NewStringType()
	fieldNameMap["description"] = "Description"
	fields["type"] = vapiBindings_.NewEnumType("com.vmware.vapi.metadata.source_type", reflect.TypeOf(vapiMetadata_.SourceTypeEnum(vapiMetadata_.SourceType_FILE)))
	fieldNameMap["type"] = "Type_"
	fields["filepath"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["filepath"] = "Filepath"
	fields["address"] = vapiBindings_.NewOptionalType(vapiBindings_.NewUriType())
	fieldNameMap["address"] = "Address"
	var validators = []vapiBindings_.Validator{}
	uv1 := vapiBindings_.NewUnionValidator("type",
		map[string][]vapiBindings_.FieldData{
			"FILE": []vapiBindings_.FieldData{
				vapiBindings_.NewFieldData("filepath", true),
			},
			"REMOTE": []vapiBindings_.FieldData{
				vapiBindings_.NewFieldData("address", true),
			},
		},
	)
	validators = append(validators, uv1)
	return vapiBindings_.NewStructType("com.vmware.vapi.metadata.routing.source.info", fields, reflect.TypeOf(SourceInfo{}), fieldNameMap, validators)
}

func SourceCreateSpecBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["description"] = vapiBindings_.NewStringType()
	fieldNameMap["description"] = "Description"
	fields["type"] = vapiBindings_.NewEnumType("com.vmware.vapi.metadata.source_type", reflect.TypeOf(vapiMetadata_.SourceTypeEnum(vapiMetadata_.SourceType_FILE)))
	fieldNameMap["type"] = "Type_"
	fields["filepath"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["filepath"] = "Filepath"
	fields["address"] = vapiBindings_.NewOptionalType(vapiBindings_.NewUriType())
	fieldNameMap["address"] = "Address"
	var validators = []vapiBindings_.Validator{}
	uv1 := vapiBindings_.NewUnionValidator("type",
		map[string][]vapiBindings_.FieldData{
			"FILE": []vapiBindings_.FieldData{
				vapiBindings_.NewFieldData("filepath", true),
			},
			"REMOTE": []vapiBindings_.FieldData{
				vapiBindings_.NewFieldData("address", true),
			},
		},
	)
	validators = append(validators, uv1)
	return vapiBindings_.NewStructType("com.vmware.vapi.metadata.routing.source.create_spec", fields, reflect.TypeOf(SourceCreateSpec{}), fieldNameMap, validators)
}
