/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: ServiceStatus
 * Used by client-side stubs.
 */

package serviceStatus

import (
)

// The ``ServiceStatus`` interface provides methods to get the Key Provider Service health status.
type ServiceStatusClient interface {


    // Return the Key Provider Service health in the given cluster.
    //
    // @param clusterParam Identifier of the cluster.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @return The Key Provider Service health status in the entire cluster.
    // @throws Error For any other error.
    // @throws InvalidArgument If the cluster id is empty.
    // @throws NotFound If the cluster is not found.
    // @throws Unauthenticated If the caller is not authenticated.
    Get(clusterParam string) (Info, error) 

}
