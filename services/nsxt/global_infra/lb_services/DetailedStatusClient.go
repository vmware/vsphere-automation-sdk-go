/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: DetailedStatus
 * Used by client-side stubs.
 */

package lb_services

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type DetailedStatusClient interface {

    // Get LBService detailed status information. - no enforcement point path specified: Information will be aggregated from each enforcement point. - {enforcement_point_path}: Information will be retrieved only from the given enforcement point.
    //
    // @param lbServiceIdParam LBService id (required)
    // @param enforcementPointPathParam String Path of the enforcement point (optional)
    // @param includeInstanceDetailsParam Flag to indicate whether include detail information (optional, default to false)
    // @param sourceParam Data source type. (optional)
    // @param transportNodeIdsParam The UUIDs of transport nodes (optional)
    // @return com.vmware.nsx_policy.model.AggregateLBServiceStatus
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(lbServiceIdParam string, enforcementPointPathParam *string, includeInstanceDetailsParam *bool, sourceParam *string, transportNodeIdsParam *string) (model.AggregateLBServiceStatus, error)
}
