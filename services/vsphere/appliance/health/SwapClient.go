/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Swap
 * Used by client-side stubs.
 */

package health


// ``Swap`` interface provides methods Get swap health.
type SwapClient interface {

    // Get swap health.
    // @return Swap health
    // @throws Error Generic error
	Get() (SwapHealthLevel, error)
}
