/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: CounterSets
 * Used by client-side stubs.
 */

package counterSets

import (
)

// The ``CounterSets`` interface provides methods for accessing groupings of counters. Counter-sets allow consumers to use groups of counters. The counters may relate to different resource types. When an AcqSpecs record refers to a counter-set, only the relevant counters apply.
type CounterSetsClient interface {


    // Returns information about all the counter sets.
    // @return List of counter sets.
    // @throws Error if the system reports an error while responding to the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have sufficient privileges.
    List() ([]Info, error) 


    // Returns information about a specific counter set.
    //
    // @param counterSetParam identifier of the counter set to retrieve.
    // The parameter must be an identifier for the resource type: ``com.vmware.vstats.model.CounterSet``.
    // @return Information about the desired CounterSet.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if ``counter_set`` is invalid.
    // @throws NotFound if the requested counter set could not be located.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have sufficient privileges.
    Get(counterSetParam string) (Info, error) 

}
