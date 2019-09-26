package security_test

import (
	"github.com/stretchr/testify/assert"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/security"
	"testing"
)

const VALID_SAML_DOMAIN = "valid-saml-domain"

type UserPwdAuthzHandler struct {
}

func NewUserPwdAuthzHandler() *UserPwdAuthzHandler {
	return &UserPwdAuthzHandler{}
}

func (u *UserPwdAuthzHandler) Authorize(serviceId string, operationId string, secCtx core.SecurityContext) (bool, error) {
	userId := secCtx.Property(security.AUTHN_IDENTITY)
	return userId.(*security.UserIdentity).UserName() == "validuser", nil
}

type SamlTokenAuthzHandler struct {
}

func NewSamlTokenAuthzHandler() *SamlTokenAuthzHandler {
	return &SamlTokenAuthzHandler{}
}

func (s *SamlTokenAuthzHandler) Authorize(serviceId string, operationId string, secCtx core.SecurityContext) (bool, error) {
	userId := secCtx.Property(security.AUTHN_IDENTITY)
	return userId.(*security.UserIdentity).UserName() == "valid-saml-user" && *userId.(*security.UserIdentity).Domain() == VALID_SAML_DOMAIN, nil
}

type SessionAuthzHandler struct {
}

func NewSessionAuthzHandler() *SessionAuthzHandler {
	return &SessionAuthzHandler{}
}

func (s *SessionAuthzHandler) Authorize(serviceId string, operationId string, secCtx core.SecurityContext) (bool, error) {
	userId := secCtx.Property(security.AUTHN_IDENTITY)
	return userId.(*security.UserIdentity).UserName() == "valid-session-user", nil
}

func TestValidUserPwdAuthzHandler(t *testing.T) {
	securityContext := core.NewSecurityContextImpl()
	securityContext.SetProperty(security.AUTHENTICATION_SCHEME_ID, security.USER_PASSWORD_SCHEME_ID)
	securityContext.SetProperty(security.USER_KEY, "validuser")
	securityContext.SetProperty(security.PASSWORD_KEY, "password")
	securityContext.SetProperty(security.AUTHN_IDENTITY, security.NewUserIdentity("validuser"))
	authzHandler := NewUserPwdAuthzHandler()
	result, err := authzHandler.Authorize("com.vmware.pkg", "op", securityContext)
	assert.True(t, result)
	assert.Nil(t, err)
}

func TestUnAuthorizedUserPwdAuthzHandler(t *testing.T) {
	securityContext := core.NewSecurityContextImpl()
	securityContext.SetProperty(security.AUTHENTICATION_SCHEME_ID, security.USER_PASSWORD_SCHEME_ID)
	securityContext.SetProperty(security.USER_KEY, "unauthorizeduser")
	securityContext.SetProperty(security.PASSWORD_KEY, "password")
	securityContext.SetProperty(security.AUTHN_IDENTITY, security.NewUserIdentity("unauthorizeduser"))
	authzHandler := NewUserPwdAuthzHandler()
	result, err := authzHandler.Authorize("com.vmware.pkg", "op", securityContext)
	assert.False(t, result)
	assert.Nil(t, err)
}

func TestValidSamlHokAuthzHandler(t *testing.T) {
	securityContext := core.NewSecurityContextImpl()
	securityContext.SetProperty(security.AUTHENTICATION_SCHEME_ID, security.SAML_HOK_SCHEME_ID)
	securityContext.SetProperty(security.PRIVATE_KEY, "key")
	securityContext.SetProperty(security.SAML_TOKEN, "valid-saml-token")
	domain := VALID_SAML_DOMAIN
	uI := security.NewUserIdentity("valid-saml-user")
	uI.SetDomain(&domain)
	securityContext.SetProperty(security.AUTHN_IDENTITY, uI)
	authzHandler := NewSamlTokenAuthzHandler()
	result, err := authzHandler.Authorize("com.vmware.pkg", "op", securityContext)
	assert.True(t, result)
	assert.Nil(t, err)
}

func TestInValidUserSamlHokAuthzHandler(t *testing.T) {
	securityContext := core.NewSecurityContextImpl()
	securityContext.SetProperty(security.AUTHENTICATION_SCHEME_ID, security.SAML_HOK_SCHEME_ID)
	securityContext.SetProperty(security.PRIVATE_KEY, "key")
	securityContext.SetProperty(security.SAML_TOKEN, "invalid-saml-token")
	domain := VALID_SAML_DOMAIN
	uI := security.NewUserIdentity("invalid-saml-user")
	uI.SetDomain(&domain)
	securityContext.SetProperty(security.AUTHN_IDENTITY, uI)
	authzHandler := NewSamlTokenAuthzHandler()
	result, err := authzHandler.Authorize("com.vmware.pkg", "op", securityContext)
	assert.False(t, result)
	assert.Nil(t, err)
}

func TestInValidDomainSamlHokAuthzHandler(t *testing.T) {
	securityContext := core.NewSecurityContextImpl()
	securityContext.SetProperty(security.AUTHENTICATION_SCHEME_ID, security.SAML_HOK_SCHEME_ID)
	securityContext.SetProperty(security.PRIVATE_KEY, "key")
	securityContext.SetProperty(security.SAML_TOKEN, "valid-saml-token")
	domain := "invalid-saml-domain"
	uI := security.NewUserIdentity("valid-saml-user")
	uI.SetDomain(&domain)
	securityContext.SetProperty(security.AUTHN_IDENTITY, uI)
	authzHandler := NewSamlTokenAuthzHandler()
	result, err := authzHandler.Authorize("com.vmware.pkg", "op", securityContext)
	assert.False(t, result)
	assert.Nil(t, err)
}

func TestValidSamlBearerHandler(t *testing.T) {
	securityContext := core.NewSecurityContextImpl()
	securityContext.SetProperty(security.AUTHENTICATION_SCHEME_ID, security.SAML_HOK_SCHEME_ID)
	securityContext.SetProperty(security.SAML_TOKEN, "valid-saml-token")
	domain := VALID_SAML_DOMAIN
	uI := security.NewUserIdentity("valid-saml-user")
	uI.SetDomain(&domain)
	securityContext.SetProperty(security.AUTHN_IDENTITY, uI)
	authzHandler := NewSamlTokenAuthzHandler()
	result, err := authzHandler.Authorize("com.vmware.pkg", "op", securityContext)
	assert.True(t, result)
	assert.Nil(t, err)
}

func TestInValidUserSamlBearerHandler(t *testing.T) {
	securityContext := core.NewSecurityContextImpl()
	securityContext.SetProperty(security.AUTHENTICATION_SCHEME_ID, security.SAML_HOK_SCHEME_ID)
	securityContext.SetProperty(security.SAML_TOKEN, "invalid-saml-token")
	domain := VALID_SAML_DOMAIN
	uI := security.NewUserIdentity("invalid-saml-user")
	uI.SetDomain(&domain)
	securityContext.SetProperty(security.AUTHN_IDENTITY, uI)
	authzHandler := NewSamlTokenAuthzHandler()
	result, err := authzHandler.Authorize("com.vmware.pkg", "op", securityContext)
	assert.False(t, result)
	assert.Nil(t, err)
}

func TestInValidDomainSamlBearerHandler(t *testing.T) {
	securityContext := core.NewSecurityContextImpl()
	securityContext.SetProperty(security.AUTHENTICATION_SCHEME_ID, security.SAML_HOK_SCHEME_ID)
	securityContext.SetProperty(security.SAML_TOKEN, "valid-saml-token")
	domain := "invalid-saml-domain"
	uI := security.NewUserIdentity("valid-saml-user")
	uI.SetDomain(&domain)
	securityContext.SetProperty(security.AUTHN_IDENTITY, uI)
	authzHandler := NewSamlTokenAuthzHandler()
	result, err := authzHandler.Authorize("com.vmware.pkg", "op", securityContext)
	assert.False(t, result)
	assert.Nil(t, err)
}

func TestValidSessionHandler(t *testing.T) {
	securityContext := core.NewSecurityContextImpl()
	securityContext.SetProperty(security.AUTHENTICATION_SCHEME_ID, security.SESSION_SCHEME_ID)
	securityContext.SetProperty(security.SESSION_ID, "session-id")
	uI := security.NewUserIdentity("valid-session-user")

	securityContext.SetProperty(security.AUTHN_IDENTITY, uI)
	authzHandler := NewSessionAuthzHandler()
	result, err := authzHandler.Authorize("com.vmware.pkg", "op", securityContext)
	assert.True(t, result)
	assert.Nil(t, err)
}

func TestInValidSessionHandler(t *testing.T) {
	securityContext := core.NewSecurityContextImpl()
	securityContext.SetProperty(security.AUTHENTICATION_SCHEME_ID, security.SESSION_SCHEME_ID)
	securityContext.SetProperty(security.SESSION_ID, "invalid-session-id")
	securityContext.SetProperty(security.AUTHN_IDENTITY, security.NewUserIdentity("invalid-session-user"))
	authzHandler := NewSessionAuthzHandler()
	result, err := authzHandler.Authorize("com.vmware.pkg", "op", securityContext)
	assert.False(t, result)
	assert.Nil(t, err)
}
