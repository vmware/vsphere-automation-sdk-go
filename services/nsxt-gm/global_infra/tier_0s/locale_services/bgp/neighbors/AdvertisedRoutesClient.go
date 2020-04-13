/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: AdvertisedRoutes
 * Used by client-side stubs.
 */

package neighbors

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type AdvertisedRoutesClient interface {

    // Returns routes advertised by BGP neighbor from all edge transport nodes on which this neighbor is currently enabled. The query parameter \"source=cached\" is not supported.
    //
    // @param tier0IdParam (required)
    // @param localeServiceIdParam (required)
    // @param neighborIdParam (required)
    // @param countParam Number of routes to retrieve (optional, default to 1000)
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param enforcementPointPathParam Enforcement point path (optional)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_global_policy.model.BgpNeighborRoutesListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(tier0IdParam string, localeServiceIdParam string, neighborIdParam string, countParam *int64, cursorParam *string, enforcementPointPathParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.BgpNeighborRoutesListResult, error)
}
