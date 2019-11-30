package refreshtoken

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/args/keys"
)

func CreateOAuthRefreshTokenFromProperties(authMap map[string]interface{}) (Info, error) {
	cspurl := authMap[keys.AuthSchemeOAuthCSPURLKey]
	token := authMap[keys.AuthSchemeOAuthRefreshTokenKey]
	if cspurl == nil && token == nil {
		return nil, &Error{}
	} else if cspurl == nil {
		return nil, &CSPURLError{}
	} else if token == nil {
		return nil, &RefreshTokenError{}
	}

	return &info{cspURL: cspurl.(string), refreshToken: token.(string)}, nil
}

func Create(token string, cspurl string) (Info, error) {
	if len(cspurl) <= 0 && len(token) <= 0 {
		return nil, &Error{}
	} else if len(cspurl) <= 0 {
		return nil, &CSPURLError{}
	} else if len(token) <= 0 {
		return nil, &RefreshTokenError{}
	}

	return &info{cspURL: cspurl, refreshToken: token}, nil
}
