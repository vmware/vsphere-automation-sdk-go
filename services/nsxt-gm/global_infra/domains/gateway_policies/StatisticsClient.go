/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Statistics
 * Used by client-side stubs.
 */

package gateway_policies

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type StatisticsClient interface {

    // Get statistics of a gateay policy. - no enforcement point path specified: Stats will be evaluated on each enforcement. point. - {enforcement_point_path}: Stats are evaluated only on the given enforcement point.
    //
    // @param domainIdParam (required)
    // @param gatewayPolicyIdParam (required)
    // @param enforcementPointPathParam String Path of the enforcement point (optional)
    // @return com.vmware.nsx_global_policy.model.SecurityPolicyStatisticsListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(domainIdParam string, gatewayPolicyIdParam string, enforcementPointPathParam *string) (model.SecurityPolicyStatisticsListResult, error)
}
