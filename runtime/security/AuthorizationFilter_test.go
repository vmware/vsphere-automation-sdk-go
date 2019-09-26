package security_test

import (
	"github.com/stretchr/testify/assert"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/common"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/l10n"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/log"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/metadata"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/security"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/security/test"
	"testing"
)

func setup(privileges []string) *security.AuthorizationFilter {

	authzfiles := []string{"test/resources/vmodl.test.uber.annotations.lifecycle.filtering.internal_metadata_authz.json"}

	var authzRoleProvider test.AuthorizationProvider
	var privilegeProv security.PrivilegeProvider
	var err error
	var apimeta *metadata.ApiMetadata

	authzRoleProvider = test.NewAuthorizationPrivilegeProvider()
	authzRoleProvider.SetPrivileges("validuser", privileges)

	apimeta, err = metadata.NewApiMetadata(authzfiles, authzfiles)
	if err != nil {
		log.Error(err)
	}

	privilegeProv = security.NewPrivilegeProviderImpl(apimeta)

	var permissionValid security.PermissionValidator
	permissionValid = test.NewPermissionValidatorImpl(authzRoleProvider)

	authzFilter, _ := security.NewAuthorizationFilter(NewMockProvider(), privilegeProv, permissionValid)
	authzFilter.AddHandler(NewUserPwdAuthzHandler())

	return authzFilter
}

func createUserPassSecurityCtx(userKey string, password string) *core.SecurityContextImpl {
	securityCtx := core.NewSecurityContextImpl()
	securityCtx.SetProperty(security.AUTHENTICATION_SCHEME_ID, security.USER_PASSWORD_SCHEME_ID)
	securityCtx.SetProperty(security.AUTHN_IDENTITY, security.NewUserIdentity(userKey))
	securityCtx.SetProperty(security.USER_KEY, userKey)
	securityCtx.SetProperty(security.PASSWORD_KEY, password)
	return securityCtx
}

func TestValidUser(t *testing.T) {

	securityCtx := createUserPassSecurityCtx("validuser", "password")
	execCtx := core.NewExecutionContext(common.NewDefaultApplicationContext(), securityCtx)
	privileges := []string{"vapi.id.parameter", "dummy.vapi.id.parameter", "System.Read"}

	authzFilter := setup(privileges)

	dataVal := data.NewStructValue("vmodl.test.uber.privileges.uber_privileges.operation_with_id_parameter", nil)
	dataVal.SetField("id", data.NewStringValue("id"))

	methodResult := authzFilter.Invoke("vmodl.test.uber.privileges.uber_privileges", "operation_with_id_parameter", dataVal, execCtx)
	assert.Nil(t, methodResult.Error())
	assert.Equal(t, data.NewIntegerValue(10), methodResult.Output())
}

func TestValidUserOnlyOperationPrivilege(t *testing.T) {
	securityCtx := createUserPassSecurityCtx("validuser", "password")
	execCtx := core.NewExecutionContext(common.NewDefaultApplicationContext(), securityCtx)
	privileges := []string{"vapi.id.parameter", "dummy.vapi.id.parameter", "System.Read"}

	authzFilter := setup(privileges)

	methodResult := authzFilter.Invoke("vmodl.test.uber.privileges.uber_privileges", "operation_with_id_parameter", data.NewVoidValue(), execCtx)
	assert.Nil(t, methodResult.Error())
	assert.Equal(t, data.NewIntegerValue(10), methodResult.Output())
}

func TestValidUserOnlyOperationPrivilege2(t *testing.T) {
	securityCtx := createUserPassSecurityCtx("validuser", "password")
	execCtx := core.NewExecutionContext(common.NewDefaultApplicationContext(), securityCtx)
	privileges := []string{"dummy.vapi.id.parameter", "System.Read"}

	authzFilter := setup(privileges)

	methodResult := authzFilter.Invoke("vmodl.test.uber.privileges.uber_privileges", "operation_with_id_parameter", nil, execCtx)
	assert.Nil(t, methodResult.Error())
	assert.Equal(t, data.NewIntegerValue(10), methodResult.Output())
}

func TestAuthzNotPermitted(t *testing.T) {

	securityCtx := createUserPassSecurityCtx("validuser", "password")
	execCtx := core.NewExecutionContext(common.NewDefaultApplicationContext(), securityCtx)
	privileges := []string{"System.Write", "System.Read"}

	authzFilter := setup(privileges)

	methodResult := authzFilter.Invoke("vmodl.test.uber.privileges.uber_privileges", "operation_with_id_parameter", data.NewVoidValue(), execCtx)
	assert.Equal(t, methodResult.Error().Name(), "com.vmware.vapi.std.errors.unauthorized")
	assert.Equal(t, methodResult.Output(), nil)
}

func TestValidUserLevel2(t *testing.T) {
	securityCtx := createUserPassSecurityCtx("validuser", "password")
	execCtx := core.NewExecutionContext(common.NewDefaultApplicationContext(), securityCtx)
	privileges := []string{"vapi.id.field", "System.Read"}

	authzFilter := setup(privileges)

	dataVal := data.NewStructValue("vmodl.test.uber.privileges.uber_privileges.get_spec", nil)
	dataVal.SetField("id", data.NewStringValue("id"))

	methodResult := authzFilter.Invoke("vmodl.test.uber.privileges.uber_privileges", "operation_with_struct_parameter1", dataVal, execCtx)
	assert.Nil(t, methodResult.Error())
	assert.Equal(t, data.NewIntegerValue(10), methodResult.Output())
}

func TestValidUserWithMap(t *testing.T) {
	securityCtx := createUserPassSecurityCtx("validuser", "password")
	execCtx := core.NewExecutionContext(common.NewDefaultApplicationContext(), securityCtx)
	privileges := []string{"vapi.id.map.key.field", "System.Read"}

	authzFilter := setup(privileges)

	dataVal := data.NewStructValue("vmodl.test.uber.privileges.uber_privileges.map_spec-map_field#key", nil)
	dataVal.SetField("key", data.NewStringValue("map_field-key"))

	methodResult := authzFilter.Invoke("vmodl.test.uber.privileges.uber_privileges", "operation_with_struct_parameter5", dataVal, execCtx)
	assert.Nil(t, methodResult.Error())
	assert.Equal(t, data.NewIntegerValue(10), methodResult.Output())
}

func TestValidUserLevelNested(t *testing.T) {
	securityCtx := createUserPassSecurityCtx("validuser", "password")
	execCtx := core.NewExecutionContext(common.NewDefaultApplicationContext(), securityCtx)
	privileges := []string{"vapi.id.field", "System.Read", "vapi.id.parameter"}

	authzFilter := setup(privileges)

	dataVal := data.NewStructValue("vmodl.test.uber.privileges.uber_privileges.operation_mix", nil)

	subDataVal := data.NewStructValue("vmodl.test.uber.privileges.uber_privileges.get_spec", nil)
	subDataVal.SetField("id", data.NewStringValue("get_spec.id"))

	dataVal.SetField("spec", subDataVal)
	dataVal.SetField("id", data.NewStringValue("operation_mix.id"))

	methodResult := authzFilter.Invoke("vmodl.test.uber.privileges.uber_privileges", "operation_mix", dataVal, execCtx)
	assert.Nil(t, methodResult.Error())
	assert.Equal(t, data.NewIntegerValue(10), methodResult.Output())
}

func TestValidUserNoAuthzForOperation(t *testing.T) {
	securityCtx := createUserPassSecurityCtx("validuser", "password")
	execCtx := core.NewExecutionContext(common.NewDefaultApplicationContext(), securityCtx)
	privileges := []string{}

	authzFilter := setup(privileges)

	methodResult := authzFilter.Invoke("vmodl.test.uber.annotations.crud.uber_CRUD", "list", nil, execCtx)
	assert.Nil(t, methodResult.Error())
	assert.Equal(t, data.NewIntegerValue(10), methodResult.Output())
}

func TestNoauthOp(t *testing.T) {

	execCtx := core.NewExecutionContext(common.NewDefaultApplicationContext(), nil)
	privileges := []string{"vapi.id.parameter", "dummy.vapi.id.parameter", "System.Read"}

	authzFilter := setup(privileges)

	methodResult := authzFilter.Invoke("vmodl.test.uber.privileges.uber_privileges", "simple_operation", data.NewVoidValue(), execCtx)
	assert.Nil(t, methodResult.Error())
	assert.Equal(t, methodResult.Output(), data.NewIntegerValue(10))
}

func TestNoAuthOpUnknownOp(t *testing.T) {
	execCtx := core.NewExecutionContext(common.NewDefaultApplicationContext(), nil)
	privileges := []string{}

	authzFilter := setup(privileges)
	methodResult := authzFilter.Invoke("com", "op", data.NewVoidValue(), execCtx)
	assert.Nil(t, methodResult.Error())
	assert.Equal(t, methodResult.Output(), data.NewIntegerValue(10))
}

func TestUnknownOpWithValidUser(t *testing.T) {
	securityCtx := createUserPassSecurityCtx("validuser", "password")
	appCtx := common.NewDefaultApplicationContext()
	ctx := core.NewExecutionContext(appCtx, securityCtx)
	privileges := []string{}

	authzFilter := setup(privileges)
	methodResult := authzFilter.Invoke("com.pkg.svc", "op1", data.NewVoidValue(), ctx)
	assert.Nil(t, methodResult.Output())
	assert.Equal(t, methodResult.Error().Name(), "com.vmware.vapi.std.errors.unauthorized")
}

func TestInvalidUser(t *testing.T) {
	securityCtx := createUserPassSecurityCtx("invaliduser", "password")
	appCtx := common.NewDefaultApplicationContext()
	ctx := core.NewExecutionContext(appCtx, securityCtx)
	privileges := []string{}

	authzFilter := setup(privileges)

	methodResult := authzFilter.Invoke("vmodl.test.uber.privileges.uber_privileges", "operation_with_id_parameter", data.NewVoidValue(), ctx)
	assert.Equal(t, methodResult.Error().Name(), "com.vmware.vapi.std.errors.unauthorized")
	assert.Equal(t, methodResult.Output(), nil)
}

func TestNosecurityCtxInRetrieveUserIdentity(t *testing.T) {
	appCtx := common.NewDefaultApplicationContext()
	ctx := core.NewExecutionContext(appCtx, nil)
	_, err := security.RetrieveUserIdentity(ctx)

	expected_err := l10n.NewRuntimeError("vapi.security.authentication.failed", make(map[string]string))
	assert.Equal(t, err, expected_err)
}

func TestRetrieveUserIdentityAssertionFailed(t *testing.T) {
	securityCtx := createUserPassSecurityCtx("validuser", "password")
	// overriding the AUTHN_IDENTITY
	securityCtx.SetProperty(security.AUTHN_IDENTITY, nil)

	appCtx := common.NewDefaultApplicationContext()
	ctx := core.NewExecutionContext(appCtx, securityCtx)
	_, err := security.RetrieveUserIdentity(ctx)

	expected_err := l10n.NewRuntimeError("vapi.security.authentication.failed", make(map[string]string))
	assert.Equal(t, expected_err, err)
}
