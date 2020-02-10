/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Consolecli
 * Used by client-side stubs.
 */

package access


// ``Consolecli`` interface provides methods Get/Set enabled state of CLI.
type ConsolecliClient interface {

    // Set enabled state of the console-based controlled CLI (TTY1).
    //
    // @param enabledParam Console-based controlled CLI is enabled.
    // @throws Error Generic error
	Set(enabledParam bool) error

    // Get enabled state of the console-based controlled CLI (TTY1).
    // @return Console-based controlled CLI is enabled.
    // @throws Error Generic error
	Get() (bool, error)
}
