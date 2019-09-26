/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Operations
 * Used by client-side stubs.
 */

package operations

import (
)

// Status of operations in the guest OS. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type OperationsClient interface {


    // Get information about the guest operation status. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param vmParam Identifier of the virtual machine.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @return guest operations readiness.
    // @throws NotFound if the virtual machine is not found.
    Get(vmParam string) (Info, error) 

}
