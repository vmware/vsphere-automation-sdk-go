/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: ReplicationStatus
 * Used by client-side stubs.
 */

package topology


// The ``ReplicationStatus`` interface provides methods to retrieve replication status information of vCenter and Platform Services Controller nodes of type VCSA_EMBEDDED/PSC_EXTERNAL (see NodesInfo#type). This interface was added in vSphere API 6.7.2.
type ReplicationStatusClient interface {

    // Returns the replication information of vCenter and Platform Services Controller nodes of type VCSA_EMBEDDED/PSC_EXTERNAL (see NodesInfo#type) matching the ReplicationStatusFilterSpec. This method was added in vSphere API 6.7.2.
    //
    // @param filterParam  Specification of matching vCenter and Platform Services Controller nodes for which information should be returned.
    // If null, the behavior is equivalent to a ReplicationStatusFilterSpec with all properties null which means all vCenter and Platform Services Controller nodes of type VCSA_EMBEDDED/PSC_EXTERNAL match the filter.
    // @return Commonly used replication information about vCenter and Platform Services Controller nodes matching the ReplicationStatusFilterSpec.
    // @throws Unauthenticated  if the user can not be authenticated.
    // @throws Unauthorized  if the user doesn't have the required privileges.
    // @throws InvalidArgument  if the ReplicationStatusFilterSpec#nodes property contains a invalid value.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
	List(filterParam *ReplicationStatusFilterSpec) ([]ReplicationStatusSummary, error)
}
