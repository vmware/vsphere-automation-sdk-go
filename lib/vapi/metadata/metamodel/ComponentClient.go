/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Component
 * Used by client-side stubs.
 */

package metamodel


// The ``Component`` interface providers methods to retrieve metamodel information of a component element. 
//
//  A component defines a set of functionality that is deployed together and versioned together. For example, all the interfaces that belong to VMware Content Library are part of a single component. A component element describes a component. A component element contains one or more package elements. 
//
//  The methods for package elements are provided by interface Package.
type ComponentClient interface {

    // Returns the identifiers for the component elements that are registered with the infrastructure.
    // @return The list of identifiers for the component elements that are registered with the infrastructure.
    // The return value will contain identifiers for the resource type: ``com.vmware.vapi.component``.
	List() ([]string, error)

    // Retrieves metamodel information about the component element corresponding to ``component_id``. 
    //
    //  The ComponentData contains the metamodel information about the component and it's fingerprint. It contains information about all the package elements that are contained in this component element.
    //
    // @param componentIdParam Identifier of the component element.
    // The parameter must be an identifier for the resource type: ``com.vmware.vapi.component``.
    // @return The ComponentData instance that corresponds to ``component_id``.
    // @throws NotFound if the component element associated with ``component_id`` is not registered with the infrastructure.
	Get(componentIdParam string) (ComponentData, error)

    // Retrieves the fingerprint computed from the metamodel metadata of the component element corresponding to ``component_id``. 
    //
    //  The fingerprint provides clients an efficient way to check if the metadata for a particular component element has been modified on the server. The client can do this by comparing the result of this operation with the fingerprint returned in the result of Component#get.
    //
    // @param componentIdParam Identifier of the component element.
    // The parameter must be an identifier for the resource type: ``com.vmware.vapi.component``.
    // @return The fingerprint computed from the metamodel metadata of the component element.
    // @throws NotFound if the component element associated with ``component_id`` is not registered with the infrastructure.
	Fingerprint(componentIdParam string) (string, error)
}
