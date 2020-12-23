/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Preferences
 * Used by client-side stubs.
 */

package onboarding

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type PreferencesClient interface {

    // Get user onboarding preferences for a site on global manager.
    //
    // @param siteIdParam (required)
    // @return com.vmware.nsx_global_policy.model.SiteOnboardingPreference
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(siteIdParam string) (model.SiteOnboardingPreference, error)

    // Update user onboarding preferences to allow or reject site onboarding on global manager.
    //
    // @param siteIdParam (required)
    // @param siteOnboardingPreferenceParam (required)
    // @return com.vmware.nsx_global_policy.model.SiteOnboardingPreference
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(siteIdParam string, siteOnboardingPreferenceParam model.SiteOnboardingPreference) (model.SiteOnboardingPreference, error)
}
