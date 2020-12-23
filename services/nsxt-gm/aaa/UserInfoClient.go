/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: UserInfo
 * Used by client-side stubs.
 */

package aaa

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type UserInfoClient interface {

    // This API will return the name and role information of the user invoking this API request. This API is available for all NSX users no matter their authentication method (Local account, VIDM, LDAP etc). The permissions parameter of the NsxRole has been deprecated.
    // @return com.vmware.nsx_global_policy.model.UserInfo
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get() (model.UserInfo, error)
}
