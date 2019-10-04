/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: DistributedSwitchCompatibility
 * Used by client-side stubs.
 */

package distributedSwitchCompatibility

import (
)

// The ``DistributedSwitchCompatibility`` interface provides methods to get Namespaces compatibility information of Distributed Switches in this vCenter. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type DistributedSwitchCompatibilityClient interface {


    // Returns Namespaces compatibility information of Distributed Switches in vCenter associated with the vCenter cluster, matching the FilterSpec. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param clusterParam Identifier of a vCenter Cluster. Only Distributed Switches associated with the vCenter Cluster will be considered by the filter.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param filterParam Specification of matching Distributed Switches for which information should be returned.
    // If null, the behavior is equivalent to a FilterSpec with all properties null which means all Distributed Switches match the filter.
    // @return Namespaces compatibility information for Distributed Switches matching the the FilterSpec.
    List(clusterParam string, filterParam *FilterSpec) ([]Summary, error) 

}
