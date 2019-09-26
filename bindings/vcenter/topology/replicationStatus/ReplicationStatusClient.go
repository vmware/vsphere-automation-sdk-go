/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: ReplicationStatus
 * Used by client-side stubs.
 */

package replicationStatus

import (
)

// The ``ReplicationStatus`` interface provides methods to retrieve replication status information of vCenter and Platform Services Controller nodes of type VCSA_EMBEDDED/PSC_EXTERNAL (see nodes.Info#type).
type ReplicationStatusClient interface {


    // Returns the replication information of vCenter and Platform Services Controller nodes of type VCSA_EMBEDDED/PSC_EXTERNAL (see nodes.Info#type) matching the FilterSpec.
    //
    // @param filterParam  Specification of matching vCenter and Platform Services Controller nodes for which information should be returned.
    // If null, the behavior is equivalent to a FilterSpec with all properties null which means all vCenter and Platform Services Controller nodes of type VCSA_EMBEDDED/PSC_EXTERNAL match the filter.
    // @return Commonly used replication information about vCenter and Platform Services Controller nodes matching the FilterSpec.
    // @throws Unauthenticated  if the user can not be authenticated.
    // @throws Unauthorized  if the user doesn't have the required privileges.
    // @throws InvalidArgument  if the FilterSpec#nodes property contains a invalid value.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
    List(filterParam *FilterSpec) ([]Summary, error) 

}
