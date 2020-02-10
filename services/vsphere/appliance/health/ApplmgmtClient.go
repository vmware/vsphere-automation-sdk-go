/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Applmgmt
 * Used by client-side stubs.
 */

package health


// ``Applmgmt`` interface provides methods Get health status of applmgmt services.
type ApplmgmtClient interface {

    // Get health status of applmgmt services.
    // @return health status
    // @throws Error Generic error
	Get() (string, error)
}
