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

// The ``Host`` interface provides methods to list tags that can be attached to hosts. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type HostClient interface {


    // Returns information about the tags that can be attached to virtual machines. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @return The list of tags that can be attached to hosts.
    List() ([]Summary, error) 

}
