/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: ConfigOnboarding
 * Used by client-side stubs.
 */

package lmonboarding

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type ConfigOnboardingClient interface {

    // Complete config-onboarding process by sending details from GM whether GM was success or failure. With GM status provided to LM, LM may take necessary actions for next steps.
    //
    // @param policyLMOnboardingStatusParam (required)
    // @return com.vmware.nsx_policy.model.PolicyLMOnboardingStatus
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Completeonboarding(policyLMOnboardingStatusParam model.PolicyLMOnboardingStatus) (model.PolicyLMOnboardingStatus, error)

    // Invoke config-onboarding process
    //
    // @param policyLMOnboardingMetadataParam (required)
    // @return com.vmware.nsx_policy.model.PolicyLMOnboardingStatus
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Startonboarding(policyLMOnboardingMetadataParam model.PolicyLMOnboardingMetadata) (model.PolicyLMOnboardingStatus, error)
}
