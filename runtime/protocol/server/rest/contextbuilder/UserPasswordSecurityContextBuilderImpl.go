/* Copyright © 2020, 2023 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package contextbuilder

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/security"
)

type UserPasswordSecurityContextBuilderImpl struct {
	AuthType string
}

func NewUserPasswordSecurityContextBuilderImpl() *UserPasswordSecurityContextBuilderImpl {
	return &UserPasswordSecurityContextBuilderImpl{
		AuthType: "Basic",
	}
}

func (sCtx *UserPasswordSecurityContextBuilderImpl) BuildSecurityContext(r *http.Request) (core.SecurityContext, error) {
	secCtx := core.NewSecurityContextImpl()

	// Ref: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Authorization
	val := r.Header.Get(lib.AUTH_HEADER)
	user_pass_base64 := strings.TrimSpace(strings.TrimPrefix(val, sCtx.AuthType))
	user_pass, err := base64.StdEncoding.DecodeString(user_pass_base64)
	if err != nil {
		return nil, err
	}

	user, pass := decodeUserPass(string(user_pass))
	secCtx.SetProperty(security.USER_KEY, user)
	secCtx.SetProperty(security.PASSWORD_KEY, pass)
	secCtx.SetProperty(security.AUTHENTICATION_SCHEME_ID, security.USER_PASSWORD_SCHEME_ID)

	return secCtx, nil
}

func (sCtx *UserPasswordSecurityContextBuilderImpl) CanHandle(r *http.Request) bool {
	if val, ok := r.Header[lib.AUTH_HEADER]; ok && strings.HasPrefix(val[0], sCtx.AuthType) {
		return true
	}
	return false
}

func decodeUserPass(userPass string) (string, string) {
	// see https://www.rfc-editor.org/rfc/rfc7617#section-2
	// "a user-id containing a colon character is invalid, as the first colon
	// in a user-pass string separates user-id and password from one another;
	// text after the first colon is part of the password"
	firstColon := strings.IndexByte(userPass, ':')
	if firstColon < 0 {
		return userPass, ""
	}
	return userPass[:firstColon], userPass[firstColon+1:]
}
