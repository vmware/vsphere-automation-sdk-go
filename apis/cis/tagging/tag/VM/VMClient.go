/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: VM
 * Used by client-side stubs.
 */

package VM

import (
)

// The ``VM`` interface provides methods to list tags that can be attached to virtual machines. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type VMClient interface {


    // Returns information about the tags that can be attached to virtual machines. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @return The list of tags that can be attached to virtual machines available on this vCenter server.
    List() ([]Summary, error) 

}
