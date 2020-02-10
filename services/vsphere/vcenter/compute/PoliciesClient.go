/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Policies
 * Used by client-side stubs.
 */

package compute

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
)

// The ``Policies`` interface provides methods to manage compute policies in VMware Cloud on AWS. Usage beyond VMware Cloud on AWS is not supported. A compute policy defines the intended behavior for a collection of vSphere objects identified by a tag. A compute policy is an instance of a capability. It is created by providing a value for the creation type specified by the capability. See policies.CapabilitiesInfo#createSpecType. **Warning:** This interface is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type PoliciesClient interface {

    // Creates a new compute policy in VMware Cloud on AWS. Usage beyond VMware Cloud on AWS is not supported. **Warning:** This method is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    //
    // @param specParam Specification for the new policy to be created. The new policy will be an instance of the capability that has the creation type (see policies.CapabilitiesInfo#createSpecType) equal to the type of the specified value (see ``spec``).
    // The parameter must contain all the properties defined in policies.CreateSpec.
    // @return The identifier of the newly created policy. Use this identifier to get or destroy the policy.
    // The return value will be an identifier for the resource type: ``com.vmware.vcenter.compute.Policy``.
    // @throws AlreadyExists if a policy with the same name is already present on this vCenter server.
    // @throws InvalidArgument if a parameter passed in the spec is invalid.
    // @throws UnableToAllocateResource if more than 100 policies are created.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``ComputePolicy.Manage``.
	Create(specParam *data.StructValue) (string, error)

    // Returns information about the compute policies available in this vCenter server in VMware Cloud on AWS. Usage beyond VMware Cloud on AWS is not supported. **Warning:** This method is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    // @return The list of compute policies available on this vCenter server.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
	List() ([]PoliciesSummary, error)

    // Returns information about a specific compute policy in VMware Cloud on AWS. Usage beyond VMware Cloud on AWS is not supported. **Warning:** This method is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    //
    // @param policyParam Identifier of the policy for which information should be retrieved.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.compute.Policy``.
    // @return Detailed information about the specified compute policy. The returned value can be converted to the information type of the capability that this policy is based on. See policies.CapabilitiesInfo#infoType.
    // The return value will contain all the properties defined in policies.Info.
    // @throws NotFound if a policy with this identifier does not exist.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
	Get(policyParam string) (*data.StructValue, error)

    // Deletes a specific compute policy in VMware Cloud on AWS. Usage beyond VMware Cloud on AWS is not supported. **Warning:** This method is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    //
    // @param policyParam Identifier of the policy to be deleted.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.compute.Policy``.
    // @throws NotFound if a policy with this identifier does not exist.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``ComputePolicy.Manage``.
	Delete(policyParam string) error
}
