/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Applmgmt
 * Used by client-side stubs.
 */

package applmgmt

import (
)

// ``Applmgmt`` interface provides methods Get health status of applmgmt services.
type ApplmgmtClient interface {


    // Get health status of applmgmt services.
    // @return health status
    // @throws Error Generic error
    Get() (string, error) 

}
