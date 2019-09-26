/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Folder
 * Used by client-side stubs.
 */

package folder

import (
)

// The Folder interface provides methods for manipulating a vCenter Server folder.
type FolderClient interface {


    // Returns information about at most 1000 visible (subject to permission checks) folders in vCenter matching the FilterSpec.
    //
    // @param filterParam Specification of matching folders for which information should be returned.
    // If null, the behavior is equivalent to a FilterSpec with all properties null which means all folders match the filter.
    // @return Commonly used information about the folders matching the FilterSpec.
    // @throws InvalidArgument if the FilterSpec#type property contains a value that is not supported by the server.
    // @throws UnableToAllocateResource if more than 1000 folders match the FilterSpec.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    List(filterParam *FilterSpec) ([]Summary, error) 

}
