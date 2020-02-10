/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Network
 * Used by client-side stubs.
 */

package vcenter


// The Network interface provides methods for manipulating a vCenter Server network.
type NetworkClient interface {

    // Returns information about at most 1000 visible (subject to permission checks) networks in vCenter matching the NetworkFilterSpec.
    //
    // @param filterParam Specification of matching networks for which information should be returned.
    // If null, the behavior is equivalent to a NetworkFilterSpec with all properties null which means all networks match the filter.
    // @return Commonly used information about the networks matching the NetworkFilterSpec.
    // @throws InvalidArgument if the NetworkFilterSpec#types property contains a value that is not supported by the server.
    // @throws UnableToAllocateResource if more than 1000 networks match the NetworkFilterSpec.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
	List(filterParam *NetworkFilterSpec) ([]NetworkSummary, error)
}
