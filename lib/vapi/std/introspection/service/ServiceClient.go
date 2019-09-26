/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Service
 * Used by client-side stubs.
 */

package service

import (
)

// The Service service provides operations to retrieve information about the services exposed by a vAPI provider. A provider is a container that exposes one or more vAPI services.
type ServiceClient interface {


    // Returns the set of service identifiers.
    // @return set of service identifiers
    // The return value will contain identifiers for the resource type: ``com.vmware.vapi.service``.
    List() (map[string]bool, error) 


    // Returns the Info for the specified service
    //
    // @param idParam service identifier
    // The parameter must be an identifier for the resource type: ``com.vmware.vapi.service``.
    // @return Info for the specified service
    // @throws NotFound If the service identifier does not exist
    Get(idParam string) (Info, error) 

}
