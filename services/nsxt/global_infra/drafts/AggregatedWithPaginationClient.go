/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Aggregated_with_pagination
 * Used by client-side stubs.
 */

package drafts

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type Aggregated_with_paginationClient interface {

    // Get a paginated aggregated configuration of a given draft. This aggregated configuration is the differnece between the current published firewall configuration and a firewall configuration stored in a given draft. For an initial API call, if request_id is present in a response, then this is a paginated aggregated configuration of a given draft, containing all the security policies from the aggregated configuration. Using this request_id, more granular aggregated configuration, at security policy level, can be fetched from subsequent API calls. Absence of request_id suggests that whole aggregated configuration has been returned as a response to initial API call, as the size of aggregated configuration is not big enough to need pagination.
    //
    // @param draftIdParam (required)
    // @param requestIdParam Request identifier to track subsequent API calls (optional)
    // @param rootPathParam Path of the root object of subtree (optional)
    // @return com.vmware.nsx_policy.model.PolicyDraftPaginatedAggregatedConfigurationResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(draftIdParam string, requestIdParam *string, rootPathParam *string) (model.PolicyDraftPaginatedAggregatedConfigurationResult, error)
}
