/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Events
 * Used by client-side stubs.
 */

package events

import (
)

// The ``Events`` interface provides methods to get Kubernetes events related to a particular namespace. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type EventsClient interface {


    // Returns Kubernetes events related to a specific namespace. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param namespaceParam Identifier for the Namespace.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.namespaces.Instance``.
    // @return List of latest Kubernetes events by 'lastTimestamp', with a limit of 1000 events.
    // @throws NotFound if namespace could not be located.
    // @throws Unsupported if the specified cluster does not have Namespaces enabled.
    // @throws Error if the system reports an error while responding to the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have System.Read privilege.
    Get(namespaceParam string) ([]Event, error) 

}
