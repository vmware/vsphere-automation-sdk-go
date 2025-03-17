// Copyright (c) 2020-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

package rest

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"

	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/security"
)

type UserPwdSecContextSerializer struct {
}

func NewUserPwdSecContextSerializer() *UserPwdSecContextSerializer {
	return &UserPwdSecContextSerializer{}
}

// Serialize authorization headers for username security context
func (u *UserPwdSecContextSerializer) Serialize(ctx core.SecurityContext) (map[string]interface{}, error) {
	username, err := GetSecurityCtxStrValue(ctx, security.USER_KEY)
	if err != nil {
		return nil, err
	}

	password, err := GetSecurityCtxStrValue(ctx, security.PASSWORD_KEY)
	if err != nil {
		return nil, err
	}

	if username == nil || password == nil {
		err := errors.New("username and password are required for UserPwdSecContextSerializer")
		return nil, err
	}

	credentialString := fmt.Sprintf("%s:%s", *username, *password)
	base64EncodedVal := base64.StdEncoding.EncodeToString([]byte(credentialString))
	return map[string]interface{}{"Authorization": "Basic " + base64EncodedVal}, nil
}

var _ protocol.SecurityContextSerializer = NewUserPwdSecContextSerializer()

type SessionSecContextSerializer struct {
}

func NewSessionSecContextSerializer() *SessionSecContextSerializer {
	return &SessionSecContextSerializer{}
}

// Serialize authorization headers for session security context
func (s *SessionSecContextSerializer) Serialize(ctx core.SecurityContext) (map[string]interface{}, error) {
	sessionID, err := GetSecurityCtxStrValue(ctx, security.SESSION_ID)
	if err != nil {
		return nil, err
	}

	if sessionID == nil {
		err := errors.New("Session ID is required for SessionSecContextSerializer")
		return nil, err
	}

	return map[string]interface{}{security.SESSION_ID_KEY: *sessionID}, nil
}

var _ protocol.SecurityContextSerializer = NewSessionSecContextSerializer()

type OauthSecContextSerializer struct {
}

func NewOauthSecContextSerializer() *OauthSecContextSerializer {
	return &OauthSecContextSerializer{}
}

// Serialize authorization headers for oauth security context.
func (o *OauthSecContextSerializer) Serialize(ctx core.SecurityContext) (map[string]interface{}, error) {
	oauthToken, err := GetSecurityCtxStrValue(ctx, security.ACCESS_TOKEN)
	if err != nil {
		return nil, err
	}

	if oauthToken == nil {
		err := errors.New("Oauth token is required for OauthSecContextSerializer")
		return nil, err
	}

	return map[string]interface{}{security.CSP_AUTH_TOKEN_KEY: *oauthToken}, nil
}

var _ protocol.SecurityContextSerializer = NewOauthSecContextSerializer()

type SamlBearerSecContextSerializer struct {
}

func NewSamlBearerSecContextSerializer() *SamlBearerSecContextSerializer {
	return &SamlBearerSecContextSerializer{}
}

// Serialize a SAML bearer token into an HTTP Authorization header
func (s *SamlBearerSecContextSerializer) Serialize(ctx core.SecurityContext) (map[string]interface{}, error) {
	// see https://wiki.eng.vmware.com/SSO/REST
	// see https://wiki.eng.vmware.com/VAPI/APIEndpoint/REST_API_Authentication
	samlToken, err := GetSecurityCtxStrValue(ctx, security.SAML_TOKEN)
	if err != nil {
		return nil, err
	}
	if samlToken == nil || *samlToken == "" {
		err := errors.New("missing SAML token")
		return nil, err
	}
	encodedToken, err := encodeSamlToken(*samlToken)
	if err != nil {
		return nil, err
	}
	headerVal := fmt.Sprintf("SIGN token=\"%s\"", encodedToken)
	return map[string]interface{}{lib.AUTH_HEADER: headerVal}, nil
}

func encodeSamlToken(samlToken string) (string, error) {
	gzipped, err := gzipString(samlToken)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(gzipped), nil
}

func gzipString(s string) ([]byte, error) {
	var buffer bytes.Buffer
	gzipWriter := gzip.NewWriter(&buffer)
	_, err := gzipWriter.Write([]byte(s))
	if err != nil {
		return nil, err
	}
	err = gzipWriter.Close()
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

var _ protocol.SecurityContextSerializer = NewSamlBearerSecContextSerializer()
