/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Update
 * Used by client-side stubs.
 */

package update

import (
)

// ``Update`` interface provides methods Performs update repository configuration.
type UpdateClient interface {


    // Set update repository configuration.
    //
    // @param configParam update related configuration
    // @throws Error Generic error
    Set(configParam UpdateStructSet) error 


    // Get update repository configuration.
    // @return update related configuration
    // @throws Error Generic error
    Get() (UpdateStructGet, error) 

}
