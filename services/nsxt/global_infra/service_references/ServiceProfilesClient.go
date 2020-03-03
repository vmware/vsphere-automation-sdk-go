/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: ServiceProfiles
 * Used by client-side stubs.
 */

package service_references

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type ServiceProfilesClient interface {

    // This API can be used to delete service profile with given service-profile-id
    //
    // @param serviceReferenceIdParam Id of Service Reference (required)
    // @param serviceProfileIdParam Service profile id (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(serviceReferenceIdParam string, serviceProfileIdParam string) error

    // This API can be used to read service profile with given service-profile-id
    //
    // @param serviceReferenceIdParam Id of Service Reference (required)
    // @param serviceProfileIdParam Service profile id (required)
    // @return com.vmware.nsx_policy.model.PolicyServiceProfile
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(serviceReferenceIdParam string, serviceProfileIdParam string) (model.PolicyServiceProfile, error)

    // List all the service profiles available for given service reference
    //
    // @param serviceReferenceIdParam Service reference id (required)
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_policy.model.PolicyServiceProfileListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(serviceReferenceIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.PolicyServiceProfileListResult, error)

    // Create Service profile to specify vendor template attri- butes for a given 3rd party service.
    //
    // @param serviceReferenceIdParam Service reference id (required)
    // @param serviceProfileIdParam Service profile id (required)
    // @param policyServiceProfileParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(serviceReferenceIdParam string, serviceProfileIdParam string, policyServiceProfileParam model.PolicyServiceProfile) error

    // Create or update Service profile to specify vendor temp- late attributes for a given 3rd party service.
    //
    // @param serviceReferenceIdParam Service reference id (required)
    // @param serviceProfileIdParam Service profile id (required)
    // @param policyServiceProfileParam (required)
    // @return com.vmware.nsx_policy.model.PolicyServiceProfile
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(serviceReferenceIdParam string, serviceProfileIdParam string, policyServiceProfileParam model.PolicyServiceProfile) (model.PolicyServiceProfile, error)
}
