/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Storage
 * Used by client-side stubs.
 */

package storage

import (
)

// ``Storage`` interface provides methods Get storage health.
type StorageClient interface {


    // Get storage health.
    // @return Storage health.
    // @throws Error Generic error
    Get() (HealthLevel, error) 

}
