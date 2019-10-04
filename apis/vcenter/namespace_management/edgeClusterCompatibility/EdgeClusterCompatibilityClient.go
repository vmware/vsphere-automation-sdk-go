/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: EdgeClusterCompatibility
 * Used by client-side stubs.
 */

package edgeClusterCompatibility

import (
)

// The ``EdgeClusterCompatibility`` interface provides methods to get Namespaces compatibility information of NSX Edge Clusters. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type EdgeClusterCompatibilityClient interface {


    // Returns Namespaces compatibility information of NSX-T Edge Clusters matching the FilterSpec. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param clusterParam Identifier of a vCenter Cluster. Only Edge Clusters that are associated with the particular vCenter Cluster will be considered by the filter.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param distributedSwitchParam Identifier of a Distributed Switch. Only Edge Clusters that are associated with the particular Distributed Switch will be considered by the filter.
    // The parameter must be an identifier for the resource type: ``vSphereDistributedSwitch``.
    // @param filterParam Specification of matching Edge Clusters for which information should be returned.
    // If null, the behavior is equivalent to a FilterSpec with all properties null which means all Edge Clusters match the filter.
    // @return List of summaries of Edge Clusters associated with the given vCenter Cluster and Distributed Switch matching the FilterSpec.
    List(clusterParam string, distributedSwitchParam string, filterParam *FilterSpec) ([]Summary, error) 

}
