/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Status
 * Used by client-side stubs.
 */

package vms

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type StatusClient interface {

    // This API will list all VMs and statuses based on transport node ID of idfw enabled compute collection.
    //
    // @param transportNodeIdParam Transport node id (required)
    // @param enforcementPointPathParam String Path of the enforcement point (optional)
    // @return com.vmware.nsx_policy.model.IdfwVirtualMachineStatusListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(transportNodeIdParam string, enforcementPointPathParam *string) (model.IdfwVirtualMachineStatusListResult, error)
}
