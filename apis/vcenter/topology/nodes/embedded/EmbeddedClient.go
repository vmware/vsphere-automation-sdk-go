/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Embedded
 * Used by client-side stubs.
 */

package embedded

import (
)

// The ``Embedded`` interface provides methods to manage the vCenter Server node in the topology. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type EmbeddedClient interface {


    // Decommission the vCenter Server node. Decommissioning a vCenter Server node from the topology is irreversible. Any workloads associated with the node will be lost. Make sure to move the workloads before decommissioning the vCenter Server node. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param hostnameParam FQDN or IP address of vCenter Server node to be decommissioned.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.VCenter.name``.
    // @param specParam Information needed to decommission the vCenter Server node.
    // @param onlyPrecheckParam Flag indicating whether only a pre-check should be performed.
    // If null the full decommission will be performed.
    // @param repairReplicationParam Flag indicating whether broken replication should be repaired after decommission.
    // If null Broken replication will be repaired after decommission of vCenter Server node.
    // @throws Error if any other error occurs during the execution of the operation.
    // @throws Unsupported if the appliance is in not a vCenter Server node.
    // @throws InvalidArgument if passed arguments are invalid.
    // @throws UnverifiedPeer If the SSL certificate of the foreign vCenter Server node cannot be validated. 
    //  The value of the data property of errors.Error will be a class that contains all the properties defined in null.
    // @throws Unauthenticated if the caller is not authenticated.
    Decommission(hostnameParam string, specParam DecommissionSpec, onlyPrecheckParam *bool, repairReplicationParam *bool) error 

}
