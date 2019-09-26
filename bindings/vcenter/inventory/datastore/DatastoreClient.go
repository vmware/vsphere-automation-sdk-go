/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Datastore
 * Used by client-side stubs.
 */

package datastore

import (
)

// The ``Datastore`` interface provides methods to retrieve information about datastores.
type DatastoreClient interface {


    // Returns datastore information for the specified datastores. The key in the return value map is the datastore identifier and the value in the map is the datastore information.
    //
    // @param datastoresParam  Identifiers of the datastores for which information will be returned.
    // The parameter must contain identifiers for the resource type: ``Datastore``.
    // @return Datastore information for the specified datastores. The key in the return value map is the datastore identifier and the value in the map is the datastore information.
    // The key in the return value map will be an identifier for the resource type: ``Datastore``.
    // @throws NotFound  if no datastore can be found for one or more of the datastore identifiers in ``datastores``
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
    // * The resource ``Datastore`` referenced by the parameter ``datastores`` requires ``System.Read``.
    Find(datastoresParam []string) (map[string]*Info, error) 

}
