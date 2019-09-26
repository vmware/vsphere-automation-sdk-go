package introspection_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/provider/introspection"
)

func TestServiceIdentifier(t *testing.T) {
	serviceApiInterface, _ := introspection.NewServiceApiInterface("test3", nil)
	interfaceId := serviceApiInterface.Identifier()
	assert.Equal(t, interfaceId, core.NewInterfaceIdentifier("com.vmware.vapi.std.introspection.service"))
}
