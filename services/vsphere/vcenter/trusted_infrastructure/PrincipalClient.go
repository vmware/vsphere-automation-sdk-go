/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Principal
 * Used by client-side stubs.
 */

package trusted_infrastructure


// The ``Principal`` interface contains information about the certificates which sign the tokens used by vCenter for authentication. This interface was added in vSphere API 7.0.0.
type PrincipalClient interface {

    // Returns information about the STS used by this vCenter instance. This method was added in vSphere API 7.0.0.
    // @return \\\\@{link Info} a summary containing the certificates used to sign tokens and the solution user used to retrieve them.
    // @throws Error if there is a generic error.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``TrustedAdmin.ReadStsInfo``.
	Get() (PrincipalInfo, error)
}
