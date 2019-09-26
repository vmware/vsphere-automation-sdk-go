/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Cluster
 * Used by client-side stubs.
 */

package cluster

import (
)

// The ``Cluster`` interface provides methods to manage clusters in the vCenter Server.
type ClusterClient interface {


    // Returns information about at most 1000 visible (subject to permission checks) clusters in vCenter matching the FilterSpec.
    //
    // @param filterParam Specification of matching clusters for which information should be returned.
    // If null, the behavior is equivalent to a FilterSpec with all properties null which means all clusters match the filter.
    // @return Commonly used information about the clusters matching the FilterSpec.
    // @throws UnableToAllocateResource if more than 1000 clusters match the FilterSpec.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    List(filterParam *FilterSpec) ([]Summary, error) 


    // Retrieves information about the cluster corresponding to ``cluster``.
    //
    // @param clusterParam Identifier of the cluster.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @return The Info instances that corresponds to the ``cluster``.
    // @throws NotFound if there is no cluster associated with ``cluster`` in the system.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the session id is missing from the request or the corresponding session object cannot be found.
    // @throws Unauthorized if the user doesn't not have the required privileges.
    Get(clusterParam string) (Info, error) 


    // Create a new cluster in the vCenter inventory. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param specParam Specification for the new cluster to be created.
    // @return The identifier of the newly created cluster
    // The return value will be an identifier for the resource type: ``ClusterComputeResource``.
    // @throws AlreadyExists if the cluster with the same name is already present.
    // @throws InvalidArgument if the cluster name is empty or invalid as per the underlying implementation.
    // @throws InvalidElementType if the parent folder does not support vSphere compute resource as its children type.
    // @throws Unauthenticated if the session id is missing from the request or the corresponding session object cannot be found.
    // @throws Unauthorized if the user doesn't not have the required privileges.
    Create(specParam CreateSpec) (string, error) 


    // Delete an empty cluster from the vCenter Server. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param clusterParam Identifier of the cluster to be deleted.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @throws Error For all vimfaults thrown by VMOMI but yet to be figured out. TODO Replace this with specific errors.
    // @throws NotFound if there is no cluster associated with ``cluster`` in the system.
    // @throws ResourceInUse if the cluster associated with ``cluster`` is not empty.
    // @throws Unauthenticated if the session id is missing from the request or the corresponding session object cannot be found.
    // @throws Unauthorized if the user doesn't not have the required privileges.
    Delete(clusterParam string) error 

}
