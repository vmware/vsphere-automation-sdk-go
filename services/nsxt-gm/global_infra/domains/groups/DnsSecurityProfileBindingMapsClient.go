/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: DnsSecurityProfileBindingMaps
 * Used by client-side stubs.
 */

package groups

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type DnsSecurityProfileBindingMapsClient interface {

    // API will delete DNS security profile binding map
    //
    // @param domainIdParam Domain ID (required)
    // @param groupIdParam Group ID (required)
    // @param dnsSecurityProfileBindingMapIdParam DNS security profile binding map ID (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(domainIdParam string, groupIdParam string, dnsSecurityProfileBindingMapIdParam string) error

    // API will get DNS security profile binding map
    //
    // @param domainIdParam Domain ID (required)
    // @param groupIdParam Group ID (required)
    // @param dnsSecurityProfileBindingMapIdParam DNS security profile binding map ID (required)
    // @return com.vmware.nsx_global_policy.model.DnsSecurityProfileBindingMap
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(domainIdParam string, groupIdParam string, dnsSecurityProfileBindingMapIdParam string) (model.DnsSecurityProfileBindingMap, error)

    // API will get DNS security profile binding map
    //
    // @param domainIdParam (required)
    // @param groupIdParam (required)
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_global_policy.model.DnsSecurityProfileBindingMapListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(domainIdParam string, groupIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.DnsSecurityProfileBindingMapListResult, error)

    // API will create or update DNS security profile binding map
    //
    // @param domainIdParam Domain ID (required)
    // @param groupIdParam Group ID (required)
    // @param dnsSecurityProfileBindingMapIdParam DNS security profile binding map ID (required)
    // @param dnsSecurityProfileBindingMapParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(domainIdParam string, groupIdParam string, dnsSecurityProfileBindingMapIdParam string, dnsSecurityProfileBindingMapParam model.DnsSecurityProfileBindingMap) error

    // API will update DNS security profile binding map
    //
    // @param domainIdParam Domain ID (required)
    // @param groupIdParam Group ID (required)
    // @param dnsSecurityProfileBindingMapIdParam DNS security profile binding map ID (required)
    // @param dnsSecurityProfileBindingMapParam (required)
    // @return com.vmware.nsx_global_policy.model.DnsSecurityProfileBindingMap
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(domainIdParam string, groupIdParam string, dnsSecurityProfileBindingMapIdParam string, dnsSecurityProfileBindingMapParam model.DnsSecurityProfileBindingMap) (model.DnsSecurityProfileBindingMap, error)
}
