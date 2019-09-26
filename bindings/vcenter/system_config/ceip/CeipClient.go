/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Ceip
 * Used by client-side stubs.
 */

package ceip

import (
)

// The ``Ceip`` interface provides methods to get and set the CEIP (Customer Experience Improvement Program) value. If CEIP is set to True, the customer agrees to participate in the CEIP program. Setting the CEIP value to True or False affects the whole MxN deployment.
type CeipClient interface {


    // Set CEIP (Customer Experience Improvement Program) value.
    //
    // @param valueParam CEIP boolean value to be set.
    // @throws Error if any error occurs during the execution of the set operation.
    Set(valueParam bool) error 


    // Get CEIP (Customer Experience Improvement Program) value.
    // @return CEIP bool value.
    // @throws Error if any error occurs during the execution of the get operation.
    Get() (bool, error) 

}
