/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: DeploymentType
 * Used by client-side stubs.
 */

package deploymentType

import (
)

// The ``DeploymentType`` interface provides methods to get/set the type of the appliance.
type DeploymentTypeClient interface {


    // Get the type of the vCenter appliance.
    // @return The type of the vCenter appliance.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws NotAllowedInCurrentState if appliance is not in CONFIGURED state.
    Get() (Info, error) 


    // Reconfigure the type of the vCenter appliance.
    //
    // @param specParam ReconfigureSpec to set the appliance type.
    // @throws Unsupported if the appliance is in CONFIGURED state and if not changing the type from VCSA_EMBEDDED to VCSA_EXTERNAL.
    // @throws InvalidArgument if passed arguments are invalid.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthenticated if external PSC credentials are not valid.
    // @throws NotAllowedInCurrentState if the appliance is not in INITIALIZED or CONFIGURED state.
    Reconfigure(specParam ReconfigureSpec) error 

}
