/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: FileIntegrity
 * Used by client-side stubs.
 */

package fileIntegrity

import (
)

// The ``FileIntegrity`` interface provides methods to perform file integrity operations in the appliance.
type FileIntegrityClient interface {


    // Create file integrity baseline for the appliance.
    // @throws Error Generic error
    Baseline() error 


    // Get state for scheduled file integrity check.
    // @return File integrity periodic check is enabled.
    // @throws Error Generic error
    Get() (bool, error) 


    // Set enabled state for scheduled file integrity check.
    //
    // @param enabledParam File integrity periodic check is enabled.
    // @throws Error Generic error
    Set(enabledParam bool) error 

}
