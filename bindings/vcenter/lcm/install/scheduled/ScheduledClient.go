/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Scheduled
 * Used by client-side stubs.
 */

package scheduled

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/vcenter/lcm"
)

type ScheduledClient interface {


    // Retrieves information of the install operation that has the given id. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param taskParam  The identifier of the task.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.lcm.task``.
    // @return The task information of the given identifier.
    // @throws NotFound  If the id doesn't exist.
    Get(taskParam string) (lcm.InstallSpec, error) 


    // Updates the information of the task that has the given identifier. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param taskParam  The identifier of the task.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.lcm.task``.
    // @param specParam  The configuration that contains the time and deployment information to be deployed at the given start time.
    // @throws NotFound  If the identifier doesn't exist.
    Set(taskParam string, specParam lcm.InstallSpec) error 

}
