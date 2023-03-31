/* Copyright © 2019-2020 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package bindings

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"reflect"
)

type BindingType interface {
	Definition() data.DataDefinition
	Type() data.DataType
}

type VoidType struct{}

func (i VoidType) Definition() data.DataDefinition {
	return data.NewVoidDefinition()
}

func (i VoidType) Type() data.DataType {
	return data.VOID
}

func NewVoidType() VoidType {
	return VoidType{}
}

type IntegerType struct {
}

func (i IntegerType) Definition() data.DataDefinition {
	return data.NewIntegerDefinition()
}

func (i IntegerType) Type() data.DataType {
	return data.INTEGER
}

func NewIntegerType() IntegerType {
	return IntegerType{}
}

//implements BindingType
type StringType struct{}

func NewStringType() StringType {
	return StringType{}
}
func (s StringType) Definition() data.DataDefinition {
	return data.NewStringDefinition()
}

func (i StringType) Type() data.DataType {
	return data.STRING
}

type BooleanType struct{}

func NewBooleanType() BooleanType {
	return BooleanType{}
}
func (b BooleanType) Definition() data.DataDefinition {
	return data.NewBooleanDefinition()
}

func (i BooleanType) Type() data.DataType {
	return data.BOOLEAN
}

type OptionalType struct {
	elementType BindingType
}

func NewOptionalType(elementType BindingType) OptionalType {
	return OptionalType{elementType: elementType}
}

func (o OptionalType) Definition() data.DataDefinition {
	return data.NewOptionalDefinition(o.elementType.Definition())
}

func (i OptionalType) Type() data.DataType {
	return data.OPTIONAL
}

func (o OptionalType) ElementType() BindingType {
	return o.elementType
}

// ListType Representation of List IDL in Golang Binding
type ListType struct {
	elementType BindingType
	//this is necessary when ListType is the top most type for conversion.
	// When list is part of a struct, this is not used.
	// When list is part of a struct, struct will be the top most type for conversion.
	// so bindingstruct is not necessary.
	bindingStruct reflect.Type
}

func NewListType(elementType BindingType, bindingStruct reflect.Type) ListType {
	return ListType{elementType: elementType, bindingStruct: bindingStruct}
}

func (l ListType) SetBindingStruct(typ reflect.Type) {
	l.bindingStruct = typ
}

func (l ListType) BindingStruct() reflect.Type {
	return l.bindingStruct
}

func (l ListType) ElementType() BindingType {
	return l.elementType
}

func (l ListType) Definition() data.DataDefinition {
	return data.NewListDefinition(l.elementType.Definition())
}

func (i ListType) Type() data.DataType {
	return data.LIST
}

type OpaqueType struct {
}

func NewOpaqueType() OpaqueType {
	return OpaqueType{}
}

func (o OpaqueType) Definition() data.DataDefinition {
	return data.NewOpaqueDefinition()
}

func (i OpaqueType) Type() data.DataType {
	return data.OPAQUE
}

type StructType struct {
	name              string
	fields            map[string]BindingType
	bindingStruct     reflect.Type
	canonicalFieldMap map[string]string
	validators        []Validator
}

func NewStructType(name string,
	fields map[string]BindingType,
	bindingClass reflect.Type,
	canonicalFieldMap map[string]string,
	validators []Validator) StructType {
	return StructType{name: name, fields: fields, bindingStruct: bindingClass,
		canonicalFieldMap: canonicalFieldMap, validators: validators}
}

func (s StructType) Name() string {
	return s.name
}

func (s StructType) BindingStruct() reflect.Type {
	return s.bindingStruct
}

func (s StructType) Field(fieldName string) BindingType {
	return s.fields[fieldName]
}

func (s StructType) CanonicalField(fieldName string) string {
	return s.canonicalFieldMap[fieldName]
}

func (s StructType) FieldNames() []string {
	var keys = make([]string, 0)
	for key, _ := range s.fields {
		keys = append(keys, key)
	}
	return keys
}

func (s StructType) Definition() data.DataDefinition {
	fieldDefMap := make(map[string]data.DataDefinition)

	for key, field := range s.fields {
		fieldDefMap[key] = field.Definition()
	}
	var result = data.NewStructDefinition(s.name, fieldDefMap)
	return result
}

func (i StructType) Type() data.DataType {
	return data.STRUCTURE
}

func (s StructType) Validate(structValue *data.StructValue) []error {
	if s.validators != nil {
		for _, v := range s.validators {
			msgs := v.Validate(structValue)
			if msgs != nil || len(msgs) > 0 {
				return msgs
			}
		}
	}
	return nil
}

type MapType struct {
	KeyType       BindingType
	ValueType     BindingType
	bindingStruct reflect.Type
}

func NewMapType(keyType BindingType, valueType BindingType, bindingStruct reflect.Type) MapType {
	return MapType{KeyType: keyType, ValueType: valueType, bindingStruct: bindingStruct}
}

func (m MapType) Definition() data.DataDefinition {
	fieldDefs := make(map[string]data.DataDefinition)
	fieldDefs[lib.MAP_KEY_FIELD] = m.KeyType.Definition()
	fieldDefs[lib.MAP_VALUE_FIELD] = m.ValueType.Definition()
	elementDef := data.NewStructDefinition(lib.MAP_ENTRY, fieldDefs)
	return data.NewListDefinition(elementDef)
}

func (i MapType) Type() data.DataType {
	return data.LIST
}

type IdType struct {
	ResourceTypes      []string
	ResourceTypeHolder string
}

func NewIdType(resourceTypes []string, typeHolder string) IdType {
	return IdType{ResourceTypes: resourceTypes, ResourceTypeHolder: typeHolder}
}

func (i IdType) Definition() data.DataDefinition {
	return data.NewStringDefinition()
}

func (i IdType) Type() data.DataType {
	return data.STRING
}

type EnumType struct {
	name          string
	bindingStruct reflect.Type
}

func (e EnumType) Name() string {
	return e.name
}

func (e EnumType) BindingStruct() reflect.Type {
	return e.bindingStruct
}

func NewEnumType(name string, bindingStruct reflect.Type) EnumType {
	return EnumType{name: name, bindingStruct: bindingStruct}
}

func (e EnumType) Definition() data.DataDefinition {
	return data.NewStringDefinition()
}

func (i EnumType) Type() data.DataType {
	return data.STRING
}

type SetType struct {
	elementType   BindingType
	bindingStruct reflect.Type
}

func NewSetType(elementType BindingType, bindingStruct reflect.Type) SetType {
	return SetType{elementType: elementType, bindingStruct: bindingStruct}
}
func (s SetType) ElementType() BindingType {
	return s.elementType
}
func (s SetType) BindingStruct() reflect.Type {
	return s.bindingStruct
}

func (s SetType) Definition() data.DataDefinition {
	return data.NewListDefinition(s.elementType.Definition())
}

func (i SetType) Type() data.DataType {
	return data.LIST
}

type ErrorType struct {
	name              string
	fields            map[string]BindingType
	bindingStruct     reflect.Type
	canonicalFieldMap map[string]string
}

func NewErrorType(name string, fields map[string]BindingType, bindingClass reflect.Type, canonicalFieldMap map[string]string) ErrorType {
	return ErrorType{name: name, fields: fields, bindingStruct: bindingClass, canonicalFieldMap: canonicalFieldMap}
}
func (e ErrorType) Name() string {
	return e.name
}
func (e ErrorType) BindingStruct() reflect.Type {
	return e.bindingStruct
}

func (e ErrorType) Field(fieldName string) BindingType {
	return e.fields[fieldName]
}

func (e ErrorType) FieldNames() []string {
	var keys = make([]string, 0)
	for key, _ := range e.fields {
		keys = append(keys, key)
	}
	return keys
}

func (e ErrorType) Definition() data.DataDefinition {
	fieldDefMap := make(map[string]data.DataDefinition)

	for key, field := range e.fields {
		fieldDefMap[key] = field.Definition()
	}
	var result = data.NewErrorDefinition(e.name, fieldDefMap)
	return result
}

func (i ErrorType) Type() data.DataType {
	return data.ERROR
}

type DynamicStructType struct {
	name          string
	fields        map[string]BindingType
	bindingStruct reflect.Type
	validator     Validator
}

func NewDynamicStructType(hasFieldsOfTypes []ReferenceType) DynamicStructType {
	return DynamicStructType{
		name:          "vmware.vapi.dynamic_struct",
		bindingStruct: StructBindingType,
		validator:     NewHasFieldsOfValidator(hasFieldsOfTypes),
	}
}

func (d DynamicStructType) Name() string {
	return d.name
}
func (d DynamicStructType) BindingStruct() reflect.Type {
	return d.bindingStruct
}

func (d DynamicStructType) Field(fieldName string) BindingType {
	return d.fields[fieldName]
}

func (d DynamicStructType) FieldNames() []string {
	var keys = make([]string, 0)
	for key, _ := range d.fields {
		keys = append(keys, key)
	}
	return keys
}

func (d DynamicStructType) Definition() data.DataDefinition {
	return data.NewDynamicStructDefinition()
}

func (i DynamicStructType) Type() data.DataType {
	return data.DYNAMIC_STRUCTURE
}

func (d DynamicStructType) Validate(structValue *data.StructValue) []error {
	if d.validator != nil {
		return d.validator.Validate(structValue)
	}
	return nil
}

type DoubleType struct {
}

func NewDoubleType() DoubleType {
	return DoubleType{}
}

func (d DoubleType) Definition() data.DataDefinition {
	return data.NewDoubleDefinition()
}

func (i DoubleType) Type() data.DataType {
	return data.DOUBLE
}

type DateTimeType struct {
}

func NewDateTimeType() DateTimeType {
	return DateTimeType{}
}

func (d DateTimeType) Definition() data.DataDefinition {
	return data.NewStringDefinition()
}

func (i DateTimeType) Type() data.DataType {
	return data.STRING
}

type BlobType struct {
}

func NewBlobType() BlobType {
	return BlobType{}
}

func (b BlobType) Definition() data.DataDefinition {
	return data.NewBlobDefinition()
}

func (i BlobType) Type() data.DataType {
	return data.BLOB
}

type SecretType struct {
}

func NewSecretType() SecretType {
	return SecretType{}
}

func (s SecretType) Definition() data.DataDefinition {
	return data.NewSecretDefinition()
}

func (i SecretType) Type() data.DataType {
	return data.SECRET
}

type UriType struct {
}

func NewUriType() UriType {
	return UriType{}
}

func (u UriType) Definition() data.DataDefinition {
	return data.NewStringDefinition()
}

func (i UriType) Type() data.DataType {
	return data.STRING
}

type AnyErrorType struct {
}

func NewAnyErrorType() AnyErrorType {
	return AnyErrorType{}
}

func (e AnyErrorType) Definition() data.DataDefinition {
	fieldDefMap := make(map[string]data.DataDefinition)
	var result = data.NewStructDefinition("Exception", fieldDefMap)
	return result
}

func (i AnyErrorType) Type() data.DataType {
	return data.ANY_ERROR
}

type BindingTypeFunction func() BindingType
type ReferenceType struct {
	Fn BindingTypeFunction
}

func NewReferenceType(fn BindingTypeFunction) ReferenceType {
	return ReferenceType{Fn: fn}
}

func (r ReferenceType) Resolve() BindingType {
	return r.Fn()
}

func (r ReferenceType) Definition() data.DataDefinition {
	return r.Fn().Definition()
}

func (i ReferenceType) Type() data.DataType {
	return data.STRUCTURE_REF
}
