/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Remote
 * Used by client-side stubs.
 */

package remote

import (
)

// The ``Remote`` interface provides methods to retrieve remote key server certificate.
type RemoteClient interface {


    // Return the remote server certificates. 
    //
    //  Contacts the configured key servers and attempts to retrieve their certificates. These certificates might not yet be trusted. 
    //
    //  If the returned certificates are to be considered trustworthy, then it must be added to the list of trusted server certificates by adding to the certificates returned by ServerCertificates#get and invoking ServerCertificates#set with the updated array of certificates.
    //
    // @param providerParam Identifier of the provider.
    // The parameter must be an identifier for the resource type: ``com.vmware.esx.kms.providers``.
    // @param specParam Filter spec.
    // If null, the behavior is equivalent to a FilterSpec with all properties null
    // @return Summary of server certificates.
    // @throws InvalidArgument if the provider Id is empty.
    // @throws NotFound if the provider is not found.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
    // @throws Error For any other error.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``SECURITY_MGMT``.
    List(providerParam string, specParam *FilterSpec) ([]Summary, error) 

}
