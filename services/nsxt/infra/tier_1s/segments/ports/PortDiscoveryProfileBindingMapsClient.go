/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: PortDiscoveryProfileBindingMaps
 * Used by client-side stubs.
 */

package ports

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type PortDiscoveryProfileBindingMapsClient interface {

    // API will delete Port Discovery Profile Binding Profile
    //
    // @param tier1IdParam Tier-1 ID (required)
    // @param segmentIdParam Segment ID (required)
    // @param portIdParam Port ID (required)
    // @param portDiscoveryProfileBindingMapIdParam Port Discovery Profile Binding Map ID (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(tier1IdParam string, segmentIdParam string, portIdParam string, portDiscoveryProfileBindingMapIdParam string) error

    // API will get Port Discovery Profile Binding Map
    //
    // @param tier1IdParam Tier-1 ID (required)
    // @param segmentIdParam Segment ID (required)
    // @param portIdParam Port ID (required)
    // @param portDiscoveryProfileBindingMapIdParam Port Discovery Profile Binding Map ID (required)
    // @return com.vmware.nsx_policy.model.PortDiscoveryProfileBindingMap
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(tier1IdParam string, segmentIdParam string, portIdParam string, portDiscoveryProfileBindingMapIdParam string) (model.PortDiscoveryProfileBindingMap, error)

    // API will list all Port Discovery Profile Binding Maps in current port id.
    //
    // @param tier1IdParam (required)
    // @param segmentIdParam (required)
    // @param portIdParam (required)
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_policy.model.PortDiscoveryProfileBindingMapListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(tier1IdParam string, segmentIdParam string, portIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.PortDiscoveryProfileBindingMapListResult, error)

    // API will create Port Discovery Profile Binding Map
    //
    // @param tier1IdParam Tier-1 ID (required)
    // @param segmentIdParam Segment ID (required)
    // @param portIdParam Port ID (required)
    // @param portDiscoveryProfileBindingMapIdParam Port Discovery Profile Binding Map ID (required)
    // @param portDiscoveryProfileBindingMapParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(tier1IdParam string, segmentIdParam string, portIdParam string, portDiscoveryProfileBindingMapIdParam string, portDiscoveryProfileBindingMapParam model.PortDiscoveryProfileBindingMap) error

    // API will update Port Discovery Profile Binding Map
    //
    // @param tier1IdParam Tier-1 ID (required)
    // @param segmentIdParam Segment ID (required)
    // @param portIdParam Port ID (required)
    // @param portDiscoveryProfileBindingMapIdParam Port Discovery Profile Binding Map ID (required)
    // @param portDiscoveryProfileBindingMapParam (required)
    // @return com.vmware.nsx_policy.model.PortDiscoveryProfileBindingMap
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(tier1IdParam string, segmentIdParam string, portIdParam string, portDiscoveryProfileBindingMapIdParam string, portDiscoveryProfileBindingMapParam model.PortDiscoveryProfileBindingMap) (model.PortDiscoveryProfileBindingMap, error)
}
