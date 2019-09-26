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
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
)

// The ``Policies`` interface provides methods to manage compute policies. A compute policy defines the intended behavior for a collection of vSphere objects identified by a tag. A compute policy is an instance of a capability. It is created by providing a value for the creation type specified by the capability. See capabilities.Info#createSpecType. **Warning:** This interface is available as technical preview. It may be changed in a future release.
type PoliciesClient interface {


    // Creates a new compute policy. **Warning:** This method is available as technical preview. It may be changed in a future release.
    //
    // @param specParam Specification for the new policy to be created. The new policy will be an instance of the capability that has the creation type (see capabilities.Info#createSpecType) equal to the type of the specified value (see ``spec``).
    // The parameter must contain all the properties defined in CreateSpec.
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


    // Returns information about the compute policies available in this vCenter server. **Warning:** This method is available as technical preview. It may be changed in a future release.
    // @return The list of compute policies available on this vCenter server.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
    List() ([]Summary, error) 


    // Returns information about a specific compute policy. **Warning:** This method is available as technical preview. It may be changed in a future release.
    //
    // @param policyParam Identifier of the policy for which information should be retrieved.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.compute.Policy``.
    // @return Detailed information about the specified compute policy. The returned value can be converted to the information type of the capability that this policy is based on. See capabilities.Info#infoType.
    // The return value will contain all the properties defined in Info.
    // @throws NotFound if a policy with this identifier does not exist.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
    Get(policyParam string) (*data.StructValue, error) 


    // Deletes a specific compute policy. **Warning:** This method is available as technical preview. It may be changed in a future release.
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
