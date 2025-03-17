// Copyright (c) 2020-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

package rest

import "net/http"

// RequestProcessor Provides access to request object right before execution
// Deprecated: use core.RequestProcessor instead
type RequestProcessor interface {
	Process(*http.Request) error
}
