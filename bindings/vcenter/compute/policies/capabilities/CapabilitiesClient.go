/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Capabilities
 * Used by client-side stubs.
 */

package capabilities

import (
)

// The ``Capabilities`` interface provides methods to manage compute policy capabilities. The description of the capability provides information about the intent of a policy based on this capability. A capability provides a type to create a policy (see Policies#create). A capability also provides a type that describes the information returned when retrieving information about a policy (see Policies#get). **Warning:** This interface is available as technical preview. It may be changed in a future release.
type CapabilitiesClient interface {


    // Returns information about the compute policy capabilities available in this vCenter server. **Warning:** This method is available as technical preview. It may be changed in a future release.
    // @return The list of compute policy capabilities available on this vCenter server.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
    List() ([]Summary, error) 


    // Returns information about a specific compute policy capability. **Warning:** This method is available as technical preview. It may be changed in a future release.
    //
    // @param capabilityParam Identifier of the capability for which information should be retrieved.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.compute.policies.Capability``.
    // @return Detailed information about the capability.
    // @throws NotFound if a capability with this identifier does not exist.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
    Get(capabilityParam string) (Info, error) 

}
