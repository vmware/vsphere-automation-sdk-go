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
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vcenter/lcm/tasks"
)

type ScheduledClient interface {


    // Retrieves information of the tasks that are scheduled to be executed in the future. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @return The information of the tasks that are scheduled to be executed in the future.
    List() ([]tasks.Info, error) 

}
