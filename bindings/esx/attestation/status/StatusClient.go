/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Status
 * Used by client-side stubs.
 */

package status

import (
)

// The ``Status`` interface provides methods to get the attestation service health status.
type StatusClient interface {


    // Return the servie status. 
    //
    //  Returns the attestation service health status.
    // @return The service status information.
    // @throws Error if there is a generic error.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``SECURITY_MGMT|READ_ONLY``.
    Get() (Info, error) 

}
