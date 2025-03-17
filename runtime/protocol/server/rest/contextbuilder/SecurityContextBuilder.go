// Copyright (c) 2020-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

package contextbuilder

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"net/http"
)

type SecurityContextBuilder interface {
	BuildSecurityContext(r *http.Request) (core.SecurityContext, error)
	CanHandle(r *http.Request) bool
}
