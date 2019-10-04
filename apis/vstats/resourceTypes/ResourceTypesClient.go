/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: ResourceTypes
 * Used by client-side stubs.
 */

package resourceTypes

import (
)

// The ``ResourceTypes`` interface provides list of resource types. Resource refers to any item or concept that could have measurable properties. It is a prerequisite that a resource can be identified. 
//
//  Example resource types are virtual machine, virtual disk etc.
type ResourceTypesClient interface {


    // Returns a list of available resource types.
    // @return List of resource types.
    // @throws Error if the system reports an error while responding to the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have sufficient privileges.
    List() ([]Summary, error) 

}
