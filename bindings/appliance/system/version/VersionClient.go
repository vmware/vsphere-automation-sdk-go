/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Version
 * Used by client-side stubs.
 */

package version

import (
)

// ``Version`` interface provides methods Get the appliance version.
type VersionClient interface {


    // Get the version.
    // @return version information about the appliance
    // @throws Error Generic error
    Get() (VersionStruct, error) 

}
