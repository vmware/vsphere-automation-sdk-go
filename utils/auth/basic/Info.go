package basic

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/security"
)

type info struct {
	userName string
	password string
}

func (basicAuthInfo *info) UserName() string {
	return basicAuthInfo.userName
}

func (basicAuthInfo *info) Password() string {
	return basicAuthInfo.password
}

func (basicAuthInfo *info) GetName() string {
	return Name
}

func (basicAuthInfo *info) GetAuthInterface() interface{} {
	return Info(basicAuthInfo)
}

func (basicAuthInfo *info) GetSecurityContext() (core.SecurityContext, error) {
	return security.NewUserPasswordSecurityContext(basicAuthInfo.UserName(), basicAuthInfo.Password()), nil
}
