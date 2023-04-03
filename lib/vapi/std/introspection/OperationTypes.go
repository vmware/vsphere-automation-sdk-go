// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Operation.
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

// The OperationDataDefinition structure describes a vAPI data type.
type OperationDataDefinition struct {
	// Data type of the value.
	Type_ OperationDataDefinitionDataTypeEnum
	// Contains the element definition for generic data types like List and Optional.
	ElementDefinition *OperationDataDefinition
	// Fully qualified name of the structure.
	Name *string
	// Fields of the structure type. The key of the map is the canonical name of the field and the value is the OperationDataDefinition for the field. The order of the structure fields defined in IDL is not maintained by the Operation service.
	Fields map[string]OperationDataDefinition
}

func (s *OperationDataDefinition) GetType__() vapiBindings_.BindingType {
	return OperationDataDefinitionBindingType()
}

func (s *OperationDataDefinition) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for OperationDataDefinition._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// The OperationDataDefinitionDataTypeEnum enumeration provides values representing the data types supported by the vAPI infrastructure.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type OperationDataDefinitionDataTypeEnum string

const (
	// Indicates the value is a binary type.
	OperationDataDefinitionDataType_BINARY OperationDataDefinitionDataTypeEnum = "BINARY"
	// Indicates the value is a boolean type. The possible values are True and False equivalent of the language used to invoke this operation.
	OperationDataDefinitionDataType_BOOLEAN OperationDataDefinitionDataTypeEnum = "BOOLEAN"
	// Indicates the value is a double type. It is a 64 bit floating point number.
	OperationDataDefinitionDataType_DOUBLE OperationDataDefinitionDataTypeEnum = "DOUBLE"
	// Indicates the value is a dynamic structure. This means, any data of type OperationDataDefinitionDataTypeEnum#DataType_STRUCTURE can be used.
	OperationDataDefinitionDataType_DYNAMIC_STRUCTURE OperationDataDefinitionDataTypeEnum = "DYNAMIC_STRUCTURE"
	// Indicates the value is a specific error type.
	OperationDataDefinitionDataType_ERROR OperationDataDefinitionDataTypeEnum = "ERROR"
	// Indicates the value is arbitrary error type. This means, any data of type OperationDataDefinitionDataTypeEnum#DataType_ERROR can be used.
	OperationDataDefinitionDataType_ANY_ERROR OperationDataDefinitionDataTypeEnum = "ANY_ERROR"
	// Indicates the value is a list data type. Any value of this type can have zero or more elements in the list.
	OperationDataDefinitionDataType_LIST OperationDataDefinitionDataTypeEnum = "LIST"
	// Indicates the value is a long data type. It is a 64 bit signed integer number.
	OperationDataDefinitionDataType_LONG OperationDataDefinitionDataTypeEnum = "LONG"
	// Indicates the value is an opaque type. This means, data of any OperationDataDefinitionDataTypeEnum can be used.
	OperationDataDefinitionDataType_OPAQUE OperationDataDefinitionDataTypeEnum = "OPAQUE"
	// Indicates the value is an optional data type. Any value of this type can be null.
	OperationDataDefinitionDataType_OPTIONAL OperationDataDefinitionDataTypeEnum = "OPTIONAL"
	// Indicates the value is a secret data type. This is used for sensitive information. The server will not log any data of this type and if possible wipe the data from the memory after usage.
	OperationDataDefinitionDataType_SECRET OperationDataDefinitionDataTypeEnum = "SECRET"
	// Indicates the value is a string data type. This is a unicode string.
	OperationDataDefinitionDataType_STRING OperationDataDefinitionDataTypeEnum = "STRING"
	// Indicates the value is a structure data type. A structure has string identifier and a set of fields with corresponding values.
	OperationDataDefinitionDataType_STRUCTURE OperationDataDefinitionDataTypeEnum = "STRUCTURE"
	// Indicates the value is a structure reference. This is used to break circular dependencies in the type references. This just has a string identifier of the structure. Clients have to maintain a list of structures already visited and use that to resolve this reference.
	OperationDataDefinitionDataType_STRUCTURE_REF OperationDataDefinitionDataTypeEnum = "STRUCTURE_REF"
	// Indicates the value is a void data type.
	OperationDataDefinitionDataType_VOID OperationDataDefinitionDataTypeEnum = "VOID"
)

func (d OperationDataDefinitionDataTypeEnum) OperationDataDefinitionDataTypeEnum() bool {
	switch d {
	case OperationDataDefinitionDataType_BINARY:
		return true
	case OperationDataDefinitionDataType_BOOLEAN:
		return true
	case OperationDataDefinitionDataType_DOUBLE:
		return true
	case OperationDataDefinitionDataType_DYNAMIC_STRUCTURE:
		return true
	case OperationDataDefinitionDataType_ERROR:
		return true
	case OperationDataDefinitionDataType_ANY_ERROR:
		return true
	case OperationDataDefinitionDataType_LIST:
		return true
	case OperationDataDefinitionDataType_LONG:
		return true
	case OperationDataDefinitionDataType_OPAQUE:
		return true
	case OperationDataDefinitionDataType_OPTIONAL:
		return true
	case OperationDataDefinitionDataType_SECRET:
		return true
	case OperationDataDefinitionDataType_STRING:
		return true
	case OperationDataDefinitionDataType_STRUCTURE:
		return true
	case OperationDataDefinitionDataType_STRUCTURE_REF:
		return true
	case OperationDataDefinitionDataType_VOID:
		return true
	default:
		return false
	}
}

// Information about a vAPI operation.
type OperationInfo struct {
	// OperationDataDefinition describing the operation input.
	//
	//  The OperationDataDefinition#type of this field will be OperationDataDefinitionDataTypeEnum#DataType_STRUCTURE. The keys of OperationDataDefinition#fields are the names of the operation parameters, and the values of OperationDataDefinition#fields describe the type of the operation parameters.
	InputDefinition OperationDataDefinition
	// OperationDataDefinition describing the operation output.
	OutputDefinition OperationDataDefinition
	// List of OperationDataDefinition describing the errors that the operation might report.
	//
	//  The OperationDataDefinition#type of every element in this list will be OperationDataDefinitionDataTypeEnum#DataType_ERROR.
	ErrorDefinitions []OperationDataDefinition
}

func (s *OperationInfo) GetType__() vapiBindings_.BindingType {
	return OperationInfoBindingType()
}

func (s *OperationInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for OperationInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

func operationListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = vapiBindings_.NewIdType([]string{"com.vmware.vapi.service"}, "")
	fieldNameMap["service_id"] = "ServiceId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func OperationListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewSetType(vapiBindings_.NewIdType([]string{"com.vmware.vapi.operation"}, ""), reflect.TypeOf(map[string]bool{}))
}

func operationListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = vapiBindings_.NewIdType([]string{"com.vmware.vapi.service"}, "")
	fieldNameMap["service_id"] = "ServiceId"
	paramsTypeMap["service_id"] = vapiBindings_.NewIdType([]string{"com.vmware.vapi.service"}, "")
	queryParams["service_id"] = "service_id"
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
		"/vapi/std/introspection/operation",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.not_found": 404})
}

func operationGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = vapiBindings_.NewIdType([]string{"com.vmware.vapi.service"}, "")
	fields["operation_id"] = vapiBindings_.NewIdType([]string{"com.vmware.vapi.operation"}, "")
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["operation_id"] = "OperationId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func OperationGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(OperationInfoBindingType)
}

func operationGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = vapiBindings_.NewIdType([]string{"com.vmware.vapi.service"}, "")
	fields["operation_id"] = vapiBindings_.NewIdType([]string{"com.vmware.vapi.operation"}, "")
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["operation_id"] = "OperationId"
	paramsTypeMap["service_id"] = vapiBindings_.NewIdType([]string{"com.vmware.vapi.service"}, "")
	paramsTypeMap["operation_id"] = vapiBindings_.NewIdType([]string{"com.vmware.vapi.operation"}, "")
	queryParams["service_id"] = "service_id"
	queryParams["operation_id"] = "operation_id"
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
		"operation_id",
		"",
		"GET",
		"/vapi/std/introspection/operation",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.not_found": 404})
}

func OperationDataDefinitionBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = vapiBindings_.NewEnumType("com.vmware.vapi.std.introspection.operation.data_definition.data_type", reflect.TypeOf(OperationDataDefinitionDataTypeEnum(OperationDataDefinitionDataType_BINARY)))
	fieldNameMap["type"] = "Type_"
	fields["element_definition"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(OperationDataDefinitionBindingType))
	fieldNameMap["element_definition"] = "ElementDefinition"
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewReferenceType(OperationDataDefinitionBindingType), reflect.TypeOf(map[string]OperationDataDefinition{})))
	fieldNameMap["fields"] = "Fields"
	var validators = []vapiBindings_.Validator{}
	uv1 := vapiBindings_.NewUnionValidator("type",
		map[string][]vapiBindings_.FieldData{
			"OPTIONAL": []vapiBindings_.FieldData{
				vapiBindings_.NewFieldData("element_definition", true),
			},
			"LIST": []vapiBindings_.FieldData{
				vapiBindings_.NewFieldData("element_definition", true),
			},
			"STRUCTURE": []vapiBindings_.FieldData{
				vapiBindings_.NewFieldData("name", true),
				vapiBindings_.NewFieldData("fields", true),
			},
			"STRUCTURE_REF": []vapiBindings_.FieldData{
				vapiBindings_.NewFieldData("name", true),
			},
			"ERROR": []vapiBindings_.FieldData{
				vapiBindings_.NewFieldData("name", true),
				vapiBindings_.NewFieldData("fields", true),
			},
			"BINARY":            []vapiBindings_.FieldData{},
			"BOOLEAN":           []vapiBindings_.FieldData{},
			"DOUBLE":            []vapiBindings_.FieldData{},
			"DYNAMIC_STRUCTURE": []vapiBindings_.FieldData{},
			"ANY_ERROR":         []vapiBindings_.FieldData{},
			"LONG":              []vapiBindings_.FieldData{},
			"OPAQUE":            []vapiBindings_.FieldData{},
			"SECRET":            []vapiBindings_.FieldData{},
			"STRING":            []vapiBindings_.FieldData{},
			"VOID":              []vapiBindings_.FieldData{},
		},
	)
	validators = append(validators, uv1)
	return vapiBindings_.NewStructType("com.vmware.vapi.std.introspection.operation.data_definition", fields, reflect.TypeOf(OperationDataDefinition{}), fieldNameMap, validators)
}

func OperationInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["input_definition"] = vapiBindings_.NewReferenceType(OperationDataDefinitionBindingType)
	fieldNameMap["input_definition"] = "InputDefinition"
	fields["output_definition"] = vapiBindings_.NewReferenceType(OperationDataDefinitionBindingType)
	fieldNameMap["output_definition"] = "OutputDefinition"
	fields["error_definitions"] = vapiBindings_.NewListType(vapiBindings_.NewReferenceType(OperationDataDefinitionBindingType), reflect.TypeOf([]OperationDataDefinition{}))
	fieldNameMap["error_definitions"] = "ErrorDefinitions"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vapi.std.introspection.operation.info", fields, reflect.TypeOf(OperationInfo{}), fieldNameMap, validators)
}
