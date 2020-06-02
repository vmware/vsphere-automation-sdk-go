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

    // Filter the firewall exclude list by the given object, to check whether the object is a member of this exclude list.
    //
    // @param intentPathParam Path of the intent object to be searched in the exclude list (required)
    // @param deepCheckParam Check all parents (optional, default to false)
    // @param enforcementPointPathParam Path of the enforcement point (optional)
    // @return com.vmware.nsx_policy.model.PolicyResourceReference
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Filter(intentPathParam string, deepCheckParam *bool, enforcementPointPathParam *string) (model.PolicyResourceReference, error)

    // Read exclude list for firewall
    // @return com.vmware.nsx_policy.model.PolicyExcludeList
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get() (model.PolicyExcludeList, error)

    // Patch exclusion list for security policy.
    //
    // @param policyExcludeListParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(policyExcludeListParam model.PolicyExcludeList) error

    // Update the exclusion list for security policy
    //
    // @param policyExcludeListParam (required)
    // @return com.vmware.nsx_policy.model.PolicyExcludeList
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(policyExcludeListParam model.PolicyExcludeList) (model.PolicyExcludeList, error)
}
