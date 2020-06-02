/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: MacSetsNsxt
 * Used by client-side stubs.
 */

package mac_sets

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type MacSetsNsxtClient interface {

    // Read an MACSet
    //
    // @param enforcementPointNameParam Enforcement Point Name (required)
    // @param macSetNameParam MACSet name (required)
    // @return com.vmware.nsx_policy.model.GenericPolicyRealizedResource
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(enforcementPointNameParam string, macSetNameParam string) (model.GenericPolicyRealizedResource, error)

    // Paginated list of all Realized MACSets
    //
    // @param enforcementPointNameParam Enforcement Point Name (required)
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_policy.model.GenericPolicyRealizedResourceListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(enforcementPointNameParam string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.GenericPolicyRealizedResourceListResult, error)
}
