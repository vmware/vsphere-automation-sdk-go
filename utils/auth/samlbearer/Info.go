package samlbearer

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/security"
)

type info struct {
	Token string
}

func (saml *info) GetSAMLBearerToken() string {
	return saml.Token
}

func (saml *info) GetName() string {
	return Name
}

func (saml *info) GetAuthInterface() interface{} {
	return saml
}

func (saml *info) GetSecurityContext() (core.SecurityContext, error) {
	securityContext := core.NewSecurityContextImpl()
	securityContext.SetProperty(security.AUTHENTICATION_SCHEME_ID, security.SAML_BEARER_SCHEME_ID)
	securityContext.SetProperty(security.SAML_TOKEN, saml.Token)
	return securityContext, nil
}
