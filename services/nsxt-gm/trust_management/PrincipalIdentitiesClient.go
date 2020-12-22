/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: PrincipalIdentities
 * Used by client-side stubs.
 */

package trust_management

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type PrincipalIdentitiesClient interface {

    // Associates a principal's name with a certificate that is used to authenticate. The combination name and node_id needs to be unique across token-based and certificate-based principal identities. Deprecated, use POST /trust-management/principal-identities/with-certificate instead.
    //
    // @param principalIdentityParam (required)
    // @return com.vmware.nsx_global_policy.model.PrincipalIdentity
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Create(principalIdentityParam model.PrincipalIdentity) (model.PrincipalIdentity, error)

    // Delete a principal identity. It does not delete the certificate.
    //
    // @param principalIdentityIdParam Unique id of the principal identity to delete (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(principalIdentityIdParam string) error

    // Get a stored principal identity
    //
    // @param principalIdentityIdParam ID of the principal identity to get (required)
    // @return com.vmware.nsx_global_policy.model.PrincipalIdentity
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(principalIdentityIdParam string) (model.PrincipalIdentity, error)

    // Returns the list of principals registered with a certificate.
    // @return com.vmware.nsx_global_policy.model.PrincipalIdentityList
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List() (model.PrincipalIdentityList, error)

    // Update a principal identity's certificate
    //
    // @param updatePrincipalIdentityCertificateRequestParam (required)
    // @return com.vmware.nsx_global_policy.model.PrincipalIdentity
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Updatecertificate(updatePrincipalIdentityCertificateRequestParam model.UpdatePrincipalIdentityCertificateRequest) (model.PrincipalIdentity, error)
}
