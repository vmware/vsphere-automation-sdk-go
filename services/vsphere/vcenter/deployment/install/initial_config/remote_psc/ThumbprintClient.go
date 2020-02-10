/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Thumbprint
 * Used by client-side stubs.
 */

package remote_psc


// The ``Thumbprint`` interface provides methods to get the thumbprint of the remote PSC. This interface was added in vSphere API 6.7.
type ThumbprintClient interface {

    // Gets the SHA1 thumbprint of the remote PSC. This method was added in vSphere API 6.7.
    //
    // @param specParam Information used to connect to the remote PSC.
    // @return The thumbprint of the specified remote PSC
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Error on exception.
	Get(specParam ThumbprintRemoteSpec) (string, error)
}
