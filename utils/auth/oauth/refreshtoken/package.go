/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Package refreshtoken represents OAuth Authentication Scheme
// by Refresh Token, that uses Refresh Token for authentication.
package refreshtoken

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/args/keys"
)

// CreateOAuthRefreshTokenFromProperties defines and returns new Info for
// OAuth Authentication Scheme by RefreshToken from map of string to interface.
func CreateOAuthRefreshTokenFromProperties(authMap map[string]interface{}) (Info, error) {
	cspurl := authMap[keys.AuthSchemeOAuthCSPURLKey]
	token := authMap[keys.AuthSchemeOAuthRefreshTokenKey]
	if cspurl == nil && token == nil {
		return nil, &Error{}
	} else if cspurl == nil {
		return nil, &CSPURLError{}
	} else if token == nil {
		return nil, &TokenError{}
	}

	return &info{cspURL: cspurl.(string), refreshToken: token.(string)}, nil
}

// Create defines and returns new Info for
// OAuth Authentication Scheme by RefreshToken from map of string to interface.
func Create(token string, cspurl string) (Info, error) {
	if len(cspurl) <= 0 && len(token) <= 0 {
		return nil, &Error{}
	} else if len(cspurl) <= 0 {
		return nil, &CSPURLError{}
	} else if len(token) <= 0 {
		return nil, &TokenError{}
	}

	return &info{cspURL: cspurl, refreshToken: token}, nil
}
