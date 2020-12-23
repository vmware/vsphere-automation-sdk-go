/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Crls
 * Used by client-side stubs.
 */

package trust_management

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type CrlsClient interface {

    // Deletes an existing CRL.
    //
    // @param crlIdParam ID of CRL to delete (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(crlIdParam string) error

    // Returns information about the specified CRL. For additional information, include the ?details=true modifier at the end of the request URI.
    //
    // @param crlIdParam ID of CRL to read (required)
    // @param detailsParam whether to expand the pem data and show all its details (optional, default to false)
    // @return com.vmware.nsx_global_policy.model.Crl
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(crlIdParam string, detailsParam *bool) (model.Crl, error)

    // Adds a new certificate revocation list (CRL). The CRL is used to verify the client certificate status against the revocation lists published by the CA. For this reason, the administrator needs to add the CRL in certificate repository as well.
    //
    // @param crlObjectDataParam (required)
    // @return com.vmware.nsx_global_policy.model.CrlList
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Importcrl(crlObjectDataParam model.CrlObjectData) (model.CrlList, error)

    // Returns information about all CRLs. For additional information, include the ?details=true modifier at the end of the request URI.
    //
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param detailsParam whether to expand the pem data and show all its details (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @param type_Param Type of certificate to return (optional)
    // @return com.vmware.nsx_global_policy.model.CrlList
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(cursorParam *string, detailsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, type_Param *string) (model.CrlList, error)

    // Updates an existing CRL.
    //
    // @param crlIdParam ID of CRL to update (required)
    // @param crlParam (required)
    // @return com.vmware.nsx_global_policy.model.Crl
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(crlIdParam string, crlParam model.Crl) (model.Crl, error)
}
