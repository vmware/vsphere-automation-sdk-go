/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: ServiceStatus
 * Used by client-side stubs.
 */

package kms


// The ``ServiceStatus`` interface provides methods to get the Key Provider Service health status. This interface was added in vSphere API 7.0.0.
type ServiceStatusClient interface {

    // Return the Key Provider Service health in the given cluster. This method was added in vSphere API 7.0.0.
    //
    // @param clusterParam Identifier of the cluster.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @return The Key Provider Service health status in the entire cluster.
    // @throws Error For any other error.
    // @throws InvalidArgument If the cluster id is empty.
    // @throws NotFound If the cluster is not found.
    // @throws Unauthenticated If the caller is not authenticated.
	Get(clusterParam string) (ServiceStatusInfo, error)
}
