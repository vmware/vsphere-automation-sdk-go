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

// The ``VM`` interface provides methods to enumerate the policies that are associated with virtual machines and their status. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type VMClient interface {


    // Returns policies associated with virtual machines and their compliance status. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @return The policies associated with virtual machines with their compliance status.
    // @throws Unauthorized if the user does not have the required privileges.
    List() (ListResult, error) 

}
