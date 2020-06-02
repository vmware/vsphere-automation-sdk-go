/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Status
 * Used by client-side stubs.
 */

package compliance

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type StatusClient interface {

    // Returns the compliance status and details of non compliant configuration
    // @return com.vmware.nsx_policy.model.PolicyComplianceStatus
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get() (model.PolicyComplianceStatus, error)
}
