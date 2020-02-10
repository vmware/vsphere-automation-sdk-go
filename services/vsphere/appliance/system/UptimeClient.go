/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Uptime
 * Used by client-side stubs.
 */

package system


// ``Uptime`` interface provides methods Get the system uptime.
type UptimeClient interface {

    // Get the system uptime.
    // @return system uptime
    // @throws Error Generic error
	Get() (float64, error)
}
