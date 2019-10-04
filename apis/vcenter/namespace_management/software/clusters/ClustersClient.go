/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Clusters
 * Used by client-side stubs.
 */

package clusters

import (
)

// The ``Clusters`` interface provides methods to upgrade the vSphere clusters. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ClustersClient interface {


    // Upgrade the cluster to a specific version. This operation upgrades the components on master VM based on the slected version. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param clusterParam Identifier for the cluster which will be upgraded.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param specParam Specification for upgrading the cluster.
    // @throws Error if the system reports an error while responding to the request.
    // @throws Unsupported if pre-check failed of the cluster.
    // @throws NotFound if cluster could not be located.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have Namespaces.Manage privilege.
    Upgrade(clusterParam string, specParam UpgradeSpec) error 


    // Upgrade a set of clusters to its corresponding specific version. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param specsParam Specifications for upgrading selected clusters.
    // The key in the parameter map must be an identifier for the resource type: ``ClusterComputeResource``.
    // @return Pre-check result when invoking upgrade for each cluster.
    // The key in the return value map will be an identifier for the resource type: ``ClusterComputeResource``.
    // @throws Error if the system reports an error while responding to the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have Namespaces.Manage privilege on all specified clusters.
    UpgradeMultiple(specsParam map[string]UpgradeSpec) (map[string]Result, error) 


    // Returns upgrade related information of a specific cluster. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param clusterParam Identifier for the cluster which will be upgraded.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @return Information about the upgrade of the specified WCP enabled cluster.
    // @throws NotFound if cluster could not be located.
    // @throws Error if the system reports an error while responding to the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have System.Read privilege.
    Get(clusterParam string) (Info, error) 


    // Returns upgrade related information about all WCP enabled clusters. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @return List of upgrade summary of all WCP enabled clusters.
    // @throws Error if the system reports an error while responding to the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have System.Read privilege.
    List() ([]Summary, error) 

}
