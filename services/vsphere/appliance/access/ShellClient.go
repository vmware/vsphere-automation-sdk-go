/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Shell
 * Used by client-side stubs.
 */

package access


// ``Shell`` interface provides methods Get/Set enabled state of BASH.
type ShellClient interface {

    // Set enabled state of BASH, that is, access to BASH from within the controlled CLI.
    //
    // @param configParam Shell configuration
    // @throws Error Generic error
	Set(configParam ShellShellConfig) error

    // Get enabled state of BASH, that is, access to BASH from within the controlled CLI.
    // @return Current shell configuration.
    // @throws Error Generic error
	Get() (ShellShellConfig, error)
}
