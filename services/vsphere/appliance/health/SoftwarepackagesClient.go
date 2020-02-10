/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Softwarepackages
 * Used by client-side stubs.
 */

package health


// ``Softwarepackages`` interface provides methods Get information on available software updates available in remote VUM repository.
type SoftwarepackagesClient interface {

    // Get information on available software updates available in the remote vSphere Update Manager repository. Red indicates that security updates are available. Orange indicates that non-security updates are available. Green indicates that there are no updates available. Gray indicates that there was an error retreiving information on software updates.
    // @return software updates available.
    // @throws Error Generic error
	Get() (SoftwarepackagesHealthLevel, error)
}
