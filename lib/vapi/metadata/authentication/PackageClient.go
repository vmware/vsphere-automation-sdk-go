/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Package
 * Used by client-side stubs.
 */

package authentication


// The ``Package`` interface provides methods to retrieve authentication information of a package element. 
//
//  A package element is said to contain authentication information if there is a default authentication assigned to all service elements contained in the package element or if one of the service element contained in this package element has authentication information.
type PackageClient interface {

    // Returns the identifiers for the package elements that have authentication information.
    // @return The list of identifiers for the package elements that have authentication information.
    // The return value will contain identifiers for the resource type: ``com.vmware.vapi.package``.
	List() ([]string, error)

    // Retrieves authentication information about the package element corresponding to ``package_id``.
    //
    // @param packageIdParam Identifier of the package element.
    // The parameter must be an identifier for the resource type: ``com.vmware.vapi.package``.
    // @return The PackageInfo instance that corresponds to ``package_id``
    // @throws NotFound if the package element associated with ``package_id`` does not have any authentication information.
	Get(packageIdParam string) (PackageInfo, error)
}
