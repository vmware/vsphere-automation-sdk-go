/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: GroupVmDetails
 * Used by client-side stubs.
 */

package idfw

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type GroupVmDetailsClient interface {

    // Get all Identity Firewall Group VM details for a given Group.
    //
    // @param groupPathParam String Path of the group (required)
    // @param enforcementPointPathParam String Path of the enforcement point (optional)
    // @return com.vmware.nsx_policy.model.PolicyIdfwGroupVmDetailListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(groupPathParam string, enforcementPointPathParam *string) (model.PolicyIdfwGroupVmDetailListResult, error)
}
