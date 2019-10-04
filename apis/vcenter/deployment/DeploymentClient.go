/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Deployment
 * Used by client-side stubs.
 */

package deployment

import (
)

// The ``Deployment`` interface provides methods to get the status of the vCenter appliance deployment.
type DeploymentClient interface {


    // Get the current status of the appliance deployment.
    // @return Info structure containing the status information about the appliance.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws NotFound if appliance state cannot be determined.
    Get() (Info, error) 


    // Rollback a failed appliance so it can be configured once again.
    // @throws Unsupported if the appliance is not in FAILED state.
    // @throws Unauthenticated if the caller is not authenticated.
    Rollback() error 

}
