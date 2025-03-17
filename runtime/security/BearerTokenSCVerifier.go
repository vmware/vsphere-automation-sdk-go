// Copyright (c) 2019-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

package security

import (
	"errors"
)

// Verifies subject confirmation data of token matches bearer token.
// Deprecated: This validation is now part of JsonSsoVerifier processor. Use it instead.
type BearerTokenSCVerifier struct {
}

func NewBearerTokenSCVerifier() *BearerTokenSCVerifier {
	return &BearerTokenSCVerifier{}
}

// verify if the json payload that is passed contains valid bearer token and corresponding scheme ID
func (j *BearerTokenSCVerifier) Process(jsonMessage map[string]interface{}) error {

	securityContext, err := GetSecurityContext(jsonMessage)
	if err != nil {
		return err
	}
	if securityContext == nil {
		// anonymous requests have no security context
		return nil
	}

	if securityContext[AUTHENTICATION_SCHEME_ID] != SAML_BEARER_SCHEME_ID {
		return nil
	}

	if !isSamlBearerToken(securityContext[SAML_TOKEN].(string)) {
		return errors.New("SAML token passed is not of type Bearer")
	}

	return nil
}
