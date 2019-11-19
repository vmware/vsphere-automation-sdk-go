/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Structure
 * Used by client-side stubs.
 */

package metamodel


// The ``Structure`` interface providers methods to retrieve metamodel information about a structure element in the interface definition language.
type StructureClient interface {

    // Returns the identifiers for the structure elements that are contained in all the package elements and service elements.
    // @return The list of identifiers for the structure elements.
    // The return value will contain identifiers for the resource type: ``com.vmware.vapi.structure``.
	List() ([]string, error)

    // Retrieves information about the structure element corresponding to ``structure_id``. 
    //
    //  The StructureInfo contains the metamodel information about the structure element. It contains information about all the field elements and enumeration elements contained in this structure element.
    //
    // @param structureIdParam Identifier of the structure element.
    // The parameter must be an identifier for the resource type: ``com.vmware.vapi.structure``.
    // @return The StructureInfo instance that corresponds to ``structure_id``.
    // @throws NotFound if the structure element associated with ``structure_id`` is not contained in any of the package elements or service elements.
	Get(structureIdParam string) (StructureInfo, error)
}
