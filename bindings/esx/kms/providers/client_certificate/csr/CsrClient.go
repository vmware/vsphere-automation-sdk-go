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

// The ``Csr`` interface provides methods to create a certificate signing request(CSR).
type CsrClient interface {


    // Generate a certificate signing request (CSR) for the client certificate. This overwrites any existing CSR. 
    //
    //  The CSR will allow the certificate to be signed by a third party. Once the CSR is signed, certificate may be added as a client certificate. 
    //
    //  Calling the API repeatedly will result in a generating a new CSR each time.
    //
    // @param providerParam Identifier of the provider.
    // The parameter must be an identifier for the resource type: ``com.vmware.esx.kms.providers``.
    // @return A structure containing the CSR.
    // @throws InvalidArgument if the provider Id is empty.
    // @throws NotFound if the provider is not found.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
    // @throws Error if any other error occurs.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``SECURITY_MGMT``.
    Create(providerParam string) (Info, error) 


    // Get existing certificate signing request (CSR) for the client certificate and optionally the private key. 
    //
    //  The CSR will allow the certificate to be signed by a third party. Once the CSR is signed, certificate may be added as a client certificate. 
    //
    // @param providerParam Identifier of the provider.
    // The parameter must be an identifier for the resource type: ``com.vmware.esx.kms.providers``.
    // @param specParam Specification to include additional information.
    // If null, the behavior is equivalent to a GetSpec with all properties null
    // @return A structure containing the CSR.
    // @throws InvalidArgument if the provider Id is empty.
    // @throws NotFound if the provider not found.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
    // @throws Error if any other error occurs.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``SECURITY_MGMT``.
    Get(providerParam string, specParam *GetSpec) (Info, error) 


    // Set the certificate signing request (CSR) and private key. 
    //
    //  The CSR will allow the certificate to be signed by a third party. Once the CSR is signed, certificate may be added as a client certificate. 
    //
    // @param providerParam Identifier of the provider.
    // The parameter must be an identifier for the resource type: ``com.vmware.esx.kms.providers``.
    // @param specParam Specification describing new CSR and private key.
    // @throws InvalidArgument if the provider Id is empty.
    // @throws NotFound if the provider not found.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
    // @throws Error if any other error occurs.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``SECURITY_MGMT``.
    Set(providerParam string, specParam SetSpec) error 

}
