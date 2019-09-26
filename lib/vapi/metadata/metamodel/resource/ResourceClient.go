/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Resource
 * Used by client-side stubs.
 */

package resource

import (
)

// The Resource interface provides methods to retrieve information about resource types. 
//
//  A service is a logical grouping of operations that operate on an entity. Each entity is identifier by a namespace (or resource type) and an unique identifier.
type ResourceClient interface {


    // Returns the set of resource types present across all the service elements contained in all the package elements.
    // @return Set of resource types
    // The return value will contain identifiers for the resource type: ``com.vmware.vapi.resource``.
    List() (map[string]bool, error) 

}
