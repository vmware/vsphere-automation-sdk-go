/* Copyright © 2020 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package contextbuilder

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/security"
	"net/http"
)

type SessionSecurityContextBuilderImpl struct {
}

func NewSessionSecurityContextBuilderImpl() *SessionSecurityContextBuilderImpl {
	return &SessionSecurityContextBuilderImpl{}
}

func (sCtx *SessionSecurityContextBuilderImpl) BuildSecurityContext(r *http.Request) (core.SecurityContext, error) {
	secCtx := core.NewSecurityContextImpl()

	headerKey := http.CanonicalHeaderKey(lib.VAPI_SESSION_HEADER)
	value := r.Header[headerKey]

	// req.Header returns a list of values for each key (name)
	secCtx.SetProperty(security.SESSION_ID, value[0])
	secCtx.SetProperty(security.AUTHENTICATION_SCHEME_ID, security.SESSION_SCHEME_ID)

	return secCtx, nil
}

func (sCtx *SessionSecurityContextBuilderImpl) CanHandle(r *http.Request) bool {
	headerKey := http.CanonicalHeaderKey(lib.VAPI_SESSION_HEADER)
	if _, ok := r.Header[headerKey]; ok {
		return true
	}
	return false
}
