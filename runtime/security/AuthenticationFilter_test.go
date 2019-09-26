package security_test

import (
	"github.com/stretchr/testify/assert"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/common"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/security"
	"testing"
)

type MockupApiInterface struct {
	iFaceId  core.InterfaceIdentifier
	methodId core.MethodIdentifier
}

func (m *MockupApiInterface) Identifier() core.InterfaceIdentifier {
	return m.iFaceId
}

func (m *MockupApiInterface) Definition() core.InterfaceDefinition {
	return core.NewInterfaceDefinition(m.iFaceId, []core.MethodIdentifier{m.methodId})
}

func (m *MockupApiInterface) MethodDefinition(methodId core.MethodIdentifier) *core.MethodDefinition {
	input := data.NewStructDefinition("mock", nil)
	output := data.NewIntegerDefinition()
	md := core.NewMethodDefinition(m.methodId, input, output, []data.ErrorDefinition{})
	return &md

}

func (m *MockupApiInterface) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, input data.DataValue) core.MethodResult {
	return core.NewMethodResult(data.NewIntegerValue(10), nil)
}

type MockProvider struct {
}

func NewMockProvider() *MockProvider {
	return &MockProvider{}
}

func (m *MockProvider) Invoke(serviceId string, operationId string, inputValue data.DataValue, ctx *core.ExecutionContext) core.MethodResult {
	return core.NewMethodResult(data.NewIntegerValue(10), nil)
}

func TestUserPasswordScheme(t *testing.T) {
	authnfiles := []string{"test/resources/authn.json"}
	authFilter, err := security.NewAuthenticationFilter([]security.AuthenticationHandler{NewUserPasswordAuthenticationHandler()}, NewMockProvider(), authnfiles)
	securityCtx := core.NewSecurityContextImpl()
	securityCtx.SetProperty(security.AUTHENTICATION_SCHEME_ID, security.USER_PASSWORD_SCHEME_ID)
	securityCtx.SetProperty(security.USER_KEY, "testUser")
	securityCtx.SetProperty(security.PASSWORD_KEY, "password")
	appCtx := common.NewDefaultApplicationContext()
	ctx := core.NewExecutionContext(appCtx, securityCtx)
	inputVal := data.NewVoidValue()
	methodResult := authFilter.Invoke("com.pkg.svc", "op", inputVal, ctx)
	assert.Nil(t, err)
	assert.Equal(t, methodResult.Output(), data.NewIntegerValue(10))
	assert.Equal(t, ctx.SecurityContext().Property(security.AUTHN_IDENTITY), security.NewUserIdentity("testUser"))
}

func TestUserPasswordInvalid(t *testing.T) {
	authnfiles := []string{"test/resources/authn.json"}
	authFilter, err := security.NewAuthenticationFilter([]security.AuthenticationHandler{NewUserPasswordAuthenticationHandler()}, NewMockProvider(), authnfiles)
	securityCtx := core.NewSecurityContextImpl()
	securityCtx.SetProperty(security.AUTHENTICATION_SCHEME_ID, security.USER_PASSWORD_SCHEME_ID)
	securityCtx.SetProperty(security.USER_KEY, "testUser")
	securityCtx.SetProperty(security.PASSWORD_KEY, "passwordinvalid")
	appCtx := common.NewDefaultApplicationContext()
	ctx := core.NewExecutionContext(appCtx, securityCtx)
	inputVal := data.NewVoidValue()
	methodResult := authFilter.Invoke("com.pkg.svc", "op", inputVal, ctx)
	assert.Nil(t, err)
	assert.Equal(t, methodResult.Error().Name(), "com.vmware.vapi.std.errors.unauthenticated")
}

func TestInvalidScheme(t *testing.T) {
	authnfiles := []string{"test/resources/authn.json"}
	authFilter, err := security.NewAuthenticationFilter([]security.AuthenticationHandler{NewUserPasswordAuthenticationHandler()}, NewMockProvider(), authnfiles)
	securityCtx := core.NewSecurityContextImpl()
	securityCtx.SetProperty(security.AUTHENTICATION_SCHEME_ID, security.SAML_HOK_SCHEME_ID)
	securityCtx.SetProperty(security.USER_KEY, "testUser")
	securityCtx.SetProperty(security.PASSWORD_KEY, "passwordinvalid")
	appCtx := common.NewDefaultApplicationContext()
	ctx := core.NewExecutionContext(appCtx, securityCtx)
	inputVal := data.NewVoidValue()
	methodResult := authFilter.Invoke("com.pkg.svc", "op", inputVal, ctx)
	assert.Nil(t, err)
	assert.Equal(t, methodResult.Output(), data.NewIntegerValue(10))
	assert.Equal(t, ctx.SecurityContext().Property(security.AUTHN_IDENTITY), nil)
	var errVal *data.ErrorValue = nil
	assert.Equal(t, methodResult.Error(), errVal)
}

//second handler is the correct handler.
func TestMultipleSchemes(t *testing.T) {
	authnfiles := []string{"test/resources/authn.json"}
	authFilter, err := security.NewAuthenticationFilter([]security.AuthenticationHandler{NewUserPasswordAuthenticationHandler(), NewSessionAuthnHandler()}, NewMockProvider(), authnfiles)
	securityCtx := core.NewSecurityContextImpl()
	securityCtx.SetProperty(security.AUTHENTICATION_SCHEME_ID, security.SESSION_SCHEME_ID)
	securityCtx.SetProperty(security.SESSION_ID, "session-id")
	appCtx := common.NewDefaultApplicationContext()
	ctx := core.NewExecutionContext(appCtx, securityCtx)
	inputVal := data.NewVoidValue()
	methodResult := authFilter.Invoke("com.pkg.svc", "op2", inputVal, ctx)
	assert.Nil(t, err)
	assert.Equal(t, methodResult.Output(), data.NewIntegerValue(10))
	assert.Equal(t, ctx.SecurityContext().Property(security.AUTHN_IDENTITY), security.NewUserIdentity("session-user"))
	var errVal *data.ErrorValue = nil
	assert.Equal(t, methodResult.Error(), errVal)
}

func TestNoMatchingHandler(t *testing.T) {
	authnfiles := []string{"test/resources/authn.json"}
	authFilter, err := security.NewAuthenticationFilter([]security.AuthenticationHandler{NewUserPasswordAuthenticationHandler()}, NewMockProvider(), authnfiles)
	securityCtx := core.NewSecurityContextImpl()
	securityCtx.SetProperty(security.AUTHENTICATION_SCHEME_ID, security.SESSION_SCHEME_ID)
	securityCtx.SetProperty(security.SESSION_ID, "session-id")
	appCtx := common.NewDefaultApplicationContext()
	ctx := core.NewExecutionContext(appCtx, securityCtx)
	inputVal := data.NewVoidValue()
	methodResult := authFilter.Invoke("com.pkg.svc", "op2", inputVal, ctx)
	assert.Nil(t, err)
	assert.Equal(t, methodResult.Output(), data.NewIntegerValue(10))
	assert.Equal(t, ctx.SecurityContext().Property(security.AUTHN_IDENTITY), nil)
	var errVal *data.ErrorValue = nil
	assert.Equal(t, methodResult.Error(), errVal)
}

func TestWithUnsetScheme(t *testing.T) {
	authnfiles := []string{"test/resources/authn.json"}
	authFilter, err := security.NewAuthenticationFilter([]security.AuthenticationHandler{NewUserPasswordAuthenticationHandler()}, NewMockProvider(), authnfiles)
	securityCtx := core.NewSecurityContextImpl()
	appCtx := common.NewDefaultApplicationContext()
	ctx := core.NewExecutionContext(appCtx, securityCtx)
	inputVal := data.NewVoidValue()
	methodResult := authFilter.Invoke("com.pkg.svc", "op", inputVal, ctx)
	assert.Nil(t, err)
	assert.Equal(t, methodResult.Output(), data.NewIntegerValue(10))
	var errVal *data.ErrorValue = nil
	assert.Equal(t, methodResult.Error(), errVal)
}

func TestWithInvalidScheme(t *testing.T) {
	authnfiles := []string{"test/resources/authn.json"}
	authFilter, err := security.NewAuthenticationFilter([]security.AuthenticationHandler{NewUserPasswordAuthenticationHandler()}, NewMockProvider(), authnfiles)
	securityCtx := core.NewSecurityContextImpl()
	securityCtx.SetProperty(security.AUTHENTICATION_SCHEME_ID, 12)
	appCtx := common.NewDefaultApplicationContext()
	ctx := core.NewExecutionContext(appCtx, securityCtx)
	inputVal := data.NewVoidValue()
	methodResult := authFilter.Invoke("com.pkg.svc", "op", inputVal, ctx)
	assert.Nil(t, err)
	assert.Equal(t, methodResult.Output(), nil)
	assert.Equal(t, methodResult.Error().Name(), "com.vmware.vapi.std.errors.invalid_request")
}
