/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Statistics
 * Used by client-side stubs.
 */

package sessions

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type StatisticsClient interface {

    // Resets the statistics of the given VPN session. Since source of data is enforcement point, data is reset there.
    //
    // @param tier0IdParam (required)
    // @param localeServiceIdParam (required)
    // @param serviceIdParam (required)
    // @param sessionIdParam (required)
    // @param actionParam Action on statistics (required)
    // @param enforcementPointPathParam String Path of the enforcement point (optional)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Create(tier0IdParam string, localeServiceIdParam string, serviceIdParam string, sessionIdParam string, actionParam string, enforcementPointPathParam *string) error

    // - no enforcement point path specified: statistics are evaluated on each enforcement point. - an enforcement point path is specified: statistics are evaluated only on the given enforcement point. - source=realtime: statistics are fetched realtime from the enforcement point. - source=cached: cached statistics from enforcement point are returned.
    //
    // @param tier0IdParam (required)
    // @param localeServiceIdParam (required)
    // @param serviceIdParam (required)
    // @param sessionIdParam (required)
    // @param enforcementPointPathParam String Path of the enforcement point (optional)
    // @param sourceParam Data source type. (optional)
    // @return com.vmware.nsx_policy.model.AggregateIPSecVpnSessionStatistics
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(tier0IdParam string, localeServiceIdParam string, serviceIdParam string, sessionIdParam string, enforcementPointPathParam *string, sourceParam *string) (model.AggregateIPSecVpnSessionStatistics, error)
}
