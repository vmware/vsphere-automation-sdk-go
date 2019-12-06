/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package samlbearer

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/security"
)

type info struct {
	Token string
}

func (saml *info) SAMLBearerToken() string {
	return saml.Token
}

func (saml *info) Name() string {
	return Name
}

func (saml *info) AuthInterface() interface{} {
	return saml
}

func (saml *info) SecurityContext() (core.SecurityContext, error) {
	securityContext := core.NewSecurityContextImpl()
	securityContext.SetProperty(security.AUTHENTICATION_SCHEME_ID, security.SAML_BEARER_SCHEME_ID)
	securityContext.SetProperty(security.SAML_TOKEN, saml.Token)
	return securityContext, nil
}
