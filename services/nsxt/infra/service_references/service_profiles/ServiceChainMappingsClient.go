/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: ServiceChainMappings
 * Used by client-side stubs.
 */

package service_profiles

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type ServiceChainMappingsClient interface {

    // List all service chain mappings in the system for the given service profile. If no explicit enforcement point is provided in the request, will return for default. Else, will return for specified points.
    //
    // @param serviceReferenceIdParam Service reference id (required)
    // @param serviceProfileIdParam Service profile id (required)
    // @param enforcementPointPathParam String Path of the enforcement point (optional)
    // @return com.vmware.nsx_policy.model.ServiceChainMappingListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(serviceReferenceIdParam string, serviceProfileIdParam string, enforcementPointPathParam *string) (model.ServiceChainMappingListResult, error)
}
