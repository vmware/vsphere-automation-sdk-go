/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: ArpProxy
 * Used by client-side stubs.
 */

package interfaces

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type ArpProxyClient interface {

    // Returns ARP proxy table for a tier-0 interface. Interfaces can be of types - EXTERNAL and SERVICE. Interfaces of type LOOBACK and downlink are not supported.
    //
    // @param tier0IdParam (required)
    // @param localeServiceIdParam (required)
    // @param interfaceIdParam (required)
    // @param enforcementPointPathParam String Path of the enforcement point (optional)
    // @param sourceParam Data source type. (optional)
    // @return com.vmware.nsx_global_policy.model.PolicyArpProxyTableListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(tier0IdParam string, localeServiceIdParam string, interfaceIdParam string, enforcementPointPathParam *string, sourceParam *string) (model.PolicyArpProxyTableListResult, error)
}
