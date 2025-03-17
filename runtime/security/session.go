// Copyright (c) 2019-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

package security

import (
	"crypto"
	"encoding/json"
)

// SessionSecurityContext represents the security context needed for authentication using session ID.
type SessionSecurityContext struct {
	properties map[string]interface{}
}

func NewSessionSecurityContext(sessionID string) *SessionSecurityContext {
	properties := map[string]interface{}{}
	properties[SESSION_ID] = sessionID
	properties[AUTHENTICATION_SCHEME_ID] = SESSION_SCHEME_ID
	return &SessionSecurityContext{properties: properties}
}

func (s *SessionSecurityContext) Property(key string) interface{} {
	return s.properties[key]
}

func (s *SessionSecurityContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.properties)
}

func (s *SessionSecurityContext) GetAllProperties() map[string]interface{} {
	return s.properties
}

func (s *SessionSecurityContext) SetProperty(property string, value interface{}) {
	s.properties[property] = value
}

// UserPasswordSecurityContext represents a security context suitable for user/password authentication.
type UserPasswordSecurityContext struct {
	properties map[string]interface{}
}

func NewUserPasswordSecurityContext(username string, password string) *UserPasswordSecurityContext {
	properties := map[string]interface{}{}
	properties[USER_KEY] = username
	properties[PASSWORD_KEY] = password
	properties[AUTHENTICATION_SCHEME_ID] = USER_PASSWORD_SCHEME_ID
	return &UserPasswordSecurityContext{properties: properties}
}

func (u *UserPasswordSecurityContext) Property(key string) interface{} {
	return u.properties[key]
}

func (u *UserPasswordSecurityContext) User() string {
	return u.properties[USER_KEY].(string)
}

func (u *UserPasswordSecurityContext) Password() string {
	return u.properties[PASSWORD_KEY].(string)
}

func (u *UserPasswordSecurityContext) GetAllProperties() map[string]interface{} {
	return u.properties
}

func (u *UserPasswordSecurityContext) SetProperty(key string, value interface{}) {
	u.properties[key] = value
}

func (u *UserPasswordSecurityContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.properties)
}

// OauthSecurityContext represents a security context suitable for oauth authentication.
type OauthSecurityContext struct {
	properties map[string]interface{}
}

func NewOauthSecurityContext(accessToken string) *OauthSecurityContext {
	properties := map[string]interface{}{}
	properties[AUTHENTICATION_SCHEME_ID] = OAUTH_SCHEME_ID
	properties[ACCESS_TOKEN] = accessToken
	return &OauthSecurityContext{properties: properties}
}

func (o *OauthSecurityContext) Property(key string) interface{} {
	return o.properties[key]
}

func (o *OauthSecurityContext) Token() string {
	return o.properties[ACCESS_TOKEN].(string)
}

func (o *OauthSecurityContext) GetAllProperties() map[string]interface{} {
	return o.properties
}

func (o *OauthSecurityContext) SetProperty(key string, value interface{}) {
	o.properties[key] = value
}

func (o *OauthSecurityContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.properties)
}

// SAMLSecurityContext represents a security context for SAML tokens.
type SAMLSecurityContext struct {
	properties map[string]interface{}
}

// NewSAMLSecurityContext creates a security context for the given SAML
// holder-of-key (HoK) token.
//
// When an API request is authenticated via a SAML HoK token, the request needs
// to be signed with the private key associated with the SAML HoK token. This
// signature is meant to prove that the sender of the request indeed holds the
// private key.
//
// The given private key must be in PEM format. The given signature algorithm
// must match the type of the private key, e.g. use security.RS256,
// security.RS384 or security.RS512 for RSA private keys.
//
//	example:
//		connector := client.NewConnector(
//		"https://www.example.com/api",
//		client.WithSecurityContext(NewSAMLSecurityContext(
//			mySamlHolderOfKeyToken,
//			myPEMencodedRSAkey,
//			security.RS256)))
//		client := NewSampleClient(connector)
//		client.MyOperation()
func NewSAMLSecurityContext(samlHokToken, privateKey, signAlgorithm string) *SAMLSecurityContext {
	return &SAMLSecurityContext{
		properties: map[string]interface{}{
			AUTHENTICATION_SCHEME_ID: SAML_HOK_SCHEME_ID,
			SAML_TOKEN:               samlHokToken,
			PRIVATE_KEY:              privateKey,
			SIGNATURE_ALGORITHM:      signAlgorithm,
		},
	}
}

// NewSAMLPrivateKeySecurityContext creates a security context for the given
// SAML holder-of-key (HoK) token.
//
// When an API request is authenticated via a SAML HoK token, the request needs
// to be signed with the private key associated with the SAML HoK token. This
// signature is meant to prove that the sender of the request indeed holds the
// private key.
//
// The given private key must be in PEM format. The given signature algorithm
// must match the type of the private key, e.g. use security.RS256,
// security.RS384 or security.RS512 for RSA private keys.
//
//		example:
//	     var myRSAPrivateKey rsa.PrivateKey
//	     ...
//			connector := client.NewConnector(
//			"https://www.example.com/api",
//			client.WithSecurityContext(NewSAMLPrivateKeySecurityContext(
//				mySamlHolderOfKeyToken,
//				myRSAPrivateKey,
//				security.RS256)))
//			client := NewSampleClient(connector)
//			client.MyOperation()
func NewSAMLPrivateKeySecurityContext(samlHokToken string, privateKey crypto.PrivateKey, signAlgorithm string) *SAMLSecurityContext {
	return &SAMLSecurityContext{
		properties: map[string]interface{}{
			AUTHENTICATION_SCHEME_ID: SAML_HOK_SCHEME_ID,
			SAML_TOKEN:               samlHokToken,
			PRIVATE_KEY:              privateKey,
			SIGNATURE_ALGORITHM:      signAlgorithm,
		},
	}
}

// NewSAMLBearerSecurityContext creates a security context for the given SAML
// bearer token.
func NewSAMLBearerSecurityContext(samlBearerToken string) *SAMLSecurityContext {
	return &SAMLSecurityContext{
		properties: map[string]interface{}{
			AUTHENTICATION_SCHEME_ID: SAML_BEARER_SCHEME_ID,
			SAML_TOKEN:               samlBearerToken,
		},
	}
}

func (o *SAMLSecurityContext) Property(key string) interface{} {
	return o.properties[key]
}

func (o *SAMLSecurityContext) Token() string {
	return o.properties[SAML_TOKEN].(string)
}

func (o *SAMLSecurityContext) GetAllProperties() map[string]interface{} {
	return o.properties
}

func (o *SAMLSecurityContext) SetProperty(key string, value interface{}) {
	o.properties[key] = value
}

func (o *SAMLSecurityContext) MarshalJSON() ([]byte, error) {
	var marshallable map[string]interface{}
	if SAML_HOK_SCHEME_ID == o.properties[AUTHENTICATION_SCHEME_ID] {
		// SAML HoK requires sophisticated serialization - see JsonSsoSigner.
		// Here just include the scheme id for easier troubleshooting.
		marshallable = map[string]interface{}{AUTHENTICATION_SCHEME_ID: SAML_HOK_SCHEME_ID}
	} else {
		marshallable = o.properties
	}
	return json.Marshal(marshallable)
}
