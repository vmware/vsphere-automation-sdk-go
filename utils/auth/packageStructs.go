package auth

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/args"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/args/keys"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/auth/scheme"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/auth/scheme/basic"
	oauthRefreshToken "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/auth/scheme/oauth/refreshtoken"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/auth/scheme/samlbearer"
)

func GetAuthSchemeDetails(arguments args.Properties) (scheme.Details, error) {
	var details scheme.Details
	authType := arguments.GetPropertyValue(keys.AuthSchemeKey)
	if authType != nil {
		authScheme := scheme.GetAuthSchemeByName(authType.(string))
		switch authScheme {
		case scheme.NoAuth:
			return nil, nil
		case scheme.BasicAuth:
			basicDetails, err := basic.CreateBasicAuthDetails(arguments.Options())
			if err != nil {
				return nil, err
			}
			details = scheme.Details(basicDetails)
		case scheme.OAuthRefreshToken:
			oAuthDetails, err := oauthRefreshToken.CreateOAuthRefreshTokenDetails(arguments.Options())
			if err != nil {
				return nil, err
			}
			details = scheme.Details(oAuthDetails)
		case scheme.SAMLBearer:
			samlBearerDetails, err := samlbearer.CreateSAMLBearerAuthDetails(arguments.Options())
			if err != nil {
				return nil, err
			}
			details = scheme.Details(samlBearerDetails)
		}
	}
	return details, nil
}

func IsBasicAuth(properties args.Properties) bool {
	if properties == nil {
		return false
	}
	actualScheme, err := scheme.GetAuthSchemeFromProperties(properties)
	if err != nil {
		return false
	}
	return actualScheme == scheme.BasicAuth
}

func IsSAMLBearerAuth(properties args.Properties) bool {
	if properties == nil {
		return false
	}
	actualScheme, err := scheme.GetAuthSchemeFromProperties(properties)
	if err != nil {
		return false
	}
	return actualScheme == scheme.SAMLBearer
}

func IsOAuthRefreshToken(properties args.Properties) bool {
	if properties == nil {
		return false
	}
	actualScheme, err := scheme.GetAuthSchemeFromProperties(properties)
	if err != nil {
		return false
	}
	return actualScheme == scheme.OAuthRefreshToken
}
