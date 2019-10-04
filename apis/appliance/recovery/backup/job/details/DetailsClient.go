/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Details
 * Used by client-side stubs.
 */

package details

import (
)

// The ``Details`` interface provides methods to get the details about backup jobs.
type DetailsClient interface {


    // Returns detailed information about the current and historical backup jobs.
    //
    // @param filterParam Specification of matching backup jobs for which information should be returned.
    // If null, the behavior is equivalent to FilterSpec with all properties null which means all the backup jobs match the filter.
    // @return Map of backup job identifier to Info Structure.
    // The key in the return value map will be an identifier for the resource type: ``com.vmware.appliance.recovery.backup.job``.
    // @throws Error if any error occurs during the execution of the operation.
    List(filterParam *FilterSpec) (map[string]Info, error) 

}
