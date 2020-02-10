/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Cluster
 * Used by client-side stubs.
 */

package vcenter


// The ``Cluster`` interface provides methods to manage clusters in the vCenter Server.
type ClusterClient interface {

    // Returns information about at most 1000 visible (subject to permission checks) clusters in vCenter matching the ClusterFilterSpec.
    //
    // @param filterParam Specification of matching clusters for which information should be returned.
    // If null, the behavior is equivalent to a ClusterFilterSpec with all properties null which means all clusters match the filter.
    // @return Commonly used information about the clusters matching the ClusterFilterSpec.
    // @throws UnableToAllocateResource if more than 1000 clusters match the ClusterFilterSpec.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
	List(filterParam *ClusterFilterSpec) ([]ClusterSummary, error)

    // Retrieves information about the cluster corresponding to ``cluster``.
    //
    // @param clusterParam Identifier of the cluster.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @return The ClusterInfo instances that corresponds to the ``cluster``.
    // @throws NotFound if there is no cluster associated with ``cluster`` in the system.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the session id is missing from the request or the corresponding session object cannot be found.
    // @throws Unauthorized if the user doesn't not have the required privileges.
	Get(clusterParam string) (ClusterInfo, error)
}
