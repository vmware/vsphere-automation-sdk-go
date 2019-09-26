package security_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/security"
)

func TestNewSessionSecurityContext(t *testing.T) {
	var securityContext core.SecurityContext = security.NewSessionSecurityContext("xyz")
	assert.NotNil(t, securityContext)
	assert.Equal(t, securityContext.Property("sessionId"), "xyz")
	assert.Equal(t, securityContext.Property("schemeId"), "com.vmware.vapi.std.security.session_id")
}

func TestNewSessionSecurityContextMarshalJSON(t *testing.T) {
	var securityContext = security.NewSessionSecurityContext("xyz")
	var json, _ = securityContext.MarshalJSON()
	assert.Equal(t, string(json), "{\"schemeId\":\"com.vmware.vapi.std.security.session_id\",\"sessionId\":\"xyz\"}")
}

func TestNewSessionSecurityContextGetAllProperties(t *testing.T) {
	var securityContext = security.NewSessionSecurityContext("xyz")
	securityContext.SetProperty("p1", "value1")
	assert.Equal(t, securityContext.GetAllProperties()["p1"], "value1")
}

func TestNewUserPasswordSecurityContext(t *testing.T) {
	var userPassContext core.SecurityContext = security.NewUserPasswordSecurityContext("administrator@vsphere.local", "password")
	assert.NotNil(t, userPassContext)
	assert.Equal(t, userPassContext.Property("userName"), "administrator@vsphere.local")
	assert.Equal(t, userPassContext.Property("password"), "password")
	assert.Equal(t, userPassContext.Property("schemeId"), "com.vmware.vapi.std.security.user_pass")
}

func TestNewUserPasswordSecurityContextUser(t *testing.T) {
	var userPassContext = security.NewUserPasswordSecurityContext("administrator@vsphere.local", "password")
	assert.Equal(t, userPassContext.User(), "administrator@vsphere.local")
}

func TestNewUserPasswordSecurityContextPassword(t *testing.T) {
	var userPassContext = security.NewUserPasswordSecurityContext("administrator@vsphere.local", "password")
	assert.Equal(t, userPassContext.Password(), "password")
}

func TestNewUserPasswordSecurityContextMarshalJSON(t *testing.T) {
	var userPassContext = security.NewUserPasswordSecurityContext("administrator@vsphere.local", "password")
	var json, _ = userPassContext.MarshalJSON()
	assert.Equal(t, string(json), "{\"password\":\"password\",\"schemeId\":\"com.vmware.vapi.std.security.user_pass\",\"userName\":\"administrator@vsphere.local\"}")
}

func TestNewUserPasswordSecurityContextGetAllProperties(t *testing.T) {
	var userPassContext = security.NewUserPasswordSecurityContext("administrator@vsphere.local", "password")
	userPassContext.SetProperty("p1", "value1")
	assert.Equal(t, userPassContext.GetAllProperties()["p1"], "value1")
}

func TestNewOauthSecurityContext(t *testing.T) {
	securityContext := security.NewOauthSecurityContext("xyz")
	assert.NotNil(t, securityContext)
	assert.Equal(t, securityContext.Property(security.ACCESS_TOKEN), "xyz")
	assert.Equal(t, securityContext.Property(security.AUTHENTICATION_SCHEME_ID), security.OAUTH_SCHEME_ID)
}

func TestNewOauthSecurityContextMarshalJSON(t *testing.T) {
	var securityContext = security.NewOauthSecurityContext("xyz")
	var json, _ = securityContext.MarshalJSON()
	assert.Equal(t, "{\"accessToken\":\"xyz\",\"schemeId\":\"com.vmware.vapi.std.security.oauth\"}", string(json))
}

func TestNewOauthSecurityContextGetAllProperties(t *testing.T) {
	var securityContext = security.NewOauthSecurityContext("xyz")
	securityContext.SetProperty("p1", "value1")
	assert.Equal(t, securityContext.GetAllProperties()["p1"], "value1")
}
