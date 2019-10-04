/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Health
 * Used by client-side stubs.
 */

package health

import (
)

// The ``Health`` interface provides methods to retrieve health status for a container registry. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type HealthClient interface {


    // Returns the health information of a container registry in the vCenter. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param registryParam Identifier of the registry.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.content.Registry``.
    // @return Health information of the registry.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the registry does not exist.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user is not a member of the Administrators
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
    // * The resource ``com.vmware.vcenter.content.Registry`` referenced by the parameter ``registry`` requires ``System.Read``.
    Get(registryParam string) (Info, error) 

}
