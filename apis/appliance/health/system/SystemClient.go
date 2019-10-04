/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: System
 * Used by client-side stubs.
 */

package system

import (
    "time"
)

// ``System`` interface provides methods Get overall health of the system.
type SystemClient interface {


    // Get last check timestamp of the health of the system.
    // @return System health last check timestamp
    // @throws Error Generic error
    Lastcheck() (time.Time, error) 


    // Get overall health of system.
    // @return System health
    // @throws Error Generic error
    Get() (HealthLevel, error) 

}
