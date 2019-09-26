/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Compliance
 * Used by client-side stubs.
 */

package compliance

import (
)

// The Compliance interface provides methods related to all the associated entities of given compliance statuses.
type ComplianceClient interface {


    // Returns compliance information about entities matching the filter FilterSpec. Entities without storage policy association are not returned.
    //
    // @param filterParam compliance status of matching entities for which information should be returned.
    // @return compliance information about entities matching the filter FilterSpec.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if the FilterSpec#status property contains a value that is not supported by the server.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    List(filterParam FilterSpec) ([]Summary, error) 

}
