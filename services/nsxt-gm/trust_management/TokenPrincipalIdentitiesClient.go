/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: TokenPrincipalIdentities
 * Used by client-side stubs.
 */

package trust_management

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type TokenPrincipalIdentitiesClient interface {

    // Register a principal identity that is going to be authenticated through a token. The combination name and node_id needs to be unique across token-based and certificate-based principal identities.
    //
    // @param tokenBasedPrincipalIdentityParam (required)
    // @return com.vmware.nsx_global_policy.model.TokenBasedPrincipalIdentity
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Create(tokenBasedPrincipalIdentityParam model.TokenBasedPrincipalIdentity) (model.TokenBasedPrincipalIdentity, error)

    // Delete a token-based principal identity.
    //
    // @param principalIdentityIdParam Unique id of the token-based principal identity to delete (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(principalIdentityIdParam string) error

    // Get a stored token-based principal identity
    //
    // @param principalIdentityIdParam ID of the principal identity to get (required)
    // @return com.vmware.nsx_global_policy.model.TokenBasedPrincipalIdentity
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(principalIdentityIdParam string) (model.TokenBasedPrincipalIdentity, error)

    // Return the list of token-based principal identities. | These don't have certificate or role information.
    // @return com.vmware.nsx_global_policy.model.TokenBasedPrincipalIdentityListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List() (model.TokenBasedPrincipalIdentityListResult, error)
}
