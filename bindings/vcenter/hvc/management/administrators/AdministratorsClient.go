/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Administrators
 * Used by client-side stubs.
 */

package administrators

import (
)

// The ``Administrators`` provides methods to update, delete, and list groups in the local sso group. This is limited to the Hybrid Linked Mode service. **Warning:** This interface is available as technical preview. It may be changed in a future release.
type AdministratorsClient interface {


    // Add the local sso group with the new group. **Warning:** This method is available as technical preview. It may be changed in a future release.
    //
    // @param groupNameParam Name of the new group to be added. Ex - xyz\\\\@abc.com where xyz is the group name and abc.com is the domain name
    // @throws Unauthorized If the user is not authorized.
    // @throws Error if the system reports an error while responding to the request.
    Add(groupNameParam string) error 


    // Remove the group from the local sso group. **Warning:** This method is available as technical preview. It may be changed in a future release.
    //
    // @param groupNameParam Name of the group to be removed. Ex - xyz\\\\@abc.com where xyz is the group name and abc.com is the domain name
    // @throws Unauthorized If the user is not authorized.
    // @throws Error if the system reports an error while responding to the request.
    Remove(groupNameParam string) error 


    // Sets the groups in the local sso group. **Warning:** This method is available as technical preview. It may be changed in a future release.
    //
    // @param groupNamesParam Names the groups to be in the CloudAdminGroup Ex - xyz\\\\@abc.com where xyz is the group name and abc.com is the domain name
    // @throws Unauthorized If the user is not authorized.
    // @throws Error if the system reports an error while responding to the request.
    Set(groupNamesParam map[string]bool) error 


    // Enumerates the set of all the groups in the local sso group. **Warning:** This method is available as technical preview. It may be changed in a future release.
    // @return The map with bool value of all the groups.
    // @throws Error if the system reports an error while responding to the request.
    Get() (map[string]bool, error) 

}
