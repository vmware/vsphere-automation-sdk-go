/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: GetRoot
 * Used by client-side stubs.
 */

package getRoot

import (
)

// The ``GetRoot`` interface provides methods to get VMCA's trusted root cert. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type GetRootClient interface {


    // Returns the rhttpproxy TLS certificate. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @return Root certificate.
    // @throws Error if failed to fetch root certificate
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``CertificateAuthority.Manage`` and ``CertificateAuthority.Administer``.
    GetRoot() (string, error) 

}
