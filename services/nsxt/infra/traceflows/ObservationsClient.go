/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Observations
 * Used by client-side stubs.
 */

package traceflows

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type ObservationsClient interface {

    // Read traceflow observations for id traceflow-id. Traceflow configuration will be cleaned up by the system after two hours of inactivity.
    //
    // @param traceflowIdParam (required)
    // @param enforcementPointPathParam Enforcement point path (optional)
    // @return com.vmware.nsx_policy.model.TraceflowObservationListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(traceflowIdParam string, enforcementPointPathParam *string) (model.TraceflowObservationListResult, error)
}
