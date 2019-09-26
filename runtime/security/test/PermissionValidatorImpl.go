package test

/* **********************************************************
 * Copyright 2019 VMware, Inc.  All rights reserved.
 *      -- VMware Confidential
 * **********************************************************/

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/log"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/security"
)

type PermissionValidatorImpl struct {
	authzProvider AuthorizationProvider
}

func NewPermissionValidatorImpl(authzProvider AuthorizationProvider) *PermissionValidatorImpl {
	permissionValidatorImpl := PermissionValidatorImpl{authzProvider: authzProvider}
	return &permissionValidatorImpl
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (perValid *PermissionValidatorImpl) Validate(username string, groups []string, resIdPrivileges map[security.ResourceIdentifier][]string) bool {

	// Check if Role exists
	// userPrivileges represents privileges avaiable by the user username
	userPrivileges, ok := perValid.authzProvider.Privileges(username)
	if !ok {
		return false
	}

	requiredPrivileges := []string{}
	for _, v := range resIdPrivileges {
		requiredPrivileges = append(requiredPrivileges, v...)
	}

	log.Debugf("User Privileges: %s", userPrivileges)
	log.Debugf("Required Privileges for given Operation: %s", requiredPrivileges)

	for _, val := range requiredPrivileges {
		if !contains(userPrivileges, val) {
			return false
		}
	}

	return true
}
