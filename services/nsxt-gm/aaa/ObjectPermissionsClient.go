/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: ObjectPermissions
 * Used by client-side stubs.
 */

package aaa

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type ObjectPermissionsClient interface {

    // Delete object-permissions entries
    //
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param inheritanceDisabledParam Does children of this object inherit this rule (optional, default to false)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param pathPrefixParam Path prefix (optional)
    // @param roleNameParam Role name (optional)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, inheritanceDisabledParam *bool, pageSizeParam *int64, pathPrefixParam *string, roleNameParam *string, sortAscendingParam *bool, sortByParam *string) error

    // Get list of Object-level RBAC entries.
    //
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param inheritanceDisabledParam Does children of this object inherit this rule (optional, default to false)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param pathPrefixParam Path prefix (optional)
    // @param roleNameParam Role name (optional)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_global_policy.model.ObjectRolePermissionGroupListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, inheritanceDisabledParam *bool, pageSizeParam *int64, pathPrefixParam *string, roleNameParam *string, sortAscendingParam *bool, sortByParam *string) (model.ObjectRolePermissionGroupListResult, error)

    // Create/update object permission mappings
    //
    // @param objectRolePermissionGroupParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(objectRolePermissionGroupParam model.ObjectRolePermissionGroup) error
}
