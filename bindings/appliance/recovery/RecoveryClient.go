/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Recovery
 * Used by client-side stubs.
 */

package recovery

import (
)

// The ``Recovery`` interface provides methods to invoke an appliance recovery (backup and restore).
type RecoveryClient interface {


    // Gets the properties of the appliance Recovery subsystem.
    // @return Structure containing the properties of the Recovery subsystem.
    // @throws Error if any error occurs during the execution of the operation.
    Get() (Info, error) 

}
