/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: ImportFlag
 * Used by client-side stubs.
 */

package importFlag

import (
)

// The ``ImportFlag`` interface provides methods for retrieving information about the import flags supported by the deployment platform. Import flags can be specified in a libraryItem.ResourcePoolDeploymentSpec to customize an OVF deployment.
type ImportFlagClient interface {


    // Returns information about the import flags supported by the deployment platform. 
    //
    //  The supported flags are: 
    //
    // LAX
    //     Lax mode parsing of the OVF descriptor.
    //  
    //
    //  Future server versions might support additional flags.
    //
    // @param rpParam  The identifier of resource pool target for retrieving the import flag(s).
    // The parameter must be an identifier for the resource type: ``ResourcePool``.
    // @return A array of supported import flags.
    // @throws NotFound  If the resource pool associated with ``rp`` does not exist.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
    // * The resource ``ResourcePool`` referenced by the parameter ``rp`` requires ``System.Read``.
    List(rpParam string) ([]Info, error) 

}
