package bindings_test

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"reflect"
)

type Properties struct {
	IntVal              int64
	IntValPtr           *int64
	MyItem              Item
	OptionalItem        *Item
	MyItemList          []Item
	MyOptionalItemList  []*Item
	MyMap               map[string]int64
	MyItemMap           map[string]Item
	MyItemOptionalValue map[string]*Item
	MyItemOptionalList  []Item
	//StrVal             string
	//BoolVal            bool
	//OptVal             *int
	//ListVal            []int
}

type Item struct {
	Id string
}

func getItemType() bindings.BindingType {
	var fields = make(map[string]bindings.BindingType)
	fields["Id"] = bindings.NewStringType()
	canonicalFieldMap := make(map[string]string)
	canonicalFieldMap["Id"] = "Id"
	return bindings.NewStructType("Item", fields, reflect.TypeOf(Item{}), canonicalFieldMap, nil)
}

type OptionalPds struct {
	Id string
	Ds *data.StructValue
}

func getOptionalPdsBindingType() bindings.BindingType {
	var fields = make(map[string]bindings.BindingType)
	fields["Id"] = bindings.NewStringType()
	fields["Ds"] = bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
	canonicalFieldMap := make(map[string]string)
	canonicalFieldMap["Id"] = "Id"
	canonicalFieldMap["Ds"] = "Ds"
	return bindings.NewStructType("Pds", fields, reflect.TypeOf(OptionalPds{}), canonicalFieldMap, nil)
}

type Pds struct {
	Id string
	Ds *data.StructValue
}

func getPdsBindingType() bindings.BindingType {
	var fields = make(map[string]bindings.BindingType)
	fields["Id"] = bindings.NewStringType()
	fields["Ds"] = bindings.NewDynamicStructType(nil, bindings.JSONRPC)
	canonicalFieldMap := make(map[string]string)
	canonicalFieldMap["Id"] = "Id"
	canonicalFieldMap["Ds"] = "Ds"
	return bindings.NewStructType("Pds", fields, reflect.TypeOf(Pds{}), canonicalFieldMap, nil)
}

func NewItem(id string) Item {
	return Item{Id: id}
}

func getPropertiesBindingType() bindings.BindingType {
	var fields = make(map[string]bindings.BindingType)
	fields["IntVal"] = bindings.NewIntegerType()
	fields["IntValPtr"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["MyItem"] = getItemType()
	fields["OptionalItem"] = bindings.NewOptionalType(getItemType())
	fields["MyItemList"] = bindings.NewListType(getItemType(), reflect.TypeOf([]Item{}))
	fields["MyOptionalItemList"] = bindings.NewListType(bindings.NewOptionalType(getItemType()), reflect.TypeOf([]*Item{}))
	fields["MyMap"] = bindings.NewMapType(bindings.NewStringType(), bindings.NewIntegerType(), reflect.TypeOf(map[string]int64{}))
	fields["MyItemMap"] = bindings.NewMapType(bindings.NewStringType(), getItemType(), reflect.TypeOf(map[string]Item{}))
	fields["MyItemOptionalValue"] = bindings.NewMapType(bindings.NewStringType(), bindings.NewOptionalType(getItemType()), reflect.TypeOf(map[string]*Item{}))
	fields["MyItemOptionalList"] = bindings.NewOptionalType(bindings.NewListType(getItemType(), reflect.TypeOf([]Item{})))
	//fields["StrVal"] = NewStringType()
	//fields["BoolVal"] = NewBooleanType()
	//fields["OptVal"] = NewOptionalType(NewIntegerType())
	//fields["ListVal"] = NewListType(NewIntegerType(), reflect.TypeOf(0))
	canonicalFieldMap := make(map[string]string)
	canonicalFieldMap["IntVal"] = "IntVal"
	canonicalFieldMap["IntValPtr"] = "IntValPtr"
	canonicalFieldMap["MyItem"] = "MyItem"
	canonicalFieldMap["OptionalItem"] = "OptionalItem"
	canonicalFieldMap["MyItemList"] = "MyItemList"
	canonicalFieldMap["MyOptionalItemList"] = "MyOptionalItemList"
	canonicalFieldMap["MyMap"] = "MyMap"
	canonicalFieldMap["MyItemMap"] = "MyItemMap"
	canonicalFieldMap["MyItemOptionalValue"] = "MyItemOptionalValue"
	canonicalFieldMap["MyItemOptionalList"] = "MyItemOptionalList"
	var structType = bindings.NewStructType("Properties", fields,
		reflect.TypeOf(Properties{}), canonicalFieldMap, nil)
	return structType
}

type SampleStruct struct {
	Id *string
}

func getSampleStructBindingType() bindings.BindingType {
	var fields = make(map[string]bindings.BindingType)
	fields["Id"] = bindings.NewOptionalType(bindings.NewStringType())
	canonicalFieldMap := make(map[string]string)
	canonicalFieldMap["Id"] = "Id"
	return bindings.NewStructType("SampleStruct", fields,
		reflect.TypeOf(SampleStruct{}), canonicalFieldMap, nil)
}

type SampleStructForBinary struct {
	Blob []byte
}

func getSampleStructBindingTypeForBinary() bindings.BindingType {
	var fields = make(map[string]bindings.BindingType)
	fields["Blob"] = bindings.NewBlobType()
	canonicalFieldMap := make(map[string]string)
	canonicalFieldMap["Blob"] = "Blob"
	return bindings.NewStructType("SampleStruct", fields,
		reflect.TypeOf(SampleStructForBinary{}), canonicalFieldMap, nil)
}

type SampleStructStringValueForSecret struct {
	Password string
}

func getSampleStructStringValueForSecretValueBindingType() bindings.BindingType {
	var fields = make(map[string]bindings.BindingType)
	fields["Password"] = bindings.NewSecretType()
	canonicalFieldMap := make(map[string]string)
	canonicalFieldMap["Password"] = "Password"
	return bindings.NewStructType("SampleStructStringValueForSecret", fields,
		reflect.TypeOf(SampleStructStringValueForSecret{}), canonicalFieldMap, nil)
}

type WeekDay string

const (
	WEEKDAY_SUNDAY WeekDay = "SUNDAY"
	WEEKDAY_MONDAY WeekDay = "MONDAY"
)

func (day WeekDay) WeekDay() bool {
	switch day {
	case WEEKDAY_SUNDAY, WEEKDAY_MONDAY:
		return true
	default:
		return false
	}
}

type TestStructWithMap struct {
	TestMap map[string]int64
}

func getTestStructWithMapBindingType() bindings.BindingType {
	var fields = make(map[string]bindings.BindingType)
	fields["TestMap"] = bindings.NewMapType(bindings.NewStringType(),
		bindings.NewIntegerType(), reflect.TypeOf(map[string]int64{}))
	canonicalFieldMap := make(map[string]string)
	canonicalFieldMap["TestMap"] = "TestMap"
	return bindings.NewStructType("TestItem", fields,
		reflect.TypeOf(TestStructWithMap{}), canonicalFieldMap, nil)
}

type TestStructWithSet struct {
	TestStrSet  map[string]bool
	TestLongSet map[int64]bool
}

func getTestStructWithSetBindingType() bindings.BindingType {
	var fields = make(map[string]bindings.BindingType)
	fields["TestStrSet"] = bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{}))
	fields["TestLongSet"] = bindings.NewSetType(bindings.NewIntegerType(), reflect.TypeOf(map[int64]bool{}))
	canonicalFieldMap := make(map[string]string)
	canonicalFieldMap["TestStrSet"] = "TestStrSet"
	canonicalFieldMap["TestLongSet"] = "TestLongSet"
	return bindings.NewStructType("TestStructWithSet", fields,
		reflect.TypeOf(TestStructWithSet{}), canonicalFieldMap, nil)
}

type CrossReferenceStruct struct {
	Two *CrossReferenceStructSecond
}

func CrossReferenceStructType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["two"] = bindings.NewOptionalType(bindings.NewReferenceType(CrossReferenceStructSecondType))
	fieldNameMap["two"] = "Two"

	return bindings.NewStructType("vmodl.test.uber.structures.uber_recursive.cross_reference_struct",
		fields, reflect.TypeOf(CrossReferenceStruct{}), fieldNameMap, nil)
}

type CrossReferenceStructSecond struct {
	One *CrossReferenceStruct
}

func CrossReferenceStructSecondType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["one"] = bindings.NewOptionalType(bindings.NewReferenceType(CrossReferenceStructType))
	fieldNameMap["one"] = "One"
	return bindings.NewStructType("vmodl.test.uber.structures.uber_recursive.cross_reference_struct_second",
		fields, reflect.TypeOf(CrossReferenceStructSecond{}), fieldNameMap, nil)
}

type NestedSelfRecursiveStruct struct {
	Data int64
	Next *NestedSelfRecursiveStruct
}

func NestedSelfRecursiveStructType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["data"] = bindings.NewIntegerType()
	fieldNameMap["data"] = "Data"
	fields["next"] = bindings.NewOptionalType(bindings.NewReferenceType(NestedSelfRecursiveStructType))
	fieldNameMap["next"] = "Next"
	return bindings.NewStructType("vmodl.test.uber.structures.uber_recursive.nested_self_recursive_struct",
		fields, reflect.TypeOf(NestedSelfRecursiveStruct{}), fieldNameMap, nil)
}

type MapOfSelfRecursiveStructHolder struct {
	Data  int64
	Roots map[string]NestedSelfRecursiveStruct
}

func MapOfSelfRecursiveStructHolderType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["data"] = bindings.NewIntegerType()
	fieldNameMap["data"] = "Data"
	fields["roots"] = bindings.NewMapType(bindings.NewStringType(),
		bindings.NewReferenceType(NestedSelfRecursiveStructType),
		reflect.TypeOf(make(map[string]NestedSelfRecursiveStruct)))
	fieldNameMap["roots"] = "Roots"

	return bindings.NewStructType("vmodl.test.uber.structures.uber_recursive.map_of_self_recursive_struct_holder",
		fields, reflect.TypeOf(MapOfSelfRecursiveStructHolder{}), fieldNameMap, nil)
}

type MultiLevelNestingSelfRecursiveHolder struct {
	Data  int64
	Roots map[string][]MapOfSelfRecursiveStructHolder
}

func MultiLevelNestingSelfRecursiveHolderType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["data"] = bindings.NewIntegerType()
	fieldNameMap["data"] = "Data"
	fields["roots"] = bindings.NewMapType(bindings.NewStringType(),
		bindings.NewListType(bindings.NewReferenceType(MapOfSelfRecursiveStructHolderType),
			reflect.TypeOf([]MapOfSelfRecursiveStructHolder{})),
		reflect.TypeOf(make(map[string][]MapOfSelfRecursiveStructHolder)))
	fieldNameMap["roots"] = "Roots"

	return bindings.NewStructType(
		"vmodl.test.uber.structures.uber_recursive.multi_level_nestig_self_recursive_holder",
		fields, reflect.TypeOf(MultiLevelNestingSelfRecursiveHolder{}), fieldNameMap, nil)
}

type TestUnionStruct struct {
	TestEnum  string
	LongVal   *int64
	StringVal *string
}

func getTestUnionStructBindingType() bindings.BindingType {
	var fields = make(map[string]bindings.BindingType)
	fields["test_enum"] = bindings.NewStringType()
	fields["long_val"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["string_val"] = bindings.NewOptionalType(bindings.NewStringType())
	canonicalFieldMap := make(map[string]string)
	canonicalFieldMap["test_enum"] = "TestEnum"
	canonicalFieldMap["long_val"] = "LongVal"
	canonicalFieldMap["string_val"] = "StringVal"
	unionValidator := bindings.NewUnionValidator("test_enum",
		map[string][]bindings.FieldData{
			"LONG":   []bindings.FieldData{bindings.NewFieldData("long_val", true)},
			"STRING": []bindings.FieldData{bindings.NewFieldData("string_val", true)},
			"NONE":   []bindings.FieldData{},
		},
	)
	return bindings.NewStructType("TestUnionStruct", fields, reflect.TypeOf(TestUnionStruct{}),
		canonicalFieldMap, []bindings.Validator{unionValidator})
}

func getItemDynamicStructBindingType() bindings.BindingType {
	return bindings.NewDynamicStructType(
		[]bindings.ReferenceType{bindings.ReferenceType{getItemType}}, bindings.JSONRPC)
}

type TestIsOneOfStruct struct {
	IsOneOfVal string
	SameVal    *string
}

func getTestIsOneOfStructBindingType() bindings.BindingType {
	var fields = make(map[string]bindings.BindingType)
	fields["is_one_of_val"] = bindings.NewStringType()
	fields["some_val"] = bindings.NewOptionalType(bindings.NewStringType())
	canonicalFieldMap := make(map[string]string)
	canonicalFieldMap["is_one_of_val"] = "IsOneOfVal"
	canonicalFieldMap["some_val"] = "SameVal"
	isValidator := bindings.NewIsOneOfValidator("is_one_of_val",
		[]string{"val1", "val2"},
	)
	return bindings.NewStructType("TestIsOneOfStruct", fields, reflect.TypeOf(TestIsOneOfStruct{}),
		canonicalFieldMap, []bindings.Validator{isValidator})
}
