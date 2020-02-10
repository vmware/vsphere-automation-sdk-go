/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: ResourceAddresses
 * Used by client-side stubs.
 */

package vstats


// The ``ResourceAddresses`` interface provides methods to perform resource addressing queries. **Warning:** This interface is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type ResourceAddressesClient interface {

    // Returns the list of Resource Addresses matching the filter parameters. **Warning:** This method is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    //
    // @param filterParam Criteria for selecting records to return.
    // If null all records will be returned.
    // @return Matching resource addresses.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if any of the specified parameters are invalid.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have sufficient privileges.
	List(filterParam *ResourceAddressesFilterSpec) (ResourceAddressesListResult, error)

    // Returns information about a specific Resource Address. **Warning:** This method is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    //
    // @param idParam Resource Address ID.
    // The parameter must be an identifier for the resource type: ``com.vmware.vstats.model.RsrcAddr``.
    // @return Information about the desired Resource Address.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if ``id`` is invalid.
    // @throws NotFound if Resource Address could not be located.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have sufficient privileges.
	Get(idParam string) (ResourceAddressesInfo, error)
}
