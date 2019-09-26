/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: EndorsementKeys
 * Used by client-side stubs.
 */

package endorsementKeys

import (
)

// The ``EndorsementKeys`` interface provides methods to manage Trusted Platform Module (TPM) Endorsement Keys (EK).
type EndorsementKeysClient interface {


    // Return a list of configured TPM endorsement keys.
    // @return A list of configured endorsement keys.
    // @throws Error if there is a generic error.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``SECURITY_MGMT``.
    List() ([]Summary, error) 


    // Add a new TPM endorsement key.
    //
    // @param specParam The configuration.
    // @throws AlreadyExists if the endorsement key name exists.
    // @throws Error if there is a generic error.
    // @throws InvalidArgument if the configuration is invalid.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``SECURITY_MGMT``.
    // * The resource ``com.vmware.esx.attestation.tpm2.endorsement_keys`` referenced by the property CreateSpec#name requires ``SECURITY_MGMT``.
    Create(specParam CreateSpec) error 


    // Remove a TPM endorsement key.
    //
    // @param nameParam The endorsement key name.
    // The parameter must be an identifier for the resource type: ``com.vmware.esx.attestation.tpm2.endorsement_keys``.
    // @throws Error if there is a generic error.
    // @throws InvalidArgument if the name is invalid.
    // @throws NotFound if the name is not found.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``SECURITY_MGMT``.
    // * The resource ``com.vmware.esx.attestation.tpm2.endorsement_keys`` referenced by the parameter ``name`` requires ``SECURITY_MGMT``.
    Delete(nameParam string) error 


    // Get the TPM endorsement key details.
    //
    // @param nameParam The endorsement key name.
    // The parameter must be an identifier for the resource type: ``com.vmware.esx.attestation.tpm2.endorsement_keys``.
    // @return The endorsement key info.
    // @throws Error if there is a generic error.
    // @throws InvalidArgument if the name is invalid.
    // @throws NotFound if the endorsement key is not found.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``SECURITY_MGMT``.
    // * The resource ``com.vmware.esx.attestation.tpm2.endorsement_keys`` referenced by the parameter ``name`` requires ``SECURITY_MGMT``.
    Get(nameParam string) (Info, error) 

}
