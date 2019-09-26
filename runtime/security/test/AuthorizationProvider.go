package test

/* **********************************************************
 * Copyright 2019 VMware, Inc.  All rights reserved.
 *      -- VMware Confidential
 * **********************************************************/

type AuthorizationProvider interface {
	SetPrivileges(role string, arr []string)
	Privileges(role string) ([]string, bool)
}

type AuthorizationPrivilegeProvider struct {
	privileges map[string][]string
}

func NewAuthorizationPrivilegeProvider() *AuthorizationPrivilegeProvider {
	authzProvider := AuthorizationPrivilegeProvider{privileges: map[string][]string{}}
	return &authzProvider
}

func (a *AuthorizationPrivilegeProvider) SetPrivileges(role string, arr []string) {
	a.privileges[role] = arr
}

func (a *AuthorizationPrivilegeProvider) Privileges(role string) ([]string, bool) {
	if v, found := a.privileges[role]; found {
		return v, found
	} else {
		return nil, false
	}
}
