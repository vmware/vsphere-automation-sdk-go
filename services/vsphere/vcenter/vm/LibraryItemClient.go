/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: LibraryItem
 * Used by client-side stubs.
 */

package vm


// The ``LibraryItem`` interface provides methods to identify virtual machines managed by Content Library. This interface was added in vSphere API 6.9.1.
type LibraryItemClient interface {

    // Returns the information about the library item associated with the virtual machine. This method was added in vSphere API 6.9.1.
    //
    // @param vmParam Identifier of the virtual machine.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @return Information about the library item associated with the virtual machine.
    // @throws NotFound  if the virtual machine is not found.
    // @throws Unauthenticated  if the user that requested the method cannot be authenticated.
    // @throws Unauthorized  if the user that requested the method is not authorized to perform the method.
	Get(vmParam string) (LibraryItemInfo, error)
}
