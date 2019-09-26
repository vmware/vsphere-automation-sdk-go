/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: LibraryItem
 * Used by client-side stubs.
 */

package libraryItem

import (
)

// The ``LibraryItem`` interface provides methods to identify virtual machines managed by Content Library.
type LibraryItemClient interface {


    // Returns the information about the library item associated with the virtual machine.
    //
    // @param vmParam Identifier of the virtual machine.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @return Information about the library item associated with the virtual machine.
    // @throws NotFound  if the virtual machine is not found.
    // @throws Unauthenticated  if the user that requested the method cannot be authenticated.
    // @throws Unauthorized  if the user that requested the method is not authorized to perform the method.
    Get(vmParam string) (Info, error) 

}
