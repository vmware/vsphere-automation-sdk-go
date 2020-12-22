/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: FeatureSummary
 * Used by client-side stubs.
 */

package onboarding

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type FeatureSummaryClient interface {

    // Get consolidated list of conflicting entities summary for each supported feature for a site with an example.
    //
    // @param siteIdParam (required)
    // @return com.vmware.nsx_global_policy.model.ConflictingEntityListResponse
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(siteIdParam string) (model.ConflictingEntityListResponse, error)
}
