package security_test

import (
	"errors"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/security"
)

type UserPasswordAuthenticationHandler struct {
}

func NewUserPasswordAuthenticationHandler() *UserPasswordAuthenticationHandler {
	return &UserPasswordAuthenticationHandler{}
}

func (u *UserPasswordAuthenticationHandler) Authenticate(ctx core.SecurityContext) (*security.UserIdentity, error) {
	if ctx.Property(security.AUTHENTICATION_SCHEME_ID).(string) == security.USER_PASSWORD_SCHEME_ID {
		if ctx.Property(security.USER_KEY) == "testUser" && ctx.Property(security.PASSWORD_KEY) == "password" {
			return security.NewUserIdentity("testUser"), nil
		} else {
			return nil, errors.New("Authentication failed")
		}
	}
	return nil, nil
}

func (u *UserPasswordAuthenticationHandler) SupportedScheme() string {
	return security.USER_PASSWORD_SCHEME_ID
}
