/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: TlsCsr
 * Used by client-side stubs.
 */

package tlsCsr

import (
)

// The ``TlsCsr`` interface provides methods to generate certificate signing request.
type TlsCsrClient interface {


    // Generates a CSR with the given Spec.
    //
    // @param specParam The information needed to create a CSR.
    // @return A Certificate Signing Request.
    // @throws Error If CSR could not be created for given spec for a generic error.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``CertificateManagement.Manage`` and ``CertificateManagement.Administer``.
    Create(specParam Spec) (Info, error) 

}
