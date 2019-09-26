package common

import (
	"github.com/stretchr/testify/assert"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/lib"
	"testing"
)

func TestCreateOpId(t *testing.T) {
	opId := NewOpId()
	assert.NotNil(t, opId)
}

func TestNewDefaultApplicationContext(t *testing.T) {
	appContext := NewDefaultApplicationContext()
	assert.NotNil(t, appContext)
	assert.NotNil(t, appContext.GetProperty(lib.OPID))
}
