/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: DetailedStatus
 * Used by client-side stubs.
 */

package lb_pools

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type DetailedStatusClient interface {

    // Get LBPool detailed status information. - no enforcement point path specified: Information will be aggregated from each enforcement point. - {enforcement_point_path}: Information will be retrieved only from the given enforcement point.
    //
    // @param lbServiceIdParam LBService id (required)
    // @param lbPoolIdParam LBPool id (required)
    // @param enforcementPointPathParam String Path of the enforcement point (optional)
    // @param sourceParam Data source type. (optional)
    // @return com.vmware.nsx_policy.model.AggregateLBPoolStatus
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(lbServiceIdParam string, lbPoolIdParam string, enforcementPointPathParam *string, sourceParam *string) (model.AggregateLBPoolStatus, error)
}
