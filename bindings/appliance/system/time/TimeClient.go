/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Time
 * Used by client-side stubs.
 */

package time

import (
)

// ``Time`` interface provides methods Gets system time.
type TimeClient interface {


    // Get system time.
    // @return System time
    // @throws Error Generic error
    Get() (SystemTimeStruct, error) 

}
