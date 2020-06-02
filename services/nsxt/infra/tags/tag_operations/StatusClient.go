/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Status
 * Used by client-side stubs.
 */

package tag_operations

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type StatusClient interface {

    // Get status of tag bulk operation with details of tag operation on each virtual machine.
    //
    // @param operationIdParam (required)
    // @return com.vmware.nsx_policy.model.TagBulkOperationStatus
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(operationIdParam string) (model.TagBulkOperationStatus, error)
}
