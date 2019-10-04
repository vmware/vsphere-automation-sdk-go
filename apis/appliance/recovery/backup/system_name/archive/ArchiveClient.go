/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Archive
 * Used by client-side stubs.
 */

package archive

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/appliance/recovery/backup"
)

// The ``Archive`` interface provides methods to get the backup information.
type ArchiveClient interface {


    // Returns the information for backup corresponding to given backup location and system name.
    //
    // @param specParam LocationSpec Structure.
    // @param systemNameParam System name identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.recovery.backup.system_name``.
    // @param archiveParam Archive identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.recovery.backup.system_name.archive``.
    // @return Info Structure.
    // @throws NotFound if backup does not exist.
    // @throws Error if any error occurs during the execution of the operation.
    Get(specParam backup.LocationSpec, systemNameParam string, archiveParam string) (Info, error) 


    // Returns information about backup archives corresponding to given backup location and system name, which match the FilterSpec.
    //
    // @param locSpecParam LocationSpec Structure.
    // @param systemNameParam System name identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.recovery.backup.system_name``.
    // @param filterSpecParam Specification of matching backups for which information should be returned.
    // @return Commonly used information about the backup archives.
    // @throws NotFound if combination of ``loc_spec`` and system name does not refer to an existing location on the backup server.
    // @throws Error if any error occurs during the execution of the operation.
    List(locSpecParam backup.LocationSpec, systemNameParam string, filterSpecParam FilterSpec) ([]Summary, error) 

}
