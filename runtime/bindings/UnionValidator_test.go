package bindings_test

import (
	"github.com/stretchr/testify/assert"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"testing"
)

func TestUnionValidatorPositive(t *testing.T) {
	unionValidator := bindings.NewUnionValidator("test_enum",
		map[string][]bindings.FieldData{
			"LONG":   []bindings.FieldData{bindings.NewFieldData("long_val", true)},
			"STRING": []bindings.FieldData{bindings.NewFieldData("string_val", true)},
			"NONE":   []bindings.FieldData{},
		},
	)

	testStruct := data.NewStructValue(
		"TestStruct",
		map[string]data.DataValue{
			"test_enum": data.NewStringValue("LONG"),
			"long_val":  data.NewOptionalValue(data.NewIntegerValue(100)),
		})
	msgList := unionValidator.Validate(testStruct)
	assert.Equal(t, 0, len(msgList))

	testStruct = data.NewStructValue(
		"TestStruct",
		map[string]data.DataValue{
			"test_enum":  data.NewStringValue("STRING"),
			"string_val": data.NewStringValue("someVal"),
		})
	msgList = unionValidator.Validate(testStruct)
	assert.Equal(t, 0, len(msgList))

	testStruct = data.NewStructValue(
		"TestStruct",
		map[string]data.DataValue{
			"test_enum": data.NewStringValue("NONE"),
		})
	msgList = unionValidator.Validate(testStruct)
	assert.Equal(t, 0, len(msgList))

	testStruct = data.NewStructValue(
		"TestStruct",
		map[string]data.DataValue{
			"test_enum": data.NewStringValue("NONE"),
			"long_val":  data.NewOptionalValue(nil),
		})
	msgList = unionValidator.Validate(testStruct)
	assert.Equal(t, 0, len(msgList))

	//Optional tag
	testStruct = data.NewStructValue(
		"TestStruct",
		map[string]data.DataValue{
			"test_enum":  data.NewOptionalValue(data.NewStringValue("STRING")),
			"string_val": data.NewStringValue("someVal"),
		})
	msgList = unionValidator.Validate(testStruct)
	assert.Equal(t, 0, len(msgList))

	testStruct = data.NewStructValue(
		"TestStruct",
		map[string]data.DataValue{
			"test_enum": data.NewOptionalValue(data.NewStringValue("NONE")),
		})
	msgList = unionValidator.Validate(testStruct)
	assert.Equal(t, 0, len(msgList))

	//Unset tag
	testStruct = data.NewStructValue(
		"TestStruct",
		map[string]data.DataValue{
			"test_enum": data.NewOptionalValue(nil),
		})
	msgList = unionValidator.Validate(testStruct)
	assert.Equal(t, 0, len(msgList))
}

func TestUnionValidatorNegative(t *testing.T) {
	unionValidator := bindings.NewUnionValidator("test_enum",
		map[string][]bindings.FieldData{
			"LONG":   []bindings.FieldData{bindings.NewFieldData("long_val", true)},
			"STRING": []bindings.FieldData{bindings.NewFieldData("string_val", true)},
			"NONE":   []bindings.FieldData{},
		},
	)

	testStruct := data.NewStructValue(
		"TestStruct",
		map[string]data.DataValue{
			"test_enum":  data.NewStringValue("LONG"),
			"string_val": data.NewOptionalValue(data.NewStringValue("someVal")),
		})
	msgList := unionValidator.Validate(testStruct)
	assert.Equal(t, 1, len(msgList))
	assert.Contains(t, msgList[0].Error(), "missing a required field 'long_val' for this case")

	testStruct = data.NewStructValue(
		"TestStruct",
		map[string]data.DataValue{
			"test_enum":  data.NewStringValue("None"),
			"string_val": data.NewOptionalValue(data.NewStringValue("someVal")),
		})
	msgList = unionValidator.Validate(testStruct)
	assert.Equal(t, 1, len(msgList))
	assert.Contains(t, msgList[0].Error(), "field 'string_val' not required for this case")

	//Optional union tag
	testStruct = data.NewStructValue(
		"TestStruct",
		map[string]data.DataValue{
			"test_enum":  data.NewOptionalValue(data.NewStringValue("LONG")),
			"string_val": data.NewOptionalValue(data.NewStringValue("someVal")),
		})
	msgList = unionValidator.Validate(testStruct)
	assert.Equal(t, 1, len(msgList))
	assert.Contains(t, msgList[0].Error(), "missing a required field 'long_val' for this case")

	testStruct = data.NewStructValue(
		"TestStruct",
		map[string]data.DataValue{
			"test_enum":  data.NewOptionalValue(data.NewStringValue("None")),
			"string_val": data.NewOptionalValue(data.NewStringValue("someVal")),
		})
	msgList = unionValidator.Validate(testStruct)
	assert.Equal(t, 1, len(msgList))
	assert.Contains(t, msgList[0].Error(), "field 'string_val' not required for this case")

	//Unset tag
	testStruct = data.NewStructValue(
		"TestStruct",
		map[string]data.DataValue{
			"test_enum":  data.NewOptionalValue(nil),
			"string_val": data.NewStringValue("someVal"),
		})
	msgList = unionValidator.Validate(testStruct)
	assert.Equal(t, 1, len(msgList))
	assert.Contains(t, msgList[0].Error(), "field 'string_val' not required for this case")
}

func TestUnionValidatorPositiveOptionalCase(t *testing.T) {
	unionValidator := bindings.NewUnionValidator("test_enum",
		map[string][]bindings.FieldData{
			"LONG": []bindings.FieldData{
				bindings.NewFieldData("long_val", true),
				bindings.NewFieldData("opt_long_val", false)},
			"STRING": []bindings.FieldData{
				bindings.NewFieldData("string_val", true)},
			"NONE": []bindings.FieldData{},
		},
	)

	testStruct := data.NewStructValue(
		"TestStruct",
		map[string]data.DataValue{
			"test_enum": data.NewStringValue("LONG"),
			"long_val":  data.NewIntegerValue(100),
		})
	msgList := unionValidator.Validate(testStruct)
	assert.Equal(t, 0, len(msgList))

	testStruct = data.NewStructValue(
		"TestStruct",
		map[string]data.DataValue{
			"test_enum":    data.NewStringValue("LONG"),
			"long_val":     data.NewIntegerValue(100),
			"opt_long_val": data.NewOptionalValue(data.NewIntegerValue(100)),
		})
	msgList = unionValidator.Validate(testStruct)
	assert.Equal(t, 0, len(msgList))

	testStruct = data.NewStructValue(
		"TestStruct",
		map[string]data.DataValue{
			"test_enum": data.NewStringValue("NONE"),
		})
	msgList = unionValidator.Validate(testStruct)
	assert.Equal(t, 0, len(msgList))

	testStruct = data.NewStructValue(
		"TestStruct",
		map[string]data.DataValue{
			"test_enum":    data.NewStringValue("NONE"),
			"opt_long_val": data.NewOptionalValue(nil),
		})
	msgList = unionValidator.Validate(testStruct)
	assert.Equal(t, 0, len(msgList))
}

func TestUnionValidatorNegativeOptionalCase(t *testing.T) {
	unionValidator := bindings.NewUnionValidator("test_enum",
		map[string][]bindings.FieldData{
			"LONG": []bindings.FieldData{
				bindings.NewFieldData("long_val", true),
				bindings.NewFieldData("opt_long_val", false)},
			"STRING": []bindings.FieldData{
				bindings.NewFieldData("string_val", true)},
			"NONE": []bindings.FieldData{},
		},
	)

	testStruct := data.NewStructValue(
		"TestStruct",
		map[string]data.DataValue{
			"test_enum":    data.NewStringValue("LONG"),
			"opt_long_val": data.NewOptionalValue(data.NewIntegerValue(100)),
		})
	msgList := unionValidator.Validate(testStruct)
	assert.Equal(t, 1, len(msgList))
	assert.Contains(t, msgList[0].Error(), "missing a required field 'long_val' for this case")

	testStruct = data.NewStructValue(
		"TestStruct",
		map[string]data.DataValue{
			"test_enum":    data.NewStringValue("STRING"),
			"string_val":   data.NewOptionalValue(data.NewStringValue("someVal")),
			"opt_long_val": data.NewOptionalValue(data.NewIntegerValue(100)),
		})
	msgList = unionValidator.Validate(testStruct)
	assert.Equal(t, 1, len(msgList))
	assert.Contains(t, msgList[0].Error(), "field 'opt_long_val' not required for this case")
}

func TestUnionValidatorForStructVal(t *testing.T) {
	testUnionType := getTestUnionStructBindingType()
	tc := bindings.NewTypeConverter()

	testStruct := data.NewStructValue(
		"TestStruct",
		map[string]data.DataValue{
			"test_enum": data.NewStringValue("LONG"),
			"long_val":  data.NewOptionalValue(data.NewIntegerValue(100)),
		})
	_, err := tc.ConvertToGolang(testStruct, testUnionType)
	assert.Nil(t, err)

	testStruct.SetField("test_enum", data.NewStringValue("STRING"))
	_, msgList := tc.ConvertToGolang(testStruct, testUnionType)
	assert.NotNil(t, msgList)
	assert.Equal(t, 1, len(msgList))
	assert.Equal(t, "Structure 'TestStruct' has a union that is missing a required field 'string_val' for this case",
		msgList[0].Error())
}

func TestUnionValidatorForGoVal(t *testing.T) {
	l := int64(100)
	testUnionType := getTestUnionStructBindingType()
	tc := bindings.NewTypeConverter()

	tv := TestUnionStruct{
		TestEnum: "LONG",
		LongVal:  &l,
	}
	_, err := tc.ConvertToVapi(tv, testUnionType)
	assert.Nil(t, err)

	tv.TestEnum = "STRING"
	_, msgList := tc.ConvertToVapi(tv, testUnionType)
	assert.NotNil(t, msgList)
	assert.Equal(t, 1, len(msgList))
	assert.Equal(t, "Structure 'TestUnionStruct' has a union that is missing a required field 'string_val' for this case",
		msgList[0].Error())
}
