/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Providers
 * Used by client-side stubs.
 */

package providers

import (
)

// The ``Providers`` interface provides methods to create, update and delete Key Providers that handoff to key servers.
type ProvidersClient interface {


    // Return a list of summary of Key Providers.
    //
    // @param clusterParam Identifier of the cluster.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @return List of providers.
    // @throws InvalidArgument If the cluster id is empty.
    // @throws NotFound If the cluster is not found.
    // @throws Unauthenticated If the caller is not authenticated.
    // @throws Error For any other error.
    List(clusterParam string) ([]Summary, error) 


    // Add a new Key Provider.
    //
    // @param clusterParam Identifier of the cluster.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param specParam Provider information.
    // @throws AlreadyExists If the provider already exists.
    // @throws InvalidArgument If the spec is invalid or cluster id is empty.
    // @throws NotFound If the cluster is not found.
    // @throws Unauthenticated If the caller is not authenticated.
    // @throws Error For any other error.
    Create(clusterParam string, specParam CreateSpec) error 


    // Update an existing Key Provider.
    //
    // @param clusterParam Identifier of the cluster.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param providerParam Identifier of the provider.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider``.
    // @param specParam Provider information.
    // @throws InvalidArgument If the cluster or provider id is empty, or the spec is invalid.
    // @throws NotFound If the cluster or provider is not found.
    // @throws Unauthenticated If the caller is not authenticated.
    // @throws Error For any other error.
    Update(clusterParam string, providerParam string, specParam UpdateSpec) error 


    // Remove a Key Provider.
    //
    // @param clusterParam Identifier of the cluster.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param providerParam Identifier of the provider.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider``.
    // @throws InvalidArgument If the cluster or provider id is empty.
    // @throws NotFound If the cluster or provider is not found.
    // @throws Unauthenticated If the caller is not authenticated.
    // @throws Error For any other error.
    Delete(clusterParam string, providerParam string) error 


    // Return information about a Key Provider.
    //
    // @param clusterParam Identifier of the cluster.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param providerParam Identifier of the provider.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider``.
    // @return Provider information.
    // @throws InvalidArgument If the cluster or provider id is empty.
    // @throws NotFound If the cluster or provider is not found.
    // @throws Unauthenticated If the caller is not authenticated.
    // @throws Error For any other error.
    Get(clusterParam string, providerParam string) (Info, error) 

}
