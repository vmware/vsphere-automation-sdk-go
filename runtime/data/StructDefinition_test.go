/* **********************************************************
 * Copyright 2018 VMware, Inc.  All rights reserved. -- VMware Confidential
 * **********************************************************/
package data_test

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/l10n"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/lib"
)

var SD_NAME_STRUCT_WITH_ONE_FIELD = "struct with one field"
var NAME_STRUCT_WITH_NO_FIELDS = "struct with no fields"
var STRUCT_DEF_MAP_KEY = "one"

func getStructDefWithNoFields() data.StructDefinition {
	var mapNoFields = make(map[string]data.DataDefinition)
	var result = data.NewStructDefinition(NAME_STRUCT_WITH_NO_FIELDS, mapNoFields)
	return result
}

func getStructDefWithStringField() data.StructDefinition {
	var mapOneFields = make(map[string]data.DataDefinition, 1)
	mapOneFields[STRUCT_DEF_MAP_KEY] = data.NewStringDefinition()
	var structDef = data.NewStructDefinition(SD_NAME_STRUCT_WITH_ONE_FIELD, mapOneFields)
	return structDef
}

func getStructDefWithOptStringField() data.StructDefinition {
	var map1 = make(map[string]data.DataDefinition, 1)
	var optionalDef = data.NewOptionalDefinition(data.NewStringDefinition())
	map1[STRUCT_DEF_MAP_KEY] = optionalDef
	var structDef = data.NewStructDefinition(NAME_STRUCT_WITH_ONE_FIELD, map1)
	return structDef
}

var structDefWithNoFields = getStructDefWithNoFields()
var structDefWithStringField = getStructDefWithStringField()
var structDefWithOptStringField = getStructDefWithOptStringField()

//func TestConstructionNilName(t *testing.T) {
//	var fieldMap= make(map[string]DataDefinition, 0)
//	var _, error= NewStructDefinition("", fieldMap)
//	if error == nil {
//		t.Error("Expected error but did not get one")
//	}
//}

func TestStructGetType(t *testing.T) {
	var structDef = getStructDefWithNoFields()
	if structDef.Type() != data.STRUCTURE {
		t.Error("Expected STRUCTURE datatype but found " + structDef.Type().String())
	}
}

func TestGetNameReturnsStructureName(t *testing.T) {

	if NAME_STRUCT_WITH_NO_FIELDS != structDefWithNoFields.Name() {
		t.Error("Expected " + NAME_STRUCT_WITH_NO_FIELDS + " but got " + structDefWithNoFields.Name())
	}
}

func TestGetFieldNamesReturnsEmptySetForStructDefWithNoFields(t *testing.T) {
	var structDef = getStructDefWithNoFields()
	if len(structDef.FieldNames()) > 0 {
		t.Error("Expected empty collection")
	}
}

func TestGetFieldNamesReturnsAllFields(t *testing.T) {
	var structDef = getStructDefWithStringField()
	if len(structDef.FieldNames()) != 1 || structDef.FieldNames()[0] != STRUCT_DEF_MAP_KEY {
		t.Error("Expected non empty slice with value " + STRUCT_DEF_MAP_KEY)
	}
}

//TODO: sreeshas
//when stringdef is a implemented as singleton, this test should change to reflect that.
func TestGetFieldReturnsFieldWhenNameMatches(t *testing.T) {
	var structDef = getStructDefWithStringField()
	var stringDef = structDef.Field(STRUCT_DEF_MAP_KEY)
	if stringDef.Type().String() != data.NewStringDefinition().Type().String() {
		t.Error("Expected StringDefinition")
	}
}

func TestGetFieldReturnsNilWhenNameDoesNotMatch(t *testing.T) {
	var structDef = getStructDefWithStringField()
	if structDef.Field("no such field") != nil {
		t.Error("Expected nil")
	}
}

func TestHasFieldReturnsTrueWhenNameMatches(t *testing.T) {
	var structDef = getStructDefWithStringField()
	if structDef.HasField(STRUCT_DEF_MAP_KEY) == false {
		t.Error("Expected true")
	}
}

func TestHasFieldReturnsFalseWhenNameDoesNotMatch(t *testing.T) {
	var structDef = getStructDefWithStringField()
	if structDef.HasField(STRUCT_DEF_MAP_KEY+NAME_STRUCT_WITH_ONE_FIELD) == true {
		t.Error("Expected false")
	}
}

func TestNullIsNotInstanceOfStructDef(t *testing.T) {
	var structDef = getStructDefWithStringField()
	if structDef.ValidInstanceOf(nil) == true {
		t.Error("Expected false")
	}
}

func TestStringValueIsNotInstanceOfStructDef(t *testing.T) {
	var anyString = "anystring"
	var value = data.NewStringValue(anyString)
	var structDef = getStructDefWithStringField()
	if structDef.ValidInstanceOf(value) == true {
		t.Error("Expected false")
	}
}

func TestStructWithNoFieldsIsInstanceOfStructDefWithNoFields(t *testing.T) {
	var structDef = getStructDefWithNoFields()
	var value = data.NewStructValue(structDef.Name(), nil)
	if structDef.ValidInstanceOf(value) == false {
		t.Error("Expected true")
	}
}

func TestStructIsNotInstanceOfStructDefWithDifferentName(t *testing.T) {
	var value = data.NewStructValue(structDefWithNoFields.Name()+"z", nil)
	if structDefWithNoFields.ValidInstanceOf(value) == true {
		t.Error("Expected false")
	}
}

func TestStructWithAllRequiredFieldsIsInstanceOf(t *testing.T) {
	var structValue = data.NewStructValue(structDefWithStringField.Name(), nil)
	var str = ""
	var stringValue = data.NewStringValue(str)

	structValue.SetField(STRUCT_DEF_MAP_KEY, stringValue)
	if structDefWithStringField.ValidInstanceOf(structValue) == false {
		t.Error("Expected true")
	}
}

func TestStructWithMissingFieldIsNotInstanceOf(t *testing.T) {
	var value = data.NewStructValue(structDefWithStringField.Name(), nil)
	if structDefWithStringField.ValidInstanceOf(value) == true {
		t.Error("Expected false")
	}
}

func TestStructWithMissingOptionalFieldIsNotInstanceOf(t *testing.T) {
	var value = data.NewStructValue(structDefWithOptStringField.Name(), nil)
	if structDefWithOptStringField.ValidInstanceOf(value) == true {
		t.Error("Expected false")
	}
}

func TestStructWithUnsetOptionalFieldIsInstanceOf(t *testing.T) {
	var value = data.NewStructValue(structDefWithOptStringField.Name(), nil)
	value.SetField(STRUCT_DEF_MAP_KEY, data.NewOptionalValue(nil))
	if structDefWithOptStringField.ValidInstanceOf(value) == false {
		t.Error("Expected true")
	}
}

func TestStructWithSetOptionalFieldIsInstanceOf(t *testing.T) {
	var value = data.NewStructValue(structDefWithOptStringField.Name(), nil)
	var somestring = "something"
	var strValue = data.NewStringValue(somestring)
	var optValue = data.NewOptionalValue(strValue)
	value.SetField(STRUCT_DEF_MAP_KEY, optValue)
	if structDefWithOptStringField.ValidInstanceOf(value) == false {
		t.Error("Expected true")
	}
}

func TestStructWithExtraFieldIsInstanceOf(t *testing.T) {
	var value = data.NewStructValue(structDefWithStringField.Name(), nil)
	var str = ""
	var strValue = data.NewStringValue(str)
	value.SetField(STRUCT_DEF_MAP_KEY, strValue)
	value.SetField("s"+STRUCT_DEF_MAP_KEY, strValue)
	if structDefWithStringField.ValidInstanceOf(value) == false {
		t.Error("Expected true")
	}

}

func TestStructWithFieldOfIncorrectTypeIsNotInstanceOf(t *testing.T) {
	var value = data.NewStructValue(structDefWithStringField.Name(), nil)
	value.SetField(STRUCT_DEF_MAP_KEY, data.NewOptionalValue(nil))
	if structDefWithStringField.ValidInstanceOf(value) == true {
		t.Error("Expected false")
	}
}

func TestCompleteValueIgnoresNonSetOptionalValue(t *testing.T) {
	var optValue = data.NewOptionalValue(nil)
	structDefWithOptStringField.CompleteValue(optValue)
	if optValue.IsSet() == true {
		t.Error("Expected false")
	}
}

func TestStructDefCompleteValueIgnoresSetOptionalValue(t *testing.T) {
	var str = "something"
	var strValue = data.NewStringValue(str)
	var optValue = data.NewOptionalValue(strValue)
	structDefWithOptStringField.CompleteValue(optValue)
	if optValue.IsSet() == false {
		t.Error("Expected true")
	}

}

func TestStructCompleteValueIgnoresSetOptionalField(t *testing.T) {
	var str = "something"
	var strValue = data.NewStringValue(str)
	var optValue = data.NewOptionalValue(strValue)
	var structValue = data.NewStructValue(structDefWithOptStringField.Name(), nil)
	structValue.SetField(STRUCT_DEF_MAP_KEY, optValue)
	// optional field is present - validation passes
	if structDefWithOptStringField.ValidInstanceOf(structValue) == false {
		t.Error("Expected true")
	}
	structDefWithOptStringField.CompleteValue(structValue)
	// nothing has changed - validation still passes
	if structDefWithOptStringField.ValidInstanceOf(structValue) == false {
		t.Error("Expected true")
	}

}

func TestStructCompleteValueIgnoresUnSetOptionalField(t *testing.T) {
	var structValue = data.NewStructValue(structDefWithOptStringField.Name(), nil)
	structValue.SetField(STRUCT_DEF_MAP_KEY, data.NewOptionalValue(nil))
	// optional field is present - validation passes
	if structDefWithOptStringField.ValidInstanceOf(structValue) == false {
		t.Error("Expected true")
	}
	structDefWithOptStringField.CompleteValue(structValue)
	//// nothing has changed - validation still passes
	if structDefWithOptStringField.ValidInstanceOf(structValue) == false {
		t.Error("Expected true")
	}
}

func TestStructCompleteValueIgnoresRequiredField(t *testing.T) {
	var structValue = data.NewStructValue(structDefWithStringField.Name(), nil)
	var str = "something"
	var strValue = data.NewStringValue(str)
	structValue.SetField(STRUCT_DEF_MAP_KEY, strValue)
	// optional field is present - validation passes
	if structDefWithStringField.ValidInstanceOf(structValue) == false {
		t.Error("Expected true")
	}
	structDefWithStringField.CompleteValue(structValue)
	var stValue, _ = structValue.Field(STRUCT_DEF_MAP_KEY)

	if stValue.(*data.StringValue).Value() != strValue.Value() {
		t.Error("Expected value mismatch")
	}
	if len(structDefWithStringField.FieldNames()) != 1 {
		t.Error("Expected 1")
	}
	//// nothing has changed - validation still passes
	if structDefWithStringField.ValidInstanceOf(structValue) == false {
		t.Error("Expected true")
	}
}

func TestStructCompleteValueIgnoresMissingRequiredField(t *testing.T) {
	var structValue = data.NewStructValue(structDefWithStringField.Name(), nil)
	// required field is missing - validation fails
	if structDefWithStringField.ValidInstanceOf(structValue) == true {
		t.Error("Expected false")
	}
	structDefWithStringField.CompleteValue(structValue)
	// nothing has changed - validation still fails
	if structDefWithStringField.ValidInstanceOf(structValue) == true {
		t.Error("Expected false")
	}
	if len(structValue.FieldNames()) != 0 {
		t.Error("Expected empty list")
	}
}

func TestStructCompleteValueIgnoresUndefinedField(t *testing.T) {
	//Todo
	//when integervalue is implemented
}

func TestStructCompleteValueIgnoresFieldOfIncorrectType(t *testing.T) {
	//TODO
	// when integervalue is implemented
}

func TestStructCompleteValueFillsMissingOptionalField(t *testing.T) {
	var structValue = data.NewStructValue(structDefWithOptStringField.Name(), nil)
	structDefWithOptStringField.CompleteValue(structValue)
	// field was added with unset optional value
	var optValue, _ = structValue.Field(STRUCT_DEF_MAP_KEY)
	if optValue == nil || optValue.Type() != data.OPTIONAL {
		t.Error("Expected non nil optional value")
	}
}

func TestStructCompleteValueFillsMissingOptionalFieldIfStructValueHasIncorrectName(t *testing.T) {
	var structValue = data.NewStructValue(structDefWithOptStringField.Name()+"incorrect", nil)
	structDefWithOptStringField.CompleteValue(structValue)
	// field was added with unset optional value
	var optValue, _ = structValue.Field(STRUCT_DEF_MAP_KEY)
	if optValue == nil {
		t.Error("Expected non nil value")
	}
}

func TestStructCompareItToItself(t *testing.T) {
	if structDefWithOptStringField.Equals(structDefWithOptStringField) == false {
		t.Error("Expected true")
	}
}

func TestCompareStructDefWithNoFieldsToOtherStructDefWithNoFields(t *testing.T) {
	var structDef = getStructDefWithNoFields()
	var emptymap = make(map[string]data.DataDefinition)
	var nameStructNoField = "struct with no fields"
	var struct1Def = data.NewStructDefinition(nameStructNoField, emptymap)
	if structDef.Equals(struct1Def) == false {
		t.Error("Expected true")
	}
}

func TestCompareStructDefWithStringFieldToOtherStructDefWithSameField(t *testing.T) {
	var myMap = make(map[string]data.DataDefinition)
	myMap[STRUCT_DEF_MAP_KEY] = data.NewStringDefinition()
	var structDef = data.NewStructDefinition(NAME_STRUCT_WITH_ONE_FIELD, myMap)
	if structDef.Equals(getStructDefWithStringField()) == false {
		t.Error("Expected true")
	}
}

func TestCompareStructToNill(t *testing.T) {
	if structDefWithOptStringField.Equals(nil) == true {
		t.Error("Expected false")
	}
}

func TestCompareStructDefToStringDef(t *testing.T) {
	if structDefWithStringField.Equals(data.NewStringDefinition()) == true {
		t.Error("Expected false")
	}
}

func TestCompateStructDefToOptionalDef(t *testing.T) {
	var optDef = data.NewOptionalDefinition(nil)
	if structDefWithStringField.Equals(optDef) == true {
		t.Error("Expected false")
	}
}

func TestCompareStructDefToStructDefWithDifferentName(t *testing.T) {
	var xyz = "xyz"
	var myMap = make(map[string]data.DataDefinition)
	var dataDef = data.NewStructDefinition(xyz, myMap)
	if structDefWithNoFields.Equals(dataDef) == true {
		t.Error("Expected false")
	}
}

func TestCompareStructDefToStructDefWithSameNameButDifferentNumberOfFields(t *testing.T) {
	var myMap = make(map[string]data.DataDefinition, 2)
	var sd = data.NewStringDefinition()
	myMap["xyz"] = sd
	myMap["abc"] = sd
	var otherStruct = data.NewStructDefinition(SD_NAME_STRUCT_WITH_ONE_FIELD, myMap)
	if otherStruct.Equals(structDefWithStringField) == true {
		t.Error("Expected false")
	}
}

func TestCompareStructDefWithStringFieldToStructDefWithStringFieldIfFieldNamesDiffer(t *testing.T) {
	var myMap = make(map[string]data.DataDefinition, 1)
	var sd = data.NewStringDefinition()
	myMap["xyz"] = sd
	var otherStruct = data.NewStructDefinition(SD_NAME_STRUCT_WITH_ONE_FIELD, myMap)
	if otherStruct.Equals(structDefWithStringField) == true {
		t.Error("Expected false")
	}
}

//func TestCompareStructDefWithOptionalFieldToStructDefWithStringFieldEvenIfFieldNamesMatch(t *testing.T) {
//	var myMap= make(map[string]DataDefinition, 1)
//	var od = NewOptionalDefinition(nil)
//	myMap[STRUCT_DEF_MAP_KEY] = od
//	var otherStruct, _ = NewStructDefinition(SD_NAME_STRUCT_WITH_ONE_FIELD, myMap)
//	if otherStruct.Equals(structDefWithStringField) == true {
//		t.Error("Expected false")
//	}
//}

func TestValidateStructValueValid(t *testing.T) {
	var structValue = data.NewStructValue(structDefWithStringField.Name(), nil)
	var str = "some test string"
	var strValue = data.NewStringValue(str)
	structValue.SetField(STRUCT_DEF_MAP_KEY, strValue)
	if len(structDefWithStringField.Validate(structValue)) != 0 {
		t.Error("Expected no errors")
	}
}

func TestValidateStructValueInvalidForName(t *testing.T) {
	var structValue = data.NewStructValue(structDefWithStringField.Name()+"1", nil)
	var str = "some test string"
	var strValue = data.NewStringValue(str)
	structValue.SetField(STRUCT_DEF_MAP_KEY, strValue)
	var errors = structDefWithStringField.Validate(structValue)
	if errors[0].(*l10n.Error).ID() != "vapi.data.structure.name.mismatch" {
		t.Error("Expected vapi.data.structure.name.mismatch")
	}
}

func TestValidateStructValueInvalidForFieldValue(t *testing.T) {
	var structValue = data.NewStructValue(structDefWithStringField.Name(), nil)
	var str = "some test string"
	var strValue = data.NewStringValue(str)
	var optValue = data.NewOptionalValue(strValue)
	structValue.SetField(STRUCT_DEF_MAP_KEY, optValue)
	var errors = structDefWithStringField.Validate(structValue)
	assert.Equal(t, 2, len(errors))
	assert.Equal(t, "vapi.data.validate.mismatch", errors[0].(*l10n.Error).ID())
	assert.Equal(t, "vapi.data.structure.field.invalid", errors[1].(*l10n.Error).ID())
}

func TestValidateStructValueInvalidForMissingField(t *testing.T) {
	var structValue = data.NewStructValue(structDefWithStringField.Name(), nil)
	var errors = structDefWithStringField.Validate(structValue)
	if len(errors) != 1 {
		t.Error("Expected 1")
	}
	if errors[0].(*l10n.Error).ID() != "vapi.data.structure.field.missing" {
		t.Error("Expected vapi.data.structure.field.missing")
	}

}

func TestValidateStructValue(t *testing.T) {
	var fields = make(map[string]data.DataValue)
	var structValue = data.NewStructValue(lib.OPERATION_INPUT, fields)
	var structDefinition = data.NewStructDefinition(lib.OPERATION_INPUT, make(map[string]data.DataDefinition))
	var errors = structDefinition.Validate(structValue)
	if len(errors) != 0 {
		t.Error("Expected no errors")
	}
}

func TestValidateStructValueOptional(t *testing.T) {
	var fields = make(map[string]data.DataValue)
	var structValue = data.NewStructValue(lib.OPERATION_INPUT, fields)
	fields["return_total_count"] = data.NewOptionalValue(nil)
	var defMap = make(map[string]data.DataDefinition)
	defMap["return_total_count"] = data.NewOptionalDefinition(data.NewBooleanDefinition())
	var structDefinition = data.NewStructDefinition(lib.OPERATION_INPUT, defMap)
	var errors = structDefinition.Validate(structValue)
	if len(errors) != 0 {
		t.Error("Expected no errors")
	}
}

func TestStructDefinitionString(t *testing.T) {
	var structDef = getStructDefWithStringField()
	assert.Equal(t, structDef.String(), "STRUCTURE")
}
