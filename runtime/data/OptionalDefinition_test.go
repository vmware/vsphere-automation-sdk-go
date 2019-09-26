/* **********************************************************
 * Copyright 2018 VMware, Inc.  All rights reserved. -- VMware Confidential
 * **********************************************************/
package data_test

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"testing"
)

var NAME_STRUCT_WITH_ONE_OPTIONAL_FIELD = "struct with one optional field"
var NAME_STRUCT_WITH_ONE_FIELD = "struct with one field"
var OPT_DEF_NAME_FIELD_1 = "one"

func getOptionalStructureWithOptionalString() data.OptionalDefinition {
	var fields = make(map[string]data.DataDefinition, 1)
	var optDef = data.NewOptionalDefinition(data.NewStringDefinition())
	fields[OPT_DEF_NAME_FIELD_1] = optDef
	var structDef = data.NewStructDefinition(NAME_STRUCT_WITH_ONE_OPTIONAL_FIELD, fields)
	var result = data.NewOptionalDefinition(structDef)
	return result
}

func getOptionalStructureWithString() data.OptionalDefinition {
	var fields = make(map[string]data.DataDefinition, 1)
	fields[OPT_DEF_NAME_FIELD_1] = data.NewStringDefinition()
	var structDef = data.NewStructDefinition(NAME_STRUCT_WITH_ONE_FIELD, fields)
	var result = data.NewOptionalDefinition(structDef)
	return result
}

//func TestNullElementTypeIsForbidden(t *testing.T)  {
//	var _, error = NewOptionalDefinition(nil)
//	if error == nil {
//		t.Error("Expected error. Did not get an error. optionalElementType cannot be null")
//	}
//}

func TestGetTypeReturnsOptional(t *testing.T) {
	var optDef = data.NewOptionalDefinition(data.NewStringDefinition())
	if optDef.Type() != data.OPTIONAL {
		t.Error("Expected optional but got " + optDef.Type().String())
	}
}

func TestGetElementTypeReturnsCorrectDefinition(t *testing.T) {
	var optDef = data.NewOptionalDefinition(data.NewStringDefinition())
	if optDef.Type() != data.OPTIONAL {
		t.Error("Expected OPTIONAL type but got " + optDef.Type().String())
	}
}

func TestCompleteValuePopulatesUnSetOptionalValue(t *testing.T) {
	var structValue = data.NewStructValue(NAME_STRUCT_WITH_ONE_OPTIONAL_FIELD, nil)

	var value = data.NewOptionalValue(structValue)
	getOptionalStructureWithOptionalString().CompleteValue(value)
	var structValue2 = (value.Value()).(*data.StructValue)
	var optionalValue, _ = structValue2.Optional(OPT_DEF_NAME_FIELD_1)
	if optionalValue.IsSet() == true {
		t.Error("Expected optionalValue to be unset")
	}
}

func TestCompleteValueIgnoresSetOptionalValue(t *testing.T) {
	var structValue = data.NewStructValue(NAME_STRUCT_WITH_ONE_OPTIONAL_FIELD, nil)
	var str = ""
	var stringValue = data.NewStringValue(str)
	var oValue = data.NewOptionalValue(stringValue)

	structValue.SetField(OPT_DEF_NAME_FIELD_1, oValue)
	var optValue = data.NewOptionalValue(structValue)
	getOptionalStructureWithOptionalString().CompleteValue(optValue)
	var structValue2 = (optValue.Value()).(*data.StructValue)
	var optionalValue, _ = structValue2.Optional(OPT_DEF_NAME_FIELD_1)
	if optionalValue.IsSet() != true {
		t.Error("Expected optionalValue to be set")
	}
}

func TestCompleteValueIgnoresNil(t *testing.T) {
	getOptionalStructureWithOptionalString().CompleteValue(nil)
}

func TestCompleteValueIgnoresSetOptionalField(t *testing.T) {
	var structValue = data.NewStructValue(NAME_STRUCT_WITH_ONE_OPTIONAL_FIELD, nil)
	var str = ""
	var stringValue = data.NewStringValue(str)
	structValue.SetField(OPT_DEF_NAME_FIELD_1, data.NewOptionalValue(stringValue))
	var value = data.NewOptionalValue(structValue)
	var opStructOpString = getOptionalStructureWithOptionalString()
	if opStructOpString.ValidInstanceOf(value) == false {
		t.Error("Expected true but got false instead")
	}
	opStructOpString.CompleteValue(value)
	if opStructOpString.ValidInstanceOf(value) == false {
		t.Error("Expected true but got false instead")
	}

}

func TestCompleteValueIgnoresUnSetOptionalField(t *testing.T) {
	var structValue = data.NewStructValue(NAME_STRUCT_WITH_ONE_OPTIONAL_FIELD, nil)
	structValue.SetField(OPT_DEF_NAME_FIELD_1, data.NewOptionalValue(nil))
	var opStructOpString = getOptionalStructureWithOptionalString()
	var value = data.NewOptionalValue(structValue)
	if opStructOpString.ValidInstanceOf(value) == false {
		t.Error("Expected true but got false instead")
	}
	opStructOpString.CompleteValue(value)
	if opStructOpString.ValidInstanceOf(value) == false {
		t.Error("Expected true but got false instead")
	}
}

func TestCompleteValueIgnoresRequiredField(t *testing.T) {
	var structValue = data.NewStructValue(NAME_STRUCT_WITH_ONE_FIELD, nil)
	var str = ""
	var stringValue = data.NewStringValue(str)
	structValue.SetField(OPT_DEF_NAME_FIELD_1, stringValue)
	var value = data.NewOptionalValue(structValue)
	//required field is present
	var optionalStructureWithStringField = getOptionalStructureWithString()
	if optionalStructureWithStringField.ValidInstanceOf(value) == false {
		t.Error("expected true got false")
	}
	optionalStructureWithStringField.CompleteValue(value)
	if optionalStructureWithStringField.ValidInstanceOf(value) == false {
		t.Error("expected true got false")
	}
}

func TestCompleteValueIgnoresMissingRequiredField(t *testing.T) {
	var structValue = data.NewStructValue(NAME_STRUCT_WITH_ONE_FIELD, nil)
	var value = data.NewOptionalValue(structValue)
	// required field is missing
	var optionalStructureWithStringField = getOptionalStructureWithString()
	if optionalStructureWithStringField.ValidInstanceOf(value) == true {
		t.Error("expected false got true")
	}
	optionalStructureWithStringField.CompleteValue(value)
	if optionalStructureWithStringField.ValidInstanceOf(value) == true {
		t.Error("expected false got true")
	}
}

func TestCompleteValueIgnoresWrongValueType(t *testing.T) {
	var str = ""
	var stringValue = data.NewStringValue(str)
	var optionalStructureWithStringField = getOptionalStructureWithString()
	if optionalStructureWithStringField.ValidInstanceOf(stringValue) == true {
		t.Error("expected false got true")
	}
	optionalStructureWithStringField.CompleteValue(stringValue)
	if optionalStructureWithStringField.ValidInstanceOf(stringValue) == true {
		t.Error("expected false got true")
	}

}
