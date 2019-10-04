/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Uptime
 * Used by client-side stubs.
 */

package uptime

import (
)

// ``Uptime`` interface provides methods Get the system uptime.
type UptimeClient interface {


    // Get the system uptime.
    // @return system uptime
    // @throws Error Generic error
    Get() (float64, error) 

}
