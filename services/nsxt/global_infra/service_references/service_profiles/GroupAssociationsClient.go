/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: GroupAssociations
 * Used by client-side stubs.
 */

package service_profiles

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type GroupAssociationsClient interface {

    // List of Groups used in Redirection rules for a given Service Profile.
    //
    // @param serviceReferenceIdParam Service reference id (required)
    // @param serviceProfileIdParam Service profile id (required)
    // @param enforcementPointPathParam String Path of the enforcement point (optional)
    // @return com.vmware.nsx_policy.model.ServiceProfileGroups
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(serviceReferenceIdParam string, serviceProfileIdParam string, enforcementPointPathParam *string) (model.ServiceProfileGroups, error)
}
