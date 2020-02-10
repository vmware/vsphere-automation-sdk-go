/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: TlsCsr
 * Used by client-side stubs.
 */

package vcenter


// The ``TlsCsr`` interface provides methods to generate certificate signing request. This interface was added in vSphere API 6.7.2.
type TlsCsrClient interface {

    // Generates a CSR with the given Spec. This method was added in vSphere API 6.7.2.
    //
    // @param specParam The information needed to create a CSR.
    // @return A Certificate Signing Request.
    // @throws Error If CSR could not be created for given spec for a generic error.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``CertificateManagement.Manage`` and ``CertificateManagement.Administer``.
	Create(specParam TlsCsrSpec) (TlsCsrInfo, error)
}
