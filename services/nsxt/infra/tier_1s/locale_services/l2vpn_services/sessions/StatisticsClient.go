/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Statistics
 * Used by client-side stubs.
 */

package sessions

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type StatisticsClient interface {

    // - no enforcement point path specified: statistics are evaluated on each enforcement point. - an enforcement point path is specified: statistics are evaluated only on the given enforcement point. - source=realtime: statistics are fetched realtime from the enforcement point. - source=cached: cached statistics from the enforcement point are returned.
    //
    // @param tier1IdParam (required)
    // @param localeServiceIdParam (required)
    // @param serviceIdParam (required)
    // @param sessionIdParam (required)
    // @param enforcementPointPathParam String Path of the enforcement point (optional)
    // @param sourceParam Data source type. (optional)
    // @return com.vmware.nsx_policy.model.AggregateL2VPNSessionStatistics
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(tier1IdParam string, localeServiceIdParam string, serviceIdParam string, sessionIdParam string, enforcementPointPathParam *string, sourceParam *string) (model.AggregateL2VPNSessionStatistics, error)
}
