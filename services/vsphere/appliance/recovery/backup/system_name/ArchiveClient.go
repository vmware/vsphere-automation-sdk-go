/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Archive
 * Used by client-side stubs.
 */

package system_name

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/vsphere/appliance/recovery/backup"
)

// The ``Archive`` interface provides methods to get the backup information. This interface was added in vSphere API 6.7.
type ArchiveClient interface {

    // Returns the information for backup corresponding to given backup location and system name. This method was added in vSphere API 6.7.
    //
    // @param specParam LocationSpec Structure.
    // @param systemNameParam System name identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.recovery.backup.system_name``.
    // @param archiveParam Archive identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.recovery.backup.system_name.archive``.
    // @return Info Structure.
    // @throws NotFound if backup does not exist.
    // @throws Error if any error occurs during the execution of the operation.
	Get(specParam backup.LocationSpec, systemNameParam string, archiveParam string) (ArchiveInfo, error)

    // Returns information about backup archives corresponding to given backup location and system name, which match the ArchiveFilterSpec. This method was added in vSphere API 6.7.
    //
    // @param locSpecParam LocationSpec Structure.
    // @param systemNameParam System name identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.recovery.backup.system_name``.
    // @param filterSpecParam Specification of matching backups for which information should be returned.
    // @return Commonly used information about the backup archives.
    // @throws NotFound if combination of ``loc_spec`` and system name does not refer to an existing location on the backup server.
    // @throws Error if any error occurs during the execution of the operation.
	List(locSpecParam backup.LocationSpec, systemNameParam string, filterSpecParam ArchiveFilterSpec) ([]ArchiveSummary, error)
}
