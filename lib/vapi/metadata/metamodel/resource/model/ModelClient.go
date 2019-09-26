/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Model
 * Used by client-side stubs.
 */

package model

import (
)

// The ``Model`` interface provides methods to retrieve information about models. 
//
//  A structure is used as a model if it is used for persisting data about an entity. Some of the fields in the model structure are also used for creating indexes for querying. 
//
//  One or more services can operate on the same resource type. One or more services can provide the model structure for an entity of this resource type. Using ``Model`` interface you can retrieve the list of all the structure elements that are model structures for a given resource type.
type ModelClient interface {


    // Returns the set of identifiers for the structure elements that are models for the resource type corresponding to ``resource_id``. 
    //
    //  The Structure interface provides methods to retrieve more details about the structure elements corresponding to the identifiers returned by this method.
    //
    // @param resourceIdParam Identifier of the resource type.
    // The parameter must be an identifier for the resource type: ``com.vmware.vapi.resource``.
    // @return The set of identifiers for the models that are associated with the resource type in ``resource_id``.
    // The return value will contain identifiers for the resource type: ``com.vmware.vapi.structure``.
    // @throws NotFound if the resource type associated with ``resource_id`` does not exist.
    List(resourceIdParam string) (map[string]bool, error) 

}
