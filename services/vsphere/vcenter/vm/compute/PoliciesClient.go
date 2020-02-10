/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Policies
 * Used by client-side stubs.
 */

package compute


// The ``Policies`` interface provides methods to query the status of policies on virtual machines in VMware Cloud on AWS. Usage beyond VMware Cloud on AWS is not supported. **Warning:** This interface is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type PoliciesClient interface {

    // Returns information about the compliance of a virtual machine with a compute policy in VMware Cloud on AWS. Usage beyond VMware Cloud on AWS is not supported. **Warning:** This method is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    //
    // @param vmParam Identifier of the virtual machine to query the status for.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param policyParam Identifier of the policy to query the status for.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.compute.Resources.COMPUTE_POLICY``.
    // @return Information about the compliance of the specified virtual machine with the specified compute policy.
    // @throws NotFound if a virtual machine with the given identifier does not exist, or if a policy with the given identifier does not exist.
    // @throws Unauthorized if the user doesn't have the required privileges.
	Get(vmParam string, policyParam string) (PoliciesInfo, error)
}
