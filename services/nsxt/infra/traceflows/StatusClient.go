/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Status
 * Used by client-side stubs.
 */

package traceflows

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type StatusClient interface {

    // Read traceflow status with id traceflow-id
    //
    // @param traceflowIdParam (required)
    // @param traceflowStatusRequestParam (required)
    // @return com.vmware.nsx_policy.model.TraceflowStatusResponse
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(traceflowIdParam string, traceflowStatusRequestParam model.TraceflowStatusRequest) (model.TraceflowStatusResponse, error)
}
