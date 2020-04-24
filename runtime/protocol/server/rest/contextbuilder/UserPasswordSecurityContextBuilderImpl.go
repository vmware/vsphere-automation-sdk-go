/* Copyright © 2020 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package contextbuilder

import (
	"encoding/base64"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/security"
	"net/http"
	"strings"
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
	value := r.Header[lib.AUTH_HEADER]
	// req.Header returns a list of values for each key (name)
	val := value[0]
	base64_index := len(sCtx.AuthType + " ")
	user_pass_base64 := val[base64_index:]
	user_pass, err := base64.StdEncoding.DecodeString(user_pass_base64)
	if err != nil {
		return nil, err
	}

	temp := strings.Split(string(user_pass), ":")
	user, pass := temp[0], temp[1]
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
