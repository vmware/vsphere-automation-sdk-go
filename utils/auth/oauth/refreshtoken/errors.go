/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package refreshtoken

// Error represents basic error for OAuth Authentication by Refresh Token.
type Error struct {
	ErrorDetails error
}

func (authError *Error) Error() string {
	return Name + " Error: CSPURL and RefreshToken are missing."
}

// CSPURLError represents the Error for CSP URL for OAuth Authentication
// Scheme by Refresh Token.
type CSPURLError struct{}

func (cspUrlError *CSPURLError) Error() string {
	return Name + " Error: CSPURL is not available."
}

// TokenError represents the Error for Refresh Token for OAuth Authentication
// Scheme by Refresh Token.
type TokenError struct{}

func (refreshTokenError *TokenError) Error() string {
	return Name + " Error: RefreshToken is not available."
}
