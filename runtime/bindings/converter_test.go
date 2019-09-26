package bindings_test

import (
	"encoding/base64"
	"fmt"
	"github.com/stretchr/testify/assert"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/l10n"
	"net/url"
	"reflect"
	"testing"
	"time"
)

type LocalizableMessage struct {
	Id             string
	DefaultMessage string
	Args           []string
}

func LocalizableMessageType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = bindings.NewStringType()
	fieldNameMap["id"] = "Id"

	fields["default_message"] = bindings.NewStringType()
	fieldNameMap["default_message"] = "DefaultMessage"

	fields["args"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["args"] = "Args"

	return bindings.NewStructType("com.vmware.vapi.std.localizable_message",
		fields, reflect.TypeOf(LocalizableMessage{}), fieldNameMap, nil)
}

func TimedOutType() bindings.ErrorType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["messages"] = bindings.NewListType(LocalizableMessageType(), reflect.TypeOf([]LocalizableMessage{}))
	fieldNameMap["messages"] = "Messages"
	fields["data"] = bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
	fieldNameMap["data"] = "Data"
	return bindings.NewErrorType("com.vmware.vapi.std.errors.timed_out",
		fields, reflect.TypeOf(TimedOut{}), fieldNameMap)
}

type AnyErrorProperties struct {
	ErrorField *data.ErrorValue
}

var timeOutDef = bindings.CreateStdErrorDefinition("com.vmware.vapi.std.errors.timed_out")
var timeOutMessage = l10n.NewRuntimeError("vapi.connection", map[string]string{"host": "localhost"})
var timeOutValue = bindings.CreateErrorValueFromMessages(timeOutDef, []error{timeOutMessage})

type TimedOut struct {
	Messages []LocalizableMessage
	Data     *data.StructValue
}

func getBindingTypeAnyErrorProperties() bindings.StructType {
	var fields = make(map[string]bindings.BindingType)
	fields["ErrorField"] = bindings.NewAnyErrorType()

	canonicalFieldMap := make(map[string]string)
	canonicalFieldMap["ErrorField"] = "ErrorField"

	var structType = bindings.NewStructType("AnyErrorProperties", fields,
		reflect.TypeOf(AnyErrorProperties{}), canonicalFieldMap, nil)
	return structType
}

func getProperties() Properties {
	item1 := NewItem("hello")
	item2 := NewItem("world")
	itemlist := []Item{item1, item2}
	item3 := NewItem("golang")
	item4 := NewItem("support")
	itemlist2 := []*Item{&item3, &item4}
	itemMap := make(map[string]Item)
	itemMap["Itemkey1"] = item1
	itemMap["Itemkey2"] = item2
	itemOptionalValue := make(map[string]*Item)
	itemOptionalValue["Itemkey1"] = &item1
	itemOptionalValue["Itemkey2"] = &item2
	x := int64(10)
	ma := make(map[string]int64)
	ma["hello"] = 1
	ma["world"] = 2
	var optionalItemList []Item = nil
	return Properties{int64(10),
		&x,
		item1,
		&item1,
		itemlist,
		itemlist2,
		ma,
		itemMap,
		itemOptionalValue,
		optionalItemList}

}

func TestIntegerConversion(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	var bindingType = bindings.NewIntegerType()
	dataVal, err := typeConverter.ConvertToVapi(int64(10), bindingType)
	assert.Nil(t, err)
	assert.Equal(t, dataVal.(*data.IntegerValue).Value(), int64(10))
	goval, err := typeConverter.ConvertToGolang(dataVal, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, int64(10), goval)
}

func TestIntegerPtrConversion(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	var bindingType = bindings.NewOptionalType(bindings.NewIntegerType())
	var x int64 = 10
	dataVal, err := typeConverter.ConvertToVapi(&x, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, dataVal.(*data.OptionalValue).Value().(*data.IntegerValue).Value(), int64(10))
	goval, err := typeConverter.ConvertToGolang(dataVal, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, x, *(goval.(*int64)))
}

func TestIntegerNilConversion(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	var bindingType = bindings.NewOptionalType(bindings.NewIntegerType())
	var intNil *int64 = nil
	dataValNil, err := typeConverter.ConvertToVapi(intNil, bindingType)
	assert.Nil(t, err)
	goValNil, err := typeConverter.ConvertToGolang(dataValNil, bindingType)
	assert.Equal(t, intNil, goValNil)
}

func TestDoubleConversion(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	var bindingType = bindings.NewDoubleType()
	dataVal, err := typeConverter.ConvertToVapi(float64(3.14), bindingType)
	assert.Nil(t, err)
	assert.Equal(t, dataVal.(*data.DoubleValue).Value(), float64(3.14))
	goval, err := typeConverter.ConvertToGolang(dataVal, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, float64(3.14), goval)
}

func TestDoubleOptionalConversion(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	var bindingType = bindings.NewOptionalType(bindings.NewDoubleType())
	var x = 3.14
	dataVal, err := typeConverter.ConvertToVapi(&x, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, dataVal.(*data.OptionalValue).Value().(*data.DoubleValue).Value(), float64(3.14))
	goval, err := typeConverter.ConvertToGolang(dataVal, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, float64(3.14), *(goval.(*float64)))
}

func TestNilDoubleConversion(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	var bindingType = bindings.NewOptionalType(bindings.NewDoubleType())
	var y *float64 = nil
	dataVal, err := typeConverter.ConvertToVapi(y, bindingType)
	assert.Nil(t, err)
	assert.False(t, dataVal.(*data.OptionalValue).IsSet())
	goval, err := typeConverter.ConvertToGolang(dataVal, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, goval, y)
}

func TestBlobConversion(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	var bindingType = bindings.NewBlobType()
	hw := []byte("helloworld")
	dataVal, err := typeConverter.ConvertToVapi(hw, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, dataVal.(*data.BlobValue).Value(), hw)
	goval, err := typeConverter.ConvertToGolang(dataVal, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, hw, goval)
}

func TestBlobPtrConversion(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	var bindingType = bindings.NewOptionalType(bindings.NewBlobType())
	hw := []byte("helloworld")
	dataVal, err := typeConverter.ConvertToVapi(&hw, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, dataVal.(*data.OptionalValue).Value().(*data.BlobValue).Value(), hw)
	goval, err := typeConverter.ConvertToGolang(dataVal, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, hw, (goval.([]byte)))
}

func TestBlobPtrNilConversion(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	var bindingType = bindings.NewOptionalType(bindings.NewBlobType())
	var x []byte = nil
	dataVal, err := typeConverter.ConvertToVapi(&x, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, dataVal.(*data.OptionalValue).Value().(*data.BlobValue).Value(), x)
	goval, err := typeConverter.ConvertToGolang(dataVal, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, x, goval)
}

// Special case implemented for: VAPI-1648
// See https://reviewboard.eng.vmware.com/r/1469420/
func TestBlobInMapConversion(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	var bindingType = bindings.NewMapType(bindings.NewStringType(), bindings.NewBlobType(), reflect.TypeOf(make(map[string][]byte)))
	input := map[string][]byte{"a": []byte("helloworld")}
	dataVal, err := typeConverter.ConvertToVapi(input, bindingType)
	assert.Nil(t, err)
	goval, err := typeConverter.ConvertToGolang(dataVal, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, input, goval)
	assert.Contains(t, goval, "a")
}

func TestStringConversion(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	var bindingType = bindings.NewStringType()
	dataVal, err := typeConverter.ConvertToVapi("test", bindingType)
	assert.Nil(t, err)
	assert.Equal(t, dataVal.(*data.StringValue).Value(), "test")
	goval, err := typeConverter.ConvertToGolang(dataVal, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, goval, "test")
}

func TestStringPtrConversion(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	var bindingType = bindings.NewOptionalType(bindings.NewStringType())
	x := "hello"
	dataVal, err := typeConverter.ConvertToVapi(&x, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, dataVal.(*data.OptionalValue).Value().(*data.StringValue).Value(), "hello")
	goval, err := typeConverter.ConvertToGolang(dataVal, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, "hello", *(goval.(*string)))
}

func TestStringNilConversion(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	var bindingType = bindings.NewOptionalType(bindings.NewStringType())
	var x *string = nil
	dataVal, err := typeConverter.ConvertToVapi(x, bindingType)
	assert.Nil(t, err)
	goval, err := typeConverter.ConvertToGolang(dataVal, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, x, goval)
}

func TestSecretConversion(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	var bindingType = bindings.NewSecretType()
	dataVal, err := typeConverter.ConvertToVapi("test", bindingType)
	assert.Nil(t, err)
	assert.Equal(t, dataVal.(*data.SecretValue).Value(), "test")
	goval, err := typeConverter.ConvertToGolang(dataVal, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, goval, "test")
}

func TestSecretPtrConversion(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	var bindingType = bindings.NewOptionalType(bindings.NewSecretType())
	x := "hello"
	dataVal, err := typeConverter.ConvertToVapi(&x, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, dataVal.(*data.OptionalValue).Value().(*data.SecretValue).Value(), "hello")
	goval, err := typeConverter.ConvertToGolang(dataVal, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, "hello", *(goval.(*string)))
}

func TestSecretNilConversion(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	var bindingType = bindings.NewOptionalType(bindings.NewSecretType())
	var x *string = nil
	dataVal, err := typeConverter.ConvertToVapi(x, bindingType)
	assert.Nil(t, err)
	goval, err := typeConverter.ConvertToGolang(dataVal, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, x, goval)
}

func TestIdConversion(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	idValue := "id"
	bindingType := bindings.NewIdType([]string{"SomeResource"}, "")
	dataVal, err := typeConverter.ConvertToVapi(idValue, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, dataVal.(*data.StringValue).Value(), "id")
	goVal, err := typeConverter.ConvertToGolang(dataVal, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, goVal.(string), "id")
}

func TestIdConversionOptional(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	idValue := "id"
	bindingType := bindings.NewOptionalType(bindings.NewIdType([]string{"SomeResource"}, ""))
	dataVal, err := typeConverter.ConvertToVapi(&idValue, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, dataVal.(*data.OptionalValue).Value().(*data.StringValue).Value(), "id")
	goVal, err := typeConverter.ConvertToGolang(dataVal, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, *(goVal.(*string)), "id")
}

func TestNilIdConversion(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	var id *string = nil
	bindingType := bindings.NewOptionalType(bindings.NewIdType([]string{"SomeResource"}, ""))
	dataVal, err := typeConverter.ConvertToVapi(id, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, dataVal.(*data.OptionalValue).IsSet(), false)
	goVal, err := typeConverter.ConvertToGolang(dataVal, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, goVal.(*string), id)
}

func TestBooleanConversion(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	var dataDefinition = bindings.NewBooleanType()
	dataVal, err := typeConverter.ConvertToVapi(true, dataDefinition)
	assert.Nil(t, err)
	assert.Equal(t, dataVal.(*data.BooleanValue).Value(), true)
	goval, err := typeConverter.ConvertToGolang(dataVal, dataDefinition)
	assert.Nil(t, err)
	assert.Equal(t, goval, true)
	dataVal, err = typeConverter.ConvertToVapi(false, dataDefinition)
	assert.Nil(t, err)
	assert.Equal(t, dataVal.(*data.BooleanValue).Value(), false)
	goval, err = typeConverter.ConvertToGolang(dataVal, dataDefinition)
	assert.Nil(t, err)
	assert.Equal(t, goval, false)
}

func TestOptionalBooleanConversion(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	var dataDef = bindings.NewOptionalType(bindings.NewBooleanType())
	x := true
	dataVal, err := typeConverter.ConvertToVapi(&x, dataDef)
	assert.Nil(t, err)
	assert.Equal(t, dataVal.(*data.OptionalValue).Value().(*data.BooleanValue).Value(), true)
	goval, err := typeConverter.ConvertToGolang(dataVal, dataDef)
	assert.Nil(t, err)
	assert.Equal(t, *(goval.(*bool)), true)
}

func TestBooleanNilConversion(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	var dataDef = bindings.NewOptionalType(bindings.NewBooleanType())
	var x *bool = nil
	dataVal, err := typeConverter.ConvertToVapi(x, dataDef)
	assert.Nil(t, err)
	goval, err := typeConverter.ConvertToGolang(dataVal, dataDef)
	assert.Nil(t, err)
	assert.Equal(t, x, goval)
}

func TestEnumConversion(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	dataDef := bindings.NewEnumType("WeekDay", reflect.TypeOf(WeekDay("SUNDAY")))
	dataVal, err := typeConverter.ConvertToVapi("MONDAY", dataDef)
	assert.Nil(t, err)
	assert.Equal(t, dataVal.(*data.StringValue).Value(), "MONDAY")
	goVal, err := typeConverter.ConvertToGolang(dataVal, dataDef)
	assert.Nil(t, err)
	assert.Equal(t, goVal.(WeekDay).WeekDay(), true)
}

func TestEnumConversionFail(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	dataDef := bindings.NewEnumType("WeekDay", reflect.TypeOf(WeekDay("SUNDAY")))
	dataVal, err := typeConverter.ConvertToVapi("TUESDAY", dataDef)
	assert.Nil(t, err)
	assert.Equal(t, dataVal.(*data.StringValue).Value(), "TUESDAY")
	goVal, err := typeConverter.ConvertToGolang(dataVal, dataDef)
	assert.Nil(t, err)
	assert.Equal(t, goVal.(WeekDay).WeekDay(), false)
}

func TestOptionalEnumConversion(t *testing.T) {
	typeConverter := bindings.NewTypeConverter()
	bindingType := bindings.NewOptionalType(bindings.NewEnumType("WeekDay", reflect.TypeOf(WEEKDAY_MONDAY)))

	monday := WEEKDAY_MONDAY
	dataValue, err := typeConverter.ConvertToVapi(&monday, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, dataValue.(*data.OptionalValue).Value().(*data.StringValue).Value(), "MONDAY")

	assert.Equal(t, dataValue.(*data.OptionalValue).IsSet(), true)
	goval, err := typeConverter.ConvertToGolang(dataValue, bindingType)
	assert.Nil(t, err)

	assert.Equal(t, goval, &monday)
	assert.Equal(t, reflect.TypeOf(goval), reflect.TypeOf(&monday))
	assert.NotNil(t, goval)
}

func TestNilEnumConversion(t *testing.T) {
	typeConverter := bindings.NewTypeConverter()
	var monday *WeekDay = nil
	bindingType := bindings.NewOptionalType(bindings.NewEnumType("WeekDay", reflect.TypeOf(WEEKDAY_MONDAY)))

	dataValue, err := typeConverter.ConvertToVapi(monday, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, dataValue.(*data.OptionalValue).IsSet(), false)

	goval, err := typeConverter.ConvertToGolang(dataValue, bindingType)
	assert.Equal(t, goval, monday)
	assert.Nil(t, err)

}

func TestDateTimeConversion(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	bindingType := bindings.NewDateTimeType()
	currentTime := time.Now().Format((bindings.VAPI_DATETIME_LAYOUT))
	t1, _ := time.Parse(bindings.VAPI_DATETIME_LAYOUT, currentTime)
	dataval, err := typeConverter.ConvertToVapi(t1, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, dataval.(*data.StringValue).Value(), t1.Format(bindings.VAPI_DATETIME_LAYOUT))
	t2, err := typeConverter.ConvertToGolang(dataval, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, t2.(time.Time).String(), t1.String())
}

func TestDateTimeOptionalConversion(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	bindingType := bindings.NewOptionalType(bindings.NewDateTimeType())
	currentTime := time.Now()
	t1, _ := time.Parse("2006-01-02T15:04:05.000Z", currentTime.String())
	dataval, err := typeConverter.ConvertToVapi(&t1, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, dataval.(*data.OptionalValue).Value().(*data.StringValue).Value(), fmt.Sprintf(t1.Format("2006-01-02T15:04:05.000Z")))
	t2, err := typeConverter.ConvertToGolang(dataval, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, t2.(*time.Time).String(), t1.String())
	var t3 *time.Time = nil
	dataval, err = typeConverter.ConvertToVapi(t3, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, dataval.(*data.OptionalValue).Value(), nil)
	t2, err = typeConverter.ConvertToGolang(dataval, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, t2.(*time.Time), t3)

}

func TestNilDateTimeConversion(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	bindingType := bindings.NewOptionalType(bindings.NewDateTimeType())
	var nilTime *time.Time = nil
	dataval, err := typeConverter.ConvertToVapi(nilTime, bindingType)
	assert.Nil(t, err)
	goVal, err := typeConverter.ConvertToGolang(dataval, bindingType)
	assert.Equal(t, nilTime, goVal)

}

func TestURIConversion(t *testing.T) {
	typeConverter := bindings.NewTypeConverter()
	bindingType := bindings.NewUriType()
	urlParsed, urlErr := url.Parse("https://en.wiktionary.org/wiki/Ῥόδος")
	if urlErr != nil {
		t.Fail()
	}
	dataValue, err := typeConverter.ConvertToVapi(*urlParsed, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, dataValue.(*data.StringValue).Value(), urlParsed.String())
	goValue, err := typeConverter.ConvertToGolang(dataValue, bindingType)
	assert.Nil(t, err)
	x := goValue.(url.URL)
	assert.Equal(t, urlParsed.String(), (&x).String())

}

func TestOptionalURIConversion(t *testing.T) {
	typeConverter := bindings.NewTypeConverter()
	bindingType := bindings.NewOptionalType(bindings.NewUriType())
	urlParsed, urlErr := url.Parse("https://en.wiktionary.org/wiki/Ῥόδος")
	if urlErr != nil {
		t.Fail()
	}
	dataValue, err := typeConverter.ConvertToVapi(urlParsed, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, dataValue.(*data.OptionalValue).Value().(*data.StringValue).Value(), urlParsed.String())
	var urlNil *url.URL = nil
	dataValue, err = typeConverter.ConvertToVapi(urlNil, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, dataValue.(*data.OptionalValue).IsSet(), false)
	goValue, err := typeConverter.ConvertToGolang(dataValue, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, urlNil, goValue.(*url.URL))

}

/** AnyError tests
All stdlib entities are copied here to prevent getting a dependency on stdlib */

func TestAnyErrorOptional(t *testing.T) {
	typeConverter := bindings.NewTypeConverter()
	bindingType := bindings.NewOptionalType(bindings.NewAnyErrorType())
	dataVal, err := typeConverter.ConvertToVapi(timeOutValue, bindingType)
	assert.Nil(t, err)
	assert.NotNil(t, dataVal)
	assert.Equal(t, dataVal.(*data.OptionalValue).Value().(*data.ErrorValue).Name(), "com.vmware.vapi.std.errors.timed_out")
	goVal, err := typeConverter.ConvertToGolang(dataVal, bindingType)
	var errorVal = goVal.(*data.ErrorValue)
	goValTimeOut, err := typeConverter.ConvertToGolang(errorVal, TimedOutType())

	assert.Nil(t, err)
	assert.NotNil(t, goVal)
	assert.Equal(t, goValTimeOut.(TimedOut).Messages[0].Id, "vapi.connection")

}

func TestAnyErrorOptionalNil(t *testing.T) {
	typeConverter := bindings.NewTypeConverter()
	bindingType := bindings.NewOptionalType(bindings.NewAnyErrorType())
	var errorVal *data.ErrorValue = nil
	dataVal, err := typeConverter.ConvertToVapi(errorVal, bindingType)
	assert.Nil(t, err)
	assert.NotNil(t, dataVal)
	assert.Equal(t, dataVal.(*data.OptionalValue).IsSet(), false)

	goval, err := typeConverter.ConvertToGolang(dataVal, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, errorVal, goval)

}

func TestAnyErrorType(t *testing.T) {
	typeConverter := bindings.NewTypeConverter()
	dataValue, err := typeConverter.ConvertToVapi(timeOutValue, bindings.AnyErrorType{})
	assert.Nil(t, err)
	assert.Equal(t, dataValue, timeOutValue)
	goVal, err := typeConverter.ConvertToGolang(dataValue, TimedOutType())
	assert.Nil(t, err)
	lm := LocalizableMessage{"vapi.connection", "Could not connect to 'localhost'", []string{"localhost"}}
	messages := goVal.(TimedOut).Messages
	assert.Equal(t, 1, len(messages))
	assert.Equal(t, lm, messages[0])

}

func TestRequiredAnyErrorTypeNil(t *testing.T) {
	typeConverter := bindings.NewTypeConverter()
	dataValue, err := typeConverter.ConvertToVapi(timeOutValue, bindings.AnyErrorType{})
	assert.Nil(t, err)
	assert.Equal(t, dataValue, timeOutValue)
	goVal, err := typeConverter.ConvertToGolang(dataValue, TimedOutType())
	assert.Nil(t, err)
	lm := LocalizableMessage{"vapi.connection", "Could not connect to 'localhost'", []string{"localhost"}}
	messages := goVal.(TimedOut).Messages
	assert.Equal(t, 1, len(messages))
	assert.Equal(t, lm, messages[0])

}

func TestAnyErrorTypeWithinStruct(t *testing.T) {
	typeConverter := bindings.NewTypeConverter()
	ap := AnyErrorProperties{ErrorField: timeOutValue}
	dataValue, err := typeConverter.ConvertToVapi(ap, getBindingTypeAnyErrorProperties())
	assert.Nil(t, err)
	goval, err := typeConverter.ConvertToGolang(dataValue, getBindingTypeAnyErrorProperties())
	assert.Nil(t, err)
	assert.Equal(t, goval, ap)
}

func TestVoidConversion(t *testing.T) {
	typeConverter := bindings.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(nil, bindings.NewVoidType())
	assert.Nil(t, err)
	assert.Equal(t, dataVal.(*data.VoidValue).Value(), nil)
	goval, err := typeConverter.ConvertToGolang(dataVal, bindings.NewVoidType())
	assert.Nil(t, err)
	assert.Equal(t, nil, goval)
}

func TestStructure(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	var nativeGo = getProperties()
	var bindingType = getPropertiesBindingType()

	vapiRepresentation, err := typeConverter.ConvertToVapi(nativeGo, bindingType)
	assert.Nil(t, err)
	nativeGoAfterConversion, err := typeConverter.ConvertToGolang(vapiRepresentation, bindingType)
	assert.Nil(t, err)
	var props = nativeGoAfterConversion.(Properties)
	assert.NotNil(t, props)

}

func TestOptionalStructConversion(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	var dataDef = bindings.NewOptionalType(getPropertiesBindingType())
	x := getProperties()
	dataVal, err := typeConverter.ConvertToVapi(&x, dataDef)
	assert.Nil(t, err)
	goval, err := typeConverter.ConvertToGolang(dataVal, dataDef)
	assert.Nil(t, err)
	var props = *(goval.(*Properties))
	assert.NotNil(t, props)
	assert.Equal(t, x, props)
}

func TestOptionalStructConversionNil(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	var bindingType = bindings.NewOptionalType(getPropertiesBindingType())
	var x *Properties = nil
	dataVal, err := typeConverter.ConvertToVapi(x, bindingType)
	assert.Nil(t, err)
	goval, err := typeConverter.ConvertToGolang(dataVal, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, x, goval)
}

func TestStructOptionalFieldsNil(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	var bindingType = getPropertiesBindingType()
	var properties = Properties{}
	var item *Item = nil
	properties.MyItemOptionalValue = map[string]*Item{"a": item, "b": item}
	properties.MyOptionalItemList = []*Item{item, item}
	properties.OptionalItem = item
	properties.MyItemMap = map[string]Item{}
	properties.MyMap = map[string]int64{}
	properties.MyItemList = []Item{}
	var intPtr *int64 = nil
	properties.IntValPtr = intPtr
	dataVal, err := typeConverter.ConvertToVapi(properties, bindingType)
	assert.Nil(t, err)
	goVal, err := typeConverter.ConvertToGolang(dataVal, bindingType)
	goValProp := goVal.(Properties)
	assert.Nil(t, err)
	assert.Equal(t, goValProp.IntValPtr, intPtr)
	assert.Equal(t, goValProp.OptionalItem, item)
	assert.Equal(t, goValProp.MyOptionalItemList[0], item)
	assert.Equal(t, len(goValProp.MyOptionalItemList), 2)
	assert.Equal(t, goValProp.MyItemOptionalValue["a"], item)

}

func TestNilListConversionError(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	var bindingType = bindings.NewListType(getItemType(), reflect.TypeOf([]Item{}))
	var itemlist []Item = nil
	_, err := typeConverter.ConvertToVapi(itemlist, bindingType)
	assert.NotNil(t, err)
	assert.Equal(t, err[0].Error(), "Value of element cannot be nil")
}

func TestNilMapConversionError(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	var bindingType = bindings.NewMapType(bindings.NewStringType(), getItemType(), reflect.TypeOf(map[string]Item{}))
	var itemMap map[string]Item = nil
	_, err := typeConverter.ConvertToVapi(itemMap, bindingType)
	assert.NotNil(t, err)
	assert.Equal(t, err[0].Error(), "Value of element cannot be nil")
}

func TestNilSetConversionError(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	var bindingType = bindings.NewMapType(getItemType(), bindings.NewBooleanType(), reflect.TypeOf(map[Item]bool{}))
	var itemSet map[Item]bool = nil
	_, err := typeConverter.ConvertToVapi(itemSet, bindingType)
	assert.NotNil(t, err)
	assert.Equal(t, err[0].Error(), "Value of element cannot be nil")
}

func TestOptionalListConversion(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	var bindingType = bindings.NewOptionalType(bindings.NewListType(getItemType(), reflect.TypeOf([]Item{})))
	var itemlist []Item = nil
	_, err := typeConverter.ConvertToVapi(itemlist, bindingType)
	assert.Nil(t, err)

}

func TestNilListConversion(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	var bindingType = bindings.NewOptionalType(bindings.NewListType(getItemType(), reflect.TypeOf([]Item{})))
	var itemlist []Item = nil
	dataVal, err := typeConverter.ConvertToVapi(itemlist, bindingType)
	assert.Nil(t, err)
	goVal, err := typeConverter.ConvertToGolang(dataVal, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, itemlist, goVal)
}

func TestListPrimitiveConversion(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	var bindingType = bindings.NewListType(bindings.NewIntegerType(), reflect.TypeOf([]int64{}))

	x := []int64{1, 2, 3}
	dataVal, err := typeConverter.ConvertToVapi(x, bindingType)
	assert.Nil(t, err)
	for index, element := range dataVal.(*data.ListValue).List() {
		assert.Equal(t, x[index], element.(*data.IntegerValue).Value())
	}
	goval, err := typeConverter.ConvertToGolang(dataVal, bindingType)
	assert.Nil(t, err)
	for index, element := range goval.([]int64) {
		assert.Equal(t, x[index], element)
	}
}

func TestListOfStructConversion(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	var bindingType = bindings.NewListType(getItemType(), reflect.TypeOf([]Item{}))
	item1 := NewItem("hello")
	item2 := NewItem("world")
	itemlist := []Item{item1, item2}

	dataVal, err := typeConverter.ConvertToVapi(itemlist, bindingType)
	assert.Nil(t, err)
	for index, element := range dataVal.(*data.ListValue).List() {
		structVal := element.(*data.StructValue)
		res, _ := structVal.String("Id")
		assert.Equal(t, itemlist[index].Id, res)
	}
	goval, err := typeConverter.ConvertToGolang(dataVal, bindingType)
	assert.Nil(t, err)
	for index, element := range goval.([]Item) {
		assert.Equal(t, element.Id, itemlist[index].Id)
	}
}

func TestSetInStruct(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	//TODO: if value is set to false, server returns true. Wrong behavior?
	input := TestStructWithSet{
		TestStrSet:  map[string]bool{"i": true, "j": true, "k": true},
		TestLongSet: map[int64]bool{3: true, 2: true, 1: true}}
	bindingType := getTestStructWithSetBindingType()
	dataval, err := typeConverter.ConvertToVapi(input, bindingType)
	assert.Nil(t, err)
	goVal, err := typeConverter.ConvertToGolang(dataval, bindingType)
	assert.Nil(t, err)
	assert.EqualValues(t, input, goVal)
	outputStrSet := goVal.(TestStructWithSet).TestStrSet
	assert.Contains(t, outputStrSet, "i")
	assert.Contains(t, outputStrSet, "j")
	assert.Contains(t, outputStrSet, "k")
	outputLongSet := goVal.(TestStructWithSet).TestLongSet
	assert.Contains(t, outputLongSet, int64(1))
	assert.Contains(t, outputLongSet, int64(2))
	assert.Contains(t, outputLongSet, int64(3))
}

func TestNilSetInStructError(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	input := TestStructWithSet{TestStrSet: map[string]bool{}}
	bindingType := getTestStructWithSetBindingType()
	_, err := typeConverter.ConvertToVapi(input, bindingType)
	assert.NotNil(t, err)
	assert.Equal(t, 2, len(err))
	assert.Equal(t, "Value of element cannot be nil", err[0].Error())
	assert.Equal(t, "Invalid field 'TestLongSet' in structure 'TestStructWithSet'", err[1].Error())
}

func TestEmptyMapInStruct(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	item := TestStructWithMap{TestMap: map[string]int64{}}
	dataVal, err := typeConverter.ConvertToVapi(item, getTestStructWithMapBindingType())
	assert.Nil(t, err)
	goval, err := typeConverter.ConvertToGolang(dataVal, getTestStructWithMapBindingType())
	assert.Nil(t, err)
	op2 := goval.(TestStructWithMap)
	assert.Equal(t, item.TestMap, op2.TestMap)
	assert.Equal(t, 0, len(op2.TestMap))
}

func TestMapOfPrimitiveConversion(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	var bindingType = bindings.NewMapType(bindings.NewStringType(), bindings.NewIntegerType(), reflect.TypeOf(make(map[string]int64)))
	mymap := make(map[string]int64)
	mymap["hello"] = 1
	mymap["world"] = 2
	dataval, err := typeConverter.ConvertToVapi(mymap, bindingType)
	assert.Nil(t, err)
	for _, element := range dataval.(*data.ListValue).List() {
		structVal := element.(*data.StructValue)
		keyv, _ := structVal.String("key")
		valv, _ := structVal.Field("value")
		val := valv.(*data.IntegerValue).Value()
		assert.Equal(t, mymap[keyv], val)
	}
	goval, err := typeConverter.ConvertToGolang(dataval, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, goval.(map[string]int64)["hello"], int64(1))
	assert.Equal(t, goval.(map[string]int64)["world"], int64(2))
}

func TestMapOfStructConversion(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	var bindingType = bindings.NewMapType(bindings.NewStringType(), getItemType(), reflect.TypeOf(make(map[string]Item)))
	item1 := NewItem("hello")
	item2 := NewItem("world")
	mymap := make(map[string]Item)
	mymap["hello"] = item1
	mymap["world"] = item2
	dataval, err := typeConverter.ConvertToVapi(mymap, bindingType)
	assert.Nil(t, err)
	for _, element := range dataval.(*data.ListValue).List() {
		structVal := element.(*data.StructValue)
		keyv, _ := structVal.String("key")
		valv, _ := structVal.Field("value")
		val, _ := valv.(*data.StructValue).String("Id")
		assert.Equal(t, mymap[keyv].Id, val)
	}
	goval, err := typeConverter.ConvertToGolang(dataval, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, goval.(map[string]Item)["hello"].Id, "hello")
	assert.Equal(t, goval.(map[string]Item)["world"].Id, "world")
}

func TestOptionalMapOfStructConversion(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	var bindingType = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), getItemType(), reflect.TypeOf(make(map[string]Item))))
	var mymap map[string]Item = nil
	dataval, err := typeConverter.ConvertToVapi(mymap, bindingType)
	assert.Nil(t, err)

	goval, err := typeConverter.ConvertToGolang(dataval, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, mymap, goval)
}

func TestListOptionalPrimitiveConversion(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	var bindingType = bindings.NewListType(bindings.NewOptionalType(bindings.NewIntegerType()), reflect.TypeOf([]*int64{}))
	a := int64(1)
	b := int64(2)
	c := int64(3)
	x := []*int64{&a, &b, &c}
	dataVal, err := typeConverter.ConvertToVapi(x, bindingType)
	assert.Nil(t, err)
	for index, element := range dataVal.(*data.ListValue).List() {
		assert.Equal(t, *(x[index]), element.(*data.OptionalValue).Value().(*data.IntegerValue).Value())
	}
	goval, err := typeConverter.ConvertToGolang(dataVal, bindingType)
	assert.Nil(t, err)
	for index, element := range goval.([]*int64) {
		assert.Equal(t, *(x[index]), *(element))
	}
}

func TestListOptionalStructConversion(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	var bindingType = bindings.NewListType(bindings.NewOptionalType(getItemType()), reflect.TypeOf([]*Item{}))
	item1 := NewItem("hello")
	item2 := NewItem("world")
	itemlist := []*Item{&item1, &item2}
	dataVal, err := typeConverter.ConvertToVapi(itemlist, bindingType)
	assert.Nil(t, err)
	for index, element := range dataVal.(*data.ListValue).List() {
		s, _ := element.(*data.OptionalValue).Value().(*data.StructValue).String("Id")
		x := itemlist[index]
		assert.Equal(t, (x).Id, s)
	}
	goval, err := typeConverter.ConvertToGolang(dataVal, bindingType)
	assert.Nil(t, err)
	for index, element := range goval.([]*Item) {
		assert.Equal(t, (itemlist[index]).Id, (element).Id)
	}
}

func TestMapOptionalStructConversion(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	var bindingType = bindings.NewMapType(bindings.NewStringType(), bindings.NewOptionalType(getItemType()), reflect.TypeOf(make(map[string]*Item)))
	x := make(map[string]*Item)
	item1 := NewItem("hello")
	x["hello"] = &item1
	dataVal, err := typeConverter.ConvertToVapi(x, bindingType)
	assert.Nil(t, err)
	var listVal = dataVal.(*data.ListValue)
	for _, element := range listVal.List() {
		structVal := element.(*data.StructValue)
		val, _ := structVal.Field("value")
		itemVal := val.(*data.OptionalValue).Value()
		y, _ := itemVal.(*data.StructValue).String("Id")
		assert.Equal(t, y, "hello")

	}

	goval, err := typeConverter.ConvertToGolang(dataVal, bindingType)
	assert.Nil(t, err)
	hello := goval.(map[string]*Item)["hello"].Id
	assert.Equal(t, hello, "hello")

}

type TestStruct struct {
	List         []int64
	OptionalList []int64
	Map          map[string]string
	OptionalMap  map[string]string
	OptionalItem *Item
}

func TStructBindingType() bindings.BindingType {
	fieldMap := map[string]bindings.BindingType{}
	fieldMap["List"] = bindings.NewListType(bindings.NewIntegerType(), reflect.TypeOf([]int64{}))
	fieldMap["OptionalList"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewIntegerType(), reflect.TypeOf([]int64{})))
	fieldMap["Map"] = bindings.NewMapType(bindings.NewStringType(), bindings.NewStringType(), reflect.TypeOf(map[string]string{}))
	fieldMap["OptionalMap"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewStringType(), reflect.TypeOf(map[string]string{})))
	fieldMap["OptionalItem"] = bindings.NewOptionalType(getItemType())
	canonMap := map[string]string{}
	canonMap["List"] = "List"
	canonMap["OptionalList"] = "OptionalList"
	canonMap["Map"] = "Map"
	canonMap["OptionalMap"] = "OptionalMap"
	canonMap["OptionalItem"] = "OptionalItem"
	return bindings.NewStructType("TestStruct", fieldMap, reflect.TypeOf(TestStruct{}), canonMap, nil)

}

func TestMapListOptionalMapOptionalListWithinAStruct(t *testing.T) {
	tc := bindings.NewTypeConverter()
	list := []int64{0, 1, 2, 3}
	mymap := map[string]string{}
	mymap["a"] = "b"
	mymap["c"] = "d"
	var optionalItem *Item = nil
	tStruct := TestStruct{List: list, OptionalList: nil, Map: mymap, OptionalMap: nil, OptionalItem: optionalItem}
	dataVal, err := tc.ConvertToVapi(tStruct, TStructBindingType())
	assert.Nil(t, err)
	goVal, err := tc.ConvertToGolang(dataVal, TStructBindingType())
	assert.Nil(t, err)
	assert.Equal(t, tStruct, goVal)
	item := Item{Id: "id"}
	tStruct.OptionalItem = &item
	tStruct.OptionalMap = mymap
	tStruct.OptionalList = list
	dataVal, err = tc.ConvertToVapi(tStruct, TStructBindingType())
	assert.Nil(t, err)
	goVal, err = tc.ConvertToGolang(dataVal, TStructBindingType())
	assert.Nil(t, err)
	assert.Equal(t, tStruct, goVal)

}

func TestSetOfPrimitives(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	x := make(map[string]bool)
	x["a"] = true
	x["b"] = true
	x["c"] = true
	bindingType := bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(make(map[string]bool)))
	dataval, err := typeConverter.ConvertToVapi(x, bindingType)
	assert.Nil(t, err)
	goVal, err := typeConverter.ConvertToGolang(dataval, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, goVal.(map[string]bool)["a"], true)
	assert.Equal(t, goVal.(map[string]bool)["b"], true)

}

func TestOptionalSetType(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	var x map[string]bool = nil
	bindingType := bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(make(map[string]bool))))
	dataval, err := typeConverter.ConvertToVapi(x, bindingType)
	assert.Nil(t, err)
	goVal, err := typeConverter.ConvertToGolang(dataval, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, x, goVal)
}

func TestSetOfStructs(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	item1 := NewItem("hello")
	item2 := NewItem("world")
	x := make(map[Item]bool)
	x[item1] = true
	x[item2] = true
	bindingType := bindings.NewSetType(getItemType(), reflect.TypeOf(make(map[Item]bool)))
	dataval, err := typeConverter.ConvertToVapi(x, bindingType)
	assert.Nil(t, err)
	goVal, err := typeConverter.ConvertToGolang(dataval, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, goVal.(map[Item]bool)[item1], true)
	assert.Equal(t, goVal.(map[Item]bool)[item2], true)
}

func TestDynamicStructureOptional(t *testing.T) {
	typeConverter := bindings.NewTypeConverter()
	bindingType := bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
	fields := make(map[string]data.DataValue)
	fields["a"] = data.NewStringValue("helloDS")
	sv := data.NewStructValue("randomStructValue", fields)
	dataVal, err := typeConverter.ConvertToVapi(sv, bindingType)
	assert.Nil(t, err)
	optionalVal := dataVal.(*data.OptionalValue)
	assert.Equal(t, optionalVal.Value().(*data.StructValue).Name(), "randomStructValue")
	goVal, err := typeConverter.ConvertToGolang(dataVal, bindingType)
	var structVal = goVal.(*data.StructValue)
	assert.Nil(t, err)
	assert.Equal(t, (structVal).Name(), "randomStructValue")
}

func TestDynamicStructureOptionalNil(t *testing.T) {
	typeConverter := bindings.NewTypeConverter()
	bindingType := bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
	var x *data.StructValue = nil
	dataVal, err := typeConverter.ConvertToVapi(x, bindingType)
	assert.Nil(t, err)
	optionalVal := dataVal.(*data.OptionalValue)
	assert.Equal(t, optionalVal.IsSet(), false)
	goVal, err := typeConverter.ConvertToGolang(dataVal, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, goVal.(*data.StructValue), x)
}

func TestDynamicStructureWithinStruct(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	bindingType := getPdsBindingType()
	fields := make(map[string]data.DataValue)
	fields["a"] = data.NewStringValue("helloDS")
	x := Pds{"hello", data.NewStructValue("randomStructValue", fields)}
	dataval, err := typeConverter.ConvertToVapi(x, bindingType)
	assert.Nil(t, err)
	goval, err := typeConverter.ConvertToGolang(dataval, bindingType)
	assert.Nil(t, err)
	govalPds := goval.(Pds)
	assert.NotNil(t, dataval)
	assert.NotNil(t, goval)
	assert.NotNil(t, govalPds)
	assert.Equal(t, govalPds.Id, "hello")
	assert.Equal(t, govalPds.Ds.Name(), "randomStructValue")
	result, _ := govalPds.Ds.String("a")
	assert.Equal(t, result, "helloDS")
}

func TestRequiredDynamicStructureWithinStructNil(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	bindingType := getPdsBindingType()
	fields := make(map[string]data.DataValue)
	fields["a"] = data.NewStringValue("helloDS")
	x := Pds{"hello", nil}
	_, err := typeConverter.ConvertToVapi(x, bindingType)
	assert.NotNil(t, err)
	assert.Equal(t, err[0].Error(), "Expected a value of type '*data.StructValue', but received a value of type 'nil'")
}

func TestDynamicStructurePermissiveModeBinaryType(t *testing.T) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetPermissive(true)
	fields := make(map[string]data.DataValue)
	encoded := base64.StdEncoding.EncodeToString([]byte("Hello, 世界"))
	fields["Blob"] = data.NewStringValue(encoded)
	structValue := data.NewStructValue("somename", fields)
	sampleStructForBinary, err := typeConverter.ConvertToGolang(structValue, getSampleStructBindingTypeForBinary())
	assert.Nil(t, err)
	assert.NotNil(t, sampleStructForBinary.(SampleStructForBinary).Blob)
	assert.Equal(t, string(sampleStructForBinary.(SampleStructForBinary).Blob), encoded)
}

/**
  With clean json format (DCLI), optional wrapper will not be present while interacting with dynamic structures.
  TypeConverter should tolerate absence of optional wrappers in permissive mode.
*/
func TestDynamicStructurePermissiveModeOptionalWrapper(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	typeConverter.SetPermissive(true)
	fields := make(map[string]data.DataValue)
	fields["Id"] = data.NewStringValue("helloDS")
	structValue := data.NewStructValue("somename", fields)
	sampleStruct, err := typeConverter.ConvertToGolang(structValue, getSampleStructBindingType())
	assert.Nil(t, err)
	assert.NotNil(t, sampleStruct.(SampleStruct).Id)
	assert.Equal(t, *sampleStruct.(SampleStruct).Id, "helloDS")
}

/**
  With clean json format (DCLI), secret wrapper will not be present while interacting with dynamic structures.
  TypeConverter should tolerate and accept stringvalue instead of secretvalue in permissive mode.
*/
func TestDynamicStructurePermissiveModeStringValueForSecretValue(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	typeConverter.SetPermissive(true)
	fields := make(map[string]data.DataValue)
	fields["Password"] = data.NewStringValue("mysupersecurepassword")
	structValue := data.NewStructValue("somename", fields)
	sampleStruct, err := typeConverter.ConvertToGolang(structValue, getSampleStructStringValueForSecretValueBindingType())
	assert.Nil(t, err)
	assert.NotNil(t, sampleStruct.(SampleStructStringValueForSecret).Password)
	assert.Equal(t, "mysupersecurepassword", sampleStruct.(SampleStructStringValueForSecret).Password)
}

func TestOptionalDynamicStructureWithinStruct(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	bindingType := getOptionalPdsBindingType()
	fields := make(map[string]data.DataValue)
	fields["a"] = data.NewStringValue("helloDS")
	ds := data.NewStructValue("randomStructValue", fields)
	x := OptionalPds{"hello", ds}
	dataval, err := typeConverter.ConvertToVapi(x, bindingType)
	assert.Nil(t, err)
	goval, err := typeConverter.ConvertToGolang(dataval, bindingType)
	assert.Nil(t, err)
	govalPds := goval.(OptionalPds)
	assert.NotNil(t, dataval)
	assert.NotNil(t, goval)
	assert.NotNil(t, govalPds)
	assert.Equal(t, govalPds.Id, "hello")
	var dsgoval = (govalPds.Ds)
	assert.Equal(t, dsgoval.Name(), "randomStructValue")
	result, _ := dsgoval.String("a")
	assert.Equal(t, result, "helloDS")
}

func TestOptionalDynamicStructureWithinStructWithNil(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	bindingType := getOptionalPdsBindingType()
	fields := make(map[string]data.DataValue)
	fields["a"] = data.NewStringValue("helloDS")

	x := OptionalPds{"hello", nil}
	dataval, err := typeConverter.ConvertToVapi(x, bindingType)
	assert.Nil(t, err)
	goval, err := typeConverter.ConvertToGolang(dataval, bindingType)
	assert.Nil(t, err)
	govalPds := goval.(OptionalPds)
	assert.NotNil(t, dataval)
	assert.NotNil(t, goval)
	assert.NotNil(t, govalPds)
	assert.Equal(t, govalPds.Id, "hello")
	assert.Nil(t, govalPds.Ds)

}

//TODO
// Test Set within a struct
// Test Enum within a struct
// Test Id within a struct

func TestOpaquePrimitive(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	var bindingType = bindings.NewOpaqueType()
	var dataVal = data.NewIntegerValue(10)
	result, err := typeConverter.ConvertToGolang(dataVal, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, result, dataVal)
	result2, err := typeConverter.ConvertToVapi(dataVal, bindingType)
	assert.Equal(t, result2, dataVal)
}

func TestOpaqueGeneric(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	listValue := data.NewListValue()
	listValue.Add(data.NewIntegerValue(10))
	var bindingType = bindings.NewOpaqueType()
	result, err := typeConverter.ConvertToGolang(listValue, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, listValue, result)
	result2, err := typeConverter.ConvertToVapi(result, bindingType)
	assert.Equal(t, result2, listValue)
}

func TestOpaqueInvalid(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(10, bindings.NewOpaqueType())
	assert.Nil(t, dataVal)
	assert.NotNil(t, err)
}

type ResourceItem struct {
	PropertyValues []data.DataValue
}

func ResourceItemBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["PropertyValues"] = bindings.NewListType(bindings.NewOptionalType(bindings.NewOpaqueType()), reflect.TypeOf([]data.DataValue{}))
	fieldNameMap["PropertyValues"] = "PropertyValues"

	return bindings.NewStructType("ResourceItem", fields, reflect.TypeOf(ResourceItem{}), fieldNameMap, nil)
}

func TestStructWithListOptionalOpaque(t *testing.T) {
	intValue := data.NewIntegerValue(10)
	stringValue := data.NewStringValue("hello")
	listVal := data.NewListValue()
	listVal.Add(intValue)
	listVal.Add(stringValue)
	propValues := []data.DataValue{intValue, stringValue, listVal, nil}
	resourceItem := ResourceItem{PropertyValues: propValues}
	typeConverter := bindings.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(resourceItem, ResourceItemBindingType())
	assert.Nil(t, err)
	goVal, err := typeConverter.ConvertToGolang(dataVal, ResourceItemBindingType())
	assert.Equal(t, goVal.(ResourceItem), resourceItem)
}

// TODO
// write a test case where map entry's  value field is missing and
// server fills in nil as value for that key.
func TestOptionalFieldsMissing(t *testing.T) {
	properties := getProperties()
	var nilItem *Item = nil
	properties.MyOptionalItemList = []*Item{nilItem, nilItem}
	properties.MyItemOptionalValue = map[string]*Item{"a": nilItem, "b": nilItem}
	propertiesBindingType := getPropertiesBindingType()
	typeConverter := bindings.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(properties, propertiesBindingType)
	structVal := dataVal.(*data.StructValue)
	newStructVal := data.NewStructValue(structVal.Name(), nil)
	//remove all optional fields from datavalue representation
	for field, fieldVal := range structVal.Fields() {
		if _, ok := fieldVal.(*data.OptionalValue); ok {
			continue
		} else {
			newStructVal.SetField(field, fieldVal)
		}
	}

	assert.NotNil(t, newStructVal)
	goVal, err := typeConverter.ConvertToGolang(newStructVal, propertiesBindingType)
	assert.Nil(t, err)
	var x *int64 = nil
	var xItem *Item = nil
	assert.Equal(t, goVal.(Properties).IntValPtr, x)
	assert.Equal(t, xItem, goVal.(Properties).OptionalItem)

}

func TestBindingTypeNil(t *testing.T) {
	typeConverter := bindings.NewTypeConverter()
	dataValue := data.NewIntegerValue(10)
	_, err := typeConverter.ConvertToGolang(dataValue, nil)
	assert.NotNil(t, err)
	assert.Equal(t, 1, len(err))
	assert.Equal(t, err[0].Error(), "Binding type cannot be nil")
	_, err = typeConverter.ConvertToVapi(2, nil)
	assert.NotNil(t, err)
	assert.Equal(t, 1, len(err))
	assert.Equal(t, err[0].Error(), "Binding type cannot be nil")
}

func TestCrossReferenceStructConversion(t *testing.T) {
	struct_one := CrossReferenceStruct{}
	struct_two := CrossReferenceStructSecond{}
	struct_one.Two = &struct_two
	struct_two.One = nil
	typeConverter := bindings.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(struct_one, CrossReferenceStructType())
	assert.Nil(t, err)
	assert.NotNil(t, dataVal)
	goVal, err := typeConverter.ConvertToGolang(dataVal, CrossReferenceStructType())
	assert.Nil(t, err)
	assert.NotNil(t, goVal)
	assert.Equal(t, *goVal.(CrossReferenceStruct).Two, struct_two)
}

type MapValueList struct {
	ItemMapList map[string][]Item
}

func MapValueListBindingType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ItemMapList"] = bindings.NewMapType(bindings.NewStringType(), bindings.NewListType(getItemType(), reflect.TypeOf([]Item{})), reflect.TypeOf(map[string][]Item{}))
	fieldNameMap["ItemMapList"] = "ItemMapList"
	return bindings.NewStructType("MapValueList", fields, reflect.TypeOf(MapValueList{}), fieldNameMap, nil)
}

func TestMapValueList(t *testing.T) {
	mvL := MapValueList{}
	item1 := Item{Id: "a"}
	mvL.ItemMapList = map[string][]Item{}
	mvL.ItemMapList["a"] = []Item{item1}
	mvL.ItemMapList["b"] = []Item{}
	typeConverter := bindings.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(mvL, MapValueListBindingType())
	assert.Nil(t, err)
	assert.NotNil(t, dataVal)
	goVal, err := typeConverter.ConvertToGolang(dataVal, MapValueListBindingType())
	assert.Nil(t, err)
	res := goVal.(MapValueList)
	assert.Equal(t, item1, res.ItemMapList["a"][0])
	assert.Equal(t, 0, len(res.ItemMapList["b"]))
}

type MapValueMap struct {
	ItemMapMap map[string]map[string]Item
}

func MapValueMapBindingType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ItemMapMap"] = bindings.NewMapType(bindings.NewStringType(), bindings.NewMapType(bindings.NewStringType(), getItemType(), reflect.TypeOf(map[string]Item{})), reflect.TypeOf(map[string]map[string]Item{}))
	fieldNameMap["ItemMapMap"] = "ItemMapMap"
	return bindings.NewStructType("MapValueMap", fields, reflect.TypeOf(MapValueMap{}), fieldNameMap, nil)
}

func TestMapValueMap(t *testing.T) {
	mvm := MapValueMap{}
	item1 := Item{Id: "a"}
	mvm.ItemMapMap = map[string]map[string]Item{}
	itemMap := map[string]Item{}
	itemMap["a"] = item1
	mvm.ItemMapMap["a"] = itemMap
	typeConverter := bindings.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(mvm, MapValueMapBindingType())
	assert.Nil(t, err)
	assert.NotNil(t, dataVal)
	goVal, err := typeConverter.ConvertToGolang(dataVal, MapValueMapBindingType())
	assert.Nil(t, err)
	res := goVal.(MapValueMap)
	assert.Equal(t, itemMap, res.ItemMapMap["a"])

}

func TestMultiLevelNestingSelfRecursiveHolder(t *testing.T) {
	var third = NestedSelfRecursiveStruct{
		Data: 3,
		Next: nil,
	}
	var second = NestedSelfRecursiveStruct{
		Data: 2,
		Next: &third,
	}
	var first = NestedSelfRecursiveStruct{
		Data: 1,
		Next: &second,
	}

	var roots = map[string]NestedSelfRecursiveStruct{"first": first, "second": second}
	var mapOfSelfRecStructHldr = MapOfSelfRecursiveStructHolder{Data: 4, Roots: roots}
	var listOfMapOfSelfRecStructHldr = []MapOfSelfRecursiveStructHolder{mapOfSelfRecStructHldr, mapOfSelfRecStructHldr}
	var roots2 = map[string][]MapOfSelfRecursiveStructHolder{"first": listOfMapOfSelfRecStructHldr, "second": listOfMapOfSelfRecStructHldr}

	var input = MultiLevelNestingSelfRecursiveHolder{Data: 5, Roots: roots2}
	typeConverter := bindings.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(input, MultiLevelNestingSelfRecursiveHolderType())
	assert.Nil(t, err)
	assert.NotEmpty(t, dataVal)

	goVal, err := typeConverter.ConvertToGolang(dataVal, MultiLevelNestingSelfRecursiveHolderType())
	assert.Nil(t, err)
	assert.NotEmpty(t, goVal)
}

func TestListOptionalNil(t *testing.T) {
	bindingType := bindings.NewListType(bindings.NewOptionalType(bindings.NewStringType()), reflect.TypeOf([]*string{}))
	var stringPtr *string = nil
	var x = []*string{stringPtr}
	typeConverter := bindings.NewTypeConverter()
	dataValue, err := typeConverter.ConvertToVapi(x, bindingType)
	assert.Nil(t, err)
	golangNative, err := typeConverter.ConvertToGolang(dataValue, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, golangNative, x)
}

func TestRestNativeDateTime(t *testing.T) {
	bindingType := bindings.NewDateTimeType()
	tc := bindings.NewTypeConverter()
	tc.SetMode(bindings.REST)
	inputTime, terr := time.Parse(time.RFC3339, "2006-01-02T15:04:05.000Z")
	assert.Nil(t, terr)
	dataValue, err := tc.ConvertToVapi(inputTime, bindingType)
	assert.Nil(t, err)
	golangNative, err := tc.ConvertToGolang(dataValue, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, golangNative, inputTime)
}
