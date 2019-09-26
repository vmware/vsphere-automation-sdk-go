/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Uuid
 * Used by client-side stubs.
 */

package uuid

import (
)

// The ``Uuid`` interface provides methods to get the UUID for the appliance. This api will be primarily used for Applmgmt UI Telemetry. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type UuidClient interface {


    // Get the vCSA UUID (Unique IDentifier for the appliance). **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @return UUID for the appliance
    // @throws Error Generic error
    Get() (string, error) 

}
