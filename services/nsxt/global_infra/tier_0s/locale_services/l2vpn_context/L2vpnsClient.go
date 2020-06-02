/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: L2vpns
 * Used by client-side stubs.
 */

package l2vpn_context

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type L2vpnsClient interface {

    //
    //
    // @param tier0IdParam (required)
    // @param localeServiceIdParam (required)
    // @param l2vpnIdParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(tier0IdParam string, localeServiceIdParam string, l2vpnIdParam string) error

    //
    //
    // @param tier0IdParam (required)
    // @param localeServiceIdParam (required)
    // @param l2vpnIdParam (required)
    // @return com.vmware.nsx_policy.model.L2Vpn
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(tier0IdParam string, localeServiceIdParam string, l2vpnIdParam string) (model.L2Vpn, error)

    //
    //
    // @param tier0IdParam (required)
    // @param localeServiceIdParam (required)
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_policy.model.L2VpnListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(tier0IdParam string, localeServiceIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.L2VpnListResult, error)

    //
    //
    // @param tier0IdParam (required)
    // @param localeServiceIdParam (required)
    // @param l2vpnIdParam (required)
    // @param l2VpnParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(tier0IdParam string, localeServiceIdParam string, l2vpnIdParam string, l2VpnParam model.L2Vpn) error

    //
    //
    // @param tier0IdParam (required)
    // @param localeServiceIdParam (required)
    // @param l2vpnIdParam (required)
    // @param l2VpnParam (required)
    // @return com.vmware.nsx_policy.model.L2Vpn
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(tier0IdParam string, localeServiceIdParam string, l2vpnIdParam string, l2VpnParam model.L2Vpn) (model.L2Vpn, error)
}
