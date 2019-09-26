package bindings_test

import (
	"github.com/stretchr/testify/assert"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"testing"
)

func TestIsOneOfValidator(t *testing.T) {
	isOneOfValidator := bindings.NewIsOneOfValidator("isOneOfField",
		[]string{"val1", "val2"},
	)

	testStruct := data.NewStructValue(
		"TestStruct",
		map[string]data.DataValue{})
	msgList := isOneOfValidator.Validate(testStruct)
	// It is not the job of the validator to verify the existence of a required field
	// and isOneOfFields cannot be optional field
	assert.Equal(t, 0, len(msgList))

	testStruct = data.NewStructValue(
		"TestStruct",
		map[string]data.DataValue{
			"isOneOfField": data.NewIntegerValue(10),
		})
	msgList = isOneOfValidator.Validate(testStruct)
	// It is not the job of the validator to verify the type of
	// the isOneOfField
	assert.Equal(t, 0, len(msgList))

	testStruct = data.NewStructValue(
		"TestStruct",
		map[string]data.DataValue{
			"isOneOfField": data.NewStringValue("val1"),
			"someField":    data.NewOptionalValue(data.NewIntegerValue(100)),
		})
	msgList = isOneOfValidator.Validate(testStruct)
	assert.Equal(t, 0, len(msgList))

	testStruct = data.NewStructValue(
		"TestStruct",
		map[string]data.DataValue{
			"isOneOfField": data.NewStringValue("wrongVal"),
			"someField":    data.NewOptionalValue(data.NewIntegerValue(100)),
		})
	msgList = isOneOfValidator.Validate(testStruct)
	assert.Equal(t, 1, len(msgList))
	assert.Contains(t, "Invalid value 'wrongVal' for field 'isOneOfField' marked with IsOneOf",
		msgList[0].Error())
}

//idl-tookit will prohibit this case
func TestIsOneOfNilValues(t *testing.T) {
	isOneOfValidator := bindings.NewIsOneOfValidator("isOneOfField",
		nil,
	)
	assert.NotNil(t, isOneOfValidator)

	testStruct := data.NewStructValue(
		"TestStruct",
		map[string]data.DataValue{
			"isOneOfField": data.NewStringValue("val1"),
		})
	msgList := isOneOfValidator.Validate(testStruct)
	assert.Equal(t, 1, len(msgList))
	assert.Contains(t, "Invalid value 'val1' for field 'isOneOfField' marked with IsOneOf",
		msgList[0].Error())
}

func TestIsOneOfValidatorForStructVal(t *testing.T) {
	testIsOneOfType := getTestIsOneOfStructBindingType()
	tc := bindings.NewTypeConverter()

	testStruct := data.NewStructValue(
		"test_is_one_of_struct",
		map[string]data.DataValue{
			"is_one_of_val": data.NewStringValue("val1"),
			"someField":     data.NewOptionalValue(data.NewIntegerValue(100)),
		})
	_, err := tc.ConvertToGolang(testStruct, testIsOneOfType)
	assert.Nil(t, err)

	testStruct.SetField("is_one_of_val", data.NewStringValue("wrong_val"))
	_, err = tc.ConvertToGolang(testStruct, testIsOneOfType)
	assert.NotNil(t, err)
	assert.Equal(t, 1, len(err))
	assert.Equal(t, "Invalid value 'wrong_val' for field 'is_one_of_val' marked with IsOneOf",
		err[0].Error())
}

func TestIsOneOfValidatorForGoVal(t *testing.T) {
	testIsOneOfType := getTestIsOneOfStructBindingType()
	tc := bindings.NewTypeConverter()

	tv := TestIsOneOfStruct{
		IsOneOfVal: "val1",
	}
	_, err := tc.ConvertToVapi(tv, testIsOneOfType)
	assert.Nil(t, err)

	tv.IsOneOfVal = "wrong_val"
	_, err = tc.ConvertToVapi(tv, testIsOneOfType)
	assert.NotNil(t, err)
	assert.Equal(t, 1, len(err))
	assert.Equal(t, "Invalid value 'wrong_val' for field 'is_one_of_val' marked with IsOneOf",
		err[0].Error())
}
