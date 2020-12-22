/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: ConsolidatedEffectiveIpAddresses
 * Used by client-side stubs.
 */

package members

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type ConsolidatedEffectiveIpAddressesClient interface {

    // Returns consolidated effective ip address members of the specified NSGroup. Applicable in case of federated environment. The response contains site-wise list of consolidated effective IP address members. In the response, for the local-site, the list will contain static and dynamicaly translated IPs. For the remote sites, the list will contain only the dynamically translated IPs. The static IPs will not be seen in the response of this API. Hence, user can refer to the local-site Ip response in the API results or the group definition to see the static IP membership of the Group. This API is applicable only for Global Groups containing (directly or via nesting) either VirtualMachine, VIF, Segment, SegmentPort or IPSet member type. Use the cursor value in the response to fetch the next page. If there is no cursor value for a response, it implies the last page in the results for the query.
    //
    // @param domainIdParam Domain id (required)
    // @param groupIdParam Group Id (required)
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param enforcementPointPathParam String Path of the enforcement point (optional)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param ipFilterParam IP address, range, or subnet (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param siteIdParam UUID of the site from which the effective IP addresses are to be fetched (optional)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_global_policy.model.ConsolidatedEffectiveIPAddressMemberListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(domainIdParam string, groupIdParam string, cursorParam *string, enforcementPointPathParam *string, includedFieldsParam *string, ipFilterParam *string, pageSizeParam *int64, siteIdParam *string, sortAscendingParam *bool, sortByParam *string) (model.ConsolidatedEffectiveIPAddressMemberListResult, error)
}
