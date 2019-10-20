package refreshtoken

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/args/keys"
)

func CreateOAuthRefreshTokenDetails(authMap map[string]interface{}) (Auth, error) {
	cspurl := authMap[keys.AuthSchemeOAuthCSPURLKey]
	token := authMap[keys.AuthSchemeOAuthRefreshTokenKey]
	if cspurl == nil && token == nil {
		return nil, &Error{}
	} else if cspurl == nil {
		return nil, &UsernameError{}
	} else if token == nil {
		return nil, &PasswordError{}
	}

	return Auth(oAuthRefreshTokenDetails{cspURL: cspurl.(string), refreshToken: token.(string)}), nil
}
