/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Statistics
 * Used by client-side stubs.
 */

package rules

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type StatisticsClient interface {

    // Get statistics of a IDS/IPS rule. - no enforcement point path specified: Stats will be evaluated on each enforcement point. - {enforcement_point_path}: Stats are evaluated only on the given enforcement point.
    //
    // @param domainIdParam Domain id (required)
    // @param idsPolicyIdParam IDS policy id (required)
    // @param ruleIdParam Rule id (required)
    // @param enforcementPointPathParam String Path of the enforcement point (optional)
    // @return com.vmware.nsx_policy.model.IdsRuleStatisticsListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(domainIdParam string, idsPolicyIdParam string, ruleIdParam string, enforcementPointPathParam *string) (model.IdsRuleStatisticsListResult, error)
}
