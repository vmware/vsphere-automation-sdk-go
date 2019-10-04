/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Csr
 * Used by client-side stubs.
 */

package csr

import (
)

// The ``Csr`` interface provides methods to create a certificate signing request (CSR) for the attestation report signing certificate.
type CsrClient interface {


    // Generate a CSR. 
    //
    //  Generate a certificate signing request (CSR) for the attestation signing certificate. 
    //
    //  Calling this method repeatedly will result in a generating a new CSR each time and generating a new private key. An existing CSR, if it exists, will be discarded.
    //
    // @param specParam The CSR specification.
    // @return A structure containing the CSR.
    // @throws Error if there is a generic error.
    // @throws InvalidArgument if the CSR specification is invalid.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``SECURITY_MGMT``.
    Create(specParam CreateSpec) (Info, error) 


    // Delete an existing CSR. 
    //
    //  Discards the existing CSR. Completes successfully if no CSR is present.
    // @throws Error if there is a generic error.
    // @throws NotFound if no CSR is available.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``SECURITY_MGMT``.
    Delete() error 


    // Get the CSR. 
    //
    //  Return the most recently generated CSR.
    // @return A structure containing the CSR.
    // @throws Error if there is a generic error.
    // @throws NotFound if a CSR has not been created.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``SECURITY_MGMT``.
    Get() (Info, error) 

}
