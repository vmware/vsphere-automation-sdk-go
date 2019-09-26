/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: ManagementNetworkCompatibility
 * Used by client-side stubs.
 */

package managementNetworkCompatibility

import (
)

// The ``ManagementNetworkCompatibility`` interface provides methods to get compatibility information for Distributed Virtual Portgroups (DVPG) and their associated vmknics. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ManagementNetworkCompatibilityClient interface {


    // List the Distributed Portgroups (DVPG) and the network configuration of vmknics associated with the given ``cluster`` and ``distributed_switch``. This information will be used in populating the ManagementNetworkSpec for the EdgeClusters#enable API. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param clusterParam Identifier of a vCenter Cluster.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param distributedSwitchParam Identifier of a vSphere Distributed Switch (VDS).
    // The parameter must be an identifier for the resource type: ``DistributedVirtualSwitch``.
    // @return Compatibility information about the DVPG and its associated vmknics from the cluster identified by ``cluster`` and the VDS identified by ``distributed_switch``.
    // @throws NotFound if the cluster identified by ``cluster`` or if the VDS identified by ``distributed_switch`` could not be located.
    // @throws Error if any other error occurs while handling the request.
    List(clusterParam string, distributedSwitchParam string) ([]Summary, error) 

}
