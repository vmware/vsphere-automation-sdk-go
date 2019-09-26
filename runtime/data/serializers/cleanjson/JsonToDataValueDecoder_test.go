package cleanjson_test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data/serializers/cleanjson"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestJsonToDataValueDecoder_Decode(t *testing.T) {
	var jsonToDataValueDecoder = cleanjson.NewJsonToDataValueDecoder()
	jsonFile, err := os.Open("test/sample.json")
	defer jsonFile.Close()
	assert.Nil(t, err)
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var jsonMap = map[string]interface{}{}
	d := json.NewDecoder(strings.NewReader(string(byteValue)))
	d.UseNumber()
	if err := d.Decode(&jsonMap); err != nil {
		t.Error(err)
		t.Fail()
	}

	actualVal, err := jsonToDataValueDecoder.Decode(jsonMap)
	structVal := actualVal.(*data.StructValue)

	strkeyVal, _ := structVal.Field("strkey")
	assert.Equal(t, strkeyVal, data.NewStringValue("val1"))

	strunicodekeyVal, _ := structVal.Field("strunicodekey")
	assert.Equal(t, strunicodekeyVal, data.NewStringValue("val©"))

	intkeyVal, _ := structVal.Field("intkey")
	assert.Equal(t, intkeyVal, data.NewIntegerValue(10))

	boolkeyVal, _ := structVal.Field("boolkey")
	assert.Equal(t, boolkeyVal, data.NewBooleanValue(true))

	floatkeyVal, _ := structVal.Field("floatkey")
	assert.Equal(t, floatkeyVal, data.NewDoubleValue(10.10))

	listvalemptykeyVal, _ := structVal.Field("listvalemptykey")
	assert.Equal(t, listvalemptykeyVal, data.NewListValue())

	listvalkeyVal, _ := structVal.Field("listvalkey")
	listval1 := data.NewListValue()
	listval1.Add(data.NewStringValue("val1"))
	listval1.Add(data.NewStringValue("val2"))
	assert.Equal(t, listvalkeyVal, listval1)

	nonekeyVal, _ := structVal.Field("nonekey")
	assert.Equal(t, nonekeyVal, data.NewOptionalValue(nil))

	dictkeyVal, _ := structVal.Field("dictkey")
	dictkeyStructVal := dictkeyVal.(*data.StructValue)
	field1Val, _ := dictkeyStructVal.Field("field1")
	assert.Equal(t, field1Val, data.NewIntegerValue(10))

	field2Val, _ := dictkeyStructVal.Field("field2")
	assert.Equal(t, field2Val, data.NewStringValue("twenty"))

	field3Val, _ := dictkeyStructVal.Field("field3")
	assert.Equal(t, field3Val, data.NewListValue())

	field4Val, _ := dictkeyStructVal.Field("field4")
	listval1 = data.NewListValue()
	listval1.Add(data.NewIntegerValue(10))
	listval1.Add(data.NewIntegerValue(20))
	assert.Equal(t, field4Val, listval1)

}

type testSruct struct {
	test int
}

func TestInvalidValueDecode(t *testing.T) {
	var jsonToDataValueDecoder = cleanjson.NewJsonToDataValueDecoder()
	var testType testSruct
	_, err := jsonToDataValueDecoder.Decode(testType)
	assert.NotNil(t, err)
	assert.Equal(t, "Could not determine appropriate DataValue for json type 'cleanjson_test.testSruct'", err.Error())
}

func TestInvalidNumberDecode(t *testing.T) {
	var jsonToDataValueDecoder = cleanjson.NewJsonToDataValueDecoder()
	var test json.Number = "123abc"
	_, err := jsonToDataValueDecoder.Decode(test)
	assert.NotNil(t, err)
	assert.Equal(t, "json.Number neither Int64 nor Float64", err.Error())
}
