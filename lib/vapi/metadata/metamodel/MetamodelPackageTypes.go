// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for package: com.vmware.vapi.metadata.metamodel.
// Includes binding types of a top level structures and enumerations.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package metamodel

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/log"
	"reflect"
)

// The ``ComponentData`` class contains the metamodel metadata information of a component element along with its fingerprint.
type ComponentData struct {
	// Metamodel information of the component element. This includes information about all the package elements contained in this component element.
	//
	//  The metamodel information about a component could be quite large if there are a lot of package elements contained in this component.
	Info ComponentInfo
	// Fingerprint of the metamodel metadata of the component component.
	//
	//  Metamodel information could change when there is an infrastructure update and new functionality is added to an existing component.
	//
	//  Since the data present in ComponentData#info could be quite large, ``fingerprint`` provides a convenient way to check if the data for a particular component is updated.
	//
	//  You should store the fingerprint associated with a component. After an update, by invoking the Component#fingerprint method, you can retrieve the new fingerprint for the component. If the new fingerprint and the previously stored fingerprint do not match, clients can use the Component#get to retrieve the new metamodel information for the component.
	Fingerprint string
}

func (s *ComponentData) GetType__() bindings.BindingType {
	return ComponentDataBindingType()
}

func (s *ComponentData) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ComponentData._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// The ``ComponentInfo`` class contains metamodel metadata information about a component element.
type ComponentInfo struct {
	// Dot separated name of the component element. The segments in the name reflect the organization of the APIs. The format of each segment is lower case with underscores. Each underscore represents a word boundary. If there are acronyms in the word, the capitalization is preserved. This format makes it easy to translate the segment into a different naming convention.
	Name string
	// Metamodel metadata information of all the package elements contained in the component element. The key in the map is the identifier of the package element and the value in the map is the metamodel information of the package element.
	Packages map[string]PackageInfo
	// Generic metadata for the component element. The key in the map is the name of the metadata element and the value is the data associated with that metadata element.
	//
	//  The MetadataIdentifier contains possible string values for keys in the map.
	Metadata map[string]ElementMap
	// English language documentation for a component. It can contain HTML markup and documentation tags (similar to Javadoc tags). The first sentence of the package documentation is a complete sentence that identifies the component by name and summarizes the purpose of the component.
	Documentation string
}

func (s *ComponentInfo) GetType__() bindings.BindingType {
	return ComponentInfoBindingType()
}

func (s *ComponentInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ComponentInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// The ``ConstantInfo`` class contains metamodel information of the constant elements.
type ConstantInfo struct {
	// Type of the constant element.
	Type_ Type
	// Value of the constant element.
	Value ConstantValue
	// English language documentation for the constant element. It can contain HTML markup and documentation tags (similar to Javadoc tags).
	Documentation string
}

func (s *ConstantInfo) GetType__() bindings.BindingType {
	return ConstantInfoBindingType()
}

func (s *ConstantInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ConstantInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// The ``ConstantValue`` class contains the metamodel information of the constant element.
type ConstantValue struct {
	// Category of the type of constant value.
	Category ConstantValueCategoryEnum
	// Primitive value of the constant element.
	PrimitiveValue *PrimitiveValue
	// List value of the constant element.
	ListValue []PrimitiveValue
}

func (s *ConstantValue) GetType__() bindings.BindingType {
	return ConstantValueBindingType()
}

func (s *ConstantValue) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ConstantValue._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// The ``Category`` enumeration class defines enumeration constants for the valid kinds of values.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type ConstantValueCategoryEnum string

const (
	// Indicates the type of constant value is primitive.
	ConstantValueCategory_PRIMITIVE ConstantValueCategoryEnum = "PRIMITIVE"
	// Indicates the type of constant value is a list.
	ConstantValueCategory_LIST ConstantValueCategoryEnum = "LIST"
)

func (c ConstantValueCategoryEnum) ConstantValueCategoryEnum() bool {
	switch c {
	case ConstantValueCategory_PRIMITIVE:
		return true
	case ConstantValueCategory_LIST:
		return true
	default:
		return false
	}
}

// The ``ElementMap`` class contains the metadata elements.
//
//  One of the sources for metadata is the annotations present in the interface definition language. When an annotation is represented in the ``ElementMap``, ``ElementMap`` describes the data specified in the arguments for the annotation.
//
//  For example, in ``\\\\@UnionCase(tag="tag", value="SELECT")``, ElementMap describes the keyword arguments tag and value.
type ElementMap struct {
	// Metamodel information of the metadata elements. The key parameter of the map is the identifier for the element and the value corresponds to the element value.
	Elements map[string]ElementValue
}

func (s *ElementMap) GetType__() bindings.BindingType {
	return ElementMapBindingType()
}

func (s *ElementMap) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ElementMap._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// The ``ElementValue`` class describes the value of the metadata element.
type ElementValue struct {
	// Type of the value.
	Type_ ElementValueTypeEnum
	// Long value of the metadata element.
	LongValue *int64
	// String value of the metadata element.
	StringValue *string
	// List of strings value of the metadata element.
	ListValue []string
	// Identifier of the structure element.
	StructureId *string
	// List of identifiers of the structure elements.
	StructureIds []string
}

func (s *ElementValue) GetType__() bindings.BindingType {
	return ElementValueBindingType()
}

func (s *ElementValue) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ElementValue._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// The ``Type`` enumeration class defines the valid types for values in metadata elements.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type ElementValueTypeEnum string

const (
	// Indicates the type of the value is a long (64 bit signed integer).
	ElementValueType_LONG ElementValueTypeEnum = "LONG"
	// Indicates the type of the value is a string (a variable length sequence of characters). The encoding is UTF-8.
	ElementValueType_STRING ElementValueTypeEnum = "STRING"
	// Indicates the type of the value is a list of strings.
	ElementValueType_STRING_LIST ElementValueTypeEnum = "STRING_LIST"
	// Indicates the type of the value is an identifier for a structure element.
	ElementValueType_STRUCTURE_REFERENCE ElementValueTypeEnum = "STRUCTURE_REFERENCE"
	// Indicates the type of the value is a list of identifiers for a structure element.
	ElementValueType_STRUCTURE_REFERENCE_LIST ElementValueTypeEnum = "STRUCTURE_REFERENCE_LIST"
)

func (t ElementValueTypeEnum) ElementValueTypeEnum() bool {
	switch t {
	case ElementValueType_LONG:
		return true
	case ElementValueType_STRING:
		return true
	case ElementValueType_STRING_LIST:
		return true
	case ElementValueType_STRUCTURE_REFERENCE:
		return true
	case ElementValueType_STRUCTURE_REFERENCE_LIST:
		return true
	default:
		return false
	}
}

// The ``EnumerationInfo`` class contains the metamodel information of an enumeration element.
type EnumerationInfo struct {
	// Dot separated name of the enumeration element. The segments in the name reflect the organization of the APIs. The format of each segment is lower case with underscores. Each underscore represents a word boundary. If there are acronyms in the word, the capitalization is preserved. This format makes it easy to translate the segment into a different naming convention.
	Name string
	// Metamodel information of all the enumeration value elements contained in this enumeration element. The order of the enumeration value elements in the list is same as the order in which they are defined in the interface definition file.
	Values []EnumerationValueInfo
	// Generic metadata elements for an enumeration element. The key in the map is the name of the metadata element and the value is the data associated with that metadata element.
	//
	//  The MetadataIdentifier contains possible string values for keys in the map.
	Metadata map[string]ElementMap
	// English language documentation for an enumeration element. It can contain HTML markup and Javadoc tags. The first sentence of the enumeration documentation is a complete sentence that identifies the enumeration by name and summarizes the purpose of the enumeration. The documentation describes the context in which the enumeration is used.
	//
	//  The documentation also contains references to the context in which the enumeration is used. But if the enumeration is used in many contexts, the references may not be present.
	Documentation string
}

func (s *EnumerationInfo) GetType__() bindings.BindingType {
	return EnumerationInfoBindingType()
}

func (s *EnumerationInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for EnumerationInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// The ``EnumerationValueInfo`` class describes the enumeration constant in the enumeration class.
type EnumerationValueInfo struct {
	// Value in the enumerated type. All the characters in the string are capitalized.
	Value string
	// Additional metadata for enumeration value in the enumerated type. The key in the map is the name of the metadata element and the value is the data associated with that metadata element.
	//
	//  The MetadataIdentifier contains possible string values for keys in the map.
	Metadata map[string]ElementMap
	// English language documentation for an enumeration value. It can contain HTML markup and documentation tags (similar to Javadoc tags). The first statement will be a noun or verb phrase that describes the purpose of the enumeration value.
	Documentation string
}

func (s *EnumerationValueInfo) GetType__() bindings.BindingType {
	return EnumerationValueInfoBindingType()
}

func (s *EnumerationValueInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for EnumerationValueInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// The ``ErrorInfo`` class contains the metadata information about the error elements contained in an operation element.
type ErrorInfo struct {
	// Identifier for the structure element corresponding to the error that is being reported by the operation.
	StructureId string
	// The English language documentation for the service element. It can contain HTML markup and Javadoc tags.
	Documentation string
}

func (s *ErrorInfo) GetType__() bindings.BindingType {
	return ErrorInfoBindingType()
}

func (s *ErrorInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ErrorInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// The ``FieldInfo`` class contains metamodel information of a field element contained in a structure element.
type FieldInfo struct {
	// Name of the field element in a canonical format. The format is lower case with underscores. Each underscore represents a word boundary. If there are acronyms in the word, the capitalization is preserved. This format makes it easy to translate the segment into a different naming convention.
	Name string
	// Type information.
	Type_ Type
	// Generic metadata elements for the field element. The key in the map is the name of the metadata element and the value is the data associated with that metadata element.
	//
	//  The MetadataIdentifier contains possible string values for keys in the map.
	Metadata map[string]ElementMap
	// English language documentation for the service element. It can contain HTML markup and Javadoc tags.
	Documentation string
}

func (s *FieldInfo) GetType__() bindings.BindingType {
	return FieldInfoBindingType()
}

func (s *FieldInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for FieldInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// The ``GenericInstantiation`` class describes the type information of a typed element when the type is an instantiation of one of the generic types provided by the infrastructure.
type GenericInstantiation struct {
	// The generic type that is being instantiated.
	GenericType GenericInstantiationGenericTypeEnum
	// Type of the element parameter if the generic type instantiation is a GenericInstantiationGenericTypeEnum#GenericType_LIST, GenericInstantiationGenericTypeEnum#GenericType_OPTIONAL or GenericInstantiationGenericTypeEnum#GenericType_SET.
	ElementType *Type
	// Type of the key parameter of the map generic type instantiation. The map generic type has a key parameter and value parameter. The type of the value parameter is described by GenericInstantiation#mapValueType..
	MapKeyType *Type
	// Type of the value parameter of the map generic type instantiation. The map generic type has a key parameter and value parameter. The type of the key parameter is described by GenericInstantiation#mapKeyType..
	MapValueType *Type
}

func (s *GenericInstantiation) GetType__() bindings.BindingType {
	return GenericInstantiationBindingType()
}

func (s *GenericInstantiation) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for GenericInstantiation._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// The ``GenericType`` enumeration class provides enumeration constants for each of the generic types provided by the infrastructure.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type GenericInstantiationGenericTypeEnum string

const (
	// Indicates the generic type is a list.
	GenericInstantiationGenericType_LIST GenericInstantiationGenericTypeEnum = "LIST"
	// Indicates the generic type is a map.
	GenericInstantiationGenericType_MAP GenericInstantiationGenericTypeEnum = "MAP"
	// Indicates the generic type is an optional.
	GenericInstantiationGenericType_OPTIONAL GenericInstantiationGenericTypeEnum = "OPTIONAL"
	// Indicates the generic type is a set.
	GenericInstantiationGenericType_SET GenericInstantiationGenericTypeEnum = "SET"
)

func (g GenericInstantiationGenericTypeEnum) GenericInstantiationGenericTypeEnum() bool {
	switch g {
	case GenericInstantiationGenericType_LIST:
		return true
	case GenericInstantiationGenericType_MAP:
		return true
	case GenericInstantiationGenericType_OPTIONAL:
		return true
	case GenericInstantiationGenericType_SET:
		return true
	default:
		return false
	}
}

// The ``OperationInfo`` class contains metamodel information of an operation element.
type OperationInfo struct {
	// Name of the operation element in a canonical format. The format is lower case with underscores. Each underscore represents a word boundary. If there are acronyms in the word, the capitalization is preserved. This format makes it easy to translate the segment into a different naming convention.
	Name string
	// Metamodel information for the parameter elements. The order of the parameters elements in the list is same as the order of the parameters declared in the interface definition file.
	Params []FieldInfo
	// Metamodel type for the output element.
	Output OperationResultInfo
	// List of error elements that might be reported by the operation element. If the operation reports the same error for more than one reason, the list contains the error element associated with the error more than once with different documentation elements.
	Errors []ErrorInfo
	// Generic metadata elements for the operation element. The key in the map is the name of the metadata element and the value is the data associated with that metadata element.
	//
	//  The MetadataIdentifier contains possible string values for key in the map.
	Metadata map[string]ElementMap
	// English language documentation for the service element. It can contain HTML markup and Javadoc tags.
	Documentation string
}

func (s *OperationInfo) GetType__() bindings.BindingType {
	return OperationInfoBindingType()
}

func (s *OperationInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for OperationInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// The ``OperationResultInfo`` class contains the metamodel information of an operation result element.
//
//  An operation accepts a list of parameters and returns a result or an error. The ``OperationResultInfo`` describes the result element of an operation.
type OperationResultInfo struct {
	// Type information of the operation result element.
	Type_ Type
	// Generic metadata elements for the service element. The key in the map is the name of the metadata element and the value is the data associated with that metadata element.
	//
	//  The MetadataIdentifier contains possible string values for keys in the map.
	Metadata map[string]ElementMap
	// English language documentation for the operation result element. It can contain HTML markup and Javadoc tags.
	Documentation string
}

func (s *OperationResultInfo) GetType__() bindings.BindingType {
	return OperationResultInfoBindingType()
}

func (s *OperationResultInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for OperationResultInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// The ``PackageInfo`` class contains the metamodel information of all the service elements, structure elements and enumeration elements contained in the package element.
type PackageInfo struct {
	// Dot separated name of the package element. The segments in the name reflect the organization of the APIs. The format of each segment is lower case with underscores. Each underscore represents a word boundary. If there are acronyms in the word, the capitalization is preserved. This format makes it easy to translate the segment into a different naming convention.
	Name string
	// Metamodel information of all the structure elements contained in the package element. The key in the map is the identifier of the structure element and the value in the map is the metamodel information for the structure element.
	//
	//  This does not include the structure elements contained in the service elements that are contained in this package element.
	Structures map[string]StructureInfo
	// Metamodel information of all the enumeration elements contained in the package element. The key in the map is the identifier of the enumeration element and the value in the map is the metamodel information for the enumeration element.
	//
	//  This does not include the enumeration elements that are contained in the service elements of this package element or structure elements of this package element.
	Enumerations map[string]EnumerationInfo
	// Metamodel information of all the service elements contained in the package element. The key in the map is the identifier of the service element and the value in the map is the metamodel information for the service element.
	Services map[string]ServiceInfo
	// Generic metadata elements for the package element. The key in the map is the name of the metadata element and the value is the data associated with that metadata element.
	//
	//  The MetadataIdentifier contains possible string values for keys in the map.
	Metadata map[string]ElementMap
	// English language documentation for a package. It can contain HTML markup and Javadoc tags. The first sentence of the package documentation is a complete sentence that identifies the package by name and summarizes the purpose of the package.
	//
	//  The primary purpose of a package documentation is to provide high-level context that will provide a framework in which the users can put the detail about the package contents.
	Documentation string
}

func (s *PackageInfo) GetType__() bindings.BindingType {
	return PackageInfoBindingType()
}

func (s *PackageInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for PackageInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// The ``PrimitiveValue`` class contains value of the constant element.
type PrimitiveValue struct {
	// Type of the constant value.
	Type_ PrimitiveValueTypeEnum
	// Boolean value of the constant.
	BooleanValue *bool
	// Double value of the constant.
	DoubleValue *float64
	// Long value of the constant.
	LongValue *int64
	// String value of the constant.
	StringValue *string
}

func (s *PrimitiveValue) GetType__() bindings.BindingType {
	return PrimitiveValueBindingType()
}

func (s *PrimitiveValue) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for PrimitiveValue._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// The ``Type`` enumeration class defines the valid types for values in constant elements.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type PrimitiveValueTypeEnum string

const (
	// Indicates the value is a boolean (true or false).
	PrimitiveValueType_BOOLEAN PrimitiveValueTypeEnum = "BOOLEAN"
	// Indicates the value is a double (64 bit floating number).
	PrimitiveValueType_DOUBLE PrimitiveValueTypeEnum = "DOUBLE"
	// Indicates the value is a long (64 bit signed integer).
	PrimitiveValueType_LONG PrimitiveValueTypeEnum = "LONG"
	// Indicates the value is a string (a variable length sequence of characters). The encoding is UTF8.
	PrimitiveValueType_STRING PrimitiveValueTypeEnum = "STRING"
)

func (t PrimitiveValueTypeEnum) PrimitiveValueTypeEnum() bool {
	switch t {
	case PrimitiveValueType_BOOLEAN:
		return true
	case PrimitiveValueType_DOUBLE:
		return true
	case PrimitiveValueType_LONG:
		return true
	case PrimitiveValueType_STRING:
		return true
	default:
		return false
	}
}

// The ``ServiceInfo`` class contains the metamodel information of all the operation elements, structure elements and enumeration elements containted in a service element.
type ServiceInfo struct {
	// Dot separated name of the service element. The segments in the name reflect the organization of the APIs. The format of each segment is lower case with underscores. Each underscore represents a word boundary. If there are acronyms in the word, the capitalization is preserved. This format makes it easy to translate the segment into a different naming convention.
	Name string
	// Metamodel information of all the operation elements contained in the service element. The key in the map is the identifier of the operation element and the value in the map is the metamodel information for the operation element.
	Operations map[string]OperationInfo
	// Metamodel information of all the structure elements contained in the service element. The key in the map is the identifier of the structure element and the value in the map is the metamodel information for the structure element.
	Structures map[string]StructureInfo
	// Metamodel information of all the enumeration elements contained in the service element. The key in the map is the identifier of the enumeration element and the value in the map is the metamodel information for the enumeration element.
	Enumerations map[string]EnumerationInfo
	// Metamodel information of all the constant elements contained in the service element. The key in the map is the name of the constant element and the value in the map is the metamodel information for the contant element.
	Constants map[string]ConstantInfo
	// Generic metadata elements for the service element. The key in the map is the name of the metadata element and the value is the data associated with that metadata element.
	//
	//  The MetadataIdentifier contains possible string values for keys in the map.
	Metadata map[string]ElementMap
	// English language documentation for the service element. It can contain HTML markup and Javadoc tags. The first sentence of the service documentation is a complete sentence that identifies the service by name and summarizes the purpose of the service. The remaining part of the documentation provides a summary of how to use the operations defined in the service.
	Documentation string
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

// The ``StructureInfo`` class contains the metamodel information of all the field elements, constant elements and enumeration elements contained in the structure element.
//
//  In the interface definition language, API designers have the ability to include all the fields from one structure to another structure. This is done by using an annotation ``\\\\@Include`` on the structure in which we want to add the fields. If this annotation is present, the list of fields in the ``StructureInfo`` will also contain the fields that are being included. The annotation information is also retained in the StructureInfo#metadata element as well.
type StructureInfo struct {
	// Dot separated name of the structure element. The segments in the name reflect the organization of the APIs. The format of each segment is lower case with underscores. Each underscore represents a word boundary. If there are acronyms in the word, the capitalization is preserved. This format makes it easy to translate the segment into a different naming convention.
	Name string
	// Type of the structure.
	Type_ StructureInfoTypeEnum
	// Metamodel information of all the enumeration elements contained in the structure element. The key in the map is the identifier of the enumeration element and the value is the metamodel information of the enumeration element.
	Enumerations map[string]EnumerationInfo
	// Metamodel information of all the constant elements contained in the structure element. The key in the map is the name of the constant element and the value in the map is the metamodel information for the constant element.
	Constants map[string]ConstantInfo
	// Metamodel information of all the field elements. The order of the field elements in the list matches the order in which the fields are defined in the service.
	Fields []FieldInfo
	// Generic metadata elements for the structure element. The key in the map is the name of the metadata element and the value is the data associated with that metadata element.
	//
	//  The MetadataIdentifier contains possible string values for keys in the map.
	Metadata map[string]ElementMap
	// English language documentation for a structure element. It can contain HTML markup and Javadoc tags. The first sentence of the structure documentation is a complete sentence that identifies the structure by name and summarizes the purpose of the structure.
	Documentation string
}

func (s *StructureInfo) GetType__() bindings.BindingType {
	return StructureInfoBindingType()
}

func (s *StructureInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for StructureInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// The ``Type`` enumeration class defines the kind of this structure element. In the interface definition language, structure element and error element have similar characteristics. The difference is that only error elements can be used to describe the exceptions of an operation element.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type StructureInfoTypeEnum string

const (
	// If the type is a structure element.
	StructureInfoType_STRUCTURE StructureInfoTypeEnum = "STRUCTURE"
	// If the type is an error element.
	StructureInfoType_ERROR StructureInfoTypeEnum = "ERROR"
)

func (t StructureInfoTypeEnum) StructureInfoTypeEnum() bool {
	switch t {
	case StructureInfoType_STRUCTURE:
		return true
	case StructureInfoType_ERROR:
		return true
	default:
		return false
	}
}

// The ``Type`` class describes the type information of a typed element in the interface definiton language. The following elements in the metamodel are typed:
//
// * Field element in a structure element. See StructureInfo#fields
// * Parameter element in an operation element. See OperationInfo#params
// * Result element in an operation element. See OperationInfo#output
//
//  The type could be one of the three following categories:
//
// * Built-in types: These are types present in the interface definition language type system. They are provided by the infrastructure.
// * User defined named type: API designers can create custom types and use them for the typed elements. These types have a unique identifier.
// * Generic type instantiation: The language infrastructure also provides generic types such as list, map, set and so on. An instantiation of one of these generic types could also be used for the typed elements.
type Type struct {
	// Category of this type.
	Category TypeCategoryEnum
	// Category of the built-in type.
	BuiltinType *TypeBuiltinTypeEnum
	// Identifier and type of the user defined type.
	UserDefinedType *UserDefinedType
	// Instantiation of one of the generic types available in the interface definition language.
	GenericInstantiation *GenericInstantiation
}

func (s *Type) GetType__() bindings.BindingType {
	return TypeBindingType()
}

func (s *Type) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for Type._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// The ``Category`` enumeration class provides enumeration constant for each category of the type.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type TypeCategoryEnum string

const (
	// The type is one of the built-in types specified in TypeBuiltinTypeEnum
	TypeCategory_BUILTIN TypeCategoryEnum = "BUILTIN"
	// The type is one of the user defined named types.
	TypeCategory_USER_DEFINED TypeCategoryEnum = "USER_DEFINED"
	// The type is an instantiation of one of the generic types.
	TypeCategory_GENERIC TypeCategoryEnum = "GENERIC"
)

func (c TypeCategoryEnum) TypeCategoryEnum() bool {
	switch c {
	case TypeCategory_BUILTIN:
		return true
	case TypeCategory_USER_DEFINED:
		return true
	case TypeCategory_GENERIC:
		return true
	default:
		return false
	}
}

// The ``BuiltinType`` enumeration class provides enumeration constant for each of the built-in types present in the interface definition language type system.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type TypeBuiltinTypeEnum string

const (
	// The built-in type is a void. The value is null.
	TypeBuiltinType_VOID TypeBuiltinTypeEnum = "VOID"
	// The built-in type is a boolean. The value is true or false.
	TypeBuiltinType_BOOLEAN TypeBuiltinTypeEnum = "BOOLEAN"
	// The built-in type is a long. The value is a 64 bit signed integer.
	TypeBuiltinType_LONG TypeBuiltinTypeEnum = "LONG"
	// The built-in type is a double. The value is a 64 bit floating point number.
	TypeBuiltinType_DOUBLE TypeBuiltinTypeEnum = "DOUBLE"
	// The built-in type is a string. The value is a variable-length sequence of zero or more unicode characters.
	TypeBuiltinType_STRING TypeBuiltinTypeEnum = "STRING"
	// The built-in type is a binary. The value is a variable-length sequence of zero or more bytes.
	TypeBuiltinType_BINARY TypeBuiltinTypeEnum = "BINARY"
	// The built-in type is a secret. The value is a variable-length sequence of zero or more unicode characters. The value contains sensitive data that should not be printed or displayed anywhere.
	TypeBuiltinType_SECRET TypeBuiltinTypeEnum = "SECRET"
	// The built-in type is a datetime. The value should be in the UTC timezone and the precision is milliseconds.
	TypeBuiltinType_DATE_TIME TypeBuiltinTypeEnum = "DATE_TIME"
	// The built-in type is an ID. The value represents an identifier for a resource.
	TypeBuiltinType_ID TypeBuiltinTypeEnum = "ID"
	// The built-in type is an URI. The value follows the IRI specification in RFC 3987.
	TypeBuiltinType_URI TypeBuiltinTypeEnum = "URI"
	// The built-in type is an arbitrary exception type. This is used if the value of a typed element can be one of any user defined named type which is an exception.
	TypeBuiltinType_ANY_ERROR TypeBuiltinTypeEnum = "ANY_ERROR"
	// The built-in type is a dynamic structure. This is used if the value of a typed element can be one of any user defined named type.
	TypeBuiltinType_DYNAMIC_STRUCTURE TypeBuiltinTypeEnum = "DYNAMIC_STRUCTURE"
	// The built-in type is an opaque. This is used if the value of a typed element could be of any type and the actual type will be known only during the execution of the API. This is mostly used in infrastructure interfaces.
	TypeBuiltinType_OPAQUE TypeBuiltinTypeEnum = "OPAQUE"
)

func (b TypeBuiltinTypeEnum) TypeBuiltinTypeEnum() bool {
	switch b {
	case TypeBuiltinType_VOID:
		return true
	case TypeBuiltinType_BOOLEAN:
		return true
	case TypeBuiltinType_LONG:
		return true
	case TypeBuiltinType_DOUBLE:
		return true
	case TypeBuiltinType_STRING:
		return true
	case TypeBuiltinType_BINARY:
		return true
	case TypeBuiltinType_SECRET:
		return true
	case TypeBuiltinType_DATE_TIME:
		return true
	case TypeBuiltinType_ID:
		return true
	case TypeBuiltinType_URI:
		return true
	case TypeBuiltinType_ANY_ERROR:
		return true
	case TypeBuiltinType_DYNAMIC_STRUCTURE:
		return true
	case TypeBuiltinType_OPAQUE:
		return true
	default:
		return false
	}
}

// The ``UserDefinedType`` class contains the metamodel type information of a typed element whose type is a user defined named type.
type UserDefinedType struct {
	// Category of the user defined named type. The named type could be a structure element or an enumeration element.
	ResourceType string
	// Identifier of the user defined named type.
	ResourceId string
}

func (s *UserDefinedType) GetType__() bindings.BindingType {
	return UserDefinedTypeBindingType()
}

func (s *UserDefinedType) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for UserDefinedType._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

func ComponentDataBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["info"] = bindings.NewReferenceType(ComponentInfoBindingType)
	fieldNameMap["info"] = "Info"
	fields["fingerprint"] = bindings.NewStringType()
	fieldNameMap["fingerprint"] = "Fingerprint"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vapi.metadata.metamodel.component_data", fields, reflect.TypeOf(ComponentData{}), fieldNameMap, validators)
}

func ComponentInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["packages"] = bindings.NewMapType(bindings.NewIdType([]string{"com.vmware.vapi.package"}, ""), bindings.NewReferenceType(PackageInfoBindingType), reflect.TypeOf(map[string]PackageInfo{}))
	fieldNameMap["packages"] = "Packages"
	fields["metadata"] = bindings.NewMapType(bindings.NewStringType(), bindings.NewReferenceType(ElementMapBindingType), reflect.TypeOf(map[string]ElementMap{}))
	fieldNameMap["metadata"] = "Metadata"
	fields["documentation"] = bindings.NewStringType()
	fieldNameMap["documentation"] = "Documentation"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vapi.metadata.metamodel.component_info", fields, reflect.TypeOf(ComponentInfo{}), fieldNameMap, validators)
}

func ConstantInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewReferenceType(TypeBindingType)
	fieldNameMap["type"] = "Type_"
	fields["value"] = bindings.NewReferenceType(ConstantValueBindingType)
	fieldNameMap["value"] = "Value"
	fields["documentation"] = bindings.NewStringType()
	fieldNameMap["documentation"] = "Documentation"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vapi.metadata.metamodel.constant_info", fields, reflect.TypeOf(ConstantInfo{}), fieldNameMap, validators)
}

func ConstantValueBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["category"] = bindings.NewEnumType("com.vmware.vapi.metadata.metamodel.constant_value.category", reflect.TypeOf(ConstantValueCategoryEnum(ConstantValueCategory_PRIMITIVE)))
	fieldNameMap["category"] = "Category"
	fields["primitive_value"] = bindings.NewOptionalType(bindings.NewReferenceType(PrimitiveValueBindingType))
	fieldNameMap["primitive_value"] = "PrimitiveValue"
	fields["list_value"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(PrimitiveValueBindingType), reflect.TypeOf([]PrimitiveValue{})))
	fieldNameMap["list_value"] = "ListValue"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("category",
		map[string][]bindings.FieldData{
			"PRIMITIVE": []bindings.FieldData{
				bindings.NewFieldData("primitive_value", true),
			},
			"LIST": []bindings.FieldData{
				bindings.NewFieldData("list_value", true),
			},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vapi.metadata.metamodel.constant_value", fields, reflect.TypeOf(ConstantValue{}), fieldNameMap, validators)
}

func ElementMapBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["elements"] = bindings.NewMapType(bindings.NewStringType(), bindings.NewReferenceType(ElementValueBindingType), reflect.TypeOf(map[string]ElementValue{}))
	fieldNameMap["elements"] = "Elements"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vapi.metadata.metamodel.element_map", fields, reflect.TypeOf(ElementMap{}), fieldNameMap, validators)
}

func ElementValueBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewEnumType("com.vmware.vapi.metadata.metamodel.element_value.type", reflect.TypeOf(ElementValueTypeEnum(ElementValueType_LONG)))
	fieldNameMap["type"] = "Type_"
	fields["long_value"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["long_value"] = "LongValue"
	fields["string_value"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["string_value"] = "StringValue"
	fields["list_value"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["list_value"] = "ListValue"
	fields["structure_id"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.vapi.structure"}, ""))
	fieldNameMap["structure_id"] = "StructureId"
	fields["structure_ids"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewIdType([]string{"com.vmware.vapi.structure"}, ""), reflect.TypeOf([]string{})))
	fieldNameMap["structure_ids"] = "StructureIds"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
		map[string][]bindings.FieldData{
			"LONG": []bindings.FieldData{
				bindings.NewFieldData("long_value", true),
			},
			"STRING": []bindings.FieldData{
				bindings.NewFieldData("string_value", true),
			},
			"STRING_LIST": []bindings.FieldData{
				bindings.NewFieldData("list_value", true),
			},
			"STRUCTURE_REFERENCE": []bindings.FieldData{
				bindings.NewFieldData("structure_id", true),
			},
			"STRUCTURE_REFERENCE_LIST": []bindings.FieldData{
				bindings.NewFieldData("structure_ids", true),
			},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vapi.metadata.metamodel.element_value", fields, reflect.TypeOf(ElementValue{}), fieldNameMap, validators)
}

func EnumerationInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["values"] = bindings.NewListType(bindings.NewReferenceType(EnumerationValueInfoBindingType), reflect.TypeOf([]EnumerationValueInfo{}))
	fieldNameMap["values"] = "Values"
	fields["metadata"] = bindings.NewMapType(bindings.NewStringType(), bindings.NewReferenceType(ElementMapBindingType), reflect.TypeOf(map[string]ElementMap{}))
	fieldNameMap["metadata"] = "Metadata"
	fields["documentation"] = bindings.NewStringType()
	fieldNameMap["documentation"] = "Documentation"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vapi.metadata.metamodel.enumeration_info", fields, reflect.TypeOf(EnumerationInfo{}), fieldNameMap, validators)
}

func EnumerationValueInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["value"] = bindings.NewStringType()
	fieldNameMap["value"] = "Value"
	fields["metadata"] = bindings.NewMapType(bindings.NewStringType(), bindings.NewReferenceType(ElementMapBindingType), reflect.TypeOf(map[string]ElementMap{}))
	fieldNameMap["metadata"] = "Metadata"
	fields["documentation"] = bindings.NewStringType()
	fieldNameMap["documentation"] = "Documentation"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vapi.metadata.metamodel.enumeration_value_info", fields, reflect.TypeOf(EnumerationValueInfo{}), fieldNameMap, validators)
}

func ErrorInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["structure_id"] = bindings.NewIdType([]string{"com.vmware.vapi.structure"}, "")
	fieldNameMap["structure_id"] = "StructureId"
	fields["documentation"] = bindings.NewStringType()
	fieldNameMap["documentation"] = "Documentation"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vapi.metadata.metamodel.error_info", fields, reflect.TypeOf(ErrorInfo{}), fieldNameMap, validators)
}

func FieldInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["type"] = bindings.NewReferenceType(TypeBindingType)
	fieldNameMap["type"] = "Type_"
	fields["metadata"] = bindings.NewMapType(bindings.NewStringType(), bindings.NewReferenceType(ElementMapBindingType), reflect.TypeOf(map[string]ElementMap{}))
	fieldNameMap["metadata"] = "Metadata"
	fields["documentation"] = bindings.NewStringType()
	fieldNameMap["documentation"] = "Documentation"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vapi.metadata.metamodel.field_info", fields, reflect.TypeOf(FieldInfo{}), fieldNameMap, validators)
}

func GenericInstantiationBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["generic_type"] = bindings.NewEnumType("com.vmware.vapi.metadata.metamodel.generic_instantiation.generic_type", reflect.TypeOf(GenericInstantiationGenericTypeEnum(GenericInstantiationGenericType_LIST)))
	fieldNameMap["generic_type"] = "GenericType"
	fields["element_type"] = bindings.NewOptionalType(bindings.NewReferenceType(TypeBindingType))
	fieldNameMap["element_type"] = "ElementType"
	fields["map_key_type"] = bindings.NewOptionalType(bindings.NewReferenceType(TypeBindingType))
	fieldNameMap["map_key_type"] = "MapKeyType"
	fields["map_value_type"] = bindings.NewOptionalType(bindings.NewReferenceType(TypeBindingType))
	fieldNameMap["map_value_type"] = "MapValueType"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("generic_type",
		map[string][]bindings.FieldData{
			"LIST": []bindings.FieldData{
				bindings.NewFieldData("element_type", true),
			},
			"OPTIONAL": []bindings.FieldData{
				bindings.NewFieldData("element_type", true),
			},
			"SET": []bindings.FieldData{
				bindings.NewFieldData("element_type", true),
			},
			"MAP": []bindings.FieldData{
				bindings.NewFieldData("map_key_type", true),
				bindings.NewFieldData("map_value_type", true),
			},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vapi.metadata.metamodel.generic_instantiation", fields, reflect.TypeOf(GenericInstantiation{}), fieldNameMap, validators)
}

func OperationInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["params"] = bindings.NewListType(bindings.NewReferenceType(FieldInfoBindingType), reflect.TypeOf([]FieldInfo{}))
	fieldNameMap["params"] = "Params"
	fields["output"] = bindings.NewReferenceType(OperationResultInfoBindingType)
	fieldNameMap["output"] = "Output"
	fields["errors"] = bindings.NewListType(bindings.NewReferenceType(ErrorInfoBindingType), reflect.TypeOf([]ErrorInfo{}))
	fieldNameMap["errors"] = "Errors"
	fields["metadata"] = bindings.NewMapType(bindings.NewStringType(), bindings.NewReferenceType(ElementMapBindingType), reflect.TypeOf(map[string]ElementMap{}))
	fieldNameMap["metadata"] = "Metadata"
	fields["documentation"] = bindings.NewStringType()
	fieldNameMap["documentation"] = "Documentation"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vapi.metadata.metamodel.operation_info", fields, reflect.TypeOf(OperationInfo{}), fieldNameMap, validators)
}

func OperationResultInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewReferenceType(TypeBindingType)
	fieldNameMap["type"] = "Type_"
	fields["metadata"] = bindings.NewMapType(bindings.NewStringType(), bindings.NewReferenceType(ElementMapBindingType), reflect.TypeOf(map[string]ElementMap{}))
	fieldNameMap["metadata"] = "Metadata"
	fields["documentation"] = bindings.NewStringType()
	fieldNameMap["documentation"] = "Documentation"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vapi.metadata.metamodel.operation_result_info", fields, reflect.TypeOf(OperationResultInfo{}), fieldNameMap, validators)
}

func PackageInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["structures"] = bindings.NewMapType(bindings.NewIdType([]string{"com.vmware.vapi.structure"}, ""), bindings.NewReferenceType(StructureInfoBindingType), reflect.TypeOf(map[string]StructureInfo{}))
	fieldNameMap["structures"] = "Structures"
	fields["enumerations"] = bindings.NewMapType(bindings.NewIdType([]string{"com.vmware.vapi.enumeration"}, ""), bindings.NewReferenceType(EnumerationInfoBindingType), reflect.TypeOf(map[string]EnumerationInfo{}))
	fieldNameMap["enumerations"] = "Enumerations"
	fields["services"] = bindings.NewMapType(bindings.NewIdType([]string{"com.vmware.vapi.service"}, ""), bindings.NewReferenceType(ServiceInfoBindingType), reflect.TypeOf(map[string]ServiceInfo{}))
	fieldNameMap["services"] = "Services"
	fields["metadata"] = bindings.NewMapType(bindings.NewStringType(), bindings.NewReferenceType(ElementMapBindingType), reflect.TypeOf(map[string]ElementMap{}))
	fieldNameMap["metadata"] = "Metadata"
	fields["documentation"] = bindings.NewStringType()
	fieldNameMap["documentation"] = "Documentation"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vapi.metadata.metamodel.package_info", fields, reflect.TypeOf(PackageInfo{}), fieldNameMap, validators)
}

func PrimitiveValueBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewEnumType("com.vmware.vapi.metadata.metamodel.primitive_value.type", reflect.TypeOf(PrimitiveValueTypeEnum(PrimitiveValueType_BOOLEAN)))
	fieldNameMap["type"] = "Type_"
	fields["boolean_value"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["boolean_value"] = "BooleanValue"
	fields["double_value"] = bindings.NewOptionalType(bindings.NewDoubleType())
	fieldNameMap["double_value"] = "DoubleValue"
	fields["long_value"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["long_value"] = "LongValue"
	fields["string_value"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["string_value"] = "StringValue"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
		map[string][]bindings.FieldData{
			"BOOLEAN": []bindings.FieldData{
				bindings.NewFieldData("boolean_value", true),
			},
			"DOUBLE": []bindings.FieldData{
				bindings.NewFieldData("double_value", true),
			},
			"LONG": []bindings.FieldData{
				bindings.NewFieldData("long_value", true),
			},
			"STRING": []bindings.FieldData{
				bindings.NewFieldData("string_value", true),
			},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vapi.metadata.metamodel.primitive_value", fields, reflect.TypeOf(PrimitiveValue{}), fieldNameMap, validators)
}

func ServiceInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["operations"] = bindings.NewMapType(bindings.NewIdType([]string{"com.vmware.vapi.operation"}, ""), bindings.NewReferenceType(OperationInfoBindingType), reflect.TypeOf(map[string]OperationInfo{}))
	fieldNameMap["operations"] = "Operations"
	fields["structures"] = bindings.NewMapType(bindings.NewIdType([]string{"com.vmware.vapi.structure"}, ""), bindings.NewReferenceType(StructureInfoBindingType), reflect.TypeOf(map[string]StructureInfo{}))
	fieldNameMap["structures"] = "Structures"
	fields["enumerations"] = bindings.NewMapType(bindings.NewIdType([]string{"com.vmware.vapi.enumeration"}, ""), bindings.NewReferenceType(EnumerationInfoBindingType), reflect.TypeOf(map[string]EnumerationInfo{}))
	fieldNameMap["enumerations"] = "Enumerations"
	fields["constants"] = bindings.NewMapType(bindings.NewStringType(), bindings.NewReferenceType(ConstantInfoBindingType), reflect.TypeOf(map[string]ConstantInfo{}))
	fieldNameMap["constants"] = "Constants"
	fields["metadata"] = bindings.NewMapType(bindings.NewStringType(), bindings.NewReferenceType(ElementMapBindingType), reflect.TypeOf(map[string]ElementMap{}))
	fieldNameMap["metadata"] = "Metadata"
	fields["documentation"] = bindings.NewStringType()
	fieldNameMap["documentation"] = "Documentation"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vapi.metadata.metamodel.service_info", fields, reflect.TypeOf(ServiceInfo{}), fieldNameMap, validators)
}

func StructureInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["type"] = bindings.NewEnumType("com.vmware.vapi.metadata.metamodel.structure_info.type", reflect.TypeOf(StructureInfoTypeEnum(StructureInfoType_STRUCTURE)))
	fieldNameMap["type"] = "Type_"
	fields["enumerations"] = bindings.NewMapType(bindings.NewIdType([]string{"com.vmware.vapi.enumeration"}, ""), bindings.NewReferenceType(EnumerationInfoBindingType), reflect.TypeOf(map[string]EnumerationInfo{}))
	fieldNameMap["enumerations"] = "Enumerations"
	fields["constants"] = bindings.NewMapType(bindings.NewStringType(), bindings.NewReferenceType(ConstantInfoBindingType), reflect.TypeOf(map[string]ConstantInfo{}))
	fieldNameMap["constants"] = "Constants"
	fields["fields"] = bindings.NewListType(bindings.NewReferenceType(FieldInfoBindingType), reflect.TypeOf([]FieldInfo{}))
	fieldNameMap["fields"] = "Fields"
	fields["metadata"] = bindings.NewMapType(bindings.NewStringType(), bindings.NewReferenceType(ElementMapBindingType), reflect.TypeOf(map[string]ElementMap{}))
	fieldNameMap["metadata"] = "Metadata"
	fields["documentation"] = bindings.NewStringType()
	fieldNameMap["documentation"] = "Documentation"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vapi.metadata.metamodel.structure_info", fields, reflect.TypeOf(StructureInfo{}), fieldNameMap, validators)
}

func TypeBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["category"] = bindings.NewEnumType("com.vmware.vapi.metadata.metamodel.type.category", reflect.TypeOf(TypeCategoryEnum(TypeCategory_BUILTIN)))
	fieldNameMap["category"] = "Category"
	fields["builtin_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vapi.metadata.metamodel.type.builtin_type", reflect.TypeOf(TypeBuiltinTypeEnum(TypeBuiltinType_VOID))))
	fieldNameMap["builtin_type"] = "BuiltinType"
	fields["user_defined_type"] = bindings.NewOptionalType(bindings.NewReferenceType(UserDefinedTypeBindingType))
	fieldNameMap["user_defined_type"] = "UserDefinedType"
	fields["generic_instantiation"] = bindings.NewOptionalType(bindings.NewReferenceType(GenericInstantiationBindingType))
	fieldNameMap["generic_instantiation"] = "GenericInstantiation"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("category",
		map[string][]bindings.FieldData{
			"BUILTIN": []bindings.FieldData{
				bindings.NewFieldData("builtin_type", true),
			},
			"USER_DEFINED": []bindings.FieldData{
				bindings.NewFieldData("user_defined_type", true),
			},
			"GENERIC": []bindings.FieldData{
				bindings.NewFieldData("generic_instantiation", true),
			},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vapi.metadata.metamodel.type", fields, reflect.TypeOf(Type{}), fieldNameMap, validators)
}

func UserDefinedTypeBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["resource_type"] = bindings.NewStringType()
	fieldNameMap["resource_type"] = "ResourceType"
	fields["resource_id"] = bindings.NewIdType([]string{"com.vmware.vapi.structure", "com.vmware.vapi.enumeration"}, "resource_type")
	fieldNameMap["resource_id"] = "ResourceId"
	var validators = []bindings.Validator{}
	isv1 := bindings.NewIsOneOfValidator(
		"resource_type",
		[]string{
			"com.vmware.vapi.structure",
			"com.vmware.vapi.enumeration",
		},
	)
	validators = append(validators, isv1)
	return bindings.NewStructType("com.vmware.vapi.metadata.metamodel.user_defined_type", fields, reflect.TypeOf(UserDefinedType{}), fieldNameMap, validators)
}
