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

// The ``Clusters`` interface provides methods to enable and disable vSphere Namespaces on a vSphere cluster. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ClustersClient interface {


    // Enable vSphere Namespaces on the cluster. This operation sets up Kubernetes instance for the cluster along with worker nodes. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param clusterParam Identifier for the cluster on which vSphere Namespaces will be enabled.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param specParam Specification for setting up the Kubernetes API server and the worker nodes.
    // @throws AlreadyExists if the cluster already has vSphere Namespaces enabled.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if resources/objects could not be located.
    // @throws InvalidArgument if ``spec`` contain any errors.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have Namespaces.Manage privilege.
    Enable(clusterParam string, specParam EnableSpec) error 


    // Disable vSphere Namespaces on the cluster. This operation tears down the Kubernetes instance and the worker nodes associated with vSphere Namespaces enabled cluster. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param clusterParam Identifier for the cluster for which vSphere Namespaces will be disabled.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if cluster could not be located.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have Namespaces.Manage privilege.
    Disable(clusterParam string) error 


    // Returns information about a specific cluster. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param clusterParam Identifier for the cluster on which vSphere Namespaces are enabled.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @return Information about the desired state of the specified cluster.
    // @throws NotFound if cluster could not be located.
    // @throws Unsupported if the specified cluster does not have vSphere Namespaces enabled.
    // @throws Error if the system reports an error while responding to the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have System.Read privilege.
    Get(clusterParam string) (Info, error) 


    // Returns information about all clusters on which vSphere Namespaces are enabled on this vCenter. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @return List of summary of all clusters with vSphere Namespaces enabled.
    // @throws Error if the system reports an error while responding to the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have System.Read privilege.
    List() ([]Summary, error) 


    // Set a new configuration on the cluster object. The specified configuration is applied in entirety and will replace the current configuration fully. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param clusterParam Identifier for the cluster on which vSphere Namespaces is enabled.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param specParam New specification for the cluster.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if ``spec`` contain any errors.
    // @throws NotAllowedInCurrentState if vSphere Namespaces is being disabled on this cluster.
    // @throws NotFound if cluster could not be located.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have Namespaces.Manage privilege.
    Set(clusterParam string, specParam SetSpec) error 


    // Update configuration on the cluster object. The specified configuration is applied partially and null fields in ``spec`` will leave those parts of configuration as-is. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param clusterParam Identifier for the cluster on which vSphere Namespaces is enabled.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param specParam New specification for the cluster.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if ``spec`` contain any errors.
    // @throws NotAllowedInCurrentState if vSphere Namespaces is being disabled on this cluster.
    // @throws NotFound if cluster could not be located.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have Namespaces.Manage privilege.
    Update(clusterParam string, specParam UpdateSpec) error 

}
