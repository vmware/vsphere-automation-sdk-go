/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: System
 * Used by client-side stubs.
 */

package health

import (
	"time"
)

// ``System`` interface provides methods Get overall health of the system.
type SystemClient interface {

    // Get last check timestamp of the health of the system.
    // @return System health last check timestamp
    // @throws Error Generic error
	Lastcheck() (time.Time, error)

    // Get overall health of system.
    // @return System health
    // @throws Error Generic error
	Get() (SystemHealthLevel, error)
}
