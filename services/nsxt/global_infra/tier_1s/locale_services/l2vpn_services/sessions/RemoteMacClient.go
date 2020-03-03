/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: RemoteMac
 * Used by client-side stubs.
 */

package sessions

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type RemoteMacClient interface {

    // Returns L2Vpn session remote macs for a logical switch. Data is fetched from enforcement point.
    //
    // @param tier1IdParam (required)
    // @param localeServiceIdParam (required)
    // @param serviceIdParam (required)
    // @param sessionIdParam (required)
    // @param enforcementPointPathParam String Path of the enforcement point (optional)
    // @param segmentPathParam Segment Path (optional)
    // @return com.vmware.nsx_policy.model.AggregateL2VpnSessionRemoteMac
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(tier1IdParam string, localeServiceIdParam string, serviceIdParam string, sessionIdParam string, enforcementPointPathParam *string, segmentPathParam *string) (model.AggregateL2VpnSessionRemoteMac, error)
}
