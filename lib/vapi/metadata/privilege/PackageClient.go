/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Package
 * Used by client-side stubs.
 */

package privilege


// The ``Package`` interface provides methods to retrieve privilege information of a package element. 
//
//  A package element is said to contain privilege information if there is a default privilege assigned to all service elements contained in the package element or if one of the operation elements contained in one of the service elements in this package element has privilege information.
type PackageClient interface {

    // Returns the identifiers for the package elements that have privilege information.
    // @return The list of identifiers for the package elements that have privilege information.
    // The return value will contain identifiers for the resource type: ``com.vmware.vapi.package``.
	List() ([]string, error)

    // Retrieves privilege information about the package element corresponding to ``package_id``.
    //
    // @param packageIdParam Identifier of the package element.
    // The parameter must be an identifier for the resource type: ``com.vmware.vapi.package``.
    // @return The PackageInfo instance that corresponds to ``package_id``
    // @throws NotFound if the package element associated with ``package_id`` does not have any privilege information.
	Get(packageIdParam string) (PackageInfo, error)
}
