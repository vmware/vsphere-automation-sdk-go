/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Load
 * Used by client-side stubs.
 */

package health


// ``Load`` interface provides methods Get load health.
type LoadClient interface {

    // Get load health.
    // @return Load health.
    // @throws Error Generic error
	Get() (LoadHealthLevel, error)
}
