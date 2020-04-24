/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: EnforcementPoints
 * Used by client-side stubs.
 */

package sites

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type EnforcementPointsClient interface {

    // Delete EnforcementPoint from Site
    //
    // @param siteIdParam (required)
    // @param enforcementpointIdParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(siteIdParam string, enforcementpointIdParam string) error

    // Full sync EnforcementPoint from Site
    //
    // @param siteIdParam (required)
    // @param enforcementPointIdParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Fullsync(siteIdParam string, enforcementPointIdParam string) error

    // Read an Enforcement Point under Infra/Site
    //
    // @param siteIdParam (required)
    // @param enforcementpointIdParam (required)
    // @return com.vmware.nsx_global_policy.model.EnforcementPoint
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(siteIdParam string, enforcementpointIdParam string) (model.EnforcementPoint, error)

    // Paginated list of all enforcementpoints under Site.
    //
    // @param siteIdParam (required)
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_global_policy.model.EnforcementPointListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(siteIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.EnforcementPointListResult, error)

    // If the passed Enforcement Point does not already exist, create a new Enforcement Point. If it already exists, patch it.
    //
    // @param siteIdParam (required)
    // @param enforcementpointIdParam (required)
    // @param enforcementPointParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(siteIdParam string, enforcementpointIdParam string, enforcementPointParam model.EnforcementPoint) error

    // Reload an Enforcement Point under Site. This will read and update fabric configs from enforcement point.
    //
    // @param siteIdParam (required)
    // @param enforcementpointIdParam (required)
    // @return com.vmware.nsx_global_policy.model.EnforcementPoint
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Reload(siteIdParam string, enforcementpointIdParam string) (model.EnforcementPoint, error)

    // If the passed Enforcement Point does not already exist, create a new Enforcement Point. If it already exists, replace it.
    //
    // @param siteIdParam (required)
    // @param enforcementpointIdParam (required)
    // @param enforcementPointParam (required)
    // @return com.vmware.nsx_global_policy.model.EnforcementPoint
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(siteIdParam string, enforcementpointIdParam string, enforcementPointParam model.EnforcementPoint) (model.EnforcementPoint, error)
}
