package samlbearer

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/security"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/auth/scheme"
)

func (saml samlBearerDetails) GetSAMLBearerToken() string {
	return saml.Token
}

func (saml samlBearerDetails) GetAuthScheme() scheme.Scheme {
	return scheme.SAMLBearer
}

func (saml samlBearerDetails) GetAuthInterface() interface{} {
	return Auth(saml)
}

func (saml samlBearerDetails) GetSecurityContext() (core.SecurityContext, error) {
	securityContext := core.NewSecurityContextImpl()
	securityContext.SetProperty(security.AUTHENTICATION_SCHEME_ID, security.SAML_BEARER_SCHEME_ID)
	securityContext.SetProperty(security.SAML_TOKEN, saml.Token)
	return securityContext, nil
}
