/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: TrustAuthorityClusters
 * Used by client-side stubs.
 */

package trustAuthorityClusters

import (
)

// The ``TrustAuthorityClusters`` interface manages all the Trust Authority Components on each Trust Authority Host in the cluster. The ``TrustAuthorityClusters`` interface transforms a ClusterComputeResource into Trust Authority Cluster and vice versa.
type TrustAuthorityClustersClient interface {


    // Updates the state of a cluster.
    //
    // @param clusterParam Cluster id.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param specParam The specification for update of a cluster.
    // @throws Error if there is a generic error.
    // @throws NotFound if ``spec`` doesn't match to any cluster compute resource.
    // @throws Unauthenticated if the user can not be authenticated.
    Update(clusterParam string, specParam UpdateSpec) error 


    // Get the result of the last Update operation which matches the cluster id.
    //
    // @param clusterParam Cluster id.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @return The Info instance which contains information about the state of the cluster.
    // @throws Error if there is a generic error.
    // @throws NotFound if ``cluster`` doesn't match to any ClusterComputeResource.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``TrustedAdmin.ReadTrustedHosts``.
    // * The resource ``ClusterComputeResource`` referenced by the parameter ``cluster`` requires ``System.View``.
    Get(clusterParam string) (Info, error) 


    // Returns a list of clusters for this vCenter instance which matches the FilterSpec.
    //
    // @param specParam Return only clusters matching the specified filters.
    // If {\\\\@term.unset} return all clusters.
    // @return List of Summary for a TrustAuthorityClusters.
    // @throws Error if there is a generic error.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``TrustedAdmin.ReadTrustedHosts``.
    // * The resource ``ClusterComputeResource`` referenced by the property FilterSpec#cluster requires ``System.View``.
    List(specParam *FilterSpec) ([]Summary, error) 

}
