/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Hidden
 * Used by client-side stubs.
 */

package hidden

import (
)

// The ``Hidden`` interface provides methods to retrieve the list of interfaces that are hidden and should not be exposed in various presentation layers of API infrastructure.
type HiddenClient interface {


    // Returns the interface identifiers that should be hidden.
    // @return The list of interface identifiers that should be hidden.
    // The return value will contain identifiers for the resource type: ``com.vmware.vapi.service``.
    List() ([]string, error) 

}
