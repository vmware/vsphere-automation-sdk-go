/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Content
 * Used by client-side stubs.
 */

package eula

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type ContentClient interface {

    // Return the content of end user license agreement in the specified format. By default, it's pure string without line break
    //
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @param valueFormatParam End User License Agreement content output format (optional)
    // @return com.vmware.nsx_policy.model.EULAContent
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, valueFormatParam *string) (model.EULAContent, error)
}
