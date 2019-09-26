package security

/* **********************************************************
 * Copyright 2019 VMware, Inc.  All rights reserved.
 *      -- VMware Confidential
 * **********************************************************/

type PermissionValidator interface {
	Validate(username string, groups []string, requiredPrivileges map[ResourceIdentifier][]string) bool
}
