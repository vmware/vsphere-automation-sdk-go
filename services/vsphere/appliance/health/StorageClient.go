/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Storage
 * Used by client-side stubs.
 */

package health


// ``Storage`` interface provides methods Get storage health.
type StorageClient interface {

    // Get storage health.
    // @return Storage health.
    // @throws Error Generic error
	Get() (StorageHealthLevel, error)
}
