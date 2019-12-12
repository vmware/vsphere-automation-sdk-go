/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: History
 * Used by client-side stubs.
 */

package backups

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type HistoryClient interface {

    // Get history of previous backup operations
    // @return com.vmware.nsx_policy.model.BackupOperationHistory
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get() (model.BackupOperationHistory, error)
}
