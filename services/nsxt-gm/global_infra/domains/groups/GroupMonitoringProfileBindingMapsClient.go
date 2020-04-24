/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: GroupMonitoringProfileBindingMaps
 * Used by client-side stubs.
 */

package groups

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type GroupMonitoringProfileBindingMapsClient interface {

    // API will delete Group Monitoring Profile Binding
    //
    // @param domainIdParam Domain ID (required)
    // @param groupIdParam Group ID (required)
    // @param groupMonitoringProfileBindingMapIdParam Group Monitoring Profile Binding Map ID (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(domainIdParam string, groupIdParam string, groupMonitoringProfileBindingMapIdParam string) error

    // API will get Group Monitoring Profile Binding Map
    //
    // @param domainIdParam Domain-ID (required)
    // @param groupIdParam Group ID (required)
    // @param groupMonitoringProfileBindingMapIdParam Group Monitoring Profile Binding Map ID (required)
    // @return com.vmware.nsx_global_policy.model.GroupMonitoringProfileBindingMap
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(domainIdParam string, groupIdParam string, groupMonitoringProfileBindingMapIdParam string) (model.GroupMonitoringProfileBindingMap, error)

    // API will list all Group Monitoring Profile Binding Maps in current group id.
    //
    // @param domainIdParam (required)
    // @param groupIdParam (required)
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_global_policy.model.GroupMonitoringProfileBindingMapListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(domainIdParam string, groupIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.GroupMonitoringProfileBindingMapListResult, error)

    // API will create group monitoring profile binding map
    //
    // @param domainIdParam Domain ID (required)
    // @param groupIdParam Group ID (required)
    // @param groupMonitoringProfileBindingMapIdParam Group Monitoring Profile Binding Map ID (required)
    // @param groupMonitoringProfileBindingMapParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(domainIdParam string, groupIdParam string, groupMonitoringProfileBindingMapIdParam string, groupMonitoringProfileBindingMapParam model.GroupMonitoringProfileBindingMap) error

    // API will update Group Monitoring Profile Binding Map
    //
    // @param domainIdParam DomainID (required)
    // @param groupIdParam Group ID (required)
    // @param groupMonitoringProfileBindingMapIdParam Group Monitoring Profile Binding Map ID (required)
    // @param groupMonitoringProfileBindingMapParam (required)
    // @return com.vmware.nsx_global_policy.model.GroupMonitoringProfileBindingMap
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(domainIdParam string, groupIdParam string, groupMonitoringProfileBindingMapIdParam string, groupMonitoringProfileBindingMapParam model.GroupMonitoringProfileBindingMap) (model.GroupMonitoringProfileBindingMap, error)
}
