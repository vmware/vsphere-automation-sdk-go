/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Settings
 * Used by client-side stubs.
 */

package settings

import (
)

// The ``Settings`` interface provides methods to get or update settings related to the TPM 2.0 attestation protocol behavior.
type SettingsClient interface {


    // Return the TPM 2.0 protocol settings.
    // @return The settings.
    // @throws Error if there is a generic error.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``SECURITY_MGMT``.
    Get() (Info, error) 


    // Set the TPM 2.0 protocol settings.
    //
    // @param specParam The settings.
    // @throws Error if there is a generic error.
    // @throws InvalidArgument if the spec is invalid.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``SECURITY_MGMT``.
    Update(specParam UpdateSpec) error 

}
