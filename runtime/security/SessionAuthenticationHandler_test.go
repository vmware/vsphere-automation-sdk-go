package security_test

import (
	"errors"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/security"
)

type SessionAuthnHandler struct {
}

func NewSessionAuthnHandler() *SessionAuthnHandler {
	return &SessionAuthnHandler{}
}

func (s *SessionAuthnHandler) Authenticate(ctx core.SecurityContext) (*security.UserIdentity, error) {
	if ctx.Property(security.AUTHENTICATION_SCHEME_ID).(string) == s.SupportedScheme() {
		if ctx.Property(security.SESSION_ID) == "session-id" {
			uIdentity := security.NewUserIdentity("session-user")
			return uIdentity, nil
		} else {
			return nil, errors.New("Authentication failed")
		}
	}
	return nil, nil
}

func (s *SessionAuthnHandler) SupportedScheme() string {
	return security.SESSION_SCHEME_ID
}
