/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: PortSecurityProfileBindingMaps
 * Used by client-side stubs.
 */

package ports

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type PortSecurityProfileBindingMapsClient interface {

    // API will delete the port security profile binding map.
    //
    // @param tier1IdParam tier-1 gateway id (required)
    // @param segmentIdParam segment id (required)
    // @param portIdParam port id (required)
    // @param portSecurityProfileBindingMapIdParam port security profile binding map id (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(tier1IdParam string, segmentIdParam string, portIdParam string, portSecurityProfileBindingMapIdParam string) error

    // API will return details of the port security profile binding map. If the security profile binding map does not exist, it will return 404.
    //
    // @param tier1IdParam tier-1 gateway id (required)
    // @param segmentIdParam segment id (required)
    // @param portIdParam port id (required)
    // @param portSecurityProfileBindingMapIdParam port security profile binding map id (required)
    // @return com.vmware.nsx_policy.model.PortSecurityProfileBindingMap
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(tier1IdParam string, segmentIdParam string, portIdParam string, portSecurityProfileBindingMapIdParam string) (model.PortSecurityProfileBindingMap, error)

    // API will list all port security profile binding maps.
    //
    // @param tier1IdParam tier-1 gateway id (required)
    // @param segmentIdParam segment id (required)
    // @param portIdParam port id (required)
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_policy.model.PortSecurityProfileBindingMapListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(tier1IdParam string, segmentIdParam string, portIdParam string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.PortSecurityProfileBindingMapListResult, error)

    // Create a new port security profile binding map if the given security profile binding map does not exist. Otherwise, patch the existing port security profile binding map.
    //
    // @param tier1IdParam tier-1 gateway id (required)
    // @param segmentIdParam segment id (required)
    // @param portIdParam port id (required)
    // @param portSecurityProfileBindingMapIdParam port security profile binding map id (required)
    // @param portSecurityProfileBindingMapParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(tier1IdParam string, segmentIdParam string, portIdParam string, portSecurityProfileBindingMapIdParam string, portSecurityProfileBindingMapParam model.PortSecurityProfileBindingMap) error

    // API will create or replace the port security profile binding map.
    //
    // @param tier1IdParam tier-1 gateway id (required)
    // @param segmentIdParam segment id (required)
    // @param portIdParam port id (required)
    // @param portSecurityProfileBindingMapIdParam port security profile binding map id (required)
    // @param portSecurityProfileBindingMapParam (required)
    // @return com.vmware.nsx_policy.model.PortSecurityProfileBindingMap
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(tier1IdParam string, segmentIdParam string, portIdParam string, portSecurityProfileBindingMapIdParam string, portSecurityProfileBindingMapParam model.PortSecurityProfileBindingMap) (model.PortSecurityProfileBindingMap, error)
}
