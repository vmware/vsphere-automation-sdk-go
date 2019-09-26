package bindings_test

import (
	"github.com/stretchr/testify/assert"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"testing"
)

func Test_CreateErrorValueFromMessages(t *testing.T) {
	errorValue := bindings.CreateErrorValueFromMessages(bindings.NOT_FOUND_ERROR_DEF, []error{})
	assert.NotNil(t, errorValue)
	dv, dvErr := errorValue.Field("error_type")
	assert.Nil(t, dvErr)
	assert.Equal(t, dv.(*data.OptionalValue).Value().(*data.StringValue).Value(), "NOT_FOUND")
}
