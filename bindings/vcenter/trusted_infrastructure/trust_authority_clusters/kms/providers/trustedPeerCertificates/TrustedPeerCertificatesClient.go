/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: TrustedPeerCertificates
 * Used by client-side stubs.
 */

package trustedPeerCertificates

import (
)

// Provides management operations for the TLS certificates trusted for communication with peer key servers. 
//
//  To obtain the currently used TLS certificates use the CurrentPeerCertificates interface
type TrustedPeerCertificatesClient interface {


    // Update trusted server certificate(s). 
    //
    //  The client will use these certificates to validate the server connection. The existing list of trusted certificates will be overwritten. 
    //
    //  The client will not trust the server connection until a server certificate has been set.
    //
    // @param clusterParam Identifier of the cluster.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param providerParam Identifier of the provider.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider``.
    // @param specParam The update spec
    // @throws InvalidArgument If one or more certificates are invalid or the cluster/provider Id is empty.
    // @throws NotFound If the cluster or provider is not found.
    // @throws Unauthenticated If the caller is not authenticated.
    // @throws Error If any other error occurs.
    Update(clusterParam string, providerParam string, specParam UpdateSpec) error 


    // Return the list of trusted server certificates.
    //
    // @param clusterParam Identifier of the cluster.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param providerParam Identifier of the provider.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider``.
    // @return Info containing server certificates, PEM.
    // @throws InvalidArgument If cluster or provider id are empty.
    // @throws NotFound If the cluster or provider is not found.
    // @throws Unauthenticated If the caller is not authenticated.
    // @throws Error If any other error occurs.
    Get(clusterParam string, providerParam string) (Info, error) 

}
