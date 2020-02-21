/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Sessions
 * Used by client-side stubs.
 */

package l2vpn_services

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type SessionsClient interface {

    // Create or patch an L2VPN session under Tier-0 from Peer Codes. In addition to the L2VPN Session, the IPSec VPN Session, along with the IKE, Tunnel, and DPD Profiles are created and owned by the system. IPSec VPN Service and Local Endpoint are created only when required, i.e., an IPSec VPN Service does not already exist, or an IPSec VPN Local Endpoint with same local address does not already exist. Updating the L2VPN Session can be performed only through this API by specifying new peer codes. Use of specific APIs to update the L2VPN Session and the different resources associated with it is not allowed, except for IPSec VPN Service and Local Endpoint, resources that are not system owned. API supported only when L2VPN Service is in Client Mode.
    //
    // @param tier0IdParam (required)
    // @param localeServiceIdParam (required)
    // @param serviceIdParam (required)
    // @param sessionIdParam (required)
    // @param l2VPNSessionDataParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Createwithpeercode(tier0IdParam string, localeServiceIdParam string, serviceIdParam string, sessionIdParam string, l2VPNSessionDataParam model.L2VPNSessionData) error

    // Delete L2VPN session under Tier-0. When L2VPN Service is in CLIENT Mode, the L2VPN Session is deleted along with its transpot tunnels and related resources.
    //
    // @param tier0IdParam (required)
    // @param localeServiceIdParam (required)
    // @param serviceIdParam (required)
    // @param sessionIdParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(tier0IdParam string, localeServiceIdParam string, serviceIdParam string, sessionIdParam string) error

    // Get L2VPN session under Tier-0.
    //
    // @param tier0IdParam (required)
    // @param localeServiceIdParam (required)
    // @param serviceIdParam (required)
    // @param sessionIdParam (required)
    // @return com.vmware.nsx_policy.model.L2VPNSession
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(tier0IdParam string, localeServiceIdParam string, serviceIdParam string, sessionIdParam string) (model.L2VPNSession, error)

    // Get paginated list of all L2VPN sessions under Tier-0.
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
    // @return com.vmware.nsx_policy.model.L2VPNSessionListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(tier0IdParam string, localeServiceIdParam string, serviceIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.L2VPNSessionListResult, error)

    // Create or patch an L2VPN session under Tier-0. API supported only when L2VPN Service is in Server Mode.
    //
    // @param tier0IdParam (required)
    // @param localeServiceIdParam (required)
    // @param serviceIdParam (required)
    // @param sessionIdParam (required)
    // @param l2VPNSessionParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(tier0IdParam string, localeServiceIdParam string, serviceIdParam string, sessionIdParam string, l2VPNSessionParam model.L2VPNSession) error

    // Create or fully replace L2VPN session under Tier-0. API supported only when L2VPN Service is in Server Mode. Revision is optional for creation and required for update.
    //
    // @param tier0IdParam (required)
    // @param localeServiceIdParam (required)
    // @param serviceIdParam (required)
    // @param sessionIdParam (required)
    // @param l2VPNSessionParam (required)
    // @return com.vmware.nsx_policy.model.L2VPNSession
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(tier0IdParam string, localeServiceIdParam string, serviceIdParam string, sessionIdParam string, l2VPNSessionParam model.L2VPNSession) (model.L2VPNSession, error)
}
