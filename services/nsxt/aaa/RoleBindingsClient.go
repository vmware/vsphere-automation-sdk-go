/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: RoleBindings
 * Used by client-side stubs.
 */

package aaa

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type RoleBindingsClient interface {

    // When assigning a user role, specify the user name with the same case as it appears in vIDM to access the NSX-T user interface. For example, if vIDM has the user name User1\\\\@example.com then the name attribute in the API call must be be User1\\\\@example.com and cannot be user1\\\\@example.com.
    //
    // @param roleBindingParam (required)
    // @return com.vmware.nsx_policy.model.RoleBinding
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Create(roleBindingParam model.RoleBinding) (model.RoleBinding, error)

    // Delete user/group's roles assignment
    //
    // @param bindingIdParam User/Group's id (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(bindingIdParam string) error

    // Delete all stale role assignments
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Deletestalebindings() error

    // Get user/group's role information
    //
    // @param bindingIdParam User/Group's id (required)
    // @return com.vmware.nsx_policy.model.RoleBinding
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(bindingIdParam string) (model.RoleBinding, error)

    // Get all users and groups with their roles
    //
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param identitySourceIdParam Identity source ID (optional)
    // @param identitySourceTypeParam Identity source type (optional)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param nameParam User/Group name (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param roleParam Role ID (optional)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @param type_Param Type (optional)
    // @return com.vmware.nsx_policy.model.RoleBindingListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(cursorParam *string, identitySourceIdParam *string, identitySourceTypeParam *string, includedFieldsParam *string, nameParam *string, pageSizeParam *int64, roleParam *string, sortAscendingParam *bool, sortByParam *string, type_Param *string) (model.RoleBindingListResult, error)

    // Update User or Group's roles
    //
    // @param bindingIdParam User/Group's id (required)
    // @param roleBindingParam (required)
    // @return com.vmware.nsx_policy.model.RoleBinding
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(bindingIdParam string, roleBindingParam model.RoleBinding) (model.RoleBinding, error)
}
