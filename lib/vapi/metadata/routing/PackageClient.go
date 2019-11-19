/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Package
 * Used by client-side stubs.
 */

package routing


// Operations to retrieve information about routing information in a vAPI package A Package is said to contain routing information if there is a default RoutingInfo assigned to all operations within a package or if one of the operations within this package has explicit routing information
type PackageClient interface {

    // List of all vAPI packages that have routing information
    // @return list of fully qualified package names
    // The return value will contain identifiers for the resource type: ``com.vmware.vapi.package``.
	List() ([]string, error)

    // Get the routing information for a vAPI package
    //
    // @param packageIdParam fully qualified package name
    // The parameter must be an identifier for the resource type: ``com.vmware.vapi.package``.
    // @return routing information for the vAPI package
    // @throws NotFound If the package name does not exist
	Get(packageIdParam string) (PackageInfo, error)
}
