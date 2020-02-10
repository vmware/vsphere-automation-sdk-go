/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Ssh
 * Used by client-side stubs.
 */

package access


// ``Ssh`` interface provides methods Get/Set enabled state of SSH-based controlled CLI.
type SshClient interface {

    // Set enabled state of the SSH-based controlled CLI.
    //
    // @param enabledParam SSH-based controlled CLI is enabled.
    // @throws Error Generic error
	Set(enabledParam bool) error

    // Get enabled state of the SSH-based controlled CLI.
    // @return SSH-based controlled CLI is enabled.
    // @throws Error Generic error
	Get() (bool, error)
}
