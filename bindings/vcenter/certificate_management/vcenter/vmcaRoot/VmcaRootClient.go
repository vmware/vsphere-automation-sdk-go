/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: VmcaRoot
 * Used by client-side stubs.
 */

package vmcaRoot

import (
)

// The ``VmcaRoot`` interface provides methods to replace VMware Certificate Authority (VMCA) root certificate.
type VmcaRootClient interface {


    // Replace Root Certificate with VMCA signed one using the given Spec. 
    //
    // After this method completes, the services using the certificate will be restarted for the new certificate to take effect.
    //
    // @param specParam The information needed to generate VMCA signed Root Certificate.
    // Default values will be set for all null parameters.
    // @throws Error If the system failed to renew the TLS certificate.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``CertificateManagement.Administer``.
    Create(specParam *CreateSpec) error 

}
