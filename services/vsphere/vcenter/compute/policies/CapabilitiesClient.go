/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Capabilities
 * Used by client-side stubs.
 */

package policies


// The ``Capabilities`` interface provides methods to manage compute policy capabilities in VMware Cloud on AWS. Usage beyond VMware Cloud on AWS is not supported. The description of the capability provides information about the intent of a policy based on this capability. A capability provides a type to create a policy (see Policies#create). A capability also provides a type that describes the information returned when retrieving information about a policy (see Policies#get). **Warning:** This interface is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type CapabilitiesClient interface {

    // Returns information about the compute policy capabilities available in this vCenter server in VMware Cloud on AWS. Usage beyond VMware Cloud on AWS is not supported. **Warning:** This method is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    // @return The list of compute policy capabilities available on this vCenter server.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
	List() ([]CapabilitiesSummary, error)

    // Returns information about a specific compute policy capability in VMware Cloud on AWS. Usage beyond VMware Cloud on AWS is not supported. **Warning:** This method is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    //
    // @param capabilityParam Identifier of the capability for which information should be retrieved.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.compute.policies.Capability``.
    // @return Detailed information about the capability.
    // @throws NotFound if a capability with this identifier does not exist.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
	Get(capabilityParam string) (CapabilitiesInfo, error)
}
