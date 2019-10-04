/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Parts
 * Used by client-side stubs.
 */

package parts

import (
)

// ``Parts`` interface provides methods Provides list of parts optional for the backup
type PartsClient interface {


    // Gets a list of the backup parts.
    // @return Information about each of the backup parts.
    // @throws Error if any error occurs during the execution of the operation.
    List() ([]Part, error) 


    // Gets the size (in MB) of the part.
    //
    // @param idParam Identifier of the part.
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.recovery.backup.parts``.
    // @return long Size of the part in megabytes.
    // @throws Error if any error occurs during the execution of the operation.
    Get(idParam string) (int64, error) 

}
