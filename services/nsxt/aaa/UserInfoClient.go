/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: UserInfo
 * Used by client-side stubs.
 */

package aaa

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type UserInfoClient interface {

    // Get information about logged-in user. The permissions parameter of the NsxRole has been deprecated.
    // @return com.vmware.nsx_policy.model.UserInfo
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get() (model.UserInfo, error)
}
