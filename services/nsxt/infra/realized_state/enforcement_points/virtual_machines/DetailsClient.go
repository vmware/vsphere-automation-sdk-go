/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Details
 * Used by client-side stubs.
 */

package virtual_machines

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type DetailsClient interface {

    // This API return optional details about a virtual machines (e.g. user login session) from the specified enforcement point. In case of NSXT, virtual-machine-id would be the value of the external_id of the virtual machine.
    //
    // @param enforcementPointNameParam (required)
    // @param virtualMachineIdParam (required)
    // @return com.vmware.nsx_policy.model.VirtualMachineDetails
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(enforcementPointNameParam string, virtualMachineIdParam string) (model.VirtualMachineDetails, error)
}
