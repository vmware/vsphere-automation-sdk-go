/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Status
 * Used by client-side stubs.
 */

package security

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type StatusClient interface {

    // Get the current service insertion status configuration.
    // @return com.vmware.nsx_policy.model.PolicySIStatusConfiguration
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get() (model.PolicySIStatusConfiguration, error)

    // Update service insertion status.
    //
    // @param policySIStatusConfigurationParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(policySIStatusConfigurationParam model.PolicySIStatusConfiguration) error

    // Update service insertion status.
    //
    // @param policySIStatusConfigurationParam (required)
    // @return com.vmware.nsx_policy.model.PolicySIStatusConfiguration
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(policySIStatusConfigurationParam model.PolicySIStatusConfiguration) (model.PolicySIStatusConfiguration, error)
}
