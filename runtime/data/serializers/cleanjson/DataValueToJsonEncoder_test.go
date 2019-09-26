package cleanjson_test

/* **********************************************************
 * Copyright 2019 VMware, Inc.  All rights reserved. -- VMware Confidential
 * **********************************************************/
import (
	"encoding/base64"
	"github.com/stretchr/testify/assert"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data/serializers/cleanjson"
	"strconv"
	"testing"
)

var dataValueToJsonEncoder = cleanjson.NewDataValueToJsonEncoder()

func TestStructValue(t *testing.T) {
	inputVal := data.NewStructValue("name", map[string]data.DataValue{
		"val1": data.NewStringValue("val1"),
		"val2": data.NewStringValue("val2"),
	})
	actualVal, err := dataValueToJsonEncoder.Encode(inputVal)
	assert.Nil(t, err)
	assert.Equal(t, actualVal, "{\"val1\":\"val1\",\"val2\":\"val2\"}")
}

func TestStructValueWithOptionalSet(t *testing.T) {
	inputVal := data.NewStructValue("name", map[string]data.DataValue{
		"val1": data.NewStringValue("val1"),
		"val2": data.NewOptionalValue(data.NewStringValue("val2"))})

	actualVal, err := dataValueToJsonEncoder.Encode(inputVal)
	assert.Nil(t, err)
	assert.Equal(t, actualVal, "{\"val1\":\"val1\",\"val2\":\"val2\"}")

}

func TestStructValueWithOptionalUnSet(t *testing.T) {
	inputVal := data.NewStructValue("name", map[string]data.DataValue{
		"val1": data.NewStringValue("val1"),
		"val2": data.NewOptionalValue(nil)})
	actualVal, err := dataValueToJsonEncoder.Encode(inputVal)
	assert.Nil(t, err)
	assert.Equal(t, actualVal, "{\"val1\":\"val1\"}")

}

func TestErrorValue(t *testing.T) {
	inputVal := data.NewErrorValue("name", nil)
	inputVal.SetField("val1", data.NewStringValue("val1"))
	inputVal.SetField("val2", data.NewStringValue("val2"))
	actualVal, err := dataValueToJsonEncoder.Encode(inputVal)
	assert.Equal(t, actualVal, "{\"val1\":\"val1\",\"val2\":\"val2\"}")
	assert.Nil(t, err)
}

func TestErrorValueWithEmptyOptional(t *testing.T) {
	inputVal := data.NewErrorValue("name", nil)
	inputVal.SetField("val1", data.NewStringValue("val1"))
	inputVal.SetField("val2", data.NewOptionalValue(nil))
	actualVal, err := dataValueToJsonEncoder.Encode(inputVal)
	assert.Equal(t, actualVal, "{\"val1\":\"val1\"}")
	assert.Nil(t, err)
}

func TestListValue(t *testing.T) {
	inputVal := data.NewListValue()
	inputVal.Add(nil)
	inputVal.Add(data.NewStringValue("val1"))
	inputVal.Add(data.NewStringValue("val2"))
	actualVal, err := dataValueToJsonEncoder.Encode(inputVal)
	assert.Nil(t, err)
	assert.Equal(t, actualVal, "[\"val1\",\"val2\"]")
}

func TestListValueWithEmptyOptional(t *testing.T) {
	inputVal := data.NewListValue()
	inputVal.Add(data.NewStringValue("val1"))
	inputVal.Add(data.NewOptionalValue(nil))
	actualVal, err := dataValueToJsonEncoder.Encode(inputVal)
	assert.Nil(t, err)
	assert.Equal(t, actualVal, "[\"val1\"]")
}

func TestOptionalValue(t *testing.T) {
	// when optional value is set
	inputVal := data.NewOptionalValue(data.NewStringValue("val1"))
	actualVal, err := dataValueToJsonEncoder.Encode(inputVal)
	actualVal, _ = strconv.Unquote(actualVal)
	expectedVal := "val1"
	assert.Nil(t, err)
	assert.Equal(t, actualVal, expectedVal)

	// when optional value is not set
	inputVal = data.NewOptionalValue(nil)
	actualVal, err = dataValueToJsonEncoder.Encode(inputVal)
	expectedVal = "null"
	assert.Nil(t, err)
	assert.Equal(t, expectedVal, actualVal)
}

func TestStringValue(t *testing.T) {
	inputVal := data.NewStringValue("val1")
	actualVal, err := dataValueToJsonEncoder.Encode(inputVal)
	actualVal, _ = strconv.Unquote(actualVal)
	expectedVal := "val1"
	assert.Nil(t, err)
	assert.Equal(t, expectedVal, actualVal)
}

func TestIntegerValue(t *testing.T) {
	inputVal := data.NewIntegerValue(10)
	actualVal, err := dataValueToJsonEncoder.Encode(inputVal)
	expectedVal := "10"
	assert.Nil(t, err)
	assert.Equal(t, expectedVal, actualVal)
}

func TestBooleanValue(t *testing.T) {
	inputVal := data.NewBooleanValue(true)
	actualVal, err := dataValueToJsonEncoder.Encode(inputVal)
	expectedVal := "true"
	assert.Nil(t, err)
	assert.Equal(t, expectedVal, actualVal)
}

func TestVoidValue(t *testing.T) {
	inputVal := data.NewVoidValue()
	actualVal, err := dataValueToJsonEncoder.Encode(inputVal)
	expectedVal := "null"
	assert.Nil(t, err)
	assert.Equal(t, expectedVal, actualVal)
}

func TestBinaryValue(t *testing.T) {
	inputVal := data.NewBlobValue([]byte("asdlkfu"))
	actualVal, err := dataValueToJsonEncoder.Encode(inputVal)
	expectedVal := base64.StdEncoding.EncodeToString([]byte("asdlkfu"))
	actualVal, _ = strconv.Unquote(actualVal)
	assert.Nil(t, err)
	assert.Equal(t, expectedVal, actualVal)
}

func TestBinaryValue2(t *testing.T) {
	inputVal := data.NewBlobValue([]byte("0\x82\x04\x00"))
	actualVal, err := dataValueToJsonEncoder.Encode(inputVal)
	expectedVal := base64.StdEncoding.EncodeToString([]byte("0\x82\x04\x00"))
	actualVal, _ = strconv.Unquote(actualVal)
	assert.Nil(t, err)
	assert.Equal(t, expectedVal, actualVal)
}

func TestSecretValue(t *testing.T) {
	inputVal := data.NewSecretValue("password")
	actualVal, err := dataValueToJsonEncoder.Encode(inputVal)
	actualVal, _ = strconv.Unquote(actualVal)
	expectedVal := "password"
	assert.Nil(t, err)
	assert.Equal(t, expectedVal, actualVal)
}

func TestDoubleValue(t *testing.T) {
	inputVal := data.NewDoubleValue(10.10)
	actualVal, err := dataValueToJsonEncoder.Encode(inputVal)
	expectedVal := "10.1"
	assert.Nil(t, err)
	assert.Equal(t, expectedVal, actualVal)
}

func TestUnknownDataValueType(t *testing.T) {
	inputVal := 1
	_, err := dataValueToJsonEncoder.Encode(inputVal)
	assert.NotNil(t, err)
	assert.Equal(t, "Error serializing DataValue of type 'int'", err.Error())
}
