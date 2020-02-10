/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Counters
 * Used by client-side stubs.
 */

package vstats


// The ``Counters`` interface provides methods to perform various Counter related operations. Counter is derived from metric. It applies the metric to a particular class of a resource. **Warning:** This interface is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type CountersClient interface {

    // Returns information about all counters matching the filter parameters. **Warning:** This method is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    //
    // @param filterParam Filters the returned records.
    // When null no filtering will be applied.
    // @return List of Counters.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if any of the specified parameters are invalid.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have sufficient privileges.
	List(filterParam *CountersFilterSpec) ([]CountersInfo, error)

    // Returns information about a specific Counter. **Warning:** This method is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    //
    // @param cidParam Counter ID.
    // The parameter must be an identifier for the resource type: ``com.vmware.vstats.model.Counter``.
    // @return Information about the requested counter.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if ``cid`` is invalid.
    // @throws NotFound if Counter could not be located.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have sufficient privileges.
	Get(cidParam string) (CountersInfo, error)
}
