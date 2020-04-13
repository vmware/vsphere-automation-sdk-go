/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Status
 * Used by client-side stubs.
 */

package dns_forwarder

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type StatusClient interface {

    // Get current status of tier-1 DNS forwarder. - no enforcement point path specified: Status will be evaluated on each enforcement point. - {enforcement_point_path}: Status will be evaluated only on the given enforcement point.
    //
    // @param tier1IdParam Tier-1 id (required)
    // @param enforcementPointPathParam String Path of the enforcement point (optional)
    // @return com.vmware.nsx_global_policy.model.AggregateDNSForwarderStatus
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(tier1IdParam string, enforcementPointPathParam *string) (model.AggregateDNSForwarderStatus, error)
}
