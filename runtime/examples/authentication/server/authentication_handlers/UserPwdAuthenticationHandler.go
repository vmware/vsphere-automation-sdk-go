/* Copyright © 2021-2022 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package auth

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std"
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/security"
)

type UserPasswordAuthenticationHandler struct {
}

func NewUserPwdAuthenticationHandler() *UserPasswordAuthenticationHandler {
	return &UserPasswordAuthenticationHandler{}
}

func getUnauthenticatedError() error {
	err := errors.NewUnauthenticated()
	err.Messages = []std.LocalizableMessage{{Id: "example.authentication.userpass.failed", DefaultMessage: "authentication failed", Args: []string{}}}
	return err
}

func (u *UserPasswordAuthenticationHandler) Authenticate(ctx core.SecurityContext) (*security.UserIdentity, error) {
	const validUserName = "username"
	const validPassword = "password"

	if ctx.Property(security.AUTHENTICATION_SCHEME_ID) == u.SupportedScheme() {
		if ctx.Property(security.USER_KEY) == validUserName && ctx.Property(security.PASSWORD_KEY) == validPassword {
			return security.NewUserIdentity(validUserName), nil
		} else {
			return nil, getUnauthenticatedError()
		}
	}
	//Returning nil as error so the AuthenticationFilter can continue and try
	//with the next auth handler if one exists
	return nil, nil
}

func (*UserPasswordAuthenticationHandler) SupportedScheme() string {
	return security.USER_PASSWORD_SCHEME_ID
}
