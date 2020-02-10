/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: CounterSets
 * Used by client-side stubs.
 */

package vstats


// The ``CounterSets`` interface provides methods for accessing groupings of counters. Counter-sets allow consumers to use groups of counters. The counters may relate to different resource types. When an AcqSpecs record refers to a counter-set, only the relevant counters apply. **Warning:** This interface is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type CounterSetsClient interface {

    // Returns information about all the counter sets. **Warning:** This method is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    // @return List of counter sets.
    // @throws Error if the system reports an error while responding to the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have sufficient privileges.
	List() ([]CounterSetsInfo, error)

    // Returns information about a specific counter set. **Warning:** This method is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    //
    // @param counterSetParam identifier of the counter set to retrieve.
    // The parameter must be an identifier for the resource type: ``com.vmware.vstats.model.CounterSet``.
    // @return Information about the desired CounterSet.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if ``counter_set`` is invalid.
    // @throws NotFound if the requested counter set could not be located.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have sufficient privileges.
	Get(counterSetParam string) (CounterSetsInfo, error)
}
