/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Folder
 * Used by client-side stubs.
 */

package vcenter


// The Folder interface provides methods for manipulating a vCenter Server folder.
type FolderClient interface {

    // Returns information about at most 1000 visible (subject to permission checks) folders in vCenter matching the FolderFilterSpec.
    //
    // @param filterParam Specification of matching folders for which information should be returned.
    // If null, the behavior is equivalent to a FolderFilterSpec with all properties null which means all folders match the filter.
    // @return Commonly used information about the folders matching the FolderFilterSpec.
    // @throws InvalidArgument if the FolderFilterSpec#type property contains a value that is not supported by the server.
    // @throws UnableToAllocateResource if more than 1000 folders match the FolderFilterSpec.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
	List(filterParam *FolderFilterSpec) ([]FolderSummary, error)
}
