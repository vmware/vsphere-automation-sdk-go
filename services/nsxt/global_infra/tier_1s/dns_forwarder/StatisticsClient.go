/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Statistics
 * Used by client-side stubs.
 */

package dns_forwarder

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type StatisticsClient interface {

    // Get statistics of tier-1 DNS forwarder. - no enforcement point path specified: Statistics will be evaluated on each enforcement point. - {enforcement_point_path}: Statistics are evaluated only on the given enforcement point.
    //
    // @param tier1IdParam Tier-1 id (required)
    // @param enforcementPointPathParam String Path of the enforcement point (optional)
    // @return com.vmware.nsx_policy.model.AggregateDNSForwarderStatistics
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(tier1IdParam string, enforcementPointPathParam *string) (model.AggregateDNSForwarderStatistics, error)
}
