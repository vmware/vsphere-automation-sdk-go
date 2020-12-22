/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: ExcludeList
 * Used by client-side stubs.
 */

package security

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type ExcludeListClient interface {

    // Read exclude list for service insertion
    // @return com.vmware.nsx_policy.model.PolicySIExcludeList
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get() (model.PolicySIExcludeList, error)

    // Patch service insertion exclusion list for security policy.
    //
    // @param policySIExcludeListParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(policySIExcludeListParam model.PolicySIExcludeList) error

    // Update the exclusion list for service insertion policy
    //
    // @param policySIExcludeListParam (required)
    // @return com.vmware.nsx_policy.model.PolicySIExcludeList
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(policySIExcludeListParam model.PolicySIExcludeList) (model.PolicySIExcludeList, error)
}
