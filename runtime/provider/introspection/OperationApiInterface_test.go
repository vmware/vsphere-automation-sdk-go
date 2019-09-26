package introspection_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/provider/introspection"
)

func TestOperationIdentifier(t *testing.T) {
	operationApiInterface, _ := introspection.NewOperationApiInterface("test1", nil)
	assert.Equal(t, operationApiInterface.Identifier().Name(), "com.vmware.vapi.std.introspection.operation")
}
