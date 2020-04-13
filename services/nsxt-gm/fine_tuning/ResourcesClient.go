/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Resources
 * Used by client-side stubs.
 */

package fine_tuning

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type ResourcesClient interface {

    // This API provides field names of attributes in NSX types that are owned by Policy, as opposed to those owned by the enforcement point. For any type on NSX, some of the attributes of that type may be owned and set by Policy when realizing the intent, while some others may be owned and set by the enforcement point itself. This information can be used to disable updates to Policy owned attributes by the advanced networking UI, while allowing tweaking to the attributes owned by the management plane.
    //
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @param type_Param Type query (optional)
    // @return com.vmware.nsx_global_policy.model.ResourceInfoListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, type_Param *string) (model.ResourceInfoListResult, error)
}
