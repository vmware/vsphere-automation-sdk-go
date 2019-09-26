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

// ``Storage`` interface provides methods Appliance storage configuration
type StorageClient interface {


    // Get disk to partition mapping.
    // @return list of mapping items
    // @throws Error Generic error
    List() ([]StorageMapping, error) 


    // Resize all partitions to 100 percent of disk size.
    // @throws Error Generic error
    Resize() error 


    // Resize all partitions to 100 percent of disk size.
    // @return List of the partitions with the size before and after resizing
    // @throws Error Generic error
    ResizeEx() (map[string]StorageChange, error) 

}
