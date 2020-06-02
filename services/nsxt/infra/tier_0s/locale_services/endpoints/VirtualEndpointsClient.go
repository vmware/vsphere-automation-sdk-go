/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: VirtualEndpoints
 * Used by client-side stubs.
 */

package endpoints

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type VirtualEndpointsClient interface {

    // Delete virtual endpoint
    //
    // @param tier0IdParam Tier-0 id (required)
    // @param localeServiceIdParam Locale service id (required)
    // @param virtualEndpointIdParam Virtual endpoint id (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(tier0IdParam string, localeServiceIdParam string, virtualEndpointIdParam string) error

    // Read virtual endpoint with given id under given Tier0.
    //
    // @param tier0IdParam Tier-0 id (required)
    // @param localeServiceIdParam Locale service id (required)
    // @param virtualEndpointIdParam Virtual endpoint id (required)
    // @return com.vmware.nsx_policy.model.VirtualEndpoint
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(tier0IdParam string, localeServiceIdParam string, virtualEndpointIdParam string) (model.VirtualEndpoint, error)

    // List all virtual endpoints
    //
    // @param tier0IdParam Tier-0 id (required)
    // @param localeServiceIdParam Locale service id (required)
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_policy.model.VirtualEndpointListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(tier0IdParam string, localeServiceIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.VirtualEndpointListResult, error)

    // Create or update virtual endpoint.
    //
    // @param tier0IdParam Tier-0 id (required)
    // @param localeServiceIdParam Locale service id (required)
    // @param virtualEndpointIdParam Virtual endpoint id (required)
    // @param virtualEndpointParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(tier0IdParam string, localeServiceIdParam string, virtualEndpointIdParam string, virtualEndpointParam model.VirtualEndpoint) error

    // Create or update virtual endpoint.
    //
    // @param tier0IdParam Tier-0 id (required)
    // @param localeServiceIdParam Locale service id (required)
    // @param virtualEndpointIdParam Virtual endpoint id (required)
    // @param virtualEndpointParam (required)
    // @return com.vmware.nsx_policy.model.VirtualEndpoint
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(tier0IdParam string, localeServiceIdParam string, virtualEndpointIdParam string, virtualEndpointParam model.VirtualEndpoint) (model.VirtualEndpoint, error)
}
