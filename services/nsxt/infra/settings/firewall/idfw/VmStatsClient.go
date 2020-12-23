/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: VmStats
 * Used by client-side stubs.
 */

package idfw

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type VmStatsClient interface {

    // It will get IDFW user login events for a given VM (all active plus up to 5 most recent archived entries).
    //
    // @param vmIdParam VM id (required)
    // @param enforcementPointPathParam String Path of the enforcement point (optional)
    // @return com.vmware.nsx_policy.model.IdfwVmStats
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(vmIdParam string, enforcementPointPathParam *string) (model.IdfwVmStats, error)
}
