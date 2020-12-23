/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Onboarding
 * Used by client-side stubs.
 */

package sites

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type OnboardingClient interface {

    // Verifies and reports conflicting onboarding feature for a site. The response will contain first conflicting feature for the site configuration compared to corresponding global manager configuration.
    //
    // @param siteIdParam (required)
    // @param configOnboardingConflictRequestParam (required)
    // @return com.vmware.nsx_global_policy.model.ConfigOnboardingConflictStatus
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Checkconflict(siteIdParam string, configOnboardingConflictRequestParam model.ConfigOnboardingConflictRequest) (model.ConfigOnboardingConflictStatus, error)

    // Initiate config on-boarding of a Site. The entire on-boarding is async workflow controlled by API.
    //
    // @param siteIdParam (required)
    // @param configOnboardingRequestParam (required)
    // @return com.vmware.nsx_global_policy.model.ConfigOnboardingStatus
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Startonboarding(siteIdParam string, configOnboardingRequestParam model.ConfigOnboardingRequest) (model.ConfigOnboardingStatus, error)
}
