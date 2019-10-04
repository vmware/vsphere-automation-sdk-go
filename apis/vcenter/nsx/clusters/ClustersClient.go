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

// The ``Clusters`` interface provides methods to configure NSX networking on a vSphere cluster. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ClustersClient interface {


    // Enables NSX networking on the hosts in a vSphere cluster. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param clusterParam Identifier for the vSphere cluster on which NSX networking will be enabled.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param specParam Specification for setting up the NSX networking on the hosts in the vSphere cluster.
    // @throws NotFound if ``spec``.dvSwitch EnableSpec#dvSwitch could not be located.
    // @throws AlreadyExists if the cluster already has NSX networking enabled.
    // @throws InvalidArgument if ``spec`` contain any errors.
    // @throws Error if the system reports an error while responding to the request.
    Enable(clusterParam string, specParam EnableSpec) error 

}
