/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Tasks
 * Used by client-side stubs.
 */

package tasks

import (
)

// The service of life cycle management tasks. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type TasksClient interface {


    // Removes the task of the given identifier. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param taskParam 
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.lcm.task``.
    // @throws NotFound  If the identifier doesn't exist
    Delete(taskParam string) error 

}
