// Copyright (c) 2020-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

package bindings

import "github.com/vmware/vsphere-automation-sdk-go/runtime/data"

// Structure Base interface that represents a Go language bindings structure.
type Structure interface {
	GetType__() BindingType
	GetDataValue__() (data.DataValue, []error)
}
