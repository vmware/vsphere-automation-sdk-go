/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: WithCertificate
 * Used by client-side stubs.
 */

package principal_identities

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type WithCertificateClient interface {

    // Create a principal identity with a new, unused, certificate. The combination name and node_id needs to be unique across token-based and certificate-based principal identities.
    //
    // @param principalIdentityWithCertificateParam (required)
    // @return com.vmware.nsx_global_policy.model.PrincipalIdentity
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Create(principalIdentityWithCertificateParam model.PrincipalIdentityWithCertificate) (model.PrincipalIdentity, error)
}
