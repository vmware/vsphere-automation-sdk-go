/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: PortMirroringInstances
 * Used by client-side stubs.
 */

package groups

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type PortMirroringInstancesClient interface {

    // API will delete port mirroring instance. Mirroring from source to destination ports will be stopped. This API is deprecated. Please use the following API: https://<policy-mgr>/policy/api/v1/infra/port-mirroring-profiles
    //
    // @param domainIdParam Domain ID (required)
    // @param groupIdParam Group ID (required)
    // @param portMirroringInstanceIdParam Port Mirroring Instance Id (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(domainIdParam string, groupIdParam string, portMirroringInstanceIdParam string) error

    // API will return details of port mirroring instance. If instance does not exist, it will return 404. This API is deprecated. Please use the following API: https://<policy-mgr>/policy/api/v1/infra/port-mirroring-profiles
    //
    // @param domainIdParam Domain ID (required)
    // @param groupIdParam Group ID (required)
    // @param portMirroringInstanceIdParam Port Mirroring Instance Id (required)
    // @return com.vmware.nsx_policy.model.PortMirroringInstance
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(domainIdParam string, groupIdParam string, portMirroringInstanceIdParam string) (model.PortMirroringInstance, error)

    // API will list all port mirroring instances active on current group. This API is deprecated. Please use the following API: https://<policy-mgr>/policy/api/v1/infra/port-mirroring-profiles
    //
    // @param domainIdParam Domain ID (required)
    // @param groupIdParam Group ID (required)
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_policy.model.PortMirroringInstanceListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(domainIdParam string, groupIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.PortMirroringInstanceListResult, error)

    // Create a new Port Mirroring Instance if the Port Mirroring Instance with given id does not already exist. If the Port Mirroring Instance with the given id already exists, patch with the existing Port Mirroring Instance. This API is deprecated. Please use the following API: https://<policy-mgr>/policy/api/v1/infra/port-mirroring-profiles
    //
    // @param domainIdParam Domain ID (required)
    // @param groupIdParam Group ID (required)
    // @param portMirroringInstanceIdParam Port Mirroring Instance Id (required)
    // @param portMirroringInstanceParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(domainIdParam string, groupIdParam string, portMirroringInstanceIdParam string, portMirroringInstanceParam model.PortMirroringInstance) error

    // Create or Replace port mirroring instance. Packets will be mirrored from source group to destination group. This API is deprecated. Please use the following API: https://<policy-mgr>/policy/api/v1/infra/port-mirroring-profiles
    //
    // @param domainIdParam Domain ID (required)
    // @param groupIdParam Group ID (required)
    // @param portMirroringInstanceIdParam Port Mirroring Instance Id (required)
    // @param portMirroringInstanceParam (required)
    // @return com.vmware.nsx_policy.model.PortMirroringInstance
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(domainIdParam string, groupIdParam string, portMirroringInstanceIdParam string, portMirroringInstanceParam model.PortMirroringInstance) (model.PortMirroringInstance, error)
}
