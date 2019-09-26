package bindings_test

import (
	"github.com/stretchr/testify/assert"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"reflect"
	"testing"
)


type TopLevelStruct struct {
	Dbl float64
}

func getTopLevelStructBindingType() bindings.BindingType {
	var fields = make(map[string]bindings.BindingType)
	fields["Dbl"] = bindings.NewDoubleType()
	canonicalFieldMap := make(map[string]string)
	canonicalFieldMap["Dbl"] = "Dbl"
	var structType = bindings.NewStructType("TopLevelStruct", fields,
		reflect.TypeOf(TopLevelStruct{}), canonicalFieldMap, nil)
	return structType
}

type MultipleStructReferences struct {
	Field1 TopLevelStruct
	Field2 TopLevelStruct
}

func MultipleStructReferencesBindingType() bindings.BindingType {
	var fields = make(map[string]bindings.BindingType)
	fields["Field1"] = bindings.NewReferenceType(getTopLevelStructBindingType)
	fields["Field2"] = bindings.NewReferenceType(getTopLevelStructBindingType)
	canonicalFieldMap := make(map[string]string)
	canonicalFieldMap["Field1"] = "Field1"
	canonicalFieldMap["Field2"] = "Field2"
	var structType = bindings.NewStructType("MultipleStructReferences", fields,
		reflect.TypeOf(MultipleStructReferences{}), canonicalFieldMap, nil)
	return structType
}

type ReferenceLocalRecursive struct {
	Data string
	NextNode *ReferenceLocalRecursive
}

func ReferenceLocalRecursiveBindingType() bindings.BindingType {
	var fields = make(map[string]bindings.BindingType)
	fields["Data"] = bindings.NewStringType()
	fields["NextNode"] = bindings.NewOptionalType(bindings.NewReferenceType(ReferenceLocalRecursiveBindingType))
	canonicalFieldMap := make(map[string]string)
	canonicalFieldMap["Data"] = "Data"
	canonicalFieldMap["NextNode"] = "NextNode"
	var structType = bindings.NewStructType("ReferenceLocalRecursive", fields,
		reflect.TypeOf(ReferenceLocalRecursive{}), canonicalFieldMap, nil)
	return structType
}

func TestRecursiveStructure(t *testing.T) {
	typeConverter := bindings.NewTypeConverter()
	dataDef, err  := typeConverter.ConvertToDataDefinition(ReferenceLocalRecursiveBindingType())
	assert.Nil(t, err)
	assert.NotNil(t, dataDef)
	secondNode := ReferenceLocalRecursive{Data:"second node", NextNode:nil}
	firstNode := ReferenceLocalRecursive{Data:"first node", NextNode:&secondNode}
	dataVal, err := typeConverter.ConvertToVapi(firstNode, ReferenceLocalRecursiveBindingType())
	assert.Nil(t, err)
	goVal, err := typeConverter.ConvertToGolang(dataVal, ReferenceLocalRecursiveBindingType())
	referenceLocalRecursive := goVal.(ReferenceLocalRecursive)
	assert.Nil(t, err)
	assert.NotNil(t, goVal)
	assert.Equal(t, referenceLocalRecursive, firstNode)
}

func TestMultipleReferenceStructure(t *testing.T) {
	typeConverter := bindings.NewTypeConverter()
	t1 := TopLevelStruct{Dbl:float64(3.14)}
	t2 := TopLevelStruct{Dbl:float64(6.28)}
	m1 := MultipleStructReferences{Field1:t1, Field2:t2}
	dataValue, err := typeConverter.ConvertToVapi(m1, MultipleStructReferencesBindingType())
	assert.Nil(t, err)
	assert.NotNil(t, dataValue)
	goval, err := typeConverter.ConvertToGolang(dataValue, MultipleStructReferencesBindingType())
	assert.Nil(t, err)
	assert.NotNil(t, goval)
	assert.Equal(t, m1, goval)

}


