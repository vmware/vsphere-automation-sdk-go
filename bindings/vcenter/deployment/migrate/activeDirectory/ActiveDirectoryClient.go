/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: ActiveDirectory
 * Used by client-side stubs.
 */

package activeDirectory

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/vcenter/deployment"
)

// The ``ActiveDirectory`` interface provides methods to check if the migrated vCenter Server appliance can join to the given domain using the provided credentials. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ActiveDirectoryClient interface {


    // Checks whether the provided Active Directory user has permission to join the migrated vCenter Server appliance to the domain. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param specParam Information to connect to Active Directory.
    // @return Information about the success or failure of the checks that were performed.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws InvalidArgument if passed arguments are invalid.
    // @throws NotAllowedInCurrentState if the appliance is not in INITIALIZED state.
    Check(specParam CheckSpec) (deployment.CheckInfo, error) 

}
