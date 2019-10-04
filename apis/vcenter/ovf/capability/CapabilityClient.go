/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Capability
 * Used by client-side stubs.
 */

package capability

import (
)

// The ``Capability`` interface provides methods for retrieving information about the capabilities supported by the OVF service server.
type CapabilityClient interface {


    // Returns information about the capability of the given vCenter server.
    //
    // @param serverGuidParam  Target vCenter server GUID
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.VCenter``.
    // @return CapabilityInfo of the given server.
    // @throws InvalidArgument  If the given vCenter server GUID is invalid, or does not match available vCenter servers.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
    // * The resource ``com.vmware.vcenter.VCenter`` referenced by the parameter ``server_guid`` requires ``System.Read``.
    Get(serverGuidParam string) (CapabilityInfo, error) 

}
