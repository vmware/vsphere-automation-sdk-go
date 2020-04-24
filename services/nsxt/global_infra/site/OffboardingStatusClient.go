/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: OffboardingStatus
 * Used by client-side stubs.
 */

package site

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type OffboardingStatusClient interface {

    // Get site offboarding status.
    // @return com.vmware.nsx_policy.model.SiteOffBoardingState
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get() (model.SiteOffBoardingState, error)
}
