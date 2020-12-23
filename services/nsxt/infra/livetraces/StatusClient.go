/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Status
 * Used by client-side stubs.
 */

package livetraces

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type StatusClient interface {

    // Read status for a livetrace config with the specified identifier.
    //
    // @param livetraceIdParam (required)
    // @param enforcementPointPathParam Enforcement point path (optional)
    // @return com.vmware.nsx_policy.model.LiveTraceStatus
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(livetraceIdParam string, enforcementPointPathParam *string) (model.LiveTraceStatus, error)
}
