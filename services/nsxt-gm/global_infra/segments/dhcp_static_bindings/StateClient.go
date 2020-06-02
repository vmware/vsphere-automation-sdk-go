/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: State
 * Used by client-side stubs.
 */

package dhcp_static_bindings

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type StateClient interface {

    // Read DHCP static binding state
    //
    // @param segmentIdParam (required)
    // @param bindingIdParam (required)
    // @return com.vmware.nsx_global_policy.model.DhcpStaticBindingState
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(segmentIdParam string, bindingIdParam string) (model.DhcpStaticBindingState, error)
}
