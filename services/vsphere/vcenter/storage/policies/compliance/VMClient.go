/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: VM
 * Used by client-side stubs.
 */

package compliance


// The ``VM`` class provides methods related to query virtual machines of given compliance statuses. This interface was added in vSphere API 6.7.
type VMClient interface {

    // Returns compliance information about at most 1000 virtual machines matching the filter VMFilterSpec. If there are no virtual machines matching the VMFilterSpec an empty List is returned. Virtual machines without storage policy association are not returned. This method was added in vSphere API 6.7.
    //
    // @param filterParam compliance status of matching virtual machines for which information should be returned.
    // @return compliance information about virtual machines matching the filter VMFilterSpec.
    // The key in the return value map will be an identifier for the resource type: ``VirtualMachine``.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if the VMFilterSpec#status property contains a value that is not supported by the server.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have the required privileges.
    // @throws Unsupported if the API is invoked against vCenter Server version is less than 6.5
    // @throws UnableToAllocateResource If more than 1000 results match the VMFilterSpec
	List(filterParam VMFilterSpec) (map[string]VMInfo, error)
}
