/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: CurrentPeerCertificates
 * Used by client-side stubs.
 */

package currentPeerCertificates

import (
)

// Retrieves the list of TLS certificates used by peer key servers. Those are meant for review. Following approval these certificates should be added as trusted certificates in the TrustedPeerCertificates interface
type CurrentPeerCertificatesClient interface {


    // Return the remote server certificates. 
    //
    //  Contacts the configured key servers and attempts to retrieve their certificates. These certificates might not yet be trusted. 
    //
    //  If the returned certificates are to be considered trustworthy, then it must be added to the list of trusted server certificates by adding to the certificates returned by TrustedPeerCertificates#get and invoking TrustedPeerCertificates#update with the updated array of certificates.
    //
    // @param clusterParam Identifier of the cluster.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param providerParam Identifier of the provider.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider``.
    // @param specParam Filter spec.
    // If null, the behavior is equivalent to a FilterSpec with all properties null
    // @return Summary of server certificates.
    // @throws InvalidArgument If the cluster or provider id is empty.
    // @throws NotFound If the cluster or provider is not found.
    // @throws Unauthenticated If the caller is not authenticated.
    // @throws Error For any other error.
    List(clusterParam string, providerParam string, specParam *FilterSpec) ([]Summary, error) 

}
