/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: UserStats
 * Used by client-side stubs.
 */

package idfw

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type UserStatsClient interface {

    // It will get IDFW user login events for a given user.
    //
    // @param userIdParam User id (required)
    // @param enforcementPointPathParam String Path of the enforcement point (optional)
    // @return com.vmware.nsx_policy.model.IdfwUserStats
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(userIdParam string, enforcementPointPathParam *string) (model.IdfwUserStats, error)
}
