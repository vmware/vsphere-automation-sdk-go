/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: EnforcementPoints
 * Used by client-side stubs.
 */

package deployment_zones

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type EnforcementPointsClient interface {

    // Delete EnforcementPoint. This is a deprecated API. DeploymentZone has been renamed to Site. Use DELETE /infra/sites/site-id/enforcement-points/enforcementpoint-id.
    //
    // @param deploymentZoneIdParam (required)
    // @param enforcementpointIdParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(deploymentZoneIdParam string, enforcementpointIdParam string) error

    // Read an Enforcement Point. This is a deprecated API. DeploymentZone has been renamed to Site. Use GET /infra/sites/site-id/enforcement-points/enforcementpoint-id.
    //
    // @param deploymentZoneIdParam (required)
    // @param enforcementpointIdParam (required)
    // @return com.vmware.nsx_policy.model.EnforcementPoint
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(deploymentZoneIdParam string, enforcementpointIdParam string) (model.EnforcementPoint, error)

    // Paginated list of all enforcementpoints for infra. This is a deprecated API. DeploymentZone has been renamed to Site. Use GET /infra/sites/site-id/enforcement-points.
    //
    // @param deploymentZoneIdParam (required)
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_policy.model.EnforcementPointListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(deploymentZoneIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.EnforcementPointListResult, error)

    // If the passed Enforcement Point does not already exist, create a new Enforcement Point. If it already exists, patch it. This is a deprecated API. DeploymentZone has been renamed to Site. Use PATCH /infra/sites/site-1/enforcement-points/enforcementpoint-1.
    //
    // @param deploymentZoneIdParam (required)
    // @param enforcementpointIdParam (required)
    // @param enforcementPointParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(deploymentZoneIdParam string, enforcementpointIdParam string, enforcementPointParam model.EnforcementPoint) error

    // If the passed Enforcement Point does not already exist, create a new Enforcement Point. If it already exists, replace it. This is a deprecated API. DeploymentZone has been renamed to Site. Use PUT /infra/sites/site-id/enforcement-points/enforcementpoint-id.
    //
    // @param deploymentZoneIdParam (required)
    // @param enforcementpointIdParam (required)
    // @param enforcementPointParam (required)
    // @return com.vmware.nsx_policy.model.EnforcementPoint
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(deploymentZoneIdParam string, enforcementpointIdParam string, enforcementPointParam model.EnforcementPoint) (model.EnforcementPoint, error)
}
