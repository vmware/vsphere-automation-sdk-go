/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: DetailedStatus
 * Used by client-side stubs.
 */

package sessions

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type DetailedStatusClient interface {

    // - no enforcement point path specified: detailed status is evaluated on each enforcement point. - an enforcement point path is specified: detailed status is evaluated only on the given enforcement point. - source=realtime: detailed status is fetched realtime from the enforcement point. - source=cached: cached detailed status from enforcement point is returned.
    //
    // @param tier0IdParam (required)
    // @param localeServiceIdParam (required)
    // @param serviceIdParam (required)
    // @param sessionIdParam (required)
    // @param enforcementPointPathParam String Path of the enforcement point (optional)
    // @param sourceParam Data source type. (optional)
    // @return com.vmware.nsx_policy.model.AggregateIPSecVpnSessionStatus
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(tier0IdParam string, localeServiceIdParam string, serviceIdParam string, sessionIdParam string, enforcementPointPathParam *string, sourceParam *string) (model.AggregateIPSecVpnSessionStatus, error)
}
