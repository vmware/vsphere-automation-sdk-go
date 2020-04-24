/* Copyright © 2020 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package contextbuilder

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"net/http"
)

type SecurityContextBuilder interface {
	BuildSecurityContext(r *http.Request) (core.SecurityContext, error)
	CanHandle(r *http.Request) bool
}
