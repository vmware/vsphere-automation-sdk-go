/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Leases
 * Used by client-side stubs.
 */

package dhcp_server_configs

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type LeasesClient interface {

    // Read DHCP server leases
    //
    // @param configIdParam (required)
    // @param connectivityPathParam String Path of Tier0, Tier1 or Segment (required)
    // @param addressParam IP or MAC address (optional)
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param enforcementPointPathParam String Path of the enforcement point (optional)
    // @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param segmentPathParam Segment path to retrieve lease information (optional)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @param sourceParam Data source type. (optional)
    // @return com.vmware.nsx_policy.model.DhcpLeasesResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(configIdParam string, connectivityPathParam string, addressParam *string, cursorParam *string, enforcementPointPathParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, segmentPathParam *string, sortAscendingParam *bool, sortByParam *string, sourceParam *string) (model.DhcpLeasesResult, error)
}
