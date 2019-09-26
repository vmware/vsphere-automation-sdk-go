package security_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/security"
)

func TestUserIdentityGroups(t *testing.T) {
	uIdentity := security.NewUserIdentity("saml-user")
	groups := []string{"a", "b", "c"}
	uIdentity.SetGroups(groups)
	assert.Equal(t, uIdentity.Groups(), groups)
}
