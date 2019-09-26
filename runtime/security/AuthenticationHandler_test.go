package security_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/security"
	"testing"
)

type SamlHokTokenAuthnHandler struct {
}

func NewSamlHokTokenAuthnHandler() *SamlHokTokenAuthnHandler {
	return &SamlHokTokenAuthnHandler{}
}

func (s *SamlHokTokenAuthnHandler) Authenticate(ctx core.SecurityContext) (*security.UserIdentity, error) {
	if ctx.Property(security.AUTHENTICATION_SCHEME_ID).(string) == s.SupportedScheme() {
		if ctx.Property(security.SAML_TOKEN) == "saml-token" {
			domain := "saml-domain"
			uIdentity := security.NewUserIdentity("saml-user")
			uIdentity.SetDomain(&domain)
			return uIdentity, nil
		} else {
			return nil, errors.New("Authentication failed")
		}
	}
	return nil, nil
}

func (s *SamlHokTokenAuthnHandler) SupportedScheme() string {
	return security.SAML_HOK_SCHEME_ID
}

type SamlBearerTokenAuthnHandler struct {
}

func NewSamlBearerTokenAuthnHandler() *SamlBearerTokenAuthnHandler {
	return &SamlBearerTokenAuthnHandler{}
}

func (s *SamlBearerTokenAuthnHandler) Authenticate(ctx core.SecurityContext) (*security.UserIdentity, error) {
	if ctx.Property(security.AUTHENTICATION_SCHEME_ID).(string) == s.SupportedScheme() {
		if ctx.Property(security.SAML_TOKEN) == "saml-bearer-token" {
			domain := "saml-bearer-domain"
			uIdentity := security.NewUserIdentity("saml-bearer-user")
			uIdentity.SetDomain(&domain)
			return uIdentity, nil
		} else {
			return nil, errors.New("Authentication failed")
		}
	}
	return nil, nil
}

func (s *SamlBearerTokenAuthnHandler) SupportedScheme() string {
	return security.SAML_BEARER_SCHEME_ID
}

func TestValidUserPwdHandler(t *testing.T) {
	securityContext := core.NewSecurityContextImpl()
	securityContext.SetProperty(security.AUTHENTICATION_SCHEME_ID, security.USER_PASSWORD_SCHEME_ID)
	securityContext.SetProperty(security.USER_KEY, "testUser")
	securityContext.SetProperty(security.PASSWORD_KEY, "password")
	handler := NewUserPasswordAuthenticationHandler()
	authResult, err := handler.Authenticate(securityContext)
	assert.Nil(t, err)
	assert.Equal(t, authResult.UserName(), "testUser")
	var x *string = nil
	assert.Equal(t, authResult.Domain(), x)
}

func TestInValidUserPwdHandler(t *testing.T) {
	securityContext := core.NewSecurityContextImpl()
	securityContext.SetProperty(security.AUTHENTICATION_SCHEME_ID, security.USER_PASSWORD_SCHEME_ID)
	securityContext.SetProperty(security.USER_KEY, "testUser")
	securityContext.SetProperty(security.PASSWORD_KEY, "pwd")
	handler := NewUserPasswordAuthenticationHandler()
	authResult, err := handler.Authenticate(securityContext)
	assert.NotNil(t, err)
	assert.Nil(t, authResult)
}

func TestValidSamlHokTokenAuthHandler(t *testing.T) {
	securityContext := core.NewSecurityContextImpl()
	securityContext.SetProperty(security.AUTHENTICATION_SCHEME_ID, security.SAML_HOK_SCHEME_ID)
	securityContext.SetProperty(security.SAML_TOKEN, "saml-token")
	securityContext.SetProperty(security.PRIVATE_KEY, "key")
	handler := NewSamlHokTokenAuthnHandler()
	result, err := handler.Authenticate(securityContext)
	assert.Nil(t, err)
	assert.Equal(t, result.UserName(), "saml-user")
	assert.Equal(t, *(result.Domain()), "saml-domain")

}

func TestInValidSamlHokTokenAuthHandler(t *testing.T) {
	securityContext := core.NewSecurityContextImpl()
	securityContext.SetProperty(security.AUTHENTICATION_SCHEME_ID, security.SAML_HOK_SCHEME_ID)
	securityContext.SetProperty(security.SAML_TOKEN, "invalid-saml-token")
	securityContext.SetProperty(security.PRIVATE_KEY, "key")
	handler := NewSamlHokTokenAuthnHandler()
	result, err := handler.Authenticate(securityContext)
	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func TestSamlHokInvalidScheme(t *testing.T) {
	securityContext := core.NewSecurityContextImpl()
	securityContext.SetProperty(security.AUTHENTICATION_SCHEME_ID, security.SESSION_SCHEME_ID)
	securityContext.SetProperty(security.SAML_TOKEN, "saml-token")
	securityContext.SetProperty(security.PRIVATE_KEY, "key")
	handler := NewSamlHokTokenAuthnHandler()
	result, err := handler.Authenticate(securityContext)
	assert.Nil(t, err)
	assert.Nil(t, result)
}

func TestValidSamlBearerScheme(t *testing.T) {
	securityContext := core.NewSecurityContextImpl()
	securityContext.SetProperty(security.AUTHENTICATION_SCHEME_ID, security.SAML_BEARER_SCHEME_ID)
	securityContext.SetProperty(security.SAML_TOKEN, "saml-bearer-token")
	handler := NewSamlBearerTokenAuthnHandler()
	result, err := handler.Authenticate(securityContext)
	assert.Nil(t, err)
	assert.Equal(t, result.UserName(), "saml-bearer-user")
	assert.Equal(t, *result.Domain(), "saml-bearer-domain")
}

func TestInValidSamlBearer(t *testing.T) {
	securityContext := core.NewSecurityContextImpl()
	securityContext.SetProperty(security.AUTHENTICATION_SCHEME_ID, security.SAML_BEARER_SCHEME_ID)
	securityContext.SetProperty(security.SAML_TOKEN, "invalid-saml-bearer-token")
	handler := NewSamlBearerTokenAuthnHandler()
	result, err := handler.Authenticate(securityContext)
	assert.NotNil(t, err)
	assert.Nil(t, result)

}

func TestInValidSamlBearerScheme(t *testing.T) {
	securityContext := core.NewSecurityContextImpl()
	securityContext.SetProperty(security.AUTHENTICATION_SCHEME_ID, security.SESSION_SCHEME_ID)
	securityContext.SetProperty(security.SAML_TOKEN, "saml-bearer-token")
	handler := NewSamlBearerTokenAuthnHandler()
	result, err := handler.Authenticate(securityContext)
	assert.Nil(t, err)
	assert.Nil(t, result)

}

func TestSessionScheme(t *testing.T) {
	securityContext := core.NewSecurityContextImpl()
	securityContext.SetProperty(security.AUTHENTICATION_SCHEME_ID, security.SESSION_SCHEME_ID)
	securityContext.SetProperty(security.SESSION_ID, "session-id")
	handler := NewSessionAuthnHandler()
	result, err := handler.Authenticate(securityContext)
	assert.Nil(t, err)
	assert.Equal(t, result.UserName(), "session-user")
	var x *string = nil
	assert.Equal(t, result.Domain(), x)
}

func TestSessionInvalid(t *testing.T) {
	securityContext := core.NewSecurityContextImpl()
	securityContext.SetProperty(security.AUTHENTICATION_SCHEME_ID, security.SESSION_SCHEME_ID)
	securityContext.SetProperty(security.SESSION_ID, "invalid-session-id")
	handler := NewSessionAuthnHandler()
	result, err := handler.Authenticate(securityContext)
	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func TestSessionInvalidScheme(t *testing.T) {
	securityContext := core.NewSecurityContextImpl()
	securityContext.SetProperty(security.AUTHENTICATION_SCHEME_ID, security.SAML_BEARER_SCHEME_ID)
	securityContext.SetProperty(security.SESSION_ID, "invalid-session-id")
	handler := NewSessionAuthnHandler()
	result, err := handler.Authenticate(securityContext)
	assert.Nil(t, err)
	assert.Nil(t, result)
}
