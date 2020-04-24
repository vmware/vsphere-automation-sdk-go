/* Copyright © 2020 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package contextbuilder

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/security"
	"net/http"
	"strings"
)

type OAuthSecurityContextBuilderImpl struct {
	AuthType string
}

func NewOAuthSecurityContextBuilderImpl() *OAuthSecurityContextBuilderImpl {

	return &OAuthSecurityContextBuilderImpl{
		AuthType: "Bearer",
	}
}

func (sCtx *OAuthSecurityContextBuilderImpl) BuildSecurityContext(r *http.Request) (core.SecurityContext, error) {
	secCtx := core.NewSecurityContextImpl()

	/** Ref:
	*  1. https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Authorization
	   2. https://tools.ietf.org/html/rfc6750
	*/
	value := r.Header[lib.AUTH_HEADER]
	// req.Header returns a list of values for each key (name)
	val := value[0]

	access_token_index := len(sCtx.AuthType + " ")
	access_token := val[access_token_index:]
	secCtx.SetProperty(security.ACCESS_TOKEN, access_token)
	secCtx.SetProperty(security.AUTHENTICATION_SCHEME_ID, security.OAUTH_SCHEME_ID)

	return secCtx, nil
}

func (sCtx *OAuthSecurityContextBuilderImpl) CanHandle(r *http.Request) bool {
	if val, ok := r.Header[lib.AUTH_HEADER]; ok && strings.HasPrefix(val[0], sCtx.AuthType) {
		return true
	}
	return false
}
