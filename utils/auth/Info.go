package auth

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/auth/model"
)

type info struct {
	authInterface model.AuthInfo
	authScheme    Scheme
}

func (ai *info) GetAuthScheme() Scheme {
	return ai.authScheme
}

func (ai *info) GetAuthInterface() interface{} {
	switch ai.authScheme {
	case BasicAuth, OAuthRefreshToken, SAMLBearer:
		return ai.authInterface.GetAuthInterface()
	default:
		return nil
	}
}

func (ai *info) GetSecurityContext() (core.SecurityContext, error) {
	switch ai.authScheme {
	case BasicAuth, OAuthRefreshToken, SAMLBearer:
		return ai.authInterface.GetSecurityContext()
	default:
		return nil, nil
	}
}
