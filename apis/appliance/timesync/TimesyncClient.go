/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Timesync
 * Used by client-side stubs.
 */

package timesync

import (
)

// ``Timesync`` interface provides methods Performs time synchronization configuration.
type TimesyncClient interface {


    // Set time synchronization mode.
    //
    // @param modeParam Time synchronization mode.
    // @throws Error Generic error
    Set(modeParam TimeSyncMode) error 


    // Get time synchronization mode.
    // @return Time synchronization mode.
    // @throws Error Generic error
    Get() (TimeSyncMode, error) 

}
