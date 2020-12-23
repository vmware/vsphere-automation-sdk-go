/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: OidcUris
 * Used by client-side stubs.
 */

package trust_management

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type OidcUrisClient interface {

    // This request also fetches the issuer and jwks_uri meta-data from the OIDC end-point and stores it.
    //
    // @param oidcEndPointParam (required)
    // @return com.vmware.nsx_global_policy.model.OidcEndPoint
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Create(oidcEndPointParam model.OidcEndPoint) (model.OidcEndPoint, error)

    // When ?refresh=true is added to the request, the meta-data is newly fetched from the OIDC end-point.
    //
    // @param idParam (required)
    // @param refreshParam Refresh meta-data (optional, default to false)
    // @return com.vmware.nsx_global_policy.model.OidcEndPoint
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(idParam string, refreshParam *bool) (model.OidcEndPoint, error)

    // Return the list of OpenID Connect end-points.
    // @return com.vmware.nsx_global_policy.model.OidcEndPointListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List() (model.OidcEndPointListResult, error)

    // Update a OpenID Connect end-point's thumbprint used to connect to the oidc_uri through SSL
    //
    // @param updateOidcEndPointThumbprintRequestParam (required)
    // @return com.vmware.nsx_global_policy.model.OidcEndPoint
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Updatethumbprint(updateOidcEndPointThumbprintRequestParam model.UpdateOidcEndPointThumbprintRequest) (model.OidcEndPoint, error)
}
