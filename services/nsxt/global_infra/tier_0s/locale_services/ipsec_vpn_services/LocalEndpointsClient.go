/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: LocalEndpoints
 * Used by client-side stubs.
 */

package ipsec_vpn_services

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type LocalEndpointsClient interface {

    // Delete IPSec VPN local endpoint for a given locale service under Tier-0.
    //
    // @param tier0IdParam (required)
    // @param localeServiceIdParam (required)
    // @param serviceIdParam (required)
    // @param localEndpointIdParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(tier0IdParam string, localeServiceIdParam string, serviceIdParam string, localEndpointIdParam string) error

    // Get IPSec VPN local endpoint for a given locale service under Tier-0.
    //
    // @param tier0IdParam (required)
    // @param localeServiceIdParam (required)
    // @param serviceIdParam (required)
    // @param localEndpointIdParam (required)
    // @return com.vmware.nsx_policy.model.IPSecVpnLocalEndpoint
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(tier0IdParam string, localeServiceIdParam string, serviceIdParam string, localEndpointIdParam string) (model.IPSecVpnLocalEndpoint, error)

    // Get paginated list of all IPSec VPN local endpoints for a given locale service under Tier-0.
    //
    // @param tier0IdParam (required)
    // @param localeServiceIdParam (required)
    // @param serviceIdParam (required)
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_policy.model.IPSecVpnLocalEndpointListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(tier0IdParam string, localeServiceIdParam string, serviceIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.IPSecVpnLocalEndpointListResult, error)

    // Create or patch a custom IPSec VPN local endpoint for a given locale service under Tier-0.
    //
    // @param tier0IdParam (required)
    // @param localeServiceIdParam (required)
    // @param serviceIdParam (required)
    // @param localEndpointIdParam (required)
    // @param ipSecVpnLocalEndpointParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(tier0IdParam string, localeServiceIdParam string, serviceIdParam string, localEndpointIdParam string, ipSecVpnLocalEndpointParam model.IPSecVpnLocalEndpoint) error

    // Create or fully replace IPSec VPN local endpoint for a given locale service under Tier-0. Revision is optional for creation and required for update.
    //
    // @param tier0IdParam (required)
    // @param localeServiceIdParam (required)
    // @param serviceIdParam (required)
    // @param localEndpointIdParam (required)
    // @param ipSecVpnLocalEndpointParam (required)
    // @return com.vmware.nsx_policy.model.IPSecVpnLocalEndpoint
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(tier0IdParam string, localeServiceIdParam string, serviceIdParam string, localEndpointIdParam string, ipSecVpnLocalEndpointParam model.IPSecVpnLocalEndpoint) (model.IPSecVpnLocalEndpoint, error)
}
