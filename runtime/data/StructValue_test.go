/* **********************************************************
 * Copyright (c) 2018 VMware, Inc.  All rights reserved. -- VMware Confidential
 * **********************************************************/
package data

import (
	"testing"
)

const STRUCT_NAME = "TEST_STRUCT"

func TestGetType(t *testing.T) {
	var structValue *StructValue = NewStructValue(STRUCT_NAME, nil)
	if structValue.Type() != STRUCTURE {
		t.Error("Expected STRUCTURE but got " + structValue.Type().String())
	}
}

func TestGetName(t *testing.T) {
	var structValue *StructValue = NewStructValue(STRUCT_NAME, nil)
	if structValue.Name() != STRUCT_NAME {
		t.Error("Expected " + STRUCT_NAME + " but got " + structValue.Name())
	}
}

func TestGetString(t *testing.T) {
	var structValue = NewStructValue(STRUCT_NAME, nil)
	var sampleText = "some text"
	structValue.SetStringField("tex", sampleText)
	var value, err = structValue.String("tex")
	if err != nil {
		t.Error("Unknown error while testing get string")
	}
	if value != sampleText {
		t.Error("Expected " + STRUCT_NAME + " but got " + structValue.Name())
	}
}

//Needs IntegerValue implementation
//func TestGetStringIfFieldTypeIsWrong(t *testing.T) {
//	var structValue = NewStructValue(STRUCT_NAME)
//	structValue.SetField("text")
//}

func TestGetStringFieldMissing(t *testing.T) {
	var structValue = NewStructValue(STRUCT_NAME, nil)
	var _, error = structValue.String("nosuch")
	if error == nil {
		t.Error("Expected error")
	}
}

func TestGetOptional(t *testing.T) {
	var structValue = NewStructValue(STRUCT_NAME, nil)
	var str = "some text"
	var stringValue = NewStringValue(str)
	structValue.SetField("tex", stringValue)
	structValue.SetField("ote", NewOptionalValue(nil))
	var optionalValue, _ = structValue.Optional("ote")
	if *NewOptionalValue(nil) != *optionalValue {
		t.Error("expected equal found unequal")
	}
}

func TestOptionalErrorsIfTypeIsWrong(t *testing.T) {
	var structValue = NewStructValue(STRUCT_NAME, nil)
	var str = "some text"
	var stringValue = NewStringValue(str)
	structValue.SetField("tex", stringValue)
	structValue.SetField("ote", NewOptionalValue(nil))
	var _, error = structValue.Optional("tex")
	if error == nil {
		t.Error("expected error but got nil")
	}
}

//func TestStructValueIsEqualToStructValueWithSameFieldsAndName(t *testing.T) {
//	var structValue = NewStructValue("*", nil)
//	var str1 = "___"
//	var str2 = "___"
//
//	var stringValuea  = NewStringValue(str1)
//	structValue.SetField("t", stringValuea)
//	var structValueb = NewStructValue("*", nil)
//	var stringValueb  = NewStringValue(str2)
//	structValueb.SetField("t", stringValueb)
//	if !(structValue).Equals(structValueb) {
//		t.Error("Expected two structs to be equal but they are not")
//	}
//
//}

//func TestStructValueIsNotEqualToStructValueWithDifferentFieldValue(t *testing.T){
//	var structValue = NewStructValue("*", nil)
//	var str1 = "A"
//	var str2 = "B"
//
//	var stringValuea  = NewStringValue(str1)
//	structValue.SetField("t", stringValuea)
//	var structValueb = NewStructValue("*", nil)
//	var stringValueb = NewStringValue(str2)
//	structValueb.SetField("t", stringValueb)
//	if (structValue).Equals(structValueb) {
//		t.Error("Expected two structs to be not equal but they are equal")
//	}
//}

//func TestStructValueIsNotEqualToStructValueWithDifferentFieldName(t *testing.T) {
//	var structValue = NewStructValue("*", nil)
//	var str1 = "A"
//	var str2 = "A"
//
//	var stringValuea  = NewStringValue(str1)
//	structValue.SetField("X", stringValuea)
//	var structValueb = NewStructValue("*", nil)
//	var stringValueb  = NewStringValue(str2)
//	structValueb.SetField("Y", stringValueb)
//	if (structValue).Equals(structValueb) {
//		t.Error("Expected two structs to be not equal but they are equal")
//	}
//}

//func TestStructValueIsNotEqualToStructValueWithDifferentSize(t *testing.T) {
//	var structValue = NewStructValue("*", nil)
//	var str2 = "A"
//	var structValueb = NewStructValue("*", nil)
//	var stringValueb  = NewStringValue(str2)
//	structValueb.SetField("Y", stringValueb)
//	if (structValue).Equals(structValueb) {
//		t.Error("Expected two structs to be not equal but they are equal")
//	}
//}

//func TestStructValueIsNotEqualToStructValueWithDifferentName(t *testing.T) {
//	var structValue = NewStructValue("y", nil)
//	var structValueB = NewStructValue("X", nil)
//	if (structValue).Equals(structValueB) {
//		t.Error("Expected two structs to be not equal but they are equal")
//	}
//}

//func TestStructValueIsNotEqualToStringValue(t *testing.T){
//	var structValue = NewStructValue("y", nil)
//	var str = "y"
//	var stringValue = NewStringValue(str)
//	if structValue.Equals(stringValue) {
//		t.Error("struct value and string value cannot be equal")
//	}
//
//}

//func TestEqualToItself(t *testing.T){
//	var structValue = NewStructValue("y", nil)
//	if !structValue.Equals(structValue) {
//		t.Error("Expected equal but found unequal")
//	}
//}

//func TestEqualToNil(t *testing.T) {
//	var structValue = NewStructValue("y", nil)
//	if structValue.Equals(nil){
//		t.Error("Expected not equal but found equal")
//	}
//}
