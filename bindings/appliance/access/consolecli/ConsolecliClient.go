/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Consolecli
 * Used by client-side stubs.
 */

package consolecli

import (
)

// ``Consolecli`` interface provides methods Get/Set enabled state of CLI.
type ConsolecliClient interface {


    // Set enabled state of the console-based controlled CLI (TTY1).
    //
    // @param enabledParam Console-based controlled CLI is enabled.
    // @throws Error Generic error
    Set(enabledParam bool) error 


    // Get enabled state of the console-based controlled CLI (TTY1).
    // @return Console-based controlled CLI is enabled.
    // @throws Error Generic error
    Get() (bool, error) 

}
