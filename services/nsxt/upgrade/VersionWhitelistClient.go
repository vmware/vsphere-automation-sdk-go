/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: VersionWhitelist
 * Used by client-side stubs.
 */

package upgrade

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type VersionWhitelistClient interface {

    // Get whitelist of versions for a component. Component can include HOST, EDGE, CCP, MP
    //
    // @param componentTypeParam (required)
    // @return com.vmware.nsx_policy.model.AcceptableComponentVersion
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(componentTypeParam string) (model.AcceptableComponentVersion, error)

    // Get whitelist of versions for different components
    // @return com.vmware.nsx_policy.model.AcceptableComponentVersionList
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List() (model.AcceptableComponentVersionList, error)

    // Update the version whitelist for the specified component type (HOST, EDGE, CCP, MP).
    //
    // @param componentTypeParam (required)
    // @param versionListParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(componentTypeParam string, versionListParam model.VersionList) error
}
