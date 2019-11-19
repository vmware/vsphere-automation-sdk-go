/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Component
 * Used by client-side stubs.
 */

package routing


// Operations to retrieve information about the routing information in a vAPI component. A Component is said to contain routing information if any of its packages, services or methods contain routing information
type ComponentClient interface {

    // List all the vAPI components that contain operations which have routing information.
    // @return list of fully qualified component names
    // The return value will contain identifiers for the resource type: ``com.vmware.vapi.component``.
	List() ([]string, error)

    // Get the routing information for a vAPI component
    //
    // @param componentIdParam  fully qualified component name
    // The parameter must be an identifier for the resource type: ``com.vmware.vapi.component``.
    // @return routing information for the vAPI component
    // @throws NotFound  If the component name does not exist
	Get(componentIdParam string) (ComponentData, error)

    // Checksum of all routing metadata for a vAPI component on the server
    //
    // @param componentIdParam  fully qualified component name
    // The parameter must be an identifier for the resource type: ``com.vmware.vapi.component``.
    // @return checksum of routing metadata for a vAPI component
    // @throws NotFound  If the component name does not exist
	Fingerprint(componentIdParam string) (string, error)
}
