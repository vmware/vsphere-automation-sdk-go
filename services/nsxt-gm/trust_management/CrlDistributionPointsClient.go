/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: CrlDistributionPoints
 * Used by client-side stubs.
 */

package trust_management

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type CrlDistributionPointsClient interface {

    // Create an entity that will represent a Crl Distribution Point
    //
    // @param crlDistributionPointParam (required)
    // @return com.vmware.nsx_global_policy.model.CrlDistributionPoint
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Create(crlDistributionPointParam model.CrlDistributionPoint) (model.CrlDistributionPoint, error)

    // Delete a CrlDistributionPoint. It does not delete the actual CRL.
    //
    // @param crlDistributionPointIdParam Unique id of the CrlDistributionPoint to delete (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(crlDistributionPointIdParam string) error

    //
    //
    // @param crlDistributionPointIdParam (required)
    // @return com.vmware.nsx_global_policy.model.CrlDistributionPoint
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(crlDistributionPointIdParam string) (model.CrlDistributionPoint, error)

    // Return the list of CrlDistributionPoints
    //
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_global_policy.model.CrlDistributionPointList
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.CrlDistributionPointList, error)

    //
    //
    // @param crlDistributionPointIdParam (required)
    // @param crlDistributionPointParam (required)
    // @return com.vmware.nsx_global_policy.model.CrlDistributionPoint
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(crlDistributionPointIdParam string, crlDistributionPointParam model.CrlDistributionPoint) (model.CrlDistributionPoint, error)
}
