/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: DefaultPolicy
 * Used by client-side stubs.
 */

package defaultPolicy

import (
)

// The ``DefaultPolicy`` interface provides methods related to storage policies associated with datastore object. The DefaultPolicy#get method provides information about the default storage policy associated with the specific datastore.
type DefaultPolicyClient interface {


    // Returns the identifier of the current default storage policy associated with the specified datastore.
    //
    // @param datastoreParam Identifier of the datastore for which the default policy is requested.
    // The parameter must be an identifier for the resource type: ``Datastore``.
    // @return Identifier of the default storage policy associated with the specified datastore.
    // The return value will be an identifier for the resource type: ``com.vmware.vcenter.StoragePolicy``.
    // this field is null if there is no default storage policy associated with the datastore.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the specified datastore does not exist.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user is not authenticated.
    // @throws Unauthorized if the user does not have the required priveleges.
    Get(datastoreParam string) (*string, error) 

}
