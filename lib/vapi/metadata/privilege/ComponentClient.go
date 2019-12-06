/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Component
 * Used by client-side stubs.
 */

package privilege


// The ``Component`` interface provides methods to retrieve privilege information of a component element. 
//
//  A component element is said to contain privilege information if any one of package elements in it contains privilege information.
type ComponentClient interface {

    // Returns the identifiers for the component elements that have privilege information.
    // @return The list of identifiers for the component elements that have privilege information.
    // The return value will contain identifiers for the resource type: ``com.vmware.vapi.component``.
	List() ([]string, error)

    // Retrieves privilege information about the component element corresponding to ``component_id``. 
    //
    //  The ComponentData contains the privilege information about the component element and its fingerprint. It contains information about all the package elements that belong to this component element.
    //
    // @param componentIdParam Identifier of the component element.
    // The parameter must be an identifier for the resource type: ``com.vmware.vapi.component``.
    // @return The ComponentData instance that corresponds to ``component_id``
    // @throws NotFound if the component element associated with ``component_id`` does not have any privilege information.
	Get(componentIdParam string) (ComponentData, error)

    // Retrieves the fingerprint computed from the privilege metadata of the component element corresponding to ``component_id``. 
    //
    //  The fingerprint provides clients an efficient way to check if the metadata for a particular component has been modified on the server. The client can do this by comparing the result of this operation with the fingerprint returned in the result of Component#get.
    //
    // @param componentIdParam Identifier of the component element.
    // The parameter must be an identifier for the resource type: ``com.vmware.vapi.component``.
    // @return The fingerprint computed from the privilege metadata of the component.
    // @throws NotFound if the component element associated with ``component_id`` does not have any privilege information.
	Fingerprint(componentIdParam string) (string, error)
}
