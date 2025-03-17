// Copyright (c) 2021-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

package internal

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
)

// Connector contract for implementing specific client connector type
type Connector interface {
	Address() string
	ApplicationContext() *core.ApplicationContext
	SecurityContext() core.SecurityContext
	SetSecurityContext(core.SecurityContext)
	NewExecutionContext() *core.ExecutionContext
	GetApiProvider() core.APIProvider
	TypeConverter() *bindings.TypeConverter
}
