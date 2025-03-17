// Copyright (c) 2019-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

package security

import "github.com/vmware/vsphere-automation-sdk-go/runtime/core"

// Deprecated: No longer necessary.
type AuthorizationHandler interface {
	Authorize(serviceID string, operationID string, ctx core.SecurityContext) (bool, error)
}
