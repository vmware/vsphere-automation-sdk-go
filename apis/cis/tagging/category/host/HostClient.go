/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Host
 * Used by client-side stubs.
 */

package host

import (
)

// The ``Host`` interface provides methods to list categories that have hosts as an attachable type. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type HostClient interface {


    // Returns information about the categories that have hosts as an attachable type. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @return The list of categories that have hosts as an attachable type.
    List() ([]Summary, error) 

}
