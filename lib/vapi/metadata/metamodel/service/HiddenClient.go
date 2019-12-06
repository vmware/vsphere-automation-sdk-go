/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Hidden
 * Used by client-side stubs.
 */

package service


// The ``Hidden`` interface provides methods to retrieve the list of interfaces that are hidden and should not be exposed in various presentation layers of API infrastructure.
type HiddenClient interface {

    // Returns the interface identifiers that should be hidden.
    // @return The list of interface identifiers that should be hidden.
    // The return value will contain identifiers for the resource type: ``com.vmware.vapi.service``.
	List() ([]string, error)
}
