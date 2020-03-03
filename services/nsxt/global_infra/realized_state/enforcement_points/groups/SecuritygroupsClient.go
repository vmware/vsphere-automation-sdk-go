/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Securitygroups
 * Used by client-side stubs.
 */

package groups

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type SecuritygroupsClient interface {

    // Read a Security Group and the complete tree underneath. Returns the populated Security Group object.
    //
    // @param enforcementPointNameParam Enforcement Point Name (required)
    // @param securitygroupNameParam Group Name (required)
    // @return com.vmware.nsx_policy.model.RealizedSecurityGroup
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(enforcementPointNameParam string, securitygroupNameParam string) (model.RealizedSecurityGroup, error)

    // Paginated list of all Security Groups. Returns populated Security Groups.
    //
    // @param enforcementPointNameParam Enforcement Point Name (required)
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_policy.model.RealizedSecurityGroupListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(enforcementPointNameParam string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.RealizedSecurityGroupListResult, error)
}
