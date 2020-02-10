/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Parts
 * Used by client-side stubs.
 */

package backup


// ``Parts`` interface provides methods Provides list of parts optional for the backup
type PartsClient interface {

    // Gets a list of the backup parts.
    // @return Information about each of the backup parts.
    // @throws Error if any error occurs during the execution of the operation.
	List() ([]PartsPart, error)

    // Gets the size (in MB) of the part.
    //
    // @param idParam Identifier of the part.
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.recovery.backup.parts``.
    // @return long Size of the part in megabytes.
    // @throws Error if any error occurs during the execution of the operation.
	Get(idParam string) (int64, error)
}
