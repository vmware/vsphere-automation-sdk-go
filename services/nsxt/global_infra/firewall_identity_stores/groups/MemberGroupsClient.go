/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: MemberGroups
 * Used by client-side stubs.
 */

package groups

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type MemberGroupsClient interface {

    // A member group could be either direct member of the group specified by group_id or nested member of it. Both direct member groups and nested member groups are returned.
    //
    // @param firewallIdentityStoreIdParam Firewall Identity store identifier (required)
    // @param groupIdParam Directory group identifier (required)
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param enforcementPointPathParam String Path of the enforcement point (optional)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_policy.model.DirectoryGroupMemberListResults
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(firewallIdentityStoreIdParam string, groupIdParam string, cursorParam *string, enforcementPointPathParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.DirectoryGroupMemberListResults, error)
}
