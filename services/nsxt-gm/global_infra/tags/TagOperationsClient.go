/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: TagOperations
 * Used by client-side stubs.
 */

package tags

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type TagOperationsClient interface {

    // Get details of tag bulk operation request with which tag is applied or removed on virtual machines.
    //
    // @param operationIdParam (required)
    // @return com.vmware.nsx_global_policy.model.TagBulkOperation
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(operationIdParam string) (model.TagBulkOperation, error)

    // Tag can be assigned or unassigned on multiple objects. Supported object type is restricted to Virtual Machine for now and support for other objects will be added later. Permissions for tag bulk operation would be similar to virtual machine tag permissions.
    //
    // @param operationIdParam (required)
    // @param tagBulkOperationParam (required)
    // @return com.vmware.nsx_global_policy.model.TagBulkOperation
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(operationIdParam string, tagBulkOperationParam model.TagBulkOperation) (model.TagBulkOperation, error)
}
