/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Package
 * Used by client-side stubs.
 */

package package_

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/metadata/metamodel"
)

// The ``Package`` interface provides methods to retrieve metamodel information about a package element in the interface definition language. 
//
//  A package is a logical grouping of services, structures and enumerations. A package element describes the package. It contains the service elements, structure elements and enumeration elements that are grouped together.
type PackageClient interface {


    // Returns the identifiers for the packages elements that are contained in all the registered component elements.
    // @return The list of identifiers for the package elements that are contained in all the registered component elements.
    // The return value will contain identifiers for the resource type: ``com.vmware.vapi.package``.
    List() ([]string, error) 


    // Retrieves information about the package element corresponding to ``package_id``.
    //
    // @param packageIdParam Identifier of the package element.
    // The parameter must be an identifier for the resource type: ``com.vmware.vapi.package``.
    // @return The metamodel.PackageInfo instance that corresponds to ``package_id``.
    // @throws NotFound if the package element associated with ``package_id`` does not exist.
    Get(packageIdParam string) (metamodel.PackageInfo, error) 

}
