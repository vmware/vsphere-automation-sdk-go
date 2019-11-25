/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Observations
 * Used by client-side stubs.
 */

package traceflows

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type ObservationsClient interface {

    // Read traceflow observations for id traceflow-id
    //
    // @param traceflowIdParam (required)
    // @param componentNameParam Observations having the given component name will be listed. (optional)
    // @param componentTypeParam Observations having the given component type will be listed. (optional)
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param resourceTypeParam The type of observations that will be listed. (optional)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @param transportNodeNameParam Observations having the given transport node name will be listed. (optional)
    // @return com.vmware.nsx_policy.model.TraceflowObservationListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(traceflowIdParam string, componentNameParam *string, componentTypeParam *string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, resourceTypeParam *string, sortAscendingParam *bool, sortByParam *string, transportNodeNameParam *string) (model.TraceflowObservationListResult, error)
}
