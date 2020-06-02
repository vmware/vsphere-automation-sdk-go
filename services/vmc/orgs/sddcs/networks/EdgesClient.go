/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Edges
 * Used by client-side stubs.
 */

package networks

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/vmc/model"
)

type EdgesClient interface {

    // Retrieve information about all management and compute gateways and other routers (NSX Edges).
    //
    // @param orgParam Organization identifier. (required)
    // @param sddcParam Sddc Identifier. (required)
    // @param edgeTypeParam Retrieve records matching NSX Edge type ('gatewayServices' or 'distributedRouter'). Specify gatewayServices to find management and compute gateways in your SDDC. (required)
    // @param prevEdgeIdParam Provide Edge ID as prevEdgeId to serve as reference for startIndex. (optional)
    // @param startIndexParam Start index for the current page. Default is 0. (optional)
    // @param pageSizeParam Number of records per page. Default is 256. (optional)
    // @param sortOrderAscendingParam Set to true to fetch records in ascending sorted order. (optional)
    // @param sortByParam Sort records using one of the column names (id, name, description, tenantId, size, enableFips). (optional)
    // @param filterParam Filter records matching the NSX Edge ID, name or description. (optional)
    // @param ldRnameParam (optional)
    // @return com.vmware.vmc.model.PagedEdgeList
    // @throws InvalidRequest  Bad request. Request object passed is invalid.
    // @throws Unauthorized  Forbidden. Authorization header not provided
    // @throws NotFound  Not found. Requested object not found.
	Get(orgParam string, sddcParam string, edgeTypeParam string, prevEdgeIdParam *string, startIndexParam *int64, pageSizeParam *int64, sortOrderAscendingParam *bool, sortByParam *string, filterParam *string, ldRnameParam *string) (model.PagedEdgeList, error)
}
