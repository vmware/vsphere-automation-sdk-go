/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package task

import "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/session"

// Task represents the task that is to be executed on the Host server.
type Task interface {
	Execute(sessionManager session.Manager) error
}
