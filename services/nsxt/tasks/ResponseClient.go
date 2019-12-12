/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Response
 * Used by client-side stubs.
 */

package tasks

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
)

type ResponseClient interface {

    // Get the response of a task
    //
    // @param taskIdParam ID of task to read (required)
    // @return DynamicStructure
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(taskIdParam string) (*data.StructValue, error)
}
