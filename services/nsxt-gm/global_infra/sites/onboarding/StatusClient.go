/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Status
 * Used by client-side stubs.
 */

package onboarding

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type StatusClient interface {

    // Get onboarding status for a site.
    //
    // @param siteIdParam (required)
    // @return com.vmware.nsx_global_policy.model.ConfigOnboardingStatus
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(siteIdParam string) (model.ConfigOnboardingStatus, error)
}
