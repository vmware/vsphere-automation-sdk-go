/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: State
 * Used by client-side stubs.
 */

package ports

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type StateClient interface {

    // Returns tier-1 segment port state on enforcement point
    //
    // @param tier1IdParam (required)
    // @param segmentIdParam (required)
    // @param portIdParam (required)
    // @param enforcementPointPathParam String Path of the enforcement point (optional)
    // @param sourceParam Data source type. (optional)
    // @return com.vmware.nsx_global_policy.model.SegmentPortState
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(tier1IdParam string, segmentIdParam string, portIdParam string, enforcementPointPathParam *string, sourceParam *string) (model.SegmentPortState, error)
}
