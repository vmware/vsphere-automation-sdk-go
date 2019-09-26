package security_test

/* **********************************************************
 * Copyright 2019 VMware, Inc.  All rights reserved.
 *      -- VMware Confidential
 * **********************************************************/

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/security"
)

type AuthenticationResultImpl struct {
	userIdentity    *security.UserIdentity
	securityContext core.SecurityContext
}

func NewAuthenticationResultImpl(name string) *AuthenticationResultImpl {
	userId := security.NewUserIdentity(name)
	authenticationResultImpl := AuthenticationResultImpl{userIdentity: userId, securityContext: nil}
	return &authenticationResultImpl
}

func (a *AuthenticationResultImpl) GetUserIdentity() *security.UserIdentity {
	return a.userIdentity
}

func (a *AuthenticationResultImpl) GetSecurityContext() core.SecurityContext {
	return a.securityContext
}
