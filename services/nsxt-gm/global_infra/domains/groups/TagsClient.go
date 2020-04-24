/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Tags
 * Used by client-side stubs.
 */

package groups

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type TagsClient interface {

    // Get tags used to define conditions inside a Group. Also includes tags inside nested groups.
    //
    // @param domainIdParam Domain id (required)
    // @param groupIdParam Group Id (required)
    // @return com.vmware.nsx_global_policy.model.GroupTagsList
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(domainIdParam string, groupIdParam string) (model.GroupTagsList, error)
}
