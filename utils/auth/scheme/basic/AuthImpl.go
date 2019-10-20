package basic

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/security"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/auth/scheme"
)

func (bad basicAuthDetails) GetUserName() string {
	return bad.UserName
}

func (bad basicAuthDetails) GetPassword() string {
	return bad.Password
}

func (bad basicAuthDetails) GetAuthScheme() scheme.Scheme {
	return scheme.BasicAuth
}

func (bad basicAuthDetails) GetAuthInterface() interface{} {
	return Auth(bad)
}

func (bad basicAuthDetails) GetSecurityContext() (core.SecurityContext, error) {
	return security.NewUserPasswordSecurityContext(bad.GetUserName(), bad.GetPassword()), nil
}
