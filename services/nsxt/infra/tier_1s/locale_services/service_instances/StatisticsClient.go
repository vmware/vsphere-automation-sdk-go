/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Statistics
 * Used by client-side stubs.
 */

package service_instances

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type StatisticsClient interface {

    // Get statistics for all data NICs on all runtimes associated with this Tier1 PolicyServiceInstance.
    //
    // @param tier1IdParam Tier-1 id (required)
    // @param localeServiceIdParam Locale service id (required)
    // @param serviceInstanceIdParam Tier1 Service instance id (required)
    // @param enforcementPointPathParam String Path of the enforcement point (optional)
    // @return com.vmware.nsx_policy.model.PolicyServiceInstanceStatistics
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(tier1IdParam string, localeServiceIdParam string, serviceInstanceIdParam string, enforcementPointPathParam *string) (model.PolicyServiceInstanceStatistics, error)
}
