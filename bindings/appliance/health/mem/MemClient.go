/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Mem
 * Used by client-side stubs.
 */

package mem

import (
)

// ``Mem`` interface provides methods Get memory health.
type MemClient interface {


    // Get memory health.
    // @return Memory health.
    // @throws Error Generic error
    Get() (HealthLevel, error) 

}
