/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: NSXAppliance
 * Used by client-side stubs.
 */

package NSXAppliance

import (
)

// The ``NSXAppliance`` interface provides methods to create, query and delete a nsx appliance attached to the vCenter. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type NSXApplianceClient interface {


    // Gets the NSX appliance deployment status on the current vCenter. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @return Returns a Info
    // @throws Error for any other unspecified error.
    // @throws NotFound if the resource cannot be queried.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user is not a member of the Administrators group.
    Get() (Info, error) 


    // Starts the NSX appliance deployment on the current vCenter. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param specParam Specification for deploying the NSX appliance
    // @throws Error for any other unspecified error.
    // @throws InvalidArgument if the input spec is not valid.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user is not a member of the Administrators group.
    // @throws NotAllowedInCurrentState if the operation is already in progress.
    // @throws AlreadyExists if the deployment already exists.
    Create(specParam InstallSpec) error 


    // Deletes the NSX appliance deployed on the current vCenter. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @throws Error for any other unspecified error.
    // @throws NotFound if the resource is not found.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user is not a member of the Administrators group.
    // @throws NotAllowedInCurrentState if the operation is already in progress.
    Delete() error 

}
