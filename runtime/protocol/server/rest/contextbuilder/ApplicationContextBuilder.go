/* Copyright © 2020 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package contextbuilder

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"net/http"
)

type ApplicationContextBuilder interface {
	BuildApplicationContext(r *http.Request) (*core.ApplicationContext, error)
	CanHandle(r *http.Request) bool
}
