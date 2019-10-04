/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Swap
 * Used by client-side stubs.
 */

package swap

import (
)

// ``Swap`` interface provides methods Get swap health.
type SwapClient interface {


    // Get swap health.
    // @return Swap health
    // @throws Error Generic error
    Get() (HealthLevel, error) 

}
