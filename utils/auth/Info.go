/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package auth

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/auth/model"
)

type info struct {
	authInterface model.AuthInfo
	authScheme    Scheme
}

func (ai *info) AuthScheme() Scheme {
	return ai.authScheme
}

func (ai *info) AuthInterface() interface{} {
	switch ai.authScheme {
	case BasicAuth, OAuthRefreshToken, SAMLBearer:
		return ai.authInterface.AuthInterface()
	default:
		return nil
	}
}

func (ai *info) SecurityContext() (core.SecurityContext, error) {
	switch ai.authScheme {
	case BasicAuth, OAuthRefreshToken, SAMLBearer:
		return ai.authInterface.SecurityContext()
	default:
		return nil, nil
	}
}
