/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: LocalFilesystem
 * Used by client-side stubs.
 */

package guest


// The ``LocalFilesystem`` interface provides methods for retrieving information about the guest operating system local file systems. This interface was added in vSphere API 6.7.
type LocalFilesystemClient interface {

    // Returns details of the local file systems in the guest operating system. This method was added in vSphere API 6.7.
    //
    // @param vmParam Identifier of the virtual machine.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @return Information about the local file systems configured in the guest operating system.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine is not found.
    // @throws ServiceUnavailable if VMware Tools is not running.
    // @throws ServiceUnavailable if VMware Tools has not provided any data.
	Get(vmParam string) (map[string]LocalFilesystemInfo, error)
}
