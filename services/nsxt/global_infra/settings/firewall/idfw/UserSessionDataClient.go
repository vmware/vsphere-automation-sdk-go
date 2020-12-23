/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: UserSessionData
 * Used by client-side stubs.
 */

package idfw

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type UserSessionDataClient interface {

    // It will get user session data.
    //
    // @param enforcementPointPathParam String Path of the enforcement point (optional)
    // @return com.vmware.nsx_policy.model.IdfwUserSessionDataAndMappings
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(enforcementPointPathParam *string) (model.IdfwUserSessionDataAndMappings, error)
}
