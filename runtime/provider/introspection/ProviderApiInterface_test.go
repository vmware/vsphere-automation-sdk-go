package introspection_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/provider/introspection"
)

func TestProviderIdentifier(t *testing.T) {
	providerApiInterface, _ := introspection.NewProviderApiInterface("test2", nil)
	assert.Equal(t, providerApiInterface.Identifier().Name(), "com.vmware.vapi.std.introspection.provider")
}
