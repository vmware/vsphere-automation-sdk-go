/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Policies
 * Used by client-side stubs.
 */

package policies

import (
)

// The ``Policies`` interface provides methods to query the status of policies on virtual machines. **Warning:** This interface is available as technical preview. It may be changed in a future release.
type PoliciesClient interface {


    // Returns information about the compliance of a virtual machine with a compute policy. **Warning:** This method is available as technical preview. It may be changed in a future release.
    //
    // @param vmParam Identifier of the virtual machine to query the status for.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param policyParam Identifier of the policy to query the status for.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.compute.Resources.COMPUTE_POLICY``.
    // @return Information about the compliance of the specified virtual machine with the specified compute policy.
    // @throws NotFound if a virtual machine with the given identifier does not exist, or if a policy with the given identifier does not exist.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Get(vmParam string, policyParam string) (Info, error) 

}
