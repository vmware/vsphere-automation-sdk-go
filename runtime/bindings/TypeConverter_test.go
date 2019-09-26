package bindings

import (
	"github.com/stretchr/testify/assert"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"testing"
)

func TestTypeConverter_ConvertToVapi(t *testing.T) {
	typeConverter := NewTypeConverter()
	var x = 10
	var bindingType BindingType = nil
	_, err := typeConverter.ConvertToVapi(x, bindingType)

	assert.NotNil(t, err)
	assert.Equal(t, 1, len(err))
	assert.Equal(t, "Binding type cannot be nil", err[0].Error())
}

func TestTypeConverter_ConvertToDataDefinition(t *testing.T) {
	typeConverter := NewTypeConverter()
	var bindingType BindingType = nil
	_, err := typeConverter.ConvertToDataDefinition(bindingType)

	assert.NotNil(t, err)
	assert.Equal(t, 1, len(err))
	assert.Equal(t, "Binding type cannot be nil", err[0].Error())
}

func TestTypeConverter_ConvertToGolang(t *testing.T) {
	typeConverter := NewTypeConverter()
	var bindingType BindingType = nil
	var dataValue data.DataValue = data.NewStringValue("x")
	_, err := typeConverter.ConvertToGolang(dataValue, bindingType)

	assert.NotNil(t, err)
	assert.Equal(t, 1, len(err))
	assert.Equal(t, "Binding type cannot be nil", err[0].Error())
}
