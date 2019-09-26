package bindings_test

import (
	"github.com/stretchr/testify/assert"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"testing"
)

//HV = HasFieldsOfValidator
func TestHV(t *testing.T) {
	prop1 := getProperties()
	validatePropertiesToDynamicStruct(t, prop1)

	prop2 := getProperties()
	prop2.OptionalItem = nil
	validatePropertiesToDynamicStruct(t, prop2)

	prop3 := getProperties()
	prop3.MyItemList = []Item{}
	var emptyItem *Item
	prop3.MyOptionalItemList = append(prop3.MyOptionalItemList, emptyItem)
	validatePropertiesToDynamicStruct(t, prop3)

	propDataVal1 := getPropertiesStructValue(t)
	delete(propDataVal1.Fields(), "OptionalItem")
	msgList := validateDynamicStructVal(propDataVal1)
	assert.Equal(t, 0, len(msgList))
}

func TestHVMultiple(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	item := Item{"Hi"}
	itemDataVal, err := typeConverter.ConvertToVapi(item, getItemType())
	itemStruct := itemDataVal.(*data.StructValue)
	assert.Nil(t, err)
	testStructWithSet := TestStructWithSet{
		TestStrSet:  map[string]bool{"i": true, "j": true},
		TestLongSet: map[int64]bool{2: true, 1: true}}
	testStruct2DataVal, err := typeConverter.ConvertToVapi(testStructWithSet, getTestStructWithSetBindingType())
	assert.Nil(t, err)
	testStruct2 := testStruct2DataVal.(*data.StructValue)

	bothStructsFields := make(map[string]data.DataValue)
	for k, v := range itemStruct.Fields() {
		bothStructsFields[k] = v
	}
	for k, v := range testStruct2.Fields() {
		bothStructsFields[k] = v
	}
	bothStructsFields["extra_field"] = data.NewStringValue("something")
	dynamicStructVal := data.NewStructValue("", bothStructsFields)

	hv := bindings.NewHasFieldsOfValidator(
		[]bindings.ReferenceType{
			bindings.NewReferenceType(getItemType),
			bindings.NewReferenceType(getTestStructWithSetBindingType)},
		bindings.JSONRPC)
	msgList := hv.Validate(dynamicStructVal)
	assert.Equal(t, 0, len(msgList))
}

func TestHVMultipleNegative(t *testing.T) {
	var typeConverter = bindings.NewTypeConverter()
	testStructWithSet := TestStructWithSet{
		TestStrSet:  map[string]bool{"i": true, "j": true},
		TestLongSet: map[int64]bool{2: true, 1: true}}
	testStruct2DataVal, err := typeConverter.ConvertToVapi(testStructWithSet, getTestStructWithSetBindingType())
	assert.Nil(t, err)
	testStruct2 := testStruct2DataVal.(*data.StructValue)

	someFields := testStruct2.Fields()
	delete(someFields, "TestLongSet")
	someFields["Id"] = data.NewIntegerValue(10)
	dynamicStructVal := data.NewStructValue("", someFields)

	hv := bindings.NewHasFieldsOfValidator(
		[]bindings.ReferenceType{
			bindings.NewReferenceType(getItemType),
			bindings.NewReferenceType(getTestStructWithSetBindingType)},
		bindings.JSONRPC)
	msgList := hv.Validate(dynamicStructVal)
	assert.Equal(t, 5, len(msgList))
	assert.Equal(t, "Expected a value of type 'StringValue', but received a value of type '*data.IntegerValue'",
		msgList[0].Error())
	assert.Equal(t, "Invalid field 'Id' in structure 'Item'",
		msgList[1].Error())
	assert.Equal(t, "Expected valid fields of structure 'Item'",
		msgList[2].Error())
	assert.Equal(t, "Field 'TestLongSet' missing from structure 'TestStructWithSet'",
		msgList[3].Error())
	assert.Equal(t, "Expected valid fields of structure 'TestStructWithSet'",
		msgList[4].Error())
}

func TestHVMissingField(t *testing.T) {
	validateMissingField(t, "MyItem")
	validateMissingField(t, "MyItemList")
}

func TestHVMissingFieldInStruct(t *testing.T) {
	propDataVal1 := getPropertiesStructValue(t)
	myItemDataVal, _ := propDataVal1.Field("MyItem")
	myItemStruct := myItemDataVal.(*data.StructValue)
	delete(myItemStruct.Fields(), "Id")
	msgList := validateDynamicStructVal(propDataVal1)
	assert.Equal(t, 3, len(msgList))
	assert.Equal(t, msgList[0].Error(), "Field 'Id' missing from structure 'Item'")
	assert.Equal(t, msgList[1].Error(), "Invalid field 'MyItem' in structure 'Properties'")
	assert.Equal(t, msgList[2].Error(), "Expected valid fields of structure 'Properties'")
}

func TestHVMissingFieldInList(t *testing.T) {
	propDataVal1 := getPropertiesStructValue(t)
	myItemListDataVal, _ := propDataVal1.Field("MyItemList")
	myItemList := myItemListDataVal.(*data.ListValue)
	myItemDataVal := myItemList.Get(0)
	myItemStruct := myItemDataVal.(*data.StructValue)
	delete(myItemStruct.Fields(), "Id")
	msgList := validateDynamicStructVal(propDataVal1)
	assert.Equal(t, 4, len(msgList))
	assert.Equal(t, msgList[0].Error(), "Field 'Id' missing from structure 'Item'")
	assert.Equal(t, msgList[1].Error(), "The value at index '0' of the list is invalid")
	assert.Equal(t, msgList[2].Error(), "Invalid field 'MyItemList' in structure 'Properties'")
	assert.Equal(t, msgList[3].Error(), "Expected valid fields of structure 'Properties'")
}

func TestHVMissingFieldInMap(t *testing.T) {
	propDataVal1 := getPropertiesStructValue(t)
	myItemMapDataVal, _ := propDataVal1.Field("MyItemMap")
	myItemMapStructList := myItemMapDataVal.(*data.ListValue)
	myItemMapStruct := myItemMapStructList.Get(0).(*data.StructValue)
	myItemDataVal, _ := myItemMapStruct.Field("value")
	myItemStruct := myItemDataVal.(*data.StructValue)
	delete(myItemStruct.Fields(), "Id")
	msgList := validateDynamicStructVal(propDataVal1)
	assert.Equal(t, 4, len(msgList))
	assert.Equal(t, msgList[0].Error(), "Field 'Id' missing from structure 'Item'")
	assert.Contains(t, msgList[1].Error(), "Invalid value for key 'Itemkey")
	assert.Equal(t, msgList[2].Error(), "Invalid field 'MyItemMap' in structure 'Properties'")
	assert.Equal(t, msgList[3].Error(), "Expected valid fields of structure 'Properties'")
}

func TestHVForStructVal(t *testing.T) {
	itemDsType := getItemDynamicStructBindingType()
	tc := bindings.NewTypeConverter()

	testStruct := data.NewStructValue(
		"TestStruct",
		map[string]data.DataValue{
			"Id": data.NewStringValue("some_id"),
		})
	_, err := tc.ConvertToGolang(testStruct, itemDsType)
	assert.Nil(t, err)

	testStruct.SetField("Id", data.NewIntegerValue(100))
	_, msgList := tc.ConvertToGolang(testStruct, itemDsType)
	assert.Equal(t, 3, len(msgList))
	assert.Equal(t, msgList[0].Error(), "Expected a value of type 'StringValue', but received a value of type '*data.IntegerValue'")
	assert.Equal(t, msgList[1].Error(), "Invalid field 'Id' in structure 'Item'")
	assert.Equal(t, msgList[2].Error(), "Expected valid fields of structure 'Item'")
}

func TestHVForGoVal(t *testing.T) {
	itemDsType := getItemDynamicStructBindingType()
	tc := bindings.NewTypeConverter()

	item := Item{"some_id"}
	itemStruct, _ := tc.ConvertToVapi(item, getItemType())
	_, err := tc.ConvertToVapi(itemStruct, itemDsType)
	assert.Nil(t, err)

	itemStruct.(*data.StructValue).SetField("Id", data.NewIntegerValue(100))
	_, msgList := tc.ConvertToVapi(itemStruct, itemDsType)
	assert.NotNil(t, msgList)
	assert.Equal(t, 3, len(msgList))
	assert.Equal(t, msgList[0].Error(), "Expected a value of type 'StringValue', but received a value of type '*data.IntegerValue'")
	assert.Equal(t, msgList[1].Error(), "Invalid field 'Id' in structure 'Item'")
	assert.Equal(t, msgList[2].Error(), "Expected valid fields of structure 'Item'")
}

func validateMissingField(t *testing.T, fieldName string) {
	propDataVal1 := getPropertiesStructValue(t)
	delete(propDataVal1.Fields(), fieldName)
	msgList := validateDynamicStructVal(propDataVal1)
	assert.Equal(t, 2, len(msgList))
	assert.Equal(t, msgList[0].Error(), "Field '"+fieldName+"' missing from structure 'Properties'")
	assert.Equal(t, msgList[1].Error(), "Expected valid fields of structure 'Properties'")
}

func validateDynamicStructVal(prop *data.StructValue) []error {
	hv := bindings.NewHasFieldsOfValidator(
		[]bindings.ReferenceType{bindings.NewReferenceType(getPropertiesBindingType)}, bindings.JSONRPC)
	msgList := hv.Validate(prop)
	return msgList
}

func getPropertiesStructValue(t *testing.T) *data.StructValue {
	propBindingType := getPropertiesBindingType()
	prop := getProperties()
	var typeConverter = bindings.NewTypeConverter()
	propDataVal, err := typeConverter.ConvertToVapi(prop, propBindingType)
	assert.Nil(t, err)
	return propDataVal.(*data.StructValue)
}

func validatePropertiesToDynamicStruct(t *testing.T, prop Properties) {
	propBindingType := getPropertiesBindingType()
	var typeConverter = bindings.NewTypeConverter()
	propDataVal, err := typeConverter.ConvertToVapi(prop, propBindingType)
	assert.Nil(t, err)
	propStruct, ok := propDataVal.(*data.StructValue)
	assert.True(t, ok)

	hv := bindings.NewHasFieldsOfValidator(
		[]bindings.ReferenceType{bindings.NewReferenceType(getPropertiesBindingType)},
		bindings.JSONRPC)
	msgList := hv.Validate(propStruct)
	assert.Equal(t, 0, len(msgList))
}
