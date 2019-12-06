/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Package auth defines different Authentication Schemes.
package auth

import (
	"fmt"
	"strings"

	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/args"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/args/keys"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/auth/basic"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/auth/model"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/auth/oauth/refreshtoken"
	oauthRefreshToken "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/auth/oauth/refreshtoken"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/auth/samlbearer"
)

// GetAuthSchemeByName fetches Authentication Scheme by Scheme Name.
func GetAuthSchemeByName(name string) Scheme {
	for _, authScheme := range Schemes {
		if strings.ToLower(string(authScheme)) == strings.ToLower(name) {
			return authScheme
		}
	}
	return NoAuth
}

// GetAuthSchemeFromProperties fetches Authentication Scheme from Properties.
func GetAuthSchemeFromProperties(properties args.Properties) (Scheme, error) {
	if properties == nil {
		return NoAuth, fmt.Errorf("No Authentication Scheme set")
	}
	schemeName, err := properties.GetStringPropertyValue(keys.AuthSchemeKey)
	if err != nil {
		return NoAuth, err
	}
	return GetAuthSchemeByName(schemeName), nil
}

// GetAllAuthScheme fetches All Authentication Scheme Names.
func GetAllAuthScheme() []string {
	var allAuthSchemes []string
	for _, authScheme := range Schemes {
		allAuthSchemes = append(allAuthSchemes, string(authScheme))
	}
	return allAuthSchemes
}

// GetAuthSchemeInfo returns Authentication Info from Properties.
func GetAuthSchemeInfo(arguments args.Properties) (Info, error) {
	var authInterface model.AuthInfo
	authType := arguments.GetPropertyValue(keys.AuthSchemeKey)
	authScheme := GetAuthSchemeByName(authType.(string))
	if authType != nil {
		switch authScheme {
		case NoAuth:
			return nil, nil
		case BasicAuth:
			authMap := arguments.Options()
			user := authMap[keys.AuthSchemeBasicUsernameKey]
			password := authMap[keys.AuthSchemeBasicPasswordKey]
			if user == nil && password == nil {
				return nil, &basic.Error{}
			} else if user == nil {
				return nil, &basic.UsernameError{}
			} else if password == nil {
				return nil, &basic.PasswordError{}
			}
			basicDetails, err := basic.Create(user.(string), password.(string))
			if err != nil {
				return nil, err
			}
			authInterface = basicDetails
		case OAuthRefreshToken:
			authMap := arguments.Options()
			cspurl := authMap[keys.AuthSchemeOAuthCSPURLKey]
			token := authMap[keys.AuthSchemeOAuthRefreshTokenKey]
			if cspurl == nil && token == nil {
				return nil, &refreshtoken.Error{}
			} else if cspurl == nil {
				return nil, &refreshtoken.CSPURLError{}
			} else if token == nil {
				return nil, &refreshtoken.TokenError{}
			}
			oAuthDetails, err := oauthRefreshToken.Create(token.(string), cspurl.(string))
			if err != nil {
				return nil, err
			}
			authInterface = oAuthDetails
		case SAMLBearer:
			authMap := arguments.Options()
			samlToken := authMap[keys.AuthSchemeSAMLBearerTokenKey]
			if samlToken == nil {
				return nil, &samlbearer.Error{}
			}
			samlBearerDetails, err := samlbearer.Create(samlToken.(string))
			if err != nil {
				return nil, err
			}
			authInterface = samlBearerDetails
		}
	}
	return &info{authInterface: authInterface, authScheme: authScheme}, nil
}

// IsBasicAuth checks whether Basic Authentication Scheme
// is defined or not in Properties.
func IsBasicAuth(properties args.Properties) bool {
	if properties == nil {
		return false
	}
	actualScheme, err := GetAuthSchemeFromProperties(properties)
	if err != nil {
		return false
	}
	return actualScheme == BasicAuth
}

// IsSAMLBearerAuth checks whether SAML Bearer Authentication Scheme
// is defined or not in Properties.
func IsSAMLBearerAuth(properties args.Properties) bool {
	if properties == nil {
		return false
	}
	actualScheme, err := GetAuthSchemeFromProperties(properties)
	if err != nil {
		return false
	}
	return actualScheme == SAMLBearer
}

// IsOAuthRefreshToken checks whether OAuth Authentication Scheme
// by Refresh Token is defined or not in Properties.
func IsOAuthRefreshToken(properties args.Properties) bool {
	if properties == nil {
		return false
	}
	actualScheme, err := GetAuthSchemeFromProperties(properties)
	if err != nil {
		return false
	}
	return actualScheme == OAuthRefreshToken
}
