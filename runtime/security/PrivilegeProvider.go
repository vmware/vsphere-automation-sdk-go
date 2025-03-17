// Copyright (c) 2019-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

package security

import "github.com/vmware/vsphere-automation-sdk-go/runtime/data"

type PrivilegeProvider interface {
	GetPrivilegeInfo(fullyQualifiedOperName string, inputValue data.DataValue) (map[ResourceIdentifier][]string, error)
}
