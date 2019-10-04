/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: SystemName
 * Used by client-side stubs.
 */

package systemName

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/appliance/recovery/backup"
)

// The ``SystemName`` interface provides methods to enumerate system names of appliance backups.
type SystemNameClient interface {


    // Returns a list of system names for which backup archives exist under ``loc_spec``.
    //
    // @param locSpecParam LocationSpec Structure
    // @return list of system names
    // The return value will contain identifiers for the resource type: ``com.vmware.appliance.recovery.backup.system_name``.
    // @throws NotFound if ``loc_spec`` doesn't refer to an existing location on the backup server.
    // @throws Error if any error occurs during the execution of the operation.
    List(locSpecParam backup.LocationSpec) ([]string, error) 

}
