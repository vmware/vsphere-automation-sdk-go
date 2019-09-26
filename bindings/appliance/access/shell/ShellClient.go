/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Shell
 * Used by client-side stubs.
 */

package shell

import (
)

// ``Shell`` interface provides methods Get/Set enabled state of BASH.
type ShellClient interface {


    // Set enabled state of BASH, that is, access to BASH from within the controlled CLI.
    //
    // @param configParam Shell configuration
    // @throws Error Generic error
    Set(configParam ShellConfig) error 


    // Get enabled state of BASH, that is, access to BASH from within the controlled CLI.
    // @return Current shell configuration.
    // @throws Error Generic error
    Get() (ShellConfig, error) 

}
