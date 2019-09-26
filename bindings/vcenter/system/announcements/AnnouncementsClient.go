/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Announcements
 * Used by client-side stubs.
 */

package announcements

import (
)

// The ``Announcements`` interface provides methods to manage system wide announcements. **Warning:** This interface is available as technical preview. It may be changed in a future release.
type AnnouncementsClient interface {


    // Return information about all the announcements in the system. **Warning:** This method is available as technical preview. It may be changed in a future release.
    // @return Map from announcement identifier to announcement Info structures
    // The key in the return value map will be an identifier for the resource type: ``com.vmware.vcenter.system.announcements``.
    // @throws Error Generic error
    List() (map[string]Info, error) 


    // Return information about a specific announcements in the system. **Warning:** This method is available as technical preview. It may be changed in a future release.
    //
    // @param announcementParam Identifier of the announcement.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.system.announcements``.
    // @return Map from announcement identifier to announcement Info structures
    // @throws Error Generic error
    Get(announcementParam string) (Info, error) 


    // Create a new announcement. **Warning:** This method is available as technical preview. It may be changed in a future release.
    //
    // @param specParam Spec structure defining the announcement properties
    // @return Unique identifier of the new announcement
    // The return value will be an identifier for the resource type: ``com.vmware.vcenter.system.announcements``.
    // @throws Error Generic error
    Create(specParam Spec) (string, error) 


    // Set announcement properties. **Warning:** This method is available as technical preview. It may be changed in a future release.
    //
    // @param announcementParam Identifier of the announcement. If such identifier exists then the object will be updated with the new values. If such identifier doesn't exist a new object with this identifier will be created.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.system.announcements``.
    // @param specParam Spec structure defining the announcement properties
    // @throws Error Generic error
    Set(announcementParam string, specParam Spec) error 


    // Delete an announcement. **Warning:** This method is available as technical preview. It may be changed in a future release.
    //
    // @param announcementParam Identifier of the announcement
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.system.announcements``.
    // @throws NotFound If the identifier is not found
    // @throws Error Generic error
    Delete(announcementParam string) error 

}
