/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Load
 * Used by client-side stubs.
 */

package load

import (
)

// ``Load`` interface provides methods Get load health.
type LoadClient interface {


    // Get load health.
    // @return Load health.
    // @throws Error Generic error
    Get() (HealthLevel, error) 

}
