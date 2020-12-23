/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: InvalidConfigDetails
 * Used by client-side stubs.
 */

package onboarding

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type InvalidConfigDetailsClient interface {

    // Get feature summary details with invalid configuration for a feature.
    //
    // @param siteIdParam (required)
    // @param featureParam Unsupported features (required)
    // @return com.vmware.nsx_global_policy.model.FederationInvalidConfigurationDetailsResponse
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(siteIdParam string, featureParam string) (model.FederationInvalidConfigurationDetailsResponse, error)
}
