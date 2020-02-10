/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Databasestorage
 * Used by client-side stubs.
 */

package health


// ``Databasestorage`` interface provides methods Get database storage health.
type DatabasestorageClient interface {

    // Get database storage health.
    // @return Database storage health
    // @throws Error Generic error
	Get() (DatabasestorageHealthLevel, error)
}
