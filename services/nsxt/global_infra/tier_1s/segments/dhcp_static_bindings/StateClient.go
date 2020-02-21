/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: State
 * Used by client-side stubs.
 */

package dhcp_static_bindings

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type StateClient interface {

    // Read DHCP static binding state
    //
    // @param tier1IdParam (required)
    // @param segmentIdParam (required)
    // @param bindingIdParam (required)
    // @return com.vmware.nsx_policy.model.DhcpStaticBindingState
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(tier1IdParam string, segmentIdParam string, bindingIdParam string) (model.DhcpStaticBindingState, error)
}
