// Copyright (c) 2019-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

package security

type PermissionValidator interface {
	Validate(username string, groups []string, requiredPrivileges map[ResourceIdentifier][]string) (bool, error)
}
